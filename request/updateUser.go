package request

import (
	"time"
	"yuno/echo-example/model"
)

type UpdateStudentRequest struct {
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	BornDate  *time.Time `json:"born_date"`

	CountryId *int `json:"country_id"`
}

func UpdateStudentRequestToStudent(request UpdateStudentRequest, student *model.Student) {
	if request.FirstName != "" {
		student.FirstName = request.FirstName
	}
	if request.LastName != "" {
		student.LastName = request.LastName
	}
	if request.BornDate != nil {
		student.BornDate = *request.BornDate
	}
	if request.CountryId != nil {
		student.CountryId = *request.CountryId
	}
}
