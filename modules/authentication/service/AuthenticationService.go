package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	AuthenticationRepository "tubesppb-be/modules/authentication/repository"
	User "tubesppb-be/modules/users/data"
	Response "tubesppb-be/response"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	rows, _ := AuthenticationRepository.VerifyLogin(username, password)

	var userScan User.User
	var user User.User

	for rows.Next() {
		if err := rows.Scan(&userScan.UID, &userScan.Username, &userScan.Profile_picture, &userScan.Friend_mode); err != nil {
			log.Fatal(err.Error())
		} else {
			user = userScan
		}
	}

	var response Response.GeneralResponse

	if user != (User.User{}) {
		generateToken(c, user.UID, user.Username, user.Friend_mode)
		response.Message = "OK"
		response.Status = 200
		response.Data = user
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, response)
	} else {
		response.Message = "Unauthorized"
		response.Status = 401
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusUnauthorized, response)
	}
}

func Logout(c *gin.Context) {
	resetUserToken(c)

	var response Response.GeneralResponse
	response.Message = "OK"
	response.Status = 200
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}
