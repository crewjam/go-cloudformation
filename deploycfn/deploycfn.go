package deploycfn

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/crewjam/awsregion"
	cf "github.com/crewjam/go-cloudformation"
)

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
func Deploy(awsConfigProvider client.ConfigProvider, stackName string, template *cf.Template) error {
	templateBody, err := json.Marshal(template)
	if err != nil {
		return err
	}

	if awsConfigProvider == nil {
		awsSession := session.New()
		awsregion.GuessRegion(awsSession.Config)
		awsConfigProvider = awsSession
	}
	cfn := cloudformation.New(awsConfigProvider)

	describeStacksResponse, err := cfn.DescribeStacks(&cloudformation.DescribeStacksInput{
		StackName: aws.String(stackName),
	})
	doCreate := err != nil || len(describeStacksResponse.Stacks) == 0

	// seenEvents maps the events that have already been printed. If the stack
	// already exists, then we want to get all the existing events into this
	// set so we don't re-print old events.
	seenEvents := map[string]struct{}{}
	if !doCreate {
		cfn.DescribeStackEventsPages(&cloudformation.DescribeStackEventsInput{
			StackName: aws.String(stackName),
		}, func(p *cloudformation.DescribeStackEventsOutput, _ bool) bool {
			for _, stackEvent := range p.StackEvents {
				seenEvents[*stackEvent.EventId] = struct{}{}
			}
			return true
		})
	}

	if doCreate {
		_, err = cfn.CreateStack(&cloudformation.CreateStackInput{
			StackName:    aws.String(stackName),
			TemplateBody: aws.String(string(templateBody)),
		})
		if err != nil {
			return err
		}
	} else {
		_, err = cfn.UpdateStack(&cloudformation.UpdateStackInput{
			StackName:    aws.String(stackName),
			TemplateBody: aws.String(string(templateBody)),
		})
		if err != nil {
			return err
		}
	}

	lastStackStatus := ""
	for {
		// print the events for the stack
		cfn.DescribeStackEventsPages(&cloudformation.DescribeStackEventsInput{
			StackName: aws.String(stackName),
		}, func(p *cloudformation.DescribeStackEventsOutput, _ bool) bool {
			for _, stackEvent := range p.StackEvents {
				if _, ok := seenEvents[*stackEvent.EventId]; ok {
					continue
				}
				wrapStrPtr := func(s *string) string {
					if s == nil {
						return ""
					}
					return *s
				}
				log.Printf("%s\t%s\t%s\t%s\t%s\n",
					wrapStrPtr(stackEvent.ResourceStatus),
					wrapStrPtr(stackEvent.ResourceType),
					wrapStrPtr(stackEvent.PhysicalResourceId),
					wrapStrPtr(stackEvent.LogicalResourceId),
					wrapStrPtr(stackEvent.ResourceStatusReason))
				seenEvents[*stackEvent.EventId] = struct{}{}
			}
			return true
		})

		// monitor the status of the stack
		describeStacksResponse, err := cfn.DescribeStacks(&cloudformation.DescribeStacksInput{
			StackName: aws.String(stackName),
		})
		if err != nil {
			// the stack might not exist yet
			log.Printf("DescribeStacks: %s", err)
			time.Sleep(time.Second)
			continue
		}

		stackStatus := *describeStacksResponse.Stacks[0].StackStatus
		if stackStatus != lastStackStatus {
			log.Printf("Stack: %s\n", stackStatus)
			lastStackStatus = stackStatus
		}
		switch stackStatus {
		case cloudformation.StackStatusCreateComplete:
			return nil
		case cloudformation.StackStatusCreateFailed:
			return fmt.Errorf("%s", stackStatus)
		case cloudformation.StackStatusRollbackComplete:
			return fmt.Errorf("%s", stackStatus)
		case cloudformation.StackStatusUpdateRollbackComplete:
			return fmt.Errorf("%s", stackStatus)
		case cloudformation.StackStatusRollbackFailed:
			return fmt.Errorf("%s", stackStatus)
		case cloudformation.StackStatusUpdateComplete:
			return nil
		case cloudformation.StackStatusUpdateRollbackFailed:
			return fmt.Errorf("%s", stackStatus)
		default:
			time.Sleep(time.Second * 5)
			continue
		}
	}
}
