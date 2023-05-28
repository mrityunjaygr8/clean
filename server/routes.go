package server

import (
	"net/http"
	"time"

	"github.com/alexliesenfeld/health"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/rs/zerolog/hlog"
)

type request struct {
	Name string `json:"name"`
	Game string `json:"game"`
}

type response struct {
	NameResp string `json:"name"`
	GameResp string `json:"game"`
}

func (a *Server) routes() http.Handler {
	router := chi.NewRouter()

	// middleware.DefaultLogger = middleware.RequestLogger(customLogFormatter{logger: a.l})

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	// router.Use(middleware.Recoverer)
	router.Use(a.recoverPanic)

	router.Use(hlog.NewHandler(*a.l))

	router.Use(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().Str("method", r.Method).Stringer("url", r.URL).Int("status", status).Int("size", size).Dur("duration", duration).Msg("")
	}))

	router.Use(hlog.RemoteAddrHandler("ip"))
	router.Use(hlog.UserAgentHandler("user_agent"))
	router.Use(hlog.RefererHandler("referer"))
	router.Use(hlog.RequestIDHandler("req_id", "Request-id"))

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	router.Post("/ping", func(w http.ResponseWriter, r *http.Request) {
		var req request

		a.readJSON(w, r, &req)

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

	router.Post("/auth/login", a.handleAuthLogin())

	router.Post("/orgs", a.handleOrgCreate())
	router.Get("/orgs", a.handleOrgList())

	router.Post("/orgs/{orgId}/users", a.handleOrgUserCreate())
	router.Get("/orgs/{orgId}/users", a.handleOrgUserList())

	router.Get("/health", health.NewHandler(a.health))

	return router
}
