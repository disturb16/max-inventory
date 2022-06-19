package service

import (
	context "context"
	"database/sql"
	"encoding/base64"
	"os"
	"testing"

	"github.com/disturb/max-inventory/encryption"
	"github.com/disturb/max-inventory/internal/entity"
	"github.com/disturb/max-inventory/internal/repository"
	mock "github.com/stretchr/testify/mock"
)

var s Service
var repo *repository.MockRepository

func TestMain(m *testing.M) {
	repo = &repository.MockRepository{}
	repo.On("SaveUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	s = New(repo)

	code := m.Run()
	os.Exit(code)
}

func TestRegisterUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		Password      string
		ExpectedError error
	}{
		{
			Name:          "Register valid user",
			Email:         "test@test.com",
			Password:      "testpassword",
			ExpectedError: nil,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			repo.Mock.Test(t)

			err := s.RegisterUser(ctx, tc.Email, tc.Name, tc.Password)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestLoginUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		Password      string
		ExpectedError error
	}{
		{
			Name:          "Invalid Password Credential",
			Email:         "test@test.com",
			Password:      "invalidpassword",
			ExpectedError: ErrInvalidCredentials,
		},
		{
			Name:          "Invalid Email Credential",
			Email:         "invalidemail",
			Password:      "invalidpassword",
			ExpectedError: ErrInvalidCredentials,
		},
		{
			Name:          "Valid Credentials",
			Email:         "test@test.com",
			Password:      "validpassword",
			ExpectedError: nil,
		},
	}

	ctx := context.Background()
	pass, _ := encryption.Encrypt([]byte("validpassword"))
	passwordBased64 := base64.RawStdEncoding.EncodeToString(pass)
	repo.On("GetUserByEmail", mock.Anything, "test@test.com").Return(entity.User{Password: passwordBased64}, nil)
	repo.On("GetUserByEmail", mock.Anything, "invalidemail").Return(entity.User{}, sql.ErrNoRows)

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			repo.Mock.Test(t)

			_, err := s.LoginUser(ctx, tc.Email, tc.Password)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}
