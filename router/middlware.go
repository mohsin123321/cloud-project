package router

import (
	"log"
	"net/http"
	"runtime"

	"github.com/golang-jwt/jwt"
	"github.com/mohsin123321/cloud-project/config"
	"github.com/mohsin123321/cloud-project/error_handling"
	"github.com/mohsin123321/cloud-project/model"
)

// // log into the terminal all the informations about a call to an api
// func loggingMiddleware(h http.Handler) http.Handler {
// 	return handlers.LoggingHandler(os.Stdout, h)
// }

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

func checkAuthMdlw(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		var tokenStr string

		if tokenStr = r.Header.Get(model.TokenHeader); tokenStr == "" {
			error_handling.ThrowError(error_handling.ErrMissingToken)
		}

		if _, err = jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Config.Token.Secret), nil
		}); err != nil {
			error_handling.ThrowError(error_handling.ErrInvalidToken)
		}

		// Serve the request
		h.ServeHTTP(w, r)
	})
}
