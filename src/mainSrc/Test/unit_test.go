package Test

import "testing"

//type Student struct {
//	Name string
//	Age int
//	Absent bool
//}

func GetNamFromFile() string {
	return GetName()
}
func (s Student) GetNextAge() int {
	return s.Age + 2
}

func TestGetNextAge(t *testing.T) {
	std1 := Student {"Trung", 1, false}
	exp := 21
	if actual := std1.GetNextAge(); actual != exp {
		t.Error("failtest",actual, exp)
	}
}
func TestGetName(t *testing.T) {
	exp := "trung"
	if actual := GetName(); actual != exp {
		t.Error("failed test",actual, exp)
	}
}