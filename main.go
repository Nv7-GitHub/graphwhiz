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
	a := app.NewWithID("com.nv.graphwhiz")
	win = a.NewWindow("Graphwhiz")

	// Setup UI
	vbox := container.NewVBox()
	outputForm := widget.NewForm()

	// Put these before so that they can be captured by the file open closure
	outFileShower := widget.NewEntry()
	fileShower := widget.NewEntry()

	saveBtn := widget.NewButton("Select", func() {
		dlg := dialog.NewFileOpen(func(file fyne.URIReadCloser, err error) {
			if file != nil {
				fileShower.SetText(file.URI().Path())

				if outFileShower.Text == "" {
					ext := filepath.Ext(fileShower.Text)
					outFileShower.SetText(fileShower.Text[:len(fileShower.Text)-len(ext)] + ".png")
				}
			}
		}, win)
		dlg.SetFilter(storage.NewExtensionFileFilter([]string{".dot"}))
		dlg.Show()
	})
	hbox := container.NewBorder(nil, nil, nil, saveBtn, fileShower)

	outSaveBtn := widget.NewButton("Select", func() {
		dlg := dialog.NewFileSave(func(file fyne.URIWriteCloser, err error) {
			if file != nil {
				outFileShower.SetText(file.URI().Path())
			}
		}, win)
		dlg.Show()
	})
	outHbox := container.NewBorder(nil, nil, nil, outSaveBtn, outFileShower)

	outputTypeBox := widget.NewSelect(outputTypeList, func(_ string) {})
	outputTypeBox.SetSelectedIndex(0)

	layoutBox := widget.NewSelect(layoutList, func(_ string) {})
	layoutBox.SetSelectedIndex(0)

	check := widget.NewCheck("", func(_ bool) {})
	check.SetChecked(false)

	system := widget.NewCheck("", func(_ bool) {})
	system.SetChecked(false)

	outputForm.Append("Input DOT File", hbox)
	outputForm.Append("Output File", outHbox)
	outputForm.Append("Output Type", outputTypeBox)
	outputForm.Append("Layout", layoutBox)
	outputForm.Append("Remove original file", check)
	outputForm.Append("Use System Graphviz", system)

	var renderBtn *widget.Button
	renderBtn = widget.NewButton("Render!", func() {
		renderBtn.Disable()
		renderBtn.SetText("Rendering...")

		err := Render(fileShower.Text, outFileShower.Text, outputTypeBox.Selected, layoutBox.Selected, system.Checked)
		handle(err)

		if check.Checked {
			err = os.Remove(fileShower.Text)
			handle(err)
		}

		renderBtn.Enable()
		renderBtn.SetText("Render!")
	})

	vbox.Add(outputForm)
	vbox.Add(renderBtn)
	win.SetContent(vbox)
	win.Resize(fyne.NewSize(700, vbox.MinSize().Height))

	win.ShowAndRun()
}
