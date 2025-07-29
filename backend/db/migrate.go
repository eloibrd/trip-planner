package db

import (
	"embed"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func Migrate() (err error) {
	// get sql.DB obj
	gormInstance, err := GetDB()
	if err != nil {
		return
	}
	db, err := gormInstance.DB()
	if err != nil {
		return
	}

	// setup Goose
	goose.SetBaseFS(embedMigrations)

	if err = goose.SetDialect("postgres"); err != nil {
		return
	}

	// Run migrations
	if err = goose.Up(db, "migrations"); err != nil {
		return
	}
	return
}
