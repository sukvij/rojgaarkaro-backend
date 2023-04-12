package main

import (
	"fmt"
	authentication "rojgaarkaro-backend/authentication"
	database "rojgaarkaro-backend/database"
	userController "rojgaarkaro-backend/user/controller"

	"github.com/kataras/iris/v12"
)

func main() {
	db, err := database.Connection()
	if err != nil {
		return
	}

	fmt.Println(db)
	app := iris.New()
	// crs := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"*"},
	// 	AllowedMethods:   []string{"GET", "POST", "DELETE"},
	// 	AllowCredentials: true,
	// })
	// app.Use(crs)
	app.Post("/signin", authentication.GenerateToken)
	app.Get("/logout", authentication.VerifyMiddleware, authentication.Logout)
	userController.UserApis(app, db)
	app.Listen(":8080")
}
