package cloudformation

import (
	"encoding/json"
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type StringTest struct{}

var _ = Suite(&StringTest{})

func (testSuite *StringTest) TestStringExpressionLiteral(c *C) {
	inputBuf := `"Hello, World!"`
	re := StringExpression{}
	err := json.Unmarshal([]byte(inputBuf), &re)
	c.Assert(err, IsNil)

	buf, err := json.Marshal(re)
	c.Assert(err, IsNil)
	c.Assert(inputBuf, Equals, string(buf))

	c.Assert(re, DeepEquals, StringExpression{Literal: "Hello, World!"})
}

func (testSuite *StringTest) TestRef(c *C) {
	inputBuf := `{"Ref":"AWS::AccountId"}`
	re := StringExpression{}
	err := json.Unmarshal([]byte(inputBuf), &re)
	c.Assert(err, IsNil)

	buf, err := json.Marshal(re)
	c.Assert(err, IsNil)
	c.Assert(inputBuf, Equals, string(buf))

	c.Assert(re, DeepEquals, Ref("AWS::AccountId"))
}

func (testSuite *StringTest) TestJoinExpression(c *C) {
	inputBuf := `{"Fn::Join":[";",["arn:aws:iam::",{"Ref":"AWS::AccountId"},":server-certificate/",{"Ref":"DnsName"}]]}`
	je := StringExpression{}
	err := json.Unmarshal([]byte(inputBuf), &je)
	c.Assert(err, IsNil)

	buf, err := json.Marshal(je)
	c.Assert(err, IsNil)
	c.Assert(inputBuf, Equals, string(buf))

	expected := StringExpression{Join: &JoinExpression{
		Separator: String(";"),
		Items: []StringExpression{
			String("arn:aws:iam::"),
			Ref("AWS::AccountId"),
			String(":server-certificate/"),
			Ref("DnsName"),
		},
	}}
	c.Assert(je, DeepEquals, expected)
}

func (testSuite *StringTest) TestGetAtt(c *C) {
	inputBuf := `{"Fn::GetAtt":["ServersSecurityGroup","GroupId"]}`
	je := StringExpression{}
	err := json.Unmarshal([]byte(inputBuf), &je)
	c.Assert(err, IsNil)

	buf, err := json.Marshal(je)
	c.Assert(err, IsNil)
	c.Assert(inputBuf, Equals, string(buf))

	expected := StringExpression{GetAtt: &getAttExpression{
		Resource: "ServersSecurityGroup",
		Name:     "GroupId",
	}}
	c.Assert(je, DeepEquals, expected)
}

func (testSuite *StringTest) TestStringListLiteral(c *C) {
	inputBuf := `["foo","bar"]`
	je := StringListExpression{}
	err := json.Unmarshal([]byte(inputBuf), &je)
	c.Assert(err, IsNil)

	buf, err := json.Marshal(je)
	c.Assert(err, IsNil)
	c.Assert(inputBuf, Equals, string(buf))

	expected := StringList(String("foo"), String("bar"))
	c.Assert(je, DeepEquals, expected)
}

func (testSuite *StringTest) TestGetAZs(c *C) {
	inputBuf := `{"Fn::GetAZs":{"Ref":"AWS::Region"}}`
	je := StringListExpression{}
	err := json.Unmarshal([]byte(inputBuf), &je)
	c.Assert(err, IsNil)

	buf, err := json.Marshal(je)
	c.Assert(err, IsNil)
	c.Assert(inputBuf, Equals, string(buf))

	expected := StringListExpression{GetAZs: &getAZsExpression{
		Region: Ref("AWS::Region"),
	}}
	c.Assert(je, DeepEquals, expected)
}

func (testSuite *StringTest) TestSelect(c *C) {
	inputBuf := `{"Fn::Select":["1",{"Fn::GetAZs":{"Ref":"AWS::Region"}}]}`
	je := StringExpression{}
	err := json.Unmarshal([]byte(inputBuf), &je)
	c.Assert(err, IsNil)

	buf, err := json.Marshal(je)
	c.Assert(err, IsNil)
	c.Assert(inputBuf, Equals, string(buf))

	expected := Select(1, GetAZs(Ref("AWS::Region")))
	c.Assert(je, DeepEquals, expected)
}

func (testSuite *StringTest) TestFindInMap(c *C) {
	inputBuf := `{"Fn::FindInMap":["AWSRegionArch2AMI",{"Ref":"AWS::Region"},{"Fn::FindInMap":["AWSInstanceType2Arch",{"Ref":"InstanceType"},"Arch"]}]}`

	je := StringExpression{}
	err := json.Unmarshal([]byte(inputBuf), &je)
	c.Assert(err, IsNil)

	buf, err := json.Marshal(je)
	c.Assert(err, IsNil)
	c.Assert(inputBuf, Equals, string(buf))

	expected := FindInMap(String("AWSRegionArch2AMI"), Ref("AWS::Region"),
		FindInMap(String("AWSInstanceType2Arch"), Ref("InstanceType"), String("Arch")))
	c.Assert(je, DeepEquals, expected)
}

func (testSuite *StringTest) TestListIf(c *C) {
	inputBuf := `{"Fn::If":["Is-EC2-VPC",[{"Fn::GetAtt":["DBEC2SecurityGroup","GroupId"]}],{"Ref":"AWS::NoValue"}]}`

	je := StringListExpression{}
	err := json.Unmarshal([]byte(inputBuf), &je)
	c.Assert(err, IsNil)

	buf, err := json.Marshal(je)
	c.Assert(err, IsNil)
	c.Assert(inputBuf, Equals, string(buf))

	expected := ListIf("Is-EC2-VPC",
		StringList(GetAtt("DBEC2SecurityGroup", "GroupId")),
		NoValue())
	c.Assert(je, DeepEquals, expected)
}

/*
func (testSuite *StringTest) TestThing(c *C) {
	inputBuf := `{
      "Type": "AWS::RDS::DBInstance",
      "Properties": {
        "Engine" : "MySQL",
        "DBName" : { "Ref": "DBName" },
        "MultiAZ" : { "Ref": "MultiAZDatabase" },
        "MasterUsername": { "Ref": "DBUser" },
        "MasterUserPassword": { "Ref" : "DBPassword" },
        "DBInstanceClass": { "Ref" : "DBInstanceClass" },
        "AllocatedStorage": { "Ref" : "DBAllocatedStorage" },
        "VPCSecurityGroups": { "Fn::If" : [ "Is-EC2-VPC", [ { "Fn::GetAtt": [ "DBEC2SecurityGroup", "GroupId" ] } ], { "Ref" : "AWS::NoValue"}]},
        "DBSecurityGroups": { "Fn::If" : [ "Is-EC2-Classic", [ { "Ref": "DBSecurityGroup" } ], { "Ref" : "AWS::NoValue"}]}
      }
    }`

	je := Resource{}
	err := json.Unmarshal([]byte(inputBuf), &je)
	c.Assert(err, IsNil)

	buf, err := json.Marshal(je)
	c.Assert(err, IsNil)
	//c.Assert(inputBuf, Equals, string(buf))

	expected := Resource{
		Type: "AWS::RDS::DBInstance",
		Properties: &AWSRDSDBInstance{
			Engine:             String("MySQL"),
			DBName:             Ref("DBName"),
			MultiAZ:            Ref("MultiAZDatabase"),
			MasterUsername:     Ref("DBUser"),
			MasterUserPassword: Ref("DBPassword"),
			DBInstanceClass:    Ref("DBInstanceClass"),
			AllocatedStorage:   Ref("DBAllocatedStorage"),
			VPCSecurityGroups:  ListIf("Is-EC2-VPC", StringList(GetAtt("DBEC2SecurityGroup", "GroupId")), NoValue()),
			DBSecurityGroups:   ListIf("Is-EC2-Classic", StringList(Ref("DBSecurityGroup")), NoValue()),
		},
	}
	c.Assert(je, DeepEquals, expected)
}
*/
