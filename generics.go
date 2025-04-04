package main

import "fmt"

func testGenerics() {
	var xint uint8 = 5
	var yint uint8 = 6

	fmt.Println("Sum of x and y is: ", addTwoInt(xint, yint))

	var xfloat float64 = 5.2
	var yfloat float64 = 3.8

	fmt.Println("Sum of x and y is: ", addTwoFloat(xfloat, yfloat))
	fmt.Println("sum of x and y int using genrics", addTwo(xint, yint))
	//fmt.Println("sum of x and y float using genrics", addTwo(xfloat, yfloat))
}

func addTwoInt(a, b uint8) uint8 {
	return a + b
}

func addTwoFloat(a, b float64) float64 {
	return a + b
}

func addTwo[bro uint8 | float32](a, b bro) bro {
	return a + b
}