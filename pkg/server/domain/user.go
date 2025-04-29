package domain

import (
	"errors"
	"regexp"
)

type User struct {
	Email    string
	Password string
}

func NewUser(email, password string) (*User, error) {
	if !isValidEmail(email) {
		return nil, errors.New("invalid email format")
	}
	if len(password) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}
	return &User{
		Email:    email,
		Password: password,
	}, nil
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetPassword() string {
	return u.Password
}
