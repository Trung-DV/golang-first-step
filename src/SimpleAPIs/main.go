package main

import (
	"SimpleAPIs/models"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	fmt.Println("Build APIs by Go")
	fmt.Println("------------------------------")
	// GET /
	e := echo.New()



	e.GET("/students", GetStudent)
	e.POST("/students", CreateStudent)
	e.DELETE("/students/:name", DeleteStudent)
	e.Logger.Fatal(e.Start(":1323"))
}

func DeleteStudent(context echo.Context) error {
	studentName := context.Param("name")
	hasStudent, index := models.TableStudents.HasStudentName(studentName)
	if hasStudent{
		models.TableStudents = append(models.TableStudents[:index], models.TableStudents[index+1:]...)

		return context.JSON(http.StatusOK, "Successful")

	}
	return context.JSON(http.StatusOK, "Unsuccessful")

}

func CreateStudent(context echo.Context) error {
	var createStudent models.Student
	err := context.Bind(&createStudent)
	if (nil != err) {
		return context.JSON(http.StatusOK, "Unsuccessful")
	}
	models.TableStudents = append(models.TableStudents, createStudent)
	return context.JSON(http.StatusOK, "Successful")

}
// Handler
func GetStudent(c echo.Context) error {
	//tStudents := models.TableStudents
	//tStudents := append(models.TableStudents, {"1", 12, true,})
	return c.JSON(http.StatusOK, models.TableStudents)
}