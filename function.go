package main

import (
	"fmt";
	"errors";
)

func PrintMe(name string) {
	fmt.Println(name)
}

func add(a, b uint) uint { //uint is return type
	x := a + b
	println(x)
	return x
}

func mulANDdiv(a, b int) (int, int, error) { //we can return multiple values
	if b == 0 {
		return 0, 0, errors.New("division by zero")
	} else if a == 0 {
		return 0, 0, errors.New("mul by zero")
	} else {
		return a * b, a / b, nil //dont have to specify the nil values bcz by default it is nil and 0 for all int and float and rune values '' for string
	}
}