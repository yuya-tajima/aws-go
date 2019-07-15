package client

import (
	_"fmt"
	"net/http"
	"io/ioutil"
	"bytes"

	"github.com/yuya-tajima/aws-go/httpd/app/auth"
	"github.com/yuya-tajima/aws-go/httpd/app/util"
)

func setRegionHeader (req *http.Request) {
	req.Header.Set("Region", "ap-northeast-1")
}

func setAuthHeader (req *http.Request) {
	cred := auth.GetCred()
	req.Header.Set("Secret_Id", cred.Access_key)
	req.Header.Set("Secret_key", cred.Secret_key)
}

func setJsonHeader (req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
}


func PostRequest(url string, json []byte) ([]byte, int) {

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(json),
	)

	if err != nil {
		return util.InternalErrorJSON(err), http.StatusInternalServerError
	} else {

		setRegionHeader(req)
		setAuthHeader(req)
		setJsonHeader(req)

		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			return util.InternalErrorJSON(err), http.StatusInternalServerError
		}

		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		return body, resp.StatusCode
	}
}

func GetRequest(url string) ([]byte, int) {

	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)

	if err != nil {
		return util.InternalErrorJSON(err), http.StatusInternalServerError
	} else {

		setRegionHeader(req)
		setAuthHeader(req)
		setJsonHeader(req)

		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			return util.InternalErrorJSON(err), http.StatusInternalServerError
		}

		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		return body, resp.StatusCode
	}
}
