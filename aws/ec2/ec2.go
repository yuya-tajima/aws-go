package ec2

import (
	"fmt"
	"regexp"

	_aws "github.com/aws/aws-sdk-go/aws"
	_session "github.com/aws/aws-sdk-go/aws/session"
	_ec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Tags []Tag

type Item struct {
	InsID       string `json:"ins_id"`
	ImageID     string `json:"image_id"`
	PublicIpV4  string `json:"public_ipv4"`
	PrivateIpV4 string `json:"private_ipv4"`
	InsType     string `json:"ins_type"`
	StateCode   int64  `json:"status_code"`
	StateName   string `json:"status_name"`
	Tags        Tags   `json:"tags"`
}

type Data struct {
	Items []*Item `json:"items"`
}

type Ec2 struct {
	svc *_ec2.EC2
}

func NewEc2(session *_session.Session) *Ec2 {
	return &Ec2{
		svc: _ec2.New(session),
	}
}

func (a *Ec2) Reboot(s string, dry bool) (string, error) {

	input := &_ec2.RebootInstancesInput{
		InstanceIds: []*string{
			_aws.String(s),
		},
		DryRun: _aws.Bool(dry),
	}

	result, err := a.svc.RebootInstances(input)
	output := fmt.Sprintf("%s", result)

	return output, err
}

func (a *Ec2) Start(s string, dry bool) (string, error) {

	input := &_ec2.StartInstancesInput{
		InstanceIds: []*string{
			_aws.String(s),
		},
		DryRun: _aws.Bool(dry),
	}

	result, err := a.svc.StartInstances(input)
	output := fmt.Sprintf("%s", result)

	return output, err
}

func (a *Ec2) Stop(s string, dry bool) (string, error) {

	input := &_ec2.StopInstancesInput{
		InstanceIds: []*string{
			_aws.String(s),
		},
		DryRun: _aws.Bool(dry),
	}

	result, err := a.svc.StopInstances(input)
	output := fmt.Sprintf("%s", result)

	return output, err
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

func (a *Ec2) GetInstance(id string) (*Data, error) {

	result, err := a.DescInstanceById(id)

	var data *Data

	if err == nil {
		if len(result.Reservations) > 0 {
			data = &Data{}
			for _, r := range result.Reservations {
				for _, i := range r.Instances {
					item := createItem(i)
					data.Items = append(data.Items, item)
				}
			}
		}
	}

	return data, err
}

func (a *Ec2) GetInstances(tag string) (*Data, error) {

	result, err := a.DescInstances(tag)

	var data *Data

	if err == nil {
		if len(result.Reservations) > 0 {
			data = &Data{}
			for _, r := range result.Reservations {
				for _, i := range r.Instances {
					item := createItem(i)
					data.Items = append(data.Items, item)
				}
			}
		}
	}

	return data, err
}

func createItem(i *_ec2.Instance) *Item {

	item := Item{
		InsID:       _aws.StringValue(i.InstanceId),
		ImageID:     _aws.StringValue(i.ImageId),
		PublicIpV4:  _aws.StringValue(i.PublicIpAddress),
		PrivateIpV4: _aws.StringValue(i.PrivateIpAddress),
		InsType:     _aws.StringValue(i.InstanceType),
		StateCode:   _aws.Int64Value(i.State.Code),
		StateName:   _aws.StringValue(i.State.Name),
	}

	addTagField(i, &item)

	return &item
}

func addTagField(i *_ec2.Instance, item *Item) {
	if len(i.Tags) > 0 {
		var tags Tags
		for _, v := range i.Tags {
			tags = append(tags, Tag{
				Key:   _aws.StringValue(v.Key),
				Value: _aws.StringValue(v.Value),
			})
		}
		item.Tags = tags
	}
}

func (a *Ec2) DescInstanceById(id string) (*_ec2.DescribeInstancesOutput, error) {
	var input *_ec2.DescribeInstancesInput

	input = &_ec2.DescribeInstancesInput{
		InstanceIds: []*string{
			_aws.String(id),
		},
	}

	result, err := a.svc.DescribeInstances(input)

	return result, err
}

func (a *Ec2) DescInstances(tag string) (*_ec2.DescribeInstancesOutput, error) {
	var input *_ec2.DescribeInstancesInput

	if tag != "" {
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
