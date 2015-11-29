package cloudformation

import (
	"encoding/json"
	"fmt"
)

type Func interface {
}

type BoolFunc interface {
	Bool() *BoolExpr
}

type IntegerFunc interface {
	Integer() *IntegerExpr
}

type StringFunc interface {
	String() *StringExpr
}

type StringListFunc interface {
	StringList() *StringListExpr
}

type UnknownFunctionError struct {
	Name string
}

func (ufe UnknownFunctionError) Error() string {
	return fmt.Sprintf("unkown function %s", ufe.Name)
}

func unmarshalFunc(data []byte) (Func, error) {
	rawDecode := map[string]json.RawMessage{}
	err := json.Unmarshal(data, &rawDecode)
	if err != nil {
		return nil, err
	}
	for key := range rawDecode {
		switch key {
		case "Ref":
			f := RefFunc{}
			if err := json.Unmarshal(data, &f); err == nil {
				return f, nil
			}
		case "Fn::Join":
			f := JoinFunc{}
			if err := json.Unmarshal(data, &f); err == nil {
				return f, nil
			}
		case "Fn::Select":
			f := SelectFunc{}
			if err := json.Unmarshal(data, &f); err == nil {
				return f, nil
			}
		case "Fn::GetAtt":
			f := GetAttFunc{}
			if err := json.Unmarshal(data, &f); err == nil {
				return f, nil
			}
		case "Fn::FindInMap":
			f := FindInMapFunc{}
			if err := json.Unmarshal(data, &f); err == nil {
				return f, nil
			}
		case "Fn::Base64":
			f := Base64Func{}
			if err := json.Unmarshal(data, &f); err == nil {
				return f, nil
			}
		case "Fn::GetAZs":
			f := GetAZsFunc{}
			if err := json.Unmarshal(data, &f); err == nil {
				return f, nil
			}
		case "Fn::If":
			f := IfFunc{}
			if err := json.Unmarshal(data, &f); err == nil {
				return f, nil
			}
		default:
			return nil, UnknownFunctionError{Name: key}
		}
	}
	return nil, fmt.Errorf("cannot decode function")
}
