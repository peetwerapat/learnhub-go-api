package db

import (
	"github.com/peetwerapat/learnhub-go-api/internal/domain"
	"github.com/peetwerapat/learnhub-go-api/internal/interface/repository"
	"gorm.io/gorm"
)

type GormContentRepository struct {
	DB *gorm.DB
}

func NewGormContentRepository(db *gorm.DB) repository.ContentRepository {
	return &GormContentRepository{DB: db}
}

func (r *GormContentRepository) CreateContent(content *domain.Content) error {
	return r.DB.Create(content).Error
}

func (r *GormContentRepository) GetContents() ([]domain.Content, error) {
	var contents []domain.Content
	err := r.DB.
		Where("deleted_at IS NULL").
		Preload("User").
		Find(&contents).Error

	if err != nil {
		return nil, err
	}

	return contents, nil
}

func (r *GormContentRepository) GetContentById(id uint) (*domain.Content, error) {
	var c domain.Content
	if err := r.DB.First(&c, id).Error; err != nil {
		return nil, err
	}

	return &c, nil
}
