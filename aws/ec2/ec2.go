package ec2

import (
	"fmt"
	"regexp"

	_aws "github.com/aws/aws-sdk-go/aws"
	_session "github.com/aws/aws-sdk-go/aws/session"
	_ec2 "github.com/aws/aws-sdk-go/service/ec2"
	"github.com/yuya-tajima/aws-go/aws/util"
)

type Item struct {
	InsID       string `json:"ins_id"`
	ImageID     string `json:"image_id"`
	PublicIpV4  string `json:"public_ipv4"`
	PrivateIpV4 string `json:"private_ipv4"`
	InsType     string `json:"ins_type"`
	StateCode   int64  `json:"status_code"`
	StateName   string `json:"status_name"`
}

type Data struct {
	Items []Item `json:"items"`
}

type Ec2 struct {
	svc *_ec2.EC2
}

func NewEc2(session *_session.Session) *Ec2 {
	return &Ec2{
		svc: _ec2.New(session),
	}
}

func (a *Ec2) Reboot(s *string, dry bool) {

	input := &_ec2.RebootInstancesInput{
		InstanceIds: []*string{
			s,
		},
		DryRun: _aws.Bool(dry),
	}

	result, err := a.svc.RebootInstances(input)

	util.MaybeExitError(err)

	fmt.Println(result)
}

func (a *Ec2) Start(s *string, dry bool) {

	input := &_ec2.StartInstancesInput{
		InstanceIds: []*string{
			s,
		},
		DryRun: _aws.Bool(dry),
	}

	result, err := a.svc.StartInstances(input)

	util.MaybeExitError(err)

	fmt.Println(result)
}

func (a *Ec2) Stop(s *string, dry bool) {

	input := &_ec2.StopInstancesInput{
		InstanceIds: []*string{
			s,
		},
		DryRun: _aws.Bool(dry),
	}

	result, err := a.svc.StopInstances(input)

	util.MaybeError(err)

	fmt.Println(result)
}

func ShowDetails(i *_ec2.Instance) {
	fmt.Printf("InstanceId: %s \n", _aws.StringValue(i.InstanceId))
	fmt.Printf("ImageId: %s \n", _aws.StringValue(i.ImageId))
	fmt.Printf("InstanceType: %s \n", _aws.StringValue(i.InstanceType))
	fmt.Printf("PrivateIpAddress: %s \n", _aws.StringValue(i.PrivateIpAddress))
	fmt.Printf("PublicIpAddress: %s \n", _aws.StringValue(i.PublicIpAddress))
	fmt.Printf("State: code:%d name:%s \n", _aws.Int64Value(i.State.Code), _aws.StringValue(i.State.Name))

	if len(i.Tags) > 0 {
		fmt.Print("*** This instance is associated with the following Tags. ***\n")
		for _, v := range i.Tags {
			fmt.Printf("%s:%s\n", _aws.StringValue(v.Key), _aws.StringValue(v.Value))
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
		if _aws.StringValue(v.Key) == "Name" && r.MatchString(_aws.StringValue(v.Value)) {
			return true
		}
	}

	return false
}

func (a *Ec2) GetInstances(tag string) (*Data, error) {

	result, err := a.DescInstances(tag)

	var data *Data

	if err == nil {
		if len(result.Reservations) > 0 {
			data = &Data{}
			for _, r := range result.Reservations {
				for _, i := range r.Instances {
					if len(i.Tags) > 0 {
						fmt.Print("*** This instance is associated with the following Tags. ***\n")
						for _, v := range i.Tags {
							fmt.Printf("%s:%s\n", _aws.StringValue(v.Key), _aws.StringValue(v.Value))
						}
					} else {
						fmt.Print("This instance is not associated with any Tags. \n")
					}
					data.Items = append(data.Items, Item{
						InsID:       _aws.StringValue(i.InstanceId),
						ImageID:     _aws.StringValue(i.ImageId),
						PublicIpV4:  _aws.StringValue(i.PublicIpAddress),
						PrivateIpV4: _aws.StringValue(i.PrivateIpAddress),
						InsType:     _aws.StringValue(i.InstanceType),
						StateCode:   _aws.Int64Value(i.State.Code),
						StateName:   _aws.StringValue(i.State.Name),
					})
				}
			}
		}
	}

	return data, err

}

func (a *Ec2) DescInstances(tag string) (*_ec2.DescribeInstancesOutput, error) {
	var input *_ec2.DescribeInstancesInput

	if len(tag) > 0 {
		input = &_ec2.DescribeInstancesInput{
			Filters: []*_ec2.Filter{
				{
					Name:   _aws.String("tag:Name"),
					Values: []*string{_aws.String(tag)},
				},
			},
		}
	}

	result, err := a.svc.DescribeInstances(input)

	return result, err
}

/*
func (a *ec2) GetInstances(tag string) (instances, error) {

	result, err := a.DescInstances(tag)

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
*/
