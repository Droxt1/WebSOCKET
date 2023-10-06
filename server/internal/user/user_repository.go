package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}
type repository struct {
	db DBTX
}

// passing DBTX interface to repository struct because sometimes we want to pass a transaction to the repository instead of the whole DB connection
// so later we can inject transaction as dependency instead of the whole DB connection
func NewRepository(db DBTX) Repository {
	return &repository{db}
}

// CreateUser
// 1. create a query
// 2. execute the query
// 3. scan the result to get the last insert id
// 4. return the user with the last insert id
func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var lastInsertID int64
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&lastInsertID)
	if err != nil {
		return &User{}, err
	}
	user.ID = lastInsertID
	return user, nil
}
