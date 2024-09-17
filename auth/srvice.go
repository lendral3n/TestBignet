package auth

import (
	"tesBignet/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(email, password string) (models.User, error)
	Login(email, password string) (models.User, error)
	GetAllUsers() ([]models.User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) Register(email, password string) (models.User, error) {
	existingUser, _ := s.repo.FindByEmail(email)
	if existingUser.ID != 0 {
		return models.User{}, errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Email:    email,
		Password: string(hashedPassword),
	}
	return s.repo.Create(user)
}

func (s *service) Login(email, password string) (models.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return models.User{}, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.User{}, errors.New("invalid email or password")
	}

	return user, nil
}

func (s *service) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}
