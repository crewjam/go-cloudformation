package cloudformation

import (
	"encoding/json"
)

func Base64(value String) String {
	v := struct {
		V String `json:"Fn::Base64"`
	}{V: value}
	buf, _ := json.Marshal(&v)
	return String(rawMarker + String(buf))
}

func FindInMap(mapName String, topLevelKey String, secondLevelKey String) String {
	return rawString(struct {
		V []String `json:"Fn::FindInMap"`
	}{V: []String{mapName, topLevelKey, secondLevelKey}})
}

func GetAtt(logicalNameOfResource String, attributeName String) String {
	v := struct {
		V []String `json:"Fn::GetAtt"`
	}{V: []String{logicalNameOfResource, attributeName}}
	buf, _ := json.Marshal(&v)
	return String(rawMarker + String(buf))
}

func GetAZs(region String) String {
	v := struct {
		V String `json:"Fn::GetAZs"`
	}{V: region}
	buf, _ := json.Marshal(&v)
	return String(rawMarker + String(buf))
}

func Join(delimiter string, values ...String) String {
	v := struct {
		V interface{} `json:"Fn::Join"`
	}{V: []interface{}{delimiter, values}}
	buf, _ := json.Marshal(&v)
	return String(rawMarker + String(buf))
}

func Select(index String, values ...String) String {
	v := struct {
		V interface{} `json:"Fn::Select"`
	}{V: []interface{}{index, values}}
	buf, _ := json.Marshal(&v)
	return String(rawMarker + String(buf))
}

func Ref(logicalName string) String {
	ref := struct{ Ref string }{Ref: logicalName}
	buf, _ := json.Marshal(&ref)
	return String(rawMarker + string(buf))
}
