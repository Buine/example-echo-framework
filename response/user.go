package request

import (
	"time"
	"yuno/echo-example/model"
)

type UserResponse struct {
	Id        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BornDate  time.Time `json:"born_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	CountryName string	`json:"country_name,omitempty"` 
	CountryISO string   `json:"country_iso,omitempty"`
}

func FromUserModel(student model.Student) UserResponse {
	return UserResponse{
		Id:        student.Id,
		FirstName: student.FirstName,
		LastName:  student.LastName,
		BornDate:  student.BornDate,
		CreatedAt: student.CreatedAt,
		UpdatedAt: student.UpdatedAt,
	}
}

func FromUserModelAndCountry(student model.Student, country model.Country) UserResponse {
	return UserResponse{
		Id:        student.Id,
		FirstName: student.FirstName,
		LastName:  student.LastName,
		BornDate:  student.BornDate,
		CreatedAt: student.CreatedAt,
		UpdatedAt: student.UpdatedAt,
		CountryName: country.FirstName,
		CountryISO: country.IsoName,
	}
}
