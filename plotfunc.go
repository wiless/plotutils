package pf

import (
	"fmt"
	"image/color"
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

var plotcurves map[string]int
var currentFigure string
var err error
var gcf *plot.Plot
var figurestack []string
var figurestackIndx map[string]int
var CurrentIndex int

type Figure struct {
	*plot.Plot
}

func init() {
	figures = make(map[string]*plot.Plot)
	figurestackIndx = make(map[string]int)
	plotcurves = make(map[string]int)
}

func Fig(str ...string) {

	if len(str) == 0 {
		// if empty argument
		currentFigure = nextFigureName()
		gcf, err = plot.New()
		figures[currentFigure] = gcf
		figurestack = append(figurestack, currentFigure)
		figurestackIndx[currentFigure] = len(figurestack) - 1
		CurrentIndex = len(figurestack) - 1
	} else {
		/// Ignoring more than one args..
		var isok bool
		gcf, isok = figures[str[0]]
		currentFigure = str[0]
		CurrentIndex = figurestackIndx[str[0]]
		if !isok {
			// Create the figure
			gcf, err = plot.New()
			figurestack = append(figurestack, currentFigure)

			figurestackIndx[currentFigure] = len(figurestack) - 1
			log.Print("Creating Figure ", currentFigure)
			figures[currentFigure] = gcf
			CurrentIndex = len(figurestack) - 1
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

func getLastFigure() *plot.Plot {
	if len(figures) == 0 {
		return nil
	}

	var isok bool
	gcf, isok = figures[currentFigure]
	CurrentIndex = figurestackIndx[currentFigure]
	if !isok {
		CurrentIndex = -1
		return nil

	}

	return gcf
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

func ShowPrev() {
	log.Print(currentFigure, figurestack, "CurrentIndex = ", CurrentIndex)
	if CurrentIndex == 0 {
		CurrentIndex = len(figurestack) - 1
		Fig(figurestack[CurrentIndex])
	} else {
		CurrentIndex--
		Fig(figurestack[CurrentIndex])
	}
	Save()
}

func ShowNext() {
	log.Print(currentFigure, figurestack, "CurrentIndex = ", CurrentIndex)
	if CurrentIndex == len(figurestack)-1 {
		CurrentIndex = 0
		Fig(figurestack[CurrentIndex])
	} else {
		CurrentIndex++
		Fig(figurestack[CurrentIndex])
	}
	Save()
}

func Save() {
	if gcf == nil {
		return
	}
	// gcf.Title.TextStyle.Color = colornames.Blue

	// log.Print("Figure generated .. ", currentFigure+".png")
	// gcf.Save(10*vg.Inch, 7.5*vg.Inch, currentFigure+".png")

	ShowX11()
	// SavePng()

}

func SavePng() {
	if gcf == nil {
		return
	}
	gcf.Title.TextStyle.Color = colornames.Blue

	log.Print("Figure generated .. ", currentFigure+".png")
	gcf.Save(10*vg.Inch, 7.5*vg.Inch, currentFigure+".png")
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

func Legends(ls ...string) {
	for _, s := range ls {
		p := getLastFigure()
		quad := plotter.NewFunction(func(x float64) float64 { return x * x })
		quad.Color = color.RGBA{B: 255, A: 255}
		quad.Width = vg.Points(2)

		p.Legend.Add(s, quad)

		p.Legend.Padding.Dots(100)
		p.Legend.ThumbnailWidth = vg.Inch * .80

	}
	Save()
}

func SetXlabel(x string) {
	p := getLastFigure()
	p.X.Label.Text = x
	Save()
}

func SetXLim(xmin, xmax float64) {
	p := getLastFigure()
	p.X.Min, p.X.Max = xmin, xmax
	Save()
}
func SetYLim(ymin, ymax float64) {
	p := getLastFigure()
	p.Y.Min, p.Y.Max = ymin, ymax
	Save()
}

func SetYlabel(y string) {
	p := getLastFigure()

	p.Y.Label.Text = y
	Save()

}
func SetLogY(m *vlib.MatrixF, cols ...int) (tmpfile string) {
	p := getFigure()
	p.Y.Scale = plot.LogScale{}
	return ""
}

func PolyGon(v plotter.XYer) {
	var p *plot.Plot
	if gcf == nil {
		p = getFigure()
	} else {
		p = gcf
	}

	h, err := plotter.NewPolygon(v)
	if err != nil {
		log.Print("Error Polygon ", err)
	}
	// h.Normalize(1)
	p.Add(h)
	p.X.Max *= 1.10
	p.Add(plotter.NewGrid())

	Save()
	if !holdOn {
		gcf = nil
	}

}

func Hist(v vlib.VectorF, bins ...int) string {

	var p *plot.Plot
	if gcf == nil {
		p = getFigure()
	} else {
		p = gcf
	}

	h, err := plotter.NewHist(v, 10)

	if err != nil {
		log.Print("Error Histogram ", err)
	}
	// h.Normalize(1)

	p.Add(h)

	Save()
	if !holdOn {
		gcf = nil
	}

	return currentFigure + ".png"
}

func PlotXY(x, y vlib.VectorF) (tmpfile string) {
	if x.Len() != y.Len() {
		log.Print("PlotXY : Length of x & y not same..")
		return ""
	}
	// m.AppendColumn(ds.Col("distance").Float()).AppendColumn(ds.Col("PL").Float())
	var p *plot.Plot
	if gcf == nil {
		p = getFigure()
	} else {
		p = gcf
	}
	var m vlib.MatrixF
	log.Print(x.Len(), y.Len())

	m.AppendColumn(x).AppendColumn(y)
	p.Add(plotter.NewGrid())

	plotcurves[currentFigure] = plotcurves[currentFigure] + 1
	// nxtclr := plotutil.Color(plotcurves[currentFigure])

	// plotutil.AddLines(p, m.GetCols(cols...))
	// plines, _ := plotter.NewScatter(m.GetCols(cols...))
	plotutil.AddLines(p, m)
	// // plines, _, _ := plotter.NewLinePoints(m)
	// //
	// // plines.Color = nxtclr
	// p.Add(plines)

	// plotutil.AddLines(p, m.GetCols(0, 2))
	// plotutil.AddScatters(p, m.GetCols(cols...))
	// p.X.Label.Text = fmt.Sprint("Col ", cols[0])
	// p.Y.Label.Text = fmt.Sprint("Col ", cols[1])
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

	plotcurves[currentFigure] = plotcurves[currentFigure] + 1
	nxtclr := plotutil.Color(plotcurves[currentFigure])

	// plotutil.AddLines(p, m.GetCols(cols...))
	// plines, _ := plotter.NewScatter(m.GetCols(cols...))
	plines, _, _ := plotter.NewLinePoints(m.GetCols(cols...))

	plines.Color = nxtclr
	p.Add(plines)

	// plotutil.AddLines(p, m.GetCols(0, 2))
	// plotutil.AddScatters(p, m.GetCols(cols...))
	// p.X.Label.Text = fmt.Sprint("Col ", cols[0])
	// p.Y.Label.Text = fmt.Sprint("Col ", cols[1])
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

func Scatter(xys plotter.XYer) (tmpfile string) {
	// m.AppendColumn(ds.Col("distance").Float()).AppendColumn(ds.Col("PL").Float())
	var p *plot.Plot
	if gcf == nil {
		p = getFigure()
	} else {
		p = gcf
	}

	p.Add(plotter.NewGrid())

	plotcurves[currentFigure] = plotcurves[currentFigure] + 1
	nxtclr := plotutil.Color(plotcurves[currentFigure])

	// plotutil.AddLines(p, m.GetCols(cols...))
	plines, _ := plotter.NewScatter(xys)

	plines.Color = nxtclr
	p.Add(plines)

	// plotutil.AddLines(p, m.GetCols(0, 2))
	// plotutil.AddScatters(p, m.GetCols(cols...))
	// p.X.Label.Text = fmt.Sprint("Col ", cols[0])
	// p.Y.Label.Text = fmt.Sprint("Col ", cols[1])
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

// function DrawPolyGon(centre,radius float64 )
// // if ~exist('fmt')
// //     fmt='-r';
// // end
//
// // for k=1:length(centre)
// in:=vlib.VectorF{0,1,2,3,4,5}
//
//  for theta:=range in {
// 	 x,y:=cmplx.Rect( radius,theta)
// 	 x,y=pol2cart([0:5]*pi/3+pi/6,radius);
// 	 x=x+real(centre(k));
// 	 y=y+imag(centre(k));
// }
// 	 x(end+1)=x(1);
// 	 y(end+1)=y(1);
//
//
//  }
//  plot(x,y,fmt,'LineWidth',2) ;hold on;
// end
