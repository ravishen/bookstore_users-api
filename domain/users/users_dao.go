package users

import (
	"fmt"
	"strings"

	"github.com/ravishen/bookstore_users-api/datasources/mysql/users_db"
	"github.com/ravishen/bookstore_users-api/utils/date_utils"
	"github.com/ravishen/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

const (
	indexUniqueEmail = "email_unique"
	queryInsertUser  = "INSERT INTO users(first_name,last_name,email,date_created) VALUES(?, ?, ?, ?);"
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
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

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error creating user %s", err.Error()))
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("Unable to get last insert id")
	}
	user.Id = userId
	// current := userDB[user.Id]
	// if current != nil {
	// 	if current.Email == user.Email {
	// 		return errors.NewBadRequestError(fmt.Sprintf("%s email already exists", user.Email))

	// 	}
	// 	return errors.NewBadRequestError(fmt.Sprintf("%d userid already exists", user.Id))
	// } else {
	// 	userDB[user.Id] = &User{
	// 		Id:          user.Id,
	// 		FirstName:   user.FirstName,
	// 		LastName:    user.LastName,
	// 		Email:       user.Email,
	// 		DateCreated: date_utils.GetNowString(),
	// 	}
	// }
	return nil
}
