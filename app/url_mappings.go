package app

import (
	ping "github.com/ravishen/bookstore_users-api/controllers/ping"
	users "github.com/ravishen/bookstore_users-api/controllers/users"
)

func mapUrl() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	//	router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", users.CreateUser)

}
