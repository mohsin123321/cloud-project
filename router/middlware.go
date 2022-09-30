package router

import (
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gorilla/handlers"
	"github.com/mohsin123321/cloud-project/error_handling"
)

// log into the terminal all the informations about a call to an api
func loggingMiddleware(h http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, h)
}

// recover the panic called by an api
func recoveryPanicMdlw(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				if appErr, ok := err.(*error_handling.CustomError); ok {
					http.Error(w, appErr.Error(), appErr.Status())
				} else {
					pc, fn, line, _ := runtime.Caller(2)
					log.Printf("UNEXPECTED ERROR in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
					status := http.StatusInternalServerError
					http.Error(w, http.StatusText(status), status)
				}
			}
		}()

		h.ServeHTTP(w, r)
	})
}
