package pf

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/image/colornames"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
	"github.com/gonum/plot/vg/draw"
	"github.com/gonum/plot/vg/vgeps"
	"github.com/gonum/plot/vg/vgpdf"
	"github.com/gonum/plot/vg/vgtex"
	"github.com/wiless/vlib"
)

var holdOn bool = false

var figures map[string]*plot.Plot
var currentFigure string
var err error
var gcf *plot.Plot

type Figure struct {
	*plot.Plot
}

func init() {
	figures = make(map[string]*plot.Plot)
}

func Fig(str ...string) {

	if len(str) == 0 {
		// if empty argument
		currentFigure = nextFigureName()
		gcf, err = plot.New()
		figures[currentFigure] = gcf
	} else {
		/// Ignoring more than one args..
		var isok bool
		gcf, isok = figures[str[0]]
		currentFigure = str[0]
		if !isok {
			// Create the figure
			gcf, err = plot.New()
			log.Print("Creating Figure ", currentFigure)
			figures[currentFigure] = gcf
		}

	}
	HoldOff()
}

func nextFigureName() string {
	return fmt.Sprint("Figure ", len(figures))
}

func HoldOn() {
	holdOn = true

}

func HoldOff() {
	holdOn = false
}

func getFigure() *plot.Plot {
	if len(figures) == 0 {
		Fig(nextFigureName())
		return gcf
	}

	if !holdOn {
		Fig(nextFigureName())
		return gcf
	} else {
		var isok bool
		gcf, isok = figures[currentFigure]
		if !isok {
			log.Print("Not possible ")
		}
		return gcf
	}

}

func SetLabel(x, y string) {
	p := getFigure()

	p.X.Label.Text = x
	p.Y.Label.Text = y

	Save()
}

func SetTitle(t string) {
	if gcf != nil {

		gcf.Title.Text = t
	}
}
func Save() {
	if gcf == nil {
		return
	}
	gcf.Title.TextStyle.Color = colornames.Blue

	// log.Print("Figure generated .. ", currentFigure+".png")
	gcf.Save(10*vg.Inch, 7.5*vg.Inch, currentFigure+".png")
	ShowX11()
	// gcf.Save(10*vg.Inch, 7.5*vg.Inch, currentFigure+".png")

}

func SaveTex() {
	p := getFigure()
	c := vgtex.NewDocument(5*vg.Centimeter, 5*vg.Centimeter)
	// c := vgtex.NewDocument(10*vg.Inch, 7.5*vg.Inch)

	p.Draw(draw.New(c))
	c.FillString(p.Title.Font, vg.Point{2.5 * vg.Centimeter, 2.5 * vg.Centimeter}, "x")

	f, err := os.Create(currentFigure + ".tex")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err = c.WriteTo(f); err != nil {
		log.Fatal(err)
	}
}

func SavePDF() {
	p := getFigure()
	c := vgpdf.New(5*vg.Inch, 5*vg.Inch)
	// c := vgtex.NewDocument(10*vg.Inch, 7.5*vg.Inch)

	p.Draw(draw.New(c))
	// c.FillString(p.Title.Font, vg.Point{2.5 * vg.Centimeter, 2.5 * vg.Centimeter}, "x")

	f, err := os.Create(currentFigure + ".pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err = c.WriteTo(f); err != nil {
		log.Fatal(err)
	}
	// log.Print("Figure generated .. ", currentFigure+".pdf")
}

func SaveEPS() {
	p := getFigure()
	c := vgeps.New(5*vg.Inch, 5*vg.Inch)
	// c := vgtex.NewDocument(10*vg.Inch, 7.5*vg.Inch)

	p.Draw(draw.New(c))
	// c.FillString(p.Title.Font, vg.Point{2.5 * vg.Centimeter, 2.5 * vg.Centimeter}, "x")

	f, err := os.Create(currentFigure + ".eps")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err = c.WriteTo(f); err != nil {
		log.Fatal(err)
	}
	// log.Print("Figure generated .. ", currentFigure+".eps")

}

func SetXlabel(x string) {
	if gcf != nil {

		gcf.X.Label.Text = x
		Save()
	}
}

func SetYlabel(y string) {
	if gcf != nil {
		gcf.Y.Label.Text = y
		Save()
	}
}
func SetLogY(m *vlib.MatrixF, cols ...int) (tmpfile string) {
	p := getFigure()
	p.Y.Scale = plot.LogScale{}
	return ""
}
func Plot(m *vlib.MatrixF, cols ...int) (tmpfile string) {
	// m.AppendColumn(ds.Col("distance").Float()).AppendColumn(ds.Col("PL").Float())
	var p *plot.Plot
	if gcf == nil {
		p = getFigure()
	} else {
		p = gcf
	}

	p.Add(plotter.NewGrid())

	if len(cols) == 0 {
		cols = []int{0, 1}
	}
	plotutil.AddLines(p, m.GetCols(cols...))
	// plotutil.AddScatters(p, m.GetCols(cols...))
	p.X.Label.Text = fmt.Sprint("Col ", cols[0])
	p.Y.Label.Text = fmt.Sprint("Col ", cols[1])
	p.Title.Text = currentFigure

	// m.AppendColumn(locations.X()).AppendColumn(locations.Y()).AppendColumn(pls)
	// pb, _ := plotter.NewBubbles(m, 0, 10)
	// p.Add(pb)
	// pb.Color = colorful.LinearRgb(1, 0, 0)
	Save()
	if !holdOn {
		gcf = nil
	}
	log.Print("Figure generated .. ", currentFigure+".png")

	return currentFigure + ".png"

}
