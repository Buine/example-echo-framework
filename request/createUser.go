package request

import (
	"time"
	"yuno/echo-example/model"
)

type CreateStudentRequest struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BornDate  time.Time `json:"born_date"`

	CountryId int `json:"country_id"`
}


func CreateStudentRequestToStudent(request CreateStudentRequest) model.Student {
	return model.Student{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		BornDate:  request.BornDate,
		CountryId: request.CountryId,
	}
}