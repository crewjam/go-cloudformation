package cloudformation

import (
	"encoding/json"

	. "gopkg.in/check.v1"
)

type JoinFuncTest struct{}

var _ = Suite(&JoinFuncTest{})

func (testSuite *JoinFuncTest) TestRef(c *C) {
	inputBuf := `{"Fn::Join":["x",["y",{"Ref":"foo"},"1"]]}`
	f, err := unmarshalFunc([]byte(inputBuf))
	c.Assert(err, IsNil)
	c.Assert(f, DeepEquals, Join("x", *String("y"), *Ref("foo").String(), *String("1")))

	buf, err := json.Marshal(f)
	c.Assert(err, IsNil)
	c.Assert(string(buf), Equals, inputBuf)
}
