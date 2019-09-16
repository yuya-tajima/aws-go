package main

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/yuya-tajima/aws-go/aws/s3"
	"github.com/yuya-tajima/aws-go/aws/util"
)

func showBucketDetails(i *s3.BucketItem) {
	fmt.Printf("Name: %s \n", i.Name)
	fmt.Printf("CreationDate: %s \n", i.Cdate)
	fmt.Print("\n")
}

func showObjectItems(i *s3.ObjectItem) {
	fmt.Printf("Key: %s \n", i.Key)
	fmt.Printf("ETag: %s \n", i.Etag)
	fmt.Printf("LastModified: %s \n", i.Mdate)
	fmt.Print("\n")
}

func showObjectDirs(d *s3.ObjectDir) {
	fmt.Printf("Name: %s \n", d.Name)
	fmt.Print("\n")
}

func setMetaData(c *cli.Context) {
	c.App.Metadata = map[string]interface{}{
		"isDry": c.GlobalBool("dryrun"),
		"tag":   c.GlobalString("tag"),
	}
}

func printError(etype errType, err error) {
	switch etype {
	case otherErr:
		util.Errorf("%v", err)
	}
}
