package services

import (
	"context"
	"fmt"

	dbmodels "github.com/mrityunjaygr8/clean/db/models"
	"github.com/mrityunjaygr8/clean/internal/repository"
	"github.com/teris-io/shortid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/crypto/bcrypt"
)

type ErrAbstractUserExists struct {
	Email string
}

func (e ErrAbstractUserExists) Error() string {
	return fmt.Sprintf("User with email %s already exists", e.Email)
}

type AbstractUserServiceInterface interface {
	CreateAbstractUser(email, password string) (AbstractUser, error)
	List() ([]AbstractUser, error)
}

type AbstractUserService struct {
	ctx  context.Context
	DB   boil.ContextExecutor
	repo repository.AbstractUserRepoInterface
}

func NewAbstractUserService(ctx context.Context, db boil.ContextExecutor) AbstractUserServiceInterface {
	return &AbstractUserService{
		DB:   db,
		ctx:  ctx,
		repo: repository.NewAbstractUserRepo(ctx, db),
	}
}

type AbstractUser struct {
	Email      string
	ID         string
	InternalID int
	Password   string
}

func toAbstractUser(dbu *dbmodels.AbstractUser) AbstractUser {
	return AbstractUser{
		Email:      dbu.Email,
		ID:         dbu.ID,
		InternalID: dbu.InternalID,
		Password:   dbu.Password,
	}
}

func toAbstractUsers(dbs dbmodels.AbstractUserSlice) []AbstractUser {
	list := make([]AbstractUser, 0)

	for _, u := range dbs {
		list = append(list, toAbstractUser(u))
	}

	return list
}

func toDBAbstractUser(au AbstractUser) dbmodels.AbstractUser {
	return dbmodels.AbstractUser{
		ID:         au.ID,
		InternalID: au.InternalID,
		Email:      au.Email,
		Password:   au.Password,
	}
}

func (a *AbstractUserService) CreateAbstractUser(email, password string) (AbstractUser, error) {
	exists, err := a.repo.ExistsByEmail(email)
	if err != nil {
		return AbstractUser{}, nil
	}

	if exists {
		return AbstractUser{}, ErrAbstractUserExists{Email: email}
	}

	var dbUser dbmodels.AbstractUser
	dbUser.Email = email

	tmpId, err := shortid.Generate()
	if err != nil {
		return AbstractUser{}, err
	}

	dbUser.ID = tmpId

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return AbstractUser{}, err
	}

	dbUser.Password = string(hash)

	err = a.repo.Insert(&dbUser)
	if err != nil {
		return AbstractUser{}, err
	}

	return toAbstractUser(&dbUser), nil
}

func (a *AbstractUserService) List() ([]AbstractUser, error) {
	user, err := a.repo.List()

	if err != nil {
		return nil, err
	}

	return toAbstractUsers(user), nil
}
