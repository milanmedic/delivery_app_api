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

func Authenticate(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := &AuthHeader{}

		if err := c.ShouldBindHeader(&authHeader); err != nil {
			c.Status(http.StatusInternalServerError)
			c.Abort()
			return
		}

		if authHeader == nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		if len(authHeader.IDToken) <= 0 {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		token := strings.Split(authHeader.IDToken, " ")[1]

		valid, err := jwt_utils.ValidateToken(token)
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.Status(http.StatusUnauthorized)
				c.Abort()
				return
			}
			c.Status(http.StatusBadRequest)
			c.Abort()
			return
		}

		if !valid {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		_, claims, err := jwt_utils.ParseToken(token)
		if err != nil {
			c.Error(fmt.Errorf("Error while parsting the token. \nReason: %s", err.Error()))
			c.Status(http.StatusInternalServerError)
			c.Abort()
			return
		}

		var passed bool
		for _, role := range roles {
			if strings.Compare(claims.Role, role) == 0 {
				passed = true
				break
			}
		}
		if !passed {
			c.Error(fmt.Errorf("Unauthorized!"))
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		c.Set("user_id", claims.UserId)
		c.Next()
	}
}

func RefreshToken(c *gin.Context) {
	authHeader := &AuthHeader{}

	err := c.ShouldBindHeader(&authHeader)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	token := strings.Split(authHeader.IDToken, " ")[1]

	_, claims, err := jwt_utils.ParseToken(token)
	if err != nil {
		c.Error(fmt.Errorf("Error while parsting the token. \nReason: %s", err.Error()))
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
