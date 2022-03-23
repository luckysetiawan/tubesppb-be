package service

import (
	"fmt"
	"time"
	"tubesppb-be/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte(config.LoadConfig("JWT_KEY"))
var tokenName = "token"

type Claims struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	FriendMode bool   `json:"user_type"`
	jwt.StandardClaims
}

func generateToken(c *gin.Context, id int, name string, fridendMode bool) {
	tokenExpiryTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		ID:         id,
		Name:       name,
		FriendMode: fridendMode,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpiryTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return
	}

	// c.SetCookie(tokenName, signedToken, tokenExpiryTime, "/", "localhost", false, true)
	c.SetCookie(tokenName, signedToken, 1000, "/", "localhost", false, true)
}

func resetUserToken(c *gin.Context) {
	c.SetCookie(tokenName, "", -1, "/", "localhost", false, true)
}

func Authenticate(friendMode bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		isValidToken := validateUserToken(c, friendMode)
		if !isValidToken {
			// var response UserResponse
			// response.Message = "Unauthorized Access"
			// sendUnAuthorizedResponse(c, response)
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}

func validateUserToken(c *gin.Context, friendMode bool) bool {
	isAccessTokenValid, id, email, userType := validateTokenFromCookies(c)
	fmt.Print(id, email, userType, friendMode, isAccessTokenValid)

	if isAccessTokenValid {
		isUserValid := userType == friendMode
		fmt.Print(isUserValid)
		if isUserValid {
			return true
		}
	}
	return false
}

func validateTokenFromCookies(c *gin.Context) (bool, int, string, bool) {
	if cookie, err := c.Cookie(tokenName); err == nil {
		accessToken := cookie
		accessClaims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err == nil && parsedToken.Valid {
			return true, accessClaims.ID, accessClaims.Name, accessClaims.FriendMode
		}
	}
	return false, -1, "", false
}
