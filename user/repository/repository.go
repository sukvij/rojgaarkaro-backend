package repository

import (
	"fmt"
	userModel "rojgaarkaro-backend/user/model"

	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) (*[]userModel.User, error) {
	var result *[]userModel.User
	err := db.Find(&result).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, nil
}

func GetUser(db *gorm.DB, userId string) (*userModel.User, error) {
	var result *userModel.User
	err := db.Where("id = ?", userId).First(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(db *gorm.DB, user userModel.User) error {
	// first checck whether given email exist or not
	var result *userModel.User
	err := db.Where("email = ?", user.Email).First(&result).Error

	if err != nil && err == gorm.ErrRecordNotFound {
		fmt.Println("it mean record not found just insert record")
		err = db.Create(&user).Error
		return err
	}
	return fmt.Errorf("email aready exist")
}

func DeleteUser(db *gorm.DB, userId string) error {
	var result *userModel.User
	err := db.Where("id = ?", userId).Delete(&result).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(db *gorm.DB, user userModel.User) error {
	err := db.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}
