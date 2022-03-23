package app

import (
	authentication "tubesppb-be/modules/authentication/controller"
	masimelrowoo "tubesppb-be/modules/masimelrowoo/controller"
	user "tubesppb-be/modules/user/controller"

	"github.com/gin-gonic/gin"
)

func MapEndPoints(router *gin.Engine) {
	authentication.Routes(router)
	masimelrowoo.Routes(router)
	user.Routes(router)
}
