package main // Package name
import "fmt"

type Student struct {
	Name          string
	Age           int
	IsAbsentToday bool
}

func main() {
	//var students []Student
	students := make([]Student, 0, 1000)
	//fmt.Printf()
	fmt.Printf("Capacity: %d, Length: %d", cap(students), len(students))
	//students= append(students,{"1", 1, true,})

	//println("------")
	//student := Student{
	//	Name:          "1st",
	//	Age:           12,
	//	IsAbsentToday: false,
	//}
	//fmt.Println("-----")
	//fmt.Println("-----")
	//fmt.Printf("%+v \n", student)
	////
	//stdGo := Test.Student{
	//	"12", 32,
	//}
	//fmt.Printf("%+v \n", stdGo)

}
