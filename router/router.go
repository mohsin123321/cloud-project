package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mohsin123321/cloud-project/config"
	"github.com/mohsin123321/cloud-project/controller"
	httpSwagger "github.com/swaggo/http-swagger"
)

// routes the request to the right controller
func SetupRoutes(ctrl controller.ControllerInterface) *chi.Mux {
	r := chi.NewRouter()

	// added common middlewares
	r.Use(corsMiddleware, middleware.NoCache, iPPathLimitMiddleware, middleware.Logger, recoveryPanicMdlw)

	setupPublicRouter(r, ctrl)
	setupPrivateRouter(r, ctrl)

	return r
}

// setup all private routes that needs authentication
func setupPublicRouter(router *chi.Mux, ctrl controller.ControllerInterface) {
	// ping endpoint
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Pong"))
	})

	if config.Config.ShowDocs {
		// swagger docs endpoint
		router.Get("/docs/*", httpSwagger.WrapHandler)
	}
}

// setup all private routes that needs authentication
func setupPrivateRouter(router *chi.Mux, ctrl controller.ControllerInterface) {
	r := chi.NewRouter()

	// add authentication middleware for the private routes
	r.Use(checkAuthMdlw)

	r.Post("/device", ctrl.InsertData)
}
