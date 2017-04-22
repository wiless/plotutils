package pf

import (
	"image"

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
	app = x11ui.NewApp(false, 500, 300)
	app.SetDefaultKeys()
	childwin = app.NewChildWindow("Figure", 0, 0, 500, 300)
	rawimg = childwin.CreateRawImage(0, 0, 500, 300)

	vgopts := vgimg.NewWith(vgimg.UseImage(rawimg))
	canvas = draw.New(vgopts)

}
func Wait() {
	app.Show()
}

func CloseAll() {
	app.Close()
}
func ShowX11() {
	if gcf != nil {

		gcf.Draw(canvas)
		// draw.NewCanvas(c, w, h)
		// c := vgtex.NewDocument(5*vg.Centimeter, 5*vg.Centimeter)
		// c := vgtex.NewDocument(10*vg.Inch, 7.5*vg.Inch)

		// p.Draw(draw.New(c))
		childwin.ReDrawImage()

	}
}
