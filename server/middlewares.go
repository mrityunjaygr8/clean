package server

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func (app *Server) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.l.Error().Stack().Err(errors.WithStack(fmt.Errorf("%s", err))).Msg("")
				app.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
