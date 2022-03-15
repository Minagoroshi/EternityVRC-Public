package vrcAuth

import (
	"EternityGUI/shared"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func VerifyToken() string {
	url := shared.BaseURL + "auth"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Other")
	req.Header.Add("Cookie", "apiKey=JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26; auth="+shared.AuthToken)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	return string(body)
}
