package cloudformation

import "time"
import "encoding/json"

// AutoScalingAutoScalingGroup represents AWS::AutoScaling::AutoScalingGroup
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html
type AutoScalingAutoScalingGroup struct {
	// Contains a list of availability zones for the group.
	AvailabilityZones *StringListExpr `json:"AvailabilityZones,omitempty"`

	// The number of seconds after a scaling activity is completed before any
	// further scaling activities can start.
	Cooldown *StringExpr `json:"Cooldown,omitempty"`

	// Specifies the desired capacity for the Auto Scaling group.
	DesiredCapacity *StringExpr `json:"DesiredCapacity,omitempty"`

	// The length of time in seconds after a new EC2 instance comes into
	// service that Auto Scaling starts checking its health.
	HealthCheckGracePeriod *IntegerExpr `json:"HealthCheckGracePeriod,omitempty"`

	// The service you want the health status from, Amazon EC2 or Elastic
	// Load Balancer. Valid values are EC2 or ELB.
	HealthCheckType *StringExpr `json:"HealthCheckType,omitempty"`

	// The ID of the Amazon EC2 instance you want to use to create the Auto
	// Scaling group. Use this property if you want to create an Auto Scaling
	// group that uses an existing Amazon EC2 instance instead of a launch
	// configuration.
	InstanceId *StringExpr `json:"InstanceId,omitempty"`

	// Specifies the name of the associated
	// AWS::AutoScaling::LaunchConfiguration.
	LaunchConfigurationName *StringExpr `json:"LaunchConfigurationName,omitempty"`

	// A list of load balancers associated with this Auto Scaling group.
	LoadBalancerNames *StringListExpr `json:"LoadBalancerNames,omitempty"`

	// The maximum size of the Auto Scaling group.
	MaxSize *StringExpr `json:"MaxSize,omitempty"`

	// Enables the monitoring of group metrics of an Auto Scaling group.
	MetricsCollection *AutoScalingMetricsCollectionList `json:"MetricsCollection,omitempty"`

	// The minimum size of the Auto Scaling group.
	MinSize *StringExpr `json:"MinSize,omitempty"`

	// An embedded property that configures an Auto Scaling group to send
	// notifications when specified events take place.
	NotificationConfigurations *AutoScalingNotificationConfigurationsList `json:"NotificationConfigurations,omitempty"`

	// The name of an existing cluster placement group into which you want to
	// launch your instances. A placement group is a logical grouping of
	// instances within a single Availability Zone. You cannot specify
	// multiple Availability Zones and a placement group.
	PlacementGroup *StringExpr `json:"PlacementGroup,omitempty"`

	// The tags you want to attach to this resource.
	Tags *AutoScalingTagsList `json:"Tags,omitempty"`

	// A policy or a list of policies that are used to select the instances
	// to terminate. The policies are executed in the order that you list
	// them.
	TerminationPolicies *StringListExpr `json:"TerminationPolicies,omitempty"`

	// A list of subnet identifiers of Amazon Virtual Private Cloud (Amazon
	// VPCs).
	VPCZoneIdentifier *StringListExpr `json:"VPCZoneIdentifier,omitempty"`
}

// ResourceType returns AWS::AutoScaling::AutoScalingGroup to implement the ResourceProperties interface
func (s AutoScalingAutoScalingGroup) ResourceType() string {
	return "AWS::AutoScaling::AutoScalingGroup"
}

// AutoScalingLaunchConfiguration represents AWS::AutoScaling::LaunchConfiguration
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html
type AutoScalingLaunchConfiguration struct {
	// For Amazon EC2 instances in a VPC, indicates whether instances in the
	// Auto Scaling group receive public IP addresses. If you specify true,
	// each instance in the Auto Scaling receives a unique public IP address.
	AssociatePublicIpAddress *BoolExpr `json:"AssociatePublicIpAddress,omitempty"`

	// Specifies how block devices are exposed to the instance. You can
	// specify virtual devices and EBS volumes.
	BlockDeviceMappings *AutoScalingBlockDeviceMappingList `json:"BlockDeviceMappings,omitempty"`

	// The ID of a ClassicLink-enabled VPC to link your EC2-Classic instances
	// to. You can specify this property only for EC2-Classic instances. For
	// more information, see ClassicLink in the Amazon Elastic Compute Cloud
	// User Guide.
	ClassicLinkVPCId *StringExpr `json:"ClassicLinkVPCId,omitempty"`

	// The IDs of one or more security groups for the VPC that you specified
	// in the ClassicLinkVPCId property.
	ClassicLinkVPCSecurityGroups *StringListExpr `json:"ClassicLinkVPCSecurityGroups,omitempty"`

	// Specifies whether the launch configuration is optimized for EBS I/O.
	// This optimization provides dedicated throughput to Amazon EBS and an
	// optimized configuration stack to provide optimal EBS I/O performance.
	EbsOptimized *BoolExpr `json:"EbsOptimized,omitempty"`

	// Provides the name or the Amazon Resource Name (ARN) of the instance
	// profile associated with the IAM role for the instance. The instance
	// profile contains the IAM role.
	IamInstanceProfile *StringExpr `json:"IamInstanceProfile,omitempty"`

	// Provides the unique ID of the Amazon Machine Image (AMI) that was
	// assigned during registration.
	ImageId *StringExpr `json:"ImageId,omitempty"`

	// The ID of the Amazon EC2 instance you want to use to create the launch
	// configuration. Use this property if you want the launch configuration
	// to use settings from an existing Amazon EC2 instance.
	InstanceId *StringExpr `json:"InstanceId,omitempty"`

	// Indicates whether detailed instance monitoring is enabled for the Auto
	// Scaling group. By default, this property is set to true (enabled).
	InstanceMonitoring *BoolExpr `json:"InstanceMonitoring,omitempty"`

	// Specifies the instance type of the EC2 instance.
	InstanceType *StringExpr `json:"InstanceType,omitempty"`

	// Provides the ID of the kernel associated with the EC2 AMI.
	KernelId *StringExpr `json:"KernelId,omitempty"`

	// Provides the name of the EC2 key pair.
	KeyName *StringExpr `json:"KeyName,omitempty"`

	// The tenancy of the instance. An instance with a tenancy of dedicated
	// runs on single-tenant hardware and can only be launched in a VPC. You
	// must set the value of this parameter to dedicated if want to launch
	// dedicated instances in a shared tenancy VPC (a VPC with the instance
	// placement tenancy attribute set to default). For more information, see
	// CreateLaunchConfiguration in the Auto Scaling API Reference.
	PlacementTenancy *StringExpr `json:"PlacementTenancy,omitempty"`

	// The ID of the RAM disk to select. Some kernels require additional
	// drivers at launch. Check the kernel requirements for information about
	// whether you need to specify a RAM disk. To find kernel requirements,
	// refer to the AWS Resource Center and search for the kernel ID.
	RamDiskId *StringExpr `json:"RamDiskId,omitempty"`

	// A list that contains the EC2 security groups to assign to the Amazon
	// EC2 instances in the Auto Scaling group. The list can contain the name
	// of existing EC2 security groups or references to
	// AWS::EC2::SecurityGroup resources created in the template. If your
	// instances are launched within VPC, specify Amazon VPC security group
	// IDs.
	SecurityGroups interface{} `json:"SecurityGroups,omitempty"`

	// The spot price for this autoscaling group. If a spot price is set,
	// then the autoscaling group will launch when the current spot price is
	// less than the amount specified in the template.
	SpotPrice *StringExpr `json:"SpotPrice,omitempty"`

	// The user data available to the launched EC2 instances.
	UserData *StringExpr `json:"UserData,omitempty"`
}

// ResourceType returns AWS::AutoScaling::LaunchConfiguration to implement the ResourceProperties interface
func (s AutoScalingLaunchConfiguration) ResourceType() string {
	return "AWS::AutoScaling::LaunchConfiguration"
}

// AutoScalingLifecycleHook represents AWS::AutoScaling::LifecycleHook
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html
type AutoScalingLifecycleHook struct {
	// The name of the Auto Scaling group for the lifecycle hook.
	AutoScalingGroupName *StringExpr `json:"AutoScalingGroupName,omitempty"`

	// The action the Auto Scaling group takes when the lifecycle hook
	// timeout elapses or if an unexpected failure occurs.
	DefaultResult *StringExpr `json:"DefaultResult,omitempty"`

	// The amount of time that can elapse before the lifecycle hook times
	// out. When the lifecycle hook times out, Auto Scaling performs the
	// action that you specified in the DefaultResult property.
	HeartbeatTimeout *IntegerExpr `json:"HeartbeatTimeout,omitempty"`

	// The state of the Amazon EC2 instance to which you want to attach the
	// lifecycle hook.
	LifecycleTransition *StringExpr `json:"LifecycleTransition,omitempty"`

	// Additional information that you want to include when Auto Scaling
	// sends a message to the notification target.
	NotificationMetadata *StringExpr `json:"NotificationMetadata,omitempty"`

	// The Amazon resource name (ARN) of the notification target that Auto
	// Scaling uses to notify you when an instance is in the transition state
	// for the lifecycle hook. You can specify an Amazon SQS queue or an
	// Amazon SNS topic. The notification message includes the following
	// information: lifecycle action token, user account ID, Auto Scaling
	// group name, lifecycle hook name, instance ID, lifecycle transition,
	// and notification metadata.
	NotificationTargetARN *StringExpr `json:"NotificationTargetARN,omitempty"`

	// The ARN of the IAM role that allows the Auto Scaling group to publish
	// to the specified notification target. The role requires permissions to
	// Amazon SNS and Amazon SQS.
	RoleARN *StringExpr `json:"RoleARN,omitempty"`
}

// ResourceType returns AWS::AutoScaling::LifecycleHook to implement the ResourceProperties interface
func (s AutoScalingLifecycleHook) ResourceType() string {
	return "AWS::AutoScaling::LifecycleHook"
}

// AutoScalingScalingPolicy represents AWS::AutoScaling::ScalingPolicy
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html
type AutoScalingScalingPolicy struct {
	// Specifies whether the ScalingAdjustment is an absolute number or a
	// percentage of the current capacity. Valid values are ChangeInCapacity,
	// ExactCapacity, and PercentChangeInCapacity.
	AdjustmentType *StringExpr `json:"AdjustmentType,omitempty"`

	// The name or Amazon Resource Name (ARN) of the Auto Scaling Group that
	// you want to attach the policy to.
	AutoScalingGroupName *StringExpr `json:"AutoScalingGroupName,omitempty"`

	// The amount of time, in seconds, after a scaling activity completes
	// before any further trigger-related scaling activities can start.
	Cooldown *StringExpr `json:"Cooldown,omitempty"`

	// The minmum number of instances that are added or removed when the Auto
	// Scaling group scales up or down. You can use this property only when
	// you specify PercentChangeInCapacity for the AdjustmentType property.
	MinAdjustmentStep *IntegerExpr `json:"MinAdjustmentStep,omitempty"`

	// The number of instances by which to scale. AdjustmentType determines
	// the interpretation of this number, such as an absolute number or as a
	// percentage of the existing Auto Scaling group size. A positive
	// increment adds to the current capacity and a negative value removes
	// from the current capacity.
	ScalingAdjustment *StringExpr `json:"ScalingAdjustment,omitempty"`
}

// ResourceType returns AWS::AutoScaling::ScalingPolicy to implement the ResourceProperties interface
func (s AutoScalingScalingPolicy) ResourceType() string {
	return "AWS::AutoScaling::ScalingPolicy"
}

// AutoScalingScheduledAction represents AWS::AutoScaling::ScheduledAction
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html
type AutoScalingScheduledAction struct {
	// The name or ARN of the Auto Scaling group.
	AutoScalingGroupName *StringExpr `json:"AutoScalingGroupName,omitempty"`

	// The number of Amazon EC2 instances that should be running in the Auto
	// Scaling group.
	DesiredCapacity *IntegerExpr `json:"DesiredCapacity,omitempty"`

	// The time in UTC for this schedule to end. For example,
	// 2010-06-01T00:00:00Z.
	EndTime time.Time `json:"EndTime,omitempty"`

	// The maximum number of Amazon EC2 instances in the Auto Scaling group.
	MaxSize *IntegerExpr `json:"MaxSize,omitempty"`

	// The minimum number of Amazon EC2 instances in the Auto Scaling group.
	MinSize *IntegerExpr `json:"MinSize,omitempty"`

	// The time in UTC when recurring future actions will start. You specify
	// the start time by following the Unix cron syntax format. For more
	// information about cron syntax, go to
	// http://en.wikipedia.org/wiki/Cron.
	Recurrence *StringExpr `json:"Recurrence,omitempty"`

	// The time in UTC for this schedule to start. For example,
	// 2010-06-01T00:00:00Z.
	StartTime time.Time `json:"StartTime,omitempty"`
}

// ResourceType returns AWS::AutoScaling::ScheduledAction to implement the ResourceProperties interface
func (s AutoScalingScheduledAction) ResourceType() string {
	return "AWS::AutoScaling::ScheduledAction"
}

// CloudFormationAuthentication represents AWS::CloudFormation::Authentication
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-authentication.html
type CloudFormationAuthentication struct {
	// Specifies the access key ID for S3 authentication.
	AccessKeyId *StringExpr `json:"accessKeyId,omitempty"`

	// A comma-delimited list of Amazon S3 buckets to be associated with the
	// S3 authentication credentials.
	Buckets *StringListExpr `json:"buckets,omitempty"`

	// Specifies the password for basic authentication.
	Password *StringExpr `json:"password,omitempty"`

	// Specifies the secret key for S3 authentication.
	SecretKey *StringExpr `json:"secretKey,omitempty"`

	// Specifies whether the authentication scheme uses a user name and
	// password ("basic") or an access key ID and secret key ("S3").
	Type *StringExpr `json:"type,omitempty"`

	// A comma-delimited list of URIs to be associated with the basic
	// authentication credentials. The authorization applies to the specified
	// URIs and any more specific URI. For example, if you specify
	// http://www.example.com, the authorization will also apply to
	// http://www.example.com/test.
	Uris *StringListExpr `json:"uris,omitempty"`

	// Specifies the user name for basic authentication.
	Username *StringExpr `json:"username,omitempty"`

	// Describes the role for role-based authentication.
	RoleName *StringExpr `json:"roleName,omitempty"`
}

// ResourceType returns AWS::CloudFormation::Authentication to implement the ResourceProperties interface
func (s CloudFormationAuthentication) ResourceType() string {
	return "AWS::CloudFormation::Authentication"
}

// CloudFormationCustomResource represents AWS::CloudFormation::CustomResource
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cfn-customresource.html
type CloudFormationCustomResource struct {
	// The service token that was given to the template developer by the
	// service provider to access the service, such as an Amazon SNS topic
	// ARN or Lambda function ARN. The service token must be from the same
	// region in which you are creating the stack.
	ServiceToken *StringExpr `json:"ServiceToken,omitempty"`
}

// ResourceType returns AWS::CloudFormation::CustomResource to implement the ResourceProperties interface
func (s CloudFormationCustomResource) ResourceType() string {
	return "AWS::CloudFormation::CustomResource"
}

// CloudFormationInit represents AWS::CloudFormation::Init
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-init.html
type CloudFormationInit struct {
}

// ResourceType returns AWS::CloudFormation::Init to implement the ResourceProperties interface
func (s CloudFormationInit) ResourceType() string {
	return "AWS::CloudFormation::Init"
}

// CloudFormationStack represents AWS::CloudFormation::Stack
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stack.html
type CloudFormationStack struct {
	// A list of existing Amazon SNS topics where notifications about stack
	// events are sent.
	NotificationARNs *StringListExpr `json:"NotificationARNs,omitempty"`

	// The set of parameters passed to AWS CloudFormation when this nested
	// stack is created.
	Parameters *CloudFormationStackParameters `json:"Parameters,omitempty"`

	// The URL of a template that specifies the stack that you want to create
	// as a resource. The template must be stored on an Amazon S3 bucket, so
	// the URL must have the form:
	// https://s3.amazonaws.com/.../TemplateName.template
	TemplateURL *StringExpr `json:"TemplateURL,omitempty"`

	// The length of time, in minutes, that AWS CloudFormation waits for the
	// nested stack to reach the CREATE_COMPLETE state. The default is no
	// timeout. When AWS CloudFormation detects that the nested stack has
	// reached the CREATE_COMPLETE state, it marks the nested stack resource
	// as CREATE_COMPLETE in the parent stack and resumes creating the parent
	// stack. If the timeout period expires before the nested stack reaches
	// CREATE_COMPLETE, AWS CloudFormation marks the nested stack as failed
	// and rolls back both the nested stack and parent stack.
	TimeoutInMinutes *StringExpr `json:"TimeoutInMinutes,omitempty"`
}

// ResourceType returns AWS::CloudFormation::Stack to implement the ResourceProperties interface
func (s CloudFormationStack) ResourceType() string {
	return "AWS::CloudFormation::Stack"
}

// CloudFormationWaitCondition represents AWS::CloudFormation::WaitCondition
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waitcondition.html
type CloudFormationWaitCondition struct {
	// The number of success signals that AWS CloudFormation must receive
	// before it continues the stack creation process. When the wait
	// condition receives the requisite number of success signals, AWS
	// CloudFormation resumes the creation of the stack. If the wait
	// condition does not receive the specified number of success signals
	// before the Timeout period expires, AWS CloudFormation assumes that the
	// wait condition has failed and rolls the stack back.
	Count *StringExpr `json:"Count,omitempty"`

	// A reference to the wait condition handle used to signal this wait
	// condition. Use the Ref intrinsic function to specify an
	// AWS::CloudFormation::WaitConditionHandle resource.
	Handle *StringExpr `json:"Handle,omitempty"`

	// The length of time (in seconds) to wait for the number of signals that
	// the Count property specifies. Timeout is a minimum-bound property,
	// meaning the timeout occurs no sooner than the time you specify, but
	// can occur shortly thereafter. The maximum time that can be specified
	// for this property is 12 hours (43200 seconds).
	Timeout *StringExpr `json:"Timeout,omitempty"`
}

// ResourceType returns AWS::CloudFormation::WaitCondition to implement the ResourceProperties interface
func (s CloudFormationWaitCondition) ResourceType() string {
	return "AWS::CloudFormation::WaitCondition"
}

// CloudFormationWaitConditionHandle represents AWS::CloudFormation::WaitConditionHandle
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waitconditionhandle.html
type CloudFormationWaitConditionHandle struct {
}

// ResourceType returns AWS::CloudFormation::WaitConditionHandle to implement the ResourceProperties interface
func (s CloudFormationWaitConditionHandle) ResourceType() string {
	return "AWS::CloudFormation::WaitConditionHandle"
}

// CloudFrontDistribution represents AWS::CloudFront::Distribution
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution.html
type CloudFrontDistribution struct {
	// The distribution's configuration information.
	DistributionConfig *CloudFrontDistributionConfig `json:"DistributionConfig,omitempty"`
}

// ResourceType returns AWS::CloudFront::Distribution to implement the ResourceProperties interface
func (s CloudFrontDistribution) ResourceType() string {
	return "AWS::CloudFront::Distribution"
}

// CloudTrailTrail represents AWS::CloudTrail::Trail
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html
type CloudTrailTrail struct {
	// Indicates whether the trail is publishing events from global services,
	// such as IAM, to the log files.
	IncludeGlobalServiceEvents *BoolExpr `json:"IncludeGlobalServiceEvents,omitempty"`

	// Indicates whether the CloudTrail trail is currently logging AWS API
	// calls.
	IsLogging *BoolExpr `json:"IsLogging,omitempty"`

	// The name of the Amazon S3 bucket where CloudTrail publishes log files.
	S3BucketName *StringExpr `json:"S3BucketName,omitempty"`

	// An Amazon S3 object key prefix that precedes the name of all log
	// files.
	S3KeyPrefix *StringExpr `json:"S3KeyPrefix,omitempty"`

	// The name of an Amazon SNS topic that is notified when new log files
	// are published.
	SnsTopicName *StringExpr `json:"SnsTopicName,omitempty"`
}

// ResourceType returns AWS::CloudTrail::Trail to implement the ResourceProperties interface
func (s CloudTrailTrail) ResourceType() string {
	return "AWS::CloudTrail::Trail"
}

// CloudWatchAlarm represents AWS::CloudWatch::Alarm
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-alarm.html
type CloudWatchAlarm struct {
	// Indicates whether or not actions should be executed during any changes
	// to the alarm's state.
	ActionsEnabled *BoolExpr `json:"ActionsEnabled,omitempty"`

	// The list of actions to execute when this alarm transitions into an
	// ALARM state from any other state. Each action is specified as an
	// Amazon Resource Number (ARN). For more information about creating
	// alarms and the actions you can specify, see Creating Amazon CloudWatch
	// Alarms in the Amazon CloudWatch Developer Guide.
	AlarmActions *StringListExpr `json:"AlarmActions,omitempty"`

	// The description for the alarm.
	AlarmDescription *StringExpr `json:"AlarmDescription,omitempty"`

	// A name for the alarm. If you don't specify a name, AWS CloudFormation
	// generates a unique physical ID and uses that ID for the alarm name.
	// For more information, see Name Type.
	AlarmName *StringExpr `json:"AlarmName,omitempty"`

	// The arithmetic operation to use when comparing the specified Statistic
	// and Threshold. The specified Statistic value is used as the first
	// operand.
	ComparisonOperator *StringExpr `json:"ComparisonOperator,omitempty"`

	// The dimensions for the alarm's associated metric.
	Dimensions *CloudWatchMetricDimensionList `json:"Dimensions,omitempty"`

	// The number of periods over which data is compared to the specified
	// threshold.
	EvaluationPeriods *StringExpr `json:"EvaluationPeriods,omitempty"`

	// The list of actions to execute when this alarm transitions into an
	// INSUFFICIENT_DATA state from any other state. Each action is specified
	// as an Amazon Resource Number (ARN). Currently the only action
	// supported is publishing to an Amazon SNS topic or an Amazon Auto
	// Scaling policy.
	InsufficientDataActions *StringListExpr `json:"InsufficientDataActions,omitempty"`

	// The name for the alarm's associated metric. For more information about
	// the metrics that you can specify, see Amazon CloudWatch Namespaces,
	// Dimensions, and Metrics Reference in the Amazon CloudWatch Developer
	// Guide.
	MetricName *StringExpr `json:"MetricName,omitempty"`

	// The namespace for the alarm's associated metric.
	Namespace *StringExpr `json:"Namespace,omitempty"`

	// The list of actions to execute when this alarm transitions into an OK
	// state from any other state. Each action is specified as an Amazon
	// Resource Number (ARN). Currently the only action supported is
	// publishing to an Amazon SNS topic or an Amazon Auto Scaling policy.
	OKActions *StringListExpr `json:"OKActions,omitempty"`

	// The time over which the specified statistic is applied. You must
	// specify a time in seconds that is also a multiple of 60.
	Period *StringExpr `json:"Period,omitempty"`

	// The statistic to apply to the alarm's associated metric.
	Statistic *StringExpr `json:"Statistic,omitempty"`

	// The value against which the specified statistic is compared.
	Threshold *StringExpr `json:"Threshold,omitempty"`

	// The unit for the alarm's associated metric.
	Unit *StringExpr `json:"Unit,omitempty"`
}

// ResourceType returns AWS::CloudWatch::Alarm to implement the ResourceProperties interface
func (s CloudWatchAlarm) ResourceType() string {
	return "AWS::CloudWatch::Alarm"
}

// DataPipelinePipeline represents AWS::DataPipeline::Pipeline
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html
type DataPipelinePipeline struct {
	// Indicates whether to validate and start the pipeline or stop an active
	// pipeline. By default, the value is set to true.
	Activate *BoolExpr `json:"Activate,omitempty"`

	// A description for the pipeline.
	Description *StringExpr `json:"Description,omitempty"`

	// A name for the pipeline. Because AWS CloudFormation assigns each new
	// pipeline a unique identifier, you can use the same name for multiple
	// pipelines that are associated with your AWS account.
	Name *StringExpr `json:"Name,omitempty"`

	// Defines the variables that are in the pipeline definition. For more
	// information, see Creating a Pipeline Using Parameterized Templates in
	// the AWS Data Pipeline Developer Guide.
	ParameterObjects *DataPipelinePipelineParameterObjects `json:"ParameterObjects,omitempty"`

	// Defines the values for the parameters that are defined in the
	// ParameterObjects property. For more information, see Creating a
	// Pipeline Using Parameterized Templates in the AWS Data Pipeline
	// Developer Guide.
	ParameterValues *DataPipelinePipelineParameterValues `json:"ParameterValues,omitempty"`

	// A list of pipeline objects that make up the pipeline. For more
	// information about pipeline objects and a description of each object,
	// see Pipeline Object Reference in the AWS Data Pipeline Developer
	// Guide.
	PipelineObjects *DataPipelinePipelineObjectsList `json:"PipelineObjects,omitempty"`

	// A list of arbitrary tags (key-value pairs) to associate with the
	// pipeline, which you can use to control permissions. For more
	// information, see Controlling Access to Pipelines and Resources in the
	// AWS Data Pipeline Developer Guide.
	PipelineTags *DataPipelinePipelinePipelineTagsList `json:"PipelineTags,omitempty"`
}

// ResourceType returns AWS::DataPipeline::Pipeline to implement the ResourceProperties interface
func (s DataPipelinePipeline) ResourceType() string {
	return "AWS::DataPipeline::Pipeline"
}

// DynamoDBTable represents AWS::DynamoDB::Table
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html
type DynamoDBTable struct {
	// A list of AttributeName and AttributeType objects that describe the
	// key schema for the table and indexes.
	AttributeDefinitions *DynamoDBAttributeDefinitionsList `json:"AttributeDefinitions,omitempty"`

	// Global secondary indexes to be created on the table. You can create up
	// to 5 global secondary indexes.
	GlobalSecondaryIndexes *DynamoDBGlobalSecondaryIndexes `json:"GlobalSecondaryIndexes,omitempty"`

	// Specifies the attributes that make up the primary key for the table.
	// The attributes in the KeySchema property must also be defined in the
	// AttributeDefinitions property.
	KeySchema *DynamoDBKeySchema `json:"KeySchema,omitempty"`

	// Local secondary indexes to be created on the table. You can create up
	// to 5 local secondary indexes. Each index is scoped to a given hash key
	// value. The size of each hash key can be up to 10 gigabytes.
	LocalSecondaryIndexes *DynamoDBLocalSecondaryIndexes `json:"LocalSecondaryIndexes,omitempty"`

	// Throughput for the specified table, consisting of values for
	// ReadCapacityUnits and WriteCapacityUnits. For more information about
	// the contents of a Provisioned Throughput structure, see DynamoDB
	// Provisioned Throughput.
	ProvisionedThroughput *DynamoDBProvisionedThroughput `json:"ProvisionedThroughput,omitempty"`

	// A name for the table. If you don't specify a name, AWS CloudFormation
	// generates a unique physical ID and uses that ID for the table name.
	// For more information, see Name Type.
	TableName *StringExpr `json:"TableName,omitempty"`
}

// ResourceType returns AWS::DynamoDB::Table to implement the ResourceProperties interface
func (s DynamoDBTable) ResourceType() string {
	return "AWS::DynamoDB::Table"
}

// EC2CustomerGateway represents AWS::EC2::CustomerGateway
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customer-gateway.html
type EC2CustomerGateway struct {
	// The customer gateway's Border Gateway Protocol (BGP) Autonomous System
	// Number (ASN).
	BgpAsn *IntegerExpr `json:"BgpAsn,omitempty"`

	// The internet-routable IP address for the customer gateway's outside
	// interface. The address must be static.
	IpAddress *StringExpr `json:"IpAddress,omitempty"`

	// The tags that you want to attach to the resource.
	Tags []ResourceTag `json:"Tags,omitempty"`

	// The type of VPN connection that this customer gateway supports.
	Type *StringExpr `json:"Type,omitempty"`
}

// ResourceType returns AWS::EC2::CustomerGateway to implement the ResourceProperties interface
func (s EC2CustomerGateway) ResourceType() string {
	return "AWS::EC2::CustomerGateway"
}

// EC2DHCPOptions represents AWS::EC2::DHCPOptions
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-dhcp-options.html
type EC2DHCPOptions struct {
	// A domain name of your choice.
	DomainName *StringExpr `json:"DomainName,omitempty"`

	// The IP (IPv4) address of a domain name server. You can specify up to
	// four addresses.
	DomainNameServers *StringListExpr `json:"DomainNameServers,omitempty"`

	// The IP address (IPv4) of a NetBIOS name server. You can specify up to
	// four addresses.
	NetbiosNameServers *StringListExpr `json:"NetbiosNameServers,omitempty"`

	// An integer value indicating the NetBIOS node type:
	NetbiosNodeType interface{} `json:"NetbiosNodeType,omitempty"`

	// The IP address (IPv4) of a Network Time Protocol (NTP) server. You can
	// specify up to four addresses.
	NtpServers *StringListExpr `json:"NtpServers,omitempty"`

	// An arbitrary set of tags (key–value pairs) for this resource.
	Tags []ResourceTag `json:"Tags,omitempty"`
}

// ResourceType returns AWS::EC2::DHCPOptions to implement the ResourceProperties interface
func (s EC2DHCPOptions) ResourceType() string {
	return "AWS::EC2::DHCPOptions"
}

// EC2EIP represents AWS::EC2::EIP
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html
type EC2EIP struct {
	// The Instance ID of the Amazon EC2 instance that you want to associate
	// with this Elastic IP address.
	InstanceId *StringExpr `json:"InstanceId,omitempty"`

	// Set to vpc to allocate the address to your Virtual Private Cloud
	// (VPC). No other values are supported.
	Domain *StringExpr `json:"Domain,omitempty"`
}

// ResourceType returns AWS::EC2::EIP to implement the ResourceProperties interface
func (s EC2EIP) ResourceType() string {
	return "AWS::EC2::EIP"
}

// EC2EIPAssociation represents AWS::EC2::EIPAssociation
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip-association.html
type EC2EIPAssociation struct {
	// Allocation ID for the VPC Elastic IP address you want to associate
	// with an Amazon EC2 instance in your VPC.
	AllocationId *StringExpr `json:"AllocationId,omitempty"`

	// Elastic IP address that you want to associate with the Amazon EC2
	// instance specified by the InstanceId property. You can specify an
	// existing Elastic IP address or a reference to an Elastic IP address
	// allocated with a AWS::EC2::EIP resource.
	EIP *StringExpr `json:"EIP,omitempty"`

	// Instance ID of the Amazon EC2 instance that you want to associate with
	// the Elastic IP address specified by the EIP property.
	InstanceId *StringExpr `json:"InstanceId,omitempty"`

	// The ID of the network interface to associate with the Elastic IP
	// address (VPC only).
	NetworkInterfaceId *StringExpr `json:"NetworkInterfaceId,omitempty"`

	// The private IP address that you want to associate with the Elastic IP
	// address. The private IP address is restricted to the primary and
	// secondary private IP addresses that are associated with the network
	// interface. By default, the private IP address that is associated with
	// the EIP is the primary private IP address of the network interface.
	PrivateIpAddress *StringExpr `json:"PrivateIpAddress,omitempty"`
}

// ResourceType returns AWS::EC2::EIPAssociation to implement the ResourceProperties interface
func (s EC2EIPAssociation) ResourceType() string {
	return "AWS::EC2::EIPAssociation"
}

// EC2Instance represents AWS::EC2::Instance
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html
type EC2Instance struct {
	// Specifies the name of the Availability Zone in which the instance is
	// located.
	AvailabilityZone *StringExpr `json:"AvailabilityZone,omitempty"`

	// Defines a set of Amazon Elastic Block Store block device mappings,
	// ephemeral instance store block device mappings, or both. For more
	// information, see Amazon Elastic Block Store or Amazon EC2 Instance
	// Store in the Amazon EC2 User Guide for Linux Instances.
	BlockDeviceMappings *EC2BlockDeviceMappingPropertyList `json:"BlockDeviceMappings,omitempty"`

	// Specifies whether the instance can be terminated through the API.
	DisableApiTermination *BoolExpr `json:"DisableApiTermination,omitempty"`

	// Specifies whether the instance is optimized for Amazon Elastic Block
	// Store I/O. This optimization provides dedicated throughput to Amazon
	// EBS and an optimized configuration stack to provide optimal EBS I/O
	// performance.
	EbsOptimized *BoolExpr `json:"EbsOptimized,omitempty"`

	// The physical ID of an instance profile or a reference to an
	// AWS::IAM::InstanceProfile resource.
	IamInstanceProfile *StringExpr `json:"IamInstanceProfile,omitempty"`

	// Provides the unique ID of the Amazon Machine Image (AMI) that was
	// assigned during registration.
	ImageId *StringExpr `json:"ImageId,omitempty"`

	// Indicates whether an instance stops or terminates when you shut down
	// the instance from the instance's operating system shutdown command.
	// You can specify stop or terminate. For more information, see the
	// RunInstances command in the Amazon EC2 API Reference.
	InstanceInitiatedShutdownBehavior *StringExpr `json:"InstanceInitiatedShutdownBehavior,omitempty"`

	// The instance type, such as t2.micro. The default type is "m1.small".
	// For a list of instance types, see Instance Families and Types.
	InstanceType *StringExpr `json:"InstanceType,omitempty"`

	// The kernel ID.
	KernelId *StringExpr `json:"KernelId,omitempty"`

	// Provides the name of the Amazon EC2 key pair.
	KeyName *StringExpr `json:"KeyName,omitempty"`

	// Specifies whether monitoring is enabled for the instance.
	Monitoring *BoolExpr `json:"Monitoring,omitempty"`

	// A list of embedded objects that describe the network interfaces to
	// associate with this instance.
	NetworkInterfaces *EC2NetworkInterfaceEmbeddedList `json:"NetworkInterfaces,omitempty"`

	// The name of an existing placement group that you want to launch the
	// instance into (for cluster instances).
	PlacementGroupName *StringExpr `json:"PlacementGroupName,omitempty"`

	// The private IP address for this instance.
	PrivateIpAddress *StringExpr `json:"PrivateIpAddress,omitempty"`

	// The ID of the RAM disk to select. Some kernels require additional
	// drivers at launch. Check the kernel requirements for information about
	// whether you need to specify a RAM disk. To find kernel requirements,
	// go to the AWS Resource Center and search for the kernel ID.
	RamdiskId *StringExpr `json:"RamdiskId,omitempty"`

	// A list that contains the security group IDs for VPC security groups to
	// assign to the Amazon EC2 instance. If you specified the
	// NetworkInterfaces property, do not specify this property.
	SecurityGroupIds *StringListExpr `json:"SecurityGroupIds,omitempty"`

	// Valid only for Amazon EC2 security groups. A list that contains the
	// Amazon EC2 security groups to assign to the Amazon EC2 instance. The
	// list can contain both the name of existing Amazon EC2 security groups
	// or references to AWS::EC2::SecurityGroup resources created in the
	// template.
	SecurityGroups *StringListExpr `json:"SecurityGroups,omitempty"`

	// Controls whether source/destination checking is enabled on the
	// instance. Also determines if an instance in a VPC will perform network
	// address translation (NAT).
	SourceDestCheck *BoolExpr `json:"SourceDestCheck,omitempty"`

	// If you're using Amazon VPC, this property specifies the ID of the
	// subnet that you want to launch the instance into. If you specified the
	// NetworkInterfaces property, do not specify this property.
	SubnetId *StringExpr `json:"SubnetId,omitempty"`

	// An arbitrary set of tags (key–value pairs) for this instance.
	Tags []ResourceTag `json:"Tags,omitempty"`

	// The tenancy of the instance that you want to launch. This value can be
	// either "default" or "dedicated". An instance that has a tenancy value
	// of "dedicated" runs on single-tenant hardware and can be launched only
	// into a VPC. For more information, see Using EC2 Dedicated Instances
	// Within Your VPC in the Amazon VPC User Guide.
	Tenancy *StringExpr `json:"Tenancy,omitempty"`

	// Base64-encoded MIME user data that is made available to the instances.
	UserData *StringExpr `json:"UserData,omitempty"`

	// The Amazon EBS volumes to attach to the instance.
	Volumes *EC2MountPointList `json:"Volumes,omitempty"`
}

// ResourceType returns AWS::EC2::Instance to implement the ResourceProperties interface
func (s EC2Instance) ResourceType() string {
	return "AWS::EC2::Instance"
}

// EC2InternetGateway represents AWS::EC2::InternetGateway
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-internet-gateway.html
type EC2InternetGateway struct {
	// An arbitrary set of tags (key–value pairs) for this resource.
	Tags []ResourceTag `json:"Tags,omitempty"`
}

// ResourceType returns AWS::EC2::InternetGateway to implement the ResourceProperties interface
func (s EC2InternetGateway) ResourceType() string {
	return "AWS::EC2::InternetGateway"
}

// EC2NetworkAcl represents AWS::EC2::NetworkAcl
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl.html
type EC2NetworkAcl struct {
	// An arbitrary set of tags (key–value pairs) for this ACL.
	Tags []ResourceTag `json:"Tags,omitempty"`

	// The ID of the VPC where the network ACL will be created.
	VpcId *StringExpr `json:"VpcId,omitempty"`
}

// ResourceType returns AWS::EC2::NetworkAcl to implement the ResourceProperties interface
func (s EC2NetworkAcl) ResourceType() string {
	return "AWS::EC2::NetworkAcl"
}

// EC2NetworkAclEntry represents AWS::EC2::NetworkAclEntry
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html
type EC2NetworkAclEntry struct {
	// The CIDR range to allow or deny, in CIDR notation (e.g.,
	// 172.16.0.0/24).
	CidrBlock *StringExpr `json:"CidrBlock,omitempty"`

	// Whether this rule applies to egress traffic from the subnet ("true")
	// or ingress traffic to the subnet ("false").
	Egress *BoolExpr `json:"Egress,omitempty"`

	// The Internet Control Message Protocol (ICMP) code and type.
	Icmp *EC2ICMP `json:"Icmp,omitempty"`

	// ID of the ACL where the entry will be created.
	NetworkAclId *StringExpr `json:"NetworkAclId,omitempty"`

	// The range of port numbers for the UDP/TCP protocol.
	PortRange *EC2PortRange `json:"PortRange,omitempty"`

	// The IP protocol that the rule applies to. You must specify -1 or a
	// protocol number (go to Protocol Numbers at iana.org). You can specify
	// -1 for all protocols.
	Protocol *IntegerExpr `json:"Protocol,omitempty"`

	// Whether to allow or deny traffic that matches the rule; valid values
	// are "allow" or "deny".
	RuleAction *StringExpr `json:"RuleAction,omitempty"`

	// Rule number to assign to the entry (e.g., 100). This must be a
	// positive integer from 1 to 32766.
	RuleNumber *IntegerExpr `json:"RuleNumber,omitempty"`
}

// ResourceType returns AWS::EC2::NetworkAclEntry to implement the ResourceProperties interface
func (s EC2NetworkAclEntry) ResourceType() string {
	return "AWS::EC2::NetworkAclEntry"
}

// EC2NetworkInterface represents AWS::EC2::NetworkInterface
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html
type EC2NetworkInterface struct {
	// The description of this network interface.
	Description *StringExpr `json:"Description,omitempty"`

	// A list of security group IDs associated with this network interface.
	GroupSet *StringListExpr `json:"GroupSet,omitempty"`

	// Assigns a single private IP address to the network interface, which is
	// used as the primary private IP address. If you want to specify
	// multiple private IP address, use the PrivateIpAddresses property.
	PrivateIpAddress *StringExpr `json:"PrivateIpAddress,omitempty"`

	// Assigns a list of private IP addresses to the network interface. You
	// can specify a primary private IP address by setting the value of the
	// Primary property to true in the PrivateIpAddressSpecification
	// property. If you want Amazon EC2 to automatically assign private IP
	// addresses, use the SecondaryPrivateIpAddressCount property and do not
	// specify this property.
	PrivateIpAddresses *EC2NetworkInterfacePrivateIPSpecificationList `json:"PrivateIpAddresses,omitempty"`

	// The number of secondary private IP addresses that Amazon EC2
	// automatically assigns to the network interface. Amazon EC2 uses the
	// value of the PrivateIpAddress property as the primary private IP
	// address. If you don't specify that property, Amazon EC2 automatically
	// assigns both the primary and secondary private IP addresses.
	SecondaryPrivateIpAddressCount *IntegerExpr `json:"SecondaryPrivateIpAddressCount,omitempty"`

	// Flag indicating whether traffic to or from the instance is validated.
	SourceDestCheck *BoolExpr `json:"SourceDestCheck,omitempty"`

	// The ID of the subnet to associate with the network interface.
	SubnetId *StringExpr `json:"SubnetId,omitempty"`

	// An arbitrary set of tags (key–value pairs) for this network
	// interface.
	Tags []ResourceTag `json:"Tags,omitempty"`
}

// ResourceType returns AWS::EC2::NetworkInterface to implement the ResourceProperties interface
func (s EC2NetworkInterface) ResourceType() string {
	return "AWS::EC2::NetworkInterface"
}

// EC2NetworkInterfaceAttachment represents AWS::EC2::NetworkInterfaceAttachment
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface-attachment.html
type EC2NetworkInterfaceAttachment struct {
	// Whether to delete the network interface when the instance terminates.
	// By default, this value is set to True.
	DeleteOnTermination *BoolExpr `json:"DeleteOnTermination,omitempty"`

	// The network interface's position in the attachment order. For example,
	// the first attached network interface has a DeviceIndex of 0.
	DeviceIndex *StringExpr `json:"DeviceIndex,omitempty"`

	// The ID of the instance to which you will attach the ENI.
	InstanceId *StringExpr `json:"InstanceId,omitempty"`

	// The ID of the ENI that you want to attach.
	NetworkInterfaceId *StringExpr `json:"NetworkInterfaceId,omitempty"`
}

// ResourceType returns AWS::EC2::NetworkInterfaceAttachment to implement the ResourceProperties interface
func (s EC2NetworkInterfaceAttachment) ResourceType() string {
	return "AWS::EC2::NetworkInterfaceAttachment"
}

// EC2Route represents AWS::EC2::Route
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html
type EC2Route struct {
	// The CIDR address block used for the destination match. For example,
	// "0.0.0.0/0". Routing decisions are based on the most specific match.
	DestinationCidrBlock *StringExpr `json:"DestinationCidrBlock,omitempty"`

	// The ID of an Internet gateway or virtual private gateway that is
	// attached to your VPC. For example: "igw-eaad4883".
	GatewayId *StringExpr `json:"GatewayId,omitempty"`

	// The ID of a NAT instance in your VPC. For example, "i-1a2b3c4d".
	InstanceId *StringExpr `json:"InstanceId,omitempty"`

	// Allows the routing of network interface IDs.
	NetworkInterfaceId *StringExpr `json:"NetworkInterfaceId,omitempty"`

	// The ID of the route table where the route will be added.
	RouteTableId *StringExpr `json:"RouteTableId,omitempty"`

	// The ID of a VPC peering connection.
	VpcPeeringConnectionId *StringExpr `json:"VpcPeeringConnectionId,omitempty"`
}

// ResourceType returns AWS::EC2::Route to implement the ResourceProperties interface
func (s EC2Route) ResourceType() string {
	return "AWS::EC2::Route"
}

// EC2RouteTable represents AWS::EC2::RouteTable
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route-table.html
type EC2RouteTable struct {
	// The ID of the VPC where the route table will be created.
	VpcId *StringExpr `json:"VpcId,omitempty"`

	// An arbitrary set of tags (key–value pairs) for this route table.
	Tags []ResourceTag `json:"Tags,omitempty"`
}

// ResourceType returns AWS::EC2::RouteTable to implement the ResourceProperties interface
func (s EC2RouteTable) ResourceType() string {
	return "AWS::EC2::RouteTable"
}

// EC2SecurityGroup represents AWS::EC2::SecurityGroup
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html
type EC2SecurityGroup struct {
	// Description of the security group.
	GroupDescription *StringExpr `json:"GroupDescription,omitempty"`

	// A list of Amazon EC2 security group egress rules.
	SecurityGroupEgress *EC2SecurityGroupRuleList `json:"SecurityGroupEgress,omitempty"`

	// A list of Amazon EC2 security group ingress rules.
	SecurityGroupIngress *EC2SecurityGroupRuleList `json:"SecurityGroupIngress,omitempty"`

	// The tags that you want to attach to the resource.
	Tags []ResourceTag `json:"Tags,omitempty"`

	// The physical ID of the VPC. Can be obtained by using a reference to an
	// AWS::EC2::VPC, such as: { "Ref" : "myVPC" }.
	VpcId *StringExpr `json:"VpcId,omitempty"`
}

// ResourceType returns AWS::EC2::SecurityGroup to implement the ResourceProperties interface
func (s EC2SecurityGroup) ResourceType() string {
	return "AWS::EC2::SecurityGroup"
}

// EC2SecurityGroupEgress represents AWS::EC2::SecurityGroupEgress
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html
type EC2SecurityGroupEgress struct {
	// CIDR range.
	CidrIp *StringExpr `json:"CidrIp,omitempty"`

	// Specifies the group ID of the destination Amazon VPC security group.
	DestinationSecurityGroupId *StringExpr `json:"DestinationSecurityGroupId,omitempty"`

	// Start of port range for the TCP and UDP protocols, or an ICMP type
	// number. If you specify icmp for the IpProtocol property, you can
	// specify -1 as a wildcard (i.e., any ICMP type number).
	FromPort *IntegerExpr `json:"FromPort,omitempty"`

	// ID of the Amazon VPC security group to modify. This value can be a
	// reference to an AWS::EC2::SecurityGroup resource that has a valid
	// VpcId property or the ID of an existing Amazon VPC security group.
	GroupId *StringExpr `json:"GroupId,omitempty"`

	// IP protocol name or number. For valid values, see the IpProtocol
	// parameter in AuthorizeSecurityGroupIngress
	IpProtocol *StringExpr `json:"IpProtocol,omitempty"`

	// End of port range for the TCP and UDP protocols, or an ICMP code. If
	// you specify icmp for the IpProtocol property, you can specify -1 as a
	// wildcard (i.e., any ICMP code).
	ToPort *IntegerExpr `json:"ToPort,omitempty"`
}

// ResourceType returns AWS::EC2::SecurityGroupEgress to implement the ResourceProperties interface
func (s EC2SecurityGroupEgress) ResourceType() string {
	return "AWS::EC2::SecurityGroupEgress"
}

// EC2SecurityGroupIngress represents AWS::EC2::SecurityGroupIngress
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html
type EC2SecurityGroupIngress struct {
	// Specifies a CIDR range.
	CidrIp *StringExpr `json:"CidrIp,omitempty"`

	// Start of port range for the TCP and UDP protocols, or an ICMP type
	// number. If you specify icmp for the IpProtocol property, you can
	// specify -1 as a wildcard (i.e., any ICMP type number).
	FromPort *IntegerExpr `json:"FromPort,omitempty"`

	// ID of the Amazon EC2 or VPC security group to modify. The group must
	// belong to your account.
	GroupId *StringExpr `json:"GroupId,omitempty"`

	// Name of the Amazon EC2 security group (non-VPC security group) to
	// modify. This value can be a reference to an AWS::EC2::SecurityGroup
	// resource or the name of an existing Amazon EC2 security group.
	GroupName *StringExpr `json:"GroupName,omitempty"`

	// IP protocol name or number. For valid values, see the IpProtocol
	// parameter in AuthorizeSecurityGroupIngress
	IpProtocol *StringExpr `json:"IpProtocol,omitempty"`

	// Specifies the ID of the source security group or uses the Ref
	// intrinsic function to refer to the logical ID of a security group
	// defined in the same template.
	SourceSecurityGroupId *StringExpr `json:"SourceSecurityGroupId,omitempty"`

	// Specifies the name of the Amazon EC2 security group (non-VPC security
	// group) to allow access or uses the Ref intrinsic function to refer to
	// the logical name of a security group defined in the same template. For
	// instances in a VPC, specify the SourceSecurityGroupId property.
	SourceSecurityGroupName *StringExpr `json:"SourceSecurityGroupName,omitempty"`

	// Specifies the AWS Account ID of the owner of the Amazon EC2 security
	// group specified in the SourceSecurityGroupName property.
	SourceSecurityGroupOwnerId *StringExpr `json:"SourceSecurityGroupOwnerId,omitempty"`

	// End of port range for the TCP and UDP protocols, or an ICMP code. If
	// you specify icmp for the IpProtocol property, you can specify -1 as a
	// wildcard (i.e., any ICMP code).
	ToPort *IntegerExpr `json:"ToPort,omitempty"`
}

// ResourceType returns AWS::EC2::SecurityGroupIngress to implement the ResourceProperties interface
func (s EC2SecurityGroupIngress) ResourceType() string {
	return "AWS::EC2::SecurityGroupIngress"
}

// EC2Subnet represents AWS::EC2::Subnet
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html
type EC2Subnet struct {
	// The availability zone in which you want the subnet. Default: AWS
	// selects a zone for you (recommended).
	AvailabilityZone *StringExpr `json:"AvailabilityZone,omitempty"`

	// The CIDR block that you want the subnet to cover (for example,
	// "10.0.0.0/24").
	CidrBlock *StringExpr `json:"CidrBlock,omitempty"`

	// Indicates whether instances that are launched in this subnet receive a
	// public IP address. By default, the value is false.
	MapPublicIpOnLaunch *BoolExpr `json:"MapPublicIpOnLaunch,omitempty"`

	// An arbitrary set of tags (key–value pairs) for this subnet.
	Tags []ResourceTag `json:"Tags,omitempty"`

	// A Ref structure that contains the ID of the VPC on which you want to
	// create the subnet. The VPC ID is provided as the value of the "Ref"
	// property, as: { "Ref": "VPCID" }.
	VpcId interface{} `json:"VpcId,omitempty"`
}

// ResourceType returns AWS::EC2::Subnet to implement the ResourceProperties interface
func (s EC2Subnet) ResourceType() string {
	return "AWS::EC2::Subnet"
}

// EC2SubnetNetworkAclAssociation represents AWS::EC2::SubnetNetworkAclAssociation
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-network-acl-assoc.html
type EC2SubnetNetworkAclAssociation struct {
	// The ID representing the current association between the original
	// network ACL and the subnet.
	SubnetId *StringExpr `json:"SubnetId,omitempty"`

	// The ID of the new ACL to associate with the subnet.
	NetworkAclId *StringExpr `json:"NetworkAclId,omitempty"`
}

// ResourceType returns AWS::EC2::SubnetNetworkAclAssociation to implement the ResourceProperties interface
func (s EC2SubnetNetworkAclAssociation) ResourceType() string {
	return "AWS::EC2::SubnetNetworkAclAssociation"
}

// EC2SubnetRouteTableAssociation represents AWS::EC2::SubnetRouteTableAssociation
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-route-table-assoc.html
type EC2SubnetRouteTableAssociation struct {
	// The ID of the route table. This is commonly written as a reference to
	// a route table declared elsewhere in the template. For example:
	RouteTableId *StringExpr `json:"RouteTableId,omitempty"`

	// The ID of the subnet. This is commonly written as a reference to a
	// subnet declared elsewhere in the template. For example:
	SubnetId *StringExpr `json:"SubnetId,omitempty"`
}

// ResourceType returns AWS::EC2::SubnetRouteTableAssociation to implement the ResourceProperties interface
func (s EC2SubnetRouteTableAssociation) ResourceType() string {
	return "AWS::EC2::SubnetRouteTableAssociation"
}

// EC2Volume represents AWS::EC2::Volume
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volume.html
type EC2Volume struct {
	// The Availability Zone in which to create the new volume.
	AvailabilityZone *StringExpr `json:"AvailabilityZone,omitempty"`

	// Indicates whether the volume is encrypted. Encrypted Amazon EBS
	// volumes can only be attached to instance types that support Amazon EBS
	// encryption. Volumes that are created from encrypted snapshots are
	// automatically encrypted. You cannot create an encrypted volume from an
	// unencrypted snapshot or vice versa. If your AMI uses encrypted
	// volumes, you can only launch the AMI on supported instance types. For
	// more information, see Amazon EBS encryption in the Amazon EC2 User
	// Guide for Linux Instances.
	Encrypted *BoolExpr `json:"Encrypted,omitempty"`

	// The number of I/O operations per second (IOPS) that the volume
	// supports. This can be any integer value from 1–4000.
	Iops *IntegerExpr `json:"Iops,omitempty"`

	// The Amazon Resource Name (ARN) of the AWS Key Management Service
	// master key that is used to create the encrypted volume, such as
	// arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef.
	// If you create an encrypted volume and don't specify this property, the
	// default master key is used.
	KmsKeyId *StringExpr `json:"KmsKeyId,omitempty"`

	// The size of the volume, in gibibytes (GiBs). This can be any value
	// from 10–1024.
	Size *StringExpr `json:"Size,omitempty"`

	// The snapshot from which to create the new volume.
	SnapshotId *StringExpr `json:"SnapshotId,omitempty"`

	// An arbitrary set of tags (key–value pairs) for this volume.
	Tags []ResourceTag `json:"Tags,omitempty"`

	// The volume type. You can specify standard, io1, or gp2. If you set the
	// type to io1, you must also set the Iops property. For more information
	// about these values and the default value, see CreateVolume in the
	// Amazon EC2 API Reference.
	VolumeType *StringExpr `json:"VolumeType,omitempty"`
}

// ResourceType returns AWS::EC2::Volume to implement the ResourceProperties interface
func (s EC2Volume) ResourceType() string {
	return "AWS::EC2::Volume"
}

// EC2VolumeAttachment represents AWS::EC2::VolumeAttachment
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volumeattachment.html
type EC2VolumeAttachment struct {
	// How the device is exposed to the instance (e.g., /dev/sdh, or xvdh).
	Device *StringExpr `json:"Device,omitempty"`

	// The ID of the instance to which the volume attaches. This value can be
	// a reference to an AWS::EC2::Instance resource, or it can be the
	// physical ID of an existing EC2 instance.
	InstanceId *StringExpr `json:"InstanceId,omitempty"`

	// The ID of the Amazon EBS volume. The volume and instance must be
	// within the same Availability Zone. This value can be a reference to an
	// AWS::EC2::Volume resource, or it can be the volume ID of an existing
	// Amazon EBS volume.
	VolumeId *StringExpr `json:"VolumeId,omitempty"`
}

// ResourceType returns AWS::EC2::VolumeAttachment to implement the ResourceProperties interface
func (s EC2VolumeAttachment) ResourceType() string {
	return "AWS::EC2::VolumeAttachment"
}

// EC2VPC represents AWS::EC2::VPC
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc.html
type EC2VPC struct {
	// The CIDR block you want the VPC to cover. For example: "10.0.0.0/16".
	CidrBlock *StringExpr `json:"CidrBlock,omitempty"`

	// Specifies whether DNS resolution is supported for the VPC. If this
	// attribute is true, the Amazon DNS server resolves DNS hostnames for
	// your instances to their corresponding IP addresses; otherwise, it does
	// not. By default the value is set to true.
	EnableDnsSupport *BoolExpr `json:"EnableDnsSupport,omitempty"`

	// Specifies whether the instances launched in the VPC get DNS hostnames.
	// If this attribute is true, instances in the VPC get DNS hostnames;
	// otherwise, they do not. You can only set EnableDnsHostnames to true if
	// you also set the EnableDnsSupport attribute to true. By default, the
	// value is set to false.
	EnableDnsHostnames *BoolExpr `json:"EnableDnsHostnames,omitempty"`

	// The allowed tenancy of instances launched into the VPC.
	InstanceTenancy *StringExpr `json:"InstanceTenancy,omitempty"`

	// An arbitrary set of tags (key–value pairs) for this VPC.
	Tags []ResourceTag `json:"Tags,omitempty"`
}

// ResourceType returns AWS::EC2::VPC to implement the ResourceProperties interface
func (s EC2VPC) ResourceType() string {
	return "AWS::EC2::VPC"
}

// EC2VPCDHCPOptionsAssociation represents AWS::EC2::VPCDHCPOptionsAssociation
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-dhcp-options-assoc.html
type EC2VPCDHCPOptionsAssociation struct {
	// The ID of the DHCP options you want to associate with the VPC. Specify
	// default if you want the VPC to use no DHCP options.
	DhcpOptionsId *StringExpr `json:"DhcpOptionsId,omitempty"`

	// The ID of the VPC to associate with this DHCP options set.
	VpcId *StringExpr `json:"VpcId,omitempty"`
}

// ResourceType returns AWS::EC2::VPCDHCPOptionsAssociation to implement the ResourceProperties interface
func (s EC2VPCDHCPOptionsAssociation) ResourceType() string {
	return "AWS::EC2::VPCDHCPOptionsAssociation"
}

// EC2VPCGatewayAttachment represents AWS::EC2::VPCGatewayAttachment
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-gateway-attachment.html
type EC2VPCGatewayAttachment struct {
	// The ID of the Internet gateway.
	InternetGatewayId *StringExpr `json:"InternetGatewayId,omitempty"`

	// The ID of the VPC to associate with this gateway.
	VpcId *StringExpr `json:"VpcId,omitempty"`

	// The ID of the virtual private network (VPN) gateway to attach to the
	// VPC.
	VpnGatewayId *StringExpr `json:"VpnGatewayId,omitempty"`
}

// ResourceType returns AWS::EC2::VPCGatewayAttachment to implement the ResourceProperties interface
func (s EC2VPCGatewayAttachment) ResourceType() string {
	return "AWS::EC2::VPCGatewayAttachment"
}

// EC2VPCPeeringConnection represents AWS::EC2::VPCPeeringConnection
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html
type EC2VPCPeeringConnection struct {
	// The ID of the VPC with which you are creating the peering connection.
	PeerVpcId *StringExpr `json:"PeerVpcId,omitempty"`

	// An arbitrary set of tags (key–value pairs) for this resource.
	Tags []ResourceTag `json:"Tags,omitempty"`

	// The ID of the VPC that is requesting a peering connection.
	VpcId *StringExpr `json:"VpcId,omitempty"`
}

// ResourceType returns AWS::EC2::VPCPeeringConnection to implement the ResourceProperties interface
func (s EC2VPCPeeringConnection) ResourceType() string {
	return "AWS::EC2::VPCPeeringConnection"
}

// EC2VPNConnection represents AWS::EC2::VPNConnection
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection.html
type EC2VPNConnection struct {
	// The type of VPN connection this virtual private gateway supports.
	Type *StringExpr `json:"Type,omitempty"`

	// The ID of the customer gateway. This can either be an embedded JSON
	// object or a reference to a Gateway ID.
	CustomerGatewayId *StringExpr `json:"CustomerGatewayId,omitempty"`

	// Indicates whether the VPN connection requires static routes.
	StaticRoutesOnly *BoolExpr `json:"StaticRoutesOnly,omitempty"`

	// The tags that you want to attach to the resource.
	Tags []ResourceTag `json:"Tags,omitempty"`

	// The ID of the virtual private gateway. This can either be an embedded
	// JSON object or a reference to a Gateway ID.
	VpnGatewayId *StringExpr `json:"VpnGatewayId,omitempty"`
}

// ResourceType returns AWS::EC2::VPNConnection to implement the ResourceProperties interface
func (s EC2VPNConnection) ResourceType() string {
	return "AWS::EC2::VPNConnection"
}

// EC2VPNConnectionRoute represents AWS::EC2::VPNConnectionRoute
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection-route.html
type EC2VPNConnectionRoute struct {
	// The CIDR block that is associated with the local subnet of the
	// customer network.
	DestinationCidrBlock *StringExpr `json:"DestinationCidrBlock,omitempty"`

	// The ID of the VPN connection.
	VpnConnectionId *StringExpr `json:"VpnConnectionId,omitempty"`
}

// ResourceType returns AWS::EC2::VPNConnectionRoute to implement the ResourceProperties interface
func (s EC2VPNConnectionRoute) ResourceType() string {
	return "AWS::EC2::VPNConnectionRoute"
}

// EC2VPNGateway represents AWS::EC2::VPNGateway
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-gateway.html
type EC2VPNGateway struct {
	// The type of VPN connection this virtual private gateway supports. The
	// only valid value is "ipsec.1".
	Type *StringExpr `json:"Type,omitempty"`

	// An arbitrary set of tags (key–value pairs) for this resource.
	Tags []ResourceTag `json:"Tags,omitempty"`
}

// ResourceType returns AWS::EC2::VPNGateway to implement the ResourceProperties interface
func (s EC2VPNGateway) ResourceType() string {
	return "AWS::EC2::VPNGateway"
}

// EC2VPNGatewayRoutePropagation represents AWS::EC2::VPNGatewayRoutePropagation
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-gatewayrouteprop.html
type EC2VPNGatewayRoutePropagation struct {
	// A list of routing table IDs that are associated with a VPC. The
	// routing tables must be associated with the same VPC that the virtual
	// private gateway is attached to.
	RouteTableIds interface{} `json:"RouteTableIds,omitempty"`

	// The ID of the virtual private gateway that is attached to a VPC. The
	// virtual private gateway must be attached to the same VPC that the
	// routing tables are associated with.
	VpnGatewayId *StringExpr `json:"VpnGatewayId,omitempty"`
}

// ResourceType returns AWS::EC2::VPNGatewayRoutePropagation to implement the ResourceProperties interface
func (s EC2VPNGatewayRoutePropagation) ResourceType() string {
	return "AWS::EC2::VPNGatewayRoutePropagation"
}

// ECSCluster represents AWS::ECS::Cluster
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-cluster.html
type ECSCluster struct {
}

// ResourceType returns AWS::ECS::Cluster to implement the ResourceProperties interface
func (s ECSCluster) ResourceType() string {
	return "AWS::ECS::Cluster"
}

// ECSService represents AWS::ECS::Service
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html
type ECSService struct {
	// The name or Amazon Resource Name (ARN) of the cluster that you want to
	// run your service on. If you do not specify a cluster, Amazon ECS uses
	// the default cluster.
	Cluster *StringExpr `json:"Cluster,omitempty"`

	// The number of simultaneous tasks, which you specify by using the
	// TaskDefinition property, that you want to run on the cluster.
	DesiredCount *StringExpr `json:"DesiredCount,omitempty"`

	// A list of load balancer objects to associate with the cluster.
	LoadBalancers *EC2ContainerServiceServiceLoadBalancersList `json:"LoadBalancers,omitempty"`

	// The name or ARN of an AWS Identity and Access Management (IAM) role
	// that allows your Amazon ECS container agent to make calls to your load
	// balancer.
	Role *StringExpr `json:"Role,omitempty"`

	// The family, family and revision (family:revision), or ARN of the task
	// definition that you want to run on the cluster.
	TaskDefinition *StringExpr `json:"TaskDefinition,omitempty"`
}

// ResourceType returns AWS::ECS::Service to implement the ResourceProperties interface
func (s ECSService) ResourceType() string {
	return "AWS::ECS::Service"
}

// ECSTaskDefinition represents AWS::ECS::TaskDefinition
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html
type ECSTaskDefinition struct {
	// A list of container definitions in JSON format that describe the
	// containers that make up your task.
	ContainerDefinitions *EC2ContainerServiceTaskDefinitionContainerDefinitionsList `json:"ContainerDefinitions,omitempty"`

	// A list of volume definitions in JSON format for volumes that you can
	// use in your container definitions.
	Volumes *EC2ContainerServiceTaskDefinitionVolumesList `json:"Volumes,omitempty"`
}

// ResourceType returns AWS::ECS::TaskDefinition to implement the ResourceProperties interface
func (s ECSTaskDefinition) ResourceType() string {
	return "AWS::ECS::TaskDefinition"
}

// ElastiCacheCacheCluster represents AWS::ElastiCache::CacheCluster
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cache-cluster.html
type ElastiCacheCacheCluster struct {
	// Indicates that minor engine upgrades will be applied automatically to
	// the cache cluster during the maintenance window.
	AutoMinorVersionUpgrade *BoolExpr `json:"AutoMinorVersionUpgrade,omitempty"`

	// For Memcached cache clusters, indicates whether the nodes are created
	// in a single Availability Zone or across multiple Availability Zones in
	// the cluster's region.
	AZMode *StringExpr `json:"AZMode,omitempty"`

	// The compute and memory capacity of nodes in a cache cluster.
	CacheNodeType *StringExpr `json:"CacheNodeType,omitempty"`

	// The name of the cache parameter group that is associated with this
	// cache cluster.
	CacheParameterGroupName *StringExpr `json:"CacheParameterGroupName,omitempty"`

	// A list of cache security group names that are associated with this
	// cache cluster. If your cache cluster is in a VPC, specify the
	// VpcSecurityGroupIds property instead.
	CacheSecurityGroupNames *StringListExpr `json:"CacheSecurityGroupNames,omitempty"`

	// The cache subnet group that you associate with a cache cluster.
	CacheSubnetGroupName *StringExpr `json:"CacheSubnetGroupName,omitempty"`

	// A name for the cache cluster. If you don't specify a name, AWS
	// CloudFormation generates a unique physical ID and uses that ID for the
	// cache cluster. For more information, see Name Type.
	ClusterName *StringExpr `json:"ClusterName,omitempty"`

	// The name of the cache engine to be used for this cache cluster, such
	// as memcached or redis.
	Engine *StringExpr `json:"Engine,omitempty"`

	// The version of the cache engine to be used for this cluster.
	EngineVersion *StringExpr `json:"EngineVersion,omitempty"`

	// The Amazon Resource Name (ARN) of the Amazon Simple Notification
	// Service (SNS) topic to which notifications will be sent.
	NotificationTopicArn *StringExpr `json:"NotificationTopicArn,omitempty"`

	// The number of cache nodes that the cache cluster should have.
	NumCacheNodes *StringExpr `json:"NumCacheNodes,omitempty"`

	// The port number on which each of the cache nodes will accept
	// connections.
	Port *IntegerExpr `json:"Port,omitempty"`

	// The Amazon EC2 Availability Zone in which the cache cluster is
	// created.
	PreferredAvailabilityZone *StringExpr `json:"PreferredAvailabilityZone,omitempty"`

	// For Memcached cache clusters, the list of Availability Zones in which
	// cache nodes are created. The number of Availability Zones listed must
	// equal the number of cache nodes. For example, if you want to create
	// three nodes in two different Availability Zones, you can specify
	// ["us-east-1a", "us-east-1a", "us-east-1b"], which would create two
	// nodes in us-east-1a and one node in us-east-1b.
	PreferredAvailabilityZones *StringListExpr `json:"PreferredAvailabilityZones,omitempty"`

	// The weekly time range (in UTC) during which system maintenance can
	// occur.
	PreferredMaintenanceWindow *StringExpr `json:"PreferredMaintenanceWindow,omitempty"`

	// The ARN of the snapshot file that you want to use to seed a new Redis
	// cache cluster. If you manage a Redis instance outside of Amazon
	// ElastiCache, you can create a new cache cluster in ElastiCache by
	// using a snapshot file that is stored in an Amazon S3 bucket.
	SnapshotArns *StringListExpr `json:"SnapshotArns,omitempty"`

	// The name of a snapshot from which to restore data into a new Redis
	// cache cluster.
	SnapshotName *StringExpr `json:"SnapshotName,omitempty"`

	// For Redis cache clusters, the number of days for which ElastiCache
	// retains automatic snapshots before deleting them. For example, if you
	// set the value to 5, a snapshot that was taken today will be retained
	// for 5 days before being deleted.
	SnapshotRetentionLimit *IntegerExpr `json:"SnapshotRetentionLimit,omitempty"`

	// For Redis cache clusters, the daily time range (in UTC) during which
	// ElastiCache will begin taking a daily snapshot of your node group. For
	// example, you can specify 05:00-09:00.
	SnapshotWindow *StringExpr `json:"SnapshotWindow,omitempty"`

	// A list of VPC security group IDs. If your cache cluster isn't in a
	// VPC, specify the CacheSecurityGroupNames property instead.
	XVpcSecurityGroupIdsX *StringListExpr `json:" VpcSecurityGroupIds ,omitempty"`
}

// ResourceType returns AWS::ElastiCache::CacheCluster to implement the ResourceProperties interface
func (s ElastiCacheCacheCluster) ResourceType() string {
	return "AWS::ElastiCache::CacheCluster"
}

// ElastiCacheParameterGroup represents AWS::ElastiCache::ParameterGroup
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-parameter-group.html
type ElastiCacheParameterGroup struct {
	// The name of the cache parameter group family that the cache parameter
	// group can be used with.
	CacheParameterGroupFamily *StringExpr `json:"CacheParameterGroupFamily,omitempty"`

	// The description for the Cache Parameter Group.
	Description *StringExpr `json:"Description,omitempty"`

	// A comma-delimited list of parameter name/value pairs. For more
	// information, go to ModifyCacheParameterGroup in the Amazon ElastiCache
	// API Reference Guide.
	Properties interface{} `json:"Properties,omitempty"`
}

// ResourceType returns AWS::ElastiCache::ParameterGroup to implement the ResourceProperties interface
func (s ElastiCacheParameterGroup) ResourceType() string {
	return "AWS::ElastiCache::ParameterGroup"
}

// ElastiCacheReplicationGroup represents AWS::ElastiCache::ReplicationGroup
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html
type ElastiCacheReplicationGroup struct {
	// Indicates whether Multi-AZ is enabled. When Multi-AZ is enabled, a
	// read-only replica is automatically promoted to a read-write primary
	// cluster if the existing primary cluster fails. If you specify true,
	// you must specify a value greater than 1 for the NumCacheNodes
	// property. By default, AWS CloudFormation sets the value to true.
	AutomaticFailoverEnabled *BoolExpr `json:"AutomaticFailoverEnabled,omitempty"`

	// Currently, this property isn't used by ElastiCache.
	AutoMinorVersionUpgrade *BoolExpr `json:"AutoMinorVersionUpgrade,omitempty"`

	// The compute and memory capacity of nodes in the node group. To see
	// valid values, see CreateReplicationGroup in the Amazon ElastiCache API
	// Reference Guide.
	CacheNodeType *StringExpr `json:"CacheNodeType,omitempty"`

	// The name of the parameter group to associate with this replication
	// group.
	CacheParameterGroupName *StringExpr `json:"CacheParameterGroupName,omitempty"`

	// A list of cache security group names to associate with this
	// replication group. If you specify the SecurityGroupIds property, do
	// not specify this property; you can specify only one.
	CacheSecurityGroupNames *StringListExpr `json:"CacheSecurityGroupNames,omitempty"`

	// The name of a cache subnet group to use for this replication group.
	CacheSubnetGroupName *StringExpr `json:"CacheSubnetGroupName,omitempty"`

	// The name of the cache engine to use for the cache clusters in this
	// replication group. Currently, you can specify only redis.
	Engine *StringExpr `json:"Engine,omitempty"`

	// The version number of the cache engine to use for the cache clusters
	// in this replication group.
	EngineVersion *StringExpr `json:"EngineVersion,omitempty"`

	// The Amazon Resource Name (ARN) of the Amazon Simple Notification
	// Service topic to which notifications are sent.
	NotificationTopicArn *StringExpr `json:"NotificationTopicArn,omitempty"`

	// The number of cache clusters for this replication group. If automatic
	// failover is enabled, you must specify a value greater than 1.
	NumCacheClusters *IntegerExpr `json:"NumCacheClusters,omitempty"`

	// The port number on which each member of the replication group accepts
	// connections.
	Port *IntegerExpr `json:"Port,omitempty"`

	// A list of Availability Zones (AZs) in which the cache clusters in this
	// replication group are created.
	PreferredCacheClusterAZs *StringListExpr `json:"PreferredCacheClusterAZs,omitempty"`

	// The weekly time range during which system maintenance can occur. Use
	// the following format to specify a time range: ddd:hh24:mi-ddd:hh24:mi
	// (24H Clock UTC). For example, you can specify sun:22:00-sun:23:30 for
	// Sunday from 10 PM to 11:30 PM.
	PreferredMaintenanceWindow *StringExpr `json:"PreferredMaintenanceWindow,omitempty"`

	// The description of the replication group.
	ReplicationGroupDescription *StringExpr `json:"ReplicationGroupDescription,omitempty"`

	// A list of Amazon Virtual Private Cloud (Amazon VPC) security groups to
	// associate with this replication group. Use this property only when you
	// are creating a replication group in a VPC. If you specify the
	// CacheSecurityGroupNames property, do not specify this property; you
	// can specify only one.
	SecurityGroupIds *StringListExpr `json:"SecurityGroupIds,omitempty"`

	// A single-element string list that specifies an ARN of a Redis .rdb
	// snapshot file that is stored in Amazon Simple Storage Service (Amazon
	// S3). The snapshot file populates the node group. The Amazon S3 object
	// name in the ARN cannot contain commas. For example, you can specify
	// arn:aws:s3:::my_bucket/snapshot1.rdb.
	SnapshotArns *StringListExpr `json:"SnapshotArns,omitempty"`

	// The number of days that ElastiCache retains automatic snapshots before
	// deleting them.
	SnapshotRetentionLimit *IntegerExpr `json:"SnapshotRetentionLimit,omitempty"`

	// The time range (in UTC) when ElastiCache takes a daily snapshot of
	// your node group. For example, you can specify 05:00-09:00.
	SnapshotWindow *StringExpr `json:"SnapshotWindow,omitempty"`
}

// ResourceType returns AWS::ElastiCache::ReplicationGroup to implement the ResourceProperties interface
func (s ElastiCacheReplicationGroup) ResourceType() string {
	return "AWS::ElastiCache::ReplicationGroup"
}

// ElastiCacheSecurityGroup represents AWS::ElastiCache::SecurityGroup
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group.html
type ElastiCacheSecurityGroup struct {
	// A description for the cache security group.
	Description *StringExpr `json:"Description,omitempty"`
}

// ResourceType returns AWS::ElastiCache::SecurityGroup to implement the ResourceProperties interface
func (s ElastiCacheSecurityGroup) ResourceType() string {
	return "AWS::ElastiCache::SecurityGroup"
}

// ElastiCacheSecurityGroupIngress represents AWS::ElastiCache::SecurityGroupIngress
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html
type ElastiCacheSecurityGroupIngress struct {
	// The name of the Cache Security Group to authorize.
	CacheSecurityGroupName *StringExpr `json:"CacheSecurityGroupName,omitempty"`

	// Name of the EC2 Security Group to include in the authorization.
	EC2SecurityGroupName *StringExpr `json:"EC2SecurityGroupName,omitempty"`

	// Specifies the AWS Account ID of the owner of the EC2 security group
	// specified in the EC2SecurityGroupName property. The AWS access key ID
	// is not an acceptable value.
	EC2SecurityGroupOwnerId *StringExpr `json:"EC2SecurityGroupOwnerId,omitempty"`
}

// ResourceType returns AWS::ElastiCache::SecurityGroupIngress to implement the ResourceProperties interface
func (s ElastiCacheSecurityGroupIngress) ResourceType() string {
	return "AWS::ElastiCache::SecurityGroupIngress"
}

// ElastiCacheSubnetGroup represents AWS::ElastiCache::SubnetGroup
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-subnetgroup.html
type ElastiCacheSubnetGroup struct {
}

// ResourceType returns AWS::ElastiCache::SubnetGroup  to implement the ResourceProperties interface
func (s ElastiCacheSubnetGroup) ResourceType() string {
	return "AWS::ElastiCache::SubnetGroup "
}

// ElasticBeanstalkApplication represents AWS::ElasticBeanstalk::Application
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk.html
type ElasticBeanstalkApplication struct {
	// A name for the Elastic Beanstalk application. If you don't specify a
	// name, AWS CloudFormation generates a unique physical ID and uses that
	// ID for the application name. For more information, see Name Type.
	ApplicationName *StringExpr `json:"ApplicationName,omitempty"`

	// An optional description of this application.
	Description *StringExpr `json:"Description,omitempty"`
}

// ResourceType returns AWS::ElasticBeanstalk::Application to implement the ResourceProperties interface
func (s ElasticBeanstalkApplication) ResourceType() string {
	return "AWS::ElasticBeanstalk::Application"
}

// ElasticBeanstalkApplicationVersion represents AWS::ElasticBeanstalk::ApplicationVersion
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html
type ElasticBeanstalkApplicationVersion struct {
}

// ResourceType returns AWS::ElasticBeanstalk::ApplicationVersion to implement the ResourceProperties interface
func (s ElasticBeanstalkApplicationVersion) ResourceType() string {
	return "AWS::ElasticBeanstalk::ApplicationVersion"
}

// ElasticBeanstalkConfigurationTemplate represents AWS::ElasticBeanstalk::ConfigurationTemplate
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-beanstalk-configurationtemplate.html
type ElasticBeanstalkConfigurationTemplate struct {
}

// ResourceType returns AWS::ElasticBeanstalk::ConfigurationTemplate to implement the ResourceProperties interface
func (s ElasticBeanstalkConfigurationTemplate) ResourceType() string {
	return "AWS::ElasticBeanstalk::ConfigurationTemplate"
}

// ElasticBeanstalkEnvironment represents AWS::ElasticBeanstalk::Environment
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html
type ElasticBeanstalkEnvironment struct {
	// The name of the application that is associated with this environment.
	ApplicationName *StringExpr `json:"ApplicationName,omitempty"`

	// A prefix for your Elastic Beanstalk environment URL.
	CNAMEPrefix *StringExpr `json:"CNAMEPrefix,omitempty"`

	// A description that helps you identify this environment.
	Description *StringExpr `json:"Description,omitempty"`

	// A name for the Elastic Beanstalk environment. If you don't specify a
	// name, AWS CloudFormation generates a unique physical ID and uses that
	// ID for the environment name. For more information, see Name Type.
	EnvironmentName *StringExpr `json:"EnvironmentName,omitempty"`

	// Key-value pairs defining configuration options for this environment.
	// These options override the values that are defined in the solution
	// stack or the configuration template. If you remove any options during
	// a stack update, the removed options revert to default values.
	OptionSettings *ElasticBeanstalkOptionSettingsList `json:"OptionSettings,omitempty"`

	// The name of an Elastic Beanstalk solution stack that this
	// configuration will use. For more information, see Supported Platforms
	// in the AWS Elastic Beanstalk Developer Guide. You must specify either
	// this parameter or an Elastic Beanstalk configuration template name.
	SolutionStackName *StringExpr `json:"SolutionStackName,omitempty"`

	// The name of the Elastic Beanstalk configuration template to use with
	// the environment. You must specify either this parameter or a solution
	// stack name.
	TemplateName *StringExpr `json:"TemplateName,omitempty"`

	// Specifies the tier to use in creating this environment. The
	// environment tier that you choose determines whether Elastic Beanstalk
	// provisions resources to support a web application that handles HTTP(S)
	// requests or a web application that handles background-processing
	// tasks.
	Tier *ElasticBeanstalkEnvironmentTier `json:"Tier,omitempty"`

	// The version to associate with the environment.
	VersionLabel *StringExpr `json:"VersionLabel,omitempty"`
}

// ResourceType returns AWS::ElasticBeanstalk::Environment to implement the ResourceProperties interface
func (s ElasticBeanstalkEnvironment) ResourceType() string {
	return "AWS::ElasticBeanstalk::Environment"
}

// ElasticLoadBalancingLoadBalancer represents AWS::ElasticLoadBalancing::LoadBalancer
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html
type ElasticLoadBalancingLoadBalancer struct {
	// Captures detailed information for all requests made to your load
	// balancer, such as the time a request was received, client’s IP
	// address, latencies, request path, and server responses.
	AccessLoggingPolicy *ElasticLoadBalancingAccessLoggingPolicy `json:"AccessLoggingPolicy,omitempty"`

	// Generates one or more stickiness policies with sticky session
	// lifetimes that follow that of an application-generated cookie. These
	// policies can be associated only with HTTP/HTTPS listeners.
	AppCookieStickinessPolicy *ElasticLoadBalancingAppCookieStickinessPolicyList `json:"AppCookieStickinessPolicy,omitempty"`

	// The Availability Zones in which to create the load balancer. You can
	// specify the AvailabilityZones or Subnets property, but not both.
	AvailabilityZones *StringListExpr `json:"AvailabilityZones,omitempty"`

	// Whether deregistered or unhealthy instances can complete all in-flight
	// requests.
	ConnectionDrainingPolicy *ElasticLoadBalancingConnectionDrainingPolicy `json:"ConnectionDrainingPolicy,omitempty"`

	// Specifies how long front-end and back-end connections of your load
	// balancer can remain idle.
	ConnectionSettings *ElasticLoadBalancingConnectionSettings `json:"ConnectionSettings,omitempty"`

	// Whether cross-zone load balancing is enabled for the load balancer.
	// With cross-zone load balancing, your load balancer nodes route traffic
	// to the back-end instances across all Availability Zones. By default
	// the CrossZone property is false.
	CrossZone *BoolExpr `json:"CrossZone,omitempty"`

	// Application health check for the instances.
	HealthCheck *ElasticLoadBalancingHealthCheck `json:"HealthCheck,omitempty"`

	// A list of EC2 instance IDs for the load balancer.
	Instances *StringListExpr `json:"Instances,omitempty"`

	// Generates a stickiness policy with sticky session lifetimes controlled
	// by the lifetime of the browser (user-agent), or by a specified
	// expiration period. This policy can be associated only with HTTP/HTTPS
	// listeners.
	LBCookieStickinessPolicy *ElasticLoadBalancingLBCookieStickinessPolicyList `json:"LBCookieStickinessPolicy,omitempty"`

	// A name for the load balancer. If you don't specify a name, AWS
	// CloudFormation generates a unique physical ID and uses that ID for the
	// load balancer. The name must be unique within your set of load
	// balancers. For more information, see Name Type.
	LoadBalancerName *StringExpr `json:"LoadBalancerName,omitempty"`

	// One or more listeners for this load balancer. Each listener must be
	// registered for a specific port, and you cannot have more than one
	// listener for a given port.
	Listeners *ElasticLoadBalancingListenerList `json:"Listeners,omitempty"`

	// A list of elastic load balancing policies to apply to this elastic
	// load balancer.
	Policies *ElasticLoadBalancingPolicyList `json:"Policies,omitempty"`

	// For load balancers attached to an Amazon VPC, this parameter can be
	// used to specify the type of load balancer to use. Specify internal to
	// create an internal load balancer with a DNS name that resolves to
	// private IP addresses or internet-facing to create a load balancer with
	// a publicly resolvable DNS name, which resolves to public IP addresses.
	Scheme *StringExpr `json:"Scheme,omitempty"`

	// Required: No
	SecurityGroups interface{} `json:"SecurityGroups,omitempty"`

	// A list of subnet IDs in your virtual private cloud (VPC) to attach to
	// your load balancer. You can specify the AvailabilityZones or Subnets
	// property, but not both.
	Subnets *StringListExpr `json:"Subnets,omitempty"`

	// An arbitrary set of tags (key-value pairs) for this load balancer.
	Tags []ResourceTag `json:"Tags,omitempty"`
}

// ResourceType returns AWS::ElasticLoadBalancing::LoadBalancer to implement the ResourceProperties interface
func (s ElasticLoadBalancingLoadBalancer) ResourceType() string {
	return "AWS::ElasticLoadBalancing::LoadBalancer"
}

// IAMAccessKey represents AWS::IAM::AccessKey
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html
type IAMAccessKey struct {
	// This value is specific to AWS CloudFormation and can only be
	// incremented. Incrementing this value notifies AWS CloudFormation that
	// you want to rotate your access key. When you update your stack, AWS
	// CloudFormation will replace the existing access key with a new key.
	Serial *IntegerExpr `json:"Serial,omitempty"`

	// The status of the access key.
	Status *StringExpr `json:"Status,omitempty"`

	// The name of the user that the new key will belong to.
	UserName *StringExpr `json:"UserName,omitempty"`
}

// ResourceType returns AWS::IAM::AccessKey to implement the ResourceProperties interface
func (s IAMAccessKey) ResourceType() string {
	return "AWS::IAM::AccessKey"
}

// IAMGroup represents AWS::IAM::Group
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html
type IAMGroup struct {
	// One or more managed policy ARNs to attach to this group.
	ManagedPolicyArns *StringListExpr `json:"ManagedPolicyArns,omitempty"`

	// The path to the group. For more information about paths, see
	// Identifiers for IAM Entities in Using IAM.
	Path *StringExpr `json:"Path,omitempty"`

	// The policies to associate with this group. For information about
	// policies, see Overview of Policies in Using IAM.
	Policies *IAMPoliciesList `json:"Policies,omitempty"`
}

// ResourceType returns AWS::IAM::Group to implement the ResourceProperties interface
func (s IAMGroup) ResourceType() string {
	return "AWS::IAM::Group"
}

// IAMInstanceProfile represents AWS::IAM::InstanceProfile
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html
type IAMInstanceProfile struct {
	// The path associated with this IAM instance profile. For information
	// about IAM paths, see Friendly Names and Paths in the AWS Identity and
	// Access Management User Guide.
	Path *StringExpr `json:"Path,omitempty"`

	// The roles associated with this IAM instance profile.
	Roles interface{} `json:"Roles,omitempty"`
}

// ResourceType returns AWS::IAM::InstanceProfile to implement the ResourceProperties interface
func (s IAMInstanceProfile) ResourceType() string {
	return "AWS::IAM::InstanceProfile"
}

// IAMManagedPolicy represents AWS::IAM::ManagedPolicy
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html
type IAMManagedPolicy struct {
	// A description of the policy. For example, you can describe the
	// permissions that are defined in the policy.
	Description *StringExpr `json:"Description,omitempty"`

	// The names of groups to attach to this policy.
	Groups *StringListExpr `json:"Groups,omitempty"`

	// The path for the policy. By default, the path is /. For more
	// information, see IAM Identifiers in the Using IAM guide.
	Path *StringExpr `json:"Path,omitempty"`

	// Policies that define the permissions for this managed policy.
	PolicyDocument interface{} `json:"PolicyDocument,omitempty"`

	// The names of roles to attach to this policy.
	Roles *StringListExpr `json:"Roles,omitempty"`

	// The names of users to attach to this policy.
	Users *StringListExpr `json:"Users,omitempty"`
}

// ResourceType returns AWS::IAM::ManagedPolicy to implement the ResourceProperties interface
func (s IAMManagedPolicy) ResourceType() string {
	return "AWS::IAM::ManagedPolicy"
}

// IAMPolicy represents AWS::IAM::Policy
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html
type IAMPolicy struct {
	// The names of groups to which you want to add the policy.
	Groups *StringListExpr `json:"Groups,omitempty"`

	// A policy document that contains permissions to add to the specified
	// users or groups.
	PolicyDocument interface{} `json:"PolicyDocument,omitempty"`

	// The name of the policy.
	PolicyName *StringExpr `json:"PolicyName,omitempty"`

	// The names of AWS::IAM::Roles to attach to this policy.
	Roles *StringListExpr `json:"Roles,omitempty"`

	// The names of users for whom you want to add the policy.
	Users *StringListExpr `json:"Users,omitempty"`
}

// ResourceType returns AWS::IAM::Policy to implement the ResourceProperties interface
func (s IAMPolicy) ResourceType() string {
	return "AWS::IAM::Policy"
}

// IAMRole represents AWS::IAM::Role
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html
type IAMRole struct {
	// The IAM assume role policy that is associated with this role.
	AssumeRolePolicyDocument interface{} `json:"AssumeRolePolicyDocument,omitempty"`

	// One or more managed policy ARNs to attach to this role.
	ManagedPolicyArns *StringListExpr `json:"ManagedPolicyArns,omitempty"`

	// The path associated with this role. For information about IAM paths,
	// see Friendly Names and Paths in Using IAM.
	Path *StringExpr `json:"Path,omitempty"`

	// The policies to associate with this role. Policies can also be
	// specified externally. For sample templates that demonstrates both
	// embedded and external policies, see Template Examples.
	Policies *IAMPoliciesList `json:"Policies,omitempty"`
}

// ResourceType returns AWS::IAM::Role to implement the ResourceProperties interface
func (s IAMRole) ResourceType() string {
	return "AWS::IAM::Role"
}

// IAMUser represents AWS::IAM::User
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html
type IAMUser struct {
	// A name of a group to which you want to add the user.
	Groups *StringListExpr `json:"Groups,omitempty"`

	// Creates a login profile so that the user can access the AWS Management
	// Console.
	LoginProfile *IAMUserLoginProfile `json:"LoginProfile,omitempty"`

	// One or more managed policy ARNs to attach to this user.
	ManagedPolicyArns *StringListExpr `json:"ManagedPolicyArns,omitempty"`

	// The path for the user name. For more information about paths, see
	// Identifiers for IAM Entities in Using AWS Identity and Access
	// Management.
	Path *StringExpr `json:"Path,omitempty"`

	// The policies to associate with this user. For information about
	// policies, see Overview of Policies in [Using IAM].
	Policies *IAMPoliciesList `json:"Policies,omitempty"`
}

// ResourceType returns AWS::IAM::User to implement the ResourceProperties interface
func (s IAMUser) ResourceType() string {
	return "AWS::IAM::User"
}

// IAMUserToGroupAddition represents AWS::IAM::UserToGroupAddition
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html
type IAMUserToGroupAddition struct {
	// The name of group to add users to.
	GroupName *StringExpr `json:"GroupName,omitempty"`

	// Required: Yes
	Users interface{} `json:"Users,omitempty"`
}

// ResourceType returns AWS::IAM::UserToGroupAddition to implement the ResourceProperties interface
func (s IAMUserToGroupAddition) ResourceType() string {
	return "AWS::IAM::UserToGroupAddition"
}

// KinesisStream represents AWS::Kinesis::Stream
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-stream.html
type KinesisStream struct {
	// The number of shards that the stream uses. For greater provisioned
	// throughput, increase the number of shards.
	ShardCount *IntegerExpr `json:"ShardCount,omitempty"`
}

// ResourceType returns AWS::Kinesis::Stream to implement the ResourceProperties interface
func (s KinesisStream) ResourceType() string {
	return "AWS::Kinesis::Stream"
}

// LambdaFunction represents AWS::Lambda::Function
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html
type LambdaFunction struct {
	// The source code of your Lambda function.
	Code *LambdaFunctionCode `json:"Code,omitempty"`

	// A description of the function.
	Description *StringExpr `json:"Description,omitempty"`

	// The name of the function within your code that Lambda calls to start
	// running the code.
	Handler *StringExpr `json:"Handler,omitempty"`

	// The amount of memory, in MB, that is allocated to your Lambda
	// function. Lambda uses this value to infer the amount of CPU and memory
	// that is allocated to your function. Your function use case determines
	// your CPU and memory requirements. For example, a database operation
	// might need less memory than an image processing function. The default
	// value is 128 MB, and you must specify a value that is greater than or
	// equal to 128.
	MemorySize *IntegerExpr `json:"MemorySize,omitempty"`

	// The Amazon Resource Name (ARN) of the AWS Identity and Access
	// Management (IAM) role that Lambda assumes when it runs your code to
	// access AWS.
	Role *StringExpr `json:"Role,omitempty"`

	// The runtime environment for the Lambda function that you are
	// uploading. Currently, Lambda supports only nodejs.
	Runtime *StringExpr `json:"Runtime,omitempty"`

	// The function execution time (in seconds) after which Lambda terminates
	// the function. Because the execution time affects cost, set this value
	// based on the function's expected execution time. By default, Timeout
	// is set to 3 seconds.
	Timeout *IntegerExpr `json:"Timeout,omitempty"`
}

// ResourceType returns AWS::Lambda::Function to implement the ResourceProperties interface
func (s LambdaFunction) ResourceType() string {
	return "AWS::Lambda::Function"
}

// LogsLogGroup represents AWS::Logs::LogGroup
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html
type LogsLogGroup struct {
	// The number of days log events are kept in CloudWatch Logs. When a log
	// event expires, CloudWatch Logs automatically deletes it. For valid
	// values, see PutRetentionPolicy in the Amazon CloudWatch Logs API
	// Reference.
	RetentionInDays *IntegerExpr `json:"RetentionInDays,omitempty"`
}

// ResourceType returns AWS::Logs::LogGroup to implement the ResourceProperties interface
func (s LogsLogGroup) ResourceType() string {
	return "AWS::Logs::LogGroup"
}

// LogsMetricFilter represents AWS::Logs::MetricFilter
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html
type LogsMetricFilter struct {
	// Describes the pattern that CloudWatch Logs follows to interpret each
	// entry in a log. For example, a log entry might contain fields such as
	// timestamps, IP addresses, error codes, bytes transferred, and so on.
	// You use the pattern to specify those fields and to specify what to
	// look for in the log file. For example, if you're interested in error
	// codes that begin with 1234, your filter pattern might be [timestamps,
	// ip_addresses, error_codes = 1234*, size, ...].
	FilterPattern *StringListExpr `json:"FilterPattern,omitempty"`

	// The name of an existing log group that you want to associate with this
	// metric filter.
	LogGroupName *StringExpr `json:"LogGroupName,omitempty"`

	// Describes how to transform data from a log into a CloudWatch metric.
	MetricTransformations *CloudWatchLogsMetricFilterMetricTransformationPropertyList `json:"MetricTransformations,omitempty"`
}

// ResourceType returns AWS::Logs::MetricFilter to implement the ResourceProperties interface
func (s LogsMetricFilter) ResourceType() string {
	return "AWS::Logs::MetricFilter"
}

// OpsWorksApp represents AWS::OpsWorks::App
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html
type OpsWorksApp struct {
	// Contains the information required to retrieve an app from a
	// repository.
	AppSource *OpsWorksSource `json:"AppSource,omitempty"`

	// One or more user-defined key-value pairs to be added to the stack
	// attributes bag.
	Attributes interface{} `json:"Attributes,omitempty"`

	// A description of the app.
	Description *StringExpr `json:"Description,omitempty"`

	// The app virtual host settings, with multiple domains separated by
	// commas. For example, 'www.example.com, example.com'.
	Domains *StringListExpr `json:"Domains,omitempty"`

	// Whether to enable SSL for this app.
	EnableSsl *BoolExpr `json:"EnableSsl,omitempty"`

	// The AWS OpsWorks app name.
	Name *StringExpr `json:"Name,omitempty"`

	// The app short name, which is used internally by AWS OpsWorks and by
	// Chef recipes.
	Shortname *StringExpr `json:"Shortname,omitempty"`

	// The SSL configuration
	SslConfiguration *OpsWorksSslConfiguration `json:"SslConfiguration,omitempty"`

	// The AWS OpsWorks stack ID that this app will be associated with.
	StackId *StringExpr `json:"StackId,omitempty"`

	// The app type. Each supported type is associated with a particular
	// layer. For more information, see CreateApp in the AWS OpsWorks API
	// Reference.
	Type *StringExpr `json:"Type,omitempty"`
}

// ResourceType returns AWS::OpsWorks::App to implement the ResourceProperties interface
func (s OpsWorksApp) ResourceType() string {
	return "AWS::OpsWorks::App"
}

// OpsWorksElasticLoadBalancerAttachment represents AWS::OpsWorks::ElasticLoadBalancerAttachment
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-elbattachment.html
type OpsWorksElasticLoadBalancerAttachment struct {
	// Elastic Load Balancing load balancer name.
	ElasticLoadBalancerName *StringExpr `json:"ElasticLoadBalancerName,omitempty"`

	// The AWS OpsWorks layer ID that the Elastic Load Balancing load
	// balancer will be attached to.
	LayerId *StringExpr `json:"LayerId,omitempty"`
}

// ResourceType returns AWS::OpsWorks::ElasticLoadBalancerAttachment to implement the ResourceProperties interface
func (s OpsWorksElasticLoadBalancerAttachment) ResourceType() string {
	return "AWS::OpsWorks::ElasticLoadBalancerAttachment"
}

// OpsWorksInstance represents AWS::OpsWorks::Instance
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html
type OpsWorksInstance struct {
	// The ID of the custom AMI to be used to create the instance. The AMI
	// should be based on one of the standard AWS OpsWorks APIs.
	AmiId *StringExpr `json:"AmiId,omitempty"`

	// The instance architecture.
	Architecture *StringExpr `json:"Architecture,omitempty"`

	// For scaling instances, the type of scaling. If you specify load-based
	// scaling, do not specify a time-based scaling configuration. For valid
	// values, see CreateInstance in the AWS OpsWorks API Reference.
	AutoScalingType *StringExpr `json:"AutoScalingType,omitempty"`

	// The instance Availability Zone.
	AvailabilityZone *StringExpr `json:"AvailabilityZone,omitempty"`

	// Whether to install operating system and package updates when the
	// instance boots.
	InstallUpdatesOnBoot *BoolExpr `json:"InstallUpdatesOnBoot,omitempty"`

	// The instance type, which must be supported by AWS OpsWorks. For more
	// information, see CreateInstance in the AWS OpsWorks API Reference.
	InstanceType *StringExpr `json:"InstanceType,omitempty"`

	// The IDs of the AWS OpsWorks layers that will be associated with this
	// instance.
	LayerIds *StringListExpr `json:"LayerIds,omitempty"`

	// The instance operating system. For more information, see
	// CreateInstance in the AWS OpsWorks API Reference.
	Os *StringExpr `json:"Os,omitempty"`

	// The instance root device type.
	RootDeviceType *StringExpr `json:"RootDeviceType,omitempty"`

	// The instance SSH key name.
	SshKeyName *StringExpr `json:"SshKeyName,omitempty"`

	// The ID of the AWS OpsWorks stack that this instance will be associated
	// with.
	StackId *StringExpr `json:"StackId,omitempty"`

	// The ID of the instance's subnet. If the stack is running in a VPC, you
	// can use this parameter to override the stack's default subnet ID value
	// and direct AWS OpsWorks to launch the instance in a different subnet.
	SubnetId *StringExpr `json:"SubnetId,omitempty"`

	// The time-based scaling configuration for the instance.
	TimeBasedAutoScaling *OpsWorksTimeBasedAutoScaling `json:"TimeBasedAutoScaling,omitempty"`
}

// ResourceType returns AWS::OpsWorks::Instance to implement the ResourceProperties interface
func (s OpsWorksInstance) ResourceType() string {
	return "AWS::OpsWorks::Instance"
}

// OpsWorksLayer represents AWS::OpsWorks::Layer
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html
type OpsWorksLayer struct {
	// One or more user-defined key-value pairs to be added to the stack
	// attributes bag.
	Attributes interface{} `json:"Attributes,omitempty"`

	// Whether to automatically assign an Elastic IP address to Amazon EC2
	// instances in this layer.
	AutoAssignElasticIps *BoolExpr `json:"AutoAssignElasticIps,omitempty"`

	// For AWS OpsWorks stacks that are running in a VPC, whether to
	// automatically assign a public IP address to Amazon EC2 instances in
	// this layer.
	AutoAssignPublicIps *BoolExpr `json:"AutoAssignPublicIps,omitempty"`

	// The Amazon Resource Name (ARN) of an IAM instance profile that is to
	// be used for the Amazon EC2 instances in this layer.
	CustomInstanceProfileArn *StringExpr `json:"CustomInstanceProfileArn,omitempty"`

	// Custom event recipes for this layer.
	CustomRecipes *OpsWorksRecipes `json:"CustomRecipes,omitempty"`

	// Custom security group IDs for this layer.
	CustomSecurityGroupIds *StringListExpr `json:"CustomSecurityGroupIds,omitempty"`

	// Whether to automatically heal Amazon EC2 instances that have become
	// disconnected or timed out.
	EnableAutoHealing *BoolExpr `json:"EnableAutoHealing,omitempty"`

	// Whether to install operating system and package updates when the
	// instance boots.
	InstallUpdatesOnBoot *BoolExpr `json:"InstallUpdatesOnBoot,omitempty"`

	// The lifecycle events for the AWS OpsWorks layer.
	LifecycleEventConfiguration *OpsWorksLayerLifeCycleConfiguration `json:"LifecycleEventConfiguration,omitempty"`

	// The load-based scaling configuration for the AWS OpsWorks layer.
	LoadBasedAutoScaling *OpsWorksLoadBasedAutoScaling `json:"LoadBasedAutoScaling,omitempty"`

	// The AWS OpsWorks layer name.
	Name *StringExpr `json:"Name,omitempty"`

	// The packages for this layer.
	Packages *StringListExpr `json:"Packages,omitempty"`

	// The layer short name, which is used internally by AWS OpsWorks and by
	// Chef recipes. The short name is also used as the name for the
	// directory where your app files are installed.
	Shortname *StringExpr `json:"Shortname,omitempty"`

	// The ID of the AWS OpsWorks stack that this layer will be associated
	// with.
	StackId *StringExpr `json:"StackId,omitempty"`

	// The layer type. A stack cannot have more than one layer of the same
	// type. For more information, see CreateLayer in the AWS OpsWorks API
	// Reference.
	Type *StringExpr `json:"Type,omitempty"`

	// Describes the Amazon EBS volumes for this layer.
	VolumeConfigurations *OpsWorksVolumeConfigurationList `json:"VolumeConfigurations,omitempty"`
}

// ResourceType returns AWS::OpsWorks::Layer to implement the ResourceProperties interface
func (s OpsWorksLayer) ResourceType() string {
	return "AWS::OpsWorks::Layer"
}

// OpsWorksStack represents AWS::OpsWorks::Stack
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html
type OpsWorksStack struct {
	// One or more user-defined key-value pairs to be added to the stack
	// attributes bag.
	Attributes interface{} `json:"Attributes,omitempty"`

	// Describes the Chef configuration. For more information, see the
	// CreateStack ChefConfiguration parameter in the AWS OpsWorks API
	// Reference.
	ChefConfiguration *OpsWorksChefConfiguration `json:"ChefConfiguration,omitempty"`

	// Describes the configuration manager. When you create a stack, you use
	// the configuration manager to specify the Chef version. For supported
	// Chef versions, see the CreateStack ConfigurationManager parameter in
	// the AWS OpsWorks API Reference.
	ConfigurationManager *OpsWorksStackConfigurationManager `json:"ConfigurationManager,omitempty"`

	// Contains the information required to retrieve a cookbook from a
	// repository.
	CustomCookbooksSource *OpsWorksSource `json:"CustomCookbooksSource,omitempty"`

	// A user-defined custom JSON object. The custom JSON is used to override
	// the corresponding default stack configuration JSON values. For more
	// information, see CreateStack in the AWS OpsWorks API Reference.
	CustomJson interface{} `json:"CustomJson,omitempty"`

	// The stack's default Availability Zone, which must be in the specified
	// region.
	DefaultAvailabilityZone *StringExpr `json:"DefaultAvailabilityZone,omitempty"`

	// The Amazon Resource Name (ARN) of an IAM instance profile that is the
	// default profile for all of the stack's Amazon EC2 instances.
	DefaultInstanceProfileArn *StringExpr `json:"DefaultInstanceProfileArn,omitempty"`

	// The stack's default operating system. For more information, see
	// CreateStack in the AWS OpsWorks API Reference.
	DefaultOs *StringExpr `json:"DefaultOs,omitempty"`

	// The default root device type. This value is used by default for all
	// instances in the stack, but you can override it when you create an
	// instance. For more information, see CreateStack in the AWS OpsWorks
	// API Reference.
	DefaultRootDeviceType *StringExpr `json:"DefaultRootDeviceType,omitempty"`

	// A default SSH key for the stack instances. You can override this value
	// when you create or update an instance.
	DefaultSshKeyName *StringExpr `json:"DefaultSshKeyName,omitempty"`

	// The stack's default subnet ID. All instances are launched into this
	// subnet unless you specify another subnet ID when you create the
	// instance.
	DefaultSubnetId *StringExpr `json:"DefaultSubnetId,omitempty"`

	// The stack's host name theme, with spaces replaced by underscores. The
	// theme is used to generate host names for the stack's instances. For
	// more information, see CreateStack in the AWS OpsWorks API Reference.
	HostnameTheme *StringExpr `json:"HostnameTheme,omitempty"`

	// The name of the AWS OpsWorks stack.
	Name *StringExpr `json:"Name,omitempty"`

	// The AWS Identity and Access Management (IAM) role that AWS OpsWorks
	// uses to work with AWS resources on your behalf. You must specify an
	// Amazon Resource Name (ARN) for an existing IAM role.
	ServiceRoleArn *StringExpr `json:"ServiceRoleArn,omitempty"`

	// Whether the stack uses custom cookbooks.
	UseCustomCookbooks *BoolExpr `json:"UseCustomCookbooks,omitempty"`

	// Whether to associate the AWS OpsWorks built-in security groups with
	// the stack's layers.
	UseOpsworksSecurityGroups *BoolExpr `json:"UseOpsworksSecurityGroups,omitempty"`

	// The ID of the VPC that the stack is to be launched into, which must be
	// in the specified region. All instances are launched into this VPC. If
	// you specify this property, you must specify the DefaultSubnetId
	// property.
	VpcId *StringExpr `json:"VpcId,omitempty"`
}

// ResourceType returns AWS::OpsWorks::Stack to implement the ResourceProperties interface
func (s OpsWorksStack) ResourceType() string {
	return "AWS::OpsWorks::Stack"
}

// RedshiftCluster represents AWS::Redshift::Cluster
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html
type RedshiftCluster struct {
	// When a new version of the Amazon Redshift is released, indicates
	// whether upgrades can be applied to the engine that is running on the
	// cluster. The upgrades are applied during the maintenance window.
	AllowVersionUpgrade *BoolExpr `json:"AllowVersionUpgrade,omitempty"`

	// The number of days that automated snapshots are retained. If you set
	// the value to 0, automated snapshots are disabled.
	AutomatedSnapshotRetentionPeriod *IntegerExpr `json:"AutomatedSnapshotRetentionPeriod,omitempty"`

	// The Amazon EC2 Availability Zone in which you want to provision your
	// Amazon Redshift cluster. For example, if you have several Amazon EC2
	// instances running in a specific Availability Zone, you might want the
	// cluster to be provisioned in the same zone in order to decrease
	// network latency.
	AvailabilityZone *StringExpr `json:"AvailabilityZone,omitempty"`

	// The name of the parameter group that you want to associate with this
	// cluster.
	ClusterParameterGroupName *StringExpr `json:"ClusterParameterGroupName,omitempty"`

	// A list of security groups that you want to associate with this
	// cluster.
	ClusterSecurityGroups *StringListExpr `json:"ClusterSecurityGroups,omitempty"`

	// The name of a cluster subnet group that you want to associate with
	// this cluster.
	ClusterSubnetGroupName *StringExpr `json:"ClusterSubnetGroupName,omitempty"`

	// The type of cluster. You can specify single-node or multi-node.
	ClusterType *StringExpr `json:"ClusterType,omitempty"`

	// The Amazon Redshift engine version that you want to deploy on the
	// cluster.
	ClusterVersion *StringExpr `json:"ClusterVersion,omitempty"`

	// The name of the first database that is created when the cluster is
	// created.
	DBName *StringExpr `json:"DBName,omitempty"`

	// The Elastic IP (EIP) address for the cluster.
	ElasticIp *StringExpr `json:"ElasticIp,omitempty"`

	// Indicates whether the data in the cluster is encrypted at rest.
	Encrypted *BoolExpr `json:"Encrypted,omitempty"`

	// Specifies the name of the HSM client certificate that the Amazon
	// Redshift cluster uses to retrieve the data encryption keys stored in
	// an HSM.
	HsmClientCertificateIdentifier *StringExpr `json:"HsmClientCertificateIdentifier,omitempty"`

	// Specifies the name of the HSM configuration that contains the
	// information that the Amazon Redshift cluster can use to retrieve and
	// store keys in an HSM.
	HsmConfigurationIdentifier *StringExpr `json:"HsmConfigurationIdentifier,omitempty"`

	// The user name that is associated with the master user account for this
	// cluster.
	MasterUsername *StringExpr `json:"MasterUsername,omitempty"`

	// The password associated with the master user account for this cluster.
	MasterUserPassword *StringExpr `json:"MasterUserPassword,omitempty"`

	// The node type that is provisioned for this cluster.
	NodeType *StringExpr `json:"NodeType,omitempty"`

	// The number of compute nodes in the cluster. If you specify multi-node
	// for the ClusterType parameter, you must specify a number greater than
	// 1.
	NumberOfNodes *IntegerExpr `json:"NumberOfNodes,omitempty"`

	// When you restore from a snapshot from another AWS account, the
	// 12-digit AWS account ID that contains that snapshot.
	OwnerAccount *StringExpr `json:"OwnerAccount,omitempty"`

	// The port number on which the cluster accepts incoming connections.
	Port *IntegerExpr `json:"Port,omitempty"`

	// The weekly time range (in UTC) during which automated cluster
	// maintenance can occur. The format of the time range is
	// ddd:hh24:mi-ddd:hh24:mi.
	PreferredMaintenanceWindow *StringExpr `json:"PreferredMaintenanceWindow,omitempty"`

	// Indicates whether the cluster can be accessed from a public network.
	PubliclyAccessible *BoolExpr `json:"PubliclyAccessible,omitempty"`

	// The name of the cluster the source snapshot was created from.
	SnapshotClusterIdentifier interface{} `json:"SnapshotClusterIdentifier,omitempty"`

	// The name of the snapshot from which to create a new cluster.
	SnapshotIdentifier *StringExpr `json:"SnapshotIdentifier,omitempty"`

	// A list of VPC security groups that are associated with this cluster.
	VpcSecurityGroupIds *StringListExpr `json:"VpcSecurityGroupIds,omitempty"`
}

// ResourceType returns AWS::Redshift::Cluster to implement the ResourceProperties interface
func (s RedshiftCluster) ResourceType() string {
	return "AWS::Redshift::Cluster"
}

// RedshiftClusterParameterGroup represents AWS::Redshift::ClusterParameterGroup
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clusterparametergroup.html
type RedshiftClusterParameterGroup struct {
	// A description of the parameter group.
	Description *StringExpr `json:"Description,omitempty"`

	// The Amazon Redshift engine version that applies to this cluster
	// parameter group. The cluster engine version determines the set of
	// parameters that you can specify in the Parameters property.
	ParameterGroupFamily *StringExpr `json:"ParameterGroupFamily,omitempty"`

	// A list of parameter names and values that are allowed by the Amazon
	// Redshift engine version that you specified in the ParameterGroupFamily
	// property. For more information, see Amazon Redshift Parameter Groups
	// in the Amazon Redshift Cluster Management Guide.
	Parameters *RedshiftParameterList `json:"Parameters,omitempty"`
}

// ResourceType returns AWS::Redshift::ClusterParameterGroup to implement the ResourceProperties interface
func (s RedshiftClusterParameterGroup) ResourceType() string {
	return "AWS::Redshift::ClusterParameterGroup"
}

// RedshiftClusterSecurityGroup represents AWS::Redshift::ClusterSecurityGroup
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroup.html
type RedshiftClusterSecurityGroup struct {
	// A description of the security group.
	Description *StringExpr `json:"Description,omitempty"`
}

// ResourceType returns AWS::Redshift::ClusterSecurityGroup to implement the ResourceProperties interface
func (s RedshiftClusterSecurityGroup) ResourceType() string {
	return "AWS::Redshift::ClusterSecurityGroup"
}

// RedshiftClusterSecurityGroupIngress represents AWS::Redshift::ClusterSecurityGroupIngress
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroupingress.html
type RedshiftClusterSecurityGroupIngress struct {
	// The name of the Amazon Redshift security group that will be associated
	// with the ingress rule.
	ClusterSecurityGroupName *StringExpr `json:"ClusterSecurityGroupName,omitempty"`

	// The IP address range that has inbound access to the Amazon Redshift
	// security group.
	CIDRIP *StringExpr `json:"CIDRIP,omitempty"`

	// The Amazon EC2 security group that will be added the Amazon Redshift
	// security group.
	EC2SecurityGroupName *StringExpr `json:"EC2SecurityGroupName,omitempty"`

	// The 12-digit AWS account number of the owner of the Amazon EC2
	// security group that is specified by the EC2SecurityGroupName
	// parameter.
	EC2SecurityGroupOwnerId *StringExpr `json:"EC2SecurityGroupOwnerId,omitempty"`
}

// ResourceType returns AWS::Redshift::ClusterSecurityGroupIngress to implement the ResourceProperties interface
func (s RedshiftClusterSecurityGroupIngress) ResourceType() string {
	return "AWS::Redshift::ClusterSecurityGroupIngress"
}

// RedshiftClusterSubnetGroup represents AWS::Redshift::ClusterSubnetGroup
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersubnetgroup.html
type RedshiftClusterSubnetGroup struct {
	// A description of the subnet group.
	Description *StringExpr `json:"Description,omitempty"`

	// A list of VPC subnet IDs. You can modify a maximum of 20 subnets.
	SubnetIds *StringListExpr `json:"SubnetIds,omitempty"`
}

// ResourceType returns AWS::Redshift::ClusterSubnetGroup to implement the ResourceProperties interface
func (s RedshiftClusterSubnetGroup) ResourceType() string {
	return "AWS::Redshift::ClusterSubnetGroup"
}

// RDSDBInstance represents AWS::RDS::DBInstance
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html
type RDSDBInstance struct {
	// The allocated storage size specified in gigabytes (GB).
	AllocatedStorage *StringExpr `json:"AllocatedStorage,omitempty"`

	// Indicates whether major version upgrades are allowed. Changing this
	// parameter does not result in an outage, and the change is applied
	// asynchronously as soon as possible.
	AllowMajorVersionUpgrade *BoolExpr `json:"AllowMajorVersionUpgrade,omitempty"`

	// Indicates that minor engine upgrades will be applied automatically to
	// the DB instance during the maintenance window. The default value is
	// true.
	AutoMinorVersionUpgrade *BoolExpr `json:"AutoMinorVersionUpgrade,omitempty"`

	// The name of the Availability Zone where the DB instance is located.
	// You cannot set the AvailabilityZone parameter if the MultiAZ parameter
	// is set to true.
	AvailabilityZone *StringExpr `json:"AvailabilityZone,omitempty"`

	// The number of days for which automatic DB snapshots are retained.
	BackupRetentionPeriod *StringExpr `json:"BackupRetentionPeriod,omitempty"`

	// For supported engines, specifies the character set to associate with
	// the database instance. For more information, see Appendix: Oracle
	// Character Sets Supported in Amazon RDS in the Amazon Relational
	// Database Service User Guide.
	CharacterSetName *StringExpr `json:"CharacterSetName,omitempty"`

	// The name of the compute and memory capacity class of the DB instance.
	DBInstanceClass *StringExpr `json:"DBInstanceClass,omitempty"`

	// A name for the DB instance. If you don't specify a name, AWS
	// CloudFormation generates a unique physical ID and uses that ID for the
	// DB instance. For more information, see Name Type.
	DBInstanceIdentifier *StringExpr `json:"DBInstanceIdentifier,omitempty"`

	// The name of the initial database of this instance that was provided at
	// create time, if one was specified. This same name is returned for the
	// life of the DB instance.
	DBName *StringExpr `json:"DBName,omitempty"`

	// The name of an existing DB parameter group or a reference to an
	// AWS::RDS::DBParameterGroup resource created in the template.
	DBParameterGroupName *StringExpr `json:"DBParameterGroupName,omitempty"`

	// A list of the DB security groups to assign to the Amazon RDS instance.
	// The list can include both the name of existing DB security groups or
	// references to AWS::RDS::DBSecurityGroup resources created in the
	// template.
	DBSecurityGroups *StringListExpr `json:"DBSecurityGroups,omitempty"`

	// The identifier for the DB snapshot to restore from.
	DBSnapshotIdentifier *StringExpr `json:"DBSnapshotIdentifier,omitempty"`

	// A DB subnet group to associate with the DB instance.
	DBSubnetGroupName *StringExpr `json:"DBSubnetGroupName,omitempty"`

	// The name of the database engine that the DB instance uses. This
	// property is optional when you specify the DBSnapshotIdentifier
	// property to create DB instances.
	Engine *StringExpr `json:"Engine,omitempty"`

	// The version number of the database engine to use.
	EngineVersion *StringExpr `json:"EngineVersion,omitempty"`

	// The number of I/O operations per second (IOPS) that the database
	// should provision. This can be any integer value from 1000 to 10,000,
	// in 1000 IOPS increments.
	Iops *IntegerExpr `json:"Iops,omitempty"`

	// The Amazon Resource Name (ARN) of the AWS Key Management Service
	// master key that is used to encrypt the database instance, such as
	// arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef.
	// If you enable the StorageEncrypted property but don't specify this
	// property, the default master key is used.
	KmsKeyId *StringExpr `json:"KmsKeyId,omitempty"`

	// The license model information for the DB instance.
	LicenseModel *StringExpr `json:"LicenseModel,omitempty"`

	// The master user name for the database instance. This property is
	// optional when you specify the DBSnapshotIdentifier property to create
	// DB instances.
	MasterUsername *StringExpr `json:"MasterUsername,omitempty"`

	// The master password for the database instance. This property is
	// optional when you specify the DBSnapshotIdentifier property to create
	// DB instances.
	MasterUserPassword *StringExpr `json:"MasterUserPassword,omitempty"`

	// Specifies if the database instance is a multiple Availability Zone
	// deployment. You cannot set the AvailabilityZone parameter if the
	// MultiAZ parameter is set to true.
	MultiAZ *BoolExpr `json:"MultiAZ,omitempty"`

	// An option group that this database instance is associated with.
	OptionGroupName *StringExpr `json:"OptionGroupName,omitempty"`

	// The port for the instance.
	Port *StringExpr `json:"Port,omitempty"`

	// The daily time range during which automated backups are created if
	// automated backups are enabled, as determined by the
	// BackupRetentionPeriod.
	PreferredBackupWindow *StringExpr `json:"PreferredBackupWindow,omitempty"`

	// The weekly time range (in UTC) during which system maintenance can
	// occur.
	PreferredMaintenanceWindow *StringExpr `json:"PreferredMaintenanceWindow,omitempty"`

	// Indicates whether the database instance is an Internet-facing
	// instance. If you specify true, an instance is created with a publicly
	// resolvable DNS name, which resolves to a public IP address. If you
	// specify false, an internal instance is created with a DNS name that
	// resolves to a private IP address.
	PubliclyAccessible *BoolExpr `json:"PubliclyAccessible,omitempty"`

	// If you want to create a read replica DB instance, specify the ID of
	// the source database instance. Each database instance can have a
	// certain number of read replicas. For more information, see Working
	// with Read Replicas in the Amazon Relational Database Service Developer
	// Guide.
	SourceDBInstanceIdentifier *StringExpr `json:"SourceDBInstanceIdentifier,omitempty"`

	// Indicates whether the database instance is encrypted.
	StorageEncrypted *BoolExpr `json:"StorageEncrypted,omitempty"`

	// The storage type associated with this database instance.
	StorageType *StringExpr `json:"StorageType,omitempty"`

	// An arbitrary set of tags (key–value pairs) for this database
	// instance.
	Tags []ResourceTag `json:"Tags,omitempty"`

	// A list of the VPC security groups to assign to the Amazon RDS
	// instance. The list can include both the physical IDs of existing VPC
	// security groups or references to AWS::EC2::SecurityGroup resources
	// created in the template.
	VPCSecurityGroups *StringListExpr `json:"VPCSecurityGroups,omitempty"`
}

// ResourceType returns AWS::RDS::DBInstance to implement the ResourceProperties interface
func (s RDSDBInstance) ResourceType() string {
	return "AWS::RDS::DBInstance"
}

// RDSDBParameterGroup represents AWS::RDS::DBParameterGroup
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbparametergroup.html
type RDSDBParameterGroup struct {
	// A friendly description of the RDS parameter group. For example, "My
	// Parameter Group".
	Description interface{} `json:"Description,omitempty"`

	// The database family of this RDS parameter group. For example,
	// "MySQL5.1".
	Family interface{} `json:"Family,omitempty"`

	// The parameters to set for this RDS parameter group.
	Parameters interface{} `json:"Parameters,omitempty"`

	// The tags that you want to attach to the RDS parameter group.
	Tags *ResourceTagList `json:"Tags,omitempty"`
}

// ResourceType returns AWS::RDS::DBParameterGroup to implement the ResourceProperties interface
func (s RDSDBParameterGroup) ResourceType() string {
	return "AWS::RDS::DBParameterGroup"
}

// RDSDBSubnetGroup represents AWS::RDS::DBSubnetGroup
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsubnet-group.html
type RDSDBSubnetGroup struct {
	// The description for the DB Subnet Group.
	DBSubnetGroupDescription *StringExpr `json:"DBSubnetGroupDescription,omitempty"`

	// The EC2 Subnet IDs for the DB Subnet Group.
	SubnetIds *StringListExpr `json:"SubnetIds,omitempty"`

	// The tags that you want to attach to the RDS database subnet group.
	Tags *ResourceTagList `json:"Tags,omitempty"`
}

// ResourceType returns AWS::RDS::DBSubnetGroup to implement the ResourceProperties interface
func (s RDSDBSubnetGroup) ResourceType() string {
	return "AWS::RDS::DBSubnetGroup"
}

// RDSDBSecurityGroup represents AWS::RDS::DBSecurityGroup
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group.html
type RDSDBSecurityGroup struct {
	// The Id of VPC. Indicates which VPC this DB Security Group should
	// belong to.
	EC2VpcId *StringExpr `json:"EC2VpcId,omitempty"`

	// Network ingress authorization for an Amazon EC2 security group or an
	// IP address range.
	DBSecurityGroupIngress *RDSSecurityGroupRuleList `json:"DBSecurityGroupIngress,omitempty"`

	// Description of the security group.
	GroupDescription *StringExpr `json:"GroupDescription,omitempty"`

	// The tags that you want to attach to the Amazon RDS DB security group.
	Tags *ResourceTagList `json:"Tags,omitempty"`
}

// ResourceType returns AWS::RDS::DBSecurityGroup to implement the ResourceProperties interface
func (s RDSDBSecurityGroup) ResourceType() string {
	return "AWS::RDS::DBSecurityGroup"
}

// RDSDBSecurityGroupIngress represents AWS::RDS::DBSecurityGroupIngress
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-security-group-ingress.html
type RDSDBSecurityGroupIngress struct {
	// The IP range to authorize.
	CIDRIP *StringExpr `json:"CIDRIP,omitempty"`

	// The name (ARN) of the AWS::RDS::DBSecurityGroup to which this ingress
	// will be added.
	DBSecurityGroupName *StringExpr `json:"DBSecurityGroupName,omitempty"`

	// The ID of the VPC or EC2 security group to authorize.
	EC2SecurityGroupId *StringExpr `json:"EC2SecurityGroupId,omitempty"`

	// The name of the EC2 security group to authorize.
	EC2SecurityGroupName *StringExpr `json:"EC2SecurityGroupName,omitempty"`

	// The AWS Account Number of the owner of the EC2 security group
	// specified in the EC2SecurityGroupName parameter. The AWS Access Key ID
	// is not an acceptable value.
	EC2SecurityGroupOwnerId *StringExpr `json:"EC2SecurityGroupOwnerId,omitempty"`
}

// ResourceType returns AWS::RDS::DBSecurityGroupIngress to implement the ResourceProperties interface
func (s RDSDBSecurityGroupIngress) ResourceType() string {
	return "AWS::RDS::DBSecurityGroupIngress"
}

// RDSEventSubscription represents AWS::RDS::EventSubscription
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html
type RDSEventSubscription struct {
	// Indicates whether to activate the subscription. If you don't specify
	// this property, AWS CloudFormation activates the subscription.
	Enabled *BoolExpr `json:"Enabled,omitempty"`

	// A list of event categories that you want to subscribe to for a given
	// source type. If you don't specify this property, you are notified
	// about all event categories.
	EventCategories *StringListExpr `json:"EventCategories,omitempty"`

	// The Amazon Resource Name (ARN) of an Amazon SNS topic that you want to
	// send event notifications to.
	SnsTopicArn *StringExpr `json:"SnsTopicArn,omitempty"`

	// A list of identifiers for which Amazon RDS provides notification
	// events.
	SourceIds *StringListExpr `json:"SourceIds,omitempty"`

	// The type of source for which Amazon RDS provides notification events.
	// For example, if you want to be notified of events generated by a
	// database instance, set this parameter to db-instance. If you don't
	// specify a value, notifications are provided for all source types.
	SourceTypeX *StringExpr `json:"SourceType ,omitempty"`
}

// ResourceType returns AWS::RDS::EventSubscription to implement the ResourceProperties interface
func (s RDSEventSubscription) ResourceType() string {
	return "AWS::RDS::EventSubscription"
}

// RDSOptionGroup represents AWS::RDS::OptionGroup
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html
type RDSOptionGroup struct {
	// The name of the database engine that this option group is associated
	// with.
	EngineName *StringExpr `json:"EngineName,omitempty"`

	// The major version number of the database engine that this option group
	// is associated with.
	MajorEngineVersion *StringExpr `json:"MajorEngineVersion,omitempty"`

	// A description of the option group.
	OptionGroupDescription *StringExpr `json:"OptionGroupDescription,omitempty"`

	// The configurations for this option group.
	OptionConfigurations *RDSOptionGroupOptionConfigurations `json:"OptionConfigurations,omitempty"`

	// An arbitrary set of tags (key–value pairs) for this option group.
	Tags []ResourceTag `json:"Tags,omitempty"`
}

// ResourceType returns AWS::RDS::OptionGroup to implement the ResourceProperties interface
func (s RDSOptionGroup) ResourceType() string {
	return "AWS::RDS::OptionGroup"
}

// Route53HealthCheck represents AWS::Route53::HealthCheck
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html
type Route53HealthCheck struct {
	// An Amazon Route 53 health check.
	HealthCheckConfig *Route53HealthCheckConfig `json:"HealthCheckConfig,omitempty"`

	// An arbitrary set of tags (key–value pairs) for this health check.
	HealthCheckTags *Route53HealthCheckTagsList `json:"HealthCheckTags,omitempty"`
}

// ResourceType returns AWS::Route53::HealthCheck to implement the ResourceProperties interface
func (s Route53HealthCheck) ResourceType() string {
	return "AWS::Route53::HealthCheck"
}

// Route53HostedZone represents AWS::Route53::HostedZone
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html
type Route53HostedZone struct {
	// A complex type that contains an optional comment about your hosted
	// zone.
	HostedZoneConfig *Route53HostedZoneConfigProperty `json:"HostedZoneConfig,omitempty"`

	// An arbitrary set of tags (key–value pairs) for this hosted zone.
	HostedZoneTags *Route53HostedZoneTagsList `json:"HostedZoneTags,omitempty"`

	// The name of the domain. For resource record types that include a
	// domain name, specify a fully qualified domain name.
	Name *StringExpr `json:"Name,omitempty"`

	// One or more VPCs that you want to associate with this hosted zone.
	// When you specify this property, AWS CloudFormation creates a private
	// hosted zone.
	VPCs *Route53HostedZoneVPCsList `json:"VPCs,omitempty"`
}

// ResourceType returns AWS::Route53::HostedZone to implement the ResourceProperties interface
func (s Route53HostedZone) ResourceType() string {
	return "AWS::Route53::HostedZone"
}

// Route53RecordSet represents AWS::Route53::RecordSet
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html
type Route53RecordSet struct {
	// Alias resource record sets only: Information about the domain to which
	// you are redirecting traffic.
	AliasTarget *Route53AliasTargetProperty `json:"AliasTarget,omitempty"`

	// Any comments you want to include about the hosted zone.
	Comment *StringExpr `json:"Comment,omitempty"`

	// Designates the record set as a PRIMARY or SECONDARY failover record
	// set. When you have more than one resource performing the same
	// function, you can configure Amazon Route 53 to check the health of
	// your resources and use only health resources to respond to DNS
	// queries. You cannot create nonfailover resource record sets that have
	// the same Name and Type property values as failover resource record
	// sets. For more information, see the Failover element in the Amazon
	// Route 53 API Reference.
	Failover *StringExpr `json:"Failover,omitempty"`

	// Describes how Amazon Route 53 responds to DNS queries based on the
	// geographic origin of the query.
	GeoLocation *Route53RecordSetGeoLocationProperty `json:"GeoLocation,omitempty"`

	// The health check ID that you want to apply to this record set. Amazon
	// Route 53 returns this resource record set in response to a DNS query
	// only while record set is healthy.
	HealthCheckId *StringExpr `json:"HealthCheckId,omitempty"`

	// The ID of the hosted zone.
	HostedZoneId *StringExpr `json:"HostedZoneId,omitempty"`

	// The name of the domain for the hosted zone where you want to add the
	// record set.
	HostedZoneName *StringExpr `json:"HostedZoneName,omitempty"`

	// The name of the domain. This must be a fully specified domain, ending
	// with a period as the last label indication. If you omit the final
	// period, Amazon Route 53 assumes the domain is relative to the root.
	Name *StringExpr `json:"Name,omitempty"`

	// Latency resource record sets only: The Amazon EC2 region where the
	// resource that is specified in this resource record set resides. The
	// resource typically is an AWS resource, for example, Amazon EC2
	// instance or an Elastic Load Balancing load balancer, and is referred
	// to by an IP address or a DNS domain name, depending on the record
	// type.
	Region interface{} `json:"Region,omitempty"`

	// List of resource records to add. Each record should be in the format
	// appropriate for the record type specified by the Type property. For
	// information about different record types and their record formats, see
	// Appendix: Domain Name Format in the Amazon Route 53 Developer Guide.
	ResourceRecords *StringListExpr `json:"ResourceRecords,omitempty"`

	// A unique identifier that differentiates among multiple resource record
	// sets that have the same combination of DNS name and type.
	SetIdentifier *StringExpr `json:"SetIdentifier,omitempty"`

	// The resource record cache time to live (TTL), in seconds. If you
	// specify this property, do not specify the AliasTarget property. For
	// alias target records, the alias uses a TTL value from the target.
	TTL *StringExpr `json:"TTL,omitempty"`

	// The type of records to add.
	Type *StringExpr `json:"Type,omitempty"`

	// Weighted resource record sets only: Among resource record sets that
	// have the same combination of DNS name and type, a value that
	// determines what portion of traffic for the current resource record set
	// is routed to the associated location.
	Weight *IntegerExpr `json:"Weight,omitempty"`
}

// ResourceType returns AWS::Route53::RecordSet to implement the ResourceProperties interface
func (s Route53RecordSet) ResourceType() string {
	return "AWS::Route53::RecordSet"
}

// Route53RecordSetList represents a list of Route53RecordSet
type Route53RecordSetList []Route53RecordSet

// UnmarshalJSON sets the object from the provided JSON representation
func (l *Route53RecordSetList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := Route53RecordSet{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = Route53RecordSetList{item}
		return nil
	}
	list := []Route53RecordSet{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = Route53RecordSetList(list)
		return nil
	}
	return err
}

// Route53RecordSetGroup represents AWS::Route53::RecordSetGroup
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordsetgroup.html
type Route53RecordSetGroup struct {
	// The ID of the hosted zone.
	HostedZoneId *StringExpr `json:"HostedZoneId,omitempty"`

	// The name of the domain for the hosted zone where you want to add the
	// record set.
	HostedZoneName *StringExpr `json:"HostedZoneName,omitempty"`

	// List of resource record sets to add.
	RecordSets *Route53RecordSetList `json:"RecordSets,omitempty"`

	// Any comments you want to include about the hosted zone.
	Comment *StringExpr `json:"Comment,omitempty"`
}

// ResourceType returns AWS::Route53::RecordSetGroup to implement the ResourceProperties interface
func (s Route53RecordSetGroup) ResourceType() string {
	return "AWS::Route53::RecordSetGroup"
}

// S3Bucket represents AWS::S3::Bucket
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html
type S3Bucket struct {
	// A canned access control list (ACL) that grants predefined permissions
	// to the bucket. For more information about canned ACLs, see Canned ACLs
	// in the Amazon S3 documentation.
	AccessControl *StringExpr `json:"AccessControl,omitempty"`

	// A name for the bucket. If you don't specify a name, AWS CloudFormation
	// generates a unique physical ID and uses that ID for the bucket name.
	// For more information, see Name Type. The bucket name must contain only
	// lowercase letters, numbers, periods (.), and dashes (-).
	BucketName *StringExpr `json:"BucketName,omitempty"`

	// Rules that define cross-origin resource sharing of objects in this
	// bucket. For more information, see Enabling Cross-Origin Resource
	// Sharing in the Amazon Simple Storage Service Developer Guide.
	CorsConfiguration *S3CorsConfiguration `json:"CorsConfiguration,omitempty"`

	// Rules that define how Amazon S3 manages objects during their lifetime.
	// For more information, see Object Lifecycle Management in the Amazon
	// Simple Storage Service Developer Guide.
	LifecycleConfiguration *S3LifecycleConfiguration `json:"LifecycleConfiguration,omitempty"`

	// Settings that defines where logs are stored.
	LoggingConfiguration *S3LoggingConfiguration `json:"LoggingConfiguration,omitempty"`

	// Configuration that defines which Amazon SNS topic to send messages to
	// and what events to report.
	NotificationConfiguration *S3NotificationConfiguration `json:"NotificationConfiguration,omitempty"`

	// An arbitrary set of tags (key-value pairs) for this Amazon S3 bucket.
	Tags []ResourceTag `json:"Tags,omitempty"`

	// Enables multiple variants of all objects in this bucket. You might
	// enable versioning to prevent objects from being deleted or overwritten
	// by mistake or to archive objects so that you can retrieve previous
	// versions of them.
	VersioningConfiguration *S3VersioningConfiguration `json:"VersioningConfiguration,omitempty"`

	// Information used to configure the bucket as a static website. For more
	// information, see Hosting Websites on Amazon S3.
	WebsiteConfiguration *S3WebsiteConfigurationProperty `json:"WebsiteConfiguration,omitempty"`
}

// ResourceType returns AWS::S3::Bucket to implement the ResourceProperties interface
func (s S3Bucket) ResourceType() string {
	return "AWS::S3::Bucket"
}

// S3BucketPolicy represents AWS::S3::BucketPolicy
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-policy.html
type S3BucketPolicy struct {
	// The Amazon S3 bucket that the policy applies to.
	Bucket *StringExpr `json:"Bucket,omitempty"`

	// A policy document containing permissions to add to the specified
	// bucket.
	PolicyDocument interface{} `json:"PolicyDocument,omitempty"`
}

// ResourceType returns AWS::S3::BucketPolicy to implement the ResourceProperties interface
func (s S3BucketPolicy) ResourceType() string {
	return "AWS::S3::BucketPolicy"
}

// SDBDomain represents AWS::SDB::Domain
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-simpledb.html
type SDBDomain struct {
}

// ResourceType returns AWS::SDB::Domain to implement the ResourceProperties interface
func (s SDBDomain) ResourceType() string {
	return "AWS::SDB::Domain"
}

// SNSTopic represents AWS::SNS::Topic
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html
type SNSTopic struct {
	// A developer-defined string that can be used to identify this SNS
	// topic.
	DisplayName *StringExpr `json:"DisplayName,omitempty"`

	// The SNS subscriptions (endpoints) for this topic.
	Subscription *SNSSubscriptionList `json:"Subscription,omitempty"`

	// A name for the topic. If you don't specify a name, AWS CloudFormation
	// generates a unique physical ID and uses that ID for the topic name.
	// For more information, see Name Type.
	TopicName *StringExpr `json:"TopicName,omitempty"`
}

// ResourceType returns AWS::SNS::Topic to implement the ResourceProperties interface
func (s SNSTopic) ResourceType() string {
	return "AWS::SNS::Topic"
}

// SNSTopicPolicy represents AWS::SNS::TopicPolicy
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html
type SNSTopicPolicy struct {
	// A policy document that contains permissions to add to the specified
	// SNS topics.
	PolicyDocument interface{} `json:"PolicyDocument,omitempty"`

	// The Amazon Resource Names (ARN) of the topics to which you want to add
	// the policy. You can use the Ref function to specify an AWS::SNS::Topic
	// resource.
	Topics interface{} `json:"Topics,omitempty"`
}

// ResourceType returns AWS::SNS::TopicPolicy to implement the ResourceProperties interface
func (s SNSTopicPolicy) ResourceType() string {
	return "AWS::SNS::TopicPolicy"
}

// SQSQueue represents AWS::SQS::Queue
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html
type SQSQueue struct {
	// The time in seconds that the delivery of all messages in the queue
	// will be delayed. You can specify an integer value of 0 to 900 (15
	// minutes). The default value is 0.
	DelaySeconds *IntegerExpr `json:"DelaySeconds,omitempty"`

	// The limit of how many bytes a message can contain before Amazon SQS
	// rejects it. You can specify an integer value from 1024 bytes (1 KiB)
	// to 262144 bytes (256 KiB). The default value is 262144 (256 KiB).
	MaximumMessageSize *IntegerExpr `json:"MaximumMessageSize,omitempty"`

	// The number of seconds Amazon SQS retains a message. You can specify an
	// integer value from 60 seconds (1 minute) to 1209600 seconds (14 days).
	// The default value is 345600 seconds (4 days).
	MessageRetentionPeriod *IntegerExpr `json:"MessageRetentionPeriod,omitempty"`

	// A name for the queue. If you don't specify a name, AWS CloudFormation
	// generates a unique physical ID and uses that ID for the queue name.
	// For more information, see Name Type.
	QueueName *StringExpr `json:"QueueName,omitempty"`

	// Specifies the duration, in seconds, that the ReceiveMessage action
	// call waits until a message is in the queue in order to include it in
	// the response, as opposed to returning an empty response if a message
	// is not yet available. You can specify an integer from 1 to 20. The
	// short polling is used as the default or when you specify 0 for this
	// property. For more information, see Amazon SQS Long Poll.
	ReceiveMessageWaitTimeSeconds *IntegerExpr `json:"ReceiveMessageWaitTimeSeconds,omitempty"`

	// Specifies an existing dead letter queue to receive messages after the
	// source queue (this queue) fails to process a message a specified
	// number of times.
	RedrivePolicy *SQSRedrivePolicy `json:"RedrivePolicy,omitempty"`

	// The length of time during which the queue will be unavailable once a
	// message is delivered from the queue. This blocks other components from
	// receiving the same message and gives the initial component time to
	// process and delete the message from the queue.
	VisibilityTimeout *IntegerExpr `json:"VisibilityTimeout,omitempty"`
}

// ResourceType returns AWS::SQS::Queue to implement the ResourceProperties interface
func (s SQSQueue) ResourceType() string {
	return "AWS::SQS::Queue"
}

// SQSQueuePolicy represents AWS::SQS::QueuePolicy
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-policy.html
type SQSQueuePolicy struct {
	// A policy document containing permissions to add to the specified SQS
	// queues.
	PolicyDocument interface{} `json:"PolicyDocument,omitempty"`

	// The URLs of the queues to which you want to add the policy. You can
	// use the Ref function to specify an AWS::SQS::Queue resource.
	Queues *StringListExpr `json:"Queues,omitempty"`
}

// ResourceType returns AWS::SQS::QueuePolicy to implement the ResourceProperties interface
func (s SQSQueuePolicy) ResourceType() string {
	return "AWS::SQS::QueuePolicy"
}

// AutoScalingBlockDeviceMapping represents AWS CloudFormation AutoScaling Block Device Mapping Property Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-mapping.html
type AutoScalingBlockDeviceMapping struct {
	// The name of the device within Amazon EC2.
	DeviceName *StringExpr `json:"DeviceName,omitempty"`

	// The Amazon Elastic Block Store volume information.
	Ebs *AutoScalingEBSBlockDevice `json:"Ebs,omitempty"`

	// Suppresses the device mapping. If NoDevice is set to true for the root
	// device, the instance might fail the Amazon EC2 health check. Auto
	// Scaling launches a replacement instance if the instance fails the
	// health check.
	NoDevice *BoolExpr `json:"NoDevice,omitempty"`

	// The name of the virtual device. The name must be in the form
	// ephemeralX where X is a number starting from zero (0), for example,
	// ephemeral0.
	VirtualName *StringExpr `json:"VirtualName,omitempty"`
}

// AutoScalingBlockDeviceMappingList represents a list of AutoScalingBlockDeviceMapping
type AutoScalingBlockDeviceMappingList []AutoScalingBlockDeviceMapping

// UnmarshalJSON sets the object from the provided JSON representation
func (l *AutoScalingBlockDeviceMappingList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AutoScalingBlockDeviceMapping{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AutoScalingBlockDeviceMappingList{item}
		return nil
	}
	list := []AutoScalingBlockDeviceMapping{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AutoScalingBlockDeviceMappingList(list)
		return nil
	}
	return err
}

// AutoScalingEBSBlockDevice represents AWS CloudFormation AutoScaling EBS Block Device Property Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-template.html
type AutoScalingEBSBlockDevice struct {
	// Indicates whether to delete the volume when the instance is
	// terminated. By default, Auto Scaling uses true.
	DeleteOnTermination *BoolExpr `json:"DeleteOnTermination,omitempty"`

	// The number of I/O operations per second (IOPS) that the volume
	// supports. The maximum ratio of IOPS to volume size is 30.
	Iops *IntegerExpr `json:"Iops,omitempty"`

	// The snapshot ID of the volume to use.
	SnapshotId *StringExpr `json:"SnapshotId,omitempty"`

	// The volume size, in Gibibytes (GiB). This can be a number from 1 –
	// 1024. If the volume type is EBS optimized, the minimum value is 10.
	// For more information about specifying the volume type, see
	// EbsOptimized in AWS::AutoScaling::LaunchConfiguration.
	VolumeSize *IntegerExpr `json:"VolumeSize,omitempty"`

	// The volume type. By default, Auto Scaling uses the standard volume
	// type. For more information, see Ebs in the Auto Scaling API Reference.
	VolumeType *StringExpr `json:"VolumeType,omitempty"`
}

// AutoScalingEBSBlockDeviceList represents a list of AutoScalingEBSBlockDevice
type AutoScalingEBSBlockDeviceList []AutoScalingEBSBlockDevice

// UnmarshalJSON sets the object from the provided JSON representation
func (l *AutoScalingEBSBlockDeviceList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AutoScalingEBSBlockDevice{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AutoScalingEBSBlockDeviceList{item}
		return nil
	}
	list := []AutoScalingEBSBlockDevice{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AutoScalingEBSBlockDeviceList(list)
		return nil
	}
	return err
}

// AutoScalingMetricsCollection represents Auto Scaling MetricsCollection
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-metricscollection.html
type AutoScalingMetricsCollection struct {
	// The frequency at which Auto Scaling sends aggregated data to
	// CloudWatch. For example, you can specify 1Minute to send aggregated
	// data to CloudWatch every minute.
	Granularity *StringExpr `json:"Granularity,omitempty"`

	// The list of metrics to collect. If you don't specify any metrics, all
	// metrics are enabled.
	Metrics *StringListExpr `json:"Metrics,omitempty"`
}

// AutoScalingMetricsCollectionList represents a list of AutoScalingMetricsCollection
type AutoScalingMetricsCollectionList []AutoScalingMetricsCollection

// UnmarshalJSON sets the object from the provided JSON representation
func (l *AutoScalingMetricsCollectionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AutoScalingMetricsCollection{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AutoScalingMetricsCollectionList{item}
		return nil
	}
	list := []AutoScalingMetricsCollection{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AutoScalingMetricsCollectionList(list)
		return nil
	}
	return err
}

// AutoScalingNotificationConfigurations represents Auto Scaling NotificationConfigurations
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-notificationconfigurations.html
type AutoScalingNotificationConfigurations struct {
	// A list of event types that trigger a notification. Event types can
	// include any of the following types: autoscaling:EC2_INSTANCE_LAUNCH,
	// autoscaling:EC2_INSTANCE_LAUNCH_ERROR,
	// autoscaling:EC2_INSTANCE_TERMINATE,
	// autoscaling:EC2_INSTANCE_TERMINATE_ERROR, and
	// autoscaling:TEST_NOTIFICATION. For more information about event types,
	// see DescribeAutoScalingNotificationTypes in the Auto Scaling API
	// Reference.
	NotificationTypes *StringListExpr `json:"NotificationTypes,omitempty"`

	// The Amazon Resource Name (ARN) of the Amazon Simple Notification
	// Service (SNS) topic.
	TopicARN *StringExpr `json:"TopicARN,omitempty"`
}

// AutoScalingNotificationConfigurationsList represents a list of AutoScalingNotificationConfigurations
type AutoScalingNotificationConfigurationsList []AutoScalingNotificationConfigurations

// UnmarshalJSON sets the object from the provided JSON representation
func (l *AutoScalingNotificationConfigurationsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AutoScalingNotificationConfigurations{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AutoScalingNotificationConfigurationsList{item}
		return nil
	}
	list := []AutoScalingNotificationConfigurations{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AutoScalingNotificationConfigurationsList(list)
		return nil
	}
	return err
}

// AutoScalingTags represents Auto Scaling Tags Property Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-tags.html
type AutoScalingTags struct {
	// The key name of the tag.
	Key *StringExpr `json:"Key,omitempty"`

	// The value for the tag.
	Value *StringExpr `json:"Value,omitempty"`

	// Set to true if you want AWS CloudFormation to copy the tag to EC2
	// instances that are launched as part of the auto scaling group. Set to
	// false if you want the tag attached only to the auto scaling group and
	// not copied to any instances launched as part of the auto scaling
	// group.
	PropagateAtLaunch *BoolExpr `json:"PropagateAtLaunch,omitempty"`
}

// AutoScalingTagsList represents a list of AutoScalingTags
type AutoScalingTagsList []AutoScalingTags

// UnmarshalJSON sets the object from the provided JSON representation
func (l *AutoScalingTagsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AutoScalingTags{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AutoScalingTagsList{item}
		return nil
	}
	list := []AutoScalingTags{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AutoScalingTagsList(list)
		return nil
	}
	return err
}

// CloudFormationStackParameters represents CloudFormation Stack Parameters Property Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stack-parameters.html
type CloudFormationStackParameters struct {
}

// CloudFormationStackParametersList represents a list of CloudFormationStackParameters
type CloudFormationStackParametersList []CloudFormationStackParameters

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudFormationStackParametersList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudFormationStackParameters{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudFormationStackParametersList{item}
		return nil
	}
	list := []CloudFormationStackParameters{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudFormationStackParametersList(list)
		return nil
	}
	return err
}

// CloudFrontDistributionConfig represents CloudFront DistributionConfig
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html
type CloudFrontDistributionConfig struct {
	// CNAMEs (alternate domain names), if any, for the distribution.
	Aliases *StringListExpr `json:"Aliases,omitempty"`

	// A list of CacheBehavior types for the distribution.
	CacheBehaviors *CloudFrontDistributionConfigCacheBehaviorList `json:"CacheBehaviors,omitempty"`

	// Any comments that you want to include about the distribution.
	Comment *StringExpr `json:"Comment,omitempty"`

	// Whether CloudFront replaces HTTP status codes in the 4xx and 5xx range
	// with custom error messages before returning the response to the
	// viewer.
	CustomErrorResponses *CloudFrontDistributionConfigCustomErrorResponse `json:"CustomErrorResponses,omitempty"`

	// The default cache behavior that is triggered if you do not specify the
	// CacheBehavior property or if files don't match any of the values of
	// PathPattern in the CacheBehavior property.
	DefaultCacheBehavior *CloudFrontDefaultCacheBehavior `json:"DefaultCacheBehavior,omitempty"`

	// The object (such as index.html) that you want CloudFront to request
	// from your origin when the root URL for your distribution (such as
	// http://example.com/) is requested.
	DefaultRootObject *StringExpr `json:"DefaultRootObject,omitempty"`

	// Controls whether the distribution is enabled to accept end user
	// requests for content.
	Enabled *BoolExpr `json:"Enabled,omitempty"`

	// Controls whether access logs are written for the distribution. To turn
	// on access logs, specify this property.
	Logging *CloudFrontLogging `json:"Logging,omitempty"`

	// A list of origins for this CloudFront distribution. For each origin,
	// you can specify whether it is an Amazon S3 or custom origin.
	Origins *CloudFrontDistributionConfigOriginList `json:"Origins,omitempty"`

	// The price class that corresponds with the maximum price that you want
	// to pay for the CloudFront service. For more information, see Choosing
	// the Price Class in the Amazon CloudFront Developer Guide.
	PriceClass *StringExpr `json:"PriceClass,omitempty"`

	// Specifies restrictions on who or how viewers can access your content.
	Restrictions *CloudFrontDistributionConfigurationRestrictions `json:"Restrictions,omitempty"`

	// The certificate to use when viewers use HTTPS to request objects.
	ViewerCertificate *CloudFrontDistributionConfigurationViewerCertificate `json:"ViewerCertificate,omitempty"`
}

// CloudFrontDistributionConfigList represents a list of CloudFrontDistributionConfig
type CloudFrontDistributionConfigList []CloudFrontDistributionConfig

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudFrontDistributionConfigList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudFrontDistributionConfig{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudFrontDistributionConfigList{item}
		return nil
	}
	list := []CloudFrontDistributionConfig{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudFrontDistributionConfigList(list)
		return nil
	}
	return err
}

// CloudFrontDistributionConfigCacheBehavior represents CloudFront DistributionConfig CacheBehavior
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachebehavior.html
type CloudFrontDistributionConfigCacheBehavior struct {
	// HTTP methods that CloudFront processes and forwards to your Amazon S3
	// bucket or your custom origin. You can specify ["HEAD", "GET"], ["GET",
	// "HEAD", "OPTIONS"], or ["DELETE", "GET", "HEAD", "OPTIONS", "PATCH",
	// "POST", "PUT"]. If you don't specify a value, AWS CloudFormation
	// specifies ["HEAD", "GET"].
	AllowedMethods *StringListExpr `json:"AllowedMethods,omitempty"`

	// HTTP methods for which CloudFront caches responses. You can specify
	// ["HEAD", "GET"] or ["GET", "HEAD", "OPTIONS"]. If you don't specify a
	// value, AWS CloudFormation specifies ["HEAD", "GET"].
	CachedMethods *StringListExpr `json:"CachedMethods,omitempty"`

	// Specifies how CloudFront handles query strings or cookies.
	ForwardedValues *CloudFrontForwardedValues `json:"ForwardedValues,omitempty"`

	// The minimum amount of time that you want objects to stay in the cache
	// before CloudFront queries your origin to see whether the object has
	// been updated.
	MinTTL *StringExpr `json:"MinTTL,omitempty"`

	// The pattern to which this cache behavior applies. For example, you can
	// specify images/*.jpg.
	PathPattern *StringExpr `json:"PathPattern,omitempty"`

	// Indicates whether to use the origin that is associated with this cache
	// behavior to distribute media files in the Microsoft Smooth Streaming
	// format. If you specify true, you can still use this cache behavior to
	// distribute other content if the content matches the PathPattern value.
	SmoothStreaming *BoolExpr `json:"SmoothStreaming,omitempty"`

	// The ID value of the origin to which you want CloudFront to route
	// requests when a request matches the value of the PathPattern property.
	TargetOriginId *StringExpr `json:"TargetOriginId,omitempty"`

	// A list of AWS accounts that can create signed URLs in order to access
	// private content.
	TrustedSigners *StringListExpr `json:"TrustedSigners,omitempty"`

	// The protocol that users can use to access the files in the origin that
	// you specified in the TargetOriginId property when a request matches
	// the value of the PathPattern property.
	ViewerProtocolPolicy *StringExpr `json:"ViewerProtocolPolicy,omitempty"`
}

// CloudFrontDistributionConfigCacheBehaviorList represents a list of CloudFrontDistributionConfigCacheBehavior
type CloudFrontDistributionConfigCacheBehaviorList []CloudFrontDistributionConfigCacheBehavior

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudFrontDistributionConfigCacheBehaviorList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudFrontDistributionConfigCacheBehavior{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudFrontDistributionConfigCacheBehaviorList{item}
		return nil
	}
	list := []CloudFrontDistributionConfigCacheBehavior{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudFrontDistributionConfigCacheBehaviorList(list)
		return nil
	}
	return err
}

// CloudFrontDistributionConfigCustomErrorResponse represents CloudFront DistributionConfig CustomErrorResponse
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-customerrorresponse.html
type CloudFrontDistributionConfigCustomErrorResponse struct {
	// The minimum amount of time, in seconds, that Amazon CloudFront caches
	// the HTTP status code that you specified in the ErrorCode property. The
	// default value is 300.
	ErrorCachingMinTTL *IntegerExpr `json:"ErrorCachingMinTTL,omitempty"`

	// An HTTP status code for which you want to specify a custom error page.
	// You can specify 400, 403, 404, 405, 414, 500, 501, 502, 503, or 504.
	ErrorCode *IntegerExpr `json:"ErrorCode,omitempty"`

	// The HTTP status code that CloudFront returns to viewer along with the
	// custom error page. You can specify 200, 400, 403, 404, 405, 414, 500,
	// 501, 502, 503, or 504.
	ResponseCode *IntegerExpr `json:"ResponseCode,omitempty"`

	// The path to the custom error page that CloudFront returns to a viewer
	// when your origin returns the HTTP status code that you specified in
	// the ErrorCode property. For example, you can specify
	// /404-errors/403-forbidden.html.
	ResponsePagePath *StringExpr `json:"ResponsePagePath,omitempty"`
}

// CloudFrontDistributionConfigCustomErrorResponseList represents a list of CloudFrontDistributionConfigCustomErrorResponse
type CloudFrontDistributionConfigCustomErrorResponseList []CloudFrontDistributionConfigCustomErrorResponse

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudFrontDistributionConfigCustomErrorResponseList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudFrontDistributionConfigCustomErrorResponse{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudFrontDistributionConfigCustomErrorResponseList{item}
		return nil
	}
	list := []CloudFrontDistributionConfigCustomErrorResponse{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudFrontDistributionConfigCustomErrorResponseList(list)
		return nil
	}
	return err
}

// CloudFrontDefaultCacheBehavior represents CloudFront DefaultCacheBehavior
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-defaultcachebehavior.html
type CloudFrontDefaultCacheBehavior struct {
	// HTTP methods that CloudFront processes and forwards to your Amazon S3
	// bucket or your custom origin. You can specify ["HEAD", "GET"], ["GET",
	// "HEAD", "OPTIONS"], or ["DELETE", "GET", "HEAD", "OPTIONS", "PATCH",
	// "POST", "PUT"]. If you don't specify a value, AWS CloudFormation
	// specifies ["HEAD", "GET"].
	AllowedMethods *StringListExpr `json:"AllowedMethods,omitempty"`

	// HTTP methods for which CloudFront caches responses. You can specify
	// ["HEAD", "GET"] or ["GET", "HEAD", "OPTIONS"]. If you don't specify a
	// value, AWS CloudFormation specifies ["HEAD", "GET"].
	CachedMethods *StringListExpr `json:"CachedMethods,omitempty"`

	// Specifies how CloudFront handles query strings or cookies.
	ForwardedValues *CloudFrontForwardedValues `json:"ForwardedValues,omitempty"`

	// The minimum amount of time that you want objects to stay in the cache
	// before CloudFront queries your origin to see whether the object has
	// been updated.
	MinTTL *StringExpr `json:"MinTTL,omitempty"`

	// Indicates whether to use the origin that is associated with this cache
	// behavior to distribute media files in the Microsoft Smooth Streaming
	// format.
	SmoothStreaming *BoolExpr `json:"SmoothStreaming,omitempty"`

	// The value of ID for the origin that CloudFront routes requests to when
	// the default cache behavior is applied to a request.
	TargetOriginId *StringExpr `json:"TargetOriginId,omitempty"`

	// A list of AWS accounts that can create signed URLs in order to access
	// private content.
	TrustedSigners *StringListExpr `json:"TrustedSigners,omitempty"`

	// The protocol that users can use to access the files in the origin that
	// you specified in the TargetOriginId property when the default cache
	// behavior is applied to a request.
	ViewerProtocolPolicy *StringExpr `json:"ViewerProtocolPolicy,omitempty"`
}

// CloudFrontDefaultCacheBehaviorList represents a list of CloudFrontDefaultCacheBehavior
type CloudFrontDefaultCacheBehaviorList []CloudFrontDefaultCacheBehavior

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudFrontDefaultCacheBehaviorList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudFrontDefaultCacheBehavior{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudFrontDefaultCacheBehaviorList{item}
		return nil
	}
	list := []CloudFrontDefaultCacheBehavior{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudFrontDefaultCacheBehaviorList(list)
		return nil
	}
	return err
}

// CloudFrontLogging represents CloudFront Logging
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-logging.html
type CloudFrontLogging struct {
	// The Amazon S3 bucket address where access logs are stored, for
	// example, mybucket.s3.amazonaws.com.
	Bucket *StringExpr `json:"Bucket,omitempty"`

	// Indicates whether CloudFront includes cookies in access logs.
	IncludeCookies *BoolExpr `json:"IncludeCookies,omitempty"`

	// A prefix for the access log file names for this distribution.
	Prefix *StringExpr `json:"Prefix,omitempty"`
}

// CloudFrontLoggingList represents a list of CloudFrontLogging
type CloudFrontLoggingList []CloudFrontLogging

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudFrontLoggingList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudFrontLogging{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudFrontLoggingList{item}
		return nil
	}
	list := []CloudFrontLogging{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudFrontLoggingList(list)
		return nil
	}
	return err
}

// CloudFrontDistributionConfigOrigin represents CloudFront DistributionConfig Origin
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin.html
type CloudFrontDistributionConfigOrigin struct {
	// Origin information to specify a custom origin.
	CustomOriginConfig *CloudFrontDistributionConfigOriginCustomOrigin `json:"CustomOriginConfig,omitempty"`

	// The DNS name of the Amazon Simple Storage Service (S3) bucket or the
	// HTTP server from which you want CloudFront to get objects for this
	// origin.
	DomainName *StringExpr `json:"DomainName,omitempty"`

	// An identifier for the origin. The value of Id must be unique within
	// the distribution.
	Id *StringExpr `json:"Id,omitempty"`

	// The path that CloudFront uses to request content from an S3 bucket or
	// custom origin. The combination of the DomainName and OriginPath
	// properties must resolve to a valid path. The value must start with a
	// slash mark (/) and cannot end with a slash mark.
	OriginPath *StringExpr `json:"OriginPath,omitempty"`

	// Origin information to specify an S3 origin.
	S3OriginConfig *CloudFrontDistributionConfigOriginS3Origin `json:"S3OriginConfig,omitempty"`
}

// CloudFrontDistributionConfigOriginList represents a list of CloudFrontDistributionConfigOrigin
type CloudFrontDistributionConfigOriginList []CloudFrontDistributionConfigOrigin

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudFrontDistributionConfigOriginList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudFrontDistributionConfigOrigin{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudFrontDistributionConfigOriginList{item}
		return nil
	}
	list := []CloudFrontDistributionConfigOrigin{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudFrontDistributionConfigOriginList(list)
		return nil
	}
	return err
}

// CloudFrontDistributionConfigOriginCustomOrigin represents CloudFront DistributionConfig Origin CustomOrigin
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-customorigin.html
type CloudFrontDistributionConfigOriginCustomOrigin struct {
	// The HTTP port the custom origin listens on.
	HTTPPort *StringExpr `json:"HTTPPort,omitempty"`

	// The HTTPS port the custom origin listens on.
	HTTPSPort *StringExpr `json:"HTTPSPort,omitempty"`

	// The origin protocol policy to apply to your origin.
	OriginProtocolPolicy *StringExpr `json:"OriginProtocolPolicy,omitempty"`
}

// CloudFrontDistributionConfigOriginCustomOriginList represents a list of CloudFrontDistributionConfigOriginCustomOrigin
type CloudFrontDistributionConfigOriginCustomOriginList []CloudFrontDistributionConfigOriginCustomOrigin

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudFrontDistributionConfigOriginCustomOriginList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudFrontDistributionConfigOriginCustomOrigin{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudFrontDistributionConfigOriginCustomOriginList{item}
		return nil
	}
	list := []CloudFrontDistributionConfigOriginCustomOrigin{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudFrontDistributionConfigOriginCustomOriginList(list)
		return nil
	}
	return err
}

// CloudFrontDistributionConfigOriginS3Origin represents CloudFront DistributionConfig Origin S3Origin
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-s3origin.html
type CloudFrontDistributionConfigOriginS3Origin struct {
	// The CloudFront origin access identity to associate with the origin.
	// This is used to configure the origin so that end users can access
	// objects in an Amazon S3 bucket through CloudFront only.
	OriginAccessIdentity *StringExpr `json:"OriginAccessIdentity,omitempty"`
}

// CloudFrontDistributionConfigOriginS3OriginList represents a list of CloudFrontDistributionConfigOriginS3Origin
type CloudFrontDistributionConfigOriginS3OriginList []CloudFrontDistributionConfigOriginS3Origin

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudFrontDistributionConfigOriginS3OriginList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudFrontDistributionConfigOriginS3Origin{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudFrontDistributionConfigOriginS3OriginList{item}
		return nil
	}
	list := []CloudFrontDistributionConfigOriginS3Origin{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudFrontDistributionConfigOriginS3OriginList(list)
		return nil
	}
	return err
}

// CloudFrontDistributionConfigurationRestrictions represents CloudFront DistributionConfiguration Restrictions
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-restrictions.html
type CloudFrontDistributionConfigurationRestrictions struct {
	// The countries in which viewers are able to access your content.
	GeoRestriction *CloudFrontDistributionConfigRestrictionsGeoRestriction `json:"GeoRestriction,omitempty"`
}

// CloudFrontDistributionConfigurationRestrictionsList represents a list of CloudFrontDistributionConfigurationRestrictions
type CloudFrontDistributionConfigurationRestrictionsList []CloudFrontDistributionConfigurationRestrictions

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudFrontDistributionConfigurationRestrictionsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudFrontDistributionConfigurationRestrictions{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudFrontDistributionConfigurationRestrictionsList{item}
		return nil
	}
	list := []CloudFrontDistributionConfigurationRestrictions{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudFrontDistributionConfigurationRestrictionsList(list)
		return nil
	}
	return err
}

// CloudFrontDistributionConfigRestrictionsGeoRestriction represents CloudFront DistributionConfig Restrictions GeoRestriction
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-restrictions-georestriction.html
type CloudFrontDistributionConfigRestrictionsGeoRestriction struct {
	// The two-letter, uppercase country code for a country that you want to
	// include in your blacklist or whitelist.
	Locations *StringListExpr `json:"Locations,omitempty"`

	// The method to restrict distribution of your content:
	RestrictionType *StringExpr `json:"RestrictionType,omitempty"`

	// Prevents viewers in the countries that you specified from accessing
	// your content.
	Blacklist interface{} `json:"blacklist,omitempty"`

	// Allows viewers in the countries that you specified to access your
	// content.
	Whitelist interface{} `json:"whitelist,omitempty"`

	// No distribution restrictions by country.
	None interface{} `json:"none,omitempty"`
}

// CloudFrontDistributionConfigRestrictionsGeoRestrictionList represents a list of CloudFrontDistributionConfigRestrictionsGeoRestriction
type CloudFrontDistributionConfigRestrictionsGeoRestrictionList []CloudFrontDistributionConfigRestrictionsGeoRestriction

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudFrontDistributionConfigRestrictionsGeoRestrictionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudFrontDistributionConfigRestrictionsGeoRestriction{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudFrontDistributionConfigRestrictionsGeoRestrictionList{item}
		return nil
	}
	list := []CloudFrontDistributionConfigRestrictionsGeoRestriction{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudFrontDistributionConfigRestrictionsGeoRestrictionList(list)
		return nil
	}
	return err
}

// CloudFrontDistributionConfigurationViewerCertificate represents CloudFront DistributionConfiguration ViewerCertificate
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-viewercertificate.html
type CloudFrontDistributionConfigurationViewerCertificate struct {
	// Indicates whether to use the default certificate for your CloudFront
	// domain name when viewers use HTTPS to request your content.
	CloudFrontDefaultCertificate *BoolExpr `json:"CloudFrontDefaultCertificate,omitempty"`

	// The IAM certificate ID to use if you're using an alternate domain
	// name.
	IamCertificateId *StringExpr `json:"IamCertificateId,omitempty"`

	// The minimum version of the SSL protocol that you want CloudFront to
	// use for HTTPS connections. CloudFront serves your objects only to
	// browsers or devices that support at least the SSL version that you
	// specify.
	MinimumProtocolVersion *StringExpr `json:"MinimumProtocolVersion,omitempty"`

	// Specifies how CloudFront serves HTTPS requests.
	SslSupportMethod *StringExpr `json:"SslSupportMethod,omitempty"`
}

// CloudFrontDistributionConfigurationViewerCertificateList represents a list of CloudFrontDistributionConfigurationViewerCertificate
type CloudFrontDistributionConfigurationViewerCertificateList []CloudFrontDistributionConfigurationViewerCertificate

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudFrontDistributionConfigurationViewerCertificateList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudFrontDistributionConfigurationViewerCertificate{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudFrontDistributionConfigurationViewerCertificateList{item}
		return nil
	}
	list := []CloudFrontDistributionConfigurationViewerCertificate{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudFrontDistributionConfigurationViewerCertificateList(list)
		return nil
	}
	return err
}

// CloudFrontForwardedValues represents CloudFront ForwardedValues
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-forwardedvalues.html
type CloudFrontForwardedValues struct {
	// Forwards specified cookies to the origin of the cache behavior.
	Cookies *CloudFrontForwardedValuesCookies `json:"Cookies,omitempty"`

	// Specifies the headers that you want Amazon CloudFront to forward to
	// the origin for this cache behavior (whitelisted headers). For the
	// headers that you specify, Amazon CloudFront also caches separate
	// versions of a specified object that is based on the header values in
	// viewer requests.
	Headers *StringListExpr `json:"Headers,omitempty"`

	// Indicates whether you want CloudFront to forward query strings to the
	// origin that is associated with this cache behavior. If so, specify
	// true; if not, specify false.
	QueryString *BoolExpr `json:"QueryString,omitempty"`
}

// CloudFrontForwardedValuesList represents a list of CloudFrontForwardedValues
type CloudFrontForwardedValuesList []CloudFrontForwardedValues

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudFrontForwardedValuesList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudFrontForwardedValues{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudFrontForwardedValuesList{item}
		return nil
	}
	list := []CloudFrontForwardedValues{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudFrontForwardedValuesList(list)
		return nil
	}
	return err
}

// CloudFrontForwardedValuesCookies represents CloudFront ForwardedValues Cookies
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-forwardedvalues-cookies.html
type CloudFrontForwardedValuesCookies struct {
	// The cookies to forward to the origin of the cache behavior. You can
	// specify none, all, or whitelist.
	Forward *StringExpr `json:"Forward,omitempty"`

	// The names of cookies to forward to the origin for the cache behavior.
	WhitelistedNames *StringListExpr `json:"WhitelistedNames,omitempty"`
}

// CloudFrontForwardedValuesCookiesList represents a list of CloudFrontForwardedValuesCookies
type CloudFrontForwardedValuesCookiesList []CloudFrontForwardedValuesCookies

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudFrontForwardedValuesCookiesList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudFrontForwardedValuesCookies{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudFrontForwardedValuesCookiesList{item}
		return nil
	}
	list := []CloudFrontForwardedValuesCookies{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudFrontForwardedValuesCookiesList(list)
		return nil
	}
	return err
}

// CloudWatchMetricDimension represents CloudWatch Metric Dimension Property Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-dimension.html
type CloudWatchMetricDimension struct {
	// The name of the dimension, from 1–255 characters in length.
	Name *StringExpr `json:"Name,omitempty"`

	// The value representing the dimension measurement, from 1–255
	// characters in length.
	Value *StringExpr `json:"Value,omitempty"`
}

// CloudWatchMetricDimensionList represents a list of CloudWatchMetricDimension
type CloudWatchMetricDimensionList []CloudWatchMetricDimension

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudWatchMetricDimensionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudWatchMetricDimension{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudWatchMetricDimensionList{item}
		return nil
	}
	list := []CloudWatchMetricDimension{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudWatchMetricDimensionList(list)
		return nil
	}
	return err
}

// CloudWatchLogsMetricFilterMetricTransformationProperty represents CloudWatch Logs MetricFilter MetricTransformation Property
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-metricfilter-metrictransformation.html
type CloudWatchLogsMetricFilterMetricTransformationProperty struct {
	// The name of the CloudWatch metric to which the log information will be
	// published.
	MetricName *StringExpr `json:"MetricName,omitempty"`

	// The destination namespace of the CloudWatch metric. Namespaces are
	// containers for metrics. For example, you can add related metrics in
	// the same namespace.
	MetricNamespace *StringExpr `json:"MetricNamespace,omitempty"`

	// The value that is published to the CloudWatch metric. For example, if
	// you're counting the occurrences of a particular term like Error,
	// specify 1 for the metric value. If you're counting the number of bytes
	// transferred, reference the value that is in the log event by using $
	// followed by the name of the field that you specified in the filter
	// pattern, such as $size.
	MetricValue *StringExpr `json:"MetricValue,omitempty"`
}

// CloudWatchLogsMetricFilterMetricTransformationPropertyList represents a list of CloudWatchLogsMetricFilterMetricTransformationProperty
type CloudWatchLogsMetricFilterMetricTransformationPropertyList []CloudWatchLogsMetricFilterMetricTransformationProperty

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudWatchLogsMetricFilterMetricTransformationPropertyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudWatchLogsMetricFilterMetricTransformationProperty{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudWatchLogsMetricFilterMetricTransformationPropertyList{item}
		return nil
	}
	list := []CloudWatchLogsMetricFilterMetricTransformationProperty{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudWatchLogsMetricFilterMetricTransformationPropertyList(list)
		return nil
	}
	return err
}

// DataPipelinePipelineParameterObjects represents AWS Data Pipeline Pipeline ParameterObjects
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parameterobjects.html
type DataPipelinePipelineParameterObjects struct {
	// Key-value pairs that define the attributes of the parameter object.
	Attributes *DataPipelineParameterObjectsAttributes `json:"Attributes,omitempty"`

	// The identifier of the parameter object.
	Id *StringExpr `json:"Id,omitempty"`
}

// DataPipelinePipelineParameterObjectsList represents a list of DataPipelinePipelineParameterObjects
type DataPipelinePipelineParameterObjectsList []DataPipelinePipelineParameterObjects

// UnmarshalJSON sets the object from the provided JSON representation
func (l *DataPipelinePipelineParameterObjectsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := DataPipelinePipelineParameterObjects{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = DataPipelinePipelineParameterObjectsList{item}
		return nil
	}
	list := []DataPipelinePipelineParameterObjects{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = DataPipelinePipelineParameterObjectsList(list)
		return nil
	}
	return err
}

// DataPipelineParameterObjectsAttributes represents AWS Data Pipeline Parameter Objects Attributes
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parameterobjects-attributes.html
type DataPipelineParameterObjectsAttributes struct {
	// Specifies the name of a parameter attribute. To view parameter
	// attributes, see Creating a Pipeline Using Parameterized Templates in
	// the AWS Data Pipeline Developer Guide.
	Key *StringExpr `json:"Key,omitempty"`

	// A parameter attribute value.
	StringValue *StringExpr `json:"StringValue,omitempty"`
}

// DataPipelineParameterObjectsAttributesList represents a list of DataPipelineParameterObjectsAttributes
type DataPipelineParameterObjectsAttributesList []DataPipelineParameterObjectsAttributes

// UnmarshalJSON sets the object from the provided JSON representation
func (l *DataPipelineParameterObjectsAttributesList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := DataPipelineParameterObjectsAttributes{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = DataPipelineParameterObjectsAttributesList{item}
		return nil
	}
	list := []DataPipelineParameterObjectsAttributes{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = DataPipelineParameterObjectsAttributesList(list)
		return nil
	}
	return err
}

// DataPipelinePipelineParameterValues represents AWS Data Pipeline Pipeline ParameterValues
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parametervalues.html
type DataPipelinePipelineParameterValues struct {
	// The ID of a parameter object.
	Id *StringExpr `json:"Id,omitempty"`

	// A value to associate with the parameter object.
	StringValue *StringExpr `json:"StringValue,omitempty"`
}

// DataPipelinePipelineParameterValuesList represents a list of DataPipelinePipelineParameterValues
type DataPipelinePipelineParameterValuesList []DataPipelinePipelineParameterValues

// UnmarshalJSON sets the object from the provided JSON representation
func (l *DataPipelinePipelineParameterValuesList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := DataPipelinePipelineParameterValues{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = DataPipelinePipelineParameterValuesList{item}
		return nil
	}
	list := []DataPipelinePipelineParameterValues{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = DataPipelinePipelineParameterValuesList(list)
		return nil
	}
	return err
}

// DataPipelinePipelineObjects represents AWS Data Pipeline PipelineObjects
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobjects.html
type DataPipelinePipelineObjects struct {
	// Key-value pairs that define the properties of the object.
	Fields *DataPipelineDataPipelineObjectFields `json:"Fields,omitempty"`

	// Identifier of the object.
	Id *StringExpr `json:"Id,omitempty"`

	// Name of the object.
	Name *StringExpr `json:"Name,omitempty"`
}

// DataPipelinePipelineObjectsList represents a list of DataPipelinePipelineObjects
type DataPipelinePipelineObjectsList []DataPipelinePipelineObjects

// UnmarshalJSON sets the object from the provided JSON representation
func (l *DataPipelinePipelineObjectsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := DataPipelinePipelineObjects{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = DataPipelinePipelineObjectsList{item}
		return nil
	}
	list := []DataPipelinePipelineObjects{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = DataPipelinePipelineObjectsList(list)
		return nil
	}
	return err
}

// DataPipelineDataPipelineObjectFields represents AWS Data Pipeline Data Pipeline Object Fields
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobjects-fields.html
type DataPipelineDataPipelineObjectFields struct {
	// Specifies the name of a field for a particular object. To view fields
	// for a data pipeline object, see Pipeline Object Reference in the AWS
	// Data Pipeline Developer Guide.
	Key *StringExpr `json:"Key,omitempty"`

	// A field value that you specify as an identifier of another object in
	// the same pipeline definition.
	RefValue *StringExpr `json:"RefValue,omitempty"`

	// A field value that you specify as a string. To view valid values for a
	// particular field, see Pipeline Object Reference in the AWS Data
	// Pipeline Developer Guide.
	StringValue *StringExpr `json:"StringValue,omitempty"`
}

// DataPipelineDataPipelineObjectFieldsList represents a list of DataPipelineDataPipelineObjectFields
type DataPipelineDataPipelineObjectFieldsList []DataPipelineDataPipelineObjectFields

// UnmarshalJSON sets the object from the provided JSON representation
func (l *DataPipelineDataPipelineObjectFieldsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := DataPipelineDataPipelineObjectFields{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = DataPipelineDataPipelineObjectFieldsList{item}
		return nil
	}
	list := []DataPipelineDataPipelineObjectFields{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = DataPipelineDataPipelineObjectFieldsList(list)
		return nil
	}
	return err
}

// DataPipelinePipelinePipelineTags represents AWS Data Pipeline Pipeline PipelineTags
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelinetags.html
type DataPipelinePipelinePipelineTags struct {
	// The key name of a tag.
	Key *StringExpr `json:"Key,omitempty"`

	// The value to associate with the key name.
	Value *StringExpr `json:"Value,omitempty"`
}

// DataPipelinePipelinePipelineTagsList represents a list of DataPipelinePipelinePipelineTags
type DataPipelinePipelinePipelineTagsList []DataPipelinePipelinePipelineTags

// UnmarshalJSON sets the object from the provided JSON representation
func (l *DataPipelinePipelinePipelineTagsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := DataPipelinePipelinePipelineTags{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = DataPipelinePipelinePipelineTagsList{item}
		return nil
	}
	list := []DataPipelinePipelinePipelineTags{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = DataPipelinePipelinePipelineTagsList(list)
		return nil
	}
	return err
}

// DynamoDBAttributeDefinitions represents DynamoDB Attribute Definitions
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-attributedef.html
type DynamoDBAttributeDefinitions struct {
	// The name of an attribute. Attribute names can be 1 – 255 characters
	// long and have no character restrictions.
	AttributeName *StringExpr `json:"AttributeName,omitempty"`

	// The data type for the attribute. You can specify S for string data, N
	// for numeric data, or B for binary data.
	AttributeType *StringExpr `json:"AttributeType,omitempty"`
}

// DynamoDBAttributeDefinitionsList represents a list of DynamoDBAttributeDefinitions
type DynamoDBAttributeDefinitionsList []DynamoDBAttributeDefinitions

// UnmarshalJSON sets the object from the provided JSON representation
func (l *DynamoDBAttributeDefinitionsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := DynamoDBAttributeDefinitions{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = DynamoDBAttributeDefinitionsList{item}
		return nil
	}
	list := []DynamoDBAttributeDefinitions{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = DynamoDBAttributeDefinitionsList(list)
		return nil
	}
	return err
}

// DynamoDBGlobalSecondaryIndexes represents DynamoDB Global Secondary Indexes
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html
type DynamoDBGlobalSecondaryIndexes struct {
	// The name of the global secondary index. The index name can be 3 –
	// 255 characters long and have no character restrictions.
	IndexName *StringExpr `json:"IndexName,omitempty"`

	// The complete index key schema for the global secondary index, which
	// consists of one or more pairs of attribute names and key types.
	KeySchema *DynamoDBKeySchema `json:"KeySchema,omitempty"`

	// Attributes that are copied (projected) from the source table into the
	// index. These attributes are in addition to the primary key attributes
	// and index key attributes, which are automatically projected.
	Projection *DynamoDBProjectionObject `json:"Projection,omitempty"`

	// The provisioned throughput settings for the index.
	ProvisionedThroughput *DynamoDBProvisionedThroughput `json:"ProvisionedThroughput,omitempty"`
}

// DynamoDBGlobalSecondaryIndexesList represents a list of DynamoDBGlobalSecondaryIndexes
type DynamoDBGlobalSecondaryIndexesList []DynamoDBGlobalSecondaryIndexes

// UnmarshalJSON sets the object from the provided JSON representation
func (l *DynamoDBGlobalSecondaryIndexesList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := DynamoDBGlobalSecondaryIndexes{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = DynamoDBGlobalSecondaryIndexesList{item}
		return nil
	}
	list := []DynamoDBGlobalSecondaryIndexes{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = DynamoDBGlobalSecondaryIndexesList(list)
		return nil
	}
	return err
}

// DynamoDBKeySchema represents DynamoDB Key Schema
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-keyschema.html
type DynamoDBKeySchema struct {
	// The attribute name that is used as the primary key for this table.
	// Primary key element names can be 1 – 255 characters long and have no
	// character restrictions.
	AttributeName *StringExpr `json:"AttributeName,omitempty"`

	// Represents the attribute data, consisting of the data type and the
	// attribute value itself. You can specify HASH or RANGE.
	KeyType *StringExpr `json:"KeyType,omitempty"`
}

// DynamoDBKeySchemaList represents a list of DynamoDBKeySchema
type DynamoDBKeySchemaList []DynamoDBKeySchema

// UnmarshalJSON sets the object from the provided JSON representation
func (l *DynamoDBKeySchemaList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := DynamoDBKeySchema{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = DynamoDBKeySchemaList{item}
		return nil
	}
	list := []DynamoDBKeySchema{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = DynamoDBKeySchemaList(list)
		return nil
	}
	return err
}

// DynamoDBLocalSecondaryIndexes represents DynamoDB Local Secondary Indexes
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-lsi.html
type DynamoDBLocalSecondaryIndexes struct {
	// The name of the local secondary index. The index name can be 3 – 255
	// characters long and have no character restrictions.
	IndexName *StringExpr `json:"IndexName,omitempty"`

	// The complete index key schema for the local secondary index, which
	// consists of one or more pairs of attribute names and key types. For
	// local secondary indexes, the hash key must be the same as that of the
	// source table.
	KeySchema *DynamoDBKeySchema `json:"KeySchema,omitempty"`

	// Attributes that are copied (projected) from the source table into the
	// index. These attributes are additions to the primary key attributes
	// and index key attributes, which are automatically projected.
	Projection *DynamoDBProjectionObject `json:"Projection,omitempty"`
}

// DynamoDBLocalSecondaryIndexesList represents a list of DynamoDBLocalSecondaryIndexes
type DynamoDBLocalSecondaryIndexesList []DynamoDBLocalSecondaryIndexes

// UnmarshalJSON sets the object from the provided JSON representation
func (l *DynamoDBLocalSecondaryIndexesList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := DynamoDBLocalSecondaryIndexes{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = DynamoDBLocalSecondaryIndexesList{item}
		return nil
	}
	list := []DynamoDBLocalSecondaryIndexes{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = DynamoDBLocalSecondaryIndexesList(list)
		return nil
	}
	return err
}

// DynamoDBProjectionObject represents DynamoDB Projection Object
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-projectionobject.html
type DynamoDBProjectionObject struct {
	// The non-key attribute names that are projected into the index.
	NonKeyAttributes *StringListExpr `json:"NonKeyAttributes,omitempty"`

	// The set of attributes that are projected into the index:
	ProjectionType *StringExpr `json:"ProjectionType,omitempty"`

	// Only the index and primary keys are projected into the index.
	KEYSXONLY interface{} `json:"KEYS_ONLY,omitempty"`

	// Only the specified table attributes are projected into the index. The
	// list of projected attributes are in NonKeyAttributes.
	INCLUDE interface{} `json:"INCLUDE,omitempty"`

	// All of the table attributes are projected into the index.
	ALL interface{} `json:"ALL,omitempty"`
}

// DynamoDBProjectionObjectList represents a list of DynamoDBProjectionObject
type DynamoDBProjectionObjectList []DynamoDBProjectionObject

// UnmarshalJSON sets the object from the provided JSON representation
func (l *DynamoDBProjectionObjectList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := DynamoDBProjectionObject{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = DynamoDBProjectionObjectList{item}
		return nil
	}
	list := []DynamoDBProjectionObject{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = DynamoDBProjectionObjectList(list)
		return nil
	}
	return err
}

// DynamoDBProvisionedThroughput represents DynamoDB Provisioned Throughput
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-provisionedthroughput.html
type DynamoDBProvisionedThroughput struct {
}

// DynamoDBProvisionedThroughputList represents a list of DynamoDBProvisionedThroughput
type DynamoDBProvisionedThroughputList []DynamoDBProvisionedThroughput

// UnmarshalJSON sets the object from the provided JSON representation
func (l *DynamoDBProvisionedThroughputList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := DynamoDBProvisionedThroughput{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = DynamoDBProvisionedThroughputList{item}
		return nil
	}
	list := []DynamoDBProvisionedThroughput{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = DynamoDBProvisionedThroughputList(list)
		return nil
	}
	return err
}

// EC2BlockDeviceMappingProperty represents Amazon EC2 Block Device Mapping Property
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-mapping.html
type EC2BlockDeviceMappingProperty struct {
	// The name of the device within Amazon EC2.
	DeviceName *StringExpr `json:"DeviceName,omitempty"`

	// Required: Conditional You can specify either VirtualName or Ebs, but
	// not both.
	Ebs *ElasticBlockStoreBlockDeviceProperty `json:"Ebs,omitempty"`

	// This property can be used to unmap a defined device.
	NoDevice interface{} `json:"NoDevice,omitempty"`

	// The name of the virtual device. The name must be in the form
	// ephemeralX where X is a number starting from zero (0); for example,
	// ephemeral0.
	VirtualName *StringExpr `json:"VirtualName,omitempty"`
}

// EC2BlockDeviceMappingPropertyList represents a list of EC2BlockDeviceMappingProperty
type EC2BlockDeviceMappingPropertyList []EC2BlockDeviceMappingProperty

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2BlockDeviceMappingPropertyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2BlockDeviceMappingProperty{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2BlockDeviceMappingPropertyList{item}
		return nil
	}
	list := []EC2BlockDeviceMappingProperty{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2BlockDeviceMappingPropertyList(list)
		return nil
	}
	return err
}

// ElasticBlockStoreBlockDeviceProperty represents Amazon Elastic Block Store Block Device Property
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-template.html
type ElasticBlockStoreBlockDeviceProperty struct {
	// Determines whether to delete the volume on instance termination. The
	// default value is true.
	DeleteOnTermination *BoolExpr `json:"DeleteOnTermination,omitempty"`

	// Indicates whether the volume is encrypted. Encrypted Amazon EBS
	// volumes can only be attached to instance types that support Amazon EBS
	// encryption. Volumes that are created from encrypted snapshots are
	// automatically encrypted. You cannot create an encrypted volume from an
	// unencrypted snapshot or vice versa. If your AMI uses encrypted
	// volumes, you can only launch the AMI on supported instance types. For
	// more information, see Amazon EBS encryption in the Amazon EC2 User
	// Guide for Linux Instances.
	Encrypted *BoolExpr `json:"Encrypted,omitempty"`

	// The number of I/O operations per second (IOPS) that the volume
	// supports. This can be an integer from 100 – 2000.
	Iops *IntegerExpr `json:"Iops,omitempty"`

	// The snapshot ID of the volume to use to create a block device.
	SnapshotId *StringExpr `json:"SnapshotId,omitempty"`

	// The volume size, in gibibytes (GiB). This can be a number from 1 –
	// 1024. If the volume type is io1, the minimum value is 10.
	VolumeSize *StringExpr `json:"VolumeSize,omitempty"`

	// The volume type. You can specify standard, io1, or gp2. If you set the
	// type to io1, you must also set the Iops property. For more information
	// about these values and the default value, see CreateVolume in the
	// Amazon EC2 API Reference.
	VolumeType *StringExpr `json:"VolumeType,omitempty"`
}

// ElasticBlockStoreBlockDevicePropertyList represents a list of ElasticBlockStoreBlockDeviceProperty
type ElasticBlockStoreBlockDevicePropertyList []ElasticBlockStoreBlockDeviceProperty

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticBlockStoreBlockDevicePropertyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticBlockStoreBlockDeviceProperty{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticBlockStoreBlockDevicePropertyList{item}
		return nil
	}
	list := []ElasticBlockStoreBlockDeviceProperty{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticBlockStoreBlockDevicePropertyList(list)
		return nil
	}
	return err
}

// EC2ICMP represents EC2 ICMP Property Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-icmp.html
type EC2ICMP struct {
}

// EC2ICMPList represents a list of EC2ICMP
type EC2ICMPList []EC2ICMP

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2ICMPList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2ICMP{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2ICMPList{item}
		return nil
	}
	list := []EC2ICMP{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2ICMPList(list)
		return nil
	}
	return err
}

// EC2MountPoint represents EC2 MountPoint Property Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-mount-point.html
type EC2MountPoint struct {
	// How the device is exposed to the instance (such as /dev/sdh, or xvdh).
	Device *StringExpr `json:"Device,omitempty"`

	// The ID of the Amazon EBS volume. The volume and instance must be
	// within the same Availability Zone and the instance must be running.
	VolumeId *StringExpr `json:"VolumeId,omitempty"`
}

// EC2MountPointList represents a list of EC2MountPoint
type EC2MountPointList []EC2MountPoint

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2MountPointList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2MountPoint{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2MountPointList{item}
		return nil
	}
	list := []EC2MountPoint{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2MountPointList(list)
		return nil
	}
	return err
}

// EC2NetworkInterfaceEmbedded represents EC2 NetworkInterface Embedded Property Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html
type EC2NetworkInterfaceEmbedded struct {
	// Indicates whether the network interface receives a public IP address.
	// You can associate a public IP address with a network interface only if
	// it has a device index of eth0 and if it is a new network interface
	// (not an existing one). In other words, if you specify true, don't
	// specify a network interface ID. For more information, see Amazon EC2
	// Instance IP Addressing.
	AssociatePublicIpAddress *BoolExpr `json:"AssociatePublicIpAddress,omitempty"`

	// Whether to delete the network interface when the instance terminates.
	DeleteOnTermination *BoolExpr `json:"DeleteOnTermination,omitempty"`

	// The description of this network interface.
	Description *StringExpr `json:"Description,omitempty"`

	// The network interface's position in the attachment order.
	DeviceIndex *StringExpr `json:"DeviceIndex,omitempty"`

	// A list of security group IDs associated with this network interface.
	GroupSet *StringListExpr `json:"GroupSet,omitempty"`

	// An existing network interface ID.
	NetworkInterfaceId *StringExpr `json:"NetworkInterfaceId,omitempty"`

	// Assigns a single private IP address to the network interface, which is
	// used as the primary private IP address. If you want to specify
	// multiple private IP address, use the PrivateIpAddresses property.
	PrivateIpAddress *StringExpr `json:"PrivateIpAddress,omitempty"`

	// Assigns a list of private IP addresses to the network interface. You
	// can specify a primary private IP address by setting the value of the
	// Primary property to true in the PrivateIpAddressSpecification
	// property. If you want Amazon EC2 to automatically assign private IP
	// addresses, use the SecondaryPrivateIpCount property and do not specify
	// this property.
	PrivateIpAddresses *EC2NetworkInterfacePrivateIPSpecificationList `json:"PrivateIpAddresses,omitempty"`

	// The number of secondary private IP addresses that Amazon EC2 auto
	// assigns to the network interface. Amazon EC2 uses the value of the
	// PrivateIpAddress property as the primary private IP address. If you
	// don't specify that property, Amazon EC2 auto assigns both the primary
	// and secondary private IP addresses.
	SecondaryPrivateIpAddressCount *IntegerExpr `json:"SecondaryPrivateIpAddressCount,omitempty"`

	// The ID of the subnet to associate with the network interface.
	SubnetId *StringExpr `json:"SubnetId,omitempty"`
}

// EC2NetworkInterfaceEmbeddedList represents a list of EC2NetworkInterfaceEmbedded
type EC2NetworkInterfaceEmbeddedList []EC2NetworkInterfaceEmbedded

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2NetworkInterfaceEmbeddedList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2NetworkInterfaceEmbedded{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2NetworkInterfaceEmbeddedList{item}
		return nil
	}
	list := []EC2NetworkInterfaceEmbedded{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2NetworkInterfaceEmbeddedList(list)
		return nil
	}
	return err
}

// EC2NetworkInterfaceAssociation represents EC2 Network Interface Association
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-interface-association.html
type EC2NetworkInterfaceAssociation struct {
	// The ID of the network interface attachment.
	AttachmentID *StringExpr `json:"AttachmentID,omitempty"`

	// The ID of the instance attached to the network interface.
	InstanceID *StringExpr `json:"InstanceID,omitempty"`

	// The address of the Elastic IP address bound to the network interface.
	PublicIp *StringExpr `json:"PublicIp,omitempty"`

	// The ID of the Elastic IP address owner.
	IpOwnerId *StringExpr `json:"IpOwnerId,omitempty"`
}

// EC2NetworkInterfaceAssociationList represents a list of EC2NetworkInterfaceAssociation
type EC2NetworkInterfaceAssociationList []EC2NetworkInterfaceAssociation

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2NetworkInterfaceAssociationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2NetworkInterfaceAssociation{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2NetworkInterfaceAssociationList{item}
		return nil
	}
	list := []EC2NetworkInterfaceAssociation{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2NetworkInterfaceAssociationList(list)
		return nil
	}
	return err
}

// EC2NetworkInterfaceAttachmentType represents EC2 Network Interface Attachment
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-interface-attachment.html
type EC2NetworkInterfaceAttachmentType struct {
	// The ID of the network interface attachment.
	AttachmentID *StringExpr `json:"AttachmentID,omitempty"`

	// The ID of the instance attached to the network interface.
	InstanceID *StringExpr `json:"InstanceID,omitempty"`
}

// EC2NetworkInterfaceAttachmentTypeList represents a list of EC2NetworkInterfaceAttachmentType
type EC2NetworkInterfaceAttachmentTypeList []EC2NetworkInterfaceAttachmentType

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2NetworkInterfaceAttachmentTypeList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2NetworkInterfaceAttachmentType{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2NetworkInterfaceAttachmentTypeList{item}
		return nil
	}
	list := []EC2NetworkInterfaceAttachmentType{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2NetworkInterfaceAttachmentTypeList(list)
		return nil
	}
	return err
}

// EC2NetworkInterfaceGroupItem represents EC2 Network Interface Group Item
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-interface-groupitem.html
type EC2NetworkInterfaceGroupItem struct {
	// ID of the security group.
	Key *StringExpr `json:"Key,omitempty"`

	// Name of the security group.
	Value *StringExpr `json:"Value,omitempty"`
}

// EC2NetworkInterfaceGroupItemList represents a list of EC2NetworkInterfaceGroupItem
type EC2NetworkInterfaceGroupItemList []EC2NetworkInterfaceGroupItem

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2NetworkInterfaceGroupItemList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2NetworkInterfaceGroupItem{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2NetworkInterfaceGroupItemList{item}
		return nil
	}
	list := []EC2NetworkInterfaceGroupItem{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2NetworkInterfaceGroupItemList(list)
		return nil
	}
	return err
}

// EC2NetworkInterfacePrivateIPSpecification represents EC2 Network Interface Private IP Specification
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-interface-privateipspec.html
type EC2NetworkInterfacePrivateIPSpecification struct {
	// The private IP address of the network interface.
	PrivateIpAddress *StringExpr `json:"PrivateIpAddress,omitempty"`

	// Sets the private IP address as the primary private address. You can
	// set only one primary private IP address. If you don't specify a
	// primary private IP address, Amazon EC2 automatically assigns a primary
	// private IP address.
	Primary *BoolExpr `json:"Primary,omitempty"`
}

// EC2NetworkInterfacePrivateIPSpecificationList represents a list of EC2NetworkInterfacePrivateIPSpecification
type EC2NetworkInterfacePrivateIPSpecificationList []EC2NetworkInterfacePrivateIPSpecification

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2NetworkInterfacePrivateIPSpecificationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2NetworkInterfacePrivateIPSpecification{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2NetworkInterfacePrivateIPSpecificationList{item}
		return nil
	}
	list := []EC2NetworkInterfacePrivateIPSpecification{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2NetworkInterfacePrivateIPSpecificationList(list)
		return nil
	}
	return err
}

// EC2PortRange represents EC2 PortRange Property Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-port-range.html
type EC2PortRange struct {
}

// EC2PortRangeList represents a list of EC2PortRange
type EC2PortRangeList []EC2PortRange

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2PortRangeList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2PortRange{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2PortRangeList{item}
		return nil
	}
	list := []EC2PortRange{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2PortRangeList(list)
		return nil
	}
	return err
}

// EC2SecurityGroupRule represents EC2 Security Group Rule Property Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-rule.html
type EC2SecurityGroupRule struct {
	// Specifies a CIDR range.
	CidrIp *StringExpr `json:"CidrIp,omitempty"`

	// Specifies the GroupId of the destination Amazon VPC security group.
	DestinationSecurityGroupIdXXSecurityGroupEgressXOnlyX *StringExpr `json:"DestinationSecurityGroupId (SecurityGroupEgress only),omitempty"`

	// The start of port range for the TCP and UDP protocols, or an ICMP type
	// number. An ICMP type number of -1 indicates a wildcard (i.e., any ICMP
	// type number).
	FromPort *IntegerExpr `json:"FromPort,omitempty"`

	// An IP protocol name or number. For valid values, go to the IpProtocol
	// parameter in AuthorizeSecurityGroupIngress
	IpProtocol *StringExpr `json:"IpProtocol,omitempty"`

	// For VPC security groups only. Specifies the ID of the Amazon EC2
	// Security Group to allow access. You can use the Ref intrinsic function
	// to refer to the logical ID of a security group defined in the same
	// template.
	SourceSecurityGroupIdXXSecurityGroupIngressXOnlyX *StringExpr `json:"SourceSecurityGroupId (SecurityGroupIngress only),omitempty"`

	// For non-VPC security groups only. Specifies the name of the Amazon EC2
	// Security Group to use for access. You can use the Ref intrinsic
	// function to refer to the logical name of a security group that is
	// defined in the same template.
	SourceSecurityGroupNameXXSecurityGroupIngressXOnlyX *StringExpr `json:"SourceSecurityGroupName (SecurityGroupIngress only),omitempty"`

	// Specifies the AWS Account ID of the owner of the Amazon EC2 Security
	// Group that is specified in the SourceSecurityGroupName property.
	SourceSecurityGroupOwnerIdXXSecurityGroupIngressXOnlyX *StringExpr `json:"SourceSecurityGroupOwnerId (SecurityGroupIngress only),omitempty"`

	// The end of port range for the TCP and UDP protocols, or an ICMP code.
	// An ICMP code of -1 indicates a wildcard (i.e., any ICMP code).
	ToPort *IntegerExpr `json:"ToPort,omitempty"`
}

// EC2SecurityGroupRuleList represents a list of EC2SecurityGroupRule
type EC2SecurityGroupRuleList []EC2SecurityGroupRule

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2SecurityGroupRuleList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2SecurityGroupRule{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2SecurityGroupRuleList{item}
		return nil
	}
	list := []EC2SecurityGroupRule{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2SecurityGroupRuleList(list)
		return nil
	}
	return err
}

// EC2ContainerServiceServiceLoadBalancers represents Amazon EC2 Container Service Service LoadBalancers
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-loadbalancers.html
type EC2ContainerServiceServiceLoadBalancers struct {
	// The name of a container to use with the load balancer.
	ContainerName *StringExpr `json:"ContainerName,omitempty"`

	// The port number on the container to direct load balancer traffic to.
	// Your container instances must allow ingress traffic on this port.
	ContainerPort *IntegerExpr `json:"ContainerPort,omitempty"`

	// The name of the load balancer to associated with the Amazon ECS
	// service.
	LoadBalancerName *StringExpr `json:"LoadBalancerName,omitempty"`
}

// EC2ContainerServiceServiceLoadBalancersList represents a list of EC2ContainerServiceServiceLoadBalancers
type EC2ContainerServiceServiceLoadBalancersList []EC2ContainerServiceServiceLoadBalancers

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2ContainerServiceServiceLoadBalancersList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2ContainerServiceServiceLoadBalancers{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2ContainerServiceServiceLoadBalancersList{item}
		return nil
	}
	list := []EC2ContainerServiceServiceLoadBalancers{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2ContainerServiceServiceLoadBalancersList(list)
		return nil
	}
	return err
}

// EC2ContainerServiceTaskDefinitionContainerDefinitions represents Amazon EC2 Container Service TaskDefinition ContainerDefinitions
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html
type EC2ContainerServiceTaskDefinitionContainerDefinitions struct {
	// The CMD value to pass to the container. For more information about the
	// Docker CMD parameter, see
	// https://docs.docker.com/reference/builder/#cmd.
	Command *StringListExpr `json:"Command,omitempty"`

	// The minimum number of CPU units to reserve for the container.
	// Containers share unallocated CPU units with other containers on the
	// instance by using the same ratio as their allocated CPU units.
	Cpu *IntegerExpr `json:"Cpu,omitempty"`

	// The ENTRYPOINT value to pass to the container. For more information
	// about the Docker ENTRYPOINT parameter, see
	// https://docs.docker.com/reference/builder/#entrypoint.
	EntryPoint *StringListExpr `json:"EntryPoint,omitempty"`

	// The environment variables to pass to the container.
	Environment *EC2ContainerServiceTaskDefinitionContainerDefinitionsEnvironmentList `json:"Environment,omitempty"`

	// Indicates whether the task stops if this container fails. If you
	// specify true and the container fails, all other containers in the task
	// stop. If you specify false and the container fails, none of the other
	// containers in the task are affected. This value is true by default.
	Essential *BoolExpr `json:"Essential,omitempty"`

	// The image to use for a container, which is passed directly to the
	// Docker daemon. You can use images in the Docker Hub registry or
	// specify other repositories (repository-url/image:tag).
	Image *StringExpr `json:"Image,omitempty"`

	// The name of another container to connect to. With links, containers
	// can communicate with each other without using port mappings.
	Links *StringListExpr `json:"Links,omitempty"`

	// The number of MiB of memory to reserve for the container. If your
	// container attempts to exceed the allocated memory, the container is
	// killed.
	Memory *IntegerExpr `json:"Memory,omitempty"`

	// The mount points for data volumes in the container.
	MountPoints *EC2ContainerServiceTaskDefinitionContainerDefinitionsMountPointsList `json:"MountPoints,omitempty"`

	// A name for the container.
	Name *StringExpr `json:"Name,omitempty"`

	// A mapping of the container port to a host port. Port mappings enable
	// containers to access ports on the host container instance to send or
	// receive traffic.
	PortMappings *EC2ContainerServiceTaskDefinitionContainerDefinitionsPortMappingsList `json:"PortMappings,omitempty"`

	// Data volumes to mount from another container.
	VolumesFrom *EC2ContainerServiceTaskDefinitionContainerDefinitionsVolumesFromList `json:"VolumesFrom,omitempty"`
}

// EC2ContainerServiceTaskDefinitionContainerDefinitionsList represents a list of EC2ContainerServiceTaskDefinitionContainerDefinitions
type EC2ContainerServiceTaskDefinitionContainerDefinitionsList []EC2ContainerServiceTaskDefinitionContainerDefinitions

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2ContainerServiceTaskDefinitionContainerDefinitionsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2ContainerServiceTaskDefinitionContainerDefinitions{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2ContainerServiceTaskDefinitionContainerDefinitionsList{item}
		return nil
	}
	list := []EC2ContainerServiceTaskDefinitionContainerDefinitions{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2ContainerServiceTaskDefinitionContainerDefinitionsList(list)
		return nil
	}
	return err
}

// EC2ContainerServiceTaskDefinitionContainerDefinitionsEnvironment represents Amazon EC2 Container Service TaskDefinition ContainerDefinitions Environment
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-environment.html
type EC2ContainerServiceTaskDefinitionContainerDefinitionsEnvironment struct {
	// The name of the environment variable.
	Name *StringExpr `json:"Name,omitempty"`

	// The value of the environment variable.
	Value *StringExpr `json:"Value,omitempty"`
}

// EC2ContainerServiceTaskDefinitionContainerDefinitionsEnvironmentList represents a list of EC2ContainerServiceTaskDefinitionContainerDefinitionsEnvironment
type EC2ContainerServiceTaskDefinitionContainerDefinitionsEnvironmentList []EC2ContainerServiceTaskDefinitionContainerDefinitionsEnvironment

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2ContainerServiceTaskDefinitionContainerDefinitionsEnvironmentList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2ContainerServiceTaskDefinitionContainerDefinitionsEnvironment{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2ContainerServiceTaskDefinitionContainerDefinitionsEnvironmentList{item}
		return nil
	}
	list := []EC2ContainerServiceTaskDefinitionContainerDefinitionsEnvironment{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2ContainerServiceTaskDefinitionContainerDefinitionsEnvironmentList(list)
		return nil
	}
	return err
}

// EC2ContainerServiceTaskDefinitionContainerDefinitionsMountPoints represents Amazon EC2 Container Service TaskDefinition ContainerDefinitions MountPoints
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-mountpoints.html
type EC2ContainerServiceTaskDefinitionContainerDefinitionsMountPoints struct {
	// The path on the container that indicates where you want to mount the
	// volume.
	ContainerPath *StringExpr `json:"ContainerPath,omitempty"`

	// The name of the volume to mount.
	SourceVolume *StringExpr `json:"SourceVolume,omitempty"`

	// Indicates whether the container can write to the volume. If you
	// specify true, the container has read-only access to the volume. If you
	// specify false, the container can write to the volume. By default, the
	// value is false.
	ReadOnly *BoolExpr `json:"ReadOnly,omitempty"`
}

// EC2ContainerServiceTaskDefinitionContainerDefinitionsMountPointsList represents a list of EC2ContainerServiceTaskDefinitionContainerDefinitionsMountPoints
type EC2ContainerServiceTaskDefinitionContainerDefinitionsMountPointsList []EC2ContainerServiceTaskDefinitionContainerDefinitionsMountPoints

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2ContainerServiceTaskDefinitionContainerDefinitionsMountPointsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2ContainerServiceTaskDefinitionContainerDefinitionsMountPoints{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2ContainerServiceTaskDefinitionContainerDefinitionsMountPointsList{item}
		return nil
	}
	list := []EC2ContainerServiceTaskDefinitionContainerDefinitionsMountPoints{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2ContainerServiceTaskDefinitionContainerDefinitionsMountPointsList(list)
		return nil
	}
	return err
}

// EC2ContainerServiceTaskDefinitionContainerDefinitionsPortMappings represents Amazon EC2 Container Service TaskDefinition ContainerDefinitions PortMappings
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-portmappings.html
type EC2ContainerServiceTaskDefinitionContainerDefinitionsPortMappings struct {
	// The port number on the container that is bound to the host port.
	ContainerPort *IntegerExpr `json:"ContainerPort,omitempty"`

	// The host port number on the container instance that you want to
	// reserve for your container. You can specify a non-reserved host port
	// for your container port mapping, or you can omit the host port (or set
	// it to 0). If you specify a container port but no host port, your
	// container port is automatically assigned a host port in the 49153 to
	// 65535 port range.
	HostPort *IntegerExpr `json:"HostPort,omitempty"`
}

// EC2ContainerServiceTaskDefinitionContainerDefinitionsPortMappingsList represents a list of EC2ContainerServiceTaskDefinitionContainerDefinitionsPortMappings
type EC2ContainerServiceTaskDefinitionContainerDefinitionsPortMappingsList []EC2ContainerServiceTaskDefinitionContainerDefinitionsPortMappings

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2ContainerServiceTaskDefinitionContainerDefinitionsPortMappingsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2ContainerServiceTaskDefinitionContainerDefinitionsPortMappings{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2ContainerServiceTaskDefinitionContainerDefinitionsPortMappingsList{item}
		return nil
	}
	list := []EC2ContainerServiceTaskDefinitionContainerDefinitionsPortMappings{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2ContainerServiceTaskDefinitionContainerDefinitionsPortMappingsList(list)
		return nil
	}
	return err
}

// EC2ContainerServiceTaskDefinitionContainerDefinitionsVolumesFrom represents Amazon EC2 Container Service TaskDefinition ContainerDefinitions VolumesFrom
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-volumesfrom.html
type EC2ContainerServiceTaskDefinitionContainerDefinitionsVolumesFrom struct {
	// The name of the container that has the volumes to mount.
	SourceContainer *StringExpr `json:"SourceContainer,omitempty"`

	// Indicates whether the container can write to the volume. If you
	// specify true, the container has read-only access to the volume. If you
	// specify false, the container can write to the volume. By default, the
	// value is false.
	ReadOnly *BoolExpr `json:"ReadOnly,omitempty"`
}

// EC2ContainerServiceTaskDefinitionContainerDefinitionsVolumesFromList represents a list of EC2ContainerServiceTaskDefinitionContainerDefinitionsVolumesFrom
type EC2ContainerServiceTaskDefinitionContainerDefinitionsVolumesFromList []EC2ContainerServiceTaskDefinitionContainerDefinitionsVolumesFrom

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2ContainerServiceTaskDefinitionContainerDefinitionsVolumesFromList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2ContainerServiceTaskDefinitionContainerDefinitionsVolumesFrom{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2ContainerServiceTaskDefinitionContainerDefinitionsVolumesFromList{item}
		return nil
	}
	list := []EC2ContainerServiceTaskDefinitionContainerDefinitionsVolumesFrom{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2ContainerServiceTaskDefinitionContainerDefinitionsVolumesFromList(list)
		return nil
	}
	return err
}

// EC2ContainerServiceTaskDefinitionVolumes represents Amazon EC2 Container Service TaskDefinition Volumes
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volumes.html
type EC2ContainerServiceTaskDefinitionVolumes struct {
	// The name of the volume. To specify mount points in your container
	// definitions, use the value of this property.
	Name *StringExpr `json:"Name,omitempty"`

	// Determines whether your data volume persists on the host container
	// instance and at the location where it is stored.
	Host *EC2ContainerServiceTaskDefinitionVolumesHost `json:"Host,omitempty"`
}

// EC2ContainerServiceTaskDefinitionVolumesList represents a list of EC2ContainerServiceTaskDefinitionVolumes
type EC2ContainerServiceTaskDefinitionVolumesList []EC2ContainerServiceTaskDefinitionVolumes

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2ContainerServiceTaskDefinitionVolumesList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2ContainerServiceTaskDefinitionVolumes{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2ContainerServiceTaskDefinitionVolumesList{item}
		return nil
	}
	list := []EC2ContainerServiceTaskDefinitionVolumes{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2ContainerServiceTaskDefinitionVolumesList(list)
		return nil
	}
	return err
}

// EC2ContainerServiceTaskDefinitionVolumesHost represents Amazon EC2 Container Service TaskDefinition Volumes Host
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volumes-host.html
type EC2ContainerServiceTaskDefinitionVolumesHost struct {
	// The data volume path on the host container instance.
	SourcePath *StringExpr `json:"SourcePath,omitempty"`
}

// EC2ContainerServiceTaskDefinitionVolumesHostList represents a list of EC2ContainerServiceTaskDefinitionVolumesHost
type EC2ContainerServiceTaskDefinitionVolumesHostList []EC2ContainerServiceTaskDefinitionVolumesHost

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2ContainerServiceTaskDefinitionVolumesHostList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2ContainerServiceTaskDefinitionVolumesHost{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2ContainerServiceTaskDefinitionVolumesHostList{item}
		return nil
	}
	list := []EC2ContainerServiceTaskDefinitionVolumesHost{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2ContainerServiceTaskDefinitionVolumesHostList(list)
		return nil
	}
	return err
}

// ElasticBeanstalkEnvironmentTier represents Elastic Beanstalk Environment Tier Property Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment-tier.html
type ElasticBeanstalkEnvironmentTier struct {
}

// ElasticBeanstalkEnvironmentTierList represents a list of ElasticBeanstalkEnvironmentTier
type ElasticBeanstalkEnvironmentTierList []ElasticBeanstalkEnvironmentTier

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticBeanstalkEnvironmentTierList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticBeanstalkEnvironmentTier{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticBeanstalkEnvironmentTierList{item}
		return nil
	}
	list := []ElasticBeanstalkEnvironmentTier{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticBeanstalkEnvironmentTierList(list)
		return nil
	}
	return err
}

// ElasticBeanstalkOptionSettings represents Elastic Beanstalk OptionSettings Property Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-option-settings.html
type ElasticBeanstalkOptionSettings struct {
}

// ElasticBeanstalkOptionSettingsList represents a list of ElasticBeanstalkOptionSettings
type ElasticBeanstalkOptionSettingsList []ElasticBeanstalkOptionSettings

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticBeanstalkOptionSettingsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticBeanstalkOptionSettings{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticBeanstalkOptionSettingsList{item}
		return nil
	}
	list := []ElasticBeanstalkOptionSettings{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticBeanstalkOptionSettingsList(list)
		return nil
	}
	return err
}

// ElasticBeanstalkSourceBundle represents Elastic Beanstalk SourceBundle Property Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-sourcebundle.html
type ElasticBeanstalkSourceBundle struct {
}

// ElasticBeanstalkSourceBundleList represents a list of ElasticBeanstalkSourceBundle
type ElasticBeanstalkSourceBundleList []ElasticBeanstalkSourceBundle

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticBeanstalkSourceBundleList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticBeanstalkSourceBundle{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticBeanstalkSourceBundleList{item}
		return nil
	}
	list := []ElasticBeanstalkSourceBundle{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticBeanstalkSourceBundleList(list)
		return nil
	}
	return err
}

// ElasticBeanstalkSourceConfiguration represents Elastic Beanstalk SourceConfiguration Property Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-configurationtemplate-sourceconfiguration.html
type ElasticBeanstalkSourceConfiguration struct {
}

// ElasticBeanstalkSourceConfigurationList represents a list of ElasticBeanstalkSourceConfiguration
type ElasticBeanstalkSourceConfigurationList []ElasticBeanstalkSourceConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticBeanstalkSourceConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticBeanstalkSourceConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticBeanstalkSourceConfigurationList{item}
		return nil
	}
	list := []ElasticBeanstalkSourceConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticBeanstalkSourceConfigurationList(list)
		return nil
	}
	return err
}

// ElasticLoadBalancingAccessLoggingPolicy represents Elastic Load Balancing AccessLoggingPolicy
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-accessloggingpolicy.html
type ElasticLoadBalancingAccessLoggingPolicy struct {
	// The interval for publishing access logs in minutes. You can specify an
	// interval of either 5 minutes or 60 minutes.
	EmitInterval *IntegerExpr `json:"EmitInterval,omitempty"`

	// Whether logging is enabled for the load balancer.
	Enabled *BoolExpr `json:"Enabled,omitempty"`

	// The name of an Amazon S3 bucket where access log files are stored.
	S3BucketName *StringExpr `json:"S3BucketName,omitempty"`

	// A prefix for the all log object keys, such as
	// my-load-balancer-logs/prod. If you store log files from multiple
	// sources in a single bucket, you can use a prefix to distinguish each
	// log file and its source.
	S3BucketPrefix *StringExpr `json:"S3BucketPrefix,omitempty"`
}

// ElasticLoadBalancingAccessLoggingPolicyList represents a list of ElasticLoadBalancingAccessLoggingPolicy
type ElasticLoadBalancingAccessLoggingPolicyList []ElasticLoadBalancingAccessLoggingPolicy

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticLoadBalancingAccessLoggingPolicyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticLoadBalancingAccessLoggingPolicy{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticLoadBalancingAccessLoggingPolicyList{item}
		return nil
	}
	list := []ElasticLoadBalancingAccessLoggingPolicy{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticLoadBalancingAccessLoggingPolicyList(list)
		return nil
	}
	return err
}

// ElasticLoadBalancingAppCookieStickinessPolicy represents ElasticLoadBalancing AppCookieStickinessPolicy Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-AppCookieStickinessPolicy.html
type ElasticLoadBalancingAppCookieStickinessPolicy struct {
	// Name of the application cookie used for stickiness.
	CookieName *StringExpr `json:"CookieName,omitempty"`

	// The name of the policy being created. The name must be unique within
	// the set of policies for this Load Balancer.
	PolicyName *StringExpr `json:"PolicyName,omitempty"`
}

// ElasticLoadBalancingAppCookieStickinessPolicyList represents a list of ElasticLoadBalancingAppCookieStickinessPolicy
type ElasticLoadBalancingAppCookieStickinessPolicyList []ElasticLoadBalancingAppCookieStickinessPolicy

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticLoadBalancingAppCookieStickinessPolicyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticLoadBalancingAppCookieStickinessPolicy{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticLoadBalancingAppCookieStickinessPolicyList{item}
		return nil
	}
	list := []ElasticLoadBalancingAppCookieStickinessPolicy{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticLoadBalancingAppCookieStickinessPolicyList(list)
		return nil
	}
	return err
}

// ElasticLoadBalancingConnectionDrainingPolicy represents Elastic Load Balancing ConnectionDrainingPolicy
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-connectiondrainingpolicy.html
type ElasticLoadBalancingConnectionDrainingPolicy struct {
	// Whether or not connection draining is enabled for the load balancer.
	Enabled *BoolExpr `json:"Enabled,omitempty"`

	// The time in seconds after the load balancer closes all connections to
	// a deregistered or unhealthy instance.
	Timeout *IntegerExpr `json:"Timeout,omitempty"`
}

// ElasticLoadBalancingConnectionDrainingPolicyList represents a list of ElasticLoadBalancingConnectionDrainingPolicy
type ElasticLoadBalancingConnectionDrainingPolicyList []ElasticLoadBalancingConnectionDrainingPolicy

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticLoadBalancingConnectionDrainingPolicyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticLoadBalancingConnectionDrainingPolicy{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticLoadBalancingConnectionDrainingPolicyList{item}
		return nil
	}
	list := []ElasticLoadBalancingConnectionDrainingPolicy{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticLoadBalancingConnectionDrainingPolicyList(list)
		return nil
	}
	return err
}

// ElasticLoadBalancingConnectionSettings represents Elastic Load Balancing ConnectionSettings
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-connectionsettings.html
type ElasticLoadBalancingConnectionSettings struct {
	// The time (in seconds) that a connection to the load balancer can
	// remain idle, which means no data is sent over the connection. After
	// the specified time, the load balancer closes the connection.
	IdleTimeout *IntegerExpr `json:"IdleTimeout,omitempty"`
}

// ElasticLoadBalancingConnectionSettingsList represents a list of ElasticLoadBalancingConnectionSettings
type ElasticLoadBalancingConnectionSettingsList []ElasticLoadBalancingConnectionSettings

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticLoadBalancingConnectionSettingsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticLoadBalancingConnectionSettings{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticLoadBalancingConnectionSettingsList{item}
		return nil
	}
	list := []ElasticLoadBalancingConnectionSettings{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticLoadBalancingConnectionSettingsList(list)
		return nil
	}
	return err
}

// ElasticLoadBalancingHealthCheck represents ElasticLoadBalancing HealthCheck Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-health-check.html
type ElasticLoadBalancingHealthCheck struct {
	// Specifies the number of consecutive health probe successes required
	// before moving the instance to the Healthy state.
	HealthyThreshold *StringExpr `json:"HealthyThreshold,omitempty"`

	// Specifies the approximate interval, in seconds, between health checks
	// of an individual instance.
	Interval *StringExpr `json:"Interval,omitempty"`

	// Specifies the instance's protocol and port to check. The protocol can
	// be TCP, HTTP, HTTPS, or SSL. The range of valid ports is 1 through
	// 65535.
	Target *StringExpr `json:"Target,omitempty"`

	// Specifies the amount of time, in seconds, during which no response
	// means a failed health probe. This value must be less than the value
	// for Interval.
	Timeout *StringExpr `json:"Timeout,omitempty"`

	// Specifies the number of consecutive health probe failures required
	// before moving the instance to the Unhealthy state.
	UnhealthyThreshold *StringExpr `json:"UnhealthyThreshold,omitempty"`
}

// ElasticLoadBalancingHealthCheckList represents a list of ElasticLoadBalancingHealthCheck
type ElasticLoadBalancingHealthCheckList []ElasticLoadBalancingHealthCheck

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticLoadBalancingHealthCheckList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticLoadBalancingHealthCheck{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticLoadBalancingHealthCheckList{item}
		return nil
	}
	list := []ElasticLoadBalancingHealthCheck{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticLoadBalancingHealthCheckList(list)
		return nil
	}
	return err
}

// ElasticLoadBalancingLBCookieStickinessPolicy represents ElasticLoadBalancing LBCookieStickinessPolicy Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-LBCookieStickinessPolicy.html
type ElasticLoadBalancingLBCookieStickinessPolicy struct {
	// The time period, in seconds, after which the cookie should be
	// considered stale. If this parameter isn't specified, the sticky
	// session will last for the duration of the browser session.
	CookieExpirationPeriod *StringExpr `json:"CookieExpirationPeriod,omitempty"`

	// The name of the policy being created. The name must be unique within
	// the set of policies for this load balancer.
	PolicyName interface{} `json:"PolicyName,omitempty"`
}

// ElasticLoadBalancingLBCookieStickinessPolicyList represents a list of ElasticLoadBalancingLBCookieStickinessPolicy
type ElasticLoadBalancingLBCookieStickinessPolicyList []ElasticLoadBalancingLBCookieStickinessPolicy

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticLoadBalancingLBCookieStickinessPolicyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticLoadBalancingLBCookieStickinessPolicy{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticLoadBalancingLBCookieStickinessPolicyList{item}
		return nil
	}
	list := []ElasticLoadBalancingLBCookieStickinessPolicy{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticLoadBalancingLBCookieStickinessPolicyList(list)
		return nil
	}
	return err
}

// ElasticLoadBalancingListener represents ElasticLoadBalancing Listener Property Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html
type ElasticLoadBalancingListener struct {
	// Specifies the TCP port on which the instance server is listening. This
	// property cannot be modified for the life of the load balancer.
	InstancePort *StringExpr `json:"InstancePort,omitempty"`

	// Specifies the protocol to use for routing traffic to back-end
	// instances—HTTP, HTTPS, TCP, or SSL. This property cannot be modified
	// for the life of the load balancer.
	InstanceProtocol *StringExpr `json:"InstanceProtocol,omitempty"`

	// Specifies the external load balancer port number. This property cannot
	// be modified for the life of the load balancer.
	LoadBalancerPort *StringExpr `json:"LoadBalancerPort,omitempty"`

	// A list of ElasticLoadBalancing policy names to associate with the
	// listener.
	PolicyNames *StringListExpr `json:"PolicyNames,omitempty"`

	// Specifies the load balancer transport protocol to use for routing —
	// HTTP, HTTPS, TCP or SSL. This property cannot be modified for the life
	// of the load balancer.
	Protocol *StringExpr `json:"Protocol,omitempty"`

	// The ARN of the SSL certificate to use. For more information about SSL
	// certificates, see Managing Server Certificates in the AWS Identity and
	// Access Management documentation.
	SSLCertificateId *StringExpr `json:"SSLCertificateId,omitempty"`
}

// ElasticLoadBalancingListenerList represents a list of ElasticLoadBalancingListener
type ElasticLoadBalancingListenerList []ElasticLoadBalancingListener

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticLoadBalancingListenerList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticLoadBalancingListener{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticLoadBalancingListenerList{item}
		return nil
	}
	list := []ElasticLoadBalancingListener{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticLoadBalancingListenerList(list)
		return nil
	}
	return err
}

// ElasticLoadBalancingPolicy represents ElasticLoadBalancing Policy Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-policy.html
type ElasticLoadBalancingPolicy struct {
	// A list of arbitrary attributes for this policy.
	Attributes interface{} `json:"Attributes,omitempty"`

	// A list of instance ports for the policy. These are the ports
	// associated with the back-end server.
	InstancePorts interface{} `json:"InstancePorts,omitempty"`

	// A list of external load balancer ports for the policy.
	LoadBalancerPorts interface{} `json:"LoadBalancerPorts,omitempty"`

	// A name for this policy that is unique to the load balancer.
	PolicyName *StringExpr `json:"PolicyName,omitempty"`

	// The name of the policy type for this policy. This must be one of the
	// types reported by the Elastic Load Balancing
	// DescribeLoadBalancerPolicyTypes action.
	PolicyType *StringExpr `json:"PolicyType,omitempty"`
}

// ElasticLoadBalancingPolicyList represents a list of ElasticLoadBalancingPolicy
type ElasticLoadBalancingPolicyList []ElasticLoadBalancingPolicy

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticLoadBalancingPolicyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticLoadBalancingPolicy{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticLoadBalancingPolicyList{item}
		return nil
	}
	list := []ElasticLoadBalancingPolicy{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticLoadBalancingPolicyList(list)
		return nil
	}
	return err
}

// IAMPolicies represents IAM Policies
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html
type IAMPolicies struct {
	// A policy document that describes what actions are allowed on which
	// resources.
	PolicyDocument interface{} `json:"PolicyDocument,omitempty"`

	// The name of the policy.
	PolicyName *StringExpr `json:"PolicyName,omitempty"`
}

// IAMPoliciesList represents a list of IAMPolicies
type IAMPoliciesList []IAMPolicies

// UnmarshalJSON sets the object from the provided JSON representation
func (l *IAMPoliciesList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := IAMPolicies{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = IAMPoliciesList{item}
		return nil
	}
	list := []IAMPolicies{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = IAMPoliciesList(list)
		return nil
	}
	return err
}

// IAMUserLoginProfile represents IAM User LoginProfile
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user-loginprofile.html
type IAMUserLoginProfile struct {
	// The password for the user.
	Password *StringExpr `json:"Password,omitempty"`

	// Specifies whether the user is required to set a new password the next
	// time the user logs in to the AWS Management Console.
	PasswordResetRequired *BoolExpr `json:"PasswordResetRequired,omitempty"`
}

// IAMUserLoginProfileList represents a list of IAMUserLoginProfile
type IAMUserLoginProfileList []IAMUserLoginProfile

// UnmarshalJSON sets the object from the provided JSON representation
func (l *IAMUserLoginProfileList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := IAMUserLoginProfile{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = IAMUserLoginProfileList{item}
		return nil
	}
	list := []IAMUserLoginProfile{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = IAMUserLoginProfileList(list)
		return nil
	}
	return err
}

// LambdaFunctionCode represents AWS Lambda Function Code
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-code.html
type LambdaFunctionCode struct {
	// The name of the Amazon Simple Storage Service (Amazon S3) bucket that
	// contains the source code of your Lambda function. The S3 bucket must
	// be in the same region in which the stack is created.
	S3Bucket *StringExpr `json:"S3Bucket,omitempty"`

	// The location and name of your packaged source code (.zip file).
	S3Key *StringExpr `json:"S3Key,omitempty"`

	// If you have versioning enabled, the version ID of your packaged source
	// code.
	S3ObjectVersion *StringExpr `json:"S3ObjectVersion,omitempty"`
}

// LambdaFunctionCodeList represents a list of LambdaFunctionCode
type LambdaFunctionCodeList []LambdaFunctionCode

// UnmarshalJSON sets the object from the provided JSON representation
func (l *LambdaFunctionCodeList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := LambdaFunctionCode{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = LambdaFunctionCodeList{item}
		return nil
	}
	list := []LambdaFunctionCode{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = LambdaFunctionCodeList(list)
		return nil
	}
	return err
}

// Name represents Name Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html
type Name struct {
}

// NameList represents a list of Name
type NameList []Name

// UnmarshalJSON sets the object from the provided JSON representation
func (l *NameList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := Name{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = NameList{item}
		return nil
	}
	list := []Name{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = NameList(list)
		return nil
	}
	return err
}

// OpsWorksAutoScalingThresholds represents AWS OpsWorks AutoScalingThresholds Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling-autoscalingthresholds.html
type OpsWorksAutoScalingThresholds struct {
	// The percentage of CPU utilization that triggers the starting or
	// stopping of instances (scaling).
	CpuThreshold *IntegerExpr `json:"CpuThreshold,omitempty"`

	// The amount of time (in minutes) after a scaling event occurs that AWS
	// OpsWorks should ignore metrics and not start any additional scaling
	// events.
	IgnoreMetricsTime *IntegerExpr `json:"IgnoreMetricsTime,omitempty"`

	// The number of instances to add or remove when the load exceeds a
	// threshold.
	InstanceCount *IntegerExpr `json:"InstanceCount,omitempty"`

	// The degree of system load that triggers the starting or stopping of
	// instances (scaling). For more information about how load is computed,
	// see Load (computing).
	LoadThreshold *IntegerExpr `json:"LoadThreshold,omitempty"`

	// The percentage of memory consumption that triggers the starting or
	// stopping of instances (scaling).
	MemoryThreshold *IntegerExpr `json:"MemoryThreshold,omitempty"`

	// The amount of time, in minutes, that the load must exceed a threshold
	// before instances are added or removed.
	ThresholdsWaitTime *IntegerExpr `json:"ThresholdsWaitTime,omitempty"`
}

// OpsWorksAutoScalingThresholdsList represents a list of OpsWorksAutoScalingThresholds
type OpsWorksAutoScalingThresholdsList []OpsWorksAutoScalingThresholds

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksAutoScalingThresholdsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksAutoScalingThresholds{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksAutoScalingThresholdsList{item}
		return nil
	}
	list := []OpsWorksAutoScalingThresholds{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksAutoScalingThresholdsList(list)
		return nil
	}
	return err
}

// OpsWorksChefConfiguration represents AWS OpsWorks ChefConfiguration Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-chefconfiguration.html
type OpsWorksChefConfiguration struct {
	// The Berkshelf version.
	BerkshelfVersion *StringExpr `json:"BerkshelfVersion,omitempty"`

	// Whether to enable Berkshelf.
	ManageBerkshelf *BoolExpr `json:"ManageBerkshelf,omitempty"`
}

// OpsWorksChefConfigurationList represents a list of OpsWorksChefConfiguration
type OpsWorksChefConfigurationList []OpsWorksChefConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksChefConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksChefConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksChefConfigurationList{item}
		return nil
	}
	list := []OpsWorksChefConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksChefConfigurationList(list)
		return nil
	}
	return err
}

// OpsWorksLayerLifeCycleConfiguration represents AWS OpsWorks Layer LifeCycleConfiguration
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration.html
type OpsWorksLayerLifeCycleConfiguration struct {
	// Specifies the shutdown event configuration for a layer.
	ShutdownEventConfiguration *OpsWorksLayerLifeCycleConfigurationShutdownEventConfiguration `json:"ShutdownEventConfiguration,omitempty"`
}

// OpsWorksLayerLifeCycleConfigurationList represents a list of OpsWorksLayerLifeCycleConfiguration
type OpsWorksLayerLifeCycleConfigurationList []OpsWorksLayerLifeCycleConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksLayerLifeCycleConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksLayerLifeCycleConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksLayerLifeCycleConfigurationList{item}
		return nil
	}
	list := []OpsWorksLayerLifeCycleConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksLayerLifeCycleConfigurationList(list)
		return nil
	}
	return err
}

// OpsWorksLayerLifeCycleConfigurationShutdownEventConfiguration represents AWS OpsWorks Layer LifeCycleConfiguration ShutdownEventConfiguration
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration-shutdowneventconfiguration.html
type OpsWorksLayerLifeCycleConfigurationShutdownEventConfiguration struct {
	// Indicates whether to wait for connections to drain from the Elastic
	// Load Balancing load balancers.
	DelayUntilElbConnectionsDrained *BoolExpr `json:"DelayUntilElbConnectionsDrained,omitempty"`

	// The time, in seconds, that AWS OpsWorks waits after a shutdown event
	// has been triggered before shutting down an instance.
	ExecutionTimeout *IntegerExpr `json:"ExecutionTimeout,omitempty"`
}

// OpsWorksLayerLifeCycleConfigurationShutdownEventConfigurationList represents a list of OpsWorksLayerLifeCycleConfigurationShutdownEventConfiguration
type OpsWorksLayerLifeCycleConfigurationShutdownEventConfigurationList []OpsWorksLayerLifeCycleConfigurationShutdownEventConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksLayerLifeCycleConfigurationShutdownEventConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksLayerLifeCycleConfigurationShutdownEventConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksLayerLifeCycleConfigurationShutdownEventConfigurationList{item}
		return nil
	}
	list := []OpsWorksLayerLifeCycleConfigurationShutdownEventConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksLayerLifeCycleConfigurationShutdownEventConfigurationList(list)
		return nil
	}
	return err
}

// OpsWorksLoadBasedAutoScaling represents AWS OpsWorks LoadBasedAutoScaling Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling.html
type OpsWorksLoadBasedAutoScaling struct {
	// The threshold below which the instances are scaled down (stopped). If
	// the load falls below this threshold for a specified amount of time,
	// AWS OpsWorks stops a specified number of instances.
	DownScaling *OpsWorksAutoScalingThresholds `json:"DownScaling,omitempty"`

	// Whether to enable automatic load-based scaling for the layer.
	Enable *BoolExpr `json:"Enable,omitempty"`

	// The threshold above which the instances are scaled up (added). If the
	// load exceeds this thresholds for a specified amount of time, AWS
	// OpsWorks starts a specified number of instances.
	UpScaling *OpsWorksAutoScalingThresholds `json:"UpScaling,omitempty"`
}

// OpsWorksLoadBasedAutoScalingList represents a list of OpsWorksLoadBasedAutoScaling
type OpsWorksLoadBasedAutoScalingList []OpsWorksLoadBasedAutoScaling

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksLoadBasedAutoScalingList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksLoadBasedAutoScaling{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksLoadBasedAutoScalingList{item}
		return nil
	}
	list := []OpsWorksLoadBasedAutoScaling{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksLoadBasedAutoScalingList(list)
		return nil
	}
	return err
}

// OpsWorksRecipes represents AWS OpsWorks Recipes Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-recipes.html
type OpsWorksRecipes struct {
	// Custom recipe names to be run following a Configure event. The event
	// occurs on all of the stack's instances when an instance enters or
	// leaves the online state.
	Configure *StringListExpr `json:"Configure,omitempty"`

	// Custom recipe names to be run following a Deploy event. The event
	// occurs when you run a deploy command, typically to deploy an
	// application to a set of application server instances.
	Deploy *StringListExpr `json:"Deploy,omitempty"`

	// Custom recipe names to be run following a Setup event. This event
	// occurs on a new instance after it successfully boots.
	Setup *StringListExpr `json:"Setup,omitempty"`

	// Custom recipe names to be run following a Shutdown event. This event
	// occurs after you direct AWS OpsWorks to shut an instance down before
	// the associated Amazon EC2 instance is actually terminated.
	Shutdown *StringListExpr `json:"Shutdown,omitempty"`

	// Custom recipe names to be run following a Undeploy event. This event
	// occurs when you delete an app or run an undeploy command to remove an
	// app from a set of application server instances.
	Undeploy *StringListExpr `json:"Undeploy,omitempty"`
}

// OpsWorksRecipesList represents a list of OpsWorksRecipes
type OpsWorksRecipesList []OpsWorksRecipes

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksRecipesList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksRecipes{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksRecipesList{item}
		return nil
	}
	list := []OpsWorksRecipes{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksRecipesList(list)
		return nil
	}
	return err
}

// OpsWorksSource represents AWS OpsWorks Source Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html
type OpsWorksSource struct {
	// This parameter depends on the repository type. For Amazon S3 bundles,
	// set Password to the appropriate IAM secret access key. For HTTP
	// bundles, Git repositories, and Subversion repositories, set Password
	// to the appropriate password.
	Password *StringExpr `json:"Password,omitempty"`

	// The application's version. With AWS OpsWorks, you can deploy new
	// versions of an application. One of the simplest approaches is to have
	// branches or revisions in your repository that represent different
	// versions that can potentially be deployed.
	Revision *StringExpr `json:"Revision,omitempty"`

	// The repository's SSH key.
	SshKey *StringExpr `json:"SshKey,omitempty"`

	// The repository type.
	Type *StringExpr `json:"Type,omitempty"`

	// The source URL.
	Url *StringExpr `json:"Url,omitempty"`

	// This parameter depends on the repository type. For Amazon S3 bundles,
	// set Username to the appropriate IAM access key ID. For HTTP bundles,
	// Git repositories, and Subversion repositories, set Username to the
	// appropriate user name.
	Username *StringExpr `json:"Username,omitempty"`
}

// OpsWorksSourceList represents a list of OpsWorksSource
type OpsWorksSourceList []OpsWorksSource

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksSourceList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksSource{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksSourceList{item}
		return nil
	}
	list := []OpsWorksSource{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksSourceList(list)
		return nil
	}
	return err
}

// OpsWorksSslConfiguration represents AWS OpsWorks SslConfiguration Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-sslconfiguration.html
type OpsWorksSslConfiguration struct {
	// The contents of the certificate's domain.crt file.
	Certificate *StringExpr `json:"Certificate,omitempty"`

	// An intermediate certificate authority key or client authentication.
	Chain *StringExpr `json:"Chain,omitempty"`

	// The private key; the contents of the certificate's domain.kex file.
	PrivateKey *StringExpr `json:"PrivateKey,omitempty"`
}

// OpsWorksSslConfigurationList represents a list of OpsWorksSslConfiguration
type OpsWorksSslConfigurationList []OpsWorksSslConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksSslConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksSslConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksSslConfigurationList{item}
		return nil
	}
	list := []OpsWorksSslConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksSslConfigurationList(list)
		return nil
	}
	return err
}

// OpsWorksStackConfigurationManager represents AWS OpsWorks StackConfigurationManager Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-stackconfigmanager.html
type OpsWorksStackConfigurationManager struct {
	// The name of the configuration manager.
	Name *StringExpr `json:"Name,omitempty"`

	// The Chef version.
	Version *StringExpr `json:"Version,omitempty"`
}

// OpsWorksStackConfigurationManagerList represents a list of OpsWorksStackConfigurationManager
type OpsWorksStackConfigurationManagerList []OpsWorksStackConfigurationManager

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksStackConfigurationManagerList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksStackConfigurationManager{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksStackConfigurationManagerList{item}
		return nil
	}
	list := []OpsWorksStackConfigurationManager{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksStackConfigurationManagerList(list)
		return nil
	}
	return err
}

// OpsWorksTimeBasedAutoScaling represents AWS OpsWorks TimeBasedAutoScaling Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-timebasedautoscaling.html
type OpsWorksTimeBasedAutoScaling struct {
	// The schedule for Friday.
	Friday *StringExpr `json:"Friday,omitempty"`

	// The schedule for Monday.
	Monday *StringExpr `json:"Monday,omitempty"`

	// The schedule for Saturday.
	Saturday *StringExpr `json:"Saturday,omitempty"`

	// The schedule for Sunday.
	Sunday *StringExpr `json:"Sunday,omitempty"`

	// The schedule for Thursday.
	Thursday *StringExpr `json:"Thursday,omitempty"`

	// The schedule for Tuesday.
	Tuesday *StringExpr `json:"Tuesday,omitempty"`

	// The schedule for Wednesday.
	Wednesday *StringExpr `json:"Wednesday,omitempty"`
}

// OpsWorksTimeBasedAutoScalingList represents a list of OpsWorksTimeBasedAutoScaling
type OpsWorksTimeBasedAutoScalingList []OpsWorksTimeBasedAutoScaling

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksTimeBasedAutoScalingList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksTimeBasedAutoScaling{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksTimeBasedAutoScalingList{item}
		return nil
	}
	list := []OpsWorksTimeBasedAutoScaling{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksTimeBasedAutoScalingList(list)
		return nil
	}
	return err
}

// OpsWorksVolumeConfiguration represents AWS OpsWorks VolumeConfiguration Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfig.html
type OpsWorksVolumeConfiguration struct {
	// The number of I/O operations per second (IOPS) to provision for the
	// volume.
	Iops *IntegerExpr `json:"Iops,omitempty"`

	// The volume mount point, such as /dev/sdh.
	MountPoint *StringExpr `json:"MountPoint,omitempty"`

	// The number of disks in the volume.
	NumberOfDisks *IntegerExpr `json:"NumberOfDisks,omitempty"`

	// The volume RAID level.
	RaidLevel *IntegerExpr `json:"RaidLevel,omitempty"`

	// The volume size.
	Size *IntegerExpr `json:"Size,omitempty"`

	// The type of volume, such as magnetic or SSD. For valid values, see
	// VolumeConfiguration in the AWS OpsWorks API Reference.
	VolumeType *StringExpr `json:"VolumeType,omitempty"`
}

// OpsWorksVolumeConfigurationList represents a list of OpsWorksVolumeConfiguration
type OpsWorksVolumeConfigurationList []OpsWorksVolumeConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksVolumeConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksVolumeConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksVolumeConfigurationList{item}
		return nil
	}
	list := []OpsWorksVolumeConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksVolumeConfigurationList(list)
		return nil
	}
	return err
}

// RedshiftParameter represents Amazon Redshift Parameter Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-property-redshift-clusterparametergroup-parameter.html
type RedshiftParameter struct {
	// The name of the parameter.
	ParameterName *StringExpr `json:"ParameterName,omitempty"`

	// The value of the parameter.
	ParameterValue *StringExpr `json:"ParameterValue,omitempty"`
}

// RedshiftParameterList represents a list of RedshiftParameter
type RedshiftParameterList []RedshiftParameter

// UnmarshalJSON sets the object from the provided JSON representation
func (l *RedshiftParameterList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := RedshiftParameter{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = RedshiftParameterList{item}
		return nil
	}
	list := []RedshiftParameter{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = RedshiftParameterList(list)
		return nil
	}
	return err
}

// ResourceTag represents AWS CloudFormation Resource Tags Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html
type ResourceTag struct {
	// The key name of the tag. You can specify a value that is 1 to 128
	// Unicode characters in length and cannot be prefixed with aws:. You can
	// use any of the following characters: the set of Unicode letters,
	// digits, whitespace, _, ., /, =, +, and -.
	Key *StringExpr `json:"Key,omitempty"`

	// The value for the tag. You can specify a value that is 1 to 128
	// Unicode characters in length and cannot be prefixed with aws:. You can
	// use any of the following characters: the set of Unicode letters,
	// digits, whitespace, _, ., /, =, +, and -.
	Value *StringExpr `json:"Value,omitempty"`
}

// ResourceTagList represents a list of ResourceTag
type ResourceTagList []ResourceTag

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ResourceTagList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ResourceTag{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ResourceTagList{item}
		return nil
	}
	list := []ResourceTag{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ResourceTagList(list)
		return nil
	}
	return err
}

// RDSOptionGroupOptionConfigurations represents Amazon RDS OptionGroup OptionConfigurations
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations.html
type RDSOptionGroupOptionConfigurations struct {
	// A list of database security group names for this option. If the option
	// requires access to a port, the security groups must allow access to
	// that port. If you specify this property, don't specify the
	// VPCSecurityGroupMemberships property.
	DBSecurityGroupMemberships *StringListExpr `json:"DBSecurityGroupMemberships,omitempty"`

	// The name of the option. For more information about options, see
	// Working with Option Groups in the Amazon Relational Database Service
	// User Guide.
	OptionName *StringExpr `json:"OptionName,omitempty"`

	// The settings for this option.
	OptionSettings *RDSOptionGroupOptionConfigurationsOptionSettings `json:"OptionSettings,omitempty"`

	// The port number that this option uses.
	Port *IntegerExpr `json:"Port,omitempty"`

	// A list of VPC security group IDs for this option. If the option
	// requires access to a port, the security groups must allow access to
	// that port. If you specify this property, don't specify the
	// DBSecurityGroupMemberships property.
	VpcSecurityGroupMemberships *StringListExpr `json:"VpcSecurityGroupMemberships,omitempty"`
}

// RDSOptionGroupOptionConfigurationsList represents a list of RDSOptionGroupOptionConfigurations
type RDSOptionGroupOptionConfigurationsList []RDSOptionGroupOptionConfigurations

// UnmarshalJSON sets the object from the provided JSON representation
func (l *RDSOptionGroupOptionConfigurationsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := RDSOptionGroupOptionConfigurations{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = RDSOptionGroupOptionConfigurationsList{item}
		return nil
	}
	list := []RDSOptionGroupOptionConfigurations{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = RDSOptionGroupOptionConfigurationsList(list)
		return nil
	}
	return err
}

// RDSOptionGroupOptionConfigurationsOptionSettings represents Amazon RDS OptionGroup OptionConfigurations OptionSettings
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations-optionsettings.html
type RDSOptionGroupOptionConfigurationsOptionSettings struct {
	// The name of the option setting that you want to specify.
	Name *StringExpr `json:"Name,omitempty"`

	// The value of the option setting.
	Value *StringExpr `json:"Value,omitempty"`
}

// RDSOptionGroupOptionConfigurationsOptionSettingsList represents a list of RDSOptionGroupOptionConfigurationsOptionSettings
type RDSOptionGroupOptionConfigurationsOptionSettingsList []RDSOptionGroupOptionConfigurationsOptionSettings

// UnmarshalJSON sets the object from the provided JSON representation
func (l *RDSOptionGroupOptionConfigurationsOptionSettingsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := RDSOptionGroupOptionConfigurationsOptionSettings{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = RDSOptionGroupOptionConfigurationsOptionSettingsList{item}
		return nil
	}
	list := []RDSOptionGroupOptionConfigurationsOptionSettings{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = RDSOptionGroupOptionConfigurationsOptionSettingsList(list)
		return nil
	}
	return err
}

// RDSSecurityGroupRule represents Amazon RDS Security Group Rule
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group-rule.html
type RDSSecurityGroupRule struct {
	// The IP range to authorize.
	CIDRIP *StringExpr `json:"CIDRIP,omitempty"`

	// Id of the VPC or EC2 Security Group to authorize.
	EC2SecurityGroupId *StringExpr `json:"EC2SecurityGroupId,omitempty"`

	// Name of the EC2 Security Group to authorize.
	EC2SecurityGroupName *StringExpr `json:"EC2SecurityGroupName,omitempty"`

	// AWS Account Number of the owner of the EC2 Security Group specified in
	// the EC2SecurityGroupName parameter. The AWS Access Key ID is not an
	// acceptable value.
	EC2SecurityGroupOwnerId *StringExpr `json:"EC2SecurityGroupOwnerId,omitempty"`
}

// RDSSecurityGroupRuleList represents a list of RDSSecurityGroupRule
type RDSSecurityGroupRuleList []RDSSecurityGroupRule

// UnmarshalJSON sets the object from the provided JSON representation
func (l *RDSSecurityGroupRuleList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := RDSSecurityGroupRule{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = RDSSecurityGroupRuleList{item}
		return nil
	}
	list := []RDSSecurityGroupRule{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = RDSSecurityGroupRuleList(list)
		return nil
	}
	return err
}

// Route53AliasTargetProperty represents Route 53 AliasTarget Property
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-aliastarget.html
type Route53AliasTargetProperty struct {
	// The DNS name of the load balancer, the domain name of the CloudFront
	// distribution, or the website endpoint of the Amazon S3 bucket that is
	// the target of the alias.
	DNSName *StringExpr `json:"DNSName,omitempty"`

	// Whether Amazon Route 53 checks the health of the resource record sets
	// in the alias target when responding to DNS queries. For more
	// information about using this property, see EvaluateTargetHealth in the
	// Amazon Route 53 API Reference.
	EvaluateTargetHealth *BoolExpr `json:"EvaluateTargetHealth,omitempty"`

	// The hosted zone ID. For load balancers, use the canonical hosted zone
	// ID of the load balancer. For Amazon S3, use the hosted zone ID for
	// your bucket's website endpoint. For CloudFront, use Z2FDTNDATAQYW2.
	// For examples, see Example: Creating Alias Resource Record Sets in the
	// Amazon Route 53 API Reference.
	HostedZoneId *StringExpr `json:"HostedZoneId,omitempty"`
}

// Route53AliasTargetPropertyList represents a list of Route53AliasTargetProperty
type Route53AliasTargetPropertyList []Route53AliasTargetProperty

// UnmarshalJSON sets the object from the provided JSON representation
func (l *Route53AliasTargetPropertyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := Route53AliasTargetProperty{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = Route53AliasTargetPropertyList{item}
		return nil
	}
	list := []Route53AliasTargetProperty{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = Route53AliasTargetPropertyList(list)
		return nil
	}
	return err
}

// Route53RecordSetGeoLocationProperty represents Amazon Route 53 Record Set GeoLocation Property
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html
type Route53RecordSetGeoLocationProperty struct {
	// All DNS queries from the continent that you specified are routed to
	// this resource record set. If you specify this property, omit the
	// CountryCode and SubdivisionCode properties.
	ContinentCode *StringExpr `json:"ContinentCode,omitempty"`

	// All DNS queries from the country that you specified are routed to this
	// resource record set. If you specify this property, omit the
	// ContinentCode property.
	CountryCode *StringExpr `json:"CountryCode,omitempty"`

	// If you specified US for the country code, you can specify a state in
	// the United States. All DNS queries from the state that you specified
	// are routed to this resource record set. If you specify this property,
	// you must specify US for the CountryCode and omit the ContinentCode
	// property.
	SubdivisionCode *StringExpr `json:"SubdivisionCode,omitempty"`
}

// Route53RecordSetGeoLocationPropertyList represents a list of Route53RecordSetGeoLocationProperty
type Route53RecordSetGeoLocationPropertyList []Route53RecordSetGeoLocationProperty

// UnmarshalJSON sets the object from the provided JSON representation
func (l *Route53RecordSetGeoLocationPropertyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := Route53RecordSetGeoLocationProperty{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = Route53RecordSetGeoLocationPropertyList{item}
		return nil
	}
	list := []Route53RecordSetGeoLocationProperty{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = Route53RecordSetGeoLocationPropertyList(list)
		return nil
	}
	return err
}

// Route53HealthCheckConfig represents Amazon Route 53 HealthCheckConfig
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html
type Route53HealthCheckConfig struct {
	// The number of consecutive health checks that an endpoint must pass or
	// fail for Amazon Route 53 to change the current status of the endpoint
	// from unhealthy to healthy or healthy to unhealthy. For more
	// information, see How Amazon Route 53 Determines Whether an Endpoint
	// Is Healthy in the Amazon Route 53 Developer Guide.
	FailureThreshold *IntegerExpr `json:"FailureThreshold,omitempty"`

	// If you specified the IPAddress property, the value that you want
	// Amazon Route 53 to pass in the host header in all health checks
	// except for TCP health checks. If you don't specify an IP address, the
	// domain that Amazon Route 53 sends a DNS request to. Amazon Route 53
	// uses the IP address that the DNS returns to check the health of the
	// endpoint.
	FullyQualifiedDomainName *StringExpr `json:"FullyQualifiedDomainName,omitempty"`

	// The IPv4 IP address of the endpoint on which you want Amazon Route 53
	// to perform health checks. If you don't specify an IP address, Amazon
	// Route 53 sends a DNS request to resolve the domain name that you
	// specify in the FullyQualifiedDomainName property.
	IPAddress *StringExpr `json:"IPAddress,omitempty"`

	// The port on the endpoint on which you want Amazon Route 53 to perform
	// health checks.
	Port *IntegerExpr `json:"Port,omitempty"`

	// The number of seconds between the time that Amazon Route 53 gets a
	// response from your endpoint and the time that it sends the next
	// health-check request. Each Amazon Route 53 health checker makes
	// requests at this interval.
	RequestInterval *IntegerExpr `json:"RequestInterval,omitempty"`

	// The path that you want Amazon Route 53 to request when performing
	// health checks. The path can be any value for which your endpoint
	// returns an HTTP status code of 2xx or 3xx when the endpoint is
	// healthy, such as /docs/route53-health-check.html.
	ResourcePath *StringExpr `json:"ResourcePath,omitempty"`

	// If the value of the Type property is HTTP_STR_MATCH or HTTP_STR_MATCH,
	// the string that you want Amazon Route 53 to search for in the
	// response body from the specified resource. If the string appears in
	// the response body, Amazon Route 53 considers the resource healthy.
	SearchString *StringExpr `json:"SearchString,omitempty"`

	// The type of health check that you want to create, which indicates how
	// Amazon Route 53 determines whether an endpoint is healthy. You can
	// specify HTTP, HTTPS, HTTP_STR_MATCH, HTTPS_STR_MATCH, or TCP. For
	// information about the different types, see the Type element in the
	// Amazon Route 53 API Reference.
	Type *StringExpr `json:"Type,omitempty"`
}

// Route53HealthCheckConfigList represents a list of Route53HealthCheckConfig
type Route53HealthCheckConfigList []Route53HealthCheckConfig

// UnmarshalJSON sets the object from the provided JSON representation
func (l *Route53HealthCheckConfigList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := Route53HealthCheckConfig{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = Route53HealthCheckConfigList{item}
		return nil
	}
	list := []Route53HealthCheckConfig{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = Route53HealthCheckConfigList(list)
		return nil
	}
	return err
}

// Route53HealthCheckTags represents Amazon Route 53 HealthCheckTags
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthchecktags.html
type Route53HealthCheckTags struct {
	// The key name of the tag.
	Key *StringExpr `json:"Key,omitempty"`

	// The value for the tag.
	Value *StringExpr `json:"Value,omitempty"`
}

// Route53HealthCheckTagsList represents a list of Route53HealthCheckTags
type Route53HealthCheckTagsList []Route53HealthCheckTags

// UnmarshalJSON sets the object from the provided JSON representation
func (l *Route53HealthCheckTagsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := Route53HealthCheckTags{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = Route53HealthCheckTagsList{item}
		return nil
	}
	list := []Route53HealthCheckTags{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = Route53HealthCheckTagsList(list)
		return nil
	}
	return err
}

// Route53HostedZoneConfigProperty represents Amazon Route 53 HostedZoneConfig Property
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzoneconfig.html
type Route53HostedZoneConfigProperty struct {
	// Any comments that you want to include about the hosted zone.
	Comment *StringExpr `json:"Comment,omitempty"`
}

// Route53HostedZoneConfigPropertyList represents a list of Route53HostedZoneConfigProperty
type Route53HostedZoneConfigPropertyList []Route53HostedZoneConfigProperty

// UnmarshalJSON sets the object from the provided JSON representation
func (l *Route53HostedZoneConfigPropertyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := Route53HostedZoneConfigProperty{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = Route53HostedZoneConfigPropertyList{item}
		return nil
	}
	list := []Route53HostedZoneConfigProperty{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = Route53HostedZoneConfigPropertyList(list)
		return nil
	}
	return err
}

// Route53HostedZoneTags represents Amazon Route 53 HostedZoneTags
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzonetags.html
type Route53HostedZoneTags struct {
	// The key name of the tag.
	Key *StringExpr `json:"Key,omitempty"`

	// The value for the tag.
	Value *StringExpr `json:"Value,omitempty"`
}

// Route53HostedZoneTagsList represents a list of Route53HostedZoneTags
type Route53HostedZoneTagsList []Route53HostedZoneTags

// UnmarshalJSON sets the object from the provided JSON representation
func (l *Route53HostedZoneTagsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := Route53HostedZoneTags{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = Route53HostedZoneTagsList{item}
		return nil
	}
	list := []Route53HostedZoneTags{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = Route53HostedZoneTagsList(list)
		return nil
	}
	return err
}

// Route53HostedZoneVPCs represents Amazon Route 53 HostedZoneVPCs
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone-hostedzonevpcs.html
type Route53HostedZoneVPCs struct {
	// The ID of the Amazon VPC that you want to associate with the hosted
	// zone.
	VPCId *StringExpr `json:"VPCId,omitempty"`

	// The region in which the Amazon VPC was created as specified in the
	// VPCId property.
	VPCRegion *StringExpr `json:"VPCRegion,omitempty"`
}

// Route53HostedZoneVPCsList represents a list of Route53HostedZoneVPCs
type Route53HostedZoneVPCsList []Route53HostedZoneVPCs

// UnmarshalJSON sets the object from the provided JSON representation
func (l *Route53HostedZoneVPCsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := Route53HostedZoneVPCs{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = Route53HostedZoneVPCsList{item}
		return nil
	}
	list := []Route53HostedZoneVPCs{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = Route53HostedZoneVPCsList(list)
		return nil
	}
	return err
}

// S3CorsConfiguration represents Amazon S3 Cors Configuration
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors.html
type S3CorsConfiguration struct {
	// A set of origins and methods that you allow.
	CorsRules *S3CorsConfigurationRule `json:"CorsRules,omitempty"`
}

// S3CorsConfigurationList represents a list of S3CorsConfiguration
type S3CorsConfigurationList []S3CorsConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3CorsConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3CorsConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3CorsConfigurationList{item}
		return nil
	}
	list := []S3CorsConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3CorsConfigurationList(list)
		return nil
	}
	return err
}

// S3CorsConfigurationRule represents Amazon S3 Cors Configuration Rule
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors-corsrule.html
type S3CorsConfigurationRule struct {
	// Headers that are specified in the Access-Control-Request-Headers
	// header. These headers are allowed in a preflight OPTIONS request. In
	// response to any preflight OPTIONS request, Amazon S3 returns any
	// requested headers that are allowed.
	AllowedHeaders *StringListExpr `json:"AllowedHeaders,omitempty"`

	// An HTTP method that you allow the origin to execute. The valid values
	// are GET, PUT, HEAD, POST, and DELETE.
	AllowedMethods *StringListExpr `json:"AllowedMethods,omitempty"`

	// An origin that you allow to send cross-domain requests.
	AllowedOrigins *StringListExpr `json:"AllowedOrigins,omitempty"`

	// One or more headers in the response that are accessible to client
	// applications (for example, from a JavaScript XMLHttpRequest object).
	ExposedHeaders *StringListExpr `json:"ExposedHeaders,omitempty"`

	// A unique identifier for this rule. The value cannot be more than 255
	// characters.
	Id *StringExpr `json:"Id,omitempty"`

	// The time in seconds that your browser is to cache the preflight
	// response for the specified resource.
	MaxAge *IntegerExpr `json:"MaxAge,omitempty"`
}

// S3CorsConfigurationRuleList represents a list of S3CorsConfigurationRule
type S3CorsConfigurationRuleList []S3CorsConfigurationRule

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3CorsConfigurationRuleList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3CorsConfigurationRule{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3CorsConfigurationRuleList{item}
		return nil
	}
	list := []S3CorsConfigurationRule{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3CorsConfigurationRuleList(list)
		return nil
	}
	return err
}

// S3LifecycleConfiguration represents Amazon S3 Lifecycle Configuration
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig.html
type S3LifecycleConfiguration struct {
	// A lifecycle rule for individual objects in an S3 bucket.
	Rules *S3LifecycleRule `json:"Rules,omitempty"`
}

// S3LifecycleConfigurationList represents a list of S3LifecycleConfiguration
type S3LifecycleConfigurationList []S3LifecycleConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3LifecycleConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3LifecycleConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3LifecycleConfigurationList{item}
		return nil
	}
	list := []S3LifecycleConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3LifecycleConfigurationList(list)
		return nil
	}
	return err
}

// S3LifecycleRule represents Amazon S3 Lifecycle Rule
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html
type S3LifecycleRule struct {
	// Indicates when objects are deleted from Amazon S3 and Amazon Glacier.
	// The date value must be in ISO 8601 format. The time is always midnight
	// UTC. If you specify an expiration and transition time, you must use
	// the same time unit for both properties (either in days or by date).
	// The expiration time must also be later than the transition time.
	ExpirationDate *StringExpr `json:"ExpirationDate,omitempty"`

	// Indicates the number of days after creation when objects are deleted
	// from Amazon S3 and Amazon Glacier. If you specify an expiration and
	// transition time, you must use the same time unit for both properties
	// (either in days or by date). The expiration time must also be later
	// than the transition time.
	ExpirationInDays *IntegerExpr `json:"ExpirationInDays,omitempty"`

	// A unique identifier for this rule. The value cannot be more than 255
	// characters.
	Id *StringExpr `json:"Id,omitempty"`

	// For buckets with versioning enabled (or suspended), specifies the
	// time, in days, between when a new version of the object is uploaded to
	// the bucket and when old versions of the object expire. When object
	// versions expire, Amazon S3 permanently deletes them. If you specify a
	// transition and expiration time, the expiration time must be later than
	// the transition time.
	NoncurrentVersionExpirationInDays *IntegerExpr `json:"NoncurrentVersionExpirationInDays,omitempty"`

	// For buckets with versioning enabled (or suspended), specifies when
	// non-current objects transition to a specified storage class. If you
	// specify a transition and expiration time, the expiration time must be
	// later than the transition time.
	NoncurrentVersionTransition *S3LifecycleRuleNoncurrentVersionTransition `json:"NoncurrentVersionTransition,omitempty"`

	// Object key prefix that identifies one or more objects to which this
	// rule applies.
	Prefix *StringExpr `json:"Prefix,omitempty"`

	// Specify either Enabled or Disabled. If you specify Enabled, Amazon S3
	// executes this rule as scheduled. If you specify Disabled, Amazon S3
	// ignores this rule.
	Status *StringExpr `json:"Status,omitempty"`

	// Specifies when an object transitions to a specified storage class. If
	// you specify an expiration and transition time, you must use the same
	// time unit for both properties (either in days or by date). The
	// expiration time must also be later than the transition time.
	Transition *S3LifecycleRuleTransition `json:"Transition,omitempty"`
}

// S3LifecycleRuleList represents a list of S3LifecycleRule
type S3LifecycleRuleList []S3LifecycleRule

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3LifecycleRuleList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3LifecycleRule{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3LifecycleRuleList{item}
		return nil
	}
	list := []S3LifecycleRule{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3LifecycleRuleList(list)
		return nil
	}
	return err
}

// S3LifecycleRuleNoncurrentVersionTransition represents Amazon S3 Lifecycle Rule NoncurrentVersionTransition
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-noncurrentversiontransition.html
type S3LifecycleRuleNoncurrentVersionTransition struct {
	// The storage class to which you want the noncurrent objects to
	// transition. Currently, you can specify only Glacier.
	StorageClass *StringExpr `json:"StorageClass,omitempty"`

	// The number of days between the time that a new version of the object
	// is uploaded to the bucket and when old versions of the object are
	// transitioned to the specified storage class.
	TransitionInDays *IntegerExpr `json:"TransitionInDays,omitempty"`
}

// S3LifecycleRuleNoncurrentVersionTransitionList represents a list of S3LifecycleRuleNoncurrentVersionTransition
type S3LifecycleRuleNoncurrentVersionTransitionList []S3LifecycleRuleNoncurrentVersionTransition

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3LifecycleRuleNoncurrentVersionTransitionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3LifecycleRuleNoncurrentVersionTransition{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3LifecycleRuleNoncurrentVersionTransitionList{item}
		return nil
	}
	list := []S3LifecycleRuleNoncurrentVersionTransition{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3LifecycleRuleNoncurrentVersionTransitionList(list)
		return nil
	}
	return err
}

// S3LifecycleRuleTransition represents Amazon S3 Lifecycle Rule Transition
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-transition.html
type S3LifecycleRuleTransition struct {
	// The storage class to which you want the object to transition.
	// Currently, you can specify only Glacier.
	StorageClass *StringExpr `json:"StorageClass,omitempty"`

	// Indicates when objects are transitioned to the specified storage
	// class. The date value must be in ISO 8601 format. The time is always
	// midnight UTC.
	TransitionDate *StringExpr `json:"TransitionDate,omitempty"`

	// Indicates the number of days after creation when objects are
	// transitioned to the specified storage class.
	TransitionInDays *IntegerExpr `json:"TransitionInDays,omitempty"`
}

// S3LifecycleRuleTransitionList represents a list of S3LifecycleRuleTransition
type S3LifecycleRuleTransitionList []S3LifecycleRuleTransition

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3LifecycleRuleTransitionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3LifecycleRuleTransition{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3LifecycleRuleTransitionList{item}
		return nil
	}
	list := []S3LifecycleRuleTransition{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3LifecycleRuleTransitionList(list)
		return nil
	}
	return err
}

// S3LoggingConfiguration represents Amazon S3 Logging Configuration
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-loggingconfig.html
type S3LoggingConfiguration struct {
	// The name of an Amazon S3 bucket where Amazon S3 store server access
	// log files. You can store log files in any bucket that you own. By
	// default, logs are stored in the bucket where the LoggingConfiguration
	// property is defined.
	DestinationBucketName *StringExpr `json:"DestinationBucketName,omitempty"`

	// A prefix for the all log object keys. If you store log files from
	// multiple Amazon S3 buckets in a single bucket, you can use a prefix to
	// distinguish which log files came from which bucket.
	LogFilePrefix *StringExpr `json:"LogFilePrefix,omitempty"`
}

// S3LoggingConfigurationList represents a list of S3LoggingConfiguration
type S3LoggingConfigurationList []S3LoggingConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3LoggingConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3LoggingConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3LoggingConfigurationList{item}
		return nil
	}
	list := []S3LoggingConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3LoggingConfigurationList(list)
		return nil
	}
	return err
}

// S3NotificationConfiguration represents Amazon S3 Notification Configuration
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig.html
type S3NotificationConfiguration struct {
	// The topic to which notifications are sent and the events for which
	// notification are generated.
	TopicConfigurations *S3NotificationTopicConfigurations `json:"TopicConfigurations,omitempty"`
}

// S3NotificationConfigurationList represents a list of S3NotificationConfiguration
type S3NotificationConfigurationList []S3NotificationConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3NotificationConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3NotificationConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3NotificationConfigurationList{item}
		return nil
	}
	list := []S3NotificationConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3NotificationConfigurationList(list)
		return nil
	}
	return err
}

// S3NotificationTopicConfigurations represents Amazon S3 Notification Topic Configurations
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-topicconfig.html
type S3NotificationTopicConfigurations struct {
	// The Amazon S3 bucket event about which to send notifications. For more
	// information, see Supported Event Types in the Amazon Simple Storage
	// Service Developer Guide.
	Event *StringExpr `json:"Event,omitempty"`

	// The Amazon SNS topic to which Amazon S3 reports the specified events.
	Topic *StringExpr `json:"Topic,omitempty"`
}

// S3NotificationTopicConfigurationsList represents a list of S3NotificationTopicConfigurations
type S3NotificationTopicConfigurationsList []S3NotificationTopicConfigurations

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3NotificationTopicConfigurationsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3NotificationTopicConfigurations{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3NotificationTopicConfigurationsList{item}
		return nil
	}
	list := []S3NotificationTopicConfigurations{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3NotificationTopicConfigurationsList(list)
		return nil
	}
	return err
}

// S3VersioningConfiguration represents Amazon S3 Versioning Configuration
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-versioningconfig.html
type S3VersioningConfiguration struct {
	// The versioning state of an Amazon S3 bucket. If you enable versioning,
	// you must suspend versioning to disable it.
	Status *StringExpr `json:"Status,omitempty"`
}

// S3VersioningConfigurationList represents a list of S3VersioningConfiguration
type S3VersioningConfigurationList []S3VersioningConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3VersioningConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3VersioningConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3VersioningConfigurationList{item}
		return nil
	}
	list := []S3VersioningConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3VersioningConfigurationList(list)
		return nil
	}
	return err
}

// S3WebsiteConfigurationProperty represents Amazon S3 Website Configuration Property
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration.html
type S3WebsiteConfigurationProperty struct {
	// The name of the error document for the website.
	ErrorDocument *StringExpr `json:"ErrorDocument,omitempty"`

	// The name of the index document for the website.
	IndexDocument *StringExpr `json:"IndexDocument,omitempty"`

	// The redirect behavior for every request to this bucket's website
	// endpoint.
	RedirectAllRequestsTo *S3WebsiteConfigurationRedirectAllRequestsToProperty `json:"RedirectAllRequestsTo,omitempty"`

	// Rules that define when a redirect is applied and the redirect
	// behavior.
	RoutingRules *S3WebsiteConfigurationRoutingRulesProperty `json:"RoutingRules,omitempty"`
}

// S3WebsiteConfigurationPropertyList represents a list of S3WebsiteConfigurationProperty
type S3WebsiteConfigurationPropertyList []S3WebsiteConfigurationProperty

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3WebsiteConfigurationPropertyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3WebsiteConfigurationProperty{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3WebsiteConfigurationPropertyList{item}
		return nil
	}
	list := []S3WebsiteConfigurationProperty{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3WebsiteConfigurationPropertyList(list)
		return nil
	}
	return err
}

// S3WebsiteConfigurationRedirectAllRequestsToProperty represents Amazon S3 Website Configuration Redirect All Requests To Property
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-redirectallrequeststo.html
type S3WebsiteConfigurationRedirectAllRequestsToProperty struct {
	// Name of the host where requests are redirected.
	HostName *StringExpr `json:"HostName,omitempty"`

	// Protocol to use (http or https) when redirecting requests. The default
	// is the protocol that is used in the original request.
	Protocol *StringExpr `json:"Protocol,omitempty"`
}

// S3WebsiteConfigurationRedirectAllRequestsToPropertyList represents a list of S3WebsiteConfigurationRedirectAllRequestsToProperty
type S3WebsiteConfigurationRedirectAllRequestsToPropertyList []S3WebsiteConfigurationRedirectAllRequestsToProperty

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3WebsiteConfigurationRedirectAllRequestsToPropertyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3WebsiteConfigurationRedirectAllRequestsToProperty{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3WebsiteConfigurationRedirectAllRequestsToPropertyList{item}
		return nil
	}
	list := []S3WebsiteConfigurationRedirectAllRequestsToProperty{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3WebsiteConfigurationRedirectAllRequestsToPropertyList(list)
		return nil
	}
	return err
}

// S3WebsiteConfigurationRoutingRulesProperty represents Amazon S3 Website Configuration Routing Rules Property
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules.html
type S3WebsiteConfigurationRoutingRulesProperty struct {
	// Redirect requests to another host, to another page, or with another
	// protocol.
	RedirectRule *S3WebsiteConfigurationRoutingRulesRedirectRuleProperty `json:"RedirectRule,omitempty"`

	// Rules that define when a redirect is applied.
	RoutingRuleCondition *S3WebsiteConfigurationRoutingRulesRoutingRuleConditionProperty `json:"RoutingRuleCondition,omitempty"`
}

// S3WebsiteConfigurationRoutingRulesPropertyList represents a list of S3WebsiteConfigurationRoutingRulesProperty
type S3WebsiteConfigurationRoutingRulesPropertyList []S3WebsiteConfigurationRoutingRulesProperty

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3WebsiteConfigurationRoutingRulesPropertyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3WebsiteConfigurationRoutingRulesProperty{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3WebsiteConfigurationRoutingRulesPropertyList{item}
		return nil
	}
	list := []S3WebsiteConfigurationRoutingRulesProperty{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3WebsiteConfigurationRoutingRulesPropertyList(list)
		return nil
	}
	return err
}

// S3WebsiteConfigurationRoutingRulesRedirectRuleProperty represents Amazon S3 Website Configuration Routing Rules Redirect Rule Property
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-redirectrule.html
type S3WebsiteConfigurationRoutingRulesRedirectRuleProperty struct {
	// Name of the host where requests are redirected.
	HostName *StringExpr `json:"HostName,omitempty"`

	// The HTTP redirect code to use on the response.
	HttpRedirectCode *StringExpr `json:"HttpRedirectCode,omitempty"`

	// The protocol to use in the redirect request.
	Protocol *StringExpr `json:"Protocol,omitempty"`

	// The object key prefix to use in the redirect request. For example, to
	// redirect requests for all pages with the prefix docs/ (objects in the
	// docs/ folder) to the documents/ prefix, you can set the
	// KeyPrefixEquals property in routing condition property to docs/, and
	// set the ReplaceKeyPrefixWith property to documents/.
	ReplaceKeyPrefixWith *StringExpr `json:"ReplaceKeyPrefixWith,omitempty"`

	// The specific object key to use in the redirect request. For example,
	// redirect request to error.html.
	ReplaceKeyWith *StringExpr `json:"ReplaceKeyWith,omitempty"`
}

// S3WebsiteConfigurationRoutingRulesRedirectRulePropertyList represents a list of S3WebsiteConfigurationRoutingRulesRedirectRuleProperty
type S3WebsiteConfigurationRoutingRulesRedirectRulePropertyList []S3WebsiteConfigurationRoutingRulesRedirectRuleProperty

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3WebsiteConfigurationRoutingRulesRedirectRulePropertyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3WebsiteConfigurationRoutingRulesRedirectRuleProperty{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3WebsiteConfigurationRoutingRulesRedirectRulePropertyList{item}
		return nil
	}
	list := []S3WebsiteConfigurationRoutingRulesRedirectRuleProperty{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3WebsiteConfigurationRoutingRulesRedirectRulePropertyList(list)
		return nil
	}
	return err
}

// S3WebsiteConfigurationRoutingRulesRoutingRuleConditionProperty represents Amazon S3 Website Configuration Routing Rules Routing Rule Condition Property
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-routingrulecondition.html
type S3WebsiteConfigurationRoutingRulesRoutingRuleConditionProperty struct {
	// Applies this redirect if the error code equals this value in the event
	// of an error.
	HttpErrorCodeReturnedEquals *StringExpr `json:"HttpErrorCodeReturnedEquals,omitempty"`

	// The object key name prefix when the redirect is applied. For example,
	// to redirect requests for ExamplePage.html, set the key prefix to
	// ExamplePage.html. To redirect request for all pages with the prefix
	// docs/, set the key prefix to docs/, which identifies all objects in
	// the docs/ folder.
	KeyPrefixEquals *StringExpr `json:"KeyPrefixEquals,omitempty"`
}

// S3WebsiteConfigurationRoutingRulesRoutingRuleConditionPropertyList represents a list of S3WebsiteConfigurationRoutingRulesRoutingRuleConditionProperty
type S3WebsiteConfigurationRoutingRulesRoutingRuleConditionPropertyList []S3WebsiteConfigurationRoutingRulesRoutingRuleConditionProperty

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3WebsiteConfigurationRoutingRulesRoutingRuleConditionPropertyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3WebsiteConfigurationRoutingRulesRoutingRuleConditionProperty{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3WebsiteConfigurationRoutingRulesRoutingRuleConditionPropertyList{item}
		return nil
	}
	list := []S3WebsiteConfigurationRoutingRulesRoutingRuleConditionProperty{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3WebsiteConfigurationRoutingRulesRoutingRuleConditionPropertyList(list)
		return nil
	}
	return err
}

// SNSSubscription represents Amazon SNS Subscription Property Type
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-subscription.html
type SNSSubscription struct {
	// The subscription's endpoint (format depends on the protocol). For more
	// information, see the Subscribe Endpoint parameter in the Amazon Simple
	// Notification Service API Reference.
	Endpoint *StringExpr `json:"Endpoint,omitempty"`

	// The subscription's protocol. For more information, see the Subscribe
	// Protocol parameter in the Amazon Simple Notification Service API
	// Reference.
	Protocol *StringExpr `json:"Protocol,omitempty"`
}

// SNSSubscriptionList represents a list of SNSSubscription
type SNSSubscriptionList []SNSSubscription

// UnmarshalJSON sets the object from the provided JSON representation
func (l *SNSSubscriptionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := SNSSubscription{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = SNSSubscriptionList{item}
		return nil
	}
	list := []SNSSubscription{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = SNSSubscriptionList(list)
		return nil
	}
	return err
}

// SQSRedrivePolicy represents Amazon SQS RedrivePolicy
//
// see http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues-redrivepolicy.html
type SQSRedrivePolicy struct {
	// The Amazon Resource Name (ARN) of the dead letter queue to which the
	// messages are sent to after the maxReceiveCount value has been
	// exceeded.
	DeadLetterTargetArn *StringExpr `json:"deadLetterTargetArn,omitempty"`

	// The number of times a message is delivered to the source queue before
	// being sent to the dead letter queue.
	MaxReceiveCount *IntegerExpr `json:"maxReceiveCount,omitempty"`
}

// SQSRedrivePolicyList represents a list of SQSRedrivePolicy
type SQSRedrivePolicyList []SQSRedrivePolicy

// UnmarshalJSON sets the object from the provided JSON representation
func (l *SQSRedrivePolicyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := SQSRedrivePolicy{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = SQSRedrivePolicyList{item}
		return nil
	}
	list := []SQSRedrivePolicy{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = SQSRedrivePolicyList(list)
		return nil
	}
	return err
}

// NewResourceByType returns a new resource object correspoding with the provided type
func NewResourceByType(typeName string) ResourceProperties {
	switch typeName {
	case "AWS::AutoScaling::AutoScalingGroup":
		return &AutoScalingAutoScalingGroup{}
	case "AWS::AutoScaling::LaunchConfiguration":
		return &AutoScalingLaunchConfiguration{}
	case "AWS::AutoScaling::LifecycleHook":
		return &AutoScalingLifecycleHook{}
	case "AWS::AutoScaling::ScalingPolicy":
		return &AutoScalingScalingPolicy{}
	case "AWS::AutoScaling::ScheduledAction":
		return &AutoScalingScheduledAction{}
	case "AWS::CloudFormation::Authentication":
		return &CloudFormationAuthentication{}
	case "AWS::CloudFormation::CustomResource":
		return &CloudFormationCustomResource{}
	case "AWS::CloudFormation::Init":
		return &CloudFormationInit{}
	case "AWS::CloudFormation::Stack":
		return &CloudFormationStack{}
	case "AWS::CloudFormation::WaitCondition":
		return &CloudFormationWaitCondition{}
	case "AWS::CloudFormation::WaitConditionHandle":
		return &CloudFormationWaitConditionHandle{}
	case "AWS::CloudFront::Distribution":
		return &CloudFrontDistribution{}
	case "AWS::CloudTrail::Trail":
		return &CloudTrailTrail{}
	case "AWS::CloudWatch::Alarm":
		return &CloudWatchAlarm{}
	case "AWS::DataPipeline::Pipeline":
		return &DataPipelinePipeline{}
	case "AWS::DynamoDB::Table":
		return &DynamoDBTable{}
	case "AWS::EC2::CustomerGateway":
		return &EC2CustomerGateway{}
	case "AWS::EC2::DHCPOptions":
		return &EC2DHCPOptions{}
	case "AWS::EC2::EIP":
		return &EC2EIP{}
	case "AWS::EC2::EIPAssociation":
		return &EC2EIPAssociation{}
	case "AWS::EC2::Instance":
		return &EC2Instance{}
	case "AWS::EC2::InternetGateway":
		return &EC2InternetGateway{}
	case "AWS::EC2::NetworkAcl":
		return &EC2NetworkAcl{}
	case "AWS::EC2::NetworkAclEntry":
		return &EC2NetworkAclEntry{}
	case "AWS::EC2::NetworkInterface":
		return &EC2NetworkInterface{}
	case "AWS::EC2::NetworkInterfaceAttachment":
		return &EC2NetworkInterfaceAttachment{}
	case "AWS::EC2::Route":
		return &EC2Route{}
	case "AWS::EC2::RouteTable":
		return &EC2RouteTable{}
	case "AWS::EC2::SecurityGroup":
		return &EC2SecurityGroup{}
	case "AWS::EC2::SecurityGroupEgress":
		return &EC2SecurityGroupEgress{}
	case "AWS::EC2::SecurityGroupIngress":
		return &EC2SecurityGroupIngress{}
	case "AWS::EC2::Subnet":
		return &EC2Subnet{}
	case "AWS::EC2::SubnetNetworkAclAssociation":
		return &EC2SubnetNetworkAclAssociation{}
	case "AWS::EC2::SubnetRouteTableAssociation":
		return &EC2SubnetRouteTableAssociation{}
	case "AWS::EC2::Volume":
		return &EC2Volume{}
	case "AWS::EC2::VolumeAttachment":
		return &EC2VolumeAttachment{}
	case "AWS::EC2::VPC":
		return &EC2VPC{}
	case "AWS::EC2::VPCDHCPOptionsAssociation":
		return &EC2VPCDHCPOptionsAssociation{}
	case "AWS::EC2::VPCGatewayAttachment":
		return &EC2VPCGatewayAttachment{}
	case "AWS::EC2::VPCPeeringConnection":
		return &EC2VPCPeeringConnection{}
	case "AWS::EC2::VPNConnection":
		return &EC2VPNConnection{}
	case "AWS::EC2::VPNConnectionRoute":
		return &EC2VPNConnectionRoute{}
	case "AWS::EC2::VPNGateway":
		return &EC2VPNGateway{}
	case "AWS::EC2::VPNGatewayRoutePropagation":
		return &EC2VPNGatewayRoutePropagation{}
	case "AWS::ECS::Cluster":
		return &ECSCluster{}
	case "AWS::ECS::Service":
		return &ECSService{}
	case "AWS::ECS::TaskDefinition":
		return &ECSTaskDefinition{}
	case "AWS::ElastiCache::CacheCluster":
		return &ElastiCacheCacheCluster{}
	case "AWS::ElastiCache::ParameterGroup":
		return &ElastiCacheParameterGroup{}
	case "AWS::ElastiCache::ReplicationGroup":
		return &ElastiCacheReplicationGroup{}
	case "AWS::ElastiCache::SecurityGroup":
		return &ElastiCacheSecurityGroup{}
	case "AWS::ElastiCache::SecurityGroupIngress":
		return &ElastiCacheSecurityGroupIngress{}
	case "AWS::ElastiCache::SubnetGroup ":
		return &ElastiCacheSubnetGroup{}
	case "AWS::ElasticBeanstalk::Application":
		return &ElasticBeanstalkApplication{}
	case "AWS::ElasticBeanstalk::ApplicationVersion":
		return &ElasticBeanstalkApplicationVersion{}
	case "AWS::ElasticBeanstalk::ConfigurationTemplate":
		return &ElasticBeanstalkConfigurationTemplate{}
	case "AWS::ElasticBeanstalk::Environment":
		return &ElasticBeanstalkEnvironment{}
	case "AWS::ElasticLoadBalancing::LoadBalancer":
		return &ElasticLoadBalancingLoadBalancer{}
	case "AWS::IAM::AccessKey":
		return &IAMAccessKey{}
	case "AWS::IAM::Group":
		return &IAMGroup{}
	case "AWS::IAM::InstanceProfile":
		return &IAMInstanceProfile{}
	case "AWS::IAM::ManagedPolicy":
		return &IAMManagedPolicy{}
	case "AWS::IAM::Policy":
		return &IAMPolicy{}
	case "AWS::IAM::Role":
		return &IAMRole{}
	case "AWS::IAM::User":
		return &IAMUser{}
	case "AWS::IAM::UserToGroupAddition":
		return &IAMUserToGroupAddition{}
	case "AWS::Kinesis::Stream":
		return &KinesisStream{}
	case "AWS::Lambda::Function":
		return &LambdaFunction{}
	case "AWS::Logs::LogGroup":
		return &LogsLogGroup{}
	case "AWS::Logs::MetricFilter":
		return &LogsMetricFilter{}
	case "AWS::OpsWorks::App":
		return &OpsWorksApp{}
	case "AWS::OpsWorks::ElasticLoadBalancerAttachment":
		return &OpsWorksElasticLoadBalancerAttachment{}
	case "AWS::OpsWorks::Instance":
		return &OpsWorksInstance{}
	case "AWS::OpsWorks::Layer":
		return &OpsWorksLayer{}
	case "AWS::OpsWorks::Stack":
		return &OpsWorksStack{}
	case "AWS::Redshift::Cluster":
		return &RedshiftCluster{}
	case "AWS::Redshift::ClusterParameterGroup":
		return &RedshiftClusterParameterGroup{}
	case "AWS::Redshift::ClusterSecurityGroup":
		return &RedshiftClusterSecurityGroup{}
	case "AWS::Redshift::ClusterSecurityGroupIngress":
		return &RedshiftClusterSecurityGroupIngress{}
	case "AWS::Redshift::ClusterSubnetGroup":
		return &RedshiftClusterSubnetGroup{}
	case "AWS::RDS::DBInstance":
		return &RDSDBInstance{}
	case "AWS::RDS::DBParameterGroup":
		return &RDSDBParameterGroup{}
	case "AWS::RDS::DBSubnetGroup":
		return &RDSDBSubnetGroup{}
	case "AWS::RDS::DBSecurityGroup":
		return &RDSDBSecurityGroup{}
	case "AWS::RDS::DBSecurityGroupIngress":
		return &RDSDBSecurityGroupIngress{}
	case "AWS::RDS::EventSubscription":
		return &RDSEventSubscription{}
	case "AWS::RDS::OptionGroup":
		return &RDSOptionGroup{}
	case "AWS::Route53::HealthCheck":
		return &Route53HealthCheck{}
	case "AWS::Route53::HostedZone":
		return &Route53HostedZone{}
	case "AWS::Route53::RecordSet":
		return &Route53RecordSet{}
	case "AWS::Route53::RecordSetGroup":
		return &Route53RecordSetGroup{}
	case "AWS::S3::Bucket":
		return &S3Bucket{}
	case "AWS::S3::BucketPolicy":
		return &S3BucketPolicy{}
	case "AWS::SDB::Domain":
		return &SDBDomain{}
	case "AWS::SNS::Topic":
		return &SNSTopic{}
	case "AWS::SNS::TopicPolicy":
		return &SNSTopicPolicy{}
	case "AWS::SQS::Queue":
		return &SQSQueue{}
	case "AWS::SQS::QueuePolicy":
		return &SQSQueuePolicy{}
	}
	return nil
}
