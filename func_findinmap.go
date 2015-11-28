package cloudformation

import "encoding/json"

func FindInMap(mapName string, topLevelKey StringExpr, secondLevelKey StringExpr) FindInMapFunc {
	return FindInMapFunc{
		MapName:        mapName,
		TopLevelKey:    topLevelKey,
		SecondLevelKey: secondLevelKey,
	}
}

type FindInMapFunc struct {
	MapName        string
	TopLevelKey    StringExpr
	SecondLevelKey StringExpr
}

func (f FindInMapFunc) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FnFindInMap []interface{} `json:"Fn::FindInMap"`
	}{FnFindInMap: []interface{}{f.MapName, f.TopLevelKey, f.SecondLevelKey}})
}

func (f *FindInMapFunc) UnmarshalJSON(buf []byte) error {
	v := struct {
		FnFindInMap [3]json.RawMessage `json:"Fn::FindInMap"`
	}{}
	if err := json.Unmarshal(buf, &v); err != nil {
		return err
	}
	if err := json.Unmarshal(v.FnFindInMap[0], &f.MapName); err != nil {
		return err
	}
	if err := json.Unmarshal(v.FnFindInMap[1], &f.TopLevelKey); err != nil {
		return err
	}
	if err := json.Unmarshal(v.FnFindInMap[2], &f.SecondLevelKey); err != nil {
		return err
	}

	return nil
}

func (f FindInMapFunc) String() *StringExpr {
	return &StringExpr{Func: f}
}

var _ StringFunc = FindInMapFunc{} // FindInMapFunc must implement StringFunc
