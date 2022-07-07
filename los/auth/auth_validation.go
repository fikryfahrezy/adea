package auth

import (
	"errors"
	"unicode/utf8"
)

var (
	ErrUsernameRequired = errors.New("password required")
	ErrPasswordRequired = errors.New("password required")
)

func validateRegister(in RegisterIn) error {
	if utf8.RuneCountInString(in.Username) == 0 {
		return ErrUsernameRequired
	}
	if utf8.RuneCountInString(in.Password) == 0 {
		return ErrPasswordRequired
	}
	return nil
}

func validateLogin(in LoginIn) error {
	if utf8.RuneCountInString(in.Username) == 0 {
		return ErrUsernameRequired
	}
	if utf8.RuneCountInString(in.Password) == 0 {
		return ErrPasswordRequired
	}
	return nil
}
