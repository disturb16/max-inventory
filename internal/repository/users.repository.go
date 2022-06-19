package repository

import (
	context "context"

	"github.com/disturb/max-inventory/internal/entity"
)

const (
	qryInsertUser = `
		insert into USERS (email, name, password)
		values (?, ?, ?);
	`

	qryGetUserByEmail = `
		select id, name, email, password from USERS where email = ?;
	`
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	_, err := r.db.ExecContext(ctx, qryInsertUser, email, name, password)
	return err
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	user := entity.User{}
	err := r.db.GetContext(ctx, &user, qryGetUserByEmail, email)
	return user, err
}
