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
	commonMiddlewares(r)

	setupPublicRouter(r, ctrl)
	r.Mount("/", setupPrivateRouter(ctrl))

	return r
}

func commonMiddlewares(r *chi.Mux) {
	r.Use(corsMiddleware)
	r.Use(middleware.NoCache)
	r.Use(iDMiddleware)
	r.Use(iPPathLimitMiddleware)
	r.Use(loggerMiddleware)
	r.Use(recoveryPanicMdlw)
}

// setup all private routes that needs authentication
func setupPublicRouter(r *chi.Mux, ctrl controller.ControllerInterface) *chi.Mux {
	// ping endpoint
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Pong"))
	})

	if config.Config.ShowDocs {
		// swagger docs endpoint
		r.Get("/docs/*", httpSwagger.WrapHandler)
	}
	return r
}

// setup all private routes that needs authentication
func setupPrivateRouter(ctrl controller.ControllerInterface) *chi.Mux {
	r := chi.NewRouter()
	// add authentication middleware for the private routes
	r.Use(checkAuthMdlw)

	r.Post("/device", ctrl.InsertData)

	return r
}
