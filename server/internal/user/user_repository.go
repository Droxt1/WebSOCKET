package user

import (
	"context"
	"database/sql"
	"errors"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return &User{}, err
	}

	var lastInsertID int
	query := "INSERT INTO users(username, password, email) VALUES ($1, $2, $3) RETURNING id"

	// Check for unique constraint violations before attempting to insert
	checkQuery := "SELECT id FROM users WHERE username = $1 OR email = $2"
	var existingID int
	err = tx.QueryRowContext(ctx, checkQuery, user.Username, user.Email).Scan(&existingID)
	if err == nil {
		err := tx.Rollback()
		if err != nil {
			return nil, err
		}
		return &User{}, errors.New("User with this username or email already exists")
	} else if err != sql.ErrNoRows {
		err := tx.Rollback()
		if err != nil {
			return nil, err
		}
		return &User{}, err
	}

	err = tx.QueryRowContext(ctx, query, user.Username, user.Password, user.Email).Scan(&lastInsertID)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, err
		}
		return &User{}, err
	}

	user.ID = int64(lastInsertID)

	if err = tx.Commit(); err != nil {
		return &User{}, err
	}

	return user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	u := User{}
	query := "SELECT id, email, username, password FROM users WHERE email = $1"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Email, &u.Username, &u.Password)
	if err != nil {
		return &User{}, err
	}

	return &u, nil
}
