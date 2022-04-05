package security

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	//BECAUSE OF OAUTH REGISTRATION
	if len(password) <= 0 {
		return "", nil
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
