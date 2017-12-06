package plot

import (
	"testing"
	"fmt"
	"math"
	"github.com/CarefreeTour/ml/datatype"
	"gonum.org/v1/gonum/mat"
)

func TestGunPlot(t *testing.T) {
	fname := ""
	persist := false
	debug := true
	p,err := NewPlotter(fname, persist, debug)
	if err != nil {
		err_string := fmt.Sprintf("** err: %v\n", err)
		panic(err_string)
	}
	defer p.Close()
	p.CheckedCmd("set samples 500")
	p.CheckedCmd("unset key")
	p.CheckedCmd("plot sin(5*x)'")
	p.CheckedCmd("set terminal jpeg")
	p.CheckedCmd("set output 'TestGunPlot.jpg'")
	p.CheckedCmd("replot")
	p.CheckedCmd("q")
}

func TestGunPlotX(t *testing.T) {
	fname := ""
	persist := false
	debug := true

	p, err := NewPlotter(fname, persist, debug)
	if err != nil {
		err_string := fmt.Sprintf("** err: %v\n", err)
		panic(err_string)
	}
	defer p.Close()

	p.PlotX([]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "some data - x")
	p.SetXLabel("my x data")
	p.SetYLabel("my y data")
	p.CheckedCmd("set terminal jpeg")
	p.CheckedCmd("set output 'TestGunPlotX.jpg'")
	p.CheckedCmd("replot")

	p.CheckedCmd("q")
}

func TestGunPlotXY(t *testing.T) {
	fname := ""
	persist := false
	debug := true

	p, err := NewPlotter(fname, persist, debug)
	if err != nil {
		err_string := fmt.Sprintf("** err: %v\n", err)
		panic(err_string)
	}
	defer p.Close()

	p.PlotXY([]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "some data - x/y")
	p.SetXLabel("my x data")
	p.SetYLabel("my y data")
	p.CheckedCmd("set terminal svg")
	p.CheckedCmd("set output 'TestGunPlotXY.svg'")
	p.CheckedCmd("replot")

	p.CheckedCmd("q")
}


func TestGunPlotXYZ(t *testing.T) {
	fname := ""
	persist := false
	debug := true

	p, err := NewPlotter(fname, persist, debug)
	if err != nil {
		err_string := fmt.Sprintf("** err: %v\n", err)
		panic(err_string)
	}
	defer p.Close()

	p.CheckedCmd("set grid x")
	p.CheckedCmd("set grid y")
	p.CheckedCmd("set grid z")
	p.PlotXYZ(
		[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		"test 3d plot")
	p.SetLabels("x", "y", "z")
	p.CheckedCmd("set terminal pdf")
	p.CheckedCmd("set output 'plot005.pdf'")
	p.CheckedCmd("replot")

	p.CheckedCmd("q")
}

func TestSPlot(t *testing.T)  {
	fname := ""
	persist := false
	debug := true

	p, err := NewPlotter(fname, persist, debug)
	if err != nil {
		err_string := fmt.Sprintf("** err: %v\n", err)
		panic(err_string)
	}
	defer p.Close()

	p.SetPlotCmd("splot")
	p.SetStyle("pm3d")

	p.PlotXY([]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "some data - x/y")
	p.SetXLabel("my x data")
	p.SetYLabel("my y data")
	p.CheckedCmd("set terminal pdf")
	p.CheckedCmd("set output 'plot003.pdf'")
	p.CheckedCmd("replot")

	p.CheckedCmd("q")
}

func TestNdPlot(t *testing.T)  {
	fname := ""
	persist := false
	debug := true

	p, err := NewPlotter(fname, persist, debug)
	if err != nil {
		err_string := fmt.Sprintf("** err: %v\n", err)
		panic(err_string)
	}
	defer p.Close()

	p.CheckedCmd("set grid x")
	p.CheckedCmd("set grid y")
	p.CheckedCmd("set grid z")
	p.PlotNd(
		"test Nd plot",
		[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	p.SetLabels("x", "y", "z")
	p.CheckedCmd("set terminal pdf")
	p.CheckedCmd("set output 'plot006.pdf'")
	p.CheckedCmd("replot")

	p.CheckedCmd("q")
}

func TestFuncPlot(t *testing.T) {
	fname := ""
	persist := false
	debug := true

	p, err := NewPlotter(fname, persist, debug)
	if err != nil {
		err_string := fmt.Sprintf("** err: %v\n", err)
		panic(err_string)
	}
	defer p.Close()

	p.SetStyle("steps")
	p.PlotFunc([]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		func(x float64) float64 { return math.Exp(float64(x) + 2.) },
		"test plot-func")
	p.SetXLabel("my x data")
	p.SetYLabel("my y data")
	p.CheckedCmd("set terminal pdf")
	p.CheckedCmd("set output 'plot004.pdf'")
	p.CheckedCmd("replot")

	p.CheckedCmd("q")
}

func TestMlPlot(t *testing.T) {
	fname := ""
	persist := false
	debug := true
	p, err := NewPlotter(fname, persist, debug)
	if err != nil {
		err_string := fmt.Sprintf("** err: %v\n", err)
		panic(err_string)
	}
	defer p.Close()
	r, _, data, err := datatype.CVS2Dense("ex1data1.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	xcol := make([]float64,r)
	ycol := make([]float64,r)
	xcol=mat.Col(xcol,0,data)
	ycol=mat.Col(ycol,1,data)
	p.PlotXY(xcol,ycol, "ponit")
	p.SetXLabel("my x data")
	p.SetYLabel("my y data")
	//
	p.CheckedCmd("set grid")
	p.CheckedCmd("set title 'Bessel Functions of the First Kind'")
	p.CheckedCmd("set key top outside horizontal center")
	p.CheckedCmd("set terminal svg")
	p.CheckedCmd("set output 'fangjia.svg'")
	p.CheckedCmd("replot")
	p.CheckedCmd("q")
}