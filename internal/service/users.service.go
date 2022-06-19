package service

import (
	"context"
	"database/sql"
	"encoding/base64"
	"errors"

	"github.com/disturb/max-inventory/encryption"
	"github.com/disturb/max-inventory/internal/models"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

func (s *service) RegisterUser(ctx context.Context, email, name, password string) error {
	pass, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}

	return s.repo.SaveUser(ctx, email, name, base64.RawStdEncoding.EncodeToString(pass))
}

func (s *service) LoginUser(ctx context.Context, email, password string) (*models.User, error) {

	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrInvalidCredentials
		}

		return nil, err
	}

	pass, err := encryption.Decode(u.Password)
	if err != nil {
		return nil, err
	}

	if string(pass) != password {
		return nil, ErrInvalidCredentials
	}

	//TODO: return JWT token

	return &models.User{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
	}, nil
}
