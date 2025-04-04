package main

import (
	"fmt"
)

func testPoniters() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	var list = [5]int{1, 2, 3, 4, 5} // 1 time allocation
	fmt.Printf("list address %p\n", &list)
	fmt.Println("list in normal func", list)
	var res[5]int = addfive(&list) // pass the address and create a new array so 2 time allocation
	var res1[5]int = addfivewihoutpointer(list) // pass the address and create a new array so 2 time allocation
	fmt.Println("list in normal func after add five", list)
	fmt.Println("list in pointer func", res)
	fmt.Println("list in pointer func without pointer", res1) // this will be 11 12 13 14 15 bcz first its changed to 6 7 8 9 10 (addfive func) and then added 5 to it
	fmt.Println("list in normal func after add five without pointer", list)
	fmt.Printf("list address res %p\n", &res)
	fmt.Printf("list address res 1 %p\n", &res1)
}

func addfive(xx *[5]int)  [5]int{
	//fmt.Println("list in add five", xx)
	//if we do not use pointer then it will create a new array and the address will be different
	// and the value will not be changed in the original array
	fmt.Printf("list address in add five %p\n", xx)
	for i:= range xx{
		fmt.Printf("list in add five %v amd %p\n", xx[i], &xx[i])
		xx[i] += 5
	}
	return *xx
}

func addfivewihoutpointer(xx [5]int)  [5]int{
	// here we dont use pointer so it will create a new array and the address will be different
	// and the value will not be changed in the original array
	fmt.Printf("list address in add five without pointer %p\n", &xx)
	for i:= range xx{
		xx[i] += 5
	}
	return xx
}