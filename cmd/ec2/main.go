package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/yuya-tajima/aws-go/aws"
	"github.com/yuya-tajima/aws-go/aws/util"
)

type errType int

const (
	getErr errType = iota
	startErr
	stopErr
	rebootErr
	otherErr
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
			Name:   start,
			Usage:  "",
			Action: startAction,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "instance-id, i",
					Usage: "specific instance-id",
				},
			},
		},
		{
			Name:   stop,
			Usage:  "",
			Action: stopAction,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "instance-id, i",
					Usage: "specific instance-id",
				},
			},
		},
		{
			Name:   reboot,
			Usage:  "",
			Action: rebootAction,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "instance-id, i",
					Usage: "specific instance-id",
				},
			},
		},
		{
			Name:   desc,
			Usage:  "",
			Action: descAction,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "instance-id, i",
					Usage: "specific instance-id",
				},
			},
		},
	}

	app.Before = func(c *cli.Context) error {
		_aws = &aws.Aws{}
		_aws.SetSession(c.GlobalString("profile"))
		_aws.SetEc2Client(nil)

		setMetaData(c)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		util.ExitErrorf("%v.", err)
	}
}
