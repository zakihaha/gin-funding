package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zakihaha/gin-funding/auth"
	"github.com/zakihaha/gin-funding/campaign"
	"github.com/zakihaha/gin-funding/handler"
	"github.com/zakihaha/gin-funding/helper"
	"github.com/zakihaha/gin-funding/middleware"
	"github.com/zakihaha/gin-funding/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	helper.LoadEnvVariables()
}

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/gin_funding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)

	userService := user.NewService(userRepository)
	authService := auth.NewService()
	campaignService := campaign.NewService(campaignRepository)

	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	router := gin.Default()
	router.Static("/images", "./public/images")
	api := router.Group("/api/v1")

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", middleware.AuthMiddleware(authService, userService), userHandler.UploadAvatar)

	api.GET("/campaigns", campaignHandler.GetCampaigns)

	router.Run(os.Getenv("PORT"))
}
