package util

import (
	"fmt"
	"os"

	_awserr "github.com/aws/aws-sdk-go/aws/awserr"
)

func Errorf(msg string, args ...interface{}) {
	_errorf(msg, args)
}

func ExitErrorf(msg string, args ...interface{}) {
	_errorf(msg, args)
	os.Exit(1)
}

func _errorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
}

func _error(err error) {
	if aerr, ok := err.(_awserr.Error); ok {
		switch aerr.Code() {
		default:
			_errorf(aerr.Error())
		}
	} else {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		_errorf(err.Error())
	}
}
