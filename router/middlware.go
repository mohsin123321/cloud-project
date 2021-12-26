package router

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func loggingMiddleware(h http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, h)
}
