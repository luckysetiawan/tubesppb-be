package controller

import (
	"github.com/gin-gonic/gin"

	AuthenticationService "tubesppb-be/modules/authentication/service"
)

func Routes(router *gin.Engine) {
	router.POST("/login", AuthenticationService.Login)
}
