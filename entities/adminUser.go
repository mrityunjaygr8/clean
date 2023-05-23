package entities

import "database/sql"

type AdminUserInterface interface {
	NewAdminUser(email, password string) (*AdminUser, error)
	ListAdminUsers() (*[]AdminUser, error)
}

type AdminUserModel struct {
	DB *sql.DB
}

type AdminUser struct {
	ID         string `json:"id"`
	InternalID int    `json:"-"`
	Admin      bool   `json:"admin"`
}
