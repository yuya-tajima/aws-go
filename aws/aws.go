package aws

import (
	_aws "github.com/aws/aws-sdk-go/aws"
	_credentials "github.com/aws/aws-sdk-go/aws/credentials"
	_session "github.com/aws/aws-sdk-go/aws/session"
	"github.com/yuya-tajima/aws-go/aws/ec2"
)

type Aws struct {
	session *_session.Session
	Ec2     *ec2.Ec2
	profile string
}

type Ec2Config struct {
	Region string
}

func (a *Aws) GetProfile() string {
	return a.profile
}

func (a *Aws) SetStaticSession(id, key, token, region string) {
	cred := _credentials.NewStaticCredentials(id, key, token)

	sess := _session.Must(_session.NewSession(
		&_aws.Config{
			Credentials: cred,
			Region:      _aws.String(region),
		},
	))

	a.session = sess
	a.profile = ""
}

func (a *Aws) SetSession(profile string) {

	sess := _session.Must(_session.NewSessionWithOptions(_session.Options{
		SharedConfigState: _session.SharedConfigEnable,
		Profile:           profile,
	}))

	a.session = sess
	a.profile = profile
}

func (a *Aws) SetEc2Client(config *Ec2Config) {

	cfg := _aws.NewConfig()
	if config != nil {
		if config.Region != "" {
			cfg.WithRegion(config.Region)
		}
	}

	a.Ec2 = ec2.NewEc2(a.session, cfg)
}
