package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Gender    string
	Age       int
	Contact   string
	Email     string
	Password  string
	IsMember  bool
	Priority  int
	Verified  bool
}
