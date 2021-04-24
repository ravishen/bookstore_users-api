package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/ravishen/bookstore_users-api/utils/errors"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseErr(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record found ")
		}
		return errors.NewInternalServerError("error parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("duplicated key")
	}
	return errors.NewInternalServerError("error processing request")
}
