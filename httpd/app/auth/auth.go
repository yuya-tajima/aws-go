package auth

import (
	crand "crypto/rand"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"golang.org/x/crypto/bcrypt"
	"math"
	"math/big"
	"math/rand"
)

type userMeta struct {
	hash string
}

type users map[string]*userMeta

type creds struct {
	Access_key string
	Secret_key string
}

var (
	user users
	aws *creds
)

type Login struct {
	Id string `json:"login"`
	Pwd string `json:"password"`
}

const passLen = 12

func Init() {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
}

func AuthUser (id string, pass string) bool {
	if um, err := getUserMeta(id); err == nil {
		if err := um.passwordVerify(pass); err == nil {
			return true
		} else {
			return false
		}
	}
	return false
}

func GenerateRandom(n int) string {
	const chars = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (um *userMeta) passwordVerify(pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(um.hash), []byte(pw))
}

func NewUser (id string) (string ,error){

	password := GenerateRandom(passLen)
	hash, err := hashPassword(password)

	if err != nil {
		return "", err
	}

	um := &userMeta{
		hash: hash,
	}

	setUser(id, um)

	return password, nil
}

func setUser (id string, um *userMeta) {
	if user == nil {
		user = users{id:um}
	} else {
		user[id] = um
	}
}

func getUserMeta (id string) (*userMeta, error) {
	if um, ok := user[id]; ok {
		return um, nil
	} else {
		return nil, fmt.Errorf("%s does not exist)\n", id)
	}
}

func setCred(access_key, secret_key string) {
	aws = &creds{
		Access_key: access_key,
		Secret_key: secret_key,
	}
}

func GetCred() *creds {
	return aws
}

func SetCredByEnv() error {
	creds := credentials.NewEnvCredentials()
	credValue, err := creds.Get()
	if err != nil {
		return err
	}

	setCred(credValue.AccessKeyID, credValue.SecretAccessKey)

	return nil
}

func SetCredByFile(profile string) error {

	creds := credentials.NewSharedCredentials("", profile)
	credValue, err := creds.Get()
	if err != nil {
		return fmt.Errorf("%s (tried to find '%s')\n", err, profile)
	}

	setCred(credValue.AccessKeyID, credValue.SecretAccessKey)

	return nil
}

