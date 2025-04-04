package main

import "fmt"

func loopTest() {
	var myMap map[string]string = map[string]string{"name":"suresh","favGame":"valo","clg":"MGR"}

	//NOTE loop in map will return a random order everyTime its run
	for x:= range myMap { //the x will be the key in the map
		fmt.Println(x) //print only the x name, favGame, clg
	}

	for z, y:= range myMap {
		fmt.Printf("key = %v and value = %v\n",z, y) //z is key and y is value
	}

	var myarr = [3]uint16{56,34,12}
	//in array we dont get in random order, its only in index order
	for arr:= range myarr {
		fmt.Println(arr) //same this will print only the index
	}

	for index, value := range myarr {
		fmt.Printf("index = %v, value = %v\n",index, value)
	}

	for i:= 0; i<=10; i++ { //c like loop
		fmt.Println(i)
	}
	//go dont have while loop but we can achive it by for loop itself
	var i uint16 = 1
	
	for i <=10 { 
		i = i +1
	} //or else we can make the condition inside the in if statement with break key word

	//i++, i--, i+=10, i /=10, i *=10
	
}