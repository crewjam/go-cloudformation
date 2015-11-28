// This program emits a cloudformation document for `app` to stdout
package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/crewjam/go-cloudformation"
)

func makeTemplate() *cloudformation.Template {
	t := cloudformation.NewTemplate()
	t.Description = "example production infrastructure"
	t.Parameters["DnsName"] = cloudformation.Parameter{
		Description: "The top level DNS name for the infrastructure",
		Type:        "String",
		Default:     "preview.example.io",
	}

	t.AddResource("ServerLoadBalancer", cloudformation.AWSElasticLoadBalancingLoadBalancer{
		ConnectionDrainingPolicy: &cloudformation.ElasticLoadBalancingConnectionDrainingPolicy{
			Enabled: true,
			Timeout: 30,
		},
	})

	/*

	     "Type": "AWS::ElasticLoadBalancing::LoadBalancer",
	     "Properties": {
	       "AppCookieStickinessPolicy": [],
	       "ConnectionDrainingPolicy": {
	         "Enabled": true,
	         "Timeout": 30
	       },
	       "CrossZone": true,
	       "HealthCheck": {
	         "HealthyThreshold": "2",
	         "Interval": "60",
	         "Target": "HTTP:80/",
	         "Timeout": "5",
	         "UnhealthyThreshold": "2"
	       },
	       "Listeners": [
	         {
	           "InstancePort": "8000",
	           "InstanceProtocol": "TCP",
	           "LoadBalancerPort": "443",
	           "Protocol": "SSL",
	           "SSLCertificateId": {
	             "Fn::Join": [
	               "",
	               [
	                 "arn:aws:iam::",
	                 {
	                   "Ref": "AWS::AccountId"
	                 },
	                 ":server-certificate/",
	                 {
	                   "Ref": "DnsName"
	                 }
	               ]
	             ]
	           }
	         }
	       ],
	       "Policies": [
	         {
	           "PolicyName": "EnableProxyProtocol",
	           "PolicyType": "ProxyProtocolPolicyType",
	           "Attributes": [
	             {
	               "Name": "ProxyProtocol",
	               "Value": "true"
	             }
	           ],
	           "InstancePorts": [
	             8000
	           ]
	         }
	       ],
	       "Subnets": [
	         {
	           "Ref": "VpcSubnetA"
	         },
	         {
	           "Ref": "VpcSubnetB"
	         },
	         {
	           "Ref": "VpcSubnetC"
	         }
	       ],
	       "SecurityGroups": [
	         {
	           "Ref": "LoadBalancerSecurityGroup"
	         }
	       ]
	     }
	   },
	*/

	return t
}

func main() {
	template := makeTemplate()
	buf, err := json.MarshalIndent(template, "", "  ")
	if err != nil {
		log.Fatalf("marshal: %s", err)
	}
	os.Stdout.Write(buf)
}
