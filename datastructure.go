package main

import (
	"fmt"
	//"errors"
)

func arrayTest(){
	var arr [3]int32 //array declaration by normal and it will be initialized with 0
	//	var arr [3]int32 = [3]int32{10, 20, 30} //array declaration by normal and it will be initialized with values
	fmt.Println("this is before add the values in array",arr) //three zeros will be printed
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	fmt.Println("this is after add the values in array",arr) //10,20,30 will be printed

	arr2 := [3]int32{10, 20, 30} //this is short hand declaration of array and it will be initialized with the given values
	fmt.Println("arr2 =",arr2)

	arr3 := [...]int32{10, 20, 30, 40, 50} //this is also short hand declaration but it will automatically calculate the size of the array
	fmt.Println("arr3 =",arr3)

	//to know the memory address of the array
	fmt.Printf("Memory address of arr3 is %v\n", &arr3) //& is used to get the memory address of the variable
	fmt.Printf("Memory address of 0 is %v\n", &arr3[0]) //the 0 is the index of the array is same as the memory address of the array
	fmt.Printf("Memory address of 1 is %v\n", &arr3[1]) //the 1 is the index of the array is +4 bytes bcz we use 32bit means 4 x 8 bytes (1 bit = 8 bytes) from the memory address of the 0 index
	fmt.Printf("Memory address of 2 is %v\n", &arr3[2]) //the 2 is the index of the array is +4 bytes from the memory address of the 1 index
}

func sliceTest(){
	//by avoiding the size of the array we can use slice
	var slice []int32 = []int32{10, 20, 30} //slice declaration by normal
	fmt.Println("slice =",slice)
	fmt.Printf("before append the slice length is %v and the capacity is %v\n", len(slice), cap(slice))
	slice = append(slice, 40) //append is used to add the value in the slice
	//append will add the value in the slice and if the slice is full then it will create the new slice with increased the size of the old slice and copy the old slice values to the new slice
	//so we have give a new variable to the slice to store the new slice
	slice = append(slice, 50, 60, 70) //we can add multiple values in the slice
	fmt.Println("slice 0 to 2 =",slice[0:3]) //we can get the values from the slice by using the index
	fmt.Println("slice[5:]=",slice[5:]) //we can get the values from the slice by using the index
	fmt.Println("slice[:5]=",slice[:5]) //we can get the values from the slice by using the index
	fmt.Println("slice[6] =",slice[6]) //we can get the value from the slice by using the index
	fmt.Println("slice =",slice)
	fmt.Printf("after append the slice length is %v and the capacity is %v\n", len(slice), cap(slice))

	slice2 := []int32{10, 20, 30} //slice declaration by short hand
	fmt.Println("slice2 =",slice2)
	both := append(slice, slice2...) //we can add the two slices by using the ... operator
	fmt.Println("both =",both)

	slice3 := make([]int8, 3, 5) //make is used to create the slice with the given length and capacity
	fmt.Println("slice3 =",slice3)
	fmt.Printf("before append the slice3 length is %v and the capacity is %v\n", len(slice3), cap(slice3))
	//slice3[0] = 10 or
	slice3 = []int8{10, 20, 30, 40, 50} //we can add the values in the slice
	fmt.Println("slice3 =",slice3)

	strslice := []string{"hello", "im"}

	fmt.Println("strslice =",strslice)
	strslice = append(strslice, "suresh")
	fmt.Println("strslice =",strslice)
	//all opertain can be done in the slice of string

	//boolslice := []bool{true, false}
	//fmt.Println("boolslice =",boolslice)

	//floatslice := []float32{10.5, 20.5}

	//uintslice := []uint{10, 20}

	byteSlice := []byte{'a', 'b'} //byte values should be in single quotes
	//byte values is not a, b it is 97, 98 because byte is a number
	byteSlice = append(byteSlice, "hello"...) //we can add the string in the byte slice by using the ... operator
	//here the string will be converted into the byte slice and each character will be converted into the byte value
	//because go works in ascii values
	//ascii value of h is 104, e is 101, l is 108, l is 108, o is 111
	fmt.Println("byteSlice =",byteSlice)

	//runeSlice := []rune{'a', 'b'}
	//rune values should be in single quotes
	//same as the byte slice but rune is used for the unicode8 values

}

func mapTest(){
	var myMap map[string]string = map[string]string{"name":"suresh","favGame":"valo","clg":"MGR"}
	fmt.Println("Mymap = ", myMap)
	
	//or else
	var myMap2 = map[string]bool{"haveName":true,"havePet":false}
	fmt.Println(myMap2)
	
	myMap3 := map[int]string{5:"suresh"} // a key can be anything ig
	fmt.Println(myMap3)
	//fmt.Println(myMap3[5])

	//to access it
	fmt.Println(myMap["name"])

	//maps in go always gives a value, even its not there means its gives its default value i.e 0 for int, empty space for string
	fmt.Println(myMap["clg"])

	//we can handle the above problem by decleare a error handler when we accessing it
	var name, ok = myMap["name"]
	if ok{
		fmt.Printf("my name is %v\n",name)
	} else {
		fmt.Println("i dont have a name")
	}

	var age, ageok = myMap["age"]

	if ageok {
		fmt.Printf("my age is %v\n",age)
	} else {
		fmt.Println("i dont have a age")
	}

	//to delete something in map
	delete(myMap3, 5) //1st is variable, 2nd is key
	fmt.Println(myMap3)

}

func memoryallocation() {
	//slice in go are reference type which means while u copy a slice to another slice it will just copy the data's element pointer
	//its not create a new memory space for the data but it will create a use space for the slice not for the data's
	//so if u change the slice means i will effect the another slice to , so use make([]int, len()) and copy() for create a new slice with new data element space 

	var list []int8 = []int8{1, 2, 3, 4, 5}
	var copylist = list //get e data's pointer value
	fmt.Println("list before change the copy list",list)
	copylist[2] = 69 // affect both slice
	fmt.Printf("list %v address is\n",list,)
	fmt.Printf("copy list %v address is\n",copylist)
	fmt.Printf("addres of list %p\n",&list) // this slice stored in different place
	fmt.Printf("addres of copy list %p\n",&copylist) // and slice stored in new memory but
	//but both uses same data's element pointers 
	for x := range list {
		fmt.Printf("Value: %d, Address of list[%d]: %v, Address of x: %v\n", list[x], x, &list[x], &x)
	}
	for x := range list {
		fmt.Printf("Value: %d, Address of copylist[%d]: %v, Address of x: %v\n", copylist[x], x, &copylist[x], &x)
	}

	// as we can see the output of the both loop that both the data's are stored in the same memory location but the pointers are stored in the different memory 

	//but the array in go are value type only, so its work normally creatinig a new space for copied array

}