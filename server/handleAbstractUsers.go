package server

import (
	"database/sql"
	"net/http"

	dbmodels "github.com/mrityunjaygr8/clean/db/models"
	"github.com/teris-io/shortid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/crypto/bcrypt"
)

func (a *Server) handleAbstractUserCreate() http.HandlerFunc {
	type request struct {
		Password string `json:"password" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
	}
	type response struct {
		Email string `json:"email"`
		ID    string `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var req request

		err := a.readJSON(w, r, &req)
		if a.checkExtraFieldsError(w, err) {
			return
		}
		err = a.validator.Struct(req)
		if a.checkValidationErrors(w, err) {
			return
		}

		tx, err := a.db.BeginTx(r.Context(), &sql.TxOptions{})
		if a.checkTransactionBeginError(w, err) {
			return
		}

		var abstractUser dbmodels.AbstractUser
		abstractUser.Email = req.Email

		exists, err := abstractUser.Exists(r.Context(), tx)
		if err != nil {
			if a.transactionRollback(w, tx) {
				return
			}
			a.l.Error().Str("error-type", "error checking if abstract user exists").Err(err).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}
		if exists {
			if a.transactionRollback(w, tx) {
				return
			}
			a.writeJSON(w, http.StatusBadRequest, envelope{"error": "User with email already exists"}, nil)
			return
		}

		tmpId, err := shortid.Generate()
		if err != nil {
			if a.transactionRollback(w, tx) {
				return
			}
			a.l.Error().Str("error-type", "error creating abstract id").Err(err).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}
		abstractUser.ID = tmpId

		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
		if err != nil {
			if a.transactionRollback(w, tx) {
				return
			}
			a.l.Error().Str("error-type", "error creating hash from raw password").Err(err).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}

		abstractUser.Password = string(hash)

		err = abstractUser.Insert(r.Context(), tx, boil.Infer())
		if err != nil {
			if a.transactionRollback(w, tx) {
				return
			}
			a.l.Error().Str("error-type", "error inserting abstract user").Err(err).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}

		if ok := a.transactionCommit(w, tx); ok {
			return
		}

		resp := response{
			Email: abstractUser.Email,
			ID:    abstractUser.ID,
		}

		a.writeJSON(w, http.StatusCreated, envelope{"admin-user": resp}, nil)
	}
}

func (a *Server) handleAbstractUserList() http.HandlerFunc {
	type abstractUser struct {
		Email string `json:"email"`
		Id    string `json:"id"`
	}
	type response struct {
		Results []abstractUser `json:"abstract_users"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		abms := make([]abstractUser, 0)

		abstractUsers, err := dbmodels.AbstractUsers().All(r.Context(), a.db)
		if err != nil {
			a.l.Error().Str("error-type", "error listing admin users").Err(err).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err}, nil)
			return
		}

		for _, abm := range abstractUsers {
			tmp := abstractUser{}
			tmp.Email = abm.Email
			tmp.Id = abm.ID

			abms = append(abms, tmp)
		}

		a.writeJSON(w, http.StatusOK, envelope{"abstract-users": abms}, nil)
	}
}
