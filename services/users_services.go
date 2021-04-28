package services

import (
	"github.com/ravishen/bookstore_users-api/domain/users"
	"github.com/ravishen/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{
		Id: userId,
	}
	return user.Delete()
	// err := user.Get()
	// if err != nil {
	// 	return errors.NewBadRequestError(err.Message)
	// }
	// return nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	user := &users.User{
		Id: userId,
	}
	err := user.Get()
	if err != nil {
		return nil, errors.NewBadRequestError(err.Message)
	}
	return user, nil

}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
		current.DateCreated = user.DateCreated
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
		current.DateCreated = user.DateCreated
	}
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}
