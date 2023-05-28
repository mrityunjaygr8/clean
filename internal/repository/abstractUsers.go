package repository

import (
	"context"

	dbmodels "github.com/mrityunjaygr8/clean/db/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type AbstractUserRepoInterface interface {
	Insert(*dbmodels.AbstractUser) error
	ExistsByEmail(email string) (bool, error)
	ExistsById(id string) (bool, error)
	List() (dbmodels.AbstractUserSlice, error)
	Delete(dbmodels.AbstractUser) (int64, error)
	Update(dbmodels.AbstractUser) (int64, error)
}

func NewAbstractUserRepo(ctx context.Context, dbLike boil.ContextExecutor) AbstractUserRepoInterface {
	return &AbstractUserRepo{
		DB:  dbLike,
		ctx: ctx,
	}
}

type AbstractUserRepo struct {
	DB  boil.ContextExecutor
	ctx context.Context
}

func (a *AbstractUserRepo) Insert(au *dbmodels.AbstractUser) error {
	return au.Insert(a.ctx, a.DB, boil.Infer())
}

func (a *AbstractUserRepo) ExistsByEmail(email string) (bool, error) {
	return dbmodels.AbstractUsers(dbmodels.AbstractUserWhere.Email.EQ(email)).Exists(a.ctx, a.DB)
}

func (a *AbstractUserRepo) ExistsById(id string) (bool, error) {
	return dbmodels.AbstractUserExists(a.ctx, a.DB, id)
}

func (a *AbstractUserRepo) List() (dbmodels.AbstractUserSlice, error) {
	return dbmodels.AbstractUsers().All(a.ctx, a.DB)
}

func (a *AbstractUserRepo) Delete(au dbmodels.AbstractUser) (int64, error) {
	return au.Delete(a.ctx, a.DB)
}

func (a *AbstractUserRepo) Update(au dbmodels.AbstractUser) (int64, error) {
	return au.Update(a.ctx, a.DB, boil.Infer())
}
