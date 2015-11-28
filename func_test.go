package cloudformation

import (
	"encoding/json"

	. "gopkg.in/check.v1"
)

type FuncTest struct{}

var _ = Suite(&FuncTest{})

func (testSuite *FuncTest) TestRef(c *C) {
	inputBuf := `{"Ref": "foo"}`
	f, err := unmarshalFunc([]byte(inputBuf))
	c.Assert(err, IsNil)
	c.Assert(f, DeepEquals, RefFunc{Name: "foo"})

	buf, err := json.Marshal(f)
	c.Assert(err, IsNil)
	c.Assert(string(buf), Equals, `{"Ref":"foo"}`)
}
