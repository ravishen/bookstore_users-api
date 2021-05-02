package app

import (
	"github.com/gin-gonic/gin"
	"github.com/ravishen/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrl()
	logger.Info("Starting Application")
	router.Run(":8000")
}
