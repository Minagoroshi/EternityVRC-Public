package vrcapi

import "EternityGUI/utils"

var (
	BaseURL         = "https://api.vrchat.cloud/api/1/"
	AuthToken       = utils.ReadConfig().VRChatLogin
	FriendOnline    = make(chan string)
	FriendOffline   = make(chan string)
	FriendLocation  = make(chan string)
	VRCNotification = make(chan string)
	FriendUpdate    = make(chan string)
	C1              = make(chan string)
)
