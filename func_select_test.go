package cloudformation

import (
	"encoding/json"

	. "gopkg.in/check.v1"
)

type SelectFuncTest struct{}

var _ = Suite(&SelectFuncTest{})

func (testSuite *SelectFuncTest) TestRef(c *C) {
	inputBuf := `{"Fn::Select":["2",{"Fn::GetAZs":{"Ref":"AWS::Region"}}]}`
	f, err := unmarshalFunc([]byte(inputBuf))
	c.Assert(err, IsNil)
	c.Assert(f, DeepEquals,
		SelectFunc{Selector: "2",
			Items: *GetAZs(*Ref("AWS::Region").String()).StringList()})

	buf, err := json.Marshal(f)
	c.Assert(err, IsNil)
	c.Assert(string(buf), Equals, inputBuf)
}
