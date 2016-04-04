package cloudformation

// IAMPolicyDocument represents an IAM policy document
type IAMPolicyDocument struct {
	Version   string `json:",omitempty"`
	Statement []IAMPolicyStatement
}

// IAMPrincipal represents a principal in an IAM policy
type IAMPrincipal struct {
	Service *StringListExpr `json:",omitempty"`
}

// IAMPolicyStatement represents an IAM policy statement
type IAMPolicyStatement struct {
	Sid            string          `json:",omitempty"`
	Effect         string          `json:",omitempty"`
	Principal      *IAMPrincipal   `json:",omitempty"`
	Action         *StringListExpr `json:",omitempty"`
	Resource       *StringListExpr `json:",omitempty"`
	ConditionBlock interface{}     `json:",omitempty"`
}
