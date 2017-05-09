package pf

import (
	"image"
	"log"

	"github.com/gonum/plot/vg/draw"
	"github.com/gonum/plot/vg/vgimg"
	"github.com/wiless/x11ui"
)

var app *x11ui.Application
var childwin *x11ui.Window
var rawimg *image.RGBA
var canvas draw.Canvas

func init() {
	app = nil
}

func IsX() bool {
	return app != nil
}

func StartX() {
	app = x11ui.NewApp(false, 500, 400)
	app.SetDefaultKeys()
	childwin = app.NewChildWindow("Figure", 0, 0, 500, 400)
	// app.SetLayoutSpacing(50, 50)
	// app.NewChildWindow("Figure 2", 500, 200)
	// app.NewChildWindow("Figure 2", 500, 200)
	// app.NewChildWindow("Figure 2", 500, 200)
	app.RegisterKey("s", SavePng)
	app.RegisterKey("z", ZoomIn)
	app.RegisterKey("n", ShowNext)
	app.RegisterKey("p", ShowPrev)

	app.RegisterKey("shiftZ", ZoomOut)
	rawimg = childwin.CreateRawImage(0, 0, 500, 400)
	vgopts := vgimg.NewWith(vgimg.UseImage(rawimg))
	canvas = draw.New(vgopts)

}

func ZoomIn() {
	canvas.Scale(1.5, 1.5)
	ShowX11()
}

func ZoomOut() {
	log.Print("Scale from", canvas.Size())
	canvas.Scale(.5, .5)

	ShowX11()
}

func Wait() {
	app.Show()
}

func CloseAll() {
	app.Close()
}
func ShowX11() {
	if app == nil {
		return
	}
	if gcf != nil {

		gcf.Draw(canvas)
		// draw.NewCanvas(c, w, h)
		// c := vgtex.NewDocument(5*vg.Centimeter, 5*vg.Centimeter)
		// c := vgtex.NewDocument(10*vg.Inch, 7.5*vg.Inch)

		// p.Draw(draw.New(c))
		childwin.ReDrawImage()

	}
}
