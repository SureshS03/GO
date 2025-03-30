package main

import (
	"fmt";
	"strings";
)

func TestStrings(){
	//see go usees UTF-8 encoding by default
	// which is 8 bit, for example a = 97 which is 01100001 in binary
	// and A = 65 which is 01000001 in binary
	// résumé here é = 233 in decimal which is 11101001 in binary
	// but the UTF-8, use x bytes for each character, so the é is 2 bytes and a = 1 byte
	// so the length of the string is 2 for é and 1 for a, so the length of the string is 3 bcz go calculates the length of the string in bytes not in characters
	//the UTF-8 have different formats for each bytes for example for 1 bytes it is 0xxxxxxx
	// for 2 bytes it is 110xxxxx 10xxxxxx the first byte is 110 and the second byte is 10, so the é is 2 bytes so it is 110"00011" 10"101001"
	// and also the string is immutable, so you can't change the string, you can only create a new string

	var word string = "résumé"
	fmt.Printf("The first letter: %v and %T\n", word[0], word[0]) // 114, string are represented as byte array, so the first letter is the first byte of the string and it is 114(uint8) in decimal which is r in ascii
	fmt.Printf("The sec letter: %v and %T\n", word[1], word[1]) // it is 233(uint8) in decimal which is é in ascii, but it is not a single byte, it is 2 bytes, so the first byte is 110"00011" and the second byte is 10"100101", so the first byte 11000011 is 195
	fmt.Printf("The thrid letter: %v and %T\n", word[2], word[2]) // this will represent the second byte of é which is 10101001 which is 169 in decimal, so the first byte is 195 and the second byte is 169

	for i, v:= range word{
		fmt.Printf("The %v letter: %v \n", i, v) // here index 2 and 7 will skip, and the range func will know that the é is 2 bytes, the v is going to have the 2 bytes
		fmt.Printf("The %v letter without use v: %v \n", i, word[i]) // but here we use word[i], so this will print the extact values of word[i] so if its 2 bytes means it will skip the second byte and give only the first byte
		fmt.Println("-----------------")
	}
	fmt.Println(len(word)) // there is only 6 letters in the string but the length is 8 because the é is 2 bytes

	// to work crtly on the string we need to convert it to rune,

	var runeWord []rune = []rune("résumé") // we use rune() conversion to convert the string to slice of rune
	fmt.Println(runeWord[1]) // rune will understand the tell the crt values of é 233
	fmt.Println("The length of the rune word is: ", len(runeWord)) // the length of the rune word is 6

	//String are inmmutable, so we can't change the string, we can only create a new string
	strslice := []string{"hello","im","suresh"}
	var joinedstr string = ""
	for x:= range strslice {
		joinedstr += strslice[x] // this will create a new variable for each loop like first x = hello,then x = hello + im , now x = helloim thenx = helloim + suresh
		//this will be take too much space, so we gonna use string package in go to do this
		fmt.Println(joinedstr)
	}
	fmt.Println("after the loop :", joinedstr)
	//joinedstr[1] = "cc" this will show error bcz strings are inmutable in go, so we cant change after created

	var stradder strings.Builder
	for c := range strslice {
		stradder.WriteString(strslice[c])
	}
	var joinedstr2 = stradder.String()
	fmt.Println("result is : ",joinedstr2)
	var xstr string = "oii" // new string
	var stradder2 strings.Builder // new builder
	stradder2.WriteString("some text") //added some text to builder buffer
	fmt.Println("before add the builder ",xstr) // this will be oi
	xstr = stradder2.String() //this will re write all the string  
	fmt.Println("after add the builder ",xstr)

	// we can do this
	var len , err = stradder.WriteString(joinedstr2)
	fmt.Printf("this is the len %v and this is the error %v\n",len , err)
}