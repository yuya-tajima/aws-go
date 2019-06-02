package main

import (
	"fmt"
	"os"

	. "github.com/yuya-tajima/aws-go/aws"
)

func usage() {
	ExitErrorf("Usage: ec2 profile {start|stop|desc [tag]}")
}

func main() {

	if len(os.Args) < 3 {
		usage()
	}

	profile := os.Args[1]

	action := os.Args[2]
	if action != "start" && action != "stop" && action != "desc" {
		usage()
	}

	tag := ``
	if len(os.Args) > 3 {
		tag = fmt.Sprintf(`%s`, os.Args[3])
	}

	_aws := Aws{}

	_aws.SetSession(profile)
	_aws.SetClient()

	result, err := _aws.GetInstances()

	if err != nil {
		ExitErrorf("failed to get instances, profile '%s'  %v.", profile, err)
	}

	if len(result) > 0 {
		for _, i := range result {
			if len(tag) == 0 || HasTagName(tag, i) {
				switch action {
				case "start":
					_aws.Start(*i.InstanceId)
				case "stop":
					_aws.Stop(*i.InstanceId)
				case "desc":
					ShowDetails(i)
				}
			}
		}
	} else {
		fmt.Printf("There is no instance.\n")
	}

}
