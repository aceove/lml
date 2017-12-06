package cost

import (
	"testing"
	"github.com/CarefreeTour/ml/datatype"
	"github.com/CarefreeTour/ml/plot"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 50                  // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func TestBatchGradientDescent(t *testing.T) {
	fname := ""
	persist := false
	debug := true
	p, err := plot.NewPlotter(fname, persist, debug)
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
	xcol := make([]float64, r)
	ycol := make([]float64, r)
	xcol = mat.Col(xcol, 1, data)
	ycol = mat.Col(ycol, 2, data)

	r, _, bigdata, _ := datatype.CVS2Float64Array("ex1data2.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//array([-2.28286727,  1.03099898])

	theta := []float64{2, 3}
	step, cost := BatchGradientDescent(theta, bigdata, 0.001, 0.00001, 500)
	l := len(cost)
	cxcol := make([]float64, l)
	for i := 0; i < l; i++ {
		cxcol[i] = float64(i)
	}
	fmt.Println(step[len(step)-1])

	//
	p.CheckedCmd("set multiplot")
	p.PlotXY(cxcol, cost, "gd")
	p.SetXLabel("cost")
	p.SetYLabel("epoch")
	p.CheckedCmd("set grid")
	//p.CheckedCmd("plot 0.44103227803851236 + 0.6615484170577665*x")
	p.CheckedCmd("set title 'gd'")
	p.CheckedCmd("set key top outside horizontal center")
	p.CheckedCmd("set terminal svg")
	p.CheckedCmd("set output 'cost.svg'")
	p.CheckedCmd("replot")
	p.CheckedCmd("q")

	//fmt.Println(cost)

	/*
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
	*/
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
