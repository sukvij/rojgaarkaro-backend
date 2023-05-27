package controller

import (
	authentication "rojgaarkaro-backend/authentication"
	errorWithDetails "rojgaarkaro-backend/baseThing"
	userModel "rojgaarkaro-backend/user/model"
	userService "rojgaarkaro-backend/user/service"
	"strconv"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

var db *gorm.DB

func UserApis(app *iris.Application, DB *gorm.DB) {
	db = DB
	AllUserApis := app.Party("/user")
	{
		AllUserApis.Post("/signin/{userEmail}/{userPassword}", checkUserExistance, authentication.GenerateToken, getUserByEmail)
		AllUserApis.Get("/", authentication.VerifyMiddleware, getAllUsers)
		AllUserApis.Get("/{userId}", authentication.VerifyMiddleware, getUser)
		// AllUserApis.Get("/{userEmail}", authentication.VerifyMiddleware, getUserByEmail)
		AllUserApis.Post("/{userEmail}", authentication.VerifyMiddleware, getUserByEmail, resetPassword)
		AllUserApis.Post("/", createUser)
		AllUserApis.Delete("/{userId}", authentication.VerifyMiddleware, deleteUser)
		AllUserApis.Put("/{userId}", authentication.VerifyMiddleware, updateUser)
		// AllUserApis.Post("/upload", uploadFile)
		// AllUserApis.Get("/logout", authentication.VerifyMiddleware, authentication.Logout)
	}
}

func checkUserExistance(ctx iris.Context) {
	userEmail := ctx.Params().Get("userEmail")
	userPassword := ctx.Params().Get("userPassword")
	user := &userModel.User{Email: userEmail, Password: userPassword}
	service := &userService.Service{Db: db, User: user}
	userExist, err := service.GetUserByEmail()
	if err.Detail != "" {
		ctx.StopWithJSON(err.Status, err)
		return
	}

	if userExist.Password != userPassword {
		err1 := errorWithDetails.ErrorWithDetails{Status: 401, Detail: "wrong password"}
		ctx.StopWithJSON(err1.Status, err1)
		return
	}
	ctx.Next()
}

func getAllUsers(ctx iris.Context) {
	service := &userService.Service{Db: db, User: &userModel.User{}}
	result, err := service.GetAllUsers()
	if err.Detail != "" {
		ctx.StopWithJSON(err.Status, err)
		return
	}
	ctx.JSON(result)
}

func getUser(ctx iris.Context) {
	userId := ctx.Params().Get("userId")
	user := &userModel.User{}
	user.ID, _ = strconv.ParseInt(userId, 10, 64)
	service := &userService.Service{Db: db, User: user}
	result, err := service.GetUser()
	if err.Detail != "" {
		ctx.StopWithJSON(err.Status, err)
		return
	}
	ctx.JSON(result)
}

func resetPassword(ctx iris.Context) {
	// enter gmail address to check whether email exist or not.
	userEmail := ctx.Params().Get("userEmail")
	userPassword := ctx.Params().Get("userPassword")
	user := &userModel.User{Email: userEmail, Password: userPassword}
	service := &userService.Service{Db: db, User: user}
	updatedUser, err := service.GetUserByEmail()
	if err.Detail != "" {
		ctx.StopWithJSON(err.Status, err)
		return
	}
	service.User = updatedUser
	errs := service.UpdateUser()
	if errs.Detail != "" {
		ctx.StopWithJSON(errs.Status, errs)
		return
	}
	ctx.JSON("password reset successfully")
}

func getUserByEmail(ctx iris.Context) {
	userEmail := ctx.Params().Get("userEmail")
	user := &userModel.User{}
	user.Email = userEmail
	service := &userService.Service{Db: db, User: user}
	result, err := service.GetUserByEmail()
	if err.Detail != "" {
		ctx.StopWithJSON(err.Status, err)
		return
	}
	ctx.JSON(result)
	ctx.Next()
}

func createUser(ctx iris.Context) {
	var user userModel.User
	ctx.ReadJSON(&user)
	service := &userService.Service{Db: db, User: &user}
	err := service.CreateUser()
	if err.Detail == "" {
		ctx.JSON("user created successfully")
	} else {
		ctx.StopWithJSON(err.Status, err)
	}
}

func deleteUser(ctx iris.Context) {
	userId := ctx.Params().Get("userId")
	user := &userModel.User{}
	user.ID, _ = strconv.ParseInt(userId, 10, 64)
	service := &userService.Service{Db: db, User: user}
	err := service.DeleteUser()
	if err.Detail == "" {
		ctx.JSON("user deleted successfully")
	} else {
		ctx.StopWithJSON(err.Status, err)
	}
}

func updateUser(ctx iris.Context) {
	userId := ctx.Params().Get("userId")
	updatedUser := &userModel.User{}
	ctx.ReadJSON(&updatedUser)
	updatedUser.ID, _ = strconv.ParseInt(userId, 10, 64)
	service := &userService.Service{Db: db, User: updatedUser}
	err := service.UpdateUser()
	if err.Detail == "" {
		ctx.JSON("user updated successfully")
	} else {
		ctx.StopWithJSON(err.Status, err)
	}
}

// func uploadFile(ctx iris.Context) {

// 	userId := ctx.FormValue("id")
// 	file, fileHeader, err := ctx.FormFile("file")
// 	if err != nil {
// 		ctx.StopWithError(iris.StatusBadRequest, err)
// 		return
// 	}
// 	user := &userModel.User{}
// 	user.ID, _ = strconv.ParseInt(userId, 10, 64)
// 	service := &userService.Service{Db: db, User: user}
// 	err = service.UploadFile(fileHeader)
// 	if err != nil {
// 		ctx.JSON(err)
// 	}
// 	// ctx.Writef("File: %s uploaded!", fileHeader.Filename)
// 	defer file.Close()
// }
