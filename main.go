package main

import (
	"bwu-startup/handler"
	"bwu-startup/repository"
	"bwu-startup/service"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=postgres password=pass123 dbname=bwu_startup port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	repoUser := repository.NewUserRepository(db)
	svcUser := service.NewUserService(repoUser)
	handlerUser := handler.NewUserHandler(svcUser)

	app := gin.Default()
	UrlPrefix := app.Group("api/v1")
	UrlPrefix.POST("/user", handlerUser.RegisterUser)
	UrlPrefix.POST("/login", handlerUser.Login)
	UrlPrefix.POST("/email_check", handlerUser.CheckAvailableEmail)
	UrlPrefix.POST("/avatar", handlerUser.UploadAvatar)

	app.Run()
}
