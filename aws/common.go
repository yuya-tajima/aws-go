package aws

import (
	_session "github.com/aws/aws-sdk-go/aws/session"
)

type Aws struct {
	session *_session.Session
	*ec2
	profile string
}

func (a *Aws) GetProfile() string {
	return a.profile
}

func (a *Aws) SetSession(profile string) {

	sess := _session.Must(_session.NewSessionWithOptions(_session.Options{
		SharedConfigState: _session.SharedConfigEnable,
		Profile:           profile,
	}))

	a.session = sess
	a.profile = profile
}
