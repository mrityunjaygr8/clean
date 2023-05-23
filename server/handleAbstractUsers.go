package server

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	dbmodels "github.com/mrityunjaygr8/clean/db/models"
	"github.com/sirupsen/logrus"
	"github.com/teris-io/shortid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/crypto/bcrypt"
)

func (a *Application) handleAbstractUserCreate() http.HandlerFunc {
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
		if err != nil {
			a.logger.Errorln(err)
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

			a.writeJSON(w, http.StatusInternalServerError, envelope{"errors": errResp}, nil)

			return
		}

		tx, err := a.db.BeginTx(r.Context(), &sql.TxOptions{})
		if err != nil {
			a.logger.WithFields(logrus.Fields{
				"error-type": "error creating transaction",
			}).Errorln(err)
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}

		var abstractUser dbmodels.AbstractUser
		abstractUser.Email = req.Email

		exists, err := abstractUser.Exists(r.Context(), tx)
		if err != nil {
			txRollErr := tx.Rollback()
			if txRollErr != nil {
				a.logger.WithFields(logrus.Fields{
					"error-type": "error rolling back transaction",
				}).Errorln(txRollErr)
				a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
				return
			}
			a.logger.WithFields(logrus.Fields{
				"error-type": "error checking if abstract user exists",
			}).Errorln(err)
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}
		if exists {
			txRollErr := tx.Rollback()
			if txRollErr != nil {
				a.logger.WithFields(logrus.Fields{
					"error-type": "error rolling back transaction",
				}).Errorln(txRollErr)
				a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
				return
			}
			a.writeJSON(w, http.StatusBadRequest, envelope{"error": "User with email already exists"}, nil)
			return
		}

		tmpId, err := shortid.Generate()
		if err != nil {
			txRollErr := tx.Rollback()
			if txRollErr != nil {
				a.logger.WithFields(logrus.Fields{
					"error-type": "error rolling back transaction",
				}).Errorln(txRollErr)
				a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
				return
			}
			a.logger.WithFields(logrus.Fields{
				"error-type": "error creating abstract id",
			}).Errorln(err)
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}
		abstractUser.ID = tmpId

		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
		if err != nil {
			txRollErr := tx.Rollback()
			if txRollErr != nil {
				a.logger.WithFields(logrus.Fields{
					"error-type": "error rolling back transaction",
				}).Errorln(txRollErr)
				a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
				return
			}
			a.logger.WithFields(logrus.Fields{
				"error-type": "error creating hash from raw password",
			}).Errorln(err)
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}

		abstractUser.Password = string(hash)

		err = abstractUser.Insert(r.Context(), tx, boil.Infer())
		if err != nil {
			txRollErr := tx.Rollback()
			if txRollErr != nil {
				a.logger.WithFields(logrus.Fields{
					"error-type": "error rolling back transaction",
				}).Errorln(txRollErr)
				a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
				return
			}
			a.logger.WithFields(logrus.Fields{
				"error-type": "error inserting abstract user",
			}).Errorln(err)
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}

		txCommitErr := tx.Commit()
		if txCommitErr != nil {
			a.logger.WithFields(logrus.Fields{
				"error-type": "error committing transaction",
			}).Errorln(txCommitErr)
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}

		resp := response{
			Email: abstractUser.Email,
			ID:    abstractUser.ID,
		}

		a.writeJSON(w, http.StatusCreated, envelope{"admin-user": resp}, nil)
	}
}
func (a *Application) handleAbstractUserList() http.HandlerFunc {
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
			a.logger.WithFields(logrus.Fields{"error-type": "error listing admin users"}).Errorln(err)
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
