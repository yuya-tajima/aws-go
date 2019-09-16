package main

import (
	"fmt"
	"github.com/urfave/cli"
	"github.com/yuya-tajima/aws-go/aws"
	"github.com/yuya-tajima/aws-go/aws/util"
	"os"
)

type errType int

const (
	otherErr errType = iota
)

const (
	listBucket = "ls-bucket"
	listObject = "ls"
)

var _aws *aws.Aws

func main() {

	app := cli.NewApp()

	app.Name = "s3"
	app.Usage = "you can simply manage your s3 buckets and objects"
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
			Name:   listBucket,
			Usage:  "list buckets",
			Action: listBucketAction,
		},
		{
			Name:   listObject,
			Usage:  "list objects",
			Action: listObjectAction,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "bucket-name, b",
					Usage: "specific bucket name",
				},
				cli.StringFlag{
					Name:  "path-name, p",
					Usage: "specific path name",
				},
			},
			Before: func(c *cli.Context) error {
				b := c.String("bucket-name")
				if b == "" {
					return fmt.Errorf("--bucket-name is required")
				}
				return nil
			},
		},
	}

	app.Before = func(c *cli.Context) error {
		_aws = &aws.Aws{}
		_aws.SetSession(c.GlobalString("profile"))
		_aws.SetS3Client(nil)

		setMetaData(c)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		util.ExitErrorf("%v.", err)
	}

}
