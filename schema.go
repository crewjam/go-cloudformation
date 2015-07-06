//go:generate go run ./scraper/scrape.go -format=go -out=schema.go

package cloudformation

import "time"

type AWSAutoScalingAutoScalingGroup struct {
  AvailabilityZones []string `json:"AvailabilityZones"`
  Cooldown string `json:"Cooldown"`
  DesiredCapacity string `json:"DesiredCapacity"`
  HealthCheckGracePeriod int `json:"HealthCheckGracePeriod"`
  HealthCheckType string `json:"HealthCheckType"`
  InstanceId string `json:"InstanceId"`
  LaunchConfigurationName string `json:"LaunchConfigurationName"`
  LoadBalancerNames []string `json:"LoadBalancerNames"`
  MaxSize string `json:"MaxSize"`
  MetricsCollection []AutoScalingMetricsCollection `json:"MetricsCollection"`
  MinSize string `json:"MinSize"`
  NotificationConfigurations []AutoScalingNotificationConfigurations `json:"NotificationConfigurations"`
  PlacementGroup string `json:"PlacementGroup"`
  Tags []AutoScalingTags `json:"Tags"`
  TerminationPolicies []string `json:"TerminationPolicies"`
  VPCZoneIdentifier []string `json:"VPCZoneIdentifier"`
}

type AWSAutoScalingLaunchConfiguration struct {
  AssociatePublicIpAddress bool `json:"AssociatePublicIpAddress"`
  BlockDeviceMappings []AWSCloudFormationAutoScalingBlockDeviceMapping `json:"BlockDeviceMappings"`
  ClassicLinkVPCId string `json:"ClassicLinkVPCId"`
  ClassicLinkVPCSecurityGroups []string `json:"ClassicLinkVPCSecurityGroups"`
  EbsOptimized bool `json:"EbsOptimized"`
  IamInstanceProfile string `json:"IamInstanceProfile"`
  ImageId string `json:"ImageId"`
  InstanceId string `json:"InstanceId"`
  InstanceMonitoring bool `json:"InstanceMonitoring"`
  InstanceType string `json:"InstanceType"`
  KernelId string `json:"KernelId"`
  KeyName string `json:"KeyName"`
  PlacementTenancy string `json:"PlacementTenancy"`
  RamDiskId string `json:"RamDiskId"`
  SecurityGroups interface{} `json:"SecurityGroups"`
  SpotPrice string `json:"SpotPrice"`
  UserData string `json:"UserData"`
}

type AWSAutoScalingLifecycleHook struct {
}

type AWSAutoScalingScalingPolicy struct {
  AdjustmentType string `json:"AdjustmentType"`
  AutoScalingGroupName string `json:"AutoScalingGroupName"`
  Cooldown string `json:"Cooldown"`
  MinAdjustmentStep int `json:"MinAdjustmentStep"`
  ScalingAdjustment string `json:"ScalingAdjustment"`
}

type AWSAutoScalingScheduledAction struct {
  AutoScalingGroupName string `json:"AutoScalingGroupName"`
  DesiredCapacity int `json:"DesiredCapacity"`
  EndTime time.Time `json:"EndTime"`
  MaxSize int `json:"MaxSize"`
  MinSize int `json:"MinSize"`
  Recurrence string `json:"Recurrence"`
  StartTime time.Time `json:"StartTime"`
}

type AWSCloudFormationAuthentication struct {
  AccessKeyId string `json:"accessKeyId"`
  Buckets []string `json:"buckets"`
  Password string `json:"password"`
  SecretKey string `json:"secretKey"`
  Type string `json:"type"`
  Uris []string `json:"uris"`
  Username string `json:"username"`
  RoleName string `json:"roleName"`
}

type AWSCloudFormationCustomResource struct {
  ServiceToken string `json:"ServiceToken"`
}

type AWSCloudFormationInit struct {
}

type AWSCloudFormationStack struct {
  NotificationARNs []string `json:"NotificationARNs"`
  Parameters CloudFormationStackParameters `json:"Parameters"`
  TemplateURL string `json:"TemplateURL"`
  TimeoutInMinutes string `json:"TimeoutInMinutes"`
}

type AWSCloudFormationWaitCondition struct {
  Count string `json:"Count"`
  Handle string `json:"Handle"`
  Timeout string `json:"Timeout"`
}

type AWSCloudFormationWaitConditionHandle struct {
}

type AWSCloudFrontDistribution struct {
  DistributionConfig CloudFrontDistributionConfig `json:"DistributionConfig"`
}

type AWSCloudTrailTrail struct {
  IncludeGlobalServiceEvents bool `json:"IncludeGlobalServiceEvents"`
  IsLogging bool `json:"IsLogging"`
  S3BucketName string `json:"S3BucketName"`
  S3KeyPrefix string `json:"S3KeyPrefix"`
  SnsTopicName string `json:"SnsTopicName"`
}

type AWSCloudWatchAlarm struct {
  ActionsEnabled bool `json:"ActionsEnabled"`
  AlarmActions []string `json:"AlarmActions"`
  AlarmDescription string `json:"AlarmDescription"`
  AlarmName string `json:"AlarmName"`
  ComparisonOperator string `json:"ComparisonOperator"`
  Dimensions []CloudWatchMetricDimension `json:"Dimensions"`
  EvaluationPeriods string `json:"EvaluationPeriods"`
  InsufficientDataActions []string `json:"InsufficientDataActions"`
  MetricName string `json:"MetricName"`
  Namespace string `json:"Namespace"`
  OKActions []string `json:"OKActions"`
  Period string `json:"Period"`
  Statistic string `json:"Statistic"`
  Threshold string `json:"Threshold"`
  Unit string `json:"Unit"`
}

type AWSDataPipelinePipeline struct {
  Activate bool `json:"Activate"`
  Description string `json:"Description"`
  Name string `json:"Name"`
  ParameterObjects AWSDataPipelinePipelineParameterObjects `json:"ParameterObjects"`
  ParameterValues AWSDataPipelinePipelineParameterValues `json:"ParameterValues"`
  PipelineObjects []AWSDataPipelinePipelineObjects `json:"PipelineObjects"`
  PipelineTags AWSDataPipelinePipelinePipelineTags `json:"PipelineTags"`
}

type AWSDynamoDBTable struct {
  AttributeDefinitions DynamoDBAttributeDefinitions `json:"AttributeDefinitions"`
  GlobalSecondaryIndexes DynamoDBGlobalSecondaryIndexes `json:"GlobalSecondaryIndexes"`
  KeySchema DynamoDBKeySchema `json:"KeySchema"`
  LocalSecondaryIndexes DynamoDBLocalSecondaryIndexes `json:"LocalSecondaryIndexes"`
  ProvisionedThroughput DynamoDBProvisionedThroughput `json:"ProvisionedThroughput"`
  TableName string `json:"TableName"`
}

type AWSEC2CustomerGateway struct {
  BgpAsn int `json:"BgpAsn"`
  IpAddress string `json:"IpAddress"`
  Tags AWSCloudFormationResourceTags `json:"Tags"`
  Type string `json:"Type"`
}

type AWSEC2DHCPOptions struct {
  DomainName string `json:"DomainName"`
  DomainNameServers []string `json:"DomainNameServers"`
  NetbiosNameServers []string `json:"NetbiosNameServers"`
  NetbiosNodeType interface{} `json:"NetbiosNodeType"`
  NtpServers []string `json:"NtpServers"`
  Tags AWSCloudFormationResourceTags `json:"Tags"`
}

type AWSEC2EIP struct {
  InstanceId string `json:"InstanceId"`
  Domain string `json:"Domain"`
}

type AWSEC2EIPAssociation struct {
  AllocationId string `json:"AllocationId"`
  EIP string `json:"EIP"`
  InstanceId string `json:"InstanceId"`
  NetworkInterfaceId string `json:"NetworkInterfaceId"`
  PrivateIpAddress string `json:"PrivateIpAddress"`
}

type AWSEC2Instance struct {
  AvailabilityZone string `json:"AvailabilityZone"`
  BlockDeviceMappings []AmazonEC2BlockDeviceMappingProperty `json:"BlockDeviceMappings"`
  DisableApiTermination bool `json:"DisableApiTermination"`
  EbsOptimized bool `json:"EbsOptimized"`
  IamInstanceProfile string `json:"IamInstanceProfile"`
  ImageId string `json:"ImageId"`
  InstanceInitiatedShutdownBehavior string `json:"InstanceInitiatedShutdownBehavior"`
  InstanceType string `json:"InstanceType"`
  KernelId string `json:"KernelId"`
  KeyName string `json:"KeyName"`
  Monitoring bool `json:"Monitoring"`
  NetworkInterfaces []EC2NetworkInterfaceEmbedded `json:"NetworkInterfaces"`
  PlacementGroupName string `json:"PlacementGroupName"`
  PrivateIpAddress string `json:"PrivateIpAddress"`
  RamdiskId string `json:"RamdiskId"`
  SecurityGroupIds []string `json:"SecurityGroupIds"`
  SecurityGroups []string `json:"SecurityGroups"`
  SourceDestCheck bool `json:"SourceDestCheck"`
  SubnetId string `json:"SubnetId"`
  Tags AWSCloudFormationResourceTags `json:"Tags"`
  Tenancy string `json:"Tenancy"`
  UserData string `json:"UserData"`
  Volumes []EC2MountPoint `json:"Volumes"`
}

type AWSEC2InternetGateway struct {
  Tags AWSCloudFormationResourceTags `json:"Tags"`
}

type AWSEC2NetworkAcl struct {
  Tags AWSCloudFormationResourceTags `json:"Tags"`
  VpcId string `json:"VpcId"`
}

type AWSEC2NetworkAclEntry struct {
  CidrBlock string `json:"CidrBlock"`
  Egress bool `json:"Egress"`
  Icmp EC2ICMP `json:"Icmp"`
  NetworkAclId string `json:"NetworkAclId"`
  PortRange EC2PortRange `json:"PortRange"`
  Protocol int `json:"Protocol"`
  RuleAction string `json:"RuleAction"`
  RuleNumber int `json:"RuleNumber"`
}

type AWSEC2NetworkInterface struct {
  Description string `json:"Description"`
  GroupSet []string `json:"GroupSet"`
  PrivateIpAddress string `json:"PrivateIpAddress"`
  PrivateIpAddresses []EC2NetworkInterfacePrivateIPSpecification `json:"PrivateIpAddresses"`
  SecondaryPrivateIpAddressCount int `json:"SecondaryPrivateIpAddressCount"`
  SourceDestCheck bool `json:"SourceDestCheck"`
  SubnetId string `json:"SubnetId"`
  Tags AWSCloudFormationResourceTags `json:"Tags"`
}

type AWSEC2NetworkInterfaceAttachment struct {
  DeleteOnTermination bool `json:"DeleteOnTermination"`
  DeviceIndex string `json:"DeviceIndex"`
  InstanceId string `json:"InstanceId"`
  NetworkInterfaceId string `json:"NetworkInterfaceId"`
}

type AWSEC2Route struct {
  DestinationCidrBlock string `json:"DestinationCidrBlock"`
  GatewayId string `json:"GatewayId"`
  InstanceId string `json:"InstanceId"`
  NetworkInterfaceId string `json:"NetworkInterfaceId"`
  RouteTableId string `json:"RouteTableId"`
}

type AWSEC2RouteTable struct {
  VpcId string `json:"VpcId"`
  Tags AWSCloudFormationResourceTags `json:"Tags"`
}

type AWSEC2SecurityGroup struct {
  GroupDescription string `json:"GroupDescription"`
  SecurityGroupEgress EC2SecurityGroupRule `json:"SecurityGroupEgress"`
  SecurityGroupIngress EC2SecurityGroupRule `json:"SecurityGroupIngress"`
  Tags AWSCloudFormationResourceTags `json:"Tags"`
  VpcId string `json:"VpcId"`
}

type AWSEC2SecurityGroupEgress struct {
}

type AWSEC2SecurityGroupIngress struct {
}

type AWSEC2Subnet struct {
  AvailabilityZone string `json:"AvailabilityZone"`
  CidrBlock string `json:"CidrBlock"`
  MapPublicIpOnLaunch bool `json:"MapPublicIpOnLaunch"`
  Tags AWSCloudFormationResourceTags `json:"Tags"`
  VpcId interface{} `json:"VpcId"`
}

type AWSEC2SubnetNetworkAclAssociation struct {
  SubnetId string `json:"SubnetId"`
  NetworkAclId string `json:"NetworkAclId"`
}

type AWSEC2SubnetRouteTableAssociation struct {
  RouteTableId string `json:"RouteTableId"`
  SubnetId string `json:"SubnetId"`
}

type AWSEC2Volume struct {
  AvailabilityZone string `json:"AvailabilityZone"`
  Encrypted bool `json:"Encrypted"`
  Iops int `json:"Iops"`
  KmsKeyId string `json:"KmsKeyId"`
  Size string `json:"Size"`
  SnapshotId string `json:"SnapshotId"`
  Tags AWSCloudFormationResourceTags `json:"Tags"`
  VolumeType string `json:"VolumeType"`
}

type AWSEC2VolumeAttachment struct {
  Device string `json:"Device"`
  InstanceId string `json:"InstanceId"`
  VolumeId string `json:"VolumeId"`
}

type AWSEC2VPC struct {
  CidrBlock string `json:"CidrBlock"`
  EnableDnsSupport bool `json:"EnableDnsSupport"`
  EnableDnsHostnames bool `json:"EnableDnsHostnames"`
  InstanceTenancy string `json:"InstanceTenancy"`
  Tags AWSCloudFormationResourceTags `json:"Tags"`
}

type AWSEC2VPCDHCPOptionsAssociation struct {
  DhcpOptionsId string `json:"DhcpOptionsId"`
  VpcId string `json:"VpcId"`
}

type AWSEC2VPCGatewayAttachment struct {
  InternetGatewayId string `json:"InternetGatewayId"`
  VpcId string `json:"VpcId"`
  VpnGatewayId string `json:"VpnGatewayId"`
}

type AWSEC2VPCPeeringConnection struct {
  PeerVpcId string `json:"PeerVpcId"`
  Tags AWSCloudFormationResourceTags `json:"Tags"`
  VpcId string `json:"VpcId"`
}

type AWSEC2VPNConnection struct {
  Type string `json:"Type"`
  CustomerGatewayId string `json:"CustomerGatewayId"`
  StaticRoutesOnly bool `json:"StaticRoutesOnly"`
  Tags AWSCloudFormationResourceTags `json:"Tags"`
  VpnGatewayId string `json:"VpnGatewayId"`
}

type AWSEC2VPNConnectionRoute struct {
  DestinationCidrBlock string `json:"DestinationCidrBlock"`
  VpnConnectionId string `json:"VpnConnectionId"`
}

type AWSEC2VPNGateway struct {
  Type string `json:"Type"`
  Tags AWSCloudFormationResourceTags `json:"Tags"`
}

type AWSEC2VPNGatewayRoutePropagation struct {
  RouteTableIds interface{} `json:"RouteTableIds"`
  VpnGatewayId string `json:"VpnGatewayId"`
}

type AWSECSCluster struct {
}

type AWSECSService struct {
}

type AWSECSTaskDefinition struct {
  ContainerDefinitions []AmazonEC2ContainerServiceTaskDefinitionContainerDefinitions `json:"ContainerDefinitions"`
  Volumes []AmazonEC2ContainerServiceTaskDefinitionVolumes `json:"Volumes"`
}

type AWSElastiCacheCacheCluster struct {
}

type AWSElastiCacheParameterGroup struct {
  CacheParameterGroupFamily string `json:"CacheParameterGroupFamily"`
  Description string `json:"Description"`
  Properties interface{} `json:"Properties"`
}

type AWSElastiCacheReplicationGroup struct {
}

type AWSElastiCacheSecurityGroup struct {
  Description string `json:"Description"`
}

type AWSElastiCacheSecurityGroupIngress struct {
  CacheSecurityGroupName string `json:"CacheSecurityGroupName"`
  EC2SecurityGroupName string `json:"EC2SecurityGroupName"`
  EC2SecurityGroupOwnerId string `json:"EC2SecurityGroupOwnerId"`
}

type AWSElastiCacheSubnetGroup struct {
}

type AWSElasticBeanstalkApplication struct {
  ApplicationName string `json:"ApplicationName"`
  Description string `json:"Description"`
}

type AWSElasticBeanstalkApplicationVersion struct {
}

type AWSElasticBeanstalkConfigurationTemplate struct {
}

type AWSElasticBeanstalkEnvironment struct {
  ApplicationName string `json:"ApplicationName"`
  CNAMEPrefix string `json:"CNAMEPrefix"`
  Description string `json:"Description"`
  EnvironmentName string `json:"EnvironmentName"`
  OptionSettings []ElasticBeanstalkOptionSettings `json:"OptionSettings"`
  SolutionStackName string `json:"SolutionStackName"`
  TemplateName string `json:"TemplateName"`
  Tier ElasticBeanstalkEnvironmentTier `json:"Tier"`
  VersionLabel string `json:"VersionLabel"`
}

type AWSElasticLoadBalancingLoadBalancer struct {
  AccessLoggingPolicy ElasticLoadBalancingAccessLoggingPolicy `json:"AccessLoggingPolicy"`
  AppCookieStickinessPolicy []ElasticLoadBalancingAppCookieStickinessPolicy `json:"AppCookieStickinessPolicy"`
  AvailabilityZones []string `json:"AvailabilityZones"`
  ConnectionDrainingPolicy ElasticLoadBalancingConnectionDrainingPolicy `json:"ConnectionDrainingPolicy"`
  ConnectionSettings ElasticLoadBalancingConnectionSettings `json:"ConnectionSettings"`
  CrossZone bool `json:"CrossZone"`
  HealthCheck ElasticLoadBalancingHealthCheck `json:"HealthCheck"`
  Instances []string `json:"Instances"`
  LBCookieStickinessPolicy []ElasticLoadBalancingLBCookieStickinessPolicy `json:"LBCookieStickinessPolicy"`
  LoadBalancerName string `json:"LoadBalancerName"`
  Listeners []ElasticLoadBalancingListener `json:"Listeners"`
  Policies []ElasticLoadBalancingPolicy `json:"Policies"`
  Scheme string `json:"Scheme"`
  SecurityGroups interface{} `json:"SecurityGroups"`
  Subnets []string `json:"Subnets"`
  Tags AWSCloudFormationResourceTags `json:"Tags"`
}

type AWSIAMAccessKey struct {
  Serial int `json:"Serial"`
  Status string `json:"Status"`
  UserName string `json:"UserName"`
}

type AWSIAMGroup struct {
  ManagedPolicyArns []string `json:"ManagedPolicyArns"`
  Path string `json:"Path"`
  Policies []IAMPolicies `json:"Policies"`
}

type AWSIAMInstanceProfile struct {
  Path string `json:"Path"`
  Roles interface{} `json:"Roles"`
}

type AWSIAMManagedPolicy struct {
  Description string `json:"Description"`
  Groups []string `json:"Groups"`
  Path string `json:"Path"`
  PolicyDocument interface{} `json:"PolicyDocument"`
  Roles []string `json:"Roles"`
  Users []string `json:"Users"`
}

type AWSIAMPolicy struct {
  Groups []string `json:"Groups"`
  PolicyDocument interface{} `json:"PolicyDocument"`
  PolicyName string `json:"PolicyName"`
  Roles []string `json:"Roles"`
  Users []string `json:"Users"`
}

type AWSIAMRole struct {
  AssumeRolePolicyDocument interface{} `json:"AssumeRolePolicyDocument"`
  ManagedPolicyArns []string `json:"ManagedPolicyArns"`
  Path string `json:"Path"`
  Policies []IAMPolicies `json:"Policies"`
}

type AWSIAMUser struct {
  Groups []string `json:"Groups"`
  LoginProfile IAMUserLoginProfile `json:"LoginProfile"`
  ManagedPolicyArns []string `json:"ManagedPolicyArns"`
  Path string `json:"Path"`
  Policies []IAMPolicies `json:"Policies"`
}

type AWSIAMUserToGroupAddition struct {
  GroupName string `json:"GroupName"`
  Users interface{} `json:"Users"`
}

type AWSKinesisStream struct {
  ShardCount int `json:"ShardCount"`
}

type AWSLambdaFunction struct {
  Code AWSLambdaFunctionCode `json:"Code"`
  Description string `json:"Description"`
  Handler string `json:"Handler"`
  MemorySize int `json:"MemorySize"`
  Role string `json:"Role"`
  Runtime string `json:"Runtime"`
  Timeout int `json:"Timeout"`
}

type AWSLogsLogGroup struct {
  RetentionInDays int `json:"RetentionInDays"`
}

type AWSLogsMetricFilter struct {
}

type AWSOpsWorksApp struct {
  AppSource AWSOpsWorksSource `json:"AppSource"`
  Attributes interface{} `json:"Attributes"`
  Description string `json:"Description"`
  Domains []string `json:"Domains"`
  EnableSsl bool `json:"EnableSsl"`
  Name string `json:"Name"`
  Shortname string `json:"Shortname"`
  SslConfiguration AWSOpsWorksSslConfiguration `json:"SslConfiguration"`
  StackId string `json:"StackId"`
  Type string `json:"Type"`
}

type AWSOpsWorksElasticLoadBalancerAttachment struct {
  ElasticLoadBalancerName string `json:"ElasticLoadBalancerName"`
  LayerId string `json:"LayerId"`
}

type AWSOpsWorksInstance struct {
  AmiId string `json:"AmiId"`
  Architecture string `json:"Architecture"`
  AutoScalingType string `json:"AutoScalingType"`
  AvailabilityZone string `json:"AvailabilityZone"`
  InstallUpdatesOnBoot bool `json:"InstallUpdatesOnBoot"`
  InstanceType string `json:"InstanceType"`
  LayerIds []string `json:"LayerIds"`
  Os string `json:"Os"`
  RootDeviceType string `json:"RootDeviceType"`
  SshKeyName string `json:"SshKeyName"`
  StackId string `json:"StackId"`
  SubnetId string `json:"SubnetId"`
  TimeBasedAutoScaling AWSOpsWorksTimeBasedAutoScaling `json:"TimeBasedAutoScaling"`
}

type AWSOpsWorksLayer struct {
  Attributes interface{} `json:"Attributes"`
  AutoAssignElasticIps bool `json:"AutoAssignElasticIps"`
  AutoAssignPublicIps bool `json:"AutoAssignPublicIps"`
  CustomInstanceProfileArn string `json:"CustomInstanceProfileArn"`
  CustomRecipes AWSOpsWorksRecipes `json:"CustomRecipes"`
  CustomSecurityGroupIds []string `json:"CustomSecurityGroupIds"`
  EnableAutoHealing bool `json:"EnableAutoHealing"`
  InstallUpdatesOnBoot bool `json:"InstallUpdatesOnBoot"`
  LifecycleEventConfiguration AWSOpsWorksLayerLifeCycleConfiguration `json:"LifecycleEventConfiguration"`
  LoadBasedAutoScaling AWSOpsWorksLoadBasedAutoScaling `json:"LoadBasedAutoScaling"`
  Name string `json:"Name"`
  Packages []string `json:"Packages"`
  Shortname string `json:"Shortname"`
  StackId string `json:"StackId"`
  Type string `json:"Type"`
  VolumeConfigurations []AWSOpsWorksVolumeConfiguration `json:"VolumeConfigurations"`
}

type AWSOpsWorksStack struct {
  Attributes interface{} `json:"Attributes"`
  ChefConfiguration AWSOpsWorksChefConfiguration `json:"ChefConfiguration"`
  ConfigurationManager AWSOpsWorksStackConfigurationManager `json:"ConfigurationManager"`
  CustomCookbooksSource AWSOpsWorksSource `json:"CustomCookbooksSource"`
  CustomJson interface{} `json:"CustomJson"`
  DefaultAvailabilityZone string `json:"DefaultAvailabilityZone"`
  DefaultInstanceProfileArn string `json:"DefaultInstanceProfileArn"`
  DefaultOs string `json:"DefaultOs"`
  DefaultRootDeviceType string `json:"DefaultRootDeviceType"`
  DefaultSshKeyName string `json:"DefaultSshKeyName"`
  DefaultSubnetId string `json:"DefaultSubnetId"`
  HostnameTheme string `json:"HostnameTheme"`
  Name string `json:"Name"`
  ServiceRoleArn string `json:"ServiceRoleArn"`
  UseCustomCookbooks bool `json:"UseCustomCookbooks"`
  UseOpsworksSecurityGroups bool `json:"UseOpsworksSecurityGroups"`
  VpcId string `json:"VpcId"`
}

type AWSRedshiftCluster struct {
  AllowVersionUpgrade bool `json:"AllowVersionUpgrade"`
  AutomatedSnapshotRetentionPeriod int `json:"AutomatedSnapshotRetentionPeriod"`
  AvailabilityZone string `json:"AvailabilityZone"`
  ClusterParameterGroupName string `json:"ClusterParameterGroupName"`
  ClusterSecurityGroups []string `json:"ClusterSecurityGroups"`
  ClusterSubnetGroupName string `json:"ClusterSubnetGroupName"`
  ClusterType string `json:"ClusterType"`
  ClusterVersion string `json:"ClusterVersion"`
  DBName string `json:"DBName"`
  ElasticIp string `json:"ElasticIp"`
  Encrypted bool `json:"Encrypted"`
  HsmClientCertificateIdentifier string `json:"HsmClientCertificateIdentifier"`
  HsmConfigurationIdentifier string `json:"HsmConfigurationIdentifier"`
  MasterUsername string `json:"MasterUsername"`
  MasterUserPassword string `json:"MasterUserPassword"`
  NodeType string `json:"NodeType"`
  NumberOfNodes int `json:"NumberOfNodes"`
  OwnerAccount string `json:"OwnerAccount"`
  Port int `json:"Port"`
  PreferredMaintenanceWindow string `json:"PreferredMaintenanceWindow"`
  PubliclyAccessible bool `json:"PubliclyAccessible"`
  SnapshotClusterIdentifier interface{} `json:"SnapshotClusterIdentifier"`
  SnapshotIdentifier string `json:"SnapshotIdentifier"`
  VpcSecurityGroupIds []string `json:"VpcSecurityGroupIds"`
}

type AWSRedshiftClusterParameterGroup struct {
  Description string `json:"Description"`
  ParameterGroupFamily string `json:"ParameterGroupFamily"`
  Parameters AmazonRedshiftParameter `json:"Parameters"`
}

type AWSRedshiftClusterSecurityGroup struct {
  Description string `json:"Description"`
}

type AWSRedshiftClusterSecurityGroupIngress struct {
  ClusterSecurityGroupName string `json:"ClusterSecurityGroupName"`
  CIDRIP string `json:"CIDRIP"`
  EC2SecurityGroupName string `json:"EC2SecurityGroupName"`
  EC2SecurityGroupOwnerId string `json:"EC2SecurityGroupOwnerId"`
}

type AWSRedshiftClusterSubnetGroup struct {
  Description string `json:"Description"`
  SubnetIds []string `json:"SubnetIds"`
}

type AWSRDSDBInstance struct {
  AllocatedStorage string `json:"AllocatedStorage"`
  AllowMajorVersionUpgrade bool `json:"AllowMajorVersionUpgrade"`
  AutoMinorVersionUpgrade bool `json:"AutoMinorVersionUpgrade"`
  AvailabilityZone string `json:"AvailabilityZone"`
  BackupRetentionPeriod string `json:"BackupRetentionPeriod"`
  CharacterSetName string `json:"CharacterSetName"`
  DBInstanceClass string `json:"DBInstanceClass"`
  DBInstanceIdentifier string `json:"DBInstanceIdentifier"`
  DBName string `json:"DBName"`
  DBParameterGroupName string `json:"DBParameterGroupName"`
  DBSecurityGroups []string `json:"DBSecurityGroups"`
  DBSnapshotIdentifier string `json:"DBSnapshotIdentifier"`
  DBSubnetGroupName string `json:"DBSubnetGroupName"`
  Engine string `json:"Engine"`
  EngineVersion string `json:"EngineVersion"`
  Iops int `json:"Iops"`
  KmsKeyId string `json:"KmsKeyId"`
  LicenseModel string `json:"LicenseModel"`
  MasterUsername string `json:"MasterUsername"`
  MasterUserPassword string `json:"MasterUserPassword"`
  MultiAZ bool `json:"MultiAZ"`
  OptionGroupName string `json:"OptionGroupName"`
  Port string `json:"Port"`
  PreferredBackupWindow string `json:"PreferredBackupWindow"`
  PreferredMaintenanceWindow string `json:"PreferredMaintenanceWindow"`
  PubliclyAccessible bool `json:"PubliclyAccessible"`
  SourceDBInstanceIdentifier string `json:"SourceDBInstanceIdentifier"`
  StorageEncrypted bool `json:"StorageEncrypted"`
  StorageType string `json:"StorageType"`
  Tags AWSCloudFormationResourceTags `json:"Tags"`
  VPCSecurityGroups []string `json:"VPCSecurityGroups"`
}

type AWSRDSDBParameterGroup struct {
  Description interface{} `json:"Description"`
  Family interface{} `json:"Family"`
  Parameters interface{} `json:"Parameters"`
  Tags []AWSCloudFormationResourceTags `json:"Tags"`
}

type AWSRDSDBSubnetGroup struct {
  DBSubnetGroupDescription string `json:"DBSubnetGroupDescription"`
  SubnetIds []string `json:"SubnetIds"`
  Tags []AWSCloudFormationResourceTags `json:"Tags"`
}

type AWSRDSDBSecurityGroup struct {
  EC2VpcId string `json:"EC2VpcId"`
  DBSecurityGroupIngress []AmazonRDSSecurityGroupRule `json:"DBSecurityGroupIngress"`
  GroupDescription string `json:"GroupDescription"`
  Tags []AWSCloudFormationResourceTags `json:"Tags"`
}

type AWSRDSDBSecurityGroupIngress struct {
  CIDRIP string `json:"CIDRIP"`
  DBSecurityGroupName string `json:"DBSecurityGroupName"`
  EC2SecurityGroupId string `json:"EC2SecurityGroupId"`
  EC2SecurityGroupName string `json:"EC2SecurityGroupName"`
  EC2SecurityGroupOwnerId string `json:"EC2SecurityGroupOwnerId"`
}

type AWSRDSEventSubscription struct {
}

type AWSRDSOptionGroup struct {
  EngineName string `json:"EngineName"`
  MajorEngineVersion string `json:"MajorEngineVersion"`
  OptionGroupDescription string `json:"OptionGroupDescription"`
  OptionConfigurations AmazonRDSOptionGroupOptionConfigurations `json:"OptionConfigurations"`
  Tags AWSCloudFormationResourceTags `json:"Tags"`
}

type AWSRoute53HealthCheck struct {
  HealthCheckConfig AmazonRoute53HealthCheckConfig `json:"HealthCheckConfig"`
  HealthCheckTags []AmazonRoute53HealthCheckTags `json:"HealthCheckTags"`
}

type AWSRoute53HostedZone struct {
  HostedZoneConfig AmazonRoute53HostedZoneConfigProperty `json:"HostedZoneConfig"`
  HostedZoneTags []AmazonRoute53HostedZoneTags `json:"HostedZoneTags"`
  Name string `json:"Name"`
  VPCs []AmazonRoute53HostedZoneVPCs `json:"VPCs"`
}

type AWSRoute53RecordSet struct {
  AliasTarget Route53AliasTargetProperty `json:"AliasTarget"`
  Comment string `json:"Comment"`
  Failover string `json:"Failover"`
  GeoLocation AmazonRoute53RecordSetGeoLocationProperty `json:"GeoLocation"`
  HealthCheckId string `json:"HealthCheckId"`
  HostedZoneId string `json:"HostedZoneId"`
  HostedZoneName string `json:"HostedZoneName"`
  Name string `json:"Name"`
  Region interface{} `json:"Region"`
  ResourceRecords []string `json:"ResourceRecords"`
  SetIdentifier string `json:"SetIdentifier"`
  TTL string `json:"TTL"`
  Type string `json:"Type"`
  Weight int `json:"Weight"`
}

type AWSRoute53RecordSetGroup struct {
  HostedZoneId string `json:"HostedZoneId"`
  HostedZoneName string `json:"HostedZoneName"`
  RecordSets []AWSRoute53RecordSet `json:"RecordSets"`
  Comment string `json:"Comment"`
}

type AWSS3Bucket struct {
  AccessControl string `json:"AccessControl"`
  BucketName string `json:"BucketName"`
  CorsConfiguration AmazonS3CorsConfiguration `json:"CorsConfiguration"`
  LifecycleConfiguration AmazonS3LifecycleConfiguration `json:"LifecycleConfiguration"`
  LoggingConfiguration AmazonS3LoggingConfiguration `json:"LoggingConfiguration"`
  NotificationConfiguration AmazonS3NotificationConfiguration `json:"NotificationConfiguration"`
  Tags AWSCloudFormationResourceTags `json:"Tags"`
  VersioningConfiguration AmazonS3VersioningConfiguration `json:"VersioningConfiguration"`
  WebsiteConfiguration AmazonS3WebsiteConfigurationProperty `json:"WebsiteConfiguration"`
}

type AWSS3BucketPolicy struct {
  Bucket string `json:"Bucket"`
  PolicyDocument interface{} `json:"PolicyDocument"`
}

type AWSSDBDomain struct {
}

type AWSSNSTopic struct {
}

type AWSSNSTopicPolicy struct {
  PolicyDocument interface{} `json:"PolicyDocument"`
  Topics interface{} `json:"Topics"`
}

type AWSSQSQueue struct {
  DelaySeconds int `json:"DelaySeconds"`
  MaximumMessageSize int `json:"MaximumMessageSize"`
  MessageRetentionPeriod int `json:"MessageRetentionPeriod"`
  QueueName string `json:"QueueName"`
  ReceiveMessageWaitTimeSeconds int `json:"ReceiveMessageWaitTimeSeconds"`
  RedrivePolicy AmazonSQSRedrivePolicy `json:"RedrivePolicy"`
  VisibilityTimeout int `json:"VisibilityTimeout"`
}

type AWSSQSQueuePolicy struct {
  PolicyDocument interface{} `json:"PolicyDocument"`
  Queues []string `json:"Queues"`
}

type AWSCloudFormationAutoScalingBlockDeviceMapping struct {
}

type AWSCloudFormationAutoScalingEBSBlockDevice struct {
  DeleteOnTermination bool `json:"DeleteOnTermination"`
  Iops int `json:"Iops"`
  SnapshotId string `json:"SnapshotId"`
  VolumeSize int `json:"VolumeSize"`
  VolumeType string `json:"VolumeType"`
}

type AutoScalingMetricsCollection struct {
  Granularity string `json:"Granularity"`
  Metrics []string `json:"Metrics"`
}

type AutoScalingNotificationConfigurations struct {
  NotificationTypes []string `json:"NotificationTypes"`
  TopicARN string `json:"TopicARN"`
}

type AutoScalingTags struct {
  Key string `json:"Key"`
  Value string `json:"Value"`
  PropagateAtLaunch bool `json:"PropagateAtLaunch"`
}

type CloudFormationStackParameters struct {
}

type CloudFrontDistributionConfig struct {
}

type CloudFrontDistributionConfigCacheBehavior struct {
}

type CloudFrontDistributionConfigCustomErrorResponse struct {
}

type CloudFrontDefaultCacheBehavior struct {
}

type CloudFrontLogging struct {
}

type CloudFrontDistributionConfigOrigin struct {
}

type CloudFrontDistributionConfigOriginCustomOrigin struct {
}

type CloudFrontDistributionConfigOriginS3Origin struct {
}

type CloudFrontDistributionConfigurationRestrictions struct {
}

type CloudFrontDistributionConfigRestrictionsGeoRestriction struct {
}

type CloudFrontDistributionConfigurationViewerCertificate struct {
}

type CloudFrontForwardedValues struct {
}

type CloudFrontForwardedValuesCookies struct {
}

type CloudWatchMetricDimension struct {
  Name string `json:"Name"`
  Value string `json:"Value"`
}

type CloudWatchLogsMetricFilterMetricTransformationProperty struct {
}

type AWSDataPipelinePipelineParameterObjects struct {
  Attributes AWSDataPipelineParameterObjectsAttributes `json:"Attributes"`
  Id string `json:"Id"`
}

type AWSDataPipelineParameterObjectsAttributes struct {
  Key string `json:"Key"`
  StringValue string `json:"StringValue"`
}

type AWSDataPipelinePipelineParameterValues struct {
  Id string `json:"Id"`
  StringValue string `json:"StringValue"`
}

type AWSDataPipelinePipelineObjects struct {
  Fields AWSDataPipelineDataPipelineObjectFields `json:"Fields"`
  Id string `json:"Id"`
  Name string `json:"Name"`
}

type AWSDataPipelineDataPipelineObjectFields struct {
  Key string `json:"Key"`
  RefValue string `json:"RefValue"`
  StringValue string `json:"StringValue"`
}

type AWSDataPipelinePipelinePipelineTags struct {
  Key string `json:"Key"`
  Value string `json:"Value"`
}

type DynamoDBAttributeDefinitions struct {
  AttributeName string `json:"AttributeName"`
  AttributeType string `json:"AttributeType"`
}

type DynamoDBGlobalSecondaryIndexes struct {
  IndexName string `json:"IndexName"`
  KeySchema DynamoDBKeySchema `json:"KeySchema"`
  Projection DynamoDBProjectionObject `json:"Projection"`
  ProvisionedThroughput DynamoDBProvisionedThroughput `json:"ProvisionedThroughput"`
}

type DynamoDBKeySchema struct {
  AttributeName string `json:"AttributeName"`
  KeyType string `json:"KeyType"`
}

type DynamoDBLocalSecondaryIndexes struct {
  IndexName string `json:"IndexName"`
  KeySchema DynamoDBKeySchema `json:"KeySchema"`
  Projection DynamoDBProjectionObject `json:"Projection"`
}

type DynamoDBProjectionObject struct {
  NonKeyAttributes []string `json:"NonKeyAttributes"`
  ProjectionType string `json:"ProjectionType"`
  KEYSXONLY interface{} `json:"KEYS_ONLY"`
  INCLUDE interface{} `json:"INCLUDE"`
  ALL interface{} `json:"ALL"`
}

type DynamoDBProvisionedThroughput struct {
}

type AmazonEC2BlockDeviceMappingProperty struct {
  DeviceName string `json:"DeviceName"`
  Ebs AmazonElasticBlockStoreBlockDeviceProperty `json:"Ebs"`
  NoDevice interface{} `json:"NoDevice"`
  VirtualName string `json:"VirtualName"`
}

type AmazonElasticBlockStoreBlockDeviceProperty struct {
  DeleteOnTermination bool `json:"DeleteOnTermination"`
  Encrypted bool `json:"Encrypted"`
  Iops int `json:"Iops"`
  SnapshotId string `json:"SnapshotId"`
  VolumeSize string `json:"VolumeSize"`
  VolumeType string `json:"VolumeType"`
}

type EC2ICMP struct {
}

type EC2MountPoint struct {
  Device string `json:"Device"`
  VolumeId string `json:"VolumeId"`
}

type EC2NetworkInterfaceEmbedded struct {
  AssociatePublicIpAddress bool `json:"AssociatePublicIpAddress"`
  DeleteOnTermination bool `json:"DeleteOnTermination"`
  Description string `json:"Description"`
  DeviceIndex string `json:"DeviceIndex"`
  GroupSet []string `json:"GroupSet"`
  NetworkInterfaceId string `json:"NetworkInterfaceId"`
  PrivateIpAddress string `json:"PrivateIpAddress"`
  PrivateIpAddresses []EC2NetworkInterfacePrivateIPSpecification `json:"PrivateIpAddresses"`
  SecondaryPrivateIpAddressCount int `json:"SecondaryPrivateIpAddressCount"`
  SubnetId string `json:"SubnetId"`
}

type EC2NetworkInterfaceAssociation struct {
  AttachmentID string `json:"AttachmentID"`
  InstanceID string `json:"InstanceID"`
  PublicIp string `json:"PublicIp"`
  IpOwnerId string `json:"IpOwnerId"`
}

type EC2NetworkInterfaceAttachment struct {
  AttachmentID string `json:"AttachmentID"`
  InstanceID string `json:"InstanceID"`
}

type EC2NetworkInterfaceGroupItem struct {
  Key string `json:"Key"`
  Value string `json:"Value"`
}

type EC2NetworkInterfacePrivateIPSpecification struct {
  PrivateIpAddress string `json:"PrivateIpAddress"`
  Primary bool `json:"Primary"`
}

type EC2PortRange struct {
}

type EC2SecurityGroupRule struct {
  CidrIp string `json:"CidrIp"`
  DestinationSecurityGroupIdXXSecurityGroupEgressXOnlyX string `json:"DestinationSecurityGroupId (SecurityGroupEgress only)"`
  FromPort int `json:"FromPort"`
  IpProtocol string `json:"IpProtocol"`
  SourceSecurityGroupIdXXSecurityGroupIngressXOnlyX string `json:"SourceSecurityGroupId (SecurityGroupIngress only)"`
  SourceSecurityGroupNameXXSecurityGroupIngressXOnlyX string `json:"SourceSecurityGroupName (SecurityGroupIngress only)"`
  SourceSecurityGroupOwnerIdXXSecurityGroupIngressXOnlyX string `json:"SourceSecurityGroupOwnerId (SecurityGroupIngress only)"`
  ToPort int `json:"ToPort"`
}

type AmazonEC2ContainerServiceServiceLoadBalancers struct {
  ContainerName string `json:"ContainerName"`
  ContainerPort int `json:"ContainerPort"`
  LoadBalancerName string `json:"LoadBalancerName"`
}

type AmazonEC2ContainerServiceTaskDefinitionContainerDefinitions struct {
}

type AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsEnvironment struct {
}

type AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsMountPoints struct {
}

type AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsPortMappings struct {
}

type AmazonEC2ContainerServiceTaskDefinitionContainerDefinitionsVolumesFrom struct {
}

type AmazonEC2ContainerServiceTaskDefinitionVolumes struct {
}

type AmazonEC2ContainerServiceTaskDefinitionVolumesHost struct {
}

type ElasticBeanstalkEnvironmentTier struct {
}

type ElasticBeanstalkOptionSettings struct {
}

type ElasticBeanstalkSourceBundle struct {
}

type ElasticBeanstalkSourceConfiguration struct {
}

type ElasticLoadBalancingAccessLoggingPolicy struct {
  EmitInterval int `json:"EmitInterval"`
  Enabled bool `json:"Enabled"`
  S3BucketName string `json:"S3BucketName"`
  S3BucketPrefix string `json:"S3BucketPrefix"`
}

type ElasticLoadBalancingAppCookieStickinessPolicy struct {
  CookieName string `json:"CookieName"`
  PolicyName string `json:"PolicyName"`
}

type ElasticLoadBalancingConnectionDrainingPolicy struct {
  Enabled bool `json:"Enabled"`
  Timeout int `json:"Timeout"`
}

type ElasticLoadBalancingConnectionSettings struct {
  IdleTimeout int `json:"IdleTimeout"`
}

type ElasticLoadBalancingHealthCheck struct {
  HealthyThreshold string `json:"HealthyThreshold"`
  Interval string `json:"Interval"`
  Target string `json:"Target"`
  Timeout string `json:"Timeout"`
  UnhealthyThreshold string `json:"UnhealthyThreshold"`
}

type ElasticLoadBalancingLBCookieStickinessPolicy struct {
  CookieExpirationPeriod string `json:"CookieExpirationPeriod"`
  PolicyName interface{} `json:"PolicyName"`
}

type ElasticLoadBalancingListener struct {
  InstancePort string `json:"InstancePort"`
  InstanceProtocol string `json:"InstanceProtocol"`
  LoadBalancerPort string `json:"LoadBalancerPort"`
  PolicyNames []string `json:"PolicyNames"`
  Protocol string `json:"Protocol"`
  SSLCertificateId string `json:"SSLCertificateId"`
}

type ElasticLoadBalancingPolicy struct {
  Attributes interface{} `json:"Attributes"`
  InstancePorts interface{} `json:"InstancePorts"`
  LoadBalancerPorts interface{} `json:"LoadBalancerPorts"`
  PolicyName string `json:"PolicyName"`
  PolicyType string `json:"PolicyType"`
}

type IAMPolicies struct {
  PolicyDocument interface{} `json:"PolicyDocument"`
  PolicyName string `json:"PolicyName"`
}

type IAMUserLoginProfile struct {
  Password string `json:"Password"`
  PasswordResetRequired bool `json:"PasswordResetRequired"`
}

type AWSLambdaFunctionCode struct {
  S3Bucket string `json:"S3Bucket"`
  S3Key string `json:"S3Key"`
  S3ObjectVersion string `json:"S3ObjectVersion"`
}

type Name struct {
}

type AWSOpsWorksAutoScalingThresholds struct {
  CpuThreshold int `json:"CpuThreshold"`
  IgnoreMetricsTime int `json:"IgnoreMetricsTime"`
  InstanceCount int `json:"InstanceCount"`
  LoadThreshold int `json:"LoadThreshold"`
  MemoryThreshold int `json:"MemoryThreshold"`
  ThresholdsWaitTime int `json:"ThresholdsWaitTime"`
}

type AWSOpsWorksChefConfiguration struct {
  BerkshelfVersion string `json:"BerkshelfVersion"`
  ManageBerkshelf bool `json:"ManageBerkshelf"`
}

type AWSOpsWorksLayerLifeCycleConfiguration struct {
  ShutdownEventConfiguration AWSOpsWorksLayerLifeCycleConfigurationShutdownEventConfiguration `json:"ShutdownEventConfiguration"`
}

type AWSOpsWorksLayerLifeCycleConfigurationShutdownEventConfiguration struct {
  DelayUntilElbConnectionsDrained bool `json:"DelayUntilElbConnectionsDrained"`
  ExecutionTimeout int `json:"ExecutionTimeout"`
}

type AWSOpsWorksLoadBasedAutoScaling struct {
  DownScaling AWSOpsWorksAutoScalingThresholds `json:"DownScaling"`
  Enable bool `json:"Enable"`
  UpScaling AWSOpsWorksAutoScalingThresholds `json:"UpScaling"`
}

type AWSOpsWorksRecipes struct {
  Configure []string `json:"Configure"`
  Deploy []string `json:"Deploy"`
  Setup []string `json:"Setup"`
  Shutdown []string `json:"Shutdown"`
  Undeploy []string `json:"Undeploy"`
}

type AWSOpsWorksSource struct {
  Password string `json:"Password"`
  Revision string `json:"Revision"`
  SshKey string `json:"SshKey"`
  Type string `json:"Type"`
  Url string `json:"Url"`
  Username string `json:"Username"`
}

type AWSOpsWorksSslConfiguration struct {
  Certificate string `json:"Certificate"`
  Chain string `json:"Chain"`
  PrivateKey string `json:"PrivateKey"`
}

type AWSOpsWorksStackConfigurationManager struct {
  Name string `json:"Name"`
  Version string `json:"Version"`
}

type AWSOpsWorksTimeBasedAutoScaling struct {
}

type AWSOpsWorksVolumeConfiguration struct {
  Iops int `json:"Iops"`
  MountPoint string `json:"MountPoint"`
  NumberOfDisks int `json:"NumberOfDisks"`
  RaidLevel int `json:"RaidLevel"`
  Size int `json:"Size"`
  VolumeType string `json:"VolumeType"`
}

type AmazonRedshiftParameter struct {
  ParameterName string `json:"ParameterName"`
  ParameterValue string `json:"ParameterValue"`
}

type AWSCloudFormationResourceTags struct {
  Key string `json:"Key"`
  Value string `json:"Value"`
}

type AmazonRDSOptionGroupOptionConfigurations struct {
  DBSecurityGroupMemberships []string `json:"DBSecurityGroupMemberships"`
  OptionName string `json:"OptionName"`
  OptionSettings AmazonRDSOptionGroupOptionConfigurationsOptionSettings `json:"OptionSettings"`
  Port int `json:"Port"`
  VpcSecurityGroupMemberships []string `json:"VpcSecurityGroupMemberships"`
}

type AmazonRDSOptionGroupOptionConfigurationsOptionSettings struct {
}

type AmazonRDSSecurityGroupRule struct {
  CIDRIP string `json:"CIDRIP"`
  EC2SecurityGroupId string `json:"EC2SecurityGroupId"`
  EC2SecurityGroupName string `json:"EC2SecurityGroupName"`
  EC2SecurityGroupOwnerId string `json:"EC2SecurityGroupOwnerId"`
}

type Route53AliasTargetProperty struct {
  DNSName string `json:"DNSName"`
  EvaluateTargetHealth bool `json:"EvaluateTargetHealth"`
  HostedZoneId string `json:"HostedZoneId"`
}

type AmazonRoute53RecordSetGeoLocationProperty struct {
  ContinentCode string `json:"ContinentCode"`
  CountryCode string `json:"CountryCode"`
  SubdivisionCode string `json:"SubdivisionCode"`
}

type AmazonRoute53HealthCheckConfig struct {
  FailureThreshold int `json:"FailureThreshold"`
  FullyQualifiedDomainName string `json:"FullyQualifiedDomainName"`
  IPAddress string `json:"IPAddress"`
  Port int `json:"Port"`
  RequestInterval int `json:"RequestInterval"`
  ResourcePath string `json:"ResourcePath"`
  SearchString string `json:"SearchString"`
  Type string `json:"Type"`
}

type AmazonRoute53HealthCheckTags struct {
  Key string `json:"Key"`
  Value string `json:"Value"`
}

type AmazonRoute53HostedZoneConfigProperty struct {
  Comment string `json:"Comment"`
}

type AmazonRoute53HostedZoneTags struct {
  Key string `json:"Key"`
  Value string `json:"Value"`
}

type AmazonRoute53HostedZoneVPCs struct {
  VPCId string `json:"VPCId"`
  VPCRegion string `json:"VPCRegion"`
}

type AmazonS3CorsConfiguration struct {
  CorsRules AmazonS3CorsConfigurationRule `json:"CorsRules"`
}

type AmazonS3CorsConfigurationRule struct {
  AllowedHeaders []string `json:"AllowedHeaders"`
  AllowedMethods []string `json:"AllowedMethods"`
  AllowedOrigins []string `json:"AllowedOrigins"`
  ExposedHeaders []string `json:"ExposedHeaders"`
  Id string `json:"Id"`
  MaxAge int `json:"MaxAge"`
}

type AmazonS3LifecycleConfiguration struct {
  Rules AmazonS3LifecycleRule `json:"Rules"`
}

type AmazonS3LifecycleRule struct {
  ExpirationDate string `json:"ExpirationDate"`
  ExpirationInDays int `json:"ExpirationInDays"`
  Id string `json:"Id"`
  NoncurrentVersionExpirationInDays int `json:"NoncurrentVersionExpirationInDays"`
  NoncurrentVersionTransition AmazonS3LifecycleRuleNoncurrentVersionTransition `json:"NoncurrentVersionTransition"`
  Prefix string `json:"Prefix"`
  Status string `json:"Status"`
  Transition AmazonS3LifecycleRuleTransition `json:"Transition"`
}

type AmazonS3LifecycleRuleNoncurrentVersionTransition struct {
  StorageClass string `json:"StorageClass"`
  TransitionInDays int `json:"TransitionInDays"`
}

type AmazonS3LifecycleRuleTransition struct {
  StorageClass string `json:"StorageClass"`
  TransitionDate string `json:"TransitionDate"`
  TransitionInDays int `json:"TransitionInDays"`
}

type AmazonS3LoggingConfiguration struct {
  DestinationBucketName string `json:"DestinationBucketName"`
  LogFilePrefix string `json:"LogFilePrefix"`
}

type AmazonS3NotificationConfiguration struct {
  TopicConfigurations AmazonS3NotificationTopicConfigurations `json:"TopicConfigurations"`
}

type AmazonS3NotificationTopicConfigurations struct {
  Event string `json:"Event"`
  Topic string `json:"Topic"`
}

type AmazonS3VersioningConfiguration struct {
  Status string `json:"Status"`
}

type AmazonS3WebsiteConfigurationProperty struct {
  ErrorDocument string `json:"ErrorDocument"`
  IndexDocument string `json:"IndexDocument"`
  RedirectAllRequestsTo AmazonS3WebsiteConfigurationRedirectAllRequestsToProperty `json:"RedirectAllRequestsTo"`
  RoutingRules AmazonS3WebsiteConfigurationRoutingRulesProperty `json:"RoutingRules"`
}

type AmazonS3WebsiteConfigurationRedirectAllRequestsToProperty struct {
  HostName string `json:"HostName"`
  Protocol string `json:"Protocol"`
}

type AmazonS3WebsiteConfigurationRoutingRulesProperty struct {
  RedirectRule AmazonS3WebsiteConfigurationRoutingRulesRedirectRuleProperty `json:"RedirectRule"`
  RoutingRuleCondition AmazonS3WebsiteConfigurationRoutingRulesRoutingRuleConditionProperty `json:"RoutingRuleCondition"`
}

type AmazonS3WebsiteConfigurationRoutingRulesRedirectRuleProperty struct {
  HostName string `json:"HostName"`
  HttpRedirectCode string `json:"HttpRedirectCode"`
  Protocol string `json:"Protocol"`
  ReplaceKeyPrefixWith string `json:"ReplaceKeyPrefixWith"`
  ReplaceKeyWith string `json:"ReplaceKeyWith"`
}

type AmazonS3WebsiteConfigurationRoutingRulesRoutingRuleConditionProperty struct {
  HttpErrorCodeReturnedEquals string `json:"HttpErrorCodeReturnedEquals"`
  KeyPrefixEquals string `json:"KeyPrefixEquals"`
}

type AmazonSNSSubscription struct {
  Endpoint string `json:"Endpoint"`
  Protocol string `json:"Protocol"`
}

type AmazonSQSRedrivePolicy struct {
  DeadLetterTargetArn string `json:"deadLetterTargetArn"`
  MaxReceiveCount int `json:"maxReceiveCount"`
}
