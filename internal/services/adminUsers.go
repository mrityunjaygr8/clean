package services

import (
	"context"

	dbmodels "github.com/mrityunjaygr8/clean/db/models"
	"github.com/mrityunjaygr8/clean/internal/repository"
	"github.com/teris-io/shortid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type AdminUserServiceInterface interface {
	CreateAdminUser(email, password string, admin bool) (AdminUser, error)
	List() ([]AdminUser, error)
}

type AdminUserService struct {
	ctx  context.Context
	DB   boil.ContextExecutor
	repo repository.AdminUsersRepoInterface
}

func NewAdminUserService(ctx context.Context, db boil.ContextExecutor) AdminUserServiceInterface {
	return &AdminUserService{
		DB:   db,
		ctx:  ctx,
		repo: repository.NewAdminUsersRepo(ctx, db),
	}
}

type AdminUser struct {
	UserID     int
	ID         string
	InternalID int
	Admin      bool
	Email      string
}

func toAdminUser(dbu *dbmodels.AdminUser) AdminUser {
	a := AdminUser{
		ID:         dbu.ID,
		InternalID: dbu.InternalID,
		Admin:      dbu.Admin,
	}

	if dbu.R != nil {
		a.Email = dbu.R.User.Email
	}

	return a
}

func toAdminUsers(dbs dbmodels.AdminUserSlice) []AdminUser {
	list := make([]AdminUser, 0)

	for _, u := range dbs {
		list = append(list, toAdminUser(u))
	}

	return list
}

func toDBAdminUser(au AdminUser) dbmodels.AdminUser {
	return dbmodels.AdminUser{
		ID:         au.ID,
		InternalID: au.InternalID,
		Admin:      au.Admin,
		UserID:     au.UserID,
	}
}

func (a *AdminUserService) CreateAdminUser(email, password string, admin bool) (AdminUser, error) {
	abstractService := NewAbstractUserService(a.ctx, a.DB)
	abs, err := abstractService.CreateAbstractUser(email, password)

	if err != nil {
		return AdminUser{}, err
	}

	var dbUser dbmodels.AdminUser
	dbUser.UserID = abs.InternalID
	dbUser.Admin = admin

	tmpId, err := shortid.Generate()
	if err != nil {
		return AdminUser{}, err
	}

	dbUser.ID = tmpId

	err = a.repo.Insert(&dbUser)
	if err != nil {
		return AdminUser{}, err
	}

	u := toAdminUser(&dbUser)
	u.Email = abs.Email

	return u, nil
}

func (a *AdminUserService) List() ([]AdminUser, error) {
	user, err := a.repo.List()

	if err != nil {
		return nil, err
	}

	return toAdminUsers(user), nil
}
