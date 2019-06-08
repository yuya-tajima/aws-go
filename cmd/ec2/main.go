package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/yuya-tajima/aws-go/aws"
	"github.com/yuya-tajima/aws-go/aws/ec2"
	"github.com/yuya-tajima/aws-go/aws/util"
)

type errType int

const (
	getErr errType = iota
	startErr
	stopErr
	rebootErr
)
const (
	start  = "start"
	stop   = "stop"
	reboot = "reboot"
	desc   = "desc"
)

var _aws *aws.Aws

func main() {

	app := cli.NewApp()

	app.Name = "ec2"
	app.Usage = "you can simply manage your ec2 instances"
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "profile, p",
			Value:  "default",
			Usage:  "specific named profile",
			EnvVar: "AWS_DEFAULT_PROFILE",
		},
		cli.BoolFlag{
			Name:  "dryrun, d",
			Usage: "perform a trial run",
		},
		cli.StringFlag{
			Name:  "tag, t",
			Usage: "Name tag value associated with Name tag",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  start,
			Usage: "",
			Action: func(c *cli.Context) error {
				isDry := isDryRun(c)
				tag := getNameTag(c)
				id := c.String("instance-id")

				if len(id) > 0 {
					_aws.Ec2.Start(&id, isDry)
				} else {
					result, err := _aws.Ec2.GetInstances(tag)
					if err != nil {
						printExitError(getErr, err)
					}
                    fmt.Printf("%v\n", result)
                    /*
					if len(result) > 0 {
						for _, i := range result {
							_aws.Ec2.Start(i.InstanceId, isDry)
						}
					}
                    */
				}

				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "instance-id, id",
					Usage: "specific instance-id",
				},
			},
		},
		{
			Name:  stop,
			Usage: "",
			Action: func(c *cli.Context) error {
				isDry := isDryRun(c)
				tag := getNameTag(c)
				id := c.String("instance-id")
				if len(id) > 0 {
					_aws.Ec2.Stop(&id, isDry)
				} else {
					result, err := _aws.Ec2.GetInstances(tag)
					if err != nil {
						printExitError(getErr, err)
					}
                    fmt.Printf("%v\n", result)
                    /*
                    fmt,Printf("%v\n", result)
					if len(result) > 0 {
						for _, i := range result {
							_aws.Ec2.Stop(i.InstanceId, isDry)
						}
					} else {
						printNoInstance()
					}
                    */

				}
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "instance-id, id",
					Usage: "specific instance-id",
				},
			},
		},
		{
			Name:  reboot,
			Usage: "",
			Action: func(c *cli.Context) error {
				isDry := isDryRun(c)
				tag := getNameTag(c)
				id := c.String("instance-id")
				if len(id) > 0 {
					_aws.Ec2.Reboot(&id, isDry)
				} else {
					result, err := _aws.Ec2.GetInstances(tag)
					if err != nil {
						printExitError(getErr, err)
					}
                    fmt.Printf("%v\n", result)
                    /*
					if len(result) > 0 {
						for _, i := range result {
							_aws.Ec2.Reboot(i.InstanceId, isDry)
						}
					} else {
						printNoInstance()
					}
                    */
				}

				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "instance-id, id",
					Usage: "specific instance-id",
				},
			},
		},
		{
			Name:  desc,
			Usage: "",
			Action: func(c *cli.Context) error {
				tag := getNameTag(c)

				result, err := _aws.Ec2.GetInstances(tag)
				if err != nil {
					printExitError(getErr, err)
				}

				if result != nil {
                    for _, v := range result.Items {
                        showDetails(v)
                    }
				}
				return nil
			},
		},
	}

	app.Before = func(c *cli.Context) error {
		_aws = &aws.Aws{}
		_aws.SetSession(c.GlobalString("profile"))
		_aws.SetEc2Client()
		setMetaData(c)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		util.ExitErrorf("%v.", err)
	}
}

func showDetails(i ec2.Item) {
	fmt.Printf("InstanceId: %s \n", i.InsID)
	fmt.Printf("ImageId: %s \n", i.ImageID)
	fmt.Printf("InstanceType: %s \n", i.InsType)
	fmt.Printf("PrivateIpAddress: %s \n", i.PrivateIpV4)
	fmt.Printf("PublicIpAddress: %s \n", i.PublicIpV4)
	fmt.Printf("State: code:%d name:%s \n", i.StateCode, i.StateName)
}

func printNoInstance() {
	fmt.Printf("There is no instance.\n")
}

func printExitError(etype errType, err error) {
	switch etype {
	case getErr:
		util.ExitErrorf("failed to get instances, profile '%s' %v.", _aws.GetProfile(), err)
	}
}

func getNameTag(c *cli.Context) string {
	return c.App.Metadata["tag"].(string)
}

func isDryRun(c *cli.Context) bool {
	return c.App.Metadata["isDry"].(bool)
}

func setMetaData(c *cli.Context) {
	c.App.Metadata = map[string]interface{}{
		"isDry": c.GlobalBool("dryrun"),
		"tag":   c.GlobalString("tag"),
	}
}
