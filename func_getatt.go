package cloudformation

import "encoding/json"

func GetAtt(resource, name string) GetAttFunc {
	return GetAttFunc{Resource: resource, Name: name}
}

type GetAttFunc struct {
	Resource string
	Name     string
}

func (f GetAttFunc) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FnGetAtt []string `json:"Fn::GetAtt"`
	}{FnGetAtt: []string{f.Resource, f.Name}})
}

func (f *GetAttFunc) UnmarshalJSON(buf []byte) error {
	v := struct {
		FnGetAtt [2]string `json:"Fn::GetAtt"`
	}{}
	if err := json.Unmarshal(buf, &v); err != nil {
		return err
	}
	f.Resource = v.FnGetAtt[0]
	f.Name = v.FnGetAtt[1]
	return nil
}

func (f GetAttFunc) String() *StringExpr {
	return &StringExpr{Func: f}
}

var _ StringFunc = GetAttFunc{} // GetAttFunc must implement StringFunc
