package repository

import (
	"context"

	dbmodels "github.com/mrityunjaygr8/clean/db/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AdminUsersRepoInterface interface {
	Insert(au *dbmodels.AdminUser) error
	ExistsByEmail(email string) (bool, error)
	ExistsById(id string) (bool, error)
	List() (dbmodels.AdminUserSlice, error)
	Delete(au dbmodels.AdminUser) (int64, error)
	Update(au dbmodels.AdminUser) (int64, error)
}

type AdminUsersRepo struct {
	DB  boil.ContextExecutor
	ctx context.Context
}

func NewAdminUsersRepo(ctx context.Context, db boil.ContextExecutor) AdminUsersRepoInterface {
	return &AdminUsersRepo{
		DB:  db,
		ctx: ctx,
	}
}

func (a *AdminUsersRepo) Insert(au *dbmodels.AdminUser) error {
	return au.Insert(a.ctx, a.DB, boil.Infer())
}
func (a *AdminUsersRepo) ExistsByEmail(email string) (bool, error) {
	return dbmodels.AdminUsers(qm.InnerJoin(dbmodels.TableNames.AbstractUsers+" d on "+dbmodels.AdminUserColumns.UserID+"="+dbmodels.AbstractUserColumns.InternalID), dbmodels.AbstractUserWhere.Email.EQ(email)).Exists(a.ctx, a.DB)
}

func (a *AdminUsersRepo) ExistsById(id string) (bool, error) {
	return dbmodels.AdminUserExists(a.ctx, a.DB, id)
}

func (a *AdminUsersRepo) List() (dbmodels.AdminUserSlice, error) {
	return dbmodels.AdminUsers(qm.Load(dbmodels.AdminUserRels.User)).All(a.ctx, a.DB)
}

func (a *AdminUsersRepo) Delete(au dbmodels.AdminUser) (int64, error) {
	return au.Delete(a.ctx, a.DB)
}

func (a *AdminUsersRepo) Update(au dbmodels.AdminUser) (int64, error) {
	return au.Update(a.ctx, a.DB, boil.Infer())
}
