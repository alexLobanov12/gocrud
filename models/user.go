package models

import (
	"errors"
)

type Message struct {
	Message string `json:"message"`
}

type Login struct {
	Login string `json:"login"`
	Password string `json:"password"`
}

type LoginError struct {
	Status int      `json:"status"`
	Errors []string `json:"errors"`
}

func (l *Login) ValidateLogin() []error {
	var errorsArr []error
	if l.Login == "" {
		errorsArr = append(errorsArr, errors.New("login empty"))
	}

	if l.Password == "" {
		errorsArr = append(errorsArr, errors.New("password empty"))
	}

	return errorsArr
}

/*
const (
	ErrorMaxLimitChar
)
*/