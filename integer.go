package cloudformation

import (
	"encoding/json"
	"strconv"
)

// IntegerExpr is a intean expression. If the value is computed then
// Func will be non-nill. If it is a literal `true` or `false` then
// the Literal gives the value. Typically instances of this function
// are created by Integer() or one of the function constructors. Ex:
//
//   type LocalBalancer struct {
//     CrossZone *IntegerExpr
//   }
//
//   lb := LocalBalancer{CrossZone: Integer(true)}
//   lb2 := LocalBalancer{CrossZone: Ref("LoadBalancerCrossZone").Integer()}
//
type IntegerExpr struct {
	Func    IntegerFunc
	Literal int
}

func (x IntegerExpr) MarshalJSON() ([]byte, error) {
	if x.Func != nil {
		return json.Marshal(x.Func)
	} else {
		return json.Marshal(x.Literal)
	}
}

func (x *IntegerExpr) UnmarshalJSON(data []byte) error {
	var v int
	err := json.Unmarshal(data, &v)
	if err == nil {
		x.Func = nil
		x.Literal = v
		return nil
	}

	// Cloudformation allows int values to be represented as strings
	var strValue string
	if err := json.Unmarshal(data, &strValue); err == nil {
		if v, err := strconv.ParseInt(strValue, 10, 64); err == nil {
			x.Func = nil
			x.Literal = int(v)
			return nil
		}
	}

	// Perhaps we have a serialized function call (like `{"Ref": "Foo"}`)
	// so we'll try to unmarshal it with UnmarshalFunc. Not all Funcs also
	// implement IntegerFunc, so we have to make sure that the referenced
	// function actually works in the intean context
	funcCall, err2 := unmarshalFunc(data)
	if err2 == nil {
		intFunc, ok := funcCall.(IntegerFunc)
		if ok {
			x.Func = intFunc
			return nil
		}
	} else if unknownFunctionErr, ok := err2.(UnknownFunctionError); ok {
		return unknownFunctionErr
	}

	// Return the original error trying to unmarshal the literal expression,
	// which will be the most expressive.
	return err
}

// Integer returns a new IntegerExpr representing the literal value v.
func Integer(v int) *IntegerExpr {
	return &IntegerExpr{Literal: v}
}
