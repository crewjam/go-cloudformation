package cloudformation

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

type StringExpression struct {
	Literal   string
	Ref       *RefExpression
	Join      *JoinExpression
	GetAtt    *getAttExpression
	Select    *selectExpression
	FindInMap *findInMapExpression
	Base64    *base64Expression
	If        *ifExpression
}

func (se StringExpression) MarshalJSON() ([]byte, error) {
	switch {
	case se.Ref != nil:
		return json.Marshal(se.Ref)
	case se.Join != nil:
		return json.Marshal(se.Join)
	case se.GetAtt != nil:
		return json.Marshal(se.GetAtt)
	case se.Select != nil:
		return json.Marshal(se.Select)
	case se.FindInMap != nil:
		return json.Marshal(se.FindInMap)
	case se.Base64 != nil:
		return json.Marshal(se.Base64)
	case se.If != nil:
		return json.Marshal(se.If)
	default:
		return json.Marshal(se.Literal)
	}
}

func (se *StringExpression) UnmarshalJSON(buf []byte) error {
	v := map[string]interface{}{}
	if err := json.Unmarshal(buf, &v); err == nil && len(v) == 1 {
		for key, _ := range v {
			switch key {
			case "Ref":
				return json.Unmarshal(buf, &se.Ref)
			case "Fn::Join":
				return json.Unmarshal(buf, &se.Join)
			case "Fn::GetAtt":
				return json.Unmarshal(buf, &se.GetAtt)
			case "Fn::Select":
				return json.Unmarshal(buf, &se.Select)
			case "Fn::FindInMap":
				return json.Unmarshal(buf, &se.FindInMap)
			case "Fn::Base64":
				return json.Unmarshal(buf, &se.Base64)
			case "Fn::If":
				return json.Unmarshal(buf, &se.If)
			}
			break
		}
	}

	return json.Unmarshal(buf, &se.Literal)
}

type RefExpression struct {
	Ref string
}

func String(v string) StringExpression {
	return StringExpression{Literal: v}
}

func Ref(ref string) StringExpression {
	return StringExpression{Ref: &RefExpression{Ref: ref}}
}

type JoinExpression struct {
	Separator StringExpression
	Items     []StringExpression
}

func (je JoinExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FnJoin []interface{} `json:"Fn::Join"`
	}{FnJoin: []interface{}{je.Separator, je.Items}})
}

func (je *JoinExpression) UnmarshalJSON(buf []byte) error {
	v := struct {
		FnJoin [2]interface{} `json:"Fn::Join"`
	}{}
	if err := json.Unmarshal(buf, &v); err != nil {
		return err
	}

	//log.Panicf("JoinExpression.UnmarshalJSON(%s)", buf)

	// The first item in Fn::Join is a StringExpression. The second
	// item is a list of StringExpression. The only golang type that
	// works for this is [2]interface{}. After we get the two types
	// split we need to unmarshal them again into their real types.
	// TODO(ross): this is an ugly hack.

	separatorBuf, err := json.Marshal(v.FnJoin[0])
	if err != nil {
		return err
	}
	if err := json.Unmarshal(separatorBuf, &je.Separator); err != nil {
		return err
	}

	// items need to be marshaled and re-unmarshaled. Ugh,
	// this is ugly.
	itemsBuf, err := json.Marshal(v.FnJoin[1])
	if err != nil {
		return err
	}

	if err := json.Unmarshal(itemsBuf, &je.Items); err != nil {
		return err
	}

	return nil
}

func GetAtt(resource, name string) StringExpression {
	return StringExpression{GetAtt: &getAttExpression{
		Resource: resource,
		Name:     name,
	}}
}

type getAttExpression struct {
	Resource string
	Name     string
}

func (ga getAttExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FnGetAtt []string `json:"Fn::GetAtt"`
	}{FnGetAtt: []string{ga.Resource, ga.Name}})
}

func (ga *getAttExpression) UnmarshalJSON(buf []byte) error {
	v := struct {
		FnGetAtt [2]string `json:"Fn::GetAtt"`
	}{}

	if err := json.Unmarshal(buf, &v); err != nil {
		return err
	}
	ga.Resource = v.FnGetAtt[0]
	ga.Name = v.FnGetAtt[1]
	return nil
}

func StringList(items ...StringExpression) StringListExpression {
	return StringListExpression{Literal: items}
}

type StringListExpression struct {
	Literal []StringExpression
	GetAZs  *getAZsExpression
	If      *listIfExpression
	NoValue *noValueExpression
}

func (se StringListExpression) MarshalJSON() ([]byte, error) {
	switch {
	case se.GetAZs != nil:
		return json.Marshal(se.GetAZs)
	case se.If != nil:
		return json.Marshal(se.If)
	case se.NoValue != nil:
		return json.Marshal(se.NoValue)
	default:
		return json.Marshal(se.Literal)
	}
}

func (se *StringListExpression) UnmarshalJSON(buf []byte) error {
	v := map[string]interface{}{}
	if err := json.Unmarshal(buf, &v); err == nil && len(v) == 1 {
		for key, value := range v {
			switch key {
			case "Fn::GetAZs":
				return json.Unmarshal(buf, &se.GetAZs)
			case "Fn::If":
				return json.Unmarshal(buf, &se.If)
			case "Ref":
				if value == "AWS::NoValue" {
					se.NoValue = &noValueExpression{}
					return nil
				}
			}
			break
		}
	}

	bestErr := json.Unmarshal(buf, &se.Literal)
	if bestErr != nil {
		log.Printf("buf: %q err=%s", string(buf), bestErr)
	}

	if false {
		if bestErr == nil {
			return nil
		}

		log.Printf("StringListExpression fallback %s", buf)

		// Annonyingly, CFN occasionally accepts a single expression when
		// a list of strings is called for.
		stringExpr := StringExpression{}
		err := json.Unmarshal(buf, &stringExpr)
		if err == nil {
			se.Literal = []StringExpression{stringExpr}
			return nil
		}
		log.Printf("single string error: %s", err)
	}

	// return the error from trying to unmarshal the literal, which
	// might be the most descriptive
	return bestErr

}

func GetAZs(region StringExpression) StringListExpression {
	return StringListExpression{GetAZs: &getAZsExpression{Region: region}}
}

type getAZsExpression struct {
	Region StringExpression
}

func (ga getAZsExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FnGetAZs StringExpression `json:"Fn::GetAZs"`
	}{FnGetAZs: ga.Region})
}

func (ga *getAZsExpression) UnmarshalJSON(buf []byte) error {
	v := struct {
		FnGetAZs StringExpression `json:"Fn::GetAZs"`
	}{}

	if err := json.Unmarshal(buf, &v); err != nil {
		return err
	}
	ga.Region = v.FnGetAZs
	return nil
}

func Select(selector int, items StringListExpression) StringExpression {
	return StringExpression{Select: &selectExpression{
		Selector: String(fmt.Sprintf("%d", selector)),
		Items:    items,
	}}
}

type selectExpression struct {
	Selector StringExpression
	Items    StringListExpression
}

func (se selectExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FnSelect []interface{} `json:"Fn::Select"`
	}{FnSelect: []interface{}{se.Selector, se.Items}})
}

func (se *selectExpression) UnmarshalJSON(buf []byte) error {
	v := struct {
		FnSelect [2]interface{} `json:"Fn::Select"`
	}{}
	if err := json.Unmarshal(buf, &v); err != nil {
		return err
	}

	// The first item in Fn::Select is a StringExpression. The second
	// item is a list of StringExpression. The only golang type that
	// works for this is [2]interface{}. After we get the two types
	// split we need to unmarshal them again into their real types.
	// TODO(ross): this is an ugly hack.

	separatorBuf, err := json.Marshal(v.FnSelect[0])
	if err != nil {
		return err
	}
	if err := json.Unmarshal(separatorBuf, &se.Selector); err != nil {
		return err
	}

	// items need to be marshaled and re-unmarshaled. Ugh,
	// this is ugly.
	itemsBuf, err := json.Marshal(v.FnSelect[1])
	if err != nil {
		return err
	}

	if err := json.Unmarshal(itemsBuf, &se.Items); err != nil {
		return err
	}

	return nil
}

type Integer int

func (li *Integer) UnmarshalJSON(buf []byte) error {
	var v int
	if err := json.Unmarshal(buf, &v); err == nil {
		*li = Integer(v)
		return nil
	}
	var v2 string
	if err := json.Unmarshal(buf, &v2); err != nil {
		return err
	}
	i, err := strconv.ParseInt(v2, 10, 64)
	if err != nil {
		return err
	}
	*li = Integer(int(i))
	return nil
}

type Bool struct {
	Literal bool
	Ref     string
}

func (li Bool) MarshalJSON() ([]byte, error) {
	if li.Ref != "" {
		return json.Marshal(struct {
			Ref string
		}{Ref: li.Ref})
	}
	return json.Marshal(li.Literal)
}

func (li *Bool) UnmarshalJSON(buf []byte) error {
	v3 := map[string]string{}
	if err := json.Unmarshal(buf, &v3); err == nil {
		ref, ok := v3["Ref"]
		if ok {
			li.Ref = ref
			return nil
		}
	}

	var v2 string
	if err := json.Unmarshal(buf, &v2); err == nil {
		i, err := strconv.ParseBool(v2)
		if err == nil {
			li.Literal = i
			return nil
		}
	}

	return json.Unmarshal(buf, &li.Literal)
}

func FindInMap(mapName, topLevelKey, secondLevelKey StringExpression) StringExpression {
	return StringExpression{FindInMap: &findInMapExpression{
		MapName:        mapName,
		TopLevelKey:    topLevelKey,
		SecondLevelKey: secondLevelKey,
	}}
}

type findInMapExpression struct {
	MapName        StringExpression
	TopLevelKey    StringExpression
	SecondLevelKey StringExpression
}

func (se findInMapExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FindInMap []StringExpression `json:"Fn::FindInMap"`
	}{FindInMap: []StringExpression{
		se.MapName,
		se.TopLevelKey,
		se.SecondLevelKey,
	}})
}

func (se *findInMapExpression) UnmarshalJSON(buf []byte) error {
	v := struct {
		FindInMap [3]StringExpression `json:"Fn::FindInMap"`
	}{}
	if err := json.Unmarshal(buf, &v); err != nil {
		return err
	}

	se.MapName = v.FindInMap[0]
	se.TopLevelKey = v.FindInMap[1]
	se.SecondLevelKey = v.FindInMap[2]
	return nil
}

func Base64(s StringExpression) StringExpression {
	return StringExpression{Base64: &base64Expression{Value: s}}
}

type base64Expression struct {
	Value StringExpression
}

func (se base64Expression) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FnBase64 StringExpression `json:"Fn::Base64"`
	}{FnBase64: se.Value})
}

func (se *base64Expression) UnmarshalJSON(buf []byte) error {
	v := struct {
		FnBase64 StringExpression `json:"Fn::Base64"`
	}{}
	if err := json.Unmarshal(buf, &v); err != nil {
		return err
	}
	se.Value = v.FnBase64
	return nil
}

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

func ListIf(condition string, valueIfTrue, valueIfFalse StringListExpression) StringListExpression {
	return StringListExpression{If: &listIfExpression{
		Condition:    condition,
		ValueIfTrue:  valueIfTrue,
		ValueIfFalse: valueIfFalse,
	}}
}

func NoValue() StringListExpression {
	return StringListExpression{NoValue: &noValueExpression{}}
}

type noValueExpression struct {
}

func (se noValueExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Ref string
	}{Ref: "AWS::NoValue"})
}

type listIfExpression struct {
	Condition    string
	ValueIfTrue  StringListExpression
	ValueIfFalse StringListExpression
}

func (se listIfExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FnIf []interface{} `json:"Fn::If"`
	}{FnIf: []interface{}{se.Condition, se.ValueIfTrue, se.ValueIfFalse}})
}

func (se *listIfExpression) UnmarshalJSON(buf []byte) error {
	v := struct {
		FnIf [3]interface{} `json:"Fn::If"`
	}{}
	if err := json.Unmarshal(buf, &v); err != nil {
		return err
	}

	buf2, err := json.Marshal(v.FnIf[0])
	if err != nil {
		return err
	}
	if err := json.Unmarshal(buf2, &se.Condition); err != nil {
		return err
	}

	buf2, err = json.Marshal(v.FnIf[1])
	if err != nil {
		return err
	}
	if err := json.Unmarshal(buf2, &se.ValueIfTrue); err != nil {
		return err
	}

	buf2, err = json.Marshal(v.FnIf[2])
	if err != nil {
		return err
	}
	if err := json.Unmarshal(buf2, &se.ValueIfFalse); err != nil {
		return err
	}

	return nil
}
