package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	users "github.com/ravishen/bookstore_users-api/domain/users"
	"github.com/ravishen/bookstore_users-api/services"
	"github.com/ravishen/bookstore_users-api/utils/errors"
)

func getUserId(c *gin.Context) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		return 0, err
	}
	return userId, nil

}

func Update(c *gin.Context) {
	var user users.User
	userId, userErr := getUserId(c)
	if userErr != nil {
		c.JSON(http.StatusNotFound, userErr)
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("invalid json body")
		c.JSON(restError.Status, restError)

		return
	}
	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch
	result, err := services.UserService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))

}

func Get(c *gin.Context) {
	userId, userErr := getUserId(c)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.UserService.GetUser(userId)
	if getErr != nil {
		err := errors.NewBadRequestError(getErr.Message)
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

func Create(c *gin.Context) {
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
	result, err_svc := services.UserService.CreateUser(user)
	if err_svc != nil {
		c.JSON(err_svc.Status, err_svc)
	}
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))

}
func SearchUser(c *gin.Context) {

}
func Delete(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	if err2 := services.UserService.DeleteUser(userId); err2 != nil {
		c.JSON(err.Status, err2)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "Deleted"})
	return
}

func Search(c *gin.Context) {
	users, err := services.UserService.Search(c.Query("status"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
	return

}
