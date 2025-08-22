package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/peetwerapat/learnhub-go-api/config"
	_ "github.com/peetwerapat/learnhub-go-api/docs"
	"github.com/peetwerapat/learnhub-go-api/internal/infrastructure/db"
	"github.com/peetwerapat/learnhub-go-api/internal/interface/controller"
	"github.com/peetwerapat/learnhub-go-api/internal/usecase"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// CORS config
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PACTH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	// Init repositories
	userRepo := db.NewGormUserRepository(config.DB)

	// Init usecases
	userUC := usecase.NewUserUsecase(userRepo)

	// Init controllers
	controller.NewAuthController(r, userUC)

	return r
}
