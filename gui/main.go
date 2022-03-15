package gui

import (
	"EternityGUI/cmd/vrchat"
	vrcfriends "EternityGUI/cmd/vrchat/vrcapi/friends"
	vrcusers "EternityGUI/cmd/vrchat/vrcapi/users"
	vrcworlds "EternityGUI/cmd/vrchat/vrcapi/worlds"
	"EternityGUI/shared"
	"EternityGUI/utils"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
	"github.com/buger/jsonparser"
	"log"
	"net/url"
	"strconv"
	"strings"
)

//TODO: Bot Handler Functionality

func VRChatScreen(win fyne.Window) fyne.CanvasObject {
	//Define Data Variables
	var targetID, body, contentID string
	var idData, AvatarIdData, AuthorNameData, DescriptionData, ReleaseStatusData, FeaturedData, statusDescriptionData, statusData,
		ImageURLData, trustData, AssetURLData, fdata, fImageURLData, fstatusDescriptionData, fstatusData, ftrustData, fidData, IconURLData []string
	var count, fcount int
	var img fyne.Resource

	searchType := 0

	//Gui Variables
	MoreLink := "https://discord.gg/cope"
	image := canvas.NewImageFromResource(img)
	image.FillMode = canvas.ImageFillContain
	image.SetMinSize(fyne.NewSize(300, 300))
	displayText := widget.NewLabel("")
	/*r, _ := fyne.LoadResourceFromURLString("https://www.welovesolo.com/wp-content/uploads/vecteezy/11/hrtfeqrzylm.jpg")
	Backround := canvas.NewImageFromResource(r)
	Backround.FillMode = canvas.ImageFillStretch*/
	clip := ""
	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		//Open browser containing more item information
		widget.NewToolbarAction(theme.LoginIcon(), func() {
			u, _ := url.Parse(MoreLink)
			_ = fyne.CurrentApp().OpenURL(u)
		}),
		//Copy
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {
			err := clipboard.WriteAll(displayText.Text + "\n" + clip)
			if err != nil {
				return
			}
		}),
		//Toolbar Actions
		//0 = None
		//1 = Player
		//2 = World
		//3 = Avatar
		//4 = Friend
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			switch searchType {
			case 0: // 0 = None

				dialog.ShowInformation("No selection", "You have nothing selected, please select something from the search menu", win)
			case 1: // 1 = Player
				items :=
					container.NewVBox(widget.NewButton("Copy UserID", func() {
						err := clipboard.WriteAll(contentID)
						if err != nil {
							return
						}
					}),
						widget.NewButton("Send Friend Request", func() {
							res := vrcfriends.SendFriendRequest(targetID, "")
							fmt.Println(res)
						}))
				//Target Options
				dlg := dialog.NewCustom("Target Options", "Dismiss", items, win)
				dlg.Show()

			case 2: // 2 = World

				items := container.NewVBox(widget.NewButton("Copy World ID", func() {
					err := clipboard.WriteAll(contentID)
					if err != nil {
						return
					}
				}))

				dialog.ShowCustom("Target Options", "Done", items, win)
			case 3: //Avatars

				items := container.NewVBox(widget.NewButton("Copy Avatar ID", func() {
					err := clipboard.WriteAll(contentID)
					if err != nil {
						return
					}
				}))

				dialog.ShowCustom("Target Options", "Done", items, win)

			case 4: // Friends
				items :=
					container.NewVBox(widget.NewButton("Copy UserID", func() {
						err := clipboard.WriteAll(contentID)
						if err != nil {
							return
						}
					}))
				//Target Options
				dlg := dialog.NewCustom("Target Options", "Dismiss", items, win)

				dlg.Show()

			}
		}),
	)

	displayGroup := container.NewHBox(image, container.NewCenter(container.NewVBox(toolbar, displayText)))
	//Search Box
	var selection string
	Search := widget.NewEntry()
	SearchOptions := widget.NewRadioGroup([]string{"Player Search", "World Search", "Avatar Search" /*, "Friend Search"*/}, func(value string) {
		log.Println("Selected", value)
		selection = value
	})
	//idk
	SearchOptions.Horizontal = true
	SearchOptions.Required = true
	SearchOptions.Selected = "Player Search"

	data := make([]string, 1)
	for i := range data {
		data[i] = "No search results"
	}

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(theme.AccountIcon()), widget.NewLabel(""))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(data[id])

		},
	)

	//Make Friend List
	fdata, fidData, fImageURLData, fstatusData, fstatusDescriptionData, ftrustData, fcount = vrcfriends.MakeFriendList()
	s := strconv.Itoa(fcount)
	FriendList := widget.NewList(
		func() int {
			return len(fdata)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(theme.AccountIcon()), widget.NewLabel(""))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(fdata[id])
		},
	)
	FriendCard := widget.NewCard("Friends List", "", widget.NewRichTextWithText("Friends Online: "+s))

	// Search Switch
	S := container.NewVBox(widget.NewCard("Search", "Search functions for VRChat",
		container.NewVBox(SearchOptions,
			Search,
			widget.NewButton("Search", func() {
				config := utils.ReadConfig()
				token := config.VRChatLogin
				switch selection {
				case "Player Search": //Player Search | Search Type 1
					list.UnselectAll()
					s := url.QueryEscape(Search.Text)
					body = vrcusers.PlayerSearch(s, token)
					count = strings.Count(body, "currentAvatarThumbnailImageUrl")
					//Make Data Arrays
					data = make([]string, count)
					idData = make([]string, count)
					ImageURLData = make([]string, count)
					statusDescriptionData = make([]string, count)
					trustData = make([]string, count)
					//tagsData = make([]string, count)
					//bioData = make([]string, count)
					statusData = make([]string, count)
					log.Println("Results:", count)
					//Populate Data
					for i := range data {
						n := "[" + strconv.Itoa(i) + "]"

						result, _ := jsonparser.GetString([]byte(body), n, "displayName")
						data[i] = result

						resultID, _ := jsonparser.GetString([]byte(body), n, "id")
						idData[i] = "User ID: " + resultID

						resultImageUrl, _, _, _ := jsonparser.Get([]byte(body), n, "currentAvatarImageUrl")
						ImageURLData[i] = string(resultImageUrl)
						//resultBio, _, _, _ := jsonparser.Get([]byte(body), n, "bio")

						resultStatus, _ := jsonparser.GetString([]byte(body), n, "status")
						resultStatusDescription, _ := jsonparser.GetString([]byte(body), n, "statusDescription")

						resultTags, _, _, _ := jsonparser.Get([]byte(body), n, "tags")
						//tagsData[i] =  resultTags
						resultTrust := vrchat.TagsConverter(resultTags)
						trustData[i] = "Trust: " + resultTrust

						statusData[i] = "Status: " + string(resultStatus)
						statusDescriptionData[i] = "Status Description: " + string(resultStatusDescription)

					}
					searchType = 1

				case "World Search": // World Search | Search Type 2
					list.UnselectAll()
					s := url.QueryEscape(Search.Text)
					log.Println(s)
					body = vrcworlds.WorldSearch(s, token)
					count = strings.Count(body, "labsPublicationDate")
					data = make([]string, count)
					ImageURLData = make([]string, count)
					idData = make([]string, count)
					AuthorNameData = make([]string, count)
					log.Println("Results:", count)
					for i := range data {
						n := "[" + strconv.Itoa(i) + "]"

						result, _ := jsonparser.GetString([]byte(body), n, "name")
						resultID, _ := jsonparser.GetString([]byte(body), n, "id")

						resultAuthorName, _ := jsonparser.GetString([]byte(body), n, "authorName")
						resultImage, _, _, _ := jsonparser.Get([]byte(body), n, "thumbnailImageUrl")

						data[i] = result
						idData[i] = "World ID: " + resultID

						AuthorNameData[i] = "Author Name: " + resultAuthorName
						ImageURLData[i] = string(resultImage)
					}
					searchType = 2
				case "Avatar Search": //Avatar Search | Search Type 3
					list.UnselectAll()
					body = vrchat.AvatarSearch(Search.Text)
					count = strings.Count(body, "ThumbnailImageURL")
					if count > 100 {
						count = 100
					}
					data = make([]string, count)
					AvatarIdData = make([]string, count)
					AuthorNameData = make([]string, count)
					DescriptionData = make([]string, count)
					FeaturedData = make([]string, count)
					ReleaseStatusData = make([]string, count)
					ImageURLData = make([]string, count)
					AssetURLData = make([]string, count)
					log.Println("Results:", count)
					for i := range data {
						n := "[" + strconv.Itoa(i) + "]"

						resultAvatarName, _ := jsonparser.GetString([]byte(body), "results", n, "AvatarName")
						data[i] = resultAvatarName

						resultAvatarID, _ := jsonparser.GetString([]byte(body), "results", n, "_id")
						AvatarIdData[i] = "Avatar ID: " + resultAvatarID

						resultAuthorName, _ := jsonparser.GetString([]byte(body), "results", n, "AuthorName")
						AuthorNameData[i] = "Author Name:  " + resultAuthorName

						resultDescription, _ := jsonparser.GetString([]byte(body), "results", n, "Description")
						DescriptionData[i] = "Description: " + resultDescription

						resultFeatured, _ := jsonparser.GetString([]byte(body), "results", n, "Featured")
						FeaturedData[i] = "Featured:  " + resultFeatured

						resultImageURL, _, _, _ := jsonparser.Get([]byte(body), "results", n, "ThumbnailImageURL")
						ImageURLData[i] = string(resultImageURL)

						resultAssetURL, _, _, _ := jsonparser.Get([]byte(body), "results", n, "AssetURL")
						AssetURLData[i] = "URL: " + string(resultAssetURL)

						resultReleaseStatus, _ := jsonparser.GetString([]byte(body), "results", n, "ReleaseStatus")
						ReleaseStatusData[i] = "Release Status:  " + resultReleaseStatus
					}

					searchType = 3

				default: //Default Search | Player Search | Search Type 1
					list.UnselectAll()
					s := url.QueryEscape(Search.Text)
					body = vrcusers.PlayerSearch(s, token)
					count = strings.Count(body, "currentAvatarThumbnailImageUrl")
					data = make([]string, count)
					idData = make([]string, count)
					ImageURLData = make([]string, count)
					statusDescriptionData = make([]string, count)
					trustData = make([]string, count)
					//tagsData = make([]string, count)
					//bioData = make([]string, count)
					statusData = make([]string, count)
					IconURLData = make([]string, count)
					log.Println("Results:", count)
					go func() {
						for i := range data {
							n := "[" + strconv.Itoa(i) + "]"
							result, _ := jsonparser.GetString([]byte(body), n, "displayName")
							resultID, _ := jsonparser.GetString([]byte(body), n, "id")
							//resultBio, _, _, _ := jsonparser.Get([]byte(body), n, "bio")
							resultStatus, _ := jsonparser.GetString([]byte(body), n, "status")
							resultStatusDescription, _ := jsonparser.GetString([]byte(body), n, "statusDescription")
							resultTags, _, _, _ := jsonparser.Get([]byte(body), n, "tags")
							resultTrust := vrchat.TagsConverter(resultTags)
							data[i] = result
							idData[i] = "User ID: " + resultID
							//tagsData[i] =  resultTags
							statusData[i] = "Status: " + string(resultStatus)
							statusDescriptionData[i] = "Status Description: " + string(resultStatusDescription)
							trustData[i] = "Trust: " + resultTrust
							resultImageUrl, _, _, _ := jsonparser.Get([]byte(body), n, "currentAvatarImageUrl")
							ImageURLData[i] = string(resultImageUrl)
						}
					}()
					//TODO: Preview Icons
					/*
						icon = make([]fyne.Resource, count)
						go func() {
						for i := range data {
							n := "[" + strconv.Itoa(i) + "]"
							resultImageUrl, _, _, _ := jsonparser.Get([]byte(body), n, "currentAvatarImageUrl")
							ImageURLData[i] = string(resultImageUrl)
							IconURLData[i] = ImageURLData[i]
							icon[i], _ = fyne.LoadResourceFromURLString(ImageURLData[i])
							log.Println("Loaded Icon", i)

						}
					}()*/
					searchType = 1
				}
				list.Refresh()
			}))))

	list.OnSelected = func(id widget.ListItemID) {
		// 0 = No Search
		// 1 = Player Search
		// 2 = World Search
		// 3 = Avatar Search
		FriendList.UnselectAll()

		//Reset Search Type if item from friends list was selected
		if searchType == 4 {
			switch selection {
			case "Player Search":
				searchType = 1
			case "World Search":
				searchType = 2
			case "Avatar Search":
				searchType = 3

			}
		}

		switch searchType {
		case 0:
			displayText.SetText(data[id])
		case 1: //Player Search
			log.Println(ImageURLData[id])
			r, _ := fyne.LoadResourceFromURLString(ImageURLData[id])
			image.Resource = r
			image.Refresh()
			go displayText.SetText(data[id] + "\n" +
				trustData[id] + "\n" +
				idData[id] + "\n" +
				statusData[id] + "\n" +
				statusDescriptionData[id])
			contentID = idData[id]
			uid := strings.ReplaceAll(idData[id], "User ID: ", "")
			MoreLink = "https://vrchat.com/home/user/" + uid
			targetID = uid
			clip = ImageURLData[id]
			dlg := dialog.NewCustom("Result", "Dismiss", displayGroup, win)
			dlg.Show()
			log.Println(MoreLink)
			log.Println("Update Success")
		case 2: //World Search
			contentID = idData[id]
			go displayText.SetText(data[id] + "\n" +
				idData[id] + "\n" +
				AuthorNameData[id])
			r, _ := fyne.LoadResourceFromURLString(ImageURLData[id])
			image.Resource = r
			image.Refresh()
			clip = ImageURLData[id]
			dlg := dialog.NewCustom("Result", "Dismiss", displayGroup, win)
			dlg.Show()
			log.Println("Update Success")
		case 3: //Avatar Search
			log.Println(ImageURLData[id])
			log.Println("Loading Image")
			r, _ := fyne.LoadResourceFromURLString(ImageURLData[id])
			image.Resource = r
			image.Refresh()
			log.Println("Setting Label Text")
			contentID = AvatarIdData[id]
			go displayText.SetText(data[id] + "\n" +
				AvatarIdData[id] + "\n" +
				AuthorNameData[id] + "\n" +
				DescriptionData[id] + "\n" +
				ReleaseStatusData[id] + "\n")
			log.Println("Performing String Operations")
			MoreLink = strings.ReplaceAll(AssetURLData[id], "URL: ", "")
			clip = AssetURLData[id] + "\n" + "Thumbnail: " + ImageURLData[id]
			dlg := dialog.NewCustom("Result", "Dismiss", displayGroup, win)
			dlg.Show()
			log.Println("Update Success")
		}
	}

	FriendList.OnSelected = func(id widget.ListItemID) {
		list.UnselectAll()
		log.Println("Loading Image Resource")
		r, _ := fyne.LoadResourceFromURLString(fImageURLData[id])
		image.Resource = r
		image.Refresh()
		log.Println("Setting Label Text")
		contentID = fidData[id]
		go displayText.SetText(fdata[id] + "\n" +
			ftrustData[id] + "\n" +
			fidData[id] + "\n" +
			fstatusData[id] + "\n" +
			fstatusDescriptionData[id])
		uid := strings.ReplaceAll(fidData[id], "User ID: ", "")
		MoreLink = "https://vrchat.com/home/user/" + uid
		targetID = uid
		clip = fImageURLData[id]
		dlg := dialog.NewCustom("Result", "Dismiss", displayGroup, win)
		dlg.Show()
		//time.Sleep(100 * time.Millisecond)
		//dlg.Resize(fyne.NewSize(975, 575))
		log.Println(MoreLink)
		log.Println("Update Success")

		searchType = 4

	}

	//Notifications Group

	nData := make([]string, 1)
	for i := range data {
		nData[i] = "Start of notifications history"
	}
	nList :=
		widget.NewList(
			func() int {
				return len(nData)
			},
			func() fyne.CanvasObject {
				return container.NewHBox(widget.NewIcon(theme.DocumentIcon()), widget.NewLabel(""))
			},
			func(id widget.ListItemID, item fyne.CanvasObject) {
				item.(*fyne.Container).Objects[1].(*widget.Label).SetText(nData[id])

			},
		)
	go func() {
		for {
			select {
			case online := <-shared.FriendOnline:
				log.Println("Received", online)
				nData = append([]string{online}, nData...)
				nList.Refresh()
			case offline := <-shared.FriendOffline:
				log.Println("Received", offline)
				nData = append([]string{offline}, nData...)
				nList.Refresh()
			case location := <-shared.FriendLocation:
				log.Println("Received", location)
				nData = append([]string{location}, nData...)
				nList.Refresh()
			case notification := <-shared.VRCNotification:
				log.Println("Received", notification)
				nData = append([]string{notification}, nData...)
				nList.Refresh()
			case update := <-shared.FriendUpdate:
				log.Println("Received", update)
				nData = append([]string{update}, nData...)
				nList.Refresh()
			case test := <-shared.C1:
				log.Println("Received", test)
			}
		}
	}()

	//Home Group

	//Profile Group

	//Gui Grouping
	//rect := canvas.NewRectangle(color.NRGBA{0x14, 0x14, 0x14, 0xff})
	ListGroup := list
	SearchGroup := container.NewBorder(container.NewVBox(S, widget.NewSeparator()), nil, nil, nil, ListGroup)
	FriendGroup := container.NewBorder(widget.NewButton("Refresh", func() {
		fdata, fidData, fImageURLData, fstatusData, fstatusDescriptionData, ftrustData, fcount = vrcfriends.MakeFriendList()
		s = strconv.Itoa(fcount)
		FriendCard.SetContent(widget.NewRichTextWithText("Friends Online: " + s))
		FriendCard.Refresh()
		FriendList.Refresh()
		log.Println(s)
	}), nil, nil, nil, FriendList)
	tabOptions :=
		container.NewAppTabs(
			//TODO:	container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("")),
			container.NewTabItemWithIcon("Search", theme.SearchIcon(), SearchGroup),
			container.NewTabItemWithIcon("Logs", theme.InfoIcon(), widget.NewCard("VRC Logs", "List of friend updates and notifications received from VRChat websocket", nList)),
			container.NewTabItemWithIcon("Fun Stuff", theme.WarningIcon(), widget.NewCard("Fun Stuff", "Fun Stuff for VRChat, we are not responsible for any account bans. ", FunScreen(win))),
			//TODO:	container.NewTabItemWithIcon("Moderation", theme.CheckButtonCheckedIcon(), widget.NewLabel("Nothing here")),
			//TODO:	container.NewTabItemWithIcon("Profile", theme.AccountIcon(), widget.NewLabel("Nothing here")),
		)
	tabOptions.SetTabLocation(container.TabLocationTop)
	MainGroup := container.NewHSplit(container.NewMax(container.NewBorder(FriendCard, nil, nil, nil, FriendGroup)), tabOptions)
	MainGroup.SetOffset(0.25)

	//Return to main
	return MainGroup

}
