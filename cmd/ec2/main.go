package main

import (
	_ "fmt"
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

func getAws(d map[string]interface{}) *aws.Aws {
	return d["aws"].(*aws.Aws)
}

func getTag(d map[string]interface{}) string {
	return d["tag"].(string)
}

func isDryRun(d map[string]interface{}) bool {
	return d["isDry"].(bool)
}

func setMetaData(c *cli.Context, isDry bool, _aws *aws.Aws, tag string) {
	c.App.Metadata = map[string]interface{}{
		"isDry": isDry,
		"aws":   _aws,
		"tag":   tag,
	}
}

func main() {

	app := cli.NewApp()

	app.Name = "ec2"

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
			Usage: "value associated with Name tag",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  start,
			Usage: "",
			Action: func(c *cli.Context) error {
				_aws := getAws(c.App.Metadata)
				isDry := isDryRun(c.App.Metadata)
				result, err := _aws.GetInstances()
				if err != nil {
					aws.ExitErrorf("failed to get instances, profile '%s' %v.", _aws.GetProfile(), err)
				}
				tag := getTag(c.App.Metadata)
				if len(result) > 0 {
					for _, i := range result {
						if len(tag) == 0 || aws.HasTagName(tag, i) {
							_aws.Start(i.InstanceId, isDry)
						}
					}
				}
				return nil
			},
		},
		{
			Name:  stop,
			Usage: "",
			Action: func(c *cli.Context) error {
				_aws := getAws(c.App.Metadata)
				isDry := isDryRun(c.App.Metadata)
				result, err := _aws.GetInstances()
				if err != nil {
					aws.ExitErrorf("failed to get instances, profile '%s' %v.", _aws.GetProfile(), err)
				}
				tag := getTag(c.App.Metadata)
				if len(result) > 0 {
					for _, i := range result {
						if len(tag) == 0 || aws.HasTagName(tag, i) {
							_aws.Stop(i.InstanceId, isDry)
						}
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
				_aws := getAws(c.App.Metadata)
				isDry := isDryRun(c.App.Metadata)
				result, err := _aws.GetInstances()
				if err != nil {
					aws.ExitErrorf("failed to get instances, profile '%s' %v.", _aws.GetProfile(), err)
				}
				tag := getTag(c.App.Metadata)
				if len(result) > 0 {
					for _, i := range result {
						if len(tag) == 0 || aws.HasTagName(tag, i) {
							_aws.Reboot(i.InstanceId, isDry)
						}
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
				_aws := getAws(c.App.Metadata)
				result, err := _aws.GetInstances()
				if err != nil {
					aws.ExitErrorf("failed to get instances, profile '%s' %v.", _aws.GetProfile(), err)
				}
				tag := getTag(c.App.Metadata)
				if len(result) > 0 {
					for _, i := range result {
						if len(tag) == 0 || aws.HasTagName(tag, i) {
							aws.ShowDetails(i)
						}
					}
				} else {
					fmt.Printf("There is no instance.\n")
				}
				return nil
			},
		},
	}

	app.Before = func(c *cli.Context) error {
		_aws := &aws.Aws{}
		_aws.SetSession(c.GlobalString("profile"))
		_aws.SetClient()
		isDry := c.GlobalBool("dryrun")
		tag := c.GlobalString("tag")
		setMetaData(c, isDry, _aws, tag)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		//
	}

}
