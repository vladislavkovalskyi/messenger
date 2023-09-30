package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/image/colornames"
	"image/color"
	"messenger/tools"
)

func SetContent(window fyne.Window, screenX float32, screenY float32) {
	welcomeLabel := canvas.NewText("Dora | Authorization", color.White)
	welcomeLabel.TextSize = 30
	welcomeContainer := container.New(layout.NewCenterLayout(), welcomeLabel)

	nicknameEntry := widget.NewEntry()
	nicknameEntry.PlaceHolder = "Введите свой никнейм"
	nicknameEntry.Resize(fyne.NewSize(200, 40))
	nicknameEntry.Move(fyne.NewPos(screenX/2-screenX+250, 1))

	acceptButton := widget.NewButton("Подключиться", func() { fmt.Println("Here will be connection code") })
	acceptButton.Resize(fyne.NewSize(200, 40))
	acceptButton.Move(fyne.NewPos(screenX/2-screenX+450, 1))
	connectionContainer := container.NewWithoutLayout(nicknameEntry, acceptButton)

	descriptionLabel1 := canvas.NewText("Это проект, который позволяет использовать", colornames.Lightblue)
	descriptionLabel1.TextSize = 13
	descriptionLabel2 := canvas.NewText("защищенный и анонимный online-chatting.", colornames.Lightblue)
	descriptionLabel2.TextSize = 13
	descriptionLabel3 := canvas.NewText("С любовью, Kovalskyi", colornames.White)
	descriptionLabel3.TextSize = 15
	descriptionContainer := container.New(layout.NewGridLayoutWithRows(3), descriptionLabel1, descriptionLabel2, descriptionLabel3)

	contentContainer := container.NewVBox(
		welcomeContainer,
		layout.NewSpacer(),
		container.NewCenter(connectionContainer),
		layout.NewSpacer(),
		container.NewCenter(descriptionContainer),
	)

	window.SetContent(contentContainer)
}

func Authorization(application fyne.App) {
	screenWidth, screenHeight, _ := tools.GetScreenResolution(.4)

	window := application.NewWindow("Dora")
	window.Resize(fyne.NewSize(screenWidth, screenHeight))
	SetContent(window, screenWidth, screenHeight)

	window.ShowAndRun()
}
