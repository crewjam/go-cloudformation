package cloudformation

// RESOURCE SPECIFICATION VERSION: 1.4.1
// GENERATED: 2017-06-10 14:07:05.489345898 -0700 PDT
import "time"
import "encoding/json"
import _ "gopkg.in/go-playground/validator.v9" // Used for struct level validation tags

// CustomResourceProvider allows extend the NewResourceByType factory method
// with their own resource types.
type CustomResourceProvider func(customResourceType string) ResourceProperties

var customResourceProviders []CustomResourceProvider

// RegisterCustomResourceProvider registers a custom resource provider with
// go-cloudformation. Multiple
// providers may be registered. The first provider that returns a non-nil
// interface will be used and there is no check for a uniquely registered
// resource type.
func RegisterCustomResourceProvider(provider CustomResourceProvider) {
	customResourceProviders = append(customResourceProviders, provider)
}

//
//  ____                            _   _
// |  _ \ _ __ ___  _ __   ___ _ __| |_(_) ___  ___
// | |_) | '__/ _ \| '_ \ / _ \ '__| __| |/ _ \/ __|
// |  __/| | | (_) | |_) |  __/ |  | |_| |  __/\__ \
// |_|   |_|  \___/| .__/ \___|_|   \__|_|\___||___/
//                 |_|
//

// APIGatewayAPIKeyStageKey represents the AWS::ApiGateway::ApiKey.StageKey CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-apikey-stagekey.html
type APIGatewayAPIKeyStageKey struct {
	// RestAPIID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-apikey-stagekey.html#cfn-apigateway-apikey-stagekey-restapiid
	RestAPIID *StringExpr `json:"RestApiId,omitempty"`
	// StageName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-apikey-stagekey.html#cfn-apigateway-apikey-stagekey-stagename
	StageName *StringExpr `json:"StageName,omitempty"`
}

// APIGatewayAPIKeyStageKeyList represents a list of APIGatewayAPIKeyStageKey
type APIGatewayAPIKeyStageKeyList []APIGatewayAPIKeyStageKey

// UnmarshalJSON sets the object from the provided JSON representation
func (l *APIGatewayAPIKeyStageKeyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := APIGatewayAPIKeyStageKey{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = APIGatewayAPIKeyStageKeyList{item}
		return nil
	}
	list := []APIGatewayAPIKeyStageKey{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = APIGatewayAPIKeyStageKeyList(list)
		return nil
	}
	return err
}

// APIGatewayDeploymentMethodSetting represents the AWS::ApiGateway::Deployment.MethodSetting CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription-methodsetting.html
type APIGatewayDeploymentMethodSetting struct {
	// CacheDataEncrypted docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-cachedataencrypted
	CacheDataEncrypted *BoolExpr `json:"CacheDataEncrypted,omitempty"`
	// CacheTTLInSeconds docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-cachettlinseconds
	CacheTTLInSeconds *IntegerExpr `json:"CacheTtlInSeconds,omitempty"`
	// CachingEnabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-cachingenabled
	CachingEnabled *BoolExpr `json:"CachingEnabled,omitempty"`
	// DataTraceEnabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-datatraceenabled
	DataTraceEnabled *BoolExpr `json:"DataTraceEnabled,omitempty"`
	// HTTPMethod docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-httpmethod
	HTTPMethod *StringExpr `json:"HttpMethod,omitempty"`
	// LoggingLevel docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-logginglevel
	LoggingLevel *StringExpr `json:"LoggingLevel,omitempty"`
	// MetricsEnabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-metricsenabled
	MetricsEnabled *BoolExpr `json:"MetricsEnabled,omitempty"`
	// ResourcePath docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-resourcepath
	ResourcePath *StringExpr `json:"ResourcePath,omitempty"`
	// ThrottlingBurstLimit docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-throttlingburstlimit
	ThrottlingBurstLimit *IntegerExpr `json:"ThrottlingBurstLimit,omitempty"`
	// ThrottlingRateLimit docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-throttlingratelimit
	ThrottlingRateLimit *IntegerExpr `json:"ThrottlingRateLimit,omitempty"`
}

// APIGatewayDeploymentMethodSettingList represents a list of APIGatewayDeploymentMethodSetting
type APIGatewayDeploymentMethodSettingList []APIGatewayDeploymentMethodSetting

// UnmarshalJSON sets the object from the provided JSON representation
func (l *APIGatewayDeploymentMethodSettingList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := APIGatewayDeploymentMethodSetting{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = APIGatewayDeploymentMethodSettingList{item}
		return nil
	}
	list := []APIGatewayDeploymentMethodSetting{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = APIGatewayDeploymentMethodSettingList(list)
		return nil
	}
	return err
}

// APIGatewayDeploymentStageDescription represents the AWS::ApiGateway::Deployment.StageDescription CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html
type APIGatewayDeploymentStageDescription struct {
	// CacheClusterEnabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-cacheclusterenabled
	CacheClusterEnabled *BoolExpr `json:"CacheClusterEnabled,omitempty"`
	// CacheClusterSize docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-cacheclustersize
	CacheClusterSize *StringExpr `json:"CacheClusterSize,omitempty"`
	// CacheDataEncrypted docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-cachedataencrypted
	CacheDataEncrypted *BoolExpr `json:"CacheDataEncrypted,omitempty"`
	// CacheTTLInSeconds docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-cachettlinseconds
	CacheTTLInSeconds *IntegerExpr `json:"CacheTtlInSeconds,omitempty"`
	// CachingEnabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-cachingenabled
	CachingEnabled *BoolExpr `json:"CachingEnabled,omitempty"`
	// ClientCertificateID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-clientcertificateid
	ClientCertificateID *StringExpr `json:"ClientCertificateId,omitempty"`
	// DataTraceEnabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-datatraceenabled
	DataTraceEnabled *BoolExpr `json:"DataTraceEnabled,omitempty"`
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-description
	Description *StringExpr `json:"Description,omitempty"`
	// LoggingLevel docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-logginglevel
	LoggingLevel *StringExpr `json:"LoggingLevel,omitempty"`
	// MethodSettings docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-methodsettings
	MethodSettings *APIGatewayDeploymentMethodSettingList `json:"MethodSettings,omitempty"`
	// MetricsEnabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-metricsenabled
	MetricsEnabled *BoolExpr `json:"MetricsEnabled,omitempty"`
	// StageName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-stagename
	StageName *StringExpr `json:"StageName,omitempty"`
	// ThrottlingBurstLimit docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-throttlingburstlimit
	ThrottlingBurstLimit *IntegerExpr `json:"ThrottlingBurstLimit,omitempty"`
	// ThrottlingRateLimit docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-throttlingratelimit
	ThrottlingRateLimit *IntegerExpr `json:"ThrottlingRateLimit,omitempty"`
	// Variables docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-variables
	Variables interface{} `json:"Variables,omitempty"`
}

// APIGatewayDeploymentStageDescriptionList represents a list of APIGatewayDeploymentStageDescription
type APIGatewayDeploymentStageDescriptionList []APIGatewayDeploymentStageDescription

// UnmarshalJSON sets the object from the provided JSON representation
func (l *APIGatewayDeploymentStageDescriptionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := APIGatewayDeploymentStageDescription{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = APIGatewayDeploymentStageDescriptionList{item}
		return nil
	}
	list := []APIGatewayDeploymentStageDescription{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = APIGatewayDeploymentStageDescriptionList(list)
		return nil
	}
	return err
}

// APIGatewayMethodIntegration represents the AWS::ApiGateway::Method.Integration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html
type APIGatewayMethodIntegration struct {
	// CacheKeyParameters docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html#cfn-apigateway-method-integration-cachekeyparameters
	CacheKeyParameters *StringListExpr `json:"CacheKeyParameters,omitempty"`
	// CacheNamespace docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html#cfn-apigateway-method-integration-cachenamespace
	CacheNamespace *StringExpr `json:"CacheNamespace,omitempty"`
	// Credentials docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html#cfn-apigateway-method-integration-credentials
	Credentials *StringExpr `json:"Credentials,omitempty"`
	// IntegrationHTTPMethod docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html#cfn-apigateway-method-integration-integrationhttpmethod
	IntegrationHTTPMethod *StringExpr `json:"IntegrationHttpMethod,omitempty"`
	// IntegrationResponses docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html#cfn-apigateway-method-integration-integrationresponses
	IntegrationResponses *APIGatewayMethodIntegrationResponseList `json:"IntegrationResponses,omitempty"`
	// PassthroughBehavior docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html#cfn-apigateway-method-integration-passthroughbehavior
	PassthroughBehavior *StringExpr `json:"PassthroughBehavior,omitempty"`
	// RequestParameters docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html#cfn-apigateway-method-integration-requestparameters
	RequestParameters interface{} `json:"RequestParameters,omitempty"`
	// RequestTemplates docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html#cfn-apigateway-method-integration-requesttemplates
	RequestTemplates interface{} `json:"RequestTemplates,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html#cfn-apigateway-method-integration-type
	Type *StringExpr `json:"Type,omitempty"`
	// URI docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html#cfn-apigateway-method-integration-uri
	URI *StringExpr `json:"Uri,omitempty"`
}

// APIGatewayMethodIntegrationList represents a list of APIGatewayMethodIntegration
type APIGatewayMethodIntegrationList []APIGatewayMethodIntegration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *APIGatewayMethodIntegrationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := APIGatewayMethodIntegration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = APIGatewayMethodIntegrationList{item}
		return nil
	}
	list := []APIGatewayMethodIntegration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = APIGatewayMethodIntegrationList(list)
		return nil
	}
	return err
}

// APIGatewayMethodIntegrationResponse represents the AWS::ApiGateway::Method.IntegrationResponse CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration-integrationresponse.html
type APIGatewayMethodIntegrationResponse struct {
	// ResponseParameters docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration-integrationresponse.html#cfn-apigateway-method-integration-integrationresponse-responseparameters
	ResponseParameters interface{} `json:"ResponseParameters,omitempty"`
	// ResponseTemplates docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration-integrationresponse.html#cfn-apigateway-method-integration-integrationresponse-responsetemplates
	ResponseTemplates interface{} `json:"ResponseTemplates,omitempty"`
	// SelectionPattern docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration-integrationresponse.html#cfn-apigateway-method-integration-integrationresponse-selectionpattern
	SelectionPattern *StringExpr `json:"SelectionPattern,omitempty"`
	// StatusCode docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration-integrationresponse.html#cfn-apigateway-method-integration-integrationresponse-statuscode
	StatusCode *StringExpr `json:"StatusCode,omitempty"`
}

// APIGatewayMethodIntegrationResponseList represents a list of APIGatewayMethodIntegrationResponse
type APIGatewayMethodIntegrationResponseList []APIGatewayMethodIntegrationResponse

// UnmarshalJSON sets the object from the provided JSON representation
func (l *APIGatewayMethodIntegrationResponseList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := APIGatewayMethodIntegrationResponse{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = APIGatewayMethodIntegrationResponseList{item}
		return nil
	}
	list := []APIGatewayMethodIntegrationResponse{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = APIGatewayMethodIntegrationResponseList(list)
		return nil
	}
	return err
}

// APIGatewayMethodMethodResponse represents the AWS::ApiGateway::Method.MethodResponse CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-methodresponse.html
type APIGatewayMethodMethodResponse struct {
	// ResponseModels docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-methodresponse.html#cfn-apigateway-method-methodresponse-responsemodels
	ResponseModels interface{} `json:"ResponseModels,omitempty"`
	// ResponseParameters docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-methodresponse.html#cfn-apigateway-method-methodresponse-responseparameters
	ResponseParameters interface{} `json:"ResponseParameters,omitempty"`
	// StatusCode docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-methodresponse.html#cfn-apigateway-method-methodresponse-statuscode
	StatusCode *StringExpr `json:"StatusCode,omitempty"`
}

// APIGatewayMethodMethodResponseList represents a list of APIGatewayMethodMethodResponse
type APIGatewayMethodMethodResponseList []APIGatewayMethodMethodResponse

// UnmarshalJSON sets the object from the provided JSON representation
func (l *APIGatewayMethodMethodResponseList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := APIGatewayMethodMethodResponse{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = APIGatewayMethodMethodResponseList{item}
		return nil
	}
	list := []APIGatewayMethodMethodResponse{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = APIGatewayMethodMethodResponseList(list)
		return nil
	}
	return err
}

// APIGatewayRestAPIS3Location represents the AWS::ApiGateway::RestApi.S3Location CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-restapi-bodys3location.html
type APIGatewayRestAPIS3Location struct {
	// Bucket docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-restapi-bodys3location.html#cfn-apigateway-restapi-s3location-bucket
	Bucket *StringExpr `json:"Bucket,omitempty"`
	// ETag docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-restapi-bodys3location.html#cfn-apigateway-restapi-s3location-etag
	ETag *StringExpr `json:"ETag,omitempty"`
	// Key docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-restapi-bodys3location.html#cfn-apigateway-restapi-s3location-key
	Key *StringExpr `json:"Key,omitempty"`
	// Version docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-restapi-bodys3location.html#cfn-apigateway-restapi-s3location-version
	Version *StringExpr `json:"Version,omitempty"`
}

// APIGatewayRestAPIS3LocationList represents a list of APIGatewayRestAPIS3Location
type APIGatewayRestAPIS3LocationList []APIGatewayRestAPIS3Location

// UnmarshalJSON sets the object from the provided JSON representation
func (l *APIGatewayRestAPIS3LocationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := APIGatewayRestAPIS3Location{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = APIGatewayRestAPIS3LocationList{item}
		return nil
	}
	list := []APIGatewayRestAPIS3Location{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = APIGatewayRestAPIS3LocationList(list)
		return nil
	}
	return err
}

// APIGatewayStageMethodSetting represents the AWS::ApiGateway::Stage.MethodSetting CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-stage-methodsetting.html
type APIGatewayStageMethodSetting struct {
	// CacheDataEncrypted docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-stage-methodsetting.html#cfn-apigateway-stage-methodsetting-cachedataencrypted
	CacheDataEncrypted *BoolExpr `json:"CacheDataEncrypted,omitempty"`
	// CacheTTLInSeconds docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-stage-methodsetting.html#cfn-apigateway-stage-methodsetting-cachettlinseconds
	CacheTTLInSeconds *IntegerExpr `json:"CacheTtlInSeconds,omitempty"`
	// CachingEnabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-stage-methodsetting.html#cfn-apigateway-stage-methodsetting-cachingenabled
	CachingEnabled *BoolExpr `json:"CachingEnabled,omitempty"`
	// DataTraceEnabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-stage-methodsetting.html#cfn-apigateway-stage-methodsetting-datatraceenabled
	DataTraceEnabled *BoolExpr `json:"DataTraceEnabled,omitempty"`
	// HTTPMethod docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-stage-methodsetting.html#cfn-apigateway-stage-methodsetting-httpmethod
	HTTPMethod *StringExpr `json:"HttpMethod,omitempty"`
	// LoggingLevel docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-stage-methodsetting.html#cfn-apigateway-stage-methodsetting-logginglevel
	LoggingLevel *StringExpr `json:"LoggingLevel,omitempty"`
	// MetricsEnabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-stage-methodsetting.html#cfn-apigateway-stage-methodsetting-metricsenabled
	MetricsEnabled *BoolExpr `json:"MetricsEnabled,omitempty"`
	// ResourcePath docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-stage-methodsetting.html#cfn-apigateway-stage-methodsetting-resourcepath
	ResourcePath *StringExpr `json:"ResourcePath,omitempty"`
	// ThrottlingBurstLimit docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-stage-methodsetting.html#cfn-apigateway-stage-methodsetting-throttlingburstlimit
	ThrottlingBurstLimit *IntegerExpr `json:"ThrottlingBurstLimit,omitempty"`
	// ThrottlingRateLimit docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-stage-methodsetting.html#cfn-apigateway-stage-methodsetting-throttlingratelimit
	ThrottlingRateLimit *IntegerExpr `json:"ThrottlingRateLimit,omitempty"`
}

// APIGatewayStageMethodSettingList represents a list of APIGatewayStageMethodSetting
type APIGatewayStageMethodSettingList []APIGatewayStageMethodSetting

// UnmarshalJSON sets the object from the provided JSON representation
func (l *APIGatewayStageMethodSettingList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := APIGatewayStageMethodSetting{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = APIGatewayStageMethodSettingList{item}
		return nil
	}
	list := []APIGatewayStageMethodSetting{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = APIGatewayStageMethodSettingList(list)
		return nil
	}
	return err
}

// APIGatewayUsagePlanAPIStage represents the AWS::ApiGateway::UsagePlan.ApiStage CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-apistage.html
type APIGatewayUsagePlanAPIStage struct {
	// APIID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-apistage.html#cfn-apigateway-usageplan-apistage-apiid
	APIID *StringExpr `json:"ApiId,omitempty"`
	// Stage docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-apistage.html#cfn-apigateway-usageplan-apistage-stage
	Stage *StringExpr `json:"Stage,omitempty"`
}

// APIGatewayUsagePlanAPIStageList represents a list of APIGatewayUsagePlanAPIStage
type APIGatewayUsagePlanAPIStageList []APIGatewayUsagePlanAPIStage

// UnmarshalJSON sets the object from the provided JSON representation
func (l *APIGatewayUsagePlanAPIStageList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := APIGatewayUsagePlanAPIStage{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = APIGatewayUsagePlanAPIStageList{item}
		return nil
	}
	list := []APIGatewayUsagePlanAPIStage{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = APIGatewayUsagePlanAPIStageList(list)
		return nil
	}
	return err
}

// APIGatewayUsagePlanQuotaSettings represents the AWS::ApiGateway::UsagePlan.QuotaSettings CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-quotasettings.html
type APIGatewayUsagePlanQuotaSettings struct {
	// Limit docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-quotasettings.html#cfn-apigateway-usageplan-quotasettings-limit
	Limit *IntegerExpr `json:"Limit,omitempty"`
	// Offset docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-quotasettings.html#cfn-apigateway-usageplan-quotasettings-offset
	Offset *IntegerExpr `json:"Offset,omitempty"`
	// Period docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-quotasettings.html#cfn-apigateway-usageplan-quotasettings-period
	Period *StringExpr `json:"Period,omitempty"`
}

// APIGatewayUsagePlanQuotaSettingsList represents a list of APIGatewayUsagePlanQuotaSettings
type APIGatewayUsagePlanQuotaSettingsList []APIGatewayUsagePlanQuotaSettings

// UnmarshalJSON sets the object from the provided JSON representation
func (l *APIGatewayUsagePlanQuotaSettingsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := APIGatewayUsagePlanQuotaSettings{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = APIGatewayUsagePlanQuotaSettingsList{item}
		return nil
	}
	list := []APIGatewayUsagePlanQuotaSettings{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = APIGatewayUsagePlanQuotaSettingsList(list)
		return nil
	}
	return err
}

// APIGatewayUsagePlanThrottleSettings represents the AWS::ApiGateway::UsagePlan.ThrottleSettings CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-throttlesettings.html
type APIGatewayUsagePlanThrottleSettings struct {
	// BurstLimit docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-throttlesettings.html#cfn-apigateway-usageplan-throttlesettings-burstlimit
	BurstLimit *IntegerExpr `json:"BurstLimit,omitempty"`
	// RateLimit docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-throttlesettings.html#cfn-apigateway-usageplan-throttlesettings-ratelimit
	RateLimit *IntegerExpr `json:"RateLimit,omitempty"`
}

// APIGatewayUsagePlanThrottleSettingsList represents a list of APIGatewayUsagePlanThrottleSettings
type APIGatewayUsagePlanThrottleSettingsList []APIGatewayUsagePlanThrottleSettings

// UnmarshalJSON sets the object from the provided JSON representation
func (l *APIGatewayUsagePlanThrottleSettingsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := APIGatewayUsagePlanThrottleSettings{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = APIGatewayUsagePlanThrottleSettingsList{item}
		return nil
	}
	list := []APIGatewayUsagePlanThrottleSettings{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = APIGatewayUsagePlanThrottleSettingsList(list)
		return nil
	}
	return err
}

// ApplicationAutoScalingScalingPolicyStepAdjustment represents the AWS::ApplicationAutoScaling::ScalingPolicy.StepAdjustment CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-stepadjustment.html
type ApplicationAutoScalingScalingPolicyStepAdjustment struct {
	// MetricIntervalLowerBound docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-stepadjustment.html#cfn-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-stepadjustment-metricintervallowerbound
	MetricIntervalLowerBound *IntegerExpr `json:"MetricIntervalLowerBound,omitempty"`
	// MetricIntervalUpperBound docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-stepadjustment.html#cfn-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-stepadjustment-metricintervalupperbound
	MetricIntervalUpperBound *IntegerExpr `json:"MetricIntervalUpperBound,omitempty"`
	// ScalingAdjustment docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-stepadjustment.html#cfn-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-stepadjustment-scalingadjustment
	ScalingAdjustment *IntegerExpr `json:"ScalingAdjustment,omitempty" validate:"dive,required"`
}

// ApplicationAutoScalingScalingPolicyStepAdjustmentList represents a list of ApplicationAutoScalingScalingPolicyStepAdjustment
type ApplicationAutoScalingScalingPolicyStepAdjustmentList []ApplicationAutoScalingScalingPolicyStepAdjustment

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ApplicationAutoScalingScalingPolicyStepAdjustmentList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ApplicationAutoScalingScalingPolicyStepAdjustment{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ApplicationAutoScalingScalingPolicyStepAdjustmentList{item}
		return nil
	}
	list := []ApplicationAutoScalingScalingPolicyStepAdjustment{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ApplicationAutoScalingScalingPolicyStepAdjustmentList(list)
		return nil
	}
	return err
}

// ApplicationAutoScalingScalingPolicyStepScalingPolicyConfiguration represents the AWS::ApplicationAutoScaling::ScalingPolicy.StepScalingPolicyConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration.html
type ApplicationAutoScalingScalingPolicyStepScalingPolicyConfiguration struct {
	// AdjustmentType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-adjustmenttype
	AdjustmentType *StringExpr `json:"AdjustmentType,omitempty"`
	// Cooldown docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-cooldown
	Cooldown *IntegerExpr `json:"Cooldown,omitempty"`
	// MetricAggregationType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-metricaggregationtype
	MetricAggregationType *StringExpr `json:"MetricAggregationType,omitempty"`
	// MinAdjustmentMagnitude docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-minadjustmentmagnitude
	MinAdjustmentMagnitude *IntegerExpr `json:"MinAdjustmentMagnitude,omitempty"`
	// StepAdjustments docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-stepadjustments
	StepAdjustments *ApplicationAutoScalingScalingPolicyStepAdjustmentList `json:"StepAdjustments,omitempty"`
}

// ApplicationAutoScalingScalingPolicyStepScalingPolicyConfigurationList represents a list of ApplicationAutoScalingScalingPolicyStepScalingPolicyConfiguration
type ApplicationAutoScalingScalingPolicyStepScalingPolicyConfigurationList []ApplicationAutoScalingScalingPolicyStepScalingPolicyConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ApplicationAutoScalingScalingPolicyStepScalingPolicyConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ApplicationAutoScalingScalingPolicyStepScalingPolicyConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ApplicationAutoScalingScalingPolicyStepScalingPolicyConfigurationList{item}
		return nil
	}
	list := []ApplicationAutoScalingScalingPolicyStepScalingPolicyConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ApplicationAutoScalingScalingPolicyStepScalingPolicyConfigurationList(list)
		return nil
	}
	return err
}

// AutoScalingAutoScalingGroupMetricsCollection represents the AWS::AutoScaling::AutoScalingGroup.MetricsCollection CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-metricscollection.html
type AutoScalingAutoScalingGroupMetricsCollection struct {
	// Granularity docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-metricscollection.html#cfn-as-metricscollection-granularity
	Granularity *StringExpr `json:"Granularity,omitempty" validate:"dive,required"`
	// Metrics docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-metricscollection.html#cfn-as-metricscollection-metrics
	Metrics *StringListExpr `json:"Metrics,omitempty"`
}

// AutoScalingAutoScalingGroupMetricsCollectionList represents a list of AutoScalingAutoScalingGroupMetricsCollection
type AutoScalingAutoScalingGroupMetricsCollectionList []AutoScalingAutoScalingGroupMetricsCollection

// UnmarshalJSON sets the object from the provided JSON representation
func (l *AutoScalingAutoScalingGroupMetricsCollectionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AutoScalingAutoScalingGroupMetricsCollection{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AutoScalingAutoScalingGroupMetricsCollectionList{item}
		return nil
	}
	list := []AutoScalingAutoScalingGroupMetricsCollection{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AutoScalingAutoScalingGroupMetricsCollectionList(list)
		return nil
	}
	return err
}

// AutoScalingAutoScalingGroupNotificationConfigurations represents the AWS::AutoScaling::AutoScalingGroup.NotificationConfigurations CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-notificationconfigurations.html
type AutoScalingAutoScalingGroupNotificationConfigurations struct {
	// NotificationTypes docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-notificationconfigurations.html#cfn-as-group-notificationconfigurations-notificationtypes
	NotificationTypes *StringListExpr `json:"NotificationTypes,omitempty"`
	// TopicARN docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-notificationconfigurations.html#cfn-autoscaling-autoscalinggroup-notificationconfigurations-topicarn
	TopicARN *StringExpr `json:"TopicARN,omitempty" validate:"dive,required"`
}

// AutoScalingAutoScalingGroupNotificationConfigurationsList represents a list of AutoScalingAutoScalingGroupNotificationConfigurations
type AutoScalingAutoScalingGroupNotificationConfigurationsList []AutoScalingAutoScalingGroupNotificationConfigurations

// UnmarshalJSON sets the object from the provided JSON representation
func (l *AutoScalingAutoScalingGroupNotificationConfigurationsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AutoScalingAutoScalingGroupNotificationConfigurations{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AutoScalingAutoScalingGroupNotificationConfigurationsList{item}
		return nil
	}
	list := []AutoScalingAutoScalingGroupNotificationConfigurations{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AutoScalingAutoScalingGroupNotificationConfigurationsList(list)
		return nil
	}
	return err
}

// AutoScalingAutoScalingGroupTagProperty represents the AWS::AutoScaling::AutoScalingGroup.TagProperty CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-tags.html
type AutoScalingAutoScalingGroupTagProperty struct {
	// Key docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-tags.html#cfn-as-tags-Key
	Key *StringExpr `json:"Key,omitempty" validate:"dive,required"`
	// PropagateAtLaunch docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-tags.html#cfn-as-tags-PropagateAtLaunch
	PropagateAtLaunch *BoolExpr `json:"PropagateAtLaunch,omitempty" validate:"dive,required"`
	// Value docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-tags.html#cfn-as-tags-Value
	Value *StringExpr `json:"Value,omitempty" validate:"dive,required"`
}

// AutoScalingAutoScalingGroupTagPropertyList represents a list of AutoScalingAutoScalingGroupTagProperty
type AutoScalingAutoScalingGroupTagPropertyList []AutoScalingAutoScalingGroupTagProperty

// UnmarshalJSON sets the object from the provided JSON representation
func (l *AutoScalingAutoScalingGroupTagPropertyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AutoScalingAutoScalingGroupTagProperty{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AutoScalingAutoScalingGroupTagPropertyList{item}
		return nil
	}
	list := []AutoScalingAutoScalingGroupTagProperty{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AutoScalingAutoScalingGroupTagPropertyList(list)
		return nil
	}
	return err
}

// AutoScalingLaunchConfigurationBlockDevice represents the AWS::AutoScaling::LaunchConfiguration.BlockDevice CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-template.html
type AutoScalingLaunchConfigurationBlockDevice struct {
	// DeleteOnTermination docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-template.html#cfn-as-launchconfig-blockdev-template-deleteonterm
	DeleteOnTermination *BoolExpr `json:"DeleteOnTermination,omitempty"`
	// Encrypted docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-template.html#cfn-as-launchconfig-blockdev-template-encrypted
	Encrypted *BoolExpr `json:"Encrypted,omitempty"`
	// Iops docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-template.html#cfn-as-launchconfig-blockdev-template-iops
	Iops *IntegerExpr `json:"Iops,omitempty"`
	// SnapshotID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-template.html#cfn-as-launchconfig-blockdev-template-snapshotid
	SnapshotID *StringExpr `json:"SnapshotId,omitempty"`
	// VolumeSize docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-template.html#cfn-as-launchconfig-blockdev-template-volumesize
	VolumeSize *IntegerExpr `json:"VolumeSize,omitempty"`
	// VolumeType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-template.html#cfn-as-launchconfig-blockdev-template-volumetype
	VolumeType *StringExpr `json:"VolumeType,omitempty"`
}

// AutoScalingLaunchConfigurationBlockDeviceList represents a list of AutoScalingLaunchConfigurationBlockDevice
type AutoScalingLaunchConfigurationBlockDeviceList []AutoScalingLaunchConfigurationBlockDevice

// UnmarshalJSON sets the object from the provided JSON representation
func (l *AutoScalingLaunchConfigurationBlockDeviceList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AutoScalingLaunchConfigurationBlockDevice{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AutoScalingLaunchConfigurationBlockDeviceList{item}
		return nil
	}
	list := []AutoScalingLaunchConfigurationBlockDevice{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AutoScalingLaunchConfigurationBlockDeviceList(list)
		return nil
	}
	return err
}

// AutoScalingLaunchConfigurationBlockDeviceMapping represents the AWS::AutoScaling::LaunchConfiguration.BlockDeviceMapping CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-mapping.html
type AutoScalingLaunchConfigurationBlockDeviceMapping struct {
	// DeviceName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-mapping.html#cfn-as-launchconfig-blockdev-mapping-devicename
	DeviceName *StringExpr `json:"DeviceName,omitempty" validate:"dive,required"`
	// Ebs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-mapping.html#cfn-as-launchconfig-blockdev-mapping-ebs
	Ebs *AutoScalingLaunchConfigurationBlockDevice `json:"Ebs,omitempty"`
	// NoDevice docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-mapping.html#cfn-as-launchconfig-blockdev-mapping-nodevice
	NoDevice *BoolExpr `json:"NoDevice,omitempty"`
	// VirtualName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-mapping.html#cfn-as-launchconfig-blockdev-mapping-virtualname
	VirtualName *StringExpr `json:"VirtualName,omitempty"`
}

// AutoScalingLaunchConfigurationBlockDeviceMappingList represents a list of AutoScalingLaunchConfigurationBlockDeviceMapping
type AutoScalingLaunchConfigurationBlockDeviceMappingList []AutoScalingLaunchConfigurationBlockDeviceMapping

// UnmarshalJSON sets the object from the provided JSON representation
func (l *AutoScalingLaunchConfigurationBlockDeviceMappingList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AutoScalingLaunchConfigurationBlockDeviceMapping{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AutoScalingLaunchConfigurationBlockDeviceMappingList{item}
		return nil
	}
	list := []AutoScalingLaunchConfigurationBlockDeviceMapping{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AutoScalingLaunchConfigurationBlockDeviceMappingList(list)
		return nil
	}
	return err
}

// AutoScalingScalingPolicyStepAdjustment represents the AWS::AutoScaling::ScalingPolicy.StepAdjustment CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-stepadjustments.html
type AutoScalingScalingPolicyStepAdjustment struct {
	// MetricIntervalLowerBound docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-stepadjustments.html#cfn-autoscaling-scalingpolicy-stepadjustment-metricintervallowerbound
	MetricIntervalLowerBound *IntegerExpr `json:"MetricIntervalLowerBound,omitempty"`
	// MetricIntervalUpperBound docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-stepadjustments.html#cfn-autoscaling-scalingpolicy-stepadjustment-metricintervalupperbound
	MetricIntervalUpperBound *IntegerExpr `json:"MetricIntervalUpperBound,omitempty"`
	// ScalingAdjustment docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-stepadjustments.html#cfn-autoscaling-scalingpolicy-stepadjustment-scalingadjustment
	ScalingAdjustment *IntegerExpr `json:"ScalingAdjustment,omitempty" validate:"dive,required"`
}

// AutoScalingScalingPolicyStepAdjustmentList represents a list of AutoScalingScalingPolicyStepAdjustment
type AutoScalingScalingPolicyStepAdjustmentList []AutoScalingScalingPolicyStepAdjustment

// UnmarshalJSON sets the object from the provided JSON representation
func (l *AutoScalingScalingPolicyStepAdjustmentList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := AutoScalingScalingPolicyStepAdjustment{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = AutoScalingScalingPolicyStepAdjustmentList{item}
		return nil
	}
	list := []AutoScalingScalingPolicyStepAdjustment{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = AutoScalingScalingPolicyStepAdjustmentList(list)
		return nil
	}
	return err
}

// CertificateManagerCertificateDomainValidationOption represents the AWS::CertificateManager::Certificate.DomainValidationOption CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-certificate-domainvalidationoption.html
type CertificateManagerCertificateDomainValidationOption struct {
	// DomainName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-certificate-domainvalidationoption.html#cfn-certificatemanager-certificate-domainvalidationoptions-domainname
	DomainName *StringExpr `json:"DomainName,omitempty" validate:"dive,required"`
	// ValidationDomain docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-certificate-domainvalidationoption.html#cfn-certificatemanager-certificate-domainvalidationoption-validationdomain
	ValidationDomain *StringExpr `json:"ValidationDomain,omitempty" validate:"dive,required"`
}

// CertificateManagerCertificateDomainValidationOptionList represents a list of CertificateManagerCertificateDomainValidationOption
type CertificateManagerCertificateDomainValidationOptionList []CertificateManagerCertificateDomainValidationOption

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CertificateManagerCertificateDomainValidationOptionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CertificateManagerCertificateDomainValidationOption{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CertificateManagerCertificateDomainValidationOptionList{item}
		return nil
	}
	list := []CertificateManagerCertificateDomainValidationOption{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CertificateManagerCertificateDomainValidationOptionList(list)
		return nil
	}
	return err
}

// CloudFrontDistributionCacheBehavior represents the AWS::CloudFront::Distribution.CacheBehavior CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachebehavior.html
type CloudFrontDistributionCacheBehavior struct {
	// AllowedMethods docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachebehavior.html#cfn-cloudfront-cachebehavior-allowedmethods
	AllowedMethods *StringListExpr `json:"AllowedMethods,omitempty"`
	// CachedMethods docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachebehavior.html#cfn-cloudfront-cachebehavior-cachedmethods
	CachedMethods *StringListExpr `json:"CachedMethods,omitempty"`
	// Compress docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachebehavior.html#cfn-cloudfront-cachebehavior-compress
	Compress *BoolExpr `json:"Compress,omitempty"`
	// DefaultTTL docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachebehavior.html#cfn-cloudfront-cachebehavior-defaultttl
	DefaultTTL *IntegerExpr `json:"DefaultTTL,omitempty"`
	// ForwardedValues docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachebehavior.html#cfn-cloudfront-cachebehavior-forwardedvalues
	ForwardedValues *CloudFrontDistributionForwardedValues `json:"ForwardedValues,omitempty" validate:"dive,required"`
	// MaxTTL docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachebehavior.html#cfn-cloudfront-cachebehavior-maxttl
	MaxTTL *IntegerExpr `json:"MaxTTL,omitempty"`
	// MinTTL docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachebehavior.html#cfn-cloudfront-cachebehavior-minttl
	MinTTL *IntegerExpr `json:"MinTTL,omitempty"`
	// PathPattern docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachebehavior.html#cfn-cloudfront-cachebehavior-pathpattern
	PathPattern *StringExpr `json:"PathPattern,omitempty" validate:"dive,required"`
	// SmoothStreaming docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachebehavior.html#cfn-cloudfront-cachebehavior-smoothstreaming
	SmoothStreaming *BoolExpr `json:"SmoothStreaming,omitempty"`
	// TargetOriginID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachebehavior.html#cfn-cloudfront-cachebehavior-targetoriginid
	TargetOriginID *StringExpr `json:"TargetOriginId,omitempty" validate:"dive,required"`
	// TrustedSigners docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachebehavior.html#cfn-cloudfront-cachebehavior-trustedsigners
	TrustedSigners *StringListExpr `json:"TrustedSigners,omitempty"`
	// ViewerProtocolPolicy docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachebehavior.html#cfn-cloudfront-cachebehavior-viewerprotocolpolicy
	ViewerProtocolPolicy *StringExpr `json:"ViewerProtocolPolicy,omitempty" validate:"dive,required"`
}

// CloudFrontDistributionCacheBehaviorList represents a list of CloudFrontDistributionCacheBehavior
type CloudFrontDistributionCacheBehaviorList []CloudFrontDistributionCacheBehavior

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudFrontDistributionCacheBehaviorList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudFrontDistributionCacheBehavior{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudFrontDistributionCacheBehaviorList{item}
		return nil
	}
	list := []CloudFrontDistributionCacheBehavior{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudFrontDistributionCacheBehaviorList(list)
		return nil
	}
	return err
}

// CloudFrontDistributionCookies represents the AWS::CloudFront::Distribution.Cookies CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-forwardedvalues-cookies.html
type CloudFrontDistributionCookies struct {
	// Forward docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-forwardedvalues-cookies.html#cfn-cloudfront-forwardedvalues-cookies-forward
	Forward *StringExpr `json:"Forward,omitempty" validate:"dive,required"`
	// WhitelistedNames docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-forwardedvalues-cookies.html#cfn-cloudfront-forwardedvalues-cookies-whitelistednames
	WhitelistedNames *StringListExpr `json:"WhitelistedNames,omitempty"`
}

// CloudFrontDistributionCookiesList represents a list of CloudFrontDistributionCookies
type CloudFrontDistributionCookiesList []CloudFrontDistributionCookies

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudFrontDistributionCookiesList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudFrontDistributionCookies{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudFrontDistributionCookiesList{item}
		return nil
	}
	list := []CloudFrontDistributionCookies{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudFrontDistributionCookiesList(list)
		return nil
	}
	return err
}

// CloudFrontDistributionCustomErrorResponse represents the AWS::CloudFront::Distribution.CustomErrorResponse CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-customerrorresponse.html
type CloudFrontDistributionCustomErrorResponse struct {
	// ErrorCachingMinTTL docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-customerrorresponse.html#cfn-cloudfront-distributionconfig-customerrorresponse-errorcachingminttl
	ErrorCachingMinTTL *IntegerExpr `json:"ErrorCachingMinTTL,omitempty"`
	// ErrorCode docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-customerrorresponse.html#cfn-cloudfront-distributionconfig-customerrorresponse-errorcode
	ErrorCode *IntegerExpr `json:"ErrorCode,omitempty" validate:"dive,required"`
	// ResponseCode docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-customerrorresponse.html#cfn-cloudfront-distributionconfig-customerrorresponse-responsecode
	ResponseCode *IntegerExpr `json:"ResponseCode,omitempty"`
	// ResponsePagePath docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-customerrorresponse.html#cfn-cloudfront-distributionconfig-customerrorresponse-responsepagepath
	ResponsePagePath *StringExpr `json:"ResponsePagePath,omitempty"`
}

// CloudFrontDistributionCustomErrorResponseList represents a list of CloudFrontDistributionCustomErrorResponse
type CloudFrontDistributionCustomErrorResponseList []CloudFrontDistributionCustomErrorResponse

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudFrontDistributionCustomErrorResponseList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudFrontDistributionCustomErrorResponse{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudFrontDistributionCustomErrorResponseList{item}
		return nil
	}
	list := []CloudFrontDistributionCustomErrorResponse{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudFrontDistributionCustomErrorResponseList(list)
		return nil
	}
	return err
}

// CloudFrontDistributionCustomOriginConfig represents the AWS::CloudFront::Distribution.CustomOriginConfig CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-customorigin.html
type CloudFrontDistributionCustomOriginConfig struct {
	// HTTPPort docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-customorigin.html#cfn-cloudfront-customorigin-httpport
	HTTPPort *IntegerExpr `json:"HTTPPort,omitempty"`
	// HTTPSPort docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-customorigin.html#cfn-cloudfront-customorigin-httpsport
	HTTPSPort *IntegerExpr `json:"HTTPSPort,omitempty"`
	// OriginProtocolPolicy docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-customorigin.html#cfn-cloudfront-customorigin-originprotocolpolicy
	OriginProtocolPolicy *StringExpr `json:"OriginProtocolPolicy,omitempty" validate:"dive,required"`
	// OriginSSLProtocols docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-customorigin.html#cfn-cloudfront-customorigin-originsslprotocols
	OriginSSLProtocols *StringListExpr `json:"OriginSSLProtocols,omitempty"`
}

// CloudFrontDistributionCustomOriginConfigList represents a list of CloudFrontDistributionCustomOriginConfig
type CloudFrontDistributionCustomOriginConfigList []CloudFrontDistributionCustomOriginConfig

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudFrontDistributionCustomOriginConfigList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudFrontDistributionCustomOriginConfig{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudFrontDistributionCustomOriginConfigList{item}
		return nil
	}
	list := []CloudFrontDistributionCustomOriginConfig{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudFrontDistributionCustomOriginConfigList(list)
		return nil
	}
	return err
}

// CloudFrontDistributionDefaultCacheBehavior represents the AWS::CloudFront::Distribution.DefaultCacheBehavior CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-defaultcachebehavior.html
type CloudFrontDistributionDefaultCacheBehavior struct {
	// AllowedMethods docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-defaultcachebehavior.html#cfn-cloudfront-defaultcachebehavior-allowedmethods
	AllowedMethods *StringListExpr `json:"AllowedMethods,omitempty"`
	// CachedMethods docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-defaultcachebehavior.html#cfn-cloudfront-defaultcachebehavior-cachedmethods
	CachedMethods *StringListExpr `json:"CachedMethods,omitempty"`
	// Compress docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-defaultcachebehavior.html#cfn-cloudfront-defaultcachebehavior-compress
	Compress *BoolExpr `json:"Compress,omitempty"`
	// DefaultTTL docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-defaultcachebehavior.html#cfn-cloudfront-defaultcachebehavior-defaultttl
	DefaultTTL *IntegerExpr `json:"DefaultTTL,omitempty"`
	// ForwardedValues docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-defaultcachebehavior.html#cfn-cloudfront-defaultcachebehavior-forwardedvalues
	ForwardedValues *CloudFrontDistributionForwardedValues `json:"ForwardedValues,omitempty" validate:"dive,required"`
	// MaxTTL docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-defaultcachebehavior.html#cfn-cloudfront-defaultcachebehavior-maxttl
	MaxTTL *IntegerExpr `json:"MaxTTL,omitempty"`
	// MinTTL docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-defaultcachebehavior.html#cfn-cloudfront-defaultcachebehavior-minttl
	MinTTL *IntegerExpr `json:"MinTTL,omitempty"`
	// SmoothStreaming docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-defaultcachebehavior.html#cfn-cloudfront-defaultcachebehavior-smoothstreaming
	SmoothStreaming *BoolExpr `json:"SmoothStreaming,omitempty"`
	// TargetOriginID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-defaultcachebehavior.html#cfn-cloudfront-defaultcachebehavior-targetoriginid
	TargetOriginID *StringExpr `json:"TargetOriginId,omitempty" validate:"dive,required"`
	// TrustedSigners docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-defaultcachebehavior.html#cfn-cloudfront-defaultcachebehavior-trustedsigners
	TrustedSigners *StringListExpr `json:"TrustedSigners,omitempty"`
	// ViewerProtocolPolicy docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-defaultcachebehavior.html#cfn-cloudfront-defaultcachebehavior-viewerprotocolpolicy
	ViewerProtocolPolicy *StringExpr `json:"ViewerProtocolPolicy,omitempty" validate:"dive,required"`
}

// CloudFrontDistributionDefaultCacheBehaviorList represents a list of CloudFrontDistributionDefaultCacheBehavior
type CloudFrontDistributionDefaultCacheBehaviorList []CloudFrontDistributionDefaultCacheBehavior

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudFrontDistributionDefaultCacheBehaviorList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudFrontDistributionDefaultCacheBehavior{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudFrontDistributionDefaultCacheBehaviorList{item}
		return nil
	}
	list := []CloudFrontDistributionDefaultCacheBehavior{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudFrontDistributionDefaultCacheBehaviorList(list)
		return nil
	}
	return err
}

// CloudFrontDistributionDistributionConfig represents the AWS::CloudFront::Distribution.DistributionConfig CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html
type CloudFrontDistributionDistributionConfig struct {
	// Aliases docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html#cfn-cloudfront-distributionconfig-aliases
	Aliases *StringListExpr `json:"Aliases,omitempty"`
	// CacheBehaviors docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html#cfn-cloudfront-distributionconfig-cachebehaviors
	CacheBehaviors *CloudFrontDistributionCacheBehaviorList `json:"CacheBehaviors,omitempty"`
	// Comment docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html#cfn-cloudfront-distributionconfig-comment
	Comment *StringExpr `json:"Comment,omitempty"`
	// CustomErrorResponses docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html#cfn-cloudfront-distributionconfig-customerrorresponses
	CustomErrorResponses *CloudFrontDistributionCustomErrorResponseList `json:"CustomErrorResponses,omitempty"`
	// DefaultCacheBehavior docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html#cfn-cloudfront-distributionconfig-defaultcachebehavior
	DefaultCacheBehavior *CloudFrontDistributionDefaultCacheBehavior `json:"DefaultCacheBehavior,omitempty" validate:"dive,required"`
	// DefaultRootObject docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html#cfn-cloudfront-distributionconfig-defaultrootobject
	DefaultRootObject *StringExpr `json:"DefaultRootObject,omitempty"`
	// Enabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html#cfn-cloudfront-distributionconfig-enabled
	Enabled *BoolExpr `json:"Enabled,omitempty" validate:"dive,required"`
	// HTTPVersion docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html#cfn-cloudfront-distributionconfig-httpversion
	HTTPVersion *StringExpr `json:"HttpVersion,omitempty"`
	// Logging docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html#cfn-cloudfront-distributionconfig-logging
	Logging *CloudFrontDistributionLogging `json:"Logging,omitempty"`
	// Origins docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html#cfn-cloudfront-distributionconfig-origins
	Origins *CloudFrontDistributionOriginList `json:"Origins,omitempty" validate:"dive,required"`
	// PriceClass docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html#cfn-cloudfront-distributionconfig-priceclass
	PriceClass *StringExpr `json:"PriceClass,omitempty"`
	// Restrictions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html#cfn-cloudfront-distributionconfig-restrictions
	Restrictions *CloudFrontDistributionRestrictions `json:"Restrictions,omitempty"`
	// ViewerCertificate docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html#cfn-cloudfront-distributionconfig-viewercertificate
	ViewerCertificate *CloudFrontDistributionViewerCertificate `json:"ViewerCertificate,omitempty"`
	// WebACLID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html#cfn-cloudfront-distributionconfig-webaclid
	WebACLID *StringExpr `json:"WebACLId,omitempty"`
}

// CloudFrontDistributionDistributionConfigList represents a list of CloudFrontDistributionDistributionConfig
type CloudFrontDistributionDistributionConfigList []CloudFrontDistributionDistributionConfig

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudFrontDistributionDistributionConfigList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudFrontDistributionDistributionConfig{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudFrontDistributionDistributionConfigList{item}
		return nil
	}
	list := []CloudFrontDistributionDistributionConfig{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudFrontDistributionDistributionConfigList(list)
		return nil
	}
	return err
}

// CloudFrontDistributionForwardedValues represents the AWS::CloudFront::Distribution.ForwardedValues CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-forwardedvalues.html
type CloudFrontDistributionForwardedValues struct {
	// Cookies docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-forwardedvalues.html#cfn-cloudfront-forwardedvalues-cookies
	Cookies *CloudFrontDistributionCookies `json:"Cookies,omitempty"`
	// Headers docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-forwardedvalues.html#cfn-cloudfront-forwardedvalues-headers
	Headers *StringListExpr `json:"Headers,omitempty"`
	// QueryString docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-forwardedvalues.html#cfn-cloudfront-forwardedvalues-querystring
	QueryString *BoolExpr `json:"QueryString,omitempty" validate:"dive,required"`
	// QueryStringCacheKeys docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-forwardedvalues.html#cfn-cloudfront-forwardedvalues-querystringcachekeys
	QueryStringCacheKeys *StringListExpr `json:"QueryStringCacheKeys,omitempty"`
}

// CloudFrontDistributionForwardedValuesList represents a list of CloudFrontDistributionForwardedValues
type CloudFrontDistributionForwardedValuesList []CloudFrontDistributionForwardedValues

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudFrontDistributionForwardedValuesList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudFrontDistributionForwardedValues{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudFrontDistributionForwardedValuesList{item}
		return nil
	}
	list := []CloudFrontDistributionForwardedValues{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudFrontDistributionForwardedValuesList(list)
		return nil
	}
	return err
}

// CloudFrontDistributionGeoRestriction represents the AWS::CloudFront::Distribution.GeoRestriction CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-restrictions-georestriction.html
type CloudFrontDistributionGeoRestriction struct {
	// Locations docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-restrictions-georestriction.html#cfn-cloudfront-distributionconfig-restrictions-georestriction-locations
	Locations *StringListExpr `json:"Locations,omitempty"`
	// RestrictionType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-restrictions-georestriction.html#cfn-cloudfront-distributionconfig-restrictions-georestriction-restrictiontype
	RestrictionType *StringExpr `json:"RestrictionType,omitempty" validate:"dive,required"`
}

// CloudFrontDistributionGeoRestrictionList represents a list of CloudFrontDistributionGeoRestriction
type CloudFrontDistributionGeoRestrictionList []CloudFrontDistributionGeoRestriction

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudFrontDistributionGeoRestrictionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudFrontDistributionGeoRestriction{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudFrontDistributionGeoRestrictionList{item}
		return nil
	}
	list := []CloudFrontDistributionGeoRestriction{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudFrontDistributionGeoRestrictionList(list)
		return nil
	}
	return err
}

// CloudFrontDistributionLogging represents the AWS::CloudFront::Distribution.Logging CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-logging.html
type CloudFrontDistributionLogging struct {
	// Bucket docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-logging.html#cfn-cloudfront-logging-bucket
	Bucket *StringExpr `json:"Bucket,omitempty" validate:"dive,required"`
	// IncludeCookies docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-logging.html#cfn-cloudfront-logging-includecookies
	IncludeCookies *BoolExpr `json:"IncludeCookies,omitempty"`
	// Prefix docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-logging.html#cfn-cloudfront-logging-prefix
	Prefix *StringExpr `json:"Prefix,omitempty"`
}

// CloudFrontDistributionLoggingList represents a list of CloudFrontDistributionLogging
type CloudFrontDistributionLoggingList []CloudFrontDistributionLogging

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudFrontDistributionLoggingList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudFrontDistributionLogging{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudFrontDistributionLoggingList{item}
		return nil
	}
	list := []CloudFrontDistributionLogging{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudFrontDistributionLoggingList(list)
		return nil
	}
	return err
}

// CloudFrontDistributionOrigin represents the AWS::CloudFront::Distribution.Origin CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin.html
type CloudFrontDistributionOrigin struct {
	// CustomOriginConfig docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin.html#cfn-cloudfront-origin-customorigin
	CustomOriginConfig *CloudFrontDistributionCustomOriginConfig `json:"CustomOriginConfig,omitempty"`
	// DomainName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin.html#cfn-cloudfront-origin-domainname
	DomainName *StringExpr `json:"DomainName,omitempty" validate:"dive,required"`
	// ID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin.html#cfn-cloudfront-origin-id
	ID *StringExpr `json:"Id,omitempty" validate:"dive,required"`
	// OriginCustomHeaders docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin.html#cfn-cloudfront-origin-origincustomheaders
	OriginCustomHeaders *CloudFrontDistributionOriginCustomHeaderList `json:"OriginCustomHeaders,omitempty"`
	// OriginPath docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin.html#cfn-cloudfront-origin-originpath
	OriginPath *StringExpr `json:"OriginPath,omitempty"`
	// S3OriginConfig docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin.html#cfn-cloudfront-origin-s3origin
	S3OriginConfig *CloudFrontDistributionS3OriginConfig `json:"S3OriginConfig,omitempty"`
}

// CloudFrontDistributionOriginList represents a list of CloudFrontDistributionOrigin
type CloudFrontDistributionOriginList []CloudFrontDistributionOrigin

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudFrontDistributionOriginList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudFrontDistributionOrigin{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudFrontDistributionOriginList{item}
		return nil
	}
	list := []CloudFrontDistributionOrigin{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudFrontDistributionOriginList(list)
		return nil
	}
	return err
}

// CloudFrontDistributionOriginCustomHeader represents the AWS::CloudFront::Distribution.OriginCustomHeader CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin-origincustomheader.html
type CloudFrontDistributionOriginCustomHeader struct {
	// HeaderName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin-origincustomheader.html#cfn-cloudfront-origin-origincustomheader-headername
	HeaderName *StringExpr `json:"HeaderName,omitempty" validate:"dive,required"`
	// HeaderValue docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin-origincustomheader.html#cfn-cloudfront-origin-origincustomheader-headervalue
	HeaderValue *StringExpr `json:"HeaderValue,omitempty" validate:"dive,required"`
}

// CloudFrontDistributionOriginCustomHeaderList represents a list of CloudFrontDistributionOriginCustomHeader
type CloudFrontDistributionOriginCustomHeaderList []CloudFrontDistributionOriginCustomHeader

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudFrontDistributionOriginCustomHeaderList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudFrontDistributionOriginCustomHeader{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudFrontDistributionOriginCustomHeaderList{item}
		return nil
	}
	list := []CloudFrontDistributionOriginCustomHeader{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudFrontDistributionOriginCustomHeaderList(list)
		return nil
	}
	return err
}

// CloudFrontDistributionRestrictions represents the AWS::CloudFront::Distribution.Restrictions CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-restrictions.html
type CloudFrontDistributionRestrictions struct {
	// GeoRestriction docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-restrictions.html#cfn-cloudfront-distributionconfig-restrictions-georestriction
	GeoRestriction *CloudFrontDistributionGeoRestriction `json:"GeoRestriction,omitempty" validate:"dive,required"`
}

// CloudFrontDistributionRestrictionsList represents a list of CloudFrontDistributionRestrictions
type CloudFrontDistributionRestrictionsList []CloudFrontDistributionRestrictions

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudFrontDistributionRestrictionsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudFrontDistributionRestrictions{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudFrontDistributionRestrictionsList{item}
		return nil
	}
	list := []CloudFrontDistributionRestrictions{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudFrontDistributionRestrictionsList(list)
		return nil
	}
	return err
}

// CloudFrontDistributionS3OriginConfig represents the AWS::CloudFront::Distribution.S3OriginConfig CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-s3origin.html
type CloudFrontDistributionS3OriginConfig struct {
	// OriginAccessIDentity docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-s3origin.html#cfn-cloudfront-s3origin-originaccessidentity
	OriginAccessIDentity *StringExpr `json:"OriginAccessIdentity,omitempty"`
}

// CloudFrontDistributionS3OriginConfigList represents a list of CloudFrontDistributionS3OriginConfig
type CloudFrontDistributionS3OriginConfigList []CloudFrontDistributionS3OriginConfig

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudFrontDistributionS3OriginConfigList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudFrontDistributionS3OriginConfig{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudFrontDistributionS3OriginConfigList{item}
		return nil
	}
	list := []CloudFrontDistributionS3OriginConfig{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudFrontDistributionS3OriginConfigList(list)
		return nil
	}
	return err
}

// CloudFrontDistributionViewerCertificate represents the AWS::CloudFront::Distribution.ViewerCertificate CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-viewercertificate.html
type CloudFrontDistributionViewerCertificate struct {
	// AcmCertificateArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-viewercertificate.html#cfn-cloudfront-distributionconfig-viewercertificate-acmcertificatearn
	AcmCertificateArn *StringExpr `json:"AcmCertificateArn,omitempty"`
	// CloudFrontDefaultCertificate docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-viewercertificate.html#cfn-cloudfront-distributionconfig-viewercertificate-cloudfrontdefaultcertificate
	CloudFrontDefaultCertificate *BoolExpr `json:"CloudFrontDefaultCertificate,omitempty"`
	// IamCertificateID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-viewercertificate.html#cfn-cloudfront-distributionconfig-viewercertificate-iamcertificateid
	IamCertificateID *StringExpr `json:"IamCertificateId,omitempty"`
	// MinimumProtocolVersion docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-viewercertificate.html#cfn-cloudfront-distributionconfig-viewercertificate-sslsupportmethod
	MinimumProtocolVersion *StringExpr `json:"MinimumProtocolVersion,omitempty"`
	// SslSupportMethod docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-viewercertificate.html#cfn-cloudfront-distributionconfig-viewercertificate-minimumprotocolversion
	SslSupportMethod *StringExpr `json:"SslSupportMethod,omitempty"`
}

// CloudFrontDistributionViewerCertificateList represents a list of CloudFrontDistributionViewerCertificate
type CloudFrontDistributionViewerCertificateList []CloudFrontDistributionViewerCertificate

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudFrontDistributionViewerCertificateList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudFrontDistributionViewerCertificate{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudFrontDistributionViewerCertificateList{item}
		return nil
	}
	list := []CloudFrontDistributionViewerCertificate{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudFrontDistributionViewerCertificateList(list)
		return nil
	}
	return err
}

// CloudWatchAlarmDimension represents the AWS::CloudWatch::Alarm.Dimension CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-dimension.html
type CloudWatchAlarmDimension struct {
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-dimension.html#cfn-cloudwatch-alarm-dimension-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// Value docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-dimension.html#cfn-cloudwatch-alarm-dimension-value
	Value *StringExpr `json:"Value,omitempty" validate:"dive,required"`
}

// CloudWatchAlarmDimensionList represents a list of CloudWatchAlarmDimension
type CloudWatchAlarmDimensionList []CloudWatchAlarmDimension

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CloudWatchAlarmDimensionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CloudWatchAlarmDimension{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CloudWatchAlarmDimensionList{item}
		return nil
	}
	list := []CloudWatchAlarmDimension{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CloudWatchAlarmDimensionList(list)
		return nil
	}
	return err
}

// CodeBuildProjectArtifacts represents the AWS::CodeBuild::Project.Artifacts CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-artifacts.html
type CodeBuildProjectArtifacts struct {
	// Location docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-artifacts.html#cfn-codebuild-project-artifacts-location
	Location *StringExpr `json:"Location,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-artifacts.html#cfn-codebuild-project-artifacts-name
	Name *StringExpr `json:"Name,omitempty"`
	// NamespaceType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-artifacts.html#cfn-codebuild-project-artifacts-namespacetype
	NamespaceType *StringExpr `json:"NamespaceType,omitempty"`
	// Packaging docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-artifacts.html#cfn-codebuild-project-artifacts-packaging
	Packaging *StringExpr `json:"Packaging,omitempty"`
	// Path docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-artifacts.html#cfn-codebuild-project-artifacts-path
	Path *StringExpr `json:"Path,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-artifacts.html#cfn-codebuild-project-artifacts-type
	Type *StringExpr `json:"Type,omitempty"`
}

// CodeBuildProjectArtifactsList represents a list of CodeBuildProjectArtifacts
type CodeBuildProjectArtifactsList []CodeBuildProjectArtifacts

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CodeBuildProjectArtifactsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CodeBuildProjectArtifacts{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CodeBuildProjectArtifactsList{item}
		return nil
	}
	list := []CodeBuildProjectArtifacts{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CodeBuildProjectArtifactsList(list)
		return nil
	}
	return err
}

// CodeBuildProjectEnvironment represents the AWS::CodeBuild::Project.Environment CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environment.html
type CodeBuildProjectEnvironment struct {
	// ComputeType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environment.html#cfn-codebuild-project-environment-computetype
	ComputeType *StringExpr `json:"ComputeType,omitempty"`
	// EnvironmentVariables docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environment.html#cfn-codebuild-project-environment-environmentvariables
	EnvironmentVariables *CodeBuildProjectEnvironmentVariableList `json:"EnvironmentVariables,omitempty"`
	// Image docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environment.html#cfn-codebuild-project-environment-image
	Image *StringExpr `json:"Image,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environment.html#cfn-codebuild-project-environment-type
	Type *StringExpr `json:"Type,omitempty"`
}

// CodeBuildProjectEnvironmentList represents a list of CodeBuildProjectEnvironment
type CodeBuildProjectEnvironmentList []CodeBuildProjectEnvironment

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CodeBuildProjectEnvironmentList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CodeBuildProjectEnvironment{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CodeBuildProjectEnvironmentList{item}
		return nil
	}
	list := []CodeBuildProjectEnvironment{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CodeBuildProjectEnvironmentList(list)
		return nil
	}
	return err
}

// CodeBuildProjectEnvironmentVariable represents the AWS::CodeBuild::Project.EnvironmentVariable CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environmentvariable.html
type CodeBuildProjectEnvironmentVariable struct {
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environmentvariable.html#cfn-codebuild-project-environmentvariable-name
	Name *StringExpr `json:"Name,omitempty"`
	// Value docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environmentvariable.html#cfn-codebuild-project-environmentvariable-value
	Value *StringExpr `json:"Value,omitempty"`
}

// CodeBuildProjectEnvironmentVariableList represents a list of CodeBuildProjectEnvironmentVariable
type CodeBuildProjectEnvironmentVariableList []CodeBuildProjectEnvironmentVariable

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CodeBuildProjectEnvironmentVariableList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CodeBuildProjectEnvironmentVariable{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CodeBuildProjectEnvironmentVariableList{item}
		return nil
	}
	list := []CodeBuildProjectEnvironmentVariable{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CodeBuildProjectEnvironmentVariableList(list)
		return nil
	}
	return err
}

// CodeBuildProjectSource represents the AWS::CodeBuild::Project.Source CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html
type CodeBuildProjectSource struct {
	// Auth docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html#cfn-codebuild-project-source-auth
	Auth *CodeBuildProjectSourceAuth `json:"Auth,omitempty"`
	// BuildSpec docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html#cfn-codebuild-project-source-buildspec
	BuildSpec *StringExpr `json:"BuildSpec,omitempty"`
	// Location docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html#cfn-codebuild-project-source-location
	Location *StringExpr `json:"Location,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html#cfn-codebuild-project-source-type
	Type *StringExpr `json:"Type,omitempty"`
}

// CodeBuildProjectSourceList represents a list of CodeBuildProjectSource
type CodeBuildProjectSourceList []CodeBuildProjectSource

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CodeBuildProjectSourceList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CodeBuildProjectSource{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CodeBuildProjectSourceList{item}
		return nil
	}
	list := []CodeBuildProjectSource{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CodeBuildProjectSourceList(list)
		return nil
	}
	return err
}

// CodeBuildProjectSourceAuth represents the AWS::CodeBuild::Project.SourceAuth CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-sourceauth.html
type CodeBuildProjectSourceAuth struct {
	// Resource docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-sourceauth.html#cfn-codebuild-project-sourceauth-resource
	Resource *StringExpr `json:"Resource,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-sourceauth.html#cfn-codebuild-project-sourceauth-type
	Type *StringExpr `json:"Type,omitempty"`
}

// CodeBuildProjectSourceAuthList represents a list of CodeBuildProjectSourceAuth
type CodeBuildProjectSourceAuthList []CodeBuildProjectSourceAuth

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CodeBuildProjectSourceAuthList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CodeBuildProjectSourceAuth{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CodeBuildProjectSourceAuthList{item}
		return nil
	}
	list := []CodeBuildProjectSourceAuth{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CodeBuildProjectSourceAuthList(list)
		return nil
	}
	return err
}

// CodeCommitRepositoryRepositoryTrigger represents the AWS::CodeCommit::Repository.RepositoryTrigger CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html
type CodeCommitRepositoryRepositoryTrigger struct {
	// Branches docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html#cfn-codecommit-repository-repositorytrigger-branches
	Branches *StringListExpr `json:"Branches,omitempty"`
	// CustomData docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html#cfn-codecommit-repository-repositorytrigger-customdata
	CustomData *StringExpr `json:"CustomData,omitempty"`
	// DestinationArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html#cfn-codecommit-repository-repositorytrigger-destinationarn
	DestinationArn *StringExpr `json:"DestinationArn,omitempty"`
	// Events docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html#cfn-codecommit-repository-repositorytrigger-events
	Events *StringListExpr `json:"Events,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html#cfn-codecommit-repository-repositorytrigger-name
	Name *StringExpr `json:"Name,omitempty"`
}

// CodeCommitRepositoryRepositoryTriggerList represents a list of CodeCommitRepositoryRepositoryTrigger
type CodeCommitRepositoryRepositoryTriggerList []CodeCommitRepositoryRepositoryTrigger

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CodeCommitRepositoryRepositoryTriggerList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CodeCommitRepositoryRepositoryTrigger{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CodeCommitRepositoryRepositoryTriggerList{item}
		return nil
	}
	list := []CodeCommitRepositoryRepositoryTrigger{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CodeCommitRepositoryRepositoryTriggerList(list)
		return nil
	}
	return err
}

// CodeDeployDeploymentConfigMinimumHealthyHosts represents the AWS::CodeDeploy::DeploymentConfig.MinimumHealthyHosts CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-minimumhealthyhosts.html
type CodeDeployDeploymentConfigMinimumHealthyHosts struct {
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-minimumhealthyhosts.html#cfn-codedeploy-deploymentconfig-minimumhealthyhosts-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
	// Value docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-minimumhealthyhosts.html#cfn-codedeploy-deploymentconfig-minimumhealthyhosts-value
	Value *IntegerExpr `json:"Value,omitempty" validate:"dive,required"`
}

// CodeDeployDeploymentConfigMinimumHealthyHostsList represents a list of CodeDeployDeploymentConfigMinimumHealthyHosts
type CodeDeployDeploymentConfigMinimumHealthyHostsList []CodeDeployDeploymentConfigMinimumHealthyHosts

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CodeDeployDeploymentConfigMinimumHealthyHostsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CodeDeployDeploymentConfigMinimumHealthyHosts{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CodeDeployDeploymentConfigMinimumHealthyHostsList{item}
		return nil
	}
	list := []CodeDeployDeploymentConfigMinimumHealthyHosts{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CodeDeployDeploymentConfigMinimumHealthyHostsList(list)
		return nil
	}
	return err
}

// CodeDeployDeploymentGroupAlarm represents the AWS::CodeDeploy::DeploymentGroup.Alarm CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-alarm.html
type CodeDeployDeploymentGroupAlarm struct {
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-alarm.html#cfn-codedeploy-deploymentgroup-alarm-name
	Name *StringExpr `json:"Name,omitempty"`
}

// CodeDeployDeploymentGroupAlarmList represents a list of CodeDeployDeploymentGroupAlarm
type CodeDeployDeploymentGroupAlarmList []CodeDeployDeploymentGroupAlarm

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CodeDeployDeploymentGroupAlarmList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CodeDeployDeploymentGroupAlarm{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CodeDeployDeploymentGroupAlarmList{item}
		return nil
	}
	list := []CodeDeployDeploymentGroupAlarm{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CodeDeployDeploymentGroupAlarmList(list)
		return nil
	}
	return err
}

// CodeDeployDeploymentGroupAlarmConfiguration represents the AWS::CodeDeploy::DeploymentGroup.AlarmConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-alarmconfiguration.html
type CodeDeployDeploymentGroupAlarmConfiguration struct {
	// Alarms docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-alarmconfiguration.html#cfn-codedeploy-deploymentgroup-alarmconfiguration-alarms
	Alarms *CodeDeployDeploymentGroupAlarmList `json:"Alarms,omitempty"`
	// Enabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-alarmconfiguration.html#cfn-codedeploy-deploymentgroup-alarmconfiguration-enabled
	Enabled *BoolExpr `json:"Enabled,omitempty"`
	// IgnorePollAlarmFailure docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-alarmconfiguration.html#cfn-codedeploy-deploymentgroup-alarmconfiguration-ignorepollalarmfailure
	IgnorePollAlarmFailure *BoolExpr `json:"IgnorePollAlarmFailure,omitempty"`
}

// CodeDeployDeploymentGroupAlarmConfigurationList represents a list of CodeDeployDeploymentGroupAlarmConfiguration
type CodeDeployDeploymentGroupAlarmConfigurationList []CodeDeployDeploymentGroupAlarmConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CodeDeployDeploymentGroupAlarmConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CodeDeployDeploymentGroupAlarmConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CodeDeployDeploymentGroupAlarmConfigurationList{item}
		return nil
	}
	list := []CodeDeployDeploymentGroupAlarmConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CodeDeployDeploymentGroupAlarmConfigurationList(list)
		return nil
	}
	return err
}

// CodeDeployDeploymentGroupDeployment represents the AWS::CodeDeploy::DeploymentGroup.Deployment CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment.html
type CodeDeployDeploymentGroupDeployment struct {
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment.html#cfn-properties-codedeploy-deploymentgroup-deployment-description
	Description *StringExpr `json:"Description,omitempty"`
	// IgnoreApplicationStopFailures docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment.html#cfn-properties-codedeploy-deploymentgroup-deployment-ignoreapplicationstopfailures
	IgnoreApplicationStopFailures *BoolExpr `json:"IgnoreApplicationStopFailures,omitempty"`
	// Revision docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment.html#cfn-properties-codedeploy-deploymentgroup-deployment-revision
	Revision *CodeDeployDeploymentGroupRevisionLocation `json:"Revision,omitempty" validate:"dive,required"`
}

// CodeDeployDeploymentGroupDeploymentList represents a list of CodeDeployDeploymentGroupDeployment
type CodeDeployDeploymentGroupDeploymentList []CodeDeployDeploymentGroupDeployment

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CodeDeployDeploymentGroupDeploymentList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CodeDeployDeploymentGroupDeployment{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CodeDeployDeploymentGroupDeploymentList{item}
		return nil
	}
	list := []CodeDeployDeploymentGroupDeployment{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CodeDeployDeploymentGroupDeploymentList(list)
		return nil
	}
	return err
}

// CodeDeployDeploymentGroupEC2TagFilter represents the AWS::CodeDeploy::DeploymentGroup.EC2TagFilter CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ec2tagfilters.html
type CodeDeployDeploymentGroupEC2TagFilter struct {
	// Key docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ec2tagfilters.html#cfn-properties-codedeploy-deploymentgroup-ec2tagfilters-key
	Key *StringExpr `json:"Key,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ec2tagfilters.html#cfn-properties-codedeploy-deploymentgroup-ec2tagfilters-type
	Type *StringExpr `json:"Type,omitempty"`
	// Value docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ec2tagfilters.html#cfn-properties-codedeploy-deploymentgroup-ec2tagfilters-value
	Value *StringExpr `json:"Value,omitempty"`
}

// CodeDeployDeploymentGroupEC2TagFilterList represents a list of CodeDeployDeploymentGroupEC2TagFilter
type CodeDeployDeploymentGroupEC2TagFilterList []CodeDeployDeploymentGroupEC2TagFilter

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CodeDeployDeploymentGroupEC2TagFilterList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CodeDeployDeploymentGroupEC2TagFilter{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CodeDeployDeploymentGroupEC2TagFilterList{item}
		return nil
	}
	list := []CodeDeployDeploymentGroupEC2TagFilter{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CodeDeployDeploymentGroupEC2TagFilterList(list)
		return nil
	}
	return err
}

// CodeDeployDeploymentGroupGitHubLocation represents the AWS::CodeDeploy::DeploymentGroup.GitHubLocation CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision-githublocation.html
type CodeDeployDeploymentGroupGitHubLocation struct {
	// CommitID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision-githublocation.html#cfn-properties-codedeploy-deploymentgroup-deployment-revision-githublocation-commitid
	CommitID *StringExpr `json:"CommitId,omitempty" validate:"dive,required"`
	// Repository docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision-githublocation.html#cfn-properties-codedeploy-deploymentgroup-deployment-revision-githublocation-repository
	Repository *StringExpr `json:"Repository,omitempty" validate:"dive,required"`
}

// CodeDeployDeploymentGroupGitHubLocationList represents a list of CodeDeployDeploymentGroupGitHubLocation
type CodeDeployDeploymentGroupGitHubLocationList []CodeDeployDeploymentGroupGitHubLocation

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CodeDeployDeploymentGroupGitHubLocationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CodeDeployDeploymentGroupGitHubLocation{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CodeDeployDeploymentGroupGitHubLocationList{item}
		return nil
	}
	list := []CodeDeployDeploymentGroupGitHubLocation{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CodeDeployDeploymentGroupGitHubLocationList(list)
		return nil
	}
	return err
}

// CodeDeployDeploymentGroupRevisionLocation represents the AWS::CodeDeploy::DeploymentGroup.RevisionLocation CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision.html
type CodeDeployDeploymentGroupRevisionLocation struct {
	// GitHubLocation docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision.html#cfn-properties-codedeploy-deploymentgroup-deployment-revision-githublocation
	GitHubLocation *CodeDeployDeploymentGroupGitHubLocation `json:"GitHubLocation,omitempty"`
	// RevisionType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision.html#cfn-properties-codedeploy-deploymentgroup-deployment-revision-revisiontype
	RevisionType *StringExpr `json:"RevisionType,omitempty"`
	// S3Location docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision.html#cfn-properties-codedeploy-deploymentgroup-deployment-revision-s3location
	S3Location *CodeDeployDeploymentGroupS3Location `json:"S3Location,omitempty"`
}

// CodeDeployDeploymentGroupRevisionLocationList represents a list of CodeDeployDeploymentGroupRevisionLocation
type CodeDeployDeploymentGroupRevisionLocationList []CodeDeployDeploymentGroupRevisionLocation

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CodeDeployDeploymentGroupRevisionLocationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CodeDeployDeploymentGroupRevisionLocation{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CodeDeployDeploymentGroupRevisionLocationList{item}
		return nil
	}
	list := []CodeDeployDeploymentGroupRevisionLocation{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CodeDeployDeploymentGroupRevisionLocationList(list)
		return nil
	}
	return err
}

// CodeDeployDeploymentGroupS3Location represents the AWS::CodeDeploy::DeploymentGroup.S3Location CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision-s3location.html
type CodeDeployDeploymentGroupS3Location struct {
	// Bucket docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision-s3location.html#cfn-properties-codedeploy-deploymentgroup-deployment-revision-s3location-bucket
	Bucket *StringExpr `json:"Bucket,omitempty" validate:"dive,required"`
	// BundleType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision-s3location.html#cfn-properties-codedeploy-deploymentgroup-deployment-revision-s3location-bundletype
	BundleType *StringExpr `json:"BundleType,omitempty"`
	// ETag docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision-s3location.html#cfn-properties-codedeploy-deploymentgroup-deployment-revision-s3location-etag
	ETag *StringExpr `json:"ETag,omitempty"`
	// Key docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision-s3location.html#cfn-properties-codedeploy-deploymentgroup-deployment-revision-s3location-key
	Key *StringExpr `json:"Key,omitempty" validate:"dive,required"`
	// Version docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision-s3location.html#cfn-properties-codedeploy-deploymentgroup-deployment-revision-s3location-value
	Version *StringExpr `json:"Version,omitempty"`
}

// CodeDeployDeploymentGroupS3LocationList represents a list of CodeDeployDeploymentGroupS3Location
type CodeDeployDeploymentGroupS3LocationList []CodeDeployDeploymentGroupS3Location

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CodeDeployDeploymentGroupS3LocationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CodeDeployDeploymentGroupS3Location{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CodeDeployDeploymentGroupS3LocationList{item}
		return nil
	}
	list := []CodeDeployDeploymentGroupS3Location{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CodeDeployDeploymentGroupS3LocationList(list)
		return nil
	}
	return err
}

// CodeDeployDeploymentGroupTagFilter represents the AWS::CodeDeploy::DeploymentGroup.TagFilter CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-onpremisesinstancetagfilters.html
type CodeDeployDeploymentGroupTagFilter struct {
	// Key docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-onpremisesinstancetagfilters.html#cfn-properties-codedeploy-deploymentgroup-onpremisesinstancetagfilters-key
	Key *StringExpr `json:"Key,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-onpremisesinstancetagfilters.html#cfn-properties-codedeploy-deploymentgroup-onpremisesinstancetagfilters-type
	Type *StringExpr `json:"Type,omitempty"`
	// Value docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-onpremisesinstancetagfilters.html#cfn-properties-codedeploy-deploymentgroup-onpremisesinstancetagfilters-value
	Value *StringExpr `json:"Value,omitempty"`
}

// CodeDeployDeploymentGroupTagFilterList represents a list of CodeDeployDeploymentGroupTagFilter
type CodeDeployDeploymentGroupTagFilterList []CodeDeployDeploymentGroupTagFilter

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CodeDeployDeploymentGroupTagFilterList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CodeDeployDeploymentGroupTagFilter{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CodeDeployDeploymentGroupTagFilterList{item}
		return nil
	}
	list := []CodeDeployDeploymentGroupTagFilter{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CodeDeployDeploymentGroupTagFilterList(list)
		return nil
	}
	return err
}

// CodeDeployDeploymentGroupTriggerConfig represents the AWS::CodeDeploy::DeploymentGroup.TriggerConfig CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-triggerconfig.html
type CodeDeployDeploymentGroupTriggerConfig struct {
	// TriggerEvents docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-triggerconfig.html#cfn-codedeploy-deploymentgroup-triggerconfig-triggerevents
	TriggerEvents *StringListExpr `json:"TriggerEvents,omitempty"`
	// TriggerName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-triggerconfig.html#cfn-codedeploy-deploymentgroup-triggerconfig-triggername
	TriggerName *StringExpr `json:"TriggerName,omitempty"`
	// TriggerTargetArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-triggerconfig.html#cfn-codedeploy-deploymentgroup-triggerconfig-triggertargetarn
	TriggerTargetArn *StringExpr `json:"TriggerTargetArn,omitempty"`
}

// CodeDeployDeploymentGroupTriggerConfigList represents a list of CodeDeployDeploymentGroupTriggerConfig
type CodeDeployDeploymentGroupTriggerConfigList []CodeDeployDeploymentGroupTriggerConfig

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CodeDeployDeploymentGroupTriggerConfigList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CodeDeployDeploymentGroupTriggerConfig{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CodeDeployDeploymentGroupTriggerConfigList{item}
		return nil
	}
	list := []CodeDeployDeploymentGroupTriggerConfig{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CodeDeployDeploymentGroupTriggerConfigList(list)
		return nil
	}
	return err
}

// CodePipelineCustomActionTypeArtifactDetails represents the AWS::CodePipeline::CustomActionType.ArtifactDetails CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-artifactdetails.html
type CodePipelineCustomActionTypeArtifactDetails struct {
	// MaximumCount docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-artifactdetails.html#cfn-codepipeline-customactiontype-artifactdetails-maximumcount
	MaximumCount *IntegerExpr `json:"MaximumCount,omitempty" validate:"dive,required"`
	// MinimumCount docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-artifactdetails.html#cfn-codepipeline-customactiontype-artifactdetails-minimumcount
	MinimumCount *IntegerExpr `json:"MinimumCount,omitempty" validate:"dive,required"`
}

// CodePipelineCustomActionTypeArtifactDetailsList represents a list of CodePipelineCustomActionTypeArtifactDetails
type CodePipelineCustomActionTypeArtifactDetailsList []CodePipelineCustomActionTypeArtifactDetails

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CodePipelineCustomActionTypeArtifactDetailsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CodePipelineCustomActionTypeArtifactDetails{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CodePipelineCustomActionTypeArtifactDetailsList{item}
		return nil
	}
	list := []CodePipelineCustomActionTypeArtifactDetails{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CodePipelineCustomActionTypeArtifactDetailsList(list)
		return nil
	}
	return err
}

// CodePipelineCustomActionTypeConfigurationProperties represents the AWS::CodePipeline::CustomActionType.ConfigurationProperties CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-configurationproperties.html
type CodePipelineCustomActionTypeConfigurationProperties struct {
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-configurationproperties.html#cfn-codepipeline-customactiontype-configurationproperties-description
	Description *StringExpr `json:"Description,omitempty"`
	// Key docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-configurationproperties.html#cfn-codepipeline-customactiontype-configurationproperties-key
	Key *BoolExpr `json:"Key,omitempty" validate:"dive,required"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-configurationproperties.html#cfn-codepipeline-customactiontype-configurationproperties-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// Queryable docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-configurationproperties.html#cfn-codepipeline-customactiontype-configurationproperties-queryable
	Queryable *BoolExpr `json:"Queryable,omitempty"`
	// Required docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-configurationproperties.html#cfn-codepipeline-customactiontype-configurationproperties-required
	Required *BoolExpr `json:"Required,omitempty" validate:"dive,required"`
	// Secret docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-configurationproperties.html#cfn-codepipeline-customactiontype-configurationproperties-secret
	Secret *BoolExpr `json:"Secret,omitempty" validate:"dive,required"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-configurationproperties.html#cfn-codepipeline-customactiontype-configurationproperties-type
	Type *StringExpr `json:"Type,omitempty"`
}

// CodePipelineCustomActionTypeConfigurationPropertiesList represents a list of CodePipelineCustomActionTypeConfigurationProperties
type CodePipelineCustomActionTypeConfigurationPropertiesList []CodePipelineCustomActionTypeConfigurationProperties

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CodePipelineCustomActionTypeConfigurationPropertiesList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CodePipelineCustomActionTypeConfigurationProperties{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CodePipelineCustomActionTypeConfigurationPropertiesList{item}
		return nil
	}
	list := []CodePipelineCustomActionTypeConfigurationProperties{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CodePipelineCustomActionTypeConfigurationPropertiesList(list)
		return nil
	}
	return err
}

// CodePipelineCustomActionTypeSettings represents the AWS::CodePipeline::CustomActionType.Settings CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-settings.html
type CodePipelineCustomActionTypeSettings struct {
	// EntityURLTemplate docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-settings.html#cfn-codepipeline-customactiontype-settings-entityurltemplate
	EntityURLTemplate *StringExpr `json:"EntityUrlTemplate,omitempty"`
	// ExecutionURLTemplate docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-settings.html#cfn-codepipeline-customactiontype-settings-executionurltemplate
	ExecutionURLTemplate *StringExpr `json:"ExecutionUrlTemplate,omitempty"`
	// RevisionURLTemplate docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-settings.html#cfn-codepipeline-customactiontype-settings-revisionurltemplate
	RevisionURLTemplate *StringExpr `json:"RevisionUrlTemplate,omitempty"`
	// ThirdPartyConfigurationURL docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-settings.html#cfn-codepipeline-customactiontype-settings-thirdpartyconfigurationurl
	ThirdPartyConfigurationURL *StringExpr `json:"ThirdPartyConfigurationUrl,omitempty"`
}

// CodePipelineCustomActionTypeSettingsList represents a list of CodePipelineCustomActionTypeSettings
type CodePipelineCustomActionTypeSettingsList []CodePipelineCustomActionTypeSettings

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CodePipelineCustomActionTypeSettingsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CodePipelineCustomActionTypeSettings{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CodePipelineCustomActionTypeSettingsList{item}
		return nil
	}
	list := []CodePipelineCustomActionTypeSettings{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CodePipelineCustomActionTypeSettingsList(list)
		return nil
	}
	return err
}

// CodePipelinePipelineActionDeclaration represents the AWS::CodePipeline::Pipeline.ActionDeclaration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions.html
type CodePipelinePipelineActionDeclaration struct {
	// ActionTypeID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions.html#cfn-codepipeline-pipeline-stages-actions-actiontypeid
	ActionTypeID *CodePipelinePipelineActionTypeID `json:"ActionTypeId,omitempty" validate:"dive,required"`
	// Configuration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions.html#cfn-codepipeline-pipeline-stages-actions-configuration
	Configuration interface{} `json:"Configuration,omitempty"`
	// InputArtifacts docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions.html#cfn-codepipeline-pipeline-stages-actions-inputartifacts
	InputArtifacts *CodePipelinePipelineInputArtifactList `json:"InputArtifacts,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions.html#cfn-codepipeline-pipeline-stages-actions-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// OutputArtifacts docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions.html#cfn-codepipeline-pipeline-stages-actions-outputartifacts
	OutputArtifacts *CodePipelinePipelineOutputArtifactList `json:"OutputArtifacts,omitempty"`
	// RoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions.html#cfn-codepipeline-pipeline-stages-actions-rolearn
	RoleArn *StringExpr `json:"RoleArn,omitempty"`
	// RunOrder docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions.html#cfn-codepipeline-pipeline-stages-actions-runorder
	RunOrder *IntegerExpr `json:"RunOrder,omitempty"`
}

// CodePipelinePipelineActionDeclarationList represents a list of CodePipelinePipelineActionDeclaration
type CodePipelinePipelineActionDeclarationList []CodePipelinePipelineActionDeclaration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CodePipelinePipelineActionDeclarationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CodePipelinePipelineActionDeclaration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CodePipelinePipelineActionDeclarationList{item}
		return nil
	}
	list := []CodePipelinePipelineActionDeclaration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CodePipelinePipelineActionDeclarationList(list)
		return nil
	}
	return err
}

// CodePipelinePipelineActionTypeID represents the AWS::CodePipeline::Pipeline.ActionTypeId CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions-actiontypeid.html
type CodePipelinePipelineActionTypeID struct {
	// Category docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions-actiontypeid.html#cfn-codepipeline-pipeline-stages-actions-actiontypeid-category
	Category *StringExpr `json:"Category,omitempty" validate:"dive,required"`
	// Owner docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions-actiontypeid.html#cfn-codepipeline-pipeline-stages-actions-actiontypeid-owner
	Owner *StringExpr `json:"Owner,omitempty" validate:"dive,required"`
	// Provider docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions-actiontypeid.html#cfn-codepipeline-pipeline-stages-actions-actiontypeid-provider
	Provider *StringExpr `json:"Provider,omitempty" validate:"dive,required"`
	// Version docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions-actiontypeid.html#cfn-codepipeline-pipeline-stages-actions-actiontypeid-version
	Version *StringExpr `json:"Version,omitempty" validate:"dive,required"`
}

// CodePipelinePipelineActionTypeIDList represents a list of CodePipelinePipelineActionTypeID
type CodePipelinePipelineActionTypeIDList []CodePipelinePipelineActionTypeID

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CodePipelinePipelineActionTypeIDList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CodePipelinePipelineActionTypeID{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CodePipelinePipelineActionTypeIDList{item}
		return nil
	}
	list := []CodePipelinePipelineActionTypeID{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CodePipelinePipelineActionTypeIDList(list)
		return nil
	}
	return err
}

// CodePipelinePipelineArtifactStore represents the AWS::CodePipeline::Pipeline.ArtifactStore CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore.html
type CodePipelinePipelineArtifactStore struct {
	// EncryptionKey docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore.html#cfn-codepipeline-pipeline-artifactstore-encryptionkey
	EncryptionKey *CodePipelinePipelineEncryptionKey `json:"EncryptionKey,omitempty"`
	// Location docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore.html#cfn-codepipeline-pipeline-artifactstore-location
	Location *StringExpr `json:"Location,omitempty" validate:"dive,required"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore.html#cfn-codepipeline-pipeline-artifactstore-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
}

// CodePipelinePipelineArtifactStoreList represents a list of CodePipelinePipelineArtifactStore
type CodePipelinePipelineArtifactStoreList []CodePipelinePipelineArtifactStore

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CodePipelinePipelineArtifactStoreList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CodePipelinePipelineArtifactStore{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CodePipelinePipelineArtifactStoreList{item}
		return nil
	}
	list := []CodePipelinePipelineArtifactStore{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CodePipelinePipelineArtifactStoreList(list)
		return nil
	}
	return err
}

// CodePipelinePipelineBlockerDeclaration represents the AWS::CodePipeline::Pipeline.BlockerDeclaration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-blockers.html
type CodePipelinePipelineBlockerDeclaration struct {
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-blockers.html#cfn-codepipeline-pipeline-stages-blockers-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-blockers.html#cfn-codepipeline-pipeline-stages-blockers-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
}

// CodePipelinePipelineBlockerDeclarationList represents a list of CodePipelinePipelineBlockerDeclaration
type CodePipelinePipelineBlockerDeclarationList []CodePipelinePipelineBlockerDeclaration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CodePipelinePipelineBlockerDeclarationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CodePipelinePipelineBlockerDeclaration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CodePipelinePipelineBlockerDeclarationList{item}
		return nil
	}
	list := []CodePipelinePipelineBlockerDeclaration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CodePipelinePipelineBlockerDeclarationList(list)
		return nil
	}
	return err
}

// CodePipelinePipelineEncryptionKey represents the AWS::CodePipeline::Pipeline.EncryptionKey CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore-encryptionkey.html
type CodePipelinePipelineEncryptionKey struct {
	// ID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore-encryptionkey.html#cfn-codepipeline-pipeline-artifactstore-encryptionkey-id
	ID *StringExpr `json:"Id,omitempty" validate:"dive,required"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore-encryptionkey.html#cfn-codepipeline-pipeline-artifactstore-encryptionkey-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
}

// CodePipelinePipelineEncryptionKeyList represents a list of CodePipelinePipelineEncryptionKey
type CodePipelinePipelineEncryptionKeyList []CodePipelinePipelineEncryptionKey

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CodePipelinePipelineEncryptionKeyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CodePipelinePipelineEncryptionKey{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CodePipelinePipelineEncryptionKeyList{item}
		return nil
	}
	list := []CodePipelinePipelineEncryptionKey{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CodePipelinePipelineEncryptionKeyList(list)
		return nil
	}
	return err
}

// CodePipelinePipelineInputArtifact represents the AWS::CodePipeline::Pipeline.InputArtifact CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions-inputartifacts.html
type CodePipelinePipelineInputArtifact struct {
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions-inputartifacts.html#cfn-codepipeline-pipeline-stages-actions-inputartifacts-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
}

// CodePipelinePipelineInputArtifactList represents a list of CodePipelinePipelineInputArtifact
type CodePipelinePipelineInputArtifactList []CodePipelinePipelineInputArtifact

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CodePipelinePipelineInputArtifactList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CodePipelinePipelineInputArtifact{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CodePipelinePipelineInputArtifactList{item}
		return nil
	}
	list := []CodePipelinePipelineInputArtifact{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CodePipelinePipelineInputArtifactList(list)
		return nil
	}
	return err
}

// CodePipelinePipelineOutputArtifact represents the AWS::CodePipeline::Pipeline.OutputArtifact CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions-outputartifacts.html
type CodePipelinePipelineOutputArtifact struct {
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions-outputartifacts.html#cfn-codepipeline-pipeline-stages-actions-outputartifacts-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
}

// CodePipelinePipelineOutputArtifactList represents a list of CodePipelinePipelineOutputArtifact
type CodePipelinePipelineOutputArtifactList []CodePipelinePipelineOutputArtifact

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CodePipelinePipelineOutputArtifactList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CodePipelinePipelineOutputArtifact{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CodePipelinePipelineOutputArtifactList{item}
		return nil
	}
	list := []CodePipelinePipelineOutputArtifact{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CodePipelinePipelineOutputArtifactList(list)
		return nil
	}
	return err
}

// CodePipelinePipelineStageDeclaration represents the AWS::CodePipeline::Pipeline.StageDeclaration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages.html
type CodePipelinePipelineStageDeclaration struct {
	// Actions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages.html#cfn-codepipeline-pipeline-stages-actions
	Actions *CodePipelinePipelineActionDeclarationList `json:"Actions,omitempty" validate:"dive,required"`
	// Blockers docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages.html#cfn-codepipeline-pipeline-stages-blockers
	Blockers *CodePipelinePipelineBlockerDeclarationList `json:"Blockers,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages.html#cfn-codepipeline-pipeline-stages-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
}

// CodePipelinePipelineStageDeclarationList represents a list of CodePipelinePipelineStageDeclaration
type CodePipelinePipelineStageDeclarationList []CodePipelinePipelineStageDeclaration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CodePipelinePipelineStageDeclarationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CodePipelinePipelineStageDeclaration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CodePipelinePipelineStageDeclarationList{item}
		return nil
	}
	list := []CodePipelinePipelineStageDeclaration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CodePipelinePipelineStageDeclarationList(list)
		return nil
	}
	return err
}

// CodePipelinePipelineStageTransition represents the AWS::CodePipeline::Pipeline.StageTransition CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-disableinboundstagetransitions.html
type CodePipelinePipelineStageTransition struct {
	// Reason docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-disableinboundstagetransitions.html#cfn-codepipeline-pipeline-disableinboundstagetransitions-reason
	Reason *StringExpr `json:"Reason,omitempty" validate:"dive,required"`
	// StageName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-disableinboundstagetransitions.html#cfn-codepipeline-pipeline-disableinboundstagetransitions-stagename
	StageName *StringExpr `json:"StageName,omitempty" validate:"dive,required"`
}

// CodePipelinePipelineStageTransitionList represents a list of CodePipelinePipelineStageTransition
type CodePipelinePipelineStageTransitionList []CodePipelinePipelineStageTransition

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CodePipelinePipelineStageTransitionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CodePipelinePipelineStageTransition{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CodePipelinePipelineStageTransitionList{item}
		return nil
	}
	list := []CodePipelinePipelineStageTransition{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CodePipelinePipelineStageTransitionList(list)
		return nil
	}
	return err
}

// CognitoIDentityPoolCognitoIDentityProvider represents the AWS::Cognito::IdentityPool.CognitoIdentityProvider CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-cognitoidentityprovider.html
type CognitoIDentityPoolCognitoIDentityProvider struct {
	// ClientID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-cognitoidentityprovider.html#cfn-cognito-identitypool-cognitoidentityprovider-clientid
	ClientID *StringExpr `json:"ClientId,omitempty"`
	// ProviderName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-cognitoidentityprovider.html#cfn-cognito-identitypool-cognitoidentityprovider-providername
	ProviderName *StringExpr `json:"ProviderName,omitempty"`
	// ServerSideTokenCheck docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-cognitoidentityprovider.html#cfn-cognito-identitypool-cognitoidentityprovider-serversidetokencheck
	ServerSideTokenCheck *BoolExpr `json:"ServerSideTokenCheck,omitempty"`
}

// CognitoIDentityPoolCognitoIDentityProviderList represents a list of CognitoIDentityPoolCognitoIDentityProvider
type CognitoIDentityPoolCognitoIDentityProviderList []CognitoIDentityPoolCognitoIDentityProvider

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CognitoIDentityPoolCognitoIDentityProviderList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CognitoIDentityPoolCognitoIDentityProvider{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CognitoIDentityPoolCognitoIDentityProviderList{item}
		return nil
	}
	list := []CognitoIDentityPoolCognitoIDentityProvider{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CognitoIDentityPoolCognitoIDentityProviderList(list)
		return nil
	}
	return err
}

// CognitoIDentityPoolCognitoStreams represents the AWS::Cognito::IdentityPool.CognitoStreams CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-cognitostreams.html
type CognitoIDentityPoolCognitoStreams struct {
	// RoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-cognitostreams.html#cfn-cognito-identitypool-cognitostreams-rolearn
	RoleArn *StringExpr `json:"RoleArn,omitempty"`
	// StreamName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-cognitostreams.html#cfn-cognito-identitypool-cognitostreams-streamname
	StreamName *StringExpr `json:"StreamName,omitempty"`
	// StreamingStatus docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-cognitostreams.html#cfn-cognito-identitypool-cognitostreams-streamingstatus
	StreamingStatus *StringExpr `json:"StreamingStatus,omitempty"`
}

// CognitoIDentityPoolCognitoStreamsList represents a list of CognitoIDentityPoolCognitoStreams
type CognitoIDentityPoolCognitoStreamsList []CognitoIDentityPoolCognitoStreams

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CognitoIDentityPoolCognitoStreamsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CognitoIDentityPoolCognitoStreams{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CognitoIDentityPoolCognitoStreamsList{item}
		return nil
	}
	list := []CognitoIDentityPoolCognitoStreams{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CognitoIDentityPoolCognitoStreamsList(list)
		return nil
	}
	return err
}

// CognitoIDentityPoolPushSync represents the AWS::Cognito::IdentityPool.PushSync CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-pushsync.html
type CognitoIDentityPoolPushSync struct {
	// ApplicationArns docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-pushsync.html#cfn-cognito-identitypool-pushsync-applicationarns
	ApplicationArns *StringListExpr `json:"ApplicationArns,omitempty"`
	// RoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-pushsync.html#cfn-cognito-identitypool-pushsync-rolearn
	RoleArn *StringExpr `json:"RoleArn,omitempty"`
}

// CognitoIDentityPoolPushSyncList represents a list of CognitoIDentityPoolPushSync
type CognitoIDentityPoolPushSyncList []CognitoIDentityPoolPushSync

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CognitoIDentityPoolPushSyncList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CognitoIDentityPoolPushSync{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CognitoIDentityPoolPushSyncList{item}
		return nil
	}
	list := []CognitoIDentityPoolPushSync{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CognitoIDentityPoolPushSyncList(list)
		return nil
	}
	return err
}

// CognitoIDentityPoolRoleAttachmentMappingRule represents the AWS::Cognito::IdentityPoolRoleAttachment.MappingRule CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-mappingrule.html
type CognitoIDentityPoolRoleAttachmentMappingRule struct {
	// Claim docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-mappingrule.html#cfn-cognito-identitypoolroleattachment-mappingrule-claim
	Claim *StringExpr `json:"Claim,omitempty" validate:"dive,required"`
	// MatchType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-mappingrule.html#cfn-cognito-identitypoolroleattachment-mappingrule-matchtype
	MatchType *StringExpr `json:"MatchType,omitempty" validate:"dive,required"`
	// RoleARN docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-mappingrule.html#cfn-cognito-identitypoolroleattachment-mappingrule-rolearn
	RoleARN *StringExpr `json:"RoleARN,omitempty" validate:"dive,required"`
	// Value docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-mappingrule.html#cfn-cognito-identitypoolroleattachment-mappingrule-value
	Value *StringExpr `json:"Value,omitempty" validate:"dive,required"`
}

// CognitoIDentityPoolRoleAttachmentMappingRuleList represents a list of CognitoIDentityPoolRoleAttachmentMappingRule
type CognitoIDentityPoolRoleAttachmentMappingRuleList []CognitoIDentityPoolRoleAttachmentMappingRule

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CognitoIDentityPoolRoleAttachmentMappingRuleList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CognitoIDentityPoolRoleAttachmentMappingRule{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CognitoIDentityPoolRoleAttachmentMappingRuleList{item}
		return nil
	}
	list := []CognitoIDentityPoolRoleAttachmentMappingRule{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CognitoIDentityPoolRoleAttachmentMappingRuleList(list)
		return nil
	}
	return err
}

// CognitoIDentityPoolRoleAttachmentRoleMapping represents the AWS::Cognito::IdentityPoolRoleAttachment.RoleMapping CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rolemapping.html
type CognitoIDentityPoolRoleAttachmentRoleMapping struct {
	// AmbiguousRoleResolution docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rolemapping.html#cfn-cognito-identitypoolroleattachment-rolemapping-ambiguousroleresolution
	AmbiguousRoleResolution *StringExpr `json:"AmbiguousRoleResolution,omitempty"`
	// RulesConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rolemapping.html#cfn-cognito-identitypoolroleattachment-rolemapping-rulesconfiguration
	RulesConfiguration *CognitoIDentityPoolRoleAttachmentRulesConfigurationType `json:"RulesConfiguration,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rolemapping.html#cfn-cognito-identitypoolroleattachment-rolemapping-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
}

// CognitoIDentityPoolRoleAttachmentRoleMappingList represents a list of CognitoIDentityPoolRoleAttachmentRoleMapping
type CognitoIDentityPoolRoleAttachmentRoleMappingList []CognitoIDentityPoolRoleAttachmentRoleMapping

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CognitoIDentityPoolRoleAttachmentRoleMappingList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CognitoIDentityPoolRoleAttachmentRoleMapping{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CognitoIDentityPoolRoleAttachmentRoleMappingList{item}
		return nil
	}
	list := []CognitoIDentityPoolRoleAttachmentRoleMapping{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CognitoIDentityPoolRoleAttachmentRoleMappingList(list)
		return nil
	}
	return err
}

// CognitoIDentityPoolRoleAttachmentRulesConfigurationType represents the AWS::Cognito::IdentityPoolRoleAttachment.RulesConfigurationType CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rulesconfigurationtype.html
type CognitoIDentityPoolRoleAttachmentRulesConfigurationType struct {
}

// CognitoIDentityPoolRoleAttachmentRulesConfigurationTypeList represents a list of CognitoIDentityPoolRoleAttachmentRulesConfigurationType
type CognitoIDentityPoolRoleAttachmentRulesConfigurationTypeList []CognitoIDentityPoolRoleAttachmentRulesConfigurationType

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CognitoIDentityPoolRoleAttachmentRulesConfigurationTypeList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CognitoIDentityPoolRoleAttachmentRulesConfigurationType{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CognitoIDentityPoolRoleAttachmentRulesConfigurationTypeList{item}
		return nil
	}
	list := []CognitoIDentityPoolRoleAttachmentRulesConfigurationType{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CognitoIDentityPoolRoleAttachmentRulesConfigurationTypeList(list)
		return nil
	}
	return err
}

// CognitoUserPoolAdminCreateUserConfig represents the AWS::Cognito::UserPool.AdminCreateUserConfig CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-admincreateuserconfig.html
type CognitoUserPoolAdminCreateUserConfig struct {
	// AllowAdminCreateUserOnly docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-admincreateuserconfig.html#cfn-cognito-userpool-admincreateuserconfig-allowadmincreateuseronly
	AllowAdminCreateUserOnly *BoolExpr `json:"AllowAdminCreateUserOnly,omitempty"`
	// InviteMessageTemplate docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-admincreateuserconfig.html#cfn-cognito-userpool-admincreateuserconfig-invitemessagetemplate
	InviteMessageTemplate *CognitoUserPoolInviteMessageTemplate `json:"InviteMessageTemplate,omitempty"`
	// UnusedAccountValidityDays docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-admincreateuserconfig.html#cfn-cognito-userpool-admincreateuserconfig-unusedaccountvaliditydays
	UnusedAccountValidityDays *IntegerExpr `json:"UnusedAccountValidityDays,omitempty"`
}

// CognitoUserPoolAdminCreateUserConfigList represents a list of CognitoUserPoolAdminCreateUserConfig
type CognitoUserPoolAdminCreateUserConfigList []CognitoUserPoolAdminCreateUserConfig

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CognitoUserPoolAdminCreateUserConfigList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CognitoUserPoolAdminCreateUserConfig{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CognitoUserPoolAdminCreateUserConfigList{item}
		return nil
	}
	list := []CognitoUserPoolAdminCreateUserConfig{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CognitoUserPoolAdminCreateUserConfigList(list)
		return nil
	}
	return err
}

// CognitoUserPoolDeviceConfiguration represents the AWS::Cognito::UserPool.DeviceConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-deviceconfiguration.html
type CognitoUserPoolDeviceConfiguration struct {
	// ChallengeRequiredOnNewDevice docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-deviceconfiguration.html#cfn-cognito-userpool-deviceconfiguration-challengerequiredonnewdevice
	ChallengeRequiredOnNewDevice *BoolExpr `json:"ChallengeRequiredOnNewDevice,omitempty"`
	// DeviceOnlyRememberedOnUserPrompt docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-deviceconfiguration.html#cfn-cognito-userpool-deviceconfiguration-deviceonlyrememberedonuserprompt
	DeviceOnlyRememberedOnUserPrompt *BoolExpr `json:"DeviceOnlyRememberedOnUserPrompt,omitempty"`
}

// CognitoUserPoolDeviceConfigurationList represents a list of CognitoUserPoolDeviceConfiguration
type CognitoUserPoolDeviceConfigurationList []CognitoUserPoolDeviceConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CognitoUserPoolDeviceConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CognitoUserPoolDeviceConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CognitoUserPoolDeviceConfigurationList{item}
		return nil
	}
	list := []CognitoUserPoolDeviceConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CognitoUserPoolDeviceConfigurationList(list)
		return nil
	}
	return err
}

// CognitoUserPoolEmailConfiguration represents the AWS::Cognito::UserPool.EmailConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-emailconfiguration.html
type CognitoUserPoolEmailConfiguration struct {
	// ReplyToEmailAddress docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-emailconfiguration.html#cfn-cognito-userpool-emailconfiguration-replytoemailaddress
	ReplyToEmailAddress *StringExpr `json:"ReplyToEmailAddress,omitempty"`
	// SourceArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-emailconfiguration.html#cfn-cognito-userpool-emailconfiguration-sourcearn
	SourceArn *StringExpr `json:"SourceArn,omitempty"`
}

// CognitoUserPoolEmailConfigurationList represents a list of CognitoUserPoolEmailConfiguration
type CognitoUserPoolEmailConfigurationList []CognitoUserPoolEmailConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CognitoUserPoolEmailConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CognitoUserPoolEmailConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CognitoUserPoolEmailConfigurationList{item}
		return nil
	}
	list := []CognitoUserPoolEmailConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CognitoUserPoolEmailConfigurationList(list)
		return nil
	}
	return err
}

// CognitoUserPoolInviteMessageTemplate represents the AWS::Cognito::UserPool.InviteMessageTemplate CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-invitemessagetemplate.html
type CognitoUserPoolInviteMessageTemplate struct {
	// EmailMessage docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-invitemessagetemplate.html#cfn-cognito-userpool-invitemessagetemplate-emailmessage
	EmailMessage *StringExpr `json:"EmailMessage,omitempty"`
	// EmailSubject docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-invitemessagetemplate.html#cfn-cognito-userpool-invitemessagetemplate-emailsubject
	EmailSubject *StringExpr `json:"EmailSubject,omitempty"`
	// SMSMessage docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-invitemessagetemplate.html#cfn-cognito-userpool-invitemessagetemplate-smsmessage
	SMSMessage *StringExpr `json:"SMSMessage,omitempty"`
}

// CognitoUserPoolInviteMessageTemplateList represents a list of CognitoUserPoolInviteMessageTemplate
type CognitoUserPoolInviteMessageTemplateList []CognitoUserPoolInviteMessageTemplate

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CognitoUserPoolInviteMessageTemplateList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CognitoUserPoolInviteMessageTemplate{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CognitoUserPoolInviteMessageTemplateList{item}
		return nil
	}
	list := []CognitoUserPoolInviteMessageTemplate{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CognitoUserPoolInviteMessageTemplateList(list)
		return nil
	}
	return err
}

// CognitoUserPoolLambdaConfig represents the AWS::Cognito::UserPool.LambdaConfig CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html
type CognitoUserPoolLambdaConfig struct {
	// CreateAuthChallenge docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-createauthchallenge
	CreateAuthChallenge *StringExpr `json:"CreateAuthChallenge,omitempty"`
	// CustomMessage docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-custommessage
	CustomMessage *StringExpr `json:"CustomMessage,omitempty"`
	// DefineAuthChallenge docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-defineauthchallenge
	DefineAuthChallenge *StringExpr `json:"DefineAuthChallenge,omitempty"`
	// PostAuthentication docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-postauthentication
	PostAuthentication *StringExpr `json:"PostAuthentication,omitempty"`
	// PostConfirmation docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-postconfirmation
	PostConfirmation *StringExpr `json:"PostConfirmation,omitempty"`
	// PreAuthentication docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-preauthentication
	PreAuthentication *StringExpr `json:"PreAuthentication,omitempty"`
	// PreSignUp docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-presignup
	PreSignUp *StringExpr `json:"PreSignUp,omitempty"`
	// VerifyAuthChallengeResponse docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-verifyauthchallengeresponse
	VerifyAuthChallengeResponse *StringExpr `json:"VerifyAuthChallengeResponse,omitempty"`
}

// CognitoUserPoolLambdaConfigList represents a list of CognitoUserPoolLambdaConfig
type CognitoUserPoolLambdaConfigList []CognitoUserPoolLambdaConfig

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CognitoUserPoolLambdaConfigList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CognitoUserPoolLambdaConfig{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CognitoUserPoolLambdaConfigList{item}
		return nil
	}
	list := []CognitoUserPoolLambdaConfig{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CognitoUserPoolLambdaConfigList(list)
		return nil
	}
	return err
}

// CognitoUserPoolNumberAttributeConstraints represents the AWS::Cognito::UserPool.NumberAttributeConstraints CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-numberattributeconstraints.html
type CognitoUserPoolNumberAttributeConstraints struct {
	// MaxValue docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-numberattributeconstraints.html#cfn-cognito-userpool-numberattributeconstraints-maxvalue
	MaxValue *StringExpr `json:"MaxValue,omitempty"`
	// MinValue docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-numberattributeconstraints.html#cfn-cognito-userpool-numberattributeconstraints-minvalue
	MinValue *StringExpr `json:"MinValue,omitempty"`
}

// CognitoUserPoolNumberAttributeConstraintsList represents a list of CognitoUserPoolNumberAttributeConstraints
type CognitoUserPoolNumberAttributeConstraintsList []CognitoUserPoolNumberAttributeConstraints

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CognitoUserPoolNumberAttributeConstraintsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CognitoUserPoolNumberAttributeConstraints{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CognitoUserPoolNumberAttributeConstraintsList{item}
		return nil
	}
	list := []CognitoUserPoolNumberAttributeConstraints{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CognitoUserPoolNumberAttributeConstraintsList(list)
		return nil
	}
	return err
}

// CognitoUserPoolPasswordPolicy represents the AWS::Cognito::UserPool.PasswordPolicy CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-passwordpolicy.html
type CognitoUserPoolPasswordPolicy struct {
	// MinimumLength docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-passwordpolicy.html#cfn-cognito-userpool-passwordpolicy-minimumlength
	MinimumLength *IntegerExpr `json:"MinimumLength,omitempty"`
	// RequireLowercase docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-passwordpolicy.html#cfn-cognito-userpool-passwordpolicy-requirelowercase
	RequireLowercase *BoolExpr `json:"RequireLowercase,omitempty"`
	// RequireNumbers docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-passwordpolicy.html#cfn-cognito-userpool-passwordpolicy-requirenumbers
	RequireNumbers *BoolExpr `json:"RequireNumbers,omitempty"`
	// RequireSymbols docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-passwordpolicy.html#cfn-cognito-userpool-passwordpolicy-requiresymbols
	RequireSymbols *BoolExpr `json:"RequireSymbols,omitempty"`
	// RequireUppercase docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-passwordpolicy.html#cfn-cognito-userpool-passwordpolicy-requireuppercase
	RequireUppercase *BoolExpr `json:"RequireUppercase,omitempty"`
}

// CognitoUserPoolPasswordPolicyList represents a list of CognitoUserPoolPasswordPolicy
type CognitoUserPoolPasswordPolicyList []CognitoUserPoolPasswordPolicy

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CognitoUserPoolPasswordPolicyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CognitoUserPoolPasswordPolicy{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CognitoUserPoolPasswordPolicyList{item}
		return nil
	}
	list := []CognitoUserPoolPasswordPolicy{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CognitoUserPoolPasswordPolicyList(list)
		return nil
	}
	return err
}

// CognitoUserPoolPolicies represents the AWS::Cognito::UserPool.Policies CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-policies.html
type CognitoUserPoolPolicies struct {
	// PasswordPolicy docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-policies.html#cfn-cognito-userpool-policies-passwordpolicy
	PasswordPolicy *CognitoUserPoolPasswordPolicy `json:"PasswordPolicy,omitempty"`
}

// CognitoUserPoolPoliciesList represents a list of CognitoUserPoolPolicies
type CognitoUserPoolPoliciesList []CognitoUserPoolPolicies

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CognitoUserPoolPoliciesList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CognitoUserPoolPolicies{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CognitoUserPoolPoliciesList{item}
		return nil
	}
	list := []CognitoUserPoolPolicies{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CognitoUserPoolPoliciesList(list)
		return nil
	}
	return err
}

// CognitoUserPoolSchemaAttribute represents the AWS::Cognito::UserPool.SchemaAttribute CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-schemaattribute.html
type CognitoUserPoolSchemaAttribute struct {
	// AttributeDataType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-schemaattribute.html#cfn-cognito-userpool-schemaattribute-attributedatatype
	AttributeDataType *StringExpr `json:"AttributeDataType,omitempty"`
	// DeveloperOnlyAttribute docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-schemaattribute.html#cfn-cognito-userpool-schemaattribute-developeronlyattribute
	DeveloperOnlyAttribute *BoolExpr `json:"DeveloperOnlyAttribute,omitempty"`
	// Mutable docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-schemaattribute.html#cfn-cognito-userpool-schemaattribute-mutable
	Mutable *BoolExpr `json:"Mutable,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-schemaattribute.html#cfn-cognito-userpool-schemaattribute-name
	Name *StringExpr `json:"Name,omitempty"`
	// NumberAttributeConstraints docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-schemaattribute.html#cfn-cognito-userpool-schemaattribute-numberattributeconstraints
	NumberAttributeConstraints *CognitoUserPoolNumberAttributeConstraints `json:"NumberAttributeConstraints,omitempty"`
	// Required docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-schemaattribute.html#cfn-cognito-userpool-schemaattribute-required
	Required *BoolExpr `json:"Required,omitempty"`
	// StringAttributeConstraints docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-schemaattribute.html#cfn-cognito-userpool-schemaattribute-stringattributeconstraints
	StringAttributeConstraints *CognitoUserPoolStringAttributeConstraints `json:"StringAttributeConstraints,omitempty"`
}

// CognitoUserPoolSchemaAttributeList represents a list of CognitoUserPoolSchemaAttribute
type CognitoUserPoolSchemaAttributeList []CognitoUserPoolSchemaAttribute

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CognitoUserPoolSchemaAttributeList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CognitoUserPoolSchemaAttribute{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CognitoUserPoolSchemaAttributeList{item}
		return nil
	}
	list := []CognitoUserPoolSchemaAttribute{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CognitoUserPoolSchemaAttributeList(list)
		return nil
	}
	return err
}

// CognitoUserPoolSmsConfiguration represents the AWS::Cognito::UserPool.SmsConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-smsconfiguration.html
type CognitoUserPoolSmsConfiguration struct {
	// ExternalID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-smsconfiguration.html#cfn-cognito-userpool-smsconfiguration-externalid
	ExternalID *StringExpr `json:"ExternalId,omitempty"`
	// SnsCallerArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-smsconfiguration.html#cfn-cognito-userpool-smsconfiguration-snscallerarn
	SnsCallerArn *StringExpr `json:"SnsCallerArn,omitempty"`
}

// CognitoUserPoolSmsConfigurationList represents a list of CognitoUserPoolSmsConfiguration
type CognitoUserPoolSmsConfigurationList []CognitoUserPoolSmsConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CognitoUserPoolSmsConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CognitoUserPoolSmsConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CognitoUserPoolSmsConfigurationList{item}
		return nil
	}
	list := []CognitoUserPoolSmsConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CognitoUserPoolSmsConfigurationList(list)
		return nil
	}
	return err
}

// CognitoUserPoolStringAttributeConstraints represents the AWS::Cognito::UserPool.StringAttributeConstraints CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-stringattributeconstraints.html
type CognitoUserPoolStringAttributeConstraints struct {
	// MaxLength docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-stringattributeconstraints.html#cfn-cognito-userpool-stringattributeconstraints-maxlength
	MaxLength *StringExpr `json:"MaxLength,omitempty"`
	// MinLength docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-stringattributeconstraints.html#cfn-cognito-userpool-stringattributeconstraints-minlength
	MinLength *StringExpr `json:"MinLength,omitempty"`
}

// CognitoUserPoolStringAttributeConstraintsList represents a list of CognitoUserPoolStringAttributeConstraints
type CognitoUserPoolStringAttributeConstraintsList []CognitoUserPoolStringAttributeConstraints

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CognitoUserPoolStringAttributeConstraintsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CognitoUserPoolStringAttributeConstraints{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CognitoUserPoolStringAttributeConstraintsList{item}
		return nil
	}
	list := []CognitoUserPoolStringAttributeConstraints{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CognitoUserPoolStringAttributeConstraintsList(list)
		return nil
	}
	return err
}

// CognitoUserPoolUserAttributeType represents the AWS::Cognito::UserPoolUser.AttributeType CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpooluser-attributetype.html
type CognitoUserPoolUserAttributeType struct {
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpooluser-attributetype.html#cfn-cognito-userpooluser-attributetype-name
	Name *StringExpr `json:"Name,omitempty"`
	// Value docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpooluser-attributetype.html#cfn-cognito-userpooluser-attributetype-value
	Value *StringExpr `json:"Value,omitempty"`
}

// CognitoUserPoolUserAttributeTypeList represents a list of CognitoUserPoolUserAttributeType
type CognitoUserPoolUserAttributeTypeList []CognitoUserPoolUserAttributeType

// UnmarshalJSON sets the object from the provided JSON representation
func (l *CognitoUserPoolUserAttributeTypeList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := CognitoUserPoolUserAttributeType{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = CognitoUserPoolUserAttributeTypeList{item}
		return nil
	}
	list := []CognitoUserPoolUserAttributeType{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = CognitoUserPoolUserAttributeTypeList(list)
		return nil
	}
	return err
}

// ConfigConfigRuleScope represents the AWS::Config::ConfigRule.Scope CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-scope.html
type ConfigConfigRuleScope struct {
	// ComplianceResourceID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-scope.html#cfn-config-configrule-scope-complianceresourceid
	ComplianceResourceID *StringExpr `json:"ComplianceResourceId,omitempty"`
	// ComplianceResourceTypes docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-scope.html#cfn-config-configrule-scope-complianceresourcetypes
	ComplianceResourceTypes *StringListExpr `json:"ComplianceResourceTypes,omitempty"`
	// TagKey docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-scope.html#cfn-config-configrule-scope-tagkey
	TagKey *StringExpr `json:"TagKey,omitempty"`
	// TagValue docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-scope.html#cfn-config-configrule-scope-tagvalue
	TagValue *StringExpr `json:"TagValue,omitempty"`
}

// ConfigConfigRuleScopeList represents a list of ConfigConfigRuleScope
type ConfigConfigRuleScopeList []ConfigConfigRuleScope

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ConfigConfigRuleScopeList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ConfigConfigRuleScope{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ConfigConfigRuleScopeList{item}
		return nil
	}
	list := []ConfigConfigRuleScope{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ConfigConfigRuleScopeList(list)
		return nil
	}
	return err
}

// ConfigConfigRuleSource represents the AWS::Config::ConfigRule.Source CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source.html
type ConfigConfigRuleSource struct {
	// Owner docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source.html#cfn-config-configrule-source-owner
	Owner *StringExpr `json:"Owner,omitempty" validate:"dive,required"`
	// SourceDetails docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source.html#cfn-config-configrule-source-sourcedetails
	SourceDetails *ConfigConfigRuleSourceDetailList `json:"SourceDetails,omitempty"`
	// SourceIDentifier docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source.html#cfn-config-configrule-source-sourceidentifier
	SourceIDentifier *StringExpr `json:"SourceIdentifier,omitempty" validate:"dive,required"`
}

// ConfigConfigRuleSourceList represents a list of ConfigConfigRuleSource
type ConfigConfigRuleSourceList []ConfigConfigRuleSource

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ConfigConfigRuleSourceList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ConfigConfigRuleSource{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ConfigConfigRuleSourceList{item}
		return nil
	}
	list := []ConfigConfigRuleSource{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ConfigConfigRuleSourceList(list)
		return nil
	}
	return err
}

// ConfigConfigRuleSourceDetail represents the AWS::Config::ConfigRule.SourceDetail CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source-sourcedetails.html
type ConfigConfigRuleSourceDetail struct {
	// EventSource docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source-sourcedetails.html#cfn-config-configrule-source-sourcedetail-eventsource
	EventSource *StringExpr `json:"EventSource,omitempty" validate:"dive,required"`
	// MaximumExecutionFrequency docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source-sourcedetails.html#cfn-config-configrule-sourcedetail-maximumexecutionfrequency
	MaximumExecutionFrequency *StringExpr `json:"MaximumExecutionFrequency,omitempty"`
	// MessageType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source-sourcedetails.html#cfn-config-configrule-source-sourcedetail-messagetype
	MessageType *StringExpr `json:"MessageType,omitempty" validate:"dive,required"`
}

// ConfigConfigRuleSourceDetailList represents a list of ConfigConfigRuleSourceDetail
type ConfigConfigRuleSourceDetailList []ConfigConfigRuleSourceDetail

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ConfigConfigRuleSourceDetailList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ConfigConfigRuleSourceDetail{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ConfigConfigRuleSourceDetailList{item}
		return nil
	}
	list := []ConfigConfigRuleSourceDetail{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ConfigConfigRuleSourceDetailList(list)
		return nil
	}
	return err
}

// ConfigConfigurationRecorderRecordingGroup represents the AWS::Config::ConfigurationRecorder.RecordingGroup CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordinggroup.html
type ConfigConfigurationRecorderRecordingGroup struct {
	// AllSupported docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordinggroup.html#cfn-config-configurationrecorder-recordinggroup-allsupported
	AllSupported *BoolExpr `json:"AllSupported,omitempty"`
	// IncludeGlobalResourceTypes docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordinggroup.html#cfn-config-configurationrecorder-recordinggroup-includeglobalresourcetypes
	IncludeGlobalResourceTypes *BoolExpr `json:"IncludeGlobalResourceTypes,omitempty"`
	// ResourceTypes docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordinggroup.html#cfn-config-configurationrecorder-recordinggroup-resourcetypes
	ResourceTypes *StringListExpr `json:"ResourceTypes,omitempty"`
}

// ConfigConfigurationRecorderRecordingGroupList represents a list of ConfigConfigurationRecorderRecordingGroup
type ConfigConfigurationRecorderRecordingGroupList []ConfigConfigurationRecorderRecordingGroup

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ConfigConfigurationRecorderRecordingGroupList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ConfigConfigurationRecorderRecordingGroup{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ConfigConfigurationRecorderRecordingGroupList{item}
		return nil
	}
	list := []ConfigConfigurationRecorderRecordingGroup{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ConfigConfigurationRecorderRecordingGroupList(list)
		return nil
	}
	return err
}

// ConfigDeliveryChannelConfigSnapshotDeliveryProperties represents the AWS::Config::DeliveryChannel.ConfigSnapshotDeliveryProperties CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-deliverychannel-configsnapshotdeliveryproperties.html
type ConfigDeliveryChannelConfigSnapshotDeliveryProperties struct {
	// DeliveryFrequency docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-deliverychannel-configsnapshotdeliveryproperties.html#cfn-config-deliverychannel-configsnapshotdeliveryproperties-deliveryfrequency
	DeliveryFrequency *StringExpr `json:"DeliveryFrequency,omitempty"`
}

// ConfigDeliveryChannelConfigSnapshotDeliveryPropertiesList represents a list of ConfigDeliveryChannelConfigSnapshotDeliveryProperties
type ConfigDeliveryChannelConfigSnapshotDeliveryPropertiesList []ConfigDeliveryChannelConfigSnapshotDeliveryProperties

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ConfigDeliveryChannelConfigSnapshotDeliveryPropertiesList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ConfigDeliveryChannelConfigSnapshotDeliveryProperties{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ConfigDeliveryChannelConfigSnapshotDeliveryPropertiesList{item}
		return nil
	}
	list := []ConfigDeliveryChannelConfigSnapshotDeliveryProperties{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ConfigDeliveryChannelConfigSnapshotDeliveryPropertiesList(list)
		return nil
	}
	return err
}

// DataPipelinePipelineField represents the AWS::DataPipeline::Pipeline.Field CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobjects-fields.html
type DataPipelinePipelineField struct {
	// Key docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobjects-fields.html#cfn-datapipeline-pipeline-pipelineobjects-fields-key
	Key *StringExpr `json:"Key,omitempty" validate:"dive,required"`
	// RefValue docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobjects-fields.html#cfn-datapipeline-pipeline-pipelineobjects-fields-refvalue
	RefValue *StringExpr `json:"RefValue,omitempty"`
	// StringValue docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobjects-fields.html#cfn-datapipeline-pipeline-pipelineobjects-fields-stringvalue
	StringValue *StringExpr `json:"StringValue,omitempty"`
}

// DataPipelinePipelineFieldList represents a list of DataPipelinePipelineField
type DataPipelinePipelineFieldList []DataPipelinePipelineField

// UnmarshalJSON sets the object from the provided JSON representation
func (l *DataPipelinePipelineFieldList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := DataPipelinePipelineField{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = DataPipelinePipelineFieldList{item}
		return nil
	}
	list := []DataPipelinePipelineField{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = DataPipelinePipelineFieldList(list)
		return nil
	}
	return err
}

// DataPipelinePipelineParameterAttribute represents the AWS::DataPipeline::Pipeline.ParameterAttribute CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parameterobjects-attributes.html
type DataPipelinePipelineParameterAttribute struct {
	// Key docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parameterobjects-attributes.html#cfn-datapipeline-pipeline-parameterobjects-attribtues-key
	Key *StringExpr `json:"Key,omitempty" validate:"dive,required"`
	// StringValue docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parameterobjects-attributes.html#cfn-datapipeline-pipeline-parameterobjects-attribtues-stringvalue
	StringValue *StringExpr `json:"StringValue,omitempty" validate:"dive,required"`
}

// DataPipelinePipelineParameterAttributeList represents a list of DataPipelinePipelineParameterAttribute
type DataPipelinePipelineParameterAttributeList []DataPipelinePipelineParameterAttribute

// UnmarshalJSON sets the object from the provided JSON representation
func (l *DataPipelinePipelineParameterAttributeList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := DataPipelinePipelineParameterAttribute{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = DataPipelinePipelineParameterAttributeList{item}
		return nil
	}
	list := []DataPipelinePipelineParameterAttribute{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = DataPipelinePipelineParameterAttributeList(list)
		return nil
	}
	return err
}

// DataPipelinePipelineParameterObject represents the AWS::DataPipeline::Pipeline.ParameterObject CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parameterobjects.html
type DataPipelinePipelineParameterObject struct {
	// Attributes docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parameterobjects.html#cfn-datapipeline-pipeline-parameterobjects-attributes
	Attributes *DataPipelinePipelineParameterAttributeList `json:"Attributes,omitempty" validate:"dive,required"`
	// ID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parameterobjects.html#cfn-datapipeline-pipeline-parameterobject-id
	ID *StringExpr `json:"Id,omitempty" validate:"dive,required"`
}

// DataPipelinePipelineParameterObjectList represents a list of DataPipelinePipelineParameterObject
type DataPipelinePipelineParameterObjectList []DataPipelinePipelineParameterObject

// UnmarshalJSON sets the object from the provided JSON representation
func (l *DataPipelinePipelineParameterObjectList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := DataPipelinePipelineParameterObject{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = DataPipelinePipelineParameterObjectList{item}
		return nil
	}
	list := []DataPipelinePipelineParameterObject{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = DataPipelinePipelineParameterObjectList(list)
		return nil
	}
	return err
}

// DataPipelinePipelineParameterValue represents the AWS::DataPipeline::Pipeline.ParameterValue CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parametervalues.html
type DataPipelinePipelineParameterValue struct {
	// ID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parametervalues.html#cfn-datapipeline-pipeline-parametervalues-id
	ID *StringExpr `json:"Id,omitempty" validate:"dive,required"`
	// StringValue docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parametervalues.html#cfn-datapipeline-pipeline-parametervalues-stringvalue
	StringValue *StringExpr `json:"StringValue,omitempty" validate:"dive,required"`
}

// DataPipelinePipelineParameterValueList represents a list of DataPipelinePipelineParameterValue
type DataPipelinePipelineParameterValueList []DataPipelinePipelineParameterValue

// UnmarshalJSON sets the object from the provided JSON representation
func (l *DataPipelinePipelineParameterValueList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := DataPipelinePipelineParameterValue{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = DataPipelinePipelineParameterValueList{item}
		return nil
	}
	list := []DataPipelinePipelineParameterValue{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = DataPipelinePipelineParameterValueList(list)
		return nil
	}
	return err
}

// DataPipelinePipelinePipelineObject represents the AWS::DataPipeline::Pipeline.PipelineObject CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobjects.html
type DataPipelinePipelinePipelineObject struct {
	// Fields docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobjects.html#cfn-datapipeline-pipeline-pipelineobjects-fields
	Fields *DataPipelinePipelineFieldList `json:"Fields,omitempty" validate:"dive,required"`
	// ID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobjects.html#cfn-datapipeline-pipeline-pipelineobjects-id
	ID *StringExpr `json:"Id,omitempty" validate:"dive,required"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobjects.html#cfn-datapipeline-pipeline-pipelineobjects-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
}

// DataPipelinePipelinePipelineObjectList represents a list of DataPipelinePipelinePipelineObject
type DataPipelinePipelinePipelineObjectList []DataPipelinePipelinePipelineObject

// UnmarshalJSON sets the object from the provided JSON representation
func (l *DataPipelinePipelinePipelineObjectList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := DataPipelinePipelinePipelineObject{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = DataPipelinePipelinePipelineObjectList{item}
		return nil
	}
	list := []DataPipelinePipelinePipelineObject{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = DataPipelinePipelinePipelineObjectList(list)
		return nil
	}
	return err
}

// DataPipelinePipelinePipelineTag represents the AWS::DataPipeline::Pipeline.PipelineTag CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelinetags.html
type DataPipelinePipelinePipelineTag struct {
	// Key docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelinetags.html#cfn-datapipeline-pipeline-pipelinetags-key
	Key *StringExpr `json:"Key,omitempty" validate:"dive,required"`
	// Value docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelinetags.html#cfn-datapipeline-pipeline-pipelinetags-value
	Value *StringExpr `json:"Value,omitempty" validate:"dive,required"`
}

// DataPipelinePipelinePipelineTagList represents a list of DataPipelinePipelinePipelineTag
type DataPipelinePipelinePipelineTagList []DataPipelinePipelinePipelineTag

// UnmarshalJSON sets the object from the provided JSON representation
func (l *DataPipelinePipelinePipelineTagList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := DataPipelinePipelinePipelineTag{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = DataPipelinePipelinePipelineTagList{item}
		return nil
	}
	list := []DataPipelinePipelinePipelineTag{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = DataPipelinePipelinePipelineTagList(list)
		return nil
	}
	return err
}

// DirectoryServiceMicrosoftADVPCSettings represents the AWS::DirectoryService::MicrosoftAD.VpcSettings CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-microsoftad-vpcsettings.html
type DirectoryServiceMicrosoftADVPCSettings struct {
	// SubnetIDs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-microsoftad-vpcsettings.html#cfn-directoryservice-microsoftad-vpcsettings-subnetids
	SubnetIDs *StringListExpr `json:"SubnetIds,omitempty" validate:"dive,required"`
	// VPCID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-microsoftad-vpcsettings.html#cfn-directoryservice-microsoftad-vpcsettings-vpcid
	VPCID *StringExpr `json:"VpcId,omitempty" validate:"dive,required"`
}

// DirectoryServiceMicrosoftADVPCSettingsList represents a list of DirectoryServiceMicrosoftADVPCSettings
type DirectoryServiceMicrosoftADVPCSettingsList []DirectoryServiceMicrosoftADVPCSettings

// UnmarshalJSON sets the object from the provided JSON representation
func (l *DirectoryServiceMicrosoftADVPCSettingsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := DirectoryServiceMicrosoftADVPCSettings{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = DirectoryServiceMicrosoftADVPCSettingsList{item}
		return nil
	}
	list := []DirectoryServiceMicrosoftADVPCSettings{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = DirectoryServiceMicrosoftADVPCSettingsList(list)
		return nil
	}
	return err
}

// DirectoryServiceSimpleADVPCSettings represents the AWS::DirectoryService::SimpleAD.VpcSettings CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-simplead-vpcsettings.html
type DirectoryServiceSimpleADVPCSettings struct {
	// SubnetIDs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-simplead-vpcsettings.html#cfn-directoryservice-simplead-vpcsettings-subnetids
	SubnetIDs *StringListExpr `json:"SubnetIds,omitempty" validate:"dive,required"`
	// VPCID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-simplead-vpcsettings.html#cfn-directoryservice-simplead-vpcsettings-vpcid
	VPCID *StringExpr `json:"VpcId,omitempty" validate:"dive,required"`
}

// DirectoryServiceSimpleADVPCSettingsList represents a list of DirectoryServiceSimpleADVPCSettings
type DirectoryServiceSimpleADVPCSettingsList []DirectoryServiceSimpleADVPCSettings

// UnmarshalJSON sets the object from the provided JSON representation
func (l *DirectoryServiceSimpleADVPCSettingsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := DirectoryServiceSimpleADVPCSettings{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = DirectoryServiceSimpleADVPCSettingsList{item}
		return nil
	}
	list := []DirectoryServiceSimpleADVPCSettings{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = DirectoryServiceSimpleADVPCSettingsList(list)
		return nil
	}
	return err
}

// DynamoDBTableAttributeDefinition represents the AWS::DynamoDB::Table.AttributeDefinition CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-attributedef.html
type DynamoDBTableAttributeDefinition struct {
	// AttributeName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-attributedef.html#cfn-dynamodb-attributedef-attributename
	AttributeName *StringExpr `json:"AttributeName,omitempty" validate:"dive,required"`
	// AttributeType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-attributedef.html#cfn-dynamodb-attributedef-attributename-attributetype
	AttributeType *StringExpr `json:"AttributeType,omitempty" validate:"dive,required"`
}

// DynamoDBTableAttributeDefinitionList represents a list of DynamoDBTableAttributeDefinition
type DynamoDBTableAttributeDefinitionList []DynamoDBTableAttributeDefinition

// UnmarshalJSON sets the object from the provided JSON representation
func (l *DynamoDBTableAttributeDefinitionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := DynamoDBTableAttributeDefinition{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = DynamoDBTableAttributeDefinitionList{item}
		return nil
	}
	list := []DynamoDBTableAttributeDefinition{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = DynamoDBTableAttributeDefinitionList(list)
		return nil
	}
	return err
}

// DynamoDBTableGlobalSecondaryIndex represents the AWS::DynamoDB::Table.GlobalSecondaryIndex CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html
type DynamoDBTableGlobalSecondaryIndex struct {
	// IndexName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html#cfn-dynamodb-gsi-indexname
	IndexName *StringExpr `json:"IndexName,omitempty" validate:"dive,required"`
	// KeySchema docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html#cfn-dynamodb-gsi-keyschema
	KeySchema *DynamoDBTableKeySchemaList `json:"KeySchema,omitempty" validate:"dive,required"`
	// Projection docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html#cfn-dynamodb-gsi-projection
	Projection *DynamoDBTableProjection `json:"Projection,omitempty" validate:"dive,required"`
	// ProvisionedThroughput docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html#cfn-dynamodb-gsi-provisionedthroughput
	ProvisionedThroughput *DynamoDBTableProvisionedThroughput `json:"ProvisionedThroughput,omitempty" validate:"dive,required"`
}

// DynamoDBTableGlobalSecondaryIndexList represents a list of DynamoDBTableGlobalSecondaryIndex
type DynamoDBTableGlobalSecondaryIndexList []DynamoDBTableGlobalSecondaryIndex

// UnmarshalJSON sets the object from the provided JSON representation
func (l *DynamoDBTableGlobalSecondaryIndexList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := DynamoDBTableGlobalSecondaryIndex{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = DynamoDBTableGlobalSecondaryIndexList{item}
		return nil
	}
	list := []DynamoDBTableGlobalSecondaryIndex{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = DynamoDBTableGlobalSecondaryIndexList(list)
		return nil
	}
	return err
}

// DynamoDBTableKeySchema represents the AWS::DynamoDB::Table.KeySchema CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-keyschema.html
type DynamoDBTableKeySchema struct {
	// AttributeName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-keyschema.html#aws-properties-dynamodb-keyschema-attributename
	AttributeName *StringExpr `json:"AttributeName,omitempty" validate:"dive,required"`
	// KeyType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-keyschema.html#aws-properties-dynamodb-keyschema-keytype
	KeyType *StringExpr `json:"KeyType,omitempty" validate:"dive,required"`
}

// DynamoDBTableKeySchemaList represents a list of DynamoDBTableKeySchema
type DynamoDBTableKeySchemaList []DynamoDBTableKeySchema

// UnmarshalJSON sets the object from the provided JSON representation
func (l *DynamoDBTableKeySchemaList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := DynamoDBTableKeySchema{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = DynamoDBTableKeySchemaList{item}
		return nil
	}
	list := []DynamoDBTableKeySchema{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = DynamoDBTableKeySchemaList(list)
		return nil
	}
	return err
}

// DynamoDBTableLocalSecondaryIndex represents the AWS::DynamoDB::Table.LocalSecondaryIndex CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-lsi.html
type DynamoDBTableLocalSecondaryIndex struct {
	// IndexName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-lsi.html#cfn-dynamodb-lsi-indexname
	IndexName *StringExpr `json:"IndexName,omitempty" validate:"dive,required"`
	// KeySchema docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-lsi.html#cfn-dynamodb-lsi-keyschema
	KeySchema *DynamoDBTableKeySchemaList `json:"KeySchema,omitempty" validate:"dive,required"`
	// Projection docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-lsi.html#cfn-dynamodb-lsi-projection
	Projection *DynamoDBTableProjection `json:"Projection,omitempty" validate:"dive,required"`
}

// DynamoDBTableLocalSecondaryIndexList represents a list of DynamoDBTableLocalSecondaryIndex
type DynamoDBTableLocalSecondaryIndexList []DynamoDBTableLocalSecondaryIndex

// UnmarshalJSON sets the object from the provided JSON representation
func (l *DynamoDBTableLocalSecondaryIndexList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := DynamoDBTableLocalSecondaryIndex{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = DynamoDBTableLocalSecondaryIndexList{item}
		return nil
	}
	list := []DynamoDBTableLocalSecondaryIndex{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = DynamoDBTableLocalSecondaryIndexList(list)
		return nil
	}
	return err
}

// DynamoDBTableProjection represents the AWS::DynamoDB::Table.Projection CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-projectionobject.html
type DynamoDBTableProjection struct {
	// NonKeyAttributes docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-projectionobject.html#cfn-dynamodb-projectionobj-nonkeyatt
	NonKeyAttributes *StringListExpr `json:"NonKeyAttributes,omitempty"`
	// ProjectionType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-projectionobject.html#cfn-dynamodb-projectionobj-projtype
	ProjectionType *StringExpr `json:"ProjectionType,omitempty"`
}

// DynamoDBTableProjectionList represents a list of DynamoDBTableProjection
type DynamoDBTableProjectionList []DynamoDBTableProjection

// UnmarshalJSON sets the object from the provided JSON representation
func (l *DynamoDBTableProjectionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := DynamoDBTableProjection{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = DynamoDBTableProjectionList{item}
		return nil
	}
	list := []DynamoDBTableProjection{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = DynamoDBTableProjectionList(list)
		return nil
	}
	return err
}

// DynamoDBTableProvisionedThroughput represents the AWS::DynamoDB::Table.ProvisionedThroughput CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-provisionedthroughput.html
type DynamoDBTableProvisionedThroughput struct {
	// ReadCapacityUnits docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-provisionedthroughput.html#cfn-dynamodb-provisionedthroughput-readcapacityunits
	ReadCapacityUnits *IntegerExpr `json:"ReadCapacityUnits,omitempty" validate:"dive,required"`
	// WriteCapacityUnits docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-provisionedthroughput.html#cfn-dynamodb-provisionedthroughput-writecapacityunits
	WriteCapacityUnits *IntegerExpr `json:"WriteCapacityUnits,omitempty" validate:"dive,required"`
}

// DynamoDBTableProvisionedThroughputList represents a list of DynamoDBTableProvisionedThroughput
type DynamoDBTableProvisionedThroughputList []DynamoDBTableProvisionedThroughput

// UnmarshalJSON sets the object from the provided JSON representation
func (l *DynamoDBTableProvisionedThroughputList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := DynamoDBTableProvisionedThroughput{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = DynamoDBTableProvisionedThroughputList{item}
		return nil
	}
	list := []DynamoDBTableProvisionedThroughput{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = DynamoDBTableProvisionedThroughputList(list)
		return nil
	}
	return err
}

// DynamoDBTableStreamSpecification represents the AWS::DynamoDB::Table.StreamSpecification CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-streamspecification.html
type DynamoDBTableStreamSpecification struct {
	// StreamViewType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-streamspecification.html#cfn-dynamodb-streamspecification-streamviewtype
	StreamViewType *StringExpr `json:"StreamViewType,omitempty" validate:"dive,required"`
}

// DynamoDBTableStreamSpecificationList represents a list of DynamoDBTableStreamSpecification
type DynamoDBTableStreamSpecificationList []DynamoDBTableStreamSpecification

// UnmarshalJSON sets the object from the provided JSON representation
func (l *DynamoDBTableStreamSpecificationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := DynamoDBTableStreamSpecification{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = DynamoDBTableStreamSpecificationList{item}
		return nil
	}
	list := []DynamoDBTableStreamSpecification{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = DynamoDBTableStreamSpecificationList(list)
		return nil
	}
	return err
}

// EC2InstanceAssociationParameter represents the AWS::EC2::Instance.AssociationParameter CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations-associationparameters.html
type EC2InstanceAssociationParameter struct {
	// Key docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations-associationparameters.html#cfn-ec2-instance-ssmassociations-associationparameters-key
	Key *StringExpr `json:"Key,omitempty" validate:"dive,required"`
	// Value docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations-associationparameters.html#cfn-ec2-instance-ssmassociations-associationparameters-value
	Value *StringListExpr `json:"Value,omitempty" validate:"dive,required"`
}

// EC2InstanceAssociationParameterList represents a list of EC2InstanceAssociationParameter
type EC2InstanceAssociationParameterList []EC2InstanceAssociationParameter

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2InstanceAssociationParameterList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2InstanceAssociationParameter{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2InstanceAssociationParameterList{item}
		return nil
	}
	list := []EC2InstanceAssociationParameter{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2InstanceAssociationParameterList(list)
		return nil
	}
	return err
}

// EC2InstanceBlockDeviceMapping represents the AWS::EC2::Instance.BlockDeviceMapping CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-mapping.html
type EC2InstanceBlockDeviceMapping struct {
	// DeviceName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-mapping.html#cfn-ec2-blockdev-mapping-devicename
	DeviceName *StringExpr `json:"DeviceName,omitempty" validate:"dive,required"`
	// Ebs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-mapping.html#cfn-ec2-blockdev-mapping-ebs
	Ebs *EC2InstanceEbs `json:"Ebs,omitempty"`
	// NoDevice docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-mapping.html#cfn-ec2-blockdev-mapping-nodevice
	NoDevice *EC2InstanceNoDevice `json:"NoDevice,omitempty"`
	// VirtualName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-mapping.html#cfn-ec2-blockdev-mapping-virtualname
	VirtualName *StringExpr `json:"VirtualName,omitempty"`
}

// EC2InstanceBlockDeviceMappingList represents a list of EC2InstanceBlockDeviceMapping
type EC2InstanceBlockDeviceMappingList []EC2InstanceBlockDeviceMapping

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2InstanceBlockDeviceMappingList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2InstanceBlockDeviceMapping{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2InstanceBlockDeviceMappingList{item}
		return nil
	}
	list := []EC2InstanceBlockDeviceMapping{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2InstanceBlockDeviceMappingList(list)
		return nil
	}
	return err
}

// EC2InstanceEbs represents the AWS::EC2::Instance.Ebs CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-template.html
type EC2InstanceEbs struct {
	// DeleteOnTermination docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-template.html#cfn-ec2-blockdev-template-deleteontermination
	DeleteOnTermination *BoolExpr `json:"DeleteOnTermination,omitempty"`
	// Encrypted docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-template.html#cfn-ec2-blockdev-template-encrypted
	Encrypted *BoolExpr `json:"Encrypted,omitempty"`
	// Iops docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-template.html#cfn-ec2-blockdev-template-iops
	Iops *IntegerExpr `json:"Iops,omitempty"`
	// SnapshotID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-template.html#cfn-ec2-blockdev-template-snapshotid
	SnapshotID *StringExpr `json:"SnapshotId,omitempty"`
	// VolumeSize docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-template.html#cfn-ec2-blockdev-template-volumesize
	VolumeSize *IntegerExpr `json:"VolumeSize,omitempty"`
	// VolumeType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-template.html#cfn-ec2-blockdev-template-volumetype
	VolumeType *StringExpr `json:"VolumeType,omitempty"`
}

// EC2InstanceEbsList represents a list of EC2InstanceEbs
type EC2InstanceEbsList []EC2InstanceEbs

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2InstanceEbsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2InstanceEbs{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2InstanceEbsList{item}
		return nil
	}
	list := []EC2InstanceEbs{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2InstanceEbsList(list)
		return nil
	}
	return err
}

// EC2InstanceInstanceIPv6Address represents the AWS::EC2::Instance.InstanceIpv6Address CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-instanceipv6address.html
type EC2InstanceInstanceIPv6Address struct {
	// IPv6Address docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-instanceipv6address.html#cfn-ec2-instance-instanceipv6address-ipv6address
	IPv6Address *StringExpr `json:"Ipv6Address,omitempty" validate:"dive,required"`
}

// EC2InstanceInstanceIPv6AddressList represents a list of EC2InstanceInstanceIPv6Address
type EC2InstanceInstanceIPv6AddressList []EC2InstanceInstanceIPv6Address

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2InstanceInstanceIPv6AddressList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2InstanceInstanceIPv6Address{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2InstanceInstanceIPv6AddressList{item}
		return nil
	}
	list := []EC2InstanceInstanceIPv6Address{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2InstanceInstanceIPv6AddressList(list)
		return nil
	}
	return err
}

// EC2InstanceNetworkInterface represents the AWS::EC2::Instance.NetworkInterface CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html
type EC2InstanceNetworkInterface struct {
	// AssociatePublicIPAddress docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html#aws-properties-ec2-network-iface-embedded-associatepubip
	AssociatePublicIPAddress *BoolExpr `json:"AssociatePublicIpAddress,omitempty"`
	// DeleteOnTermination docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html#aws-properties-ec2-network-iface-embedded-delete
	DeleteOnTermination *BoolExpr `json:"DeleteOnTermination,omitempty"`
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html#aws-properties-ec2-network-iface-embedded-description
	Description *StringExpr `json:"Description,omitempty"`
	// DeviceIndex docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html#aws-properties-ec2-network-iface-embedded-deviceindex
	DeviceIndex *StringExpr `json:"DeviceIndex,omitempty" validate:"dive,required"`
	// GroupSet docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html#aws-properties-ec2-network-iface-embedded-groupset
	GroupSet *StringListExpr `json:"GroupSet,omitempty"`
	// IPv6AddressCount docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html#cfn-ec2-instance-networkinterface-ipv6addresscount
	IPv6AddressCount *IntegerExpr `json:"Ipv6AddressCount,omitempty"`
	// IPv6Addresses docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html#cfn-ec2-instance-networkinterface-ipv6addresses
	IPv6Addresses *EC2InstanceInstanceIPv6AddressList `json:"Ipv6Addresses,omitempty"`
	// NetworkInterfaceID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html#aws-properties-ec2-network-iface-embedded-network-iface
	NetworkInterfaceID *StringExpr `json:"NetworkInterfaceId,omitempty"`
	// PrivateIPAddress docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html#aws-properties-ec2-network-iface-embedded-privateipaddress
	PrivateIPAddress *StringExpr `json:"PrivateIpAddress,omitempty"`
	// PrivateIPAddresses docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html#aws-properties-ec2-network-iface-embedded-privateipaddresses
	PrivateIPAddresses *EC2InstancePrivateIPAddressSpecificationList `json:"PrivateIpAddresses,omitempty"`
	// SecondaryPrivateIPAddressCount docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html#aws-properties-ec2-network-iface-embedded-secondprivateip
	SecondaryPrivateIPAddressCount *IntegerExpr `json:"SecondaryPrivateIpAddressCount,omitempty"`
	// SubnetID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html#aws-properties-ec2-network-iface-embedded-subnetid
	SubnetID *StringExpr `json:"SubnetId,omitempty"`
}

// EC2InstanceNetworkInterfaceList represents a list of EC2InstanceNetworkInterface
type EC2InstanceNetworkInterfaceList []EC2InstanceNetworkInterface

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2InstanceNetworkInterfaceList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2InstanceNetworkInterface{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2InstanceNetworkInterfaceList{item}
		return nil
	}
	list := []EC2InstanceNetworkInterface{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2InstanceNetworkInterfaceList(list)
		return nil
	}
	return err
}

// EC2InstanceNoDevice represents the AWS::EC2::Instance.NoDevice CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-nodevice.html
type EC2InstanceNoDevice struct {
}

// EC2InstanceNoDeviceList represents a list of EC2InstanceNoDevice
type EC2InstanceNoDeviceList []EC2InstanceNoDevice

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2InstanceNoDeviceList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2InstanceNoDevice{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2InstanceNoDeviceList{item}
		return nil
	}
	list := []EC2InstanceNoDevice{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2InstanceNoDeviceList(list)
		return nil
	}
	return err
}

// EC2InstancePrivateIPAddressSpecification represents the AWS::EC2::Instance.PrivateIpAddressSpecification CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-interface-privateipspec.html
type EC2InstancePrivateIPAddressSpecification struct {
	// Primary docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-interface-privateipspec.html#cfn-ec2-networkinterface-privateipspecification-primary
	Primary *BoolExpr `json:"Primary,omitempty" validate:"dive,required"`
	// PrivateIPAddress docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-interface-privateipspec.html#cfn-ec2-networkinterface-privateipspecification-privateipaddress
	PrivateIPAddress *StringExpr `json:"PrivateIpAddress,omitempty" validate:"dive,required"`
}

// EC2InstancePrivateIPAddressSpecificationList represents a list of EC2InstancePrivateIPAddressSpecification
type EC2InstancePrivateIPAddressSpecificationList []EC2InstancePrivateIPAddressSpecification

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2InstancePrivateIPAddressSpecificationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2InstancePrivateIPAddressSpecification{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2InstancePrivateIPAddressSpecificationList{item}
		return nil
	}
	list := []EC2InstancePrivateIPAddressSpecification{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2InstancePrivateIPAddressSpecificationList(list)
		return nil
	}
	return err
}

// EC2InstanceSsmAssociation represents the AWS::EC2::Instance.SsmAssociation CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations.html
type EC2InstanceSsmAssociation struct {
	// AssociationParameters docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations.html#cfn-ec2-instance-ssmassociations-associationparameters
	AssociationParameters *EC2InstanceAssociationParameterList `json:"AssociationParameters,omitempty"`
	// DocumentName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations.html#cfn-ec2-instance-ssmassociations-documentname
	DocumentName *StringExpr `json:"DocumentName,omitempty" validate:"dive,required"`
}

// EC2InstanceSsmAssociationList represents a list of EC2InstanceSsmAssociation
type EC2InstanceSsmAssociationList []EC2InstanceSsmAssociation

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2InstanceSsmAssociationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2InstanceSsmAssociation{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2InstanceSsmAssociationList{item}
		return nil
	}
	list := []EC2InstanceSsmAssociation{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2InstanceSsmAssociationList(list)
		return nil
	}
	return err
}

// EC2InstanceVolume represents the AWS::EC2::Instance.Volume CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-mount-point.html
type EC2InstanceVolume struct {
	// Device docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-mount-point.html#cfn-ec2-mountpoint-device
	Device *StringExpr `json:"Device,omitempty" validate:"dive,required"`
	// VolumeID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-mount-point.html#cfn-ec2-mountpoint-volumeid
	VolumeID *StringExpr `json:"VolumeId,omitempty" validate:"dive,required"`
}

// EC2InstanceVolumeList represents a list of EC2InstanceVolume
type EC2InstanceVolumeList []EC2InstanceVolume

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2InstanceVolumeList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2InstanceVolume{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2InstanceVolumeList{item}
		return nil
	}
	list := []EC2InstanceVolume{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2InstanceVolumeList(list)
		return nil
	}
	return err
}

// EC2NetworkACLEntryIcmp represents the AWS::EC2::NetworkAclEntry.Icmp CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkaclentry-icmp.html
type EC2NetworkACLEntryIcmp struct {
	// Code docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkaclentry-icmp.html#cfn-ec2-networkaclentry-icmp-code
	Code *IntegerExpr `json:"Code,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkaclentry-icmp.html#cfn-ec2-networkaclentry-icmp-type
	Type *IntegerExpr `json:"Type,omitempty"`
}

// EC2NetworkACLEntryIcmpList represents a list of EC2NetworkACLEntryIcmp
type EC2NetworkACLEntryIcmpList []EC2NetworkACLEntryIcmp

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2NetworkACLEntryIcmpList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2NetworkACLEntryIcmp{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2NetworkACLEntryIcmpList{item}
		return nil
	}
	list := []EC2NetworkACLEntryIcmp{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2NetworkACLEntryIcmpList(list)
		return nil
	}
	return err
}

// EC2NetworkACLEntryPortRange represents the AWS::EC2::NetworkAclEntry.PortRange CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkaclentry-portrange.html
type EC2NetworkACLEntryPortRange struct {
	// From docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkaclentry-portrange.html#cfn-ec2-networkaclentry-portrange-from
	From *IntegerExpr `json:"From,omitempty"`
	// To docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkaclentry-portrange.html#cfn-ec2-networkaclentry-portrange-to
	To *IntegerExpr `json:"To,omitempty"`
}

// EC2NetworkACLEntryPortRangeList represents a list of EC2NetworkACLEntryPortRange
type EC2NetworkACLEntryPortRangeList []EC2NetworkACLEntryPortRange

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2NetworkACLEntryPortRangeList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2NetworkACLEntryPortRange{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2NetworkACLEntryPortRangeList{item}
		return nil
	}
	list := []EC2NetworkACLEntryPortRange{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2NetworkACLEntryPortRangeList(list)
		return nil
	}
	return err
}

// EC2NetworkInterfaceInstanceIPv6Address represents the AWS::EC2::NetworkInterface.InstanceIpv6Address CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterface-instanceipv6address.html
type EC2NetworkInterfaceInstanceIPv6Address struct {
	// IPv6Address docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterface-instanceipv6address.html#cfn-ec2-networkinterface-instanceipv6address-ipv6address
	IPv6Address *StringExpr `json:"Ipv6Address,omitempty" validate:"dive,required"`
}

// EC2NetworkInterfaceInstanceIPv6AddressList represents a list of EC2NetworkInterfaceInstanceIPv6Address
type EC2NetworkInterfaceInstanceIPv6AddressList []EC2NetworkInterfaceInstanceIPv6Address

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2NetworkInterfaceInstanceIPv6AddressList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2NetworkInterfaceInstanceIPv6Address{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2NetworkInterfaceInstanceIPv6AddressList{item}
		return nil
	}
	list := []EC2NetworkInterfaceInstanceIPv6Address{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2NetworkInterfaceInstanceIPv6AddressList(list)
		return nil
	}
	return err
}

// EC2NetworkInterfacePrivateIPAddressSpecification represents the AWS::EC2::NetworkInterface.PrivateIpAddressSpecification CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-interface-privateipspec.html
type EC2NetworkInterfacePrivateIPAddressSpecification struct {
	// Primary docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-interface-privateipspec.html#cfn-ec2-networkinterface-privateipspecification-primary
	Primary *BoolExpr `json:"Primary,omitempty" validate:"dive,required"`
	// PrivateIPAddress docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-interface-privateipspec.html#cfn-ec2-networkinterface-privateipspecification-privateipaddress
	PrivateIPAddress *StringExpr `json:"PrivateIpAddress,omitempty" validate:"dive,required"`
}

// EC2NetworkInterfacePrivateIPAddressSpecificationList represents a list of EC2NetworkInterfacePrivateIPAddressSpecification
type EC2NetworkInterfacePrivateIPAddressSpecificationList []EC2NetworkInterfacePrivateIPAddressSpecification

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2NetworkInterfacePrivateIPAddressSpecificationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2NetworkInterfacePrivateIPAddressSpecification{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2NetworkInterfacePrivateIPAddressSpecificationList{item}
		return nil
	}
	list := []EC2NetworkInterfacePrivateIPAddressSpecification{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2NetworkInterfacePrivateIPAddressSpecificationList(list)
		return nil
	}
	return err
}

// EC2SecurityGroupRule represents the AWS::EC2::SecurityGroup.Rule CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-rule.html
type EC2SecurityGroupRule struct {
	// CidrIP docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-rule.html#cfn-ec2-security-group-rule-cidrip
	CidrIP *StringExpr `json:"CidrIp,omitempty"`
	// FromPort docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-rule.html#cfn-ec2-security-group-rule-fromport
	FromPort *IntegerExpr `json:"FromPort,omitempty"`
	// IPProtocol docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-rule.html#cfn-ec2-security-group-rule-ipprotocol
	IPProtocol *StringExpr `json:"IpProtocol,omitempty" validate:"dive,required"`
	// SourceSecurityGroupID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-rule.html#cfn-ec2-security-group-rule-sourcesecuritygroupid
	SourceSecurityGroupID *StringExpr `json:"SourceSecurityGroupId,omitempty"`
	// SourceSecurityGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-rule.html#cfn-ec2-security-group-rule-sourcesecuritygroupname
	SourceSecurityGroupName *StringExpr `json:"SourceSecurityGroupName,omitempty"`
	// SourceSecurityGroupOwnerID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-rule.html#cfn-ec2-security-group-rule-sourcesecuritygroupownerid
	SourceSecurityGroupOwnerID *StringExpr `json:"SourceSecurityGroupOwnerId,omitempty"`
	// ToPort docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-rule.html#cfn-ec2-security-group-rule-toport
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

// EC2SpotFleetBlockDeviceMapping represents the AWS::EC2::SpotFleet.BlockDeviceMapping CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-blockdevicemappings.html
type EC2SpotFleetBlockDeviceMapping struct {
	// DeviceName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-blockdevicemappings.html#cfn-ec2-spotfleet-blockdevicemapping-devicename
	DeviceName *StringExpr `json:"DeviceName,omitempty" validate:"dive,required"`
	// Ebs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-blockdevicemappings.html#cfn-ec2-spotfleet-blockdevicemapping-ebs
	Ebs *EC2SpotFleetEbsBlockDevice `json:"Ebs,omitempty"`
	// NoDevice docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-blockdevicemappings.html#cfn-ec2-spotfleet-blockdevicemapping-nodevice
	NoDevice *StringExpr `json:"NoDevice,omitempty"`
	// VirtualName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-blockdevicemappings.html#cfn-ec2-spotfleet-blockdevicemapping-virtualname
	VirtualName *StringExpr `json:"VirtualName,omitempty"`
}

// EC2SpotFleetBlockDeviceMappingList represents a list of EC2SpotFleetBlockDeviceMapping
type EC2SpotFleetBlockDeviceMappingList []EC2SpotFleetBlockDeviceMapping

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2SpotFleetBlockDeviceMappingList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2SpotFleetBlockDeviceMapping{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2SpotFleetBlockDeviceMappingList{item}
		return nil
	}
	list := []EC2SpotFleetBlockDeviceMapping{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2SpotFleetBlockDeviceMappingList(list)
		return nil
	}
	return err
}

// EC2SpotFleetEbsBlockDevice represents the AWS::EC2::SpotFleet.EbsBlockDevice CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-blockdevicemappings-ebs.html
type EC2SpotFleetEbsBlockDevice struct {
	// DeleteOnTermination docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-blockdevicemappings-ebs.html#cfn-ec2-spotfleet-ebsblockdevice-deleteontermination
	DeleteOnTermination *BoolExpr `json:"DeleteOnTermination,omitempty"`
	// Encrypted docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-blockdevicemappings-ebs.html#cfn-ec2-spotfleet-ebsblockdevice-encrypted
	Encrypted *BoolExpr `json:"Encrypted,omitempty"`
	// Iops docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-blockdevicemappings-ebs.html#cfn-ec2-spotfleet-ebsblockdevice-iops
	Iops *IntegerExpr `json:"Iops,omitempty"`
	// SnapshotID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-blockdevicemappings-ebs.html#cfn-ec2-spotfleet-ebsblockdevice-snapshotid
	SnapshotID *StringExpr `json:"SnapshotId,omitempty"`
	// VolumeSize docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-blockdevicemappings-ebs.html#cfn-ec2-spotfleet-ebsblockdevice-volumesize
	VolumeSize *IntegerExpr `json:"VolumeSize,omitempty"`
	// VolumeType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-blockdevicemappings-ebs.html#cfn-ec2-spotfleet-ebsblockdevice-volumetype
	VolumeType *StringExpr `json:"VolumeType,omitempty"`
}

// EC2SpotFleetEbsBlockDeviceList represents a list of EC2SpotFleetEbsBlockDevice
type EC2SpotFleetEbsBlockDeviceList []EC2SpotFleetEbsBlockDevice

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2SpotFleetEbsBlockDeviceList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2SpotFleetEbsBlockDevice{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2SpotFleetEbsBlockDeviceList{item}
		return nil
	}
	list := []EC2SpotFleetEbsBlockDevice{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2SpotFleetEbsBlockDeviceList(list)
		return nil
	}
	return err
}

// EC2SpotFleetGroupIDentifier represents the AWS::EC2::SpotFleet.GroupIdentifier CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-securitygroups.html
type EC2SpotFleetGroupIDentifier struct {
	// GroupID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-securitygroups.html#cfn-ec2-spotfleet-groupidentifier-groupid
	GroupID *StringExpr `json:"GroupId,omitempty" validate:"dive,required"`
}

// EC2SpotFleetGroupIDentifierList represents a list of EC2SpotFleetGroupIDentifier
type EC2SpotFleetGroupIDentifierList []EC2SpotFleetGroupIDentifier

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2SpotFleetGroupIDentifierList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2SpotFleetGroupIDentifier{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2SpotFleetGroupIDentifierList{item}
		return nil
	}
	list := []EC2SpotFleetGroupIDentifier{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2SpotFleetGroupIDentifierList(list)
		return nil
	}
	return err
}

// EC2SpotFleetIamInstanceProfileSpecification represents the AWS::EC2::SpotFleet.IamInstanceProfileSpecification CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-iaminstanceprofile.html
type EC2SpotFleetIamInstanceProfileSpecification struct {
	// Arn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-iaminstanceprofile.html#cfn-ec2-spotfleet-iaminstanceprofilespecification-arn
	Arn *StringExpr `json:"Arn,omitempty"`
}

// EC2SpotFleetIamInstanceProfileSpecificationList represents a list of EC2SpotFleetIamInstanceProfileSpecification
type EC2SpotFleetIamInstanceProfileSpecificationList []EC2SpotFleetIamInstanceProfileSpecification

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2SpotFleetIamInstanceProfileSpecificationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2SpotFleetIamInstanceProfileSpecification{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2SpotFleetIamInstanceProfileSpecificationList{item}
		return nil
	}
	list := []EC2SpotFleetIamInstanceProfileSpecification{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2SpotFleetIamInstanceProfileSpecificationList(list)
		return nil
	}
	return err
}

// EC2SpotFleetInstanceIPv6Address represents the AWS::EC2::SpotFleet.InstanceIpv6Address CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-instanceipv6address.html
type EC2SpotFleetInstanceIPv6Address struct {
	// IPv6Address docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-instanceipv6address.html#cfn-ec2-spotfleet-instanceipv6address-ipv6address
	IPv6Address *StringExpr `json:"Ipv6Address,omitempty" validate:"dive,required"`
}

// EC2SpotFleetInstanceIPv6AddressList represents a list of EC2SpotFleetInstanceIPv6Address
type EC2SpotFleetInstanceIPv6AddressList []EC2SpotFleetInstanceIPv6Address

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2SpotFleetInstanceIPv6AddressList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2SpotFleetInstanceIPv6Address{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2SpotFleetInstanceIPv6AddressList{item}
		return nil
	}
	list := []EC2SpotFleetInstanceIPv6Address{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2SpotFleetInstanceIPv6AddressList(list)
		return nil
	}
	return err
}

// EC2SpotFleetInstanceNetworkInterfaceSpecification represents the AWS::EC2::SpotFleet.InstanceNetworkInterfaceSpecification CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces.html
type EC2SpotFleetInstanceNetworkInterfaceSpecification struct {
	// AssociatePublicIPAddress docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces.html#cfn-ec2-spotfleet-instancenetworkinterfacespecification-associatepublicipaddress
	AssociatePublicIPAddress *BoolExpr `json:"AssociatePublicIpAddress,omitempty"`
	// DeleteOnTermination docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces.html#cfn-ec2-spotfleet-instancenetworkinterfacespecification-deleteontermination
	DeleteOnTermination *BoolExpr `json:"DeleteOnTermination,omitempty"`
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces.html#cfn-ec2-spotfleet-instancenetworkinterfacespecification-description
	Description *StringExpr `json:"Description,omitempty"`
	// DeviceIndex docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces.html#cfn-ec2-spotfleet-instancenetworkinterfacespecification-deviceindex
	DeviceIndex *IntegerExpr `json:"DeviceIndex,omitempty"`
	// Groups docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces.html#cfn-ec2-spotfleet-instancenetworkinterfacespecification-groups
	Groups *StringListExpr `json:"Groups,omitempty"`
	// IPv6AddressCount docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces.html#cfn-ec2-spotfleet-instancenetworkinterfacespecification-ipv6addresscount
	IPv6AddressCount *IntegerExpr `json:"Ipv6AddressCount,omitempty"`
	// IPv6Addresses docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces.html#cfn-ec2-spotfleet-instancenetworkinterfacespecification-ipv6addresses
	IPv6Addresses *EC2SpotFleetInstanceIPv6AddressList `json:"Ipv6Addresses,omitempty"`
	// NetworkInterfaceID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces.html#cfn-ec2-spotfleet-instancenetworkinterfacespecification-networkinterfaceid
	NetworkInterfaceID *StringExpr `json:"NetworkInterfaceId,omitempty"`
	// PrivateIPAddresses docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces.html#cfn-ec2-spotfleet-instancenetworkinterfacespecification-privateipaddresses
	PrivateIPAddresses *EC2SpotFleetPrivateIPAddressSpecificationList `json:"PrivateIpAddresses,omitempty"`
	// SecondaryPrivateIPAddressCount docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces.html#cfn-ec2-spotfleet-instancenetworkinterfacespecification-secondaryprivateipaddresscount
	SecondaryPrivateIPAddressCount *IntegerExpr `json:"SecondaryPrivateIpAddressCount,omitempty"`
	// SubnetID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces.html#cfn-ec2-spotfleet-instancenetworkinterfacespecification-subnetid
	SubnetID *StringExpr `json:"SubnetId,omitempty"`
}

// EC2SpotFleetInstanceNetworkInterfaceSpecificationList represents a list of EC2SpotFleetInstanceNetworkInterfaceSpecification
type EC2SpotFleetInstanceNetworkInterfaceSpecificationList []EC2SpotFleetInstanceNetworkInterfaceSpecification

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2SpotFleetInstanceNetworkInterfaceSpecificationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2SpotFleetInstanceNetworkInterfaceSpecification{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2SpotFleetInstanceNetworkInterfaceSpecificationList{item}
		return nil
	}
	list := []EC2SpotFleetInstanceNetworkInterfaceSpecification{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2SpotFleetInstanceNetworkInterfaceSpecificationList(list)
		return nil
	}
	return err
}

// EC2SpotFleetPrivateIPAddressSpecification represents the AWS::EC2::SpotFleet.PrivateIpAddressSpecification CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces-privateipaddresses.html
type EC2SpotFleetPrivateIPAddressSpecification struct {
	// Primary docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces-privateipaddresses.html#cfn-ec2-spotfleet-privateipaddressspecification-primary
	Primary *BoolExpr `json:"Primary,omitempty"`
	// PrivateIPAddress docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces-privateipaddresses.html#cfn-ec2-spotfleet-privateipaddressspecification-privateipaddress
	PrivateIPAddress *StringExpr `json:"PrivateIpAddress,omitempty" validate:"dive,required"`
}

// EC2SpotFleetPrivateIPAddressSpecificationList represents a list of EC2SpotFleetPrivateIPAddressSpecification
type EC2SpotFleetPrivateIPAddressSpecificationList []EC2SpotFleetPrivateIPAddressSpecification

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2SpotFleetPrivateIPAddressSpecificationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2SpotFleetPrivateIPAddressSpecification{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2SpotFleetPrivateIPAddressSpecificationList{item}
		return nil
	}
	list := []EC2SpotFleetPrivateIPAddressSpecification{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2SpotFleetPrivateIPAddressSpecificationList(list)
		return nil
	}
	return err
}

// EC2SpotFleetSpotFleetLaunchSpecification represents the AWS::EC2::SpotFleet.SpotFleetLaunchSpecification CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications.html
type EC2SpotFleetSpotFleetLaunchSpecification struct {
	// BlockDeviceMappings docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-blockdevicemappings
	BlockDeviceMappings *EC2SpotFleetBlockDeviceMappingList `json:"BlockDeviceMappings,omitempty"`
	// EbsOptimized docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-ebsoptimized
	EbsOptimized *BoolExpr `json:"EbsOptimized,omitempty"`
	// IamInstanceProfile docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-iaminstanceprofile
	IamInstanceProfile *EC2SpotFleetIamInstanceProfileSpecification `json:"IamInstanceProfile,omitempty"`
	// ImageID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-imageid
	ImageID *StringExpr `json:"ImageId,omitempty" validate:"dive,required"`
	// InstanceType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-instancetype
	InstanceType *StringExpr `json:"InstanceType,omitempty" validate:"dive,required"`
	// KernelID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-kernelid
	KernelID *StringExpr `json:"KernelId,omitempty"`
	// KeyName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-keyname
	KeyName *StringExpr `json:"KeyName,omitempty"`
	// Monitoring docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-monitoring
	Monitoring *EC2SpotFleetSpotFleetMonitoring `json:"Monitoring,omitempty"`
	// NetworkInterfaces docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-networkinterfaces
	NetworkInterfaces *EC2SpotFleetInstanceNetworkInterfaceSpecificationList `json:"NetworkInterfaces,omitempty"`
	// Placement docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-placement
	Placement *EC2SpotFleetSpotPlacement `json:"Placement,omitempty"`
	// RamdiskID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-ramdiskid
	RamdiskID *StringExpr `json:"RamdiskId,omitempty"`
	// SecurityGroups docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-securitygroups
	SecurityGroups *EC2SpotFleetGroupIDentifierList `json:"SecurityGroups,omitempty"`
	// SpotPrice docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-spotprice
	SpotPrice *StringExpr `json:"SpotPrice,omitempty"`
	// SubnetID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-subnetid
	SubnetID *StringExpr `json:"SubnetId,omitempty"`
	// UserData docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-userdata
	UserData *StringExpr `json:"UserData,omitempty"`
	// WeightedCapacity docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-weightedcapacity
	WeightedCapacity *IntegerExpr `json:"WeightedCapacity,omitempty"`
}

// EC2SpotFleetSpotFleetLaunchSpecificationList represents a list of EC2SpotFleetSpotFleetLaunchSpecification
type EC2SpotFleetSpotFleetLaunchSpecificationList []EC2SpotFleetSpotFleetLaunchSpecification

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2SpotFleetSpotFleetLaunchSpecificationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2SpotFleetSpotFleetLaunchSpecification{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2SpotFleetSpotFleetLaunchSpecificationList{item}
		return nil
	}
	list := []EC2SpotFleetSpotFleetLaunchSpecification{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2SpotFleetSpotFleetLaunchSpecificationList(list)
		return nil
	}
	return err
}

// EC2SpotFleetSpotFleetMonitoring represents the AWS::EC2::SpotFleet.SpotFleetMonitoring CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-monitoring.html
type EC2SpotFleetSpotFleetMonitoring struct {
	// Enabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-monitoring.html#cfn-ec2-spotfleet-spotfleetmonitoring-enabled
	Enabled *BoolExpr `json:"Enabled,omitempty"`
}

// EC2SpotFleetSpotFleetMonitoringList represents a list of EC2SpotFleetSpotFleetMonitoring
type EC2SpotFleetSpotFleetMonitoringList []EC2SpotFleetSpotFleetMonitoring

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2SpotFleetSpotFleetMonitoringList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2SpotFleetSpotFleetMonitoring{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2SpotFleetSpotFleetMonitoringList{item}
		return nil
	}
	list := []EC2SpotFleetSpotFleetMonitoring{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2SpotFleetSpotFleetMonitoringList(list)
		return nil
	}
	return err
}

// EC2SpotFleetSpotFleetRequestConfigData represents the AWS::EC2::SpotFleet.SpotFleetRequestConfigData CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html
type EC2SpotFleetSpotFleetRequestConfigData struct {
	// AllocationStrategy docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-allocationstrategy
	AllocationStrategy *StringExpr `json:"AllocationStrategy,omitempty"`
	// ExcessCapacityTerminationPolicy docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-excesscapacityterminationpolicy
	ExcessCapacityTerminationPolicy *StringExpr `json:"ExcessCapacityTerminationPolicy,omitempty"`
	// IamFleetRole docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-iamfleetrole
	IamFleetRole *StringExpr `json:"IamFleetRole,omitempty" validate:"dive,required"`
	// LaunchSpecifications docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications
	LaunchSpecifications *EC2SpotFleetSpotFleetLaunchSpecificationList `json:"LaunchSpecifications,omitempty" validate:"dive,required"`
	// SpotPrice docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-spotprice
	SpotPrice *StringExpr `json:"SpotPrice,omitempty" validate:"dive,required"`
	// TargetCapacity docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-targetcapacity
	TargetCapacity *IntegerExpr `json:"TargetCapacity,omitempty" validate:"dive,required"`
	// TerminateInstancesWithExpiration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-terminateinstanceswithexpiration
	TerminateInstancesWithExpiration *BoolExpr `json:"TerminateInstancesWithExpiration,omitempty"`
	// ValidFrom docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-validfrom
	ValidFrom *StringExpr `json:"ValidFrom,omitempty"`
	// ValidUntil docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-validuntil
	ValidUntil *StringExpr `json:"ValidUntil,omitempty"`
}

// EC2SpotFleetSpotFleetRequestConfigDataList represents a list of EC2SpotFleetSpotFleetRequestConfigData
type EC2SpotFleetSpotFleetRequestConfigDataList []EC2SpotFleetSpotFleetRequestConfigData

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2SpotFleetSpotFleetRequestConfigDataList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2SpotFleetSpotFleetRequestConfigData{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2SpotFleetSpotFleetRequestConfigDataList{item}
		return nil
	}
	list := []EC2SpotFleetSpotFleetRequestConfigData{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2SpotFleetSpotFleetRequestConfigDataList(list)
		return nil
	}
	return err
}

// EC2SpotFleetSpotPlacement represents the AWS::EC2::SpotFleet.SpotPlacement CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-placement.html
type EC2SpotFleetSpotPlacement struct {
	// AvailabilityZone docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-placement.html#cfn-ec2-spotfleet-spotplacement-availabilityzone
	AvailabilityZone *StringExpr `json:"AvailabilityZone,omitempty"`
	// GroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-placement.html#cfn-ec2-spotfleet-spotplacement-groupname
	GroupName *StringExpr `json:"GroupName,omitempty"`
}

// EC2SpotFleetSpotPlacementList represents a list of EC2SpotFleetSpotPlacement
type EC2SpotFleetSpotPlacementList []EC2SpotFleetSpotPlacement

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EC2SpotFleetSpotPlacementList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EC2SpotFleetSpotPlacement{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EC2SpotFleetSpotPlacementList{item}
		return nil
	}
	list := []EC2SpotFleetSpotPlacement{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EC2SpotFleetSpotPlacementList(list)
		return nil
	}
	return err
}

// ECSServiceDeploymentConfiguration represents the AWS::ECS::Service.DeploymentConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentconfiguration.html
type ECSServiceDeploymentConfiguration struct {
	// MaximumPercent docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentconfiguration.html#cfn-ecs-service-deploymentconfiguration-maximumpercent
	MaximumPercent *IntegerExpr `json:"MaximumPercent,omitempty"`
	// MinimumHealthyPercent docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentconfiguration.html#cfn-ecs-service-deploymentconfiguration-minimumhealthypercent
	MinimumHealthyPercent *IntegerExpr `json:"MinimumHealthyPercent,omitempty"`
}

// ECSServiceDeploymentConfigurationList represents a list of ECSServiceDeploymentConfiguration
type ECSServiceDeploymentConfigurationList []ECSServiceDeploymentConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ECSServiceDeploymentConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ECSServiceDeploymentConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ECSServiceDeploymentConfigurationList{item}
		return nil
	}
	list := []ECSServiceDeploymentConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ECSServiceDeploymentConfigurationList(list)
		return nil
	}
	return err
}

// ECSServiceLoadBalancer represents the AWS::ECS::Service.LoadBalancer CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-loadbalancers.html
type ECSServiceLoadBalancer struct {
	// ContainerName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-loadbalancers.html#cfn-ecs-service-loadbalancers-containername
	ContainerName *StringExpr `json:"ContainerName,omitempty"`
	// ContainerPort docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-loadbalancers.html#cfn-ecs-service-loadbalancers-containerport
	ContainerPort *IntegerExpr `json:"ContainerPort,omitempty" validate:"dive,required"`
	// LoadBalancerName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-loadbalancers.html#cfn-ecs-service-loadbalancers-loadbalancername
	LoadBalancerName *StringExpr `json:"LoadBalancerName,omitempty"`
	// TargetGroupArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-loadbalancers.html#cfn-ecs-service-loadbalancers-targetgrouparn
	TargetGroupArn *StringExpr `json:"TargetGroupArn,omitempty"`
}

// ECSServiceLoadBalancerList represents a list of ECSServiceLoadBalancer
type ECSServiceLoadBalancerList []ECSServiceLoadBalancer

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ECSServiceLoadBalancerList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ECSServiceLoadBalancer{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ECSServiceLoadBalancerList{item}
		return nil
	}
	list := []ECSServiceLoadBalancer{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ECSServiceLoadBalancerList(list)
		return nil
	}
	return err
}

// ECSServicePlacementConstraint represents the AWS::ECS::Service.PlacementConstraint CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-placementconstraint.html
type ECSServicePlacementConstraint struct {
	// Expression docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-placementconstraint.html#cfn-ecs-service-placementconstraint-expression
	Expression *StringExpr `json:"Expression,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-placementconstraint.html#cfn-ecs-service-placementconstraint-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
}

// ECSServicePlacementConstraintList represents a list of ECSServicePlacementConstraint
type ECSServicePlacementConstraintList []ECSServicePlacementConstraint

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ECSServicePlacementConstraintList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ECSServicePlacementConstraint{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ECSServicePlacementConstraintList{item}
		return nil
	}
	list := []ECSServicePlacementConstraint{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ECSServicePlacementConstraintList(list)
		return nil
	}
	return err
}

// ECSServicePlacementStrategy represents the AWS::ECS::Service.PlacementStrategy CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-placementstrategy.html
type ECSServicePlacementStrategy struct {
	// Field docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-placementstrategy.html#cfn-ecs-service-placementstrategy-field
	Field *StringExpr `json:"Field,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-placementstrategy.html#cfn-ecs-service-placementstrategy-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
}

// ECSServicePlacementStrategyList represents a list of ECSServicePlacementStrategy
type ECSServicePlacementStrategyList []ECSServicePlacementStrategy

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ECSServicePlacementStrategyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ECSServicePlacementStrategy{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ECSServicePlacementStrategyList{item}
		return nil
	}
	list := []ECSServicePlacementStrategy{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ECSServicePlacementStrategyList(list)
		return nil
	}
	return err
}

// ECSTaskDefinitionContainerDefinition represents the AWS::ECS::TaskDefinition.ContainerDefinition CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html
type ECSTaskDefinitionContainerDefinition struct {
	// Command docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-command
	Command *StringListExpr `json:"Command,omitempty"`
	// CPU docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-cpu
	CPU *IntegerExpr `json:"Cpu,omitempty"`
	// DisableNetworking docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-disablenetworking
	DisableNetworking *BoolExpr `json:"DisableNetworking,omitempty"`
	// DNSSearchDomains docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-dnssearchdomains
	DNSSearchDomains *StringListExpr `json:"DnsSearchDomains,omitempty"`
	// DNSServers docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-dnsservers
	DNSServers *StringListExpr `json:"DnsServers,omitempty"`
	// DockerLabels docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-dockerlabels
	DockerLabels interface{} `json:"DockerLabels,omitempty"`
	// DockerSecurityOptions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-dockersecurityoptions
	DockerSecurityOptions *StringListExpr `json:"DockerSecurityOptions,omitempty"`
	// EntryPoint docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-entrypoint
	EntryPoint *StringListExpr `json:"EntryPoint,omitempty"`
	// Environment docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-environment
	Environment *ECSTaskDefinitionKeyValuePairList `json:"Environment,omitempty"`
	// Essential docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-essential
	Essential *BoolExpr `json:"Essential,omitempty"`
	// ExtraHosts docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-extrahosts
	ExtraHosts *ECSTaskDefinitionHostEntryList `json:"ExtraHosts,omitempty"`
	// Hostname docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-hostname
	Hostname *StringExpr `json:"Hostname,omitempty"`
	// Image docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-image
	Image *StringExpr `json:"Image,omitempty"`
	// Links docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-links
	Links *StringListExpr `json:"Links,omitempty"`
	// LogConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-logconfiguration
	LogConfiguration *ECSTaskDefinitionLogConfiguration `json:"LogConfiguration,omitempty"`
	// Memory docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-memory
	Memory *IntegerExpr `json:"Memory,omitempty"`
	// MemoryReservation docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-memoryreservation
	MemoryReservation *IntegerExpr `json:"MemoryReservation,omitempty"`
	// MountPoints docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-mountpoints
	MountPoints *ECSTaskDefinitionMountPointList `json:"MountPoints,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-name
	Name *StringExpr `json:"Name,omitempty"`
	// PortMappings docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-portmappings
	PortMappings *ECSTaskDefinitionPortMappingList `json:"PortMappings,omitempty"`
	// Privileged docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-privileged
	Privileged *BoolExpr `json:"Privileged,omitempty"`
	// ReadonlyRootFilesystem docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-readonlyrootfilesystem
	ReadonlyRootFilesystem *BoolExpr `json:"ReadonlyRootFilesystem,omitempty"`
	// Ulimits docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-ulimits
	Ulimits *ECSTaskDefinitionUlimitList `json:"Ulimits,omitempty"`
	// User docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-user
	User *StringExpr `json:"User,omitempty"`
	// VolumesFrom docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-volumesfrom
	VolumesFrom *ECSTaskDefinitionVolumeFromList `json:"VolumesFrom,omitempty"`
	// WorkingDirectory docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-workingdirectory
	WorkingDirectory *StringExpr `json:"WorkingDirectory,omitempty"`
}

// ECSTaskDefinitionContainerDefinitionList represents a list of ECSTaskDefinitionContainerDefinition
type ECSTaskDefinitionContainerDefinitionList []ECSTaskDefinitionContainerDefinition

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ECSTaskDefinitionContainerDefinitionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ECSTaskDefinitionContainerDefinition{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ECSTaskDefinitionContainerDefinitionList{item}
		return nil
	}
	list := []ECSTaskDefinitionContainerDefinition{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ECSTaskDefinitionContainerDefinitionList(list)
		return nil
	}
	return err
}

// ECSTaskDefinitionHostEntry represents the AWS::ECS::TaskDefinition.HostEntry CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-hostentry.html
type ECSTaskDefinitionHostEntry struct {
	// Hostname docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-hostentry.html#cfn-ecs-taskdefinition-containerdefinition-hostentry-hostname
	Hostname *StringExpr `json:"Hostname,omitempty" validate:"dive,required"`
	// IPAddress docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-hostentry.html#cfn-ecs-taskdefinition-containerdefinition-hostentry-ipaddress
	IPAddress *StringExpr `json:"IpAddress,omitempty" validate:"dive,required"`
}

// ECSTaskDefinitionHostEntryList represents a list of ECSTaskDefinitionHostEntry
type ECSTaskDefinitionHostEntryList []ECSTaskDefinitionHostEntry

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ECSTaskDefinitionHostEntryList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ECSTaskDefinitionHostEntry{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ECSTaskDefinitionHostEntryList{item}
		return nil
	}
	list := []ECSTaskDefinitionHostEntry{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ECSTaskDefinitionHostEntryList(list)
		return nil
	}
	return err
}

// ECSTaskDefinitionHostVolumeProperties represents the AWS::ECS::TaskDefinition.HostVolumeProperties CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volumes-host.html
type ECSTaskDefinitionHostVolumeProperties struct {
	// SourcePath docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volumes-host.html#cfn-ecs-taskdefinition-volumes-host-sourcepath
	SourcePath *StringExpr `json:"SourcePath,omitempty"`
}

// ECSTaskDefinitionHostVolumePropertiesList represents a list of ECSTaskDefinitionHostVolumeProperties
type ECSTaskDefinitionHostVolumePropertiesList []ECSTaskDefinitionHostVolumeProperties

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ECSTaskDefinitionHostVolumePropertiesList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ECSTaskDefinitionHostVolumeProperties{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ECSTaskDefinitionHostVolumePropertiesList{item}
		return nil
	}
	list := []ECSTaskDefinitionHostVolumeProperties{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ECSTaskDefinitionHostVolumePropertiesList(list)
		return nil
	}
	return err
}

// ECSTaskDefinitionKeyValuePair represents the AWS::ECS::TaskDefinition.KeyValuePair CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-environment.html
type ECSTaskDefinitionKeyValuePair struct {
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-environment.html#cfn-ecs-taskdefinition-containerdefinition-environment-name
	Name *StringExpr `json:"Name,omitempty"`
	// Value docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-environment.html#cfn-ecs-taskdefinition-containerdefinition-environment-value
	Value *StringExpr `json:"Value,omitempty"`
}

// ECSTaskDefinitionKeyValuePairList represents a list of ECSTaskDefinitionKeyValuePair
type ECSTaskDefinitionKeyValuePairList []ECSTaskDefinitionKeyValuePair

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ECSTaskDefinitionKeyValuePairList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ECSTaskDefinitionKeyValuePair{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ECSTaskDefinitionKeyValuePairList{item}
		return nil
	}
	list := []ECSTaskDefinitionKeyValuePair{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ECSTaskDefinitionKeyValuePairList(list)
		return nil
	}
	return err
}

// ECSTaskDefinitionLogConfiguration represents the AWS::ECS::TaskDefinition.LogConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-logconfiguration.html
type ECSTaskDefinitionLogConfiguration struct {
	// LogDriver docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-logconfiguration.html#cfn-ecs-taskdefinition-containerdefinition-logconfiguration-logdriver
	LogDriver *StringExpr `json:"LogDriver,omitempty" validate:"dive,required"`
	// Options docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-logconfiguration.html#cfn-ecs-taskdefinition-containerdefinition-logconfiguration-options
	Options interface{} `json:"Options,omitempty"`
}

// ECSTaskDefinitionLogConfigurationList represents a list of ECSTaskDefinitionLogConfiguration
type ECSTaskDefinitionLogConfigurationList []ECSTaskDefinitionLogConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ECSTaskDefinitionLogConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ECSTaskDefinitionLogConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ECSTaskDefinitionLogConfigurationList{item}
		return nil
	}
	list := []ECSTaskDefinitionLogConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ECSTaskDefinitionLogConfigurationList(list)
		return nil
	}
	return err
}

// ECSTaskDefinitionMountPoint represents the AWS::ECS::TaskDefinition.MountPoint CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-mountpoints.html
type ECSTaskDefinitionMountPoint struct {
	// ContainerPath docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-mountpoints.html#cfn-ecs-taskdefinition-containerdefinition-mountpoints-containerpath
	ContainerPath *StringExpr `json:"ContainerPath,omitempty"`
	// ReadOnly docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-mountpoints.html#cfn-ecs-taskdefinition-containerdefinition-mountpoints-readonly
	ReadOnly *BoolExpr `json:"ReadOnly,omitempty"`
	// SourceVolume docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-mountpoints.html#cfn-ecs-taskdefinition-containerdefinition-mountpoints-sourcevolume
	SourceVolume *StringExpr `json:"SourceVolume,omitempty"`
}

// ECSTaskDefinitionMountPointList represents a list of ECSTaskDefinitionMountPoint
type ECSTaskDefinitionMountPointList []ECSTaskDefinitionMountPoint

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ECSTaskDefinitionMountPointList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ECSTaskDefinitionMountPoint{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ECSTaskDefinitionMountPointList{item}
		return nil
	}
	list := []ECSTaskDefinitionMountPoint{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ECSTaskDefinitionMountPointList(list)
		return nil
	}
	return err
}

// ECSTaskDefinitionPortMapping represents the AWS::ECS::TaskDefinition.PortMapping CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-portmappings.html
type ECSTaskDefinitionPortMapping struct {
	// ContainerPort docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-portmappings.html#cfn-ecs-taskdefinition-containerdefinition-portmappings-containerport
	ContainerPort *IntegerExpr `json:"ContainerPort,omitempty"`
	// HostPort docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-portmappings.html#cfn-ecs-taskdefinition-containerdefinition-portmappings-readonly
	HostPort *IntegerExpr `json:"HostPort,omitempty"`
	// Protocol docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-portmappings.html#cfn-ecs-taskdefinition-containerdefinition-portmappings-sourcevolume
	Protocol *StringExpr `json:"Protocol,omitempty"`
}

// ECSTaskDefinitionPortMappingList represents a list of ECSTaskDefinitionPortMapping
type ECSTaskDefinitionPortMappingList []ECSTaskDefinitionPortMapping

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ECSTaskDefinitionPortMappingList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ECSTaskDefinitionPortMapping{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ECSTaskDefinitionPortMappingList{item}
		return nil
	}
	list := []ECSTaskDefinitionPortMapping{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ECSTaskDefinitionPortMappingList(list)
		return nil
	}
	return err
}

// ECSTaskDefinitionTaskDefinitionPlacementConstraint represents the AWS::ECS::TaskDefinition.TaskDefinitionPlacementConstraint CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-taskdefinitionplacementconstraint.html
type ECSTaskDefinitionTaskDefinitionPlacementConstraint struct {
	// Expression docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-taskdefinitionplacementconstraint.html#cfn-ecs-taskdefinition-taskdefinitionplacementconstraint-expression
	Expression *StringExpr `json:"Expression,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-taskdefinitionplacementconstraint.html#cfn-ecs-taskdefinition-taskdefinitionplacementconstraint-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
}

// ECSTaskDefinitionTaskDefinitionPlacementConstraintList represents a list of ECSTaskDefinitionTaskDefinitionPlacementConstraint
type ECSTaskDefinitionTaskDefinitionPlacementConstraintList []ECSTaskDefinitionTaskDefinitionPlacementConstraint

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ECSTaskDefinitionTaskDefinitionPlacementConstraintList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ECSTaskDefinitionTaskDefinitionPlacementConstraint{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ECSTaskDefinitionTaskDefinitionPlacementConstraintList{item}
		return nil
	}
	list := []ECSTaskDefinitionTaskDefinitionPlacementConstraint{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ECSTaskDefinitionTaskDefinitionPlacementConstraintList(list)
		return nil
	}
	return err
}

// ECSTaskDefinitionUlimit represents the AWS::ECS::TaskDefinition.Ulimit CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-ulimit.html
type ECSTaskDefinitionUlimit struct {
	// HardLimit docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-ulimit.html#cfn-ecs-taskdefinition-containerdefinition-ulimit-hardlimit
	HardLimit *IntegerExpr `json:"HardLimit,omitempty" validate:"dive,required"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-ulimit.html#cfn-ecs-taskdefinition-containerdefinition-ulimit-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// SoftLimit docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-ulimit.html#cfn-ecs-taskdefinition-containerdefinition-ulimit-softlimit
	SoftLimit *IntegerExpr `json:"SoftLimit,omitempty" validate:"dive,required"`
}

// ECSTaskDefinitionUlimitList represents a list of ECSTaskDefinitionUlimit
type ECSTaskDefinitionUlimitList []ECSTaskDefinitionUlimit

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ECSTaskDefinitionUlimitList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ECSTaskDefinitionUlimit{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ECSTaskDefinitionUlimitList{item}
		return nil
	}
	list := []ECSTaskDefinitionUlimit{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ECSTaskDefinitionUlimitList(list)
		return nil
	}
	return err
}

// ECSTaskDefinitionVolume represents the AWS::ECS::TaskDefinition.Volume CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volumes.html
type ECSTaskDefinitionVolume struct {
	// Host docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volumes.html#cfn-ecs-taskdefinition-volumes-host
	Host *ECSTaskDefinitionHostVolumeProperties `json:"Host,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volumes.html#cfn-ecs-taskdefinition-volumes-name
	Name *StringExpr `json:"Name,omitempty"`
}

// ECSTaskDefinitionVolumeList represents a list of ECSTaskDefinitionVolume
type ECSTaskDefinitionVolumeList []ECSTaskDefinitionVolume

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ECSTaskDefinitionVolumeList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ECSTaskDefinitionVolume{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ECSTaskDefinitionVolumeList{item}
		return nil
	}
	list := []ECSTaskDefinitionVolume{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ECSTaskDefinitionVolumeList(list)
		return nil
	}
	return err
}

// ECSTaskDefinitionVolumeFrom represents the AWS::ECS::TaskDefinition.VolumeFrom CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-volumesfrom.html
type ECSTaskDefinitionVolumeFrom struct {
	// ReadOnly docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-volumesfrom.html#cfn-ecs-taskdefinition-containerdefinition-volumesfrom-readonly
	ReadOnly *BoolExpr `json:"ReadOnly,omitempty"`
	// SourceContainer docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-volumesfrom.html#cfn-ecs-taskdefinition-containerdefinition-volumesfrom-sourcecontainer
	SourceContainer *StringExpr `json:"SourceContainer,omitempty"`
}

// ECSTaskDefinitionVolumeFromList represents a list of ECSTaskDefinitionVolumeFrom
type ECSTaskDefinitionVolumeFromList []ECSTaskDefinitionVolumeFrom

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ECSTaskDefinitionVolumeFromList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ECSTaskDefinitionVolumeFrom{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ECSTaskDefinitionVolumeFromList{item}
		return nil
	}
	list := []ECSTaskDefinitionVolumeFrom{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ECSTaskDefinitionVolumeFromList(list)
		return nil
	}
	return err
}

// EFSFileSystemElasticFileSystemTag represents the AWS::EFS::FileSystem.ElasticFileSystemTag CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-filesystemtags.html
type EFSFileSystemElasticFileSystemTag struct {
	// Key docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-filesystemtags.html#cfn-efs-filesystem-filesystemtags-key
	Key *StringExpr `json:"Key,omitempty" validate:"dive,required"`
	// Value docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-filesystemtags.html#cfn-efs-filesystem-filesystemtags-value
	Value *StringExpr `json:"Value,omitempty" validate:"dive,required"`
}

// EFSFileSystemElasticFileSystemTagList represents a list of EFSFileSystemElasticFileSystemTag
type EFSFileSystemElasticFileSystemTagList []EFSFileSystemElasticFileSystemTag

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EFSFileSystemElasticFileSystemTagList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EFSFileSystemElasticFileSystemTag{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EFSFileSystemElasticFileSystemTagList{item}
		return nil
	}
	list := []EFSFileSystemElasticFileSystemTag{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EFSFileSystemElasticFileSystemTagList(list)
		return nil
	}
	return err
}

// EMRClusterApplication represents the AWS::EMR::Cluster.Application CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-application.html
type EMRClusterApplication struct {
	// AdditionalInfo docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-application.html#cfn-emr-cluster-application-additionalinfo
	AdditionalInfo interface{} `json:"AdditionalInfo,omitempty"`
	// Args docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-application.html#cfn-emr-cluster-application-args
	Args *StringListExpr `json:"Args,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-application.html#cfn-emr-cluster-application-name
	Name *StringExpr `json:"Name,omitempty"`
	// Version docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-application.html#cfn-emr-cluster-application-version
	Version *StringExpr `json:"Version,omitempty"`
}

// EMRClusterApplicationList represents a list of EMRClusterApplication
type EMRClusterApplicationList []EMRClusterApplication

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRClusterApplicationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRClusterApplication{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRClusterApplicationList{item}
		return nil
	}
	list := []EMRClusterApplication{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRClusterApplicationList(list)
		return nil
	}
	return err
}

// EMRClusterAutoScalingPolicy represents the AWS::EMR::Cluster.AutoScalingPolicy CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-autoscalingpolicy.html
type EMRClusterAutoScalingPolicy struct {
	// Constraints docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-autoscalingpolicy.html#cfn-elasticmapreduce-cluster-autoscalingpolicy-constraints
	Constraints *EMRClusterScalingConstraints `json:"Constraints,omitempty" validate:"dive,required"`
	// Rules docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-autoscalingpolicy.html#cfn-elasticmapreduce-cluster-autoscalingpolicy-rules
	Rules *EMRClusterScalingRuleList `json:"Rules,omitempty" validate:"dive,required"`
}

// EMRClusterAutoScalingPolicyList represents a list of EMRClusterAutoScalingPolicy
type EMRClusterAutoScalingPolicyList []EMRClusterAutoScalingPolicy

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRClusterAutoScalingPolicyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRClusterAutoScalingPolicy{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRClusterAutoScalingPolicyList{item}
		return nil
	}
	list := []EMRClusterAutoScalingPolicy{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRClusterAutoScalingPolicyList(list)
		return nil
	}
	return err
}

// EMRClusterBootstrapActionConfig represents the AWS::EMR::Cluster.BootstrapActionConfig CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-bootstrapactionconfig.html
type EMRClusterBootstrapActionConfig struct {
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-bootstrapactionconfig.html#cfn-emr-cluster-bootstrapactionconfig-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// ScriptBootstrapAction docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-bootstrapactionconfig.html#cfn-emr-cluster-bootstrapactionconfig-scriptbootstrapaction
	ScriptBootstrapAction *EMRClusterScriptBootstrapActionConfig `json:"ScriptBootstrapAction,omitempty" validate:"dive,required"`
}

// EMRClusterBootstrapActionConfigList represents a list of EMRClusterBootstrapActionConfig
type EMRClusterBootstrapActionConfigList []EMRClusterBootstrapActionConfig

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRClusterBootstrapActionConfigList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRClusterBootstrapActionConfig{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRClusterBootstrapActionConfigList{item}
		return nil
	}
	list := []EMRClusterBootstrapActionConfig{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRClusterBootstrapActionConfigList(list)
		return nil
	}
	return err
}

// EMRClusterCloudWatchAlarmDefinition represents the AWS::EMR::Cluster.CloudWatchAlarmDefinition CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-cloudwatchalarmdefinition.html
type EMRClusterCloudWatchAlarmDefinition struct {
	// ComparisonOperator docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-cluster-cloudwatchalarmdefinition-comparisonoperator
	ComparisonOperator *StringExpr `json:"ComparisonOperator,omitempty" validate:"dive,required"`
	// Dimensions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-cluster-cloudwatchalarmdefinition-dimensions
	Dimensions *EMRClusterMetricDimensionList `json:"Dimensions,omitempty"`
	// EvaluationPeriods docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-cluster-cloudwatchalarmdefinition-evaluationperiods
	EvaluationPeriods *IntegerExpr `json:"EvaluationPeriods,omitempty"`
	// MetricName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-cluster-cloudwatchalarmdefinition-metricname
	MetricName *StringExpr `json:"MetricName,omitempty" validate:"dive,required"`
	// Namespace docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-cluster-cloudwatchalarmdefinition-namespace
	Namespace *StringExpr `json:"Namespace,omitempty"`
	// Period docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-cluster-cloudwatchalarmdefinition-period
	Period *IntegerExpr `json:"Period,omitempty" validate:"dive,required"`
	// Statistic docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-cluster-cloudwatchalarmdefinition-statistic
	Statistic *StringExpr `json:"Statistic,omitempty"`
	// Threshold docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-cluster-cloudwatchalarmdefinition-threshold
	Threshold *IntegerExpr `json:"Threshold,omitempty" validate:"dive,required"`
	// Unit docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-cluster-cloudwatchalarmdefinition-unit
	Unit *StringExpr `json:"Unit,omitempty"`
}

// EMRClusterCloudWatchAlarmDefinitionList represents a list of EMRClusterCloudWatchAlarmDefinition
type EMRClusterCloudWatchAlarmDefinitionList []EMRClusterCloudWatchAlarmDefinition

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRClusterCloudWatchAlarmDefinitionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRClusterCloudWatchAlarmDefinition{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRClusterCloudWatchAlarmDefinitionList{item}
		return nil
	}
	list := []EMRClusterCloudWatchAlarmDefinition{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRClusterCloudWatchAlarmDefinitionList(list)
		return nil
	}
	return err
}

// EMRClusterConfiguration represents the AWS::EMR::Cluster.Configuration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-configuration.html
type EMRClusterConfiguration struct {
	// Classification docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-configuration.html#cfn-emr-cluster-configuration-classification
	Classification *StringExpr `json:"Classification,omitempty"`
	// ConfigurationProperties docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-configuration.html#cfn-emr-cluster-configuration-configurationproperties
	ConfigurationProperties interface{} `json:"ConfigurationProperties,omitempty"`
	// Configurations docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-configuration.html#cfn-emr-cluster-configuration-configurations
	Configurations *EMRClusterConfigurationList `json:"Configurations,omitempty"`
}

// EMRClusterConfigurationList represents a list of EMRClusterConfiguration
type EMRClusterConfigurationList []EMRClusterConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRClusterConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRClusterConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRClusterConfigurationList{item}
		return nil
	}
	list := []EMRClusterConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRClusterConfigurationList(list)
		return nil
	}
	return err
}

// EMRClusterEbsBlockDeviceConfig represents the AWS::EMR::Cluster.EbsBlockDeviceConfig CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig.html
type EMRClusterEbsBlockDeviceConfig struct {
	// VolumeSpecification docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification
	VolumeSpecification *EMRClusterVolumeSpecification `json:"VolumeSpecification,omitempty" validate:"dive,required"`
	// VolumesPerInstance docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfig-volumesperinstance
	VolumesPerInstance *IntegerExpr `json:"VolumesPerInstance,omitempty"`
}

// EMRClusterEbsBlockDeviceConfigList represents a list of EMRClusterEbsBlockDeviceConfig
type EMRClusterEbsBlockDeviceConfigList []EMRClusterEbsBlockDeviceConfig

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRClusterEbsBlockDeviceConfigList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRClusterEbsBlockDeviceConfig{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRClusterEbsBlockDeviceConfigList{item}
		return nil
	}
	list := []EMRClusterEbsBlockDeviceConfig{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRClusterEbsBlockDeviceConfigList(list)
		return nil
	}
	return err
}

// EMRClusterEbsConfiguration represents the AWS::EMR::Cluster.EbsConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration.html
type EMRClusterEbsConfiguration struct {
	// EbsBlockDeviceConfigs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfigs
	EbsBlockDeviceConfigs *EMRClusterEbsBlockDeviceConfigList `json:"EbsBlockDeviceConfigs,omitempty"`
	// EbsOptimized docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration.html#cfn-emr-ebsconfiguration-ebsoptimized
	EbsOptimized *BoolExpr `json:"EbsOptimized,omitempty"`
}

// EMRClusterEbsConfigurationList represents a list of EMRClusterEbsConfiguration
type EMRClusterEbsConfigurationList []EMRClusterEbsConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRClusterEbsConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRClusterEbsConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRClusterEbsConfigurationList{item}
		return nil
	}
	list := []EMRClusterEbsConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRClusterEbsConfigurationList(list)
		return nil
	}
	return err
}

// EMRClusterInstanceGroupConfig represents the AWS::EMR::Cluster.InstanceGroupConfig CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig-instancegroupconfig.html
type EMRClusterInstanceGroupConfig struct {
	// AutoScalingPolicy docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig-instancegroupconfig.html#cfn-elasticmapreduce-cluster-instancegroupconfig-autoscalingpolicy
	AutoScalingPolicy *EMRClusterAutoScalingPolicy `json:"AutoScalingPolicy,omitempty"`
	// BidPrice docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig-instancegroupconfig.html#cfn-emr-cluster-jobflowinstancesconfig-instancegroupconfig-bidprice
	BidPrice *StringExpr `json:"BidPrice,omitempty"`
	// Configurations docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig-instancegroupconfig.html#cfn-emr-cluster-jobflowinstancesconfig-instancegroupconfig-configurations
	Configurations *EMRClusterConfigurationList `json:"Configurations,omitempty"`
	// EbsConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig-instancegroupconfig.html#cfn-emr-cluster-jobflowinstancesconfig-instancegroupconfigConfigurations-ebsconfiguration
	EbsConfiguration *EMRClusterEbsConfiguration `json:"EbsConfiguration,omitempty"`
	// InstanceCount docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig-instancegroupconfig.html#cfn-emr-cluster-jobflowinstancesconfig-instancegroupconfig-instancecount
	InstanceCount *IntegerExpr `json:"InstanceCount,omitempty" validate:"dive,required"`
	// InstanceType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig-instancegroupconfig.html#cfn-emr-cluster-jobflowinstancesconfig-instancegroupconfig-instancetype
	InstanceType *StringExpr `json:"InstanceType,omitempty" validate:"dive,required"`
	// Market docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig-instancegroupconfig.html#cfn-emr-cluster-jobflowinstancesconfig-instancegroupconfig-market
	Market *StringExpr `json:"Market,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig-instancegroupconfig.html#cfn-emr-cluster-jobflowinstancesconfig-instancegroupconfig-name
	Name *StringExpr `json:"Name,omitempty"`
}

// EMRClusterInstanceGroupConfigList represents a list of EMRClusterInstanceGroupConfig
type EMRClusterInstanceGroupConfigList []EMRClusterInstanceGroupConfig

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRClusterInstanceGroupConfigList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRClusterInstanceGroupConfig{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRClusterInstanceGroupConfigList{item}
		return nil
	}
	list := []EMRClusterInstanceGroupConfig{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRClusterInstanceGroupConfigList(list)
		return nil
	}
	return err
}

// EMRClusterJobFlowInstancesConfig represents the AWS::EMR::Cluster.JobFlowInstancesConfig CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html
type EMRClusterJobFlowInstancesConfig struct {
	// AdditionalMasterSecurityGroups docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-additionalmastersecuritygroups
	AdditionalMasterSecurityGroups *StringListExpr `json:"AdditionalMasterSecurityGroups,omitempty"`
	// AdditionalSlaveSecurityGroups docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-additionalslavesecuritygroups
	AdditionalSlaveSecurityGroups *StringListExpr `json:"AdditionalSlaveSecurityGroups,omitempty"`
	// CoreInstanceGroup docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-coreinstancegroup
	CoreInstanceGroup *EMRClusterInstanceGroupConfig `json:"CoreInstanceGroup,omitempty"`
	// Ec2KeyName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-ec2keyname
	Ec2KeyName *StringExpr `json:"Ec2KeyName,omitempty"`
	// Ec2SubnetID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-ec2subnetid
	Ec2SubnetID *StringExpr `json:"Ec2SubnetId,omitempty"`
	// EmrManagedMasterSecurityGroup docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-emrmanagedmastersecuritygroup
	EmrManagedMasterSecurityGroup *StringExpr `json:"EmrManagedMasterSecurityGroup,omitempty"`
	// EmrManagedSlaveSecurityGroup docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-emrmanagedslavesecuritygroup
	EmrManagedSlaveSecurityGroup *StringExpr `json:"EmrManagedSlaveSecurityGroup,omitempty"`
	// HadoopVersion docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-hadoopversion
	HadoopVersion *StringExpr `json:"HadoopVersion,omitempty"`
	// MasterInstanceGroup docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-masterinstancegroup
	MasterInstanceGroup *EMRClusterInstanceGroupConfig `json:"MasterInstanceGroup,omitempty" validate:"dive,required"`
	// Placement docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-placement
	Placement *EMRClusterPlacementType `json:"Placement,omitempty"`
	// ServiceAccessSecurityGroup docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-serviceaccesssecuritygroup
	ServiceAccessSecurityGroup *StringExpr `json:"ServiceAccessSecurityGroup,omitempty"`
	// TerminationProtected docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-terminationprotected
	TerminationProtected *BoolExpr `json:"TerminationProtected,omitempty"`
}

// EMRClusterJobFlowInstancesConfigList represents a list of EMRClusterJobFlowInstancesConfig
type EMRClusterJobFlowInstancesConfigList []EMRClusterJobFlowInstancesConfig

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRClusterJobFlowInstancesConfigList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRClusterJobFlowInstancesConfig{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRClusterJobFlowInstancesConfigList{item}
		return nil
	}
	list := []EMRClusterJobFlowInstancesConfig{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRClusterJobFlowInstancesConfigList(list)
		return nil
	}
	return err
}

// EMRClusterMetricDimension represents the AWS::EMR::Cluster.MetricDimension CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-metricdimension.html
type EMRClusterMetricDimension struct {
	// Key docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-metricdimension.html#cfn-elasticmapreduce-cluster-metricdimension-key
	Key *StringExpr `json:"Key,omitempty" validate:"dive,required"`
	// Value docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-metricdimension.html#cfn-elasticmapreduce-cluster-metricdimension-value
	Value *StringExpr `json:"Value,omitempty" validate:"dive,required"`
}

// EMRClusterMetricDimensionList represents a list of EMRClusterMetricDimension
type EMRClusterMetricDimensionList []EMRClusterMetricDimension

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRClusterMetricDimensionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRClusterMetricDimension{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRClusterMetricDimensionList{item}
		return nil
	}
	list := []EMRClusterMetricDimension{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRClusterMetricDimensionList(list)
		return nil
	}
	return err
}

// EMRClusterPlacementType represents the AWS::EMR::Cluster.PlacementType CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig-placementtype.html
type EMRClusterPlacementType struct {
	// AvailabilityZone docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig-placementtype.html#aws-properties-emr-cluster-jobflowinstancesconfig-placementtype
	AvailabilityZone *StringExpr `json:"AvailabilityZone,omitempty" validate:"dive,required"`
}

// EMRClusterPlacementTypeList represents a list of EMRClusterPlacementType
type EMRClusterPlacementTypeList []EMRClusterPlacementType

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRClusterPlacementTypeList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRClusterPlacementType{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRClusterPlacementTypeList{item}
		return nil
	}
	list := []EMRClusterPlacementType{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRClusterPlacementTypeList(list)
		return nil
	}
	return err
}

// EMRClusterScalingAction represents the AWS::EMR::Cluster.ScalingAction CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingaction.html
type EMRClusterScalingAction struct {
	// Market docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingaction.html#cfn-elasticmapreduce-cluster-scalingaction-market
	Market *StringExpr `json:"Market,omitempty"`
	// SimpleScalingPolicyConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingaction.html#cfn-elasticmapreduce-cluster-scalingaction-simplescalingpolicyconfiguration
	SimpleScalingPolicyConfiguration *EMRClusterSimpleScalingPolicyConfiguration `json:"SimpleScalingPolicyConfiguration,omitempty" validate:"dive,required"`
}

// EMRClusterScalingActionList represents a list of EMRClusterScalingAction
type EMRClusterScalingActionList []EMRClusterScalingAction

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRClusterScalingActionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRClusterScalingAction{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRClusterScalingActionList{item}
		return nil
	}
	list := []EMRClusterScalingAction{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRClusterScalingActionList(list)
		return nil
	}
	return err
}

// EMRClusterScalingConstraints represents the AWS::EMR::Cluster.ScalingConstraints CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingconstraints.html
type EMRClusterScalingConstraints struct {
	// MaxCapacity docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingconstraints.html#cfn-elasticmapreduce-cluster-scalingconstraints-maxcapacity
	MaxCapacity *IntegerExpr `json:"MaxCapacity,omitempty" validate:"dive,required"`
	// MinCapacity docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingconstraints.html#cfn-elasticmapreduce-cluster-scalingconstraints-mincapacity
	MinCapacity *IntegerExpr `json:"MinCapacity,omitempty" validate:"dive,required"`
}

// EMRClusterScalingConstraintsList represents a list of EMRClusterScalingConstraints
type EMRClusterScalingConstraintsList []EMRClusterScalingConstraints

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRClusterScalingConstraintsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRClusterScalingConstraints{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRClusterScalingConstraintsList{item}
		return nil
	}
	list := []EMRClusterScalingConstraints{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRClusterScalingConstraintsList(list)
		return nil
	}
	return err
}

// EMRClusterScalingRule represents the AWS::EMR::Cluster.ScalingRule CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingrule.html
type EMRClusterScalingRule struct {
	// Action docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingrule.html#cfn-elasticmapreduce-cluster-scalingrule-action
	Action *EMRClusterScalingAction `json:"Action,omitempty" validate:"dive,required"`
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingrule.html#cfn-elasticmapreduce-cluster-scalingrule-description
	Description *StringExpr `json:"Description,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingrule.html#cfn-elasticmapreduce-cluster-scalingrule-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// Trigger docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingrule.html#cfn-elasticmapreduce-cluster-scalingrule-trigger
	Trigger *EMRClusterScalingTrigger `json:"Trigger,omitempty" validate:"dive,required"`
}

// EMRClusterScalingRuleList represents a list of EMRClusterScalingRule
type EMRClusterScalingRuleList []EMRClusterScalingRule

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRClusterScalingRuleList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRClusterScalingRule{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRClusterScalingRuleList{item}
		return nil
	}
	list := []EMRClusterScalingRule{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRClusterScalingRuleList(list)
		return nil
	}
	return err
}

// EMRClusterScalingTrigger represents the AWS::EMR::Cluster.ScalingTrigger CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingtrigger.html
type EMRClusterScalingTrigger struct {
	// CloudWatchAlarmDefinition docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingtrigger.html#cfn-elasticmapreduce-cluster-scalingtrigger-cloudwatchalarmdefinition
	CloudWatchAlarmDefinition *EMRClusterCloudWatchAlarmDefinition `json:"CloudWatchAlarmDefinition,omitempty" validate:"dive,required"`
}

// EMRClusterScalingTriggerList represents a list of EMRClusterScalingTrigger
type EMRClusterScalingTriggerList []EMRClusterScalingTrigger

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRClusterScalingTriggerList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRClusterScalingTrigger{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRClusterScalingTriggerList{item}
		return nil
	}
	list := []EMRClusterScalingTrigger{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRClusterScalingTriggerList(list)
		return nil
	}
	return err
}

// EMRClusterScriptBootstrapActionConfig represents the AWS::EMR::Cluster.ScriptBootstrapActionConfig CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-bootstrapactionconfig-scriptbootstrapactionconfig.html
type EMRClusterScriptBootstrapActionConfig struct {
	// Args docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-bootstrapactionconfig-scriptbootstrapactionconfig.html#cfn-emr-cluster-bootstrapactionconfig-scriptbootstrapaction-args
	Args *StringListExpr `json:"Args,omitempty"`
	// Path docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-bootstrapactionconfig-scriptbootstrapactionconfig.html#cfn-emr-cluster-bootstrapactionconfig-scriptbootstrapaction-path
	Path *StringExpr `json:"Path,omitempty" validate:"dive,required"`
}

// EMRClusterScriptBootstrapActionConfigList represents a list of EMRClusterScriptBootstrapActionConfig
type EMRClusterScriptBootstrapActionConfigList []EMRClusterScriptBootstrapActionConfig

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRClusterScriptBootstrapActionConfigList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRClusterScriptBootstrapActionConfig{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRClusterScriptBootstrapActionConfigList{item}
		return nil
	}
	list := []EMRClusterScriptBootstrapActionConfig{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRClusterScriptBootstrapActionConfigList(list)
		return nil
	}
	return err
}

// EMRClusterSimpleScalingPolicyConfiguration represents the AWS::EMR::Cluster.SimpleScalingPolicyConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-simplescalingpolicyconfiguration.html
type EMRClusterSimpleScalingPolicyConfiguration struct {
	// AdjustmentType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-simplescalingpolicyconfiguration.html#cfn-elasticmapreduce-cluster-simplescalingpolicyconfiguration-adjustmenttype
	AdjustmentType *StringExpr `json:"AdjustmentType,omitempty"`
	// CoolDown docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-simplescalingpolicyconfiguration.html#cfn-elasticmapreduce-cluster-simplescalingpolicyconfiguration-cooldown
	CoolDown *IntegerExpr `json:"CoolDown,omitempty"`
	// ScalingAdjustment docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-simplescalingpolicyconfiguration.html#cfn-elasticmapreduce-cluster-simplescalingpolicyconfiguration-scalingadjustment
	ScalingAdjustment *IntegerExpr `json:"ScalingAdjustment,omitempty" validate:"dive,required"`
}

// EMRClusterSimpleScalingPolicyConfigurationList represents a list of EMRClusterSimpleScalingPolicyConfiguration
type EMRClusterSimpleScalingPolicyConfigurationList []EMRClusterSimpleScalingPolicyConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRClusterSimpleScalingPolicyConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRClusterSimpleScalingPolicyConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRClusterSimpleScalingPolicyConfigurationList{item}
		return nil
	}
	list := []EMRClusterSimpleScalingPolicyConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRClusterSimpleScalingPolicyConfigurationList(list)
		return nil
	}
	return err
}

// EMRClusterVolumeSpecification represents the AWS::EMR::Cluster.VolumeSpecification CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification.html
type EMRClusterVolumeSpecification struct {
	// Iops docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification-iops
	Iops *IntegerExpr `json:"Iops,omitempty"`
	// SizeInGB docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification-sizeingb
	SizeInGB *IntegerExpr `json:"SizeInGB,omitempty" validate:"dive,required"`
	// VolumeType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification-volumetype
	VolumeType *StringExpr `json:"VolumeType,omitempty" validate:"dive,required"`
}

// EMRClusterVolumeSpecificationList represents a list of EMRClusterVolumeSpecification
type EMRClusterVolumeSpecificationList []EMRClusterVolumeSpecification

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRClusterVolumeSpecificationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRClusterVolumeSpecification{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRClusterVolumeSpecificationList{item}
		return nil
	}
	list := []EMRClusterVolumeSpecification{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRClusterVolumeSpecificationList(list)
		return nil
	}
	return err
}

// EMRInstanceGroupConfigAutoScalingPolicy represents the AWS::EMR::InstanceGroupConfig.AutoScalingPolicy CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-autoscalingpolicy.html
type EMRInstanceGroupConfigAutoScalingPolicy struct {
	// Constraints docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-autoscalingpolicy.html#cfn-elasticmapreduce-instancegroupconfig-autoscalingpolicy-constraints
	Constraints *EMRInstanceGroupConfigScalingConstraints `json:"Constraints,omitempty" validate:"dive,required"`
	// Rules docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-autoscalingpolicy.html#cfn-elasticmapreduce-instancegroupconfig-autoscalingpolicy-rules
	Rules *EMRInstanceGroupConfigScalingRuleList `json:"Rules,omitempty" validate:"dive,required"`
}

// EMRInstanceGroupConfigAutoScalingPolicyList represents a list of EMRInstanceGroupConfigAutoScalingPolicy
type EMRInstanceGroupConfigAutoScalingPolicyList []EMRInstanceGroupConfigAutoScalingPolicy

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRInstanceGroupConfigAutoScalingPolicyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRInstanceGroupConfigAutoScalingPolicy{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRInstanceGroupConfigAutoScalingPolicyList{item}
		return nil
	}
	list := []EMRInstanceGroupConfigAutoScalingPolicy{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRInstanceGroupConfigAutoScalingPolicyList(list)
		return nil
	}
	return err
}

// EMRInstanceGroupConfigCloudWatchAlarmDefinition represents the AWS::EMR::InstanceGroupConfig.CloudWatchAlarmDefinition CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html
type EMRInstanceGroupConfigCloudWatchAlarmDefinition struct {
	// ComparisonOperator docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition-comparisonoperator
	ComparisonOperator *StringExpr `json:"ComparisonOperator,omitempty" validate:"dive,required"`
	// Dimensions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition-dimensions
	Dimensions *EMRInstanceGroupConfigMetricDimensionList `json:"Dimensions,omitempty"`
	// EvaluationPeriods docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition-evaluationperiods
	EvaluationPeriods *IntegerExpr `json:"EvaluationPeriods,omitempty"`
	// MetricName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition-metricname
	MetricName *StringExpr `json:"MetricName,omitempty" validate:"dive,required"`
	// Namespace docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition-namespace
	Namespace *StringExpr `json:"Namespace,omitempty"`
	// Period docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition-period
	Period *IntegerExpr `json:"Period,omitempty" validate:"dive,required"`
	// Statistic docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition-statistic
	Statistic *StringExpr `json:"Statistic,omitempty"`
	// Threshold docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition-threshold
	Threshold *IntegerExpr `json:"Threshold,omitempty" validate:"dive,required"`
	// Unit docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition-unit
	Unit *StringExpr `json:"Unit,omitempty"`
}

// EMRInstanceGroupConfigCloudWatchAlarmDefinitionList represents a list of EMRInstanceGroupConfigCloudWatchAlarmDefinition
type EMRInstanceGroupConfigCloudWatchAlarmDefinitionList []EMRInstanceGroupConfigCloudWatchAlarmDefinition

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRInstanceGroupConfigCloudWatchAlarmDefinitionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRInstanceGroupConfigCloudWatchAlarmDefinition{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRInstanceGroupConfigCloudWatchAlarmDefinitionList{item}
		return nil
	}
	list := []EMRInstanceGroupConfigCloudWatchAlarmDefinition{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRInstanceGroupConfigCloudWatchAlarmDefinitionList(list)
		return nil
	}
	return err
}

// EMRInstanceGroupConfigConfiguration represents the AWS::EMR::InstanceGroupConfig.Configuration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-configuration.html
type EMRInstanceGroupConfigConfiguration struct {
	// Classification docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-configuration.html#cfn-emr-cluster-configuration-classification
	Classification *StringExpr `json:"Classification,omitempty"`
	// ConfigurationProperties docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-configuration.html#cfn-emr-cluster-configuration-configurationproperties
	ConfigurationProperties interface{} `json:"ConfigurationProperties,omitempty"`
	// Configurations docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-configuration.html#cfn-emr-cluster-configuration-configurations
	Configurations *EMRInstanceGroupConfigConfigurationList `json:"Configurations,omitempty"`
}

// EMRInstanceGroupConfigConfigurationList represents a list of EMRInstanceGroupConfigConfiguration
type EMRInstanceGroupConfigConfigurationList []EMRInstanceGroupConfigConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRInstanceGroupConfigConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRInstanceGroupConfigConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRInstanceGroupConfigConfigurationList{item}
		return nil
	}
	list := []EMRInstanceGroupConfigConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRInstanceGroupConfigConfigurationList(list)
		return nil
	}
	return err
}

// EMRInstanceGroupConfigEbsBlockDeviceConfig represents the AWS::EMR::InstanceGroupConfig.EbsBlockDeviceConfig CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig.html
type EMRInstanceGroupConfigEbsBlockDeviceConfig struct {
	// VolumeSpecification docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification
	VolumeSpecification *EMRInstanceGroupConfigVolumeSpecification `json:"VolumeSpecification,omitempty" validate:"dive,required"`
	// VolumesPerInstance docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfig-volumesperinstance
	VolumesPerInstance *IntegerExpr `json:"VolumesPerInstance,omitempty"`
}

// EMRInstanceGroupConfigEbsBlockDeviceConfigList represents a list of EMRInstanceGroupConfigEbsBlockDeviceConfig
type EMRInstanceGroupConfigEbsBlockDeviceConfigList []EMRInstanceGroupConfigEbsBlockDeviceConfig

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRInstanceGroupConfigEbsBlockDeviceConfigList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRInstanceGroupConfigEbsBlockDeviceConfig{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRInstanceGroupConfigEbsBlockDeviceConfigList{item}
		return nil
	}
	list := []EMRInstanceGroupConfigEbsBlockDeviceConfig{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRInstanceGroupConfigEbsBlockDeviceConfigList(list)
		return nil
	}
	return err
}

// EMRInstanceGroupConfigEbsConfiguration represents the AWS::EMR::InstanceGroupConfig.EbsConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration.html
type EMRInstanceGroupConfigEbsConfiguration struct {
	// EbsBlockDeviceConfigs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfigs
	EbsBlockDeviceConfigs *EMRInstanceGroupConfigEbsBlockDeviceConfigList `json:"EbsBlockDeviceConfigs,omitempty"`
	// EbsOptimized docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration.html#cfn-emr-ebsconfiguration-ebsoptimized
	EbsOptimized *BoolExpr `json:"EbsOptimized,omitempty"`
}

// EMRInstanceGroupConfigEbsConfigurationList represents a list of EMRInstanceGroupConfigEbsConfiguration
type EMRInstanceGroupConfigEbsConfigurationList []EMRInstanceGroupConfigEbsConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRInstanceGroupConfigEbsConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRInstanceGroupConfigEbsConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRInstanceGroupConfigEbsConfigurationList{item}
		return nil
	}
	list := []EMRInstanceGroupConfigEbsConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRInstanceGroupConfigEbsConfigurationList(list)
		return nil
	}
	return err
}

// EMRInstanceGroupConfigMetricDimension represents the AWS::EMR::InstanceGroupConfig.MetricDimension CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-metricdimension.html
type EMRInstanceGroupConfigMetricDimension struct {
	// Key docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-metricdimension.html#cfn-elasticmapreduce-instancegroupconfig-metricdimension-key
	Key *StringExpr `json:"Key,omitempty" validate:"dive,required"`
	// Value docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-metricdimension.html#cfn-elasticmapreduce-instancegroupconfig-metricdimension-value
	Value *StringExpr `json:"Value,omitempty" validate:"dive,required"`
}

// EMRInstanceGroupConfigMetricDimensionList represents a list of EMRInstanceGroupConfigMetricDimension
type EMRInstanceGroupConfigMetricDimensionList []EMRInstanceGroupConfigMetricDimension

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRInstanceGroupConfigMetricDimensionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRInstanceGroupConfigMetricDimension{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRInstanceGroupConfigMetricDimensionList{item}
		return nil
	}
	list := []EMRInstanceGroupConfigMetricDimension{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRInstanceGroupConfigMetricDimensionList(list)
		return nil
	}
	return err
}

// EMRInstanceGroupConfigScalingAction represents the AWS::EMR::InstanceGroupConfig.ScalingAction CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingaction.html
type EMRInstanceGroupConfigScalingAction struct {
	// Market docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingaction.html#cfn-elasticmapreduce-instancegroupconfig-scalingaction-market
	Market *StringExpr `json:"Market,omitempty"`
	// SimpleScalingPolicyConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingaction.html#cfn-elasticmapreduce-instancegroupconfig-scalingaction-simplescalingpolicyconfiguration
	SimpleScalingPolicyConfiguration *EMRInstanceGroupConfigSimpleScalingPolicyConfiguration `json:"SimpleScalingPolicyConfiguration,omitempty" validate:"dive,required"`
}

// EMRInstanceGroupConfigScalingActionList represents a list of EMRInstanceGroupConfigScalingAction
type EMRInstanceGroupConfigScalingActionList []EMRInstanceGroupConfigScalingAction

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRInstanceGroupConfigScalingActionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRInstanceGroupConfigScalingAction{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRInstanceGroupConfigScalingActionList{item}
		return nil
	}
	list := []EMRInstanceGroupConfigScalingAction{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRInstanceGroupConfigScalingActionList(list)
		return nil
	}
	return err
}

// EMRInstanceGroupConfigScalingConstraints represents the AWS::EMR::InstanceGroupConfig.ScalingConstraints CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingconstraints.html
type EMRInstanceGroupConfigScalingConstraints struct {
	// MaxCapacity docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingconstraints.html#cfn-elasticmapreduce-instancegroupconfig-scalingconstraints-maxcapacity
	MaxCapacity *IntegerExpr `json:"MaxCapacity,omitempty" validate:"dive,required"`
	// MinCapacity docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingconstraints.html#cfn-elasticmapreduce-instancegroupconfig-scalingconstraints-mincapacity
	MinCapacity *IntegerExpr `json:"MinCapacity,omitempty" validate:"dive,required"`
}

// EMRInstanceGroupConfigScalingConstraintsList represents a list of EMRInstanceGroupConfigScalingConstraints
type EMRInstanceGroupConfigScalingConstraintsList []EMRInstanceGroupConfigScalingConstraints

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRInstanceGroupConfigScalingConstraintsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRInstanceGroupConfigScalingConstraints{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRInstanceGroupConfigScalingConstraintsList{item}
		return nil
	}
	list := []EMRInstanceGroupConfigScalingConstraints{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRInstanceGroupConfigScalingConstraintsList(list)
		return nil
	}
	return err
}

// EMRInstanceGroupConfigScalingRule represents the AWS::EMR::InstanceGroupConfig.ScalingRule CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingrule.html
type EMRInstanceGroupConfigScalingRule struct {
	// Action docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingrule.html#cfn-elasticmapreduce-instancegroupconfig-scalingrule-action
	Action *EMRInstanceGroupConfigScalingAction `json:"Action,omitempty" validate:"dive,required"`
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingrule.html#cfn-elasticmapreduce-instancegroupconfig-scalingrule-description
	Description *StringExpr `json:"Description,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingrule.html#cfn-elasticmapreduce-instancegroupconfig-scalingrule-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// Trigger docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingrule.html#cfn-elasticmapreduce-instancegroupconfig-scalingrule-trigger
	Trigger *EMRInstanceGroupConfigScalingTrigger `json:"Trigger,omitempty" validate:"dive,required"`
}

// EMRInstanceGroupConfigScalingRuleList represents a list of EMRInstanceGroupConfigScalingRule
type EMRInstanceGroupConfigScalingRuleList []EMRInstanceGroupConfigScalingRule

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRInstanceGroupConfigScalingRuleList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRInstanceGroupConfigScalingRule{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRInstanceGroupConfigScalingRuleList{item}
		return nil
	}
	list := []EMRInstanceGroupConfigScalingRule{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRInstanceGroupConfigScalingRuleList(list)
		return nil
	}
	return err
}

// EMRInstanceGroupConfigScalingTrigger represents the AWS::EMR::InstanceGroupConfig.ScalingTrigger CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingtrigger.html
type EMRInstanceGroupConfigScalingTrigger struct {
	// CloudWatchAlarmDefinition docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingtrigger.html#cfn-elasticmapreduce-instancegroupconfig-scalingtrigger-cloudwatchalarmdefinition
	CloudWatchAlarmDefinition *EMRInstanceGroupConfigCloudWatchAlarmDefinition `json:"CloudWatchAlarmDefinition,omitempty" validate:"dive,required"`
}

// EMRInstanceGroupConfigScalingTriggerList represents a list of EMRInstanceGroupConfigScalingTrigger
type EMRInstanceGroupConfigScalingTriggerList []EMRInstanceGroupConfigScalingTrigger

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRInstanceGroupConfigScalingTriggerList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRInstanceGroupConfigScalingTrigger{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRInstanceGroupConfigScalingTriggerList{item}
		return nil
	}
	list := []EMRInstanceGroupConfigScalingTrigger{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRInstanceGroupConfigScalingTriggerList(list)
		return nil
	}
	return err
}

// EMRInstanceGroupConfigSimpleScalingPolicyConfiguration represents the AWS::EMR::InstanceGroupConfig.SimpleScalingPolicyConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-simplescalingpolicyconfiguration.html
type EMRInstanceGroupConfigSimpleScalingPolicyConfiguration struct {
	// AdjustmentType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-simplescalingpolicyconfiguration.html#cfn-elasticmapreduce-instancegroupconfig-simplescalingpolicyconfiguration-adjustmenttype
	AdjustmentType *StringExpr `json:"AdjustmentType,omitempty"`
	// CoolDown docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-simplescalingpolicyconfiguration.html#cfn-elasticmapreduce-instancegroupconfig-simplescalingpolicyconfiguration-cooldown
	CoolDown *IntegerExpr `json:"CoolDown,omitempty"`
	// ScalingAdjustment docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-simplescalingpolicyconfiguration.html#cfn-elasticmapreduce-instancegroupconfig-simplescalingpolicyconfiguration-scalingadjustment
	ScalingAdjustment *IntegerExpr `json:"ScalingAdjustment,omitempty" validate:"dive,required"`
}

// EMRInstanceGroupConfigSimpleScalingPolicyConfigurationList represents a list of EMRInstanceGroupConfigSimpleScalingPolicyConfiguration
type EMRInstanceGroupConfigSimpleScalingPolicyConfigurationList []EMRInstanceGroupConfigSimpleScalingPolicyConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRInstanceGroupConfigSimpleScalingPolicyConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRInstanceGroupConfigSimpleScalingPolicyConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRInstanceGroupConfigSimpleScalingPolicyConfigurationList{item}
		return nil
	}
	list := []EMRInstanceGroupConfigSimpleScalingPolicyConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRInstanceGroupConfigSimpleScalingPolicyConfigurationList(list)
		return nil
	}
	return err
}

// EMRInstanceGroupConfigVolumeSpecification represents the AWS::EMR::InstanceGroupConfig.VolumeSpecification CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification.html
type EMRInstanceGroupConfigVolumeSpecification struct {
	// Iops docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification-iops
	Iops *IntegerExpr `json:"Iops,omitempty"`
	// SizeInGB docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification-sizeingb
	SizeInGB *IntegerExpr `json:"SizeInGB,omitempty" validate:"dive,required"`
	// VolumeType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification-volumetype
	VolumeType *StringExpr `json:"VolumeType,omitempty" validate:"dive,required"`
}

// EMRInstanceGroupConfigVolumeSpecificationList represents a list of EMRInstanceGroupConfigVolumeSpecification
type EMRInstanceGroupConfigVolumeSpecificationList []EMRInstanceGroupConfigVolumeSpecification

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRInstanceGroupConfigVolumeSpecificationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRInstanceGroupConfigVolumeSpecification{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRInstanceGroupConfigVolumeSpecificationList{item}
		return nil
	}
	list := []EMRInstanceGroupConfigVolumeSpecification{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRInstanceGroupConfigVolumeSpecificationList(list)
		return nil
	}
	return err
}

// EMRStepHadoopJarStepConfig represents the AWS::EMR::Step.HadoopJarStepConfig CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-step-hadoopjarstepconfig.html
type EMRStepHadoopJarStepConfig struct {
	// Args docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-step-hadoopjarstepconfig.html#cfn-elasticmapreduce-step-hadoopjarstepconfig-args
	Args *StringListExpr `json:"Args,omitempty"`
	// Jar docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-step-hadoopjarstepconfig.html#cfn-elasticmapreduce-step-hadoopjarstepconfig-jar
	Jar *StringExpr `json:"Jar,omitempty" validate:"dive,required"`
	// MainClass docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-step-hadoopjarstepconfig.html#cfn-elasticmapreduce-step-hadoopjarstepconfig-mainclass
	MainClass *StringExpr `json:"MainClass,omitempty"`
	// StepProperties docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-step-hadoopjarstepconfig.html#cfn-elasticmapreduce-step-hadoopjarstepconfig-stepproperties
	StepProperties *EMRStepKeyValueList `json:"StepProperties,omitempty"`
}

// EMRStepHadoopJarStepConfigList represents a list of EMRStepHadoopJarStepConfig
type EMRStepHadoopJarStepConfigList []EMRStepHadoopJarStepConfig

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRStepHadoopJarStepConfigList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRStepHadoopJarStepConfig{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRStepHadoopJarStepConfigList{item}
		return nil
	}
	list := []EMRStepHadoopJarStepConfig{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRStepHadoopJarStepConfigList(list)
		return nil
	}
	return err
}

// EMRStepKeyValue represents the AWS::EMR::Step.KeyValue CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-step-keyvalue.html
type EMRStepKeyValue struct {
	// Key docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-step-keyvalue.html#cfn-elasticmapreduce-step-keyvalue-key
	Key *StringExpr `json:"Key,omitempty"`
	// Value docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-step-keyvalue.html#cfn-elasticmapreduce-step-keyvalue-value
	Value *StringExpr `json:"Value,omitempty"`
}

// EMRStepKeyValueList represents a list of EMRStepKeyValue
type EMRStepKeyValueList []EMRStepKeyValue

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EMRStepKeyValueList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EMRStepKeyValue{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EMRStepKeyValueList{item}
		return nil
	}
	list := []EMRStepKeyValue{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EMRStepKeyValueList(list)
		return nil
	}
	return err
}

// ElastiCacheReplicationGroupNodeGroupConfiguration represents the AWS::ElastiCache::ReplicationGroup.NodeGroupConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-replicationgroup-nodegroupconfiguration.html
type ElastiCacheReplicationGroupNodeGroupConfiguration struct {
	// PrimaryAvailabilityZone docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-replicationgroup-nodegroupconfiguration.html#cfn-elasticache-replicationgroup-nodegroupconfiguration-primaryavailabilityzone
	PrimaryAvailabilityZone *StringExpr `json:"PrimaryAvailabilityZone,omitempty"`
	// ReplicaAvailabilityZones docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-replicationgroup-nodegroupconfiguration.html#cfn-elasticache-replicationgroup-nodegroupconfiguration-replicaavailabilityzones
	ReplicaAvailabilityZones *StringListExpr `json:"ReplicaAvailabilityZones,omitempty"`
	// ReplicaCount docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-replicationgroup-nodegroupconfiguration.html#cfn-elasticache-replicationgroup-nodegroupconfiguration-replicacount
	ReplicaCount *IntegerExpr `json:"ReplicaCount,omitempty"`
	// Slots docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-replicationgroup-nodegroupconfiguration.html#cfn-elasticache-replicationgroup-nodegroupconfiguration-slots
	Slots *StringExpr `json:"Slots,omitempty"`
}

// ElastiCacheReplicationGroupNodeGroupConfigurationList represents a list of ElastiCacheReplicationGroupNodeGroupConfiguration
type ElastiCacheReplicationGroupNodeGroupConfigurationList []ElastiCacheReplicationGroupNodeGroupConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElastiCacheReplicationGroupNodeGroupConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElastiCacheReplicationGroupNodeGroupConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElastiCacheReplicationGroupNodeGroupConfigurationList{item}
		return nil
	}
	list := []ElastiCacheReplicationGroupNodeGroupConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElastiCacheReplicationGroupNodeGroupConfigurationList(list)
		return nil
	}
	return err
}

// ElasticBeanstalkApplicationVersionSourceBundle represents the AWS::ElasticBeanstalk::ApplicationVersion.SourceBundle CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-sourcebundle.html
type ElasticBeanstalkApplicationVersionSourceBundle struct {
	// S3Bucket docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-sourcebundle.html#cfn-beanstalk-sourcebundle-s3bucket
	S3Bucket *StringExpr `json:"S3Bucket,omitempty" validate:"dive,required"`
	// S3Key docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-sourcebundle.html#cfn-beanstalk-sourcebundle-s3key
	S3Key *StringExpr `json:"S3Key,omitempty" validate:"dive,required"`
}

// ElasticBeanstalkApplicationVersionSourceBundleList represents a list of ElasticBeanstalkApplicationVersionSourceBundle
type ElasticBeanstalkApplicationVersionSourceBundleList []ElasticBeanstalkApplicationVersionSourceBundle

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticBeanstalkApplicationVersionSourceBundleList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticBeanstalkApplicationVersionSourceBundle{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticBeanstalkApplicationVersionSourceBundleList{item}
		return nil
	}
	list := []ElasticBeanstalkApplicationVersionSourceBundle{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticBeanstalkApplicationVersionSourceBundleList(list)
		return nil
	}
	return err
}

// ElasticBeanstalkConfigurationTemplateConfigurationOptionSetting represents the AWS::ElasticBeanstalk::ConfigurationTemplate.ConfigurationOptionSetting CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-option-settings.html
type ElasticBeanstalkConfigurationTemplateConfigurationOptionSetting struct {
	// Namespace docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-option-settings.html#cfn-beanstalk-optionsettings-namespace
	Namespace *StringExpr `json:"Namespace,omitempty" validate:"dive,required"`
	// OptionName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-option-settings.html#cfn-beanstalk-optionsettings-optionname
	OptionName *StringExpr `json:"OptionName,omitempty" validate:"dive,required"`
	// Value docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-option-settings.html#cfn-beanstalk-optionsettings-value
	Value *StringExpr `json:"Value,omitempty" validate:"dive,required"`
}

// ElasticBeanstalkConfigurationTemplateConfigurationOptionSettingList represents a list of ElasticBeanstalkConfigurationTemplateConfigurationOptionSetting
type ElasticBeanstalkConfigurationTemplateConfigurationOptionSettingList []ElasticBeanstalkConfigurationTemplateConfigurationOptionSetting

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticBeanstalkConfigurationTemplateConfigurationOptionSettingList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticBeanstalkConfigurationTemplateConfigurationOptionSetting{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticBeanstalkConfigurationTemplateConfigurationOptionSettingList{item}
		return nil
	}
	list := []ElasticBeanstalkConfigurationTemplateConfigurationOptionSetting{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticBeanstalkConfigurationTemplateConfigurationOptionSettingList(list)
		return nil
	}
	return err
}

// ElasticBeanstalkConfigurationTemplateSourceConfiguration represents the AWS::ElasticBeanstalk::ConfigurationTemplate.SourceConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-configurationtemplate-sourceconfiguration.html
type ElasticBeanstalkConfigurationTemplateSourceConfiguration struct {
	// ApplicationName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-configurationtemplate-sourceconfiguration.html#cfn-beanstalk-configurationtemplate-sourceconfiguration-applicationname
	ApplicationName *StringExpr `json:"ApplicationName,omitempty" validate:"dive,required"`
	// TemplateName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-configurationtemplate-sourceconfiguration.html#cfn-beanstalk-configurationtemplate-sourceconfiguration-templatename
	TemplateName *StringExpr `json:"TemplateName,omitempty" validate:"dive,required"`
}

// ElasticBeanstalkConfigurationTemplateSourceConfigurationList represents a list of ElasticBeanstalkConfigurationTemplateSourceConfiguration
type ElasticBeanstalkConfigurationTemplateSourceConfigurationList []ElasticBeanstalkConfigurationTemplateSourceConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticBeanstalkConfigurationTemplateSourceConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticBeanstalkConfigurationTemplateSourceConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticBeanstalkConfigurationTemplateSourceConfigurationList{item}
		return nil
	}
	list := []ElasticBeanstalkConfigurationTemplateSourceConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticBeanstalkConfigurationTemplateSourceConfigurationList(list)
		return nil
	}
	return err
}

// ElasticBeanstalkEnvironmentOptionSettings represents the AWS::ElasticBeanstalk::Environment.OptionSettings CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-option-settings.html
type ElasticBeanstalkEnvironmentOptionSettings struct {
	// Namespace docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-option-settings.html#cfn-beanstalk-optionsettings-namespace
	Namespace *StringExpr `json:"Namespace,omitempty" validate:"dive,required"`
	// OptionName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-option-settings.html#cfn-beanstalk-optionsettings-optionname
	OptionName *StringExpr `json:"OptionName,omitempty" validate:"dive,required"`
	// Value docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-option-settings.html#cfn-beanstalk-optionsettings-value
	Value *StringExpr `json:"Value,omitempty" validate:"dive,required"`
}

// ElasticBeanstalkEnvironmentOptionSettingsList represents a list of ElasticBeanstalkEnvironmentOptionSettings
type ElasticBeanstalkEnvironmentOptionSettingsList []ElasticBeanstalkEnvironmentOptionSettings

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticBeanstalkEnvironmentOptionSettingsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticBeanstalkEnvironmentOptionSettings{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticBeanstalkEnvironmentOptionSettingsList{item}
		return nil
	}
	list := []ElasticBeanstalkEnvironmentOptionSettings{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticBeanstalkEnvironmentOptionSettingsList(list)
		return nil
	}
	return err
}

// ElasticBeanstalkEnvironmentTier represents the AWS::ElasticBeanstalk::Environment.Tier CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment-tier.html
type ElasticBeanstalkEnvironmentTier struct {
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment-tier.html#cfn-beanstalk-env-tier-name
	Name *StringExpr `json:"Name,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment-tier.html#cfn-beanstalk-env-tier-type
	Type *StringExpr `json:"Type,omitempty"`
	// Version docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment-tier.html#cfn-beanstalk-env-tier-version
	Version *StringExpr `json:"Version,omitempty"`
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

// ElasticLoadBalancingLoadBalancerAccessLoggingPolicy represents the AWS::ElasticLoadBalancing::LoadBalancer.AccessLoggingPolicy CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-accessloggingpolicy.html
type ElasticLoadBalancingLoadBalancerAccessLoggingPolicy struct {
	// EmitInterval docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-accessloggingpolicy.html#cfn-elb-accessloggingpolicy-emitinterval
	EmitInterval *IntegerExpr `json:"EmitInterval,omitempty"`
	// Enabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-accessloggingpolicy.html#cfn-elb-accessloggingpolicy-enabled
	Enabled *BoolExpr `json:"Enabled,omitempty" validate:"dive,required"`
	// S3BucketName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-accessloggingpolicy.html#cfn-elb-accessloggingpolicy-s3bucketname
	S3BucketName *StringExpr `json:"S3BucketName,omitempty" validate:"dive,required"`
	// S3BucketPrefix docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-accessloggingpolicy.html#cfn-elb-accessloggingpolicy-s3bucketprefix
	S3BucketPrefix *StringExpr `json:"S3BucketPrefix,omitempty"`
}

// ElasticLoadBalancingLoadBalancerAccessLoggingPolicyList represents a list of ElasticLoadBalancingLoadBalancerAccessLoggingPolicy
type ElasticLoadBalancingLoadBalancerAccessLoggingPolicyList []ElasticLoadBalancingLoadBalancerAccessLoggingPolicy

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticLoadBalancingLoadBalancerAccessLoggingPolicyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticLoadBalancingLoadBalancerAccessLoggingPolicy{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticLoadBalancingLoadBalancerAccessLoggingPolicyList{item}
		return nil
	}
	list := []ElasticLoadBalancingLoadBalancerAccessLoggingPolicy{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticLoadBalancingLoadBalancerAccessLoggingPolicyList(list)
		return nil
	}
	return err
}

// ElasticLoadBalancingLoadBalancerAppCookieStickinessPolicy represents the AWS::ElasticLoadBalancing::LoadBalancer.AppCookieStickinessPolicy CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-AppCookieStickinessPolicy.html
type ElasticLoadBalancingLoadBalancerAppCookieStickinessPolicy struct {
	// CookieName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-AppCookieStickinessPolicy.html#cfn-elb-appcookiestickinesspolicy-cookiename
	CookieName *StringExpr `json:"CookieName,omitempty" validate:"dive,required"`
	// PolicyName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-AppCookieStickinessPolicy.html#cfn-elb-appcookiestickinesspolicy-policyname
	PolicyName *StringExpr `json:"PolicyName,omitempty" validate:"dive,required"`
}

// ElasticLoadBalancingLoadBalancerAppCookieStickinessPolicyList represents a list of ElasticLoadBalancingLoadBalancerAppCookieStickinessPolicy
type ElasticLoadBalancingLoadBalancerAppCookieStickinessPolicyList []ElasticLoadBalancingLoadBalancerAppCookieStickinessPolicy

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticLoadBalancingLoadBalancerAppCookieStickinessPolicyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticLoadBalancingLoadBalancerAppCookieStickinessPolicy{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticLoadBalancingLoadBalancerAppCookieStickinessPolicyList{item}
		return nil
	}
	list := []ElasticLoadBalancingLoadBalancerAppCookieStickinessPolicy{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticLoadBalancingLoadBalancerAppCookieStickinessPolicyList(list)
		return nil
	}
	return err
}

// ElasticLoadBalancingLoadBalancerConnectionDrainingPolicy represents the AWS::ElasticLoadBalancing::LoadBalancer.ConnectionDrainingPolicy CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-connectiondrainingpolicy.html
type ElasticLoadBalancingLoadBalancerConnectionDrainingPolicy struct {
	// Enabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-connectiondrainingpolicy.html#cfn-elb-connectiondrainingpolicy-enabled
	Enabled *BoolExpr `json:"Enabled,omitempty" validate:"dive,required"`
	// Timeout docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-connectiondrainingpolicy.html#cfn-elb-connectiondrainingpolicy-timeout
	Timeout *IntegerExpr `json:"Timeout,omitempty"`
}

// ElasticLoadBalancingLoadBalancerConnectionDrainingPolicyList represents a list of ElasticLoadBalancingLoadBalancerConnectionDrainingPolicy
type ElasticLoadBalancingLoadBalancerConnectionDrainingPolicyList []ElasticLoadBalancingLoadBalancerConnectionDrainingPolicy

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticLoadBalancingLoadBalancerConnectionDrainingPolicyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticLoadBalancingLoadBalancerConnectionDrainingPolicy{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticLoadBalancingLoadBalancerConnectionDrainingPolicyList{item}
		return nil
	}
	list := []ElasticLoadBalancingLoadBalancerConnectionDrainingPolicy{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticLoadBalancingLoadBalancerConnectionDrainingPolicyList(list)
		return nil
	}
	return err
}

// ElasticLoadBalancingLoadBalancerConnectionSettings represents the AWS::ElasticLoadBalancing::LoadBalancer.ConnectionSettings CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-connectionsettings.html
type ElasticLoadBalancingLoadBalancerConnectionSettings struct {
	// IDleTimeout docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-connectionsettings.html#cfn-elb-connectionsettings-idletimeout
	IDleTimeout *IntegerExpr `json:"IdleTimeout,omitempty" validate:"dive,required"`
}

// ElasticLoadBalancingLoadBalancerConnectionSettingsList represents a list of ElasticLoadBalancingLoadBalancerConnectionSettings
type ElasticLoadBalancingLoadBalancerConnectionSettingsList []ElasticLoadBalancingLoadBalancerConnectionSettings

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticLoadBalancingLoadBalancerConnectionSettingsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticLoadBalancingLoadBalancerConnectionSettings{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticLoadBalancingLoadBalancerConnectionSettingsList{item}
		return nil
	}
	list := []ElasticLoadBalancingLoadBalancerConnectionSettings{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticLoadBalancingLoadBalancerConnectionSettingsList(list)
		return nil
	}
	return err
}

// ElasticLoadBalancingLoadBalancerHealthCheck represents the AWS::ElasticLoadBalancing::LoadBalancer.HealthCheck CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-health-check.html
type ElasticLoadBalancingLoadBalancerHealthCheck struct {
	// HealthyThreshold docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-health-check.html#cfn-elb-healthcheck-healthythreshold
	HealthyThreshold *StringExpr `json:"HealthyThreshold,omitempty" validate:"dive,required"`
	// Interval docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-health-check.html#cfn-elb-healthcheck-interval
	Interval *StringExpr `json:"Interval,omitempty" validate:"dive,required"`
	// Target docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-health-check.html#cfn-elb-healthcheck-target
	Target *StringExpr `json:"Target,omitempty" validate:"dive,required"`
	// Timeout docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-health-check.html#cfn-elb-healthcheck-timeout
	Timeout *StringExpr `json:"Timeout,omitempty" validate:"dive,required"`
	// UnhealthyThreshold docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-health-check.html#cfn-elb-healthcheck-unhealthythreshold
	UnhealthyThreshold *StringExpr `json:"UnhealthyThreshold,omitempty" validate:"dive,required"`
}

// ElasticLoadBalancingLoadBalancerHealthCheckList represents a list of ElasticLoadBalancingLoadBalancerHealthCheck
type ElasticLoadBalancingLoadBalancerHealthCheckList []ElasticLoadBalancingLoadBalancerHealthCheck

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticLoadBalancingLoadBalancerHealthCheckList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticLoadBalancingLoadBalancerHealthCheck{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticLoadBalancingLoadBalancerHealthCheckList{item}
		return nil
	}
	list := []ElasticLoadBalancingLoadBalancerHealthCheck{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticLoadBalancingLoadBalancerHealthCheckList(list)
		return nil
	}
	return err
}

// ElasticLoadBalancingLoadBalancerLBCookieStickinessPolicy represents the AWS::ElasticLoadBalancing::LoadBalancer.LBCookieStickinessPolicy CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-LBCookieStickinessPolicy.html
type ElasticLoadBalancingLoadBalancerLBCookieStickinessPolicy struct {
	// CookieExpirationPeriod docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-LBCookieStickinessPolicy.html#cfn-elb-lbcookiestickinesspolicy-cookieexpirationperiod
	CookieExpirationPeriod *StringExpr `json:"CookieExpirationPeriod,omitempty"`
	// PolicyName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-LBCookieStickinessPolicy.html#cfn-elb-lbcookiestickinesspolicy-policyname
	PolicyName *StringExpr `json:"PolicyName,omitempty"`
}

// ElasticLoadBalancingLoadBalancerLBCookieStickinessPolicyList represents a list of ElasticLoadBalancingLoadBalancerLBCookieStickinessPolicy
type ElasticLoadBalancingLoadBalancerLBCookieStickinessPolicyList []ElasticLoadBalancingLoadBalancerLBCookieStickinessPolicy

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticLoadBalancingLoadBalancerLBCookieStickinessPolicyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticLoadBalancingLoadBalancerLBCookieStickinessPolicy{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticLoadBalancingLoadBalancerLBCookieStickinessPolicyList{item}
		return nil
	}
	list := []ElasticLoadBalancingLoadBalancerLBCookieStickinessPolicy{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticLoadBalancingLoadBalancerLBCookieStickinessPolicyList(list)
		return nil
	}
	return err
}

// ElasticLoadBalancingLoadBalancerListeners represents the AWS::ElasticLoadBalancing::LoadBalancer.Listeners CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html
type ElasticLoadBalancingLoadBalancerListeners struct {
	// InstancePort docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-instanceport
	InstancePort *StringExpr `json:"InstancePort,omitempty" validate:"dive,required"`
	// InstanceProtocol docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-instanceprotocol
	InstanceProtocol *StringExpr `json:"InstanceProtocol,omitempty"`
	// LoadBalancerPort docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-loadbalancerport
	LoadBalancerPort *StringExpr `json:"LoadBalancerPort,omitempty" validate:"dive,required"`
	// PolicyNames docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-policynames
	PolicyNames *StringListExpr `json:"PolicyNames,omitempty"`
	// Protocol docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-protocol
	Protocol *StringExpr `json:"Protocol,omitempty" validate:"dive,required"`
	// SSLCertificateID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-sslcertificateid
	SSLCertificateID *StringExpr `json:"SSLCertificateId,omitempty"`
}

// ElasticLoadBalancingLoadBalancerListenersList represents a list of ElasticLoadBalancingLoadBalancerListeners
type ElasticLoadBalancingLoadBalancerListenersList []ElasticLoadBalancingLoadBalancerListeners

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticLoadBalancingLoadBalancerListenersList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticLoadBalancingLoadBalancerListeners{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticLoadBalancingLoadBalancerListenersList{item}
		return nil
	}
	list := []ElasticLoadBalancingLoadBalancerListeners{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticLoadBalancingLoadBalancerListenersList(list)
		return nil
	}
	return err
}

// ElasticLoadBalancingLoadBalancerPolicies represents the AWS::ElasticLoadBalancing::LoadBalancer.Policies CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-policy.html
type ElasticLoadBalancingLoadBalancerPolicies struct {
	// Attributes docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-policy.html#cfn-ec2-elb-policy-attributes
	Attributes []*interface{} `json:"Attributes,omitempty" validate:"dive,required"`
	// InstancePorts docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-policy.html#cfn-ec2-elb-policy-instanceports
	InstancePorts *StringListExpr `json:"InstancePorts,omitempty"`
	// LoadBalancerPorts docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-policy.html#cfn-ec2-elb-policy-loadbalancerports
	LoadBalancerPorts *StringListExpr `json:"LoadBalancerPorts,omitempty"`
	// PolicyName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-policy.html#cfn-ec2-elb-policy-policyname
	PolicyName *StringExpr `json:"PolicyName,omitempty" validate:"dive,required"`
	// PolicyType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-policy.html#cfn-ec2-elb-policy-policytype
	PolicyType *StringExpr `json:"PolicyType,omitempty" validate:"dive,required"`
}

// ElasticLoadBalancingLoadBalancerPoliciesList represents a list of ElasticLoadBalancingLoadBalancerPolicies
type ElasticLoadBalancingLoadBalancerPoliciesList []ElasticLoadBalancingLoadBalancerPolicies

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticLoadBalancingLoadBalancerPoliciesList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticLoadBalancingLoadBalancerPolicies{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticLoadBalancingLoadBalancerPoliciesList{item}
		return nil
	}
	list := []ElasticLoadBalancingLoadBalancerPolicies{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticLoadBalancingLoadBalancerPoliciesList(list)
		return nil
	}
	return err
}

// ElasticLoadBalancingV2ListenerAction represents the AWS::ElasticLoadBalancingV2::Listener.Action CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-defaultactions.html
type ElasticLoadBalancingV2ListenerAction struct {
	// TargetGroupArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-defaultactions.html#cfn-elasticloadbalancingv2-listener-defaultactions-targetgrouparn
	TargetGroupArn *StringExpr `json:"TargetGroupArn,omitempty" validate:"dive,required"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-defaultactions.html#cfn-elasticloadbalancingv2-listener-defaultactions-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
}

// ElasticLoadBalancingV2ListenerActionList represents a list of ElasticLoadBalancingV2ListenerAction
type ElasticLoadBalancingV2ListenerActionList []ElasticLoadBalancingV2ListenerAction

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticLoadBalancingV2ListenerActionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticLoadBalancingV2ListenerAction{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticLoadBalancingV2ListenerActionList{item}
		return nil
	}
	list := []ElasticLoadBalancingV2ListenerAction{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticLoadBalancingV2ListenerActionList(list)
		return nil
	}
	return err
}

// ElasticLoadBalancingV2ListenerCertificate represents the AWS::ElasticLoadBalancingV2::Listener.Certificate CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-certificates.html
type ElasticLoadBalancingV2ListenerCertificate struct {
	// CertificateArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-certificates.html#cfn-elasticloadbalancingv2-listener-certificates-certificatearn
	CertificateArn *StringExpr `json:"CertificateArn,omitempty"`
}

// ElasticLoadBalancingV2ListenerCertificateList represents a list of ElasticLoadBalancingV2ListenerCertificate
type ElasticLoadBalancingV2ListenerCertificateList []ElasticLoadBalancingV2ListenerCertificate

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticLoadBalancingV2ListenerCertificateList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticLoadBalancingV2ListenerCertificate{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticLoadBalancingV2ListenerCertificateList{item}
		return nil
	}
	list := []ElasticLoadBalancingV2ListenerCertificate{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticLoadBalancingV2ListenerCertificateList(list)
		return nil
	}
	return err
}

// ElasticLoadBalancingV2ListenerRuleAction represents the AWS::ElasticLoadBalancingV2::ListenerRule.Action CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-actions.html
type ElasticLoadBalancingV2ListenerRuleAction struct {
	// TargetGroupArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-actions.html#cfn-elasticloadbalancingv2-listener-actions-targetgrouparn
	TargetGroupArn *StringExpr `json:"TargetGroupArn,omitempty" validate:"dive,required"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-actions.html#cfn-elasticloadbalancingv2-listener-actions-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
}

// ElasticLoadBalancingV2ListenerRuleActionList represents a list of ElasticLoadBalancingV2ListenerRuleAction
type ElasticLoadBalancingV2ListenerRuleActionList []ElasticLoadBalancingV2ListenerRuleAction

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticLoadBalancingV2ListenerRuleActionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticLoadBalancingV2ListenerRuleAction{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticLoadBalancingV2ListenerRuleActionList{item}
		return nil
	}
	list := []ElasticLoadBalancingV2ListenerRuleAction{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticLoadBalancingV2ListenerRuleActionList(list)
		return nil
	}
	return err
}

// ElasticLoadBalancingV2ListenerRuleRuleCondition represents the AWS::ElasticLoadBalancingV2::ListenerRule.RuleCondition CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-conditions.html
type ElasticLoadBalancingV2ListenerRuleRuleCondition struct {
	// Field docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-conditions.html#cfn-elasticloadbalancingv2-listenerrule-conditions-field
	Field *StringExpr `json:"Field,omitempty"`
	// Values docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-conditions.html#cfn-elasticloadbalancingv2-listenerrule-conditions-values
	Values *StringListExpr `json:"Values,omitempty"`
}

// ElasticLoadBalancingV2ListenerRuleRuleConditionList represents a list of ElasticLoadBalancingV2ListenerRuleRuleCondition
type ElasticLoadBalancingV2ListenerRuleRuleConditionList []ElasticLoadBalancingV2ListenerRuleRuleCondition

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticLoadBalancingV2ListenerRuleRuleConditionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticLoadBalancingV2ListenerRuleRuleCondition{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticLoadBalancingV2ListenerRuleRuleConditionList{item}
		return nil
	}
	list := []ElasticLoadBalancingV2ListenerRuleRuleCondition{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticLoadBalancingV2ListenerRuleRuleConditionList(list)
		return nil
	}
	return err
}

// ElasticLoadBalancingV2LoadBalancerLoadBalancerAttribute represents the AWS::ElasticLoadBalancingV2::LoadBalancer.LoadBalancerAttribute CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-loadbalancer-loadbalancerattributes.html
type ElasticLoadBalancingV2LoadBalancerLoadBalancerAttribute struct {
	// Key docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-loadbalancer-loadbalancerattributes.html#cfn-elasticloadbalancingv2-loadbalancer-loadbalancerattributes-key
	Key *StringExpr `json:"Key,omitempty"`
	// Value docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-loadbalancer-loadbalancerattributes.html#cfn-elasticloadbalancingv2-loadbalancer-loadbalancerattributes-value
	Value *StringExpr `json:"Value,omitempty"`
}

// ElasticLoadBalancingV2LoadBalancerLoadBalancerAttributeList represents a list of ElasticLoadBalancingV2LoadBalancerLoadBalancerAttribute
type ElasticLoadBalancingV2LoadBalancerLoadBalancerAttributeList []ElasticLoadBalancingV2LoadBalancerLoadBalancerAttribute

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticLoadBalancingV2LoadBalancerLoadBalancerAttributeList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticLoadBalancingV2LoadBalancerLoadBalancerAttribute{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticLoadBalancingV2LoadBalancerLoadBalancerAttributeList{item}
		return nil
	}
	list := []ElasticLoadBalancingV2LoadBalancerLoadBalancerAttribute{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticLoadBalancingV2LoadBalancerLoadBalancerAttributeList(list)
		return nil
	}
	return err
}

// ElasticLoadBalancingV2TargetGroupMatcher represents the AWS::ElasticLoadBalancingV2::TargetGroup.Matcher CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-matcher.html
type ElasticLoadBalancingV2TargetGroupMatcher struct {
	// HTTPCode docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-matcher.html#cfn-elasticloadbalancingv2-targetgroup-matcher-httpcode
	HTTPCode *StringExpr `json:"HttpCode,omitempty" validate:"dive,required"`
}

// ElasticLoadBalancingV2TargetGroupMatcherList represents a list of ElasticLoadBalancingV2TargetGroupMatcher
type ElasticLoadBalancingV2TargetGroupMatcherList []ElasticLoadBalancingV2TargetGroupMatcher

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticLoadBalancingV2TargetGroupMatcherList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticLoadBalancingV2TargetGroupMatcher{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticLoadBalancingV2TargetGroupMatcherList{item}
		return nil
	}
	list := []ElasticLoadBalancingV2TargetGroupMatcher{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticLoadBalancingV2TargetGroupMatcherList(list)
		return nil
	}
	return err
}

// ElasticLoadBalancingV2TargetGroupTargetDescription represents the AWS::ElasticLoadBalancingV2::TargetGroup.TargetDescription CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-targetdescription.html
type ElasticLoadBalancingV2TargetGroupTargetDescription struct {
	// ID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-targetdescription.html#cfn-elasticloadbalancingv2-targetgroup-targetdescription-id
	ID *StringExpr `json:"Id,omitempty" validate:"dive,required"`
	// Port docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-targetdescription.html#cfn-elasticloadbalancingv2-targetgroup-targetdescription-port
	Port *IntegerExpr `json:"Port,omitempty"`
}

// ElasticLoadBalancingV2TargetGroupTargetDescriptionList represents a list of ElasticLoadBalancingV2TargetGroupTargetDescription
type ElasticLoadBalancingV2TargetGroupTargetDescriptionList []ElasticLoadBalancingV2TargetGroupTargetDescription

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticLoadBalancingV2TargetGroupTargetDescriptionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticLoadBalancingV2TargetGroupTargetDescription{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticLoadBalancingV2TargetGroupTargetDescriptionList{item}
		return nil
	}
	list := []ElasticLoadBalancingV2TargetGroupTargetDescription{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticLoadBalancingV2TargetGroupTargetDescriptionList(list)
		return nil
	}
	return err
}

// ElasticLoadBalancingV2TargetGroupTargetGroupAttribute represents the AWS::ElasticLoadBalancingV2::TargetGroup.TargetGroupAttribute CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-targetgroupattributes.html
type ElasticLoadBalancingV2TargetGroupTargetGroupAttribute struct {
	// Key docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-targetgroupattributes.html#cfn-elasticloadbalancingv2-targetgroup-targetgroupattributes-key
	Key *StringExpr `json:"Key,omitempty"`
	// Value docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-targetgroupattributes.html#cfn-elasticloadbalancingv2-targetgroup-targetgroupattributes-value
	Value *StringExpr `json:"Value,omitempty"`
}

// ElasticLoadBalancingV2TargetGroupTargetGroupAttributeList represents a list of ElasticLoadBalancingV2TargetGroupTargetGroupAttribute
type ElasticLoadBalancingV2TargetGroupTargetGroupAttributeList []ElasticLoadBalancingV2TargetGroupTargetGroupAttribute

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticLoadBalancingV2TargetGroupTargetGroupAttributeList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticLoadBalancingV2TargetGroupTargetGroupAttribute{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticLoadBalancingV2TargetGroupTargetGroupAttributeList{item}
		return nil
	}
	list := []ElasticLoadBalancingV2TargetGroupTargetGroupAttribute{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticLoadBalancingV2TargetGroupTargetGroupAttributeList(list)
		return nil
	}
	return err
}

// ElasticsearchDomainEBSOptions represents the AWS::Elasticsearch::Domain.EBSOptions CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-ebsoptions.html
type ElasticsearchDomainEBSOptions struct {
	// EBSEnabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-ebsoptions.html#cfn-elasticsearch-domain-ebsoptions-ebsenabled
	EBSEnabled *BoolExpr `json:"EBSEnabled,omitempty"`
	// Iops docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-ebsoptions.html#cfn-elasticsearch-domain-ebsoptions-iops
	Iops *IntegerExpr `json:"Iops,omitempty"`
	// VolumeSize docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-ebsoptions.html#cfn-elasticsearch-domain-ebsoptions-volumesize
	VolumeSize *IntegerExpr `json:"VolumeSize,omitempty"`
	// VolumeType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-ebsoptions.html#cfn-elasticsearch-domain-ebsoptions-volumetype
	VolumeType *StringExpr `json:"VolumeType,omitempty"`
}

// ElasticsearchDomainEBSOptionsList represents a list of ElasticsearchDomainEBSOptions
type ElasticsearchDomainEBSOptionsList []ElasticsearchDomainEBSOptions

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticsearchDomainEBSOptionsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticsearchDomainEBSOptions{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticsearchDomainEBSOptionsList{item}
		return nil
	}
	list := []ElasticsearchDomainEBSOptions{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticsearchDomainEBSOptionsList(list)
		return nil
	}
	return err
}

// ElasticsearchDomainElasticsearchClusterConfig represents the AWS::Elasticsearch::Domain.ElasticsearchClusterConfig CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-elasticsearchclusterconfig.html
type ElasticsearchDomainElasticsearchClusterConfig struct {
	// DedicatedMasterCount docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-elasticsearchclusterconfig.html#cfn-elasticsearch-domain-elasticseachclusterconfig-dedicatedmastercount
	DedicatedMasterCount *IntegerExpr `json:"DedicatedMasterCount,omitempty"`
	// DedicatedMasterEnabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-elasticsearchclusterconfig.html#cfn-elasticsearch-domain-elasticseachclusterconfig-dedicatedmasterenabled
	DedicatedMasterEnabled *BoolExpr `json:"DedicatedMasterEnabled,omitempty"`
	// DedicatedMasterType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-elasticsearchclusterconfig.html#cfn-elasticsearch-domain-elasticseachclusterconfig-dedicatedmastertype
	DedicatedMasterType *StringExpr `json:"DedicatedMasterType,omitempty"`
	// InstanceCount docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-elasticsearchclusterconfig.html#cfn-elasticsearch-domain-elasticseachclusterconfig-instancecount
	InstanceCount *IntegerExpr `json:"InstanceCount,omitempty"`
	// InstanceType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-elasticsearchclusterconfig.html#cfn-elasticsearch-domain-elasticseachclusterconfig-instnacetype
	InstanceType *StringExpr `json:"InstanceType,omitempty"`
	// ZoneAwarenessEnabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-elasticsearchclusterconfig.html#cfn-elasticsearch-domain-elasticseachclusterconfig-zoneawarenessenabled
	ZoneAwarenessEnabled *BoolExpr `json:"ZoneAwarenessEnabled,omitempty"`
}

// ElasticsearchDomainElasticsearchClusterConfigList represents a list of ElasticsearchDomainElasticsearchClusterConfig
type ElasticsearchDomainElasticsearchClusterConfigList []ElasticsearchDomainElasticsearchClusterConfig

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticsearchDomainElasticsearchClusterConfigList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticsearchDomainElasticsearchClusterConfig{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticsearchDomainElasticsearchClusterConfigList{item}
		return nil
	}
	list := []ElasticsearchDomainElasticsearchClusterConfig{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticsearchDomainElasticsearchClusterConfigList(list)
		return nil
	}
	return err
}

// ElasticsearchDomainSnapshotOptions represents the AWS::Elasticsearch::Domain.SnapshotOptions CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-snapshotoptions.html
type ElasticsearchDomainSnapshotOptions struct {
	// AutomatedSnapshotStartHour docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-snapshotoptions.html#cfn-elasticsearch-domain-snapshotoptions-automatedsnapshotstarthour
	AutomatedSnapshotStartHour *IntegerExpr `json:"AutomatedSnapshotStartHour,omitempty"`
}

// ElasticsearchDomainSnapshotOptionsList represents a list of ElasticsearchDomainSnapshotOptions
type ElasticsearchDomainSnapshotOptionsList []ElasticsearchDomainSnapshotOptions

// UnmarshalJSON sets the object from the provided JSON representation
func (l *ElasticsearchDomainSnapshotOptionsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := ElasticsearchDomainSnapshotOptions{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = ElasticsearchDomainSnapshotOptionsList{item}
		return nil
	}
	list := []ElasticsearchDomainSnapshotOptions{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = ElasticsearchDomainSnapshotOptionsList(list)
		return nil
	}
	return err
}

// EventsRuleTarget represents the AWS::Events::Rule.Target CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html
type EventsRuleTarget struct {
	// Arn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-arn
	Arn *StringExpr `json:"Arn,omitempty" validate:"dive,required"`
	// ID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-id
	ID *StringExpr `json:"Id,omitempty" validate:"dive,required"`
	// Input docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-input
	Input *StringExpr `json:"Input,omitempty"`
	// InputPath docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-inputpath
	InputPath *StringExpr `json:"InputPath,omitempty"`
	// RoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-rolearn
	RoleArn *StringExpr `json:"RoleArn,omitempty"`
}

// EventsRuleTargetList represents a list of EventsRuleTarget
type EventsRuleTargetList []EventsRuleTarget

// UnmarshalJSON sets the object from the provided JSON representation
func (l *EventsRuleTargetList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := EventsRuleTarget{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = EventsRuleTargetList{item}
		return nil
	}
	list := []EventsRuleTarget{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = EventsRuleTargetList(list)
		return nil
	}
	return err
}

// GameLiftAliasRoutingStrategy represents the AWS::GameLift::Alias.RoutingStrategy CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-alias-routingstrategy.html
type GameLiftAliasRoutingStrategy struct {
	// FleetID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-alias-routingstrategy.html#cfn-gamelift-alias-routingstrategy-fleetid
	FleetID *StringExpr `json:"FleetId,omitempty"`
	// Message docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-alias-routingstrategy.html#cfn-gamelift-alias-routingstrategy-message
	Message *StringExpr `json:"Message,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-alias-routingstrategy.html#cfn-gamelift-alias-routingstrategy-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
}

// GameLiftAliasRoutingStrategyList represents a list of GameLiftAliasRoutingStrategy
type GameLiftAliasRoutingStrategyList []GameLiftAliasRoutingStrategy

// UnmarshalJSON sets the object from the provided JSON representation
func (l *GameLiftAliasRoutingStrategyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := GameLiftAliasRoutingStrategy{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = GameLiftAliasRoutingStrategyList{item}
		return nil
	}
	list := []GameLiftAliasRoutingStrategy{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = GameLiftAliasRoutingStrategyList(list)
		return nil
	}
	return err
}

// GameLiftBuildS3Location represents the AWS::GameLift::Build.S3Location CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-build-storagelocation.html
type GameLiftBuildS3Location struct {
	// Bucket docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-build-storagelocation.html#cfn-gamelift-build-storage-bucket
	Bucket *StringExpr `json:"Bucket,omitempty" validate:"dive,required"`
	// Key docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-build-storagelocation.html#cfn-gamelift-build-storage-key
	Key *StringExpr `json:"Key,omitempty" validate:"dive,required"`
	// RoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-build-storagelocation.html#cfn-gamelift-build-storage-rolearn
	RoleArn *StringExpr `json:"RoleArn,omitempty" validate:"dive,required"`
}

// GameLiftBuildS3LocationList represents a list of GameLiftBuildS3Location
type GameLiftBuildS3LocationList []GameLiftBuildS3Location

// UnmarshalJSON sets the object from the provided JSON representation
func (l *GameLiftBuildS3LocationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := GameLiftBuildS3Location{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = GameLiftBuildS3LocationList{item}
		return nil
	}
	list := []GameLiftBuildS3Location{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = GameLiftBuildS3LocationList(list)
		return nil
	}
	return err
}

// GameLiftFleetIPPermission represents the AWS::GameLift::Fleet.IpPermission CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-ec2inboundpermission.html
type GameLiftFleetIPPermission struct {
	// FromPort docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-ec2inboundpermission.html#cfn-gamelift-fleet-ec2inboundpermissions-fromport
	FromPort *IntegerExpr `json:"FromPort,omitempty" validate:"dive,required"`
	// IPRange docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-ec2inboundpermission.html#cfn-gamelift-fleet-ec2inboundpermissions-iprange
	IPRange *StringExpr `json:"IpRange,omitempty" validate:"dive,required"`
	// Protocol docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-ec2inboundpermission.html#cfn-gamelift-fleet-ec2inboundpermissions-protocol
	Protocol *StringExpr `json:"Protocol,omitempty" validate:"dive,required"`
	// ToPort docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-ec2inboundpermission.html#cfn-gamelift-fleet-ec2inboundpermissions-toport
	ToPort *IntegerExpr `json:"ToPort,omitempty" validate:"dive,required"`
}

// GameLiftFleetIPPermissionList represents a list of GameLiftFleetIPPermission
type GameLiftFleetIPPermissionList []GameLiftFleetIPPermission

// UnmarshalJSON sets the object from the provided JSON representation
func (l *GameLiftFleetIPPermissionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := GameLiftFleetIPPermission{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = GameLiftFleetIPPermissionList{item}
		return nil
	}
	list := []GameLiftFleetIPPermission{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = GameLiftFleetIPPermissionList(list)
		return nil
	}
	return err
}

// IAMGroupPolicy represents the AWS::IAM::Group.Policy CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html
type IAMGroupPolicy struct {
	// PolicyDocument docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html#cfn-iam-policies-policydocument
	PolicyDocument interface{} `json:"PolicyDocument,omitempty" validate:"dive,required"`
	// PolicyName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html#cfn-iam-policies-policyname
	PolicyName *StringExpr `json:"PolicyName,omitempty" validate:"dive,required"`
}

// IAMGroupPolicyList represents a list of IAMGroupPolicy
type IAMGroupPolicyList []IAMGroupPolicy

// UnmarshalJSON sets the object from the provided JSON representation
func (l *IAMGroupPolicyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := IAMGroupPolicy{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = IAMGroupPolicyList{item}
		return nil
	}
	list := []IAMGroupPolicy{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = IAMGroupPolicyList(list)
		return nil
	}
	return err
}

// IAMRolePolicy represents the AWS::IAM::Role.Policy CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html
type IAMRolePolicy struct {
	// PolicyDocument docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html#cfn-iam-policies-policydocument
	PolicyDocument interface{} `json:"PolicyDocument,omitempty" validate:"dive,required"`
	// PolicyName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html#cfn-iam-policies-policyname
	PolicyName *StringExpr `json:"PolicyName,omitempty" validate:"dive,required"`
}

// IAMRolePolicyList represents a list of IAMRolePolicy
type IAMRolePolicyList []IAMRolePolicy

// UnmarshalJSON sets the object from the provided JSON representation
func (l *IAMRolePolicyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := IAMRolePolicy{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = IAMRolePolicyList{item}
		return nil
	}
	list := []IAMRolePolicy{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = IAMRolePolicyList(list)
		return nil
	}
	return err
}

// IAMUserLoginProfile represents the AWS::IAM::User.LoginProfile CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user-loginprofile.html
type IAMUserLoginProfile struct {
	// Password docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user-loginprofile.html#cfn-iam-user-loginprofile-password
	Password *StringExpr `json:"Password,omitempty" validate:"dive,required"`
	// PasswordResetRequired docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user-loginprofile.html#cfn-iam-user-loginprofile-passwordresetrequired
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

// IAMUserPolicy represents the AWS::IAM::User.Policy CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html
type IAMUserPolicy struct {
	// PolicyDocument docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html#cfn-iam-policies-policydocument
	PolicyDocument interface{} `json:"PolicyDocument,omitempty" validate:"dive,required"`
	// PolicyName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html#cfn-iam-policies-policyname
	PolicyName *StringExpr `json:"PolicyName,omitempty" validate:"dive,required"`
}

// IAMUserPolicyList represents a list of IAMUserPolicy
type IAMUserPolicyList []IAMUserPolicy

// UnmarshalJSON sets the object from the provided JSON representation
func (l *IAMUserPolicyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := IAMUserPolicy{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = IAMUserPolicyList{item}
		return nil
	}
	list := []IAMUserPolicy{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = IAMUserPolicyList(list)
		return nil
	}
	return err
}

// IoTThingAttributePayload represents the AWS::IoT::Thing.AttributePayload CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-thing-attributepayload.html
type IoTThingAttributePayload struct {
	// Attributes docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-thing-attributepayload.html#cfn-iot-thing-attributepayload-attributes
	Attributes interface{} `json:"Attributes,omitempty"`
}

// IoTThingAttributePayloadList represents a list of IoTThingAttributePayload
type IoTThingAttributePayloadList []IoTThingAttributePayload

// UnmarshalJSON sets the object from the provided JSON representation
func (l *IoTThingAttributePayloadList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := IoTThingAttributePayload{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = IoTThingAttributePayloadList{item}
		return nil
	}
	list := []IoTThingAttributePayload{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = IoTThingAttributePayloadList(list)
		return nil
	}
	return err
}

// IoTTopicRuleAction represents the AWS::IoT::TopicRule.Action CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html
type IoTTopicRuleAction struct {
	// CloudwatchAlarm docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-cloudwatchalarm
	CloudwatchAlarm *IoTTopicRuleCloudwatchAlarmAction `json:"CloudwatchAlarm,omitempty"`
	// CloudwatchMetric docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-cloudwatchmetric
	CloudwatchMetric *IoTTopicRuleCloudwatchMetricAction `json:"CloudwatchMetric,omitempty"`
	// DynamoDB docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-dynamodb
	DynamoDB *IoTTopicRuleDynamoDBAction `json:"DynamoDB,omitempty"`
	// Elasticsearch docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-elasticsearch
	Elasticsearch *IoTTopicRuleElasticsearchAction `json:"Elasticsearch,omitempty"`
	// Firehose docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-firehose
	Firehose *IoTTopicRuleFirehoseAction `json:"Firehose,omitempty"`
	// Kinesis docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-kinesis
	Kinesis *IoTTopicRuleKinesisAction `json:"Kinesis,omitempty"`
	// Lambda docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-lambda
	Lambda *IoTTopicRuleLambdaAction `json:"Lambda,omitempty"`
	// Republish docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-republish
	Republish *IoTTopicRuleRepublishAction `json:"Republish,omitempty"`
	// S3 docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-s3
	S3 *IoTTopicRuleS3Action `json:"S3,omitempty"`
	// Sns docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-sns
	Sns *IoTTopicRuleSnsAction `json:"Sns,omitempty"`
	// Sqs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-sqs
	Sqs *IoTTopicRuleSqsAction `json:"Sqs,omitempty"`
}

// IoTTopicRuleActionList represents a list of IoTTopicRuleAction
type IoTTopicRuleActionList []IoTTopicRuleAction

// UnmarshalJSON sets the object from the provided JSON representation
func (l *IoTTopicRuleActionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := IoTTopicRuleAction{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = IoTTopicRuleActionList{item}
		return nil
	}
	list := []IoTTopicRuleAction{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = IoTTopicRuleActionList(list)
		return nil
	}
	return err
}

// IoTTopicRuleCloudwatchAlarmAction represents the AWS::IoT::TopicRule.CloudwatchAlarmAction CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchalarm.html
type IoTTopicRuleCloudwatchAlarmAction struct {
	// AlarmName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchalarm.html#cfn-iot-cloudwatchalarm-alarmname
	AlarmName *StringExpr `json:"AlarmName,omitempty" validate:"dive,required"`
	// RoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchalarm.html#cfn-iot-cloudwatchalarm-rolearn
	RoleArn *StringExpr `json:"RoleArn,omitempty" validate:"dive,required"`
	// StateReason docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchalarm.html#cfn-iot-cloudwatchalarm-statereason
	StateReason *StringExpr `json:"StateReason,omitempty" validate:"dive,required"`
	// StateValue docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchalarm.html#cfn-iot-cloudwatchalarm-statevalue
	StateValue *StringExpr `json:"StateValue,omitempty" validate:"dive,required"`
}

// IoTTopicRuleCloudwatchAlarmActionList represents a list of IoTTopicRuleCloudwatchAlarmAction
type IoTTopicRuleCloudwatchAlarmActionList []IoTTopicRuleCloudwatchAlarmAction

// UnmarshalJSON sets the object from the provided JSON representation
func (l *IoTTopicRuleCloudwatchAlarmActionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := IoTTopicRuleCloudwatchAlarmAction{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = IoTTopicRuleCloudwatchAlarmActionList{item}
		return nil
	}
	list := []IoTTopicRuleCloudwatchAlarmAction{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = IoTTopicRuleCloudwatchAlarmActionList(list)
		return nil
	}
	return err
}

// IoTTopicRuleCloudwatchMetricAction represents the AWS::IoT::TopicRule.CloudwatchMetricAction CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchmetric.html
type IoTTopicRuleCloudwatchMetricAction struct {
	// MetricName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchmetric.html#cfn-iot-cloudwatchmetric-metricname
	MetricName *StringExpr `json:"MetricName,omitempty" validate:"dive,required"`
	// MetricNamespace docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchmetric.html#cfn-iot-cloudwatchmetric-metricnamespace
	MetricNamespace *StringExpr `json:"MetricNamespace,omitempty" validate:"dive,required"`
	// MetricTimestamp docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchmetric.html#cfn-iot-cloudwatchmetric-metrictimestamp
	MetricTimestamp *StringExpr `json:"MetricTimestamp,omitempty"`
	// MetricUnit docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchmetric.html#cfn-iot-cloudwatchmetric-metricunit
	MetricUnit *StringExpr `json:"MetricUnit,omitempty" validate:"dive,required"`
	// MetricValue docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchmetric.html#cfn-iot-cloudwatchmetric-metricvalue
	MetricValue *StringExpr `json:"MetricValue,omitempty" validate:"dive,required"`
	// RoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchmetric.html#cfn-iot-cloudwatchmetric-rolearn
	RoleArn *StringExpr `json:"RoleArn,omitempty" validate:"dive,required"`
}

// IoTTopicRuleCloudwatchMetricActionList represents a list of IoTTopicRuleCloudwatchMetricAction
type IoTTopicRuleCloudwatchMetricActionList []IoTTopicRuleCloudwatchMetricAction

// UnmarshalJSON sets the object from the provided JSON representation
func (l *IoTTopicRuleCloudwatchMetricActionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := IoTTopicRuleCloudwatchMetricAction{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = IoTTopicRuleCloudwatchMetricActionList{item}
		return nil
	}
	list := []IoTTopicRuleCloudwatchMetricAction{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = IoTTopicRuleCloudwatchMetricActionList(list)
		return nil
	}
	return err
}

// IoTTopicRuleDynamoDBAction represents the AWS::IoT::TopicRule.DynamoDBAction CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-dynamodb.html
type IoTTopicRuleDynamoDBAction struct {
	// HashKeyField docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-dynamodb.html#cfn-iot-dynamodb-hashkeyfield
	HashKeyField *StringExpr `json:"HashKeyField,omitempty" validate:"dive,required"`
	// HashKeyValue docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-dynamodb.html#cfn-iot-dynamodb-hashkeyvalue
	HashKeyValue *StringExpr `json:"HashKeyValue,omitempty" validate:"dive,required"`
	// PayloadField docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-dynamodb.html#cfn-iot-dynamodb-payloadfield
	PayloadField *StringExpr `json:"PayloadField,omitempty"`
	// RangeKeyField docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-dynamodb.html#cfn-iot-dynamodb-rangekeyfield
	RangeKeyField *StringExpr `json:"RangeKeyField,omitempty" validate:"dive,required"`
	// RangeKeyValue docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-dynamodb.html#cfn-iot-dynamodb-rangekeyvalue
	RangeKeyValue *StringExpr `json:"RangeKeyValue,omitempty" validate:"dive,required"`
	// RoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-dynamodb.html#cfn-iot-dynamodb-rolearn
	RoleArn *StringExpr `json:"RoleArn,omitempty" validate:"dive,required"`
	// TableName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-dynamodb.html#cfn-iot-dynamodb-tablename
	TableName *StringExpr `json:"TableName,omitempty" validate:"dive,required"`
}

// IoTTopicRuleDynamoDBActionList represents a list of IoTTopicRuleDynamoDBAction
type IoTTopicRuleDynamoDBActionList []IoTTopicRuleDynamoDBAction

// UnmarshalJSON sets the object from the provided JSON representation
func (l *IoTTopicRuleDynamoDBActionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := IoTTopicRuleDynamoDBAction{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = IoTTopicRuleDynamoDBActionList{item}
		return nil
	}
	list := []IoTTopicRuleDynamoDBAction{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = IoTTopicRuleDynamoDBActionList(list)
		return nil
	}
	return err
}

// IoTTopicRuleElasticsearchAction represents the AWS::IoT::TopicRule.ElasticsearchAction CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-elasticsearch.html
type IoTTopicRuleElasticsearchAction struct {
	// Endpoint docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-elasticsearch.html#cfn-iot-elasticsearch-endpoint
	Endpoint *StringExpr `json:"Endpoint,omitempty" validate:"dive,required"`
	// ID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-elasticsearch.html#cfn-iot-elasticsearch-id
	ID *StringExpr `json:"Id,omitempty" validate:"dive,required"`
	// Index docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-elasticsearch.html#cfn-iot-elasticsearch-index
	Index *StringExpr `json:"Index,omitempty" validate:"dive,required"`
	// RoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-elasticsearch.html#cfn-iot-elasticsearch-rolearn
	RoleArn *StringExpr `json:"RoleArn,omitempty" validate:"dive,required"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-elasticsearch.html#cfn-iot-elasticsearch-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
}

// IoTTopicRuleElasticsearchActionList represents a list of IoTTopicRuleElasticsearchAction
type IoTTopicRuleElasticsearchActionList []IoTTopicRuleElasticsearchAction

// UnmarshalJSON sets the object from the provided JSON representation
func (l *IoTTopicRuleElasticsearchActionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := IoTTopicRuleElasticsearchAction{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = IoTTopicRuleElasticsearchActionList{item}
		return nil
	}
	list := []IoTTopicRuleElasticsearchAction{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = IoTTopicRuleElasticsearchActionList(list)
		return nil
	}
	return err
}

// IoTTopicRuleFirehoseAction represents the AWS::IoT::TopicRule.FirehoseAction CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-firehose.html
type IoTTopicRuleFirehoseAction struct {
	// DeliveryStreamName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-firehose.html#cfn-iot-firehose-deliverystreamname
	DeliveryStreamName *StringExpr `json:"DeliveryStreamName,omitempty" validate:"dive,required"`
	// RoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-firehose.html#cfn-iot-firehose-rolearn
	RoleArn *StringExpr `json:"RoleArn,omitempty" validate:"dive,required"`
	// Separator docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-firehose.html#cfn-iot-firehose-separator
	Separator *StringExpr `json:"Separator,omitempty"`
}

// IoTTopicRuleFirehoseActionList represents a list of IoTTopicRuleFirehoseAction
type IoTTopicRuleFirehoseActionList []IoTTopicRuleFirehoseAction

// UnmarshalJSON sets the object from the provided JSON representation
func (l *IoTTopicRuleFirehoseActionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := IoTTopicRuleFirehoseAction{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = IoTTopicRuleFirehoseActionList{item}
		return nil
	}
	list := []IoTTopicRuleFirehoseAction{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = IoTTopicRuleFirehoseActionList(list)
		return nil
	}
	return err
}

// IoTTopicRuleKinesisAction represents the AWS::IoT::TopicRule.KinesisAction CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-kinesis.html
type IoTTopicRuleKinesisAction struct {
	// PartitionKey docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-kinesis.html#cfn-iot-kinesis-partitionkey
	PartitionKey *StringExpr `json:"PartitionKey,omitempty"`
	// RoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-kinesis.html#cfn-iot-kinesis-rolearn
	RoleArn *StringExpr `json:"RoleArn,omitempty" validate:"dive,required"`
	// StreamName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-kinesis.html#cfn-iot-kinesis-streamname
	StreamName *StringExpr `json:"StreamName,omitempty" validate:"dive,required"`
}

// IoTTopicRuleKinesisActionList represents a list of IoTTopicRuleKinesisAction
type IoTTopicRuleKinesisActionList []IoTTopicRuleKinesisAction

// UnmarshalJSON sets the object from the provided JSON representation
func (l *IoTTopicRuleKinesisActionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := IoTTopicRuleKinesisAction{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = IoTTopicRuleKinesisActionList{item}
		return nil
	}
	list := []IoTTopicRuleKinesisAction{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = IoTTopicRuleKinesisActionList(list)
		return nil
	}
	return err
}

// IoTTopicRuleLambdaAction represents the AWS::IoT::TopicRule.LambdaAction CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-lambda.html
type IoTTopicRuleLambdaAction struct {
	// FunctionArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-lambda.html#cfn-iot-lambda-functionarn
	FunctionArn *StringExpr `json:"FunctionArn,omitempty" validate:"dive,required"`
}

// IoTTopicRuleLambdaActionList represents a list of IoTTopicRuleLambdaAction
type IoTTopicRuleLambdaActionList []IoTTopicRuleLambdaAction

// UnmarshalJSON sets the object from the provided JSON representation
func (l *IoTTopicRuleLambdaActionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := IoTTopicRuleLambdaAction{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = IoTTopicRuleLambdaActionList{item}
		return nil
	}
	list := []IoTTopicRuleLambdaAction{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = IoTTopicRuleLambdaActionList(list)
		return nil
	}
	return err
}

// IoTTopicRuleRepublishAction represents the AWS::IoT::TopicRule.RepublishAction CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-republish.html
type IoTTopicRuleRepublishAction struct {
	// RoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-republish.html#cfn-iot-republish-rolearn
	RoleArn *StringExpr `json:"RoleArn,omitempty" validate:"dive,required"`
	// Topic docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-republish.html#cfn-iot-republish-topic
	Topic *StringExpr `json:"Topic,omitempty" validate:"dive,required"`
}

// IoTTopicRuleRepublishActionList represents a list of IoTTopicRuleRepublishAction
type IoTTopicRuleRepublishActionList []IoTTopicRuleRepublishAction

// UnmarshalJSON sets the object from the provided JSON representation
func (l *IoTTopicRuleRepublishActionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := IoTTopicRuleRepublishAction{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = IoTTopicRuleRepublishActionList{item}
		return nil
	}
	list := []IoTTopicRuleRepublishAction{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = IoTTopicRuleRepublishActionList(list)
		return nil
	}
	return err
}

// IoTTopicRuleS3Action represents the AWS::IoT::TopicRule.S3Action CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-s3.html
type IoTTopicRuleS3Action struct {
	// BucketName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-s3.html#cfn-iot-s3-bucketname
	BucketName *StringExpr `json:"BucketName,omitempty" validate:"dive,required"`
	// Key docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-s3.html#cfn-iot-s3-key
	Key *StringExpr `json:"Key,omitempty" validate:"dive,required"`
	// RoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-s3.html#cfn-iot-s3-rolearn
	RoleArn *StringExpr `json:"RoleArn,omitempty" validate:"dive,required"`
}

// IoTTopicRuleS3ActionList represents a list of IoTTopicRuleS3Action
type IoTTopicRuleS3ActionList []IoTTopicRuleS3Action

// UnmarshalJSON sets the object from the provided JSON representation
func (l *IoTTopicRuleS3ActionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := IoTTopicRuleS3Action{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = IoTTopicRuleS3ActionList{item}
		return nil
	}
	list := []IoTTopicRuleS3Action{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = IoTTopicRuleS3ActionList(list)
		return nil
	}
	return err
}

// IoTTopicRuleSnsAction represents the AWS::IoT::TopicRule.SnsAction CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-sns.html
type IoTTopicRuleSnsAction struct {
	// MessageFormat docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-sns.html#cfn-iot-sns-snsaction
	MessageFormat *StringExpr `json:"MessageFormat,omitempty"`
	// RoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-sns.html#cfn-iot-sns-rolearn
	RoleArn *StringExpr `json:"RoleArn,omitempty" validate:"dive,required"`
	// TargetArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-sns.html#cfn-iot-sns-targetarn
	TargetArn *StringExpr `json:"TargetArn,omitempty" validate:"dive,required"`
}

// IoTTopicRuleSnsActionList represents a list of IoTTopicRuleSnsAction
type IoTTopicRuleSnsActionList []IoTTopicRuleSnsAction

// UnmarshalJSON sets the object from the provided JSON representation
func (l *IoTTopicRuleSnsActionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := IoTTopicRuleSnsAction{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = IoTTopicRuleSnsActionList{item}
		return nil
	}
	list := []IoTTopicRuleSnsAction{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = IoTTopicRuleSnsActionList(list)
		return nil
	}
	return err
}

// IoTTopicRuleSqsAction represents the AWS::IoT::TopicRule.SqsAction CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-sqs.html
type IoTTopicRuleSqsAction struct {
	// QueueURL docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-sqs.html#cfn-iot-sqs-queueurl
	QueueURL *StringExpr `json:"QueueUrl,omitempty" validate:"dive,required"`
	// RoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-sqs.html#cfn-iot-sqs-rolearn
	RoleArn *StringExpr `json:"RoleArn,omitempty" validate:"dive,required"`
	// UseBase64 docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-sqs.html#cfn-iot-sqs-usebase64
	UseBase64 *BoolExpr `json:"UseBase64,omitempty"`
}

// IoTTopicRuleSqsActionList represents a list of IoTTopicRuleSqsAction
type IoTTopicRuleSqsActionList []IoTTopicRuleSqsAction

// UnmarshalJSON sets the object from the provided JSON representation
func (l *IoTTopicRuleSqsActionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := IoTTopicRuleSqsAction{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = IoTTopicRuleSqsActionList{item}
		return nil
	}
	list := []IoTTopicRuleSqsAction{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = IoTTopicRuleSqsActionList(list)
		return nil
	}
	return err
}

// IoTTopicRuleTopicRulePayload represents the AWS::IoT::TopicRule.TopicRulePayload CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrulepayload.html
type IoTTopicRuleTopicRulePayload struct {
	// Actions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrulepayload.html#cfn-iot-topicrulepayload-actions
	Actions *IoTTopicRuleActionList `json:"Actions,omitempty" validate:"dive,required"`
	// AwsIotSQLVersion docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrulepayload.html#cfn-iot-topicrulepayload-awsiotsqlversion
	AwsIotSQLVersion *StringExpr `json:"AwsIotSqlVersion,omitempty"`
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrulepayload.html#cfn-iot-topicrulepayload-description
	Description *StringExpr `json:"Description,omitempty"`
	// RuleDisabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrulepayload.html#cfn-iot-topicrulepayload-ruledisabled
	RuleDisabled *BoolExpr `json:"RuleDisabled,omitempty" validate:"dive,required"`
	// SQL docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrulepayload.html#cfn-iot-topicrulepayload-sql
	SQL *StringExpr `json:"Sql,omitempty" validate:"dive,required"`
}

// IoTTopicRuleTopicRulePayloadList represents a list of IoTTopicRuleTopicRulePayload
type IoTTopicRuleTopicRulePayloadList []IoTTopicRuleTopicRulePayload

// UnmarshalJSON sets the object from the provided JSON representation
func (l *IoTTopicRuleTopicRulePayloadList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := IoTTopicRuleTopicRulePayload{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = IoTTopicRuleTopicRulePayloadList{item}
		return nil
	}
	list := []IoTTopicRuleTopicRulePayload{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = IoTTopicRuleTopicRulePayloadList(list)
		return nil
	}
	return err
}

// KinesisFirehoseDeliveryStreamBufferingHints represents the AWS::KinesisFirehose::DeliveryStream.BufferingHints CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-bufferinghints.html
type KinesisFirehoseDeliveryStreamBufferingHints struct {
	// IntervalInSeconds docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-bufferinghints.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-bufferinghints-intervalinseconds
	IntervalInSeconds *IntegerExpr `json:"IntervalInSeconds,omitempty" validate:"dive,required"`
	// SizeInMBs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-bufferinghints.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-bufferinghints-sizeinmbs
	SizeInMBs *IntegerExpr `json:"SizeInMBs,omitempty" validate:"dive,required"`
}

// KinesisFirehoseDeliveryStreamBufferingHintsList represents a list of KinesisFirehoseDeliveryStreamBufferingHints
type KinesisFirehoseDeliveryStreamBufferingHintsList []KinesisFirehoseDeliveryStreamBufferingHints

// UnmarshalJSON sets the object from the provided JSON representation
func (l *KinesisFirehoseDeliveryStreamBufferingHintsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := KinesisFirehoseDeliveryStreamBufferingHints{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = KinesisFirehoseDeliveryStreamBufferingHintsList{item}
		return nil
	}
	list := []KinesisFirehoseDeliveryStreamBufferingHints{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = KinesisFirehoseDeliveryStreamBufferingHintsList(list)
		return nil
	}
	return err
}

// KinesisFirehoseDeliveryStreamCloudWatchLoggingOptions represents the AWS::KinesisFirehose::DeliveryStream.CloudWatchLoggingOptions CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-destination-cloudwatchloggingoptions.html
type KinesisFirehoseDeliveryStreamCloudWatchLoggingOptions struct {
	// Enabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-destination-cloudwatchloggingoptions.html#cfn-kinesisfirehose-kinesisdeliverystream-destination-cloudwatchloggingoptions-enabled
	Enabled *BoolExpr `json:"Enabled,omitempty"`
	// LogGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-destination-cloudwatchloggingoptions.html#cfn-kinesisfirehose-kinesisdeliverystream-destination-cloudwatchloggingoptions-loggroupname
	LogGroupName *StringExpr `json:"LogGroupName,omitempty"`
	// LogStreamName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-destination-cloudwatchloggingoptions.html#cfn-kinesisfirehose-kinesisdeliverystream-destination-cloudwatchloggingoptions-logstreamname
	LogStreamName *StringExpr `json:"LogStreamName,omitempty"`
}

// KinesisFirehoseDeliveryStreamCloudWatchLoggingOptionsList represents a list of KinesisFirehoseDeliveryStreamCloudWatchLoggingOptions
type KinesisFirehoseDeliveryStreamCloudWatchLoggingOptionsList []KinesisFirehoseDeliveryStreamCloudWatchLoggingOptions

// UnmarshalJSON sets the object from the provided JSON representation
func (l *KinesisFirehoseDeliveryStreamCloudWatchLoggingOptionsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := KinesisFirehoseDeliveryStreamCloudWatchLoggingOptions{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = KinesisFirehoseDeliveryStreamCloudWatchLoggingOptionsList{item}
		return nil
	}
	list := []KinesisFirehoseDeliveryStreamCloudWatchLoggingOptions{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = KinesisFirehoseDeliveryStreamCloudWatchLoggingOptionsList(list)
		return nil
	}
	return err
}

// KinesisFirehoseDeliveryStreamCopyCommand represents the AWS::KinesisFirehose::DeliveryStream.CopyCommand CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-copycommand.html
type KinesisFirehoseDeliveryStreamCopyCommand struct {
	// CopyOptions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-copycommand.html#cfn-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-copycommand-copyoptions
	CopyOptions *StringExpr `json:"CopyOptions,omitempty"`
	// DataTableColumns docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-copycommand.html#cfn-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-copycommand-datatablecolumns
	DataTableColumns *StringExpr `json:"DataTableColumns,omitempty"`
	// DataTableName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-copycommand.html#cfn-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-copycommand-datatablename
	DataTableName *StringExpr `json:"DataTableName,omitempty" validate:"dive,required"`
}

// KinesisFirehoseDeliveryStreamCopyCommandList represents a list of KinesisFirehoseDeliveryStreamCopyCommand
type KinesisFirehoseDeliveryStreamCopyCommandList []KinesisFirehoseDeliveryStreamCopyCommand

// UnmarshalJSON sets the object from the provided JSON representation
func (l *KinesisFirehoseDeliveryStreamCopyCommandList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := KinesisFirehoseDeliveryStreamCopyCommand{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = KinesisFirehoseDeliveryStreamCopyCommandList{item}
		return nil
	}
	list := []KinesisFirehoseDeliveryStreamCopyCommand{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = KinesisFirehoseDeliveryStreamCopyCommandList(list)
		return nil
	}
	return err
}

// KinesisFirehoseDeliveryStreamElasticsearchBufferingHints represents the AWS::KinesisFirehose::DeliveryStream.ElasticsearchBufferingHints CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration.html
type KinesisFirehoseDeliveryStreamElasticsearchBufferingHints struct {
	// IntervalInSeconds docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration-bufferinghints-intervalinseconds
	IntervalInSeconds *IntegerExpr `json:"IntervalInSeconds,omitempty" validate:"dive,required"`
	// SizeInMBs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration-bufferinghints-sizeinmbs
	SizeInMBs *IntegerExpr `json:"SizeInMBs,omitempty" validate:"dive,required"`
}

// KinesisFirehoseDeliveryStreamElasticsearchBufferingHintsList represents a list of KinesisFirehoseDeliveryStreamElasticsearchBufferingHints
type KinesisFirehoseDeliveryStreamElasticsearchBufferingHintsList []KinesisFirehoseDeliveryStreamElasticsearchBufferingHints

// UnmarshalJSON sets the object from the provided JSON representation
func (l *KinesisFirehoseDeliveryStreamElasticsearchBufferingHintsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := KinesisFirehoseDeliveryStreamElasticsearchBufferingHints{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = KinesisFirehoseDeliveryStreamElasticsearchBufferingHintsList{item}
		return nil
	}
	list := []KinesisFirehoseDeliveryStreamElasticsearchBufferingHints{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = KinesisFirehoseDeliveryStreamElasticsearchBufferingHintsList(list)
		return nil
	}
	return err
}

// KinesisFirehoseDeliveryStreamElasticsearchDestinationConfiguration represents the AWS::KinesisFirehose::DeliveryStream.ElasticsearchDestinationConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration.html
type KinesisFirehoseDeliveryStreamElasticsearchDestinationConfiguration struct {
	// BufferingHints docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration-bufferinghints
	BufferingHints *KinesisFirehoseDeliveryStreamElasticsearchBufferingHints `json:"BufferingHints,omitempty" validate:"dive,required"`
	// CloudWatchLoggingOptions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration-cloudwatchloggingoptions
	CloudWatchLoggingOptions *KinesisFirehoseDeliveryStreamCloudWatchLoggingOptions `json:"CloudWatchLoggingOptions,omitempty"`
	// DomainARN docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration-domainarn
	DomainARN *StringExpr `json:"DomainARN,omitempty" validate:"dive,required"`
	// IndexName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration-indexname
	IndexName *StringExpr `json:"IndexName,omitempty" validate:"dive,required"`
	// IndexRotationPeriod docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration-indexrotationperiod
	IndexRotationPeriod *StringExpr `json:"IndexRotationPeriod,omitempty" validate:"dive,required"`
	// RetryOptions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration-retryoptions
	RetryOptions *KinesisFirehoseDeliveryStreamElasticsearchRetryOptions `json:"RetryOptions,omitempty" validate:"dive,required"`
	// RoleARN docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration-rolearn
	RoleARN *StringExpr `json:"RoleARN,omitempty" validate:"dive,required"`
	// S3BackupMode docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration-s3backupmode
	S3BackupMode *StringExpr `json:"S3BackupMode,omitempty" validate:"dive,required"`
	// S3Configuration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration-s3configuration
	S3Configuration *KinesisFirehoseDeliveryStreamS3DestinationConfiguration `json:"S3Configuration,omitempty" validate:"dive,required"`
	// TypeName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration-typename
	TypeName *StringExpr `json:"TypeName,omitempty" validate:"dive,required"`
}

// KinesisFirehoseDeliveryStreamElasticsearchDestinationConfigurationList represents a list of KinesisFirehoseDeliveryStreamElasticsearchDestinationConfiguration
type KinesisFirehoseDeliveryStreamElasticsearchDestinationConfigurationList []KinesisFirehoseDeliveryStreamElasticsearchDestinationConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *KinesisFirehoseDeliveryStreamElasticsearchDestinationConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := KinesisFirehoseDeliveryStreamElasticsearchDestinationConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = KinesisFirehoseDeliveryStreamElasticsearchDestinationConfigurationList{item}
		return nil
	}
	list := []KinesisFirehoseDeliveryStreamElasticsearchDestinationConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = KinesisFirehoseDeliveryStreamElasticsearchDestinationConfigurationList(list)
		return nil
	}
	return err
}

// KinesisFirehoseDeliveryStreamElasticsearchRetryOptions represents the AWS::KinesisFirehose::DeliveryStream.ElasticsearchRetryOptions CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration-retryoptions.html
type KinesisFirehoseDeliveryStreamElasticsearchRetryOptions struct {
	// DurationInSeconds docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration-retryoptions.html#cfn-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration-retryoptions-durationinseconds
	DurationInSeconds *IntegerExpr `json:"DurationInSeconds,omitempty" validate:"dive,required"`
}

// KinesisFirehoseDeliveryStreamElasticsearchRetryOptionsList represents a list of KinesisFirehoseDeliveryStreamElasticsearchRetryOptions
type KinesisFirehoseDeliveryStreamElasticsearchRetryOptionsList []KinesisFirehoseDeliveryStreamElasticsearchRetryOptions

// UnmarshalJSON sets the object from the provided JSON representation
func (l *KinesisFirehoseDeliveryStreamElasticsearchRetryOptionsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := KinesisFirehoseDeliveryStreamElasticsearchRetryOptions{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = KinesisFirehoseDeliveryStreamElasticsearchRetryOptionsList{item}
		return nil
	}
	list := []KinesisFirehoseDeliveryStreamElasticsearchRetryOptions{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = KinesisFirehoseDeliveryStreamElasticsearchRetryOptionsList(list)
		return nil
	}
	return err
}

// KinesisFirehoseDeliveryStreamEncryptionConfiguration represents the AWS::KinesisFirehose::DeliveryStream.EncryptionConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-encryptionconfiguration.html
type KinesisFirehoseDeliveryStreamEncryptionConfiguration struct {
	// KMSEncryptionConfig docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-encryptionconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-encryptionconfiguration-kmsencryptionconfig
	KMSEncryptionConfig *KinesisFirehoseDeliveryStreamKMSEncryptionConfig `json:"KMSEncryptionConfig,omitempty"`
	// NoEncryptionConfig docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-encryptionconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-encryptionconfiguration-noencryptionconfig
	NoEncryptionConfig *StringExpr `json:"NoEncryptionConfig,omitempty"`
}

// KinesisFirehoseDeliveryStreamEncryptionConfigurationList represents a list of KinesisFirehoseDeliveryStreamEncryptionConfiguration
type KinesisFirehoseDeliveryStreamEncryptionConfigurationList []KinesisFirehoseDeliveryStreamEncryptionConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *KinesisFirehoseDeliveryStreamEncryptionConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := KinesisFirehoseDeliveryStreamEncryptionConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = KinesisFirehoseDeliveryStreamEncryptionConfigurationList{item}
		return nil
	}
	list := []KinesisFirehoseDeliveryStreamEncryptionConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = KinesisFirehoseDeliveryStreamEncryptionConfigurationList(list)
		return nil
	}
	return err
}

// KinesisFirehoseDeliveryStreamKMSEncryptionConfig represents the AWS::KinesisFirehose::DeliveryStream.KMSEncryptionConfig CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-encryptionconfiguration-kmsencryptionconfig.html
type KinesisFirehoseDeliveryStreamKMSEncryptionConfig struct {
	// AWSKMSKeyARN docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-encryptionconfiguration-kmsencryptionconfig.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-encryptionconfiguration-kmsencryptionconfig-awskmskeyarn
	AWSKMSKeyARN *StringExpr `json:"AWSKMSKeyARN,omitempty" validate:"dive,required"`
}

// KinesisFirehoseDeliveryStreamKMSEncryptionConfigList represents a list of KinesisFirehoseDeliveryStreamKMSEncryptionConfig
type KinesisFirehoseDeliveryStreamKMSEncryptionConfigList []KinesisFirehoseDeliveryStreamKMSEncryptionConfig

// UnmarshalJSON sets the object from the provided JSON representation
func (l *KinesisFirehoseDeliveryStreamKMSEncryptionConfigList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := KinesisFirehoseDeliveryStreamKMSEncryptionConfig{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = KinesisFirehoseDeliveryStreamKMSEncryptionConfigList{item}
		return nil
	}
	list := []KinesisFirehoseDeliveryStreamKMSEncryptionConfig{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = KinesisFirehoseDeliveryStreamKMSEncryptionConfigList(list)
		return nil
	}
	return err
}

// KinesisFirehoseDeliveryStreamRedshiftDestinationConfiguration represents the AWS::KinesisFirehose::DeliveryStream.RedshiftDestinationConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration.html
type KinesisFirehoseDeliveryStreamRedshiftDestinationConfiguration struct {
	// CloudWatchLoggingOptions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-cloudwatchloggingoptions
	CloudWatchLoggingOptions *KinesisFirehoseDeliveryStreamCloudWatchLoggingOptions `json:"CloudWatchLoggingOptions,omitempty"`
	// ClusterJDBCURL docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-clusterjdbcurl
	ClusterJDBCURL *StringExpr `json:"ClusterJDBCURL,omitempty" validate:"dive,required"`
	// CopyCommand docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-copycommand
	CopyCommand *KinesisFirehoseDeliveryStreamCopyCommand `json:"CopyCommand,omitempty" validate:"dive,required"`
	// Password docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-password
	Password *StringExpr `json:"Password,omitempty" validate:"dive,required"`
	// RoleARN docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-rolearn
	RoleARN *StringExpr `json:"RoleARN,omitempty" validate:"dive,required"`
	// S3Configuration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-s3configuration
	S3Configuration *KinesisFirehoseDeliveryStreamS3DestinationConfiguration `json:"S3Configuration,omitempty" validate:"dive,required"`
	// Username docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-redshiftdestinationconfiguration-usename
	Username *StringExpr `json:"Username,omitempty" validate:"dive,required"`
}

// KinesisFirehoseDeliveryStreamRedshiftDestinationConfigurationList represents a list of KinesisFirehoseDeliveryStreamRedshiftDestinationConfiguration
type KinesisFirehoseDeliveryStreamRedshiftDestinationConfigurationList []KinesisFirehoseDeliveryStreamRedshiftDestinationConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *KinesisFirehoseDeliveryStreamRedshiftDestinationConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := KinesisFirehoseDeliveryStreamRedshiftDestinationConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = KinesisFirehoseDeliveryStreamRedshiftDestinationConfigurationList{item}
		return nil
	}
	list := []KinesisFirehoseDeliveryStreamRedshiftDestinationConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = KinesisFirehoseDeliveryStreamRedshiftDestinationConfigurationList(list)
		return nil
	}
	return err
}

// KinesisFirehoseDeliveryStreamS3DestinationConfiguration represents the AWS::KinesisFirehose::DeliveryStream.S3DestinationConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration.html
type KinesisFirehoseDeliveryStreamS3DestinationConfiguration struct {
	// BucketARN docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-bucketarn
	BucketARN *StringExpr `json:"BucketARN,omitempty" validate:"dive,required"`
	// BufferingHints docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-bufferinghints
	BufferingHints *KinesisFirehoseDeliveryStreamBufferingHints `json:"BufferingHints,omitempty" validate:"dive,required"`
	// CloudWatchLoggingOptions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-cloudwatchloggingoptions
	CloudWatchLoggingOptions *KinesisFirehoseDeliveryStreamCloudWatchLoggingOptions `json:"CloudWatchLoggingOptions,omitempty"`
	// CompressionFormat docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-compressionformat
	CompressionFormat *StringExpr `json:"CompressionFormat,omitempty" validate:"dive,required"`
	// EncryptionConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-encryptionconfiguration
	EncryptionConfiguration *KinesisFirehoseDeliveryStreamEncryptionConfiguration `json:"EncryptionConfiguration,omitempty"`
	// Prefix docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-prefix
	Prefix *StringExpr `json:"Prefix,omitempty" validate:"dive,required"`
	// RoleARN docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-rolearn
	RoleARN *StringExpr `json:"RoleARN,omitempty" validate:"dive,required"`
}

// KinesisFirehoseDeliveryStreamS3DestinationConfigurationList represents a list of KinesisFirehoseDeliveryStreamS3DestinationConfiguration
type KinesisFirehoseDeliveryStreamS3DestinationConfigurationList []KinesisFirehoseDeliveryStreamS3DestinationConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *KinesisFirehoseDeliveryStreamS3DestinationConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := KinesisFirehoseDeliveryStreamS3DestinationConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = KinesisFirehoseDeliveryStreamS3DestinationConfigurationList{item}
		return nil
	}
	list := []KinesisFirehoseDeliveryStreamS3DestinationConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = KinesisFirehoseDeliveryStreamS3DestinationConfigurationList(list)
		return nil
	}
	return err
}

// LambdaFunctionCode represents the AWS::Lambda::Function.Code CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-code.html
type LambdaFunctionCode struct {
	// S3Bucket docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-code.html#cfn-lambda-function-code-s3bucket
	S3Bucket *StringExpr `json:"S3Bucket,omitempty"`
	// S3Key docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-code.html#cfn-lambda-function-code-s3key
	S3Key *StringExpr `json:"S3Key,omitempty"`
	// S3ObjectVersion docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-code.html#cfn-lambda-function-code-s3objectversion
	S3ObjectVersion *StringExpr `json:"S3ObjectVersion,omitempty"`
	// ZipFile docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-code.html#cfn-lambda-function-code-zipfile
	ZipFile *StringExpr `json:"ZipFile,omitempty"`
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

// LambdaFunctionDeadLetterConfig represents the AWS::Lambda::Function.DeadLetterConfig CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-deadletterconfig.html
type LambdaFunctionDeadLetterConfig struct {
	// TargetArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-deadletterconfig.html#cfn-lambda-function-deadletterconfig-targetarn
	TargetArn *StringExpr `json:"TargetArn,omitempty"`
}

// LambdaFunctionDeadLetterConfigList represents a list of LambdaFunctionDeadLetterConfig
type LambdaFunctionDeadLetterConfigList []LambdaFunctionDeadLetterConfig

// UnmarshalJSON sets the object from the provided JSON representation
func (l *LambdaFunctionDeadLetterConfigList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := LambdaFunctionDeadLetterConfig{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = LambdaFunctionDeadLetterConfigList{item}
		return nil
	}
	list := []LambdaFunctionDeadLetterConfig{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = LambdaFunctionDeadLetterConfigList(list)
		return nil
	}
	return err
}

// LambdaFunctionEnvironment represents the AWS::Lambda::Function.Environment CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-environment.html
type LambdaFunctionEnvironment struct {
	// Variables docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-environment.html#cfn-lambda-function-environment-variables
	Variables interface{} `json:"Variables,omitempty"`
}

// LambdaFunctionEnvironmentList represents a list of LambdaFunctionEnvironment
type LambdaFunctionEnvironmentList []LambdaFunctionEnvironment

// UnmarshalJSON sets the object from the provided JSON representation
func (l *LambdaFunctionEnvironmentList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := LambdaFunctionEnvironment{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = LambdaFunctionEnvironmentList{item}
		return nil
	}
	list := []LambdaFunctionEnvironment{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = LambdaFunctionEnvironmentList(list)
		return nil
	}
	return err
}

// LambdaFunctionTracingConfig represents the AWS::Lambda::Function.TracingConfig CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-tracingconfig.html
type LambdaFunctionTracingConfig struct {
	// Mode docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-tracingconfig.html#cfn-lambda-function-tracingconfig-mode
	Mode *StringExpr `json:"Mode,omitempty"`
}

// LambdaFunctionTracingConfigList represents a list of LambdaFunctionTracingConfig
type LambdaFunctionTracingConfigList []LambdaFunctionTracingConfig

// UnmarshalJSON sets the object from the provided JSON representation
func (l *LambdaFunctionTracingConfigList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := LambdaFunctionTracingConfig{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = LambdaFunctionTracingConfigList{item}
		return nil
	}
	list := []LambdaFunctionTracingConfig{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = LambdaFunctionTracingConfigList(list)
		return nil
	}
	return err
}

// LambdaFunctionVPCConfig represents the AWS::Lambda::Function.VpcConfig CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-vpcconfig.html
type LambdaFunctionVPCConfig struct {
	// SecurityGroupIDs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-vpcconfig.html#cfn-lambda-function-vpcconfig-securitygroupids
	SecurityGroupIDs *StringListExpr `json:"SecurityGroupIds,omitempty" validate:"dive,required"`
	// SubnetIDs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-vpcconfig.html#cfn-lambda-function-vpcconfig-subnetids
	SubnetIDs *StringListExpr `json:"SubnetIds,omitempty" validate:"dive,required"`
}

// LambdaFunctionVPCConfigList represents a list of LambdaFunctionVPCConfig
type LambdaFunctionVPCConfigList []LambdaFunctionVPCConfig

// UnmarshalJSON sets the object from the provided JSON representation
func (l *LambdaFunctionVPCConfigList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := LambdaFunctionVPCConfig{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = LambdaFunctionVPCConfigList{item}
		return nil
	}
	list := []LambdaFunctionVPCConfig{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = LambdaFunctionVPCConfigList(list)
		return nil
	}
	return err
}

// LogsMetricFilterMetricTransformation represents the AWS::Logs::MetricFilter.MetricTransformation CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-metricfilter-metrictransformation.html
type LogsMetricFilterMetricTransformation struct {
	// MetricName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-metricfilter-metrictransformation.html#cfn-cwl-metricfilter-metrictransformation-metricname
	MetricName *StringExpr `json:"MetricName,omitempty" validate:"dive,required"`
	// MetricNamespace docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-metricfilter-metrictransformation.html#cfn-cwl-metricfilter-metrictransformation-metricnamespace
	MetricNamespace *StringExpr `json:"MetricNamespace,omitempty" validate:"dive,required"`
	// MetricValue docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-metricfilter-metrictransformation.html#cfn-cwl-metricfilter-metrictransformation-metricvalue
	MetricValue *StringExpr `json:"MetricValue,omitempty" validate:"dive,required"`
}

// LogsMetricFilterMetricTransformationList represents a list of LogsMetricFilterMetricTransformation
type LogsMetricFilterMetricTransformationList []LogsMetricFilterMetricTransformation

// UnmarshalJSON sets the object from the provided JSON representation
func (l *LogsMetricFilterMetricTransformationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := LogsMetricFilterMetricTransformation{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = LogsMetricFilterMetricTransformationList{item}
		return nil
	}
	list := []LogsMetricFilterMetricTransformation{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = LogsMetricFilterMetricTransformationList(list)
		return nil
	}
	return err
}

// OpsWorksAppDataSource represents the AWS::OpsWorks::App.DataSource CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-datasource.html
type OpsWorksAppDataSource struct {
	// Arn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-datasource.html#cfn-opsworks-app-datasource-arn
	Arn *StringExpr `json:"Arn,omitempty"`
	// DatabaseName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-datasource.html#cfn-opsworks-app-datasource-databasename
	DatabaseName *StringExpr `json:"DatabaseName,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-datasource.html#cfn-opsworks-app-datasource-type
	Type *StringExpr `json:"Type,omitempty"`
}

// OpsWorksAppDataSourceList represents a list of OpsWorksAppDataSource
type OpsWorksAppDataSourceList []OpsWorksAppDataSource

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksAppDataSourceList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksAppDataSource{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksAppDataSourceList{item}
		return nil
	}
	list := []OpsWorksAppDataSource{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksAppDataSourceList(list)
		return nil
	}
	return err
}

// OpsWorksAppEnvironmentVariable represents the AWS::OpsWorks::App.EnvironmentVariable CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-environment.html
type OpsWorksAppEnvironmentVariable struct {
	// Key docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-environment.html#cfn-opsworks-app-environment-key
	Key *StringExpr `json:"Key,omitempty" validate:"dive,required"`
	// Secure docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-environment.html#cfn-opsworks-app-environment-secure
	Secure *BoolExpr `json:"Secure,omitempty"`
	// Value docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-environment.html#value
	Value *StringExpr `json:"Value,omitempty" validate:"dive,required"`
}

// OpsWorksAppEnvironmentVariableList represents a list of OpsWorksAppEnvironmentVariable
type OpsWorksAppEnvironmentVariableList []OpsWorksAppEnvironmentVariable

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksAppEnvironmentVariableList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksAppEnvironmentVariable{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksAppEnvironmentVariableList{item}
		return nil
	}
	list := []OpsWorksAppEnvironmentVariable{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksAppEnvironmentVariableList(list)
		return nil
	}
	return err
}

// OpsWorksAppSource represents the AWS::OpsWorks::App.Source CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html
type OpsWorksAppSource struct {
	// Password docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html#cfn-opsworks-custcookbooksource-pw
	Password *StringExpr `json:"Password,omitempty"`
	// Revision docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html#cfn-opsworks-custcookbooksource-revision
	Revision *StringExpr `json:"Revision,omitempty"`
	// SSHKey docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html#cfn-opsworks-custcookbooksource-sshkey
	SSHKey *StringExpr `json:"SshKey,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html#cfn-opsworks-custcookbooksource-type
	Type *StringExpr `json:"Type,omitempty"`
	// URL docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html#cfn-opsworks-custcookbooksource-url
	URL *StringExpr `json:"Url,omitempty"`
	// Username docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html#cfn-opsworks-custcookbooksource-username
	Username *StringExpr `json:"Username,omitempty"`
}

// OpsWorksAppSourceList represents a list of OpsWorksAppSource
type OpsWorksAppSourceList []OpsWorksAppSource

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksAppSourceList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksAppSource{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksAppSourceList{item}
		return nil
	}
	list := []OpsWorksAppSource{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksAppSourceList(list)
		return nil
	}
	return err
}

// OpsWorksAppSslConfiguration represents the AWS::OpsWorks::App.SslConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-sslconfiguration.html
type OpsWorksAppSslConfiguration struct {
	// Certificate docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-sslconfiguration.html#cfn-opsworks-app-sslconfig-certificate
	Certificate *StringExpr `json:"Certificate,omitempty"`
	// Chain docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-sslconfiguration.html#cfn-opsworks-app-sslconfig-chain
	Chain *StringExpr `json:"Chain,omitempty"`
	// PrivateKey docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-sslconfiguration.html#cfn-opsworks-app-sslconfig-privatekey
	PrivateKey *StringExpr `json:"PrivateKey,omitempty"`
}

// OpsWorksAppSslConfigurationList represents a list of OpsWorksAppSslConfiguration
type OpsWorksAppSslConfigurationList []OpsWorksAppSslConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksAppSslConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksAppSslConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksAppSslConfigurationList{item}
		return nil
	}
	list := []OpsWorksAppSslConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksAppSslConfigurationList(list)
		return nil
	}
	return err
}

// OpsWorksInstanceBlockDeviceMapping represents the AWS::OpsWorks::Instance.BlockDeviceMapping CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-blockdevicemapping.html
type OpsWorksInstanceBlockDeviceMapping struct {
	// DeviceName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-blockdevicemapping.html#cfn-opsworks-instance-blockdevicemapping-devicename
	DeviceName *StringExpr `json:"DeviceName,omitempty"`
	// Ebs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-blockdevicemapping.html#cfn-opsworks-instance-blockdevicemapping-ebs
	Ebs *OpsWorksInstanceEbsBlockDevice `json:"Ebs,omitempty"`
	// NoDevice docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-blockdevicemapping.html#cfn-opsworks-instance-blockdevicemapping-nodevice
	NoDevice *StringExpr `json:"NoDevice,omitempty"`
	// VirtualName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-blockdevicemapping.html#cfn-opsworks-instance-blockdevicemapping-virtualname
	VirtualName *StringExpr `json:"VirtualName,omitempty"`
}

// OpsWorksInstanceBlockDeviceMappingList represents a list of OpsWorksInstanceBlockDeviceMapping
type OpsWorksInstanceBlockDeviceMappingList []OpsWorksInstanceBlockDeviceMapping

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksInstanceBlockDeviceMappingList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksInstanceBlockDeviceMapping{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksInstanceBlockDeviceMappingList{item}
		return nil
	}
	list := []OpsWorksInstanceBlockDeviceMapping{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksInstanceBlockDeviceMappingList(list)
		return nil
	}
	return err
}

// OpsWorksInstanceEbsBlockDevice represents the AWS::OpsWorks::Instance.EbsBlockDevice CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-ebsblockdevice.html
type OpsWorksInstanceEbsBlockDevice struct {
	// DeleteOnTermination docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-ebsblockdevice.html#cfn-opsworks-instance-ebsblockdevice-deleteontermination
	DeleteOnTermination *BoolExpr `json:"DeleteOnTermination,omitempty"`
	// Iops docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-ebsblockdevice.html#cfn-opsworks-instance-ebsblockdevice-iops
	Iops *IntegerExpr `json:"Iops,omitempty"`
	// SnapshotID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-ebsblockdevice.html#cfn-opsworks-instance-ebsblockdevice-snapshotid
	SnapshotID *StringExpr `json:"SnapshotId,omitempty"`
	// VolumeSize docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-ebsblockdevice.html#cfn-opsworks-instance-ebsblockdevice-volumesize
	VolumeSize *IntegerExpr `json:"VolumeSize,omitempty"`
	// VolumeType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-ebsblockdevice.html#cfn-opsworks-instance-ebsblockdevice-volumetype
	VolumeType *StringExpr `json:"VolumeType,omitempty"`
}

// OpsWorksInstanceEbsBlockDeviceList represents a list of OpsWorksInstanceEbsBlockDevice
type OpsWorksInstanceEbsBlockDeviceList []OpsWorksInstanceEbsBlockDevice

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksInstanceEbsBlockDeviceList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksInstanceEbsBlockDevice{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksInstanceEbsBlockDeviceList{item}
		return nil
	}
	list := []OpsWorksInstanceEbsBlockDevice{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksInstanceEbsBlockDeviceList(list)
		return nil
	}
	return err
}

// OpsWorksInstanceTimeBasedAutoScaling represents the AWS::OpsWorks::Instance.TimeBasedAutoScaling CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-timebasedautoscaling.html
type OpsWorksInstanceTimeBasedAutoScaling struct {
	// Friday docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-timebasedautoscaling.html#cfn-opsworks-instance-timebasedautoscaling-friday
	Friday interface{} `json:"Friday,omitempty"`
	// Monday docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-timebasedautoscaling.html#cfn-opsworks-instance-timebasedautoscaling-monday
	Monday interface{} `json:"Monday,omitempty"`
	// Saturday docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-timebasedautoscaling.html#cfn-opsworks-instance-timebasedautoscaling-saturday
	Saturday interface{} `json:"Saturday,omitempty"`
	// Sunday docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-timebasedautoscaling.html#cfn-opsworks-instance-timebasedautoscaling-sunday
	Sunday interface{} `json:"Sunday,omitempty"`
	// Thursday docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-timebasedautoscaling.html#cfn-opsworks-instance-timebasedautoscaling-thursday
	Thursday interface{} `json:"Thursday,omitempty"`
	// Tuesday docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-timebasedautoscaling.html#cfn-opsworks-instance-timebasedautoscaling-tuesday
	Tuesday interface{} `json:"Tuesday,omitempty"`
	// Wednesday docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-timebasedautoscaling.html#cfn-opsworks-instance-timebasedautoscaling-wednesday
	Wednesday interface{} `json:"Wednesday,omitempty"`
}

// OpsWorksInstanceTimeBasedAutoScalingList represents a list of OpsWorksInstanceTimeBasedAutoScaling
type OpsWorksInstanceTimeBasedAutoScalingList []OpsWorksInstanceTimeBasedAutoScaling

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksInstanceTimeBasedAutoScalingList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksInstanceTimeBasedAutoScaling{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksInstanceTimeBasedAutoScalingList{item}
		return nil
	}
	list := []OpsWorksInstanceTimeBasedAutoScaling{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksInstanceTimeBasedAutoScalingList(list)
		return nil
	}
	return err
}

// OpsWorksLayerAutoScalingThresholds represents the AWS::OpsWorks::Layer.AutoScalingThresholds CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling-autoscalingthresholds.html
type OpsWorksLayerAutoScalingThresholds struct {
	// CPUThreshold docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling-autoscalingthresholds.html#cfn-opsworks-layer-loadbasedautoscaling-autoscalingthresholds-cputhreshold
	CPUThreshold *IntegerExpr `json:"CpuThreshold,omitempty"`
	// IgnoreMetricsTime docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling-autoscalingthresholds.html#cfn-opsworks-layer-loadbasedautoscaling-autoscalingthresholds-ignoremetricstime
	IgnoreMetricsTime *IntegerExpr `json:"IgnoreMetricsTime,omitempty"`
	// InstanceCount docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling-autoscalingthresholds.html#cfn-opsworks-layer-loadbasedautoscaling-autoscalingthresholds-instancecount
	InstanceCount *IntegerExpr `json:"InstanceCount,omitempty"`
	// LoadThreshold docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling-autoscalingthresholds.html#cfn-opsworks-layer-loadbasedautoscaling-autoscalingthresholds-loadthreshold
	LoadThreshold *IntegerExpr `json:"LoadThreshold,omitempty"`
	// MemoryThreshold docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling-autoscalingthresholds.html#cfn-opsworks-layer-loadbasedautoscaling-autoscalingthresholds-memorythreshold
	MemoryThreshold *IntegerExpr `json:"MemoryThreshold,omitempty"`
	// ThresholdsWaitTime docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling-autoscalingthresholds.html#cfn-opsworks-layer-loadbasedautoscaling-autoscalingthresholds-thresholdwaittime
	ThresholdsWaitTime *IntegerExpr `json:"ThresholdsWaitTime,omitempty"`
}

// OpsWorksLayerAutoScalingThresholdsList represents a list of OpsWorksLayerAutoScalingThresholds
type OpsWorksLayerAutoScalingThresholdsList []OpsWorksLayerAutoScalingThresholds

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksLayerAutoScalingThresholdsList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksLayerAutoScalingThresholds{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksLayerAutoScalingThresholdsList{item}
		return nil
	}
	list := []OpsWorksLayerAutoScalingThresholds{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksLayerAutoScalingThresholdsList(list)
		return nil
	}
	return err
}

// OpsWorksLayerLifecycleEventConfiguration represents the AWS::OpsWorks::Layer.LifecycleEventConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration.html
type OpsWorksLayerLifecycleEventConfiguration struct {
	// ShutdownEventConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration.html#cfn-opsworks-layer-lifecycleconfiguration-shutdowneventconfiguration
	ShutdownEventConfiguration *OpsWorksLayerShutdownEventConfiguration `json:"ShutdownEventConfiguration,omitempty"`
}

// OpsWorksLayerLifecycleEventConfigurationList represents a list of OpsWorksLayerLifecycleEventConfiguration
type OpsWorksLayerLifecycleEventConfigurationList []OpsWorksLayerLifecycleEventConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksLayerLifecycleEventConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksLayerLifecycleEventConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksLayerLifecycleEventConfigurationList{item}
		return nil
	}
	list := []OpsWorksLayerLifecycleEventConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksLayerLifecycleEventConfigurationList(list)
		return nil
	}
	return err
}

// OpsWorksLayerLoadBasedAutoScaling represents the AWS::OpsWorks::Layer.LoadBasedAutoScaling CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling.html
type OpsWorksLayerLoadBasedAutoScaling struct {
	// DownScaling docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling.html#cfn-opsworks-layer-loadbasedautoscaling-downscaling
	DownScaling *OpsWorksLayerAutoScalingThresholds `json:"DownScaling,omitempty"`
	// Enable docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling.html#cfn-opsworks-layer-loadbasedautoscaling-enable
	Enable *BoolExpr `json:"Enable,omitempty"`
	// UpScaling docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling.html#cfn-opsworks-layer-loadbasedautoscaling-upscaling
	UpScaling *OpsWorksLayerAutoScalingThresholds `json:"UpScaling,omitempty"`
}

// OpsWorksLayerLoadBasedAutoScalingList represents a list of OpsWorksLayerLoadBasedAutoScaling
type OpsWorksLayerLoadBasedAutoScalingList []OpsWorksLayerLoadBasedAutoScaling

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksLayerLoadBasedAutoScalingList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksLayerLoadBasedAutoScaling{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksLayerLoadBasedAutoScalingList{item}
		return nil
	}
	list := []OpsWorksLayerLoadBasedAutoScaling{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksLayerLoadBasedAutoScalingList(list)
		return nil
	}
	return err
}

// OpsWorksLayerRecipes represents the AWS::OpsWorks::Layer.Recipes CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-recipes.html
type OpsWorksLayerRecipes struct {
	// Configure docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-recipes.html#cfn-opsworks-layer-customrecipes-configure
	Configure *StringListExpr `json:"Configure,omitempty"`
	// Deploy docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-recipes.html#cfn-opsworks-layer-customrecipes-deploy
	Deploy *StringListExpr `json:"Deploy,omitempty"`
	// Setup docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-recipes.html#cfn-opsworks-layer-customrecipes-setup
	Setup *StringListExpr `json:"Setup,omitempty"`
	// Shutdown docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-recipes.html#cfn-opsworks-layer-customrecipes-shutdown
	Shutdown *StringListExpr `json:"Shutdown,omitempty"`
	// Undeploy docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-recipes.html#cfn-opsworks-layer-customrecipes-undeploy
	Undeploy *StringListExpr `json:"Undeploy,omitempty"`
}

// OpsWorksLayerRecipesList represents a list of OpsWorksLayerRecipes
type OpsWorksLayerRecipesList []OpsWorksLayerRecipes

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksLayerRecipesList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksLayerRecipes{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksLayerRecipesList{item}
		return nil
	}
	list := []OpsWorksLayerRecipes{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksLayerRecipesList(list)
		return nil
	}
	return err
}

// OpsWorksLayerShutdownEventConfiguration represents the AWS::OpsWorks::Layer.ShutdownEventConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration-shutdowneventconfiguration.html
type OpsWorksLayerShutdownEventConfiguration struct {
	// DelayUntilElbConnectionsDrained docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration-shutdowneventconfiguration.html#cfn-opsworks-layer-lifecycleconfiguration-shutdowneventconfiguration-delayuntilelbconnectionsdrained
	DelayUntilElbConnectionsDrained *BoolExpr `json:"DelayUntilElbConnectionsDrained,omitempty"`
	// ExecutionTimeout docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration-shutdowneventconfiguration.html#cfn-opsworks-layer-lifecycleconfiguration-shutdowneventconfiguration-executiontimeout
	ExecutionTimeout *IntegerExpr `json:"ExecutionTimeout,omitempty"`
}

// OpsWorksLayerShutdownEventConfigurationList represents a list of OpsWorksLayerShutdownEventConfiguration
type OpsWorksLayerShutdownEventConfigurationList []OpsWorksLayerShutdownEventConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksLayerShutdownEventConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksLayerShutdownEventConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksLayerShutdownEventConfigurationList{item}
		return nil
	}
	list := []OpsWorksLayerShutdownEventConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksLayerShutdownEventConfigurationList(list)
		return nil
	}
	return err
}

// OpsWorksLayerVolumeConfiguration represents the AWS::OpsWorks::Layer.VolumeConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html
type OpsWorksLayerVolumeConfiguration struct {
	// Iops docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html#cfn-opsworks-layer-volconfig-iops
	Iops *IntegerExpr `json:"Iops,omitempty"`
	// MountPoint docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html#cfn-opsworks-layer-volconfig-mountpoint
	MountPoint *StringExpr `json:"MountPoint,omitempty"`
	// NumberOfDisks docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html#cfn-opsworks-layer-volconfig-numberofdisks
	NumberOfDisks *IntegerExpr `json:"NumberOfDisks,omitempty"`
	// RaidLevel docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html#cfn-opsworks-layer-volconfig-raidlevel
	RaidLevel *IntegerExpr `json:"RaidLevel,omitempty"`
	// Size docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html#cfn-opsworks-layer-volconfig-size
	Size *IntegerExpr `json:"Size,omitempty"`
	// VolumeType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html#cfn-opsworks-layer-volconfig-volumetype
	VolumeType *StringExpr `json:"VolumeType,omitempty"`
}

// OpsWorksLayerVolumeConfigurationList represents a list of OpsWorksLayerVolumeConfiguration
type OpsWorksLayerVolumeConfigurationList []OpsWorksLayerVolumeConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksLayerVolumeConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksLayerVolumeConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksLayerVolumeConfigurationList{item}
		return nil
	}
	list := []OpsWorksLayerVolumeConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksLayerVolumeConfigurationList(list)
		return nil
	}
	return err
}

// OpsWorksStackChefConfiguration represents the AWS::OpsWorks::Stack.ChefConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-chefconfiguration.html
type OpsWorksStackChefConfiguration struct {
	// BerkshelfVersion docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-chefconfiguration.html#cfn-opsworks-chefconfiguration-berkshelfversion
	BerkshelfVersion *StringExpr `json:"BerkshelfVersion,omitempty"`
	// ManageBerkshelf docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-chefconfiguration.html#cfn-opsworks-chefconfiguration-berkshelfversion
	ManageBerkshelf *BoolExpr `json:"ManageBerkshelf,omitempty"`
}

// OpsWorksStackChefConfigurationList represents a list of OpsWorksStackChefConfiguration
type OpsWorksStackChefConfigurationList []OpsWorksStackChefConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksStackChefConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksStackChefConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksStackChefConfigurationList{item}
		return nil
	}
	list := []OpsWorksStackChefConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksStackChefConfigurationList(list)
		return nil
	}
	return err
}

// OpsWorksStackElasticIP represents the AWS::OpsWorks::Stack.ElasticIp CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-elasticip.html
type OpsWorksStackElasticIP struct {
	// IP docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-elasticip.html#cfn-opsworks-stack-elasticip-ip
	IP *StringExpr `json:"Ip,omitempty" validate:"dive,required"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-elasticip.html#cfn-opsworks-stack-elasticip-name
	Name *StringExpr `json:"Name,omitempty"`
}

// OpsWorksStackElasticIPList represents a list of OpsWorksStackElasticIP
type OpsWorksStackElasticIPList []OpsWorksStackElasticIP

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksStackElasticIPList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksStackElasticIP{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksStackElasticIPList{item}
		return nil
	}
	list := []OpsWorksStackElasticIP{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksStackElasticIPList(list)
		return nil
	}
	return err
}

// OpsWorksStackRdsDbInstance represents the AWS::OpsWorks::Stack.RdsDbInstance CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-rdsdbinstance.html
type OpsWorksStackRdsDbInstance struct {
	// DbPassword docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-rdsdbinstance.html#cfn-opsworks-stack-rdsdbinstance-dbpassword
	DbPassword *StringExpr `json:"DbPassword,omitempty" validate:"dive,required"`
	// DbUser docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-rdsdbinstance.html#cfn-opsworks-stack-rdsdbinstance-dbuser
	DbUser *StringExpr `json:"DbUser,omitempty" validate:"dive,required"`
	// RdsDbInstanceArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-rdsdbinstance.html#cfn-opsworks-stack-rdsdbinstance-rdsdbinstancearn
	RdsDbInstanceArn *StringExpr `json:"RdsDbInstanceArn,omitempty" validate:"dive,required"`
}

// OpsWorksStackRdsDbInstanceList represents a list of OpsWorksStackRdsDbInstance
type OpsWorksStackRdsDbInstanceList []OpsWorksStackRdsDbInstance

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksStackRdsDbInstanceList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksStackRdsDbInstance{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksStackRdsDbInstanceList{item}
		return nil
	}
	list := []OpsWorksStackRdsDbInstance{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksStackRdsDbInstanceList(list)
		return nil
	}
	return err
}

// OpsWorksStackSource represents the AWS::OpsWorks::Stack.Source CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html
type OpsWorksStackSource struct {
	// Password docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html#cfn-opsworks-custcookbooksource-password
	Password *StringExpr `json:"Password,omitempty"`
	// Revision docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html#cfn-opsworks-custcookbooksource-revision
	Revision *StringExpr `json:"Revision,omitempty"`
	// SSHKey docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html#cfn-opsworks-custcookbooksource-sshkey
	SSHKey *StringExpr `json:"SshKey,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html#cfn-opsworks-custcookbooksource-type
	Type *StringExpr `json:"Type,omitempty"`
	// URL docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html#cfn-opsworks-custcookbooksource-url
	URL *StringExpr `json:"Url,omitempty"`
	// Username docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html#cfn-opsworks-custcookbooksource-username
	Username *StringExpr `json:"Username,omitempty"`
}

// OpsWorksStackSourceList represents a list of OpsWorksStackSource
type OpsWorksStackSourceList []OpsWorksStackSource

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksStackSourceList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksStackSource{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksStackSourceList{item}
		return nil
	}
	list := []OpsWorksStackSource{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksStackSourceList(list)
		return nil
	}
	return err
}

// OpsWorksStackStackConfigurationManager represents the AWS::OpsWorks::Stack.StackConfigurationManager CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-stackconfigmanager.html
type OpsWorksStackStackConfigurationManager struct {
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-stackconfigmanager.html#cfn-opsworks-configmanager-name
	Name *StringExpr `json:"Name,omitempty"`
	// Version docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-stackconfigmanager.html#cfn-opsworks-configmanager-version
	Version *StringExpr `json:"Version,omitempty"`
}

// OpsWorksStackStackConfigurationManagerList represents a list of OpsWorksStackStackConfigurationManager
type OpsWorksStackStackConfigurationManagerList []OpsWorksStackStackConfigurationManager

// UnmarshalJSON sets the object from the provided JSON representation
func (l *OpsWorksStackStackConfigurationManagerList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := OpsWorksStackStackConfigurationManager{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = OpsWorksStackStackConfigurationManagerList{item}
		return nil
	}
	list := []OpsWorksStackStackConfigurationManager{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = OpsWorksStackStackConfigurationManagerList(list)
		return nil
	}
	return err
}

// RDSDBSecurityGroupIngressProperty represents the AWS::RDS::DBSecurityGroup.Ingress CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group-rule.html
type RDSDBSecurityGroupIngressProperty struct {
	// CIDRIP docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group-rule.html#cfn-rds-securitygroup-cidrip
	CIDRIP *StringExpr `json:"CIDRIP,omitempty"`
	// EC2SecurityGroupID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group-rule.html#cfn-rds-securitygroup-ec2securitygroupid
	EC2SecurityGroupID *StringExpr `json:"EC2SecurityGroupId,omitempty"`
	// EC2SecurityGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group-rule.html#cfn-rds-securitygroup-ec2securitygroupname
	EC2SecurityGroupName *StringExpr `json:"EC2SecurityGroupName,omitempty"`
	// EC2SecurityGroupOwnerID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group-rule.html#cfn-rds-securitygroup-ec2securitygroupownerid
	EC2SecurityGroupOwnerID *StringExpr `json:"EC2SecurityGroupOwnerId,omitempty"`
}

// RDSDBSecurityGroupIngressPropertyList represents a list of RDSDBSecurityGroupIngressProperty
type RDSDBSecurityGroupIngressPropertyList []RDSDBSecurityGroupIngressProperty

// UnmarshalJSON sets the object from the provided JSON representation
func (l *RDSDBSecurityGroupIngressPropertyList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := RDSDBSecurityGroupIngressProperty{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = RDSDBSecurityGroupIngressPropertyList{item}
		return nil
	}
	list := []RDSDBSecurityGroupIngressProperty{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = RDSDBSecurityGroupIngressPropertyList(list)
		return nil
	}
	return err
}

// RDSOptionGroupOptionConfiguration represents the AWS::RDS::OptionGroup.OptionConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations.html
type RDSOptionGroupOptionConfiguration struct {
	// DBSecurityGroupMemberships docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations.html#cfn-rds-optiongroup-optionconfigurations-dbsecuritygroupmemberships
	DBSecurityGroupMemberships *StringListExpr `json:"DBSecurityGroupMemberships,omitempty"`
	// OptionName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations.html#cfn-rds-optiongroup-optionconfigurations-optionname
	OptionName *StringExpr `json:"OptionName,omitempty" validate:"dive,required"`
	// OptionSettings docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations.html#cfn-rds-optiongroup-optionconfigurations-optionsettings
	OptionSettings *RDSOptionGroupOptionSetting `json:"OptionSettings,omitempty"`
	// Port docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations.html#cfn-rds-optiongroup-optionconfigurations-port
	Port *IntegerExpr `json:"Port,omitempty"`
	// VPCSecurityGroupMemberships docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations.html#cfn-rds-optiongroup-optionconfigurations-vpcsecuritygroupmemberships
	VPCSecurityGroupMemberships *StringListExpr `json:"VpcSecurityGroupMemberships,omitempty"`
}

// RDSOptionGroupOptionConfigurationList represents a list of RDSOptionGroupOptionConfiguration
type RDSOptionGroupOptionConfigurationList []RDSOptionGroupOptionConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *RDSOptionGroupOptionConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := RDSOptionGroupOptionConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = RDSOptionGroupOptionConfigurationList{item}
		return nil
	}
	list := []RDSOptionGroupOptionConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = RDSOptionGroupOptionConfigurationList(list)
		return nil
	}
	return err
}

// RDSOptionGroupOptionSetting represents the AWS::RDS::OptionGroup.OptionSetting CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations-optionsettings.html
type RDSOptionGroupOptionSetting struct {
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations-optionsettings.html#cfn-rds-optiongroup-optionconfigurations-optionsettings-name
	Name *StringExpr `json:"Name,omitempty"`
	// Value docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations-optionsettings.html#cfn-rds-optiongroup-optionconfigurations-optionsettings-value
	Value *StringExpr `json:"Value,omitempty"`
}

// RDSOptionGroupOptionSettingList represents a list of RDSOptionGroupOptionSetting
type RDSOptionGroupOptionSettingList []RDSOptionGroupOptionSetting

// UnmarshalJSON sets the object from the provided JSON representation
func (l *RDSOptionGroupOptionSettingList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := RDSOptionGroupOptionSetting{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = RDSOptionGroupOptionSettingList{item}
		return nil
	}
	list := []RDSOptionGroupOptionSetting{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = RDSOptionGroupOptionSettingList(list)
		return nil
	}
	return err
}

// RedshiftClusterParameterGroupParameter represents the AWS::Redshift::ClusterParameterGroup.Parameter CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-property-redshift-clusterparametergroup-parameter.html
type RedshiftClusterParameterGroupParameter struct {
	// ParameterName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-property-redshift-clusterparametergroup-parameter.html#cfn-redshift-clusterparametergroup-parameter-parametername
	ParameterName *StringExpr `json:"ParameterName,omitempty" validate:"dive,required"`
	// ParameterValue docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-property-redshift-clusterparametergroup-parameter.html#cfn-redshift-clusterparametergroup-parameter-parametervalue
	ParameterValue *StringExpr `json:"ParameterValue,omitempty" validate:"dive,required"`
}

// RedshiftClusterParameterGroupParameterList represents a list of RedshiftClusterParameterGroupParameter
type RedshiftClusterParameterGroupParameterList []RedshiftClusterParameterGroupParameter

// UnmarshalJSON sets the object from the provided JSON representation
func (l *RedshiftClusterParameterGroupParameterList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := RedshiftClusterParameterGroupParameter{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = RedshiftClusterParameterGroupParameterList{item}
		return nil
	}
	list := []RedshiftClusterParameterGroupParameter{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = RedshiftClusterParameterGroupParameterList(list)
		return nil
	}
	return err
}

// Route53HealthCheckAlarmIDentifier represents the AWS::Route53::HealthCheck.AlarmIdentifier CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-alarmidentifier.html
type Route53HealthCheckAlarmIDentifier struct {
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-alarmidentifier.html#cfn-route53-healthcheck-alarmidentifier-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// Region docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-alarmidentifier.html#cfn-route53-healthcheck-alarmidentifier-region
	Region *StringExpr `json:"Region,omitempty" validate:"dive,required"`
}

// Route53HealthCheckAlarmIDentifierList represents a list of Route53HealthCheckAlarmIDentifier
type Route53HealthCheckAlarmIDentifierList []Route53HealthCheckAlarmIDentifier

// UnmarshalJSON sets the object from the provided JSON representation
func (l *Route53HealthCheckAlarmIDentifierList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := Route53HealthCheckAlarmIDentifier{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = Route53HealthCheckAlarmIDentifierList{item}
		return nil
	}
	list := []Route53HealthCheckAlarmIDentifier{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = Route53HealthCheckAlarmIDentifierList(list)
		return nil
	}
	return err
}

// Route53HealthCheckHealthCheckConfig represents the AWS::Route53::HealthCheck.HealthCheckConfig CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html
type Route53HealthCheckHealthCheckConfig struct {
	// AlarmIDentifier docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-alarmidentifier
	AlarmIDentifier *Route53HealthCheckAlarmIDentifier `json:"AlarmIdentifier,omitempty"`
	// ChildHealthChecks docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-childhealthchecks
	ChildHealthChecks *StringListExpr `json:"ChildHealthChecks,omitempty"`
	// EnableSNI docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-enablesni
	EnableSNI *BoolExpr `json:"EnableSNI,omitempty"`
	// FailureThreshold docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-failurethreshold
	FailureThreshold *IntegerExpr `json:"FailureThreshold,omitempty"`
	// FullyQualifiedDomainName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-fullyqualifieddomainname
	FullyQualifiedDomainName *StringExpr `json:"FullyQualifiedDomainName,omitempty"`
	// HealthThreshold docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-healththreshold
	HealthThreshold *IntegerExpr `json:"HealthThreshold,omitempty"`
	// IPAddress docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-ipaddress
	IPAddress *StringExpr `json:"IPAddress,omitempty"`
	// InsufficientDataHealthStatus docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-insufficientdatahealthstatus
	InsufficientDataHealthStatus *StringExpr `json:"InsufficientDataHealthStatus,omitempty"`
	// Inverted docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-inverted
	Inverted *BoolExpr `json:"Inverted,omitempty"`
	// MeasureLatency docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-measurelatency
	MeasureLatency *BoolExpr `json:"MeasureLatency,omitempty"`
	// Port docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-port
	Port *IntegerExpr `json:"Port,omitempty"`
	// RequestInterval docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-requestinterval
	RequestInterval *IntegerExpr `json:"RequestInterval,omitempty"`
	// ResourcePath docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-resourcepath
	ResourcePath *StringExpr `json:"ResourcePath,omitempty"`
	// SearchString docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-searchstring
	SearchString *StringExpr `json:"SearchString,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
}

// Route53HealthCheckHealthCheckConfigList represents a list of Route53HealthCheckHealthCheckConfig
type Route53HealthCheckHealthCheckConfigList []Route53HealthCheckHealthCheckConfig

// UnmarshalJSON sets the object from the provided JSON representation
func (l *Route53HealthCheckHealthCheckConfigList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := Route53HealthCheckHealthCheckConfig{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = Route53HealthCheckHealthCheckConfigList{item}
		return nil
	}
	list := []Route53HealthCheckHealthCheckConfig{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = Route53HealthCheckHealthCheckConfigList(list)
		return nil
	}
	return err
}

// Route53HealthCheckHealthCheckTag represents the AWS::Route53::HealthCheck.HealthCheckTag CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthchecktag.html
type Route53HealthCheckHealthCheckTag struct {
	// Key docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthchecktag.html#cfn-route53-healthchecktags-key
	Key *StringExpr `json:"Key,omitempty" validate:"dive,required"`
	// Value docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthchecktag.html#cfn-route53-healthchecktags-value
	Value *StringExpr `json:"Value,omitempty" validate:"dive,required"`
}

// Route53HealthCheckHealthCheckTagList represents a list of Route53HealthCheckHealthCheckTag
type Route53HealthCheckHealthCheckTagList []Route53HealthCheckHealthCheckTag

// UnmarshalJSON sets the object from the provided JSON representation
func (l *Route53HealthCheckHealthCheckTagList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := Route53HealthCheckHealthCheckTag{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = Route53HealthCheckHealthCheckTagList{item}
		return nil
	}
	list := []Route53HealthCheckHealthCheckTag{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = Route53HealthCheckHealthCheckTagList(list)
		return nil
	}
	return err
}

// Route53HostedZoneHostedZoneConfig represents the AWS::Route53::HostedZone.HostedZoneConfig CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzoneconfig.html
type Route53HostedZoneHostedZoneConfig struct {
	// Comment docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzoneconfig.html#cfn-route53-hostedzone-hostedzoneconfig-comment
	Comment *StringExpr `json:"Comment,omitempty"`
}

// Route53HostedZoneHostedZoneConfigList represents a list of Route53HostedZoneHostedZoneConfig
type Route53HostedZoneHostedZoneConfigList []Route53HostedZoneHostedZoneConfig

// UnmarshalJSON sets the object from the provided JSON representation
func (l *Route53HostedZoneHostedZoneConfigList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := Route53HostedZoneHostedZoneConfig{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = Route53HostedZoneHostedZoneConfigList{item}
		return nil
	}
	list := []Route53HostedZoneHostedZoneConfig{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = Route53HostedZoneHostedZoneConfigList(list)
		return nil
	}
	return err
}

// Route53HostedZoneHostedZoneTag represents the AWS::Route53::HostedZone.HostedZoneTag CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzonetags.html
type Route53HostedZoneHostedZoneTag struct {
	// Key docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzonetags.html#cfn-route53-hostedzonetags-key
	Key *StringExpr `json:"Key,omitempty" validate:"dive,required"`
	// Value docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzonetags.html#cfn-route53-hostedzonetags-value
	Value *StringExpr `json:"Value,omitempty" validate:"dive,required"`
}

// Route53HostedZoneHostedZoneTagList represents a list of Route53HostedZoneHostedZoneTag
type Route53HostedZoneHostedZoneTagList []Route53HostedZoneHostedZoneTag

// UnmarshalJSON sets the object from the provided JSON representation
func (l *Route53HostedZoneHostedZoneTagList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := Route53HostedZoneHostedZoneTag{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = Route53HostedZoneHostedZoneTagList{item}
		return nil
	}
	list := []Route53HostedZoneHostedZoneTag{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = Route53HostedZoneHostedZoneTagList(list)
		return nil
	}
	return err
}

// Route53HostedZoneVPC represents the AWS::Route53::HostedZone.VPC CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone-hostedzonevpcs.html
type Route53HostedZoneVPC struct {
	// VPCID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone-hostedzonevpcs.html#cfn-route53-hostedzone-hostedzonevpcs-vpcid
	VPCID *StringExpr `json:"VPCId,omitempty" validate:"dive,required"`
	// VPCRegion docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone-hostedzonevpcs.html#cfn-route53-hostedzone-hostedzonevpcs-vpcregion
	VPCRegion *StringExpr `json:"VPCRegion,omitempty" validate:"dive,required"`
}

// Route53HostedZoneVPCList represents a list of Route53HostedZoneVPC
type Route53HostedZoneVPCList []Route53HostedZoneVPC

// UnmarshalJSON sets the object from the provided JSON representation
func (l *Route53HostedZoneVPCList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := Route53HostedZoneVPC{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = Route53HostedZoneVPCList{item}
		return nil
	}
	list := []Route53HostedZoneVPC{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = Route53HostedZoneVPCList(list)
		return nil
	}
	return err
}

// Route53RecordSetAliasTarget represents the AWS::Route53::RecordSet.AliasTarget CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-aliastarget.html
type Route53RecordSetAliasTarget struct {
	// DNSName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-aliastarget.html#cfn-route53-aliastarget-dnshostname
	DNSName *StringExpr `json:"DNSName,omitempty" validate:"dive,required"`
	// EvaluateTargetHealth docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-aliastarget.html#cfn-route53-aliastarget-evaluatetargethealth
	EvaluateTargetHealth *BoolExpr `json:"EvaluateTargetHealth,omitempty"`
	// HostedZoneID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-aliastarget.html#cfn-route53-aliastarget-hostedzoneid
	HostedZoneID *StringExpr `json:"HostedZoneId,omitempty" validate:"dive,required"`
}

// Route53RecordSetAliasTargetList represents a list of Route53RecordSetAliasTarget
type Route53RecordSetAliasTargetList []Route53RecordSetAliasTarget

// UnmarshalJSON sets the object from the provided JSON representation
func (l *Route53RecordSetAliasTargetList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := Route53RecordSetAliasTarget{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = Route53RecordSetAliasTargetList{item}
		return nil
	}
	list := []Route53RecordSetAliasTarget{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = Route53RecordSetAliasTargetList(list)
		return nil
	}
	return err
}

// Route53RecordSetGeoLocation represents the AWS::Route53::RecordSet.GeoLocation CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html
type Route53RecordSetGeoLocation struct {
	// ContinentCode docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html#cfn-route53-recordset-geolocation-continentcode
	ContinentCode *StringExpr `json:"ContinentCode,omitempty"`
	// CountryCode docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html#cfn-route53-recordset-geolocation-countrycode
	CountryCode *StringExpr `json:"CountryCode,omitempty"`
	// SubdivisionCode docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html#cfn-route53-recordset-geolocation-subdivisioncode
	SubdivisionCode *StringExpr `json:"SubdivisionCode,omitempty"`
}

// Route53RecordSetGeoLocationList represents a list of Route53RecordSetGeoLocation
type Route53RecordSetGeoLocationList []Route53RecordSetGeoLocation

// UnmarshalJSON sets the object from the provided JSON representation
func (l *Route53RecordSetGeoLocationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := Route53RecordSetGeoLocation{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = Route53RecordSetGeoLocationList{item}
		return nil
	}
	list := []Route53RecordSetGeoLocation{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = Route53RecordSetGeoLocationList(list)
		return nil
	}
	return err
}

// Route53RecordSetGroupAliasTarget represents the AWS::Route53::RecordSetGroup.AliasTarget CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-aliastarget.html
type Route53RecordSetGroupAliasTarget struct {
	// DNSName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-aliastarget.html#cfn-route53-aliastarget-dnshostname
	DNSName *StringExpr `json:"DNSName,omitempty" validate:"dive,required"`
	// EvaluateTargetHealth docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-aliastarget.html#cfn-route53-aliastarget-evaluatetargethealth
	EvaluateTargetHealth *BoolExpr `json:"EvaluateTargetHealth,omitempty"`
	// HostedZoneID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-aliastarget.html#cfn-route53-aliastarget-hostedzoneid
	HostedZoneID *StringExpr `json:"HostedZoneId,omitempty" validate:"dive,required"`
}

// Route53RecordSetGroupAliasTargetList represents a list of Route53RecordSetGroupAliasTarget
type Route53RecordSetGroupAliasTargetList []Route53RecordSetGroupAliasTarget

// UnmarshalJSON sets the object from the provided JSON representation
func (l *Route53RecordSetGroupAliasTargetList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := Route53RecordSetGroupAliasTarget{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = Route53RecordSetGroupAliasTargetList{item}
		return nil
	}
	list := []Route53RecordSetGroupAliasTarget{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = Route53RecordSetGroupAliasTargetList(list)
		return nil
	}
	return err
}

// Route53RecordSetGroupGeoLocation represents the AWS::Route53::RecordSetGroup.GeoLocation CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html
type Route53RecordSetGroupGeoLocation struct {
	// ContinentCode docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html#cfn-route53-recordsetgroup-geolocation-continentcode
	ContinentCode *StringExpr `json:"ContinentCode,omitempty"`
	// CountryCode docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html#cfn-route53-recordset-geolocation-countrycode
	CountryCode *StringExpr `json:"CountryCode,omitempty"`
	// SubdivisionCode docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html#cfn-route53-recordset-geolocation-subdivisioncode
	SubdivisionCode *StringExpr `json:"SubdivisionCode,omitempty"`
}

// Route53RecordSetGroupGeoLocationList represents a list of Route53RecordSetGroupGeoLocation
type Route53RecordSetGroupGeoLocationList []Route53RecordSetGroupGeoLocation

// UnmarshalJSON sets the object from the provided JSON representation
func (l *Route53RecordSetGroupGeoLocationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := Route53RecordSetGroupGeoLocation{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = Route53RecordSetGroupGeoLocationList{item}
		return nil
	}
	list := []Route53RecordSetGroupGeoLocation{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = Route53RecordSetGroupGeoLocationList(list)
		return nil
	}
	return err
}

// Route53RecordSetGroupRecordSet represents the AWS::Route53::RecordSetGroup.RecordSet CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html
type Route53RecordSetGroupRecordSet struct {
	// AliasTarget docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-aliastarget
	AliasTarget *Route53RecordSetGroupAliasTarget `json:"AliasTarget,omitempty"`
	// Comment docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-comment
	Comment *StringExpr `json:"Comment,omitempty"`
	// Failover docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-failover
	Failover *StringExpr `json:"Failover,omitempty"`
	// GeoLocation docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-geolocation
	GeoLocation *Route53RecordSetGroupGeoLocation `json:"GeoLocation,omitempty"`
	// HealthCheckID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-healthcheckid
	HealthCheckID *StringExpr `json:"HealthCheckId,omitempty"`
	// HostedZoneID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-hostedzoneid
	HostedZoneID *StringExpr `json:"HostedZoneId,omitempty"`
	// HostedZoneName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-hostedzonename
	HostedZoneName *StringExpr `json:"HostedZoneName,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// Region docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-region
	Region *StringExpr `json:"Region,omitempty"`
	// ResourceRecords docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-resourcerecords
	ResourceRecords *StringListExpr `json:"ResourceRecords,omitempty"`
	// SetIDentifier docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-setidentifier
	SetIDentifier *StringExpr `json:"SetIdentifier,omitempty"`
	// TTL docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-ttl
	TTL *StringExpr `json:"TTL,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
	// Weight docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-weight
	Weight *IntegerExpr `json:"Weight,omitempty"`
}

// Route53RecordSetGroupRecordSetList represents a list of Route53RecordSetGroupRecordSet
type Route53RecordSetGroupRecordSetList []Route53RecordSetGroupRecordSet

// UnmarshalJSON sets the object from the provided JSON representation
func (l *Route53RecordSetGroupRecordSetList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := Route53RecordSetGroupRecordSet{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = Route53RecordSetGroupRecordSetList{item}
		return nil
	}
	list := []Route53RecordSetGroupRecordSet{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = Route53RecordSetGroupRecordSetList(list)
		return nil
	}
	return err
}

// S3BucketCorsConfiguration represents the AWS::S3::Bucket.CorsConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors.html
type S3BucketCorsConfiguration struct {
	// CorsRules docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors.html#cfn-s3-bucket-cors-corsrule
	CorsRules *S3BucketCorsRuleList `json:"CorsRules,omitempty" validate:"dive,required"`
}

// S3BucketCorsConfigurationList represents a list of S3BucketCorsConfiguration
type S3BucketCorsConfigurationList []S3BucketCorsConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3BucketCorsConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3BucketCorsConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3BucketCorsConfigurationList{item}
		return nil
	}
	list := []S3BucketCorsConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3BucketCorsConfigurationList(list)
		return nil
	}
	return err
}

// S3BucketCorsRule represents the AWS::S3::Bucket.CorsRule CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors-corsrule.html
type S3BucketCorsRule struct {
	// AllowedHeaders docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors-corsrule.html#cfn-s3-bucket-cors-corsrule-allowedheaders
	AllowedHeaders *StringListExpr `json:"AllowedHeaders,omitempty"`
	// AllowedMethods docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors-corsrule.html#cfn-s3-bucket-cors-corsrule-allowedmethods
	AllowedMethods *StringListExpr `json:"AllowedMethods,omitempty" validate:"dive,required"`
	// AllowedOrigins docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors-corsrule.html#cfn-s3-bucket-cors-corsrule-allowedorigins
	AllowedOrigins *StringListExpr `json:"AllowedOrigins,omitempty" validate:"dive,required"`
	// ExposedHeaders docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors-corsrule.html#cfn-s3-bucket-cors-corsrule-exposedheaders
	ExposedHeaders *StringListExpr `json:"ExposedHeaders,omitempty"`
	// ID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors-corsrule.html#cfn-s3-bucket-cors-corsrule-id
	ID *StringExpr `json:"Id,omitempty"`
	// MaxAge docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors-corsrule.html#cfn-s3-bucket-cors-corsrule-maxage
	MaxAge *IntegerExpr `json:"MaxAge,omitempty"`
}

// S3BucketCorsRuleList represents a list of S3BucketCorsRule
type S3BucketCorsRuleList []S3BucketCorsRule

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3BucketCorsRuleList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3BucketCorsRule{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3BucketCorsRuleList{item}
		return nil
	}
	list := []S3BucketCorsRule{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3BucketCorsRuleList(list)
		return nil
	}
	return err
}

// S3BucketFilterRule represents the AWS::S3::Bucket.FilterRule CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter-s3key-rules.html
type S3BucketFilterRule struct {
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter-s3key-rules.html#cfn-s3-bucket-notificationconfiguraiton-config-filter-s3key-rules-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// Value docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter-s3key-rules.html#cfn-s3-bucket-notificationconfiguraiton-config-filter-s3key-rules-value
	Value *StringExpr `json:"Value,omitempty" validate:"dive,required"`
}

// S3BucketFilterRuleList represents a list of S3BucketFilterRule
type S3BucketFilterRuleList []S3BucketFilterRule

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3BucketFilterRuleList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3BucketFilterRule{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3BucketFilterRuleList{item}
		return nil
	}
	list := []S3BucketFilterRule{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3BucketFilterRuleList(list)
		return nil
	}
	return err
}

// S3BucketLambdaConfiguration represents the AWS::S3::Bucket.LambdaConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-lambdaconfig.html
type S3BucketLambdaConfiguration struct {
	// Event docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-lambdaconfig.html#cfn-s3-bucket-notificationconfig-lambdaconfig-event
	Event *StringExpr `json:"Event,omitempty" validate:"dive,required"`
	// Filter docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-lambdaconfig.html#cfn-s3-bucket-notificationconfig-lambdaconfig-filter
	Filter *S3BucketNotificationFilter `json:"Filter,omitempty"`
	// Function docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-lambdaconfig.html#cfn-s3-bucket-notificationconfig-lambdaconfig-function
	Function *StringExpr `json:"Function,omitempty" validate:"dive,required"`
}

// S3BucketLambdaConfigurationList represents a list of S3BucketLambdaConfiguration
type S3BucketLambdaConfigurationList []S3BucketLambdaConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3BucketLambdaConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3BucketLambdaConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3BucketLambdaConfigurationList{item}
		return nil
	}
	list := []S3BucketLambdaConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3BucketLambdaConfigurationList(list)
		return nil
	}
	return err
}

// S3BucketLifecycleConfiguration represents the AWS::S3::Bucket.LifecycleConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig.html
type S3BucketLifecycleConfiguration struct {
	// Rules docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig.html#cfn-s3-bucket-lifecycleconfig-rules
	Rules *S3BucketRuleList `json:"Rules,omitempty" validate:"dive,required"`
}

// S3BucketLifecycleConfigurationList represents a list of S3BucketLifecycleConfiguration
type S3BucketLifecycleConfigurationList []S3BucketLifecycleConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3BucketLifecycleConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3BucketLifecycleConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3BucketLifecycleConfigurationList{item}
		return nil
	}
	list := []S3BucketLifecycleConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3BucketLifecycleConfigurationList(list)
		return nil
	}
	return err
}

// S3BucketLoggingConfiguration represents the AWS::S3::Bucket.LoggingConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-loggingconfig.html
type S3BucketLoggingConfiguration struct {
	// DestinationBucketName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-loggingconfig.html#cfn-s3-bucket-loggingconfig-destinationbucketname
	DestinationBucketName *StringExpr `json:"DestinationBucketName,omitempty"`
	// LogFilePrefix docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-loggingconfig.html#cfn-s3-bucket-loggingconfig-logfileprefix
	LogFilePrefix *StringExpr `json:"LogFilePrefix,omitempty"`
}

// S3BucketLoggingConfigurationList represents a list of S3BucketLoggingConfiguration
type S3BucketLoggingConfigurationList []S3BucketLoggingConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3BucketLoggingConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3BucketLoggingConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3BucketLoggingConfigurationList{item}
		return nil
	}
	list := []S3BucketLoggingConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3BucketLoggingConfigurationList(list)
		return nil
	}
	return err
}

// S3BucketNoncurrentVersionTransition represents the AWS::S3::Bucket.NoncurrentVersionTransition CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-noncurrentversiontransition.html
type S3BucketNoncurrentVersionTransition struct {
	// StorageClass docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-noncurrentversiontransition.html#cfn-s3-bucket-lifecycleconfig-rule-noncurrentversiontransition-storageclass
	StorageClass *StringExpr `json:"StorageClass,omitempty" validate:"dive,required"`
	// TransitionInDays docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-noncurrentversiontransition.html#cfn-s3-bucket-lifecycleconfig-rule-noncurrentversiontransition-transitionindays
	TransitionInDays *IntegerExpr `json:"TransitionInDays,omitempty" validate:"dive,required"`
}

// S3BucketNoncurrentVersionTransitionList represents a list of S3BucketNoncurrentVersionTransition
type S3BucketNoncurrentVersionTransitionList []S3BucketNoncurrentVersionTransition

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3BucketNoncurrentVersionTransitionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3BucketNoncurrentVersionTransition{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3BucketNoncurrentVersionTransitionList{item}
		return nil
	}
	list := []S3BucketNoncurrentVersionTransition{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3BucketNoncurrentVersionTransitionList(list)
		return nil
	}
	return err
}

// S3BucketNotificationConfiguration represents the AWS::S3::Bucket.NotificationConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig.html
type S3BucketNotificationConfiguration struct {
	// LambdaConfigurations docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig.html#cfn-s3-bucket-notificationconfig-lambdaconfig
	LambdaConfigurations *S3BucketLambdaConfigurationList `json:"LambdaConfigurations,omitempty"`
	// QueueConfigurations docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig.html#cfn-s3-bucket-notificationconfig-queueconfig
	QueueConfigurations *S3BucketQueueConfigurationList `json:"QueueConfigurations,omitempty"`
	// TopicConfigurations docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig.html#cfn-s3-bucket-notificationconfig-topicconfig
	TopicConfigurations *S3BucketTopicConfigurationList `json:"TopicConfigurations,omitempty"`
}

// S3BucketNotificationConfigurationList represents a list of S3BucketNotificationConfiguration
type S3BucketNotificationConfigurationList []S3BucketNotificationConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3BucketNotificationConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3BucketNotificationConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3BucketNotificationConfigurationList{item}
		return nil
	}
	list := []S3BucketNotificationConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3BucketNotificationConfigurationList(list)
		return nil
	}
	return err
}

// S3BucketNotificationFilter represents the AWS::S3::Bucket.NotificationFilter CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter.html
type S3BucketNotificationFilter struct {
	// S3Key docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter.html#cfn-s3-bucket-notificationconfiguraiton-config-filter-s3key
	S3Key *S3BucketS3KeyFilter `json:"S3Key,omitempty" validate:"dive,required"`
}

// S3BucketNotificationFilterList represents a list of S3BucketNotificationFilter
type S3BucketNotificationFilterList []S3BucketNotificationFilter

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3BucketNotificationFilterList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3BucketNotificationFilter{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3BucketNotificationFilterList{item}
		return nil
	}
	list := []S3BucketNotificationFilter{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3BucketNotificationFilterList(list)
		return nil
	}
	return err
}

// S3BucketQueueConfiguration represents the AWS::S3::Bucket.QueueConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-queueconfig.html
type S3BucketQueueConfiguration struct {
	// Event docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-queueconfig.html#cfn-s3-bucket-notificationconfig-queueconfig-event
	Event *StringExpr `json:"Event,omitempty" validate:"dive,required"`
	// Filter docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-queueconfig.html#cfn-s3-bucket-notificationconfig-queueconfig-filter
	Filter *S3BucketNotificationFilter `json:"Filter,omitempty"`
	// Queue docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-queueconfig.html#cfn-s3-bucket-notificationconfig-queueconfig-queue
	Queue *StringExpr `json:"Queue,omitempty" validate:"dive,required"`
}

// S3BucketQueueConfigurationList represents a list of S3BucketQueueConfiguration
type S3BucketQueueConfigurationList []S3BucketQueueConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3BucketQueueConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3BucketQueueConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3BucketQueueConfigurationList{item}
		return nil
	}
	list := []S3BucketQueueConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3BucketQueueConfigurationList(list)
		return nil
	}
	return err
}

// S3BucketRedirectAllRequestsTo represents the AWS::S3::Bucket.RedirectAllRequestsTo CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-redirectallrequeststo.html
type S3BucketRedirectAllRequestsTo struct {
	// HostName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-redirectallrequeststo.html#cfn-s3-websiteconfiguration-redirectallrequeststo-hostname
	HostName *StringExpr `json:"HostName,omitempty" validate:"dive,required"`
	// Protocol docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-redirectallrequeststo.html#cfn-s3-websiteconfiguration-redirectallrequeststo-protocol
	Protocol *StringExpr `json:"Protocol,omitempty"`
}

// S3BucketRedirectAllRequestsToList represents a list of S3BucketRedirectAllRequestsTo
type S3BucketRedirectAllRequestsToList []S3BucketRedirectAllRequestsTo

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3BucketRedirectAllRequestsToList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3BucketRedirectAllRequestsTo{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3BucketRedirectAllRequestsToList{item}
		return nil
	}
	list := []S3BucketRedirectAllRequestsTo{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3BucketRedirectAllRequestsToList(list)
		return nil
	}
	return err
}

// S3BucketRedirectRule represents the AWS::S3::Bucket.RedirectRule CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-redirectrule.html
type S3BucketRedirectRule struct {
	// HostName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-redirectrule.html#cfn-s3-websiteconfiguration-redirectrule-hostname
	HostName *StringExpr `json:"HostName,omitempty"`
	// HTTPRedirectCode docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-redirectrule.html#cfn-s3-websiteconfiguration-redirectrule-httpredirectcode
	HTTPRedirectCode *StringExpr `json:"HttpRedirectCode,omitempty"`
	// Protocol docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-redirectrule.html#cfn-s3-websiteconfiguration-redirectrule-protocol
	Protocol *StringExpr `json:"Protocol,omitempty"`
	// ReplaceKeyPrefixWith docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-redirectrule.html#cfn-s3-websiteconfiguration-redirectrule-replacekeyprefixwith
	ReplaceKeyPrefixWith *StringExpr `json:"ReplaceKeyPrefixWith,omitempty"`
	// ReplaceKeyWith docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-redirectrule.html#cfn-s3-websiteconfiguration-redirectrule-replacekeywith
	ReplaceKeyWith *StringExpr `json:"ReplaceKeyWith,omitempty"`
}

// S3BucketRedirectRuleList represents a list of S3BucketRedirectRule
type S3BucketRedirectRuleList []S3BucketRedirectRule

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3BucketRedirectRuleList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3BucketRedirectRule{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3BucketRedirectRuleList{item}
		return nil
	}
	list := []S3BucketRedirectRule{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3BucketRedirectRuleList(list)
		return nil
	}
	return err
}

// S3BucketReplicationConfiguration represents the AWS::S3::Bucket.ReplicationConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration.html
type S3BucketReplicationConfiguration struct {
	// Role docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration.html#cfn-s3-bucket-replicationconfiguration-role
	Role *StringExpr `json:"Role,omitempty" validate:"dive,required"`
	// Rules docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration.html#cfn-s3-bucket-replicationconfiguration-rules
	Rules *S3BucketReplicationRuleList `json:"Rules,omitempty" validate:"dive,required"`
}

// S3BucketReplicationConfigurationList represents a list of S3BucketReplicationConfiguration
type S3BucketReplicationConfigurationList []S3BucketReplicationConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3BucketReplicationConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3BucketReplicationConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3BucketReplicationConfigurationList{item}
		return nil
	}
	list := []S3BucketReplicationConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3BucketReplicationConfigurationList(list)
		return nil
	}
	return err
}

// S3BucketReplicationDestination represents the AWS::S3::Bucket.ReplicationDestination CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules-destination.html
type S3BucketReplicationDestination struct {
	// Bucket docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules-destination.html#cfn-s3-bucket-replicationconfiguration-rules-destination-bucket
	Bucket *StringExpr `json:"Bucket,omitempty" validate:"dive,required"`
	// StorageClass docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules-destination.html#cfn-s3-bucket-replicationconfiguration-rules-destination-storageclass
	StorageClass *StringExpr `json:"StorageClass,omitempty"`
}

// S3BucketReplicationDestinationList represents a list of S3BucketReplicationDestination
type S3BucketReplicationDestinationList []S3BucketReplicationDestination

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3BucketReplicationDestinationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3BucketReplicationDestination{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3BucketReplicationDestinationList{item}
		return nil
	}
	list := []S3BucketReplicationDestination{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3BucketReplicationDestinationList(list)
		return nil
	}
	return err
}

// S3BucketReplicationRule represents the AWS::S3::Bucket.ReplicationRule CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules.html
type S3BucketReplicationRule struct {
	// Destination docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules.html#cfn-s3-bucket-replicationconfiguration-rules-destination
	Destination *S3BucketReplicationDestination `json:"Destination,omitempty" validate:"dive,required"`
	// ID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules.html#cfn-s3-bucket-replicationconfiguration-rules-id
	ID *StringExpr `json:"Id,omitempty"`
	// Prefix docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules.html#cfn-s3-bucket-replicationconfiguration-rules-prefix
	Prefix *StringExpr `json:"Prefix,omitempty" validate:"dive,required"`
	// Status docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules.html#cfn-s3-bucket-replicationconfiguration-rules-status
	Status *StringExpr `json:"Status,omitempty" validate:"dive,required"`
}

// S3BucketReplicationRuleList represents a list of S3BucketReplicationRule
type S3BucketReplicationRuleList []S3BucketReplicationRule

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3BucketReplicationRuleList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3BucketReplicationRule{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3BucketReplicationRuleList{item}
		return nil
	}
	list := []S3BucketReplicationRule{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3BucketReplicationRuleList(list)
		return nil
	}
	return err
}

// S3BucketRoutingRule represents the AWS::S3::Bucket.RoutingRule CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules.html
type S3BucketRoutingRule struct {
	// RedirectRule docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules.html#cfn-s3-websiteconfiguration-routingrules-redirectrule
	RedirectRule *S3BucketRedirectRule `json:"RedirectRule,omitempty" validate:"dive,required"`
	// RoutingRuleCondition docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules.html#cfn-s3-websiteconfiguration-routingrules-routingrulecondition
	RoutingRuleCondition *S3BucketRoutingRuleCondition `json:"RoutingRuleCondition,omitempty"`
}

// S3BucketRoutingRuleList represents a list of S3BucketRoutingRule
type S3BucketRoutingRuleList []S3BucketRoutingRule

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3BucketRoutingRuleList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3BucketRoutingRule{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3BucketRoutingRuleList{item}
		return nil
	}
	list := []S3BucketRoutingRule{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3BucketRoutingRuleList(list)
		return nil
	}
	return err
}

// S3BucketRoutingRuleCondition represents the AWS::S3::Bucket.RoutingRuleCondition CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-routingrulecondition.html
type S3BucketRoutingRuleCondition struct {
	// HTTPErrorCodeReturnedEquals docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-routingrulecondition.html#cfn-s3-websiteconfiguration-routingrules-routingrulecondition-httperrorcodereturnedequals
	HTTPErrorCodeReturnedEquals *StringExpr `json:"HttpErrorCodeReturnedEquals,omitempty"`
	// KeyPrefixEquals docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-routingrulecondition.html#cfn-s3-websiteconfiguration-routingrules-routingrulecondition-keyprefixequals
	KeyPrefixEquals *StringExpr `json:"KeyPrefixEquals,omitempty"`
}

// S3BucketRoutingRuleConditionList represents a list of S3BucketRoutingRuleCondition
type S3BucketRoutingRuleConditionList []S3BucketRoutingRuleCondition

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3BucketRoutingRuleConditionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3BucketRoutingRuleCondition{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3BucketRoutingRuleConditionList{item}
		return nil
	}
	list := []S3BucketRoutingRuleCondition{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3BucketRoutingRuleConditionList(list)
		return nil
	}
	return err
}

// S3BucketRule represents the AWS::S3::Bucket.Rule CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html
type S3BucketRule struct {
	// ExpirationDate docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-expirationdate
	ExpirationDate time.Time `json:"ExpirationDate,omitempty"`
	// ExpirationInDays docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-expirationindays
	ExpirationInDays *IntegerExpr `json:"ExpirationInDays,omitempty"`
	// ID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-id
	ID *StringExpr `json:"Id,omitempty"`
	// NoncurrentVersionExpirationInDays docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-noncurrentversionexpirationindays
	NoncurrentVersionExpirationInDays *IntegerExpr `json:"NoncurrentVersionExpirationInDays,omitempty"`
	// NoncurrentVersionTransition docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-noncurrentversiontransition
	NoncurrentVersionTransition *S3BucketNoncurrentVersionTransition `json:"NoncurrentVersionTransition,omitempty"`
	// NoncurrentVersionTransitions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-noncurrentversiontransitions
	NoncurrentVersionTransitions *S3BucketNoncurrentVersionTransition `json:"NoncurrentVersionTransitions,omitempty"`
	// Prefix docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-prefix
	Prefix *StringExpr `json:"Prefix,omitempty"`
	// Status docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-status
	Status *StringExpr `json:"Status,omitempty" validate:"dive,required"`
	// Transition docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-transition
	Transition *S3BucketTransition `json:"Transition,omitempty"`
	// Transitions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-transitions
	Transitions *S3BucketTransition `json:"Transitions,omitempty"`
}

// S3BucketRuleList represents a list of S3BucketRule
type S3BucketRuleList []S3BucketRule

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3BucketRuleList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3BucketRule{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3BucketRuleList{item}
		return nil
	}
	list := []S3BucketRule{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3BucketRuleList(list)
		return nil
	}
	return err
}

// S3BucketS3KeyFilter represents the AWS::S3::Bucket.S3KeyFilter CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter-s3key.html
type S3BucketS3KeyFilter struct {
	// Rules docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter-s3key.html#cfn-s3-bucket-notificationconfiguraiton-config-filter-s3key-rules
	Rules *S3BucketFilterRuleList `json:"Rules,omitempty" validate:"dive,required"`
}

// S3BucketS3KeyFilterList represents a list of S3BucketS3KeyFilter
type S3BucketS3KeyFilterList []S3BucketS3KeyFilter

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3BucketS3KeyFilterList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3BucketS3KeyFilter{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3BucketS3KeyFilterList{item}
		return nil
	}
	list := []S3BucketS3KeyFilter{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3BucketS3KeyFilterList(list)
		return nil
	}
	return err
}

// S3BucketTopicConfiguration represents the AWS::S3::Bucket.TopicConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-topicconfig.html
type S3BucketTopicConfiguration struct {
	// Event docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-topicconfig.html#cfn-s3-bucket-notificationconfig-topicconfig-event
	Event *StringExpr `json:"Event,omitempty" validate:"dive,required"`
	// Filter docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-topicconfig.html#cfn-s3-bucket-notificationconfig-topicconfig-filter
	Filter *S3BucketNotificationFilter `json:"Filter,omitempty"`
	// Topic docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-topicconfig.html#cfn-s3-bucket-notificationconfig-topicconfig-topic
	Topic *StringExpr `json:"Topic,omitempty" validate:"dive,required"`
}

// S3BucketTopicConfigurationList represents a list of S3BucketTopicConfiguration
type S3BucketTopicConfigurationList []S3BucketTopicConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3BucketTopicConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3BucketTopicConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3BucketTopicConfigurationList{item}
		return nil
	}
	list := []S3BucketTopicConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3BucketTopicConfigurationList(list)
		return nil
	}
	return err
}

// S3BucketTransition represents the AWS::S3::Bucket.Transition CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-transition.html
type S3BucketTransition struct {
	// StorageClass docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-transition.html#cfn-s3-bucket-lifecycleconfig-rule-transition-storageclass
	StorageClass *StringExpr `json:"StorageClass,omitempty" validate:"dive,required"`
	// TransitionDate docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-transition.html#cfn-s3-bucket-lifecycleconfig-rule-transition-transitiondate
	TransitionDate time.Time `json:"TransitionDate,omitempty"`
	// TransitionInDays docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-transition.html#cfn-s3-bucket-lifecycleconfig-rule-transition-transitionindays
	TransitionInDays *IntegerExpr `json:"TransitionInDays,omitempty"`
}

// S3BucketTransitionList represents a list of S3BucketTransition
type S3BucketTransitionList []S3BucketTransition

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3BucketTransitionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3BucketTransition{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3BucketTransitionList{item}
		return nil
	}
	list := []S3BucketTransition{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3BucketTransitionList(list)
		return nil
	}
	return err
}

// S3BucketVersioningConfiguration represents the AWS::S3::Bucket.VersioningConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-versioningconfig.html
type S3BucketVersioningConfiguration struct {
	// Status docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-versioningconfig.html#cfn-s3-bucket-versioningconfig-status
	Status *StringExpr `json:"Status,omitempty" validate:"dive,required"`
}

// S3BucketVersioningConfigurationList represents a list of S3BucketVersioningConfiguration
type S3BucketVersioningConfigurationList []S3BucketVersioningConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3BucketVersioningConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3BucketVersioningConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3BucketVersioningConfigurationList{item}
		return nil
	}
	list := []S3BucketVersioningConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3BucketVersioningConfigurationList(list)
		return nil
	}
	return err
}

// S3BucketWebsiteConfiguration represents the AWS::S3::Bucket.WebsiteConfiguration CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration.html
type S3BucketWebsiteConfiguration struct {
	// ErrorDocument docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration.html#cfn-s3-websiteconfiguration-errordocument
	ErrorDocument *StringExpr `json:"ErrorDocument,omitempty"`
	// IndexDocument docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration.html#cfn-s3-websiteconfiguration-indexdocument
	IndexDocument *StringExpr `json:"IndexDocument,omitempty"`
	// RedirectAllRequestsTo docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration.html#cfn-s3-websiteconfiguration-redirectallrequeststo
	RedirectAllRequestsTo *S3BucketRedirectAllRequestsTo `json:"RedirectAllRequestsTo,omitempty"`
	// RoutingRules docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration.html#cfn-s3-websiteconfiguration-routingrules
	RoutingRules *S3BucketRoutingRuleList `json:"RoutingRules,omitempty"`
}

// S3BucketWebsiteConfigurationList represents a list of S3BucketWebsiteConfiguration
type S3BucketWebsiteConfigurationList []S3BucketWebsiteConfiguration

// UnmarshalJSON sets the object from the provided JSON representation
func (l *S3BucketWebsiteConfigurationList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := S3BucketWebsiteConfiguration{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = S3BucketWebsiteConfigurationList{item}
		return nil
	}
	list := []S3BucketWebsiteConfiguration{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = S3BucketWebsiteConfigurationList(list)
		return nil
	}
	return err
}

// SNSTopicSubscription represents the AWS::SNS::Topic.Subscription CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-subscription.html
type SNSTopicSubscription struct {
	// Endpoint docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-subscription.html#cfn-sns-topic-subscription-endpoint
	Endpoint *StringExpr `json:"Endpoint,omitempty" validate:"dive,required"`
	// Protocol docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-subscription.html#cfn-sns-topic-subscription-protocol
	Protocol *StringExpr `json:"Protocol,omitempty" validate:"dive,required"`
}

// SNSTopicSubscriptionList represents a list of SNSTopicSubscription
type SNSTopicSubscriptionList []SNSTopicSubscription

// UnmarshalJSON sets the object from the provided JSON representation
func (l *SNSTopicSubscriptionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := SNSTopicSubscription{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = SNSTopicSubscriptionList{item}
		return nil
	}
	list := []SNSTopicSubscription{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = SNSTopicSubscriptionList(list)
		return nil
	}
	return err
}

// SSMAssociationParameterValues represents the AWS::SSM::Association.ParameterValues CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-parametervalues.html
type SSMAssociationParameterValues struct {
	// ParameterValues docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-parametervalues.html#cfn-ssm-association-parametervalues-parametervalues
	ParameterValues *StringListExpr `json:"ParameterValues,omitempty" validate:"dive,required"`
}

// SSMAssociationParameterValuesList represents a list of SSMAssociationParameterValues
type SSMAssociationParameterValuesList []SSMAssociationParameterValues

// UnmarshalJSON sets the object from the provided JSON representation
func (l *SSMAssociationParameterValuesList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := SSMAssociationParameterValues{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = SSMAssociationParameterValuesList{item}
		return nil
	}
	list := []SSMAssociationParameterValues{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = SSMAssociationParameterValuesList(list)
		return nil
	}
	return err
}

// SSMAssociationTarget represents the AWS::SSM::Association.Target CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-target.html
type SSMAssociationTarget struct {
	// Key docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-target.html#cfn-ssm-association-target-key
	Key *StringExpr `json:"Key,omitempty" validate:"dive,required"`
	// Values docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-target.html#cfn-ssm-association-target-values
	Values *StringListExpr `json:"Values,omitempty" validate:"dive,required"`
}

// SSMAssociationTargetList represents a list of SSMAssociationTarget
type SSMAssociationTargetList []SSMAssociationTarget

// UnmarshalJSON sets the object from the provided JSON representation
func (l *SSMAssociationTargetList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := SSMAssociationTarget{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = SSMAssociationTargetList{item}
		return nil
	}
	list := []SSMAssociationTarget{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = SSMAssociationTargetList(list)
		return nil
	}
	return err
}

// WAFByteMatchSetByteMatchTuple represents the AWS::WAF::ByteMatchSet.ByteMatchTuple CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples.html
type WAFByteMatchSetByteMatchTuple struct {
	// FieldToMatch docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples.html#cfn-waf-bytematchset-bytematchtuples-fieldtomatch
	FieldToMatch *WAFByteMatchSetFieldToMatch `json:"FieldToMatch,omitempty" validate:"dive,required"`
	// PositionalConstraint docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples.html#cfn-waf-bytematchset-bytematchtuples-positionalconstraint
	PositionalConstraint *StringExpr `json:"PositionalConstraint,omitempty" validate:"dive,required"`
	// TargetString docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples.html#cfn-waf-bytematchset-bytematchtuples-targetstring
	TargetString *StringExpr `json:"TargetString,omitempty"`
	// TargetStringBase64 docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples.html#cfn-waf-bytematchset-bytematchtuples-targetstringbase64
	TargetStringBase64 *StringExpr `json:"TargetStringBase64,omitempty"`
	// TextTransformation docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples.html#cfn-waf-bytematchset-bytematchtuples-texttransformation
	TextTransformation *StringExpr `json:"TextTransformation,omitempty" validate:"dive,required"`
}

// WAFByteMatchSetByteMatchTupleList represents a list of WAFByteMatchSetByteMatchTuple
type WAFByteMatchSetByteMatchTupleList []WAFByteMatchSetByteMatchTuple

// UnmarshalJSON sets the object from the provided JSON representation
func (l *WAFByteMatchSetByteMatchTupleList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := WAFByteMatchSetByteMatchTuple{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = WAFByteMatchSetByteMatchTupleList{item}
		return nil
	}
	list := []WAFByteMatchSetByteMatchTuple{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = WAFByteMatchSetByteMatchTupleList(list)
		return nil
	}
	return err
}

// WAFByteMatchSetFieldToMatch represents the AWS::WAF::ByteMatchSet.FieldToMatch CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples-fieldtomatch.html
type WAFByteMatchSetFieldToMatch struct {
	// Data docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples-fieldtomatch.html#cfn-waf-bytematchset-bytematchtuples-fieldtomatch-data
	Data *StringExpr `json:"Data,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples-fieldtomatch.html#cfn-waf-bytematchset-bytematchtuples-fieldtomatch-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
}

// WAFByteMatchSetFieldToMatchList represents a list of WAFByteMatchSetFieldToMatch
type WAFByteMatchSetFieldToMatchList []WAFByteMatchSetFieldToMatch

// UnmarshalJSON sets the object from the provided JSON representation
func (l *WAFByteMatchSetFieldToMatchList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := WAFByteMatchSetFieldToMatch{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = WAFByteMatchSetFieldToMatchList{item}
		return nil
	}
	list := []WAFByteMatchSetFieldToMatch{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = WAFByteMatchSetFieldToMatchList(list)
		return nil
	}
	return err
}

// WAFIPSetIPSetDescriptor represents the AWS::WAF::IPSet.IPSetDescriptor CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-ipset-ipsetdescriptors.html
type WAFIPSetIPSetDescriptor struct {
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-ipset-ipsetdescriptors.html#cfn-waf-ipset-ipsetdescriptors-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
	// Value docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-ipset-ipsetdescriptors.html#cfn-waf-ipset-ipsetdescriptors-value
	Value *StringExpr `json:"Value,omitempty" validate:"dive,required"`
}

// WAFIPSetIPSetDescriptorList represents a list of WAFIPSetIPSetDescriptor
type WAFIPSetIPSetDescriptorList []WAFIPSetIPSetDescriptor

// UnmarshalJSON sets the object from the provided JSON representation
func (l *WAFIPSetIPSetDescriptorList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := WAFIPSetIPSetDescriptor{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = WAFIPSetIPSetDescriptorList{item}
		return nil
	}
	list := []WAFIPSetIPSetDescriptor{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = WAFIPSetIPSetDescriptorList(list)
		return nil
	}
	return err
}

// WAFRulePredicate represents the AWS::WAF::Rule.Predicate CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-rule-predicates.html
type WAFRulePredicate struct {
	// DataID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-rule-predicates.html#cfn-waf-rule-predicates-dataid
	DataID *StringExpr `json:"DataId,omitempty" validate:"dive,required"`
	// Negated docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-rule-predicates.html#cfn-waf-rule-predicates-negated
	Negated *BoolExpr `json:"Negated,omitempty" validate:"dive,required"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-rule-predicates.html#cfn-waf-rule-predicates-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
}

// WAFRulePredicateList represents a list of WAFRulePredicate
type WAFRulePredicateList []WAFRulePredicate

// UnmarshalJSON sets the object from the provided JSON representation
func (l *WAFRulePredicateList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := WAFRulePredicate{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = WAFRulePredicateList{item}
		return nil
	}
	list := []WAFRulePredicate{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = WAFRulePredicateList(list)
		return nil
	}
	return err
}

// WAFSizeConstraintSetFieldToMatch represents the AWS::WAF::SizeConstraintSet.FieldToMatch CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sizeconstraintset-sizeconstraint-fieldtomatch.html
type WAFSizeConstraintSetFieldToMatch struct {
	// Data docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sizeconstraintset-sizeconstraint-fieldtomatch.html#cfn-waf-sizeconstraintset-sizeconstraint-fieldtomatch-data
	Data *StringExpr `json:"Data,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sizeconstraintset-sizeconstraint-fieldtomatch.html#cfn-waf-sizeconstraintset-sizeconstraint-fieldtomatch-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
}

// WAFSizeConstraintSetFieldToMatchList represents a list of WAFSizeConstraintSetFieldToMatch
type WAFSizeConstraintSetFieldToMatchList []WAFSizeConstraintSetFieldToMatch

// UnmarshalJSON sets the object from the provided JSON representation
func (l *WAFSizeConstraintSetFieldToMatchList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := WAFSizeConstraintSetFieldToMatch{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = WAFSizeConstraintSetFieldToMatchList{item}
		return nil
	}
	list := []WAFSizeConstraintSetFieldToMatch{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = WAFSizeConstraintSetFieldToMatchList(list)
		return nil
	}
	return err
}

// WAFSizeConstraintSetSizeConstraint represents the AWS::WAF::SizeConstraintSet.SizeConstraint CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sizeconstraintset-sizeconstraint.html
type WAFSizeConstraintSetSizeConstraint struct {
	// ComparisonOperator docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sizeconstraintset-sizeconstraint.html#cfn-waf-sizeconstraintset-sizeconstraint-comparisonoperator
	ComparisonOperator *StringExpr `json:"ComparisonOperator,omitempty" validate:"dive,required"`
	// FieldToMatch docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sizeconstraintset-sizeconstraint.html#cfn-waf-sizeconstraintset-sizeconstraint-fieldtomatch
	FieldToMatch *WAFSizeConstraintSetFieldToMatch `json:"FieldToMatch,omitempty" validate:"dive,required"`
	// Size docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sizeconstraintset-sizeconstraint.html#cfn-waf-sizeconstraintset-sizeconstraint-size
	Size *IntegerExpr `json:"Size,omitempty" validate:"dive,required"`
	// TextTransformation docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sizeconstraintset-sizeconstraint.html#cfn-waf-sizeconstraintset-sizeconstraint-texttransformation
	TextTransformation *StringExpr `json:"TextTransformation,omitempty" validate:"dive,required"`
}

// WAFSizeConstraintSetSizeConstraintList represents a list of WAFSizeConstraintSetSizeConstraint
type WAFSizeConstraintSetSizeConstraintList []WAFSizeConstraintSetSizeConstraint

// UnmarshalJSON sets the object from the provided JSON representation
func (l *WAFSizeConstraintSetSizeConstraintList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := WAFSizeConstraintSetSizeConstraint{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = WAFSizeConstraintSetSizeConstraintList{item}
		return nil
	}
	list := []WAFSizeConstraintSetSizeConstraint{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = WAFSizeConstraintSetSizeConstraintList(list)
		return nil
	}
	return err
}

// WAFSQLInjectionMatchSetFieldToMatch represents the AWS::WAF::SqlInjectionMatchSet.FieldToMatch CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples-fieldtomatch.html
type WAFSQLInjectionMatchSetFieldToMatch struct {
	// Data docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples-fieldtomatch.html#cfn-waf-sizeconstraintset-sizeconstraint-fieldtomatch-data
	Data *StringExpr `json:"Data,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples-fieldtomatch.html#cfn-waf-sizeconstraintset-sizeconstraint-fieldtomatch-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
}

// WAFSQLInjectionMatchSetFieldToMatchList represents a list of WAFSQLInjectionMatchSetFieldToMatch
type WAFSQLInjectionMatchSetFieldToMatchList []WAFSQLInjectionMatchSetFieldToMatch

// UnmarshalJSON sets the object from the provided JSON representation
func (l *WAFSQLInjectionMatchSetFieldToMatchList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := WAFSQLInjectionMatchSetFieldToMatch{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = WAFSQLInjectionMatchSetFieldToMatchList{item}
		return nil
	}
	list := []WAFSQLInjectionMatchSetFieldToMatch{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = WAFSQLInjectionMatchSetFieldToMatchList(list)
		return nil
	}
	return err
}

// WAFSQLInjectionMatchSetSQLInjectionMatchTuple represents the AWS::WAF::SqlInjectionMatchSet.SqlInjectionMatchTuple CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sqlinjectionmatchset-sqlinjectionmatchtuples.html
type WAFSQLInjectionMatchSetSQLInjectionMatchTuple struct {
	// FieldToMatch docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sqlinjectionmatchset-sqlinjectionmatchtuples.html#cfn-waf-sqlinjectionmatchset-sqlinjectionmatchtuples-fieldtomatch
	FieldToMatch *WAFSQLInjectionMatchSetFieldToMatch `json:"FieldToMatch,omitempty" validate:"dive,required"`
	// TextTransformation docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sqlinjectionmatchset-sqlinjectionmatchtuples.html#cfn-waf-sqlinjectionmatchset-sqlinjectionmatchtuples-texttransformation
	TextTransformation *StringExpr `json:"TextTransformation,omitempty" validate:"dive,required"`
}

// WAFSQLInjectionMatchSetSQLInjectionMatchTupleList represents a list of WAFSQLInjectionMatchSetSQLInjectionMatchTuple
type WAFSQLInjectionMatchSetSQLInjectionMatchTupleList []WAFSQLInjectionMatchSetSQLInjectionMatchTuple

// UnmarshalJSON sets the object from the provided JSON representation
func (l *WAFSQLInjectionMatchSetSQLInjectionMatchTupleList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := WAFSQLInjectionMatchSetSQLInjectionMatchTuple{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = WAFSQLInjectionMatchSetSQLInjectionMatchTupleList{item}
		return nil
	}
	list := []WAFSQLInjectionMatchSetSQLInjectionMatchTuple{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = WAFSQLInjectionMatchSetSQLInjectionMatchTupleList(list)
		return nil
	}
	return err
}

// WAFWebACLActivatedRule represents the AWS::WAF::WebACL.ActivatedRule CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-webacl-rules.html
type WAFWebACLActivatedRule struct {
	// Action docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-webacl-rules.html#cfn-waf-webacl-rules-action
	Action *WAFWebACLWafAction `json:"Action,omitempty" validate:"dive,required"`
	// Priority docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-webacl-rules.html#cfn-waf-webacl-rules-priority
	Priority *IntegerExpr `json:"Priority,omitempty" validate:"dive,required"`
	// RuleID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-webacl-rules.html#cfn-waf-webacl-rules-ruleid
	RuleID *StringExpr `json:"RuleId,omitempty" validate:"dive,required"`
}

// WAFWebACLActivatedRuleList represents a list of WAFWebACLActivatedRule
type WAFWebACLActivatedRuleList []WAFWebACLActivatedRule

// UnmarshalJSON sets the object from the provided JSON representation
func (l *WAFWebACLActivatedRuleList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := WAFWebACLActivatedRule{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = WAFWebACLActivatedRuleList{item}
		return nil
	}
	list := []WAFWebACLActivatedRule{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = WAFWebACLActivatedRuleList(list)
		return nil
	}
	return err
}

// WAFWebACLWafAction represents the AWS::WAF::WebACL.WafAction CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-webacl-action.html
type WAFWebACLWafAction struct {
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-webacl-action.html#cfn-waf-webacl-action-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
}

// WAFWebACLWafActionList represents a list of WAFWebACLWafAction
type WAFWebACLWafActionList []WAFWebACLWafAction

// UnmarshalJSON sets the object from the provided JSON representation
func (l *WAFWebACLWafActionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := WAFWebACLWafAction{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = WAFWebACLWafActionList{item}
		return nil
	}
	list := []WAFWebACLWafAction{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = WAFWebACLWafActionList(list)
		return nil
	}
	return err
}

// WAFXSSMatchSetFieldToMatch represents the AWS::WAF::XssMatchSet.FieldToMatch CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-xssmatchset-xssmatchtuple-fieldtomatch.html
type WAFXSSMatchSetFieldToMatch struct {
	// Data docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-xssmatchset-xssmatchtuple-fieldtomatch.html#cfn-waf-xssmatchset-xssmatchtuple-fieldtomatch-data
	Data *StringExpr `json:"Data,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-xssmatchset-xssmatchtuple-fieldtomatch.html#cfn-waf-xssmatchset-xssmatchtuple-fieldtomatch-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
}

// WAFXSSMatchSetFieldToMatchList represents a list of WAFXSSMatchSetFieldToMatch
type WAFXSSMatchSetFieldToMatchList []WAFXSSMatchSetFieldToMatch

// UnmarshalJSON sets the object from the provided JSON representation
func (l *WAFXSSMatchSetFieldToMatchList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := WAFXSSMatchSetFieldToMatch{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = WAFXSSMatchSetFieldToMatchList{item}
		return nil
	}
	list := []WAFXSSMatchSetFieldToMatch{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = WAFXSSMatchSetFieldToMatchList(list)
		return nil
	}
	return err
}

// WAFXSSMatchSetXSSMatchTuple represents the AWS::WAF::XssMatchSet.XssMatchTuple CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-xssmatchset-xssmatchtuple.html
type WAFXSSMatchSetXSSMatchTuple struct {
	// FieldToMatch docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-xssmatchset-xssmatchtuple.html#cfn-waf-xssmatchset-xssmatchtuple-fieldtomatch
	FieldToMatch *WAFXSSMatchSetFieldToMatch `json:"FieldToMatch,omitempty" validate:"dive,required"`
	// TextTransformation docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-xssmatchset-xssmatchtuple.html#cfn-waf-xssmatchset-xssmatchtuple-texttransformation
	TextTransformation *StringExpr `json:"TextTransformation,omitempty" validate:"dive,required"`
}

// WAFXSSMatchSetXSSMatchTupleList represents a list of WAFXSSMatchSetXSSMatchTuple
type WAFXSSMatchSetXSSMatchTupleList []WAFXSSMatchSetXSSMatchTuple

// UnmarshalJSON sets the object from the provided JSON representation
func (l *WAFXSSMatchSetXSSMatchTupleList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := WAFXSSMatchSetXSSMatchTuple{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = WAFXSSMatchSetXSSMatchTupleList{item}
		return nil
	}
	list := []WAFXSSMatchSetXSSMatchTuple{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = WAFXSSMatchSetXSSMatchTupleList(list)
		return nil
	}
	return err
}

// WAFRegionalByteMatchSetByteMatchTuple represents the AWS::WAFRegional::ByteMatchSet.ByteMatchTuple CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-bytematchset-bytematchtuple.html
type WAFRegionalByteMatchSetByteMatchTuple struct {
	// FieldToMatch docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-bytematchset-bytematchtuple.html#cfn-wafregional-bytematchset-bytematchtuple-fieldtomatch
	FieldToMatch *WAFRegionalByteMatchSetFieldToMatch `json:"FieldToMatch,omitempty" validate:"dive,required"`
	// PositionalConstraint docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-bytematchset-bytematchtuple.html#cfn-wafregional-bytematchset-bytematchtuple-positionalconstraint
	PositionalConstraint *StringExpr `json:"PositionalConstraint,omitempty" validate:"dive,required"`
	// TargetString docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-bytematchset-bytematchtuple.html#cfn-wafregional-bytematchset-bytematchtuple-targetstring
	TargetString *StringExpr `json:"TargetString,omitempty"`
	// TargetStringBase64 docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-bytematchset-bytematchtuple.html#cfn-wafregional-bytematchset-bytematchtuple-targetstringbase64
	TargetStringBase64 *StringExpr `json:"TargetStringBase64,omitempty"`
	// TextTransformation docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-bytematchset-bytematchtuple.html#cfn-wafregional-bytematchset-bytematchtuple-texttransformation
	TextTransformation *StringExpr `json:"TextTransformation,omitempty" validate:"dive,required"`
}

// WAFRegionalByteMatchSetByteMatchTupleList represents a list of WAFRegionalByteMatchSetByteMatchTuple
type WAFRegionalByteMatchSetByteMatchTupleList []WAFRegionalByteMatchSetByteMatchTuple

// UnmarshalJSON sets the object from the provided JSON representation
func (l *WAFRegionalByteMatchSetByteMatchTupleList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := WAFRegionalByteMatchSetByteMatchTuple{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = WAFRegionalByteMatchSetByteMatchTupleList{item}
		return nil
	}
	list := []WAFRegionalByteMatchSetByteMatchTuple{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = WAFRegionalByteMatchSetByteMatchTupleList(list)
		return nil
	}
	return err
}

// WAFRegionalByteMatchSetFieldToMatch represents the AWS::WAFRegional::ByteMatchSet.FieldToMatch CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-bytematchset-fieldtomatch.html
type WAFRegionalByteMatchSetFieldToMatch struct {
	// Data docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-bytematchset-fieldtomatch.html#cfn-wafregional-bytematchset-fieldtomatch-data
	Data *StringExpr `json:"Data,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-bytematchset-fieldtomatch.html#cfn-wafregional-bytematchset-fieldtomatch-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
}

// WAFRegionalByteMatchSetFieldToMatchList represents a list of WAFRegionalByteMatchSetFieldToMatch
type WAFRegionalByteMatchSetFieldToMatchList []WAFRegionalByteMatchSetFieldToMatch

// UnmarshalJSON sets the object from the provided JSON representation
func (l *WAFRegionalByteMatchSetFieldToMatchList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := WAFRegionalByteMatchSetFieldToMatch{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = WAFRegionalByteMatchSetFieldToMatchList{item}
		return nil
	}
	list := []WAFRegionalByteMatchSetFieldToMatch{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = WAFRegionalByteMatchSetFieldToMatchList(list)
		return nil
	}
	return err
}

// WAFRegionalIPSetIPSetDescriptor represents the AWS::WAFRegional::IPSet.IPSetDescriptor CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-ipset-ipsetdescriptor.html
type WAFRegionalIPSetIPSetDescriptor struct {
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-ipset-ipsetdescriptor.html#cfn-wafregional-ipset-ipsetdescriptor-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
	// Value docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-ipset-ipsetdescriptor.html#cfn-wafregional-ipset-ipsetdescriptor-value
	Value *StringExpr `json:"Value,omitempty" validate:"dive,required"`
}

// WAFRegionalIPSetIPSetDescriptorList represents a list of WAFRegionalIPSetIPSetDescriptor
type WAFRegionalIPSetIPSetDescriptorList []WAFRegionalIPSetIPSetDescriptor

// UnmarshalJSON sets the object from the provided JSON representation
func (l *WAFRegionalIPSetIPSetDescriptorList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := WAFRegionalIPSetIPSetDescriptor{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = WAFRegionalIPSetIPSetDescriptorList{item}
		return nil
	}
	list := []WAFRegionalIPSetIPSetDescriptor{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = WAFRegionalIPSetIPSetDescriptorList(list)
		return nil
	}
	return err
}

// WAFRegionalRulePredicate represents the AWS::WAFRegional::Rule.Predicate CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-rule-predicate.html
type WAFRegionalRulePredicate struct {
	// DataID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-rule-predicate.html#cfn-wafregional-rule-predicate-dataid
	DataID *StringExpr `json:"DataId,omitempty" validate:"dive,required"`
	// Negated docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-rule-predicate.html#cfn-wafregional-rule-predicate-negated
	Negated *BoolExpr `json:"Negated,omitempty" validate:"dive,required"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-rule-predicate.html#cfn-wafregional-rule-predicate-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
}

// WAFRegionalRulePredicateList represents a list of WAFRegionalRulePredicate
type WAFRegionalRulePredicateList []WAFRegionalRulePredicate

// UnmarshalJSON sets the object from the provided JSON representation
func (l *WAFRegionalRulePredicateList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := WAFRegionalRulePredicate{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = WAFRegionalRulePredicateList{item}
		return nil
	}
	list := []WAFRegionalRulePredicate{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = WAFRegionalRulePredicateList(list)
		return nil
	}
	return err
}

// WAFRegionalSizeConstraintSetFieldToMatch represents the AWS::WAFRegional::SizeConstraintSet.FieldToMatch CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sizeconstraintset-fieldtomatch.html
type WAFRegionalSizeConstraintSetFieldToMatch struct {
	// Data docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sizeconstraintset-fieldtomatch.html#cfn-wafregional-sizeconstraintset-fieldtomatch-data
	Data *StringExpr `json:"Data,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sizeconstraintset-fieldtomatch.html#cfn-wafregional-sizeconstraintset-fieldtomatch-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
}

// WAFRegionalSizeConstraintSetFieldToMatchList represents a list of WAFRegionalSizeConstraintSetFieldToMatch
type WAFRegionalSizeConstraintSetFieldToMatchList []WAFRegionalSizeConstraintSetFieldToMatch

// UnmarshalJSON sets the object from the provided JSON representation
func (l *WAFRegionalSizeConstraintSetFieldToMatchList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := WAFRegionalSizeConstraintSetFieldToMatch{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = WAFRegionalSizeConstraintSetFieldToMatchList{item}
		return nil
	}
	list := []WAFRegionalSizeConstraintSetFieldToMatch{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = WAFRegionalSizeConstraintSetFieldToMatchList(list)
		return nil
	}
	return err
}

// WAFRegionalSizeConstraintSetSizeConstraint represents the AWS::WAFRegional::SizeConstraintSet.SizeConstraint CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sizeconstraintset-sizeconstraint.html
type WAFRegionalSizeConstraintSetSizeConstraint struct {
	// ComparisonOperator docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sizeconstraintset-sizeconstraint.html#cfn-wafregional-sizeconstraintset-sizeconstraint-comparisonoperator
	ComparisonOperator *StringExpr `json:"ComparisonOperator,omitempty" validate:"dive,required"`
	// FieldToMatch docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sizeconstraintset-sizeconstraint.html#cfn-wafregional-sizeconstraintset-sizeconstraint-fieldtomatch
	FieldToMatch *WAFRegionalSizeConstraintSetFieldToMatch `json:"FieldToMatch,omitempty" validate:"dive,required"`
	// Size docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sizeconstraintset-sizeconstraint.html#cfn-wafregional-sizeconstraintset-sizeconstraint-size
	Size *IntegerExpr `json:"Size,omitempty" validate:"dive,required"`
	// TextTransformation docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sizeconstraintset-sizeconstraint.html#cfn-wafregional-sizeconstraintset-sizeconstraint-texttransformation
	TextTransformation *StringExpr `json:"TextTransformation,omitempty" validate:"dive,required"`
}

// WAFRegionalSizeConstraintSetSizeConstraintList represents a list of WAFRegionalSizeConstraintSetSizeConstraint
type WAFRegionalSizeConstraintSetSizeConstraintList []WAFRegionalSizeConstraintSetSizeConstraint

// UnmarshalJSON sets the object from the provided JSON representation
func (l *WAFRegionalSizeConstraintSetSizeConstraintList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := WAFRegionalSizeConstraintSetSizeConstraint{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = WAFRegionalSizeConstraintSetSizeConstraintList{item}
		return nil
	}
	list := []WAFRegionalSizeConstraintSetSizeConstraint{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = WAFRegionalSizeConstraintSetSizeConstraintList(list)
		return nil
	}
	return err
}

// WAFRegionalSQLInjectionMatchSetFieldToMatch represents the AWS::WAFRegional::SqlInjectionMatchSet.FieldToMatch CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sqlinjectionmatchset-fieldtomatch.html
type WAFRegionalSQLInjectionMatchSetFieldToMatch struct {
	// Data docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sqlinjectionmatchset-fieldtomatch.html#cfn-wafregional-sqlinjectionmatchset-fieldtomatch-data
	Data *StringExpr `json:"Data,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sqlinjectionmatchset-fieldtomatch.html#cfn-wafregional-sqlinjectionmatchset-fieldtomatch-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
}

// WAFRegionalSQLInjectionMatchSetFieldToMatchList represents a list of WAFRegionalSQLInjectionMatchSetFieldToMatch
type WAFRegionalSQLInjectionMatchSetFieldToMatchList []WAFRegionalSQLInjectionMatchSetFieldToMatch

// UnmarshalJSON sets the object from the provided JSON representation
func (l *WAFRegionalSQLInjectionMatchSetFieldToMatchList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := WAFRegionalSQLInjectionMatchSetFieldToMatch{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = WAFRegionalSQLInjectionMatchSetFieldToMatchList{item}
		return nil
	}
	list := []WAFRegionalSQLInjectionMatchSetFieldToMatch{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = WAFRegionalSQLInjectionMatchSetFieldToMatchList(list)
		return nil
	}
	return err
}

// WAFRegionalSQLInjectionMatchSetSQLInjectionMatchTuple represents the AWS::WAFRegional::SqlInjectionMatchSet.SqlInjectionMatchTuple CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sqlinjectionmatchset-sqlinjectionmatchtuple.html
type WAFRegionalSQLInjectionMatchSetSQLInjectionMatchTuple struct {
	// FieldToMatch docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sqlinjectionmatchset-sqlinjectionmatchtuple.html#cfn-wafregional-sqlinjectionmatchset-sqlinjectionmatchtuple-fieldtomatch
	FieldToMatch *WAFRegionalSQLInjectionMatchSetFieldToMatch `json:"FieldToMatch,omitempty" validate:"dive,required"`
	// TextTransformation docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sqlinjectionmatchset-sqlinjectionmatchtuple.html#cfn-wafregional-sqlinjectionmatchset-sqlinjectionmatchtuple-texttransformation
	TextTransformation *StringExpr `json:"TextTransformation,omitempty" validate:"dive,required"`
}

// WAFRegionalSQLInjectionMatchSetSQLInjectionMatchTupleList represents a list of WAFRegionalSQLInjectionMatchSetSQLInjectionMatchTuple
type WAFRegionalSQLInjectionMatchSetSQLInjectionMatchTupleList []WAFRegionalSQLInjectionMatchSetSQLInjectionMatchTuple

// UnmarshalJSON sets the object from the provided JSON representation
func (l *WAFRegionalSQLInjectionMatchSetSQLInjectionMatchTupleList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := WAFRegionalSQLInjectionMatchSetSQLInjectionMatchTuple{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = WAFRegionalSQLInjectionMatchSetSQLInjectionMatchTupleList{item}
		return nil
	}
	list := []WAFRegionalSQLInjectionMatchSetSQLInjectionMatchTuple{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = WAFRegionalSQLInjectionMatchSetSQLInjectionMatchTupleList(list)
		return nil
	}
	return err
}

// WAFRegionalWebACLAction represents the AWS::WAFRegional::WebACL.Action CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-webacl-action.html
type WAFRegionalWebACLAction struct {
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-webacl-action.html#cfn-wafregional-webacl-action-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
}

// WAFRegionalWebACLActionList represents a list of WAFRegionalWebACLAction
type WAFRegionalWebACLActionList []WAFRegionalWebACLAction

// UnmarshalJSON sets the object from the provided JSON representation
func (l *WAFRegionalWebACLActionList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := WAFRegionalWebACLAction{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = WAFRegionalWebACLActionList{item}
		return nil
	}
	list := []WAFRegionalWebACLAction{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = WAFRegionalWebACLActionList(list)
		return nil
	}
	return err
}

// WAFRegionalWebACLRule represents the AWS::WAFRegional::WebACL.Rule CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-webacl-rule.html
type WAFRegionalWebACLRule struct {
	// Action docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-webacl-rule.html#cfn-wafregional-webacl-rule-action
	Action *WAFRegionalWebACLAction `json:"Action,omitempty" validate:"dive,required"`
	// Priority docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-webacl-rule.html#cfn-wafregional-webacl-rule-priority
	Priority *IntegerExpr `json:"Priority,omitempty" validate:"dive,required"`
	// RuleID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-webacl-rule.html#cfn-wafregional-webacl-rule-ruleid
	RuleID *StringExpr `json:"RuleId,omitempty" validate:"dive,required"`
}

// WAFRegionalWebACLRuleList represents a list of WAFRegionalWebACLRule
type WAFRegionalWebACLRuleList []WAFRegionalWebACLRule

// UnmarshalJSON sets the object from the provided JSON representation
func (l *WAFRegionalWebACLRuleList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := WAFRegionalWebACLRule{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = WAFRegionalWebACLRuleList{item}
		return nil
	}
	list := []WAFRegionalWebACLRule{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = WAFRegionalWebACLRuleList(list)
		return nil
	}
	return err
}

// WAFRegionalXSSMatchSetFieldToMatch represents the AWS::WAFRegional::XssMatchSet.FieldToMatch CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-fieldtomatch.html
type WAFRegionalXSSMatchSetFieldToMatch struct {
	// Data docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-fieldtomatch.html#cfn-wafregional-xssmatchset-fieldtomatch-data
	Data *StringExpr `json:"Data,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-fieldtomatch.html#cfn-wafregional-xssmatchset-fieldtomatch-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
}

// WAFRegionalXSSMatchSetFieldToMatchList represents a list of WAFRegionalXSSMatchSetFieldToMatch
type WAFRegionalXSSMatchSetFieldToMatchList []WAFRegionalXSSMatchSetFieldToMatch

// UnmarshalJSON sets the object from the provided JSON representation
func (l *WAFRegionalXSSMatchSetFieldToMatchList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := WAFRegionalXSSMatchSetFieldToMatch{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = WAFRegionalXSSMatchSetFieldToMatchList{item}
		return nil
	}
	list := []WAFRegionalXSSMatchSetFieldToMatch{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = WAFRegionalXSSMatchSetFieldToMatchList(list)
		return nil
	}
	return err
}

// WAFRegionalXSSMatchSetXSSMatchTuple represents the AWS::WAFRegional::XssMatchSet.XssMatchTuple CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-xssmatchtuple.html
type WAFRegionalXSSMatchSetXSSMatchTuple struct {
	// FieldToMatch docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-xssmatchtuple.html#cfn-wafregional-xssmatchset-xssmatchtuple-fieldtomatch
	FieldToMatch *WAFRegionalXSSMatchSetFieldToMatch `json:"FieldToMatch,omitempty" validate:"dive,required"`
	// TextTransformation docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-xssmatchtuple.html#cfn-wafregional-xssmatchset-xssmatchtuple-texttransformation
	TextTransformation *StringExpr `json:"TextTransformation,omitempty" validate:"dive,required"`
}

// WAFRegionalXSSMatchSetXSSMatchTupleList represents a list of WAFRegionalXSSMatchSetXSSMatchTuple
type WAFRegionalXSSMatchSetXSSMatchTupleList []WAFRegionalXSSMatchSetXSSMatchTuple

// UnmarshalJSON sets the object from the provided JSON representation
func (l *WAFRegionalXSSMatchSetXSSMatchTupleList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := WAFRegionalXSSMatchSetXSSMatchTuple{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = WAFRegionalXSSMatchSetXSSMatchTupleList{item}
		return nil
	}
	list := []WAFRegionalXSSMatchSetXSSMatchTuple{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = WAFRegionalXSSMatchSetXSSMatchTupleList(list)
		return nil
	}
	return err
}

// Tag represents the Tag CloudFormation property type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html
type Tag struct {
	// Key docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html#cfn-resource-tags-key
	Key *StringExpr `json:"Key,omitempty" validate:"dive,required"`
	// Value docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html#cfn-resource-tags-value
	Value *StringExpr `json:"Value,omitempty" validate:"dive,required"`
}

// TagList represents a list of Tag
type TagList []Tag

// UnmarshalJSON sets the object from the provided JSON representation
func (l *TagList) UnmarshalJSON(buf []byte) error {
	// Cloudformation allows a single object when a list of objects is expected
	item := Tag{}
	if err := json.Unmarshal(buf, &item); err == nil {
		*l = TagList{item}
		return nil
	}
	list := []Tag{}
	err := json.Unmarshal(buf, &list)
	if err == nil {
		*l = TagList(list)
		return nil
	}
	return err
}

//
//  ____
// |  _ \ ___  ___  ___  _   _ _ __ ___ ___  ___
// | |_) / _ \/ __|/ _ \| | | | '__/ __/ _ \/ __|
// |  _ <  __/\__ \ (_) | |_| | | | (_|  __/\__ \
// |_| \_\___||___/\___/ \__,_|_|  \___\___||___/
//

// APIGatewayAccount represents the AWS::ApiGateway::Account CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-account.html
type APIGatewayAccount struct {
	// CloudWatchRoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-account.html#cfn-apigateway-account-cloudwatchrolearn
	CloudWatchRoleArn *StringExpr `json:"CloudWatchRoleArn,omitempty"`
}

// CfnResourceType returns AWS::ApiGateway::Account to implement the ResourceProperties interface
func (s APIGatewayAccount) CfnResourceType() string {

	return "AWS::ApiGateway::Account"
}

// APIGatewayAPIKey represents the AWS::ApiGateway::ApiKey CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html
type APIGatewayAPIKey struct {
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html#cfn-apigateway-apigateway-apikey-description
	Description *StringExpr `json:"Description,omitempty"`
	// Enabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html#cfn-apigateway-apigateway-apikey-enabled
	Enabled *BoolExpr `json:"Enabled,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html#cfn-apigateway-apigateway-apikey-name
	Name *StringExpr `json:"Name,omitempty"`
	// StageKeys docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html#cfn-apigateway-apigateway-apikey-stagekeys
	StageKeys *APIGatewayAPIKeyStageKeyList `json:"StageKeys,omitempty"`
}

// CfnResourceType returns AWS::ApiGateway::ApiKey to implement the ResourceProperties interface
func (s APIGatewayAPIKey) CfnResourceType() string {

	return "AWS::ApiGateway::ApiKey"
}

// APIGatewayAuthorizer represents the AWS::ApiGateway::Authorizer CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html
type APIGatewayAuthorizer struct {
	// AuthorizerCredentials docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html#cfn-apigateway-authorizer-authorizercredentials
	AuthorizerCredentials *StringExpr `json:"AuthorizerCredentials,omitempty"`
	// AuthorizerResultTTLInSeconds docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html#cfn-apigateway-authorizer-authorizerresultttlinseconds
	AuthorizerResultTTLInSeconds *IntegerExpr `json:"AuthorizerResultTtlInSeconds,omitempty"`
	// AuthorizerURI docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html#cfn-apigateway-authorizer-authorizeruri
	AuthorizerURI *StringExpr `json:"AuthorizerUri,omitempty"`
	// IDentitySource docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html#cfn-apigateway-authorizer-identitysource
	IDentitySource *StringExpr `json:"IdentitySource,omitempty"`
	// IDentityValidationExpression docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html#cfn-apigateway-authorizer-identityvalidationexpression
	IDentityValidationExpression *StringExpr `json:"IdentityValidationExpression,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html#cfn-apigateway-authorizer-name
	Name *StringExpr `json:"Name,omitempty"`
	// ProviderARNs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html#cfn-apigateway-authorizer-providerarns
	ProviderARNs *StringListExpr `json:"ProviderARNs,omitempty"`
	// RestAPIID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html#cfn-apigateway-authorizer-restapiid
	RestAPIID *StringExpr `json:"RestApiId,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html#cfn-apigateway-authorizer-type
	Type *StringExpr `json:"Type,omitempty"`
}

// CfnResourceType returns AWS::ApiGateway::Authorizer to implement the ResourceProperties interface
func (s APIGatewayAuthorizer) CfnResourceType() string {

	return "AWS::ApiGateway::Authorizer"
}

// APIGatewayBasePathMapping represents the AWS::ApiGateway::BasePathMapping CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-basepathmapping.html
type APIGatewayBasePathMapping struct {
	// BasePath docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-basepathmapping.html#cfn-apigateway-basepathmapping-basepath
	BasePath *StringExpr `json:"BasePath,omitempty"`
	// DomainName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-basepathmapping.html#cfn-apigateway-basepathmapping-domainname
	DomainName *StringExpr `json:"DomainName,omitempty"`
	// RestAPIID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-basepathmapping.html#cfn-apigateway-basepathmapping-restapiid
	RestAPIID *StringExpr `json:"RestApiId,omitempty"`
	// Stage docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-basepathmapping.html#cfn-apigateway-basepathmapping-stage
	Stage *StringExpr `json:"Stage,omitempty"`
}

// CfnResourceType returns AWS::ApiGateway::BasePathMapping to implement the ResourceProperties interface
func (s APIGatewayBasePathMapping) CfnResourceType() string {

	return "AWS::ApiGateway::BasePathMapping"
}

// APIGatewayClientCertificate represents the AWS::ApiGateway::ClientCertificate CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-clientcertificate.html
type APIGatewayClientCertificate struct {
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-clientcertificate.html#cfn-apigateway-clientcertificate-description
	Description *StringExpr `json:"Description,omitempty"`
}

// CfnResourceType returns AWS::ApiGateway::ClientCertificate to implement the ResourceProperties interface
func (s APIGatewayClientCertificate) CfnResourceType() string {

	return "AWS::ApiGateway::ClientCertificate"
}

// APIGatewayDeployment represents the AWS::ApiGateway::Deployment CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html
type APIGatewayDeployment struct {
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html#cfn-apigateway-deployment-description
	Description *StringExpr `json:"Description,omitempty"`
	// RestAPIID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html#cfn-apigateway-deployment-restapiid
	RestAPIID *StringExpr `json:"RestApiId,omitempty" validate:"dive,required"`
	// StageDescription docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html#cfn-apigateway-deployment-stagedescription
	StageDescription *APIGatewayDeploymentStageDescription `json:"StageDescription,omitempty"`
	// StageName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html#cfn-apigateway-deployment-stagename
	StageName *StringExpr `json:"StageName,omitempty"`
}

// CfnResourceType returns AWS::ApiGateway::Deployment to implement the ResourceProperties interface
func (s APIGatewayDeployment) CfnResourceType() string {

	return "AWS::ApiGateway::Deployment"
}

// APIGatewayMethod represents the AWS::ApiGateway::Method CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html
type APIGatewayMethod struct {
	// APIKeyRequired docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-apikeyrequired
	APIKeyRequired *BoolExpr `json:"ApiKeyRequired,omitempty"`
	// AuthorizationType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-authorizationtype
	AuthorizationType *StringExpr `json:"AuthorizationType,omitempty"`
	// AuthorizerID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-authorizerid
	AuthorizerID *StringExpr `json:"AuthorizerId,omitempty"`
	// HTTPMethod docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-httpmethod
	HTTPMethod *StringExpr `json:"HttpMethod,omitempty" validate:"dive,required"`
	// Integration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-integration
	Integration *APIGatewayMethodIntegration `json:"Integration,omitempty"`
	// MethodResponses docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-methodresponses
	MethodResponses *APIGatewayMethodMethodResponseList `json:"MethodResponses,omitempty"`
	// RequestModels docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-requestmodels
	RequestModels interface{} `json:"RequestModels,omitempty"`
	// RequestParameters docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-requestparameters
	RequestParameters interface{} `json:"RequestParameters,omitempty"`
	// ResourceID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-resourceid
	ResourceID *StringExpr `json:"ResourceId,omitempty"`
	// RestAPIID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-restapiid
	RestAPIID *StringExpr `json:"RestApiId,omitempty"`
}

// CfnResourceType returns AWS::ApiGateway::Method to implement the ResourceProperties interface
func (s APIGatewayMethod) CfnResourceType() string {

	return "AWS::ApiGateway::Method"
}

// APIGatewayModel represents the AWS::ApiGateway::Model CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-model.html
type APIGatewayModel struct {
	// ContentType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-model.html#cfn-apigateway-model-contenttype
	ContentType *StringExpr `json:"ContentType,omitempty"`
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-model.html#cfn-apigateway-model-description
	Description *StringExpr `json:"Description,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-model.html#cfn-apigateway-model-name
	Name *StringExpr `json:"Name,omitempty"`
	// RestAPIID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-model.html#cfn-apigateway-model-restapiid
	RestAPIID *StringExpr `json:"RestApiId,omitempty" validate:"dive,required"`
	// Schema docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-model.html#cfn-apigateway-model-schema
	Schema interface{} `json:"Schema,omitempty"`
}

// CfnResourceType returns AWS::ApiGateway::Model to implement the ResourceProperties interface
func (s APIGatewayModel) CfnResourceType() string {

	return "AWS::ApiGateway::Model"
}

// APIGatewayResource represents the AWS::ApiGateway::Resource CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html
type APIGatewayResource struct {
	// ParentID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html#cfn-apigateway-resource-parentid
	ParentID *StringExpr `json:"ParentId,omitempty" validate:"dive,required"`
	// PathPart docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html#cfn-apigateway-resource-pathpart
	PathPart *StringExpr `json:"PathPart,omitempty" validate:"dive,required"`
	// RestAPIID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html#cfn-apigateway-resource-restapiid
	RestAPIID *StringExpr `json:"RestApiId,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::ApiGateway::Resource to implement the ResourceProperties interface
func (s APIGatewayResource) CfnResourceType() string {

	return "AWS::ApiGateway::Resource"
}

// APIGatewayRestAPI represents the AWS::ApiGateway::RestApi CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html
type APIGatewayRestAPI struct {
	// Body docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-body
	Body interface{} `json:"Body,omitempty"`
	// BodyS3Location docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-bodys3location
	BodyS3Location *APIGatewayRestAPIS3Location `json:"BodyS3Location,omitempty"`
	// CloneFrom docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-clonefrom
	CloneFrom *StringExpr `json:"CloneFrom,omitempty"`
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-description
	Description *StringExpr `json:"Description,omitempty"`
	// FailOnWarnings docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-failonwarnings
	FailOnWarnings *BoolExpr `json:"FailOnWarnings,omitempty"`
	// Mode docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-mode
	Mode *StringExpr `json:"Mode,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-name
	Name *StringExpr `json:"Name,omitempty"`
	// Parameters docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-parameters
	Parameters interface{} `json:"Parameters,omitempty"`
}

// CfnResourceType returns AWS::ApiGateway::RestApi to implement the ResourceProperties interface
func (s APIGatewayRestAPI) CfnResourceType() string {

	return "AWS::ApiGateway::RestApi"
}

// APIGatewayStage represents the AWS::ApiGateway::Stage CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html
type APIGatewayStage struct {
	// CacheClusterEnabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-cacheclusterenabled
	CacheClusterEnabled *BoolExpr `json:"CacheClusterEnabled,omitempty"`
	// CacheClusterSize docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-cacheclustersize
	CacheClusterSize *StringExpr `json:"CacheClusterSize,omitempty"`
	// ClientCertificateID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-clientcertificateid
	ClientCertificateID *StringExpr `json:"ClientCertificateId,omitempty"`
	// DeploymentID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-deploymentid
	DeploymentID *StringExpr `json:"DeploymentId,omitempty"`
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-description
	Description *StringExpr `json:"Description,omitempty"`
	// MethodSettings docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-methodsettings
	MethodSettings *APIGatewayStageMethodSettingList `json:"MethodSettings,omitempty"`
	// RestAPIID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-restapiid
	RestAPIID *StringExpr `json:"RestApiId,omitempty"`
	// StageName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-stagename
	StageName *StringExpr `json:"StageName,omitempty"`
	// Variables docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-variables
	Variables interface{} `json:"Variables,omitempty"`
}

// CfnResourceType returns AWS::ApiGateway::Stage to implement the ResourceProperties interface
func (s APIGatewayStage) CfnResourceType() string {

	return "AWS::ApiGateway::Stage"
}

// APIGatewayUsagePlan represents the AWS::ApiGateway::UsagePlan CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html
type APIGatewayUsagePlan struct {
	// APIStages docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html#cfn-apigateway-usageplan-apistages
	APIStages *APIGatewayUsagePlanAPIStageList `json:"ApiStages,omitempty"`
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html#cfn-apigateway-usageplan-description
	Description *StringExpr `json:"Description,omitempty"`
	// Quota docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html#cfn-apigateway-usageplan-quota
	Quota *APIGatewayUsagePlanQuotaSettings `json:"Quota,omitempty"`
	// Throttle docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html#cfn-apigateway-usageplan-throttle
	Throttle *APIGatewayUsagePlanThrottleSettings `json:"Throttle,omitempty"`
	// UsagePlanName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html#cfn-apigateway-usageplan-usageplanname
	UsagePlanName *StringExpr `json:"UsagePlanName,omitempty"`
}

// CfnResourceType returns AWS::ApiGateway::UsagePlan to implement the ResourceProperties interface
func (s APIGatewayUsagePlan) CfnResourceType() string {

	return "AWS::ApiGateway::UsagePlan"
}

// ApplicationAutoScalingScalableTarget represents the AWS::ApplicationAutoScaling::ScalableTarget CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html
type ApplicationAutoScalingScalableTarget struct {
	// MaxCapacity docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html#cfn-applicationautoscaling-scalabletarget-maxcapacity
	MaxCapacity *IntegerExpr `json:"MaxCapacity,omitempty" validate:"dive,required"`
	// MinCapacity docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html#cfn-applicationautoscaling-scalabletarget-mincapacity
	MinCapacity *IntegerExpr `json:"MinCapacity,omitempty" validate:"dive,required"`
	// ResourceID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html#cfn-applicationautoscaling-scalabletarget-resourceid
	ResourceID *StringExpr `json:"ResourceId,omitempty" validate:"dive,required"`
	// RoleARN docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html#cfn-applicationautoscaling-scalabletarget-rolearn
	RoleARN *StringExpr `json:"RoleARN,omitempty" validate:"dive,required"`
	// ScalableDimension docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html#cfn-applicationautoscaling-scalabletarget-scalabledimension
	ScalableDimension *StringExpr `json:"ScalableDimension,omitempty" validate:"dive,required"`
	// ServiceNamespace docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html#cfn-applicationautoscaling-scalabletarget-servicenamespace
	ServiceNamespace *StringExpr `json:"ServiceNamespace,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::ApplicationAutoScaling::ScalableTarget to implement the ResourceProperties interface
func (s ApplicationAutoScalingScalableTarget) CfnResourceType() string {

	return "AWS::ApplicationAutoScaling::ScalableTarget"
}

// ApplicationAutoScalingScalingPolicy represents the AWS::ApplicationAutoScaling::ScalingPolicy CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html
type ApplicationAutoScalingScalingPolicy struct {
	// PolicyName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-policyname
	PolicyName *StringExpr `json:"PolicyName,omitempty" validate:"dive,required"`
	// PolicyType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-policytype
	PolicyType *StringExpr `json:"PolicyType,omitempty" validate:"dive,required"`
	// ResourceID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-resourceid
	ResourceID *StringExpr `json:"ResourceId,omitempty"`
	// ScalableDimension docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-scalabledimension
	ScalableDimension *StringExpr `json:"ScalableDimension,omitempty"`
	// ScalingTargetID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-scalingtargetid
	ScalingTargetID *StringExpr `json:"ScalingTargetId,omitempty"`
	// ServiceNamespace docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-servicenamespace
	ServiceNamespace *StringExpr `json:"ServiceNamespace,omitempty"`
	// StepScalingPolicyConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration
	StepScalingPolicyConfiguration *ApplicationAutoScalingScalingPolicyStepScalingPolicyConfiguration `json:"StepScalingPolicyConfiguration,omitempty"`
}

// CfnResourceType returns AWS::ApplicationAutoScaling::ScalingPolicy to implement the ResourceProperties interface
func (s ApplicationAutoScalingScalingPolicy) CfnResourceType() string {

	return "AWS::ApplicationAutoScaling::ScalingPolicy"
}

// AutoScalingAutoScalingGroup represents the AWS::AutoScaling::AutoScalingGroup CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html
type AutoScalingAutoScalingGroup struct {
	// AvailabilityZones docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-availabilityzones
	AvailabilityZones *StringListExpr `json:"AvailabilityZones,omitempty"`
	// Cooldown docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-cooldown
	Cooldown *StringExpr `json:"Cooldown,omitempty"`
	// DesiredCapacity docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-desiredcapacity
	DesiredCapacity *StringExpr `json:"DesiredCapacity,omitempty"`
	// HealthCheckGracePeriod docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-healthcheckgraceperiod
	HealthCheckGracePeriod *IntegerExpr `json:"HealthCheckGracePeriod,omitempty"`
	// HealthCheckType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-healthchecktype
	HealthCheckType *StringExpr `json:"HealthCheckType,omitempty"`
	// InstanceID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-instanceid
	InstanceID *StringExpr `json:"InstanceId,omitempty"`
	// LaunchConfigurationName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-launchconfigurationname
	LaunchConfigurationName *StringExpr `json:"LaunchConfigurationName,omitempty"`
	// LoadBalancerNames docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-loadbalancernames
	LoadBalancerNames *StringListExpr `json:"LoadBalancerNames,omitempty"`
	// MaxSize docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-maxsize
	MaxSize *StringExpr `json:"MaxSize,omitempty" validate:"dive,required"`
	// MetricsCollection docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-metricscollection
	MetricsCollection *AutoScalingAutoScalingGroupMetricsCollection `json:"MetricsCollection,omitempty"`
	// MinSize docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-minsize
	MinSize *StringExpr `json:"MinSize,omitempty" validate:"dive,required"`
	// NotificationConfigurations docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-notificationconfigurations
	NotificationConfigurations *AutoScalingAutoScalingGroupNotificationConfigurations `json:"NotificationConfigurations,omitempty"`
	// PlacementGroup docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-placementgroup
	PlacementGroup *StringExpr `json:"PlacementGroup,omitempty"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-tags
	Tags *AutoScalingAutoScalingGroupTagPropertyList `json:"Tags,omitempty"`
	// TargetGroupARNs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-targetgrouparns
	TargetGroupARNs *StringListExpr `json:"TargetGroupARNs,omitempty"`
	// TerminationPolicies docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-termpolicy
	TerminationPolicies *StringListExpr `json:"TerminationPolicies,omitempty"`
	// VPCZoneIDentifier docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-vpczoneidentifier
	VPCZoneIDentifier *StringListExpr `json:"VPCZoneIdentifier,omitempty"`
}

// CfnResourceType returns AWS::AutoScaling::AutoScalingGroup to implement the ResourceProperties interface
func (s AutoScalingAutoScalingGroup) CfnResourceType() string {

	return "AWS::AutoScaling::AutoScalingGroup"
}

// AutoScalingLaunchConfiguration represents the AWS::AutoScaling::LaunchConfiguration CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html
type AutoScalingLaunchConfiguration struct {
	// AssociatePublicIPAddress docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cf-as-launchconfig-associatepubip
	AssociatePublicIPAddress *BoolExpr `json:"AssociatePublicIpAddress,omitempty"`
	// BlockDeviceMappings docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-blockdevicemappings
	BlockDeviceMappings *AutoScalingLaunchConfigurationBlockDeviceMappingList `json:"BlockDeviceMappings,omitempty"`
	// ClassicLinkVPCID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-classiclinkvpcid
	ClassicLinkVPCID *StringExpr `json:"ClassicLinkVPCId,omitempty"`
	// ClassicLinkVPCSecurityGroups docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-classiclinkvpcsecuritygroups
	ClassicLinkVPCSecurityGroups *StringListExpr `json:"ClassicLinkVPCSecurityGroups,omitempty"`
	// EbsOptimized docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-ebsoptimized
	EbsOptimized *BoolExpr `json:"EbsOptimized,omitempty"`
	// IamInstanceProfile docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-iaminstanceprofile
	IamInstanceProfile *StringExpr `json:"IamInstanceProfile,omitempty"`
	// ImageID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-imageid
	ImageID *StringExpr `json:"ImageId,omitempty" validate:"dive,required"`
	// InstanceID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-instanceid
	InstanceID *StringExpr `json:"InstanceId,omitempty"`
	// InstanceMonitoring docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-instancemonitoring
	InstanceMonitoring *BoolExpr `json:"InstanceMonitoring,omitempty"`
	// InstanceType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-instancetype
	InstanceType *StringExpr `json:"InstanceType,omitempty" validate:"dive,required"`
	// KernelID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-kernelid
	KernelID *StringExpr `json:"KernelId,omitempty"`
	// KeyName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-keyname
	KeyName *StringExpr `json:"KeyName,omitempty"`
	// PlacementTenancy docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-placementtenancy
	PlacementTenancy *StringExpr `json:"PlacementTenancy,omitempty"`
	// RAMDiskID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-ramdiskid
	RAMDiskID *StringExpr `json:"RamDiskId,omitempty"`
	// SecurityGroups docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-securitygroups
	SecurityGroups *StringListExpr `json:"SecurityGroups,omitempty"`
	// SpotPrice docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-spotprice
	SpotPrice *StringExpr `json:"SpotPrice,omitempty"`
	// UserData docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-userdata
	UserData *StringExpr `json:"UserData,omitempty"`
}

// CfnResourceType returns AWS::AutoScaling::LaunchConfiguration to implement the ResourceProperties interface
func (s AutoScalingLaunchConfiguration) CfnResourceType() string {

	return "AWS::AutoScaling::LaunchConfiguration"
}

// AutoScalingLifecycleHook represents the AWS::AutoScaling::LifecycleHook CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html
type AutoScalingLifecycleHook struct {
	// AutoScalingGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html#cfn-as-lifecyclehook-autoscalinggroupname
	AutoScalingGroupName *StringExpr `json:"AutoScalingGroupName,omitempty" validate:"dive,required"`
	// DefaultResult docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html#cfn-as-lifecyclehook-defaultresult
	DefaultResult *StringExpr `json:"DefaultResult,omitempty"`
	// HeartbeatTimeout docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html#cfn-as-lifecyclehook-heartbeattimeout
	HeartbeatTimeout *IntegerExpr `json:"HeartbeatTimeout,omitempty"`
	// LifecycleTransition docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html#cfn-as-lifecyclehook-lifecycletransition
	LifecycleTransition *StringExpr `json:"LifecycleTransition,omitempty" validate:"dive,required"`
	// NotificationMetadata docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html#cfn-as-lifecyclehook-notificationmetadata
	NotificationMetadata *StringExpr `json:"NotificationMetadata,omitempty"`
	// NotificationTargetARN docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html#cfn-as-lifecyclehook-notificationtargetarn
	NotificationTargetARN *StringExpr `json:"NotificationTargetARN,omitempty"`
	// RoleARN docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html#cfn-as-lifecyclehook-rolearn
	RoleARN *StringExpr `json:"RoleARN,omitempty"`
}

// CfnResourceType returns AWS::AutoScaling::LifecycleHook to implement the ResourceProperties interface
func (s AutoScalingLifecycleHook) CfnResourceType() string {

	return "AWS::AutoScaling::LifecycleHook"
}

// AutoScalingScalingPolicy represents the AWS::AutoScaling::ScalingPolicy CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html
type AutoScalingScalingPolicy struct {
	// AdjustmentType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-adjustmenttype
	AdjustmentType *StringExpr `json:"AdjustmentType,omitempty" validate:"dive,required"`
	// AutoScalingGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-autoscalinggroupname
	AutoScalingGroupName *StringExpr `json:"AutoScalingGroupName,omitempty" validate:"dive,required"`
	// Cooldown docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-cooldown
	Cooldown *StringExpr `json:"Cooldown,omitempty"`
	// EstimatedInstanceWarmup docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-estimatedinstancewarmup
	EstimatedInstanceWarmup *IntegerExpr `json:"EstimatedInstanceWarmup,omitempty"`
	// MetricAggregationType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-metricaggregationtype
	MetricAggregationType *StringExpr `json:"MetricAggregationType,omitempty"`
	// MinAdjustmentMagnitude docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-minadjustmentmagnitude
	MinAdjustmentMagnitude *IntegerExpr `json:"MinAdjustmentMagnitude,omitempty"`
	// PolicyType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-policytype
	PolicyType *StringExpr `json:"PolicyType,omitempty"`
	// ScalingAdjustment docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-scalingadjustment
	ScalingAdjustment *IntegerExpr `json:"ScalingAdjustment,omitempty"`
	// StepAdjustments docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-stepadjustments
	StepAdjustments *AutoScalingScalingPolicyStepAdjustmentList `json:"StepAdjustments,omitempty"`
}

// CfnResourceType returns AWS::AutoScaling::ScalingPolicy to implement the ResourceProperties interface
func (s AutoScalingScalingPolicy) CfnResourceType() string {

	return "AWS::AutoScaling::ScalingPolicy"
}

// AutoScalingScheduledAction represents the AWS::AutoScaling::ScheduledAction CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html
type AutoScalingScheduledAction struct {
	// AutoScalingGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-asgname
	AutoScalingGroupName *StringExpr `json:"AutoScalingGroupName,omitempty" validate:"dive,required"`
	// DesiredCapacity docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-desiredcapacity
	DesiredCapacity *IntegerExpr `json:"DesiredCapacity,omitempty"`
	// EndTime docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-endtime
	EndTime time.Time `json:"EndTime,omitempty"`
	// MaxSize docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-maxsize
	MaxSize *IntegerExpr `json:"MaxSize,omitempty"`
	// MinSize docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-minsize
	MinSize *IntegerExpr `json:"MinSize,omitempty"`
	// Recurrence docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-recurrence
	Recurrence *StringExpr `json:"Recurrence,omitempty"`
	// StartTime docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-starttime
	StartTime time.Time `json:"StartTime,omitempty"`
}

// CfnResourceType returns AWS::AutoScaling::ScheduledAction to implement the ResourceProperties interface
func (s AutoScalingScheduledAction) CfnResourceType() string {

	return "AWS::AutoScaling::ScheduledAction"
}

// CertificateManagerCertificate represents the AWS::CertificateManager::Certificate CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html
type CertificateManagerCertificate struct {
	// DomainName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html#cfn-certificatemanager-certificate-domainname
	DomainName *StringExpr `json:"DomainName,omitempty" validate:"dive,required"`
	// DomainValidationOptions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html#cfn-certificatemanager-certificate-domainvalidationoptions
	DomainValidationOptions *CertificateManagerCertificateDomainValidationOptionList `json:"DomainValidationOptions,omitempty"`
	// SubjectAlternativeNames docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html#cfn-certificatemanager-certificate-subjectalternativenames
	SubjectAlternativeNames *StringListExpr `json:"SubjectAlternativeNames,omitempty"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html#cfn-certificatemanager-certificate-tags
	Tags *TagList `json:"Tags,omitempty"`
}

// CfnResourceType returns AWS::CertificateManager::Certificate to implement the ResourceProperties interface
func (s CertificateManagerCertificate) CfnResourceType() string {

	return "AWS::CertificateManager::Certificate"
}

// CloudFormationCustomResource represents the AWS::CloudFormation::CustomResource CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cfn-customresource.html
type CloudFormationCustomResource struct {
	// ServiceToken docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cfn-customresource.html#cfn-customresource-servicetoken
	ServiceToken *StringExpr `json:"ServiceToken,omitempty" validate:"dive,required"`

	// The user-defined Custom::* name to use for the resource.  If empty,
	// the default "AWS::CloudFormation::CustomResource" value will be used.
	// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-custom-resources.html
	ResourceTypeName string
}

// CfnResourceType returns AWS::CloudFormation::CustomResource to implement the ResourceProperties interface
func (s CloudFormationCustomResource) CfnResourceType() string {
	if "" != s.ResourceTypeName {
		return s.ResourceTypeName
	}
	return "AWS::CloudFormation::CustomResource"
}

// CloudFormationStack represents the AWS::CloudFormation::Stack CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stack.html
type CloudFormationStack struct {
	// NotificationARNs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stack.html#cfn-cloudformation-stack-notificationarns
	NotificationARNs *StringListExpr `json:"NotificationARNs,omitempty"`
	// Parameters docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stack.html#cfn-cloudformation-stack-parameters
	Parameters interface{} `json:"Parameters,omitempty"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stack.html#cfn-cloudformation-stack-tags
	Tags *TagList `json:"Tags,omitempty"`
	// TemplateURL docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stack.html#cfn-cloudformation-stack-templateurl
	TemplateURL *StringExpr `json:"TemplateURL,omitempty" validate:"dive,required"`
	// TimeoutInMinutes docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stack.html#cfn-cloudformation-stack-timeoutinminutes
	TimeoutInMinutes *IntegerExpr `json:"TimeoutInMinutes,omitempty"`
}

// CfnResourceType returns AWS::CloudFormation::Stack to implement the ResourceProperties interface
func (s CloudFormationStack) CfnResourceType() string {

	return "AWS::CloudFormation::Stack"
}

// CloudFormationWaitCondition represents the AWS::CloudFormation::WaitCondition CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waitcondition.html
type CloudFormationWaitCondition struct {
	// Count docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waitcondition.html#cfn-waitcondition-count
	Count *IntegerExpr `json:"Count,omitempty"`
	// Handle docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waitcondition.html#cfn-waitcondition-handle
	Handle *StringExpr `json:"Handle,omitempty" validate:"dive,required"`
	// Timeout docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waitcondition.html#cfn-waitcondition-timeout
	Timeout *StringExpr `json:"Timeout,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::CloudFormation::WaitCondition to implement the ResourceProperties interface
func (s CloudFormationWaitCondition) CfnResourceType() string {

	return "AWS::CloudFormation::WaitCondition"
}

// CloudFormationWaitConditionHandle represents the AWS::CloudFormation::WaitConditionHandle CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waitconditionhandle.html
type CloudFormationWaitConditionHandle struct {
}

// CfnResourceType returns AWS::CloudFormation::WaitConditionHandle to implement the ResourceProperties interface
func (s CloudFormationWaitConditionHandle) CfnResourceType() string {

	return "AWS::CloudFormation::WaitConditionHandle"
}

// CloudFrontDistribution represents the AWS::CloudFront::Distribution CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution.html
type CloudFrontDistribution struct {
	// DistributionConfig docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution.html#cfn-cloudfront-distribution-distributionconfig
	DistributionConfig *CloudFrontDistributionDistributionConfig `json:"DistributionConfig,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::CloudFront::Distribution to implement the ResourceProperties interface
func (s CloudFrontDistribution) CfnResourceType() string {

	return "AWS::CloudFront::Distribution"
}

// CloudTrailTrail represents the AWS::CloudTrail::Trail CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html
type CloudTrailTrail struct {
	// CloudWatchLogsLogGroupArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-cloudwatchlogsloggrouparn
	CloudWatchLogsLogGroupArn *StringExpr `json:"CloudWatchLogsLogGroupArn,omitempty"`
	// CloudWatchLogsRoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-cloudwatchlogsrolearn
	CloudWatchLogsRoleArn *StringExpr `json:"CloudWatchLogsRoleArn,omitempty"`
	// EnableLogFileValidation docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-enablelogfilevalidation
	EnableLogFileValidation *BoolExpr `json:"EnableLogFileValidation,omitempty"`
	// IncludeGlobalServiceEvents docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-includeglobalserviceevents
	IncludeGlobalServiceEvents *BoolExpr `json:"IncludeGlobalServiceEvents,omitempty"`
	// IsLogging docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-islogging
	IsLogging *BoolExpr `json:"IsLogging,omitempty" validate:"dive,required"`
	// IsMultiRegionTrail docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-ismultiregiontrail
	IsMultiRegionTrail *BoolExpr `json:"IsMultiRegionTrail,omitempty"`
	// KMSKeyID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-kmskeyid
	KMSKeyID *StringExpr `json:"KMSKeyId,omitempty"`
	// S3BucketName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-s3bucketname
	S3BucketName *StringExpr `json:"S3BucketName,omitempty" validate:"dive,required"`
	// S3KeyPrefix docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-s3keyprefix
	S3KeyPrefix *StringExpr `json:"S3KeyPrefix,omitempty"`
	// SnsTopicName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-snstopicname
	SnsTopicName *StringExpr `json:"SnsTopicName,omitempty"`
}

// CfnResourceType returns AWS::CloudTrail::Trail to implement the ResourceProperties interface
func (s CloudTrailTrail) CfnResourceType() string {

	return "AWS::CloudTrail::Trail"
}

// CloudWatchAlarm represents the AWS::CloudWatch::Alarm CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-alarm.html
type CloudWatchAlarm struct {
	// ActionsEnabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-alarm.html#cfn-cloudwatch-alarms-actionsenabled
	ActionsEnabled *BoolExpr `json:"ActionsEnabled,omitempty"`
	// AlarmActions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-alarm.html#cfn-cloudwatch-alarms-alarmactions
	AlarmActions *StringListExpr `json:"AlarmActions,omitempty"`
	// AlarmDescription docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-alarm.html#cfn-cloudwatch-alarms-alarmdescription
	AlarmDescription *StringExpr `json:"AlarmDescription,omitempty"`
	// AlarmName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-alarm.html#cfn-cloudwatch-alarms-alarmname
	AlarmName *StringExpr `json:"AlarmName,omitempty"`
	// ComparisonOperator docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-alarm.html#cfn-cloudwatch-alarms-comparisonoperator
	ComparisonOperator *StringExpr `json:"ComparisonOperator,omitempty" validate:"dive,required"`
	// Dimensions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-alarm.html#cfn-cloudwatch-alarms-dimension
	Dimensions *CloudWatchAlarmDimensionList `json:"Dimensions,omitempty"`
	// EvaluateLowSampleCountPercentile docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-alarm.html#cfn-cloudwatch-alarms-evaluatelowsamplecountpercentile
	EvaluateLowSampleCountPercentile *StringExpr `json:"EvaluateLowSampleCountPercentile,omitempty"`
	// EvaluationPeriods docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-alarm.html#cfn-cloudwatch-alarms-evaluationperiods
	EvaluationPeriods *IntegerExpr `json:"EvaluationPeriods,omitempty" validate:"dive,required"`
	// ExtendedStatistic docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-alarm.html#cfn-cloudwatch-alarms-extendedstatistic
	ExtendedStatistic *StringExpr `json:"ExtendedStatistic,omitempty"`
	// InsufficientDataActions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-alarm.html#cfn-cloudwatch-alarms-insufficientdataactions
	InsufficientDataActions *StringListExpr `json:"InsufficientDataActions,omitempty"`
	// MetricName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-alarm.html#cfn-cloudwatch-alarms-metricname
	MetricName *StringExpr `json:"MetricName,omitempty" validate:"dive,required"`
	// Namespace docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-alarm.html#cfn-cloudwatch-alarms-namespace
	Namespace *StringExpr `json:"Namespace,omitempty" validate:"dive,required"`
	// OKActions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-alarm.html#cfn-cloudwatch-alarms-okactions
	OKActions *StringListExpr `json:"OKActions,omitempty"`
	// Period docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-alarm.html#cfn-cloudwatch-alarms-period
	Period *IntegerExpr `json:"Period,omitempty" validate:"dive,required"`
	// Statistic docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-alarm.html#cfn-cloudwatch-alarms-statistic
	Statistic *StringExpr `json:"Statistic,omitempty"`
	// Threshold docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-alarm.html#cfn-cloudwatch-alarms-threshold
	Threshold *IntegerExpr `json:"Threshold,omitempty" validate:"dive,required"`
	// TreatMissingData docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-alarm.html#cfn-cloudwatch-alarms-treatmissingdata
	TreatMissingData *StringExpr `json:"TreatMissingData,omitempty"`
	// Unit docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-alarm.html#cfn-cloudwatch-alarms-unit
	Unit *StringExpr `json:"Unit,omitempty"`
}

// CfnResourceType returns AWS::CloudWatch::Alarm to implement the ResourceProperties interface
func (s CloudWatchAlarm) CfnResourceType() string {

	return "AWS::CloudWatch::Alarm"
}

// CodeBuildProject represents the AWS::CodeBuild::Project CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html
type CodeBuildProject struct {
	// Artifacts docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-artifacts
	Artifacts *CodeBuildProjectArtifacts `json:"Artifacts,omitempty"`
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-description
	Description *StringExpr `json:"Description,omitempty"`
	// EncryptionKey docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-encryptionkey
	EncryptionKey *StringExpr `json:"EncryptionKey,omitempty"`
	// Environment docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-environment
	Environment *CodeBuildProjectEnvironment `json:"Environment,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-name
	Name *StringExpr `json:"Name,omitempty"`
	// ServiceRole docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-servicerole
	ServiceRole *StringExpr `json:"ServiceRole,omitempty"`
	// Source docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-source
	Source *CodeBuildProjectSource `json:"Source,omitempty"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-tags
	Tags *TagList `json:"Tags,omitempty"`
	// TimeoutInMinutes docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-timeoutinminutes
	TimeoutInMinutes *IntegerExpr `json:"TimeoutInMinutes,omitempty"`
}

// CfnResourceType returns AWS::CodeBuild::Project to implement the ResourceProperties interface
func (s CodeBuildProject) CfnResourceType() string {

	return "AWS::CodeBuild::Project"
}

// CodeCommitRepository represents the AWS::CodeCommit::Repository CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codecommit-repository.html
type CodeCommitRepository struct {
	// RepositoryDescription docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codecommit-repository.html#cfn-codecommit-repository-repositorydescription
	RepositoryDescription *StringExpr `json:"RepositoryDescription,omitempty"`
	// RepositoryName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codecommit-repository.html#cfn-codecommit-repository-repositoryname
	RepositoryName *StringExpr `json:"RepositoryName,omitempty" validate:"dive,required"`
	// Triggers docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codecommit-repository.html#cfn-codecommit-repository-triggers
	Triggers *CodeCommitRepositoryRepositoryTriggerList `json:"Triggers,omitempty"`
}

// CfnResourceType returns AWS::CodeCommit::Repository to implement the ResourceProperties interface
func (s CodeCommitRepository) CfnResourceType() string {

	return "AWS::CodeCommit::Repository"
}

// CodeDeployApplication represents the AWS::CodeDeploy::Application CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-application.html
type CodeDeployApplication struct {
	// ApplicationName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-application.html#cfn-codedeploy-application-applicationname
	ApplicationName *StringExpr `json:"ApplicationName,omitempty"`
}

// CfnResourceType returns AWS::CodeDeploy::Application to implement the ResourceProperties interface
func (s CodeDeployApplication) CfnResourceType() string {

	return "AWS::CodeDeploy::Application"
}

// CodeDeployDeploymentConfig represents the AWS::CodeDeploy::DeploymentConfig CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html
type CodeDeployDeploymentConfig struct {
	// DeploymentConfigName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html#cfn-codedeploy-deploymentconfig-deploymentconfigname
	DeploymentConfigName *StringExpr `json:"DeploymentConfigName,omitempty"`
	// MinimumHealthyHosts docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html#cfn-codedeploy-deploymentconfig-minimumhealthyhosts
	MinimumHealthyHosts *CodeDeployDeploymentConfigMinimumHealthyHosts `json:"MinimumHealthyHosts,omitempty"`
}

// CfnResourceType returns AWS::CodeDeploy::DeploymentConfig to implement the ResourceProperties interface
func (s CodeDeployDeploymentConfig) CfnResourceType() string {

	return "AWS::CodeDeploy::DeploymentConfig"
}

// CodeDeployDeploymentGroup represents the AWS::CodeDeploy::DeploymentGroup CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html
type CodeDeployDeploymentGroup struct {
	// AlarmConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-alarmconfiguration
	AlarmConfiguration *CodeDeployDeploymentGroupAlarmConfiguration `json:"AlarmConfiguration,omitempty"`
	// ApplicationName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-applicationname
	ApplicationName *StringExpr `json:"ApplicationName,omitempty" validate:"dive,required"`
	// AutoScalingGroups docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-autoscalinggroups
	AutoScalingGroups *StringListExpr `json:"AutoScalingGroups,omitempty"`
	// Deployment docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-deployment
	Deployment *CodeDeployDeploymentGroupDeployment `json:"Deployment,omitempty"`
	// DeploymentConfigName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-deploymentconfigname
	DeploymentConfigName *StringExpr `json:"DeploymentConfigName,omitempty"`
	// DeploymentGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-deploymentgroupname
	DeploymentGroupName *StringExpr `json:"DeploymentGroupName,omitempty"`
	// Ec2TagFilters docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-ec2tagfilters
	Ec2TagFilters *CodeDeployDeploymentGroupEC2TagFilterList `json:"Ec2TagFilters,omitempty"`
	// OnPremisesInstanceTagFilters docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-onpremisesinstancetagfilters
	OnPremisesInstanceTagFilters *CodeDeployDeploymentGroupTagFilterList `json:"OnPremisesInstanceTagFilters,omitempty"`
	// ServiceRoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-servicerolearn
	ServiceRoleArn *StringExpr `json:"ServiceRoleArn,omitempty" validate:"dive,required"`
	// TriggerConfigurations docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-triggerconfigurations
	TriggerConfigurations *CodeDeployDeploymentGroupTriggerConfigList `json:"TriggerConfigurations,omitempty"`
}

// CfnResourceType returns AWS::CodeDeploy::DeploymentGroup to implement the ResourceProperties interface
func (s CodeDeployDeploymentGroup) CfnResourceType() string {

	return "AWS::CodeDeploy::DeploymentGroup"
}

// CodePipelineCustomActionType represents the AWS::CodePipeline::CustomActionType CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html
type CodePipelineCustomActionType struct {
	// Category docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-category
	Category *StringExpr `json:"Category,omitempty" validate:"dive,required"`
	// ConfigurationProperties docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-configurationproperties
	ConfigurationProperties *CodePipelineCustomActionTypeConfigurationPropertiesList `json:"ConfigurationProperties,omitempty"`
	// InputArtifactDetails docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-inputartifactdetails
	InputArtifactDetails *CodePipelineCustomActionTypeArtifactDetails `json:"InputArtifactDetails,omitempty" validate:"dive,required"`
	// OutputArtifactDetails docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-outputartifactdetails
	OutputArtifactDetails *CodePipelineCustomActionTypeArtifactDetails `json:"OutputArtifactDetails,omitempty" validate:"dive,required"`
	// Provider docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-provider
	Provider *StringExpr `json:"Provider,omitempty" validate:"dive,required"`
	// Settings docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-settings
	Settings *CodePipelineCustomActionTypeSettings `json:"Settings,omitempty"`
	// Version docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-version
	Version *StringExpr `json:"Version,omitempty"`
}

// CfnResourceType returns AWS::CodePipeline::CustomActionType to implement the ResourceProperties interface
func (s CodePipelineCustomActionType) CfnResourceType() string {

	return "AWS::CodePipeline::CustomActionType"
}

// CodePipelinePipeline represents the AWS::CodePipeline::Pipeline CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html
type CodePipelinePipeline struct {
	// ArtifactStore docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-artifactstore
	ArtifactStore *CodePipelinePipelineArtifactStore `json:"ArtifactStore,omitempty" validate:"dive,required"`
	// DisableInboundStageTransitions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-disableinboundstagetransitions
	DisableInboundStageTransitions *CodePipelinePipelineStageTransitionList `json:"DisableInboundStageTransitions,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-name
	Name *StringExpr `json:"Name,omitempty"`
	// RestartExecutionOnUpdate docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-restartexecutiononupdate
	RestartExecutionOnUpdate *BoolExpr `json:"RestartExecutionOnUpdate,omitempty"`
	// RoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-rolearn
	RoleArn *StringExpr `json:"RoleArn,omitempty" validate:"dive,required"`
	// Stages docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-stages
	Stages *CodePipelinePipelineStageDeclarationList `json:"Stages,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::CodePipeline::Pipeline to implement the ResourceProperties interface
func (s CodePipelinePipeline) CfnResourceType() string {

	return "AWS::CodePipeline::Pipeline"
}

// CognitoIDentityPool represents the AWS::Cognito::IdentityPool CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html
type CognitoIDentityPool struct {
	// AllowUnauthenticatedIDentities docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-allowunauthenticatedidentities
	AllowUnauthenticatedIDentities *BoolExpr `json:"AllowUnauthenticatedIdentities,omitempty" validate:"dive,required"`
	// CognitoEvents docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-cognitoevents
	CognitoEvents interface{} `json:"CognitoEvents,omitempty"`
	// CognitoIDentityProviders docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-cognitoidentityproviders
	CognitoIDentityProviders *CognitoIDentityPoolCognitoIDentityProviderList `json:"CognitoIdentityProviders,omitempty"`
	// CognitoStreams docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-cognitostreams
	CognitoStreams *CognitoIDentityPoolCognitoStreams `json:"CognitoStreams,omitempty"`
	// DeveloperProviderName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-developerprovidername
	DeveloperProviderName *StringExpr `json:"DeveloperProviderName,omitempty"`
	// IDentityPoolName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-identitypoolname
	IDentityPoolName *StringExpr `json:"IdentityPoolName,omitempty"`
	// OpenIDConnectProviderARNs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-openidconnectproviderarns
	OpenIDConnectProviderARNs *StringListExpr `json:"OpenIdConnectProviderARNs,omitempty"`
	// PushSync docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-pushsync
	PushSync *CognitoIDentityPoolPushSync `json:"PushSync,omitempty"`
	// SamlProviderARNs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-samlproviderarns
	SamlProviderARNs *StringListExpr `json:"SamlProviderARNs,omitempty"`
	// SupportedLoginProviders docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-supportedloginproviders
	SupportedLoginProviders interface{} `json:"SupportedLoginProviders,omitempty"`
}

// CfnResourceType returns AWS::Cognito::IdentityPool to implement the ResourceProperties interface
func (s CognitoIDentityPool) CfnResourceType() string {

	return "AWS::Cognito::IdentityPool"
}

// CognitoIDentityPoolRoleAttachment represents the AWS::Cognito::IdentityPoolRoleAttachment CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolroleattachment.html
type CognitoIDentityPoolRoleAttachment struct {
	// IDentityPoolID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolroleattachment.html#cfn-cognito-identitypoolroleattachment-identitypoolid
	IDentityPoolID *StringExpr `json:"IdentityPoolId,omitempty" validate:"dive,required"`
	// RoleMappings docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolroleattachment.html#cfn-cognito-identitypoolroleattachment-rolemappings
	RoleMappings interface{} `json:"RoleMappings,omitempty"`
	// Roles docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolroleattachment.html#cfn-cognito-identitypoolroleattachment-roles
	Roles interface{} `json:"Roles,omitempty"`
}

// CfnResourceType returns AWS::Cognito::IdentityPoolRoleAttachment to implement the ResourceProperties interface
func (s CognitoIDentityPoolRoleAttachment) CfnResourceType() string {

	return "AWS::Cognito::IdentityPoolRoleAttachment"
}

// CognitoUserPool represents the AWS::Cognito::UserPool CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html
type CognitoUserPool struct {
	// AdminCreateUserConfig docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-admincreateuserconfig
	AdminCreateUserConfig *CognitoUserPoolAdminCreateUserConfig `json:"AdminCreateUserConfig,omitempty"`
	// AliasAttributes docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-aliasattributes
	AliasAttributes *StringListExpr `json:"AliasAttributes,omitempty"`
	// AutoVerifiedAttributes docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-autoverifiedattributes
	AutoVerifiedAttributes *StringListExpr `json:"AutoVerifiedAttributes,omitempty"`
	// DeviceConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-deviceconfiguration
	DeviceConfiguration *CognitoUserPoolDeviceConfiguration `json:"DeviceConfiguration,omitempty"`
	// EmailConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-emailconfiguration
	EmailConfiguration *CognitoUserPoolEmailConfiguration `json:"EmailConfiguration,omitempty"`
	// EmailVerificationMessage docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-emailverificationmessage
	EmailVerificationMessage *StringExpr `json:"EmailVerificationMessage,omitempty"`
	// EmailVerificationSubject docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-emailverificationsubject
	EmailVerificationSubject *StringExpr `json:"EmailVerificationSubject,omitempty"`
	// LambdaConfig docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-lambdaconfig
	LambdaConfig *CognitoUserPoolLambdaConfig `json:"LambdaConfig,omitempty"`
	// MfaConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-mfaconfiguration
	MfaConfiguration *StringExpr `json:"MfaConfiguration,omitempty"`
	// Policies docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-policies
	Policies *CognitoUserPoolPolicies `json:"Policies,omitempty"`
	// Schema docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-schema
	Schema *CognitoUserPoolSchemaAttributeList `json:"Schema,omitempty"`
	// SmsAuthenticationMessage docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-smsauthenticationmessage
	SmsAuthenticationMessage *StringExpr `json:"SmsAuthenticationMessage,omitempty"`
	// SmsConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-smsconfiguration
	SmsConfiguration *CognitoUserPoolSmsConfiguration `json:"SmsConfiguration,omitempty"`
	// SmsVerificationMessage docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-smsverificationmessage
	SmsVerificationMessage *StringExpr `json:"SmsVerificationMessage,omitempty"`
	// UserPoolName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-userpoolname
	UserPoolName *StringExpr `json:"UserPoolName,omitempty"`
	// UserPoolTags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-userpooltags
	UserPoolTags interface{} `json:"UserPoolTags,omitempty"`
}

// CfnResourceType returns AWS::Cognito::UserPool to implement the ResourceProperties interface
func (s CognitoUserPool) CfnResourceType() string {

	return "AWS::Cognito::UserPool"
}

// CognitoUserPoolClient represents the AWS::Cognito::UserPoolClient CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html
type CognitoUserPoolClient struct {
	// ClientName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-clientname
	ClientName *StringExpr `json:"ClientName,omitempty"`
	// ExplicitAuthFlows docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-explicitauthflows
	ExplicitAuthFlows *StringListExpr `json:"ExplicitAuthFlows,omitempty"`
	// GenerateSecret docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-generatesecret
	GenerateSecret *BoolExpr `json:"GenerateSecret,omitempty"`
	// ReadAttributes docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-readattributes
	ReadAttributes *StringListExpr `json:"ReadAttributes,omitempty"`
	// RefreshTokenValidity docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-refreshtokenvalidity
	RefreshTokenValidity *IntegerExpr `json:"RefreshTokenValidity,omitempty"`
	// UserPoolID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-userpoolid
	UserPoolID *StringExpr `json:"UserPoolId,omitempty" validate:"dive,required"`
	// WriteAttributes docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-writeattributes
	WriteAttributes *StringListExpr `json:"WriteAttributes,omitempty"`
}

// CfnResourceType returns AWS::Cognito::UserPoolClient to implement the ResourceProperties interface
func (s CognitoUserPoolClient) CfnResourceType() string {

	return "AWS::Cognito::UserPoolClient"
}

// CognitoUserPoolGroup represents the AWS::Cognito::UserPoolGroup CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolgroup.html
type CognitoUserPoolGroup struct {
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolgroup.html#cfn-cognito-userpoolgroup-description
	Description *StringExpr `json:"Description,omitempty"`
	// GroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolgroup.html#cfn-cognito-userpoolgroup-groupname
	GroupName *StringExpr `json:"GroupName,omitempty"`
	// Precedence docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolgroup.html#cfn-cognito-userpoolgroup-precedence
	Precedence *IntegerExpr `json:"Precedence,omitempty"`
	// RoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolgroup.html#cfn-cognito-userpoolgroup-rolearn
	RoleArn *StringExpr `json:"RoleArn,omitempty"`
	// UserPoolID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolgroup.html#cfn-cognito-userpoolgroup-userpoolid
	UserPoolID *StringExpr `json:"UserPoolId,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::Cognito::UserPoolGroup to implement the ResourceProperties interface
func (s CognitoUserPoolGroup) CfnResourceType() string {

	return "AWS::Cognito::UserPoolGroup"
}

// CognitoUserPoolUser represents the AWS::Cognito::UserPoolUser CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html
type CognitoUserPoolUser struct {
	// DesiredDeliveryMediums docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-desireddeliverymediums
	DesiredDeliveryMediums *StringListExpr `json:"DesiredDeliveryMediums,omitempty"`
	// ForceAliasCreation docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-forcealiascreation
	ForceAliasCreation *BoolExpr `json:"ForceAliasCreation,omitempty"`
	// MessageAction docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-messageaction
	MessageAction *StringExpr `json:"MessageAction,omitempty"`
	// UserAttributes docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-userattributes
	UserAttributes *CognitoUserPoolUserAttributeTypeList `json:"UserAttributes,omitempty"`
	// UserPoolID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-userpoolid
	UserPoolID *StringExpr `json:"UserPoolId,omitempty" validate:"dive,required"`
	// Username docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-username
	Username *StringExpr `json:"Username,omitempty"`
	// ValidationData docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-validationdata
	ValidationData *CognitoUserPoolUserAttributeTypeList `json:"ValidationData,omitempty"`
}

// CfnResourceType returns AWS::Cognito::UserPoolUser to implement the ResourceProperties interface
func (s CognitoUserPoolUser) CfnResourceType() string {

	return "AWS::Cognito::UserPoolUser"
}

// CognitoUserPoolUserToGroupAttachment represents the AWS::Cognito::UserPoolUserToGroupAttachment CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolusertogroupattachment.html
type CognitoUserPoolUserToGroupAttachment struct {
	// GroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolusertogroupattachment.html#cfn-cognito-userpoolusertogroupattachment-groupname
	GroupName *StringExpr `json:"GroupName,omitempty" validate:"dive,required"`
	// UserPoolID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolusertogroupattachment.html#cfn-cognito-userpoolusertogroupattachment-userpoolid
	UserPoolID *StringExpr `json:"UserPoolId,omitempty" validate:"dive,required"`
	// Username docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolusertogroupattachment.html#cfn-cognito-userpoolusertogroupattachment-username
	Username *StringExpr `json:"Username,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::Cognito::UserPoolUserToGroupAttachment to implement the ResourceProperties interface
func (s CognitoUserPoolUserToGroupAttachment) CfnResourceType() string {

	return "AWS::Cognito::UserPoolUserToGroupAttachment"
}

// ConfigConfigRule represents the AWS::Config::ConfigRule CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configrule.html
type ConfigConfigRule struct {
	// ConfigRuleName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configrule.html#cfn-config-configrule-configrulename
	ConfigRuleName *StringExpr `json:"ConfigRuleName,omitempty"`
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configrule.html#cfn-config-configrule-description
	Description *StringExpr `json:"Description,omitempty"`
	// InputParameters docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configrule.html#cfn-config-configrule-inputparameters
	InputParameters interface{} `json:"InputParameters,omitempty"`
	// MaximumExecutionFrequency docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configrule.html#cfn-config-configrule-maximumexecutionfrequency
	MaximumExecutionFrequency *StringExpr `json:"MaximumExecutionFrequency,omitempty"`
	// Scope docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configrule.html#cfn-config-configrule-scope
	Scope *ConfigConfigRuleScope `json:"Scope,omitempty"`
	// Source docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configrule.html#cfn-config-configrule-source
	Source *ConfigConfigRuleSource `json:"Source,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::Config::ConfigRule to implement the ResourceProperties interface
func (s ConfigConfigRule) CfnResourceType() string {

	return "AWS::Config::ConfigRule"
}

// ConfigConfigurationRecorder represents the AWS::Config::ConfigurationRecorder CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationrecorder.html
type ConfigConfigurationRecorder struct {
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationrecorder.html#cfn-config-configurationrecorder-name
	Name *StringExpr `json:"Name,omitempty"`
	// RecordingGroup docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationrecorder.html#cfn-config-configurationrecorder-recordinggroup
	RecordingGroup *ConfigConfigurationRecorderRecordingGroup `json:"RecordingGroup,omitempty"`
	// RoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationrecorder.html#cfn-config-configurationrecorder-rolearn
	RoleArn *StringExpr `json:"RoleArn,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::Config::ConfigurationRecorder to implement the ResourceProperties interface
func (s ConfigConfigurationRecorder) CfnResourceType() string {

	return "AWS::Config::ConfigurationRecorder"
}

// ConfigDeliveryChannel represents the AWS::Config::DeliveryChannel CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html
type ConfigDeliveryChannel struct {
	// ConfigSnapshotDeliveryProperties docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html#cfn-config-deliverychannel-configsnapshotdeliveryproperties
	ConfigSnapshotDeliveryProperties *ConfigDeliveryChannelConfigSnapshotDeliveryProperties `json:"ConfigSnapshotDeliveryProperties,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html#cfn-config-deliverychannel-name
	Name *StringExpr `json:"Name,omitempty"`
	// S3BucketName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html#cfn-config-deliverychannel-s3bucketname
	S3BucketName *StringExpr `json:"S3BucketName,omitempty" validate:"dive,required"`
	// S3KeyPrefix docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html#cfn-config-deliverychannel-s3keyprefix
	S3KeyPrefix *StringExpr `json:"S3KeyPrefix,omitempty"`
	// SnsTopicARN docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html#cfn-config-deliverychannel-snstopicarn
	SnsTopicARN *StringExpr `json:"SnsTopicARN,omitempty"`
}

// CfnResourceType returns AWS::Config::DeliveryChannel to implement the ResourceProperties interface
func (s ConfigDeliveryChannel) CfnResourceType() string {

	return "AWS::Config::DeliveryChannel"
}

// DataPipelinePipeline represents the AWS::DataPipeline::Pipeline CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html
type DataPipelinePipeline struct {
	// Activate docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html#cfn-datapipeline-pipeline-activate
	Activate *BoolExpr `json:"Activate,omitempty"`
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html#cfn-datapipeline-pipeline-description
	Description *StringExpr `json:"Description,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html#cfn-datapipeline-pipeline-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// ParameterObjects docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html#cfn-datapipeline-pipeline-parameterobjects
	ParameterObjects *DataPipelinePipelineParameterObjectList `json:"ParameterObjects,omitempty" validate:"dive,required"`
	// ParameterValues docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html#cfn-datapipeline-pipeline-parametervalues
	ParameterValues *DataPipelinePipelineParameterValueList `json:"ParameterValues,omitempty"`
	// PipelineObjects docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html#cfn-datapipeline-pipeline-pipelineobjects
	PipelineObjects *DataPipelinePipelinePipelineObjectList `json:"PipelineObjects,omitempty"`
	// PipelineTags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html#cfn-datapipeline-pipeline-pipelinetags
	PipelineTags *DataPipelinePipelinePipelineTagList `json:"PipelineTags,omitempty"`
}

// CfnResourceType returns AWS::DataPipeline::Pipeline to implement the ResourceProperties interface
func (s DataPipelinePipeline) CfnResourceType() string {

	return "AWS::DataPipeline::Pipeline"
}

// DirectoryServiceMicrosoftAD represents the AWS::DirectoryService::MicrosoftAD CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html
type DirectoryServiceMicrosoftAD struct {
	// CreateAlias docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html#cfn-directoryservice-microsoftad-createalias
	CreateAlias *BoolExpr `json:"CreateAlias,omitempty"`
	// EnableSso docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html#cfn-directoryservice-microsoftad-enablesso
	EnableSso *BoolExpr `json:"EnableSso,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html#cfn-directoryservice-microsoftad-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// Password docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html#cfn-directoryservice-microsoftad-password
	Password *StringExpr `json:"Password,omitempty" validate:"dive,required"`
	// ShortName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html#cfn-directoryservice-microsoftad-shortname
	ShortName *StringExpr `json:"ShortName,omitempty"`
	// VPCSettings docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html#cfn-directoryservice-microsoftad-vpcsettings
	VPCSettings *DirectoryServiceMicrosoftADVPCSettings `json:"VpcSettings,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::DirectoryService::MicrosoftAD to implement the ResourceProperties interface
func (s DirectoryServiceMicrosoftAD) CfnResourceType() string {

	return "AWS::DirectoryService::MicrosoftAD"
}

// DirectoryServiceSimpleAD represents the AWS::DirectoryService::SimpleAD CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html
type DirectoryServiceSimpleAD struct {
	// CreateAlias docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html#cfn-directoryservice-simplead-createalias
	CreateAlias *BoolExpr `json:"CreateAlias,omitempty"`
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html#cfn-directoryservice-simplead-description
	Description *StringExpr `json:"Description,omitempty"`
	// EnableSso docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html#cfn-directoryservice-simplead-enablesso
	EnableSso *BoolExpr `json:"EnableSso,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html#cfn-directoryservice-simplead-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// Password docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html#cfn-directoryservice-simplead-password
	Password *StringExpr `json:"Password,omitempty" validate:"dive,required"`
	// ShortName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html#cfn-directoryservice-simplead-shortname
	ShortName *StringExpr `json:"ShortName,omitempty"`
	// Size docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html#cfn-directoryservice-simplead-size
	Size *StringExpr `json:"Size,omitempty" validate:"dive,required"`
	// VPCSettings docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-simplead.html#cfn-directoryservice-simplead-vpcsettings
	VPCSettings *DirectoryServiceSimpleADVPCSettings `json:"VpcSettings,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::DirectoryService::SimpleAD to implement the ResourceProperties interface
func (s DirectoryServiceSimpleAD) CfnResourceType() string {

	return "AWS::DirectoryService::SimpleAD"
}

// DynamoDBTable represents the AWS::DynamoDB::Table CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html
type DynamoDBTable struct {
	// AttributeDefinitions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-attributedef
	AttributeDefinitions *DynamoDBTableAttributeDefinitionList `json:"AttributeDefinitions,omitempty" validate:"dive,required"`
	// GlobalSecondaryIndexes docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-gsi
	GlobalSecondaryIndexes *DynamoDBTableGlobalSecondaryIndexList `json:"GlobalSecondaryIndexes,omitempty"`
	// KeySchema docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-keyschema
	KeySchema *DynamoDBTableKeySchemaList `json:"KeySchema,omitempty" validate:"dive,required"`
	// LocalSecondaryIndexes docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-lsi
	LocalSecondaryIndexes *DynamoDBTableLocalSecondaryIndexList `json:"LocalSecondaryIndexes,omitempty"`
	// ProvisionedThroughput docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-provisionedthroughput
	ProvisionedThroughput *DynamoDBTableProvisionedThroughput `json:"ProvisionedThroughput,omitempty" validate:"dive,required"`
	// StreamSpecification docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-streamspecification
	StreamSpecification *DynamoDBTableStreamSpecification `json:"StreamSpecification,omitempty"`
	// TableName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-tablename
	TableName *StringExpr `json:"TableName,omitempty"`
}

// CfnResourceType returns AWS::DynamoDB::Table to implement the ResourceProperties interface
func (s DynamoDBTable) CfnResourceType() string {

	return "AWS::DynamoDB::Table"
}

// EC2CustomerGateway represents the AWS::EC2::CustomerGateway CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customer-gateway.html
type EC2CustomerGateway struct {
	// BgpAsn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customer-gateway.html#cfn-ec2-customergateway-bgpasn
	BgpAsn *IntegerExpr `json:"BgpAsn,omitempty" validate:"dive,required"`
	// IPAddress docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customer-gateway.html#cfn-ec2-customergateway-ipaddress
	IPAddress *StringExpr `json:"IpAddress,omitempty" validate:"dive,required"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customer-gateway.html#cfn-ec2-customergateway-tags
	Tags *TagList `json:"Tags,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customer-gateway.html#cfn-ec2-customergateway-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::EC2::CustomerGateway to implement the ResourceProperties interface
func (s EC2CustomerGateway) CfnResourceType() string {

	return "AWS::EC2::CustomerGateway"
}

// EC2DHCPOptions represents the AWS::EC2::DHCPOptions CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-dhcp-options.html
type EC2DHCPOptions struct {
	// DomainName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-dhcp-options.html#cfn-ec2-dhcpoptions-domainname
	DomainName *StringExpr `json:"DomainName,omitempty"`
	// DomainNameServers docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-dhcp-options.html#cfn-ec2-dhcpoptions-domainnameservers
	DomainNameServers *StringListExpr `json:"DomainNameServers,omitempty"`
	// NetbiosNameServers docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-dhcp-options.html#cfn-ec2-dhcpoptions-netbiosnameservers
	NetbiosNameServers *StringListExpr `json:"NetbiosNameServers,omitempty"`
	// NetbiosNodeType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-dhcp-options.html#cfn-ec2-dhcpoptions-netbiosnodetype
	NetbiosNodeType *IntegerExpr `json:"NetbiosNodeType,omitempty"`
	// NtpServers docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-dhcp-options.html#cfn-ec2-dhcpoptions-ntpservers
	NtpServers *StringExpr `json:"NtpServers,omitempty"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-dhcp-options.html#cfn-ec2-dhcpoptions-tags
	Tags *TagList `json:"Tags,omitempty"`
}

// CfnResourceType returns AWS::EC2::DHCPOptions to implement the ResourceProperties interface
func (s EC2DHCPOptions) CfnResourceType() string {

	return "AWS::EC2::DHCPOptions"
}

// EC2EIP represents the AWS::EC2::EIP CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html
type EC2EIP struct {
	// Domain docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html#cfn-ec2-eip-domain
	Domain *StringExpr `json:"Domain,omitempty"`
	// InstanceID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html#cfn-ec2-eip-instanceid
	InstanceID *StringExpr `json:"InstanceId,omitempty"`
}

// CfnResourceType returns AWS::EC2::EIP to implement the ResourceProperties interface
func (s EC2EIP) CfnResourceType() string {

	return "AWS::EC2::EIP"
}

// EC2EIPAssociation represents the AWS::EC2::EIPAssociation CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip-association.html
type EC2EIPAssociation struct {
	// AllocationID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip-association.html#cfn-ec2-eipassociation-allocationid
	AllocationID *StringExpr `json:"AllocationId,omitempty"`
	// EIP docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip-association.html#cfn-ec2-eipassociation-eip
	EIP *StringExpr `json:"EIP,omitempty"`
	// InstanceID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip-association.html#cfn-ec2-eipassociation-instanceid
	InstanceID *StringExpr `json:"InstanceId,omitempty"`
	// NetworkInterfaceID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip-association.html#cfn-ec2-eipassociation-networkinterfaceid
	NetworkInterfaceID *StringExpr `json:"NetworkInterfaceId,omitempty"`
	// PrivateIPAddress docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip-association.html#cfn-ec2-eipassociation-PrivateIpAddress
	PrivateIPAddress *StringExpr `json:"PrivateIpAddress,omitempty"`
}

// CfnResourceType returns AWS::EC2::EIPAssociation to implement the ResourceProperties interface
func (s EC2EIPAssociation) CfnResourceType() string {

	return "AWS::EC2::EIPAssociation"
}

// EC2FlowLog represents the AWS::EC2::FlowLog CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html
type EC2FlowLog struct {
	// DeliverLogsPermissionArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-deliverlogspermissionarn
	DeliverLogsPermissionArn *StringExpr `json:"DeliverLogsPermissionArn,omitempty" validate:"dive,required"`
	// LogGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-loggroupname
	LogGroupName *StringExpr `json:"LogGroupName,omitempty" validate:"dive,required"`
	// ResourceID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-resourceid
	ResourceID *StringExpr `json:"ResourceId,omitempty" validate:"dive,required"`
	// ResourceType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-resourcetype
	ResourceType *StringExpr `json:"ResourceType,omitempty" validate:"dive,required"`
	// TrafficType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-traffictype
	TrafficType *StringExpr `json:"TrafficType,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::EC2::FlowLog to implement the ResourceProperties interface
func (s EC2FlowLog) CfnResourceType() string {

	return "AWS::EC2::FlowLog"
}

// EC2Host represents the AWS::EC2::Host CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html
type EC2Host struct {
	// AutoPlacement docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-autoplacement
	AutoPlacement *StringExpr `json:"AutoPlacement,omitempty"`
	// AvailabilityZone docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-availabilityzone
	AvailabilityZone *StringExpr `json:"AvailabilityZone,omitempty" validate:"dive,required"`
	// InstanceType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-instancetype
	InstanceType *StringExpr `json:"InstanceType,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::EC2::Host to implement the ResourceProperties interface
func (s EC2Host) CfnResourceType() string {

	return "AWS::EC2::Host"
}

// EC2Instance represents the AWS::EC2::Instance CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html
type EC2Instance struct {
	// AdditionalInfo docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-additionalinfo
	AdditionalInfo *StringExpr `json:"AdditionalInfo,omitempty"`
	// Affinity docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-affinity
	Affinity *StringExpr `json:"Affinity,omitempty"`
	// AvailabilityZone docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-availabilityzone
	AvailabilityZone *StringExpr `json:"AvailabilityZone,omitempty"`
	// BlockDeviceMappings docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-blockdevicemappings
	BlockDeviceMappings *EC2InstanceBlockDeviceMappingList `json:"BlockDeviceMappings,omitempty"`
	// DisableAPITermination docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-disableapitermination
	DisableAPITermination *BoolExpr `json:"DisableApiTermination,omitempty"`
	// EbsOptimized docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-ebsoptimized
	EbsOptimized *BoolExpr `json:"EbsOptimized,omitempty"`
	// HostID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-hostid
	HostID *StringExpr `json:"HostId,omitempty"`
	// IamInstanceProfile docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-iaminstanceprofile
	IamInstanceProfile *StringExpr `json:"IamInstanceProfile,omitempty"`
	// ImageID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-imageid
	ImageID *StringExpr `json:"ImageId,omitempty" validate:"dive,required"`
	// InstanceInitiatedShutdownBehavior docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-instanceinitiatedshutdownbehavior
	InstanceInitiatedShutdownBehavior *StringExpr `json:"InstanceInitiatedShutdownBehavior,omitempty"`
	// InstanceType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-instancetype
	InstanceType *StringExpr `json:"InstanceType,omitempty"`
	// IPv6AddressCount docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-ipv6addresscount
	IPv6AddressCount *IntegerExpr `json:"Ipv6AddressCount,omitempty"`
	// IPv6Addresses docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-ipv6addresses
	IPv6Addresses *EC2InstanceInstanceIPv6AddressList `json:"Ipv6Addresses,omitempty"`
	// KernelID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-kernelid
	KernelID *StringExpr `json:"KernelId,omitempty"`
	// KeyName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-keyname
	KeyName *StringExpr `json:"KeyName,omitempty"`
	// Monitoring docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-monitoring
	Monitoring *BoolExpr `json:"Monitoring,omitempty"`
	// NetworkInterfaces docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-networkinterfaces
	NetworkInterfaces *EC2InstanceNetworkInterfaceList `json:"NetworkInterfaces,omitempty"`
	// PlacementGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-placementgroupname
	PlacementGroupName *StringExpr `json:"PlacementGroupName,omitempty"`
	// PrivateIPAddress docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-privateipaddress
	PrivateIPAddress *StringExpr `json:"PrivateIpAddress,omitempty"`
	// RamdiskID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-ramdiskid
	RamdiskID *StringExpr `json:"RamdiskId,omitempty"`
	// SecurityGroupIDs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-securitygroupids
	SecurityGroupIDs *StringListExpr `json:"SecurityGroupIds,omitempty"`
	// SecurityGroups docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-securitygroups
	SecurityGroups *StringListExpr `json:"SecurityGroups,omitempty"`
	// SourceDestCheck docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-sourcedestcheck
	SourceDestCheck *BoolExpr `json:"SourceDestCheck,omitempty"`
	// SsmAssociations docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-ssmassociations
	SsmAssociations *EC2InstanceSsmAssociationList `json:"SsmAssociations,omitempty"`
	// SubnetID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-subnetid
	SubnetID *StringExpr `json:"SubnetId,omitempty"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-tags
	Tags *TagList `json:"Tags,omitempty"`
	// Tenancy docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-tenancy
	Tenancy *StringExpr `json:"Tenancy,omitempty"`
	// UserData docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-userdata
	UserData *StringExpr `json:"UserData,omitempty"`
	// Volumes docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-volumes
	Volumes *EC2InstanceVolumeList `json:"Volumes,omitempty"`
}

// CfnResourceType returns AWS::EC2::Instance to implement the ResourceProperties interface
func (s EC2Instance) CfnResourceType() string {

	return "AWS::EC2::Instance"
}

// EC2InternetGateway represents the AWS::EC2::InternetGateway CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-internet-gateway.html
type EC2InternetGateway struct {
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-internet-gateway.html#cfn-ec2-internetgateway-tags
	Tags *TagList `json:"Tags,omitempty"`
}

// CfnResourceType returns AWS::EC2::InternetGateway to implement the ResourceProperties interface
func (s EC2InternetGateway) CfnResourceType() string {

	return "AWS::EC2::InternetGateway"
}

// EC2NatGateway represents the AWS::EC2::NatGateway CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-natgateway.html
type EC2NatGateway struct {
	// AllocationID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-natgateway.html#cfn-ec2-natgateway-allocationid
	AllocationID *StringExpr `json:"AllocationId,omitempty" validate:"dive,required"`
	// SubnetID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-natgateway.html#cfn-ec2-natgateway-subnetid
	SubnetID *StringExpr `json:"SubnetId,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::EC2::NatGateway to implement the ResourceProperties interface
func (s EC2NatGateway) CfnResourceType() string {

	return "AWS::EC2::NatGateway"
}

// EC2NetworkACL represents the AWS::EC2::NetworkAcl CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl.html
type EC2NetworkACL struct {
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl.html#cfn-ec2-networkacl-tags
	Tags *TagList `json:"Tags,omitempty"`
	// VPCID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl.html#cfn-ec2-networkacl-vpcid
	VPCID *StringExpr `json:"VpcId,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::EC2::NetworkAcl to implement the ResourceProperties interface
func (s EC2NetworkACL) CfnResourceType() string {

	return "AWS::EC2::NetworkAcl"
}

// EC2NetworkACLEntry represents the AWS::EC2::NetworkAclEntry CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html
type EC2NetworkACLEntry struct {
	// CidrBlock docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html#cfn-ec2-networkaclentry-cidrblock
	CidrBlock *StringExpr `json:"CidrBlock,omitempty" validate:"dive,required"`
	// Egress docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html#cfn-ec2-networkaclentry-egress
	Egress *BoolExpr `json:"Egress,omitempty"`
	// Icmp docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html#cfn-ec2-networkaclentry-icmp
	Icmp *EC2NetworkACLEntryIcmp `json:"Icmp,omitempty"`
	// IPv6CidrBlock docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html#cfn-ec2-networkaclentry-ipv6cidrblock
	IPv6CidrBlock *StringExpr `json:"Ipv6CidrBlock,omitempty"`
	// NetworkACLID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html#cfn-ec2-networkaclentry-networkaclid
	NetworkACLID *StringExpr `json:"NetworkAclId,omitempty" validate:"dive,required"`
	// PortRange docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html#cfn-ec2-networkaclentry-portrange
	PortRange *EC2NetworkACLEntryPortRange `json:"PortRange,omitempty"`
	// Protocol docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html#cfn-ec2-networkaclentry-protocol
	Protocol *IntegerExpr `json:"Protocol,omitempty" validate:"dive,required"`
	// RuleAction docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html#cfn-ec2-networkaclentry-ruleaction
	RuleAction *StringExpr `json:"RuleAction,omitempty" validate:"dive,required"`
	// RuleNumber docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html#cfn-ec2-networkaclentry-rulenumber
	RuleNumber *IntegerExpr `json:"RuleNumber,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::EC2::NetworkAclEntry to implement the ResourceProperties interface
func (s EC2NetworkACLEntry) CfnResourceType() string {

	return "AWS::EC2::NetworkAclEntry"
}

// EC2NetworkInterface represents the AWS::EC2::NetworkInterface CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html
type EC2NetworkInterface struct {
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-awsec2networkinterface-description
	Description *StringExpr `json:"Description,omitempty"`
	// GroupSet docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-awsec2networkinterface-groupset
	GroupSet *StringListExpr `json:"GroupSet,omitempty"`
	// IPv6AddressCount docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-ec2-networkinterface-ipv6addresscount
	IPv6AddressCount *IntegerExpr `json:"Ipv6AddressCount,omitempty"`
	// IPv6Addresses docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-ec2-networkinterface-ipv6addresses
	IPv6Addresses *EC2NetworkInterfaceInstanceIPv6Address `json:"Ipv6Addresses,omitempty"`
	// PrivateIPAddress docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-awsec2networkinterface-privateipaddress
	PrivateIPAddress *StringExpr `json:"PrivateIpAddress,omitempty"`
	// PrivateIPAddresses docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-awsec2networkinterface-privateipaddresses
	PrivateIPAddresses *EC2NetworkInterfacePrivateIPAddressSpecificationList `json:"PrivateIpAddresses,omitempty"`
	// SecondaryPrivateIPAddressCount docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-awsec2networkinterface-secondaryprivateipcount
	SecondaryPrivateIPAddressCount *IntegerExpr `json:"SecondaryPrivateIpAddressCount,omitempty"`
	// SourceDestCheck docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-awsec2networkinterface-sourcedestcheck
	SourceDestCheck *BoolExpr `json:"SourceDestCheck,omitempty"`
	// SubnetID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-awsec2networkinterface-subnetid
	SubnetID *StringExpr `json:"SubnetId,omitempty" validate:"dive,required"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-awsec2networkinterface-tags
	Tags *TagList `json:"Tags,omitempty"`
}

// CfnResourceType returns AWS::EC2::NetworkInterface to implement the ResourceProperties interface
func (s EC2NetworkInterface) CfnResourceType() string {

	return "AWS::EC2::NetworkInterface"
}

// EC2NetworkInterfaceAttachment represents the AWS::EC2::NetworkInterfaceAttachment CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface-attachment.html
type EC2NetworkInterfaceAttachment struct {
	// DeleteOnTermination docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface-attachment.html#cfn-ec2-network-interface-attachment-deleteonterm
	DeleteOnTermination *BoolExpr `json:"DeleteOnTermination,omitempty"`
	// DeviceIndex docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface-attachment.html#cfn-ec2-network-interface-attachment-deviceindex
	DeviceIndex *StringExpr `json:"DeviceIndex,omitempty" validate:"dive,required"`
	// InstanceID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface-attachment.html#cfn-ec2-network-interface-attachment-instanceid
	InstanceID *StringExpr `json:"InstanceId,omitempty" validate:"dive,required"`
	// NetworkInterfaceID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface-attachment.html#cfn-ec2-network-interface-attachment-networkinterfaceid
	NetworkInterfaceID *StringExpr `json:"NetworkInterfaceId,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::EC2::NetworkInterfaceAttachment to implement the ResourceProperties interface
func (s EC2NetworkInterfaceAttachment) CfnResourceType() string {

	return "AWS::EC2::NetworkInterfaceAttachment"
}

// EC2PlacementGroup represents the AWS::EC2::PlacementGroup CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-placementgroup.html
type EC2PlacementGroup struct {
	// Strategy docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-placementgroup.html#cfn-ec2-placementgroup-strategy
	Strategy *StringExpr `json:"Strategy,omitempty"`
}

// CfnResourceType returns AWS::EC2::PlacementGroup to implement the ResourceProperties interface
func (s EC2PlacementGroup) CfnResourceType() string {

	return "AWS::EC2::PlacementGroup"
}

// EC2Route represents the AWS::EC2::Route CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html
type EC2Route struct {
	// DestinationCidrBlock docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-destinationcidrblock
	DestinationCidrBlock *StringExpr `json:"DestinationCidrBlock,omitempty" validate:"dive,required"`
	// DestinationIPv6CidrBlock docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-destinationipv6cidrblock
	DestinationIPv6CidrBlock *StringExpr `json:"DestinationIpv6CidrBlock,omitempty"`
	// GatewayID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-gatewayid
	GatewayID *StringExpr `json:"GatewayId,omitempty"`
	// InstanceID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-instanceid
	InstanceID *StringExpr `json:"InstanceId,omitempty"`
	// NatGatewayID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-natgatewayid
	NatGatewayID *StringExpr `json:"NatGatewayId,omitempty"`
	// NetworkInterfaceID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-networkinterfaceid
	NetworkInterfaceID *StringExpr `json:"NetworkInterfaceId,omitempty"`
	// RouteTableID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-routetableid
	RouteTableID *StringExpr `json:"RouteTableId,omitempty" validate:"dive,required"`
	// VPCPeeringConnectionID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-vpcpeeringconnectionid
	VPCPeeringConnectionID *StringExpr `json:"VpcPeeringConnectionId,omitempty"`
}

// CfnResourceType returns AWS::EC2::Route to implement the ResourceProperties interface
func (s EC2Route) CfnResourceType() string {

	return "AWS::EC2::Route"
}

// EC2RouteTable represents the AWS::EC2::RouteTable CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route-table.html
type EC2RouteTable struct {
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route-table.html#cfn-ec2-routetable-tags
	Tags *TagList `json:"Tags,omitempty"`
	// VPCID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route-table.html#cfn-ec2-routetable-vpcid
	VPCID *StringExpr `json:"VpcId,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::EC2::RouteTable to implement the ResourceProperties interface
func (s EC2RouteTable) CfnResourceType() string {

	return "AWS::EC2::RouteTable"
}

// EC2SecurityGroup represents the AWS::EC2::SecurityGroup CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html
type EC2SecurityGroup struct {
	// GroupDescription docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html#cfn-ec2-securitygroup-groupdescription
	GroupDescription *StringExpr `json:"GroupDescription,omitempty" validate:"dive,required"`
	// GroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html#cfn-ec2-securitygroup-groupname
	GroupName *StringExpr `json:"GroupName,omitempty"`
	// SecurityGroupEgress docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html#cfn-ec2-securitygroup-securitygroupegress
	SecurityGroupEgress *EC2SecurityGroupRuleList `json:"SecurityGroupEgress,omitempty"`
	// SecurityGroupIngress docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html#cfn-ec2-securitygroup-securitygroupingress
	SecurityGroupIngress *EC2SecurityGroupRuleList `json:"SecurityGroupIngress,omitempty"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html#cfn-ec2-securitygroup-tags
	Tags *TagList `json:"Tags,omitempty"`
	// VPCID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html#cfn-ec2-securitygroup-vpcid
	VPCID *StringExpr `json:"VpcId,omitempty"`
}

// CfnResourceType returns AWS::EC2::SecurityGroup to implement the ResourceProperties interface
func (s EC2SecurityGroup) CfnResourceType() string {

	return "AWS::EC2::SecurityGroup"
}

// EC2SecurityGroupEgress represents the AWS::EC2::SecurityGroupEgress CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html
type EC2SecurityGroupEgress struct {
	// CidrIP docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html#cfn-ec2-securitygroupegress-cidrip
	CidrIP *StringExpr `json:"CidrIp,omitempty"`
	// CidrIPv6 docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html#cfn-ec2-securitygroupegress-cidripv6
	CidrIPv6 *StringExpr `json:"CidrIpv6,omitempty"`
	// DestinationPrefixListID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html#cfn-ec2-securitygroupegress-destinationprefixlistid
	DestinationPrefixListID *StringExpr `json:"DestinationPrefixListId,omitempty"`
	// DestinationSecurityGroupID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html#cfn-ec2-securitygroupegress-destinationsecuritygroupid
	DestinationSecurityGroupID *StringExpr `json:"DestinationSecurityGroupId,omitempty"`
	// FromPort docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html#cfn-ec2-securitygroupegress-fromport
	FromPort *IntegerExpr `json:"FromPort,omitempty"`
	// GroupID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html#cfn-ec2-securitygroupegress-groupid
	GroupID *StringExpr `json:"GroupId,omitempty" validate:"dive,required"`
	// IPProtocol docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html#cfn-ec2-securitygroupegress-ipprotocol
	IPProtocol *StringExpr `json:"IpProtocol,omitempty" validate:"dive,required"`
	// ToPort docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html#cfn-ec2-securitygroupegress-toport
	ToPort *IntegerExpr `json:"ToPort,omitempty"`
}

// CfnResourceType returns AWS::EC2::SecurityGroupEgress to implement the ResourceProperties interface
func (s EC2SecurityGroupEgress) CfnResourceType() string {

	return "AWS::EC2::SecurityGroupEgress"
}

// EC2SecurityGroupIngress represents the AWS::EC2::SecurityGroupIngress CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html
type EC2SecurityGroupIngress struct {
	// CidrIP docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-cidrip
	CidrIP *StringExpr `json:"CidrIp,omitempty"`
	// CidrIPv6 docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-cidripv6
	CidrIPv6 *StringExpr `json:"CidrIpv6,omitempty"`
	// FromPort docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-fromport
	FromPort *IntegerExpr `json:"FromPort,omitempty"`
	// GroupID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-groupid
	GroupID *StringExpr `json:"GroupId,omitempty"`
	// GroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-groupname
	GroupName *StringExpr `json:"GroupName,omitempty"`
	// IPProtocol docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-ipprotocol
	IPProtocol *StringExpr `json:"IpProtocol,omitempty" validate:"dive,required"`
	// SourceSecurityGroupID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-sourcesecuritygroupid
	SourceSecurityGroupID *StringExpr `json:"SourceSecurityGroupId,omitempty"`
	// SourceSecurityGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-sourcesecuritygroupname
	SourceSecurityGroupName *StringExpr `json:"SourceSecurityGroupName,omitempty"`
	// SourceSecurityGroupOwnerID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-sourcesecuritygroupownerid
	SourceSecurityGroupOwnerID *StringExpr `json:"SourceSecurityGroupOwnerId,omitempty"`
	// ToPort docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-toport
	ToPort *IntegerExpr `json:"ToPort,omitempty"`
}

// CfnResourceType returns AWS::EC2::SecurityGroupIngress to implement the ResourceProperties interface
func (s EC2SecurityGroupIngress) CfnResourceType() string {

	return "AWS::EC2::SecurityGroupIngress"
}

// EC2SpotFleet represents the AWS::EC2::SpotFleet CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-spotfleet.html
type EC2SpotFleet struct {
	// SpotFleetRequestConfigData docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-spotfleet.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata
	SpotFleetRequestConfigData *EC2SpotFleetSpotFleetRequestConfigData `json:"SpotFleetRequestConfigData,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::EC2::SpotFleet to implement the ResourceProperties interface
func (s EC2SpotFleet) CfnResourceType() string {

	return "AWS::EC2::SpotFleet"
}

// EC2Subnet represents the AWS::EC2::Subnet CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html
type EC2Subnet struct {
	// AvailabilityZone docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-availabilityzone
	AvailabilityZone *StringExpr `json:"AvailabilityZone,omitempty"`
	// CidrBlock docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-cidrblock
	CidrBlock *StringExpr `json:"CidrBlock,omitempty" validate:"dive,required"`
	// MapPublicIPOnLaunch docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-mappubliciponlaunch
	MapPublicIPOnLaunch *BoolExpr `json:"MapPublicIpOnLaunch,omitempty"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-tags
	Tags *TagList `json:"Tags,omitempty"`
	// VPCID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-awsec2subnet-prop-vpcid
	VPCID *StringExpr `json:"VpcId,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::EC2::Subnet to implement the ResourceProperties interface
func (s EC2Subnet) CfnResourceType() string {

	return "AWS::EC2::Subnet"
}

// EC2SubnetCidrBlock represents the AWS::EC2::SubnetCidrBlock CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetcidrblock.html
type EC2SubnetCidrBlock struct {
	// IPv6CidrBlock docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetcidrblock.html#cfn-ec2-subnetcidrblock-ipv6cidrblock
	IPv6CidrBlock *StringExpr `json:"Ipv6CidrBlock,omitempty" validate:"dive,required"`
	// SubnetID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetcidrblock.html#cfn-ec2-subnetcidrblock-subnetid
	SubnetID *StringExpr `json:"SubnetId,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::EC2::SubnetCidrBlock to implement the ResourceProperties interface
func (s EC2SubnetCidrBlock) CfnResourceType() string {

	return "AWS::EC2::SubnetCidrBlock"
}

// EC2SubnetNetworkACLAssociation represents the AWS::EC2::SubnetNetworkAclAssociation CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-network-acl-assoc.html
type EC2SubnetNetworkACLAssociation struct {
	// NetworkACLID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-network-acl-assoc.html#cfn-ec2-subnetnetworkaclassociation-networkaclid
	NetworkACLID *StringExpr `json:"NetworkAclId,omitempty" validate:"dive,required"`
	// SubnetID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-network-acl-assoc.html#cfn-ec2-subnetnetworkaclassociation-associationid
	SubnetID *StringExpr `json:"SubnetId,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::EC2::SubnetNetworkAclAssociation to implement the ResourceProperties interface
func (s EC2SubnetNetworkACLAssociation) CfnResourceType() string {

	return "AWS::EC2::SubnetNetworkAclAssociation"
}

// EC2SubnetRouteTableAssociation represents the AWS::EC2::SubnetRouteTableAssociation CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-route-table-assoc.html
type EC2SubnetRouteTableAssociation struct {
	// RouteTableID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-route-table-assoc.html#cfn-ec2-subnetroutetableassociation-routetableid
	RouteTableID *StringExpr `json:"RouteTableId,omitempty" validate:"dive,required"`
	// SubnetID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-route-table-assoc.html#cfn-ec2-subnetroutetableassociation-subnetid
	SubnetID *StringExpr `json:"SubnetId,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::EC2::SubnetRouteTableAssociation to implement the ResourceProperties interface
func (s EC2SubnetRouteTableAssociation) CfnResourceType() string {

	return "AWS::EC2::SubnetRouteTableAssociation"
}

// EC2VPC represents the AWS::EC2::VPC CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc.html
type EC2VPC struct {
	// CidrBlock docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc.html#cfn-aws-ec2-vpc-cidrblock
	CidrBlock *StringExpr `json:"CidrBlock,omitempty" validate:"dive,required"`
	// EnableDNSHostnames docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc.html#cfn-aws-ec2-vpc-EnableDnsHostnames
	EnableDNSHostnames *BoolExpr `json:"EnableDnsHostnames,omitempty"`
	// EnableDNSSupport docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc.html#cfn-aws-ec2-vpc-EnableDnsSupport
	EnableDNSSupport *BoolExpr `json:"EnableDnsSupport,omitempty"`
	// InstanceTenancy docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc.html#cfn-aws-ec2-vpc-instancetenancy
	InstanceTenancy *StringExpr `json:"InstanceTenancy,omitempty"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc.html#cfn-aws-ec2-vpc-tags
	Tags *TagList `json:"Tags,omitempty"`
}

// CfnResourceType returns AWS::EC2::VPC to implement the ResourceProperties interface
func (s EC2VPC) CfnResourceType() string {

	return "AWS::EC2::VPC"
}

// EC2VPCCidrBlock represents the AWS::EC2::VPCCidrBlock CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpccidrblock.html
type EC2VPCCidrBlock struct {
	// AmazonProvidedIPv6CidrBlock docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpccidrblock.html#cfn-ec2-vpccidrblock-amazonprovidedipv6cidrblock
	AmazonProvidedIPv6CidrBlock *BoolExpr `json:"AmazonProvidedIpv6CidrBlock,omitempty"`
	// VPCID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpccidrblock.html#cfn-ec2-vpccidrblock-vpcid
	VPCID *StringExpr `json:"VpcId,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::EC2::VPCCidrBlock to implement the ResourceProperties interface
func (s EC2VPCCidrBlock) CfnResourceType() string {

	return "AWS::EC2::VPCCidrBlock"
}

// EC2VPCDHCPOptionsAssociation represents the AWS::EC2::VPCDHCPOptionsAssociation CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-dhcp-options-assoc.html
type EC2VPCDHCPOptionsAssociation struct {
	// DhcpOptionsID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-dhcp-options-assoc.html#cfn-ec2-vpcdhcpoptionsassociation-dhcpoptionsid
	DhcpOptionsID *StringExpr `json:"DhcpOptionsId,omitempty" validate:"dive,required"`
	// VPCID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-dhcp-options-assoc.html#cfn-ec2-vpcdhcpoptionsassociation-vpcid
	VPCID *StringExpr `json:"VpcId,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::EC2::VPCDHCPOptionsAssociation to implement the ResourceProperties interface
func (s EC2VPCDHCPOptionsAssociation) CfnResourceType() string {

	return "AWS::EC2::VPCDHCPOptionsAssociation"
}

// EC2VPCEndpoint represents the AWS::EC2::VPCEndpoint CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html
type EC2VPCEndpoint struct {
	// PolicyDocument docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-policydocument
	PolicyDocument interface{} `json:"PolicyDocument,omitempty"`
	// RouteTableIDs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-routetableids
	RouteTableIDs *StringListExpr `json:"RouteTableIds,omitempty"`
	// ServiceName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-servicename
	ServiceName *StringExpr `json:"ServiceName,omitempty" validate:"dive,required"`
	// VPCID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-vpcid
	VPCID *StringExpr `json:"VpcId,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::EC2::VPCEndpoint to implement the ResourceProperties interface
func (s EC2VPCEndpoint) CfnResourceType() string {

	return "AWS::EC2::VPCEndpoint"
}

// EC2VPCGatewayAttachment represents the AWS::EC2::VPCGatewayAttachment CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-gateway-attachment.html
type EC2VPCGatewayAttachment struct {
	// InternetGatewayID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-gateway-attachment.html#cfn-ec2-vpcgatewayattachment-internetgatewayid
	InternetGatewayID *StringExpr `json:"InternetGatewayId,omitempty"`
	// VPCID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-gateway-attachment.html#cfn-ec2-vpcgatewayattachment-vpcid
	VPCID *StringExpr `json:"VpcId,omitempty" validate:"dive,required"`
	// VpnGatewayID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-gateway-attachment.html#cfn-ec2-vpcgatewayattachment-vpngatewayid
	VpnGatewayID *StringExpr `json:"VpnGatewayId,omitempty"`
}

// CfnResourceType returns AWS::EC2::VPCGatewayAttachment to implement the ResourceProperties interface
func (s EC2VPCGatewayAttachment) CfnResourceType() string {

	return "AWS::EC2::VPCGatewayAttachment"
}

// EC2VPCPeeringConnection represents the AWS::EC2::VPCPeeringConnection CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html
type EC2VPCPeeringConnection struct {
	// PeerOwnerID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html#cfn-ec2-vpcpeeringconnection-peerownerid
	PeerOwnerID *StringExpr `json:"PeerOwnerId,omitempty"`
	// PeerRoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html#cfn-ec2-vpcpeeringconnection-peerrolearn
	PeerRoleArn *StringExpr `json:"PeerRoleArn,omitempty"`
	// PeerVPCID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html#cfn-ec2-vpcpeeringconnection-peervpcid
	PeerVPCID *StringExpr `json:"PeerVpcId,omitempty" validate:"dive,required"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html#cfn-ec2-vpcpeeringconnection-tags
	Tags *TagList `json:"Tags,omitempty"`
	// VPCID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html#cfn-ec2-vpcpeeringconnection-vpcid
	VPCID *StringExpr `json:"VpcId,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::EC2::VPCPeeringConnection to implement the ResourceProperties interface
func (s EC2VPCPeeringConnection) CfnResourceType() string {

	return "AWS::EC2::VPCPeeringConnection"
}

// EC2VPNConnection represents the AWS::EC2::VPNConnection CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection.html
type EC2VPNConnection struct {
	// CustomerGatewayID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection.html#cfn-ec2-vpnconnection-customergatewayid
	CustomerGatewayID *StringExpr `json:"CustomerGatewayId,omitempty" validate:"dive,required"`
	// StaticRoutesOnly docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection.html#cfn-ec2-vpnconnection-StaticRoutesOnly
	StaticRoutesOnly *BoolExpr `json:"StaticRoutesOnly,omitempty"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection.html#cfn-ec2-vpnconnection-tags
	Tags *TagList `json:"Tags,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection.html#cfn-ec2-vpnconnection-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
	// VpnGatewayID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection.html#cfn-ec2-vpnconnection-vpngatewayid
	VpnGatewayID *StringExpr `json:"VpnGatewayId,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::EC2::VPNConnection to implement the ResourceProperties interface
func (s EC2VPNConnection) CfnResourceType() string {

	return "AWS::EC2::VPNConnection"
}

// EC2VPNConnectionRoute represents the AWS::EC2::VPNConnectionRoute CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection-route.html
type EC2VPNConnectionRoute struct {
	// DestinationCidrBlock docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection-route.html#cfn-ec2-vpnconnectionroute-cidrblock
	DestinationCidrBlock *StringExpr `json:"DestinationCidrBlock,omitempty" validate:"dive,required"`
	// VpnConnectionID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection-route.html#cfn-ec2-vpnconnectionroute-connectionid
	VpnConnectionID *StringExpr `json:"VpnConnectionId,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::EC2::VPNConnectionRoute to implement the ResourceProperties interface
func (s EC2VPNConnectionRoute) CfnResourceType() string {

	return "AWS::EC2::VPNConnectionRoute"
}

// EC2VPNGateway represents the AWS::EC2::VPNGateway CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-gateway.html
type EC2VPNGateway struct {
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-gateway.html#cfn-ec2-vpngateway-tags
	Tags *TagList `json:"Tags,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-gateway.html#cfn-ec2-vpngateway-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::EC2::VPNGateway to implement the ResourceProperties interface
func (s EC2VPNGateway) CfnResourceType() string {

	return "AWS::EC2::VPNGateway"
}

// EC2VPNGatewayRoutePropagation represents the AWS::EC2::VPNGatewayRoutePropagation CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-gatewayrouteprop.html
type EC2VPNGatewayRoutePropagation struct {
	// RouteTableIDs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-gatewayrouteprop.html#cfn-ec2-vpngatewayrouteprop-routetableids
	RouteTableIDs *StringListExpr `json:"RouteTableIds,omitempty" validate:"dive,required"`
	// VpnGatewayID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-gatewayrouteprop.html#cfn-ec2-vpngatewayrouteprop-vpngatewayid
	VpnGatewayID *StringExpr `json:"VpnGatewayId,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::EC2::VPNGatewayRoutePropagation to implement the ResourceProperties interface
func (s EC2VPNGatewayRoutePropagation) CfnResourceType() string {

	return "AWS::EC2::VPNGatewayRoutePropagation"
}

// EC2Volume represents the AWS::EC2::Volume CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volume.html
type EC2Volume struct {
	// AutoEnableIO docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volume.html#cfn-ec2-ebs-volume-autoenableio
	AutoEnableIO *BoolExpr `json:"AutoEnableIO,omitempty"`
	// AvailabilityZone docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volume.html#cfn-ec2-ebs-volume-availabilityzone
	AvailabilityZone *StringExpr `json:"AvailabilityZone,omitempty" validate:"dive,required"`
	// Encrypted docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volume.html#cfn-ec2-ebs-volume-encrypted
	Encrypted *BoolExpr `json:"Encrypted,omitempty"`
	// Iops docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volume.html#cfn-ec2-ebs-volume-iops
	Iops *IntegerExpr `json:"Iops,omitempty"`
	// KmsKeyID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volume.html#cfn-ec2-ebs-volume-kmskeyid
	KmsKeyID *StringExpr `json:"KmsKeyId,omitempty"`
	// Size docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volume.html#cfn-ec2-ebs-volume-size
	Size *IntegerExpr `json:"Size,omitempty"`
	// SnapshotID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volume.html#cfn-ec2-ebs-volume-snapshotid
	SnapshotID *StringExpr `json:"SnapshotId,omitempty"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volume.html#cfn-ec2-ebs-volume-tags
	Tags *TagList `json:"Tags,omitempty"`
	// VolumeType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volume.html#cfn-ec2-ebs-volume-volumetype
	VolumeType *StringExpr `json:"VolumeType,omitempty"`
}

// CfnResourceType returns AWS::EC2::Volume to implement the ResourceProperties interface
func (s EC2Volume) CfnResourceType() string {

	return "AWS::EC2::Volume"
}

// EC2VolumeAttachment represents the AWS::EC2::VolumeAttachment CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volumeattachment.html
type EC2VolumeAttachment struct {
	// Device docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volumeattachment.html#cfn-ec2-ebs-volumeattachment-device
	Device *StringExpr `json:"Device,omitempty" validate:"dive,required"`
	// InstanceID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volumeattachment.html#cfn-ec2-ebs-volumeattachment-instanceid
	InstanceID *StringExpr `json:"InstanceId,omitempty" validate:"dive,required"`
	// VolumeID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volumeattachment.html#cfn-ec2-ebs-volumeattachment-volumeid
	VolumeID *StringExpr `json:"VolumeId,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::EC2::VolumeAttachment to implement the ResourceProperties interface
func (s EC2VolumeAttachment) CfnResourceType() string {

	return "AWS::EC2::VolumeAttachment"
}

// ECRRepository represents the AWS::ECR::Repository CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html
type ECRRepository struct {
	// RepositoryName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html#cfn-ecr-repository-repositoryname
	RepositoryName *StringExpr `json:"RepositoryName,omitempty"`
	// RepositoryPolicyText docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html#cfn-ecr-repository-repositorypolicytext
	RepositoryPolicyText interface{} `json:"RepositoryPolicyText,omitempty"`
}

// CfnResourceType returns AWS::ECR::Repository to implement the ResourceProperties interface
func (s ECRRepository) CfnResourceType() string {

	return "AWS::ECR::Repository"
}

// ECSCluster represents the AWS::ECS::Cluster CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-cluster.html
type ECSCluster struct {
	// ClusterName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-cluster.html#cfn-ecs-cluster-clustername
	ClusterName *StringExpr `json:"ClusterName,omitempty"`
}

// CfnResourceType returns AWS::ECS::Cluster to implement the ResourceProperties interface
func (s ECSCluster) CfnResourceType() string {

	return "AWS::ECS::Cluster"
}

// ECSService represents the AWS::ECS::Service CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html
type ECSService struct {
	// Cluster docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-cluster
	Cluster *StringExpr `json:"Cluster,omitempty"`
	// DeploymentConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-deploymentconfiguration
	DeploymentConfiguration *ECSServiceDeploymentConfiguration `json:"DeploymentConfiguration,omitempty"`
	// DesiredCount docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-desiredcount
	DesiredCount *IntegerExpr `json:"DesiredCount,omitempty"`
	// LoadBalancers docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-loadbalancers
	LoadBalancers *ECSServiceLoadBalancerList `json:"LoadBalancers,omitempty"`
	// PlacementConstraints docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-placementconstraints
	PlacementConstraints *ECSServicePlacementConstraintList `json:"PlacementConstraints,omitempty"`
	// PlacementStrategies docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-placementstrategies
	PlacementStrategies *ECSServicePlacementStrategyList `json:"PlacementStrategies,omitempty"`
	// Role docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-role
	Role *StringExpr `json:"Role,omitempty"`
	// ServiceName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-servicename
	ServiceName *StringExpr `json:"ServiceName,omitempty"`
	// TaskDefinition docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-taskdefinition
	TaskDefinition *StringExpr `json:"TaskDefinition,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::ECS::Service to implement the ResourceProperties interface
func (s ECSService) CfnResourceType() string {

	return "AWS::ECS::Service"
}

// ECSTaskDefinition represents the AWS::ECS::TaskDefinition CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html
type ECSTaskDefinition struct {
	// ContainerDefinitions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-containerdefinitions
	ContainerDefinitions *ECSTaskDefinitionContainerDefinitionList `json:"ContainerDefinitions,omitempty"`
	// Family docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-family
	Family *StringExpr `json:"Family,omitempty"`
	// NetworkMode docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-networkmode
	NetworkMode *StringExpr `json:"NetworkMode,omitempty"`
	// PlacementConstraints docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-placementconstraints
	PlacementConstraints *ECSTaskDefinitionTaskDefinitionPlacementConstraintList `json:"PlacementConstraints,omitempty"`
	// TaskRoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-taskrolearn
	TaskRoleArn *StringExpr `json:"TaskRoleArn,omitempty"`
	// Volumes docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-volumes
	Volumes *ECSTaskDefinitionVolumeList `json:"Volumes,omitempty"`
}

// CfnResourceType returns AWS::ECS::TaskDefinition to implement the ResourceProperties interface
func (s ECSTaskDefinition) CfnResourceType() string {

	return "AWS::ECS::TaskDefinition"
}

// EFSFileSystem represents the AWS::EFS::FileSystem CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html
type EFSFileSystem struct {
	// FileSystemTags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-filesystemtags
	FileSystemTags *EFSFileSystemElasticFileSystemTagList `json:"FileSystemTags,omitempty"`
	// PerformanceMode docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-performancemode
	PerformanceMode *StringExpr `json:"PerformanceMode,omitempty"`
}

// CfnResourceType returns AWS::EFS::FileSystem to implement the ResourceProperties interface
func (s EFSFileSystem) CfnResourceType() string {

	return "AWS::EFS::FileSystem"
}

// EFSMountTarget represents the AWS::EFS::MountTarget CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html
type EFSMountTarget struct {
	// FileSystemID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html#cfn-efs-mounttarget-filesystemid
	FileSystemID *StringExpr `json:"FileSystemId,omitempty" validate:"dive,required"`
	// IPAddress docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html#cfn-efs-mounttarget-ipaddress
	IPAddress *StringExpr `json:"IpAddress,omitempty"`
	// SecurityGroups docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html#cfn-efs-mounttarget-securitygroups
	SecurityGroups *StringListExpr `json:"SecurityGroups,omitempty" validate:"dive,required"`
	// SubnetID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html#cfn-efs-mounttarget-subnetid
	SubnetID *StringExpr `json:"SubnetId,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::EFS::MountTarget to implement the ResourceProperties interface
func (s EFSMountTarget) CfnResourceType() string {

	return "AWS::EFS::MountTarget"
}

// EMRCluster represents the AWS::EMR::Cluster CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html
type EMRCluster struct {
	// AdditionalInfo docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-additionalinfo
	AdditionalInfo interface{} `json:"AdditionalInfo,omitempty"`
	// Applications docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-applications
	Applications *EMRClusterApplicationList `json:"Applications,omitempty"`
	// AutoScalingRole docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-elasticmapreduce-cluster-autoscalingrole
	AutoScalingRole *StringExpr `json:"AutoScalingRole,omitempty"`
	// BootstrapActions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-bootstrapactions
	BootstrapActions *EMRClusterBootstrapActionConfigList `json:"BootstrapActions,omitempty"`
	// Configurations docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-configurations
	Configurations *EMRClusterConfigurationList `json:"Configurations,omitempty"`
	// Instances docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-instances
	Instances *EMRClusterJobFlowInstancesConfig `json:"Instances,omitempty" validate:"dive,required"`
	// JobFlowRole docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-jobflowrole
	JobFlowRole *StringExpr `json:"JobFlowRole,omitempty" validate:"dive,required"`
	// LogURI docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-loguri
	LogURI *StringExpr `json:"LogUri,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// ReleaseLabel docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-releaselabel
	ReleaseLabel *StringExpr `json:"ReleaseLabel,omitempty"`
	// ScaleDownBehavior docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-elasticmapreduce-cluster-scaledownbehavior
	ScaleDownBehavior *StringExpr `json:"ScaleDownBehavior,omitempty"`
	// SecurityConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-securityconfiguration
	SecurityConfiguration *StringExpr `json:"SecurityConfiguration,omitempty"`
	// ServiceRole docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-servicerole
	ServiceRole *StringExpr `json:"ServiceRole,omitempty" validate:"dive,required"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-elasticmapreduce-cluster-tags
	Tags *TagList `json:"Tags,omitempty"`
	// VisibleToAllUsers docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-visibletoallusers
	VisibleToAllUsers *BoolExpr `json:"VisibleToAllUsers,omitempty"`
}

// CfnResourceType returns AWS::EMR::Cluster to implement the ResourceProperties interface
func (s EMRCluster) CfnResourceType() string {

	return "AWS::EMR::Cluster"
}

// EMRInstanceGroupConfig represents the AWS::EMR::InstanceGroupConfig CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancegroupconfig.html
type EMRInstanceGroupConfig struct {
	// AutoScalingPolicy docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancegroupconfig.html#cfn-elasticmapreduce-instancegroupconfig-autoscalingpolicy
	AutoScalingPolicy *EMRInstanceGroupConfigAutoScalingPolicy `json:"AutoScalingPolicy,omitempty"`
	// BidPrice docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancegroupconfig.html#cfn-emr-instancegroupconfig-bidprice
	BidPrice *StringExpr `json:"BidPrice,omitempty"`
	// Configurations docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancegroupconfig.html#cfn-emr-instancegroupconfig-configurations
	Configurations *EMRInstanceGroupConfigConfigurationList `json:"Configurations,omitempty"`
	// EbsConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancegroupconfig.html#cfn-emr-instancegroupconfig-ebsconfiguration
	EbsConfiguration *EMRInstanceGroupConfigEbsConfiguration `json:"EbsConfiguration,omitempty"`
	// InstanceCount docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancegroupconfig.html#cfn-emr-instancegroupconfiginstancecount-
	InstanceCount *IntegerExpr `json:"InstanceCount,omitempty" validate:"dive,required"`
	// InstanceRole docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancegroupconfig.html#cfn-emr-instancegroupconfig-instancerole
	InstanceRole *StringExpr `json:"InstanceRole,omitempty" validate:"dive,required"`
	// InstanceType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancegroupconfig.html#cfn-emr-instancegroupconfig-instancetype
	InstanceType *StringExpr `json:"InstanceType,omitempty" validate:"dive,required"`
	// JobFlowID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancegroupconfig.html#cfn-emr-instancegroupconfig-jobflowid
	JobFlowID *StringExpr `json:"JobFlowId,omitempty" validate:"dive,required"`
	// Market docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancegroupconfig.html#cfn-emr-instancegroupconfig-market
	Market *StringExpr `json:"Market,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancegroupconfig.html#cfn-emr-instancegroupconfig-name
	Name *StringExpr `json:"Name,omitempty"`
}

// CfnResourceType returns AWS::EMR::InstanceGroupConfig to implement the ResourceProperties interface
func (s EMRInstanceGroupConfig) CfnResourceType() string {

	return "AWS::EMR::InstanceGroupConfig"
}

// EMRSecurityConfiguration represents the AWS::EMR::SecurityConfiguration CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-securityconfiguration.html
type EMRSecurityConfiguration struct {
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-securityconfiguration.html#cfn-emr-securityconfiguration-name
	Name *StringExpr `json:"Name,omitempty"`
	// SecurityConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-securityconfiguration.html#cfn-emr-securityconfiguration-securityconfiguration
	SecurityConfiguration interface{} `json:"SecurityConfiguration,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::EMR::SecurityConfiguration to implement the ResourceProperties interface
func (s EMRSecurityConfiguration) CfnResourceType() string {

	return "AWS::EMR::SecurityConfiguration"
}

// EMRStep represents the AWS::EMR::Step CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-step.html
type EMRStep struct {
	// ActionOnFailure docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-step.html#cfn-elasticmapreduce-step-actiononfailure
	ActionOnFailure *StringExpr `json:"ActionOnFailure,omitempty" validate:"dive,required"`
	// HadoopJarStep docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-step.html#cfn-elasticmapreduce-step-hadoopjarstep
	HadoopJarStep *EMRStepHadoopJarStepConfig `json:"HadoopJarStep,omitempty" validate:"dive,required"`
	// JobFlowID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-step.html#cfn-elasticmapreduce-step-jobflowid
	JobFlowID *StringExpr `json:"JobFlowId,omitempty" validate:"dive,required"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-step.html#cfn-elasticmapreduce-step-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::EMR::Step to implement the ResourceProperties interface
func (s EMRStep) CfnResourceType() string {

	return "AWS::EMR::Step"
}

// ElastiCacheCacheCluster represents the AWS::ElastiCache::CacheCluster CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cache-cluster.html
type ElastiCacheCacheCluster struct {
	// AZMode docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cache-cluster.html#cfn-elasticache-cachecluster-azmode
	AZMode *StringExpr `json:"AZMode,omitempty"`
	// AutoMinorVersionUpgrade docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cache-cluster.html#cfn-elasticache-cachecluster-autominorversionupgrade
	AutoMinorVersionUpgrade *BoolExpr `json:"AutoMinorVersionUpgrade,omitempty"`
	// CacheNodeType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cache-cluster.html#cfn-elasticache-cachecluster-cachenodetype
	CacheNodeType *StringExpr `json:"CacheNodeType,omitempty" validate:"dive,required"`
	// CacheParameterGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cache-cluster.html#cfn-elasticache-cachecluster-cacheparametergroupname
	CacheParameterGroupName *StringExpr `json:"CacheParameterGroupName,omitempty"`
	// CacheSecurityGroupNames docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cache-cluster.html#cfn-elasticache-cachecluster-cachesecuritygroupnames
	CacheSecurityGroupNames *StringListExpr `json:"CacheSecurityGroupNames,omitempty"`
	// CacheSubnetGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cache-cluster.html#cfn-elasticache-cachecluster-cachesubnetgroupname
	CacheSubnetGroupName *StringExpr `json:"CacheSubnetGroupName,omitempty"`
	// ClusterName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cache-cluster.html#cfn-elasticache-cachecluster-clustername
	ClusterName *StringExpr `json:"ClusterName,omitempty"`
	// Engine docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cache-cluster.html#cfn-elasticache-cachecluster-engine
	Engine *StringExpr `json:"Engine,omitempty" validate:"dive,required"`
	// EngineVersion docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cache-cluster.html#cfn-elasticache-cachecluster-engineversion
	EngineVersion *StringExpr `json:"EngineVersion,omitempty"`
	// NotificationTopicArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cache-cluster.html#cfn-elasticache-cachecluster-notificationtopicarn
	NotificationTopicArn *StringExpr `json:"NotificationTopicArn,omitempty"`
	// NumCacheNodes docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cache-cluster.html#cfn-elasticache-cachecluster-numcachenodes
	NumCacheNodes *IntegerExpr `json:"NumCacheNodes,omitempty" validate:"dive,required"`
	// Port docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cache-cluster.html#cfn-elasticache-cachecluster-port
	Port *IntegerExpr `json:"Port,omitempty"`
	// PreferredAvailabilityZone docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cache-cluster.html#cfn-elasticache-cachecluster-preferredavailabilityzone
	PreferredAvailabilityZone *StringExpr `json:"PreferredAvailabilityZone,omitempty"`
	// PreferredAvailabilityZones docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cache-cluster.html#cfn-elasticache-cachecluster-preferredavailabilityzones
	PreferredAvailabilityZones *StringListExpr `json:"PreferredAvailabilityZones,omitempty"`
	// PreferredMaintenanceWindow docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cache-cluster.html#cfn-elasticache-cachecluster-preferredmaintenancewindow
	PreferredMaintenanceWindow *StringExpr `json:"PreferredMaintenanceWindow,omitempty"`
	// SnapshotArns docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cache-cluster.html#cfn-elasticache-cachecluster-snapshotarns
	SnapshotArns *StringListExpr `json:"SnapshotArns,omitempty"`
	// SnapshotName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cache-cluster.html#cfn-elasticache-cachecluster-snapshotname
	SnapshotName *StringExpr `json:"SnapshotName,omitempty"`
	// SnapshotRetentionLimit docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cache-cluster.html#cfn-elasticache-cachecluster-snapshotretentionlimit
	SnapshotRetentionLimit *IntegerExpr `json:"SnapshotRetentionLimit,omitempty"`
	// SnapshotWindow docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cache-cluster.html#cfn-elasticache-cachecluster-snapshotwindow
	SnapshotWindow *StringExpr `json:"SnapshotWindow,omitempty"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cache-cluster.html#cfn-elasticache-cachecluster-tags
	Tags *TagList `json:"Tags,omitempty"`
	// VPCSecurityGroupIDs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cache-cluster.html#cfn-elasticache-cachecluster-vpcsecuritygroupids
	VPCSecurityGroupIDs *StringListExpr `json:"VpcSecurityGroupIds,omitempty"`
}

// CfnResourceType returns AWS::ElastiCache::CacheCluster to implement the ResourceProperties interface
func (s ElastiCacheCacheCluster) CfnResourceType() string {

	return "AWS::ElastiCache::CacheCluster"
}

// ElastiCacheParameterGroup represents the AWS::ElastiCache::ParameterGroup CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-parameter-group.html
type ElastiCacheParameterGroup struct {
	// CacheParameterGroupFamily docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-parameter-group.html#cfn-elasticache-parametergroup-cacheparametergroupfamily
	CacheParameterGroupFamily *StringExpr `json:"CacheParameterGroupFamily,omitempty" validate:"dive,required"`
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-parameter-group.html#cfn-elasticache-parametergroup-description
	Description *StringExpr `json:"Description,omitempty" validate:"dive,required"`
	// Properties docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-parameter-group.html#cfn-elasticache-parametergroup-properties
	Properties interface{} `json:"Properties,omitempty"`
}

// CfnResourceType returns AWS::ElastiCache::ParameterGroup to implement the ResourceProperties interface
func (s ElastiCacheParameterGroup) CfnResourceType() string {

	return "AWS::ElastiCache::ParameterGroup"
}

// ElastiCacheReplicationGroup represents the AWS::ElastiCache::ReplicationGroup CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html
type ElastiCacheReplicationGroup struct {
	// AutoMinorVersionUpgrade docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-autominorversionupgrade
	AutoMinorVersionUpgrade *BoolExpr `json:"AutoMinorVersionUpgrade,omitempty"`
	// AutomaticFailoverEnabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-automaticfailoverenabled
	AutomaticFailoverEnabled *BoolExpr `json:"AutomaticFailoverEnabled,omitempty"`
	// CacheNodeType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-cachenodetype
	CacheNodeType *StringExpr `json:"CacheNodeType,omitempty"`
	// CacheParameterGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-cacheparametergroupname
	CacheParameterGroupName *StringExpr `json:"CacheParameterGroupName,omitempty"`
	// CacheSecurityGroupNames docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-cachesecuritygroupnames
	CacheSecurityGroupNames *StringListExpr `json:"CacheSecurityGroupNames,omitempty"`
	// CacheSubnetGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-cachesubnetgroupname
	CacheSubnetGroupName *StringExpr `json:"CacheSubnetGroupName,omitempty"`
	// Engine docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-engine
	Engine *StringExpr `json:"Engine,omitempty"`
	// EngineVersion docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-engineversion
	EngineVersion *StringExpr `json:"EngineVersion,omitempty"`
	// NodeGroupConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-nodegroupconfiguration
	NodeGroupConfiguration *ElastiCacheReplicationGroupNodeGroupConfigurationList `json:"NodeGroupConfiguration,omitempty"`
	// NotificationTopicArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-notificationtopicarn
	NotificationTopicArn *StringExpr `json:"NotificationTopicArn,omitempty"`
	// NumCacheClusters docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-numcacheclusters
	NumCacheClusters *IntegerExpr `json:"NumCacheClusters,omitempty"`
	// NumNodeGroups docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-numnodegroups
	NumNodeGroups *IntegerExpr `json:"NumNodeGroups,omitempty"`
	// Port docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-port
	Port *IntegerExpr `json:"Port,omitempty"`
	// PreferredCacheClusterAZs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-preferredcacheclusterazs
	PreferredCacheClusterAZs *StringListExpr `json:"PreferredCacheClusterAZs,omitempty"`
	// PreferredMaintenanceWindow docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-preferredmaintenancewindow
	PreferredMaintenanceWindow *StringExpr `json:"PreferredMaintenanceWindow,omitempty"`
	// PrimaryClusterID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-primaryclusterid
	PrimaryClusterID *StringExpr `json:"PrimaryClusterId,omitempty"`
	// ReplicasPerNodeGroup docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-replicaspernodegroup
	ReplicasPerNodeGroup *IntegerExpr `json:"ReplicasPerNodeGroup,omitempty"`
	// ReplicationGroupDescription docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-replicationgroupdescription
	ReplicationGroupDescription *StringExpr `json:"ReplicationGroupDescription,omitempty" validate:"dive,required"`
	// ReplicationGroupID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-replicationgroupid
	ReplicationGroupID *StringExpr `json:"ReplicationGroupId,omitempty"`
	// SecurityGroupIDs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-securitygroupids
	SecurityGroupIDs *StringListExpr `json:"SecurityGroupIds,omitempty"`
	// SnapshotArns docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-snapshotarns
	SnapshotArns *StringListExpr `json:"SnapshotArns,omitempty"`
	// SnapshotName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-snapshotname
	SnapshotName *StringExpr `json:"SnapshotName,omitempty"`
	// SnapshotRetentionLimit docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-snapshotretentionlimit
	SnapshotRetentionLimit *IntegerExpr `json:"SnapshotRetentionLimit,omitempty"`
	// SnapshotWindow docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-snapshotwindow
	SnapshotWindow *StringExpr `json:"SnapshotWindow,omitempty"`
	// SnapshottingClusterID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-snapshottingclusterid
	SnapshottingClusterID *StringExpr `json:"SnapshottingClusterId,omitempty"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-tags
	Tags *TagList `json:"Tags,omitempty"`
}

// CfnResourceType returns AWS::ElastiCache::ReplicationGroup to implement the ResourceProperties interface
func (s ElastiCacheReplicationGroup) CfnResourceType() string {

	return "AWS::ElastiCache::ReplicationGroup"
}

// ElastiCacheSecurityGroup represents the AWS::ElastiCache::SecurityGroup CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group.html
type ElastiCacheSecurityGroup struct {
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group.html#cfn-elasticache-securitygroup-description
	Description *StringExpr `json:"Description,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::ElastiCache::SecurityGroup to implement the ResourceProperties interface
func (s ElastiCacheSecurityGroup) CfnResourceType() string {

	return "AWS::ElastiCache::SecurityGroup"
}

// ElastiCacheSecurityGroupIngress represents the AWS::ElastiCache::SecurityGroupIngress CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html
type ElastiCacheSecurityGroupIngress struct {
	// CacheSecurityGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html#cfn-elasticache-securitygroupingress-cachesecuritygroupname
	CacheSecurityGroupName *StringExpr `json:"CacheSecurityGroupName,omitempty" validate:"dive,required"`
	// EC2SecurityGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html#cfn-elasticache-securitygroupingress-ec2securitygroupname
	EC2SecurityGroupName *StringExpr `json:"EC2SecurityGroupName,omitempty" validate:"dive,required"`
	// EC2SecurityGroupOwnerID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html#cfn-elasticache-securitygroupingress-ec2securitygroupownerid
	EC2SecurityGroupOwnerID *StringExpr `json:"EC2SecurityGroupOwnerId,omitempty"`
}

// CfnResourceType returns AWS::ElastiCache::SecurityGroupIngress to implement the ResourceProperties interface
func (s ElastiCacheSecurityGroupIngress) CfnResourceType() string {

	return "AWS::ElastiCache::SecurityGroupIngress"
}

// ElastiCacheSubnetGroup represents the AWS::ElastiCache::SubnetGroup CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-subnetgroup.html
type ElastiCacheSubnetGroup struct {
	// CacheSubnetGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-subnetgroup.html#cfn-elasticache-subnetgroup-cachesubnetgroupname
	CacheSubnetGroupName *StringExpr `json:"CacheSubnetGroupName,omitempty"`
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-subnetgroup.html#cfn-elasticache-subnetgroup-description
	Description *StringExpr `json:"Description,omitempty" validate:"dive,required"`
	// SubnetIDs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-subnetgroup.html#cfn-elasticache-subnetgroup-subnetids
	SubnetIDs *StringListExpr `json:"SubnetIds,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::ElastiCache::SubnetGroup to implement the ResourceProperties interface
func (s ElastiCacheSubnetGroup) CfnResourceType() string {

	return "AWS::ElastiCache::SubnetGroup"
}

// ElasticBeanstalkApplication represents the AWS::ElasticBeanstalk::Application CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk.html
type ElasticBeanstalkApplication struct {
	// ApplicationName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk.html#cfn-elasticbeanstalk-application-name
	ApplicationName *StringExpr `json:"ApplicationName,omitempty"`
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk.html#cfn-elasticbeanstalk-application-description
	Description *StringExpr `json:"Description,omitempty"`
}

// CfnResourceType returns AWS::ElasticBeanstalk::Application to implement the ResourceProperties interface
func (s ElasticBeanstalkApplication) CfnResourceType() string {

	return "AWS::ElasticBeanstalk::Application"
}

// ElasticBeanstalkApplicationVersion represents the AWS::ElasticBeanstalk::ApplicationVersion CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html
type ElasticBeanstalkApplicationVersion struct {
	// ApplicationName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html#cfn-elasticbeanstalk-applicationversion-applicationname
	ApplicationName *StringExpr `json:"ApplicationName,omitempty" validate:"dive,required"`
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html#cfn-elasticbeanstalk-applicationversion-description
	Description *StringExpr `json:"Description,omitempty"`
	// SourceBundle docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html#cfn-elasticbeanstalk-applicationversion-sourcebundle
	SourceBundle *ElasticBeanstalkApplicationVersionSourceBundle `json:"SourceBundle,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::ElasticBeanstalk::ApplicationVersion to implement the ResourceProperties interface
func (s ElasticBeanstalkApplicationVersion) CfnResourceType() string {

	return "AWS::ElasticBeanstalk::ApplicationVersion"
}

// ElasticBeanstalkConfigurationTemplate represents the AWS::ElasticBeanstalk::ConfigurationTemplate CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-beanstalk-configurationtemplate.html
type ElasticBeanstalkConfigurationTemplate struct {
	// ApplicationName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-beanstalk-configurationtemplate.html#cfn-elasticbeanstalk-configurationtemplate-applicationname
	ApplicationName *StringExpr `json:"ApplicationName,omitempty" validate:"dive,required"`
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-beanstalk-configurationtemplate.html#cfn-elasticbeanstalk-configurationtemplate-description
	Description *StringExpr `json:"Description,omitempty"`
	// EnvironmentID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-beanstalk-configurationtemplate.html#cfn-elasticbeanstalk-configurationtemplate-environmentid
	EnvironmentID *StringExpr `json:"EnvironmentId,omitempty"`
	// OptionSettings docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-beanstalk-configurationtemplate.html#cfn-elasticbeanstalk-configurationtemplate-optionsettings
	OptionSettings *ElasticBeanstalkConfigurationTemplateConfigurationOptionSettingList `json:"OptionSettings,omitempty"`
	// SolutionStackName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-beanstalk-configurationtemplate.html#cfn-elasticbeanstalk-configurationtemplate-solutionstackname
	SolutionStackName *StringExpr `json:"SolutionStackName,omitempty"`
	// SourceConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-beanstalk-configurationtemplate.html#cfn-elasticbeanstalk-configurationtemplate-sourceconfiguration
	SourceConfiguration *ElasticBeanstalkConfigurationTemplateSourceConfiguration `json:"SourceConfiguration,omitempty"`
}

// CfnResourceType returns AWS::ElasticBeanstalk::ConfigurationTemplate to implement the ResourceProperties interface
func (s ElasticBeanstalkConfigurationTemplate) CfnResourceType() string {

	return "AWS::ElasticBeanstalk::ConfigurationTemplate"
}

// ElasticBeanstalkEnvironment represents the AWS::ElasticBeanstalk::Environment CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html
type ElasticBeanstalkEnvironment struct {
	// ApplicationName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-applicationname
	ApplicationName *StringExpr `json:"ApplicationName,omitempty" validate:"dive,required"`
	// CNAMEPrefix docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-cnameprefix
	CNAMEPrefix *StringExpr `json:"CNAMEPrefix,omitempty"`
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-description
	Description *StringExpr `json:"Description,omitempty"`
	// EnvironmentName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-name
	EnvironmentName *StringExpr `json:"EnvironmentName,omitempty"`
	// OptionSettings docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-optionsettings
	OptionSettings *ElasticBeanstalkEnvironmentOptionSettingsList `json:"OptionSettings,omitempty"`
	// SolutionStackName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-solutionstackname
	SolutionStackName *StringExpr `json:"SolutionStackName,omitempty"`
	// TemplateName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-templatename
	TemplateName *StringExpr `json:"TemplateName,omitempty"`
	// Tier docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-tier
	Tier *ElasticBeanstalkEnvironmentTier `json:"Tier,omitempty"`
	// VersionLabel docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-versionlabel
	VersionLabel *StringExpr `json:"VersionLabel,omitempty"`
}

// CfnResourceType returns AWS::ElasticBeanstalk::Environment to implement the ResourceProperties interface
func (s ElasticBeanstalkEnvironment) CfnResourceType() string {

	return "AWS::ElasticBeanstalk::Environment"
}

// ElasticLoadBalancingLoadBalancer represents the AWS::ElasticLoadBalancing::LoadBalancer CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html
type ElasticLoadBalancingLoadBalancer struct {
	// AccessLoggingPolicy docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-accessloggingpolicy
	AccessLoggingPolicy *ElasticLoadBalancingLoadBalancerAccessLoggingPolicy `json:"AccessLoggingPolicy,omitempty"`
	// AppCookieStickinessPolicy docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-appcookiestickinesspolicy
	AppCookieStickinessPolicy *ElasticLoadBalancingLoadBalancerAppCookieStickinessPolicyList `json:"AppCookieStickinessPolicy,omitempty"`
	// AvailabilityZones docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-availabilityzones
	AvailabilityZones *StringListExpr `json:"AvailabilityZones,omitempty"`
	// ConnectionDrainingPolicy docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-connectiondrainingpolicy
	ConnectionDrainingPolicy *ElasticLoadBalancingLoadBalancerConnectionDrainingPolicy `json:"ConnectionDrainingPolicy,omitempty"`
	// ConnectionSettings docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-connectionsettings
	ConnectionSettings *ElasticLoadBalancingLoadBalancerConnectionSettings `json:"ConnectionSettings,omitempty"`
	// CrossZone docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-crosszone
	CrossZone *BoolExpr `json:"CrossZone,omitempty"`
	// HealthCheck docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-healthcheck
	HealthCheck *ElasticLoadBalancingLoadBalancerHealthCheck `json:"HealthCheck,omitempty"`
	// Instances docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-instances
	Instances *StringListExpr `json:"Instances,omitempty"`
	// LBCookieStickinessPolicy docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-lbcookiestickinesspolicy
	LBCookieStickinessPolicy *ElasticLoadBalancingLoadBalancerLBCookieStickinessPolicyList `json:"LBCookieStickinessPolicy,omitempty"`
	// Listeners docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-listeners
	Listeners *ElasticLoadBalancingLoadBalancerListenersList `json:"Listeners,omitempty" validate:"dive,required"`
	// LoadBalancerName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-elbname
	LoadBalancerName *StringExpr `json:"LoadBalancerName,omitempty"`
	// Policies docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-policies
	Policies *ElasticLoadBalancingLoadBalancerPoliciesList `json:"Policies,omitempty"`
	// Scheme docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-scheme
	Scheme *StringExpr `json:"Scheme,omitempty"`
	// SecurityGroups docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-securitygroups
	SecurityGroups *StringListExpr `json:"SecurityGroups,omitempty"`
	// Subnets docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-subnets
	Subnets *StringListExpr `json:"Subnets,omitempty"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-elasticloadbalancing-loadbalancer-tags
	Tags *TagList `json:"Tags,omitempty"`
}

// CfnResourceType returns AWS::ElasticLoadBalancing::LoadBalancer to implement the ResourceProperties interface
func (s ElasticLoadBalancingLoadBalancer) CfnResourceType() string {

	return "AWS::ElasticLoadBalancing::LoadBalancer"
}

// ElasticLoadBalancingV2Listener represents the AWS::ElasticLoadBalancingV2::Listener CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html
type ElasticLoadBalancingV2Listener struct {
	// Certificates docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html#cfn-elasticloadbalancingv2-listener-certificates
	Certificates *ElasticLoadBalancingV2ListenerCertificateList `json:"Certificates,omitempty"`
	// DefaultActions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html#cfn-elasticloadbalancingv2-listener-defaultactions
	DefaultActions *ElasticLoadBalancingV2ListenerActionList `json:"DefaultActions,omitempty" validate:"dive,required"`
	// LoadBalancerArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html#cfn-elasticloadbalancingv2-listener-loadbalancerarn
	LoadBalancerArn *StringExpr `json:"LoadBalancerArn,omitempty" validate:"dive,required"`
	// Port docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html#cfn-elasticloadbalancingv2-listener-port
	Port *IntegerExpr `json:"Port,omitempty" validate:"dive,required"`
	// Protocol docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html#cfn-elasticloadbalancingv2-listener-protocol
	Protocol *StringExpr `json:"Protocol,omitempty" validate:"dive,required"`
	// SslPolicy docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html#cfn-elasticloadbalancingv2-listener-sslpolicy
	SslPolicy *StringExpr `json:"SslPolicy,omitempty"`
}

// CfnResourceType returns AWS::ElasticLoadBalancingV2::Listener to implement the ResourceProperties interface
func (s ElasticLoadBalancingV2Listener) CfnResourceType() string {

	return "AWS::ElasticLoadBalancingV2::Listener"
}

// ElasticLoadBalancingV2ListenerRule represents the AWS::ElasticLoadBalancingV2::ListenerRule CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html
type ElasticLoadBalancingV2ListenerRule struct {
	// Actions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html#cfn-elasticloadbalancingv2-listenerrule-actions
	Actions *ElasticLoadBalancingV2ListenerRuleActionList `json:"Actions,omitempty" validate:"dive,required"`
	// Conditions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html#cfn-elasticloadbalancingv2-listenerrule-conditions
	Conditions *ElasticLoadBalancingV2ListenerRuleRuleConditionList `json:"Conditions,omitempty" validate:"dive,required"`
	// ListenerArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html#cfn-elasticloadbalancingv2-listenerrule-listenerarn
	ListenerArn *StringExpr `json:"ListenerArn,omitempty" validate:"dive,required"`
	// Priority docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html#cfn-elasticloadbalancingv2-listenerrule-priority
	Priority *IntegerExpr `json:"Priority,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::ElasticLoadBalancingV2::ListenerRule to implement the ResourceProperties interface
func (s ElasticLoadBalancingV2ListenerRule) CfnResourceType() string {

	return "AWS::ElasticLoadBalancingV2::ListenerRule"
}

// ElasticLoadBalancingV2LoadBalancer represents the AWS::ElasticLoadBalancingV2::LoadBalancer CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html
type ElasticLoadBalancingV2LoadBalancer struct {
	// IPAddressType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-ipaddresstype
	IPAddressType *StringExpr `json:"IpAddressType,omitempty"`
	// LoadBalancerAttributes docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-loadbalancerattributes
	LoadBalancerAttributes *ElasticLoadBalancingV2LoadBalancerLoadBalancerAttributeList `json:"LoadBalancerAttributes,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-name
	Name *StringExpr `json:"Name,omitempty"`
	// Scheme docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-scheme
	Scheme *StringExpr `json:"Scheme,omitempty"`
	// SecurityGroups docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-securitygroups
	SecurityGroups *StringListExpr `json:"SecurityGroups,omitempty"`
	// Subnets docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-subnets
	Subnets *StringListExpr `json:"Subnets,omitempty"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-tags
	Tags *TagList `json:"Tags,omitempty"`
}

// CfnResourceType returns AWS::ElasticLoadBalancingV2::LoadBalancer to implement the ResourceProperties interface
func (s ElasticLoadBalancingV2LoadBalancer) CfnResourceType() string {

	return "AWS::ElasticLoadBalancingV2::LoadBalancer"
}

// ElasticLoadBalancingV2TargetGroup represents the AWS::ElasticLoadBalancingV2::TargetGroup CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html
type ElasticLoadBalancingV2TargetGroup struct {
	// HealthCheckIntervalSeconds docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-healthcheckintervalseconds
	HealthCheckIntervalSeconds *IntegerExpr `json:"HealthCheckIntervalSeconds,omitempty"`
	// HealthCheckPath docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-healthcheckpath
	HealthCheckPath *StringExpr `json:"HealthCheckPath,omitempty"`
	// HealthCheckPort docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-healthcheckport
	HealthCheckPort *StringExpr `json:"HealthCheckPort,omitempty"`
	// HealthCheckProtocol docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-healthcheckprotocol
	HealthCheckProtocol *StringExpr `json:"HealthCheckProtocol,omitempty"`
	// HealthCheckTimeoutSeconds docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-healthchecktimeoutseconds
	HealthCheckTimeoutSeconds *IntegerExpr `json:"HealthCheckTimeoutSeconds,omitempty"`
	// HealthyThresholdCount docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-healthythresholdcount
	HealthyThresholdCount *IntegerExpr `json:"HealthyThresholdCount,omitempty"`
	// Matcher docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-matcher
	Matcher *ElasticLoadBalancingV2TargetGroupMatcher `json:"Matcher,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-name
	Name *StringExpr `json:"Name,omitempty"`
	// Port docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-port
	Port *IntegerExpr `json:"Port,omitempty" validate:"dive,required"`
	// Protocol docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-protocol
	Protocol *StringExpr `json:"Protocol,omitempty" validate:"dive,required"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-tags
	Tags *TagList `json:"Tags,omitempty"`
	// TargetGroupAttributes docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-targetgroupattributes
	TargetGroupAttributes *ElasticLoadBalancingV2TargetGroupTargetGroupAttributeList `json:"TargetGroupAttributes,omitempty"`
	// Targets docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-targets
	Targets *ElasticLoadBalancingV2TargetGroupTargetDescriptionList `json:"Targets,omitempty"`
	// UnhealthyThresholdCount docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-unhealthythresholdcount
	UnhealthyThresholdCount *IntegerExpr `json:"UnhealthyThresholdCount,omitempty"`
	// VPCID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-vpcid
	VPCID *StringExpr `json:"VpcId,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::ElasticLoadBalancingV2::TargetGroup to implement the ResourceProperties interface
func (s ElasticLoadBalancingV2TargetGroup) CfnResourceType() string {

	return "AWS::ElasticLoadBalancingV2::TargetGroup"
}

// ElasticsearchDomain represents the AWS::Elasticsearch::Domain CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html
type ElasticsearchDomain struct {
	// AccessPolicies docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-accesspolicies
	AccessPolicies interface{} `json:"AccessPolicies,omitempty"`
	// AdvancedOptions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-advancedoptions
	AdvancedOptions interface{} `json:"AdvancedOptions,omitempty"`
	// DomainName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-domainname
	DomainName *StringExpr `json:"DomainName,omitempty"`
	// EBSOptions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-ebsoptions
	EBSOptions *ElasticsearchDomainEBSOptions `json:"EBSOptions,omitempty"`
	// ElasticsearchClusterConfig docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-elasticsearchclusterconfig
	ElasticsearchClusterConfig *ElasticsearchDomainElasticsearchClusterConfig `json:"ElasticsearchClusterConfig,omitempty"`
	// ElasticsearchVersion docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-elasticsearchversion
	ElasticsearchVersion *StringExpr `json:"ElasticsearchVersion,omitempty"`
	// SnapshotOptions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-snapshotoptions
	SnapshotOptions *ElasticsearchDomainSnapshotOptions `json:"SnapshotOptions,omitempty"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-tags
	Tags *TagList `json:"Tags,omitempty"`
}

// CfnResourceType returns AWS::Elasticsearch::Domain to implement the ResourceProperties interface
func (s ElasticsearchDomain) CfnResourceType() string {

	return "AWS::Elasticsearch::Domain"
}

// EventsRule represents the AWS::Events::Rule CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html
type EventsRule struct {
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-description
	Description *StringExpr `json:"Description,omitempty"`
	// EventPattern docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-eventpattern
	EventPattern interface{} `json:"EventPattern,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-name
	Name *StringExpr `json:"Name,omitempty"`
	// RoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-rolearn
	RoleArn *StringExpr `json:"RoleArn,omitempty"`
	// ScheduleExpression docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-scheduleexpression
	ScheduleExpression *StringExpr `json:"ScheduleExpression,omitempty"`
	// State docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-state
	State *StringExpr `json:"State,omitempty"`
	// Targets docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-targets
	Targets *EventsRuleTargetList `json:"Targets,omitempty"`
}

// CfnResourceType returns AWS::Events::Rule to implement the ResourceProperties interface
func (s EventsRule) CfnResourceType() string {

	return "AWS::Events::Rule"
}

// GameLiftAlias represents the AWS::GameLift::Alias CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html
type GameLiftAlias struct {
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html#cfn-gamelift-alias-description
	Description *StringExpr `json:"Description,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html#cfn-gamelift-alias-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// RoutingStrategy docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html#cfn-gamelift-alias-routingstrategy
	RoutingStrategy *GameLiftAliasRoutingStrategy `json:"RoutingStrategy,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::GameLift::Alias to implement the ResourceProperties interface
func (s GameLiftAlias) CfnResourceType() string {

	return "AWS::GameLift::Alias"
}

// GameLiftBuild represents the AWS::GameLift::Build CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-build.html
type GameLiftBuild struct {
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-build.html#cfn-gamelift-build-name
	Name *StringExpr `json:"Name,omitempty"`
	// StorageLocation docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-build.html#cfn-gamelift-build-storagelocation
	StorageLocation *GameLiftBuildS3Location `json:"StorageLocation,omitempty"`
	// Version docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-build.html#cfn-gamelift-build-version
	Version *StringExpr `json:"Version,omitempty"`
}

// CfnResourceType returns AWS::GameLift::Build to implement the ResourceProperties interface
func (s GameLiftBuild) CfnResourceType() string {

	return "AWS::GameLift::Build"
}

// GameLiftFleet represents the AWS::GameLift::Fleet CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html
type GameLiftFleet struct {
	// BuildID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-buildid
	BuildID *StringExpr `json:"BuildId,omitempty" validate:"dive,required"`
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-description
	Description *StringExpr `json:"Description,omitempty"`
	// DesiredEC2Instances docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-desiredec2instances
	DesiredEC2Instances *IntegerExpr `json:"DesiredEC2Instances,omitempty" validate:"dive,required"`
	// EC2InboundPermissions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-ec2inboundpermissions
	EC2InboundPermissions *GameLiftFleetIPPermissionList `json:"EC2InboundPermissions,omitempty"`
	// EC2InstanceType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-ec2instancetype
	EC2InstanceType *StringExpr `json:"EC2InstanceType,omitempty" validate:"dive,required"`
	// LogPaths docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-logpaths
	LogPaths *StringListExpr `json:"LogPaths,omitempty"`
	// MaxSize docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-maxsize
	MaxSize *IntegerExpr `json:"MaxSize,omitempty"`
	// MinSize docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-minsize
	MinSize *IntegerExpr `json:"MinSize,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// ServerLaunchParameters docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-serverlaunchparameters
	ServerLaunchParameters *StringExpr `json:"ServerLaunchParameters,omitempty"`
	// ServerLaunchPath docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-serverlaunchpath
	ServerLaunchPath *StringExpr `json:"ServerLaunchPath,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::GameLift::Fleet to implement the ResourceProperties interface
func (s GameLiftFleet) CfnResourceType() string {

	return "AWS::GameLift::Fleet"
}

// IAMAccessKey represents the AWS::IAM::AccessKey CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html
type IAMAccessKey struct {
	// Serial docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html#cfn-iam-accesskey-serial
	Serial *IntegerExpr `json:"Serial,omitempty"`
	// Status docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html#cfn-iam-accesskey-status
	Status *StringExpr `json:"Status,omitempty"`
	// UserName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html#cfn-iam-accesskey-username
	UserName *StringExpr `json:"UserName,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::IAM::AccessKey to implement the ResourceProperties interface
func (s IAMAccessKey) CfnResourceType() string {

	return "AWS::IAM::AccessKey"
}

// IAMGroup represents the AWS::IAM::Group CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html
type IAMGroup struct {
	// GroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html#cfn-iam-group-groupname
	GroupName *StringExpr `json:"GroupName,omitempty"`
	// ManagedPolicyArns docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html#cfn-iam-group-managepolicyarns
	ManagedPolicyArns *StringListExpr `json:"ManagedPolicyArns,omitempty"`
	// Path docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html#cfn-iam-group-path
	Path *StringExpr `json:"Path,omitempty"`
	// Policies docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html#cfn-iam-group-policies
	Policies *IAMGroupPolicyList `json:"Policies,omitempty"`
}

// CfnResourceType returns AWS::IAM::Group to implement the ResourceProperties interface
func (s IAMGroup) CfnResourceType() string {

	return "AWS::IAM::Group"
}

// IAMInstanceProfile represents the AWS::IAM::InstanceProfile CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html
type IAMInstanceProfile struct {
	// InstanceProfileName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html#cfn-iam-instanceprofile-instanceprofilename
	InstanceProfileName *StringExpr `json:"InstanceProfileName,omitempty"`
	// Path docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html#cfn-iam-instanceprofile-path
	Path *StringExpr `json:"Path,omitempty"`
	// Roles docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html#cfn-iam-instanceprofile-roles
	Roles *StringListExpr `json:"Roles,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::IAM::InstanceProfile to implement the ResourceProperties interface
func (s IAMInstanceProfile) CfnResourceType() string {

	return "AWS::IAM::InstanceProfile"
}

// IAMManagedPolicy represents the AWS::IAM::ManagedPolicy CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html
type IAMManagedPolicy struct {
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html#cfn-iam-managedpolicy-description
	Description *StringExpr `json:"Description,omitempty"`
	// Groups docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html#cfn-iam-managedpolicy-groups
	Groups *StringListExpr `json:"Groups,omitempty"`
	// ManagedPolicyName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html#cfn-iam-managedpolicy-managedpolicyname
	ManagedPolicyName *StringExpr `json:"ManagedPolicyName,omitempty"`
	// Path docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html#cfn-ec2-dhcpoptions-path
	Path *StringExpr `json:"Path,omitempty"`
	// PolicyDocument docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html#cfn-iam-managedpolicy-policydocument
	PolicyDocument interface{} `json:"PolicyDocument,omitempty" validate:"dive,required"`
	// Roles docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html#cfn-iam-managedpolicy-roles
	Roles *StringListExpr `json:"Roles,omitempty"`
	// Users docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html#cfn-iam-managedpolicy-users
	Users *StringListExpr `json:"Users,omitempty"`
}

// CfnResourceType returns AWS::IAM::ManagedPolicy to implement the ResourceProperties interface
func (s IAMManagedPolicy) CfnResourceType() string {

	return "AWS::IAM::ManagedPolicy"
}

// IAMPolicy represents the AWS::IAM::Policy CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html
type IAMPolicy struct {
	// Groups docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html#cfn-iam-policy-groups
	Groups *StringListExpr `json:"Groups,omitempty"`
	// PolicyDocument docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html#cfn-iam-policy-policydocument
	PolicyDocument interface{} `json:"PolicyDocument,omitempty" validate:"dive,required"`
	// PolicyName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html#cfn-iam-policy-policyname
	PolicyName *StringExpr `json:"PolicyName,omitempty" validate:"dive,required"`
	// Roles docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html#cfn-iam-policy-roles
	Roles *StringListExpr `json:"Roles,omitempty"`
	// Users docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html#cfn-iam-policy-users
	Users *StringListExpr `json:"Users,omitempty"`
}

// CfnResourceType returns AWS::IAM::Policy to implement the ResourceProperties interface
func (s IAMPolicy) CfnResourceType() string {

	return "AWS::IAM::Policy"
}

// IAMRole represents the AWS::IAM::Role CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html
type IAMRole struct {
	// AssumeRolePolicyDocument docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html#cfn-iam-role-assumerolepolicydocument
	AssumeRolePolicyDocument interface{} `json:"AssumeRolePolicyDocument,omitempty" validate:"dive,required"`
	// ManagedPolicyArns docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html#cfn-iam-role-managepolicyarns
	ManagedPolicyArns *StringListExpr `json:"ManagedPolicyArns,omitempty"`
	// Path docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html#cfn-iam-role-path
	Path *StringExpr `json:"Path,omitempty"`
	// Policies docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html#cfn-iam-role-policies
	Policies *IAMRolePolicyList `json:"Policies,omitempty"`
	// RoleName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html#cfn-iam-role-rolename
	RoleName *StringExpr `json:"RoleName,omitempty"`
}

// CfnResourceType returns AWS::IAM::Role to implement the ResourceProperties interface
func (s IAMRole) CfnResourceType() string {

	return "AWS::IAM::Role"
}

// IAMUser represents the AWS::IAM::User CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html
type IAMUser struct {
	// Groups docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html#cfn-iam-user-groups
	Groups *StringListExpr `json:"Groups,omitempty"`
	// LoginProfile docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html#cfn-iam-user-loginprofile
	LoginProfile *IAMUserLoginProfile `json:"LoginProfile,omitempty"`
	// ManagedPolicyArns docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html#cfn-iam-user-managepolicyarns
	ManagedPolicyArns *StringListExpr `json:"ManagedPolicyArns,omitempty"`
	// Path docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html#cfn-iam-user-path
	Path *StringExpr `json:"Path,omitempty"`
	// Policies docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html#cfn-iam-user-policies
	Policies *IAMUserPolicyList `json:"Policies,omitempty"`
	// UserName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html#cfn-iam-user-username
	UserName *StringExpr `json:"UserName,omitempty"`
}

// CfnResourceType returns AWS::IAM::User to implement the ResourceProperties interface
func (s IAMUser) CfnResourceType() string {

	return "AWS::IAM::User"
}

// IAMUserToGroupAddition represents the AWS::IAM::UserToGroupAddition CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html
type IAMUserToGroupAddition struct {
	// GroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html#cfn-iam-addusertogroup-groupname
	GroupName *StringExpr `json:"GroupName,omitempty" validate:"dive,required"`
	// Users docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html#cfn-iam-addusertogroup-users
	Users *StringListExpr `json:"Users,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::IAM::UserToGroupAddition to implement the ResourceProperties interface
func (s IAMUserToGroupAddition) CfnResourceType() string {

	return "AWS::IAM::UserToGroupAddition"
}

// IoTCertificate represents the AWS::IoT::Certificate CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-certificate.html
type IoTCertificate struct {
	// CertificateSigningRequest docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-certificate.html#cfn-iot-certificate-certificatesigningrequest
	CertificateSigningRequest *StringExpr `json:"CertificateSigningRequest,omitempty" validate:"dive,required"`
	// Status docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-certificate.html#cfn-iot-certificate-status
	Status *StringExpr `json:"Status,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::IoT::Certificate to implement the ResourceProperties interface
func (s IoTCertificate) CfnResourceType() string {

	return "AWS::IoT::Certificate"
}

// IoTPolicy represents the AWS::IoT::Policy CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policy.html
type IoTPolicy struct {
	// PolicyDocument docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policy.html#cfn-iot-policy-policydocument
	PolicyDocument interface{} `json:"PolicyDocument,omitempty" validate:"dive,required"`
	// PolicyName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policy.html#cfn-iot-policy-policyname
	PolicyName *StringExpr `json:"PolicyName,omitempty"`
}

// CfnResourceType returns AWS::IoT::Policy to implement the ResourceProperties interface
func (s IoTPolicy) CfnResourceType() string {

	return "AWS::IoT::Policy"
}

// IoTPolicyPrincipalAttachment represents the AWS::IoT::PolicyPrincipalAttachment CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policyprincipalattachment.html
type IoTPolicyPrincipalAttachment struct {
	// PolicyName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policyprincipalattachment.html#cfn-iot-policyprincipalattachment-policyname
	PolicyName *StringExpr `json:"PolicyName,omitempty" validate:"dive,required"`
	// Principal docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policyprincipalattachment.html#cfn-iot-policyprincipalattachment-principal
	Principal *StringExpr `json:"Principal,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::IoT::PolicyPrincipalAttachment to implement the ResourceProperties interface
func (s IoTPolicyPrincipalAttachment) CfnResourceType() string {

	return "AWS::IoT::PolicyPrincipalAttachment"
}

// IoTThing represents the AWS::IoT::Thing CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thing.html
type IoTThing struct {
	// AttributePayload docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thing.html#cfn-iot-thing-attributepayload
	AttributePayload *IoTThingAttributePayload `json:"AttributePayload,omitempty"`
	// ThingName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thing.html#cfn-iot-thing-thingname
	ThingName *StringExpr `json:"ThingName,omitempty"`
}

// CfnResourceType returns AWS::IoT::Thing to implement the ResourceProperties interface
func (s IoTThing) CfnResourceType() string {

	return "AWS::IoT::Thing"
}

// IoTThingPrincipalAttachment represents the AWS::IoT::ThingPrincipalAttachment CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thingprincipalattachment.html
type IoTThingPrincipalAttachment struct {
	// Principal docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thingprincipalattachment.html#cfn-iot-thingprincipalattachment-principal
	Principal *StringExpr `json:"Principal,omitempty" validate:"dive,required"`
	// ThingName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thingprincipalattachment.html#cfn-iot-thingprincipalattachment-thingname
	ThingName *StringExpr `json:"ThingName,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::IoT::ThingPrincipalAttachment to implement the ResourceProperties interface
func (s IoTThingPrincipalAttachment) CfnResourceType() string {

	return "AWS::IoT::ThingPrincipalAttachment"
}

// IoTTopicRule represents the AWS::IoT::TopicRule CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-topicrule.html
type IoTTopicRule struct {
	// RuleName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-topicrule.html#cfn-iot-topicrule-rulename
	RuleName *StringExpr `json:"RuleName,omitempty"`
	// TopicRulePayload docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-topicrule.html#cfn-iot-topicrule-topicrulepayload
	TopicRulePayload *IoTTopicRuleTopicRulePayload `json:"TopicRulePayload,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::IoT::TopicRule to implement the ResourceProperties interface
func (s IoTTopicRule) CfnResourceType() string {

	return "AWS::IoT::TopicRule"
}

// KMSAlias represents the AWS::KMS::Alias CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html
type KMSAlias struct {
	// AliasName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html#cfn-kms-alias-aliasname
	AliasName *StringExpr `json:"AliasName,omitempty" validate:"dive,required"`
	// TargetKeyID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html#cfn-kms-alias-targetkeyid
	TargetKeyID *StringExpr `json:"TargetKeyId,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::KMS::Alias to implement the ResourceProperties interface
func (s KMSAlias) CfnResourceType() string {

	return "AWS::KMS::Alias"
}

// KMSKey represents the AWS::KMS::Key CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html
type KMSKey struct {
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-description
	Description *StringExpr `json:"Description,omitempty"`
	// EnableKeyRotation docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-enablekeyrotation
	EnableKeyRotation *BoolExpr `json:"EnableKeyRotation,omitempty"`
	// Enabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-enabled
	Enabled *BoolExpr `json:"Enabled,omitempty"`
	// KeyPolicy docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keypolicy
	KeyPolicy interface{} `json:"KeyPolicy,omitempty" validate:"dive,required"`
	// KeyUsage docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keyusage
	KeyUsage *StringExpr `json:"KeyUsage,omitempty"`
}

// CfnResourceType returns AWS::KMS::Key to implement the ResourceProperties interface
func (s KMSKey) CfnResourceType() string {

	return "AWS::KMS::Key"
}

// KinesisStream represents the AWS::Kinesis::Stream CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-stream.html
type KinesisStream struct {
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-stream.html#cfn-kinesis-stream-name
	Name *StringExpr `json:"Name,omitempty"`
	// ShardCount docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-stream.html#cfn-kinesis-stream-shardcount
	ShardCount *IntegerExpr `json:"ShardCount,omitempty" validate:"dive,required"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-stream.html#cfn-kinesis-stream-tags
	Tags *TagList `json:"Tags,omitempty"`
}

// CfnResourceType returns AWS::Kinesis::Stream to implement the ResourceProperties interface
func (s KinesisStream) CfnResourceType() string {

	return "AWS::Kinesis::Stream"
}

// KinesisFirehoseDeliveryStream represents the AWS::KinesisFirehose::DeliveryStream CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html
type KinesisFirehoseDeliveryStream struct {
	// DeliveryStreamName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverstream-deliverystreamname
	DeliveryStreamName *StringExpr `json:"DeliveryStreamName,omitempty"`
	// ElasticsearchDestinationConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverstream-elasticsearchdestinationconfiguration
	ElasticsearchDestinationConfiguration *KinesisFirehoseDeliveryStreamElasticsearchDestinationConfiguration `json:"ElasticsearchDestinationConfiguration,omitempty"`
	// RedshiftDestinationConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-redshiftdestinationconfiguration
	RedshiftDestinationConfiguration *KinesisFirehoseDeliveryStreamRedshiftDestinationConfiguration `json:"RedshiftDestinationConfiguration,omitempty"`
	// S3DestinationConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-s3destinationconfiguration
	S3DestinationConfiguration *KinesisFirehoseDeliveryStreamS3DestinationConfiguration `json:"S3DestinationConfiguration,omitempty"`
}

// CfnResourceType returns AWS::KinesisFirehose::DeliveryStream to implement the ResourceProperties interface
func (s KinesisFirehoseDeliveryStream) CfnResourceType() string {

	return "AWS::KinesisFirehose::DeliveryStream"
}

// LambdaAlias represents the AWS::Lambda::Alias CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-alias.html
type LambdaAlias struct {
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-alias.html#cfn-lambda-alias-description
	Description *StringExpr `json:"Description,omitempty"`
	// FunctionName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-alias.html#cfn-lambda-alias-functionname
	FunctionName *StringExpr `json:"FunctionName,omitempty" validate:"dive,required"`
	// FunctionVersion docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-alias.html#cfn-lambda-alias-functionversion
	FunctionVersion *StringExpr `json:"FunctionVersion,omitempty" validate:"dive,required"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-alias.html#cfn-lambda-alias-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::Lambda::Alias to implement the ResourceProperties interface
func (s LambdaAlias) CfnResourceType() string {

	return "AWS::Lambda::Alias"
}

// LambdaEventSourceMapping represents the AWS::Lambda::EventSourceMapping CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html
type LambdaEventSourceMapping struct {
	// BatchSize docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-batchsize
	BatchSize *IntegerExpr `json:"BatchSize,omitempty"`
	// Enabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-enabled
	Enabled *BoolExpr `json:"Enabled,omitempty"`
	// EventSourceArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-eventsourcearn
	EventSourceArn *StringExpr `json:"EventSourceArn,omitempty" validate:"dive,required"`
	// FunctionName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-functionname
	FunctionName *StringExpr `json:"FunctionName,omitempty" validate:"dive,required"`
	// StartingPosition docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html#cfn-lambda-eventsourcemapping-startingposition
	StartingPosition *StringExpr `json:"StartingPosition,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::Lambda::EventSourceMapping to implement the ResourceProperties interface
func (s LambdaEventSourceMapping) CfnResourceType() string {

	return "AWS::Lambda::EventSourceMapping"
}

// LambdaFunction represents the AWS::Lambda::Function CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html
type LambdaFunction struct {
	// Code docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-code
	Code *LambdaFunctionCode `json:"Code,omitempty" validate:"dive,required"`
	// DeadLetterConfig docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-deadletterconfig
	DeadLetterConfig *LambdaFunctionDeadLetterConfig `json:"DeadLetterConfig,omitempty"`
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-description
	Description *StringExpr `json:"Description,omitempty"`
	// Environment docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-environment
	Environment *LambdaFunctionEnvironment `json:"Environment,omitempty"`
	// FunctionName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-functionname
	FunctionName *StringExpr `json:"FunctionName,omitempty"`
	// Handler docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-handler
	Handler *StringExpr `json:"Handler,omitempty" validate:"dive,required"`
	// KmsKeyArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-kmskeyarn
	KmsKeyArn *StringExpr `json:"KmsKeyArn,omitempty"`
	// MemorySize docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-memorysize
	MemorySize *IntegerExpr `json:"MemorySize,omitempty"`
	// Role docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-role
	Role *StringExpr `json:"Role,omitempty" validate:"dive,required"`
	// Runtime docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-runtime
	Runtime *StringExpr `json:"Runtime,omitempty" validate:"dive,required"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-tags
	Tags *TagList `json:"Tags,omitempty"`
	// Timeout docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-timeout
	Timeout *IntegerExpr `json:"Timeout,omitempty"`
	// TracingConfig docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-tracingconfig
	TracingConfig *LambdaFunctionTracingConfig `json:"TracingConfig,omitempty"`
	// VPCConfig docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-vpcconfig
	VPCConfig *LambdaFunctionVPCConfig `json:"VpcConfig,omitempty"`
}

// CfnResourceType returns AWS::Lambda::Function to implement the ResourceProperties interface
func (s LambdaFunction) CfnResourceType() string {

	return "AWS::Lambda::Function"
}

// LambdaPermission represents the AWS::Lambda::Permission CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html
type LambdaPermission struct {
	// Action docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-action
	Action *StringExpr `json:"Action,omitempty" validate:"dive,required"`
	// FunctionName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-functionname
	FunctionName *StringExpr `json:"FunctionName,omitempty" validate:"dive,required"`
	// Principal docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-principal
	Principal *StringExpr `json:"Principal,omitempty" validate:"dive,required"`
	// SourceAccount docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-sourceaccount
	SourceAccount *StringExpr `json:"SourceAccount,omitempty"`
	// SourceArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-sourcearn
	SourceArn *StringExpr `json:"SourceArn,omitempty"`
}

// CfnResourceType returns AWS::Lambda::Permission to implement the ResourceProperties interface
func (s LambdaPermission) CfnResourceType() string {

	return "AWS::Lambda::Permission"
}

// LambdaVersion represents the AWS::Lambda::Version CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-version.html
type LambdaVersion struct {
	// CodeSha256 docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-version.html#cfn-lambda-version-codesha256
	CodeSha256 *StringExpr `json:"CodeSha256,omitempty"`
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-version.html#cfn-lambda-version-description
	Description *StringExpr `json:"Description,omitempty"`
	// FunctionName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-version.html#cfn-lambda-version-functionname
	FunctionName *StringExpr `json:"FunctionName,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::Lambda::Version to implement the ResourceProperties interface
func (s LambdaVersion) CfnResourceType() string {

	return "AWS::Lambda::Version"
}

// LogsDestination represents the AWS::Logs::Destination CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html
type LogsDestination struct {
	// DestinationName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html#cfn-logs-destination-destinationname
	DestinationName *StringExpr `json:"DestinationName,omitempty" validate:"dive,required"`
	// DestinationPolicy docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html#cfn-logs-destination-destinationpolicy
	DestinationPolicy *StringExpr `json:"DestinationPolicy,omitempty" validate:"dive,required"`
	// RoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html#cfn-logs-destination-rolearn
	RoleArn *StringExpr `json:"RoleArn,omitempty" validate:"dive,required"`
	// TargetArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html#cfn-logs-destination-targetarn
	TargetArn *StringExpr `json:"TargetArn,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::Logs::Destination to implement the ResourceProperties interface
func (s LogsDestination) CfnResourceType() string {

	return "AWS::Logs::Destination"
}

// LogsLogGroup represents the AWS::Logs::LogGroup CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html
type LogsLogGroup struct {
	// LogGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html#cfn-cwl-loggroup-loggroupname
	LogGroupName *StringExpr `json:"LogGroupName,omitempty"`
	// RetentionInDays docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html#cfn-cwl-loggroup-retentionindays
	RetentionInDays *IntegerExpr `json:"RetentionInDays,omitempty"`
}

// CfnResourceType returns AWS::Logs::LogGroup to implement the ResourceProperties interface
func (s LogsLogGroup) CfnResourceType() string {

	return "AWS::Logs::LogGroup"
}

// LogsLogStream represents the AWS::Logs::LogStream CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-logstream.html
type LogsLogStream struct {
	// LogGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-logstream.html#cfn-logs-logstream-loggroupname
	LogGroupName *StringExpr `json:"LogGroupName,omitempty" validate:"dive,required"`
	// LogStreamName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-logstream.html#cfn-logs-logstream-logstreamname
	LogStreamName *StringExpr `json:"LogStreamName,omitempty"`
}

// CfnResourceType returns AWS::Logs::LogStream to implement the ResourceProperties interface
func (s LogsLogStream) CfnResourceType() string {

	return "AWS::Logs::LogStream"
}

// LogsMetricFilter represents the AWS::Logs::MetricFilter CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html
type LogsMetricFilter struct {
	// FilterPattern docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html#cfn-cwl-metricfilter-filterpattern
	FilterPattern *StringExpr `json:"FilterPattern,omitempty" validate:"dive,required"`
	// LogGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html#cfn-cwl-metricfilter-loggroupname
	LogGroupName *StringExpr `json:"LogGroupName,omitempty" validate:"dive,required"`
	// MetricTransformations docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html#cfn-cwl-metricfilter-metrictransformations
	MetricTransformations *LogsMetricFilterMetricTransformationList `json:"MetricTransformations,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::Logs::MetricFilter to implement the ResourceProperties interface
func (s LogsMetricFilter) CfnResourceType() string {

	return "AWS::Logs::MetricFilter"
}

// LogsSubscriptionFilter represents the AWS::Logs::SubscriptionFilter CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html
type LogsSubscriptionFilter struct {
	// DestinationArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-cwl-subscriptionfilter-destinationarn
	DestinationArn *StringExpr `json:"DestinationArn,omitempty" validate:"dive,required"`
	// FilterPattern docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-cwl-subscriptionfilter-filterpattern
	FilterPattern *StringExpr `json:"FilterPattern,omitempty" validate:"dive,required"`
	// LogGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-cwl-subscriptionfilter-loggroupname
	LogGroupName *StringExpr `json:"LogGroupName,omitempty" validate:"dive,required"`
	// RoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-cwl-subscriptionfilter-rolearn
	RoleArn *StringExpr `json:"RoleArn,omitempty"`
}

// CfnResourceType returns AWS::Logs::SubscriptionFilter to implement the ResourceProperties interface
func (s LogsSubscriptionFilter) CfnResourceType() string {

	return "AWS::Logs::SubscriptionFilter"
}

// OpsWorksApp represents the AWS::OpsWorks::App CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html
type OpsWorksApp struct {
	// AppSource docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-appsource
	AppSource *OpsWorksAppSource `json:"AppSource,omitempty"`
	// Attributes docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-attributes
	Attributes interface{} `json:"Attributes,omitempty"`
	// DataSources docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-datasources
	DataSources *OpsWorksAppDataSourceList `json:"DataSources,omitempty"`
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-description
	Description *StringExpr `json:"Description,omitempty"`
	// Domains docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-domains
	Domains *StringListExpr `json:"Domains,omitempty"`
	// EnableSsl docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-enablessl
	EnableSsl *BoolExpr `json:"EnableSsl,omitempty"`
	// Environment docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-environment
	Environment *OpsWorksAppEnvironmentVariableList `json:"Environment,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// Shortname docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-shortname
	Shortname *StringExpr `json:"Shortname,omitempty"`
	// SslConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-sslconfiguration
	SslConfiguration *OpsWorksAppSslConfiguration `json:"SslConfiguration,omitempty"`
	// StackID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-stackid
	StackID *StringExpr `json:"StackId,omitempty" validate:"dive,required"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-app.html#cfn-opsworks-app-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::OpsWorks::App to implement the ResourceProperties interface
func (s OpsWorksApp) CfnResourceType() string {

	return "AWS::OpsWorks::App"
}

// OpsWorksElasticLoadBalancerAttachment represents the AWS::OpsWorks::ElasticLoadBalancerAttachment CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-elbattachment.html
type OpsWorksElasticLoadBalancerAttachment struct {
	// ElasticLoadBalancerName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-elbattachment.html#cfn-opsworks-elbattachment-elbname
	ElasticLoadBalancerName *StringExpr `json:"ElasticLoadBalancerName,omitempty" validate:"dive,required"`
	// LayerID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-elbattachment.html#cfn-opsworks-elbattachment-layerid
	LayerID *StringExpr `json:"LayerId,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::OpsWorks::ElasticLoadBalancerAttachment to implement the ResourceProperties interface
func (s OpsWorksElasticLoadBalancerAttachment) CfnResourceType() string {

	return "AWS::OpsWorks::ElasticLoadBalancerAttachment"
}

// OpsWorksInstance represents the AWS::OpsWorks::Instance CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html
type OpsWorksInstance struct {
	// AgentVersion docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-agentversion
	AgentVersion *StringExpr `json:"AgentVersion,omitempty"`
	// AmiID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-amiid
	AmiID *StringExpr `json:"AmiId,omitempty"`
	// Architecture docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-architecture
	Architecture *StringExpr `json:"Architecture,omitempty"`
	// AutoScalingType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-autoscalingtype
	AutoScalingType *StringExpr `json:"AutoScalingType,omitempty"`
	// AvailabilityZone docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-availabilityzone
	AvailabilityZone *StringExpr `json:"AvailabilityZone,omitempty"`
	// BlockDeviceMappings docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-blockdevicemappings
	BlockDeviceMappings *OpsWorksInstanceBlockDeviceMappingList `json:"BlockDeviceMappings,omitempty"`
	// EbsOptimized docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-ebsoptimized
	EbsOptimized *BoolExpr `json:"EbsOptimized,omitempty"`
	// ElasticIPs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-elasticips
	ElasticIPs *StringListExpr `json:"ElasticIps,omitempty"`
	// Hostname docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-hostname
	Hostname *StringExpr `json:"Hostname,omitempty"`
	// InstallUpdatesOnBoot docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-installupdatesonboot
	InstallUpdatesOnBoot *BoolExpr `json:"InstallUpdatesOnBoot,omitempty"`
	// InstanceType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-instancetype
	InstanceType *StringExpr `json:"InstanceType,omitempty" validate:"dive,required"`
	// LayerIDs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-layerids
	LayerIDs *StringListExpr `json:"LayerIds,omitempty" validate:"dive,required"`
	// Os docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-os
	Os *StringExpr `json:"Os,omitempty"`
	// RootDeviceType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-rootdevicetype
	RootDeviceType *StringExpr `json:"RootDeviceType,omitempty"`
	// SSHKeyName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-sshkeyname
	SSHKeyName *StringExpr `json:"SshKeyName,omitempty"`
	// StackID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-stackid
	StackID *StringExpr `json:"StackId,omitempty" validate:"dive,required"`
	// SubnetID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-subnetid
	SubnetID *StringExpr `json:"SubnetId,omitempty"`
	// Tenancy docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-tenancy
	Tenancy *StringExpr `json:"Tenancy,omitempty"`
	// TimeBasedAutoScaling docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-timebasedautoscaling
	TimeBasedAutoScaling *OpsWorksInstanceTimeBasedAutoScaling `json:"TimeBasedAutoScaling,omitempty"`
	// VirtualizationType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-virtualizationtype
	VirtualizationType *StringExpr `json:"VirtualizationType,omitempty"`
	// Volumes docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-volumes
	Volumes *StringListExpr `json:"Volumes,omitempty"`
}

// CfnResourceType returns AWS::OpsWorks::Instance to implement the ResourceProperties interface
func (s OpsWorksInstance) CfnResourceType() string {

	return "AWS::OpsWorks::Instance"
}

// OpsWorksLayer represents the AWS::OpsWorks::Layer CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html
type OpsWorksLayer struct {
	// Attributes docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-attributes
	Attributes interface{} `json:"Attributes,omitempty"`
	// AutoAssignElasticIPs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-autoassignelasticips
	AutoAssignElasticIPs *BoolExpr `json:"AutoAssignElasticIps,omitempty" validate:"dive,required"`
	// AutoAssignPublicIPs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-autoassignpublicips
	AutoAssignPublicIPs *BoolExpr `json:"AutoAssignPublicIps,omitempty" validate:"dive,required"`
	// CustomInstanceProfileArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-custominstanceprofilearn
	CustomInstanceProfileArn *StringExpr `json:"CustomInstanceProfileArn,omitempty"`
	// CustomJSON docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-customjson
	CustomJSON interface{} `json:"CustomJson,omitempty"`
	// CustomRecipes docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-customrecipes
	CustomRecipes *OpsWorksLayerRecipes `json:"CustomRecipes,omitempty"`
	// CustomSecurityGroupIDs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-customsecuritygroupids
	CustomSecurityGroupIDs *StringListExpr `json:"CustomSecurityGroupIds,omitempty"`
	// EnableAutoHealing docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-enableautohealing
	EnableAutoHealing *BoolExpr `json:"EnableAutoHealing,omitempty" validate:"dive,required"`
	// InstallUpdatesOnBoot docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-installupdatesonboot
	InstallUpdatesOnBoot *BoolExpr `json:"InstallUpdatesOnBoot,omitempty"`
	// LifecycleEventConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-lifecycleeventconfiguration
	LifecycleEventConfiguration *OpsWorksLayerLifecycleEventConfiguration `json:"LifecycleEventConfiguration,omitempty"`
	// LoadBasedAutoScaling docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-loadbasedautoscaling
	LoadBasedAutoScaling *OpsWorksLayerLoadBasedAutoScaling `json:"LoadBasedAutoScaling,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// Packages docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-packages
	Packages *StringListExpr `json:"Packages,omitempty"`
	// Shortname docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-shortname
	Shortname *StringExpr `json:"Shortname,omitempty" validate:"dive,required"`
	// StackID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-stackid
	StackID *StringExpr `json:"StackId,omitempty" validate:"dive,required"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
	// UseEbsOptimizedInstances docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-useebsoptimizedinstances
	UseEbsOptimizedInstances *BoolExpr `json:"UseEbsOptimizedInstances,omitempty"`
	// VolumeConfigurations docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-volumeconfigurations
	VolumeConfigurations *OpsWorksLayerVolumeConfigurationList `json:"VolumeConfigurations,omitempty"`
}

// CfnResourceType returns AWS::OpsWorks::Layer to implement the ResourceProperties interface
func (s OpsWorksLayer) CfnResourceType() string {

	return "AWS::OpsWorks::Layer"
}

// OpsWorksStack represents the AWS::OpsWorks::Stack CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html
type OpsWorksStack struct {
	// AgentVersion docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-agentversion
	AgentVersion *StringExpr `json:"AgentVersion,omitempty"`
	// Attributes docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-attributes
	Attributes interface{} `json:"Attributes,omitempty"`
	// ChefConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-chefconfiguration
	ChefConfiguration *OpsWorksStackChefConfiguration `json:"ChefConfiguration,omitempty"`
	// CloneAppIDs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-cloneappids
	CloneAppIDs *StringListExpr `json:"CloneAppIds,omitempty"`
	// ClonePermissions docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-clonepermissions
	ClonePermissions *BoolExpr `json:"ClonePermissions,omitempty"`
	// ConfigurationManager docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-configmanager
	ConfigurationManager *OpsWorksStackStackConfigurationManager `json:"ConfigurationManager,omitempty"`
	// CustomCookbooksSource docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-custcookbooksource
	CustomCookbooksSource *OpsWorksStackSource `json:"CustomCookbooksSource,omitempty"`
	// CustomJSON docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-custjson
	CustomJSON interface{} `json:"CustomJson,omitempty"`
	// DefaultAvailabilityZone docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-defaultaz
	DefaultAvailabilityZone *StringExpr `json:"DefaultAvailabilityZone,omitempty"`
	// DefaultInstanceProfileArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-defaultinstanceprof
	DefaultInstanceProfileArn *StringExpr `json:"DefaultInstanceProfileArn,omitempty" validate:"dive,required"`
	// DefaultOs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-defaultos
	DefaultOs *StringExpr `json:"DefaultOs,omitempty"`
	// DefaultRootDeviceType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-defaultrootdevicetype
	DefaultRootDeviceType *StringExpr `json:"DefaultRootDeviceType,omitempty"`
	// DefaultSSHKeyName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-defaultsshkeyname
	DefaultSSHKeyName *StringExpr `json:"DefaultSshKeyName,omitempty"`
	// DefaultSubnetID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#defaultsubnet
	DefaultSubnetID *StringExpr `json:"DefaultSubnetId,omitempty"`
	// EcsClusterArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-ecsclusterarn
	EcsClusterArn *StringExpr `json:"EcsClusterArn,omitempty"`
	// ElasticIPs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-elasticips
	ElasticIPs *OpsWorksStackElasticIPList `json:"ElasticIps,omitempty"`
	// HostnameTheme docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-hostnametheme
	HostnameTheme *StringExpr `json:"HostnameTheme,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// RdsDbInstances docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-rdsdbinstances
	RdsDbInstances *OpsWorksStackRdsDbInstanceList `json:"RdsDbInstances,omitempty"`
	// ServiceRoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-servicerolearn
	ServiceRoleArn *StringExpr `json:"ServiceRoleArn,omitempty" validate:"dive,required"`
	// SourceStackID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-sourcestackid
	SourceStackID *StringExpr `json:"SourceStackId,omitempty"`
	// UseCustomCookbooks docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#usecustcookbooks
	UseCustomCookbooks *BoolExpr `json:"UseCustomCookbooks,omitempty"`
	// UseOpsworksSecurityGroups docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-useopsworkssecuritygroups
	UseOpsworksSecurityGroups *BoolExpr `json:"UseOpsworksSecurityGroups,omitempty"`
	// VPCID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html#cfn-opsworks-stack-vpcid
	VPCID *StringExpr `json:"VpcId,omitempty"`
}

// CfnResourceType returns AWS::OpsWorks::Stack to implement the ResourceProperties interface
func (s OpsWorksStack) CfnResourceType() string {

	return "AWS::OpsWorks::Stack"
}

// OpsWorksUserProfile represents the AWS::OpsWorks::UserProfile CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-userprofile.html
type OpsWorksUserProfile struct {
	// AllowSelfManagement docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-userprofile.html#cfn-opsworks-userprofile-allowselfmanagement
	AllowSelfManagement *BoolExpr `json:"AllowSelfManagement,omitempty"`
	// IamUserArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-userprofile.html#cfn-opsworks-userprofile-iamuserarn
	IamUserArn *StringExpr `json:"IamUserArn,omitempty" validate:"dive,required"`
	// SSHPublicKey docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-userprofile.html#cfn-opsworks-userprofile-sshpublickey
	SSHPublicKey *StringExpr `json:"SshPublicKey,omitempty"`
	// SSHUsername docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-userprofile.html#cfn-opsworks-userprofile-sshusername
	SSHUsername *StringExpr `json:"SshUsername,omitempty"`
}

// CfnResourceType returns AWS::OpsWorks::UserProfile to implement the ResourceProperties interface
func (s OpsWorksUserProfile) CfnResourceType() string {

	return "AWS::OpsWorks::UserProfile"
}

// OpsWorksVolume represents the AWS::OpsWorks::Volume CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-volume.html
type OpsWorksVolume struct {
	// Ec2VolumeID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-volume.html#cfn-opsworks-volume-ec2volumeid
	Ec2VolumeID *StringExpr `json:"Ec2VolumeId,omitempty" validate:"dive,required"`
	// MountPoint docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-volume.html#cfn-opsworks-volume-mountpoint
	MountPoint *StringExpr `json:"MountPoint,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-volume.html#cfn-opsworks-volume-name
	Name *StringExpr `json:"Name,omitempty"`
	// StackID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-volume.html#cfn-opsworks-volume-stackid
	StackID *StringExpr `json:"StackId,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::OpsWorks::Volume to implement the ResourceProperties interface
func (s OpsWorksVolume) CfnResourceType() string {

	return "AWS::OpsWorks::Volume"
}

// RDSDBCluster represents the AWS::RDS::DBCluster CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html
type RDSDBCluster struct {
	// AvailabilityZones docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-availabilityzones
	AvailabilityZones *StringExpr `json:"AvailabilityZones,omitempty"`
	// BackupRetentionPeriod docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-backuprententionperiod
	BackupRetentionPeriod *IntegerExpr `json:"BackupRetentionPeriod,omitempty"`
	// DBClusterParameterGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-dbclusterparametergroupname
	DBClusterParameterGroupName *StringExpr `json:"DBClusterParameterGroupName,omitempty"`
	// DBSubnetGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-dbsubnetgroupname
	DBSubnetGroupName *StringExpr `json:"DBSubnetGroupName,omitempty"`
	// DatabaseName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-databasename
	DatabaseName *StringExpr `json:"DatabaseName,omitempty"`
	// Engine docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-engine
	Engine *StringExpr `json:"Engine,omitempty" validate:"dive,required"`
	// EngineVersion docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-engineversion
	EngineVersion *StringExpr `json:"EngineVersion,omitempty"`
	// KmsKeyID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-kmskeyid
	KmsKeyID *StringExpr `json:"KmsKeyId,omitempty"`
	// MasterUserPassword docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-masteruserpassword
	MasterUserPassword *StringExpr `json:"MasterUserPassword,omitempty"`
	// MasterUsername docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-masterusername
	MasterUsername *StringExpr `json:"MasterUsername,omitempty"`
	// Port docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-port
	Port *IntegerExpr `json:"Port,omitempty"`
	// PreferredBackupWindow docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-preferredbackupwindow
	PreferredBackupWindow *StringExpr `json:"PreferredBackupWindow,omitempty"`
	// PreferredMaintenanceWindow docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-preferredmaintenancewindow
	PreferredMaintenanceWindow *StringExpr `json:"PreferredMaintenanceWindow,omitempty"`
	// SnapshotIDentifier docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-snapshotidentifier
	SnapshotIDentifier *StringExpr `json:"SnapshotIdentifier,omitempty"`
	// StorageEncrypted docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-storageencrypted
	StorageEncrypted *BoolExpr `json:"StorageEncrypted,omitempty"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-tags
	Tags *TagList `json:"Tags,omitempty"`
	// VPCSecurityGroupIDs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-vpcsecuritygroupids
	VPCSecurityGroupIDs *StringListExpr `json:"VpcSecurityGroupIds,omitempty"`
}

// CfnResourceType returns AWS::RDS::DBCluster to implement the ResourceProperties interface
func (s RDSDBCluster) CfnResourceType() string {

	return "AWS::RDS::DBCluster"
}

// RDSDBClusterParameterGroup represents the AWS::RDS::DBClusterParameterGroup CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterparametergroup.html
type RDSDBClusterParameterGroup struct {
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterparametergroup.html#cfn-rds-dbclusterparametergroup-description
	Description *StringExpr `json:"Description,omitempty" validate:"dive,required"`
	// Family docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterparametergroup.html#cfn-rds-dbclusterparametergroup-family
	Family *StringExpr `json:"Family,omitempty" validate:"dive,required"`
	// Parameters docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterparametergroup.html#cfn-rds-dbclusterparametergroup-parameters
	Parameters interface{} `json:"Parameters,omitempty" validate:"dive,required"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterparametergroup.html#cfn-rds-dbclusterparametergroup-tags
	Tags *TagList `json:"Tags,omitempty"`
}

// CfnResourceType returns AWS::RDS::DBClusterParameterGroup to implement the ResourceProperties interface
func (s RDSDBClusterParameterGroup) CfnResourceType() string {

	return "AWS::RDS::DBClusterParameterGroup"
}

// RDSDBInstance represents the AWS::RDS::DBInstance CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html
type RDSDBInstance struct {
	// AllocatedStorage docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-allocatedstorage
	AllocatedStorage *StringExpr `json:"AllocatedStorage,omitempty"`
	// AllowMajorVersionUpgrade docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-allowmajorversionupgrade
	AllowMajorVersionUpgrade *BoolExpr `json:"AllowMajorVersionUpgrade,omitempty"`
	// AutoMinorVersionUpgrade docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-autominorversionupgrade
	AutoMinorVersionUpgrade *BoolExpr `json:"AutoMinorVersionUpgrade,omitempty"`
	// AvailabilityZone docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-availabilityzone
	AvailabilityZone *StringExpr `json:"AvailabilityZone,omitempty"`
	// BackupRetentionPeriod docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-backupretentionperiod
	BackupRetentionPeriod *StringExpr `json:"BackupRetentionPeriod,omitempty"`
	// CharacterSetName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-charactersetname
	CharacterSetName *StringExpr `json:"CharacterSetName,omitempty"`
	// DBClusterIDentifier docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-dbclusteridentifier
	DBClusterIDentifier *StringExpr `json:"DBClusterIdentifier,omitempty"`
	// DBInstanceClass docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-dbinstanceclass
	DBInstanceClass *StringExpr `json:"DBInstanceClass,omitempty" validate:"dive,required"`
	// DBInstanceIDentifier docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-dbinstanceidentifier
	DBInstanceIDentifier *StringExpr `json:"DBInstanceIdentifier,omitempty"`
	// DBName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-dbname
	DBName *StringExpr `json:"DBName,omitempty"`
	// DBParameterGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-dbparametergroupname
	DBParameterGroupName *StringExpr `json:"DBParameterGroupName,omitempty"`
	// DBSecurityGroups docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-dbsecuritygroups
	DBSecurityGroups *StringListExpr `json:"DBSecurityGroups,omitempty"`
	// DBSnapshotIDentifier docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-dbsnapshotidentifier
	DBSnapshotIDentifier *StringExpr `json:"DBSnapshotIdentifier,omitempty"`
	// DBSubnetGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-dbsubnetgroupname
	DBSubnetGroupName *StringExpr `json:"DBSubnetGroupName,omitempty"`
	// Domain docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-domain
	Domain *StringExpr `json:"Domain,omitempty"`
	// DomainIAMRoleName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-domainiamrolename
	DomainIAMRoleName *StringExpr `json:"DomainIAMRoleName,omitempty"`
	// Engine docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-engine
	Engine *StringExpr `json:"Engine,omitempty"`
	// EngineVersion docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-engineversion
	EngineVersion *StringExpr `json:"EngineVersion,omitempty"`
	// Iops docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-iops
	Iops *IntegerExpr `json:"Iops,omitempty"`
	// KmsKeyID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-kmskeyid
	KmsKeyID *StringExpr `json:"KmsKeyId,omitempty"`
	// LicenseModel docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-licensemodel
	LicenseModel *StringExpr `json:"LicenseModel,omitempty"`
	// MasterUserPassword docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-masteruserpassword
	MasterUserPassword *StringExpr `json:"MasterUserPassword,omitempty"`
	// MasterUsername docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-masterusername
	MasterUsername *StringExpr `json:"MasterUsername,omitempty"`
	// MonitoringInterval docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-monitoringinterval
	MonitoringInterval *IntegerExpr `json:"MonitoringInterval,omitempty"`
	// MonitoringRoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-monitoringrolearn
	MonitoringRoleArn *StringExpr `json:"MonitoringRoleArn,omitempty"`
	// MultiAZ docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-multiaz
	MultiAZ *BoolExpr `json:"MultiAZ,omitempty"`
	// OptionGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-optiongroupname
	OptionGroupName *StringExpr `json:"OptionGroupName,omitempty"`
	// Port docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-port
	Port *StringExpr `json:"Port,omitempty"`
	// PreferredBackupWindow docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-preferredbackupwindow
	PreferredBackupWindow *StringExpr `json:"PreferredBackupWindow,omitempty"`
	// PreferredMaintenanceWindow docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-preferredmaintenancewindow
	PreferredMaintenanceWindow *StringExpr `json:"PreferredMaintenanceWindow,omitempty"`
	// PubliclyAccessible docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-publiclyaccessible
	PubliclyAccessible *BoolExpr `json:"PubliclyAccessible,omitempty"`
	// SourceDBInstanceIDentifier docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-sourcedbinstanceidentifier
	SourceDBInstanceIDentifier *StringExpr `json:"SourceDBInstanceIdentifier,omitempty"`
	// StorageEncrypted docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-storageencrypted
	StorageEncrypted *BoolExpr `json:"StorageEncrypted,omitempty"`
	// StorageType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-storagetype
	StorageType *StringExpr `json:"StorageType,omitempty"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-tags
	Tags *TagList `json:"Tags,omitempty"`
	// Timezone docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-timezone
	Timezone *StringExpr `json:"Timezone,omitempty"`
	// VPCSecurityGroups docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-vpcsecuritygroups
	VPCSecurityGroups *StringListExpr `json:"VPCSecurityGroups,omitempty"`
}

// CfnResourceType returns AWS::RDS::DBInstance to implement the ResourceProperties interface
func (s RDSDBInstance) CfnResourceType() string {

	return "AWS::RDS::DBInstance"
}

// RDSDBParameterGroup represents the AWS::RDS::DBParameterGroup CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbparametergroup.html
type RDSDBParameterGroup struct {
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbparametergroup.html#cfn-rds-dbparametergroup-description
	Description *StringExpr `json:"Description,omitempty" validate:"dive,required"`
	// Family docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbparametergroup.html#cfn-rds-dbparametergroup-family
	Family *StringExpr `json:"Family,omitempty" validate:"dive,required"`
	// Parameters docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbparametergroup.html#cfn-rds-dbparametergroup-parameters
	Parameters interface{} `json:"Parameters,omitempty"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbparametergroup.html#cfn-rds-dbparametergroup-tags
	Tags *TagList `json:"Tags,omitempty"`
}

// CfnResourceType returns AWS::RDS::DBParameterGroup to implement the ResourceProperties interface
func (s RDSDBParameterGroup) CfnResourceType() string {

	return "AWS::RDS::DBParameterGroup"
}

// RDSDBSecurityGroup represents the AWS::RDS::DBSecurityGroup CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group.html
type RDSDBSecurityGroup struct {
	// DBSecurityGroupIngress docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group.html#cfn-rds-dbsecuritygroup-dbsecuritygroupingress
	DBSecurityGroupIngress RDSDBSecurityGroupIngressPropertyList `json:"DBSecurityGroupIngress,omitempty" validate:"dive,required"`
	// EC2VPCID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group.html#cfn-rds-dbsecuritygroup-ec2vpcid
	EC2VPCID *StringExpr `json:"EC2VpcId,omitempty"`
	// GroupDescription docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group.html#cfn-rds-dbsecuritygroup-groupdescription
	GroupDescription *StringExpr `json:"GroupDescription,omitempty" validate:"dive,required"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group.html#cfn-rds-dbsecuritygroup-tags
	Tags *TagList `json:"Tags,omitempty"`
}

// CfnResourceType returns AWS::RDS::DBSecurityGroup to implement the ResourceProperties interface
func (s RDSDBSecurityGroup) CfnResourceType() string {

	return "AWS::RDS::DBSecurityGroup"
}

// RDSDBSecurityGroupIngress represents the AWS::RDS::DBSecurityGroupIngress CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-security-group-ingress.html
type RDSDBSecurityGroupIngress struct {
	// CIDRIP docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-security-group-ingress.html#cfn-rds-securitygroup-ingress-cidrip
	CIDRIP *StringExpr `json:"CIDRIP,omitempty"`
	// DBSecurityGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-security-group-ingress.html#cfn-rds-securitygroup-ingress-dbsecuritygroupname
	DBSecurityGroupName *StringExpr `json:"DBSecurityGroupName,omitempty" validate:"dive,required"`
	// EC2SecurityGroupID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-security-group-ingress.html#cfn-rds-securitygroup-ingress-ec2securitygroupid
	EC2SecurityGroupID *StringExpr `json:"EC2SecurityGroupId,omitempty"`
	// EC2SecurityGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-security-group-ingress.html#cfn-rds-securitygroup-ingress-ec2securitygroupname
	EC2SecurityGroupName *StringExpr `json:"EC2SecurityGroupName,omitempty"`
	// EC2SecurityGroupOwnerID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-security-group-ingress.html#cfn-rds-securitygroup-ingress-ec2securitygroupownerid
	EC2SecurityGroupOwnerID *StringExpr `json:"EC2SecurityGroupOwnerId,omitempty"`
}

// CfnResourceType returns AWS::RDS::DBSecurityGroupIngress to implement the ResourceProperties interface
func (s RDSDBSecurityGroupIngress) CfnResourceType() string {

	return "AWS::RDS::DBSecurityGroupIngress"
}

// RDSDBSubnetGroup represents the AWS::RDS::DBSubnetGroup CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsubnet-group.html
type RDSDBSubnetGroup struct {
	// DBSubnetGroupDescription docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsubnet-group.html#cfn-rds-dbsubnetgroup-dbsubnetgroupdescription
	DBSubnetGroupDescription *StringExpr `json:"DBSubnetGroupDescription,omitempty" validate:"dive,required"`
	// SubnetIDs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsubnet-group.html#cfn-rds-dbsubnetgroup-subnetids
	SubnetIDs *StringListExpr `json:"SubnetIds,omitempty" validate:"dive,required"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsubnet-group.html#cfn-rds-dbsubnetgroup-tags
	Tags *TagList `json:"Tags,omitempty"`
}

// CfnResourceType returns AWS::RDS::DBSubnetGroup to implement the ResourceProperties interface
func (s RDSDBSubnetGroup) CfnResourceType() string {

	return "AWS::RDS::DBSubnetGroup"
}

// RDSEventSubscription represents the AWS::RDS::EventSubscription CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html
type RDSEventSubscription struct {
	// Enabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html#cfn-rds-eventsubscription-enabled
	Enabled *BoolExpr `json:"Enabled,omitempty"`
	// EventCategories docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html#cfn-rds-eventsubscription-eventcategories
	EventCategories *StringListExpr `json:"EventCategories,omitempty"`
	// SnsTopicArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html#cfn-rds-eventsubscription-snstopicarn
	SnsTopicArn *StringExpr `json:"SnsTopicArn,omitempty" validate:"dive,required"`
	// SourceIDs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html#cfn-rds-eventsubscription-sourceids
	SourceIDs *StringListExpr `json:"SourceIds,omitempty"`
	// SourceType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html#cfn-rds-eventsubscription-sourcetype
	SourceType *StringExpr `json:"SourceType,omitempty"`
}

// CfnResourceType returns AWS::RDS::EventSubscription to implement the ResourceProperties interface
func (s RDSEventSubscription) CfnResourceType() string {

	return "AWS::RDS::EventSubscription"
}

// RDSOptionGroup represents the AWS::RDS::OptionGroup CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html
type RDSOptionGroup struct {
	// EngineName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html#cfn-rds-optiongroup-enginename
	EngineName *StringExpr `json:"EngineName,omitempty" validate:"dive,required"`
	// MajorEngineVersion docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html#cfn-rds-optiongroup-majorengineversion
	MajorEngineVersion *StringExpr `json:"MajorEngineVersion,omitempty" validate:"dive,required"`
	// OptionConfigurations docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html#cfn-rds-optiongroup-optionconfigurations
	OptionConfigurations *RDSOptionGroupOptionConfigurationList `json:"OptionConfigurations,omitempty" validate:"dive,required"`
	// OptionGroupDescription docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html#cfn-rds-optiongroup-optiongroupdescription
	OptionGroupDescription *StringExpr `json:"OptionGroupDescription,omitempty" validate:"dive,required"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html#cfn-rds-optiongroup-tags
	Tags *TagList `json:"Tags,omitempty"`
}

// CfnResourceType returns AWS::RDS::OptionGroup to implement the ResourceProperties interface
func (s RDSOptionGroup) CfnResourceType() string {

	return "AWS::RDS::OptionGroup"
}

// RedshiftCluster represents the AWS::Redshift::Cluster CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html
type RedshiftCluster struct {
	// AllowVersionUpgrade docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-allowversionupgrade
	AllowVersionUpgrade *BoolExpr `json:"AllowVersionUpgrade,omitempty"`
	// AutomatedSnapshotRetentionPeriod docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-automatedsnapshotretentionperiod
	AutomatedSnapshotRetentionPeriod *IntegerExpr `json:"AutomatedSnapshotRetentionPeriod,omitempty"`
	// AvailabilityZone docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-availabilityzone
	AvailabilityZone *StringExpr `json:"AvailabilityZone,omitempty"`
	// ClusterParameterGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-clusterparametergroupname
	ClusterParameterGroupName *StringExpr `json:"ClusterParameterGroupName,omitempty"`
	// ClusterSecurityGroups docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-clustersecuritygroups
	ClusterSecurityGroups *StringListExpr `json:"ClusterSecurityGroups,omitempty"`
	// ClusterSubnetGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-clustersubnetgroupname
	ClusterSubnetGroupName *StringExpr `json:"ClusterSubnetGroupName,omitempty"`
	// ClusterType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-clustertype
	ClusterType *StringExpr `json:"ClusterType,omitempty" validate:"dive,required"`
	// ClusterVersion docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-clusterversion
	ClusterVersion *StringExpr `json:"ClusterVersion,omitempty"`
	// DBName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-dbname
	DBName *StringExpr `json:"DBName,omitempty" validate:"dive,required"`
	// ElasticIP docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-elasticip
	ElasticIP *StringExpr `json:"ElasticIp,omitempty"`
	// Encrypted docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-encrypted
	Encrypted *BoolExpr `json:"Encrypted,omitempty"`
	// HsmClientCertificateIDentifier docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-hsmclientcertidentifier
	HsmClientCertificateIDentifier *StringExpr `json:"HsmClientCertificateIdentifier,omitempty"`
	// HsmConfigurationIDentifier docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-HsmConfigurationIdentifier
	HsmConfigurationIDentifier *StringExpr `json:"HsmConfigurationIdentifier,omitempty"`
	// IamRoles docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-iamroles
	IamRoles *StringListExpr `json:"IamRoles,omitempty"`
	// KmsKeyID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-kmskeyid
	KmsKeyID *StringExpr `json:"KmsKeyId,omitempty"`
	// MasterUserPassword docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-masteruserpassword
	MasterUserPassword *StringExpr `json:"MasterUserPassword,omitempty" validate:"dive,required"`
	// MasterUsername docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-masterusername
	MasterUsername *StringExpr `json:"MasterUsername,omitempty" validate:"dive,required"`
	// NodeType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-nodetype
	NodeType *StringExpr `json:"NodeType,omitempty" validate:"dive,required"`
	// NumberOfNodes docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-nodetype
	NumberOfNodes *IntegerExpr `json:"NumberOfNodes,omitempty"`
	// OwnerAccount docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-owneraccount
	OwnerAccount *StringExpr `json:"OwnerAccount,omitempty"`
	// Port docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-port
	Port *IntegerExpr `json:"Port,omitempty"`
	// PreferredMaintenanceWindow docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-preferredmaintenancewindow
	PreferredMaintenanceWindow *StringExpr `json:"PreferredMaintenanceWindow,omitempty"`
	// PubliclyAccessible docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-publiclyaccessible
	PubliclyAccessible *BoolExpr `json:"PubliclyAccessible,omitempty"`
	// SnapshotClusterIDentifier docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-snapshotclusteridentifier
	SnapshotClusterIDentifier *StringExpr `json:"SnapshotClusterIdentifier,omitempty"`
	// SnapshotIDentifier docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-snapshotidentifier
	SnapshotIDentifier *StringExpr `json:"SnapshotIdentifier,omitempty"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-tags
	Tags *TagList `json:"Tags,omitempty"`
	// VPCSecurityGroupIDs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-vpcsecuritygroupids
	VPCSecurityGroupIDs *StringListExpr `json:"VpcSecurityGroupIds,omitempty"`
}

// CfnResourceType returns AWS::Redshift::Cluster to implement the ResourceProperties interface
func (s RedshiftCluster) CfnResourceType() string {

	return "AWS::Redshift::Cluster"
}

// RedshiftClusterParameterGroup represents the AWS::Redshift::ClusterParameterGroup CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clusterparametergroup.html
type RedshiftClusterParameterGroup struct {
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clusterparametergroup.html#cfn-redshift-clusterparametergroup-description
	Description *StringExpr `json:"Description,omitempty" validate:"dive,required"`
	// ParameterGroupFamily docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clusterparametergroup.html#cfn-redshift-clusterparametergroup-parametergroupfamily
	ParameterGroupFamily *StringExpr `json:"ParameterGroupFamily,omitempty" validate:"dive,required"`
	// Parameters docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clusterparametergroup.html#cfn-redshift-clusterparametergroup-parameters
	Parameters *RedshiftClusterParameterGroupParameterList `json:"Parameters,omitempty"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clusterparametergroup.html#cfn-redshift-clusterparametergroup-tags
	Tags *TagList `json:"Tags,omitempty"`
}

// CfnResourceType returns AWS::Redshift::ClusterParameterGroup to implement the ResourceProperties interface
func (s RedshiftClusterParameterGroup) CfnResourceType() string {

	return "AWS::Redshift::ClusterParameterGroup"
}

// RedshiftClusterSecurityGroup represents the AWS::Redshift::ClusterSecurityGroup CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroup.html
type RedshiftClusterSecurityGroup struct {
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroup.html#cfn-redshift-clustersecuritygroup-description
	Description *StringExpr `json:"Description,omitempty" validate:"dive,required"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroup.html#cfn-redshift-clustersecuritygroup-tags
	Tags *TagList `json:"Tags,omitempty"`
}

// CfnResourceType returns AWS::Redshift::ClusterSecurityGroup to implement the ResourceProperties interface
func (s RedshiftClusterSecurityGroup) CfnResourceType() string {

	return "AWS::Redshift::ClusterSecurityGroup"
}

// RedshiftClusterSecurityGroupIngress represents the AWS::Redshift::ClusterSecurityGroupIngress CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroupingress.html
type RedshiftClusterSecurityGroupIngress struct {
	// CIDRIP docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroupingress.html#cfn-redshift-clustersecuritygroupingress-cidrip
	CIDRIP *StringExpr `json:"CIDRIP,omitempty"`
	// ClusterSecurityGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroupingress.html#cfn-redshift-clustersecuritygroupingress-clustersecuritygroupname
	ClusterSecurityGroupName *StringExpr `json:"ClusterSecurityGroupName,omitempty" validate:"dive,required"`
	// EC2SecurityGroupName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroupingress.html#cfn-redshift-clustersecuritygroupingress-ec2securitygroupname
	EC2SecurityGroupName *StringExpr `json:"EC2SecurityGroupName,omitempty"`
	// EC2SecurityGroupOwnerID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroupingress.html#cfn-redshift-clustersecuritygroupingress-ec2securitygroupownerid
	EC2SecurityGroupOwnerID *StringExpr `json:"EC2SecurityGroupOwnerId,omitempty"`
}

// CfnResourceType returns AWS::Redshift::ClusterSecurityGroupIngress to implement the ResourceProperties interface
func (s RedshiftClusterSecurityGroupIngress) CfnResourceType() string {

	return "AWS::Redshift::ClusterSecurityGroupIngress"
}

// RedshiftClusterSubnetGroup represents the AWS::Redshift::ClusterSubnetGroup CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersubnetgroup.html
type RedshiftClusterSubnetGroup struct {
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersubnetgroup.html#cfn-redshift-clustersubnetgroup-description
	Description *StringExpr `json:"Description,omitempty" validate:"dive,required"`
	// SubnetIDs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersubnetgroup.html#cfn-redshift-clustersubnetgroup-subnetids
	SubnetIDs *StringListExpr `json:"SubnetIds,omitempty" validate:"dive,required"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersubnetgroup.html#cfn-redshift-clustersubnetgroup-tags
	Tags *TagList `json:"Tags,omitempty"`
}

// CfnResourceType returns AWS::Redshift::ClusterSubnetGroup to implement the ResourceProperties interface
func (s RedshiftClusterSubnetGroup) CfnResourceType() string {

	return "AWS::Redshift::ClusterSubnetGroup"
}

// Route53HealthCheck represents the AWS::Route53::HealthCheck CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html
type Route53HealthCheck struct {
	// HealthCheckConfig docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html#cfn-route53-healthcheck-healthcheckconfig
	HealthCheckConfig *Route53HealthCheckHealthCheckConfig `json:"HealthCheckConfig,omitempty" validate:"dive,required"`
	// HealthCheckTags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html#cfn-route53-healthcheck-healthchecktags
	HealthCheckTags *Route53HealthCheckHealthCheckTagList `json:"HealthCheckTags,omitempty"`
}

// CfnResourceType returns AWS::Route53::HealthCheck to implement the ResourceProperties interface
func (s Route53HealthCheck) CfnResourceType() string {

	return "AWS::Route53::HealthCheck"
}

// Route53HostedZone represents the AWS::Route53::HostedZone CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html
type Route53HostedZone struct {
	// HostedZoneConfig docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-hostedzoneconfig
	HostedZoneConfig *Route53HostedZoneHostedZoneConfig `json:"HostedZoneConfig,omitempty"`
	// HostedZoneTags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-hostedzonetags
	HostedZoneTags *Route53HostedZoneHostedZoneTagList `json:"HostedZoneTags,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// VPCs docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-vpcs
	VPCs *Route53HostedZoneVPCList `json:"VPCs,omitempty"`
}

// CfnResourceType returns AWS::Route53::HostedZone to implement the ResourceProperties interface
func (s Route53HostedZone) CfnResourceType() string {

	return "AWS::Route53::HostedZone"
}

// Route53RecordSet represents the AWS::Route53::RecordSet CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html
type Route53RecordSet struct {
	// AliasTarget docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-aliastarget
	AliasTarget *Route53RecordSetAliasTarget `json:"AliasTarget,omitempty"`
	// Comment docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-comment
	Comment *StringExpr `json:"Comment,omitempty"`
	// Failover docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-failover
	Failover *StringExpr `json:"Failover,omitempty"`
	// GeoLocation docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-geolocation
	GeoLocation *Route53RecordSetGeoLocation `json:"GeoLocation,omitempty"`
	// HealthCheckID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-healthcheckid
	HealthCheckID *StringExpr `json:"HealthCheckId,omitempty"`
	// HostedZoneID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-hostedzoneid
	HostedZoneID *StringExpr `json:"HostedZoneId,omitempty"`
	// HostedZoneName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-hostedzonename
	HostedZoneName *StringExpr `json:"HostedZoneName,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// Region docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-region
	Region *StringExpr `json:"Region,omitempty"`
	// ResourceRecords docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-resourcerecords
	ResourceRecords *StringListExpr `json:"ResourceRecords,omitempty"`
	// SetIDentifier docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-setidentifier
	SetIDentifier *StringExpr `json:"SetIdentifier,omitempty"`
	// TTL docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-ttl
	TTL *StringExpr `json:"TTL,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
	// Weight docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset.html#cfn-route53-recordset-weight
	Weight *IntegerExpr `json:"Weight,omitempty"`
}

// CfnResourceType returns AWS::Route53::RecordSet to implement the ResourceProperties interface
func (s Route53RecordSet) CfnResourceType() string {

	return "AWS::Route53::RecordSet"
}

// Route53RecordSetGroup represents the AWS::Route53::RecordSetGroup CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html
type Route53RecordSetGroup struct {
	// Comment docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html#cfn-route53-recordsetgroup-comment
	Comment *StringExpr `json:"Comment,omitempty"`
	// HostedZoneID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html#cfn-route53-recordsetgroup-hostedzoneid
	HostedZoneID *StringExpr `json:"HostedZoneId,omitempty"`
	// HostedZoneName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html#cfn-route53-recordsetgroup-hostedzonename
	HostedZoneName *StringExpr `json:"HostedZoneName,omitempty"`
	// RecordSets docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html#cfn-route53-recordsetgroup-recordsets
	RecordSets *Route53RecordSetGroupRecordSetList `json:"RecordSets,omitempty"`
}

// CfnResourceType returns AWS::Route53::RecordSetGroup to implement the ResourceProperties interface
func (s Route53RecordSetGroup) CfnResourceType() string {

	return "AWS::Route53::RecordSetGroup"
}

// S3Bucket represents the AWS::S3::Bucket CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html
type S3Bucket struct {
	// AccessControl docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-accesscontrol
	AccessControl *StringExpr `json:"AccessControl,omitempty"`
	// BucketName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-name
	BucketName *StringExpr `json:"BucketName,omitempty"`
	// CorsConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-crossoriginconfig
	CorsConfiguration *S3BucketCorsConfiguration `json:"CorsConfiguration,omitempty"`
	// LifecycleConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-lifecycleconfig
	LifecycleConfiguration *S3BucketLifecycleConfiguration `json:"LifecycleConfiguration,omitempty"`
	// LoggingConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-loggingconfig
	LoggingConfiguration *S3BucketLoggingConfiguration `json:"LoggingConfiguration,omitempty"`
	// NotificationConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-notification
	NotificationConfiguration *S3BucketNotificationConfiguration `json:"NotificationConfiguration,omitempty"`
	// ReplicationConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-replicationconfiguration
	ReplicationConfiguration *S3BucketReplicationConfiguration `json:"ReplicationConfiguration,omitempty"`
	// Tags docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-tags
	Tags *TagList `json:"Tags,omitempty"`
	// VersioningConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-versioning
	VersioningConfiguration *S3BucketVersioningConfiguration `json:"VersioningConfiguration,omitempty"`
	// WebsiteConfiguration docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-websiteconfiguration
	WebsiteConfiguration *S3BucketWebsiteConfiguration `json:"WebsiteConfiguration,omitempty"`
}

// CfnResourceType returns AWS::S3::Bucket to implement the ResourceProperties interface
func (s S3Bucket) CfnResourceType() string {

	return "AWS::S3::Bucket"
}

// S3BucketPolicy represents the AWS::S3::BucketPolicy CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-policy.html
type S3BucketPolicy struct {
	// Bucket docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-policy.html#cfn-s3-bucketpolicy-bucket
	Bucket *StringExpr `json:"Bucket,omitempty" validate:"dive,required"`
	// PolicyDocument docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-policy.html#cfn-s3-bucketpolicy-policydocument
	PolicyDocument interface{} `json:"PolicyDocument,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::S3::BucketPolicy to implement the ResourceProperties interface
func (s S3BucketPolicy) CfnResourceType() string {

	return "AWS::S3::BucketPolicy"
}

// SDBDomain represents the AWS::SDB::Domain CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-simpledb.html
type SDBDomain struct {
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-simpledb.html#cfn-sdb-domain-description
	Description *StringExpr `json:"Description,omitempty"`
}

// CfnResourceType returns AWS::SDB::Domain to implement the ResourceProperties interface
func (s SDBDomain) CfnResourceType() string {

	return "AWS::SDB::Domain"
}

// SNSSubscription represents the AWS::SNS::Subscription CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html
type SNSSubscription struct {
	// Endpoint docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-endpoint
	Endpoint *StringExpr `json:"Endpoint,omitempty"`
	// Protocol docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-protocol
	Protocol *StringExpr `json:"Protocol,omitempty"`
	// TopicArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#topicarn
	TopicArn *StringExpr `json:"TopicArn,omitempty"`
}

// CfnResourceType returns AWS::SNS::Subscription to implement the ResourceProperties interface
func (s SNSSubscription) CfnResourceType() string {

	return "AWS::SNS::Subscription"
}

// SNSTopic represents the AWS::SNS::Topic CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html
type SNSTopic struct {
	// DisplayName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html#cfn-sns-topic-displayname
	DisplayName *StringExpr `json:"DisplayName,omitempty"`
	// Subscription docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html#cfn-sns-topic-subscription
	Subscription *SNSTopicSubscriptionList `json:"Subscription,omitempty"`
	// TopicName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html#cfn-sns-topic-topicname
	TopicName *StringExpr `json:"TopicName,omitempty"`
}

// CfnResourceType returns AWS::SNS::Topic to implement the ResourceProperties interface
func (s SNSTopic) CfnResourceType() string {

	return "AWS::SNS::Topic"
}

// SNSTopicPolicy represents the AWS::SNS::TopicPolicy CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html
type SNSTopicPolicy struct {
	// PolicyDocument docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html#cfn-sns-topicpolicy-policydocument
	PolicyDocument interface{} `json:"PolicyDocument,omitempty" validate:"dive,required"`
	// Topics docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html#cfn-sns-topicpolicy-topics
	Topics *StringListExpr `json:"Topics,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::SNS::TopicPolicy to implement the ResourceProperties interface
func (s SNSTopicPolicy) CfnResourceType() string {

	return "AWS::SNS::TopicPolicy"
}

// SQSQueue represents the AWS::SQS::Queue CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html
type SQSQueue struct {
	// ContentBasedDeduplication docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html#aws-sqs-queue-contentbaseddeduplication
	ContentBasedDeduplication *BoolExpr `json:"ContentBasedDeduplication,omitempty"`
	// DelaySeconds docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html#aws-sqs-queue-delayseconds
	DelaySeconds *IntegerExpr `json:"DelaySeconds,omitempty"`
	// FifoQueue docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html#aws-sqs-queue-fifoqueue
	FifoQueue *BoolExpr `json:"FifoQueue,omitempty"`
	// MaximumMessageSize docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html#aws-sqs-queue-maxmesgsize
	MaximumMessageSize *IntegerExpr `json:"MaximumMessageSize,omitempty"`
	// MessageRetentionPeriod docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html#aws-sqs-queue-msgretentionperiod
	MessageRetentionPeriod *IntegerExpr `json:"MessageRetentionPeriod,omitempty"`
	// QueueName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html#aws-sqs-queue-name
	QueueName *StringExpr `json:"QueueName,omitempty"`
	// ReceiveMessageWaitTimeSeconds docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html#aws-sqs-queue-receivemsgwaittime
	ReceiveMessageWaitTimeSeconds *IntegerExpr `json:"ReceiveMessageWaitTimeSeconds,omitempty"`
	// RedrivePolicy docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html#aws-sqs-queue-redrive
	RedrivePolicy interface{} `json:"RedrivePolicy,omitempty"`
	// VisibilityTimeout docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html#aws-sqs-queue-visiblitytimeout
	VisibilityTimeout *IntegerExpr `json:"VisibilityTimeout,omitempty"`
}

// CfnResourceType returns AWS::SQS::Queue to implement the ResourceProperties interface
func (s SQSQueue) CfnResourceType() string {

	return "AWS::SQS::Queue"
}

// SQSQueuePolicy represents the AWS::SQS::QueuePolicy CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-policy.html
type SQSQueuePolicy struct {
	// PolicyDocument docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-policy.html#cfn-sqs-queuepolicy-policydoc
	PolicyDocument interface{} `json:"PolicyDocument,omitempty" validate:"dive,required"`
	// Queues docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-policy.html#cfn-sqs-queuepolicy-queues
	Queues *StringListExpr `json:"Queues,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::SQS::QueuePolicy to implement the ResourceProperties interface
func (s SQSQueuePolicy) CfnResourceType() string {

	return "AWS::SQS::QueuePolicy"
}

// SSMAssociation represents the AWS::SSM::Association CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html
type SSMAssociation struct {
	// DocumentVersion docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-documentversion
	DocumentVersion *StringExpr `json:"DocumentVersion,omitempty"`
	// InstanceID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-instanceid
	InstanceID *StringExpr `json:"InstanceId,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// Parameters docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-parameters
	Parameters interface{} `json:"Parameters,omitempty"`
	// ScheduleExpression docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-scheduleexpression
	ScheduleExpression *StringExpr `json:"ScheduleExpression,omitempty"`
	// Targets docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-targets
	Targets *SSMAssociationTargetList `json:"Targets,omitempty"`
}

// CfnResourceType returns AWS::SSM::Association to implement the ResourceProperties interface
func (s SSMAssociation) CfnResourceType() string {

	return "AWS::SSM::Association"
}

// SSMDocument represents the AWS::SSM::Document CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-document.html
type SSMDocument struct {
	// Content docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-document.html#cfn-ssm-document-content
	Content interface{} `json:"Content,omitempty" validate:"dive,required"`
	// DocumentType docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-document.html#cfn-ssm-document-documenttype
	DocumentType *StringExpr `json:"DocumentType,omitempty"`
}

// CfnResourceType returns AWS::SSM::Document to implement the ResourceProperties interface
func (s SSMDocument) CfnResourceType() string {

	return "AWS::SSM::Document"
}

// SSMParameter represents the AWS::SSM::Parameter CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-parameter.html
type SSMParameter struct {
	// Description docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-parameter.html#cfn-ssm-parameter-description
	Description *StringExpr `json:"Description,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-parameter.html#cfn-ssm-parameter-name
	Name *StringExpr `json:"Name,omitempty"`
	// Type docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-parameter.html#cfn-ssm-parameter-type
	Type *StringExpr `json:"Type,omitempty" validate:"dive,required"`
	// Value docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-parameter.html#cfn-ssm-parameter-value
	Value *StringExpr `json:"Value,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::SSM::Parameter to implement the ResourceProperties interface
func (s SSMParameter) CfnResourceType() string {

	return "AWS::SSM::Parameter"
}

// StepFunctionsActivity represents the AWS::StepFunctions::Activity CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-activity.html
type StepFunctionsActivity struct {
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-activity.html#cfn-stepfunctions-activity-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::StepFunctions::Activity to implement the ResourceProperties interface
func (s StepFunctionsActivity) CfnResourceType() string {

	return "AWS::StepFunctions::Activity"
}

// StepFunctionsStateMachine represents the AWS::StepFunctions::StateMachine CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html
type StepFunctionsStateMachine struct {
	// DefinitionString docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definitionstring
	DefinitionString *StringExpr `json:"DefinitionString,omitempty" validate:"dive,required"`
	// RoleArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-rolearn
	RoleArn *StringExpr `json:"RoleArn,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::StepFunctions::StateMachine to implement the ResourceProperties interface
func (s StepFunctionsStateMachine) CfnResourceType() string {

	return "AWS::StepFunctions::StateMachine"
}

// WAFByteMatchSet represents the AWS::WAF::ByteMatchSet CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-bytematchset.html
type WAFByteMatchSet struct {
	// ByteMatchTuples docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-bytematchset.html#cfn-waf-bytematchset-bytematchtuples
	ByteMatchTuples *WAFByteMatchSetByteMatchTupleList `json:"ByteMatchTuples,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-bytematchset.html#cfn-waf-bytematchset-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::WAF::ByteMatchSet to implement the ResourceProperties interface
func (s WAFByteMatchSet) CfnResourceType() string {

	return "AWS::WAF::ByteMatchSet"
}

// WAFIPSet represents the AWS::WAF::IPSet CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-ipset.html
type WAFIPSet struct {
	// IPSetDescriptors docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-ipset.html#cfn-waf-ipset-ipsetdescriptors
	IPSetDescriptors *WAFIPSetIPSetDescriptorList `json:"IPSetDescriptors,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-ipset.html#cfn-waf-ipset-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::WAF::IPSet to implement the ResourceProperties interface
func (s WAFIPSet) CfnResourceType() string {

	return "AWS::WAF::IPSet"
}

// WAFRule represents the AWS::WAF::Rule CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-rule.html
type WAFRule struct {
	// MetricName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-rule.html#cfn-waf-rule-metricname
	MetricName *StringExpr `json:"MetricName,omitempty" validate:"dive,required"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-rule.html#cfn-waf-rule-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// Predicates docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-rule.html#cfn-waf-rule-predicates
	Predicates *WAFRulePredicateList `json:"Predicates,omitempty"`
}

// CfnResourceType returns AWS::WAF::Rule to implement the ResourceProperties interface
func (s WAFRule) CfnResourceType() string {

	return "AWS::WAF::Rule"
}

// WAFSizeConstraintSet represents the AWS::WAF::SizeConstraintSet CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sizeconstraintset.html
type WAFSizeConstraintSet struct {
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sizeconstraintset.html#cfn-waf-sizeconstraintset-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// SizeConstraints docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sizeconstraintset.html#cfn-waf-sizeconstraintset-sizeconstraints
	SizeConstraints *WAFSizeConstraintSetSizeConstraintList `json:"SizeConstraints,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::WAF::SizeConstraintSet to implement the ResourceProperties interface
func (s WAFSizeConstraintSet) CfnResourceType() string {

	return "AWS::WAF::SizeConstraintSet"
}

// WAFSQLInjectionMatchSet represents the AWS::WAF::SqlInjectionMatchSet CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sqlinjectionmatchset.html
type WAFSQLInjectionMatchSet struct {
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sqlinjectionmatchset.html#cfn-waf-sqlinjectionmatchset-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// SQLInjectionMatchTuples docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sqlinjectionmatchset.html#cfn-waf-sqlinjectionmatchset-sqlinjectionmatchtuples
	SQLInjectionMatchTuples *WAFSQLInjectionMatchSetSQLInjectionMatchTupleList `json:"SqlInjectionMatchTuples,omitempty"`
}

// CfnResourceType returns AWS::WAF::SqlInjectionMatchSet to implement the ResourceProperties interface
func (s WAFSQLInjectionMatchSet) CfnResourceType() string {

	return "AWS::WAF::SqlInjectionMatchSet"
}

// WAFWebACL represents the AWS::WAF::WebACL CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-webacl.html
type WAFWebACL struct {
	// DefaultAction docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-webacl.html#cfn-waf-webacl-defaultaction
	DefaultAction *WAFWebACLWafAction `json:"DefaultAction,omitempty" validate:"dive,required"`
	// MetricName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-webacl.html#cfn-waf-webacl-metricname
	MetricName *StringExpr `json:"MetricName,omitempty" validate:"dive,required"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-webacl.html#cfn-waf-webacl-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// Rules docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-webacl.html#cfn-waf-webacl-rules
	Rules *WAFWebACLActivatedRuleList `json:"Rules,omitempty"`
}

// CfnResourceType returns AWS::WAF::WebACL to implement the ResourceProperties interface
func (s WAFWebACL) CfnResourceType() string {

	return "AWS::WAF::WebACL"
}

// WAFXSSMatchSet represents the AWS::WAF::XssMatchSet CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-xssmatchset.html
type WAFXSSMatchSet struct {
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-xssmatchset.html#cfn-waf-xssmatchset-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// XSSMatchTuples docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-xssmatchset.html#cfn-waf-xssmatchset-xssmatchtuples
	XSSMatchTuples *WAFXSSMatchSetXSSMatchTupleList `json:"XssMatchTuples,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::WAF::XssMatchSet to implement the ResourceProperties interface
func (s WAFXSSMatchSet) CfnResourceType() string {

	return "AWS::WAF::XssMatchSet"
}

// WAFRegionalByteMatchSet represents the AWS::WAFRegional::ByteMatchSet CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-bytematchset.html
type WAFRegionalByteMatchSet struct {
	// ByteMatchTuples docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-bytematchset.html#cfn-wafregional-bytematchset-bytematchtuples
	ByteMatchTuples *WAFRegionalByteMatchSetByteMatchTupleList `json:"ByteMatchTuples,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-bytematchset.html#cfn-wafregional-bytematchset-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::WAFRegional::ByteMatchSet to implement the ResourceProperties interface
func (s WAFRegionalByteMatchSet) CfnResourceType() string {

	return "AWS::WAFRegional::ByteMatchSet"
}

// WAFRegionalIPSet represents the AWS::WAFRegional::IPSet CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ipset.html
type WAFRegionalIPSet struct {
	// IPSetDescriptors docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ipset.html#cfn-wafregional-ipset-ipsetdescriptors
	IPSetDescriptors *WAFRegionalIPSetIPSetDescriptorList `json:"IPSetDescriptors,omitempty"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ipset.html#cfn-wafregional-ipset-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::WAFRegional::IPSet to implement the ResourceProperties interface
func (s WAFRegionalIPSet) CfnResourceType() string {

	return "AWS::WAFRegional::IPSet"
}

// WAFRegionalRule represents the AWS::WAFRegional::Rule CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-rule.html
type WAFRegionalRule struct {
	// MetricName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-rule.html#cfn-wafregional-rule-metricname
	MetricName *StringExpr `json:"MetricName,omitempty" validate:"dive,required"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-rule.html#cfn-wafregional-rule-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// Predicates docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-rule.html#cfn-wafregional-rule-predicates
	Predicates *WAFRegionalRulePredicateList `json:"Predicates,omitempty"`
}

// CfnResourceType returns AWS::WAFRegional::Rule to implement the ResourceProperties interface
func (s WAFRegionalRule) CfnResourceType() string {

	return "AWS::WAFRegional::Rule"
}

// WAFRegionalSizeConstraintSet represents the AWS::WAFRegional::SizeConstraintSet CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sizeconstraintset.html
type WAFRegionalSizeConstraintSet struct {
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sizeconstraintset.html#cfn-wafregional-sizeconstraintset-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// SizeConstraints docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sizeconstraintset.html#cfn-wafregional-sizeconstraintset-sizeconstraints
	SizeConstraints *WAFRegionalSizeConstraintSetSizeConstraintList `json:"SizeConstraints,omitempty"`
}

// CfnResourceType returns AWS::WAFRegional::SizeConstraintSet to implement the ResourceProperties interface
func (s WAFRegionalSizeConstraintSet) CfnResourceType() string {

	return "AWS::WAFRegional::SizeConstraintSet"
}

// WAFRegionalSQLInjectionMatchSet represents the AWS::WAFRegional::SqlInjectionMatchSet CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sqlinjectionmatchset.html
type WAFRegionalSQLInjectionMatchSet struct {
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sqlinjectionmatchset.html#cfn-wafregional-sqlinjectionmatchset-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// SQLInjectionMatchTuples docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sqlinjectionmatchset.html#cfn-wafregional-sqlinjectionmatchset-sqlinjectionmatchtuples
	SQLInjectionMatchTuples *WAFRegionalSQLInjectionMatchSetSQLInjectionMatchTupleList `json:"SqlInjectionMatchTuples,omitempty"`
}

// CfnResourceType returns AWS::WAFRegional::SqlInjectionMatchSet to implement the ResourceProperties interface
func (s WAFRegionalSQLInjectionMatchSet) CfnResourceType() string {

	return "AWS::WAFRegional::SqlInjectionMatchSet"
}

// WAFRegionalWebACL represents the AWS::WAFRegional::WebACL CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webacl.html
type WAFRegionalWebACL struct {
	// DefaultAction docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webacl.html#cfn-wafregional-webacl-defaultaction
	DefaultAction *WAFRegionalWebACLAction `json:"DefaultAction,omitempty" validate:"dive,required"`
	// MetricName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webacl.html#cfn-wafregional-webacl-metricname
	MetricName *StringExpr `json:"MetricName,omitempty" validate:"dive,required"`
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webacl.html#cfn-wafregional-webacl-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// Rules docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webacl.html#cfn-wafregional-webacl-rules
	Rules *WAFRegionalWebACLRuleList `json:"Rules,omitempty"`
}

// CfnResourceType returns AWS::WAFRegional::WebACL to implement the ResourceProperties interface
func (s WAFRegionalWebACL) CfnResourceType() string {

	return "AWS::WAFRegional::WebACL"
}

// WAFRegionalWebACLAssociation represents the AWS::WAFRegional::WebACLAssociation CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webaclassociation.html
type WAFRegionalWebACLAssociation struct {
	// ResourceArn docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webaclassociation.html#cfn-wafregional-webaclassociation-resourcearn
	ResourceArn *StringExpr `json:"ResourceArn,omitempty" validate:"dive,required"`
	// WebACLID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webaclassociation.html#cfn-wafregional-webaclassociation-webaclid
	WebACLID *StringExpr `json:"WebACLId,omitempty" validate:"dive,required"`
}

// CfnResourceType returns AWS::WAFRegional::WebACLAssociation to implement the ResourceProperties interface
func (s WAFRegionalWebACLAssociation) CfnResourceType() string {

	return "AWS::WAFRegional::WebACLAssociation"
}

// WAFRegionalXSSMatchSet represents the AWS::WAFRegional::XssMatchSet CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-xssmatchset.html
type WAFRegionalXSSMatchSet struct {
	// Name docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-xssmatchset.html#cfn-wafregional-xssmatchset-name
	Name *StringExpr `json:"Name,omitempty" validate:"dive,required"`
	// XSSMatchTuples docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-xssmatchset.html#cfn-wafregional-xssmatchset-xssmatchtuples
	XSSMatchTuples *WAFRegionalXSSMatchSetXSSMatchTupleList `json:"XssMatchTuples,omitempty"`
}

// CfnResourceType returns AWS::WAFRegional::XssMatchSet to implement the ResourceProperties interface
func (s WAFRegionalXSSMatchSet) CfnResourceType() string {

	return "AWS::WAFRegional::XssMatchSet"
}

// WorkSpacesWorkspace represents the AWS::WorkSpaces::Workspace CloudFormation resource type
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspace.html
type WorkSpacesWorkspace struct {
	// BundleID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspace.html#cfn-workspaces-workspace-bundleid
	BundleID *StringExpr `json:"BundleId,omitempty" validate:"dive,required"`
	// DirectoryID docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspace.html#cfn-workspaces-workspace-directoryid
	DirectoryID *StringExpr `json:"DirectoryId,omitempty" validate:"dive,required"`
	// RootVolumeEncryptionEnabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspace.html#cfn-workspaces-workspace-rootvolumeencryptionenabled
	RootVolumeEncryptionEnabled *BoolExpr `json:"RootVolumeEncryptionEnabled,omitempty"`
	// UserName docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspace.html#cfn-workspaces-workspace-username
	UserName *StringExpr `json:"UserName,omitempty" validate:"dive,required"`
	// UserVolumeEncryptionEnabled docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspace.html#cfn-workspaces-workspace-uservolumeencryptionenabled
	UserVolumeEncryptionEnabled *BoolExpr `json:"UserVolumeEncryptionEnabled,omitempty"`
	// VolumeEncryptionKey docs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspace.html#cfn-workspaces-workspace-volumeencryptionkey
	VolumeEncryptionKey *StringExpr `json:"VolumeEncryptionKey,omitempty"`
}

// CfnResourceType returns AWS::WorkSpaces::Workspace to implement the ResourceProperties interface
func (s WorkSpacesWorkspace) CfnResourceType() string {

	return "AWS::WorkSpaces::Workspace"
}

// NewResourceByType returns a new resource object correspoding with the provided type
func NewResourceByType(typeName string) ResourceProperties {
	switch typeName {
	case "AWS::ApiGateway::Account":
		return &APIGatewayAccount{}
	case "AWS::ApiGateway::ApiKey":
		return &APIGatewayAPIKey{}
	case "AWS::ApiGateway::Authorizer":
		return &APIGatewayAuthorizer{}
	case "AWS::ApiGateway::BasePathMapping":
		return &APIGatewayBasePathMapping{}
	case "AWS::ApiGateway::ClientCertificate":
		return &APIGatewayClientCertificate{}
	case "AWS::ApiGateway::Deployment":
		return &APIGatewayDeployment{}
	case "AWS::ApiGateway::Method":
		return &APIGatewayMethod{}
	case "AWS::ApiGateway::Model":
		return &APIGatewayModel{}
	case "AWS::ApiGateway::Resource":
		return &APIGatewayResource{}
	case "AWS::ApiGateway::RestApi":
		return &APIGatewayRestAPI{}
	case "AWS::ApiGateway::Stage":
		return &APIGatewayStage{}
	case "AWS::ApiGateway::UsagePlan":
		return &APIGatewayUsagePlan{}
	case "AWS::ApplicationAutoScaling::ScalableTarget":
		return &ApplicationAutoScalingScalableTarget{}
	case "AWS::ApplicationAutoScaling::ScalingPolicy":
		return &ApplicationAutoScalingScalingPolicy{}
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
	case "AWS::CertificateManager::Certificate":
		return &CertificateManagerCertificate{}
	case "AWS::CloudFormation::CustomResource":
		return &CloudFormationCustomResource{}
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
	case "AWS::CodeBuild::Project":
		return &CodeBuildProject{}
	case "AWS::CodeCommit::Repository":
		return &CodeCommitRepository{}
	case "AWS::CodeDeploy::Application":
		return &CodeDeployApplication{}
	case "AWS::CodeDeploy::DeploymentConfig":
		return &CodeDeployDeploymentConfig{}
	case "AWS::CodeDeploy::DeploymentGroup":
		return &CodeDeployDeploymentGroup{}
	case "AWS::CodePipeline::CustomActionType":
		return &CodePipelineCustomActionType{}
	case "AWS::CodePipeline::Pipeline":
		return &CodePipelinePipeline{}
	case "AWS::Cognito::IdentityPool":
		return &CognitoIDentityPool{}
	case "AWS::Cognito::IdentityPoolRoleAttachment":
		return &CognitoIDentityPoolRoleAttachment{}
	case "AWS::Cognito::UserPool":
		return &CognitoUserPool{}
	case "AWS::Cognito::UserPoolClient":
		return &CognitoUserPoolClient{}
	case "AWS::Cognito::UserPoolGroup":
		return &CognitoUserPoolGroup{}
	case "AWS::Cognito::UserPoolUser":
		return &CognitoUserPoolUser{}
	case "AWS::Cognito::UserPoolUserToGroupAttachment":
		return &CognitoUserPoolUserToGroupAttachment{}
	case "AWS::Config::ConfigRule":
		return &ConfigConfigRule{}
	case "AWS::Config::ConfigurationRecorder":
		return &ConfigConfigurationRecorder{}
	case "AWS::Config::DeliveryChannel":
		return &ConfigDeliveryChannel{}
	case "AWS::DataPipeline::Pipeline":
		return &DataPipelinePipeline{}
	case "AWS::DirectoryService::MicrosoftAD":
		return &DirectoryServiceMicrosoftAD{}
	case "AWS::DirectoryService::SimpleAD":
		return &DirectoryServiceSimpleAD{}
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
	case "AWS::EC2::FlowLog":
		return &EC2FlowLog{}
	case "AWS::EC2::Host":
		return &EC2Host{}
	case "AWS::EC2::Instance":
		return &EC2Instance{}
	case "AWS::EC2::InternetGateway":
		return &EC2InternetGateway{}
	case "AWS::EC2::NatGateway":
		return &EC2NatGateway{}
	case "AWS::EC2::NetworkAcl":
		return &EC2NetworkACL{}
	case "AWS::EC2::NetworkAclEntry":
		return &EC2NetworkACLEntry{}
	case "AWS::EC2::NetworkInterface":
		return &EC2NetworkInterface{}
	case "AWS::EC2::NetworkInterfaceAttachment":
		return &EC2NetworkInterfaceAttachment{}
	case "AWS::EC2::PlacementGroup":
		return &EC2PlacementGroup{}
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
	case "AWS::EC2::SpotFleet":
		return &EC2SpotFleet{}
	case "AWS::EC2::Subnet":
		return &EC2Subnet{}
	case "AWS::EC2::SubnetCidrBlock":
		return &EC2SubnetCidrBlock{}
	case "AWS::EC2::SubnetNetworkAclAssociation":
		return &EC2SubnetNetworkACLAssociation{}
	case "AWS::EC2::SubnetRouteTableAssociation":
		return &EC2SubnetRouteTableAssociation{}
	case "AWS::EC2::VPC":
		return &EC2VPC{}
	case "AWS::EC2::VPCCidrBlock":
		return &EC2VPCCidrBlock{}
	case "AWS::EC2::VPCDHCPOptionsAssociation":
		return &EC2VPCDHCPOptionsAssociation{}
	case "AWS::EC2::VPCEndpoint":
		return &EC2VPCEndpoint{}
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
	case "AWS::EC2::Volume":
		return &EC2Volume{}
	case "AWS::EC2::VolumeAttachment":
		return &EC2VolumeAttachment{}
	case "AWS::ECR::Repository":
		return &ECRRepository{}
	case "AWS::ECS::Cluster":
		return &ECSCluster{}
	case "AWS::ECS::Service":
		return &ECSService{}
	case "AWS::ECS::TaskDefinition":
		return &ECSTaskDefinition{}
	case "AWS::EFS::FileSystem":
		return &EFSFileSystem{}
	case "AWS::EFS::MountTarget":
		return &EFSMountTarget{}
	case "AWS::EMR::Cluster":
		return &EMRCluster{}
	case "AWS::EMR::InstanceGroupConfig":
		return &EMRInstanceGroupConfig{}
	case "AWS::EMR::SecurityConfiguration":
		return &EMRSecurityConfiguration{}
	case "AWS::EMR::Step":
		return &EMRStep{}
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
	case "AWS::ElastiCache::SubnetGroup":
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
	case "AWS::ElasticLoadBalancingV2::Listener":
		return &ElasticLoadBalancingV2Listener{}
	case "AWS::ElasticLoadBalancingV2::ListenerRule":
		return &ElasticLoadBalancingV2ListenerRule{}
	case "AWS::ElasticLoadBalancingV2::LoadBalancer":
		return &ElasticLoadBalancingV2LoadBalancer{}
	case "AWS::ElasticLoadBalancingV2::TargetGroup":
		return &ElasticLoadBalancingV2TargetGroup{}
	case "AWS::Elasticsearch::Domain":
		return &ElasticsearchDomain{}
	case "AWS::Events::Rule":
		return &EventsRule{}
	case "AWS::GameLift::Alias":
		return &GameLiftAlias{}
	case "AWS::GameLift::Build":
		return &GameLiftBuild{}
	case "AWS::GameLift::Fleet":
		return &GameLiftFleet{}
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
	case "AWS::IoT::Certificate":
		return &IoTCertificate{}
	case "AWS::IoT::Policy":
		return &IoTPolicy{}
	case "AWS::IoT::PolicyPrincipalAttachment":
		return &IoTPolicyPrincipalAttachment{}
	case "AWS::IoT::Thing":
		return &IoTThing{}
	case "AWS::IoT::ThingPrincipalAttachment":
		return &IoTThingPrincipalAttachment{}
	case "AWS::IoT::TopicRule":
		return &IoTTopicRule{}
	case "AWS::KMS::Alias":
		return &KMSAlias{}
	case "AWS::KMS::Key":
		return &KMSKey{}
	case "AWS::Kinesis::Stream":
		return &KinesisStream{}
	case "AWS::KinesisFirehose::DeliveryStream":
		return &KinesisFirehoseDeliveryStream{}
	case "AWS::Lambda::Alias":
		return &LambdaAlias{}
	case "AWS::Lambda::EventSourceMapping":
		return &LambdaEventSourceMapping{}
	case "AWS::Lambda::Function":
		return &LambdaFunction{}
	case "AWS::Lambda::Permission":
		return &LambdaPermission{}
	case "AWS::Lambda::Version":
		return &LambdaVersion{}
	case "AWS::Logs::Destination":
		return &LogsDestination{}
	case "AWS::Logs::LogGroup":
		return &LogsLogGroup{}
	case "AWS::Logs::LogStream":
		return &LogsLogStream{}
	case "AWS::Logs::MetricFilter":
		return &LogsMetricFilter{}
	case "AWS::Logs::SubscriptionFilter":
		return &LogsSubscriptionFilter{}
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
	case "AWS::OpsWorks::UserProfile":
		return &OpsWorksUserProfile{}
	case "AWS::OpsWorks::Volume":
		return &OpsWorksVolume{}
	case "AWS::RDS::DBCluster":
		return &RDSDBCluster{}
	case "AWS::RDS::DBClusterParameterGroup":
		return &RDSDBClusterParameterGroup{}
	case "AWS::RDS::DBInstance":
		return &RDSDBInstance{}
	case "AWS::RDS::DBParameterGroup":
		return &RDSDBParameterGroup{}
	case "AWS::RDS::DBSecurityGroup":
		return &RDSDBSecurityGroup{}
	case "AWS::RDS::DBSecurityGroupIngress":
		return &RDSDBSecurityGroupIngress{}
	case "AWS::RDS::DBSubnetGroup":
		return &RDSDBSubnetGroup{}
	case "AWS::RDS::EventSubscription":
		return &RDSEventSubscription{}
	case "AWS::RDS::OptionGroup":
		return &RDSOptionGroup{}
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
	case "AWS::SNS::Subscription":
		return &SNSSubscription{}
	case "AWS::SNS::Topic":
		return &SNSTopic{}
	case "AWS::SNS::TopicPolicy":
		return &SNSTopicPolicy{}
	case "AWS::SQS::Queue":
		return &SQSQueue{}
	case "AWS::SQS::QueuePolicy":
		return &SQSQueuePolicy{}
	case "AWS::SSM::Association":
		return &SSMAssociation{}
	case "AWS::SSM::Document":
		return &SSMDocument{}
	case "AWS::SSM::Parameter":
		return &SSMParameter{}
	case "AWS::StepFunctions::Activity":
		return &StepFunctionsActivity{}
	case "AWS::StepFunctions::StateMachine":
		return &StepFunctionsStateMachine{}
	case "AWS::WAF::ByteMatchSet":
		return &WAFByteMatchSet{}
	case "AWS::WAF::IPSet":
		return &WAFIPSet{}
	case "AWS::WAF::Rule":
		return &WAFRule{}
	case "AWS::WAF::SizeConstraintSet":
		return &WAFSizeConstraintSet{}
	case "AWS::WAF::SqlInjectionMatchSet":
		return &WAFSQLInjectionMatchSet{}
	case "AWS::WAF::WebACL":
		return &WAFWebACL{}
	case "AWS::WAF::XssMatchSet":
		return &WAFXSSMatchSet{}
	case "AWS::WAFRegional::ByteMatchSet":
		return &WAFRegionalByteMatchSet{}
	case "AWS::WAFRegional::IPSet":
		return &WAFRegionalIPSet{}
	case "AWS::WAFRegional::Rule":
		return &WAFRegionalRule{}
	case "AWS::WAFRegional::SizeConstraintSet":
		return &WAFRegionalSizeConstraintSet{}
	case "AWS::WAFRegional::SqlInjectionMatchSet":
		return &WAFRegionalSQLInjectionMatchSet{}
	case "AWS::WAFRegional::WebACL":
		return &WAFRegionalWebACL{}
	case "AWS::WAFRegional::WebACLAssociation":
		return &WAFRegionalWebACLAssociation{}
	case "AWS::WAFRegional::XssMatchSet":
		return &WAFRegionalXSSMatchSet{}
	case "AWS::WorkSpaces::Workspace":
		return &WorkSpacesWorkspace{}

	default:
		for _, eachProvider := range customResourceProviders {
			customType := eachProvider(typeName)
			if nil != customType {
				return customType
			}
		}
	}
	return nil
}
