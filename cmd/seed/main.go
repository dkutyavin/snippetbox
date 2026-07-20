package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"snippetbox.dekutyavin.net/internal/backends/database"
	"snippetbox.dekutyavin.net/internal/config"
	"snippetbox.dekutyavin.net/internal/models"
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

	snippets := models.SnippetModel{DB: db}
	count, err := snippets.Count()
	if err != nil {
		log.Fatal(err.Error())
	}

	if count != 0 {
		log.Print("There are existing snippets. Skipping seed...")
		os.Exit(0)
	}

	_, err = snippets.Insert("Oh snail", "Oh snail, climb Mt. Fuji, but slowly, slowly. Kobayashi Issa", 356)
	if err != nil {
		log.Fatal(err.Error())
	}
}
