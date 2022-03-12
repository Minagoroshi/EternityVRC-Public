package utils

import (
	"github.com/hugolgst/rich-go/client"
	"log"
	"time"
)

//DiscordRPC initiates a rich presence on discord
func DiscordRPC() {
	err := client.Login("931669767222329344")
	if err != nil {
		log.Println(err)
	}

	now := time.Now()
	err = client.SetActivity(client.Activity{
		State:      "Allocating Females",
		Details:    "Active",
		LargeImage: "lrg",
		LargeText:  "Skid",
		SmallImage: "small",
		SmallText:  "You're Broke",

		Timestamps: &client.Timestamps{
			Start: &now,
		},
		Buttons: []*client.Button{
			{
				Label: "Purchase",
				Url:   "https://sellix.io/EternityBot",
			},
			{
				Label: "Discord",
				Url:   "https://discord.gg/cope",
			},
		},
	})

	if err != nil {
		log.Println(err)
	}

	// Discord will only show the presence if the app is running
	// Sleep for a few seconds to see the update
	log.Println("Starting Discord RPC")
}
