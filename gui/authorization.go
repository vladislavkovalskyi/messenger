package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func Authorization(application fyne.App) {
	window := application.NewWindow("Dora | Auth")
	window.Resize(fyne.NewSize(600, 300))
	window.SetContent(widget.NewLabel("Hello World!"))
	window.ShowAndRun()
}
