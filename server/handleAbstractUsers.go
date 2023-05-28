package server

import (
	"database/sql"
	"net/http"

	"github.com/mrityunjaygr8/clean/internal/services"
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

		service := services.NewAbstractUserService(r.Context(), tx)

		abstractUser, err := service.CreateAbstractUser(req.Email, req.Password)
		if err != nil {
			if _, ok := err.(services.ErrAbstractUserExists); ok {
				a.l.Error().Err(err).Msg("User Exists")
				errMsg := make(map[string]interface{})
				errMsg["Email"] = err.Error()
				a.writeJSON(w, http.StatusBadRequest, envelope{"errors": errMsg}, nil)
				return
			}

			a.l.Error().Err(err).Str("email", req.Email).Msg("")
			errMsg := make(map[string]interface{})
			errMsg["Extra"] = err.Error()
			a.writeJSON(w, http.StatusInternalServerError, envelope{"errors": errMsg}, nil)
			return
		}
		if ok := a.transactionCommit(w, tx); ok {
			return
		}

		resp := response{
			Email: abstractUser.Email,
			ID:    abstractUser.ID,
		}

		a.writeJSON(w, http.StatusCreated, envelope{"abstract-user": resp}, nil)
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

		service := services.NewAbstractUserService(r.Context(), a.db)

		abstractUsers, err := service.List()
		if err != nil {
			a.l.Error().Err(err).Msg("")
			errMsg := make(map[string]interface{})
			errMsg["Extra"] = err.Error()
			a.writeJSON(w, http.StatusInternalServerError, envelope{"errors": errMsg}, nil)
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
