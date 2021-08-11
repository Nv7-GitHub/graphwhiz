package main

import (
	"github.com/andlabs/ui"
)

var win *ui.Window

func setupUI() {
	win = ui.NewWindow("Screenlapse", 640, 0, true)
	win.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		win.Destroy()
		return true
	})
	win.SetMargined(true)

	// Setup UI
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	outputGroup := ui.NewGroup("Settings")
	outputGroup.SetMargined(true)
	outputForm := ui.NewForm()
	outputForm.SetPadded(true)

	filename := ""
	hbox := ui.NewHorizontalBox()
	fileShower := ui.NewEntry()
	saveBtn := ui.NewButton("Select")
	saveBtn.OnClicked(func(*ui.Button) {
		filename = ui.OpenFile(win)
		fileShower.SetText(filename)
	})
	hbox.Append(fileShower, true)
	hbox.Append(saveBtn, false)

	outFilename := ""
	outHbox := ui.NewHorizontalBox()
	outFileShower := ui.NewEntry()
	outSaveBtn := ui.NewButton("Select")
	outSaveBtn.OnClicked(func(*ui.Button) {
		outFilename = ui.SaveFile(win)
		outFileShower.SetText(filename)
	})
	outHbox.Append(outFileShower, true)
	outHbox.Append(outSaveBtn, false)

	outputTypeBox := ui.NewCombobox()
	for _, format := range outputTypeList {
		outputTypeBox.Append(format)
	}
	outputTypeBox.SetSelected(0)

	layoutBox := ui.NewCombobox()
	for _, layout := range layoutList {
		layoutBox.Append(layout)
	}
	layoutBox.SetSelected(0)

	outputForm.Append("Input DOT File", hbox, false)
	outputForm.Append("Output File", outHbox, false)
	outputForm.Append("Output Type", outputTypeBox, false)
	outputForm.Append("Layout", layoutBox, false)
	outputGroup.SetChild(outputForm)

	renderBtn := ui.NewButton("Render!")
	renderBtn.OnClicked(func(*ui.Button) {
		renderBtn.Disable()
		renderBtn.SetText("Converting...")
		err := Render(filename, outFilename, outputTypeList[outputTypeBox.Selected()], layoutList[layoutBox.Selected()])
		handle(err)
		renderBtn.Enable()
	})

	vbox.Append(outputGroup, false)
	vbox.Append(renderBtn, false)
	win.SetChild(vbox)

	win.Show()
}
