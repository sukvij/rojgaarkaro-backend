package model

import (
	"rojgaarkaro-backend/baseThing/model"
)

type User struct {
	model.BaseModel
	FirstName string `json:"first_name" `
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Age       int    `json:"age"`
	Contact   string
	Email     string `validate:"required"`
	Password  string
	IsMember  bool
	Priority  int
	Verified  bool
}
