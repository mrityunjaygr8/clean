package server

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	dbmodels "github.com/mrityunjaygr8/clean/db/models"
	"github.com/sirupsen/logrus"
	"github.com/teris-io/shortid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

func (a *Application) handleAdminUserCreate() http.HandlerFunc {
	type request struct {
		Password string `json:"password" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Admin    bool   `json:"admin" validate:"boolean"`
	}
	type response struct {
		Email string `json:"email"`
		ID    string `json:"id"`
		Admin bool   `json:"admin"`
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

		exists, err := dbmodels.AbstractUsers(dbmodels.AbstractUserWhere.Email.EQ(req.Email)).Exists(r.Context(), tx)
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
		var abstractUser dbmodels.AbstractUser
		abstractUser.Email = req.Email

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

		var adminUser dbmodels.AdminUser
		adminUser.Admin = req.Admin

		tmpId, err = shortid.Generate()
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
				"error-type": "error creating admin id",
			}).Errorln(err)
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}

		adminUser.ID = tmpId
		adminUser.UserID = abstractUser.InternalID

		err = adminUser.Insert(r.Context(), tx, boil.Infer())
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
				"error-type": "error inserting admin user",
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
			ID:    adminUser.ID,
			Admin: adminUser.Admin,
		}

		a.writeJSON(w, http.StatusCreated, envelope{"admin-user": resp}, nil)
	}
}

func (a *Application) handleAdminUserList() http.HandlerFunc {
	type adminUser struct {
		Email string `json:"email"`
		Id    string `json:"id"`
		Admin bool   `json:"admin"`
	}
	type response struct {
		Results []adminUser `json:"admin_users"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		adms := make([]adminUser, 0)

		adminUsers, err := dbmodels.AdminUsers(qm.Load(qm.Rels(dbmodels.AdminUserRels.User))).All(r.Context(), a.db)
		if err != nil {
			a.logger.WithFields(logrus.Fields{"error-type": "error listing admin users"}).Errorln(err)
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err}, nil)
			return
		}

		for _, adm := range adminUsers {
			tmp := adminUser{}
			tmp.Admin = adm.Admin
			tmp.Email = adm.R.User.Email
			tmp.Id = adm.ID

			adms = append(adms, tmp)
		}

		a.writeJSON(w, http.StatusOK, envelope{"admin-users": adms}, nil)
	}
}

func (a *Application) handleAdminUserRetrieve() http.HandlerFunc {
	type response struct {
		Email string `json:"email"`
		Id    string `json:"id"`
		Admin bool   `json:"admin"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "adminUserId")
		adm, err := dbmodels.AdminUsers(dbmodels.AdminUserWhere.ID.EQ(id), qm.Load(qm.Rels(dbmodels.AdminUserRels.User))).One(r.Context(), a.db)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				a.writeJSON(w, http.StatusNotFound, envelope{}, nil)
				return

			}
			a.logger.WithFields(logrus.Fields{"error-type": "error retrieving admin users", "admin-user-id": id}).Errorln(err)
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err}, nil)
			return
		}

		resp := response{
			Email: adm.R.User.Email,
			Id:    adm.ID,
			Admin: adm.Admin,
		}

		a.writeJSON(w, http.StatusOK, envelope{"admin-user": resp}, nil)
	}
}

func (a *Application) handleAdminUserDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "adminUserId")

		adms, err := dbmodels.AdminUsers(dbmodels.AdminUserWhere.ID.EQ(id), qm.Load(qm.Rels(dbmodels.AdminUserRels.User))).All(r.Context(), a.db)
		if err != nil {
			a.logger.WithFields(logrus.Fields{"error-type": "error counting admin users", "admin-user-id": id}).Errorln(err)
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err}, nil)
			return
		}

		if len(adms) == 0 {
			a.writeJSON(w, http.StatusNotFound, envelope{}, nil)
			return
		}

		if len(adms) > 1 {
			a.logger.WithFields(logrus.Fields{
				"error-type":    "error id collision",
				"admin-user-id": id,
			}).Errorln(errors.New("admin user id collision"))
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}

		admin := adms[0]
		abstract := admin.R.GetUser()
		abstractId := admin.R.User.ID

		tx, err := a.db.BeginTx(r.Context(), &sql.TxOptions{})
		if err != nil {
			a.logger.WithFields(logrus.Fields{
				"error-type": "error creating transaction",
			}).Errorln(err)
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}
		_, err = admin.Delete(r.Context(), tx)
		if err != nil {
			a.logger.WithFields(logrus.Fields{"error-type": "error deleting admin users", "admin-user-id": id}).Errorln(err)
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err}, nil)
			return
		}

		_, err = abstract.Delete(r.Context(), tx)
		if err != nil {
			a.logger.WithFields(logrus.Fields{"error-type": "error deleting abstract users", "abstract-user-id": abstractId}).Errorln(err)
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err}, nil)
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

		a.writeJSON(w, http.StatusNoContent, envelope{}, nil)
	}
}

func (a *Application) handleAdminUserUpdate() http.HandlerFunc {
	type request struct {
		Admin bool `json:"admin"`
	}

	type response struct {
		Email string `json:"email"`
		Id    string `json:"id"`
		Admin bool   `json:"admin"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "adminUserId")

		var req request
		err := a.readJSON(w, r, &req)

		if err != nil {
			a.logger.Errorln(err)
			if strings.Contains(err.Error(), "unknown field") {
				a.writeJSON(w, http.StatusBadRequest, envelope{"error": "The request contains unknown fields"}, nil)
				return
			}
		}

		adms, err := dbmodels.AdminUsers(dbmodels.AdminUserWhere.ID.EQ(id), qm.Load(qm.Rels(dbmodels.AdminUserRels.User))).All(r.Context(), a.db)
		if err != nil {
			a.logger.WithFields(logrus.Fields{"error-type": "error counting admin users", "admin-user-id": id}).Errorln(err)
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err}, nil)
			return
		}

		if len(adms) == 0 {
			a.writeJSON(w, http.StatusNotFound, envelope{}, nil)
			return
		}

		if len(adms) > 1 {
			a.logger.WithFields(logrus.Fields{
				"error-type":    "error id collision",
				"admin-user-id": id,
			}).Errorln(errors.New("admin user id collision"))
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}

		admin := adms[0]
		admin.Admin = req.Admin

		tx, err := a.db.BeginTx(r.Context(), &sql.TxOptions{})
		if err != nil {
			a.logger.WithFields(logrus.Fields{
				"error-type": "error creating transaction",
			}).Errorln(err)
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}
		_, err = admin.Update(r.Context(), tx, boil.Infer())
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
				"error-type": "error updating admin user",
			}).Errorln(err)
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err}, nil)
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
			Email: admin.R.User.Email,
			Id:    admin.ID,
			Admin: admin.Admin,
		}

		a.writeJSON(w, http.StatusOK, envelope{"admin-user": resp}, nil)
	}
}

func (a *Application) handleAdminUserUpdatePassword() http.HandlerFunc {
	type request struct {
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "adminUserId")

		var req request
		err := a.readJSON(w, r, &req)

		if err != nil {
			a.logger.Errorln(err)
			if strings.Contains(err.Error(), "unknown field") {
				a.writeJSON(w, http.StatusBadRequest, envelope{"error": "The request contains unknown fields"}, nil)
				return
			}
		}

		adms, err := dbmodels.AdminUsers(dbmodels.AdminUserWhere.ID.EQ(id), qm.Load(qm.Rels(dbmodels.AdminUserRels.User))).All(r.Context(), a.db)
		if err != nil {
			a.logger.WithFields(logrus.Fields{"error-type": "error counting admin users", "admin-user-id": id}).Errorln(err)
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err}, nil)
			return
		}

		if len(adms) == 0 {
			a.writeJSON(w, http.StatusNotFound, envelope{}, nil)
			return
		}

		if len(adms) > 1 {
			a.logger.WithFields(logrus.Fields{
				"error-type":    "error id collision",
				"admin-user-id": id,
			}).Errorln(errors.New("admin user id collision"))
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}

		admin := adms[0]
		abstract := admin.R.GetUser()

		tx, err := a.db.BeginTx(r.Context(), &sql.TxOptions{})
		if err != nil {
			a.logger.WithFields(logrus.Fields{
				"error-type": "error creating transaction",
			}).Errorln(err)
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}

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

		abstract.Password = string(hash)
		_, err = abstract.Update(r.Context(), tx, boil.Infer())

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
		txCommitErr := tx.Commit()
		if txCommitErr != nil {
			a.logger.WithFields(logrus.Fields{
				"error-type": "error committing transaction",
			}).Errorln(txCommitErr)
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}

		a.writeJSON(w, http.StatusOK, envelope{"status": fmt.Sprintf("Password for %s updated successfully", admin.R.User.Email)}, nil)
	}
}
