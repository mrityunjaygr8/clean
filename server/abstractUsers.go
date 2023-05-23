package server

import (
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

func (a *Application) ListAbstractUsers(w http.ResponseWriter, r *http.Request) {
	u, err := a.entities.AbstractUser.ListAbstractUsers()
	if err != nil {
		a.logger.Errorln(err)
		a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err}, nil)
		return
	}

	a.writeJSON(w, http.StatusOK, envelope{"resp": u}, nil)
}

func (a *Application) AddAbstractUser(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Password string `json:"password" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
	}

	type response struct {
		Email string `json:"email"`
		ID    string `json:"id"`
	}
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
			a.logger.Println(e.Field(), e.Value())
		}

		a.logger.Errorln(errResp)
		a.writeJSON(w, http.StatusInternalServerError, envelope{"errors": errResp}, nil)

		return
	}

	u, err := a.entities.AbstractUser.NewAbstractUser(req.Email, req.Password)
	if err != nil {
		a.logger.Errorln(err)
		a.writeJSON(w, http.StatusInternalServerError, envelope{"error": err.Error()}, nil)
		return
	}

	resp := response{
		Email: u.Email,
		ID:    u.ID,
	}

	a.writeJSON(w, http.StatusOK, envelope{"resp": resp}, nil)
}
