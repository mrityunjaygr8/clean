package entities

import (
	"context"
	"database/sql"
	"fmt"

	dbmodels "github.com/mrityunjaygr8/clean/db/models"
	"github.com/teris-io/shortid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/crypto/bcrypt"
)

type AbstractUserInterface interface {
	NewAbstractUser(email, password string) (*AbstractUser, error)
	ListAbstractUsers() (*[]AbstractUser, error)
}

type AbstractUserModel struct {
	DB *sql.DB
}

type AbstractUser struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"-"`
	InternalID int    `json:"-"`
}

type ErrAbstractUserExists struct {
	User string `json:"user"`
}

func (e ErrAbstractUserExists) Error() string {
	return fmt.Sprintf("Abstract User Already Exists: %s", e.User)
}

func (a AbstractUserModel) NewAbstractUser(email, password string) (*AbstractUser, error) {
	exists, err := dbmodels.AbstractUsers(dbmodels.AbstractUserWhere.Email.EQ(email)).Exists(context.Background(), a.DB)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, ErrAbstractUserExists{User: email}
	}

	u := &AbstractUser{
		Email: email,
	}

	err = u.SetPassword(password)
	if err != nil {
		return nil, err
	}
	err = u.SetId()
	if err != nil {
		return nil, err
	}

	var dbUser dbmodels.AbstractUser
	dbUser.Password = u.Password
	dbUser.Email = u.Email
	dbUser.ID = u.ID

	err = dbUser.Insert(context.Background(), a.DB, boil.Infer())
	if err != nil {
		return nil, err
	}

	u.InternalID = dbUser.InternalID

	return u, nil
}

func (a *AbstractUser) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}

	a.Password = string(hash)

	return nil
}

func (a AbstractUser) CompareHashAndPassword(rawPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(rawPassword))
	if err != nil {
		return err
	}
	return nil

}

func (a *AbstractUser) SetId() error {
	sid, err := shortid.Generate()
	if err != nil {
		return err
	}

	a.ID = sid

	return nil
}

func (a AbstractUserModel) ListAbstractUsers() (*[]AbstractUser, error) {
	var aus []AbstractUser

	err := dbmodels.AbstractUsers().Bind(context.Background(), a.DB, &aus)
	if err != nil {
		return &[]AbstractUser{}, err
	}

	return &aus, nil
}

func (a AbstractUserModel) Exists(email string) (bool, error) {
	exists, err := dbmodels.AbstractUsers(dbmodels.AbstractUserWhere.Email.EQ(email)).Exists(context.Background(), a.DB)
	if err != nil {
		return false, err
	}

	return exists, nil
}
