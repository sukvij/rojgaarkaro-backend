package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName  string
	LastName   string
	Gender     string
	Age        int
	Contact    string
	Email      string
	Password   string
	Verified   bool
	CountryId  int
	StateId    int
	DistrictId int
	TehsilId   int
	VillageId  int
}
