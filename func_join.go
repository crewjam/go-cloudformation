package cloudformation

import "encoding/json"

func Join(separator string, items ...StringExpr) JoinFunc {
	return JoinFunc{Separator: separator, Items: items}
}

type JoinFunc struct {
	Separator string
	Items     []StringExpr
}

func (f JoinFunc) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FnJoin []interface{} `json:"Fn::Join"`
	}{FnJoin: []interface{}{f.Separator, f.Items}})
}

func (f *JoinFunc) UnmarshalJSON(buf []byte) error {
	v := struct {
		FnJoin [2]json.RawMessage `json:"Fn::Join"`
	}{}
	if err := json.Unmarshal(buf, &v); err != nil {
		return err
	}
	if err := json.Unmarshal(v.FnJoin[0], &f.Separator); err != nil {
		return err
	}
	if err := json.Unmarshal(v.FnJoin[1], &f.Items); err != nil {
		return err
	}

	return nil
}

func (f JoinFunc) String() *StringExpr {
	return &StringExpr{Func: f}
}

var _ StringFunc = JoinFunc{} // JoinFunc must implement StringFunc
