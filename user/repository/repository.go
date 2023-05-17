package repository

import (
	userModel "rojgaarkaro-backend/user/model"

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
	// var result *[]userModel.User
	result := &[]userModel.User{}
	err := repository.Db.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repository *Repository) GetUser() (*userModel.User, error) {
	result := &userModel.User{}
	userId := repository.User.ID
	err := repository.Db.Where("id = ?", userId).First(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repository *Repository) GetUserByEmail() (*userModel.User, error) {
	result := &userModel.User{}
	userEmail := repository.User.Email
	err := repository.Db.Where("email = ?", userEmail).First(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repository *Repository) CreateUser() error {
	err := repository.Db.Create(&repository.User).Error
	return err
}

func (repository *Repository) DeleteUser() error {
	result := &userModel.User{}
	err := repository.Db.Where("id = ?", repository.User.ID).Delete(&result).Error
	return err
}

func (repository *Repository) UpdateUser() error {
	err := repository.Db.Save(&repository.User).Error
	return err
}
