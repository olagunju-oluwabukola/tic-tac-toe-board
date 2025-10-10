package main

import (
	"fmt"
	"math"
)
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)

}

func fx(){
	hypot:= func(x,y float64) float64{
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(10,12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

//closures

func adder()func (int) int {
sum := 0
return func (x int) int {
	sum += x
	return sum
}
}

func cs (){
	pos, neg := adder(), adder()
	for i:=0; i<10; i++{
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}