package model

import "gorm.io/gorm"

type Country struct {
	gorm.Model
	Id        int     `json:"id" gorm:"column:id;primaryKey"`
	FirstName string  `json:"name" gorm:"column:name"`
	IsoName   string  `json:"iso_name" gorm:"column:iso_name"`
}
