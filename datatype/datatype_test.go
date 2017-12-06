package datatype

import (
	"testing"
	"fmt"
)
func TestCVS2Float64Array(t *testing.T) {
	r, c, data, err := CVS2Float64Array("ex1data1.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(r, c)
	fmt.Println(data)
}
func TestCVS2Dense(t *testing.T) {
	r, c, data, err := CVS2Dense("ex1data1.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(r, c)
	fmt.Println(data)
}
