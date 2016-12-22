package deploycfn

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/crewjam/awsregion"
	cfn "github.com/crewjam/go-cloudformation"
)

type DeployInput struct {
	Session        client.ConfigProvider
	StackName      string
	Template       *cfn.Template
	Parameters     map[string]string
	TemplateBucket string
}

// Deploy creates or updates the specified CloudFormation template to AWS.
//
// If you secify nil for awsConfigProvider, a default one will be used.
//
// This is a handy function to call from your main, i.e.:
//
//   func main() {
//   	dnsName := flag.String("name", "example.com", "the DNS name")
//   	flag.Parse()
//
//   	template := makeTemplate(*dnsName)
//   	if err := deploycfn.Deploy(nil, "Example", template); err != nil {
//   		log.Fatalf("deploy: %s", err)
//   	}
//   }
//
func Deploy(input DeployInput) error {
	templateBody, err := json.Marshal(input.Template)
	if err != nil {
		return err
	}

	if input.Session == nil {
		awsSession := session.New()
		awsregion.GuessRegion(awsSession.Config)
		input.Session = awsSession
	}
	cfnSvc := cloudformation.New(input.Session)

	describeStacksResponse, err := cfnSvc.DescribeStacks(&cloudformation.DescribeStacksInput{
		StackName: aws.String(input.StackName),
	})
	doCreate := err != nil || len(describeStacksResponse.Stacks) == 0

	watcher := &StackEventWatcher{
		Service:   cfnSvc,
		StackName: input.StackName,
	}
	if !doCreate {
		watcher, err = NewStackEventWatcher(input.Session, input.StackName)
		if err != nil {
			return err
		}
	}

	needCapabilityIam := false
	for _, resource := range input.Template.Resources {
		switch resource.Properties.CfnResourceType() {
		case "AWS::IAM::AccessKey":
			needCapabilityIam = true
		case "AWS::IAM::Group":
			needCapabilityIam = true
		case "AWS::IAM::InstanceProfile":
			needCapabilityIam = true
		case "AWS::IAM::Policy":
			needCapabilityIam = true
		case "AWS::IAM::Role":
			needCapabilityIam = true
		case "AWS::IAM::User":
			needCapabilityIam = true
		case "AWS::IAM::UserToGroupAddition":
			needCapabilityIam = true
		}
	}

	capabilities := []*string{}
	if needCapabilityIam {
		capabilities = append(capabilities, aws.String(cloudformation.CapabilityCapabilityIam))
	}

	parameters := []*cloudformation.Parameter{}
	for key, value := range input.Parameters {
		parameters = append(parameters, &cloudformation.Parameter{
			ParameterKey:   aws.String(key),
			ParameterValue: aws.String(value),
		})
	}
	if !doCreate {
		for key := range input.Template.Parameters {
			if _, ok := input.Parameters[key]; ok {
				continue
			}

			parameterExistsInCurrentStack := false
			for _, p := range describeStacksResponse.Stacks[0].Parameters {
				if *p.ParameterKey == key {
					parameterExistsInCurrentStack = true
				}
			}
			if parameterExistsInCurrentStack {
				parameters = append(parameters, &cloudformation.Parameter{
					ParameterKey:     aws.String(key),
					UsePreviousValue: aws.Bool(true),
				})
			}
		}
	}

	var templateBodyStr = aws.String(string(templateBody))
	var templateURL *string

	const maxTemplateSize = 51200
	if input.TemplateBucket != "" && len(templateBody) > maxTemplateSize {
		s3svc := s3.New(input.Session)

		sha1sum := sha1.Sum(templateBody)
		key := fmt.Sprintf("cfn/%x/%s.template", sha1sum[:], input.StackName)
		_, err := s3svc.PutObject(&s3.PutObjectInput{
			Bucket: &input.TemplateBucket,
			Key:    &key,
			Body:   bytes.NewReader(templateBody),
		})
		if err != nil {
			return err
		}
		templateBodyStr = nil
		templateURL = aws.String(fmt.Sprintf("https://%s.s3.amazonaws.com/%s", input.TemplateBucket, key))
	}

	if doCreate {
		_, err = cfnSvc.CreateStack(&cloudformation.CreateStackInput{
			StackName:    aws.String(input.StackName),
			TemplateBody: templateBodyStr,
			TemplateURL:  templateURL,
			Capabilities: capabilities,
			Parameters:   parameters,
		})
		if err != nil {
			return err
		}
	} else {
		_, err = cfnSvc.UpdateStack(&cloudformation.UpdateStackInput{
			StackName:    aws.String(input.StackName),
			TemplateBody: templateBodyStr,
			TemplateURL:  templateURL,
			Capabilities: capabilities,
			Parameters:   parameters,
		})
		if err != nil {
			return err
		}
	}

	if err := watcher.Watch(); err != nil {
		return err
	}

	return nil
}
