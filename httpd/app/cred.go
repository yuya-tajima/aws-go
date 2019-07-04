package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/credentials"
)

func setCredByEnv() error {
	creds := credentials.NewEnvCredentials()
	credValue, err := creds.Get()
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	access_key = credValue.AccessKeyID
	secret_key = credValue.SecretAccessKey

	return nil
}

func setCredByFile() error {

	creds := credentials.NewSharedCredentials("", profileName)
	credValue, err := creds.Get()
	if err != nil {
		return fmt.Errorf("%s (tried to find '%s')\n", err, profileName)
	}

	access_key = credValue.AccessKeyID
	secret_key = credValue.SecretAccessKey

	return nil
}
