package cloudformation

import (
	"encoding/json"
	"fmt"
)

func NewTemplate() *Template {
	return &Template{
		AWSTemplateFormatVersion: "2010-09-09",
		Mappings:                 map[string]Mapping{},
		Parameters:               map[string]Parameter{},
		Resources:                map[string]Resource{},
	}
}

type Template struct {
	AWSTemplateFormatVersion string               `json:",omitempty"`
	Description              string               `json:",omitempty"`
	Mappings                 map[string]Mapping   `json:",omitempty"`
	Parameters               map[string]Parameter `json:",omitempty"`
	Resources                map[string]Resource  `json:",omitempty"`
}

func (t *Template) AddResource(name string, resource ResourceProperties) {
	t.Resources[name] = Resource{Properties: resource}
}

type Mapping map[string]map[string]string

type Parameter struct {
	Type        string `json:",omitempty"`
	Description string `json:",omitempty"`
	Default     string `json:",omitempty"`
}

type ResourceProperties interface {
	ResourceType() string
}

type Resource struct {
	DependsOn  string
	Properties ResourceProperties
}

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
