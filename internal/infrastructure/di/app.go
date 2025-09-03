package di

import (
	"github.com/peetwerapat/learnhub-go-api/internal/infrastructure/repository"
	"github.com/peetwerapat/learnhub-go-api/internal/usecase"
	"gorm.io/gorm"
)

type AppUseCase struct {
	UserUc    *usecase.UserUsecase
	ContentUc *usecase.ContentUsecase
}

func InitApp(dbConn *gorm.DB) *AppUseCase {
	userRepo := repository.NewGormUserRepository(dbConn)
	userUC := usecase.NewUserUsecase(userRepo)

	contentRepo := repository.NewGormContentRepository(dbConn)
	contentUC := usecase.NewContentUsecase(contentRepo)

	return &AppUseCase{
		UserUc:    userUC,
		ContentUc: contentUC,
	}
}
