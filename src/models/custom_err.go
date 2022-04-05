package models

type CustomError struct {
	message string
}

func CreateCustomError(msg string) *CustomError {
	return &CustomError{message: msg}
}

func (ce *CustomError) Error() string {
	return ce.message
}
