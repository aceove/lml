package normalized

import (
	"gonum.org/v1/gonum/stat"
	"math"
	"sort"
)

func Linear(f []float64) []float64  {
	var col []float64
	for _,v:= range f{
		col=append(col,v)
	}
	sort.Float64s(col)
	max:=col[len(col)-1]
	min:=col[0]
	dev:=max-min
	var ret []float64
	for _,v:= range f{
		ret=append(ret,(v-min)/dev)
	}
	return ret
}
func StandardDeviation(f []float64) []float64  {
	//均值
	mean := stat.Mean(f, nil)
	//方差
	variance := stat.Variance(f, nil)
	//标准差
	stddev := math.Sqrt(variance)
	var col []float64
	for _,v:= range f{
		col=append(col,(v-mean)/stddev)
	}
	return col
}