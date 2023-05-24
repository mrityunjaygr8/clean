package server

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	dbmodels "github.com/mrityunjaygr8/clean/db/models"
	"golang.org/x/crypto/bcrypt"
)

func (a *Server) handleAuthLogin() http.HandlerFunc {
	type request struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	type response struct {
		Email string `json:"email"`
		ID    string `json:"id"`
		Admin bool   `json:"admin"`
	}

	var ErrInvalidCreds = errors.New("Invalid email or password supplied")
	return func(w http.ResponseWriter, r *http.Request) {
		var req request

		err := a.readJSON(w, r, &req)
		if err != nil {
			a.l.Error().Err(err).Msg("")
			if strings.Contains(err.Error(), "unknown field") {
				a.writeJSON(w, http.StatusBadRequest, envelope{"error": "The request contains unknown fields"}, nil)
				return
			}
		}
		err = a.validator.Struct(req)
		if err != nil {
			errResp := make(map[string]interface{})
			for _, e := range err.(validator.ValidationErrors) {
				errResp[e.Field()] = e.Translate(a.trans)
			}

			a.writeJSON(w, http.StatusBadRequest, envelope{"errors": errResp}, nil)

			return
		}

		abu, err := dbmodels.AbstractUsers(dbmodels.AbstractUserWhere.Email.EQ(req.Email)).All(r.Context(), a.db)
		if err != nil {
			a.l.Error().Str("error-type", "error admin abstract user lookup").Str("email", req.Email).Err(err).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}

		if len(abu) == 0 {
			a.writeJSON(w, http.StatusBadRequest, envelope{"error": ErrInvalidCreds.Error()}, nil)
			return
		}

		if len(abu) > 1 {
			a.l.Error().Str("error-type", "error multiple abstract users with email").Str("email", req.Email).Err(err).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}

		ab := abu[0]

		err = bcrypt.CompareHashAndPassword([]byte(ab.Password), []byte(req.Password))
		if err != nil {
			if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
				a.writeJSON(w, http.StatusBadRequest, envelope{"error": ErrInvalidCreds.Error()}, nil)
				return
			}

			a.l.Error().Str("error-type", "error comparing password hash").Str("email", req.Email).Err(err).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}

		a.writeJSON(w, http.StatusOK, envelope{"status": "ok"}, nil)
	}
}
