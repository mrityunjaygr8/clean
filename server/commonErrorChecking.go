package server

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

func (a *Server) transactionRollback(w http.ResponseWriter, tx *sql.Tx) bool {
	txRollErr := tx.Rollback()
	if txRollErr != nil {
		a.l.Error().Str("error-type", "error rolling back transaction").Err(txRollErr).Msg("")
		a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
		return true
	}
	return false
}

func (a *Server) transactionCommit(w http.ResponseWriter, tx *sql.Tx) bool {
	txCommitErr := tx.Commit()
	if txCommitErr != nil {
		a.l.Error().Str("error-type", "error committing transaction").Err(txCommitErr).Msg("")
		a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
		return true
	}
	return false
}

func (a *Server) checkTransactionBeginError(w http.ResponseWriter, err error) bool {
	if err != nil {
		a.l.Error().Str("error-type", "error creating transaction").Err(err).Msg("")
		a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
		return true
	}
	return false
}

func (a *Server) checkValidationErrors(w http.ResponseWriter, err error) bool {
	if err != nil {
		errResp := make(map[string]interface{})
		for _, e := range err.(validator.ValidationErrors) {
			errResp[e.Field()] = e.Translate(a.trans)
		}

		a.writeJSON(w, http.StatusBadRequest, envelope{"errors": errResp}, nil)

		return true
	}
	return false
}

func (a *Server) checkExtraFieldsError(w http.ResponseWriter, err error) bool {
	if err != nil {
		a.l.Error().Err(err).Msg("")
		if strings.Contains(err.Error(), "unknown field") {
			a.writeJSON(w, http.StatusBadRequest, envelope{"error": "The request contains unknown fields"}, nil)
			return true
		}
	}
	return false
}
