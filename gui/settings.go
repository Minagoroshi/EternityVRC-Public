package gui

import (
	vrcwss "EternityGUI/cmd/vrchat/vrcapi/websocket"
	"EternityGUI/utils"
	"encoding/json"
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	vrc "github.com/project-vrcat/vrchat-go"
	"net/url"
	"os"
	"strings"
	"time"
)

func SettingsScreen(win fyne.Window) fyne.CanvasObject {
	MainGroup := container.NewVBox(
		widget.NewButton("Login to VRChat", func() {

			LoginEntry := widget.NewEntry()
			PasswordEntry := widget.NewPasswordEntry()
			items := []*widget.FormItem{
				widget.NewFormItem("Username", LoginEntry),
				widget.NewFormItem("Password", PasswordEntry),
			}

			dialog.ShowForm("VRChat Login", "Submit", "Cancel", items, func(b bool) {
				if !b {
					return
				}
				u := strings.TrimSpace(LoginEntry.Text)
				p := strings.TrimSpace(PasswordEntry.Text)
				vrc := vrc.NewClient()
				user, err := vrc.Login(u, p)
				username := user.DisplayName
				if err != nil {
					e := errors.New("Invalid Username or Password, try again")
					dialog.ShowError(e, win)
				} else {
					uid := user.ID
					token, _ := vrc.AuthToken()
					cfg := &utils.UserConfig{
						VRChatLogin: token,
						VRChatUID:   uid,
					}
					data, _ := json.Marshal(cfg)
					os.WriteFile("config/config.json", data, 0755)
					time.Sleep(2 * time.Second)
					dialog.ShowInformation("Success", "Successfully logged into VRChat as: "+username, win)
					go vrcwss.InitClient()
				}

			}, win)

		}),
		widget.NewButton("Discord", func() {
			app := fyne.CurrentApp()
			u, _ := url.Parse("https://discord.gg/cope")
			_ = app.OpenURL(u)
		}),
		widget.NewButton("Sellix", func() {
			app := fyne.CurrentApp()
			u, _ := url.Parse("https://sellix.io/EternityBot")
			_ = app.OpenURL(u)
		}),
		widget.NewButton("Logout", func() {
			dialog.ShowConfirm("Confirmation", "Are you sure you want to logout?", func(b bool) {
				if b != true {
					return
				}
				config := utils.ReadConfig()
				cfg := &utils.UserConfig{
					VRChatLogin: "",
					VRChatUID:   config.VRChatUID,
				}
				s, _ := json.Marshal(cfg)
				err := os.WriteFile("config/config.json", s, 0755)
				if err != nil {
					return
				}
				os.Exit(0)

			}, win,
			)
		}))

	//return container.NewVBox(widget.NewCard("Settings", "Eternity Desktop Settings", container.NewVBox(MainGroup)))
	return widget.NewCard("Config Settings", "", MainGroup)

}
