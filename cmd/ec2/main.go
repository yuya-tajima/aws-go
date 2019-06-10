package main

import (
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
			Name:  start,
			Usage: "",
			Action: func(c *cli.Context) error {
				isDry := isDryRun(c)
				tag := getNameTag(c)
				id := c.String("instance-id")

				if id != "" {
					output, err := _aws.Ec2.Start(id, isDry)
					if err != nil {
						printError(otherErr, err)
					} else {
						print(output)
					}
				} else {
					result, err := _aws.Ec2.GetInstances(tag)
					if err != nil {
						printError(getErr, err)
					} else {
						if result != nil {
							for _, i := range result.Items {
								output, err := _aws.Ec2.Start(i.InsID, isDry)
								if err != nil {
									printError(otherErr, err)
								}
								print(output)
							}
						} else {
							printNoInstance()
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

				if id != "" {
					output, err := _aws.Ec2.Stop(id, isDry)
					if err != nil {
						printError(otherErr, err)
					} else {
						print(output)
					}
				} else {
					result, err := _aws.Ec2.GetInstances(tag)
					if err != nil {
						printError(getErr, err)
					} else {
						if result != nil {
							for _, i := range result.Items {
								output, err := _aws.Ec2.Stop(i.InsID, isDry)
								if err != nil {
									printError(otherErr, err)
								}
								print(output)
							}
						} else {
							printNoInstance()
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
			Name:  reboot,
			Usage: "",
			Action: func(c *cli.Context) error {
				isDry := isDryRun(c)
				tag := getNameTag(c)
				id := c.String("instance-id")
				if id != "" {
					output, err := _aws.Ec2.Reboot(id, isDry)
					if err != nil {
						printError(otherErr, err)
					} else {
						print(output)
					}
				} else {
					result, err := _aws.Ec2.GetInstances(tag)
					if err != nil {
						printError(getErr, err)
					} else {
						if result != nil {
							for _, i := range result.Items {
								output, err := _aws.Ec2.Reboot(i.InsID, isDry)
								if err != nil {
									printError(otherErr, err)
								}
								print(output)
							}
						} else {
							printNoInstance()
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
			Name:  desc,
			Usage: "",
			Action: func(c *cli.Context) error {

				tag := getNameTag(c)
				id := c.String("instance-id")

				var result *ec2.Data
				var err error

				if id != "" {
					result, err = _aws.Ec2.GetInstance(id)
				} else {
					result, err = _aws.Ec2.GetInstances(tag)
				}

				if err != nil {
					printError(getErr, err)
				} else {
					if result != nil {
						for _, i := range result.Items {
							showDetails(i)
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
