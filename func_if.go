package cloudformation

import "encoding/json"
import "reflect"

func If(condition string, valueIfTrue, valueIfFalse StringExpr) IfFunc {
	return IfFunc{
		list:         false,
		Condition:    condition,
		ValueIfTrue:  valueIfTrue,
		ValueIfFalse: valueIfFalse,
	}
}

func IfList(condition string, valueIfTrue, valueIfFalse StringListExpr) IfFunc {
	return IfFunc{
		list:         true,
		Condition:    condition,
		ValueIfTrue:  valueIfTrue,
		ValueIfFalse: valueIfFalse,
	}
}

type IfFunc struct {
	list         bool
	Condition    string
	ValueIfTrue  interface{}
	ValueIfFalse interface{}
}

func (f IfFunc) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FnIf []interface{} `json:"Fn::If"`
	}{FnIf: []interface{}{f.Condition, f.ValueIfTrue, f.ValueIfFalse}})
}

func (f *IfFunc) UnmarshalJSON(buf []byte) error {
	v := struct {
		FnIf [3]json.RawMessage `json:"Fn::If"`
	}{}
	if err := json.Unmarshal(buf, &v); err != nil {
		return err
	}
	if err := json.Unmarshal(v.FnIf[0], &f.Condition); err != nil {
		return err
	}

	var probeValue interface{}
	if err := json.Unmarshal(v.FnIf[1], &probeValue); err != nil {
		return err
	}

	switch reflect.ValueOf(probeValue).Kind() {
	case reflect.Array:
		f.list = true
	case reflect.String:
		f.list = false
	case reflect.Map:
		expr, err := unmarshalFunc(v.FnIf[1])
		if err == nil {
			if _, ok := expr.(StringListFunc); ok {
				f.list = true
			}
		}
	}

	if f.list {
		f.ValueIfTrue = StringListExpr{}
		f.ValueIfFalse = StringListExpr{}
	} else {
		f.ValueIfTrue = StringExpr{}
		f.ValueIfFalse = StringExpr{}
	}

	if err := json.Unmarshal(v.FnIf[1], &f.ValueIfTrue); err != nil {
		return err
	}

	if err := json.Unmarshal(v.FnIf[2], &f.ValueIfFalse); err != nil {
		return err
	}
	return nil
}

func (f IfFunc) String() *StringExpr {
	if f.list {
		panic("IfFunc is a list, but being treated as a scalar")
	}
	return &StringExpr{Func: f}
}

func (f IfFunc) StringList() *StringListExpr {
	if !f.list {
		panic("IfFunc is a scaler, but being treated as a list of strings.")
	}
	return &StringListExpr{Func: f}
}

var _ StringFunc = IfFunc{}     // IfFunc must implement StringFunc
var _ StringListFunc = IfFunc{} // IfFunc must implement StringListFunc

/*

func If(condition, valueIfTrue, valueIfFalse StringExpression) StringExpression {
	return StringExpression{If: &ifExpression{
		Condition:    condition,
		ValueIfTrue:  valueIfTrue,
		ValueIfFalse: valueIfFalse,
	}}
}

type ifExpression struct {
	Condition    StringExpression
	ValueIfTrue  StringExpression
	ValueIfFalse StringExpression
}

func (se ifExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FnIf []StringExpression `json:"Fn::If"`
	}{FnIf: []StringExpression{se.Condition, se.ValueIfTrue, se.ValueIfFalse}})
}

func (se *ifExpression) UnmarshalJSON(buf []byte) error {
	v := struct {
		FnIf [3]StringExpression `json:"Fn::If"`
	}{}
	if err := json.Unmarshal(buf, &v); err != nil {
		return err
	}
	se.Condition = v.FnIf[0]
	se.ValueIfTrue = v.FnIf[1]
	se.ValueIfFalse = v.FnIf[2]
	return nil
}
*/
