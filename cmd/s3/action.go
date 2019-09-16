package main

import (
	_ "fmt"

	"github.com/urfave/cli"
	_ "github.com/yuya-tajima/aws-go/aws/s3"
)

func listBucketAction(c *cli.Context) (err error) {

	result, err := _aws.S3.GetBuckets("")
	if err != nil {
		printError(otherErr, err)
	} else {
		if result != nil {
			for _, i := range result.Items {
				showBucketDetails(i)
			}
		} else {
			//
		}
	}

	return nil
}

func listObjectAction(c *cli.Context) (err error) {

	b := c.String("bucket-name")
	p := c.String("path-name")

	result, err := _aws.S3.GetObjects(b, p)
	if err != nil {
		printError(otherErr, err)
	} else {
		if result != nil {
			for _, i := range result.Items {
				showObjectItems(i)
			}
			for _, d := range result.Dirs {
				showObjectDirs(d)
			}
		} else {
			//
		}
	}

	return nil
}
