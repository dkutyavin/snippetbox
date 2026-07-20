package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
	"snippetbox.dekutyavin.net/internal/backends/database"
	"snippetbox.dekutyavin.net/internal/config"
	"snippetbox.dekutyavin.net/internal/migrations"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := database.OpenDB(cfg.DSN)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func() {
		err := db.Close()

		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	goose.SetBaseFS(migrations.FS)
	err = goose.SetDialect("mysql")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = goose.Up(db, ".")
	if err != nil {
		log.Fatal(err.Error())
	}
}
