package users

import (
	"fmt"

	"github.com/ravishen/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user User) Save() *errors.RestErr {
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("%s email already exists", user.Email))

		}
		return errors.NewBadRequestError(fmt.Sprintf("%d userid already exists", user.Id))
	} else {
		userDB[user.Id] = &User{
			Id:          user.Id,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Email:       user.Email,
			DateCreated: user.DateCreated,
		}
	}
	return nil
}
