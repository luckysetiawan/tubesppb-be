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
	ID       int    `json:"id"`
	Name     string `json:"name"`
	UserType int    `json:"user_type"`
	jwt.StandardClaims
}

func generateToken(c *gin.Context, id int, name string, userType int) {
	tokenExpiryTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		ID:       id,
		Name:     name,
		UserType: userType,
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

func Authenticate(accessType int) gin.HandlerFunc {
	return func(c *gin.Context) {
		isValidToken := validateUserToken(c, accessType)
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

func validateUserToken(c *gin.Context, accessType int) bool {
	isAccessTokenValid, id, email, userType := validateTokenFromCookies(c)
	fmt.Print(id, email, userType, accessType, isAccessTokenValid)

	if isAccessTokenValid {
		isUserValid := userType == accessType
		fmt.Print(isUserValid)
		if isUserValid {
			return true
		}
	}
	return false
}

func validateTokenFromCookies(c *gin.Context) (bool, int, string, int) {
	if cookie, err := c.Cookie(tokenName); err == nil {
		accessToken := cookie
		accessClaims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err == nil && parsedToken.Valid {
			return true, accessClaims.ID, accessClaims.Name, accessClaims.UserType
		}
	}
	return false, -1, "", -1
}
