package util

import (
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
