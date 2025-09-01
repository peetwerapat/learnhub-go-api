package repository

import "github.com/peetwerapat/learnhub-go-api/internal/domain"

type ContentRepository interface {
	CreateContent(content *domain.Content) error
	GetContents() ([]domain.Content, error)
	GetContentById(id string) (*domain.Content, error)
}
