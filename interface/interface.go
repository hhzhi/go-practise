package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
	Des() float64
}

type V struct {
	x, y float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (f MyFloat) Des() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v V) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y + v.y)
}

func (v *V) Des() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	var a Abser
	// f := MyFloat(-math.Sqrt2)
	v := V{3, 4}

	// a = f
	a = &v

	fmt.Println(v.x)
	fmt.Println(a.Des())
}
