package scraper

// See:
// * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-resource-specification-format.html and
// * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-resource-specification.html
// for more information

// CloudFormationSchema represents the root of the
// schema
type CloudFormationSchema struct {
	PropertyTypes                map[string]PropertyTypes
	ResourceTypes                map[string]ResourceTypes
	ResourceSpecificationVersion string
}

// PropertyTypes is a definition of a property
type PropertyTypes struct {
	Documentation string
	Properties    map[string]PropertyTypeDefinition
}

// ResourceTypes is a definition of a resource
type ResourceTypes struct {
	Documentation string
	Attributes    map[string]ResourceAttribute
	Properties    map[string]PropertyTypeDefinition
}

// ResourceAttribute are outputs of CloudFormation
// reosurce
type ResourceAttribute struct {
	PrimitiveType string
}

// PropertyTypeDefinition is the definition of a property
type PropertyTypeDefinition struct {
	Required          bool
	Documentation     string
	PrimitiveType     string
	UpdateType        string
	Type              string
	DuplicatesAllowed bool
	ItemType          string
	PrimitiveItemType string
}
