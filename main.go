package main

import (
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

var win fyne.Window

func handle(err error) {
	if err != nil {
		dlg := dialog.NewError(err, win)
		dlg.SetOnClosed(func() {
			panic(err)
		})
		dlg.Show()
	}
}

func main() {
	a := app.New()
	win = a.NewWindow("Graphwhiz")

	// Setup UI
	vbox := container.NewVBox()
	outputForm := widget.NewForm()

	// Put these before so that they can be captured by the file open closure
	outFilename := ""
	outFileShower := widget.NewEntry()

	filename := ""
	hbox := container.NewHBox()
	fileShower := widget.NewEntry()
	saveBtn := widget.NewButton("Select", func() {
		dlg := dialog.NewFileOpen(func(file fyne.URIReadCloser, err error) {
			if file != nil {
				filename = file.URI().Path()
				fileShower.SetText(filename)

				if outFilename == "" {
					ext := filepath.Ext(filename)
					outFilename = filename[:len(filename)-len(ext)]
					outFileShower.SetText(outFilename)
				}
			}
		}, win)
		dlg.SetFilter(storage.NewExtensionFileFilter([]string{".dot"}))
		dlg.Show()
	})
	hbox.Add(fileShower)
	hbox.Add(saveBtn)

	outHbox := container.NewHBox()
	outSaveBtn := widget.NewButton("Select", func() {
		dlg := dialog.NewFileSave(func(file fyne.URIWriteCloser, err error) {
			if file != nil {
				outFilename = file.URI().Path()
				outFileShower.SetText(filename)
			}
		}, win)
		dlg.Show()
	})
	outHbox.Add(outFileShower)
	outHbox.Add(outSaveBtn)

	outputTypeBox := widget.NewSelect(outputTypeList, func(_ string) {})
	outputTypeBox.SetSelectedIndex(0)

	layoutBox := widget.NewSelect(layoutList, func(_ string) {})
	layoutBox.SetSelectedIndex(0)

	check := widget.NewCheck("", func(_ bool) {})
	check.SetChecked(false)

	outputForm.Append("Input DOT File", hbox)
	outputForm.Append("Output File", outHbox)
	outputForm.Append("Output Type", outputTypeBox)
	outputForm.Append("Layout", layoutBox)
	outputForm.Append("Remove original file", check)

	var renderBtn *widget.Button
	renderBtn = widget.NewButton("Render!", func() {
		renderBtn.Disable()
		renderBtn.SetText("Rendering...")

		err := Render(filename, outFilename, outputTypeBox.Selected, layoutBox.Selected)
		handle(err)

		if check.Checked {
			err = os.Remove(filename)
			handle(err)
		}

		renderBtn.Enable()
		renderBtn.SetText("Render!")
	})

	vbox.Add(outputForm)
	vbox.Add(renderBtn)
	win.SetContent(vbox)

	win.ShowAndRun()
}
