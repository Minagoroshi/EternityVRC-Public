package vrcfriends

import (
	"EternityGUI/cmd/vrchat"
	"EternityGUI/shared"
	"EternityGUI/utils"
	"github.com/buger/jsonparser"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func ListFriends() (string, bool) {
	config := utils.ReadConfig()
	token := config.VRChatLogin
	url := shared.BaseURL + "auth/user/friends?offset=0&n=100&offline=false"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "IFartedAndPooped")
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
	var b bool
	if strings.Contains(string(body), "currentAvatarImageUrl") {
		b = true
	}
	return string(body), b
}

func MakeFriendList() ([]string, []string, []string, []string, []string, []string, int) {
	fbody, b := ListFriends()
	fcount := strings.Count(fbody, "currentAvatarThumbnailImageUrl")
	fdata, fidData, fImageURLData, fstatusData, fstatusDescriptionData, ftrustData := make([]string, fcount), make([]string, fcount),
		make([]string, fcount), make([]string, fcount), make([]string, fcount), make([]string, fcount)
	if b {
		log.Println("Friend List Results", fcount)

		for i := range fdata {
			n := "[" + strconv.Itoa(i) + "]"
			fResult, _ := jsonparser.GetString([]byte(fbody), n, "displayName")
			fdata[i] = fResult
			resultID, _ := jsonparser.GetString([]byte(fbody), n, "id")
			fidData[i] = "User ID: " + resultID
			resultImageUrl, _, _, _ := jsonparser.Get([]byte(fbody), n, "currentAvatarImageUrl")
			fImageURLData[i] = string(resultImageUrl)
			resultStatus, _ := jsonparser.GetString([]byte(fbody), n, "status")
			fstatusData[i] = "Status: " + string(resultStatus)
			resultStatusDescription, _ := jsonparser.GetString([]byte(fbody), n, "statusDescription")
			fstatusDescriptionData[i] = "Status Description: " + string(resultStatusDescription)
			resultTags, _, _, _ := jsonparser.Get([]byte(fbody), n, "tags")
			resultTrust := vrchat.TagsConverter(resultTags)
			ftrustData[i] = "Trust: " + resultTrust
			////resultBio, _, _, _ := jsonparser.Get([]byte(body), n, "bio")
			////tagsData[i] =  resultTags
		}
	} /*else {
		fdata = make([]string, 1)
		fcount = 0
		for i := range fdata {
			fdata[i], fidData[i], fImageURLData[i], fstatusData[i], fstatusDescriptionData[i], ftrustData[i] =
				"You are not logged in.", "", "", "", "", ""
		}

	}*/
	return fdata, fidData, fImageURLData, fstatusData, fstatusDescriptionData, ftrustData, fcount
}
