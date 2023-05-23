package entities

import "database/sql"

type Entities struct {
	AbstractUser AbstractUserInterface
}

func New(db *sql.DB) Entities {
	return Entities{
		AbstractUser: AbstractUserModel{DB: db},
	}
}
