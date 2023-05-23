package server

import (
	"net/http"

	"github.com/alexliesenfeld/health"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type request struct {
	Name string `json:"name"`
	Game string `json:"game"`
}

type response struct {
	NameResp string `json:"name"`
	GameResp string `json:"game"`
}

func (a *Application) routes() http.Handler {
	router := chi.NewRouter()

	middleware.DefaultLogger = middleware.RequestLogger(customLogFormatter{logger: a.logger})

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	router.Post("/ping", func(w http.ResponseWriter, r *http.Request) {
		var req request

		a.readJSON(w, r, &req)
		a.logger.Println(req)

		resp := response{
			NameResp: req.Name,
			GameResp: req.Game,
		}

		a.writeJSON(w, http.StatusOK, envelope{"resp": resp}, nil)

	})

	router.Post("/abstractUser", a.handleAbstractUserCreate())
	router.Get("/abstractUser", a.handleAbstractUserList())

	router.Post("/adminUser", a.handleAdminUserCreate())
	router.Get("/adminUser", a.handleAdminUserList())
	router.Get("/adminUser/{adminUserId}", a.handleAdminUserRetrieve())
	router.Delete("/adminUser/{adminUserId}", a.handleAdminUserDelete())
	router.Put("/adminUser/{adminUserId}", a.handleAdminUserUpdate())
	router.Post("/adminUser/{adminUserId}/password", a.handleAdminUserUpdatePassword())

	router.Get("/health", health.NewHandler(a.health))

	return router
}
