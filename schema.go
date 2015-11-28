package cloudformation

import "time"
import "encoding/json"

type AWSAutoScalingAutoScalingGroup struct {
  AvailabilityZones *StringListExpression `json:"AvailabilityZones,omitempty"`  // List of strings
  Cooldown *StringExpression `json:"Cooldown,omitempty"`  // String
  DesiredCapacity *StringExpression `json:"DesiredCapacity,omitempty"`  // String
  HealthCheckGracePeriod *Integer `json:"HealthCheckGracePeriod,omitempty"`  // Integer
  HealthCheckType *StringExpression `json:"HealthCheckType,omitempty"`  // String
  InstanceId *StringExpression `json:"InstanceId,omitempty"`  // String
  LaunchConfigurationName *StringExpression `json:"LaunchConfigurationName,omitempty"`  // String
  LoadBalancerNames *StringListExpression `json:"LoadBalancerNames,omitempty"`  // List of strings
  MaxSize *StringExpression `json:"MaxSize,omitempty"`  // String
  MetricsCollection *AutoScalingMetricsCollectionList `json:"MetricsCollection,omitempty"`  // A list of Auto Scaling MetricsCollection
  MinSize *StringExpression `json:"MinSize,omitempty"`  // String
  NotificationConfigurations *AutoScalingNotificationConfigurationsList `json:"NotificationConfigurations,omitempty"`  // List of Auto Scaling NotificationConfigurations
  PlacementGroup *StringExpression `json:"PlacementGroup,omitempty"`  // String
  Tags *AutoScalingTagsList `json:"Tags,omitempty"`  // List of Auto Scaling Tags
  TerminationPolicies *StringListExpression `json:"TerminationPolicies,omitempty"`  // List of strings
  VPCZoneIdentifier *StringListExpression `json:"VPCZoneIdentifier,omitempty"`  // List of strings
}

type AWSAutoScalingAutoScalingGroupList []AWSAutoScalingAutoScalingGroup

func (l *AWSAutoScalingAutoScalingGroupList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSAutoScalingAutoScalingGroup{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSAutoScalingAutoScalingGroupList{item}
		return nil
	}
	list := []AWSAutoScalingAutoScalingGroup{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSAutoScalingAutoScalingGroupList(list)
		return nil
	}
	return err
}


type AWSAutoScalingLaunchConfiguration struct {
  AssociatePublicIpAddress *Bool `json:"AssociatePublicIpAddress,omitempty"`  // Boolean
  BlockDeviceMappings *AWSCloudFormationAutoScalingBlockDeviceMappingList `json:"BlockDeviceMappings,omitempty"`  // A list of BlockDeviceMappings
  ClassicLinkVPCId *StringExpression `json:"ClassicLinkVPCId,omitempty"`  // String
  ClassicLinkVPCSecurityGroups *StringListExpression `json:"ClassicLinkVPCSecurityGroups,omitempty"`  // List of strings
  EbsOptimized *Bool `json:"EbsOptimized,omitempty"`  // Boolean
  IamInstanceProfile *StringExpression `json:"IamInstanceProfile,omitempty"`  // String (1–1600 chars)
  ImageId *StringExpression `json:"ImageId,omitempty"`  // String
  InstanceId *StringExpression `json:"InstanceId,omitempty"`  // String
  InstanceMonitoring *Bool `json:"InstanceMonitoring,omitempty"`  // Boolean
  InstanceType *StringExpression `json:"InstanceType,omitempty"`  // String
  KernelId *StringExpression `json:"KernelId,omitempty"`  // String
  KeyName *StringExpression `json:"KeyName,omitempty"`  // String
  PlacementTenancy *StringExpression `json:"PlacementTenancy,omitempty"`  // String
  RamDiskId *StringExpression `json:"RamDiskId,omitempty"`  // String
  SecurityGroups interface{} `json:"SecurityGroups,omitempty"`  // A list of EC2 security groups
  SpotPrice *StringExpression `json:"SpotPrice,omitempty"`  // String
  UserData *StringExpression `json:"UserData,omitempty"`  // String
}

type AWSAutoScalingLaunchConfigurationList []AWSAutoScalingLaunchConfiguration

func (l *AWSAutoScalingLaunchConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSAutoScalingLaunchConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSAutoScalingLaunchConfigurationList{item}
		return nil
	}
	list := []AWSAutoScalingLaunchConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSAutoScalingLaunchConfigurationList(list)
		return nil
	}
	return err
}


type AWSAutoScalingLifecycleHook struct {
  AutoScalingGroupName *StringExpression `json:"AutoScalingGroupName,omitempty"`  // String
  DefaultResult *StringExpression `json:"DefaultResult,omitempty"`  // String
  HeartbeatTimeout *Integer `json:"HeartbeatTimeout,omitempty"`  // Integer
  LifecycleTransition *StringExpression `json:"LifecycleTransition,omitempty"`  // String
  NotificationMetadata *StringExpression `json:"NotificationMetadata,omitempty"`  // String
  NotificationTargetARN *StringExpression `json:"NotificationTargetARN,omitempty"`  // String
  RoleARN *StringExpression `json:"RoleARN,omitempty"`  // String
}

type AWSAutoScalingLifecycleHookList []AWSAutoScalingLifecycleHook

func (l *AWSAutoScalingLifecycleHookList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSAutoScalingLifecycleHook{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSAutoScalingLifecycleHookList{item}
		return nil
	}
	list := []AWSAutoScalingLifecycleHook{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSAutoScalingLifecycleHookList(list)
		return nil
	}
	return err
}


type AWSAutoScalingScalingPolicy struct {
  AdjustmentType *StringExpression `json:"AdjustmentType,omitempty"`  // String
  AutoScalingGroupName *StringExpression `json:"AutoScalingGroupName,omitempty"`  // String
  Cooldown *StringExpression `json:"Cooldown,omitempty"`  // String
  MinAdjustmentStep *Integer `json:"MinAdjustmentStep,omitempty"`  // Integer
  ScalingAdjustment *StringExpression `json:"ScalingAdjustment,omitempty"`  // String
}

type AWSAutoScalingScalingPolicyList []AWSAutoScalingScalingPolicy

func (l *AWSAutoScalingScalingPolicyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSAutoScalingScalingPolicy{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSAutoScalingScalingPolicyList{item}
		return nil
	}
	list := []AWSAutoScalingScalingPolicy{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSAutoScalingScalingPolicyList(list)
		return nil
	}
	return err
}


type AWSAutoScalingScheduledAction struct {
  AutoScalingGroupName *StringExpression `json:"AutoScalingGroupName,omitempty"`  // String
  DesiredCapacity *Integer `json:"DesiredCapacity,omitempty"`  // Integer
  EndTime time.Time `json:"EndTime,omitempty"`  // Time stamp
  MaxSize *Integer `json:"MaxSize,omitempty"`  // Integer
  MinSize *Integer `json:"MinSize,omitempty"`  // Integer
  Recurrence *StringExpression `json:"Recurrence,omitempty"`  // String
  StartTime time.Time `json:"StartTime,omitempty"`  // Time stamp
}

type AWSAutoScalingScheduledActionList []AWSAutoScalingScheduledAction

func (l *AWSAutoScalingScheduledActionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSAutoScalingScheduledAction{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSAutoScalingScheduledActionList{item}
		return nil
	}
	list := []AWSAutoScalingScheduledAction{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSAutoScalingScheduledActionList(list)
		return nil
	}
	return err
}


type AWSCloudFormationAuthentication struct {
  AccessKeyId *StringExpression `json:"accessKeyId,omitempty"`  // String
  Buckets *StringListExpression `json:"buckets,omitempty"`  // List of strings
  Password *StringExpression `json:"password,omitempty"`  // String
  SecretKey *StringExpression `json:"secretKey,omitempty"`  // String
  Type *StringExpression `json:"type,omitempty"`  // String Valid values are "basic" or "S3"
  Uris *StringListExpression `json:"uris,omitempty"`  // List of strings
  Username *StringExpression `json:"username,omitempty"`  // String
  RoleName *StringExpression `json:"roleName,omitempty"`  // String
}

type AWSCloudFormationAuthenticationList []AWSCloudFormationAuthentication

func (l *AWSCloudFormationAuthenticationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSCloudFormationAuthentication{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSCloudFormationAuthenticationList{item}
		return nil
	}
	list := []AWSCloudFormationAuthentication{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSCloudFormationAuthenticationList(list)
		return nil
	}
	return err
}


type AWSCloudFormationCustomResource struct {
  ServiceToken *StringExpression `json:"ServiceToken,omitempty"`  // String
}

type AWSCloudFormationCustomResourceList []AWSCloudFormationCustomResource

func (l *AWSCloudFormationCustomResourceList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSCloudFormationCustomResource{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSCloudFormationCustomResourceList{item}
		return nil
	}
	list := []AWSCloudFormationCustomResource{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSCloudFormationCustomResourceList(list)
		return nil
	}
	return err
}


type AWSCloudFormationInit struct {
}

type AWSCloudFormationInitList []AWSCloudFormationInit

func (l *AWSCloudFormationInitList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSCloudFormationInit{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSCloudFormationInitList{item}
		return nil
	}
	list := []AWSCloudFormationInit{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSCloudFormationInitList(list)
		return nil
	}
	return err
}


type AWSCloudFormationStack struct {
  NotificationARNs *StringListExpression `json:"NotificationARNs,omitempty"`  // List of strings
  Parameters *CloudFormationStackParameters `json:"Parameters,omitempty"`  // CloudFormation Stack Parameters Property Type
  TemplateURL *StringExpression `json:"TemplateURL,omitempty"`  // String
  TimeoutInMinutes *StringExpression `json:"TimeoutInMinutes,omitempty"`  // String
}

type AWSCloudFormationStackList []AWSCloudFormationStack

func (l *AWSCloudFormationStackList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSCloudFormationStack{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSCloudFormationStackList{item}
		return nil
	}
	list := []AWSCloudFormationStack{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSCloudFormationStackList(list)
		return nil
	}
	return err
}


type AWSCloudFormationWaitCondition struct {
  Count *StringExpression `json:"Count,omitempty"`  // String
  Handle *StringExpression `json:"Handle,omitempty"`  // String
  Timeout *StringExpression `json:"Timeout,omitempty"`  // String
}

type AWSCloudFormationWaitConditionList []AWSCloudFormationWaitCondition

func (l *AWSCloudFormationWaitConditionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSCloudFormationWaitCondition{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSCloudFormationWaitConditionList{item}
		return nil
	}
	list := []AWSCloudFormationWaitCondition{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSCloudFormationWaitConditionList(list)
		return nil
	}
	return err
}


type AWSCloudFormationWaitConditionHandle struct {
}

type AWSCloudFormationWaitConditionHandleList []AWSCloudFormationWaitConditionHandle

func (l *AWSCloudFormationWaitConditionHandleList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSCloudFormationWaitConditionHandle{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSCloudFormationWaitConditionHandleList{item}
		return nil
	}
	list := []AWSCloudFormationWaitConditionHandle{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSCloudFormationWaitConditionHandleList(list)
		return nil
	}
	return err
}


type AWSCloudFrontDistribution struct {
  DistributionConfig *CloudFrontDistributionConfig `json:"DistributionConfig,omitempty"`  // DistributionConfig type
}

type AWSCloudFrontDistributionList []AWSCloudFrontDistribution

func (l *AWSCloudFrontDistributionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSCloudFrontDistribution{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSCloudFrontDistributionList{item}
		return nil
	}
	list := []AWSCloudFrontDistribution{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSCloudFrontDistributionList(list)
		return nil
	}
	return err
}


type AWSCloudTrailTrail struct {
  IncludeGlobalServiceEvents *Bool `json:"IncludeGlobalServiceEvents,omitempty"`  // Boolean
  IsLogging *Bool `json:"IsLogging,omitempty"`  // Boolean
  S3BucketName *StringExpression `json:"S3BucketName,omitempty"`  // String
  S3KeyPrefix *StringExpression `json:"S3KeyPrefix,omitempty"`  // String
  SnsTopicName *StringExpression `json:"SnsTopicName,omitempty"`  // String
}

type AWSCloudTrailTrailList []AWSCloudTrailTrail

func (l *AWSCloudTrailTrailList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSCloudTrailTrail{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSCloudTrailTrailList{item}
		return nil
	}
	list := []AWSCloudTrailTrail{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSCloudTrailTrailList(list)
		return nil
	}
	return err
}


type AWSCloudWatchAlarm struct {
  ActionsEnabled *Bool `json:"ActionsEnabled,omitempty"`  // Boolean
  AlarmActions *StringListExpression `json:"AlarmActions,omitempty"`  // List of strings
  AlarmDescription *StringExpression `json:"AlarmDescription,omitempty"`  // String
  AlarmName *StringExpression `json:"AlarmName,omitempty"`  // String
  ComparisonOperator *StringExpression `json:"ComparisonOperator,omitempty"`  // String
  Dimensions *CloudWatchMetricDimensionList `json:"Dimensions,omitempty"`  // List of Metric Dimension
  EvaluationPeriods *StringExpression `json:"EvaluationPeriods,omitempty"`  // String
  InsufficientDataActions *StringListExpression `json:"InsufficientDataActions,omitempty"`  // List of strings
  MetricName *StringExpression `json:"MetricName,omitempty"`  // String
  Namespace *StringExpression `json:"Namespace,omitempty"`  // String
  OKActions *StringListExpression `json:"OKActions,omitempty"`  // List of strings
  Period *StringExpression `json:"Period,omitempty"`  // String
  Statistic *StringExpression `json:"Statistic,omitempty"`  // String
  Threshold *StringExpression `json:"Threshold,omitempty"`  // String
  Unit *StringExpression `json:"Unit,omitempty"`  // String
}

type AWSCloudWatchAlarmList []AWSCloudWatchAlarm

func (l *AWSCloudWatchAlarmList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSCloudWatchAlarm{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSCloudWatchAlarmList{item}
		return nil
	}
	list := []AWSCloudWatchAlarm{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSCloudWatchAlarmList(list)
		return nil
	}
	return err
}


type AWSDataPipelinePipeline struct {
  Activate *Bool `json:"Activate,omitempty"`  // Boolean
  Description *StringExpression `json:"Description,omitempty"`  // String
  Name *StringExpression `json:"Name,omitempty"`  // String
  ParameterObjects *AWSDataPipelinePipelineParameterObjects `json:"ParameterObjects,omitempty"`  // AWS Data Pipeline Pipeline ParameterObjects
  ParameterValues *AWSDataPipelinePipelineParameterValues `json:"ParameterValues,omitempty"`  // AWS Data Pipeline Pipeline ParameterValues
  PipelineObjects *AWSDataPipelinePipelineObjectsList `json:"PipelineObjects,omitempty"`  // A list of AWS Data Pipeline PipelineObjects
  PipelineTags *AWSDataPipelinePipelinePipelineTagsList `json:"PipelineTags,omitempty"`  // AWS Data Pipeline Pipeline PipelineTags
}

type AWSDataPipelinePipelineList []AWSDataPipelinePipeline

func (l *AWSDataPipelinePipelineList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSDataPipelinePipeline{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSDataPipelinePipelineList{item}
		return nil
	}
	list := []AWSDataPipelinePipeline{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSDataPipelinePipelineList(list)
		return nil
	}
	return err
}


type AWSDynamoDBTable struct {
  AttributeDefinitions *DynamoDBAttributeDefinitionsList `json:"AttributeDefinitions,omitempty"`  // DynamoDB Attribute Definitions
  GlobalSecondaryIndexes *DynamoDBGlobalSecondaryIndexes `json:"GlobalSecondaryIndexes,omitempty"`  // DynamoDB Global Secondary Indexes
  KeySchema *DynamoDBKeySchema `json:"KeySchema,omitempty"`  // DynamoDB Key Schema
  LocalSecondaryIndexes *DynamoDBLocalSecondaryIndexes `json:"LocalSecondaryIndexes,omitempty"`  // DynamoDB Local Secondary Indexes
  ProvisionedThroughput *DynamoDBProvisionedThroughput `json:"ProvisionedThroughput,omitempty"`  // DynamoDB Provisioned Throughput
  TableName *StringExpression `json:"TableName,omitempty"`  // String
}

type AWSDynamoDBTableList []AWSDynamoDBTable

func (l *AWSDynamoDBTableList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSDynamoDBTable{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSDynamoDBTableList{item}
		return nil
	}
	list := []AWSDynamoDBTable{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSDynamoDBTableList(list)
		return nil
	}
	return err
}


type AWSEC2CustomerGateway struct {
  BgpAsn Integer `json:"BgpAsn,omitempty"`  // Number BgpAsn is always an integer value
  IpAddress *StringExpression `json:"IpAddress,omitempty"`  // String
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
  Type *StringExpression `json:"Type,omitempty"`  // String
}

type AWSEC2CustomerGatewayList []AWSEC2CustomerGateway

func (l *AWSEC2CustomerGatewayList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSEC2CustomerGateway{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSEC2CustomerGatewayList{item}
		return nil
	}
	list := []AWSEC2CustomerGateway{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSEC2CustomerGatewayList(list)
		return nil
	}
	return err
}


type AWSEC2DHCPOptions struct {
  DomainName *StringExpression `json:"DomainName,omitempty"`  // String
  DomainNameServers *StringListExpression `json:"DomainNameServers,omitempty"`  // List of strings
  NetbiosNameServers *StringListExpression `json:"NetbiosNameServers,omitempty"`  // List of strings
  NetbiosNodeType interface{} `json:"NetbiosNodeType,omitempty"`  // List of numbers
  NtpServers *StringListExpression `json:"NtpServers,omitempty"`  // List of strings
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
}

type AWSEC2DHCPOptionsList []AWSEC2DHCPOptions

func (l *AWSEC2DHCPOptionsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSEC2DHCPOptions{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSEC2DHCPOptionsList{item}
		return nil
	}
	list := []AWSEC2DHCPOptions{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSEC2DHCPOptionsList(list)
		return nil
	}
	return err
}


type AWSEC2EIP struct {
  InstanceId *StringExpression `json:"InstanceId,omitempty"`  // String
  Domain *StringExpression `json:"Domain,omitempty"`  // String
}

type AWSEC2EIPList []AWSEC2EIP

func (l *AWSEC2EIPList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSEC2EIP{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSEC2EIPList{item}
		return nil
	}
	list := []AWSEC2EIP{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSEC2EIPList(list)
		return nil
	}
	return err
}


type AWSEC2EIPAssociation struct {
  AllocationId *StringExpression `json:"AllocationId,omitempty"`  // String
  EIP *StringExpression `json:"EIP,omitempty"`  // String
  InstanceId *StringExpression `json:"InstanceId,omitempty"`  // String
  NetworkInterfaceId *StringExpression `json:"NetworkInterfaceId,omitempty"`  // String
  PrivateIpAddress *StringExpression `json:"PrivateIpAddress,omitempty"`  // String
}

type AWSEC2EIPAssociationList []AWSEC2EIPAssociation

func (l *AWSEC2EIPAssociationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSEC2EIPAssociation{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSEC2EIPAssociationList{item}
		return nil
	}
	list := []AWSEC2EIPAssociation{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSEC2EIPAssociationList(list)
		return nil
	}
	return err
}


type AWSEC2Instance struct {
  AvailabilityZone *StringExpression `json:"AvailabilityZone,omitempty"`  // String
  BlockDeviceMappings *AmazonEC2BlockDeviceMappingPropertyList `json:"BlockDeviceMappings,omitempty"`  // A list of Amazon EC2 Block Device Mapping Property
  DisableApiTermination *Bool `json:"DisableApiTermination,omitempty"`  // Boolean
  EbsOptimized *Bool `json:"EbsOptimized,omitempty"`  // Boolean
  IamInstanceProfile *StringExpression `json:"IamInstanceProfile,omitempty"`  // String
  ImageId *StringExpression `json:"ImageId,omitempty"`  // String
  InstanceInitiatedShutdownBehavior *StringExpression `json:"InstanceInitiatedShutdownBehavior,omitempty"`  // String
  InstanceType *StringExpression `json:"InstanceType,omitempty"`  // String
  KernelId *StringExpression `json:"KernelId,omitempty"`  // String
  KeyName *StringExpression `json:"KeyName,omitempty"`  // String
  Monitoring *Bool `json:"Monitoring,omitempty"`  // Boolean
  NetworkInterfaces *EC2NetworkInterfaceEmbeddedList `json:"NetworkInterfaces,omitempty"`  // A list of EC2 NetworkInterface Embedded Property Type
  PlacementGroupName *StringExpression `json:"PlacementGroupName,omitempty"`  // String
  PrivateIpAddress *StringExpression `json:"PrivateIpAddress,omitempty"`  // String
  RamdiskId *StringExpression `json:"RamdiskId,omitempty"`  // String
  SecurityGroupIds *StringListExpression `json:"SecurityGroupIds,omitempty"`  // List of strings
  SecurityGroups *StringListExpression `json:"SecurityGroups,omitempty"`  // List of strings
  SourceDestCheck *Bool `json:"SourceDestCheck,omitempty"`  // Boolean
  SubnetId *StringExpression `json:"SubnetId,omitempty"`  // String
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
  Tenancy *StringExpression `json:"Tenancy,omitempty"`  // String
  UserData *StringExpression `json:"UserData,omitempty"`  // String
  Volumes *EC2MountPointList `json:"Volumes,omitempty"`  // A list of EC2 MountPoints
}

type AWSEC2InstanceList []AWSEC2Instance

func (l *AWSEC2InstanceList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSEC2Instance{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSEC2InstanceList{item}
		return nil
	}
	list := []AWSEC2Instance{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSEC2InstanceList(list)
		return nil
	}
	return err
}


type AWSEC2InternetGateway struct {
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
}

type AWSEC2InternetGatewayList []AWSEC2InternetGateway

func (l *AWSEC2InternetGatewayList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSEC2InternetGateway{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSEC2InternetGatewayList{item}
		return nil
	}
	list := []AWSEC2InternetGateway{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSEC2InternetGatewayList(list)
		return nil
	}
	return err
}


type AWSEC2NetworkAcl struct {
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
  VpcId *StringExpression `json:"VpcId,omitempty"`  // String
}

type AWSEC2NetworkAclList []AWSEC2NetworkAcl

func (l *AWSEC2NetworkAclList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSEC2NetworkAcl{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSEC2NetworkAclList{item}
		return nil
	}
	list := []AWSEC2NetworkAcl{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSEC2NetworkAclList(list)
		return nil
	}
	return err
}


type AWSEC2NetworkAclEntry struct {
  CidrBlock *StringExpression `json:"CidrBlock,omitempty"`  // String
  Egress *Bool `json:"Egress,omitempty"`  // Boolean
  Icmp *EC2ICMP `json:"Icmp,omitempty"`  // EC2 ICMP Property Type
  NetworkAclId *StringExpression `json:"NetworkAclId,omitempty"`  // String
  PortRange *EC2PortRange `json:"PortRange,omitempty"`  // EC2 PortRange Property Type
  Protocol *Integer `json:"Protocol,omitempty"`  // Number
  RuleAction *StringExpression `json:"RuleAction,omitempty"`  // String
  RuleNumber *Integer `json:"RuleNumber,omitempty"`  // Number
}

type AWSEC2NetworkAclEntryList []AWSEC2NetworkAclEntry

func (l *AWSEC2NetworkAclEntryList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSEC2NetworkAclEntry{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSEC2NetworkAclEntryList{item}
		return nil
	}
	list := []AWSEC2NetworkAclEntry{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSEC2NetworkAclEntryList(list)
		return nil
	}
	return err
}


type AWSEC2NetworkInterface struct {
  Description *StringExpression `json:"Description,omitempty"`  // String
  GroupSet *StringListExpression `json:"GroupSet,omitempty"`  // List of strings
  PrivateIpAddress *StringExpression `json:"PrivateIpAddress,omitempty"`  // String
  PrivateIpAddresses *EC2NetworkInterfacePrivateIPSpecificationList `json:"PrivateIpAddresses,omitempty"`  // list of PrivateIpAddressSpecification
  SecondaryPrivateIpAddressCount *Integer `json:"SecondaryPrivateIpAddressCount,omitempty"`  // Integer
  SourceDestCheck *Bool `json:"SourceDestCheck,omitempty"`  // Boolean
  SubnetId *StringExpression `json:"SubnetId,omitempty"`  // String
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
}

type AWSEC2NetworkInterfaceList []AWSEC2NetworkInterface

func (l *AWSEC2NetworkInterfaceList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSEC2NetworkInterface{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSEC2NetworkInterfaceList{item}
		return nil
	}
	list := []AWSEC2NetworkInterface{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSEC2NetworkInterfaceList(list)
		return nil
	}
	return err
}


type AWSEC2NetworkInterfaceAttachment struct {
  DeleteOnTermination *Bool `json:"DeleteOnTermination,omitempty"`  // Boolean
  DeviceIndex *StringExpression `json:"DeviceIndex,omitempty"`  // String
  InstanceId *StringExpression `json:"InstanceId,omitempty"`  // String
  NetworkInterfaceId *StringExpression `json:"NetworkInterfaceId,omitempty"`  // String
}

type AWSEC2NetworkInterfaceAttachmentList []AWSEC2NetworkInterfaceAttachment

func (l *AWSEC2NetworkInterfaceAttachmentList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSEC2NetworkInterfaceAttachment{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSEC2NetworkInterfaceAttachmentList{item}
		return nil
	}
	list := []AWSEC2NetworkInterfaceAttachment{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSEC2NetworkInterfaceAttachmentList(list)
		return nil
	}
	return err
}


type AWSEC2Route struct {
  DestinationCidrBlock *StringExpression `json:"DestinationCidrBlock,omitempty"`  // String
  GatewayId *StringExpression `json:"GatewayId,omitempty"`  // String
  InstanceId *StringExpression `json:"InstanceId,omitempty"`  // String
  NetworkInterfaceId *StringExpression `json:"NetworkInterfaceId,omitempty"`  // String
  RouteTableId *StringExpression `json:"RouteTableId,omitempty"`  // String
  VpcPeeringConnectionId *StringExpression `json:"VpcPeeringConnectionId,omitempty"`  // String
}

type AWSEC2RouteList []AWSEC2Route

func (l *AWSEC2RouteList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSEC2Route{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSEC2RouteList{item}
		return nil
	}
	list := []AWSEC2Route{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSEC2RouteList(list)
		return nil
	}
	return err
}


type AWSEC2RouteTable struct {
  VpcId *StringExpression `json:"VpcId,omitempty"`  // String
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
}

type AWSEC2RouteTableList []AWSEC2RouteTable

func (l *AWSEC2RouteTableList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSEC2RouteTable{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSEC2RouteTableList{item}
		return nil
	}
	list := []AWSEC2RouteTable{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSEC2RouteTableList(list)
		return nil
	}
	return err
}


type AWSEC2SecurityGroup struct {
  GroupDescription *StringExpression `json:"GroupDescription,omitempty"`  // String
  SecurityGroupEgress *EC2SecurityGroupRuleList `json:"SecurityGroupEgress,omitempty"`  // EC2 Security Group Rule
  SecurityGroupIngress *EC2SecurityGroupRuleList `json:"SecurityGroupIngress,omitempty"`  // EC2 Security Group Rule
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
  VpcId *StringExpression `json:"VpcId,omitempty"`  // String
}

type AWSEC2SecurityGroupList []AWSEC2SecurityGroup

func (l *AWSEC2SecurityGroupList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSEC2SecurityGroup{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSEC2SecurityGroupList{item}
		return nil
	}
	list := []AWSEC2SecurityGroup{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSEC2SecurityGroupList(list)
		return nil
	}
	return err
}


type AWSEC2SecurityGroupEgress struct {
  CidrIp *StringExpression `json:"CidrIp,omitempty"`  // String
  DestinationSecurityGroupId *StringExpression `json:"DestinationSecurityGroupId,omitempty"`  // String
  FromPort *Integer `json:"FromPort,omitempty"`  // Integer
  GroupId *StringExpression `json:"GroupId,omitempty"`  // String
  IpProtocol *StringExpression `json:"IpProtocol,omitempty"`  // String
  ToPort *Integer `json:"ToPort,omitempty"`  // Integer
}

type AWSEC2SecurityGroupEgressList []AWSEC2SecurityGroupEgress

func (l *AWSEC2SecurityGroupEgressList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSEC2SecurityGroupEgress{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSEC2SecurityGroupEgressList{item}
		return nil
	}
	list := []AWSEC2SecurityGroupEgress{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSEC2SecurityGroupEgressList(list)
		return nil
	}
	return err
}


type AWSEC2SecurityGroupIngress struct {
  CidrIp *StringExpression `json:"CidrIp,omitempty"`  // String
  FromPort *Integer `json:"FromPort,omitempty"`  // Integer
  GroupId *StringExpression `json:"GroupId,omitempty"`  // String
  GroupName *StringExpression `json:"GroupName,omitempty"`  // String
  IpProtocol *StringExpression `json:"IpProtocol,omitempty"`  // String
  SourceSecurityGroupId *StringExpression `json:"SourceSecurityGroupId,omitempty"`  // String
  SourceSecurityGroupName *StringExpression `json:"SourceSecurityGroupName,omitempty"`  // String
  SourceSecurityGroupOwnerId *StringExpression `json:"SourceSecurityGroupOwnerId,omitempty"`  // String
  ToPort *Integer `json:"ToPort,omitempty"`  // Integer
}

type AWSEC2SecurityGroupIngressList []AWSEC2SecurityGroupIngress

func (l *AWSEC2SecurityGroupIngressList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSEC2SecurityGroupIngress{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSEC2SecurityGroupIngressList{item}
		return nil
	}
	list := []AWSEC2SecurityGroupIngress{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSEC2SecurityGroupIngressList(list)
		return nil
	}
	return err
}


type AWSEC2Subnet struct {
  AvailabilityZone *StringExpression `json:"AvailabilityZone,omitempty"`  // String
  CidrBlock *StringExpression `json:"CidrBlock,omitempty"`  // String
  MapPublicIpOnLaunch *Bool `json:"MapPublicIpOnLaunch,omitempty"`  // Boolean
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
  VpcId interface{} `json:"VpcId,omitempty"`  // Ref ID
}

type AWSEC2SubnetList []AWSEC2Subnet

func (l *AWSEC2SubnetList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSEC2Subnet{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSEC2SubnetList{item}
		return nil
	}
	list := []AWSEC2Subnet{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSEC2SubnetList(list)
		return nil
	}
	return err
}


type AWSEC2SubnetNetworkAclAssociation struct {
  SubnetId *StringExpression `json:"SubnetId,omitempty"`  // String
  NetworkAclId *StringExpression `json:"NetworkAclId,omitempty"`  // String
}

type AWSEC2SubnetNetworkAclAssociationList []AWSEC2SubnetNetworkAclAssociation

func (l *AWSEC2SubnetNetworkAclAssociationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSEC2SubnetNetworkAclAssociation{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSEC2SubnetNetworkAclAssociationList{item}
		return nil
	}
	list := []AWSEC2SubnetNetworkAclAssociation{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSEC2SubnetNetworkAclAssociationList(list)
		return nil
	}
	return err
}


type AWSEC2SubnetRouteTableAssociation struct {
  RouteTableId *StringExpression `json:"RouteTableId,omitempty"`  // String
  SubnetId *StringExpression `json:"SubnetId,omitempty"`  // String
}

type AWSEC2SubnetRouteTableAssociationList []AWSEC2SubnetRouteTableAssociation

func (l *AWSEC2SubnetRouteTableAssociationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSEC2SubnetRouteTableAssociation{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSEC2SubnetRouteTableAssociationList{item}
		return nil
	}
	list := []AWSEC2SubnetRouteTableAssociation{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSEC2SubnetRouteTableAssociationList(list)
		return nil
	}
	return err
}


type AWSEC2Volume struct {
  AvailabilityZone *StringExpression `json:"AvailabilityZone,omitempty"`  // String
  Encrypted *Bool `json:"Encrypted,omitempty"`  // Boolean
  Iops *Integer `json:"Iops,omitempty"`  // Number
  KmsKeyId *StringExpression `json:"KmsKeyId,omitempty"`  // String
  Size *StringExpression `json:"Size,omitempty"`  // String
  SnapshotId *StringExpression `json:"SnapshotId,omitempty"`  // String
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
  VolumeType *StringExpression `json:"VolumeType,omitempty"`  // String
}

type AWSEC2VolumeList []AWSEC2Volume

func (l *AWSEC2VolumeList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSEC2Volume{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSEC2VolumeList{item}
		return nil
	}
	list := []AWSEC2Volume{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSEC2VolumeList(list)
		return nil
	}
	return err
}


type AWSEC2VolumeAttachment struct {
  Device *StringExpression `json:"Device,omitempty"`  // String
  InstanceId *StringExpression `json:"InstanceId,omitempty"`  // String
  VolumeId *StringExpression `json:"VolumeId,omitempty"`  // String
}

type AWSEC2VolumeAttachmentList []AWSEC2VolumeAttachment

func (l *AWSEC2VolumeAttachmentList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSEC2VolumeAttachment{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSEC2VolumeAttachmentList{item}
		return nil
	}
	list := []AWSEC2VolumeAttachment{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSEC2VolumeAttachmentList(list)
		return nil
	}
	return err
}


type AWSEC2VPC struct {
  CidrBlock *StringExpression `json:"CidrBlock,omitempty"`  // String
  EnableDnsSupport *Bool `json:"EnableDnsSupport,omitempty"`  // Boolean
  EnableDnsHostnames *Bool `json:"EnableDnsHostnames,omitempty"`  // Boolean
  InstanceTenancy *StringExpression `json:"InstanceTenancy,omitempty"`  // String
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
}

type AWSEC2VPCList []AWSEC2VPC

func (l *AWSEC2VPCList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSEC2VPC{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSEC2VPCList{item}
		return nil
	}
	list := []AWSEC2VPC{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSEC2VPCList(list)
		return nil
	}
	return err
}


type AWSEC2VPCDHCPOptionsAssociation struct {
  DhcpOptionsId *StringExpression `json:"DhcpOptionsId,omitempty"`  // String
  VpcId *StringExpression `json:"VpcId,omitempty"`  // String
}

type AWSEC2VPCDHCPOptionsAssociationList []AWSEC2VPCDHCPOptionsAssociation

func (l *AWSEC2VPCDHCPOptionsAssociationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSEC2VPCDHCPOptionsAssociation{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSEC2VPCDHCPOptionsAssociationList{item}
		return nil
	}
	list := []AWSEC2VPCDHCPOptionsAssociation{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSEC2VPCDHCPOptionsAssociationList(list)
		return nil
	}
	return err
}


type AWSEC2VPCGatewayAttachment struct {
  InternetGatewayId *StringExpression `json:"InternetGatewayId,omitempty"`  // String
  VpcId *StringExpression `json:"VpcId,omitempty"`  // String
  VpnGatewayId *StringExpression `json:"VpnGatewayId,omitempty"`  // String
}

type AWSEC2VPCGatewayAttachmentList []AWSEC2VPCGatewayAttachment

func (l *AWSEC2VPCGatewayAttachmentList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSEC2VPCGatewayAttachment{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSEC2VPCGatewayAttachmentList{item}
		return nil
	}
	list := []AWSEC2VPCGatewayAttachment{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSEC2VPCGatewayAttachmentList(list)
		return nil
	}
	return err
}


type AWSEC2VPCPeeringConnection struct {
  PeerVpcId *StringExpression `json:"PeerVpcId,omitempty"`  // String
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
  VpcId *StringExpression `json:"VpcId,omitempty"`  // String
}

type AWSEC2VPCPeeringConnectionList []AWSEC2VPCPeeringConnection

func (l *AWSEC2VPCPeeringConnectionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSEC2VPCPeeringConnection{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSEC2VPCPeeringConnectionList{item}
		return nil
	}
	list := []AWSEC2VPCPeeringConnection{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSEC2VPCPeeringConnectionList(list)
		return nil
	}
	return err
}


type AWSEC2VPNConnection struct {
  Type *StringExpression `json:"Type,omitempty"`  // String
  CustomerGatewayId *StringExpression `json:"CustomerGatewayId,omitempty"`  // String
  StaticRoutesOnly *Bool `json:"StaticRoutesOnly,omitempty"`  // Boolean
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
  VpnGatewayId *StringExpression `json:"VpnGatewayId,omitempty"`  // String
}

type AWSEC2VPNConnectionList []AWSEC2VPNConnection

func (l *AWSEC2VPNConnectionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSEC2VPNConnection{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSEC2VPNConnectionList{item}
		return nil
	}
	list := []AWSEC2VPNConnection{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSEC2VPNConnectionList(list)
		return nil
	}
	return err
}


type AWSEC2VPNConnectionRoute struct {
  DestinationCidrBlock *StringExpression `json:"DestinationCidrBlock,omitempty"`  // String
  VpnConnectionId *StringExpression `json:"VpnConnectionId,omitempty"`  // String
}

type AWSEC2VPNConnectionRouteList []AWSEC2VPNConnectionRoute

func (l *AWSEC2VPNConnectionRouteList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSEC2VPNConnectionRoute{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSEC2VPNConnectionRouteList{item}
		return nil
	}
	list := []AWSEC2VPNConnectionRoute{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSEC2VPNConnectionRouteList(list)
		return nil
	}
	return err
}


type AWSEC2VPNGateway struct {
  Type *StringExpression `json:"Type,omitempty"`  // String
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
}

type AWSEC2VPNGatewayList []AWSEC2VPNGateway

func (l *AWSEC2VPNGatewayList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSEC2VPNGateway{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSEC2VPNGatewayList{item}
		return nil
	}
	list := []AWSEC2VPNGateway{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSEC2VPNGatewayList(list)
		return nil
	}
	return err
}


type AWSEC2VPNGatewayRoutePropagation struct {
  RouteTableIds interface{} `json:"RouteTableIds,omitempty"`  // List of route table IDs
  VpnGatewayId *StringExpression `json:"VpnGatewayId,omitempty"`  // String
}

type AWSEC2VPNGatewayRoutePropagationList []AWSEC2VPNGatewayRoutePropagation

func (l *AWSEC2VPNGatewayRoutePropagationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSEC2VPNGatewayRoutePropagation{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSEC2VPNGatewayRoutePropagationList{item}
		return nil
	}
	list := []AWSEC2VPNGatewayRoutePropagation{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSEC2VPNGatewayRoutePropagationList(list)
		return nil
	}
	return err
}


type AWSECSCluster struct {
}

type AWSECSClusterList []AWSECSCluster

func (l *AWSECSClusterList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSECSCluster{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSECSClusterList{item}
		return nil
	}
	list := []AWSECSCluster{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSECSClusterList(list)
		return nil
	}
	return err
}


type AWSECSService struct {
  Cluster *StringExpression `json:"Cluster,omitempty"`  // String
  DesiredCount *StringExpression `json:"DesiredCount,omitempty"`  // String
  LoadBalancers *AmazonEC2ContainerServiceServiceLoadBalancersList `json:"LoadBalancers,omitempty"`  // List of Amazon EC2 Container Service Service LoadBalancers
  Role *StringExpression `json:"Role,omitempty"`  // String
  TaskDefinition *StringExpression `json:"TaskDefinition,omitempty"`  // String
}

type AWSECSServiceList []AWSECSService

func (l *AWSECSServiceList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSECSService{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSECSServiceList{item}
		return nil
	}
	list := []AWSECSService{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSECSServiceList(list)
		return nil
	}
	return err
}


type AWSECSTaskDefinition struct {
  ContainerDefinitions *AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsList `json:"ContainerDefinitions,omitempty"`  // List of Amazon EC2 Container Service TaskDefinition ContainerDefinitions
  Volumes *AmazonEC2ContainerServiceTaskDefinitionVolumesList `json:"Volumes,omitempty"`  // List of Amazon EC2 Container Service TaskDefinition Volumes
}

type AWSECSTaskDefinitionList []AWSECSTaskDefinition

func (l *AWSECSTaskDefinitionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSECSTaskDefinition{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSECSTaskDefinitionList{item}
		return nil
	}
	list := []AWSECSTaskDefinition{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSECSTaskDefinitionList(list)
		return nil
	}
	return err
}


type AWSElastiCacheCacheCluster struct {
  AutoMinorVersionUpgrade *Bool `json:"AutoMinorVersionUpgrade,omitempty"`  // Boolean
  AZMode *StringExpression `json:"AZMode,omitempty"`  // String
  CacheNodeType *StringExpression `json:"CacheNodeType,omitempty"`  // String
  CacheParameterGroupName *StringExpression `json:"CacheParameterGroupName,omitempty"`  // String
  CacheSecurityGroupNames *StringListExpression `json:"CacheSecurityGroupNames,omitempty"`  // List of strings
  CacheSubnetGroupName *StringExpression `json:"CacheSubnetGroupName,omitempty"`  // String
  ClusterName *StringExpression `json:"ClusterName,omitempty"`  // String
  Engine *StringExpression `json:"Engine,omitempty"`  // String
  EngineVersion *StringExpression `json:"EngineVersion,omitempty"`  // String
  NotificationTopicArn *StringExpression `json:"NotificationTopicArn,omitempty"`  // String
  NumCacheNodes *StringExpression `json:"NumCacheNodes,omitempty"`  // String
  Port *Integer `json:"Port,omitempty"`  // Integer
  PreferredAvailabilityZone *StringExpression `json:"PreferredAvailabilityZone,omitempty"`  // String
  PreferredAvailabilityZones *StringListExpression `json:"PreferredAvailabilityZones,omitempty"`  // List of strings
  PreferredMaintenanceWindow *StringExpression `json:"PreferredMaintenanceWindow,omitempty"`  // String
  SnapshotArns *StringListExpression `json:"SnapshotArns,omitempty"`  // List of strings
  SnapshotName *StringExpression `json:"SnapshotName,omitempty"`  // String
  SnapshotRetentionLimit *Integer `json:"SnapshotRetentionLimit,omitempty"`  // Integer
  SnapshotWindow *StringExpression `json:"SnapshotWindow,omitempty"`  // String
  XVpcSecurityGroupIdsX *StringListExpression `json:" VpcSecurityGroupIds ,omitempty"`  // List of strings
}

type AWSElastiCacheCacheClusterList []AWSElastiCacheCacheCluster

func (l *AWSElastiCacheCacheClusterList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSElastiCacheCacheCluster{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSElastiCacheCacheClusterList{item}
		return nil
	}
	list := []AWSElastiCacheCacheCluster{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSElastiCacheCacheClusterList(list)
		return nil
	}
	return err
}


type AWSElastiCacheParameterGroup struct {
  CacheParameterGroupFamily *StringExpression `json:"CacheParameterGroupFamily,omitempty"`  // String
  Description *StringExpression `json:"Description,omitempty"`  // String
  Properties interface{} `json:"Properties,omitempty"`  // JSON object
}

type AWSElastiCacheParameterGroupList []AWSElastiCacheParameterGroup

func (l *AWSElastiCacheParameterGroupList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSElastiCacheParameterGroup{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSElastiCacheParameterGroupList{item}
		return nil
	}
	list := []AWSElastiCacheParameterGroup{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSElastiCacheParameterGroupList(list)
		return nil
	}
	return err
}


type AWSElastiCacheReplicationGroup struct {
  AutomaticFailoverEnabled *Bool `json:"AutomaticFailoverEnabled,omitempty"`  // Boolean
  AutoMinorVersionUpgrade *Bool `json:"AutoMinorVersionUpgrade,omitempty"`  // Boolean
  CacheNodeType *StringExpression `json:"CacheNodeType,omitempty"`  // String
  CacheParameterGroupName *StringExpression `json:"CacheParameterGroupName,omitempty"`  // String
  CacheSecurityGroupNames *StringListExpression `json:"CacheSecurityGroupNames,omitempty"`  // List of strings
  CacheSubnetGroupName *StringExpression `json:"CacheSubnetGroupName,omitempty"`  // String
  Engine *StringExpression `json:"Engine,omitempty"`  // String
  EngineVersion *StringExpression `json:"EngineVersion,omitempty"`  // String
  NotificationTopicArn *StringExpression `json:"NotificationTopicArn,omitempty"`  // String
  NumCacheClusters *Integer `json:"NumCacheClusters,omitempty"`  // Integer
  Port *Integer `json:"Port,omitempty"`  // Integer
  PreferredCacheClusterAZs *StringListExpression `json:"PreferredCacheClusterAZs,omitempty"`  // List of strings
  PreferredMaintenanceWindow *StringExpression `json:"PreferredMaintenanceWindow,omitempty"`  // String
  ReplicationGroupDescription *StringExpression `json:"ReplicationGroupDescription,omitempty"`  // String
  SecurityGroupIds *StringListExpression `json:"SecurityGroupIds,omitempty"`  // List of strings
  SnapshotArns *StringListExpression `json:"SnapshotArns,omitempty"`  // List of strings
  SnapshotRetentionLimit *Integer `json:"SnapshotRetentionLimit,omitempty"`  // Integer
  SnapshotWindow *StringExpression `json:"SnapshotWindow,omitempty"`  // String
}

type AWSElastiCacheReplicationGroupList []AWSElastiCacheReplicationGroup

func (l *AWSElastiCacheReplicationGroupList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSElastiCacheReplicationGroup{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSElastiCacheReplicationGroupList{item}
		return nil
	}
	list := []AWSElastiCacheReplicationGroup{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSElastiCacheReplicationGroupList(list)
		return nil
	}
	return err
}


type AWSElastiCacheSecurityGroup struct {
  Description *StringExpression `json:"Description,omitempty"`  // String
}

type AWSElastiCacheSecurityGroupList []AWSElastiCacheSecurityGroup

func (l *AWSElastiCacheSecurityGroupList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSElastiCacheSecurityGroup{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSElastiCacheSecurityGroupList{item}
		return nil
	}
	list := []AWSElastiCacheSecurityGroup{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSElastiCacheSecurityGroupList(list)
		return nil
	}
	return err
}


type AWSElastiCacheSecurityGroupIngress struct {
  CacheSecurityGroupName *StringExpression `json:"CacheSecurityGroupName,omitempty"`  // String
  EC2SecurityGroupName *StringExpression `json:"EC2SecurityGroupName,omitempty"`  // String
  EC2SecurityGroupOwnerId *StringExpression `json:"EC2SecurityGroupOwnerId,omitempty"`  // String
}

type AWSElastiCacheSecurityGroupIngressList []AWSElastiCacheSecurityGroupIngress

func (l *AWSElastiCacheSecurityGroupIngressList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSElastiCacheSecurityGroupIngress{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSElastiCacheSecurityGroupIngressList{item}
		return nil
	}
	list := []AWSElastiCacheSecurityGroupIngress{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSElastiCacheSecurityGroupIngressList(list)
		return nil
	}
	return err
}


type AWSElastiCacheSubnetGroup struct {
}

type AWSElastiCacheSubnetGroupList []AWSElastiCacheSubnetGroup

func (l *AWSElastiCacheSubnetGroupList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSElastiCacheSubnetGroup{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSElastiCacheSubnetGroupList{item}
		return nil
	}
	list := []AWSElastiCacheSubnetGroup{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSElastiCacheSubnetGroupList(list)
		return nil
	}
	return err
}


type AWSElasticBeanstalkApplication struct {
  ApplicationName *StringExpression `json:"ApplicationName,omitempty"`  // String
  Description *StringExpression `json:"Description,omitempty"`  // String
}

type AWSElasticBeanstalkApplicationList []AWSElasticBeanstalkApplication

func (l *AWSElasticBeanstalkApplicationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSElasticBeanstalkApplication{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSElasticBeanstalkApplicationList{item}
		return nil
	}
	list := []AWSElasticBeanstalkApplication{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSElasticBeanstalkApplicationList(list)
		return nil
	}
	return err
}


type AWSElasticBeanstalkApplicationVersion struct {
}

type AWSElasticBeanstalkApplicationVersionList []AWSElasticBeanstalkApplicationVersion

func (l *AWSElasticBeanstalkApplicationVersionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSElasticBeanstalkApplicationVersion{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSElasticBeanstalkApplicationVersionList{item}
		return nil
	}
	list := []AWSElasticBeanstalkApplicationVersion{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSElasticBeanstalkApplicationVersionList(list)
		return nil
	}
	return err
}


type AWSElasticBeanstalkConfigurationTemplate struct {
}

type AWSElasticBeanstalkConfigurationTemplateList []AWSElasticBeanstalkConfigurationTemplate

func (l *AWSElasticBeanstalkConfigurationTemplateList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSElasticBeanstalkConfigurationTemplate{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSElasticBeanstalkConfigurationTemplateList{item}
		return nil
	}
	list := []AWSElasticBeanstalkConfigurationTemplate{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSElasticBeanstalkConfigurationTemplateList(list)
		return nil
	}
	return err
}


type AWSElasticBeanstalkEnvironment struct {
  ApplicationName *StringExpression `json:"ApplicationName,omitempty"`  // String
  CNAMEPrefix *StringExpression `json:"CNAMEPrefix,omitempty"`  // String
  Description *StringExpression `json:"Description,omitempty"`  // String
  EnvironmentName *StringExpression `json:"EnvironmentName,omitempty"`  // String
  OptionSettings *ElasticBeanstalkOptionSettingsList `json:"OptionSettings,omitempty"`  // A list of OptionSettings
  SolutionStackName *StringExpression `json:"SolutionStackName,omitempty"`  // String
  TemplateName *StringExpression `json:"TemplateName,omitempty"`  // String
  Tier *ElasticBeanstalkEnvironmentTier `json:"Tier,omitempty"`  // Elastic Beanstalk Environment Tier Property Type
  VersionLabel *StringExpression `json:"VersionLabel,omitempty"`  // String
}

type AWSElasticBeanstalkEnvironmentList []AWSElasticBeanstalkEnvironment

func (l *AWSElasticBeanstalkEnvironmentList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSElasticBeanstalkEnvironment{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSElasticBeanstalkEnvironmentList{item}
		return nil
	}
	list := []AWSElasticBeanstalkEnvironment{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSElasticBeanstalkEnvironmentList(list)
		return nil
	}
	return err
}


type AWSElasticLoadBalancingLoadBalancer struct {
  AccessLoggingPolicy *ElasticLoadBalancingAccessLoggingPolicy `json:"AccessLoggingPolicy,omitempty"`  // Elastic Load Balancing AccessLoggingPolicy
  AppCookieStickinessPolicy *ElasticLoadBalancingAppCookieStickinessPolicyList `json:"AppCookieStickinessPolicy,omitempty"`  // A list of AppCookieStickinessPolicy objects
  AvailabilityZones *StringListExpression `json:"AvailabilityZones,omitempty"`  // List of strings
  ConnectionDrainingPolicy *ElasticLoadBalancingConnectionDrainingPolicy `json:"ConnectionDrainingPolicy,omitempty"`  // Elastic Load Balancing ConnectionDrainingPolicy
  ConnectionSettings *ElasticLoadBalancingConnectionSettings `json:"ConnectionSettings,omitempty"`  // Elastic Load Balancing ConnectionSettings
  CrossZone *Bool `json:"CrossZone,omitempty"`  // Boolean
  HealthCheck *ElasticLoadBalancingHealthCheck `json:"HealthCheck,omitempty"`  // ElasticLoadBalancing HealthCheck Type
  Instances *StringListExpression `json:"Instances,omitempty"`  // List of strings
  LBCookieStickinessPolicy *ElasticLoadBalancingLBCookieStickinessPolicyList `json:"LBCookieStickinessPolicy,omitempty"`  // A list of LBCookieStickinessPolicy objects
  LoadBalancerName *StringExpression `json:"LoadBalancerName,omitempty"`  // String
  Listeners *ElasticLoadBalancingListenerList `json:"Listeners,omitempty"`  // A list of ElasticLoadBalancing Listener Property Type objects
  Policies *ElasticLoadBalancingPolicyList `json:"Policies,omitempty"`  // A list of ElasticLoadBalancing policy objects
  Scheme *StringExpression `json:"Scheme,omitempty"`  // String
  SecurityGroups interface{} `json:"SecurityGroups,omitempty"`  // A list of security groups assigned to your load balancer within your virtual private cloud (VPC)
  Subnets *StringListExpression `json:"Subnets,omitempty"`  // List of strings
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
}

type AWSElasticLoadBalancingLoadBalancerList []AWSElasticLoadBalancingLoadBalancer

func (l *AWSElasticLoadBalancingLoadBalancerList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSElasticLoadBalancingLoadBalancer{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSElasticLoadBalancingLoadBalancerList{item}
		return nil
	}
	list := []AWSElasticLoadBalancingLoadBalancer{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSElasticLoadBalancingLoadBalancerList(list)
		return nil
	}
	return err
}


type AWSIAMAccessKey struct {
  Serial *Integer `json:"Serial,omitempty"`  // Integer
  Status *StringExpression `json:"Status,omitempty"`  // String
  UserName *StringExpression `json:"UserName,omitempty"`  // String
}

type AWSIAMAccessKeyList []AWSIAMAccessKey

func (l *AWSIAMAccessKeyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSIAMAccessKey{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSIAMAccessKeyList{item}
		return nil
	}
	list := []AWSIAMAccessKey{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSIAMAccessKeyList(list)
		return nil
	}
	return err
}


type AWSIAMGroup struct {
  ManagedPolicyArns *StringListExpression `json:"ManagedPolicyArns,omitempty"`  // List of strings
  Path *StringExpression `json:"Path,omitempty"`  // String
  Policies *IAMPoliciesList `json:"Policies,omitempty"`  // List of IAM Policies
}

type AWSIAMGroupList []AWSIAMGroup

func (l *AWSIAMGroupList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSIAMGroup{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSIAMGroupList{item}
		return nil
	}
	list := []AWSIAMGroup{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSIAMGroupList(list)
		return nil
	}
	return err
}


type AWSIAMInstanceProfile struct {
  Path *StringExpression `json:"Path,omitempty"`  // String
  Roles interface{} `json:"Roles,omitempty"`  // List of references to AWS::IAM::Roles. Currently, a maximum of one role can be assigned to an instance profile
}

type AWSIAMInstanceProfileList []AWSIAMInstanceProfile

func (l *AWSIAMInstanceProfileList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSIAMInstanceProfile{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSIAMInstanceProfileList{item}
		return nil
	}
	list := []AWSIAMInstanceProfile{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSIAMInstanceProfileList(list)
		return nil
	}
	return err
}


type AWSIAMManagedPolicy struct {
  Description *StringExpression `json:"Description,omitempty"`  // String
  Groups *StringListExpression `json:"Groups,omitempty"`  // List of strings
  Path *StringExpression `json:"Path,omitempty"`  // String
  PolicyDocument interface{} `json:"PolicyDocument,omitempty"`  // JSON object
  Roles *StringListExpression `json:"Roles,omitempty"`  // List of strings
  Users *StringListExpression `json:"Users,omitempty"`  // List of strings
}

type AWSIAMManagedPolicyList []AWSIAMManagedPolicy

func (l *AWSIAMManagedPolicyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSIAMManagedPolicy{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSIAMManagedPolicyList{item}
		return nil
	}
	list := []AWSIAMManagedPolicy{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSIAMManagedPolicyList(list)
		return nil
	}
	return err
}


type AWSIAMPolicy struct {
  Groups *StringListExpression `json:"Groups,omitempty"`  // List of strings
  PolicyDocument interface{} `json:"PolicyDocument,omitempty"`  // JSON object
  PolicyName *StringExpression `json:"PolicyName,omitempty"`  // String
  Roles *StringListExpression `json:"Roles,omitempty"`  // List of strings
  Users *StringListExpression `json:"Users,omitempty"`  // List of strings
}

type AWSIAMPolicyList []AWSIAMPolicy

func (l *AWSIAMPolicyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSIAMPolicy{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSIAMPolicyList{item}
		return nil
	}
	list := []AWSIAMPolicy{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSIAMPolicyList(list)
		return nil
	}
	return err
}


type AWSIAMRole struct {
  AssumeRolePolicyDocument interface{} `json:"AssumeRolePolicyDocument,omitempty"`  // A JSON policy document
  ManagedPolicyArns *StringListExpression `json:"ManagedPolicyArns,omitempty"`  // List of strings
  Path *StringExpression `json:"Path,omitempty"`  // String
  Policies *IAMPoliciesList `json:"Policies,omitempty"`  // List of IAM Policies
}

type AWSIAMRoleList []AWSIAMRole

func (l *AWSIAMRoleList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSIAMRole{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSIAMRoleList{item}
		return nil
	}
	list := []AWSIAMRole{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSIAMRoleList(list)
		return nil
	}
	return err
}


type AWSIAMUser struct {
  Groups *StringListExpression `json:"Groups,omitempty"`  // List of strings
  LoginProfile *IAMUserLoginProfile `json:"LoginProfile,omitempty"`  // IAM User LoginProfile
  ManagedPolicyArns *StringListExpression `json:"ManagedPolicyArns,omitempty"`  // List of strings
  Path *StringExpression `json:"Path,omitempty"`  // String
  Policies *IAMPoliciesList `json:"Policies,omitempty"`  // List of IAM Policies
}

type AWSIAMUserList []AWSIAMUser

func (l *AWSIAMUserList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSIAMUser{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSIAMUserList{item}
		return nil
	}
	list := []AWSIAMUser{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSIAMUserList(list)
		return nil
	}
	return err
}


type AWSIAMUserToGroupAddition struct {
  GroupName *StringExpression `json:"GroupName,omitempty"`  // String
  Users interface{} `json:"Users,omitempty"`  // List of users
}

type AWSIAMUserToGroupAdditionList []AWSIAMUserToGroupAddition

func (l *AWSIAMUserToGroupAdditionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSIAMUserToGroupAddition{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSIAMUserToGroupAdditionList{item}
		return nil
	}
	list := []AWSIAMUserToGroupAddition{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSIAMUserToGroupAdditionList(list)
		return nil
	}
	return err
}


type AWSKinesisStream struct {
  ShardCount *Integer `json:"ShardCount,omitempty"`  // Integer
}

type AWSKinesisStreamList []AWSKinesisStream

func (l *AWSKinesisStreamList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSKinesisStream{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSKinesisStreamList{item}
		return nil
	}
	list := []AWSKinesisStream{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSKinesisStreamList(list)
		return nil
	}
	return err
}


type AWSLambdaFunction struct {
  Code *AWSLambdaFunctionCode `json:"Code,omitempty"`  // AWS Lambda Function Code
  Description *StringExpression `json:"Description,omitempty"`  // String
  Handler *StringExpression `json:"Handler,omitempty"`  // String
  MemorySize *Integer `json:"MemorySize,omitempty"`  // Integer
  Role *StringExpression `json:"Role,omitempty"`  // String
  Runtime *StringExpression `json:"Runtime,omitempty"`  // String
  Timeout *Integer `json:"Timeout,omitempty"`  // Integer
}

type AWSLambdaFunctionList []AWSLambdaFunction

func (l *AWSLambdaFunctionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSLambdaFunction{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSLambdaFunctionList{item}
		return nil
	}
	list := []AWSLambdaFunction{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSLambdaFunctionList(list)
		return nil
	}
	return err
}


type AWSLogsLogGroup struct {
  RetentionInDays *Integer `json:"RetentionInDays,omitempty"`  // Integer
}

type AWSLogsLogGroupList []AWSLogsLogGroup

func (l *AWSLogsLogGroupList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSLogsLogGroup{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSLogsLogGroupList{item}
		return nil
	}
	list := []AWSLogsLogGroup{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSLogsLogGroupList(list)
		return nil
	}
	return err
}


type AWSLogsMetricFilter struct {
  FilterPattern *StringListExpression `json:"FilterPattern,omitempty"`  // List of strings
  LogGroupName *StringExpression `json:"LogGroupName,omitempty"`  // String
  MetricTransformations *CloudWatchLogsMetricFilterMetricTransformationPropertyList `json:"MetricTransformations,omitempty"`  // A list of CloudWatch Logs MetricFilter MetricTransformation Property
}

type AWSLogsMetricFilterList []AWSLogsMetricFilter

func (l *AWSLogsMetricFilterList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSLogsMetricFilter{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSLogsMetricFilterList{item}
		return nil
	}
	list := []AWSLogsMetricFilter{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSLogsMetricFilterList(list)
		return nil
	}
	return err
}


type AWSOpsWorksApp struct {
  AppSource *AWSOpsWorksSource `json:"AppSource,omitempty"`  // AWS OpsWorks Source Type
  Attributes interface{} `json:"Attributes,omitempty"`  // A list of key-value pairs
  Description *StringExpression `json:"Description,omitempty"`  // String
  Domains *StringListExpression `json:"Domains,omitempty"`  // List of strings
  EnableSsl *Bool `json:"EnableSsl,omitempty"`  // Boolean
  Name *StringExpression `json:"Name,omitempty"`  // String
  Shortname *StringExpression `json:"Shortname,omitempty"`  // String
  SslConfiguration *AWSOpsWorksSslConfiguration `json:"SslConfiguration,omitempty"`  // AWS OpsWorks SslConfiguration Type
  StackId *StringExpression `json:"StackId,omitempty"`  // String
  Type *StringExpression `json:"Type,omitempty"`  // String
}

type AWSOpsWorksAppList []AWSOpsWorksApp

func (l *AWSOpsWorksAppList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSOpsWorksApp{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSOpsWorksAppList{item}
		return nil
	}
	list := []AWSOpsWorksApp{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSOpsWorksAppList(list)
		return nil
	}
	return err
}


type AWSOpsWorksElasticLoadBalancerAttachment struct {
  ElasticLoadBalancerName *StringExpression `json:"ElasticLoadBalancerName,omitempty"`  // String
  LayerId *StringExpression `json:"LayerId,omitempty"`  // String
}

type AWSOpsWorksElasticLoadBalancerAttachmentList []AWSOpsWorksElasticLoadBalancerAttachment

func (l *AWSOpsWorksElasticLoadBalancerAttachmentList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSOpsWorksElasticLoadBalancerAttachment{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSOpsWorksElasticLoadBalancerAttachmentList{item}
		return nil
	}
	list := []AWSOpsWorksElasticLoadBalancerAttachment{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSOpsWorksElasticLoadBalancerAttachmentList(list)
		return nil
	}
	return err
}


type AWSOpsWorksInstance struct {
  AmiId *StringExpression `json:"AmiId,omitempty"`  // String
  Architecture *StringExpression `json:"Architecture,omitempty"`  // String
  AutoScalingType *StringExpression `json:"AutoScalingType,omitempty"`  // String
  AvailabilityZone *StringExpression `json:"AvailabilityZone,omitempty"`  // String
  InstallUpdatesOnBoot *Bool `json:"InstallUpdatesOnBoot,omitempty"`  // Boolean
  InstanceType *StringExpression `json:"InstanceType,omitempty"`  // String
  LayerIds *StringListExpression `json:"LayerIds,omitempty"`  // List of strings
  Os *StringExpression `json:"Os,omitempty"`  // String
  RootDeviceType *StringExpression `json:"RootDeviceType,omitempty"`  // String
  SshKeyName *StringExpression `json:"SshKeyName,omitempty"`  // String
  StackId *StringExpression `json:"StackId,omitempty"`  // String
  SubnetId *StringExpression `json:"SubnetId,omitempty"`  // String
  TimeBasedAutoScaling *AWSOpsWorksTimeBasedAutoScaling `json:"TimeBasedAutoScaling,omitempty"`  // AWS OpsWorks TimeBasedAutoScaling Type
}

type AWSOpsWorksInstanceList []AWSOpsWorksInstance

func (l *AWSOpsWorksInstanceList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSOpsWorksInstance{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSOpsWorksInstanceList{item}
		return nil
	}
	list := []AWSOpsWorksInstance{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSOpsWorksInstanceList(list)
		return nil
	}
	return err
}


type AWSOpsWorksLayer struct {
  Attributes interface{} `json:"Attributes,omitempty"`  // A list of key-value pairs
  AutoAssignElasticIps *Bool `json:"AutoAssignElasticIps,omitempty"`  // Boolean
  AutoAssignPublicIps *Bool `json:"AutoAssignPublicIps,omitempty"`  // Boolean
  CustomInstanceProfileArn *StringExpression `json:"CustomInstanceProfileArn,omitempty"`  // String
  CustomRecipes *AWSOpsWorksRecipes `json:"CustomRecipes,omitempty"`  // AWS OpsWorks Recipes Type
  CustomSecurityGroupIds *StringListExpression `json:"CustomSecurityGroupIds,omitempty"`  // List of strings
  EnableAutoHealing *Bool `json:"EnableAutoHealing,omitempty"`  // Boolean
  InstallUpdatesOnBoot *Bool `json:"InstallUpdatesOnBoot,omitempty"`  // Boolean
  LifecycleEventConfiguration *AWSOpsWorksLayerLifeCycleConfiguration `json:"LifecycleEventConfiguration,omitempty"`  // AWS OpsWorks Layer LifeCycleConfiguration
  LoadBasedAutoScaling *AWSOpsWorksLoadBasedAutoScaling `json:"LoadBasedAutoScaling,omitempty"`  // AWS OpsWorks LoadBasedAutoScaling Type
  Name *StringExpression `json:"Name,omitempty"`  // String
  Packages *StringListExpression `json:"Packages,omitempty"`  // List of strings
  Shortname *StringExpression `json:"Shortname,omitempty"`  // String
  StackId *StringExpression `json:"StackId,omitempty"`  // String
  Type *StringExpression `json:"Type,omitempty"`  // String
  VolumeConfigurations *AWSOpsWorksVolumeConfigurationList `json:"VolumeConfigurations,omitempty"`  // A list of AWS OpsWorks VolumeConfiguration Type
}

type AWSOpsWorksLayerList []AWSOpsWorksLayer

func (l *AWSOpsWorksLayerList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSOpsWorksLayer{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSOpsWorksLayerList{item}
		return nil
	}
	list := []AWSOpsWorksLayer{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSOpsWorksLayerList(list)
		return nil
	}
	return err
}


type AWSOpsWorksStack struct {
  Attributes interface{} `json:"Attributes,omitempty"`  // A list of key-value pairs
  ChefConfiguration *AWSOpsWorksChefConfiguration `json:"ChefConfiguration,omitempty"`  // AWS OpsWorks ChefConfiguration Type
  ConfigurationManager *AWSOpsWorksStackConfigurationManager `json:"ConfigurationManager,omitempty"`  // AWS OpsWorks StackConfigurationManager Type
  CustomCookbooksSource *AWSOpsWorksSource `json:"CustomCookbooksSource,omitempty"`  // AWS OpsWorks Source Type
  CustomJson interface{} `json:"CustomJson,omitempty"`  // JSON object
  DefaultAvailabilityZone *StringExpression `json:"DefaultAvailabilityZone,omitempty"`  // String
  DefaultInstanceProfileArn *StringExpression `json:"DefaultInstanceProfileArn,omitempty"`  // String
  DefaultOs *StringExpression `json:"DefaultOs,omitempty"`  // String
  DefaultRootDeviceType *StringExpression `json:"DefaultRootDeviceType,omitempty"`  // String
  DefaultSshKeyName *StringExpression `json:"DefaultSshKeyName,omitempty"`  // String
  DefaultSubnetId *StringExpression `json:"DefaultSubnetId,omitempty"`  // String
  HostnameTheme *StringExpression `json:"HostnameTheme,omitempty"`  // String
  Name *StringExpression `json:"Name,omitempty"`  // String
  ServiceRoleArn *StringExpression `json:"ServiceRoleArn,omitempty"`  // String
  UseCustomCookbooks *Bool `json:"UseCustomCookbooks,omitempty"`  // Boolean
  UseOpsworksSecurityGroups *Bool `json:"UseOpsworksSecurityGroups,omitempty"`  // Boolean
  VpcId *StringExpression `json:"VpcId,omitempty"`  // String
}

type AWSOpsWorksStackList []AWSOpsWorksStack

func (l *AWSOpsWorksStackList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSOpsWorksStack{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSOpsWorksStackList{item}
		return nil
	}
	list := []AWSOpsWorksStack{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSOpsWorksStackList(list)
		return nil
	}
	return err
}


type AWSRedshiftCluster struct {
  AllowVersionUpgrade *Bool `json:"AllowVersionUpgrade,omitempty"`  // Boolean
  AutomatedSnapshotRetentionPeriod *Integer `json:"AutomatedSnapshotRetentionPeriod,omitempty"`  // Integer
  AvailabilityZone *StringExpression `json:"AvailabilityZone,omitempty"`  // String
  ClusterParameterGroupName *StringExpression `json:"ClusterParameterGroupName,omitempty"`  // String
  ClusterSecurityGroups *StringListExpression `json:"ClusterSecurityGroups,omitempty"`  // List of strings
  ClusterSubnetGroupName *StringExpression `json:"ClusterSubnetGroupName,omitempty"`  // String
  ClusterType *StringExpression `json:"ClusterType,omitempty"`  // String
  ClusterVersion *StringExpression `json:"ClusterVersion,omitempty"`  // String
  DBName *StringExpression `json:"DBName,omitempty"`  // String
  ElasticIp *StringExpression `json:"ElasticIp,omitempty"`  // String
  Encrypted *Bool `json:"Encrypted,omitempty"`  // Boolean
  HsmClientCertificateIdentifier *StringExpression `json:"HsmClientCertificateIdentifier,omitempty"`  // String
  HsmConfigurationIdentifier *StringExpression `json:"HsmConfigurationIdentifier,omitempty"`  // String
  MasterUsername *StringExpression `json:"MasterUsername,omitempty"`  // String
  MasterUserPassword *StringExpression `json:"MasterUserPassword,omitempty"`  // String
  NodeType *StringExpression `json:"NodeType,omitempty"`  // String
  NumberOfNodes *Integer `json:"NumberOfNodes,omitempty"`  // Integer
  OwnerAccount *StringExpression `json:"OwnerAccount,omitempty"`  // String
  Port *Integer `json:"Port,omitempty"`  // Integer
  PreferredMaintenanceWindow *StringExpression `json:"PreferredMaintenanceWindow,omitempty"`  // String
  PubliclyAccessible *Bool `json:"PubliclyAccessible,omitempty"`  // Boolean
  SnapshotClusterIdentifier interface{} `json:"SnapshotClusterIdentifier,omitempty"`  // 
  SnapshotIdentifier *StringExpression `json:"SnapshotIdentifier,omitempty"`  // String
  VpcSecurityGroupIds *StringListExpression `json:"VpcSecurityGroupIds,omitempty"`  // List of strings
}

type AWSRedshiftClusterList []AWSRedshiftCluster

func (l *AWSRedshiftClusterList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSRedshiftCluster{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSRedshiftClusterList{item}
		return nil
	}
	list := []AWSRedshiftCluster{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSRedshiftClusterList(list)
		return nil
	}
	return err
}


type AWSRedshiftClusterParameterGroup struct {
  Description *StringExpression `json:"Description,omitempty"`  // String
  ParameterGroupFamily *StringExpression `json:"ParameterGroupFamily,omitempty"`  // String
  Parameters *AmazonRedshiftParameterList `json:"Parameters,omitempty"`  // Amazon Redshift Parameter Type
}

type AWSRedshiftClusterParameterGroupList []AWSRedshiftClusterParameterGroup

func (l *AWSRedshiftClusterParameterGroupList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSRedshiftClusterParameterGroup{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSRedshiftClusterParameterGroupList{item}
		return nil
	}
	list := []AWSRedshiftClusterParameterGroup{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSRedshiftClusterParameterGroupList(list)
		return nil
	}
	return err
}


type AWSRedshiftClusterSecurityGroup struct {
  Description *StringExpression `json:"Description,omitempty"`  // String
}

type AWSRedshiftClusterSecurityGroupList []AWSRedshiftClusterSecurityGroup

func (l *AWSRedshiftClusterSecurityGroupList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSRedshiftClusterSecurityGroup{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSRedshiftClusterSecurityGroupList{item}
		return nil
	}
	list := []AWSRedshiftClusterSecurityGroup{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSRedshiftClusterSecurityGroupList(list)
		return nil
	}
	return err
}


type AWSRedshiftClusterSecurityGroupIngress struct {
  ClusterSecurityGroupName *StringExpression `json:"ClusterSecurityGroupName,omitempty"`  // String
  CIDRIP *StringExpression `json:"CIDRIP,omitempty"`  // String
  EC2SecurityGroupName *StringExpression `json:"EC2SecurityGroupName,omitempty"`  // String
  EC2SecurityGroupOwnerId *StringExpression `json:"EC2SecurityGroupOwnerId,omitempty"`  // String
}

type AWSRedshiftClusterSecurityGroupIngressList []AWSRedshiftClusterSecurityGroupIngress

func (l *AWSRedshiftClusterSecurityGroupIngressList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSRedshiftClusterSecurityGroupIngress{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSRedshiftClusterSecurityGroupIngressList{item}
		return nil
	}
	list := []AWSRedshiftClusterSecurityGroupIngress{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSRedshiftClusterSecurityGroupIngressList(list)
		return nil
	}
	return err
}


type AWSRedshiftClusterSubnetGroup struct {
  Description *StringExpression `json:"Description,omitempty"`  // String
  SubnetIds *StringListExpression `json:"SubnetIds,omitempty"`  // List of strings
}

type AWSRedshiftClusterSubnetGroupList []AWSRedshiftClusterSubnetGroup

func (l *AWSRedshiftClusterSubnetGroupList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSRedshiftClusterSubnetGroup{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSRedshiftClusterSubnetGroupList{item}
		return nil
	}
	list := []AWSRedshiftClusterSubnetGroup{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSRedshiftClusterSubnetGroupList(list)
		return nil
	}
	return err
}


type AWSRDSDBInstance struct {
  AllocatedStorage *StringExpression `json:"AllocatedStorage,omitempty"`  // String
  AllowMajorVersionUpgrade *Bool `json:"AllowMajorVersionUpgrade,omitempty"`  // Boolean
  AutoMinorVersionUpgrade *Bool `json:"AutoMinorVersionUpgrade,omitempty"`  // Boolean
  AvailabilityZone *StringExpression `json:"AvailabilityZone,omitempty"`  // String
  BackupRetentionPeriod *StringExpression `json:"BackupRetentionPeriod,omitempty"`  // String
  CharacterSetName *StringExpression `json:"CharacterSetName,omitempty"`  // String
  DBInstanceClass *StringExpression `json:"DBInstanceClass,omitempty"`  // String
  DBInstanceIdentifier *StringExpression `json:"DBInstanceIdentifier,omitempty"`  // String
  DBName *StringExpression `json:"DBName,omitempty"`  // String
  DBParameterGroupName *StringExpression `json:"DBParameterGroupName,omitempty"`  // String
  DBSecurityGroups *StringListExpression `json:"DBSecurityGroups,omitempty"`  // List of strings
  DBSnapshotIdentifier *StringExpression `json:"DBSnapshotIdentifier,omitempty"`  // String
  DBSubnetGroupName *StringExpression `json:"DBSubnetGroupName,omitempty"`  // String
  Engine *StringExpression `json:"Engine,omitempty"`  // String
  EngineVersion *StringExpression `json:"EngineVersion,omitempty"`  // String
  Iops *Integer `json:"Iops,omitempty"`  // Number
  KmsKeyId *StringExpression `json:"KmsKeyId,omitempty"`  // String
  LicenseModel *StringExpression `json:"LicenseModel,omitempty"`  // String
  MasterUsername *StringExpression `json:"MasterUsername,omitempty"`  // String
  MasterUserPassword *StringExpression `json:"MasterUserPassword,omitempty"`  // String
  MultiAZ *Bool `json:"MultiAZ,omitempty"`  // Boolean
  OptionGroupName *StringExpression `json:"OptionGroupName,omitempty"`  // String
  Port *StringExpression `json:"Port,omitempty"`  // String
  PreferredBackupWindow *StringExpression `json:"PreferredBackupWindow,omitempty"`  // String
  PreferredMaintenanceWindow *StringExpression `json:"PreferredMaintenanceWindow,omitempty"`  // String
  PubliclyAccessible *Bool `json:"PubliclyAccessible,omitempty"`  // Boolean
  SourceDBInstanceIdentifier *StringExpression `json:"SourceDBInstanceIdentifier,omitempty"`  // String
  StorageEncrypted *Bool `json:"StorageEncrypted,omitempty"`  // Boolean
  StorageType *StringExpression `json:"StorageType,omitempty"`  // String
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
  VPCSecurityGroups *StringListExpression `json:"VPCSecurityGroups,omitempty"`  // List of strings
}

type AWSRDSDBInstanceList []AWSRDSDBInstance

func (l *AWSRDSDBInstanceList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSRDSDBInstance{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSRDSDBInstanceList{item}
		return nil
	}
	list := []AWSRDSDBInstance{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSRDSDBInstanceList(list)
		return nil
	}
	return err
}


type AWSRDSDBParameterGroup struct {
  Description interface{} `json:"Description,omitempty"`  // 
  Family interface{} `json:"Family,omitempty"`  // 
  Parameters interface{} `json:"Parameters,omitempty"`  // 
  Tags *ResourceTagList `json:"Tags,omitempty"`  // A list of resource tags
}

type AWSRDSDBParameterGroupList []AWSRDSDBParameterGroup

func (l *AWSRDSDBParameterGroupList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSRDSDBParameterGroup{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSRDSDBParameterGroupList{item}
		return nil
	}
	list := []AWSRDSDBParameterGroup{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSRDSDBParameterGroupList(list)
		return nil
	}
	return err
}


type AWSRDSDBSubnetGroup struct {
  DBSubnetGroupDescription *StringExpression `json:"DBSubnetGroupDescription,omitempty"`  // String
  SubnetIds *StringListExpression `json:"SubnetIds,omitempty"`  // List of strings
  Tags *ResourceTagList `json:"Tags,omitempty"`  // A list of resource tags
}

type AWSRDSDBSubnetGroupList []AWSRDSDBSubnetGroup

func (l *AWSRDSDBSubnetGroupList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSRDSDBSubnetGroup{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSRDSDBSubnetGroupList{item}
		return nil
	}
	list := []AWSRDSDBSubnetGroup{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSRDSDBSubnetGroupList(list)
		return nil
	}
	return err
}


type AWSRDSDBSecurityGroup struct {
  EC2VpcId *StringExpression `json:"EC2VpcId,omitempty"`  // String
  DBSecurityGroupIngress *AmazonRDSSecurityGroupRuleList `json:"DBSecurityGroupIngress,omitempty"`  // List of RDS Security Group Rules
  GroupDescription *StringExpression `json:"GroupDescription,omitempty"`  // String
  Tags *ResourceTagList `json:"Tags,omitempty"`  // A list of resource tags
}

type AWSRDSDBSecurityGroupList []AWSRDSDBSecurityGroup

func (l *AWSRDSDBSecurityGroupList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSRDSDBSecurityGroup{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSRDSDBSecurityGroupList{item}
		return nil
	}
	list := []AWSRDSDBSecurityGroup{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSRDSDBSecurityGroupList(list)
		return nil
	}
	return err
}


type AWSRDSDBSecurityGroupIngress struct {
  CIDRIP *StringExpression `json:"CIDRIP,omitempty"`  // String
  DBSecurityGroupName *StringExpression `json:"DBSecurityGroupName,omitempty"`  // String
  EC2SecurityGroupId *StringExpression `json:"EC2SecurityGroupId,omitempty"`  // String
  EC2SecurityGroupName *StringExpression `json:"EC2SecurityGroupName,omitempty"`  // String
  EC2SecurityGroupOwnerId *StringExpression `json:"EC2SecurityGroupOwnerId,omitempty"`  // String
}

type AWSRDSDBSecurityGroupIngressList []AWSRDSDBSecurityGroupIngress

func (l *AWSRDSDBSecurityGroupIngressList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSRDSDBSecurityGroupIngress{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSRDSDBSecurityGroupIngressList{item}
		return nil
	}
	list := []AWSRDSDBSecurityGroupIngress{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSRDSDBSecurityGroupIngressList(list)
		return nil
	}
	return err
}


type AWSRDSEventSubscription struct {
  Enabled *Bool `json:"Enabled,omitempty"`  // Boolean
  EventCategories *StringListExpression `json:"EventCategories,omitempty"`  // List of strings
  SnsTopicArn *StringExpression `json:"SnsTopicArn,omitempty"`  // String
  SourceIds *StringListExpression `json:"SourceIds,omitempty"`  // List of strings
  SourceTypeX *StringExpression `json:"SourceType ,omitempty"`  // String
}

type AWSRDSEventSubscriptionList []AWSRDSEventSubscription

func (l *AWSRDSEventSubscriptionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSRDSEventSubscription{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSRDSEventSubscriptionList{item}
		return nil
	}
	list := []AWSRDSEventSubscription{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSRDSEventSubscriptionList(list)
		return nil
	}
	return err
}


type AWSRDSOptionGroup struct {
  EngineName *StringExpression `json:"EngineName,omitempty"`  // String
  MajorEngineVersion *StringExpression `json:"MajorEngineVersion,omitempty"`  // String
  OptionGroupDescription *StringExpression `json:"OptionGroupDescription,omitempty"`  // String
  OptionConfigurations *AmazonRDSOptionGroupOptionConfigurations `json:"OptionConfigurations,omitempty"`  // Amazon RDS OptionGroup OptionConfigurations
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
}

type AWSRDSOptionGroupList []AWSRDSOptionGroup

func (l *AWSRDSOptionGroupList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSRDSOptionGroup{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSRDSOptionGroupList{item}
		return nil
	}
	list := []AWSRDSOptionGroup{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSRDSOptionGroupList(list)
		return nil
	}
	return err
}


type AWSRoute53HealthCheck struct {
  HealthCheckConfig *AmazonRoute53HealthCheckConfig `json:"HealthCheckConfig,omitempty"`  // Amazon Route 53 HealthCheckConfig
  HealthCheckTags *AmazonRoute53HealthCheckTagsList `json:"HealthCheckTags,omitempty"`  // List of Amazon Route 53 HealthCheckTags
}

type AWSRoute53HealthCheckList []AWSRoute53HealthCheck

func (l *AWSRoute53HealthCheckList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSRoute53HealthCheck{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSRoute53HealthCheckList{item}
		return nil
	}
	list := []AWSRoute53HealthCheck{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSRoute53HealthCheckList(list)
		return nil
	}
	return err
}


type AWSRoute53HostedZone struct {
  HostedZoneConfig *AmazonRoute53HostedZoneConfigProperty `json:"HostedZoneConfig,omitempty"`  // Amazon Route 53 HostedZoneConfig Property
  HostedZoneTags *AmazonRoute53HostedZoneTagsList `json:"HostedZoneTags,omitempty"`  // List of Amazon Route 53 HostedZoneTags
  Name *StringExpression `json:"Name,omitempty"`  // String
  VPCs *AmazonRoute53HostedZoneVPCsList `json:"VPCs,omitempty"`  // List of Amazon Route 53 HostedZoneVPCs
}

type AWSRoute53HostedZoneList []AWSRoute53HostedZone

func (l *AWSRoute53HostedZoneList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSRoute53HostedZone{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSRoute53HostedZoneList{item}
		return nil
	}
	list := []AWSRoute53HostedZone{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSRoute53HostedZoneList(list)
		return nil
	}
	return err
}


type AWSRoute53RecordSet struct {
  AliasTarget *Route53AliasTargetProperty `json:"AliasTarget,omitempty"`  // AliasTarget
  Comment *StringExpression `json:"Comment,omitempty"`  // String
  Failover *StringExpression `json:"Failover,omitempty"`  // String
  GeoLocation *AmazonRoute53RecordSetGeoLocationProperty `json:"GeoLocation,omitempty"`  // Amazon Route 53 Record Set GeoLocation Property
  HealthCheckId *StringExpression `json:"HealthCheckId,omitempty"`  // String
  HostedZoneId *StringExpression `json:"HostedZoneId,omitempty"`  // String
  HostedZoneName *StringExpression `json:"HostedZoneName,omitempty"`  // String
  Name *StringExpression `json:"Name,omitempty"`  // String
  Region interface{} `json:"Region,omitempty"`  // 
  ResourceRecords *StringListExpression `json:"ResourceRecords,omitempty"`  // List of strings
  SetIdentifier *StringExpression `json:"SetIdentifier,omitempty"`  // String
  TTL *StringExpression `json:"TTL,omitempty"`  // String
  Type *StringExpression `json:"Type,omitempty"`  // String
  Weight Integer `json:"Weight,omitempty"`  // Number. Weight expects integer values
}

type AWSRoute53RecordSetList []AWSRoute53RecordSet

func (l *AWSRoute53RecordSetList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSRoute53RecordSet{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSRoute53RecordSetList{item}
		return nil
	}
	list := []AWSRoute53RecordSet{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSRoute53RecordSetList(list)
		return nil
	}
	return err
}


type AWSRoute53RecordSetGroup struct {
  HostedZoneId *StringExpression `json:"HostedZoneId,omitempty"`  // String
  HostedZoneName *StringExpression `json:"HostedZoneName,omitempty"`  // String
  RecordSets *AWSRoute53RecordSetList `json:"RecordSets,omitempty"`  // list of AWS::Route53::RecordSet
  Comment *StringExpression `json:"Comment,omitempty"`  // String
}

type AWSRoute53RecordSetGroupList []AWSRoute53RecordSetGroup

func (l *AWSRoute53RecordSetGroupList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSRoute53RecordSetGroup{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSRoute53RecordSetGroupList{item}
		return nil
	}
	list := []AWSRoute53RecordSetGroup{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSRoute53RecordSetGroupList(list)
		return nil
	}
	return err
}


type AWSS3Bucket struct {
  AccessControl *StringExpression `json:"AccessControl,omitempty"`  // String
  BucketName *StringExpression `json:"BucketName,omitempty"`  // String
  CorsConfiguration *AmazonS3CorsConfiguration `json:"CorsConfiguration,omitempty"`  // Amazon S3 Cors Configuration
  LifecycleConfiguration *AmazonS3LifecycleConfiguration `json:"LifecycleConfiguration,omitempty"`  // Amazon S3 Lifecycle Configuration
  LoggingConfiguration *AmazonS3LoggingConfiguration `json:"LoggingConfiguration,omitempty"`  // Amazon S3 Logging Configuration
  NotificationConfiguration *AmazonS3NotificationConfiguration `json:"NotificationConfiguration,omitempty"`  // Amazon S3 Notification Configuration
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
  VersioningConfiguration *AmazonS3VersioningConfiguration `json:"VersioningConfiguration,omitempty"`  // Amazon S3 Versioning Configuration
  WebsiteConfiguration *AmazonS3WebsiteConfigurationProperty `json:"WebsiteConfiguration,omitempty"`  // Website Configuration Type
}

type AWSS3BucketList []AWSS3Bucket

func (l *AWSS3BucketList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSS3Bucket{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSS3BucketList{item}
		return nil
	}
	list := []AWSS3Bucket{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSS3BucketList(list)
		return nil
	}
	return err
}


type AWSS3BucketPolicy struct {
  Bucket *StringExpression `json:"Bucket,omitempty"`  // String
  PolicyDocument interface{} `json:"PolicyDocument,omitempty"`  // JSON object
}

type AWSS3BucketPolicyList []AWSS3BucketPolicy

func (l *AWSS3BucketPolicyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSS3BucketPolicy{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSS3BucketPolicyList{item}
		return nil
	}
	list := []AWSS3BucketPolicy{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSS3BucketPolicyList(list)
		return nil
	}
	return err
}


type AWSSDBDomain struct {
}

type AWSSDBDomainList []AWSSDBDomain

func (l *AWSSDBDomainList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSSDBDomain{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSSDBDomainList{item}
		return nil
	}
	list := []AWSSDBDomain{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSSDBDomainList(list)
		return nil
	}
	return err
}


type AWSSNSTopic struct {
  DisplayName *StringExpression `json:"DisplayName,omitempty"`  // String
  Subscription *AmazonSNSSubscriptionList `json:"Subscription,omitempty"`  // List of SNS Subscriptions
  TopicName *StringExpression `json:"TopicName,omitempty"`  // String
}

type AWSSNSTopicList []AWSSNSTopic

func (l *AWSSNSTopicList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSSNSTopic{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSSNSTopicList{item}
		return nil
	}
	list := []AWSSNSTopic{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSSNSTopicList(list)
		return nil
	}
	return err
}


type AWSSNSTopicPolicy struct {
  PolicyDocument interface{} `json:"PolicyDocument,omitempty"`  // JSON object
  Topics interface{} `json:"Topics,omitempty"`  // A list of Amazon SNS topics ARNs
}

type AWSSNSTopicPolicyList []AWSSNSTopicPolicy

func (l *AWSSNSTopicPolicyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSSNSTopicPolicy{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSSNSTopicPolicyList{item}
		return nil
	}
	list := []AWSSNSTopicPolicy{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSSNSTopicPolicyList(list)
		return nil
	}
	return err
}


type AWSSQSQueue struct {
  DelaySeconds *Integer `json:"DelaySeconds,omitempty"`  // Integer
  MaximumMessageSize *Integer `json:"MaximumMessageSize,omitempty"`  // Integer
  MessageRetentionPeriod *Integer `json:"MessageRetentionPeriod,omitempty"`  // Integer
  QueueName *StringExpression `json:"QueueName,omitempty"`  // String
  ReceiveMessageWaitTimeSeconds *Integer `json:"ReceiveMessageWaitTimeSeconds,omitempty"`  // Integer
  RedrivePolicy *AmazonSQSRedrivePolicy `json:"RedrivePolicy,omitempty"`  // Amazon SQS RedrivePolicy
  VisibilityTimeout *Integer `json:"VisibilityTimeout,omitempty"`  // Integer
}

type AWSSQSQueueList []AWSSQSQueue

func (l *AWSSQSQueueList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSSQSQueue{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSSQSQueueList{item}
		return nil
	}
	list := []AWSSQSQueue{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSSQSQueueList(list)
		return nil
	}
	return err
}


type AWSSQSQueuePolicy struct {
  PolicyDocument interface{} `json:"PolicyDocument,omitempty"`  // JSON object
  Queues *StringListExpression `json:"Queues,omitempty"`  // List of strings
}

type AWSSQSQueuePolicyList []AWSSQSQueuePolicy

func (l *AWSSQSQueuePolicyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSSQSQueuePolicy{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSSQSQueuePolicyList{item}
		return nil
	}
	list := []AWSSQSQueuePolicy{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSSQSQueuePolicyList(list)
		return nil
	}
	return err
}


type AWSCloudFormationAutoScalingBlockDeviceMapping struct {
  DeviceName *StringExpression `json:"DeviceName,omitempty"`  // String
  Ebs *AWSCloudFormationAutoScalingEBSBlockDevice `json:"Ebs,omitempty"`  // AutoScaling EBS Block Device
  NoDevice *Bool `json:"NoDevice,omitempty"`  // Boolean
  VirtualName *StringExpression `json:"VirtualName,omitempty"`  // String
}

type AWSCloudFormationAutoScalingBlockDeviceMappingList []AWSCloudFormationAutoScalingBlockDeviceMapping

func (l *AWSCloudFormationAutoScalingBlockDeviceMappingList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSCloudFormationAutoScalingBlockDeviceMapping{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSCloudFormationAutoScalingBlockDeviceMappingList{item}
		return nil
	}
	list := []AWSCloudFormationAutoScalingBlockDeviceMapping{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSCloudFormationAutoScalingBlockDeviceMappingList(list)
		return nil
	}
	return err
}


type AWSCloudFormationAutoScalingEBSBlockDevice struct {
  DeleteOnTermination *Bool `json:"DeleteOnTermination,omitempty"`  // Boolean
  Iops *Integer `json:"Iops,omitempty"`  // Integer
  SnapshotId *StringExpression `json:"SnapshotId,omitempty"`  // String
  VolumeSize *Integer `json:"VolumeSize,omitempty"`  // Integer
  VolumeType *StringExpression `json:"VolumeType,omitempty"`  // String
}

type AWSCloudFormationAutoScalingEBSBlockDeviceList []AWSCloudFormationAutoScalingEBSBlockDevice

func (l *AWSCloudFormationAutoScalingEBSBlockDeviceList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSCloudFormationAutoScalingEBSBlockDevice{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSCloudFormationAutoScalingEBSBlockDeviceList{item}
		return nil
	}
	list := []AWSCloudFormationAutoScalingEBSBlockDevice{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSCloudFormationAutoScalingEBSBlockDeviceList(list)
		return nil
	}
	return err
}


type AutoScalingMetricsCollection struct {
  Granularity *StringExpression `json:"Granularity,omitempty"`  // String
  Metrics *StringListExpression `json:"Metrics,omitempty"`  // List of strings
}

type AutoScalingMetricsCollectionList []AutoScalingMetricsCollection

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


type AutoScalingNotificationConfigurations struct {
  NotificationTypes *StringListExpression `json:"NotificationTypes,omitempty"`  // List of strings
  TopicARN *StringExpression `json:"TopicARN,omitempty"`  // String
}

type AutoScalingNotificationConfigurationsList []AutoScalingNotificationConfigurations

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


type AutoScalingTags struct {
  Key *StringExpression `json:"Key,omitempty"`  // String
  Value *StringExpression `json:"Value,omitempty"`  // String
  PropagateAtLaunch *Bool `json:"PropagateAtLaunch,omitempty"`  // Boolean
}

type AutoScalingTagsList []AutoScalingTags

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


type CloudFormationStackParameters struct {
}

type CloudFormationStackParametersList []CloudFormationStackParameters

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


type CloudFrontDistributionConfig struct {
  Aliases *StringListExpression `json:"Aliases,omitempty"`  // List of strings
  CacheBehaviors *CloudFrontDistributionConfigCacheBehaviorList `json:"CacheBehaviors,omitempty"`  // List of CacheBehavior
  Comment *StringExpression `json:"Comment,omitempty"`  // String
  CustomErrorResponses *CloudFrontDistributionConfigCustomErrorResponse `json:"CustomErrorResponses,omitempty"`  // Type List of CloudFront DistributionConfig CustomErrorResponse
  DefaultCacheBehavior *CloudFrontDefaultCacheBehavior `json:"DefaultCacheBehavior,omitempty"`  // DefaultCacheBehavior type
  DefaultRootObject *StringExpression `json:"DefaultRootObject,omitempty"`  // String
  Enabled *Bool `json:"Enabled,omitempty"`  // Boolean
  Logging *CloudFrontLogging `json:"Logging,omitempty"`  // Logging type
  Origins *CloudFrontDistributionConfigOriginList `json:"Origins,omitempty"`  // List of Origins
  PriceClass *StringExpression `json:"PriceClass,omitempty"`  // String
  Restrictions *CloudFrontDistributionConfigurationRestrictions `json:"Restrictions,omitempty"`  // CloudFront DistributionConfiguration Restrictions
  ViewerCertificate *CloudFrontDistributionConfigurationViewerCertificate `json:"ViewerCertificate,omitempty"`  // CloudFront DistributionConfiguration ViewerCertificate
}

type CloudFrontDistributionConfigList []CloudFrontDistributionConfig

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


type CloudFrontDistributionConfigCacheBehavior struct {
  AllowedMethods *StringListExpression `json:"AllowedMethods,omitempty"`  // List of strings
  CachedMethods *StringListExpression `json:"CachedMethods,omitempty"`  // List of strings
  ForwardedValues *CloudFrontForwardedValues `json:"ForwardedValues,omitempty"`  // ForwardedValues type
  MinTTL *StringExpression `json:"MinTTL,omitempty"`  // String
  PathPattern *StringExpression `json:"PathPattern,omitempty"`  // String
  SmoothStreaming *Bool `json:"SmoothStreaming,omitempty"`  // Boolean
  TargetOriginId *StringExpression `json:"TargetOriginId,omitempty"`  // String
  TrustedSigners *StringListExpression `json:"TrustedSigners,omitempty"`  // List of strings
  ViewerProtocolPolicy *StringExpression `json:"ViewerProtocolPolicy,omitempty"`  // String
}

type CloudFrontDistributionConfigCacheBehaviorList []CloudFrontDistributionConfigCacheBehavior

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


type CloudFrontDistributionConfigCustomErrorResponse struct {
  ErrorCachingMinTTL *Integer `json:"ErrorCachingMinTTL,omitempty"`  // Integer
  ErrorCode *Integer `json:"ErrorCode,omitempty"`  // Integer
  ResponseCode *Integer `json:"ResponseCode,omitempty"`  // Integer
  ResponsePagePath *StringExpression `json:"ResponsePagePath,omitempty"`  // String
}

type CloudFrontDistributionConfigCustomErrorResponseList []CloudFrontDistributionConfigCustomErrorResponse

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


type CloudFrontDefaultCacheBehavior struct {
  AllowedMethods *StringListExpression `json:"AllowedMethods,omitempty"`  // List of strings
  CachedMethods *StringListExpression `json:"CachedMethods,omitempty"`  // List of strings
  ForwardedValues *CloudFrontForwardedValues `json:"ForwardedValues,omitempty"`  // ForwardedValues type
  MinTTL *StringExpression `json:"MinTTL,omitempty"`  // String
  SmoothStreaming *Bool `json:"SmoothStreaming,omitempty"`  // Boolean
  TargetOriginId *StringExpression `json:"TargetOriginId,omitempty"`  // String
  TrustedSigners *StringListExpression `json:"TrustedSigners,omitempty"`  // List of strings
  ViewerProtocolPolicy *StringExpression `json:"ViewerProtocolPolicy,omitempty"`  // String
}

type CloudFrontDefaultCacheBehaviorList []CloudFrontDefaultCacheBehavior

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


type CloudFrontLogging struct {
  Bucket *StringExpression `json:"Bucket,omitempty"`  // String
  IncludeCookies *Bool `json:"IncludeCookies,omitempty"`  // Boolean
  Prefix *StringExpression `json:"Prefix,omitempty"`  // String
}

type CloudFrontLoggingList []CloudFrontLogging

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


type CloudFrontDistributionConfigOrigin struct {
  CustomOriginConfig *CloudFrontDistributionConfigOriginCustomOrigin `json:"CustomOriginConfig,omitempty"`  // CustomOrigin type
  DomainName *StringExpression `json:"DomainName,omitempty"`  // String
  Id *StringExpression `json:"Id,omitempty"`  // String
  OriginPath *StringExpression `json:"OriginPath,omitempty"`  // String
  S3OriginConfig *CloudFrontDistributionConfigOriginS3Origin `json:"S3OriginConfig,omitempty"`  // S3Origin type
}

type CloudFrontDistributionConfigOriginList []CloudFrontDistributionConfigOrigin

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


type CloudFrontDistributionConfigOriginCustomOrigin struct {
  HTTPPort *StringExpression `json:"HTTPPort,omitempty"`  // String
  HTTPSPort *StringExpression `json:"HTTPSPort,omitempty"`  // String
  OriginProtocolPolicy *StringExpression `json:"OriginProtocolPolicy,omitempty"`  // String
}

type CloudFrontDistributionConfigOriginCustomOriginList []CloudFrontDistributionConfigOriginCustomOrigin

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


type CloudFrontDistributionConfigOriginS3Origin struct {
  OriginAccessIdentity *StringExpression `json:"OriginAccessIdentity,omitempty"`  // String
}

type CloudFrontDistributionConfigOriginS3OriginList []CloudFrontDistributionConfigOriginS3Origin

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


type CloudFrontDistributionConfigurationRestrictions struct {
  GeoRestriction *CloudFrontDistributionConfigRestrictionsGeoRestriction `json:"GeoRestriction,omitempty"`  // CloudFront DistributionConfig Restrictions GeoRestriction
}

type CloudFrontDistributionConfigurationRestrictionsList []CloudFrontDistributionConfigurationRestrictions

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


type CloudFrontDistributionConfigRestrictionsGeoRestriction struct {
  Locations *StringListExpression `json:"Locations,omitempty"`  // List of strings
  RestrictionType *StringExpression `json:"RestrictionType,omitempty"`  // String
  Blacklist interface{} `json:"blacklist,omitempty"`  // 
  Whitelist interface{} `json:"whitelist,omitempty"`  // 
  None interface{} `json:"none,omitempty"`  // 
}

type CloudFrontDistributionConfigRestrictionsGeoRestrictionList []CloudFrontDistributionConfigRestrictionsGeoRestriction

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


type CloudFrontDistributionConfigurationViewerCertificate struct {
  CloudFrontDefaultCertificate *Bool `json:"CloudFrontDefaultCertificate,omitempty"`  // Boolean
  IamCertificateId *StringExpression `json:"IamCertificateId,omitempty"`  // String
  MinimumProtocolVersion *StringExpression `json:"MinimumProtocolVersion,omitempty"`  // String
  SslSupportMethod *StringExpression `json:"SslSupportMethod,omitempty"`  // String
}

type CloudFrontDistributionConfigurationViewerCertificateList []CloudFrontDistributionConfigurationViewerCertificate

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


type CloudFrontForwardedValues struct {
  Cookies *CloudFrontForwardedValuesCookies `json:"Cookies,omitempty"`  // CloudFront ForwardedValues Cookies
  Headers *StringListExpression `json:"Headers,omitempty"`  // List of strings
  QueryString *Bool `json:"QueryString,omitempty"`  // Boolean
}

type CloudFrontForwardedValuesList []CloudFrontForwardedValues

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


type CloudFrontForwardedValuesCookies struct {
  Forward *StringExpression `json:"Forward,omitempty"`  // String
  WhitelistedNames *StringListExpression `json:"WhitelistedNames,omitempty"`  // List of strings
}

type CloudFrontForwardedValuesCookiesList []CloudFrontForwardedValuesCookies

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


type CloudWatchMetricDimension struct {
  Name *StringExpression `json:"Name,omitempty"`  // String
  Value *StringExpression `json:"Value,omitempty"`  // String
}

type CloudWatchMetricDimensionList []CloudWatchMetricDimension

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


type CloudWatchLogsMetricFilterMetricTransformationProperty struct {
  MetricName *StringExpression `json:"MetricName,omitempty"`  // String
  MetricNamespace *StringExpression `json:"MetricNamespace,omitempty"`  // String
  MetricValue *StringExpression `json:"MetricValue,omitempty"`  // String
}

type CloudWatchLogsMetricFilterMetricTransformationPropertyList []CloudWatchLogsMetricFilterMetricTransformationProperty

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


type AWSDataPipelinePipelineParameterObjects struct {
  Attributes *AWSDataPipelineParameterObjectsAttributes `json:"Attributes,omitempty"`  // AWS Data Pipeline Parameter Objects Attributes
  Id *StringExpression `json:"Id,omitempty"`  // String
}

type AWSDataPipelinePipelineParameterObjectsList []AWSDataPipelinePipelineParameterObjects

func (l *AWSDataPipelinePipelineParameterObjectsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSDataPipelinePipelineParameterObjects{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSDataPipelinePipelineParameterObjectsList{item}
		return nil
	}
	list := []AWSDataPipelinePipelineParameterObjects{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSDataPipelinePipelineParameterObjectsList(list)
		return nil
	}
	return err
}


type AWSDataPipelineParameterObjectsAttributes struct {
  Key *StringExpression `json:"Key,omitempty"`  // String
  StringValue *StringExpression `json:"StringValue,omitempty"`  // String
}

type AWSDataPipelineParameterObjectsAttributesList []AWSDataPipelineParameterObjectsAttributes

func (l *AWSDataPipelineParameterObjectsAttributesList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSDataPipelineParameterObjectsAttributes{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSDataPipelineParameterObjectsAttributesList{item}
		return nil
	}
	list := []AWSDataPipelineParameterObjectsAttributes{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSDataPipelineParameterObjectsAttributesList(list)
		return nil
	}
	return err
}


type AWSDataPipelinePipelineParameterValues struct {
  Id *StringExpression `json:"Id,omitempty"`  // String
  StringValue *StringExpression `json:"StringValue,omitempty"`  // String
}

type AWSDataPipelinePipelineParameterValuesList []AWSDataPipelinePipelineParameterValues

func (l *AWSDataPipelinePipelineParameterValuesList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSDataPipelinePipelineParameterValues{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSDataPipelinePipelineParameterValuesList{item}
		return nil
	}
	list := []AWSDataPipelinePipelineParameterValues{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSDataPipelinePipelineParameterValuesList(list)
		return nil
	}
	return err
}


type AWSDataPipelinePipelineObjects struct {
  Fields *AWSDataPipelineDataPipelineObjectFields `json:"Fields,omitempty"`  // AWS Data Pipeline Data Pipeline Object Fields
  Id *StringExpression `json:"Id,omitempty"`  // String
  Name *StringExpression `json:"Name,omitempty"`  // String
}

type AWSDataPipelinePipelineObjectsList []AWSDataPipelinePipelineObjects

func (l *AWSDataPipelinePipelineObjectsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSDataPipelinePipelineObjects{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSDataPipelinePipelineObjectsList{item}
		return nil
	}
	list := []AWSDataPipelinePipelineObjects{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSDataPipelinePipelineObjectsList(list)
		return nil
	}
	return err
}


type AWSDataPipelineDataPipelineObjectFields struct {
  Key *StringExpression `json:"Key,omitempty"`  // String
  RefValue *StringExpression `json:"RefValue,omitempty"`  // String
  StringValue *StringExpression `json:"StringValue,omitempty"`  // String
}

type AWSDataPipelineDataPipelineObjectFieldsList []AWSDataPipelineDataPipelineObjectFields

func (l *AWSDataPipelineDataPipelineObjectFieldsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSDataPipelineDataPipelineObjectFields{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSDataPipelineDataPipelineObjectFieldsList{item}
		return nil
	}
	list := []AWSDataPipelineDataPipelineObjectFields{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSDataPipelineDataPipelineObjectFieldsList(list)
		return nil
	}
	return err
}


type AWSDataPipelinePipelinePipelineTags struct {
  Key *StringExpression `json:"Key,omitempty"`  // String
  Value *StringExpression `json:"Value,omitempty"`  // String
}

type AWSDataPipelinePipelinePipelineTagsList []AWSDataPipelinePipelinePipelineTags

func (l *AWSDataPipelinePipelinePipelineTagsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSDataPipelinePipelinePipelineTags{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSDataPipelinePipelinePipelineTagsList{item}
		return nil
	}
	list := []AWSDataPipelinePipelinePipelineTags{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSDataPipelinePipelinePipelineTagsList(list)
		return nil
	}
	return err
}


type DynamoDBAttributeDefinitions struct {
  AttributeName *StringExpression `json:"AttributeName,omitempty"`  // String
  AttributeType *StringExpression `json:"AttributeType,omitempty"`  // String
}

type DynamoDBAttributeDefinitionsList []DynamoDBAttributeDefinitions

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


type DynamoDBGlobalSecondaryIndexes struct {
  IndexName *StringExpression `json:"IndexName,omitempty"`  // String
  KeySchema *DynamoDBKeySchema `json:"KeySchema,omitempty"`  // DynamoDB Key Schema
  Projection *DynamoDBProjectionObject `json:"Projection,omitempty"`  // DynamoDB Projection Object
  ProvisionedThroughput *DynamoDBProvisionedThroughput `json:"ProvisionedThroughput,omitempty"`  // DynamoDB Provisioned Throughput
}

type DynamoDBGlobalSecondaryIndexesList []DynamoDBGlobalSecondaryIndexes

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


type DynamoDBKeySchema struct {
  AttributeName *StringExpression `json:"AttributeName,omitempty"`  // String
  KeyType *StringExpression `json:"KeyType,omitempty"`  // String
}

type DynamoDBKeySchemaList []DynamoDBKeySchema

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


type DynamoDBLocalSecondaryIndexes struct {
  IndexName *StringExpression `json:"IndexName,omitempty"`  // String
  KeySchema *DynamoDBKeySchema `json:"KeySchema,omitempty"`  // DynamoDB Key Schema
  Projection *DynamoDBProjectionObject `json:"Projection,omitempty"`  // DynamoDB Projection Object
}

type DynamoDBLocalSecondaryIndexesList []DynamoDBLocalSecondaryIndexes

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


type DynamoDBProjectionObject struct {
  NonKeyAttributes *StringListExpression `json:"NonKeyAttributes,omitempty"`  // List of strings
  ProjectionType *StringExpression `json:"ProjectionType,omitempty"`  // String
  KEYSXONLY interface{} `json:"KEYS_ONLY,omitempty"`  // 
  INCLUDE interface{} `json:"INCLUDE,omitempty"`  // 
  ALL interface{} `json:"ALL,omitempty"`  // 
}

type DynamoDBProjectionObjectList []DynamoDBProjectionObject

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


type DynamoDBProvisionedThroughput struct {
}

type DynamoDBProvisionedThroughputList []DynamoDBProvisionedThroughput

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


type AmazonEC2BlockDeviceMappingProperty struct {
  DeviceName *StringExpression `json:"DeviceName,omitempty"`  // String
  Ebs *AmazonElasticBlockStoreBlockDeviceProperty `json:"Ebs,omitempty"`  // Amazon Elastic Block Store Block Device Property
  NoDevice interface{} `json:"NoDevice,omitempty"`  // an empty map: {}
  VirtualName *StringExpression `json:"VirtualName,omitempty"`  // String
}

type AmazonEC2BlockDeviceMappingPropertyList []AmazonEC2BlockDeviceMappingProperty

func (l *AmazonEC2BlockDeviceMappingPropertyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonEC2BlockDeviceMappingProperty{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonEC2BlockDeviceMappingPropertyList{item}
		return nil
	}
	list := []AmazonEC2BlockDeviceMappingProperty{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonEC2BlockDeviceMappingPropertyList(list)
		return nil
	}
	return err
}


type AmazonElasticBlockStoreBlockDeviceProperty struct {
  DeleteOnTermination *Bool `json:"DeleteOnTermination,omitempty"`  // Boolean
  Encrypted *Bool `json:"Encrypted,omitempty"`  // Boolean
  Iops *Integer `json:"Iops,omitempty"`  // Number
  SnapshotId *StringExpression `json:"SnapshotId,omitempty"`  // String
  VolumeSize *StringExpression `json:"VolumeSize,omitempty"`  // String
  VolumeType *StringExpression `json:"VolumeType,omitempty"`  // String
}

type AmazonElasticBlockStoreBlockDevicePropertyList []AmazonElasticBlockStoreBlockDeviceProperty

func (l *AmazonElasticBlockStoreBlockDevicePropertyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonElasticBlockStoreBlockDeviceProperty{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonElasticBlockStoreBlockDevicePropertyList{item}
		return nil
	}
	list := []AmazonElasticBlockStoreBlockDeviceProperty{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonElasticBlockStoreBlockDevicePropertyList(list)
		return nil
	}
	return err
}


type EC2ICMP struct {
}

type EC2ICMPList []EC2ICMP

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


type EC2MountPoint struct {
  Device *StringExpression `json:"Device,omitempty"`  // String
  VolumeId *StringExpression `json:"VolumeId,omitempty"`  // String
}

type EC2MountPointList []EC2MountPoint

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


type EC2NetworkInterfaceEmbedded struct {
  AssociatePublicIpAddress *Bool `json:"AssociatePublicIpAddress,omitempty"`  // Boolean
  DeleteOnTermination *Bool `json:"DeleteOnTermination,omitempty"`  // Boolean
  Description *StringExpression `json:"Description,omitempty"`  // String
  DeviceIndex *StringExpression `json:"DeviceIndex,omitempty"`  // String
  GroupSet *StringListExpression `json:"GroupSet,omitempty"`  // List of strings
  NetworkInterfaceId *StringExpression `json:"NetworkInterfaceId,omitempty"`  // String
  PrivateIpAddress *StringExpression `json:"PrivateIpAddress,omitempty"`  // String
  PrivateIpAddresses *EC2NetworkInterfacePrivateIPSpecificationList `json:"PrivateIpAddresses,omitempty"`  // list of PrivateIpAddressSpecification
  SecondaryPrivateIpAddressCount *Integer `json:"SecondaryPrivateIpAddressCount,omitempty"`  // Integer
  SubnetId *StringExpression `json:"SubnetId,omitempty"`  // String
}

type EC2NetworkInterfaceEmbeddedList []EC2NetworkInterfaceEmbedded

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


type EC2NetworkInterfaceAssociation struct {
  AttachmentID *StringExpression `json:"AttachmentID,omitempty"`  // String
  InstanceID *StringExpression `json:"InstanceID,omitempty"`  // String
  PublicIp *StringExpression `json:"PublicIp,omitempty"`  // String
  IpOwnerId *StringExpression `json:"IpOwnerId,omitempty"`  // String
}

type EC2NetworkInterfaceAssociationList []EC2NetworkInterfaceAssociation

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


type EC2NetworkInterfaceAttachment struct {
  AttachmentID *StringExpression `json:"AttachmentID,omitempty"`  // String
  InstanceID *StringExpression `json:"InstanceID,omitempty"`  // String
}

type EC2NetworkInterfaceAttachmentList []EC2NetworkInterfaceAttachment

func (l *EC2NetworkInterfaceAttachmentList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2NetworkInterfaceAttachment{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2NetworkInterfaceAttachmentList{item}
		return nil
	}
	list := []EC2NetworkInterfaceAttachment{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2NetworkInterfaceAttachmentList(list)
		return nil
	}
	return err
}


type EC2NetworkInterfaceGroupItem struct {
  Key *StringExpression `json:"Key,omitempty"`  // String
  Value *StringExpression `json:"Value,omitempty"`  // String
}

type EC2NetworkInterfaceGroupItemList []EC2NetworkInterfaceGroupItem

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


type EC2NetworkInterfacePrivateIPSpecification struct {
  PrivateIpAddress *StringExpression `json:"PrivateIpAddress,omitempty"`  // String
  Primary *Bool `json:"Primary,omitempty"`  // Boolean
}

type EC2NetworkInterfacePrivateIPSpecificationList []EC2NetworkInterfacePrivateIPSpecification

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


type EC2PortRange struct {
}

type EC2PortRangeList []EC2PortRange

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


type EC2SecurityGroupRule struct {
  CidrIp *StringExpression `json:"CidrIp,omitempty"`  // String
  DestinationSecurityGroupIdXXSecurityGroupEgressXOnlyX *StringExpression `json:"DestinationSecurityGroupId (SecurityGroupEgress only),omitempty"`  // String
  FromPort *Integer `json:"FromPort,omitempty"`  // Integer
  IpProtocol *StringExpression `json:"IpProtocol,omitempty"`  // String
  SourceSecurityGroupIdXXSecurityGroupIngressXOnlyX *StringExpression `json:"SourceSecurityGroupId (SecurityGroupIngress only),omitempty"`  // String
  SourceSecurityGroupNameXXSecurityGroupIngressXOnlyX *StringExpression `json:"SourceSecurityGroupName (SecurityGroupIngress only),omitempty"`  // String
  SourceSecurityGroupOwnerIdXXSecurityGroupIngressXOnlyX *StringExpression `json:"SourceSecurityGroupOwnerId (SecurityGroupIngress only),omitempty"`  // String
  ToPort *Integer `json:"ToPort,omitempty"`  // Integer
}

type EC2SecurityGroupRuleList []EC2SecurityGroupRule

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


type AmazonEC2ContainerServiceServiceLoadBalancers struct {
  ContainerName *StringExpression `json:"ContainerName,omitempty"`  // String
  ContainerPort *Integer `json:"ContainerPort,omitempty"`  // Integer
  LoadBalancerName *StringExpression `json:"LoadBalancerName,omitempty"`  // String
}

type AmazonEC2ContainerServiceServiceLoadBalancersList []AmazonEC2ContainerServiceServiceLoadBalancers

func (l *AmazonEC2ContainerServiceServiceLoadBalancersList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonEC2ContainerServiceServiceLoadBalancers{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonEC2ContainerServiceServiceLoadBalancersList{item}
		return nil
	}
	list := []AmazonEC2ContainerServiceServiceLoadBalancers{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonEC2ContainerServiceServiceLoadBalancersList(list)
		return nil
	}
	return err
}


type AmazonEC2ContainerServiceTaskDefinitionContainerDefinitions struct {
  Command *StringListExpression `json:"Command,omitempty"`  // List of strings
  Cpu *Integer `json:"Cpu,omitempty"`  // Integer
  EntryPoint *StringListExpression `json:"EntryPoint,omitempty"`  // List of strings
  Environment *AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsEnvironmentList `json:"Environment,omitempty"`  // List of Amazon EC2 Container Service TaskDefinition ContainerDefinitions Environment
  Essential *Bool `json:"Essential,omitempty"`  // Boolean
  Image *StringExpression `json:"Image,omitempty"`  // String
  Links *StringListExpression `json:"Links,omitempty"`  // List of strings
  Memory *Integer `json:"Memory,omitempty"`  // Integer
  MountPoints *AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsMountPointsList `json:"MountPoints,omitempty"`  // List of Amazon EC2 Container Service TaskDefinition ContainerDefinitions MountPoints
  Name *StringExpression `json:"Name,omitempty"`  // String
  PortMappings *AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsPortMappingsList `json:"PortMappings,omitempty"`  // List of Amazon EC2 Container Service TaskDefinition ContainerDefinitions PortMappings
  VolumesFrom *AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsVolumesFromList `json:"VolumesFrom,omitempty"`  // List of Amazon EC2 Container Service TaskDefinition ContainerDefinitions VolumesFrom
}

type AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsList []AmazonEC2ContainerServiceTaskDefinitionContainerDefinitions

func (l *AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonEC2ContainerServiceTaskDefinitionContainerDefinitions{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsList{item}
		return nil
	}
	list := []AmazonEC2ContainerServiceTaskDefinitionContainerDefinitions{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsList(list)
		return nil
	}
	return err
}


type AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsEnvironment struct {
  Name *StringExpression `json:"Name,omitempty"`  // String
  Value *StringExpression `json:"Value,omitempty"`  // String
}

type AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsEnvironmentList []AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsEnvironment

func (l *AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsEnvironmentList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsEnvironment{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsEnvironmentList{item}
		return nil
	}
	list := []AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsEnvironment{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsEnvironmentList(list)
		return nil
	}
	return err
}


type AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsMountPoints struct {
  ContainerPath *StringExpression `json:"ContainerPath,omitempty"`  // String
  SourceVolume *StringExpression `json:"SourceVolume,omitempty"`  // String
  ReadOnly *Bool `json:"ReadOnly,omitempty"`  // Boolean
}

type AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsMountPointsList []AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsMountPoints

func (l *AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsMountPointsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsMountPoints{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsMountPointsList{item}
		return nil
	}
	list := []AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsMountPoints{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsMountPointsList(list)
		return nil
	}
	return err
}


type AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsPortMappings struct {
  ContainerPort *Integer `json:"ContainerPort,omitempty"`  // Integer
  HostPort *Integer `json:"HostPort,omitempty"`  // Integer
}

type AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsPortMappingsList []AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsPortMappings

func (l *AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsPortMappingsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsPortMappings{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsPortMappingsList{item}
		return nil
	}
	list := []AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsPortMappings{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsPortMappingsList(list)
		return nil
	}
	return err
}


type AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsVolumesFrom struct {
  SourceContainer *StringExpression `json:"SourceContainer,omitempty"`  // String
  ReadOnly *Bool `json:"ReadOnly,omitempty"`  // Boolean
}

type AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsVolumesFromList []AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsVolumesFrom

func (l *AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsVolumesFromList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsVolumesFrom{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsVolumesFromList{item}
		return nil
	}
	list := []AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsVolumesFrom{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsVolumesFromList(list)
		return nil
	}
	return err
}


type AmazonEC2ContainerServiceTaskDefinitionVolumes struct {
  Name *StringExpression `json:"Name,omitempty"`  // String
  Host *AmazonEC2ContainerServiceTaskDefinitionVolumesHost `json:"Host,omitempty"`  // Amazon EC2 Container Service TaskDefinition Volumes Host
}

type AmazonEC2ContainerServiceTaskDefinitionVolumesList []AmazonEC2ContainerServiceTaskDefinitionVolumes

func (l *AmazonEC2ContainerServiceTaskDefinitionVolumesList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonEC2ContainerServiceTaskDefinitionVolumes{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonEC2ContainerServiceTaskDefinitionVolumesList{item}
		return nil
	}
	list := []AmazonEC2ContainerServiceTaskDefinitionVolumes{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonEC2ContainerServiceTaskDefinitionVolumesList(list)
		return nil
	}
	return err
}


type AmazonEC2ContainerServiceTaskDefinitionVolumesHost struct {
  SourcePath *StringExpression `json:"SourcePath,omitempty"`  // String
}

type AmazonEC2ContainerServiceTaskDefinitionVolumesHostList []AmazonEC2ContainerServiceTaskDefinitionVolumesHost

func (l *AmazonEC2ContainerServiceTaskDefinitionVolumesHostList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonEC2ContainerServiceTaskDefinitionVolumesHost{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonEC2ContainerServiceTaskDefinitionVolumesHostList{item}
		return nil
	}
	list := []AmazonEC2ContainerServiceTaskDefinitionVolumesHost{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonEC2ContainerServiceTaskDefinitionVolumesHostList(list)
		return nil
	}
	return err
}


type ElasticBeanstalkEnvironmentTier struct {
}

type ElasticBeanstalkEnvironmentTierList []ElasticBeanstalkEnvironmentTier

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


type ElasticBeanstalkOptionSettings struct {
}

type ElasticBeanstalkOptionSettingsList []ElasticBeanstalkOptionSettings

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


type ElasticBeanstalkSourceBundle struct {
}

type ElasticBeanstalkSourceBundleList []ElasticBeanstalkSourceBundle

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


type ElasticBeanstalkSourceConfiguration struct {
}

type ElasticBeanstalkSourceConfigurationList []ElasticBeanstalkSourceConfiguration

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


type ElasticLoadBalancingAccessLoggingPolicy struct {
  EmitInterval *Integer `json:"EmitInterval,omitempty"`  // Integer
  Enabled *Bool `json:"Enabled,omitempty"`  // Boolean
  S3BucketName *StringExpression `json:"S3BucketName,omitempty"`  // String
  S3BucketPrefix *StringExpression `json:"S3BucketPrefix,omitempty"`  // String
}

type ElasticLoadBalancingAccessLoggingPolicyList []ElasticLoadBalancingAccessLoggingPolicy

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


type ElasticLoadBalancingAppCookieStickinessPolicy struct {
  CookieName *StringExpression `json:"CookieName,omitempty"`  // String
  PolicyName *StringExpression `json:"PolicyName,omitempty"`  // String
}

type ElasticLoadBalancingAppCookieStickinessPolicyList []ElasticLoadBalancingAppCookieStickinessPolicy

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


type ElasticLoadBalancingConnectionDrainingPolicy struct {
  Enabled *Bool `json:"Enabled,omitempty"`  // Boolean
  Timeout *Integer `json:"Timeout,omitempty"`  // Integer
}

type ElasticLoadBalancingConnectionDrainingPolicyList []ElasticLoadBalancingConnectionDrainingPolicy

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


type ElasticLoadBalancingConnectionSettings struct {
  IdleTimeout *Integer `json:"IdleTimeout,omitempty"`  // Integer
}

type ElasticLoadBalancingConnectionSettingsList []ElasticLoadBalancingConnectionSettings

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


type ElasticLoadBalancingHealthCheck struct {
  HealthyThreshold *StringExpression `json:"HealthyThreshold,omitempty"`  // String
  Interval *StringExpression `json:"Interval,omitempty"`  // String
  Target *StringExpression `json:"Target,omitempty"`  // String
  Timeout *StringExpression `json:"Timeout,omitempty"`  // String
  UnhealthyThreshold *StringExpression `json:"UnhealthyThreshold,omitempty"`  // String
}

type ElasticLoadBalancingHealthCheckList []ElasticLoadBalancingHealthCheck

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


type ElasticLoadBalancingLBCookieStickinessPolicy struct {
  CookieExpirationPeriod *StringExpression `json:"CookieExpirationPeriod,omitempty"`  // String
  PolicyName interface{} `json:"PolicyName,omitempty"`  // 
}

type ElasticLoadBalancingLBCookieStickinessPolicyList []ElasticLoadBalancingLBCookieStickinessPolicy

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


type ElasticLoadBalancingListener struct {
  InstancePort *StringExpression `json:"InstancePort,omitempty"`  // String
  InstanceProtocol *StringExpression `json:"InstanceProtocol,omitempty"`  // String
  LoadBalancerPort *StringExpression `json:"LoadBalancerPort,omitempty"`  // String
  PolicyNames *StringListExpression `json:"PolicyNames,omitempty"`  // List of strings
  Protocol *StringExpression `json:"Protocol,omitempty"`  // String
  SSLCertificateId *StringExpression `json:"SSLCertificateId,omitempty"`  // String
}

type ElasticLoadBalancingListenerList []ElasticLoadBalancingListener

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


type ElasticLoadBalancingPolicy struct {
  Attributes interface{} `json:"Attributes,omitempty"`  // List of JSON name-value pairs
  InstancePorts interface{} `json:"InstancePorts,omitempty"`  // List of String
  LoadBalancerPorts interface{} `json:"LoadBalancerPorts,omitempty"`  // List of String
  PolicyName *StringExpression `json:"PolicyName,omitempty"`  // String
  PolicyType *StringExpression `json:"PolicyType,omitempty"`  // String
}

type ElasticLoadBalancingPolicyList []ElasticLoadBalancingPolicy

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


type IAMPolicies struct {
  PolicyDocument interface{} `json:"PolicyDocument,omitempty"`  // JSON object
  PolicyName *StringExpression `json:"PolicyName,omitempty"`  // String
}

type IAMPoliciesList []IAMPolicies

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


type IAMUserLoginProfile struct {
  Password *StringExpression `json:"Password,omitempty"`  // String
  PasswordResetRequired *Bool `json:"PasswordResetRequired,omitempty"`  // Boolean
}

type IAMUserLoginProfileList []IAMUserLoginProfile

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


type AWSLambdaFunctionCode struct {
  S3Bucket *StringExpression `json:"S3Bucket,omitempty"`  // String
  S3Key *StringExpression `json:"S3Key,omitempty"`  // String
  S3ObjectVersion *StringExpression `json:"S3ObjectVersion,omitempty"`  // String
}

type AWSLambdaFunctionCodeList []AWSLambdaFunctionCode

func (l *AWSLambdaFunctionCodeList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSLambdaFunctionCode{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSLambdaFunctionCodeList{item}
		return nil
	}
	list := []AWSLambdaFunctionCode{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSLambdaFunctionCodeList(list)
		return nil
	}
	return err
}


type Name struct {
}

type NameList []Name

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


type AWSOpsWorksAutoScalingThresholds struct {
  CpuThreshold *Integer `json:"CpuThreshold,omitempty"`  // Number
  IgnoreMetricsTime *Integer `json:"IgnoreMetricsTime,omitempty"`  // Integer
  InstanceCount *Integer `json:"InstanceCount,omitempty"`  // Integer
  LoadThreshold *Integer `json:"LoadThreshold,omitempty"`  // Number
  MemoryThreshold *Integer `json:"MemoryThreshold,omitempty"`  // Number
  ThresholdsWaitTime *Integer `json:"ThresholdsWaitTime,omitempty"`  // Integer
}

type AWSOpsWorksAutoScalingThresholdsList []AWSOpsWorksAutoScalingThresholds

func (l *AWSOpsWorksAutoScalingThresholdsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSOpsWorksAutoScalingThresholds{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSOpsWorksAutoScalingThresholdsList{item}
		return nil
	}
	list := []AWSOpsWorksAutoScalingThresholds{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSOpsWorksAutoScalingThresholdsList(list)
		return nil
	}
	return err
}


type AWSOpsWorksChefConfiguration struct {
  BerkshelfVersion *StringExpression `json:"BerkshelfVersion,omitempty"`  // String
  ManageBerkshelf *Bool `json:"ManageBerkshelf,omitempty"`  // Boolean
}

type AWSOpsWorksChefConfigurationList []AWSOpsWorksChefConfiguration

func (l *AWSOpsWorksChefConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSOpsWorksChefConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSOpsWorksChefConfigurationList{item}
		return nil
	}
	list := []AWSOpsWorksChefConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSOpsWorksChefConfigurationList(list)
		return nil
	}
	return err
}


type AWSOpsWorksLayerLifeCycleConfiguration struct {
  ShutdownEventConfiguration *AWSOpsWorksLayerLifeCycleConfigurationShutdownEventConfiguration `json:"ShutdownEventConfiguration,omitempty"`  // AWS OpsWorks Layer LifeCycleConfiguration ShutdownEventConfiguration
}

type AWSOpsWorksLayerLifeCycleConfigurationList []AWSOpsWorksLayerLifeCycleConfiguration

func (l *AWSOpsWorksLayerLifeCycleConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSOpsWorksLayerLifeCycleConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSOpsWorksLayerLifeCycleConfigurationList{item}
		return nil
	}
	list := []AWSOpsWorksLayerLifeCycleConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSOpsWorksLayerLifeCycleConfigurationList(list)
		return nil
	}
	return err
}


type AWSOpsWorksLayerLifeCycleConfigurationShutdownEventConfiguration struct {
  DelayUntilElbConnectionsDrained *Bool `json:"DelayUntilElbConnectionsDrained,omitempty"`  // Boolean
  ExecutionTimeout *Integer `json:"ExecutionTimeout,omitempty"`  // Integer
}

type AWSOpsWorksLayerLifeCycleConfigurationShutdownEventConfigurationList []AWSOpsWorksLayerLifeCycleConfigurationShutdownEventConfiguration

func (l *AWSOpsWorksLayerLifeCycleConfigurationShutdownEventConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSOpsWorksLayerLifeCycleConfigurationShutdownEventConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSOpsWorksLayerLifeCycleConfigurationShutdownEventConfigurationList{item}
		return nil
	}
	list := []AWSOpsWorksLayerLifeCycleConfigurationShutdownEventConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSOpsWorksLayerLifeCycleConfigurationShutdownEventConfigurationList(list)
		return nil
	}
	return err
}


type AWSOpsWorksLoadBasedAutoScaling struct {
  DownScaling *AWSOpsWorksAutoScalingThresholds `json:"DownScaling,omitempty"`  // AWS OpsWorks AutoScalingThresholds Type
  Enable *Bool `json:"Enable,omitempty"`  // Boolean
  UpScaling *AWSOpsWorksAutoScalingThresholds `json:"UpScaling,omitempty"`  // AWS OpsWorks AutoScalingThresholds Type
}

type AWSOpsWorksLoadBasedAutoScalingList []AWSOpsWorksLoadBasedAutoScaling

func (l *AWSOpsWorksLoadBasedAutoScalingList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSOpsWorksLoadBasedAutoScaling{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSOpsWorksLoadBasedAutoScalingList{item}
		return nil
	}
	list := []AWSOpsWorksLoadBasedAutoScaling{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSOpsWorksLoadBasedAutoScalingList(list)
		return nil
	}
	return err
}


type AWSOpsWorksRecipes struct {
  Configure *StringListExpression `json:"Configure,omitempty"`  // List of strings
  Deploy *StringListExpression `json:"Deploy,omitempty"`  // List of strings
  Setup *StringListExpression `json:"Setup,omitempty"`  // List of strings
  Shutdown *StringListExpression `json:"Shutdown,omitempty"`  // List of strings
  Undeploy *StringListExpression `json:"Undeploy,omitempty"`  // List of strings
}

type AWSOpsWorksRecipesList []AWSOpsWorksRecipes

func (l *AWSOpsWorksRecipesList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSOpsWorksRecipes{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSOpsWorksRecipesList{item}
		return nil
	}
	list := []AWSOpsWorksRecipes{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSOpsWorksRecipesList(list)
		return nil
	}
	return err
}


type AWSOpsWorksSource struct {
  Password *StringExpression `json:"Password,omitempty"`  // String
  Revision *StringExpression `json:"Revision,omitempty"`  // String
  SshKey *StringExpression `json:"SshKey,omitempty"`  // String
  Type *StringExpression `json:"Type,omitempty"`  // String
  Url *StringExpression `json:"Url,omitempty"`  // String
  Username *StringExpression `json:"Username,omitempty"`  // String
}

type AWSOpsWorksSourceList []AWSOpsWorksSource

func (l *AWSOpsWorksSourceList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSOpsWorksSource{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSOpsWorksSourceList{item}
		return nil
	}
	list := []AWSOpsWorksSource{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSOpsWorksSourceList(list)
		return nil
	}
	return err
}


type AWSOpsWorksSslConfiguration struct {
  Certificate *StringExpression `json:"Certificate,omitempty"`  // String
  Chain *StringExpression `json:"Chain,omitempty"`  // String
  PrivateKey *StringExpression `json:"PrivateKey,omitempty"`  // String
}

type AWSOpsWorksSslConfigurationList []AWSOpsWorksSslConfiguration

func (l *AWSOpsWorksSslConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSOpsWorksSslConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSOpsWorksSslConfigurationList{item}
		return nil
	}
	list := []AWSOpsWorksSslConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSOpsWorksSslConfigurationList(list)
		return nil
	}
	return err
}


type AWSOpsWorksStackConfigurationManager struct {
  Name *StringExpression `json:"Name,omitempty"`  // String
  Version *StringExpression `json:"Version,omitempty"`  // String
}

type AWSOpsWorksStackConfigurationManagerList []AWSOpsWorksStackConfigurationManager

func (l *AWSOpsWorksStackConfigurationManagerList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSOpsWorksStackConfigurationManager{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSOpsWorksStackConfigurationManagerList{item}
		return nil
	}
	list := []AWSOpsWorksStackConfigurationManager{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSOpsWorksStackConfigurationManagerList(list)
		return nil
	}
	return err
}


type AWSOpsWorksTimeBasedAutoScaling struct {
  Friday *StringExpression `json:"Friday,omitempty"`  // String to string map
  Monday *StringExpression `json:"Monday,omitempty"`  // String to string map
  Saturday *StringExpression `json:"Saturday,omitempty"`  // String to string map
  Sunday *StringExpression `json:"Sunday,omitempty"`  // String to string map
  Thursday *StringExpression `json:"Thursday,omitempty"`  // String to string map
  Tuesday *StringExpression `json:"Tuesday,omitempty"`  // String to string map
  Wednesday *StringExpression `json:"Wednesday,omitempty"`  // String to string map
}

type AWSOpsWorksTimeBasedAutoScalingList []AWSOpsWorksTimeBasedAutoScaling

func (l *AWSOpsWorksTimeBasedAutoScalingList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSOpsWorksTimeBasedAutoScaling{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSOpsWorksTimeBasedAutoScalingList{item}
		return nil
	}
	list := []AWSOpsWorksTimeBasedAutoScaling{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSOpsWorksTimeBasedAutoScalingList(list)
		return nil
	}
	return err
}


type AWSOpsWorksVolumeConfiguration struct {
  Iops *Integer `json:"Iops,omitempty"`  // Integer
  MountPoint *StringExpression `json:"MountPoint,omitempty"`  // String
  NumberOfDisks *Integer `json:"NumberOfDisks,omitempty"`  // Integer
  RaidLevel *Integer `json:"RaidLevel,omitempty"`  // Integer
  Size *Integer `json:"Size,omitempty"`  // Integer
  VolumeType *StringExpression `json:"VolumeType,omitempty"`  // String
}

type AWSOpsWorksVolumeConfigurationList []AWSOpsWorksVolumeConfiguration

func (l *AWSOpsWorksVolumeConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AWSOpsWorksVolumeConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AWSOpsWorksVolumeConfigurationList{item}
		return nil
	}
	list := []AWSOpsWorksVolumeConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AWSOpsWorksVolumeConfigurationList(list)
		return nil
	}
	return err
}


type AmazonRedshiftParameter struct {
  ParameterName *StringExpression `json:"ParameterName,omitempty"`  // String
  ParameterValue *StringExpression `json:"ParameterValue,omitempty"`  // String
}

type AmazonRedshiftParameterList []AmazonRedshiftParameter

func (l *AmazonRedshiftParameterList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonRedshiftParameter{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonRedshiftParameterList{item}
		return nil
	}
	list := []AmazonRedshiftParameter{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonRedshiftParameterList(list)
		return nil
	}
	return err
}


type ResourceTag struct {
  Key *StringExpression `json:"Key,omitempty"`  // String
  Value *StringExpression `json:"Value,omitempty"`  // String
}

type ResourceTagList []ResourceTag

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


type AmazonRDSOptionGroupOptionConfigurations struct {
  DBSecurityGroupMemberships *StringListExpression `json:"DBSecurityGroupMemberships,omitempty"`  // List of strings
  OptionName *StringExpression `json:"OptionName,omitempty"`  // String
  OptionSettings *AmazonRDSOptionGroupOptionConfigurationsOptionSettings `json:"OptionSettings,omitempty"`  // Amazon RDS OptionGroup OptionConfigurations OptionSettings
  Port *Integer `json:"Port,omitempty"`  // Integer
  VpcSecurityGroupMemberships *StringListExpression `json:"VpcSecurityGroupMemberships,omitempty"`  // List of strings
}

type AmazonRDSOptionGroupOptionConfigurationsList []AmazonRDSOptionGroupOptionConfigurations

func (l *AmazonRDSOptionGroupOptionConfigurationsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonRDSOptionGroupOptionConfigurations{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonRDSOptionGroupOptionConfigurationsList{item}
		return nil
	}
	list := []AmazonRDSOptionGroupOptionConfigurations{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonRDSOptionGroupOptionConfigurationsList(list)
		return nil
	}
	return err
}


type AmazonRDSOptionGroupOptionConfigurationsOptionSettings struct {
  Name *StringExpression `json:"Name,omitempty"`  // String
  Value *StringExpression `json:"Value,omitempty"`  // String
}

type AmazonRDSOptionGroupOptionConfigurationsOptionSettingsList []AmazonRDSOptionGroupOptionConfigurationsOptionSettings

func (l *AmazonRDSOptionGroupOptionConfigurationsOptionSettingsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonRDSOptionGroupOptionConfigurationsOptionSettings{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonRDSOptionGroupOptionConfigurationsOptionSettingsList{item}
		return nil
	}
	list := []AmazonRDSOptionGroupOptionConfigurationsOptionSettings{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonRDSOptionGroupOptionConfigurationsOptionSettingsList(list)
		return nil
	}
	return err
}


type AmazonRDSSecurityGroupRule struct {
  CIDRIP *StringExpression `json:"CIDRIP,omitempty"`  // String
  EC2SecurityGroupId *StringExpression `json:"EC2SecurityGroupId,omitempty"`  // String
  EC2SecurityGroupName *StringExpression `json:"EC2SecurityGroupName,omitempty"`  // String
  EC2SecurityGroupOwnerId *StringExpression `json:"EC2SecurityGroupOwnerId,omitempty"`  // String
}

type AmazonRDSSecurityGroupRuleList []AmazonRDSSecurityGroupRule

func (l *AmazonRDSSecurityGroupRuleList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonRDSSecurityGroupRule{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonRDSSecurityGroupRuleList{item}
		return nil
	}
	list := []AmazonRDSSecurityGroupRule{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonRDSSecurityGroupRuleList(list)
		return nil
	}
	return err
}


type Route53AliasTargetProperty struct {
  DNSName *StringExpression `json:"DNSName,omitempty"`  // String
  EvaluateTargetHealth *Bool `json:"EvaluateTargetHealth,omitempty"`  // Boolean
  HostedZoneId *StringExpression `json:"HostedZoneId,omitempty"`  // String
}

type Route53AliasTargetPropertyList []Route53AliasTargetProperty

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


type AmazonRoute53RecordSetGeoLocationProperty struct {
  ContinentCode *StringExpression `json:"ContinentCode,omitempty"`  // String
  CountryCode *StringExpression `json:"CountryCode,omitempty"`  // String
  SubdivisionCode *StringExpression `json:"SubdivisionCode,omitempty"`  // String
}

type AmazonRoute53RecordSetGeoLocationPropertyList []AmazonRoute53RecordSetGeoLocationProperty

func (l *AmazonRoute53RecordSetGeoLocationPropertyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonRoute53RecordSetGeoLocationProperty{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonRoute53RecordSetGeoLocationPropertyList{item}
		return nil
	}
	list := []AmazonRoute53RecordSetGeoLocationProperty{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonRoute53RecordSetGeoLocationPropertyList(list)
		return nil
	}
	return err
}


type AmazonRoute53HealthCheckConfig struct {
  FailureThreshold *Integer `json:"FailureThreshold,omitempty"`  // Integer
  FullyQualifiedDomainName *StringExpression `json:"FullyQualifiedDomainName,omitempty"`  // String
  IPAddress *StringExpression `json:"IPAddress,omitempty"`  // String
  Port *Integer `json:"Port,omitempty"`  // Integer
  RequestInterval *Integer `json:"RequestInterval,omitempty"`  // Integer
  ResourcePath *StringExpression `json:"ResourcePath,omitempty"`  // String
  SearchString *StringExpression `json:"SearchString,omitempty"`  // String
  Type *StringExpression `json:"Type,omitempty"`  // String
}

type AmazonRoute53HealthCheckConfigList []AmazonRoute53HealthCheckConfig

func (l *AmazonRoute53HealthCheckConfigList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonRoute53HealthCheckConfig{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonRoute53HealthCheckConfigList{item}
		return nil
	}
	list := []AmazonRoute53HealthCheckConfig{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonRoute53HealthCheckConfigList(list)
		return nil
	}
	return err
}


type AmazonRoute53HealthCheckTags struct {
  Key *StringExpression `json:"Key,omitempty"`  // String
  Value *StringExpression `json:"Value,omitempty"`  // String
}

type AmazonRoute53HealthCheckTagsList []AmazonRoute53HealthCheckTags

func (l *AmazonRoute53HealthCheckTagsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonRoute53HealthCheckTags{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonRoute53HealthCheckTagsList{item}
		return nil
	}
	list := []AmazonRoute53HealthCheckTags{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonRoute53HealthCheckTagsList(list)
		return nil
	}
	return err
}


type AmazonRoute53HostedZoneConfigProperty struct {
  Comment *StringExpression `json:"Comment,omitempty"`  // String
}

type AmazonRoute53HostedZoneConfigPropertyList []AmazonRoute53HostedZoneConfigProperty

func (l *AmazonRoute53HostedZoneConfigPropertyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonRoute53HostedZoneConfigProperty{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonRoute53HostedZoneConfigPropertyList{item}
		return nil
	}
	list := []AmazonRoute53HostedZoneConfigProperty{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonRoute53HostedZoneConfigPropertyList(list)
		return nil
	}
	return err
}


type AmazonRoute53HostedZoneTags struct {
  Key *StringExpression `json:"Key,omitempty"`  // String
  Value *StringExpression `json:"Value,omitempty"`  // String
}

type AmazonRoute53HostedZoneTagsList []AmazonRoute53HostedZoneTags

func (l *AmazonRoute53HostedZoneTagsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonRoute53HostedZoneTags{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonRoute53HostedZoneTagsList{item}
		return nil
	}
	list := []AmazonRoute53HostedZoneTags{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonRoute53HostedZoneTagsList(list)
		return nil
	}
	return err
}


type AmazonRoute53HostedZoneVPCs struct {
  VPCId *StringExpression `json:"VPCId,omitempty"`  // String
  VPCRegion *StringExpression `json:"VPCRegion,omitempty"`  // String
}

type AmazonRoute53HostedZoneVPCsList []AmazonRoute53HostedZoneVPCs

func (l *AmazonRoute53HostedZoneVPCsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonRoute53HostedZoneVPCs{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonRoute53HostedZoneVPCsList{item}
		return nil
	}
	list := []AmazonRoute53HostedZoneVPCs{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonRoute53HostedZoneVPCsList(list)
		return nil
	}
	return err
}


type AmazonS3CorsConfiguration struct {
  CorsRules *AmazonS3CorsConfigurationRule `json:"CorsRules,omitempty"`  // Amazon S3 Cors Configuration Rule
}

type AmazonS3CorsConfigurationList []AmazonS3CorsConfiguration

func (l *AmazonS3CorsConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonS3CorsConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonS3CorsConfigurationList{item}
		return nil
	}
	list := []AmazonS3CorsConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonS3CorsConfigurationList(list)
		return nil
	}
	return err
}


type AmazonS3CorsConfigurationRule struct {
  AllowedHeaders *StringListExpression `json:"AllowedHeaders,omitempty"`  // List of strings
  AllowedMethods *StringListExpression `json:"AllowedMethods,omitempty"`  // List of strings
  AllowedOrigins *StringListExpression `json:"AllowedOrigins,omitempty"`  // List of strings
  ExposedHeaders *StringListExpression `json:"ExposedHeaders,omitempty"`  // List of strings
  Id *StringExpression `json:"Id,omitempty"`  // String
  MaxAge *Integer `json:"MaxAge,omitempty"`  // Integer
}

type AmazonS3CorsConfigurationRuleList []AmazonS3CorsConfigurationRule

func (l *AmazonS3CorsConfigurationRuleList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonS3CorsConfigurationRule{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonS3CorsConfigurationRuleList{item}
		return nil
	}
	list := []AmazonS3CorsConfigurationRule{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonS3CorsConfigurationRuleList(list)
		return nil
	}
	return err
}


type AmazonS3LifecycleConfiguration struct {
  Rules *AmazonS3LifecycleRule `json:"Rules,omitempty"`  // Amazon S3 Lifecycle Rule
}

type AmazonS3LifecycleConfigurationList []AmazonS3LifecycleConfiguration

func (l *AmazonS3LifecycleConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonS3LifecycleConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonS3LifecycleConfigurationList{item}
		return nil
	}
	list := []AmazonS3LifecycleConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonS3LifecycleConfigurationList(list)
		return nil
	}
	return err
}


type AmazonS3LifecycleRule struct {
  ExpirationDate *StringExpression `json:"ExpirationDate,omitempty"`  // String
  ExpirationInDays *Integer `json:"ExpirationInDays,omitempty"`  // Integer
  Id *StringExpression `json:"Id,omitempty"`  // String
  NoncurrentVersionExpirationInDays *Integer `json:"NoncurrentVersionExpirationInDays,omitempty"`  // Integer
  NoncurrentVersionTransition *AmazonS3LifecycleRuleNoncurrentVersionTransition `json:"NoncurrentVersionTransition,omitempty"`  // Amazon S3 Lifecycle Rule NoncurrentVersionTransition
  Prefix *StringExpression `json:"Prefix,omitempty"`  // String
  Status *StringExpression `json:"Status,omitempty"`  // String
  Transition *AmazonS3LifecycleRuleTransition `json:"Transition,omitempty"`  // Amazon S3 Lifecycle Rule Transition
}

type AmazonS3LifecycleRuleList []AmazonS3LifecycleRule

func (l *AmazonS3LifecycleRuleList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonS3LifecycleRule{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonS3LifecycleRuleList{item}
		return nil
	}
	list := []AmazonS3LifecycleRule{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonS3LifecycleRuleList(list)
		return nil
	}
	return err
}


type AmazonS3LifecycleRuleNoncurrentVersionTransition struct {
  StorageClass *StringExpression `json:"StorageClass,omitempty"`  // String
  TransitionInDays *Integer `json:"TransitionInDays,omitempty"`  // Integer
}

type AmazonS3LifecycleRuleNoncurrentVersionTransitionList []AmazonS3LifecycleRuleNoncurrentVersionTransition

func (l *AmazonS3LifecycleRuleNoncurrentVersionTransitionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonS3LifecycleRuleNoncurrentVersionTransition{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonS3LifecycleRuleNoncurrentVersionTransitionList{item}
		return nil
	}
	list := []AmazonS3LifecycleRuleNoncurrentVersionTransition{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonS3LifecycleRuleNoncurrentVersionTransitionList(list)
		return nil
	}
	return err
}


type AmazonS3LifecycleRuleTransition struct {
  StorageClass *StringExpression `json:"StorageClass,omitempty"`  // String
  TransitionDate *StringExpression `json:"TransitionDate,omitempty"`  // String
  TransitionInDays *Integer `json:"TransitionInDays,omitempty"`  // Integer
}

type AmazonS3LifecycleRuleTransitionList []AmazonS3LifecycleRuleTransition

func (l *AmazonS3LifecycleRuleTransitionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonS3LifecycleRuleTransition{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonS3LifecycleRuleTransitionList{item}
		return nil
	}
	list := []AmazonS3LifecycleRuleTransition{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonS3LifecycleRuleTransitionList(list)
		return nil
	}
	return err
}


type AmazonS3LoggingConfiguration struct {
  DestinationBucketName *StringExpression `json:"DestinationBucketName,omitempty"`  // String
  LogFilePrefix *StringExpression `json:"LogFilePrefix,omitempty"`  // String
}

type AmazonS3LoggingConfigurationList []AmazonS3LoggingConfiguration

func (l *AmazonS3LoggingConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonS3LoggingConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonS3LoggingConfigurationList{item}
		return nil
	}
	list := []AmazonS3LoggingConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonS3LoggingConfigurationList(list)
		return nil
	}
	return err
}


type AmazonS3NotificationConfiguration struct {
  TopicConfigurations *AmazonS3NotificationTopicConfigurations `json:"TopicConfigurations,omitempty"`  // Amazon S3 Notification Topic Configurations
}

type AmazonS3NotificationConfigurationList []AmazonS3NotificationConfiguration

func (l *AmazonS3NotificationConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonS3NotificationConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonS3NotificationConfigurationList{item}
		return nil
	}
	list := []AmazonS3NotificationConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonS3NotificationConfigurationList(list)
		return nil
	}
	return err
}


type AmazonS3NotificationTopicConfigurations struct {
  Event *StringExpression `json:"Event,omitempty"`  // String
  Topic *StringExpression `json:"Topic,omitempty"`  // String
}

type AmazonS3NotificationTopicConfigurationsList []AmazonS3NotificationTopicConfigurations

func (l *AmazonS3NotificationTopicConfigurationsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonS3NotificationTopicConfigurations{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonS3NotificationTopicConfigurationsList{item}
		return nil
	}
	list := []AmazonS3NotificationTopicConfigurations{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonS3NotificationTopicConfigurationsList(list)
		return nil
	}
	return err
}


type AmazonS3VersioningConfiguration struct {
  Status *StringExpression `json:"Status,omitempty"`  // String
}

type AmazonS3VersioningConfigurationList []AmazonS3VersioningConfiguration

func (l *AmazonS3VersioningConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonS3VersioningConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonS3VersioningConfigurationList{item}
		return nil
	}
	list := []AmazonS3VersioningConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonS3VersioningConfigurationList(list)
		return nil
	}
	return err
}


type AmazonS3WebsiteConfigurationProperty struct {
  ErrorDocument *StringExpression `json:"ErrorDocument,omitempty"`  // String
  IndexDocument *StringExpression `json:"IndexDocument,omitempty"`  // String
  RedirectAllRequestsTo *AmazonS3WebsiteConfigurationRedirectAllRequestsToProperty `json:"RedirectAllRequestsTo,omitempty"`  // Amazon S3 Website Configuration Redirect All Requests To Property
  RoutingRules *AmazonS3WebsiteConfigurationRoutingRulesProperty `json:"RoutingRules,omitempty"`  // Amazon S3 Website Configuration Routing Rules Property
}

type AmazonS3WebsiteConfigurationPropertyList []AmazonS3WebsiteConfigurationProperty

func (l *AmazonS3WebsiteConfigurationPropertyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonS3WebsiteConfigurationProperty{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonS3WebsiteConfigurationPropertyList{item}
		return nil
	}
	list := []AmazonS3WebsiteConfigurationProperty{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonS3WebsiteConfigurationPropertyList(list)
		return nil
	}
	return err
}


type AmazonS3WebsiteConfigurationRedirectAllRequestsToProperty struct {
  HostName *StringExpression `json:"HostName,omitempty"`  // String
  Protocol *StringExpression `json:"Protocol,omitempty"`  // String
}

type AmazonS3WebsiteConfigurationRedirectAllRequestsToPropertyList []AmazonS3WebsiteConfigurationRedirectAllRequestsToProperty

func (l *AmazonS3WebsiteConfigurationRedirectAllRequestsToPropertyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonS3WebsiteConfigurationRedirectAllRequestsToProperty{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonS3WebsiteConfigurationRedirectAllRequestsToPropertyList{item}
		return nil
	}
	list := []AmazonS3WebsiteConfigurationRedirectAllRequestsToProperty{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonS3WebsiteConfigurationRedirectAllRequestsToPropertyList(list)
		return nil
	}
	return err
}


type AmazonS3WebsiteConfigurationRoutingRulesProperty struct {
  RedirectRule *AmazonS3WebsiteConfigurationRoutingRulesRedirectRuleProperty `json:"RedirectRule,omitempty"`  // Amazon S3 Website Configuration Routing Rules Redirect Rule Property
  RoutingRuleCondition *AmazonS3WebsiteConfigurationRoutingRulesRoutingRuleConditionProperty `json:"RoutingRuleCondition,omitempty"`  // Amazon S3 Website Configuration Routing Rules Routing Rule Condition Property
}

type AmazonS3WebsiteConfigurationRoutingRulesPropertyList []AmazonS3WebsiteConfigurationRoutingRulesProperty

func (l *AmazonS3WebsiteConfigurationRoutingRulesPropertyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonS3WebsiteConfigurationRoutingRulesProperty{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonS3WebsiteConfigurationRoutingRulesPropertyList{item}
		return nil
	}
	list := []AmazonS3WebsiteConfigurationRoutingRulesProperty{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonS3WebsiteConfigurationRoutingRulesPropertyList(list)
		return nil
	}
	return err
}


type AmazonS3WebsiteConfigurationRoutingRulesRedirectRuleProperty struct {
  HostName *StringExpression `json:"HostName,omitempty"`  // String
  HttpRedirectCode *StringExpression `json:"HttpRedirectCode,omitempty"`  // String
  Protocol *StringExpression `json:"Protocol,omitempty"`  // String
  ReplaceKeyPrefixWith *StringExpression `json:"ReplaceKeyPrefixWith,omitempty"`  // String
  ReplaceKeyWith *StringExpression `json:"ReplaceKeyWith,omitempty"`  // String
}

type AmazonS3WebsiteConfigurationRoutingRulesRedirectRulePropertyList []AmazonS3WebsiteConfigurationRoutingRulesRedirectRuleProperty

func (l *AmazonS3WebsiteConfigurationRoutingRulesRedirectRulePropertyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonS3WebsiteConfigurationRoutingRulesRedirectRuleProperty{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonS3WebsiteConfigurationRoutingRulesRedirectRulePropertyList{item}
		return nil
	}
	list := []AmazonS3WebsiteConfigurationRoutingRulesRedirectRuleProperty{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonS3WebsiteConfigurationRoutingRulesRedirectRulePropertyList(list)
		return nil
	}
	return err
}


type AmazonS3WebsiteConfigurationRoutingRulesRoutingRuleConditionProperty struct {
  HttpErrorCodeReturnedEquals *StringExpression `json:"HttpErrorCodeReturnedEquals,omitempty"`  // String
  KeyPrefixEquals *StringExpression `json:"KeyPrefixEquals,omitempty"`  // String
}

type AmazonS3WebsiteConfigurationRoutingRulesRoutingRuleConditionPropertyList []AmazonS3WebsiteConfigurationRoutingRulesRoutingRuleConditionProperty

func (l *AmazonS3WebsiteConfigurationRoutingRulesRoutingRuleConditionPropertyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonS3WebsiteConfigurationRoutingRulesRoutingRuleConditionProperty{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonS3WebsiteConfigurationRoutingRulesRoutingRuleConditionPropertyList{item}
		return nil
	}
	list := []AmazonS3WebsiteConfigurationRoutingRulesRoutingRuleConditionProperty{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonS3WebsiteConfigurationRoutingRulesRoutingRuleConditionPropertyList(list)
		return nil
	}
	return err
}


type AmazonSNSSubscription struct {
  Endpoint *StringExpression `json:"Endpoint,omitempty"`  // String
  Protocol *StringExpression `json:"Protocol,omitempty"`  // String
}

type AmazonSNSSubscriptionList []AmazonSNSSubscription

func (l *AmazonSNSSubscriptionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonSNSSubscription{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonSNSSubscriptionList{item}
		return nil
	}
	list := []AmazonSNSSubscription{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonSNSSubscriptionList(list)
		return nil
	}
	return err
}


type AmazonSQSRedrivePolicy struct {
  DeadLetterTargetArn *StringExpression `json:"deadLetterTargetArn,omitempty"`  // String
  MaxReceiveCount *Integer `json:"maxReceiveCount,omitempty"`  // Integer
}

type AmazonSQSRedrivePolicyList []AmazonSQSRedrivePolicy

func (l *AmazonSQSRedrivePolicyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AmazonSQSRedrivePolicy{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AmazonSQSRedrivePolicyList{item}
		return nil
	}
	list := []AmazonSQSRedrivePolicy{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AmazonSQSRedrivePolicyList(list)
		return nil
	}
	return err
}

func NewResourceByType(typeName string) interface{} {
	switch typeName {
		case "AWS::AutoScaling::AutoScalingGroup":
			return &AWSAutoScalingAutoScalingGroup{}
		case "AWS::AutoScaling::LaunchConfiguration":
			return &AWSAutoScalingLaunchConfiguration{}
		case "AWS::AutoScaling::LifecycleHook":
			return &AWSAutoScalingLifecycleHook{}
		case "AWS::AutoScaling::ScalingPolicy":
			return &AWSAutoScalingScalingPolicy{}
		case "AWS::AutoScaling::ScheduledAction":
			return &AWSAutoScalingScheduledAction{}
		case "AWS::CloudFormation::Authentication":
			return &AWSCloudFormationAuthentication{}
		case "AWS::CloudFormation::CustomResource":
			return &AWSCloudFormationCustomResource{}
		case "AWS::CloudFormation::Init":
			return &AWSCloudFormationInit{}
		case "AWS::CloudFormation::Stack":
			return &AWSCloudFormationStack{}
		case "AWS::CloudFormation::WaitCondition":
			return &AWSCloudFormationWaitCondition{}
		case "AWS::CloudFormation::WaitConditionHandle":
			return &AWSCloudFormationWaitConditionHandle{}
		case "AWS::CloudFront::Distribution":
			return &AWSCloudFrontDistribution{}
		case "AWS::CloudTrail::Trail":
			return &AWSCloudTrailTrail{}
		case "AWS::CloudWatch::Alarm":
			return &AWSCloudWatchAlarm{}
		case "AWS::DataPipeline::Pipeline":
			return &AWSDataPipelinePipeline{}
		case "AWS::DynamoDB::Table":
			return &AWSDynamoDBTable{}
		case "AWS::EC2::CustomerGateway":
			return &AWSEC2CustomerGateway{}
		case "AWS::EC2::DHCPOptions":
			return &AWSEC2DHCPOptions{}
		case "AWS::EC2::EIP":
			return &AWSEC2EIP{}
		case "AWS::EC2::EIPAssociation":
			return &AWSEC2EIPAssociation{}
		case "AWS::EC2::Instance":
			return &AWSEC2Instance{}
		case "AWS::EC2::InternetGateway":
			return &AWSEC2InternetGateway{}
		case "AWS::EC2::NetworkAcl":
			return &AWSEC2NetworkAcl{}
		case "AWS::EC2::NetworkAclEntry":
			return &AWSEC2NetworkAclEntry{}
		case "AWS::EC2::NetworkInterface":
			return &AWSEC2NetworkInterface{}
		case "AWS::EC2::NetworkInterfaceAttachment":
			return &AWSEC2NetworkInterfaceAttachment{}
		case "AWS::EC2::Route":
			return &AWSEC2Route{}
		case "AWS::EC2::RouteTable":
			return &AWSEC2RouteTable{}
		case "AWS::EC2::SecurityGroup":
			return &AWSEC2SecurityGroup{}
		case "AWS::EC2::SecurityGroupEgress":
			return &AWSEC2SecurityGroupEgress{}
		case "AWS::EC2::SecurityGroupIngress":
			return &AWSEC2SecurityGroupIngress{}
		case "AWS::EC2::Subnet":
			return &AWSEC2Subnet{}
		case "AWS::EC2::SubnetNetworkAclAssociation":
			return &AWSEC2SubnetNetworkAclAssociation{}
		case "AWS::EC2::SubnetRouteTableAssociation":
			return &AWSEC2SubnetRouteTableAssociation{}
		case "AWS::EC2::Volume":
			return &AWSEC2Volume{}
		case "AWS::EC2::VolumeAttachment":
			return &AWSEC2VolumeAttachment{}
		case "AWS::EC2::VPC":
			return &AWSEC2VPC{}
		case "AWS::EC2::VPCDHCPOptionsAssociation":
			return &AWSEC2VPCDHCPOptionsAssociation{}
		case "AWS::EC2::VPCGatewayAttachment":
			return &AWSEC2VPCGatewayAttachment{}
		case "AWS::EC2::VPCPeeringConnection":
			return &AWSEC2VPCPeeringConnection{}
		case "AWS::EC2::VPNConnection":
			return &AWSEC2VPNConnection{}
		case "AWS::EC2::VPNConnectionRoute":
			return &AWSEC2VPNConnectionRoute{}
		case "AWS::EC2::VPNGateway":
			return &AWSEC2VPNGateway{}
		case "AWS::EC2::VPNGatewayRoutePropagation":
			return &AWSEC2VPNGatewayRoutePropagation{}
		case "AWS::ECS::Cluster":
			return &AWSECSCluster{}
		case "AWS::ECS::Service":
			return &AWSECSService{}
		case "AWS::ECS::TaskDefinition":
			return &AWSECSTaskDefinition{}
		case "AWS::ElastiCache::CacheCluster":
			return &AWSElastiCacheCacheCluster{}
		case "AWS::ElastiCache::ParameterGroup":
			return &AWSElastiCacheParameterGroup{}
		case "AWS::ElastiCache::ReplicationGroup":
			return &AWSElastiCacheReplicationGroup{}
		case "AWS::ElastiCache::SecurityGroup":
			return &AWSElastiCacheSecurityGroup{}
		case "AWS::ElastiCache::SecurityGroupIngress":
			return &AWSElastiCacheSecurityGroupIngress{}
		case "AWS::ElastiCache::SubnetGroup ":
			return &AWSElastiCacheSubnetGroup{}
		case "AWS::ElasticBeanstalk::Application":
			return &AWSElasticBeanstalkApplication{}
		case "AWS::ElasticBeanstalk::ApplicationVersion":
			return &AWSElasticBeanstalkApplicationVersion{}
		case "AWS::ElasticBeanstalk::ConfigurationTemplate":
			return &AWSElasticBeanstalkConfigurationTemplate{}
		case "AWS::ElasticBeanstalk::Environment":
			return &AWSElasticBeanstalkEnvironment{}
		case "AWS::ElasticLoadBalancing::LoadBalancer":
			return &AWSElasticLoadBalancingLoadBalancer{}
		case "AWS::IAM::AccessKey":
			return &AWSIAMAccessKey{}
		case "AWS::IAM::Group":
			return &AWSIAMGroup{}
		case "AWS::IAM::InstanceProfile":
			return &AWSIAMInstanceProfile{}
		case "AWS::IAM::ManagedPolicy":
			return &AWSIAMManagedPolicy{}
		case "AWS::IAM::Policy":
			return &AWSIAMPolicy{}
		case "AWS::IAM::Role":
			return &AWSIAMRole{}
		case "AWS::IAM::User":
			return &AWSIAMUser{}
		case "AWS::IAM::UserToGroupAddition":
			return &AWSIAMUserToGroupAddition{}
		case "AWS::Kinesis::Stream":
			return &AWSKinesisStream{}
		case "AWS::Lambda::Function":
			return &AWSLambdaFunction{}
		case "AWS::Logs::LogGroup":
			return &AWSLogsLogGroup{}
		case "AWS::Logs::MetricFilter":
			return &AWSLogsMetricFilter{}
		case "AWS::OpsWorks::App":
			return &AWSOpsWorksApp{}
		case "AWS::OpsWorks::ElasticLoadBalancerAttachment":
			return &AWSOpsWorksElasticLoadBalancerAttachment{}
		case "AWS::OpsWorks::Instance":
			return &AWSOpsWorksInstance{}
		case "AWS::OpsWorks::Layer":
			return &AWSOpsWorksLayer{}
		case "AWS::OpsWorks::Stack":
			return &AWSOpsWorksStack{}
		case "AWS::Redshift::Cluster":
			return &AWSRedshiftCluster{}
		case "AWS::Redshift::ClusterParameterGroup":
			return &AWSRedshiftClusterParameterGroup{}
		case "AWS::Redshift::ClusterSecurityGroup":
			return &AWSRedshiftClusterSecurityGroup{}
		case "AWS::Redshift::ClusterSecurityGroupIngress":
			return &AWSRedshiftClusterSecurityGroupIngress{}
		case "AWS::Redshift::ClusterSubnetGroup":
			return &AWSRedshiftClusterSubnetGroup{}
		case "AWS::RDS::DBInstance":
			return &AWSRDSDBInstance{}
		case "AWS::RDS::DBParameterGroup":
			return &AWSRDSDBParameterGroup{}
		case "AWS::RDS::DBSubnetGroup":
			return &AWSRDSDBSubnetGroup{}
		case "AWS::RDS::DBSecurityGroup":
			return &AWSRDSDBSecurityGroup{}
		case "AWS::RDS::DBSecurityGroupIngress":
			return &AWSRDSDBSecurityGroupIngress{}
		case "AWS::RDS::EventSubscription":
			return &AWSRDSEventSubscription{}
		case "AWS::RDS::OptionGroup":
			return &AWSRDSOptionGroup{}
		case "AWS::Route53::HealthCheck":
			return &AWSRoute53HealthCheck{}
		case "AWS::Route53::HostedZone":
			return &AWSRoute53HostedZone{}
		case "AWS::Route53::RecordSet":
			return &AWSRoute53RecordSet{}
		case "AWS::Route53::RecordSetGroup":
			return &AWSRoute53RecordSetGroup{}
		case "AWS::S3::Bucket":
			return &AWSS3Bucket{}
		case "AWS::S3::BucketPolicy":
			return &AWSS3BucketPolicy{}
		case "AWS::SDB::Domain":
			return &AWSSDBDomain{}
		case "AWS::SNS::Topic":
			return &AWSSNSTopic{}
		case "AWS::SNS::TopicPolicy":
			return &AWSSNSTopicPolicy{}
		case "AWS::SQS::Queue":
			return &AWSSQSQueue{}
		case "AWS::SQS::QueuePolicy":
			return &AWSSQSQueuePolicy{}
	}
	return nil
}
