package users

import (
	"strings"

	"github.com/ravishen/bookstore_users-api/utils/errors"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Passsword   string `json:"password"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	user.Passsword = strings.TrimSpace(user.Passsword)
	if user.Passsword == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
