package auth_test

import (
	"go/adv-dev/internal/auth"
	"go/adv-dev/internal/user"
	"testing"
)

type MockUserRepository struct{}

func (repo *MockUserRepository) Create(u *user.User) (*user.User, error) {
	return &user.User{
		Email: "test@mail.ru",
	}, nil
}

func (repo *MockUserRepository) GetByEmail(email string) (*user.User, error) {
	return nil, nil
}

func TestRegisterSuccess(t *testing.T) {
	const initialEmail = "test@mail.ru"
	authService := auth.NewAuthService(&MockUserRepository{})
	email, err := authService.Register(initialEmail, "test1234", "test")
	if err != nil {
		t.Fatal(err)
	}
	if email != initialEmail {
		t.Fatalf("email %s does not match %s", email, initialEmail)
	}
}
