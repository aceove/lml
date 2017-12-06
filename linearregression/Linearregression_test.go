package linearregression

import (
	"testing"
	"fmt"
	"github.com/CarefreeTour/ml/datatype"
	"github.com/CarefreeTour/ml/plot"
	"gonum.org/v1/gonum/mat"
	"github.com/CarefreeTour/ml/normalized"
)
func TestPlot(t *testing.T) {
	//数据读取
	r, _, data, err := datatype.CVS2Dense("ex1data1.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	xcol := make([]float64,r)
	ycol := make([]float64,r)
	xcol=mat.Col(xcol,0,data)
	ycol=mat.Col(ycol,1,data)

	//数据绘制
	fname := ""
	persist := false
	debug := true

	p, err := plot.NewPlotter(fname, persist, debug)
	if err != nil {
		err_string := fmt.Sprintf("** err: %v\n", err)
		panic(err_string)
	}
	defer p.Close()

	p.PlotXY(xcol,ycol, "ponit")
	p.SetXLabel("my x data")
	p.SetYLabel("my y data")
	//
	p.CheckedCmd("set grid")
	p.CheckedCmd("set title 'data'")
	p.CheckedCmd("set key top outside horizontal center")
	p.CheckedCmd("set terminal svg")
	p.CheckedCmd("set output 'data.svg'")
	p.CheckedCmd("replot")

	//p.CheckedCmd("q")


	xcolnor := normalized.StandardDeviation(xcol)
	ycolnor := normalized.StandardDeviation(ycol)


	p, err = plot.NewPlotter(fname, persist, debug)
	if err != nil {
		err_string := fmt.Sprintf("** err: %v\n", err)
		panic(err_string)
	}
	defer p.Close()

	p.PlotXY(xcolnor,ycolnor, "ponit")
	p.SetXLabel("my x data")
	p.SetYLabel("my y data")
	//
	p.CheckedCmd("set grid")
	p.CheckedCmd("set title 'normalized.StandardDeviation'")
	p.CheckedCmd("set key top outside horizontal center")
	p.CheckedCmd("set terminal svg")
	p.CheckedCmd("set output 'StandardDeviation.svg'")
	p.CheckedCmd("replot")



	xcolnorlinear := normalized.Linear(xcol)
	ycolnorlinear := normalized.Linear(ycol)


	p, err = plot.NewPlotter(fname, persist, debug)
	if err != nil {
		err_string := fmt.Sprintf("** err: %v\n", err)
		panic(err_string)
	}
	defer p.Close()

	p.PlotXY(xcolnorlinear,ycolnorlinear, "ponit")
	p.SetXLabel("my x data")
	p.SetYLabel("my y data")
	//
	p.CheckedCmd("set grid")
	p.CheckedCmd("set title 'normalized.Linear'")
	p.CheckedCmd("set key top outside horizontal center")
	p.CheckedCmd("set terminal svg")
	p.CheckedCmd("set output 'Linear.svg'")
	p.CheckedCmd("replot")


	p.CheckedCmd("q")



}
