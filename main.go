package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var count = 1

func main() {

	// Starting to create an app with fyne
	a := app.New()

	// Now we will create a new window
	w := a.NewWindow("Nano Text Editor") // Add title of the app

	// Resizing the window
	w.Resize(fyne.NewSize(700, 700))

	content := container.NewVBox(
		container.NewHBox(
			widget.NewLabel("Nano Text Editor"),
		),
	)

	content.Add(widget.NewButton("Add new file", func() {
		content.Add(widget.NewLabel("New file " + strconv.Itoa(count)))
		count++
	}))

	input := widget.NewMultiLineEntry() 
	input.SetPlaceHolder("Enter Text..")
	input.Resize(fyne.NewSize(400, 400))

	saveBtn := widget.NewButton("Save", func() {

	})

	w.SetContent(
		container.NewVBox(
			content,
			input,

			container.NewHBox(
				saveBtn,
			),
		),
	)
	w.ShowAndRun() // Run the app
}
