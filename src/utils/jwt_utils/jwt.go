package jwt_utils

import (
	"time"

	"delivery_app_api.mmedic.com/m/v2/src/utils/env_utils"
	"github.com/golang-jwt/jwt"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func CreateClaims() *Claims {
	return new(Claims)
}

func CreateToken(claims *Claims) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	jwtKey := []byte(env_utils.GetEnvVar("SECRET"))
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(token string) (bool, error) {
	tkn, _, err := ParseToken(token)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, err
		}
		return false, err
	}

	if !tkn.Valid {
		return false, nil
	}

	return true, nil
}

func ParseToken(token string) (*jwt.Token, *Claims, error) {
	claims := CreateClaims()

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(env_utils.GetEnvVar("SECRET")), nil
	})

	return tkn, claims, err
}
