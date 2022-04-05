package security

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"delivery_app_api.mmedic.com/m/v2/src/utils/jwt_utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AuthHeader struct {
	IDToken string `header:"Authorization"`
}

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := &AuthHeader{}

		if err := c.ShouldBindHeader(&authHeader); err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		if authHeader == nil {
			c.Status(http.StatusUnauthorized)
			return
		}

		token := strings.Split(authHeader.IDToken, " ")[1]

		valid, err := jwt_utils.ValidateToken(token)
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.Status(http.StatusUnauthorized)
				return
			}
			c.Status(http.StatusBadRequest)
			return
		}

		if !valid {
			c.Status(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

func RefreshToken(c *gin.Context) {
	authHeader := &AuthHeader{}

	_ = c.ShouldBindHeader(&authHeader)
	token := strings.Split(authHeader.IDToken, " ")[1]

	_, claims, err := jwt_utils.ParseToken(token)
	if err != nil {
		c.Error(fmt.Errorf("Error while creating a customer. \nReason: %s", err.Error()))
		c.Status(http.StatusInternalServerError)
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 24*time.Hour {
		c.Status(http.StatusBadRequest)
		return
	}

	// Now, create a new token for the current use, with a renewed expiration time
	claims = jwt_utils.CreateClaims()
	refreshedToken, err := jwt_utils.CreateToken(claims)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": refreshedToken})
	return
}
