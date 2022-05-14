package error_handling

import "net/http"

/*
	500 Internal Server Error -  The server encountered an unexpected condition which prevented it from fulfilling the request.
*/

// ErrServerError is raised when server breaks for internal reasons
var ErrServerError = &CustomError{message: "ERR_INTERNAL_SERVER_ERROR", status: http.StatusInternalServerError}

// ErrBadSyntax is raised when user provides a form or body with missing or invalid fields.
var ErrBadSyntax = &CustomError{message: "ERR_BAD_SYNTAX", status: http.StatusBadRequest}
