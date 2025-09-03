package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/peetwerapat/learnhub-go-api/docs"
	"github.com/peetwerapat/learnhub-go-api/internal/infrastructure/di"
	"github.com/peetwerapat/learnhub-go-api/internal/interface/controller"
	"github.com/peetwerapat/learnhub-go-api/pkg/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(appUc *di.AppUseCase) *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// CORS config
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	// Auth Route
	authController := controller.NewAuthController(appUc.UserUc)
	r.POST("/register", authController.CreateUser)
	r.POST("/login", authController.Login)

	// Content Route
	contentController := controller.NewContentController(appUc.ContentUc)
	r.POST("/contents", middleware.AuthMiddleware(), contentController.CreateContent)
	r.GET("/contents", contentController.GetContents)
	r.GET("/contents/:id", contentController.GetContenById)

	return r
}
