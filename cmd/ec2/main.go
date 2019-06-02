package main

import (
	"fmt"
	"os"
	"regexp"

	. "github.com/yuya-tajima/aws-go/aws"
)

func main() {

	if len(os.Args) < 3 {
		ExitErrorf("Usage: ec2 profile {start|stop}")
	}

	profile := os.Args[1]

	action := os.Args[2]
	if action != "start" && action != "stop" {
		ExitErrorf("Usage: ec2 {start|stop}")
	}

	_aws := Aws{}

	_aws.SetSession(profile)
	_aws.SetClient()

	result, err := _aws.GetInstances()

	if err != nil {
		ExitErrorf("failed to get instances, profile '%v'  %v.", profile, err)
	}

	if len(result) > 0 {
		for _, i := range result {
			for _, v := range i.Tags {
				r := regexp.MustCompile(`test`)
				if *v.Key == "Name" && r.MatchString(*v.Value) {
					switch action {
					case "start":
						_aws.Start(*i.InstanceId)
					case "stop":
						_aws.Stop(*i.InstanceId)
					}
				}
			}
		}
	} else {
		fmt.Printf("There is no instance.\n")
	}

}
