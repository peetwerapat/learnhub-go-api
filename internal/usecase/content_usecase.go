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

func (uc *ContentUsecase) GetContents() ([]domain.Content, error) {
	return uc.contentRepo.GetContents()
}

func (uc *ContentUsecase) GetContentById(id string) (*domain.Content, error) {
	return uc.contentRepo.GetContentById(id)
}
