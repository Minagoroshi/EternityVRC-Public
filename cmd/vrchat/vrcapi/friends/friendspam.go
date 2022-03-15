package vrcfriends

import (
	"EternityGUI/shared"
	"fmt"
	vrc "github.com/project-vrcat/vrchat-go"
	"os"
	"strings"
)

func FriendSpam(uid string) {

	if _, err := os.Stat("friends.txt"); err != nil {
		_, err = os.Create("friends.txt")
		if err != nil {
			return
		}
	}

	count, _ := shared.AccManager.LoadFromFile("friends.txt")
	fmt.Println("Loaded", count, "combos")

	client := vrc.NewClient()

	for i, combo := range shared.AccManager.AccountList {

		data := strings.Split(combo, ":")

		client.Login(data[0], data[1])

		token, _ := client.AuthToken()

		body := SendFriendRequest(uid, token)
		if i > 4 {
			break
		}
		fmt.Println(body)

	}

}
