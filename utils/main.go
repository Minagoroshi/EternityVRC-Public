package utils

import (
	"encoding/json"
	"log"
	"os"
)

func Startup() {
	//Check hits folder
	if stat, err := os.Stat("./Hits"); err == nil && stat.IsDir() {
		log.Println("Hits directory loaded.")
	} else {
		log.Println("Hits directory not found, creating it.")
		err := os.Mkdir("Hits", 0755)
		if err != nil {
			return
		}
	}

	//Check friend spam account list
	if _, err := os.Stat("friends.txt"); err != nil {
		_, err = os.Create("friends.txt")
		if err != nil {
			return
		}
		log.Println("Friends List Made")
	}

	//Check config folder
	if stat, err := os.Stat("./config"); err == nil && stat.IsDir() {
		log.Println("Config directory loaded.")
	} else {
		log.Println("Config directory not found, creating it.")
		err := os.Mkdir("config", 0755)
		if err != nil {
			return
		}
	}
	//Check Config file
	if _, err := os.Stat("./config/config.json"); err == nil {
		log.Println("Config loaded.")
	} else {
		log.Println("Config file not found, creating it.")
		config := &UserConfig{
			VRChatLogin: "",
			VRChatUID:   "",
		}
		data, _ := json.Marshal(config)
		f, _ := os.Create("config/config.json")
		_, err := f.WriteString(string(data))
		if err != nil {
			return
		}
		err = f.Close()
		if err != nil {
			return
		}
		log.Println("Config loaded.")
		if err != nil {
			return
		}
	}
}

func ReadConfig() UserConfig {
	var Config UserConfig
	f, _ := os.ReadFile("config/config.json")
	err := json.Unmarshal(f, &Config)
	if err != nil {
		return UserConfig{}
	}
	return Config

}
func SliceContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

type UserConfig struct {
	VRChatLogin string `json:"vrchatlogin"`
	VRChatUID   string `json:"vrchatuid"`
}
