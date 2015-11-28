package cloudformation

import "time"
import "encoding/json"

type AWSAutoScalingAutoScalingGroup struct {
  AvailabilityZones *StringListExpr `json:"AvailabilityZones,omitempty"`  // List of strings
  Cooldown *StringExpr `json:"Cooldown,omitempty"`  // String
  DesiredCapacity *StringExpr `json:"DesiredCapacity,omitempty"`  // String
  HealthCheckGracePeriod *IntegerExpr `json:"HealthCheckGracePeriod,omitempty"`  // Integer
  HealthCheckType *StringExpr `json:"HealthCheckType,omitempty"`  // String
  InstanceId *StringExpr `json:"InstanceId,omitempty"`  // String
  LaunchConfigurationName *StringExpr `json:"LaunchConfigurationName,omitempty"`  // String
  LoadBalancerNames *StringListExpr `json:"LoadBalancerNames,omitempty"`  // List of strings
  MaxSize *StringExpr `json:"MaxSize,omitempty"`  // String
  MetricsCollection *AutoScalingMetricsCollectionList `json:"MetricsCollection,omitempty"`  // A list of Auto Scaling MetricsCollection
  MinSize *StringExpr `json:"MinSize,omitempty"`  // String
  NotificationConfigurations *AutoScalingNotificationConfigurationsList `json:"NotificationConfigurations,omitempty"`  // List of Auto Scaling NotificationConfigurations
  PlacementGroup *StringExpr `json:"PlacementGroup,omitempty"`  // String
  Tags *AutoScalingTagsList `json:"Tags,omitempty"`  // List of Auto Scaling Tags
  TerminationPolicies *StringListExpr `json:"TerminationPolicies,omitempty"`  // List of strings
  VPCZoneIdentifier *StringListExpr `json:"VPCZoneIdentifier,omitempty"`  // List of strings
}

func (s AWSAutoScalingAutoScalingGroup) ResourceType() string {
	return "AWS::AutoScaling::AutoScalingGroup"
}


type AWSAutoScalingLaunchConfiguration struct {
  AssociatePublicIpAddress *BoolExpr `json:"AssociatePublicIpAddress,omitempty"`  // Boolean
  BlockDeviceMappings *AWSCloudFormationAutoScalingBlockDeviceMappingList `json:"BlockDeviceMappings,omitempty"`  // A list of BlockDeviceMappings
  ClassicLinkVPCId *StringExpr `json:"ClassicLinkVPCId,omitempty"`  // String
  ClassicLinkVPCSecurityGroups *StringListExpr `json:"ClassicLinkVPCSecurityGroups,omitempty"`  // List of strings
  EbsOptimized *BoolExpr `json:"EbsOptimized,omitempty"`  // Boolean
  IamInstanceProfile *StringExpr `json:"IamInstanceProfile,omitempty"`  // String (1–1600 chars)
  ImageId *StringExpr `json:"ImageId,omitempty"`  // String
  InstanceId *StringExpr `json:"InstanceId,omitempty"`  // String
  InstanceMonitoring *BoolExpr `json:"InstanceMonitoring,omitempty"`  // Boolean
  InstanceType *StringExpr `json:"InstanceType,omitempty"`  // String
  KernelId *StringExpr `json:"KernelId,omitempty"`  // String
  KeyName *StringExpr `json:"KeyName,omitempty"`  // String
  PlacementTenancy *StringExpr `json:"PlacementTenancy,omitempty"`  // String
  RamDiskId *StringExpr `json:"RamDiskId,omitempty"`  // String
  SecurityGroups interface{} `json:"SecurityGroups,omitempty"`  // A list of EC2 security groups
  SpotPrice *StringExpr `json:"SpotPrice,omitempty"`  // String
  UserData *StringExpr `json:"UserData,omitempty"`  // String
}

func (s AWSAutoScalingLaunchConfiguration) ResourceType() string {
	return "AWS::AutoScaling::LaunchConfiguration"
}


type AWSAutoScalingLifecycleHook struct {
  AutoScalingGroupName *StringExpr `json:"AutoScalingGroupName,omitempty"`  // String
  DefaultResult *StringExpr `json:"DefaultResult,omitempty"`  // String
  HeartbeatTimeout *IntegerExpr `json:"HeartbeatTimeout,omitempty"`  // Integer
  LifecycleTransition *StringExpr `json:"LifecycleTransition,omitempty"`  // String
  NotificationMetadata *StringExpr `json:"NotificationMetadata,omitempty"`  // String
  NotificationTargetARN *StringExpr `json:"NotificationTargetARN,omitempty"`  // String
  RoleARN *StringExpr `json:"RoleARN,omitempty"`  // String
}

func (s AWSAutoScalingLifecycleHook) ResourceType() string {
	return "AWS::AutoScaling::LifecycleHook"
}


type AWSAutoScalingScalingPolicy struct {
  AdjustmentType *StringExpr `json:"AdjustmentType,omitempty"`  // String
  AutoScalingGroupName *StringExpr `json:"AutoScalingGroupName,omitempty"`  // String
  Cooldown *StringExpr `json:"Cooldown,omitempty"`  // String
  MinAdjustmentStep *IntegerExpr `json:"MinAdjustmentStep,omitempty"`  // Integer
  ScalingAdjustment *StringExpr `json:"ScalingAdjustment,omitempty"`  // String
}

func (s AWSAutoScalingScalingPolicy) ResourceType() string {
	return "AWS::AutoScaling::ScalingPolicy"
}


type AWSAutoScalingScheduledAction struct {
  AutoScalingGroupName *StringExpr `json:"AutoScalingGroupName,omitempty"`  // String
  DesiredCapacity *IntegerExpr `json:"DesiredCapacity,omitempty"`  // Integer
  EndTime time.Time `json:"EndTime,omitempty"`  // Time stamp
  MaxSize *IntegerExpr `json:"MaxSize,omitempty"`  // Integer
  MinSize *IntegerExpr `json:"MinSize,omitempty"`  // Integer
  Recurrence *StringExpr `json:"Recurrence,omitempty"`  // String
  StartTime time.Time `json:"StartTime,omitempty"`  // Time stamp
}

func (s AWSAutoScalingScheduledAction) ResourceType() string {
	return "AWS::AutoScaling::ScheduledAction"
}


type AWSCloudFormationAuthentication struct {
  AccessKeyId *StringExpr `json:"accessKeyId,omitempty"`  // String
  Buckets *StringListExpr `json:"buckets,omitempty"`  // List of strings
  Password *StringExpr `json:"password,omitempty"`  // String
  SecretKey *StringExpr `json:"secretKey,omitempty"`  // String
  Type *StringExpr `json:"type,omitempty"`  // String Valid values are "basic" or "S3"
  Uris *StringListExpr `json:"uris,omitempty"`  // List of strings
  Username *StringExpr `json:"username,omitempty"`  // String
  RoleName *StringExpr `json:"roleName,omitempty"`  // String
}

func (s AWSCloudFormationAuthentication) ResourceType() string {
	return "AWS::CloudFormation::Authentication"
}


type AWSCloudFormationCustomResource struct {
  ServiceToken *StringExpr `json:"ServiceToken,omitempty"`  // String
}

func (s AWSCloudFormationCustomResource) ResourceType() string {
	return "AWS::CloudFormation::CustomResource"
}


type AWSCloudFormationInit struct {
}

func (s AWSCloudFormationInit) ResourceType() string {
	return "AWS::CloudFormation::Init"
}


type AWSCloudFormationStack struct {
  NotificationARNs *StringListExpr `json:"NotificationARNs,omitempty"`  // List of strings
  Parameters *CloudFormationStackParameters `json:"Parameters,omitempty"`  // CloudFormation Stack Parameters Property Type
  TemplateURL *StringExpr `json:"TemplateURL,omitempty"`  // String
  TimeoutInMinutes *StringExpr `json:"TimeoutInMinutes,omitempty"`  // String
}

func (s AWSCloudFormationStack) ResourceType() string {
	return "AWS::CloudFormation::Stack"
}


type AWSCloudFormationWaitCondition struct {
  Count *StringExpr `json:"Count,omitempty"`  // String
  Handle *StringExpr `json:"Handle,omitempty"`  // String
  Timeout *StringExpr `json:"Timeout,omitempty"`  // String
}

func (s AWSCloudFormationWaitCondition) ResourceType() string {
	return "AWS::CloudFormation::WaitCondition"
}


type AWSCloudFormationWaitConditionHandle struct {
}

func (s AWSCloudFormationWaitConditionHandle) ResourceType() string {
	return "AWS::CloudFormation::WaitConditionHandle"
}


type AWSCloudFrontDistribution struct {
  DistributionConfig *CloudFrontDistributionConfig `json:"DistributionConfig,omitempty"`  // DistributionConfig type
}

func (s AWSCloudFrontDistribution) ResourceType() string {
	return "AWS::CloudFront::Distribution"
}


type AWSCloudTrailTrail struct {
  IncludeGlobalServiceEvents *BoolExpr `json:"IncludeGlobalServiceEvents,omitempty"`  // Boolean
  IsLogging *BoolExpr `json:"IsLogging,omitempty"`  // Boolean
  S3BucketName *StringExpr `json:"S3BucketName,omitempty"`  // String
  S3KeyPrefix *StringExpr `json:"S3KeyPrefix,omitempty"`  // String
  SnsTopicName *StringExpr `json:"SnsTopicName,omitempty"`  // String
}

func (s AWSCloudTrailTrail) ResourceType() string {
	return "AWS::CloudTrail::Trail"
}


type AWSCloudWatchAlarm struct {
  ActionsEnabled *BoolExpr `json:"ActionsEnabled,omitempty"`  // Boolean
  AlarmActions *StringListExpr `json:"AlarmActions,omitempty"`  // List of strings
  AlarmDescription *StringExpr `json:"AlarmDescription,omitempty"`  // String
  AlarmName *StringExpr `json:"AlarmName,omitempty"`  // String
  ComparisonOperator *StringExpr `json:"ComparisonOperator,omitempty"`  // String
  Dimensions *CloudWatchMetricDimensionList `json:"Dimensions,omitempty"`  // List of Metric Dimension
  EvaluationPeriods *StringExpr `json:"EvaluationPeriods,omitempty"`  // String
  InsufficientDataActions *StringListExpr `json:"InsufficientDataActions,omitempty"`  // List of strings
  MetricName *StringExpr `json:"MetricName,omitempty"`  // String
  Namespace *StringExpr `json:"Namespace,omitempty"`  // String
  OKActions *StringListExpr `json:"OKActions,omitempty"`  // List of strings
  Period *StringExpr `json:"Period,omitempty"`  // String
  Statistic *StringExpr `json:"Statistic,omitempty"`  // String
  Threshold *StringExpr `json:"Threshold,omitempty"`  // String
  Unit *StringExpr `json:"Unit,omitempty"`  // String
}

func (s AWSCloudWatchAlarm) ResourceType() string {
	return "AWS::CloudWatch::Alarm"
}


type AWSDataPipelinePipeline struct {
  Activate *BoolExpr `json:"Activate,omitempty"`  // Boolean
  Description *StringExpr `json:"Description,omitempty"`  // String
  Name *StringExpr `json:"Name,omitempty"`  // String
  ParameterObjects *AWSDataPipelinePipelineParameterObjects `json:"ParameterObjects,omitempty"`  // AWS Data Pipeline Pipeline ParameterObjects
  ParameterValues *AWSDataPipelinePipelineParameterValues `json:"ParameterValues,omitempty"`  // AWS Data Pipeline Pipeline ParameterValues
  PipelineObjects *AWSDataPipelinePipelineObjectsList `json:"PipelineObjects,omitempty"`  // A list of AWS Data Pipeline PipelineObjects
  PipelineTags *AWSDataPipelinePipelinePipelineTagsList `json:"PipelineTags,omitempty"`  // AWS Data Pipeline Pipeline PipelineTags
}

func (s AWSDataPipelinePipeline) ResourceType() string {
	return "AWS::DataPipeline::Pipeline"
}


type AWSDynamoDBTable struct {
  AttributeDefinitions *DynamoDBAttributeDefinitionsList `json:"AttributeDefinitions,omitempty"`  // DynamoDB Attribute Definitions
  GlobalSecondaryIndexes *DynamoDBGlobalSecondaryIndexes `json:"GlobalSecondaryIndexes,omitempty"`  // DynamoDB Global Secondary Indexes
  KeySchema *DynamoDBKeySchema `json:"KeySchema,omitempty"`  // DynamoDB Key Schema
  LocalSecondaryIndexes *DynamoDBLocalSecondaryIndexes `json:"LocalSecondaryIndexes,omitempty"`  // DynamoDB Local Secondary Indexes
  ProvisionedThroughput *DynamoDBProvisionedThroughput `json:"ProvisionedThroughput,omitempty"`  // DynamoDB Provisioned Throughput
  TableName *StringExpr `json:"TableName,omitempty"`  // String
}

func (s AWSDynamoDBTable) ResourceType() string {
	return "AWS::DynamoDB::Table"
}


type AWSEC2CustomerGateway struct {
  BgpAsn *IntegerExpr `json:"BgpAsn,omitempty"`  // Number BgpAsn is always an integer value
  IpAddress *StringExpr `json:"IpAddress,omitempty"`  // String
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
  Type *StringExpr `json:"Type,omitempty"`  // String
}

func (s AWSEC2CustomerGateway) ResourceType() string {
	return "AWS::EC2::CustomerGateway"
}


type AWSEC2DHCPOptions struct {
  DomainName *StringExpr `json:"DomainName,omitempty"`  // String
  DomainNameServers *StringListExpr `json:"DomainNameServers,omitempty"`  // List of strings
  NetbiosNameServers *StringListExpr `json:"NetbiosNameServers,omitempty"`  // List of strings
  NetbiosNodeType interface{} `json:"NetbiosNodeType,omitempty"`  // List of numbers
  NtpServers *StringListExpr `json:"NtpServers,omitempty"`  // List of strings
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
}

func (s AWSEC2DHCPOptions) ResourceType() string {
	return "AWS::EC2::DHCPOptions"
}


type AWSEC2EIP struct {
  InstanceId *StringExpr `json:"InstanceId,omitempty"`  // String
  Domain *StringExpr `json:"Domain,omitempty"`  // String
}

func (s AWSEC2EIP) ResourceType() string {
	return "AWS::EC2::EIP"
}


type AWSEC2EIPAssociation struct {
  AllocationId *StringExpr `json:"AllocationId,omitempty"`  // String
  EIP *StringExpr `json:"EIP,omitempty"`  // String
  InstanceId *StringExpr `json:"InstanceId,omitempty"`  // String
  NetworkInterfaceId *StringExpr `json:"NetworkInterfaceId,omitempty"`  // String
  PrivateIpAddress *StringExpr `json:"PrivateIpAddress,omitempty"`  // String
}

func (s AWSEC2EIPAssociation) ResourceType() string {
	return "AWS::EC2::EIPAssociation"
}


type AWSEC2Instance struct {
  AvailabilityZone *StringExpr `json:"AvailabilityZone,omitempty"`  // String
  BlockDeviceMappings *AmazonEC2BlockDeviceMappingPropertyList `json:"BlockDeviceMappings,omitempty"`  // A list of Amazon EC2 Block Device Mapping Property
  DisableApiTermination *BoolExpr `json:"DisableApiTermination,omitempty"`  // Boolean
  EbsOptimized *BoolExpr `json:"EbsOptimized,omitempty"`  // Boolean
  IamInstanceProfile *StringExpr `json:"IamInstanceProfile,omitempty"`  // String
  ImageId *StringExpr `json:"ImageId,omitempty"`  // String
  InstanceInitiatedShutdownBehavior *StringExpr `json:"InstanceInitiatedShutdownBehavior,omitempty"`  // String
  InstanceType *StringExpr `json:"InstanceType,omitempty"`  // String
  KernelId *StringExpr `json:"KernelId,omitempty"`  // String
  KeyName *StringExpr `json:"KeyName,omitempty"`  // String
  Monitoring *BoolExpr `json:"Monitoring,omitempty"`  // Boolean
  NetworkInterfaces *EC2NetworkInterfaceEmbeddedList `json:"NetworkInterfaces,omitempty"`  // A list of EC2 NetworkInterface Embedded Property Type
  PlacementGroupName *StringExpr `json:"PlacementGroupName,omitempty"`  // String
  PrivateIpAddress *StringExpr `json:"PrivateIpAddress,omitempty"`  // String
  RamdiskId *StringExpr `json:"RamdiskId,omitempty"`  // String
  SecurityGroupIds *StringListExpr `json:"SecurityGroupIds,omitempty"`  // List of strings
  SecurityGroups *StringListExpr `json:"SecurityGroups,omitempty"`  // List of strings
  SourceDestCheck *BoolExpr `json:"SourceDestCheck,omitempty"`  // Boolean
  SubnetId *StringExpr `json:"SubnetId,omitempty"`  // String
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
  Tenancy *StringExpr `json:"Tenancy,omitempty"`  // String
  UserData *StringExpr `json:"UserData,omitempty"`  // String
  Volumes *EC2MountPointList `json:"Volumes,omitempty"`  // A list of EC2 MountPoints
}

func (s AWSEC2Instance) ResourceType() string {
	return "AWS::EC2::Instance"
}


type AWSEC2InternetGateway struct {
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
}

func (s AWSEC2InternetGateway) ResourceType() string {
	return "AWS::EC2::InternetGateway"
}


type AWSEC2NetworkAcl struct {
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
  VpcId *StringExpr `json:"VpcId,omitempty"`  // String
}

func (s AWSEC2NetworkAcl) ResourceType() string {
	return "AWS::EC2::NetworkAcl"
}


type AWSEC2NetworkAclEntry struct {
  CidrBlock *StringExpr `json:"CidrBlock,omitempty"`  // String
  Egress *BoolExpr `json:"Egress,omitempty"`  // Boolean
  Icmp *EC2ICMP `json:"Icmp,omitempty"`  // EC2 ICMP Property Type
  NetworkAclId *StringExpr `json:"NetworkAclId,omitempty"`  // String
  PortRange *EC2PortRange `json:"PortRange,omitempty"`  // EC2 PortRange Property Type
  Protocol *IntegerExpr `json:"Protocol,omitempty"`  // Number
  RuleAction *StringExpr `json:"RuleAction,omitempty"`  // String
  RuleNumber *IntegerExpr `json:"RuleNumber,omitempty"`  // Number
}

func (s AWSEC2NetworkAclEntry) ResourceType() string {
	return "AWS::EC2::NetworkAclEntry"
}


type AWSEC2NetworkInterface struct {
  Description *StringExpr `json:"Description,omitempty"`  // String
  GroupSet *StringListExpr `json:"GroupSet,omitempty"`  // List of strings
  PrivateIpAddress *StringExpr `json:"PrivateIpAddress,omitempty"`  // String
  PrivateIpAddresses *EC2NetworkInterfacePrivateIPSpecificationList `json:"PrivateIpAddresses,omitempty"`  // list of PrivateIpAddressSpecification
  SecondaryPrivateIpAddressCount *IntegerExpr `json:"SecondaryPrivateIpAddressCount,omitempty"`  // Integer
  SourceDestCheck *BoolExpr `json:"SourceDestCheck,omitempty"`  // Boolean
  SubnetId *StringExpr `json:"SubnetId,omitempty"`  // String
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
}

func (s AWSEC2NetworkInterface) ResourceType() string {
	return "AWS::EC2::NetworkInterface"
}


type AWSEC2NetworkInterfaceAttachment struct {
  DeleteOnTermination *BoolExpr `json:"DeleteOnTermination,omitempty"`  // Boolean
  DeviceIndex *StringExpr `json:"DeviceIndex,omitempty"`  // String
  InstanceId *StringExpr `json:"InstanceId,omitempty"`  // String
  NetworkInterfaceId *StringExpr `json:"NetworkInterfaceId,omitempty"`  // String
}

func (s AWSEC2NetworkInterfaceAttachment) ResourceType() string {
	return "AWS::EC2::NetworkInterfaceAttachment"
}


type AWSEC2Route struct {
  DestinationCidrBlock *StringExpr `json:"DestinationCidrBlock,omitempty"`  // String
  GatewayId *StringExpr `json:"GatewayId,omitempty"`  // String
  InstanceId *StringExpr `json:"InstanceId,omitempty"`  // String
  NetworkInterfaceId *StringExpr `json:"NetworkInterfaceId,omitempty"`  // String
  RouteTableId *StringExpr `json:"RouteTableId,omitempty"`  // String
  VpcPeeringConnectionId *StringExpr `json:"VpcPeeringConnectionId,omitempty"`  // String
}

func (s AWSEC2Route) ResourceType() string {
	return "AWS::EC2::Route"
}


type AWSEC2RouteTable struct {
  VpcId *StringExpr `json:"VpcId,omitempty"`  // String
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
}

func (s AWSEC2RouteTable) ResourceType() string {
	return "AWS::EC2::RouteTable"
}


type AWSEC2SecurityGroup struct {
  GroupDescription *StringExpr `json:"GroupDescription,omitempty"`  // String
  SecurityGroupEgress *EC2SecurityGroupRuleList `json:"SecurityGroupEgress,omitempty"`  // EC2 Security Group Rule
  SecurityGroupIngress *EC2SecurityGroupRuleList `json:"SecurityGroupIngress,omitempty"`  // EC2 Security Group Rule
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
  VpcId *StringExpr `json:"VpcId,omitempty"`  // String
}

func (s AWSEC2SecurityGroup) ResourceType() string {
	return "AWS::EC2::SecurityGroup"
}


type AWSEC2SecurityGroupEgress struct {
  CidrIp *StringExpr `json:"CidrIp,omitempty"`  // String
  DestinationSecurityGroupId *StringExpr `json:"DestinationSecurityGroupId,omitempty"`  // String
  FromPort *IntegerExpr `json:"FromPort,omitempty"`  // Integer
  GroupId *StringExpr `json:"GroupId,omitempty"`  // String
  IpProtocol *StringExpr `json:"IpProtocol,omitempty"`  // String
  ToPort *IntegerExpr `json:"ToPort,omitempty"`  // Integer
}

func (s AWSEC2SecurityGroupEgress) ResourceType() string {
	return "AWS::EC2::SecurityGroupEgress"
}


type AWSEC2SecurityGroupIngress struct {
  CidrIp *StringExpr `json:"CidrIp,omitempty"`  // String
  FromPort *IntegerExpr `json:"FromPort,omitempty"`  // Integer
  GroupId *StringExpr `json:"GroupId,omitempty"`  // String
  GroupName *StringExpr `json:"GroupName,omitempty"`  // String
  IpProtocol *StringExpr `json:"IpProtocol,omitempty"`  // String
  SourceSecurityGroupId *StringExpr `json:"SourceSecurityGroupId,omitempty"`  // String
  SourceSecurityGroupName *StringExpr `json:"SourceSecurityGroupName,omitempty"`  // String
  SourceSecurityGroupOwnerId *StringExpr `json:"SourceSecurityGroupOwnerId,omitempty"`  // String
  ToPort *IntegerExpr `json:"ToPort,omitempty"`  // Integer
}

func (s AWSEC2SecurityGroupIngress) ResourceType() string {
	return "AWS::EC2::SecurityGroupIngress"
}


type AWSEC2Subnet struct {
  AvailabilityZone *StringExpr `json:"AvailabilityZone,omitempty"`  // String
  CidrBlock *StringExpr `json:"CidrBlock,omitempty"`  // String
  MapPublicIpOnLaunch *BoolExpr `json:"MapPublicIpOnLaunch,omitempty"`  // Boolean
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
  VpcId interface{} `json:"VpcId,omitempty"`  // Ref ID
}

func (s AWSEC2Subnet) ResourceType() string {
	return "AWS::EC2::Subnet"
}


type AWSEC2SubnetNetworkAclAssociation struct {
  SubnetId *StringExpr `json:"SubnetId,omitempty"`  // String
  NetworkAclId *StringExpr `json:"NetworkAclId,omitempty"`  // String
}

func (s AWSEC2SubnetNetworkAclAssociation) ResourceType() string {
	return "AWS::EC2::SubnetNetworkAclAssociation"
}


type AWSEC2SubnetRouteTableAssociation struct {
  RouteTableId *StringExpr `json:"RouteTableId,omitempty"`  // String
  SubnetId *StringExpr `json:"SubnetId,omitempty"`  // String
}

func (s AWSEC2SubnetRouteTableAssociation) ResourceType() string {
	return "AWS::EC2::SubnetRouteTableAssociation"
}


type AWSEC2Volume struct {
  AvailabilityZone *StringExpr `json:"AvailabilityZone,omitempty"`  // String
  Encrypted *BoolExpr `json:"Encrypted,omitempty"`  // Boolean
  Iops *IntegerExpr `json:"Iops,omitempty"`  // Number
  KmsKeyId *StringExpr `json:"KmsKeyId,omitempty"`  // String
  Size *StringExpr `json:"Size,omitempty"`  // String
  SnapshotId *StringExpr `json:"SnapshotId,omitempty"`  // String
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
  VolumeType *StringExpr `json:"VolumeType,omitempty"`  // String
}

func (s AWSEC2Volume) ResourceType() string {
	return "AWS::EC2::Volume"
}


type AWSEC2VolumeAttachment struct {
  Device *StringExpr `json:"Device,omitempty"`  // String
  InstanceId *StringExpr `json:"InstanceId,omitempty"`  // String
  VolumeId *StringExpr `json:"VolumeId,omitempty"`  // String
}

func (s AWSEC2VolumeAttachment) ResourceType() string {
	return "AWS::EC2::VolumeAttachment"
}


type AWSEC2VPC struct {
  CidrBlock *StringExpr `json:"CidrBlock,omitempty"`  // String
  EnableDnsSupport *BoolExpr `json:"EnableDnsSupport,omitempty"`  // Boolean
  EnableDnsHostnames *BoolExpr `json:"EnableDnsHostnames,omitempty"`  // Boolean
  InstanceTenancy *StringExpr `json:"InstanceTenancy,omitempty"`  // String
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
}

func (s AWSEC2VPC) ResourceType() string {
	return "AWS::EC2::VPC"
}


type AWSEC2VPCDHCPOptionsAssociation struct {
  DhcpOptionsId *StringExpr `json:"DhcpOptionsId,omitempty"`  // String
  VpcId *StringExpr `json:"VpcId,omitempty"`  // String
}

func (s AWSEC2VPCDHCPOptionsAssociation) ResourceType() string {
	return "AWS::EC2::VPCDHCPOptionsAssociation"
}


type AWSEC2VPCGatewayAttachment struct {
  InternetGatewayId *StringExpr `json:"InternetGatewayId,omitempty"`  // String
  VpcId *StringExpr `json:"VpcId,omitempty"`  // String
  VpnGatewayId *StringExpr `json:"VpnGatewayId,omitempty"`  // String
}

func (s AWSEC2VPCGatewayAttachment) ResourceType() string {
	return "AWS::EC2::VPCGatewayAttachment"
}


type AWSEC2VPCPeeringConnection struct {
  PeerVpcId *StringExpr `json:"PeerVpcId,omitempty"`  // String
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
  VpcId *StringExpr `json:"VpcId,omitempty"`  // String
}

func (s AWSEC2VPCPeeringConnection) ResourceType() string {
	return "AWS::EC2::VPCPeeringConnection"
}


type AWSEC2VPNConnection struct {
  Type *StringExpr `json:"Type,omitempty"`  // String
  CustomerGatewayId *StringExpr `json:"CustomerGatewayId,omitempty"`  // String
  StaticRoutesOnly *BoolExpr `json:"StaticRoutesOnly,omitempty"`  // Boolean
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
  VpnGatewayId *StringExpr `json:"VpnGatewayId,omitempty"`  // String
}

func (s AWSEC2VPNConnection) ResourceType() string {
	return "AWS::EC2::VPNConnection"
}


type AWSEC2VPNConnectionRoute struct {
  DestinationCidrBlock *StringExpr `json:"DestinationCidrBlock,omitempty"`  // String
  VpnConnectionId *StringExpr `json:"VpnConnectionId,omitempty"`  // String
}

func (s AWSEC2VPNConnectionRoute) ResourceType() string {
	return "AWS::EC2::VPNConnectionRoute"
}


type AWSEC2VPNGateway struct {
  Type *StringExpr `json:"Type,omitempty"`  // String
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
}

func (s AWSEC2VPNGateway) ResourceType() string {
	return "AWS::EC2::VPNGateway"
}


type AWSEC2VPNGatewayRoutePropagation struct {
  RouteTableIds interface{} `json:"RouteTableIds,omitempty"`  // List of route table IDs
  VpnGatewayId *StringExpr `json:"VpnGatewayId,omitempty"`  // String
}

func (s AWSEC2VPNGatewayRoutePropagation) ResourceType() string {
	return "AWS::EC2::VPNGatewayRoutePropagation"
}


type AWSECSCluster struct {
}

func (s AWSECSCluster) ResourceType() string {
	return "AWS::ECS::Cluster"
}


type AWSECSService struct {
  Cluster *StringExpr `json:"Cluster,omitempty"`  // String
  DesiredCount *StringExpr `json:"DesiredCount,omitempty"`  // String
  LoadBalancers *AmazonEC2ContainerServiceServiceLoadBalancersList `json:"LoadBalancers,omitempty"`  // List of Amazon EC2 Container Service Service LoadBalancers
  Role *StringExpr `json:"Role,omitempty"`  // String
  TaskDefinition *StringExpr `json:"TaskDefinition,omitempty"`  // String
}

func (s AWSECSService) ResourceType() string {
	return "AWS::ECS::Service"
}


type AWSECSTaskDefinition struct {
  ContainerDefinitions *AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsList `json:"ContainerDefinitions,omitempty"`  // List of Amazon EC2 Container Service TaskDefinition ContainerDefinitions
  Volumes *AmazonEC2ContainerServiceTaskDefinitionVolumesList `json:"Volumes,omitempty"`  // List of Amazon EC2 Container Service TaskDefinition Volumes
}

func (s AWSECSTaskDefinition) ResourceType() string {
	return "AWS::ECS::TaskDefinition"
}


type AWSElastiCacheCacheCluster struct {
  AutoMinorVersionUpgrade *BoolExpr `json:"AutoMinorVersionUpgrade,omitempty"`  // Boolean
  AZMode *StringExpr `json:"AZMode,omitempty"`  // String
  CacheNodeType *StringExpr `json:"CacheNodeType,omitempty"`  // String
  CacheParameterGroupName *StringExpr `json:"CacheParameterGroupName,omitempty"`  // String
  CacheSecurityGroupNames *StringListExpr `json:"CacheSecurityGroupNames,omitempty"`  // List of strings
  CacheSubnetGroupName *StringExpr `json:"CacheSubnetGroupName,omitempty"`  // String
  ClusterName *StringExpr `json:"ClusterName,omitempty"`  // String
  Engine *StringExpr `json:"Engine,omitempty"`  // String
  EngineVersion *StringExpr `json:"EngineVersion,omitempty"`  // String
  NotificationTopicArn *StringExpr `json:"NotificationTopicArn,omitempty"`  // String
  NumCacheNodes *StringExpr `json:"NumCacheNodes,omitempty"`  // String
  Port *IntegerExpr `json:"Port,omitempty"`  // Integer
  PreferredAvailabilityZone *StringExpr `json:"PreferredAvailabilityZone,omitempty"`  // String
  PreferredAvailabilityZones *StringListExpr `json:"PreferredAvailabilityZones,omitempty"`  // List of strings
  PreferredMaintenanceWindow *StringExpr `json:"PreferredMaintenanceWindow,omitempty"`  // String
  SnapshotArns *StringListExpr `json:"SnapshotArns,omitempty"`  // List of strings
  SnapshotName *StringExpr `json:"SnapshotName,omitempty"`  // String
  SnapshotRetentionLimit *IntegerExpr `json:"SnapshotRetentionLimit,omitempty"`  // Integer
  SnapshotWindow *StringExpr `json:"SnapshotWindow,omitempty"`  // String
  XVpcSecurityGroupIdsX *StringListExpr `json:" VpcSecurityGroupIds ,omitempty"`  // List of strings
}

func (s AWSElastiCacheCacheCluster) ResourceType() string {
	return "AWS::ElastiCache::CacheCluster"
}


type AWSElastiCacheParameterGroup struct {
  CacheParameterGroupFamily *StringExpr `json:"CacheParameterGroupFamily,omitempty"`  // String
  Description *StringExpr `json:"Description,omitempty"`  // String
  Properties interface{} `json:"Properties,omitempty"`  // JSON object
}

func (s AWSElastiCacheParameterGroup) ResourceType() string {
	return "AWS::ElastiCache::ParameterGroup"
}


type AWSElastiCacheReplicationGroup struct {
  AutomaticFailoverEnabled *BoolExpr `json:"AutomaticFailoverEnabled,omitempty"`  // Boolean
  AutoMinorVersionUpgrade *BoolExpr `json:"AutoMinorVersionUpgrade,omitempty"`  // Boolean
  CacheNodeType *StringExpr `json:"CacheNodeType,omitempty"`  // String
  CacheParameterGroupName *StringExpr `json:"CacheParameterGroupName,omitempty"`  // String
  CacheSecurityGroupNames *StringListExpr `json:"CacheSecurityGroupNames,omitempty"`  // List of strings
  CacheSubnetGroupName *StringExpr `json:"CacheSubnetGroupName,omitempty"`  // String
  Engine *StringExpr `json:"Engine,omitempty"`  // String
  EngineVersion *StringExpr `json:"EngineVersion,omitempty"`  // String
  NotificationTopicArn *StringExpr `json:"NotificationTopicArn,omitempty"`  // String
  NumCacheClusters *IntegerExpr `json:"NumCacheClusters,omitempty"`  // Integer
  Port *IntegerExpr `json:"Port,omitempty"`  // Integer
  PreferredCacheClusterAZs *StringListExpr `json:"PreferredCacheClusterAZs,omitempty"`  // List of strings
  PreferredMaintenanceWindow *StringExpr `json:"PreferredMaintenanceWindow,omitempty"`  // String
  ReplicationGroupDescription *StringExpr `json:"ReplicationGroupDescription,omitempty"`  // String
  SecurityGroupIds *StringListExpr `json:"SecurityGroupIds,omitempty"`  // List of strings
  SnapshotArns *StringListExpr `json:"SnapshotArns,omitempty"`  // List of strings
  SnapshotRetentionLimit *IntegerExpr `json:"SnapshotRetentionLimit,omitempty"`  // Integer
  SnapshotWindow *StringExpr `json:"SnapshotWindow,omitempty"`  // String
}

func (s AWSElastiCacheReplicationGroup) ResourceType() string {
	return "AWS::ElastiCache::ReplicationGroup"
}


type AWSElastiCacheSecurityGroup struct {
  Description *StringExpr `json:"Description,omitempty"`  // String
}

func (s AWSElastiCacheSecurityGroup) ResourceType() string {
	return "AWS::ElastiCache::SecurityGroup"
}


type AWSElastiCacheSecurityGroupIngress struct {
  CacheSecurityGroupName *StringExpr `json:"CacheSecurityGroupName,omitempty"`  // String
  EC2SecurityGroupName *StringExpr `json:"EC2SecurityGroupName,omitempty"`  // String
  EC2SecurityGroupOwnerId *StringExpr `json:"EC2SecurityGroupOwnerId,omitempty"`  // String
}

func (s AWSElastiCacheSecurityGroupIngress) ResourceType() string {
	return "AWS::ElastiCache::SecurityGroupIngress"
}


type AWSElastiCacheSubnetGroup struct {
}

func (s AWSElastiCacheSubnetGroup) ResourceType() string {
	return "AWS::ElastiCache::SubnetGroup "
}


type AWSElasticBeanstalkApplication struct {
  ApplicationName *StringExpr `json:"ApplicationName,omitempty"`  // String
  Description *StringExpr `json:"Description,omitempty"`  // String
}

func (s AWSElasticBeanstalkApplication) ResourceType() string {
	return "AWS::ElasticBeanstalk::Application"
}


type AWSElasticBeanstalkApplicationVersion struct {
}

func (s AWSElasticBeanstalkApplicationVersion) ResourceType() string {
	return "AWS::ElasticBeanstalk::ApplicationVersion"
}


type AWSElasticBeanstalkConfigurationTemplate struct {
}

func (s AWSElasticBeanstalkConfigurationTemplate) ResourceType() string {
	return "AWS::ElasticBeanstalk::ConfigurationTemplate"
}


type AWSElasticBeanstalkEnvironment struct {
  ApplicationName *StringExpr `json:"ApplicationName,omitempty"`  // String
  CNAMEPrefix *StringExpr `json:"CNAMEPrefix,omitempty"`  // String
  Description *StringExpr `json:"Description,omitempty"`  // String
  EnvironmentName *StringExpr `json:"EnvironmentName,omitempty"`  // String
  OptionSettings *ElasticBeanstalkOptionSettingsList `json:"OptionSettings,omitempty"`  // A list of OptionSettings
  SolutionStackName *StringExpr `json:"SolutionStackName,omitempty"`  // String
  TemplateName *StringExpr `json:"TemplateName,omitempty"`  // String
  Tier *ElasticBeanstalkEnvironmentTier `json:"Tier,omitempty"`  // Elastic Beanstalk Environment Tier Property Type
  VersionLabel *StringExpr `json:"VersionLabel,omitempty"`  // String
}

func (s AWSElasticBeanstalkEnvironment) ResourceType() string {
	return "AWS::ElasticBeanstalk::Environment"
}


type AWSElasticLoadBalancingLoadBalancer struct {
  AccessLoggingPolicy *ElasticLoadBalancingAccessLoggingPolicy `json:"AccessLoggingPolicy,omitempty"`  // Elastic Load Balancing AccessLoggingPolicy
  AppCookieStickinessPolicy *ElasticLoadBalancingAppCookieStickinessPolicyList `json:"AppCookieStickinessPolicy,omitempty"`  // A list of AppCookieStickinessPolicy objects
  AvailabilityZones *StringListExpr `json:"AvailabilityZones,omitempty"`  // List of strings
  ConnectionDrainingPolicy *ElasticLoadBalancingConnectionDrainingPolicy `json:"ConnectionDrainingPolicy,omitempty"`  // Elastic Load Balancing ConnectionDrainingPolicy
  ConnectionSettings *ElasticLoadBalancingConnectionSettings `json:"ConnectionSettings,omitempty"`  // Elastic Load Balancing ConnectionSettings
  CrossZone *BoolExpr `json:"CrossZone,omitempty"`  // Boolean
  HealthCheck *ElasticLoadBalancingHealthCheck `json:"HealthCheck,omitempty"`  // ElasticLoadBalancing HealthCheck Type
  Instances *StringListExpr `json:"Instances,omitempty"`  // List of strings
  LBCookieStickinessPolicy *ElasticLoadBalancingLBCookieStickinessPolicyList `json:"LBCookieStickinessPolicy,omitempty"`  // A list of LBCookieStickinessPolicy objects
  LoadBalancerName *StringExpr `json:"LoadBalancerName,omitempty"`  // String
  Listeners *ElasticLoadBalancingListenerList `json:"Listeners,omitempty"`  // A list of ElasticLoadBalancing Listener Property Type objects
  Policies *ElasticLoadBalancingPolicyList `json:"Policies,omitempty"`  // A list of ElasticLoadBalancing policy objects
  Scheme *StringExpr `json:"Scheme,omitempty"`  // String
  SecurityGroups interface{} `json:"SecurityGroups,omitempty"`  // A list of security groups assigned to your load balancer within your virtual private cloud (VPC)
  Subnets *StringListExpr `json:"Subnets,omitempty"`  // List of strings
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
}

func (s AWSElasticLoadBalancingLoadBalancer) ResourceType() string {
	return "AWS::ElasticLoadBalancing::LoadBalancer"
}


type AWSIAMAccessKey struct {
  Serial *IntegerExpr `json:"Serial,omitempty"`  // Integer
  Status *StringExpr `json:"Status,omitempty"`  // String
  UserName *StringExpr `json:"UserName,omitempty"`  // String
}

func (s AWSIAMAccessKey) ResourceType() string {
	return "AWS::IAM::AccessKey"
}


type AWSIAMGroup struct {
  ManagedPolicyArns *StringListExpr `json:"ManagedPolicyArns,omitempty"`  // List of strings
  Path *StringExpr `json:"Path,omitempty"`  // String
  Policies *IAMPoliciesList `json:"Policies,omitempty"`  // List of IAM Policies
}

func (s AWSIAMGroup) ResourceType() string {
	return "AWS::IAM::Group"
}


type AWSIAMInstanceProfile struct {
  Path *StringExpr `json:"Path,omitempty"`  // String
  Roles interface{} `json:"Roles,omitempty"`  // List of references to AWS::IAM::Roles. Currently, a maximum of one role can be assigned to an instance profile
}

func (s AWSIAMInstanceProfile) ResourceType() string {
	return "AWS::IAM::InstanceProfile"
}


type AWSIAMManagedPolicy struct {
  Description *StringExpr `json:"Description,omitempty"`  // String
  Groups *StringListExpr `json:"Groups,omitempty"`  // List of strings
  Path *StringExpr `json:"Path,omitempty"`  // String
  PolicyDocument interface{} `json:"PolicyDocument,omitempty"`  // JSON object
  Roles *StringListExpr `json:"Roles,omitempty"`  // List of strings
  Users *StringListExpr `json:"Users,omitempty"`  // List of strings
}

func (s AWSIAMManagedPolicy) ResourceType() string {
	return "AWS::IAM::ManagedPolicy"
}


type AWSIAMPolicy struct {
  Groups *StringListExpr `json:"Groups,omitempty"`  // List of strings
  PolicyDocument interface{} `json:"PolicyDocument,omitempty"`  // JSON object
  PolicyName *StringExpr `json:"PolicyName,omitempty"`  // String
  Roles *StringListExpr `json:"Roles,omitempty"`  // List of strings
  Users *StringListExpr `json:"Users,omitempty"`  // List of strings
}

func (s AWSIAMPolicy) ResourceType() string {
	return "AWS::IAM::Policy"
}


type AWSIAMRole struct {
  AssumeRolePolicyDocument interface{} `json:"AssumeRolePolicyDocument,omitempty"`  // A JSON policy document
  ManagedPolicyArns *StringListExpr `json:"ManagedPolicyArns,omitempty"`  // List of strings
  Path *StringExpr `json:"Path,omitempty"`  // String
  Policies *IAMPoliciesList `json:"Policies,omitempty"`  // List of IAM Policies
}

func (s AWSIAMRole) ResourceType() string {
	return "AWS::IAM::Role"
}


type AWSIAMUser struct {
  Groups *StringListExpr `json:"Groups,omitempty"`  // List of strings
  LoginProfile *IAMUserLoginProfile `json:"LoginProfile,omitempty"`  // IAM User LoginProfile
  ManagedPolicyArns *StringListExpr `json:"ManagedPolicyArns,omitempty"`  // List of strings
  Path *StringExpr `json:"Path,omitempty"`  // String
  Policies *IAMPoliciesList `json:"Policies,omitempty"`  // List of IAM Policies
}

func (s AWSIAMUser) ResourceType() string {
	return "AWS::IAM::User"
}


type AWSIAMUserToGroupAddition struct {
  GroupName *StringExpr `json:"GroupName,omitempty"`  // String
  Users interface{} `json:"Users,omitempty"`  // List of users
}

func (s AWSIAMUserToGroupAddition) ResourceType() string {
	return "AWS::IAM::UserToGroupAddition"
}


type AWSKinesisStream struct {
  ShardCount *IntegerExpr `json:"ShardCount,omitempty"`  // Integer
}

func (s AWSKinesisStream) ResourceType() string {
	return "AWS::Kinesis::Stream"
}


type AWSLambdaFunction struct {
  Code *AWSLambdaFunctionCode `json:"Code,omitempty"`  // AWS Lambda Function Code
  Description *StringExpr `json:"Description,omitempty"`  // String
  Handler *StringExpr `json:"Handler,omitempty"`  // String
  MemorySize *IntegerExpr `json:"MemorySize,omitempty"`  // Integer
  Role *StringExpr `json:"Role,omitempty"`  // String
  Runtime *StringExpr `json:"Runtime,omitempty"`  // String
  Timeout *IntegerExpr `json:"Timeout,omitempty"`  // Integer
}

func (s AWSLambdaFunction) ResourceType() string {
	return "AWS::Lambda::Function"
}


type AWSLogsLogGroup struct {
  RetentionInDays *IntegerExpr `json:"RetentionInDays,omitempty"`  // Integer
}

func (s AWSLogsLogGroup) ResourceType() string {
	return "AWS::Logs::LogGroup"
}


type AWSLogsMetricFilter struct {
  FilterPattern *StringListExpr `json:"FilterPattern,omitempty"`  // List of strings
  LogGroupName *StringExpr `json:"LogGroupName,omitempty"`  // String
  MetricTransformations *CloudWatchLogsMetricFilterMetricTransformationPropertyList `json:"MetricTransformations,omitempty"`  // A list of CloudWatch Logs MetricFilter MetricTransformation Property
}

func (s AWSLogsMetricFilter) ResourceType() string {
	return "AWS::Logs::MetricFilter"
}


type AWSOpsWorksApp struct {
  AppSource *AWSOpsWorksSource `json:"AppSource,omitempty"`  // AWS OpsWorks Source Type
  Attributes interface{} `json:"Attributes,omitempty"`  // A list of key-value pairs
  Description *StringExpr `json:"Description,omitempty"`  // String
  Domains *StringListExpr `json:"Domains,omitempty"`  // List of strings
  EnableSsl *BoolExpr `json:"EnableSsl,omitempty"`  // Boolean
  Name *StringExpr `json:"Name,omitempty"`  // String
  Shortname *StringExpr `json:"Shortname,omitempty"`  // String
  SslConfiguration *AWSOpsWorksSslConfiguration `json:"SslConfiguration,omitempty"`  // AWS OpsWorks SslConfiguration Type
  StackId *StringExpr `json:"StackId,omitempty"`  // String
  Type *StringExpr `json:"Type,omitempty"`  // String
}

func (s AWSOpsWorksApp) ResourceType() string {
	return "AWS::OpsWorks::App"
}


type AWSOpsWorksElasticLoadBalancerAttachment struct {
  ElasticLoadBalancerName *StringExpr `json:"ElasticLoadBalancerName,omitempty"`  // String
  LayerId *StringExpr `json:"LayerId,omitempty"`  // String
}

func (s AWSOpsWorksElasticLoadBalancerAttachment) ResourceType() string {
	return "AWS::OpsWorks::ElasticLoadBalancerAttachment"
}


type AWSOpsWorksInstance struct {
  AmiId *StringExpr `json:"AmiId,omitempty"`  // String
  Architecture *StringExpr `json:"Architecture,omitempty"`  // String
  AutoScalingType *StringExpr `json:"AutoScalingType,omitempty"`  // String
  AvailabilityZone *StringExpr `json:"AvailabilityZone,omitempty"`  // String
  InstallUpdatesOnBoot *BoolExpr `json:"InstallUpdatesOnBoot,omitempty"`  // Boolean
  InstanceType *StringExpr `json:"InstanceType,omitempty"`  // String
  LayerIds *StringListExpr `json:"LayerIds,omitempty"`  // List of strings
  Os *StringExpr `json:"Os,omitempty"`  // String
  RootDeviceType *StringExpr `json:"RootDeviceType,omitempty"`  // String
  SshKeyName *StringExpr `json:"SshKeyName,omitempty"`  // String
  StackId *StringExpr `json:"StackId,omitempty"`  // String
  SubnetId *StringExpr `json:"SubnetId,omitempty"`  // String
  TimeBasedAutoScaling *AWSOpsWorksTimeBasedAutoScaling `json:"TimeBasedAutoScaling,omitempty"`  // AWS OpsWorks TimeBasedAutoScaling Type
}

func (s AWSOpsWorksInstance) ResourceType() string {
	return "AWS::OpsWorks::Instance"
}


type AWSOpsWorksLayer struct {
  Attributes interface{} `json:"Attributes,omitempty"`  // A list of key-value pairs
  AutoAssignElasticIps *BoolExpr `json:"AutoAssignElasticIps,omitempty"`  // Boolean
  AutoAssignPublicIps *BoolExpr `json:"AutoAssignPublicIps,omitempty"`  // Boolean
  CustomInstanceProfileArn *StringExpr `json:"CustomInstanceProfileArn,omitempty"`  // String
  CustomRecipes *AWSOpsWorksRecipes `json:"CustomRecipes,omitempty"`  // AWS OpsWorks Recipes Type
  CustomSecurityGroupIds *StringListExpr `json:"CustomSecurityGroupIds,omitempty"`  // List of strings
  EnableAutoHealing *BoolExpr `json:"EnableAutoHealing,omitempty"`  // Boolean
  InstallUpdatesOnBoot *BoolExpr `json:"InstallUpdatesOnBoot,omitempty"`  // Boolean
  LifecycleEventConfiguration *AWSOpsWorksLayerLifeCycleConfiguration `json:"LifecycleEventConfiguration,omitempty"`  // AWS OpsWorks Layer LifeCycleConfiguration
  LoadBasedAutoScaling *AWSOpsWorksLoadBasedAutoScaling `json:"LoadBasedAutoScaling,omitempty"`  // AWS OpsWorks LoadBasedAutoScaling Type
  Name *StringExpr `json:"Name,omitempty"`  // String
  Packages *StringListExpr `json:"Packages,omitempty"`  // List of strings
  Shortname *StringExpr `json:"Shortname,omitempty"`  // String
  StackId *StringExpr `json:"StackId,omitempty"`  // String
  Type *StringExpr `json:"Type,omitempty"`  // String
  VolumeConfigurations *AWSOpsWorksVolumeConfigurationList `json:"VolumeConfigurations,omitempty"`  // A list of AWS OpsWorks VolumeConfiguration Type
}

func (s AWSOpsWorksLayer) ResourceType() string {
	return "AWS::OpsWorks::Layer"
}


type AWSOpsWorksStack struct {
  Attributes interface{} `json:"Attributes,omitempty"`  // A list of key-value pairs
  ChefConfiguration *AWSOpsWorksChefConfiguration `json:"ChefConfiguration,omitempty"`  // AWS OpsWorks ChefConfiguration Type
  ConfigurationManager *AWSOpsWorksStackConfigurationManager `json:"ConfigurationManager,omitempty"`  // AWS OpsWorks StackConfigurationManager Type
  CustomCookbooksSource *AWSOpsWorksSource `json:"CustomCookbooksSource,omitempty"`  // AWS OpsWorks Source Type
  CustomJson interface{} `json:"CustomJson,omitempty"`  // JSON object
  DefaultAvailabilityZone *StringExpr `json:"DefaultAvailabilityZone,omitempty"`  // String
  DefaultInstanceProfileArn *StringExpr `json:"DefaultInstanceProfileArn,omitempty"`  // String
  DefaultOs *StringExpr `json:"DefaultOs,omitempty"`  // String
  DefaultRootDeviceType *StringExpr `json:"DefaultRootDeviceType,omitempty"`  // String
  DefaultSshKeyName *StringExpr `json:"DefaultSshKeyName,omitempty"`  // String
  DefaultSubnetId *StringExpr `json:"DefaultSubnetId,omitempty"`  // String
  HostnameTheme *StringExpr `json:"HostnameTheme,omitempty"`  // String
  Name *StringExpr `json:"Name,omitempty"`  // String
  ServiceRoleArn *StringExpr `json:"ServiceRoleArn,omitempty"`  // String
  UseCustomCookbooks *BoolExpr `json:"UseCustomCookbooks,omitempty"`  // Boolean
  UseOpsworksSecurityGroups *BoolExpr `json:"UseOpsworksSecurityGroups,omitempty"`  // Boolean
  VpcId *StringExpr `json:"VpcId,omitempty"`  // String
}

func (s AWSOpsWorksStack) ResourceType() string {
	return "AWS::OpsWorks::Stack"
}


type AWSRedshiftCluster struct {
  AllowVersionUpgrade *BoolExpr `json:"AllowVersionUpgrade,omitempty"`  // Boolean
  AutomatedSnapshotRetentionPeriod *IntegerExpr `json:"AutomatedSnapshotRetentionPeriod,omitempty"`  // Integer
  AvailabilityZone *StringExpr `json:"AvailabilityZone,omitempty"`  // String
  ClusterParameterGroupName *StringExpr `json:"ClusterParameterGroupName,omitempty"`  // String
  ClusterSecurityGroups *StringListExpr `json:"ClusterSecurityGroups,omitempty"`  // List of strings
  ClusterSubnetGroupName *StringExpr `json:"ClusterSubnetGroupName,omitempty"`  // String
  ClusterType *StringExpr `json:"ClusterType,omitempty"`  // String
  ClusterVersion *StringExpr `json:"ClusterVersion,omitempty"`  // String
  DBName *StringExpr `json:"DBName,omitempty"`  // String
  ElasticIp *StringExpr `json:"ElasticIp,omitempty"`  // String
  Encrypted *BoolExpr `json:"Encrypted,omitempty"`  // Boolean
  HsmClientCertificateIdentifier *StringExpr `json:"HsmClientCertificateIdentifier,omitempty"`  // String
  HsmConfigurationIdentifier *StringExpr `json:"HsmConfigurationIdentifier,omitempty"`  // String
  MasterUsername *StringExpr `json:"MasterUsername,omitempty"`  // String
  MasterUserPassword *StringExpr `json:"MasterUserPassword,omitempty"`  // String
  NodeType *StringExpr `json:"NodeType,omitempty"`  // String
  NumberOfNodes *IntegerExpr `json:"NumberOfNodes,omitempty"`  // Integer
  OwnerAccount *StringExpr `json:"OwnerAccount,omitempty"`  // String
  Port *IntegerExpr `json:"Port,omitempty"`  // Integer
  PreferredMaintenanceWindow *StringExpr `json:"PreferredMaintenanceWindow,omitempty"`  // String
  PubliclyAccessible *BoolExpr `json:"PubliclyAccessible,omitempty"`  // Boolean
  SnapshotClusterIdentifier interface{} `json:"SnapshotClusterIdentifier,omitempty"`  // 
  SnapshotIdentifier *StringExpr `json:"SnapshotIdentifier,omitempty"`  // String
  VpcSecurityGroupIds *StringListExpr `json:"VpcSecurityGroupIds,omitempty"`  // List of strings
}

func (s AWSRedshiftCluster) ResourceType() string {
	return "AWS::Redshift::Cluster"
}


type AWSRedshiftClusterParameterGroup struct {
  Description *StringExpr `json:"Description,omitempty"`  // String
  ParameterGroupFamily *StringExpr `json:"ParameterGroupFamily,omitempty"`  // String
  Parameters *AmazonRedshiftParameterList `json:"Parameters,omitempty"`  // Amazon Redshift Parameter Type
}

func (s AWSRedshiftClusterParameterGroup) ResourceType() string {
	return "AWS::Redshift::ClusterParameterGroup"
}


type AWSRedshiftClusterSecurityGroup struct {
  Description *StringExpr `json:"Description,omitempty"`  // String
}

func (s AWSRedshiftClusterSecurityGroup) ResourceType() string {
	return "AWS::Redshift::ClusterSecurityGroup"
}


type AWSRedshiftClusterSecurityGroupIngress struct {
  ClusterSecurityGroupName *StringExpr `json:"ClusterSecurityGroupName,omitempty"`  // String
  CIDRIP *StringExpr `json:"CIDRIP,omitempty"`  // String
  EC2SecurityGroupName *StringExpr `json:"EC2SecurityGroupName,omitempty"`  // String
  EC2SecurityGroupOwnerId *StringExpr `json:"EC2SecurityGroupOwnerId,omitempty"`  // String
}

func (s AWSRedshiftClusterSecurityGroupIngress) ResourceType() string {
	return "AWS::Redshift::ClusterSecurityGroupIngress"
}


type AWSRedshiftClusterSubnetGroup struct {
  Description *StringExpr `json:"Description,omitempty"`  // String
  SubnetIds *StringListExpr `json:"SubnetIds,omitempty"`  // List of strings
}

func (s AWSRedshiftClusterSubnetGroup) ResourceType() string {
	return "AWS::Redshift::ClusterSubnetGroup"
}


type AWSRDSDBInstance struct {
  AllocatedStorage *StringExpr `json:"AllocatedStorage,omitempty"`  // String
  AllowMajorVersionUpgrade *BoolExpr `json:"AllowMajorVersionUpgrade,omitempty"`  // Boolean
  AutoMinorVersionUpgrade *BoolExpr `json:"AutoMinorVersionUpgrade,omitempty"`  // Boolean
  AvailabilityZone *StringExpr `json:"AvailabilityZone,omitempty"`  // String
  BackupRetentionPeriod *StringExpr `json:"BackupRetentionPeriod,omitempty"`  // String
  CharacterSetName *StringExpr `json:"CharacterSetName,omitempty"`  // String
  DBInstanceClass *StringExpr `json:"DBInstanceClass,omitempty"`  // String
  DBInstanceIdentifier *StringExpr `json:"DBInstanceIdentifier,omitempty"`  // String
  DBName *StringExpr `json:"DBName,omitempty"`  // String
  DBParameterGroupName *StringExpr `json:"DBParameterGroupName,omitempty"`  // String
  DBSecurityGroups *StringListExpr `json:"DBSecurityGroups,omitempty"`  // List of strings
  DBSnapshotIdentifier *StringExpr `json:"DBSnapshotIdentifier,omitempty"`  // String
  DBSubnetGroupName *StringExpr `json:"DBSubnetGroupName,omitempty"`  // String
  Engine *StringExpr `json:"Engine,omitempty"`  // String
  EngineVersion *StringExpr `json:"EngineVersion,omitempty"`  // String
  Iops *IntegerExpr `json:"Iops,omitempty"`  // Number
  KmsKeyId *StringExpr `json:"KmsKeyId,omitempty"`  // String
  LicenseModel *StringExpr `json:"LicenseModel,omitempty"`  // String
  MasterUsername *StringExpr `json:"MasterUsername,omitempty"`  // String
  MasterUserPassword *StringExpr `json:"MasterUserPassword,omitempty"`  // String
  MultiAZ *BoolExpr `json:"MultiAZ,omitempty"`  // Boolean
  OptionGroupName *StringExpr `json:"OptionGroupName,omitempty"`  // String
  Port *StringExpr `json:"Port,omitempty"`  // String
  PreferredBackupWindow *StringExpr `json:"PreferredBackupWindow,omitempty"`  // String
  PreferredMaintenanceWindow *StringExpr `json:"PreferredMaintenanceWindow,omitempty"`  // String
  PubliclyAccessible *BoolExpr `json:"PubliclyAccessible,omitempty"`  // Boolean
  SourceDBInstanceIdentifier *StringExpr `json:"SourceDBInstanceIdentifier,omitempty"`  // String
  StorageEncrypted *BoolExpr `json:"StorageEncrypted,omitempty"`  // Boolean
  StorageType *StringExpr `json:"StorageType,omitempty"`  // String
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
  VPCSecurityGroups *StringListExpr `json:"VPCSecurityGroups,omitempty"`  // List of strings
}

func (s AWSRDSDBInstance) ResourceType() string {
	return "AWS::RDS::DBInstance"
}


type AWSRDSDBParameterGroup struct {
  Description interface{} `json:"Description,omitempty"`  // 
  Family interface{} `json:"Family,omitempty"`  // 
  Parameters interface{} `json:"Parameters,omitempty"`  // 
  Tags *ResourceTagList `json:"Tags,omitempty"`  // A list of resource tags
}

func (s AWSRDSDBParameterGroup) ResourceType() string {
	return "AWS::RDS::DBParameterGroup"
}


type AWSRDSDBSubnetGroup struct {
  DBSubnetGroupDescription *StringExpr `json:"DBSubnetGroupDescription,omitempty"`  // String
  SubnetIds *StringListExpr `json:"SubnetIds,omitempty"`  // List of strings
  Tags *ResourceTagList `json:"Tags,omitempty"`  // A list of resource tags
}

func (s AWSRDSDBSubnetGroup) ResourceType() string {
	return "AWS::RDS::DBSubnetGroup"
}


type AWSRDSDBSecurityGroup struct {
  EC2VpcId *StringExpr `json:"EC2VpcId,omitempty"`  // String
  DBSecurityGroupIngress *AmazonRDSSecurityGroupRuleList `json:"DBSecurityGroupIngress,omitempty"`  // List of RDS Security Group Rules
  GroupDescription *StringExpr `json:"GroupDescription,omitempty"`  // String
  Tags *ResourceTagList `json:"Tags,omitempty"`  // A list of resource tags
}

func (s AWSRDSDBSecurityGroup) ResourceType() string {
	return "AWS::RDS::DBSecurityGroup"
}


type AWSRDSDBSecurityGroupIngress struct {
  CIDRIP *StringExpr `json:"CIDRIP,omitempty"`  // String
  DBSecurityGroupName *StringExpr `json:"DBSecurityGroupName,omitempty"`  // String
  EC2SecurityGroupId *StringExpr `json:"EC2SecurityGroupId,omitempty"`  // String
  EC2SecurityGroupName *StringExpr `json:"EC2SecurityGroupName,omitempty"`  // String
  EC2SecurityGroupOwnerId *StringExpr `json:"EC2SecurityGroupOwnerId,omitempty"`  // String
}

func (s AWSRDSDBSecurityGroupIngress) ResourceType() string {
	return "AWS::RDS::DBSecurityGroupIngress"
}


type AWSRDSEventSubscription struct {
  Enabled *BoolExpr `json:"Enabled,omitempty"`  // Boolean
  EventCategories *StringListExpr `json:"EventCategories,omitempty"`  // List of strings
  SnsTopicArn *StringExpr `json:"SnsTopicArn,omitempty"`  // String
  SourceIds *StringListExpr `json:"SourceIds,omitempty"`  // List of strings
  SourceTypeX *StringExpr `json:"SourceType ,omitempty"`  // String
}

func (s AWSRDSEventSubscription) ResourceType() string {
	return "AWS::RDS::EventSubscription"
}


type AWSRDSOptionGroup struct {
  EngineName *StringExpr `json:"EngineName,omitempty"`  // String
  MajorEngineVersion *StringExpr `json:"MajorEngineVersion,omitempty"`  // String
  OptionGroupDescription *StringExpr `json:"OptionGroupDescription,omitempty"`  // String
  OptionConfigurations *AmazonRDSOptionGroupOptionConfigurations `json:"OptionConfigurations,omitempty"`  // Amazon RDS OptionGroup OptionConfigurations
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
}

func (s AWSRDSOptionGroup) ResourceType() string {
	return "AWS::RDS::OptionGroup"
}


type AWSRoute53HealthCheck struct {
  HealthCheckConfig *AmazonRoute53HealthCheckConfig `json:"HealthCheckConfig,omitempty"`  // Amazon Route 53 HealthCheckConfig
  HealthCheckTags *AmazonRoute53HealthCheckTagsList `json:"HealthCheckTags,omitempty"`  // List of Amazon Route 53 HealthCheckTags
}

func (s AWSRoute53HealthCheck) ResourceType() string {
	return "AWS::Route53::HealthCheck"
}


type AWSRoute53HostedZone struct {
  HostedZoneConfig *AmazonRoute53HostedZoneConfigProperty `json:"HostedZoneConfig,omitempty"`  // Amazon Route 53 HostedZoneConfig Property
  HostedZoneTags *AmazonRoute53HostedZoneTagsList `json:"HostedZoneTags,omitempty"`  // List of Amazon Route 53 HostedZoneTags
  Name *StringExpr `json:"Name,omitempty"`  // String
  VPCs *AmazonRoute53HostedZoneVPCsList `json:"VPCs,omitempty"`  // List of Amazon Route 53 HostedZoneVPCs
}

func (s AWSRoute53HostedZone) ResourceType() string {
	return "AWS::Route53::HostedZone"
}


type AWSRoute53RecordSet struct {
  AliasTarget *Route53AliasTargetProperty `json:"AliasTarget,omitempty"`  // AliasTarget
  Comment *StringExpr `json:"Comment,omitempty"`  // String
  Failover *StringExpr `json:"Failover,omitempty"`  // String
  GeoLocation *AmazonRoute53RecordSetGeoLocationProperty `json:"GeoLocation,omitempty"`  // Amazon Route 53 Record Set GeoLocation Property
  HealthCheckId *StringExpr `json:"HealthCheckId,omitempty"`  // String
  HostedZoneId *StringExpr `json:"HostedZoneId,omitempty"`  // String
  HostedZoneName *StringExpr `json:"HostedZoneName,omitempty"`  // String
  Name *StringExpr `json:"Name,omitempty"`  // String
  Region interface{} `json:"Region,omitempty"`  // 
  ResourceRecords *StringListExpr `json:"ResourceRecords,omitempty"`  // List of strings
  SetIdentifier *StringExpr `json:"SetIdentifier,omitempty"`  // String
  TTL *StringExpr `json:"TTL,omitempty"`  // String
  Type *StringExpr `json:"Type,omitempty"`  // String
  Weight *IntegerExpr `json:"Weight,omitempty"`  // Number. Weight expects integer values
}

func (s AWSRoute53RecordSet) ResourceType() string {
	return "AWS::Route53::RecordSet"
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
  HostedZoneId *StringExpr `json:"HostedZoneId,omitempty"`  // String
  HostedZoneName *StringExpr `json:"HostedZoneName,omitempty"`  // String
  RecordSets *AWSRoute53RecordSetList `json:"RecordSets,omitempty"`  // list of AWS::Route53::RecordSet
  Comment *StringExpr `json:"Comment,omitempty"`  // String
}

func (s AWSRoute53RecordSetGroup) ResourceType() string {
	return "AWS::Route53::RecordSetGroup"
}


type AWSS3Bucket struct {
  AccessControl *StringExpr `json:"AccessControl,omitempty"`  // String
  BucketName *StringExpr `json:"BucketName,omitempty"`  // String
  CorsConfiguration *AmazonS3CorsConfiguration `json:"CorsConfiguration,omitempty"`  // Amazon S3 Cors Configuration
  LifecycleConfiguration *AmazonS3LifecycleConfiguration `json:"LifecycleConfiguration,omitempty"`  // Amazon S3 Lifecycle Configuration
  LoggingConfiguration *AmazonS3LoggingConfiguration `json:"LoggingConfiguration,omitempty"`  // Amazon S3 Logging Configuration
  NotificationConfiguration *AmazonS3NotificationConfiguration `json:"NotificationConfiguration,omitempty"`  // Amazon S3 Notification Configuration
  Tags []ResourceTag `json:"Tags,omitempty"`  // AWS CloudFormation Resource Tags
  VersioningConfiguration *AmazonS3VersioningConfiguration `json:"VersioningConfiguration,omitempty"`  // Amazon S3 Versioning Configuration
  WebsiteConfiguration *AmazonS3WebsiteConfigurationProperty `json:"WebsiteConfiguration,omitempty"`  // Website Configuration Type
}

func (s AWSS3Bucket) ResourceType() string {
	return "AWS::S3::Bucket"
}


type AWSS3BucketPolicy struct {
  Bucket *StringExpr `json:"Bucket,omitempty"`  // String
  PolicyDocument interface{} `json:"PolicyDocument,omitempty"`  // JSON object
}

func (s AWSS3BucketPolicy) ResourceType() string {
	return "AWS::S3::BucketPolicy"
}


type AWSSDBDomain struct {
}

func (s AWSSDBDomain) ResourceType() string {
	return "AWS::SDB::Domain"
}


type AWSSNSTopic struct {
  DisplayName *StringExpr `json:"DisplayName,omitempty"`  // String
  Subscription *AmazonSNSSubscriptionList `json:"Subscription,omitempty"`  // List of SNS Subscriptions
  TopicName *StringExpr `json:"TopicName,omitempty"`  // String
}

func (s AWSSNSTopic) ResourceType() string {
	return "AWS::SNS::Topic"
}


type AWSSNSTopicPolicy struct {
  PolicyDocument interface{} `json:"PolicyDocument,omitempty"`  // JSON object
  Topics interface{} `json:"Topics,omitempty"`  // A list of Amazon SNS topics ARNs
}

func (s AWSSNSTopicPolicy) ResourceType() string {
	return "AWS::SNS::TopicPolicy"
}


type AWSSQSQueue struct {
  DelaySeconds *IntegerExpr `json:"DelaySeconds,omitempty"`  // Integer
  MaximumMessageSize *IntegerExpr `json:"MaximumMessageSize,omitempty"`  // Integer
  MessageRetentionPeriod *IntegerExpr `json:"MessageRetentionPeriod,omitempty"`  // Integer
  QueueName *StringExpr `json:"QueueName,omitempty"`  // String
  ReceiveMessageWaitTimeSeconds *IntegerExpr `json:"ReceiveMessageWaitTimeSeconds,omitempty"`  // Integer
  RedrivePolicy *AmazonSQSRedrivePolicy `json:"RedrivePolicy,omitempty"`  // Amazon SQS RedrivePolicy
  VisibilityTimeout *IntegerExpr `json:"VisibilityTimeout,omitempty"`  // Integer
}

func (s AWSSQSQueue) ResourceType() string {
	return "AWS::SQS::Queue"
}


type AWSSQSQueuePolicy struct {
  PolicyDocument interface{} `json:"PolicyDocument,omitempty"`  // JSON object
  Queues *StringListExpr `json:"Queues,omitempty"`  // List of strings
}

func (s AWSSQSQueuePolicy) ResourceType() string {
	return "AWS::SQS::QueuePolicy"
}


type AWSCloudFormationAutoScalingBlockDeviceMapping struct {
  DeviceName *StringExpr `json:"DeviceName,omitempty"`  // String
  Ebs *AWSCloudFormationAutoScalingEBSBlockDevice `json:"Ebs,omitempty"`  // AutoScaling EBS Block Device
  NoDevice *BoolExpr `json:"NoDevice,omitempty"`  // Boolean
  VirtualName *StringExpr `json:"VirtualName,omitempty"`  // String
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
  DeleteOnTermination *BoolExpr `json:"DeleteOnTermination,omitempty"`  // Boolean
  Iops *IntegerExpr `json:"Iops,omitempty"`  // Integer
  SnapshotId *StringExpr `json:"SnapshotId,omitempty"`  // String
  VolumeSize *IntegerExpr `json:"VolumeSize,omitempty"`  // Integer
  VolumeType *StringExpr `json:"VolumeType,omitempty"`  // String
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
  Granularity *StringExpr `json:"Granularity,omitempty"`  // String
  Metrics *StringListExpr `json:"Metrics,omitempty"`  // List of strings
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
  NotificationTypes *StringListExpr `json:"NotificationTypes,omitempty"`  // List of strings
  TopicARN *StringExpr `json:"TopicARN,omitempty"`  // String
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
  Key *StringExpr `json:"Key,omitempty"`  // String
  Value *StringExpr `json:"Value,omitempty"`  // String
  PropagateAtLaunch *BoolExpr `json:"PropagateAtLaunch,omitempty"`  // Boolean
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
  Aliases *StringListExpr `json:"Aliases,omitempty"`  // List of strings
  CacheBehaviors *CloudFrontDistributionConfigCacheBehaviorList `json:"CacheBehaviors,omitempty"`  // List of CacheBehavior
  Comment *StringExpr `json:"Comment,omitempty"`  // String
  CustomErrorResponses *CloudFrontDistributionConfigCustomErrorResponse `json:"CustomErrorResponses,omitempty"`  // Type List of CloudFront DistributionConfig CustomErrorResponse
  DefaultCacheBehavior *CloudFrontDefaultCacheBehavior `json:"DefaultCacheBehavior,omitempty"`  // DefaultCacheBehavior type
  DefaultRootObject *StringExpr `json:"DefaultRootObject,omitempty"`  // String
  Enabled *BoolExpr `json:"Enabled,omitempty"`  // Boolean
  Logging *CloudFrontLogging `json:"Logging,omitempty"`  // Logging type
  Origins *CloudFrontDistributionConfigOriginList `json:"Origins,omitempty"`  // List of Origins
  PriceClass *StringExpr `json:"PriceClass,omitempty"`  // String
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
  AllowedMethods *StringListExpr `json:"AllowedMethods,omitempty"`  // List of strings
  CachedMethods *StringListExpr `json:"CachedMethods,omitempty"`  // List of strings
  ForwardedValues *CloudFrontForwardedValues `json:"ForwardedValues,omitempty"`  // ForwardedValues type
  MinTTL *StringExpr `json:"MinTTL,omitempty"`  // String
  PathPattern *StringExpr `json:"PathPattern,omitempty"`  // String
  SmoothStreaming *BoolExpr `json:"SmoothStreaming,omitempty"`  // Boolean
  TargetOriginId *StringExpr `json:"TargetOriginId,omitempty"`  // String
  TrustedSigners *StringListExpr `json:"TrustedSigners,omitempty"`  // List of strings
  ViewerProtocolPolicy *StringExpr `json:"ViewerProtocolPolicy,omitempty"`  // String
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
  ErrorCachingMinTTL *IntegerExpr `json:"ErrorCachingMinTTL,omitempty"`  // Integer
  ErrorCode *IntegerExpr `json:"ErrorCode,omitempty"`  // Integer
  ResponseCode *IntegerExpr `json:"ResponseCode,omitempty"`  // Integer
  ResponsePagePath *StringExpr `json:"ResponsePagePath,omitempty"`  // String
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
  AllowedMethods *StringListExpr `json:"AllowedMethods,omitempty"`  // List of strings
  CachedMethods *StringListExpr `json:"CachedMethods,omitempty"`  // List of strings
  ForwardedValues *CloudFrontForwardedValues `json:"ForwardedValues,omitempty"`  // ForwardedValues type
  MinTTL *StringExpr `json:"MinTTL,omitempty"`  // String
  SmoothStreaming *BoolExpr `json:"SmoothStreaming,omitempty"`  // Boolean
  TargetOriginId *StringExpr `json:"TargetOriginId,omitempty"`  // String
  TrustedSigners *StringListExpr `json:"TrustedSigners,omitempty"`  // List of strings
  ViewerProtocolPolicy *StringExpr `json:"ViewerProtocolPolicy,omitempty"`  // String
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
  Bucket *StringExpr `json:"Bucket,omitempty"`  // String
  IncludeCookies *BoolExpr `json:"IncludeCookies,omitempty"`  // Boolean
  Prefix *StringExpr `json:"Prefix,omitempty"`  // String
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
  DomainName *StringExpr `json:"DomainName,omitempty"`  // String
  Id *StringExpr `json:"Id,omitempty"`  // String
  OriginPath *StringExpr `json:"OriginPath,omitempty"`  // String
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
  HTTPPort *StringExpr `json:"HTTPPort,omitempty"`  // String
  HTTPSPort *StringExpr `json:"HTTPSPort,omitempty"`  // String
  OriginProtocolPolicy *StringExpr `json:"OriginProtocolPolicy,omitempty"`  // String
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
  OriginAccessIdentity *StringExpr `json:"OriginAccessIdentity,omitempty"`  // String
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
  Locations *StringListExpr `json:"Locations,omitempty"`  // List of strings
  RestrictionType *StringExpr `json:"RestrictionType,omitempty"`  // String
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
  CloudFrontDefaultCertificate *BoolExpr `json:"CloudFrontDefaultCertificate,omitempty"`  // Boolean
  IamCertificateId *StringExpr `json:"IamCertificateId,omitempty"`  // String
  MinimumProtocolVersion *StringExpr `json:"MinimumProtocolVersion,omitempty"`  // String
  SslSupportMethod *StringExpr `json:"SslSupportMethod,omitempty"`  // String
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
  Headers *StringListExpr `json:"Headers,omitempty"`  // List of strings
  QueryString *BoolExpr `json:"QueryString,omitempty"`  // Boolean
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
  Forward *StringExpr `json:"Forward,omitempty"`  // String
  WhitelistedNames *StringListExpr `json:"WhitelistedNames,omitempty"`  // List of strings
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
  Name *StringExpr `json:"Name,omitempty"`  // String
  Value *StringExpr `json:"Value,omitempty"`  // String
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
  MetricName *StringExpr `json:"MetricName,omitempty"`  // String
  MetricNamespace *StringExpr `json:"MetricNamespace,omitempty"`  // String
  MetricValue *StringExpr `json:"MetricValue,omitempty"`  // String
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
  Id *StringExpr `json:"Id,omitempty"`  // String
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
  Key *StringExpr `json:"Key,omitempty"`  // String
  StringValue *StringExpr `json:"StringValue,omitempty"`  // String
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
  Id *StringExpr `json:"Id,omitempty"`  // String
  StringValue *StringExpr `json:"StringValue,omitempty"`  // String
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
  Id *StringExpr `json:"Id,omitempty"`  // String
  Name *StringExpr `json:"Name,omitempty"`  // String
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
  Key *StringExpr `json:"Key,omitempty"`  // String
  RefValue *StringExpr `json:"RefValue,omitempty"`  // String
  StringValue *StringExpr `json:"StringValue,omitempty"`  // String
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
  Key *StringExpr `json:"Key,omitempty"`  // String
  Value *StringExpr `json:"Value,omitempty"`  // String
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
  AttributeName *StringExpr `json:"AttributeName,omitempty"`  // String
  AttributeType *StringExpr `json:"AttributeType,omitempty"`  // String
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
  IndexName *StringExpr `json:"IndexName,omitempty"`  // String
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
  AttributeName *StringExpr `json:"AttributeName,omitempty"`  // String
  KeyType *StringExpr `json:"KeyType,omitempty"`  // String
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
  IndexName *StringExpr `json:"IndexName,omitempty"`  // String
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
  NonKeyAttributes *StringListExpr `json:"NonKeyAttributes,omitempty"`  // List of strings
  ProjectionType *StringExpr `json:"ProjectionType,omitempty"`  // String
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
  DeviceName *StringExpr `json:"DeviceName,omitempty"`  // String
  Ebs *AmazonElasticBlockStoreBlockDeviceProperty `json:"Ebs,omitempty"`  // Amazon Elastic Block Store Block Device Property
  NoDevice interface{} `json:"NoDevice,omitempty"`  // an empty map: {}
  VirtualName *StringExpr `json:"VirtualName,omitempty"`  // String
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
  DeleteOnTermination *BoolExpr `json:"DeleteOnTermination,omitempty"`  // Boolean
  Encrypted *BoolExpr `json:"Encrypted,omitempty"`  // Boolean
  Iops *IntegerExpr `json:"Iops,omitempty"`  // Number
  SnapshotId *StringExpr `json:"SnapshotId,omitempty"`  // String
  VolumeSize *StringExpr `json:"VolumeSize,omitempty"`  // String
  VolumeType *StringExpr `json:"VolumeType,omitempty"`  // String
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
  Device *StringExpr `json:"Device,omitempty"`  // String
  VolumeId *StringExpr `json:"VolumeId,omitempty"`  // String
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
  AssociatePublicIpAddress *BoolExpr `json:"AssociatePublicIpAddress,omitempty"`  // Boolean
  DeleteOnTermination *BoolExpr `json:"DeleteOnTermination,omitempty"`  // Boolean
  Description *StringExpr `json:"Description,omitempty"`  // String
  DeviceIndex *StringExpr `json:"DeviceIndex,omitempty"`  // String
  GroupSet *StringListExpr `json:"GroupSet,omitempty"`  // List of strings
  NetworkInterfaceId *StringExpr `json:"NetworkInterfaceId,omitempty"`  // String
  PrivateIpAddress *StringExpr `json:"PrivateIpAddress,omitempty"`  // String
  PrivateIpAddresses *EC2NetworkInterfacePrivateIPSpecificationList `json:"PrivateIpAddresses,omitempty"`  // list of PrivateIpAddressSpecification
  SecondaryPrivateIpAddressCount *IntegerExpr `json:"SecondaryPrivateIpAddressCount,omitempty"`  // Integer
  SubnetId *StringExpr `json:"SubnetId,omitempty"`  // String
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
  AttachmentID *StringExpr `json:"AttachmentID,omitempty"`  // String
  InstanceID *StringExpr `json:"InstanceID,omitempty"`  // String
  PublicIp *StringExpr `json:"PublicIp,omitempty"`  // String
  IpOwnerId *StringExpr `json:"IpOwnerId,omitempty"`  // String
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
  AttachmentID *StringExpr `json:"AttachmentID,omitempty"`  // String
  InstanceID *StringExpr `json:"InstanceID,omitempty"`  // String
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
  Key *StringExpr `json:"Key,omitempty"`  // String
  Value *StringExpr `json:"Value,omitempty"`  // String
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
  PrivateIpAddress *StringExpr `json:"PrivateIpAddress,omitempty"`  // String
  Primary *BoolExpr `json:"Primary,omitempty"`  // Boolean
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
  CidrIp *StringExpr `json:"CidrIp,omitempty"`  // String
  DestinationSecurityGroupIdXXSecurityGroupEgressXOnlyX *StringExpr `json:"DestinationSecurityGroupId (SecurityGroupEgress only),omitempty"`  // String
  FromPort *IntegerExpr `json:"FromPort,omitempty"`  // Integer
  IpProtocol *StringExpr `json:"IpProtocol,omitempty"`  // String
  SourceSecurityGroupIdXXSecurityGroupIngressXOnlyX *StringExpr `json:"SourceSecurityGroupId (SecurityGroupIngress only),omitempty"`  // String
  SourceSecurityGroupNameXXSecurityGroupIngressXOnlyX *StringExpr `json:"SourceSecurityGroupName (SecurityGroupIngress only),omitempty"`  // String
  SourceSecurityGroupOwnerIdXXSecurityGroupIngressXOnlyX *StringExpr `json:"SourceSecurityGroupOwnerId (SecurityGroupIngress only),omitempty"`  // String
  ToPort *IntegerExpr `json:"ToPort,omitempty"`  // Integer
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
  ContainerName *StringExpr `json:"ContainerName,omitempty"`  // String
  ContainerPort *IntegerExpr `json:"ContainerPort,omitempty"`  // Integer
  LoadBalancerName *StringExpr `json:"LoadBalancerName,omitempty"`  // String
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
  Command *StringListExpr `json:"Command,omitempty"`  // List of strings
  Cpu *IntegerExpr `json:"Cpu,omitempty"`  // Integer
  EntryPoint *StringListExpr `json:"EntryPoint,omitempty"`  // List of strings
  Environment *AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsEnvironmentList `json:"Environment,omitempty"`  // List of Amazon EC2 Container Service TaskDefinition ContainerDefinitions Environment
  Essential *BoolExpr `json:"Essential,omitempty"`  // Boolean
  Image *StringExpr `json:"Image,omitempty"`  // String
  Links *StringListExpr `json:"Links,omitempty"`  // List of strings
  Memory *IntegerExpr `json:"Memory,omitempty"`  // Integer
  MountPoints *AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsMountPointsList `json:"MountPoints,omitempty"`  // List of Amazon EC2 Container Service TaskDefinition ContainerDefinitions MountPoints
  Name *StringExpr `json:"Name,omitempty"`  // String
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
  Name *StringExpr `json:"Name,omitempty"`  // String
  Value *StringExpr `json:"Value,omitempty"`  // String
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
  ContainerPath *StringExpr `json:"ContainerPath,omitempty"`  // String
  SourceVolume *StringExpr `json:"SourceVolume,omitempty"`  // String
  ReadOnly *BoolExpr `json:"ReadOnly,omitempty"`  // Boolean
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
  ContainerPort *IntegerExpr `json:"ContainerPort,omitempty"`  // Integer
  HostPort *IntegerExpr `json:"HostPort,omitempty"`  // Integer
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
  SourceContainer *StringExpr `json:"SourceContainer,omitempty"`  // String
  ReadOnly *BoolExpr `json:"ReadOnly,omitempty"`  // Boolean
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
  Name *StringExpr `json:"Name,omitempty"`  // String
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
  SourcePath *StringExpr `json:"SourcePath,omitempty"`  // String
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
  EmitInterval *IntegerExpr `json:"EmitInterval,omitempty"`  // Integer
  Enabled *BoolExpr `json:"Enabled,omitempty"`  // Boolean
  S3BucketName *StringExpr `json:"S3BucketName,omitempty"`  // String
  S3BucketPrefix *StringExpr `json:"S3BucketPrefix,omitempty"`  // String
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
  CookieName *StringExpr `json:"CookieName,omitempty"`  // String
  PolicyName *StringExpr `json:"PolicyName,omitempty"`  // String
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
  Enabled *BoolExpr `json:"Enabled,omitempty"`  // Boolean
  Timeout *IntegerExpr `json:"Timeout,omitempty"`  // Integer
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
  IdleTimeout *IntegerExpr `json:"IdleTimeout,omitempty"`  // Integer
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
  HealthyThreshold *StringExpr `json:"HealthyThreshold,omitempty"`  // String
  Interval *StringExpr `json:"Interval,omitempty"`  // String
  Target *StringExpr `json:"Target,omitempty"`  // String
  Timeout *StringExpr `json:"Timeout,omitempty"`  // String
  UnhealthyThreshold *StringExpr `json:"UnhealthyThreshold,omitempty"`  // String
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
  CookieExpirationPeriod *StringExpr `json:"CookieExpirationPeriod,omitempty"`  // String
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
  InstancePort *StringExpr `json:"InstancePort,omitempty"`  // String
  InstanceProtocol *StringExpr `json:"InstanceProtocol,omitempty"`  // String
  LoadBalancerPort *StringExpr `json:"LoadBalancerPort,omitempty"`  // String
  PolicyNames *StringListExpr `json:"PolicyNames,omitempty"`  // List of strings
  Protocol *StringExpr `json:"Protocol,omitempty"`  // String
  SSLCertificateId *StringExpr `json:"SSLCertificateId,omitempty"`  // String
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
  PolicyName *StringExpr `json:"PolicyName,omitempty"`  // String
  PolicyType *StringExpr `json:"PolicyType,omitempty"`  // String
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
  PolicyName *StringExpr `json:"PolicyName,omitempty"`  // String
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
  Password *StringExpr `json:"Password,omitempty"`  // String
  PasswordResetRequired *BoolExpr `json:"PasswordResetRequired,omitempty"`  // Boolean
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
  S3Bucket *StringExpr `json:"S3Bucket,omitempty"`  // String
  S3Key *StringExpr `json:"S3Key,omitempty"`  // String
  S3ObjectVersion *StringExpr `json:"S3ObjectVersion,omitempty"`  // String
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
  CpuThreshold *IntegerExpr `json:"CpuThreshold,omitempty"`  // Number
  IgnoreMetricsTime *IntegerExpr `json:"IgnoreMetricsTime,omitempty"`  // Integer
  InstanceCount *IntegerExpr `json:"InstanceCount,omitempty"`  // Integer
  LoadThreshold *IntegerExpr `json:"LoadThreshold,omitempty"`  // Number
  MemoryThreshold *IntegerExpr `json:"MemoryThreshold,omitempty"`  // Number
  ThresholdsWaitTime *IntegerExpr `json:"ThresholdsWaitTime,omitempty"`  // Integer
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
  BerkshelfVersion *StringExpr `json:"BerkshelfVersion,omitempty"`  // String
  ManageBerkshelf *BoolExpr `json:"ManageBerkshelf,omitempty"`  // Boolean
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
  DelayUntilElbConnectionsDrained *BoolExpr `json:"DelayUntilElbConnectionsDrained,omitempty"`  // Boolean
  ExecutionTimeout *IntegerExpr `json:"ExecutionTimeout,omitempty"`  // Integer
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
  Enable *BoolExpr `json:"Enable,omitempty"`  // Boolean
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
  Configure *StringListExpr `json:"Configure,omitempty"`  // List of strings
  Deploy *StringListExpr `json:"Deploy,omitempty"`  // List of strings
  Setup *StringListExpr `json:"Setup,omitempty"`  // List of strings
  Shutdown *StringListExpr `json:"Shutdown,omitempty"`  // List of strings
  Undeploy *StringListExpr `json:"Undeploy,omitempty"`  // List of strings
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
  Password *StringExpr `json:"Password,omitempty"`  // String
  Revision *StringExpr `json:"Revision,omitempty"`  // String
  SshKey *StringExpr `json:"SshKey,omitempty"`  // String
  Type *StringExpr `json:"Type,omitempty"`  // String
  Url *StringExpr `json:"Url,omitempty"`  // String
  Username *StringExpr `json:"Username,omitempty"`  // String
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
  Certificate *StringExpr `json:"Certificate,omitempty"`  // String
  Chain *StringExpr `json:"Chain,omitempty"`  // String
  PrivateKey *StringExpr `json:"PrivateKey,omitempty"`  // String
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
  Name *StringExpr `json:"Name,omitempty"`  // String
  Version *StringExpr `json:"Version,omitempty"`  // String
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
  Friday *StringExpr `json:"Friday,omitempty"`  // String to string map
  Monday *StringExpr `json:"Monday,omitempty"`  // String to string map
  Saturday *StringExpr `json:"Saturday,omitempty"`  // String to string map
  Sunday *StringExpr `json:"Sunday,omitempty"`  // String to string map
  Thursday *StringExpr `json:"Thursday,omitempty"`  // String to string map
  Tuesday *StringExpr `json:"Tuesday,omitempty"`  // String to string map
  Wednesday *StringExpr `json:"Wednesday,omitempty"`  // String to string map
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
  Iops *IntegerExpr `json:"Iops,omitempty"`  // Integer
  MountPoint *StringExpr `json:"MountPoint,omitempty"`  // String
  NumberOfDisks *IntegerExpr `json:"NumberOfDisks,omitempty"`  // Integer
  RaidLevel *IntegerExpr `json:"RaidLevel,omitempty"`  // Integer
  Size *IntegerExpr `json:"Size,omitempty"`  // Integer
  VolumeType *StringExpr `json:"VolumeType,omitempty"`  // String
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
  ParameterName *StringExpr `json:"ParameterName,omitempty"`  // String
  ParameterValue *StringExpr `json:"ParameterValue,omitempty"`  // String
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
  Key *StringExpr `json:"Key,omitempty"`  // String
  Value *StringExpr `json:"Value,omitempty"`  // String
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
  DBSecurityGroupMemberships *StringListExpr `json:"DBSecurityGroupMemberships,omitempty"`  // List of strings
  OptionName *StringExpr `json:"OptionName,omitempty"`  // String
  OptionSettings *AmazonRDSOptionGroupOptionConfigurationsOptionSettings `json:"OptionSettings,omitempty"`  // Amazon RDS OptionGroup OptionConfigurations OptionSettings
  Port *IntegerExpr `json:"Port,omitempty"`  // Integer
  VpcSecurityGroupMemberships *StringListExpr `json:"VpcSecurityGroupMemberships,omitempty"`  // List of strings
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
  Name *StringExpr `json:"Name,omitempty"`  // String
  Value *StringExpr `json:"Value,omitempty"`  // String
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
  CIDRIP *StringExpr `json:"CIDRIP,omitempty"`  // String
  EC2SecurityGroupId *StringExpr `json:"EC2SecurityGroupId,omitempty"`  // String
  EC2SecurityGroupName *StringExpr `json:"EC2SecurityGroupName,omitempty"`  // String
  EC2SecurityGroupOwnerId *StringExpr `json:"EC2SecurityGroupOwnerId,omitempty"`  // String
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
  DNSName *StringExpr `json:"DNSName,omitempty"`  // String
  EvaluateTargetHealth *BoolExpr `json:"EvaluateTargetHealth,omitempty"`  // Boolean
  HostedZoneId *StringExpr `json:"HostedZoneId,omitempty"`  // String
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
  ContinentCode *StringExpr `json:"ContinentCode,omitempty"`  // String
  CountryCode *StringExpr `json:"CountryCode,omitempty"`  // String
  SubdivisionCode *StringExpr `json:"SubdivisionCode,omitempty"`  // String
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
  FailureThreshold *IntegerExpr `json:"FailureThreshold,omitempty"`  // Integer
  FullyQualifiedDomainName *StringExpr `json:"FullyQualifiedDomainName,omitempty"`  // String
  IPAddress *StringExpr `json:"IPAddress,omitempty"`  // String
  Port *IntegerExpr `json:"Port,omitempty"`  // Integer
  RequestInterval *IntegerExpr `json:"RequestInterval,omitempty"`  // Integer
  ResourcePath *StringExpr `json:"ResourcePath,omitempty"`  // String
  SearchString *StringExpr `json:"SearchString,omitempty"`  // String
  Type *StringExpr `json:"Type,omitempty"`  // String
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
  Key *StringExpr `json:"Key,omitempty"`  // String
  Value *StringExpr `json:"Value,omitempty"`  // String
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
  Comment *StringExpr `json:"Comment,omitempty"`  // String
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
  Key *StringExpr `json:"Key,omitempty"`  // String
  Value *StringExpr `json:"Value,omitempty"`  // String
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
  VPCId *StringExpr `json:"VPCId,omitempty"`  // String
  VPCRegion *StringExpr `json:"VPCRegion,omitempty"`  // String
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
  AllowedHeaders *StringListExpr `json:"AllowedHeaders,omitempty"`  // List of strings
  AllowedMethods *StringListExpr `json:"AllowedMethods,omitempty"`  // List of strings
  AllowedOrigins *StringListExpr `json:"AllowedOrigins,omitempty"`  // List of strings
  ExposedHeaders *StringListExpr `json:"ExposedHeaders,omitempty"`  // List of strings
  Id *StringExpr `json:"Id,omitempty"`  // String
  MaxAge *IntegerExpr `json:"MaxAge,omitempty"`  // Integer
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
  ExpirationDate *StringExpr `json:"ExpirationDate,omitempty"`  // String
  ExpirationInDays *IntegerExpr `json:"ExpirationInDays,omitempty"`  // Integer
  Id *StringExpr `json:"Id,omitempty"`  // String
  NoncurrentVersionExpirationInDays *IntegerExpr `json:"NoncurrentVersionExpirationInDays,omitempty"`  // Integer
  NoncurrentVersionTransition *AmazonS3LifecycleRuleNoncurrentVersionTransition `json:"NoncurrentVersionTransition,omitempty"`  // Amazon S3 Lifecycle Rule NoncurrentVersionTransition
  Prefix *StringExpr `json:"Prefix,omitempty"`  // String
  Status *StringExpr `json:"Status,omitempty"`  // String
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
  StorageClass *StringExpr `json:"StorageClass,omitempty"`  // String
  TransitionInDays *IntegerExpr `json:"TransitionInDays,omitempty"`  // Integer
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
  StorageClass *StringExpr `json:"StorageClass,omitempty"`  // String
  TransitionDate *StringExpr `json:"TransitionDate,omitempty"`  // String
  TransitionInDays *IntegerExpr `json:"TransitionInDays,omitempty"`  // Integer
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
  DestinationBucketName *StringExpr `json:"DestinationBucketName,omitempty"`  // String
  LogFilePrefix *StringExpr `json:"LogFilePrefix,omitempty"`  // String
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
  Event *StringExpr `json:"Event,omitempty"`  // String
  Topic *StringExpr `json:"Topic,omitempty"`  // String
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
  Status *StringExpr `json:"Status,omitempty"`  // String
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
  ErrorDocument *StringExpr `json:"ErrorDocument,omitempty"`  // String
  IndexDocument *StringExpr `json:"IndexDocument,omitempty"`  // String
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
  HostName *StringExpr `json:"HostName,omitempty"`  // String
  Protocol *StringExpr `json:"Protocol,omitempty"`  // String
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
  HostName *StringExpr `json:"HostName,omitempty"`  // String
  HttpRedirectCode *StringExpr `json:"HttpRedirectCode,omitempty"`  // String
  Protocol *StringExpr `json:"Protocol,omitempty"`  // String
  ReplaceKeyPrefixWith *StringExpr `json:"ReplaceKeyPrefixWith,omitempty"`  // String
  ReplaceKeyWith *StringExpr `json:"ReplaceKeyWith,omitempty"`  // String
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
  HttpErrorCodeReturnedEquals *StringExpr `json:"HttpErrorCodeReturnedEquals,omitempty"`  // String
  KeyPrefixEquals *StringExpr `json:"KeyPrefixEquals,omitempty"`  // String
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
  Endpoint *StringExpr `json:"Endpoint,omitempty"`  // String
  Protocol *StringExpr `json:"Protocol,omitempty"`  // String
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
  DeadLetterTargetArn *StringExpr `json:"deadLetterTargetArn,omitempty"`  // String
  MaxReceiveCount *IntegerExpr `json:"maxReceiveCount,omitempty"`  // Integer
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

func NewResourceByType(typeName string) ResourceProperties {
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
