package security

import (
	"net/http"
	"strings"

	"delivery_app_api.mmedic.com/m/v2/src/utils/env_utils"
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

		claims := &jwt_utils.Claims{}

		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(env_utils.GetEnvVar("SECRET")), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.Status(http.StatusUnauthorized)
				return
			}
			c.Status(http.StatusBadRequest)
			return
		}

		if !tkn.Valid {
			c.Status(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
