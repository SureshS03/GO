package main

import (
	"fmt";
	//"errors";
)

func testme(num int16) {
	if num == 0 {
		fmt.Println("Number is zero")
	} else if num < 0 {
		fmt.Println("Number is negative")
	} else {
		fmt.Println("Number is positive")
	}
}

func testme2(num int16) {
	switch {
	case num == 0:
		fmt.Println("Number is zero")
	case num < 0:
		fmt.Println("Number is negative")
	default:
		fmt.Println("Number is positive")
	}
}

func testme3(num int16) {
	switch num { //this is switch condition statement 
	case 0:
		fmt.Println("Number is zero")
	case 1, 2, 3, 4, 5:
		fmt.Println("Number is between 1 and 5")
	default:
		fmt.Println("Number is greater than 5")
	}
}