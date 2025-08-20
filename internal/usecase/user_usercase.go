package usecase

import (
	"errors"

	"github.com/peetwerapat/learnhub-go-api/internal/domain"
	"github.com/peetwerapat/learnhub-go-api/internal/interface/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: repo}
}

func (uc *UserUsecase) GetUserById(id uint) (*domain.User, error) {
	return uc.userRepo.GetUserById(id)
}

func (uc *UserUsecase) CreateUser(user *domain.User) error {

	if len(user.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashed)

	return uc.userRepo.CreateUser(user)
}
