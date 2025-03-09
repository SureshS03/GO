package main

import (
	"fmt"
	//"strings"
	//"hello/bro"
)

func main() {
	fmt.Println("Hello, i'm Suresh")
	fmt.Println("-----------------")
	//var_declaration()
	printMe("Suresh")
	add(10, 20)
	mul, div, err := mulANDdiv(10, 5) // (10,0) or (0,10)
	// && and
	// || or
	// ! not
	if err != nil {
		fmt.Println(err.Error())
		return //return will stop the execution of the program
	}
	
	fmt.Printf("Multiplication: %v and the Divison is %v \n", mul, div)
	
	/*
	x:= "suresh"
	println("Length of the string is: ", len(x))
	const xx string = "γ"
	println(len(xx))
	*/
	
	testme2(10)
	testme3(0)
	arrayTest()
	sliceTest()

	//var text string = "Hello, World, World"
	//text2 := strings.Replace(text, "World", "Suresh", -1)
	//fmt.Println(text2)

}