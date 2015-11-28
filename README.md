# WARNING

This is a super duper work in progress!

# Known Issues

The following appears in an official example:

    "DBSecurityGroup": {
      "Type": "AWS::RDS::DBSecurityGroup",
      "Condition" : "Is-EC2-Classic",
      "Properties": {
        "DBSecurityGroupIngress": {
          "EC2SecurityGroupName": { "Ref": "WebServerSecurityGroup" }
        },
        "GroupDescription": "database access"
      }
    }

even though the documentation says that `DBSecurityGroupIngress` is a *list* of objects, not an object.

