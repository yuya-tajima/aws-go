package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/yuya-tajima/aws-go/aws"
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
					_aws.Start(&id, isDry)
				} else {
					result, err := _aws.GetInstances(tag)
					if err != nil {
						printExitError(getErr, err)
					}
					if len(result) > 0 {
						for _, i := range result {
							_aws.Start(i.InstanceId, isDry)
						}
					}
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
					_aws.Stop(&id, isDry)
				} else {
					result, err := _aws.GetInstances(tag)
					if err != nil {
						printExitError(getErr, err)
					}
					if len(result) > 0 {
						for _, i := range result {
							_aws.Stop(i.InstanceId, isDry)
						}
					} else {
						printNoInstance()
					}

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
					_aws.Reboot(&id, isDry)
				} else {
					result, err := _aws.GetInstances(tag)
					if err != nil {
						printExitError(getErr, err)
					}
					if len(result) > 0 {
						for _, i := range result {
							_aws.Reboot(i.InstanceId, isDry)
						}
					} else {
						printNoInstance()
					}
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
				result, err := _aws.GetInstances(tag)
				if err != nil {
					printExitError(getErr, err)
				}
				if len(result) > 0 {
					for _, i := range result {
						aws.ShowDetails(i)
					}
				} else {
					printNoInstance()
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
		aws.ExitErrorf("%v.", err)
	}
}

func printNoInstance() {
	fmt.Printf("There is no instance.\n")
}

func printExitError(etype errType, err error) {
	switch etype {
	case getErr:
		aws.ExitErrorf("failed to get instances, profile '%s' %v.", _aws.GetProfile(), err)
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
