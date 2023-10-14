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
	prefix := app.Group("api/v1")
	prefix.POST("/user", handlerUser.RegisterUser)
	prefix.POST("/login", handlerUser.Login)

	app.Run()
}
