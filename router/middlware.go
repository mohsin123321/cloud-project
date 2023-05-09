package router

import (
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/golang-jwt/jwt"
	"github.com/mohsin123321/cloud-project/config"
	"github.com/mohsin123321/cloud-project/error_handling"
	"github.com/mohsin123321/cloud-project/model"
)

// corsMiddleware set CORS for requests.
func corsMiddleware(next http.Handler) http.Handler {
	cors := cors.Handler(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodPut, http.MethodPatch, http.MethodOptions, http.MethodConnect, http.MethodTrace},
		AllowedHeaders:     []string{"*"},
		ExposedHeaders:     []string{"*"},
		AllowCredentials:   false,
		MaxAge:             0,
		OptionsPassthrough: false,
		Debug:              false,
	})

	return cors(next)
}

// iPPathLimitMiddleware puts the unique limit counter for the each user (by IP) per endpoint.
func iPPathLimitMiddleware(next http.Handler) http.Handler {
	limiter := httprate.Limit(config.Config.RateLimiter.MaxReqPerIP,
		time.Duration(config.Config.RateLimiter.SecondsWindow)*time.Second,
		httprate.WithKeyFuncs(httprate.KeyByIP, httprate.KeyByEndpoint),
		// response writer when the request limit has been reached
		httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
			err := error_handling.ErrRequestLimitReached
			http.Error(w, err.Error(), err.Status())
		}),
	)
	return limiter(next)
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

func checkAuthMdlw(next http.Handler) http.Handler {
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
		next.ServeHTTP(w, r)
	})
}
