package validations

import (
	"fmt"
	"net/mail"
	"strings"

	"delivery_app_api.mmedic.com/m/v2/src/dto"
	passwordvalidator "github.com/wagslane/go-password-validator"
)

func ValidateName(name string) error {
	if len(name) > 30 {
		return fmt.Errorf("Name longer than 30 characters.")
	}
	return nil
}

func ValidateSurname(surname string) error {
	if len(surname) > 30 {
		return fmt.Errorf("Surname longer than 30 characters.")
	}
	return nil
}

func ValidateUsername(username string) error {
	if len(username) > 30 {
		return fmt.Errorf("Name longer than 30 characters.")
	}
	return nil
}

func ValidatePassword(pass string) error {
	// PASSWORD VALIDATION
	// entropy is a float64, representing the strength in base 2 (bits)
	// if the password has enough entropy, err is nil
	// otherwise, a formatted error message is provided explaining
	// how to increase the strength of the password
	// (safe to show to the client)
	const minEntropyBits = 44
	err := passwordvalidator.Validate(pass, minEntropyBits)
	if err != nil {
		return fmt.Errorf("Weak password!. %s", err.Error())
	}
	return nil
}

func ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return fmt.Errorf("Invalid email. %s", err.Error())
	}
	return nil
}

//TODO: Validate Date

func ValidateCity(city string) error {
	if len(city) > 100 {
		return fmt.Errorf("City longer than 100 characters.")
	}
	return nil
}

func ValidatePostfix(postfix string) error {
	if len(postfix) > 1 {
		return fmt.Errorf("Postfix longer than 1 character.")
	}
	return nil
}

func ValidateStreetNum(num int) error {
	if num < 0 {
		return fmt.Errorf("Street num must be larger than 0.")
	}
	return nil
}

func ValidateStreet(street string) error {
	if len(street) > 100 {
		return fmt.Errorf("Street name longer than 100 characters.")
	}
	return nil
}

func ValidateAddress(addr dto.AddressInputDto) bool {

	if strings.Compare(addr.City, "") == 0 {
		return false
	}

	if strings.Compare(addr.Street, "") == 0 {
		return false
	}

	return true
}
