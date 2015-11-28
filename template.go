package cloudformation

import (
	"encoding/json"
	"fmt"
)

type Template struct {
	AWSTemplateFormatVersion string               `json:",omitempty"`
	Description              string               `json:",omitempty"`
	Mappings                 map[string]Mapping   `json:",omitempty"`
	Parameters               map[string]Parameter `json:",omitempty"`
	Resources                map[string]Resource  `json:",omitempty"`
}

type Mapping map[string]map[string]string

type Parameter struct {
	Type        string `json:",omitempty"`
	Description string `json:",omitempty"`
	Default     string `json:",omitempty"`
}

type Resource struct {
	Type       string      `json:",omitempty"`
	DependsOn  string      `json:",omitempty"`
	Properties interface{} `json:",omitempty"`
}

func (r *Resource) UnmarshalJSON(buf []byte) error {
	m := map[string]interface{}{}
	if err := json.Unmarshal(buf, &m); err != nil {
		return err
	}

	r.Type = m["Type"].(string)
	r.DependsOn, _ = m["DependsOn"].(string)

	r.Properties = NewResourceByType(r.Type)
	if r.Properties == nil {
		return fmt.Errorf("unknown resource type: %s", r.Type)
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
