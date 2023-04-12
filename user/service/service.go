package service

import (
	userModel "rojgaarkaro-backend/user/model"
	userRepo "rojgaarkaro-backend/user/repository"

	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) (*[]userModel.User, error) {
	// var result *[]userModel.User
	result, err := userRepo.GetAllUsers(db)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetUser(db *gorm.DB, userId string) (*userModel.User, error) {
	result, err := userRepo.GetUser(db, userId)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(db *gorm.DB, user userModel.User) error {
	err := userRepo.CreateUser(db, user)
	return err
}

func DeleteUser(db *gorm.DB, userId string) error {
	err := userRepo.DeleteUser(db, userId)

	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(db *gorm.DB, userId string, updatedUser userModel.User) error {
	user, err := GetUser(db, userId)
	if err != nil {
		return err
	}

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
	user.UpdatedAt = updatedUser.UpdatedAt

	err = userRepo.UpdateUser(db, *user)
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
