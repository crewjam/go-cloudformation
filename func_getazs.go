package cloudformation

func GetAZs(region StringExpr) GetAZsFunc {
	return GetAZsFunc{Region: region}
}

type GetAZsFunc struct {
	Region StringExpr `json:"Fn::GetAZs"`
}

func (f GetAZsFunc) StringList() *StringListExpr {
	return &StringListExpr{Func: f}
}

// Note: Fn::GetAZs does *not* implement StringFunc.
var _ StringListFunc = GetAZsFunc{} // GetAZsFunc must implement StringListFunc
