package models

type Student struct {
	Name string // `json: "name"` decode to Student
	Age int		// `json: "my_age"`
}

type Students []Student

func (students Students) HasStudentName(s string) (bool, int){
	for index, student := range students {
		if student.Name == s {
			// Found!
			return  true, index
		}
	}
	return false, 0

}

var TableStudents Students
