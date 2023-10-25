package util

import "database/sql"

func IsUsernameExist(db *sql.DB, username string) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT exists (SELECT 1 FROM users WHERE username=$1)", username).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
