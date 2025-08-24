package usecase

import (
	"github.com/peetwerapat/learnhub-go-api/internal/domain"
	"github.com/peetwerapat/learnhub-go-api/internal/interface/repository"
)

type ContentUsecase struct {
	contentRepo repository.ContentRepository
}

func NewContentUsecase(repo repository.ContentRepository) *ContentUsecase {
	return &ContentUsecase{contentRepo: repo}
}

func (uc *ContentUsecase) CreateContent(content *domain.Content) error {
	return uc.contentRepo.CreateContent(content)
}
