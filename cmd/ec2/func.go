package main

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/yuya-tajima/aws-go/aws/ec2"
	"github.com/yuya-tajima/aws-go/aws/util"
)

func showDetails(i *ec2.Item) {
	fmt.Printf("InstanceId: %s \n", i.InsID)
	fmt.Printf("ImageId: %s \n", i.ImageID)
	fmt.Printf("InstanceType: %s \n", i.InsType)
	fmt.Printf("PrivateIpAddress: %s \n", i.PrivateIpV4)
	fmt.Printf("PublicIpAddress: %s \n", i.PublicIpV4)
	fmt.Printf("State: code:%d name:%s \n", i.StateCode, i.StateName)

	if len(i.Tags) > 0 {
		fmt.Print("*** This instance is associated with the following Tags. ***\n")
		for _, v := range i.Tags {
			fmt.Printf("%s:%v\n", v.Key, v.Value)
		}
	} else {
		fmt.Print("This instance is not associated with any Tags. \n")
	}
}

func printNoInstance() {
	fmt.Printf("There is no instance.\n")
}

func print(s string) {
	if s != "" {
		fmt.Print(s)
	}
}

func printError(etype errType, err error) {
	switch etype {
	case getErr:
		util.Errorf("failed to get instances, profile '%s' %v", _aws.GetProfile(), err)
	case otherErr:
		util.Errorf("%v", err)
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
