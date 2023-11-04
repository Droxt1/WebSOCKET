package db

import "database/sql"

type DB struct {
	db *sql.DB
}

func NewDB() (*DB, error) {
	db, err := sql.Open("postgres", "postgres://root:password@localhost:5433/go-chat?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &DB{
		db}, nil

}

func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) GetDB() *sql.DB {
	return db.db
}
