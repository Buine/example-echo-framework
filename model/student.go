package model

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Id        int       `json:"id" gorm:"column:id;primaryKey"`
	FirstName string    `json:"first_name" gorm:"column:first_name"`
	LastName  string    `json:"last_name" gorm:"column:last_name"`
	BornDate  time.Time `json:"born_date" gorm:"column:born_date"`
	CountryId int       `json:"country_id" gorm:"column:country_id"`
	Country   Country   `json:"country" gorm:"foreignKey:country_id"`
}
