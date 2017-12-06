package plot

import (
	"testing"
	"math"
	"image/color/palette"
	"fmt"
	"github.com/CarefreeTour/ml/datatype"
)

func TestDraw2DFuncAuto(t *testing.T){
	p:=NewPlot("test","x","y")
	Draw2DFuncAuto(p,func(x float64) float64 { return x * x },"x1")
	Draw2DFuncAuto(p,func(x float64) float64 { return math.Pow(2, x) },"x2")
	Draw2DFuncAuto(p,func(x float64) float64 { return 10*math.Sin(x) + 50 },"x3")
	Paint(p,0,10,0,100,"Draw2DAuto.png")
}

func TestDraw2DFunc(t *testing.T){
	p:=NewPlot("test","x","y")
	c :=palette.WebSafe[28]
	Draw2DFunc(p,func(x float64) float64 { return x * x },c,"x1")
	c =palette.WebSafe[128]
	Draw2DFunc(p,func(x float64) float64 { return math.Pow(2, x) },c,"x2")
	c =palette.WebSafe[200]
	Draw2DFunc(p,func(x float64) float64 { return 10*math.Sin(x) + 50 },c,"x3")
	Paint(p,0,10,0,100,"Draw2D.png")
}
func TestDraw2DPointAuto(t *testing.T) {
	r, c, data, err := datatype.CVS2Float64Array("ex1data1.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(r, c)
	fmt.Println(data)
	p:=NewPlot("test","x","y")
	Draw2DFuncAuto(p,func(x float64) float64 { return x * x },"x1")
	Draw2DFuncAuto(p,func(x float64) float64 { return math.Pow(2, x) },"x2")
	Draw2DFuncAuto(p,func(x float64) float64 { return 10*math.Sin(x) + 50 },"x3")
	Draw2DPointAuto(p,data,"point")
	Paint(p,0,0,0,0,"TestDraw2DPointAuto.png")

}