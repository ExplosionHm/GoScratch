package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	//Just the basic example for now because I'm quite busy currently
	a := app.New()
	w := a.NewWindow("Hello")
	
	hello := widget.NewLabel("Hello Fyne!")
	con := container.NewVBox(
		hello,
		widget.NewButton("thing", func() {
			hello.SetText("Hello world")
		}),
	)
	w.SetContent(con)
	w.Resize(fyne.Size{Width: 1024, Height: 748})
	w.CenterOnScreen()
	w.ShowAndRun()
}
