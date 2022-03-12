package vrchat

import (
	"EternityGUI/cmd/requests"
	"EternityGUI/crypto"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func AvatarSearch(AvatarName string) string {

	log.Println()
	url := "https://eternitybots.ml/AvatarSearchName"
	method := "POST"
	hash := crypto.CompHash("AvatarSearch")
	jsonToMarshal := &requests.EternityPost{AvatarName: AvatarName, Username: "Tope", Sha3: hash}
	data, _ := json.Marshal(jsonToMarshal)
	payload := strings.NewReader(string(data))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Println(err)

	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)

	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)

	}

	return string(body)

}
