package usecase

import (
	"errors"
	"finance-tracker/internal/domain"
	"finance-tracker/internal/repository"
	"finance-tracker/pkg/hash"
)

// UserUsecase defines the interface for user business logic
type UserUsecase interface {
	Register(name, email, password string) error
	Login(email, password string) (*domain.User, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

// NewUserUsecase creates a new user usecase
func NewUserUsecase(r repository.UserRepository) UserUsecase {
	return &userUsecase{r}
}

// Register registers a new user
func (u *userUsecase) Register(name, email, password string) error {
	existingUser, _ := u.repo.FindByEmail(email)
	if existingUser != nil && existingUser.ID != 0 {
		return errors.New("email already registered")
	}

	hashed, err := hash.HashPassword(password)
	if err != nil {
		return err
	}

	newUser := &domain.User{
		Name:     name,
		Email:    email,
		Password: hashed,
	}

	return u.repo.Create(newUser)
}

// Login authenticates a user
func (u *userUsecase) Login(email, password string) (*domain.User, error) {
	user, err := u.repo.FindByEmail(email)
	if err != nil || user.ID == 0 {
		return nil, errors.New("email not found")
	}

	if !hash.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("incorrect password")
	}

	return user, nil
}
