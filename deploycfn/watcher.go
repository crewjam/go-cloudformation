package deploycfn

import (
	"fmt"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

// StackEventWatcher watches a CloudFormation stack for events and emits
// them to the log channel.
type StackEventWatcher struct {
	Service   *cloudformation.CloudFormation
	StackName string

	// seenEvents maps the events that have already been printed. If the stack
	// already exists, then we want to get all the existing events into this
	// set so we don't re-print old events.
	seenEvents map[string]struct{}
}

// NewStackEventWatcher returns a new StackEventWatcher that emits events for the
// specified stack. It scans all the existing events and adds them to seenEvents
// so that events the occur prior to the invocation of this function will not be
// printed.
func NewStackEventWatcher(session client.ConfigProvider, stackName string) (*StackEventWatcher, error) {
	sw := StackEventWatcher{
		Service:    cloudformation.New(session),
		StackName:  stackName,
		seenEvents: map[string]struct{}{},
	}

	err := sw.Service.DescribeStackEventsPages(&cloudformation.DescribeStackEventsInput{
		StackName: aws.String(sw.StackName),
	}, func(p *cloudformation.DescribeStackEventsOutput, _ bool) bool {
		for _, stackEvent := range p.StackEvents {
			sw.seenEvents[*stackEvent.EventId] = struct{}{}
		}
		return true
	})
	if err != nil {
		return nil, err
	}
	return &sw, nil
}

// Watch monitors the stack for events, reporting each unseen event to the
// log channel. Returns when the stack enters a non-transitional state. The
// return value is non-nil if the final state is an error state.
func (sw *StackEventWatcher) Watch() error {
	if sw.seenEvents == nil {
		sw.seenEvents = map[string]struct{}{}
	}
	lastStackStatus := ""
	for {
		// print the events for the stack
		sw.Service.DescribeStackEventsPages(&cloudformation.DescribeStackEventsInput{
			StackName: aws.String(sw.StackName),
		}, func(p *cloudformation.DescribeStackEventsOutput, _ bool) bool {
			for _, stackEvent := range p.StackEvents {
				if _, ok := sw.seenEvents[*stackEvent.EventId]; ok {
					continue
				}
				wrapStrPtr := func(s *string) string {
					if s == nil {
						return ""
					}
					return *s
				}

				l := log.WithField("Status", *stackEvent.ResourceStatus)
				if stackEvent.ResourceType != nil {
					l = l.WithField("Type", *stackEvent.ResourceType)
				}
				if stackEvent.ResourceType != nil {
					l = l.WithField("Type", *stackEvent.ResourceType)
				}
				if stackEvent.PhysicalResourceId != nil {
					l = l.WithField("PhysicalID", *stackEvent.PhysicalResourceId)
				}
				if stackEvent.LogicalResourceId != nil {
					l = l.WithField("LogicalID", *stackEvent.LogicalResourceId)
				}
				if strings.Contains(*stackEvent.ResourceStatus, "FAIL") {
					l.Error(wrapStrPtr(stackEvent.ResourceStatusReason))
				} else {
					l.Info(wrapStrPtr(stackEvent.ResourceStatusReason))
				}

				sw.seenEvents[*stackEvent.EventId] = struct{}{}
			}
			return true
		})

		// monitor the status of the stack
		describeStacksResponse, err := sw.Service.DescribeStacks(&cloudformation.DescribeStacksInput{
			StackName: aws.String(sw.StackName),
		})
		if err != nil {
			// the stack might not exist yet
			log.Errorf("DescribeStacks: %s", err)
			time.Sleep(time.Second)
			continue
		}

		stackStatus := *describeStacksResponse.Stacks[0].StackStatus
		if stackStatus != lastStackStatus {
			log.Infof("Stack: %s\n", stackStatus)
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
