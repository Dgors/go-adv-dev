package auth_test

import (
	"bytes"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"go/adv-dev/configs"
	"go/adv-dev/internal/auth"
	"go/adv-dev/internal/user"
	"go/adv-dev/pkg/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func bootstrap() (*auth.AuthHandler, sqlmock.Sqlmock, error) {
	databases, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: databases,
	}))
	if err != nil {
		return nil, nil, err
	}
	userRepo := user.NewUserRepository(&db.Db{
		DB: gormDb,
	})
	handler := auth.AuthHandler{
		Config: &configs.Config{
			Auth: configs.AuthConfig{
				Secret: "secret",
			},
		},
		AuthService: auth.NewAuthService(userRepo),
	}
	return &handler, mock, nil
}

func TestLoginSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	rows := sqlmock.NewRows([]string{"email", "password"}).
		AddRow("test@test.com", "$2a$10$IHRW1kLDSR4yDZm0SyHb9.mfKvtMDGUDHTvKOEoj3DfnxRC9uvyL.")
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	if err != nil {
		t.Fatal(err)
	}
	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "test@test.com",
		Password: "test1234",
	})
	reader := bytes.NewReader(data)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/login", reader)
	handler.Login()(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("login: got %d, want %d", w.Result().StatusCode, http.StatusOK)
	}
}
