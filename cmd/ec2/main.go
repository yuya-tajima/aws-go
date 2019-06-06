package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/yuya-tajima/aws-go/aws"
)

const (
	start  = "start"
	stop   = "stop"
	reboot = "reboot"
	desc   = "desc"
)

var _aws *aws.Aws

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

func main() {

	app := cli.NewApp()

	app.Name = "ec2"
	app.Usage = "you can simply manage your ec2"
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
				result, err := _aws.GetInstances(tag)
				if err != nil {
					aws.ExitErrorf("failed to get instances, profile '%s' %v.", _aws.GetProfile(), err)
				}
				if len(result) > 0 {
					for _, i := range result {
						_aws.Start(i.InstanceId, isDry)
					}
				}
				return nil
			},
		},
		{
			Name:  stop,
			Usage: "",
			Action: func(c *cli.Context) error {
				isDry := isDryRun(c)
				tag := getNameTag(c)
				result, err := _aws.GetInstances(tag)
				if err != nil {
					aws.ExitErrorf("failed to get instances, profile '%s' %v.", _aws.GetProfile(), err)
				}
				if len(result) > 0 {
					for _, i := range result {
						_aws.Stop(i.InstanceId, isDry)
					}
				} else {
					fmt.Printf("There is no instance.\n")
				}
				return nil
			},
		},
		{
			Name:  reboot,
			Usage: "",
			Action: func(c *cli.Context) error {
				isDry := isDryRun(c)
				tag := getNameTag(c)
				result, err := _aws.GetInstances(tag)
				if err != nil {
					aws.ExitErrorf("failed to get instances, profile '%s' %v.", _aws.GetProfile(), err)
				}
				if len(result) > 0 {
					for _, i := range result {
						_aws.Reboot(i.InstanceId, isDry)
					}
				} else {
					fmt.Printf("There is no instance.\n")
				}
				return nil
			},
		},
		{
			Name:  desc,
			Usage: "",
			Action: func(c *cli.Context) error {
				tag := getNameTag(c)
				result, err := _aws.GetInstances(tag)
				if err != nil {
					aws.ExitErrorf("failed to get instances, profile '%s' %v.", _aws.GetProfile(), err)
				}
				if len(result) > 0 {
					for _, i := range result {
						aws.ShowDetails(i)
					}
				} else {
					fmt.Printf("There is no instance.\n")
				}
				return nil
			},
		},
	}

	app.Before = func(c *cli.Context) error {
		_aws = &aws.Aws{}
		_aws.SetSession(c.GlobalString("profile"))
		_aws.SetClient()
		setMetaData(c)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		aws.ExitErrorf("%v.", err)
	}
}
