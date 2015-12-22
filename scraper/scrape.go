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
			href, _ := s.Attr("href")
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

			fmt.Fprintf(w, "// ResourceType returns %s to implement the ResourceProperties interface\n", resource.Name)
			fmt.Fprintf(w, "func (s %s) ResourceType() string {\n", resource.GoName())
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
	Properties []Property
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
		if varList.Parent().Find(".titlepage").First().Text() != "Properties" {
			return
		}

		// The variableList contains a definition list. The DT element is the
		// name of the property, the following DD element contains information
		// about it, including the type.
		varList.Find("dl dt").Each(func(i int, dt *goquery.Selection) {
			property := Property{Name: dt.Text()}

			dd := dt.Next()

			property.DocString = dd.Find("p").First().Text()
			property.DocString = regexp.MustCompile("\\s+").ReplaceAllString(property.DocString, " ")
			property.DocString = wordwrap.WrapString(property.DocString, 70)
			/*
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s\n", err)
					return
				}
				property.DocString, err = html2text.FromString(docString)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s\n", err)
					return
				}
			*/

			// Somewhere inside the <DD> element there is a span that starts
			// with `Type: ` which is our type. Grab it along with the href
			// from an anchor (for complex types, this tells us which type it
			// refers to)
			dd.Find("span").Each(func(j int, span *goquery.Selection) {
				if span.Text() == "Type" {
					property.Type = span.Parent().Text()
					property.Type = strings.TrimPrefix(property.Type, "Type: ")
					property.Type = strings.TrimPrefix(property.Type, "Type:: ") // a typo in AWS::Route53::RecordSetGroup
					property.Type = regexp.MustCompile("\\s+").ReplaceAllString(property.Type, " ")
					property.Type = strings.TrimSuffix(property.Type, ".")

					span.Parent().Find("a").Each(func(j int, a *goquery.Selection) {
						property.TypeHref, _ = a.Attr("href")
					})
				}
			})
			r.Properties = append(r.Properties, property)
		})
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
	Name      string
	Type      string
	TypeHref  string
	TypeName  string
	DocString string
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
	if p.TypeName != "" {
		if strings.HasPrefix(p.Type, "A list of") ||
			strings.HasPrefix(p.Type, "List of") ||
			strings.HasPrefix(p.Type, "list of") {
			return p.TypeName + "List"
		}
		// In various places the documentation omit the "list of"
		// when describing types, but include the phrase "list" in the
		// docstring. For example SecurityGroupEgress and SecurityGroupIngress in
		// AWS::EC2::SecurityGroup as "EC2 Security Group Rule" when
		// it should be "list of EC2 Security Group Rule"
		// c.f. http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html#cfn-ec2-securitygroup-securitygroupegress
		if strings.HasPrefix(p.DocString, "A list") {
			return p.TypeName + "List"
		}

		if p.Type == "AWS CloudFormation Resource Tags" {
			return "[]ResourceTag"
		}

		return p.TypeName
	}

	switch p.Type {
	case "String":
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
		if err := json.NewEncoder(out).Encode(tr); err != nil {
			log.Fatalf("%s", err)
		}
	default:
		fmt.Fprintf(os.Stderr, "unrecognised output format: %q", *format)
	}
}
