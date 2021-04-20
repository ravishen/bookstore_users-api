package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	users "github.com/ravishen/bookstore_users-api/domain/users"
	"github.com/ravishen/bookstore_users-api/services"
	"github.com/ravishen/bookstore_users-api/utils/errors"
)

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Status, err)
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		err := errors.NewBadRequestError("user not found")
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user users.User

	// fmt.Println(user)
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	// handle error
	// }

	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	//handle err

	// }

	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("invalid json body")
		c.JSON(restError.Status, restError)

		return
	}
	result, err_svc := services.CreateUser(user)
	if err_svc != nil {
		c.JSON(err_svc.Status, err_svc)
	}
	c.JSON(http.StatusCreated, result)

}
func SearchUser(c *gin.Context) {

}
