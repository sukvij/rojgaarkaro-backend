package controller

import (
	authentication "rojgaarkaro-backend/authentication"
	userModel "rojgaarkaro-backend/user/model"
	userService "rojgaarkaro-backend/user/service"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

var db *gorm.DB

func UserApis(app *iris.Application, DB *gorm.DB) {
	db = DB
	AllUserApis := app.Party("/user")
	{
		// AllUserApis.Post("/signin", authentication.GenerateToken())
		AllUserApis.Get("/", authentication.VerifyMiddleware, getAllUsers)
		AllUserApis.Get("/{userId}", authentication.VerifyMiddleware, getUser)
		AllUserApis.Post("/", createUser)
		AllUserApis.Delete("/{userId}", authentication.VerifyMiddleware, deleteUser)
		AllUserApis.Put("/{userId}", authentication.VerifyMiddleware, updateUser)
		// AllUserApis.Get("/logout", authentication.VerifyMiddleware, authentication.Logout)
	}
}

func getAllUsers(ctx iris.Context) {
	result, err := userService.GetAllUsers(db)
	if err != nil {
		ctx.JSON(err)
		return
	}
	ctx.JSON(result)
}

func getUser(ctx iris.Context) {
	userId := ctx.Params().Get("userId")
	result, err := userService.GetUser(db, userId)
	if err != nil {
		ctx.JSON(err.Error())

		return
	}
	ctx.JSON(result)
}

func createUser(ctx iris.Context) {
	var user userModel.User
	ctx.ReadJSON(&user)
	err := userService.CreateUser(db, user)
	if err == nil {
		ctx.JSON("user created successfully")
	} else {
		ctx.JSON(err.Error())
	}
}

func deleteUser(ctx iris.Context) {
	userId := ctx.Params().Get("userId")
	err := userService.DeleteUser(db, userId)
	if err == nil {
		ctx.JSON("user deleted successfully")
	} else {
		ctx.JSON(err.Error())
	}
}

func updateUser(ctx iris.Context) {
	userId := ctx.Params().Get("userId")
	var updatedUser userModel.User
	ctx.ReadJSON(&updatedUser)
	err := userService.UpdateUser(db, userId, updatedUser)
	if err == nil {
		ctx.JSON("user updated successfully")
	} else {
		ctx.JSON(err.Error())
	}
}
