package cloudformation

import "encoding/json"

func Select(selector string, items ...StringExpr) SelectFunc {
	return SelectFunc{Selector: selector, Items: StringListExpr{Literal: items}}
}

type SelectFunc struct {
	Selector string
	Items    StringListExpr
}

func (f SelectFunc) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FnSelect []interface{} `json:"Fn::Select"`
	}{FnSelect: []interface{}{f.Selector, f.Items}})
}

func (f *SelectFunc) UnmarshalJSON(buf []byte) error {
	v := struct {
		FnSelect [2]json.RawMessage `json:"Fn::Select"`
	}{}
	if err := json.Unmarshal(buf, &v); err != nil {
		return err
	}
	if err := json.Unmarshal(v.FnSelect[0], &f.Selector); err != nil {
		return err
	}
	if err := json.Unmarshal(v.FnSelect[1], &f.Items); err != nil {
		return err
	}

	return nil
}

func (f SelectFunc) String() *StringExpr {
	return &StringExpr{Func: f}
}

var _ StringFunc = SelectFunc{} // SelectFunc must implement StringFunc
