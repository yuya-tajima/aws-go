package aws

import (
	"fmt"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	_ec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type ec2 struct {
	svc *_ec2.EC2
}

type instances []*_ec2.Instance

func (a *Aws) SetEc2Client() {
	a.ec2 = &ec2{
		svc: _ec2.New(a.session),
	}
}

func (a *ec2) Reboot(s *string, dry bool) {

	input := &_ec2.RebootInstancesInput{
		InstanceIds: []*string{
			s,
		},
		DryRun: aws.Bool(dry),
	}

	result, err := a.svc.RebootInstances(input)

	MaybeExitError(err)

	fmt.Println(result)
}

func (a *ec2) Start(s *string, dry bool) {

	input := &_ec2.StartInstancesInput{
		InstanceIds: []*string{
			s,
		},
		DryRun: aws.Bool(dry),
	}

	result, err := a.svc.StartInstances(input)

	MaybeExitError(err)

	fmt.Println(result)
}

func (a *ec2) Stop(s *string, dry bool) {

	input := &_ec2.StopInstancesInput{
		InstanceIds: []*string{
			s,
		},
		DryRun: aws.Bool(dry),
	}

	result, err := a.svc.StopInstances(input)

	MaybeError(err)

	fmt.Println(result)
}

func ShowDetails(i *_ec2.Instance) {
	fmt.Printf("InstanceId: %s \n", aws.StringValue(i.InstanceId))
	fmt.Printf("ImageId: %s \n", aws.StringValue(i.ImageId))
	fmt.Printf("InstanceType: %s \n", aws.StringValue(i.InstanceType))
	fmt.Printf("PrivateIpAddress: %s \n", aws.StringValue(i.PrivateIpAddress))
	fmt.Printf("PublicIpAddress: %s \n", aws.StringValue(i.PublicIpAddress))
	fmt.Printf("State: code:%d name:%s \n", aws.Int64Value(i.State.Code), aws.StringValue(i.State.Name))

	if len(i.Tags) > 0 {
		fmt.Print("*** This instance is associated with the following Tags. ***\n")
		for _, v := range i.Tags {
			fmt.Printf("%s:%s\n", aws.StringValue(v.Key), aws.StringValue(v.Value))
		}
	} else {
		fmt.Print("This instance is not associated with any Tags. \n")
	}
}

func HasTagName(tag string, i *_ec2.Instance) bool {
	if len(i.Tags) == 0 {
		return false
	}

	r := regexp.MustCompile(tag)
	for _, v := range i.Tags {
		if aws.StringValue(v.Key) == "Name" && r.MatchString(aws.StringValue(v.Value)) {
			return true
		}
	}

	return false
}

func (a *ec2) GetInstances(tag string) (instances, error) {

	var input *_ec2.DescribeInstancesInput

	if len(tag) > 0 {
		input = &_ec2.DescribeInstancesInput{
			Filters: []*_ec2.Filter{
				{
					Name:   aws.String("tag:Name"),
					Values: []*string{aws.String(tag)},
				},
			},
		}
	}

	result, err := a.svc.DescribeInstances(input)

	ins := instances{}

	if err == nil {
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
