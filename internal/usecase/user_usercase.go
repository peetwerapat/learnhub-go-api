package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/peetwerapat/learnhub-go-api/internal/domain"
	"github.com/peetwerapat/learnhub-go-api/internal/interface/repository"
	"github.com/peetwerapat/learnhub-go-api/pkg/myJwt"
	"github.com/peetwerapat/learnhub-go-api/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrEmailAlreadyExist  = errors.New("email already exists")
	ErrInvalidEmailFormat = errors.New("invalid email format")
	ErrInvalidEmail       = errors.New("email not found")
	ErrInvalidPassword    = errors.New("incorrect password")
	ErrTokenCreation      = errors.New("failed to create token")
)

type UserUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: repo}
}

func (uc *UserUsecase) CreateUser(user *domain.User) error {
	existingUser, err := uc.userRepo.GetUserByEmail(user.Email)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to check existing user: %w", err)
	}

	if existingUser != nil {
		return ErrEmailAlreadyExist
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashed)

	return uc.userRepo.CreateUser(user)
}

func (uc *UserUsecase) Login(email, password string) (string, error) {
	if !utils.IsValidEmail(email) {
		return "", ErrInvalidEmailFormat
	}

	user, err := uc.userRepo.GetUserByEmail(email)
	if err != nil || user == nil {
		return "", ErrInvalidEmail
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", ErrInvalidPassword
	}

	token, err := myJwt.CreateToken(email, 24*time.Hour)
	if err != nil {
		return "", ErrTokenCreation
	}

	return token, nil
}
