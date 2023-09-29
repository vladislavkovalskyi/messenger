package gui

import "fyne.io/fyne/v2/app"

func Application() {
	application := app.New()

	Authorization(application)
}
