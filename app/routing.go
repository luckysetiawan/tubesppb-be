package app

import (
	authentication "tubesppb-be/modules/authentication/controller"

	"github.com/gin-gonic/gin"
)

func MapEndPoints(router *gin.Engine) {
	authentication.Routes(router)
}
