package aws

import (
	_session "github.com/aws/aws-sdk-go/aws/session"
	"github.com/yuya-tajima/aws-go/aws/ec2"
)

type Aws struct {
	session *_session.Session
	Ec2     *ec2.Ec2
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

func (a *Aws) SetEc2Client() {
	a.Ec2 = ec2.NewEc2(a.session)
}
