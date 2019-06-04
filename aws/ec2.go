package aws

import (
	"fmt"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type instances []*ec2.Instance

func (a *Aws) Reboot(s *string, dry bool) {

	input := &ec2.RebootInstancesInput{
		InstanceIds: []*string{
			s,
		},
		DryRun: aws.Bool(dry),
	}

	result, err := a.ec2.RebootInstances(input)

	MaybeError(err)

	fmt.Println(result)
}

func (a *Aws) Start(s *string, dry bool) {

	input := &ec2.StartInstancesInput{
		InstanceIds: []*string{
			s,
		},
		DryRun: aws.Bool(dry),
	}

	result, err := a.ec2.StartInstances(input)

	MaybeError(err)

	fmt.Println(result)
}

func (a *Aws) Stop(s *string, dry bool) {

	input := &ec2.StopInstancesInput{
		InstanceIds: []*string{
			s,
		},
		DryRun: aws.Bool(dry),
	}

	result, err := a.ec2.StopInstances(input)

	MaybeError(err)

	fmt.Println(result)
}

func (a *Aws) SetClient() {
	a.ec2 = ec2.New(a.session)
}

func ShowDetails(i *ec2.Instance) {
	fmt.Printf("InstanceId: %s \n", aws.StringValue(i.InstanceId))
	fmt.Printf("ImageId: %s \n", aws.StringValue(i.ImageId))
	fmt.Printf("InstanceType: %s \n", aws.StringValue(i.InstanceType))
	fmt.Printf("PrivateIpAddress: %s \n", aws.StringValue(i.PrivateIpAddress))
	fmt.Printf("PublicIpAddress: %s \n", aws.StringValue(i.PublicIpAddress))
	fmt.Printf("State: code:%d name:%s \n", aws.Int64Value(i.State.Code), aws.StringValue(i.State.Name))

	if len(i.Tags) > 0 {
		fmt.Print("*** The following tags are defined ***\n")
		for _, v := range i.Tags {
			fmt.Printf("%s:%s\n", aws.StringValue(v.Key), aws.StringValue(v.Value))
		}
	}
}

func HasTagName(tag string, i *ec2.Instance) bool {
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
