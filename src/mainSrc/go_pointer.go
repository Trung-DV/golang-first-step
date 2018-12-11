package main

import "fmt"
type Student struct {
	Name string
	Age int
	Absent bool
}
func main() {
	studentA := Student{"main", 12, true}
	fmt.Printf("%+v ", studentA)

	//var pStudent Student
	var p1Student *Student
	p1Student = &studentA
	fmt.Println(p1Student)
	PrintHelloFunc()
	PrintNumber(3.14523)

}

func PrintHelloFunc() {
	fmt.Println("Hello from func")
}
func PrintNumber(number float32 ) {
	fmt.Printf("Print number %f",  number)
}

func getPiNumber() float32 {
	return 3.14
}

//func (s* Student) test() {
//
//}