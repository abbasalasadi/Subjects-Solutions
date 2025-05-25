package main

import (
	"fmt"
	// "math"
)

type parameters interface{
	Area() float64
	Circomference() float64
}

type Rictangle struct{
	ShortSide float64
	LongSide float64
}

func (r Rictangle) Area() float64 {
	return r.LongSide * r.ShortSide
}

func (r Rictangle) Circomference() float64{
	return 2*(r.LongSide+r.ShortSide)
}

func main(){
	R:=Rictangle{ShortSide: 6, LongSide: 8}
	fmt.Printf("Rectangle L: %v, S: %v, Area: %v, Circomference: %v\n", R.LongSide, R.ShortSide, R.Area(), R.Circomference())
}