package auth

import (
	"errors"
	"go/adv-dev/internal/user"
	"go/adv-dev/pkg/di"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository di.IUserRepository
}

func NewAuthService(userRepository di.IUserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (service *AuthService) Register(email, password, name string) (string, error) {
	existedUser, _ := service.UserRepository.GetByEmail(email)
	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	newUser := user.NewUser(email, string(hashedPassword), name)
	_, err = service.UserRepository.Create(newUser)
	if err != nil {
		return "", err
	}
	return newUser.Email, nil
}

func (service *AuthService) Login(email, password string) (string, error) {
	existedUser, _ := service.UserRepository.GetByEmail(email)
	if existedUser == nil {
		return "", errors.New(ErrEmailOrPassword)
	}
	err := bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(password))
	if err != nil {
		return "", errors.New(ErrEmailOrPassword)
	}
	return existedUser.Email, nil
}
