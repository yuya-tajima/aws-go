package main

import (
	"github.com/urfave/cli"
	"github.com/yuya-tajima/aws-go/aws/ec2"
)

func startAction(c *cli.Context) error {

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
}

func stopAction(c *cli.Context) error {

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
}

func rebootAction(c *cli.Context) error {

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
}

func descAction(c *cli.Context) error {

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
}
