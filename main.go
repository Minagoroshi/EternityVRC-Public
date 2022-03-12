package main

import (
	"EternityGUI/cmd/vrchat/vrcapi"
	"EternityGUI/cmd/vrchat/vrcapi/authentication"
	vrcwss "EternityGUI/cmd/vrchat/vrcapi/websocket"
	"EternityGUI/gui"
	"EternityGUI/utils"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/dialog"
	"log"
	"net/url"
	"strings"
	"time"
)

func main() {
	//Startup
	utils.DiscordRPC()
	utils.Startup()
	App := app.NewWithID("Eternity")
	Icon, _ := fyne.LoadResourceFromURLString("https://i.ibb.co/Zfm7RvL/icon.png")
	MainWindow := App.NewWindow("Eternity")
	MainWindow.SetMaster()
	MainWindow.Resize(fyne.NewSize(300, 220))
	MainWindow.SetIcon(Icon)
	go MainWindow.Resize(fyne.NewSize(1200, 700))
	go MainWindow.SetMainMenu(makeMenu(App, MainWindow))
	MainWindow.SetContent(gui.VRChatScreen(MainWindow))
	t := vrcAuth.VerifyToken()
	if !strings.Contains(t, "true") {
		log.Println("Login Failed!")
		dialog.ShowInformation("Invalid Login Session", "Please goto Settings -> Config -> Login", MainWindow)
	} else {
		log.Println("Login Success!")
		go vrcwss.InitClient()
	}

	go func() {
		time.Sleep(2 * time.Second)
		vrcapi.C1 <- "Test Success!"
	}()

	MainWindow.ShowAndRun()

}

//makeMenu is our function for creation the main menu
func makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	Icon, _ := fyne.LoadResourceFromURLString("https://i.ibb.co/Zfm7RvL/icon.png")

	AppearenceItem := fyne.NewMenuItem("Appearance", func() {
		w := a.NewWindow("Appearance")
		w.SetContent(settings.NewSettings().LoadAppearanceScreen(w))
		w.Resize(fyne.NewSize(480, 480))
		w.SetIcon(Icon)
		w.Show()
	})

	SettingsItem := fyne.NewMenuItem("Config", func() {
		w := a.NewWindow("Config")
		w.SetContent(gui.SettingsScreen(w))
		w.Resize(fyne.NewSize(400, 0))
		w.SetIcon(Icon)
		w.Show()
	})

	helpMenu := fyne.NewMenu("Help",

		fyne.NewMenuItem("Discord", func() {
			u, _ := url.Parse("https://discord.gg/cope")
			_ = a.OpenURL(u)
		}))

	// A quit item will be appended to our first (File) menu
	file := fyne.NewMenu("Settings")
	if !fyne.CurrentDevice().IsMobile() {
		file.Items = append(file.Items, SettingsItem, fyne.NewMenuItemSeparator(), AppearenceItem)
	}
	return fyne.NewMainMenu(
		file,
		helpMenu,
	)
}
