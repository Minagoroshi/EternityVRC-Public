package gui

import (
	vrcfriends "EternityGUI/cmd/vrchat/vrcapi/friends"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func FunScreen(win fyne.Window) fyne.CanvasObject {

	content := container.NewVBox(widget.NewButton("Friend Spam", func() {
		label := widget.NewLabel("NOTE: 4 Accounts MAX Unless Running Eternity through Proxifier,\n" +
			" integrated proxy support will be added in the future")
		entry := widget.NewEntry()
		dialogContent := container.NewVBox(label, entry)
		dialog.ShowCustomConfirm("Friend Spam (UserID)", "Execute", "Dismiss", dialogContent, func(b bool) {
			if b {
				vrcfriends.FriendSpam(entry.Text)
			}

		}, win)
	}))

	return content

}
