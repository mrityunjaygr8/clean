package server

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	dbmodels "github.com/mrityunjaygr8/clean/db/models"
	"github.com/teris-io/shortid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (a *Server) handleOrgCreate() http.HandlerFunc {
	type request struct {
		Name string `json:"name" validate:"required"`
	}

	type response struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	}
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

		tx, err := a.db.BeginTx(r.Context(), &sql.TxOptions{})
		if err != nil {
			a.l.Error().Str("error-type", "error creating transaction").Err(err).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}
		var org dbmodels.Org
		org.Name = req.Name

		tmpId, err := shortid.Generate()
		if err != nil {
			txRollErr := tx.Rollback()
			if txRollErr != nil {
				a.l.Error().Str("error-type", "error rolling back transaction").Err(txRollErr).Msg("")
				a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
				return
			}
			a.l.Error().Str("error-type", "error creating abstract id").Err(err)
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}
		org.ID = tmpId

		err = org.Insert(r.Context(), tx, boil.Infer())
		if err != nil {
			txRollErr := tx.Rollback()
			if txRollErr != nil {
				a.l.Error().Str("error-type", "error rolling back transaction").Err(txRollErr).Msg("")
				a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
				return
			}
			a.l.Error().Str("error-type", "error inserting org").Err(err).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}
		txCommitErr := tx.Commit()
		if txCommitErr != nil {
			a.l.Error().Str("error-type", "error committing transaction").Err(txCommitErr).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}

		resp := response{
			Name: org.Name,
			ID:   org.ID,
		}

		a.writeJSON(w, http.StatusOK, envelope{"org": resp}, nil)
	}
}

func (a *Server) handleOrgList() http.HandlerFunc {
	type orgResp struct {
		Name string `json:"name"`
		Id   string `json:"id"`
	}
	type response struct {
		Results []orgResp `json:"orgs"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		orgResps := make([]orgResp, 0)

		orgs, err := dbmodels.Orgs().All(r.Context(), a.db)
		if err != nil {
			a.l.Error().Str("error-type", "error listing admin users").Err(err).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err}, nil)
			return
		}

		for _, org := range orgs {
			tmp := orgResp{}
			tmp.Name = org.Name
			tmp.Id = org.ID

			orgResps = append(orgResps, tmp)
		}

		a.writeJSON(w, http.StatusOK, envelope{"orgs": orgResps}, nil)
	}
}
