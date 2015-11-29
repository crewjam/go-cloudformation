package cloudformation

import (
	"encoding/json"
	"fmt"
)

// NewTemplate returns a new empty Template initialized with some
// default values.
func NewTemplate() *Template {
	return &Template{
		AWSTemplateFormatVersion: "2010-09-09",
		Mappings:                 map[string]Mapping{},
		Parameters:               map[string]Parameter{},
		Resources:                map[string]Resource{},
	}
}

// Template represents a cloudformation template.
type Template struct {
	AWSTemplateFormatVersion string               `json:",omitempty"`
	Description              string               `json:",omitempty"`
	Mappings                 map[string]Mapping   `json:",omitempty"`
	Parameters               map[string]Parameter `json:",omitempty"`
	Resources                map[string]Resource  `json:",omitempty"`
}

// AddResource adds the resource to the template as name, displacing
// any resource with the same name that already exists.
func (t *Template) AddResource(name string, resource ResourceProperties) {
	t.Resources[name] = Resource{Properties: resource}
}

// Mapping matches a key to a corresponding set of named values. For example,
// if you want to set values based on a region, you can create a mapping that
// uses the region name as a key and contains the values you want to specify
// for each specific region. You use the Fn::FindInMap intrinsic function to
// retrieve values in a map.
//
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/mappings-section-structure.html
type Mapping map[string]map[string]string

// Parameter represents a parameter to the template.
//
// You can use the optional Parameters section to pass values into your
// template when you create a stack. With parameters, you can create templates
// that are customized each time you create a stack. Each parameter must
// contain a value when you create a stack. You can specify a default value to
// make the parameter optional.
//
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/parameters-section-structure.html
type Parameter struct {
	Type        string `json:",omitempty"`
	Description string `json:",omitempty"`
	Default     string `json:",omitempty"`
}

// ResourceProperties is an interface that is implemented by resource objects.
type ResourceProperties interface {
	ResourceType() string
}

// Resource represents a resource in a cloudformation template. It contains resource
// metadata and, in Properties, a struct that implements ResourceProperties which
// contains the properties of the resource.
type Resource struct {
	DependsOn  string
	Properties ResourceProperties
}

// MarshalJSON returns a JSON representation of the object
func (r Resource) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type       string
		DependsOn  string `json:",omitempty"`
		Properties ResourceProperties
	}{
		Type:       r.Properties.ResourceType(),
		DependsOn:  r.DependsOn,
		Properties: r.Properties,
	})
}

// UnmarshalJSON sets the object from the provided JSON representation
func (r *Resource) UnmarshalJSON(buf []byte) error {
	m := map[string]interface{}{}
	if err := json.Unmarshal(buf, &m); err != nil {
		return err
	}

	typeName := m["Type"].(string)
	r.DependsOn, _ = m["DependsOn"].(string)

	r.Properties = NewResourceByType(typeName)
	if r.Properties == nil {
		return fmt.Errorf("unknown resource type: %s", typeName)
	}

	propertiesBuf, err := json.Marshal(m["Properties"])
	if err != nil {
		return err
	}
	if err := json.Unmarshal(propertiesBuf, r.Properties); err != nil {
		return err
	}
	return nil
}
