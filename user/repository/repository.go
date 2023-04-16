package repository

import (
	userModel "fitcare-backend/user/model"
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	Db   *gorm.DB
	User *userModel.User
}

func NewRepository(Db *gorm.DB, User *userModel.User) *Repository {
	return &Repository{Db: Db, User: User}
}

type RepositoryMethod interface {
	GetAllUsers() (*[]userModel.User, error)
	GetUser() (*userModel.User, error)
	GetUserByEmail() (*userModel.User, error)
	CreateUser() error
	DeleteUser() error
	UpdateUser() error
}

func (repository *Repository) GetAllUsers() (*[]userModel.User, error) {
	var result *[]userModel.User
	err := repository.Db.Find(&result).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, nil
}

func (repository *Repository) GetUser() (*userModel.User, error) {
	var result *userModel.User
	userId := repository.User.ID
	err := repository.Db.Where("id = ?", userId).First(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repository *Repository) GetUserByEmail() (*userModel.User, error) {
	var result *userModel.User
	userEmail := repository.User.Email
	err := repository.Db.Where("email = ?", userEmail).First(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repository *Repository) CreateUser() error {
	// first checck whether given email exist or not
	var result *userModel.User
	err := repository.Db.Where("email = ?", repository.User.Email).First(&result).Error

	if err != nil && err == gorm.ErrRecordNotFound {
		fmt.Println("it mean record not found just insert record")
		err = repository.Db.Create(&repository.User).Error
		return err
	}
	return fmt.Errorf("email aready exist")
}

func (repository *Repository) DeleteUser() error {
	var result *userModel.User
	err := repository.Db.Where("id = ?", repository.User.ID).Delete(&result).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository *Repository) UpdateUser() error {
	err := repository.Db.Save(&repository.User).Error
	if err != nil {
		return err
	}
	return nil
}
