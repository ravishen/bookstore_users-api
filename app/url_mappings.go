package app

import (
	ping "github.com/ravishen/bookstore_users-api/controllers/ping"
	users "github.com/ravishen/bookstore_users-api/controllers/users"
)

func mapUrl() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.Get)
	//	router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", users.Create)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/search", users.Search)

}
