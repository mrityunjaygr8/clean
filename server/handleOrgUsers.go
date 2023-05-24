package server

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	dbmodels "github.com/mrityunjaygr8/clean/db/models"
	"github.com/teris-io/shortid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

func (a *Server) handleOrgUserCreate() http.HandlerFunc {
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
		orgId := chi.URLParam(r, "orgId")
		var req request
		orgs, err := dbmodels.Orgs(dbmodels.OrgWhere.ID.EQ(orgId)).All(r.Context(), a.db)
		if err != nil {
			a.l.Error().Str("error-type", "error counting orgs").Str("org-id", orgId).Err(err).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err}, nil)
			return
		}

		if len(orgs) == 0 {
			a.writeJSON(w, http.StatusNotFound, envelope{}, nil)
			return
		}

		if len(orgs) > 1 {
			a.l.Error().Str("error-type", "error id collision").Str("org-id", orgId).Err(errors.New("org id collision")).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}

		org := orgs[0]

		err = a.readJSON(w, r, &req)
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

		exists, err := dbmodels.AbstractUsers(dbmodels.AbstractUserWhere.Email.EQ(req.Email)).Exists(r.Context(), tx)
		if err != nil {
			txRollErr := tx.Rollback()
			if txRollErr != nil {
				a.l.Error().Str("error-type", "error rolling back transaction").Err(txRollErr).Msg("")
				a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
				return
			}
			a.l.Error().Str("error-type", "error checking if abstract user exists").Err(err).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}
		if exists {
			txRollErr := tx.Rollback()
			if txRollErr != nil {
				a.l.Error().Str("error-type", "error rolling back transaction").Err(txRollErr).Msg("")
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
				a.l.Error().Str("error-type", "error rolling back transaction").Err(txRollErr).Msg("")
				a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
				return
			}
			a.l.Error().Str("error-type", "error creating abstract id").Err(err).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}
		abstractUser.ID = tmpId

		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
		if err != nil {
			txRollErr := tx.Rollback()
			if txRollErr != nil {
				a.l.Error().Str("error-type", "error rolling back transaction").Err(txRollErr).Msg("")
				a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
				return
			}
			a.l.Error().Str("error-type", "error creating hash from raw password").Err(err).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}

		abstractUser.Password = string(hash)

		err = abstractUser.Insert(r.Context(), tx, boil.Infer())
		if err != nil {
			txRollErr := tx.Rollback()
			if txRollErr != nil {
				a.l.Error().Str("error-type", "error rolling back transaction").Err(txRollErr).Msg("")
				a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
				return
			}
			a.l.Error().Str("error-type", "error inserting abstract user").Err(err).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}

		var orgUser dbmodels.OrgUser
		orgUser.Admin = req.Admin

		tmpId, err = shortid.Generate()
		if err != nil {
			txRollErr := tx.Rollback()
			if txRollErr != nil {
				a.l.Error().Str("error-type", "error rolling back transaction").Err(txRollErr).Msg("")
				a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
				return
			}
			a.l.Error().Str("error-type", "error creating org id").Err(err).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}

		orgUser.ID = tmpId
		orgUser.UserID = abstractUser.InternalID
		orgUser.OrgID = org.InternalID

		err = orgUser.Insert(r.Context(), tx, boil.Infer())
		if err != nil {
			txRollErr := tx.Rollback()
			if txRollErr != nil {
				a.l.Error().Str("error-type", "error rolling back transaction").Err(txRollErr).Msg("")
				a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
				return
			}
			a.l.Error().Str("error-type", "error inserting org user").Err(err).Msg("")
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
			Email: abstractUser.Email,
			ID:    orgUser.ID,
			Admin: orgUser.Admin,
		}

		a.writeJSON(w, http.StatusCreated, envelope{"org-user": resp}, nil)
	}
}

func (a *Server) handleOrgUserList() http.HandlerFunc {
	type orgUser struct {
		Email string `json:"email"`
		Id    string `json:"id"`
		Admin bool   `json:"admin"`
	}
	type response struct {
		Results []orgUser `json:"admin_users"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		orgId := chi.URLParam(r, "orgId")
		orgs, err := dbmodels.Orgs(dbmodels.OrgWhere.ID.EQ(orgId)).All(r.Context(), a.db)
		if err != nil {
			a.l.Error().Str("error-type", "error counting orgs").Str("org-id", orgId).Err(err).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err}, nil)
			return
		}

		if len(orgs) == 0 {
			a.writeJSON(w, http.StatusNotFound, envelope{}, nil)
			return
		}

		if len(orgs) > 1 {
			a.l.Error().Str("error-type", "error id collision").Str("org-id", orgId).Err(errors.New("org id collision")).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}

		org := orgs[0]
		orgUsers := make([]orgUser, 0)

		orgUsersDB, err := dbmodels.OrgUsers(qm.Load(qm.Rels(dbmodels.OrgUserRels.User)), dbmodels.OrgUserWhere.OrgID.EQ(org.InternalID)).All(r.Context(), a.db)
		if err != nil {
			a.l.Error().Str("error-type", "error listing org users").Err(err).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err}, nil)
			return
		}

		for _, orgU := range orgUsersDB {
			tmp := orgUser{}
			tmp.Admin = orgU.Admin
			tmp.Email = orgU.R.User.Email
			tmp.Id = orgU.ID

			orgUsers = append(orgUsers, tmp)
		}

		a.writeJSON(w, http.StatusOK, envelope{"org-users": orgUsers}, nil)
	}
}

//
// func (a *Application) handleAdminUserRetrieve() http.HandlerFunc {
// 	type response struct {
// 		Email string `json:"email"`
// 		Id    string `json:"id"`
// 		Admin bool   `json:"admin"`
// 	}
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		id := chi.URLParam(r, "adminUserId")
// 		adm, err := dbmodels.AdminUsers(dbmodels.AdminUserWhere.ID.EQ(id), qm.Load(qm.Rels(dbmodels.AdminUserRels.User))).One(r.Context(), a.db)
// 		if err != nil {
// 			if errors.Is(err, sql.ErrNoRows) {
// 				a.writeJSON(w, http.StatusNotFound, envelope{}, nil)
// 				return
//
// 			}
// 			a.logger.WithFields(logrus.Fields{"error-type": "error retrieving admin users", "admin-user-id": id}).Errorln(err)
// 			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err}, nil)
// 			return
// 		}
//
// 		resp := response{
// 			Email: adm.R.User.Email,
// 			Id:    adm.ID,
// 			Admin: adm.Admin,
// 		}
//
// 		a.writeJSON(w, http.StatusOK, envelope{"admin-user": resp}, nil)
// 	}
// }
//
// func (a *Application) handleAdminUserDelete() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		id := chi.URLParam(r, "adminUserId")
//
// 		adms, err := dbmodels.AdminUsers(dbmodels.AdminUserWhere.ID.EQ(id), qm.Load(qm.Rels(dbmodels.AdminUserRels.User))).All(r.Context(), a.db)
// 		if err != nil {
// 			a.logger.WithFields(logrus.Fields{"error-type": "error counting admin users", "admin-user-id": id}).Errorln(err)
// 			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err}, nil)
// 			return
// 		}
//
// 		if len(adms) == 0 {
// 			a.writeJSON(w, http.StatusNotFound, envelope{}, nil)
// 			return
// 		}
//
// 		if len(adms) > 1 {
// 			a.logger.WithFields(logrus.Fields{
// 				"error-type":    "error id collision",
// 				"admin-user-id": id,
// 			}).Errorln(errors.New("admin user id collision"))
// 			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
// 			return
// 		}
//
// 		admin := adms[0]
// 		abstract := admin.R.GetUser()
// 		abstractId := admin.R.User.ID
//
// 		tx, err := a.db.BeginTx(r.Context(), &sql.TxOptions{})
// 		if err != nil {
// 			a.logger.WithFields(logrus.Fields{
// 				"error-type": "error creating transaction",
// 			}).Errorln(err)
// 			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
// 			return
// 		}
// 		_, err = admin.Delete(r.Context(), tx)
// 		if err != nil {
// 			a.logger.WithFields(logrus.Fields{"error-type": "error deleting admin users", "admin-user-id": id}).Errorln(err)
// 			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err}, nil)
// 			return
// 		}
//
// 		_, err = abstract.Delete(r.Context(), tx)
// 		if err != nil {
// 			a.logger.WithFields(logrus.Fields{"error-type": "error deleting abstract users", "abstract-user-id": abstractId}).Errorln(err)
// 			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err}, nil)
// 			return
// 		}
//
// 		txCommitErr := tx.Commit()
// 		if txCommitErr != nil {
// 			a.logger.WithFields(logrus.Fields{
// 				"error-type": "error committing transaction",
// 			}).Errorln(txCommitErr)
// 			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
// 			return
// 		}
//
// 		a.writeJSON(w, http.StatusNoContent, envelope{}, nil)
// 	}
// }
//
// func (a *Application) handleAdminUserUpdate() http.HandlerFunc {
// 	type request struct {
// 		Admin bool `json:"admin" validate:"required,boolean"`
// 	}
//
// 	type response struct {
// 		Email string `json:"email"`
// 		Id    string `json:"id"`
// 		Admin bool   `json:"admin"`
// 	}
//
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		id := chi.URLParam(r, "adminUserId")
//
// 		var req request
// 		err := a.readJSON(w, r, &req)
//
// 		if err != nil {
// 			a.logger.Errorln(err)
// 			if strings.Contains(err.Error(), "unknown field") {
// 				a.writeJSON(w, http.StatusBadRequest, envelope{"error": "The request contains unknown fields"}, nil)
// 				return
// 			}
// 		}
// 		err = a.validator.Struct(req)
// 		if err != nil {
// 			errResp := make(map[string]interface{})
// 			for _, e := range err.(validator.ValidationErrors) {
// 				errResp[e.Field()] = e.Translate(a.trans)
// 			}
//
// 			a.writeJSON(w, http.StatusBadRequest, envelope{"errors": errResp}, nil)
//
// 			return
// 		}
//
// 		adms, err := dbmodels.AdminUsers(dbmodels.AdminUserWhere.ID.EQ(id), qm.Load(qm.Rels(dbmodels.AdminUserRels.User))).All(r.Context(), a.db)
// 		if err != nil {
// 			a.logger.WithFields(logrus.Fields{"error-type": "error counting admin users", "admin-user-id": id}).Errorln(err)
// 			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err}, nil)
// 			return
// 		}
//
// 		if len(adms) == 0 {
// 			a.writeJSON(w, http.StatusNotFound, envelope{}, nil)
// 			return
// 		}
//
// 		if len(adms) > 1 {
// 			a.logger.WithFields(logrus.Fields{
// 				"error-type":    "error id collision",
// 				"admin-user-id": id,
// 			}).Errorln(errors.New("admin user id collision"))
// 			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
// 			return
// 		}
//
// 		admin := adms[0]
// 		admin.Admin = req.Admin
//
// 		tx, err := a.db.BeginTx(r.Context(), &sql.TxOptions{})
// 		if err != nil {
// 			a.logger.WithFields(logrus.Fields{
// 				"error-type": "error creating transaction",
// 			}).Errorln(err)
// 			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
// 			return
// 		}
// 		_, err = admin.Update(r.Context(), tx, boil.Infer())
// 		if err != nil {
// 			txRollErr := tx.Rollback()
// 			if txRollErr != nil {
// 				a.logger.WithFields(logrus.Fields{
// 					"error-type": "error rolling back transaction",
// 				}).Errorln(txRollErr)
// 				a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
// 				return
// 			}
// 			a.logger.WithFields(logrus.Fields{
// 				"error-type": "error updating admin user",
// 			}).Errorln(err)
// 			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err}, nil)
// 			return
// 		}
// 		txCommitErr := tx.Commit()
// 		if txCommitErr != nil {
// 			a.logger.WithFields(logrus.Fields{
// 				"error-type": "error committing transaction",
// 			}).Errorln(txCommitErr)
// 			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
// 			return
// 		}
//
// 		resp := response{
// 			Email: admin.R.User.Email,
// 			Id:    admin.ID,
// 			Admin: admin.Admin,
// 		}
//
// 		a.writeJSON(w, http.StatusOK, envelope{"admin-user": resp}, nil)
// 	}
// }
//
// func (a *Application) handleAdminUserUpdatePassword() http.HandlerFunc {
// 	type request struct {
// 		Password string `json:"password" validate:"required"`
// 	}
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		id := chi.URLParam(r, "adminUserId")
//
// 		var req request
// 		err := a.readJSON(w, r, &req)
//
// 		if err != nil {
// 			a.logger.Errorln(err)
// 			if strings.Contains(err.Error(), "unknown field") {
// 				a.writeJSON(w, http.StatusBadRequest, envelope{"error": "The request contains unknown fields"}, nil)
// 				return
// 			}
// 		}
// 		err = a.validator.Struct(req)
// 		if err != nil {
// 			errResp := make(map[string]interface{})
// 			for _, e := range err.(validator.ValidationErrors) {
// 				errResp[e.Field()] = e.Translate(a.trans)
// 			}
//
// 			a.writeJSON(w, http.StatusBadRequest, envelope{"errors": errResp}, nil)
//
// 			return
// 		}
//
// 		adms, err := dbmodels.AdminUsers(dbmodels.AdminUserWhere.ID.EQ(id), qm.Load(qm.Rels(dbmodels.AdminUserRels.User))).All(r.Context(), a.db)
// 		if err != nil {
// 			a.logger.WithFields(logrus.Fields{"error-type": "error counting admin users", "admin-user-id": id}).Errorln(err)
// 			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err}, nil)
// 			return
// 		}
//
// 		if len(adms) == 0 {
// 			a.writeJSON(w, http.StatusNotFound, envelope{}, nil)
// 			return
// 		}
//
// 		if len(adms) > 1 {
// 			a.logger.WithFields(logrus.Fields{
// 				"error-type":    "error id collision",
// 				"admin-user-id": id,
// 			}).Errorln(errors.New("admin user id collision"))
// 			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
// 			return
// 		}
//
// 		admin := adms[0]
// 		abstract := admin.R.GetUser()
//
// 		tx, err := a.db.BeginTx(r.Context(), &sql.TxOptions{})
// 		if err != nil {
// 			a.logger.WithFields(logrus.Fields{
// 				"error-type": "error creating transaction",
// 			}).Errorln(err)
// 			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
// 			return
// 		}
//
// 		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
// 		if err != nil {
// 			txRollErr := tx.Rollback()
// 			if txRollErr != nil {
// 				a.logger.WithFields(logrus.Fields{
// 					"error-type": "error rolling back transaction",
// 				}).Errorln(txRollErr)
// 				a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
// 				return
// 			}
// 			a.logger.WithFields(logrus.Fields{
// 				"error-type": "error creating hash from raw password",
// 			}).Errorln(err)
// 			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
// 			return
// 		}
//
// 		abstract.Password = string(hash)
// 		_, err = abstract.Update(r.Context(), tx, boil.Infer())
//
// 		if err != nil {
// 			txRollErr := tx.Rollback()
// 			if txRollErr != nil {
// 				a.logger.WithFields(logrus.Fields{
// 					"error-type": "error rolling back transaction",
// 				}).Errorln(txRollErr)
// 				a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
// 				return
// 			}
// 			a.logger.WithFields(logrus.Fields{
// 				"error-type": "error creating hash from raw password",
// 			}).Errorln(err)
// 			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
// 			return
// 		}
// 		txCommitErr := tx.Commit()
// 		if txCommitErr != nil {
// 			a.logger.WithFields(logrus.Fields{
// 				"error-type": "error committing transaction",
// 			}).Errorln(txCommitErr)
// 			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
// 			return
// 		}
//
// 		a.writeJSON(w, http.StatusOK, envelope{"status": fmt.Sprintf("Password for %s updated successfully", admin.R.User.Email)}, nil)
// 	}
// }
