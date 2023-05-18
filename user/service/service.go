package service

import (
	"fmt"
	"mime/multipart"
	awsUpload "rojgaarkaro-backend/aws/upload"
	userModel "rojgaarkaro-backend/user/model"
	userRepo "rojgaarkaro-backend/user/repository"

	"gorm.io/gorm"
)

type Service struct {
	Db   *gorm.DB
	User *userModel.User
}

func NewService(db *gorm.DB, user *userModel.User) *Service {
	return &Service{
		Db:   db,
		User: user,
	}
}

type ServiceMethod interface {
	GetAllUsers() (*[]userModel.User, error)
	GetUser() (*userModel.User, error)
	CreateUser() error
	DeleteUser() error
	UpdateUser() error
	GetUserByEmail() (*userModel.User, error)
	UploadFile(*multipart.FileHeader) error
}

func (service *Service) UploadFile(fileheader *multipart.FileHeader) error {
	location, err := awsUpload.Upload(fileheader, service.User.ID)
	if err != nil {
		return err
	}

	// first we will get the user
	service.User.Image = location
	service.UpdateUser()
	repository := userRepo.NewRepository(service.Db, service.User)
	err = repository.UpdateUser()
	if err != nil {
		return err
	}
	return nil
}

func (service *Service) GetAllUsers() (*[]userModel.User, error) {
	// var result *[]userModel.User

	repository := userRepo.NewRepository(service.Db, service.User)
	result, err := repository.GetAllUsers()
	return result, err
}

func (service *Service) GetUser() (*userModel.User, error) {
	repository := userRepo.NewRepository(service.Db, service.User)
	result, err := repository.GetUser()

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return result, nil
}

func (service *Service) GetUserByEmail() (*userModel.User, error) {
	repository := userRepo.NewRepository(service.Db, service.User)
	result, err := repository.GetUserByEmail()
	fmt.Print("gmail result : ", result)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return result, nil
}

func (service *Service) CreateUser() error {
	repository := userRepo.NewRepository(service.Db, service.User)
	// first we will checck whether gmail exist or not
	_, err := service.GetUserByEmail()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = repository.CreateUser()
			return err
		}
		return err
	}
	return fmt.Errorf("email already exist")
}

func (service *Service) DeleteUser() error {
	repository := userRepo.NewRepository(service.Db, service.User)
	err := repository.DeleteUser()
	return err
}

func (service *Service) UpdateUser() error {
	user, err := service.GetUser()
	if err != nil {
		return err
	}
	updatedUser := service.User

	if updatedUser.FirstName != "" {
		user.FirstName = updatedUser.FirstName
	}

	if updatedUser.LastName != "" {
		user.LastName = updatedUser.LastName
	}

	if updatedUser.Contact != "" {
		user.Contact = updatedUser.Contact
	}

	if updatedUser.Email != "" {
		user.Email = updatedUser.Email
	}

	if updatedUser.Password != "" {
		user.Password = updatedUser.Password
	}
	if updatedUser.Image != "" {
		user.Image = updatedUser.Image
	}
	user.UpdatedAt = updatedUser.UpdatedAt

	repository := userRepo.NewRepository(service.Db, service.User)
	err = repository.UpdateUser()
	if err != nil {
		return err
	}
	return nil

	// we will update all the fields which are not null

	// types := values.Type()
	// for i := 0; i < values.NumField(); i++ {
	// 	fmt.Println(types.Field(i))
	// }

	// for i := 0; i < values.NumField(); i++ {
	// 	fmt.Println(types.Field(i).Name)
	// }

	// for i := 0; i < values.NumField(); i++ {
	// 	fmt.Println(values.Field(i).Type())
	// }

	// for i := 0; i < values.NumField(); i++ {
	// 	fmt.Println(values.Field(i))
	// }

	// valuesForUpdation := reflect.ValueOf(updatedUser)
	// types := valuesForUpdation.Type()
	// valueUser := reflect.ValueOf(user)
	// for i := 0; i < valuesForUpdation.NumField(); i++ {
	// 	// if values.Field(i).Type().String() == "int" {

	// 	// }
	// 	switch valuesForUpdation.Field(i).Type().String() {
	// 	case "int":
	// 		// we will edit this later
	// 	case "float":
	// 		// we will edit this later
	// 	case "bool":
	// 		// we will edit this later
	// 	case "string":
	// 		fieldName := types.Field(i).Name
	// 		fmt.Println(fieldName, valuesForUpdation.Field(i))
	// 		valueUser = valuesForUpdation.Field(i)
	// 	}
	// }

}
