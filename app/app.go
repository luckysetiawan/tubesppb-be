package app

import (
	"tubesppb-be/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartApplication() {
	serverPort := config.LoadConfig("PORT")

	router := gin.Default()
	router.Use(cors.Default())

	MapEndPoints(router)

	router.Run(":" + serverPort)
}
