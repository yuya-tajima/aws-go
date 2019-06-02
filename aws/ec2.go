package aws

import (
	"fmt"

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
