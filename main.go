package main

import (
	"fmt"
	"io"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)


func main() {
	// Starting to create an app with fyne
	a := app.New()

	// Now we will create a new window
	w := a.NewWindow("Nano Text Editor") // Add title of the app

	// Resizing the window
	w.Resize(fyne.NewSize(700, 700))

	content := container.NewVBox(
		container.NewHBox(
			widget.NewLabel(""),
		),
	)

	
	// Declare 'input' outside the button click event
	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter Text..")
	input.Resize(fyne.NewSize(400,400))

	var nw fyne.Window
	savedFiles := make([]string, 0)

	saveBtn := widget.NewButton("Save", func() {
			saveFileDialog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {
				textData := []byte(input.Text)
				uc.Write(textData)
				
				input = widget.NewMultiLineEntry()
				dialog.ShowInformation("Success", "File saved successfully", nw)
	
				// Update UI to show saved files
				updateSavedFilesList(content, savedFiles)
												
			}, nw)
					
			
		// changed
		now := time.Now()
		timeString := now.Format("15-04-05")

		fileName := strings.Fields(input.Text)
		var newFileName string
		if len(fileName) >= 5 {
			newFileName = strings.Join(fileName[:5], " ")
		} else {
			newFileName = input.Text
		}

		fname := fmt.Sprintf("%s %s.txt",newFileName,timeString)
		saveFileDialog.SetFileName(fname)
		saveFileDialog.Resize(fyne.NewSize(700,500))
		saveFileDialog.Show()

		savedFiles = append(savedFiles, fname)
		
				
	})

	//content.Add
	new := (widget.NewButton("Add new file", func() {

		newWindow := fyne.CurrentApp().NewWindow("New File Window")
		newWindow.Resize(fyne.NewSize(700, 700))
		nw = newWindow

		newWindow.SetContent(
			container.NewVBox(
				input,
				saveBtn,
			),
		)
		newWindow.Show()
				
	}))

	openBtn := widget.NewButton("Open File", func() {
		
		openFileDialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				readData, _ := io.ReadAll(r)

				output := fyne.NewStaticResource("New File", readData)
				viewData := widget.NewMultiLineEntry()
				viewData.SetText(string(output.StaticContent))

				w := fyne.CurrentApp().NewWindow(
					string(output.StaticName))
				w.SetContent(container.NewScroll(viewData))
				w.Resize(fyne.NewSize(400, 400))
				w.Show()
			}, w,
		)
	
		// filter only text files when perform open
		openFileDialog.SetFilter(
			storage.NewExtensionFileFilter([]string{".txt"}),
		)
		openFileDialog.Show()
	})

	// Widget to display saved files
	savedFilesLabel := widget.NewLabelWithStyle("Saved Files",fyne.TextAlignCenter,fyne.TextStyle{ Monospace: true})
	
	savedFilesList := widget.NewLabel("")

	
	
	
	w.SetContent(
	//	container.NewVBox(
			
			container.NewVBox(
				new,
				openBtn,
				content,
				
			),
		//),
	)
	content.Add(savedFilesLabel)
	content.Add(savedFilesList)

	w.ShowAndRun() // Run the app
}

func updateSavedFilesList(content *fyne.Container, savedFiles []string){
	if len(savedFiles) > 0 {
		savedFileListText := ""
		for _, file := range savedFiles {
			savedFileListText += file + "\n"
		}
		content.Objects[3].(*widget.Label).SetText(savedFileListText)
	}
}