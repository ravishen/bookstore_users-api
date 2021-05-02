package services

import (
	"github.com/ravishen/bookstore_users-api/domain/users"
	"github.com/ravishen/bookstore_users-api/utils/crypto_utils"
	"github.com/ravishen/bookstore_users-api/utils/date_utils"
	"github.com/ravishen/bookstore_users-api/utils/errors"
)

var (
	UserService userService = userService{}
)

type userService struct {
}

type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	GetUser(int64) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	Search(string) (users.Users, *errors.RestErr)
}

func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Status = users.StatusActive
	user.Passsword = crypto_utils.GetMd5(user.Passsword)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) DeleteUser(userId int64) *errors.RestErr {
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

func (s *userService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	user := &users.User{
		Id: userId,
	}
	err := user.Get()
	if err != nil {
		return nil, errors.NewBadRequestError(err.Message)
	}
	return user, nil

}

func (s *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current := &users.User{Id: user.Id}
	if err := current.Get(); err != nil {
		return nil, err
	}
	//	current, err := GetUser(user.Id)
	// if err != nil {
	// 	return nil, err
	// }
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

func (s *userService) Search(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
