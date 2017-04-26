// This program scrapes the cloudformation documentation to determine the schema
// and produces a go program to the file specified by the `-out` flag.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/mitchellh/go-wordwrap"
)

// TemplateReference describes the CloudFormation template schema
type TemplateReference struct {
	Resources []*Resource
}

func (tr *TemplateReference) Load() error {
	for _, docURI := range []string{"aws-template-resource-type-ref.html", "aws-product-property-reference.html"} {
		docReader, err := getDoc(docURI)
		if err != nil {
			return err
		}
		doc, err := goquery.NewDocumentFromReader(docReader)
		if err != nil {
			return err
		}
		doc.Find(".highlights li a").Each(func(i int, s *goquery.Selection) {
			name := s.Text()
			name = regexp.MustCompile("\\s+").ReplaceAllString(name, " ")
			name = strings.TrimSpace(name)
			href, _ := s.Attr("href")

			// hack around documentation bug, reported 20 Sept 2016
			if href == "aws-properties-ec2-networkaclentry-portrange.html" {
				name = "EC2 NetworkAclEntry PortRange"
			}
			tr.Resources = append(tr.Resources, &Resource{Name: name, Href: href})

		})
	}

	for _, resource := range tr.Resources {
		if err := resource.Load(); err != nil {
			return fmt.Errorf("%s: %s", resource.Name, err)
		}
	}
	return nil
}

func (tr *TemplateReference) WriteGo(w io.Writer) error {
	cmd := exec.Command("gofmt")
	writer, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	reader, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	go func() {
		tr.writeGo(writer)
		writer.Close()
	}()
	_, err = io.Copy(w, reader)
	if err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		return err
	}
	return err
}

const customResourceDefinitionBlock = `
// CustomResourceProvider allows extend the NewResourceByType factory method
// with their own resource types.
type CustomResourceProvider func(customResourceType string) ResourceProperties

var customResourceProviders []CustomResourceProvider

// Register a CustomResourceProvider with go-cloudformation. Multiple
// providers may be registered. The first provider that returns a non-nil
// interface will be used and there is no check for a uniquely registered
// resource type.
func RegisterCustomResourceProvider(provider CustomResourceProvider) {
	customResourceProviders = append(customResourceProviders, provider)
}
`
const customResourceDefaultLabelBlock = `
default:
  for _, eachProvider := range customResourceProviders {
    customType := eachProvider(typeName)
    if nil != customType {
      return customType
    }
  }
`

func (tr *TemplateReference) writeGo(w io.Writer) {
	fmt.Fprintf(w, "package cloudformation\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "import \"time\"\n")
	fmt.Fprintf(w, "import \"encoding/json\"\n")
	fmt.Fprint(w, customResourceDefinitionBlock)
	for _, resource := range tr.Resources {
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, "// %s represents %s\n", resource.GoName(), resource.Name)
		fmt.Fprintf(w, "//\n")
		fmt.Fprintf(w, "// see %s%s\n", rootURL, resource.Href)
		fmt.Fprintf(w, "type %s struct {\n", resource.GoName())
		for i, p := range resource.Properties {
			if i > 0 {
				fmt.Fprintf(w, "\n")
			}
			fmt.Fprintf(w, "%s", p.Comment("  // "))
			fmt.Fprintf(w, "  %s %s `json:\"%s,omitempty\"`\n", p.GoName(), p.GoType(tr), p.Name)
		}
		fmt.Fprintf(w, "}\n")
		fmt.Fprintf(w, "\n")
		if resource.IsTopLevelResource() {

			fmt.Fprintf(w, "// CfnResourceType returns %s to implement the ResourceProperties interface\n", resource.Name)
			fmt.Fprintf(w, "func (s %s) CfnResourceType() string {\n", resource.GoName())
			fmt.Fprintf(w, "	return %q\n", resource.Name)
			fmt.Fprintf(w, "}\n")
			fmt.Fprintf(w, "\n")
		}

		// Cloudformation allows a single object when a list of objects is expected. To
		// handle this we need to generate a *List type. This applies mostly only to
		// non-top-level objects. (The only exception is AWS::Route53::RecordSet which is
		// both a top level resource and a child element of AWS::Route53::RecordSetGroup
		if !resource.IsTopLevelResource() || resource.GoName() == "Route53RecordSet" {
			fmt.Fprintf(w, "// %sList represents a list of %s\n", resource.GoName(), resource.GoName())
			fmt.Fprintf(w, "type %sList []%s\n", resource.GoName(), resource.GoName())
			fmt.Fprintf(w, "\n")
			fmt.Fprintf(w, "// UnmarshalJSON sets the object from the provided JSON representation\n")
			fmt.Fprintf(w, "func (l *%sList) UnmarshalJSON(buf []byte) error {\n", resource.GoName())
			fmt.Fprintf(w, "	// Cloudformation allows a single object when a list of objects is expected\n")
			fmt.Fprintf(w, "	item := %s{}\n", resource.GoName())
			fmt.Fprintf(w, "	if err := json.Unmarshal(buf, &item); err == nil {\n")
			fmt.Fprintf(w, "		*l = %sList{item}\n", resource.GoName())
			fmt.Fprintf(w, "		return nil\n")
			fmt.Fprintf(w, "	}\n")
			fmt.Fprintf(w, "	list := []%s{}\n", resource.GoName())
			fmt.Fprintf(w, "	err := json.Unmarshal(buf, &list)\n")
			fmt.Fprintf(w, "	if err == nil {\n")
			fmt.Fprintf(w, "		*l = %sList(list)\n", resource.GoName())
			fmt.Fprintf(w, "		return nil\n")
			fmt.Fprintf(w, "	}\n")
			fmt.Fprintf(w, "	return err\n")
			fmt.Fprintf(w, "}\n")
			fmt.Fprintf(w, "\n")
		}
	}

	fmt.Fprintf(w, "// NewResourceByType returns a new resource object correspoding with the provided type\n")
	fmt.Fprintf(w, "func NewResourceByType(typeName string) ResourceProperties {\n")
	fmt.Fprintf(w, "	switch typeName {\n")

	for _, resource := range tr.Resources {
		if !resource.IsTopLevelResource() {
			continue
		}
		fmt.Fprintf(w, "		case %q:\n", resource.Name)
		fmt.Fprintf(w, "			return &%s{}\n", resource.GoName())
	}
	fmt.Fprint(w, customResourceDefaultLabelBlock)

	fmt.Fprintf(w, "	}\n")
	fmt.Fprintf(w, "	return nil\n")
	fmt.Fprintf(w, "}\n")
}

type Resource struct {
	Name       string
	Href       string
	Properties PropertyList
}

type PropertyList []*Property

func (pl PropertyList) Get(name string) *Property {
	for _, p := range pl {
		if p.Name == name {
			return p
		}
	}
	return nil
}

func (r *Resource) IsTopLevelResource() bool {
	return strings.HasPrefix(r.Name, "AWS::")
}

func (r *Resource) Load() error {
	docReader, err := getDoc(r.Href)
	if err != nil {
		return err
	}

	doc, err := goquery.NewDocumentFromReader(docReader)
	if err != nil {
		return err
	}

	// An element with the class 'variablelist' immediately preceeded by an
	// element with the text "Properties" is what we're looking for.
	doc.Find(".variablelist").Each(func(i int, varList *goquery.Selection) {
		h2 := varList.Prev()
		for {
			if h2.Length() == 0 {
				return
			}
			if goquery.NodeName(h2) != "h2" {
				h2 = h2.Prev()
				continue
			}
			break
		}

		titleText := strings.TrimSpace(h2.Text())
		switch titleText {
		case "Properties":
		case "Parameters":
		case "Members":
		default:
			return
		}

		// The variableList contains a definition list. The DT element is the
		// name of the property, the following DD element contains information
		// about it, including the type.
		varList.Find("dl dt").Each(func(i int, dt *goquery.Selection) {
			property := Property{Name: strings.TrimSpace(dt.Text())}

			dd := dt.Next()

			property.DocString = dd.Find("p").First().Text()
			property.DocString = regexp.MustCompile("\\s+").ReplaceAllString(property.DocString, " ")
			property.DocString = strings.TrimSpace(property.DocString)
			property.DocString = wordwrap.WrapString(property.DocString, 70)

			// Somewhere inside the <DD> element there is a span that starts
			// with `Type: ` which is our type. Grab it along with the href
			// from an anchor (for complex types, this tells us which type it
			// refers to)
			dd.Find("em").Each(func(j int, span *goquery.Selection) {
				if span.Text() == "Type" {
					property.Type = span.Parent().Text()
					property.Type = strings.TrimPrefix(property.Type, "Type: ")
					property.Type = strings.TrimPrefix(property.Type, "Type:: ") // a typo in AWS::Route53::RecordSetGroup
					property.Type = regexp.MustCompile("\\s+").ReplaceAllString(property.Type, " ")
					property.Type = strings.TrimSpace(property.Type)
					property.Type = strings.TrimSuffix(property.Type, ".")
					property.Type = strings.TrimSpace(property.Type)

					span.Parent().Find("a").Each(func(j int, a *goquery.Selection) {
						property.TypeHref, _ = a.Attr("href")
					})
				}
			})

			if property.Name == "Icmp" && property.TypeHref == "aws-properties-ec2-networkaclentry-portrange.html" {
				property.Type = "EC2NetworkAclEntryIcmp"
				property.TypeHref = "aws-properties-ec2-networkaclentry-icmp.html"
			}

			r.Properties = append(r.Properties, &property)
		})
	})

	// There is at least one document, aws-properties-ec2-port-range.html that uses
	// a totally different format to describe object properties.
	doc.Find(".informaltable").Each(func(i int, varList *goquery.Selection) {
		if varList.Find("th").First().Text() != "Property" {
			return
		}

		varList.Find("tbody tr").Each(func(i int, tr *goquery.Selection) {
			property := Property{}
			tr.Find("td").Each(func(i int, p *goquery.Selection) {
				switch i {
				case 0:
					property.Name = p.Find("p").Text()
				case 1:
					property.Type = p.Find("p").Text()
				case 3:
					property.DocString = p.Find("p").Text()
					property.DocString = regexp.MustCompile("\\s+").ReplaceAllString(property.DocString, " ")
					property.DocString = wordwrap.WrapString(property.DocString, 70)
				}
			})
			r.Properties = append(r.Properties, &property)
		})
	})

	// Some documents contain a "Syntax" block that contains clues as to the
	// expected types of the various properties. Parse that block and capture
	// the syntax of each property in the SyntaxExpression field of the corresponding
	// property.
	doc.Find(".programlisting").Each(func(i int, varList *goquery.Selection) {
		headlineText := varList.Parent().Find("h3").First().Text()
		if headlineText != "JSON" {
			return
		}
		nextHeadlineText := varList.Parent().Parent().Find("h2").First().Text()
		if !strings.HasPrefix(nextHeadlineText, "Syntax") {
			return
		}

		for _, line := range strings.Split(varList.Text(), "\n") {
			if !strings.Contains(line, ":") {
				continue
			}
			parts := strings.SplitN(line, ":", 2)
			propertyName := strings.Trim(parts[0], " \"")
			expression := strings.Trim(parts[1], " \",")

			p := r.Properties.Get(propertyName)
			if p == nil {
				continue
			}
			p.SyntaxExpression = expression
		}
	})

	return nil
}

func (r *Resource) GoName() string {
	rv := r.Name
	rv = strings.TrimPrefix(rv, "AWS CloudFormation ")
	rv = strings.TrimPrefix(rv, "AWS ")
	rv = strings.TrimPrefix(rv, "Amazon ")
	rv = strings.TrimPrefix(rv, "AWS::")
	rv = strings.Replace(rv, "::", "", -1)
	rv = regexp.MustCompile("\\W").ReplaceAllString(rv, "")
	rv = strings.TrimSuffix(rv, "PropertyType")
	rv = strings.TrimSuffix(rv, "Type")

	if rv == "ResourceTags" {
		rv = "ResourceTag"
	}

	// There is an object names AWS::EC2::NetworkInterfaceAttachment and an
	// object named EC2NetworkInterfaceAttachmentType. To avoid a duplicate
	// definition, we have to deconflict them here.
	if r.Name == "EC2 Network Interface Attachment" {
		return "EC2NetworkInterfaceAttachmentType"
	}

	// There is an object named AWS::SNS::Subscription and an object named
	// SNS Subscription. To avoid a duplication definition, we have to deconflict
	// them here.
	if r.Name == "Amazon SNS Subscription Property Type" {
		return "SNSSubscriptionProperty"
	}

	// Note: If we find one more conflict, we should do something more clever here

	return rv
}

const rootURL = "http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/"

// http://godoc.org/golang.org/x/net/html
// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html
// Topics -> <li><a href="aws-properties-as-group.html">AWS::AutoScaling::AutoScalingGroup</a></li>

func getDoc(url string) (io.ReadCloser, error) {
	_, cachePath := path.Split(url)
	cachePath = path.Join("./.scraper-cache", cachePath)
	os.MkdirAll(path.Dir(cachePath), 0755)

	d, err := os.Open(cachePath)
	if err == nil {
		return d, nil
	}
	if !os.IsNotExist(err) {
		return nil, err
	}
	res, err := http.Get(rootURL + url)
	if err != nil {
		return nil, err
	}

	d, err = os.Create(cachePath)
	if err != nil {
		return nil, err
	}
	io.Copy(d, res.Body)
	d.Close()

	return os.Open(cachePath)
}

type Property struct {
	Name             string
	Type             string
	TypeHref         string
	TypeName         string
	DocString        string
	SyntaxExpression string
}

func (p *Property) GoName() string {
	rv := strings.Title(p.Name)
	rv = strings.Replace(rv, ".", "", -1)
	rv = regexp.MustCompile("[^A-Za-z0-9]").ReplaceAllString(rv, "X")
	return rv
}

func (p *Property) Comment(prefix string) string {
	c := p.DocString
	c = strings.Replace(c, "\n\n: ", ": ", -1)
	c = strings.Replace(c, "\n\n:\n", ": ", -1)
	c = strings.Replace(c, "\n\n", "\n", -1)
	c = strings.Replace(c, "\n", "\n"+prefix, -1)
	return prefix + c + "\n"
}

func (p *Property) GoType(tr *TemplateReference) string {
	if p.TypeHref != "" {
		for _, res := range tr.Resources {
			if res.Href == p.TypeHref {
				p.TypeName = "*" + res.GoName()
			}
		}
		if p.TypeName == "" {
			p.TypeName = "[UNKNOWN " + p.TypeHref + "]"
		}
	}

	isMaybeList := false
	if strings.HasPrefix(p.Type, "A list of") ||
		strings.HasPrefix(p.Type, "List of") ||
		strings.HasPrefix(p.Type, "list of") ||
		strings.HasPrefix(p.Type, "Array of") ||
		strings.HasPrefix(p.SyntaxExpression, "[") {
		isMaybeList = true
	}

	if p.TypeName != "" {
		if p.Type == "AWS CloudFormation Resource Tags" {
			return "[]ResourceTag"
		}

		if isMaybeList {
			return p.TypeName + "List"
		}

		return p.TypeName
	}

	switch p.Type {
	case "String":
		if isMaybeList {
			return "*StringListExpr"
		}
		return "*StringExpr"
	case "List of strings":
		return "*StringListExpr"
	case "String list":
		return "*StringListExpr"
	case "Boolean":
		return "*BoolExpr"
	case "Integer":
		return "*IntegerExpr"
	case "Number":
		return "*IntegerExpr"
	case "Time stamp":
		return "time.Time"
	}

	if strings.HasPrefix(p.Type, "Number") {
		return "*IntegerExpr"
	}
	if strings.HasPrefix(p.Type, "String") {
		return "*StringExpr"
	}

	return "interface{}"
}

func main() {
	var format = flag.String("format", "go", "How to write the output, either `json` or `go`.")
	var outPath = flag.String("out", "", "The output path")

	flag.Parse()

	tr := TemplateReference{}
	if err := tr.Load(); err != nil {
		log.Fatal(err)
	}

	var out io.Writer
	if *outPath == "-" {
		out = os.Stdout
	} else {
		var err error
		out, err = os.OpenFile(*outPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatalf("%s: %s", *outPath, err)
		}
	}

	switch *format {
	case "go":
		if err := tr.WriteGo(out); err != nil {
			log.Fatalf("%s", err)
		}
	case "json":
		buf, err := json.MarshalIndent(tr, "", "  ")
		if err != nil {
			log.Fatalf("%s", err)
		}
		_, err = out.Write(buf)
		if err != nil {
			log.Fatalf("%s", err)
		}
	default:
		fmt.Fprintf(os.Stderr, "unrecognised output format: %q", *format)
	}
}
