package cloudformation

func Base64(value StringExpr) Base64Func {
	return Base64Func{Value: value}
}

type Base64Func struct {
	Value StringExpr `json:"Fn::Base64"`
}

func (f Base64Func) String() *StringExpr {
	return &StringExpr{Func: f}
}

var _ StringFunc = Base64Func{} // Base64Func must implement StringFunc
