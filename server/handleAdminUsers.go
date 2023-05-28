package server

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	dbmodels "github.com/mrityunjaygr8/clean/db/models"
	"github.com/mrityunjaygr8/clean/internal/services"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

func (a *Server) handleAdminUserCreate() http.HandlerFunc {
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

		service := services.NewAdminUserService(r.Context(), tx)

		user, err := service.CreateAdminUser(req.Email, req.Password, req.Admin)
		if err != nil {
			a.l.Error().Err(err).Msg("User Exists")
			errMsg := make(map[string]interface{})
			errMsg["Email"] = err.Error()
			a.writeJSON(w, http.StatusBadRequest, envelope{"errors": errMsg}, nil)
			return
		}

		if ok := a.transactionCommit(w, tx); ok {
			return
		}

		resp := response{
			Email: user.Email,
			ID:    user.ID,
			Admin: user.Admin,
		}

		a.writeJSON(w, http.StatusCreated, envelope{"admin-user": resp}, nil)
	}
}

func (a *Server) handleAdminUserList() http.HandlerFunc {
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

		service := services.NewAdminUserService(r.Context(), a.db)
		adminUsers, err := service.List()
		if err != nil {
			a.l.Error().Str("error-type", "error listing admin users").Err(err).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err}, nil)
			return
		}

		for _, adm := range adminUsers {
			tmp := adminUser{}
			tmp.Admin = adm.Admin
			tmp.Email = adm.Email
			tmp.Id = adm.ID

			adms = append(adms, tmp)
		}

		a.writeJSON(w, http.StatusOK, envelope{"admin-users": adms}, nil)
	}
}

func (a *Server) handleAdminUserRetrieve() http.HandlerFunc {
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
			a.l.Error().Str("error-type", "error retrieving admin users").Str("admin-user-id", id).Err(err).Msg("")
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

func (a *Server) handleAdminUserDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "adminUserId")

		adms, err := dbmodels.AdminUsers(dbmodels.AdminUserWhere.ID.EQ(id), qm.Load(qm.Rels(dbmodels.AdminUserRels.User))).All(r.Context(), a.db)
		if err != nil {
			a.l.Error().Str("error-type", "error counting admin users").Str("admin-user-id", id).Err(err).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err}, nil)
			return
		}

		if len(adms) == 0 {
			a.writeJSON(w, http.StatusNotFound, envelope{}, nil)
			return
		}

		if len(adms) > 1 {
			a.l.Error().Str("error-type", "error id collision").Str("admin-user-id", id).Err(errors.New("admin user id collision")).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}

		admin := adms[0]
		abstract := admin.R.GetUser()
		abstractId := admin.R.User.ID

		tx, err := a.db.BeginTx(r.Context(), &sql.TxOptions{})
		if a.checkTransactionBeginError(w, err) {
			return
		}
		_, err = admin.Delete(r.Context(), tx)
		if err != nil {
			if a.transactionRollback(w, tx) {
				return
			}
			a.l.Error().Str("error-type", "error deleting admin users").Str("admin-user-id", id).Err(err).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err}, nil)
			return
		}

		_, err = abstract.Delete(r.Context(), tx)
		if err != nil {
			if a.transactionRollback(w, tx) {
				return
			}
			a.l.Error().Str("error-type", "error deleting abstract users").Str("abstract-user-id", abstractId).Err(err).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err}, nil)
			return
		}

		if ok := a.transactionCommit(w, tx); ok {
			return
		}

		a.writeJSON(w, http.StatusNoContent, envelope{}, nil)
	}
}

func (a *Server) handleAdminUserUpdate() http.HandlerFunc {
	type request struct {
		Admin bool `json:"admin" validate:"boolean"`
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
			a.l.Error().Err(err).Msg("")
			a.writeJSON(w, http.StatusBadRequest, envelope{"error": err.Error()}, nil)
			return
		}

		if a.checkExtraFieldsError(w, err) {
			return
		}
		err = a.validator.Struct(req)
		a.l.Info().Err(err).Msg("")
		a.l.Info().Str("req", fmt.Sprintf("%v", req)).Msg("")
		if a.checkValidationErrors(w, err) {
			return
		}

		adms, err := dbmodels.AdminUsers(dbmodels.AdminUserWhere.ID.EQ(id), qm.Load(qm.Rels(dbmodels.AdminUserRels.User))).All(r.Context(), a.db)
		if err != nil {
			a.l.Error().Str("error-type", "error counting admin users").Str("admin-user-id", id).Err(err).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err}, nil)
			return
		}

		if len(adms) == 0 {
			a.writeJSON(w, http.StatusNotFound, envelope{}, nil)
			return
		}

		if len(adms) > 1 {
			a.l.Error().Str("error-type", "error id collision").Str("admin-user-id", id).Err(errors.New("admin user id collision")).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}

		admin := adms[0]
		admin.Admin = req.Admin

		tx, err := a.db.BeginTx(r.Context(), &sql.TxOptions{})
		if a.checkTransactionBeginError(w, err) {
			return
		}
		_, err = admin.Update(r.Context(), tx, boil.Infer())
		if err != nil {
			if a.transactionRollback(w, tx) {
				return
			}
			a.l.Error().Str("error-type", "error updating admin user").Err(err).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err}, nil)
			return
		}
		if ok := a.transactionCommit(w, tx); ok {
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

func (a *Server) handleAdminUserUpdatePassword() http.HandlerFunc {
	type request struct {
		Password string `json:"password" validate:"required"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "adminUserId")

		var req request
		err := a.readJSON(w, r, &req)

		if a.checkExtraFieldsError(w, err) {
			return
		}
		err = a.validator.Struct(req)
		if a.checkValidationErrors(w, err) {
			return
		}

		adms, err := dbmodels.AdminUsers(dbmodels.AdminUserWhere.ID.EQ(id), qm.Load(qm.Rels(dbmodels.AdminUserRels.User))).All(r.Context(), a.db)
		if err != nil {
			a.l.Error().Str("error-type", "error counting admin users").Str("admin-user-id", id).Err(err).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err}, nil)
			return
		}

		if len(adms) == 0 {
			a.writeJSON(w, http.StatusNotFound, envelope{}, nil)
			return
		}

		if len(adms) > 1 {
			a.l.Error().Str("error-type", "error id collision").Str("admin-user-id", id).Err(errors.New("admin user id collision")).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}

		admin := adms[0]
		abstract := admin.R.GetUser()

		tx, err := a.db.BeginTx(r.Context(), &sql.TxOptions{})
		if a.checkTransactionBeginError(w, err) {
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
		if err != nil {
			if a.transactionRollback(w, tx) {
				return
			}
			a.l.Error().Str("error-type", "error creating hash from raw password").Err(err).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}

		abstract.Password = string(hash)
		_, err = abstract.Update(r.Context(), tx, boil.Infer())

		if err != nil {
			if a.transactionRollback(w, tx) {
				return
			}
			a.l.Error().Str("error-type", "error creating hash from raw password").Err(err).Msg("")
			a.writeJSON(w, http.StatusInternalServerError, envelope{"error": http.StatusText(http.StatusInternalServerError)}, nil)
			return
		}
		if ok := a.transactionCommit(w, tx); ok {
			return
		}

		a.writeJSON(w, http.StatusOK, envelope{"status": fmt.Sprintf("Password for %s updated successfully", admin.R.User.Email)}, nil)
	}
}
