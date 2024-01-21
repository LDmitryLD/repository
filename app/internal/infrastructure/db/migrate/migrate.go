package migrate

import (
	"projects/LDmitryLD/repository/app/config"

	"github.com/jmoiron/sqlx"
)

type Migrator struct {
	db     *sqlx.DB
	dbConf config.DB
}

func NewMigrator(db *sqlx.DB, dbConf config.DB) *Migrator {
	return &Migrator{db: db, dbConf: dbConf}
}
