package router

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	"github.com/mohsin123321/cloud-project/config"
	"github.com/mohsin123321/cloud-project/error_handling"
	"github.com/mohsin123321/cloud-project/model"
	"github.com/mohsin123321/cloud-project/utility"
	"github.com/urfave/negroni"
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
			err := error_handling.ErrRequestLimitReached()
			handleError(r, w, err, error_handling.LogMessageErrorResponse)
		}),
	)
	return limiter(next)
}

// recover the panic called by an api
func recoveryPanicMdlw(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				if err, ok := err.(*error_handling.CustomError); ok {
					handleError(r, w, err, error_handling.LogMessageErrorResponse)
				} else {
					handleError(r, w, error_handling.ErrServerError(), error_handling.LogMessageErrorResponse)
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
			handleError(r, w, error_handling.ErrMissingToken(), error_handling.LogMessageErrorResponse)
			return
		}

		if _, err = jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Config.Token.Secret), nil
		}); err != nil {
			handleError(r, w, error_handling.ErrInvalidToken(), error_handling.LogMessageErrorResponse)
			return
		}

		// Serve the request
		next.ServeHTTP(w, r)
	})
}

// loggerMiddleware log success and error responses details.
func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		id, ok := r.Context().Value(model.GetCtxKeyID()).(uuid.UUID)
		if !ok {
			panic(error_handling.ErrServerError)
		}

		// response writer wrapper
		rww := negroni.NewResponseWriter(w)

		defer func() {
			elapsedTime := time.Since(start).String()
			log.Println(id.String(), r.Method, r.RemoteAddr, rww.Status(), r.URL.Path, elapsedTime)
		}()

		next.ServeHTTP(rww, r)
	})
}

// iDMiddleware add an unique id to each request, useful for logging.
//
// Each request has at least one info log, but error requests could have more than one log for each error.
// The id is printed together with each log to identify the relative request.
func iDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.NewV4()
		if err != nil {
			handleError(r, w, error_handling.ErrServerError(), error_handling.LogMessageErrorResponse)
			return
		}
		ctx := context.WithValue(r.Context(), model.GetCtxKeyID(), id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func handleError(r *http.Request, w http.ResponseWriter, err error, errLogType string) {
	err = error_handling.PropagateError(err, 2)

	id, ok := r.Context().Value(model.GetCtxKeyID()).(uuid.UUID)
	if !ok {
		panic(error_handling.ErrServerError)
	}

	appErr, ok := err.(*error_handling.CustomError)
	if !ok {
		appErr = error_handling.ErrServerError()
	}

	appErr.ID = id
	appErr.Log(errLogType)

	json := appErr.MarhsalJSON()
	w.Header().Set("Content-Type", utility.MimeTypeJSON)
	w.WriteHeader(appErr.Status())
	_, _ = w.Write(json)
}
