package main

import "fmt"

/*
	Variable declaration
	----------------------------
	var myvar
	const myconst

	bool
	string
	rune
	int, int16, int32, int64
	uint, uint16, uint32, uint64 (unsigned)
	float32, float64

*/

func var_declaration() {
	pl := fmt.Println
	var a int32 = 10 //can be 64, 32, 16,
	//a:=10 //short hand declaration
	pl("Value of a is: ", a)
	var myconst int = 100
	pl("Value of myconst is: ", myconst)
	var b int64 = 20
	c:=30
	pl("Value of b & cis: ", b, c)
	var name string = "Suresh" //or `Suresh`
	pl("My name is: ", name)
	var d float32 = 10.5 //can be 64, 32
	pl("Value of d is: ", d)
	var chnagedToint int = int(d) //float32(int)
	pl("Value of chnagedToint is: ", chnagedToint)
	var e bool = true
	pl("Value of e is: ", e)		

}