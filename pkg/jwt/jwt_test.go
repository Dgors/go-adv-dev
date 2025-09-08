package jwt_test

import (
	"go/adv-dev/pkg/jwt"
	"testing"
)

func TestJWTCreate(t *testing.T) {
	const email = "test@mail.ru"
	jwtService := jwt.NewJWT("736380ed4ca3d5036dcf44dc9dc97e1bc90c7e05310616bc0bfbede4a2f222dc")
	token, err := jwtService.Create(jwt.JWTData{
		Email: email,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatal("token is not valid")
	}
	if data.Email != email {
		t.Fatalf("email %s not equal to %s", data.Email, email)
	}
	t.Log(data.Email)
}
