package main

import (
	"fmt"
)

//struct are user defined structured things which can be anything
// Define student struct
type student struct {
	name    string
	id      uint16
	subject sub //add struct as field
	mark    uint8
}

// Define teacher struct
type teacher struct {
	name          string
	studentsTotal uint16
	studentsPass  uint16
}

// Define subject struct
type sub struct {
	name  string
	total uint8
}

//methods are the function which is work with the struct
// Implement markpercentage for student (satisfies ExamResult)
func (x student) markpercentage() float64 {
	return float64(x.mark) / float64(x.subject.total) * 100
}

// Implement markpercentage for teacher (satisfies ExamResult)
func (t teacher) markpercentage() float64 {
	return float64(t.studentsPass) / float64(t.studentsTotal) * 100
}

// Define an interface
//interface its use to achieve polymorphism in Go
//Interface is a set of methods signatures,any type have these methods can access the interface and function that uses the interface
type ExamResult interface {
	markpercentage() float64 //any type or struct have this method can access the printResult func bcz they falls under the same interface
	//interface methods should defined by its return types 
}

// Function that accepts anything implementing ExamResult
func printResult(e ExamResult) {
	fmt.Printf("Percentage: %.2f%%\n", e.markpercentage())
}

func textstruct() {
	// Create a student
	var stud = student{"Suresh", 114, sub{"Maths", 100}, 50}

	// Create a teacher
	var teach = teacher{"Mr. Kumar", 200, 180}

	// Print student details
	fmt.Println(stud) 
	fmt.Println(stud.name, stud.id, stud.subject.name)

	// Print teacher details
	//fmt.Println(teach)
	//fmt.Println(teach.name, "teaches", teach.studentsTotal)

	// Call printResult with student
	printResult(stud)

	// Call printResult with teacher
	printResult(teach)

	//we can also create private struct we declare the values straight ahead like this
	var exam = struct{
		examName string
		totalMark uint8
	}{"maths",100}
	fmt.Println(exam)
	//we cant reused this bcz its private or anonymous, if u want to use this means u have to re write the struct
}
