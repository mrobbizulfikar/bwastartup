package main

import (
	"bwastartup/auth"
	"bwastartup/handler"
	"bwastartup/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.sgDLEaGekUN2lbzpRnnOso7wiFCIbbI-jDcWS7ggUTU")
	if err != nil {
		fmt.Println("Error")
		fmt.Println("Error")
		fmt.Println("Error")
	}

	if token.Valid {
		fmt.Println("Valid")
		fmt.Println("Valid")
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
		fmt.Println("Invalid")
		fmt.Println("Invalid")
	}

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email/checks", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()
}
