package util

import (
	"errors"
	"regexp"
)

func IsUsernameValid(username string) (bool, error) {
	const usernameRegexPattern = `^[a-zA-Z0-9_-]{3,20}$`

	matcher := regexp.MustCompile(usernameRegexPattern)
	if matcher.MatchString(username) {
		return true, nil
	} else {
		return false, errors.New("username is not valid")
	}
}
