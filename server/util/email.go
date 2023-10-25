package util

import (
	"database/sql"
	"errors"
	"regexp"
)

func IsEmailValid(email string) (bool, error) {
	const emailRegexPattern = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	matcher := regexp.MustCompile(emailRegexPattern)
	if matcher.MatchString(email) {
		return true, nil
	} else {
		return false, errors.New("email address is not valid")
	}
}

func IsEmailExist(db *sql.DB, email string) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT exists (SELECT 1 FROM users WHERE email=$1)", email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
