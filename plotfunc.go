package pf

import (
	"fmt"
	"log"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
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
			figures[currentFigure] = gcf
		}

	}

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
	if len(figures) == 0 || !holdOn {
		Fig(nextFigureName())
		return gcf
	}

	return gcf
}

func SetLogY(m *vlib.MatrixF, cols ...int) (tmpfile string) {
	p := getFigure()
	p.Y.Scale = plot.LogScale{}
	return ""
}
func Plot(m *vlib.MatrixF, cols ...int) (tmpfile string) {
	// m.AppendColumn(ds.Col("distance").Float()).AppendColumn(ds.Col("PL").Float())
	p := getFigure()
	tmpfile = currentFigure
	if len(cols) == 0 {
		cols = []int{0, 1}
	}
	plotutil.AddLines(p, m.GetCols(cols...))
	// plotutil.AddScatters(p, m.GetCols(cols...))
	p.X.Label.Text = fmt.Sprint("Col ", cols[0])
	p.Y.Label.Text = fmt.Sprint("Col ", cols[1])
	p.Add(plotter.NewGrid())
	p.Title.Text = currentFigure

	p.Save(10*vg.Inch, 7.5*vg.Inch, tmpfile+".png")

	log.Print("Figure generated .. ", tmpfile+".png")
	// m.AppendColumn(locations.X()).AppendColumn(locations.Y()).AppendColumn(pls)
	// pb, _ := plotter.NewBubbles(m, 0, 10)
	// p.Add(pb)
	// pb.Color = colorful.LinearRgb(1, 0, 0)

	return tmpfile
}
