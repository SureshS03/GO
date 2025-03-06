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
