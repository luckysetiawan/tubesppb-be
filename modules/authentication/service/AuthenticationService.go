package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	dto "tubesppb-be/modules/authentication/data/dto"
	AuthenticationRepository "tubesppb-be/modules/authentication/repository"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	row := AuthenticationRepository.VerifyLogin(username, password)

	var id uint
	err := row.Scan(&id)

	var dto dto.AuthenticationDTO
	if err == nil {
		dto.Message = "Login Success"
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, dto)
	} else {
		dto.Message = "Login Failed"
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusUnauthorized, dto)
	}
}
