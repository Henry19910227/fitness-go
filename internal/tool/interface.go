package tool

import "database/sql"

type Mysql interface {
	DB() *sql.DB
}

type Migrate interface {
	Up(step *int) error
	Down(step *int) error
	Force(version int) error
	Migrate(version uint) error
	Version() (uint, bool, error)
}