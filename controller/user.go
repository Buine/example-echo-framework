package controller

import (
	"fmt"
	"net/http"
	"yuno/echo-example/database"
	"yuno/echo-example/model"
	request "yuno/echo-example/request"
	response "yuno/echo-example/response"

	"github.com/labstack/echo/v4"
)

func GetStudents(context echo.Context) error {
	connection := database.GetConnection()
	students := new([]model.Student)

	connection.Preload("Country").Find(&students)
	var listStudentsResponse = []response.UserResponse{}
	for _, student := range *students {
		listStudentsResponse = append(listStudentsResponse, response.FromUserModelAndCountry(student, student.Country))
	}

	return context.JSON(http.StatusOK, listStudentsResponse)
}

func CreateStudent(context echo.Context) error {
	requestStudent := new(request.CreateStudentRequest)
	if err := context.Bind(&requestStudent); err != nil {
		return err
	}

	student := request.CreateStudentRequestToStudent(*requestStudent)
	fmt.Println(requestStudent)
	connection := database.GetConnection()
	connection.Save(&student)
	country := new(model.Country)
	connection.First(&country, student.CountryId)
	context.JSON(http.StatusCreated, response.FromUserModelAndCountry(student, *country))

	return nil
}

func DeleteStudent(context echo.Context) error {
	studentId := context.Param("student")
	if studentId == "" {
		return context.JSON(http.StatusBadRequest, map[string]string{
			"code":    "error.InvalidInput",
			"message": "Parameter 'studentId' is empty or invalid.",
		})
	}

	connection := database.GetConnection()
	connection.Unscoped().Delete(&model.Student{}, studentId)

	return nil
}