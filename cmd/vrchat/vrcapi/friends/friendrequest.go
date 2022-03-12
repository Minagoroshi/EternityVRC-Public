package vrcfriends

import (
	"EternityGUI/cmd/vrchat/vrcapi"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

//SendFriendRequest sends a friend request to the given user ID
func SendFriendRequest(uid string) string {
	url := vrcapi.BaseURL + "user/" + uid + "/friendRequest"
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Other")
	req.Header.Add("Cookie", "apiKey=JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26; auth="+vrcapi.AuthToken)
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
	fmt.Println(req.URL)

	return string(body)

}
