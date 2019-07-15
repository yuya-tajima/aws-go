package main

import (
	"fmt"
	"github.com/yuya-tajima/aws-go/httpd/app/auth"
	"log"
)

const adminUser = "admin"
const Profile = "aws_api"

func init() {
	err := auth.SetCredByEnv()
	if err != nil {
		err := auth.SetCredByFile(Profile)
		if err != nil {
			log.Fatal(err)
		}
	}

	auth.Init()

	password, err := auth.NewUser(adminUser)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("id:%s password:%s", adminUser, password)
}
