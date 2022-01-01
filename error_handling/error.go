package error_handling

import (
	"log"
)

// StroodleError custom error type.
type CustomError struct {
	message string
	status  int // HTTP status codes as registered with IANA.
}

// Error returns error message.
func (e CustomError) Error() string {
	return e.message
}

// Status returns HTTP status code as registered with IANA.
func (e CustomError) Status() int {
	return e.status
}

// NewError returns an initialized StroodleError.
func NewError(message string, status int) error {
	return &CustomError{
		message: message,
		status:  status,
	}
}

// CheckError checks if the error is not nil, and eventually panics.
func CheckError(err error) {
	if err != nil {
		log.Printf("Error: %s", err.Error())
		panic(ErrServerError)
	}
}

// CheckDbError checks if the error is not nil and different from gorm.ErrRecordNotFound, and eventually panics.
func CheckDbError(err error) {
	if err != nil {
		log.Printf("DB error: %s", err)
		panic(ErrServerError)
	}
}

// ThrowError panics an application error that will be recovered in middleware.
func ThrowError(err error) {
	panic(err)
}
