package main

import (
	"fmt"
	"log"
	"net"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func getIp(ddns string) string{
	var rtn string
	ips, _ := net.LookupIP(ddns)

	for _, ip := range ips {
		//fmt.Printf(ip.String())
		rtn = ip.String()
		break
	}
	return rtn
}
func doReq(){
	sess, _ := session.NewSession()
	myec2 := ec2.New(sess)

	ddns := getIp("www.debian.org")
	input := ec2.ModifySecurityGroupRulesInput {
		GroupId: aws.String("sg-095963bf8f65baddc"), // Id of security Group
		SecurityGroupRules:  []*ec2.SecurityGroupRuleUpdate{ 
			{
			SecurityGroupRule: &ec2.SecurityGroupRuleRequest{
				 CidrIpv4: aws.String(ddns+"/32"),
				 Description: aws.String("DDNS ip"),
				 FromPort: aws.Int64(22), // Port 22 SSH
				 ToPort: aws.Int64(22),// Port 22 SSH
				 IpProtocol: aws.String("TCP"),
			},
			SecurityGroupRuleId: aws.String("sgr-056c6aa33ed3707b9"), // ID of security groupd Rule
			},
		},
	}
	req, err := myec2.ModifySecurityGroupRules(&input)  //Do the request
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(req)
}

func main() {
	lambda.Start(doReq)
}

			

				
			
