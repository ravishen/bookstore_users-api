package users

import (
	"github.com/ravishen/bookstore_users-api/datasources/mysql/users_db"
	"github.com/ravishen/bookstore_users-api/utils/date_utils"
	"github.com/ravishen/bookstore_users-api/utils/errors"
	"github.com/ravishen/bookstore_users-api/utils/mysql_utils"
)

var (
	userDB = make(map[int64]*User)
)

const (
	indexUniqueEmail = "email_unique"
	errorNoRows      = "no rows in result set"
	queryInsertUser  = "INSERT INTO users(first_name,last_name,email,date_created) VALUES(?, ?, ?, ?);"
	queryGetUser     = "select id,first_name,last_name,email,date_created from users where id =?;"
)

func (user *User) Get() *errors.RestErr {

	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		return mysql_utils.ParseErr(err)
	}
	// result := userDB[user.Id]
	// if result == nil {
	// 	return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	// }
	// user.Id = result.Id
	// user.FirstName = result.FirstName
	// user.LastName = result.LastName
	// user.Email = result.Email
	// user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysql_utils.ParseErr(saveErr)
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {

		return mysql_utils.ParseErr(saveErr)
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
