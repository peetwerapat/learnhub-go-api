package repository

import (
	"github.com/peetwerapat/learnhub-go-api/internal/domain"
)

type UserRepository interface {
	GetUserById(id uint) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	CreateUser(user *domain.User) error
}
