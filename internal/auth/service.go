package auth

import (
	"errors"
	"go/adv-dev/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
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
