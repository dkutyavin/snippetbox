package main

import (
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"

	"snippetbox.dekutyavin.net/internal/backends/database"
	"snippetbox.dekutyavin.net/internal/config"
	"snippetbox.dekutyavin.net/internal/models"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger         *slog.Logger
	config         config.Config
	snippets       *models.SnippetModel
	templates      map[string]*template.Template
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	db, err := database.OpenDB(cfg.DSN)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	templatesCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	sessionManager := scs.New()
	sessionManager.Store = mysqlstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour

	app := &application{
		logger:         logger,
		config:         cfg,
		snippets:       &models.SnippetModel{DB: db},
		templates:      templatesCache,
		formDecoder:    form.NewDecoder(),
		sessionManager: sessionManager,
	}

	logger.Info("starting server", slog.String("addr", cfg.Addr))

	err = http.ListenAndServe(cfg.Addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
