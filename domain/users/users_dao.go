package users

import (
	"github.com/ravishen/bookstore_users-api/datasources/mysql/users_db"
	"github.com/ravishen/bookstore_users-api/logger"
	"github.com/ravishen/bookstore_users-api/utils/date_utils"
	"github.com/ravishen/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

const (
	indexUniqueEmail = "email_unique"
	errorNoRows      = "no rows in result set"
	queryInsertUser  = "INSERT INTO users(first_name,last_name,email,date_created,password,status) VALUES(?, ?, ?, ?,?,?);"
	queryGetUser     = "select id,first_name,last_name,email,date_created from users where id =?;"
	queryUpdateUser  = "update users set first_name=?,last_name=?,email=? where id = ?;"
	queryDelete      = "delete from users where id = ?"
	findUserByStatus = "select id,first_name, last_name, email, date_created, status from users where status =? ;"
)

func (user *User) Get() *errors.RestErr {

	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("Error when trying to prepare get user statement", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {

		logger.Error("Error when trying to get user by id", err)
		return errors.NewInternalServerError("Database error")
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
		logger.Error("Error when trying to prepare save user statement", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowDBFormat()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Passsword, user.Status)
	if saveErr != nil {
		logger.Error("Error when trying to save user to database", err)
		return errors.NewInternalServerError("Database error")
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {

		logger.Error("Error when trying to fetch the last insert id", err)
		return errors.NewInternalServerError("Database error")
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

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("Error when trying to prepare update user statement", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("Error when trying to execute update user statement", err)
		return errors.NewInternalServerError("Database error")
	}
	return nil
}
func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDelete)
	if err != nil {
		logger.Error("Error when trying to prepare delete user statement", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id)
	if err != nil {
		logger.Error("Error when trying to execute delete query to database", err)
		return errors.NewInternalServerError("Database error")
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(findUserByStatus)
	if err != nil {
		logger.Error("Error when trying to prepare find by status statement", err)
		return nil, errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("Error when trying to execute find by status statement on database", err)
		return nil, errors.NewInternalServerError("Database error")
	}
	defer rows.Close()
	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("Error when trying to scan database rows", err)
			return nil, errors.NewInternalServerError("Error trying to scan database rows")
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		logger.Error("No active users found", err)
		return nil, errors.NewInternalServerError("No active users found")
	}
	return results, nil
}
