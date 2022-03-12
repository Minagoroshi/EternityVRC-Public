package vrcworlds

import (
	"EternityGUI/cmd/vrchat/vrcapi"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func WorldSearch(world string, token string) string {

	url := vrcapi.BaseURL + "worlds?sort=popularity&n=100&order=descending&offset=0&search=" + world
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Other")
	req.Header.Add("Cookie", "apiKey=JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26; auth="+token)
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
