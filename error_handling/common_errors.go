package error_handling

import (
	"net/http"
)

// ErrServerError is raised when server breaks for internal reasons.
func ErrServerError() *CustomError {
	return &CustomError{
		err:    "ERR_INTERNAL_SERVER_ERROR",
		msg:    "Internal Server Error",
		status: http.StatusInternalServerError,
	}
}

// ErrBadSyntax is raised when user provides a form or body with missing or invalid fields.
func ErrBadSyntax() *CustomError {
	return &CustomError{
		err:    "ERR_BAD_SYNTAX",
		msg:    "Provided request body fields are not valid",
		status: http.StatusBadRequest,
	}
}

// ErrInvalidToken is raised when the token contained in the request is not valid.
func ErrInvalidToken() *CustomError {
	return &CustomError{
		err:    "ERR_INVALID_TOKEN",
		msg:    "Jwt Token is invalid or expired",
		status: http.StatusUnauthorized,
	}
}

// ErrMissingToken is raised when request does not contain a jwt for an API which requires authentication.
func ErrMissingToken() *CustomError {
	return &CustomError{
		err:    "ERR_MISSING_TOKEN",
		msg:    "Jwt Token is missing in the request header",
		status: http.StatusUnauthorized,
	}

}

// ErrRequestLimitReached is raised when the request is spammed by the user
func ErrRequestLimitReached() *CustomError {
	return &CustomError{
		err:    "ERR_REQUEST_LIMIT_REACHED",
		msg:    "Request limit reached for this endpoint",
		status: http.StatusTooManyRequests,
	}

}
