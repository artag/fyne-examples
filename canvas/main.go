package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func setContentToRectangle(c fyne.Canvas) *canvas.Rectangle {
	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	rect := canvas.NewRectangle(blue)
	c.SetContent(rect)
	return rect
}

func runRectangleAnimation(rect *canvas.Rectangle) {
	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}

	for {
		time.Sleep(time.Second)

		if rect.FillColor == blue {
			rect.FillColor = green
		} else {
			rect.FillColor = blue
		}

		rect.Refresh()
	}
}

func setContentToText(c fyne.Canvas) {
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text := canvas.NewText("Text", green)
	text.TextStyle.Bold = true
	c.SetContent(text)
}

func setContentToCircle(c fyne.Canvas) {
	red := color.NRGBA{R: 0xff, G: 0x33, B: 0x33, A: 0xff}
	circle := canvas.NewCircle(color.White)
	circle.StrokeWidth = 4
	circle.StrokeColor = red
	c.SetContent(circle)
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Canvas")
	myCanvas := myWindow.Canvas()

	// rect := setContentToRectangle(myCanvas)
	// go runRectangleAnimation(rect)

	// setContentToText(myCanvas)

	setContentToCircle(myCanvas)

	myWindow.Resize(fyne.NewSize(100, 100))
	myWindow.ShowAndRun()
}
