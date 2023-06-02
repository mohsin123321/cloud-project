package error_handling

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/mohsin123321/cloud-project/dto"
)

// Log message types are used to define the type of each log.
const (
	LogMessageErrorResponse   = "ERROR RESPONSE"   // Handled errors that could happen.
	LogMessageUnexpectedError = "UNEXPECTED ERROR" // Panic errors that should never happen.
)

// StroodleError custom error type.
type CustomError struct {
	err        string
	msg        string
	status     int // HTTP status codes as registered with IANA.
	stackTrace []string
	ID         uuid.UUID // request uuid
}

// Error returns error.
func (e *CustomError) Error() string {
	return e.err
}

// Status returns error code.
func (e *CustomError) Status() int {
	return e.status
}

// Msg returns error message.
func (e *CustomError) Msg() string {
	return e.msg
}

// AddStackTraceItem appends a stack trace message
func (e *CustomError) AddStackTraceItem(item string) {
	e.stackTrace = append(e.stackTrace, item)
}

// Log displays an error with the right format using the given message.
func (e *CustomError) Log(message string) {
	log.Printf("%s: %s %s %d\n%s", message, e.ID.String(), e.Error(), e.Status(), e.GenerateStackTrace())
}

// GenerateStackTrace joins the string to create the new stack trace message
func (e *CustomError) GenerateStackTrace() string {
	if len(e.stackTrace) == 0 { // empty stacktrace
		return "[]"
	}

	// create stack trace
	return strings.Join(e.stackTrace, "\n")
}

// MarhsalJSON marshals the error in json format
func (e *CustomError) MarhsalJSON() []byte {
	json, _ := json.Marshal(dto.Error{
		Err: e.err,
		Msg: e.msg,
	})
	return json
}

// PropagateError adds the details of the error inside the stack of error
func PropagateError(err error, skips int) error {
	if err == nil {
		return nil
	}

	appErr, ok := err.(*CustomError)
	if !ok {
		appErr = ErrServerError()
	}

	pc, file, line, _ := runtime.Caller(skips)
	funcName := runtime.FuncForPC(pc).Name()

	appErr.AddStackTraceItem(fmt.Sprintf("[%s:%v:%s %s]", file, line, funcName, err.Error()))

	return appErr
}
