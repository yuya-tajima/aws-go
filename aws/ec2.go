package aws

import (
	"fmt"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type instances []*ec2.Instance

func (a *Aws) Start(s string) {

	input := &ec2.StartInstancesInput{
		InstanceIds: []*string{
			aws.String(s),
		},
	}

	result, err := a.ec2.StartInstances(input)

	MaybeError(err)

	fmt.Println(result)
}

func (a *Aws) Stop(s string) {

	input := &ec2.StopInstancesInput{
		InstanceIds: []*string{
			aws.String(s),
		},
	}

	result, err := a.ec2.StopInstances(input)

	MaybeError(err)

	fmt.Println(result)
}

func (a *Aws) SetClient() {
	a.ec2 = ec2.New(a.session)
}

func ShowDetails(i *ec2.Instance) {
	fmt.Printf("InstanceId: %s \n", *i.InstanceId)
	fmt.Printf("ImageId: %s \n", *i.ImageId)
	fmt.Printf("InstanceType: %s \n", *i.InstanceType)
	fmt.Printf("PrivateIpAddress: %s \n", *i.PrivateIpAddress)
	fmt.Printf("PublicIpAddress: %s \n", *i.PublicIpAddress)
	fmt.Printf("State: code:%d name:%s \n", *i.State.Code, *i.State.Name)

	if len(i.Tags) > 0 {
		fmt.Print("*** The following tags are defined ***\n")
		for _, v := range i.Tags {
			fmt.Printf("%v:%v\n", *v.Key, *v.Value)
		}
	}
}

func HasTagName(tag string, i *ec2.Instance) bool {
	if len(i.Tags) == 0 {
		return false
	}

	r := regexp.MustCompile(tag)
	for _, v := range i.Tags {
		if *v.Key == "Name" && r.MatchString(*v.Value) {
			return true
		}
	}

	return false
}

func (a *Aws) GetInstances() (instances, error) {

	result, err := a.ec2.DescribeInstances(nil)

	ins := instances{}

	if err != nil {
		fmt.Println("Error", err)
	} else {
		if len(result.Reservations) > 0 {
			for _, r := range result.Reservations {
				for _, i := range r.Instances {
					ins = append(ins, i)
				}
			}
		}
	}

	return ins, err
}
