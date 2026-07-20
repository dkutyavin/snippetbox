package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Addr      string
	StaticDir string
	DSN       string
}

const defaultAddr = ":4000"
const defaultStaticDir = "./ui/static"

func NewConfig() (Config, error) {
	var cfg Config

	flag.StringVar(&cfg.Addr, "addr", getAddr(), "HTTP network address")
	flag.StringVar(&cfg.StaticDir, "staticDir", getStaticDir(), "Path to ui assets directory")
	flag.StringVar(&cfg.DSN, "dsn", getDsn(), "MySQL data source name")

	flag.Parse()

	if cfg.DSN == "" {
		return Config{}, errors.New("could not form valid DSN - check your flags or make sure you set correct DB_USER, DB_PASS and DB_NAME")
	}

	return cfg, nil
}

func getAddr() string {
	fromEnv := os.Getenv("ADDR")

	if fromEnv == "" {
		return defaultAddr
	}

	return fromEnv
}

func getStaticDir() string {
	fromEnv := os.Getenv("STATIC_DIR")

	if fromEnv == "" {
		return defaultStaticDir
	}

	return fromEnv
}

func getDsn() string {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	db := os.Getenv("DB_NAME")

	if user == "" || pass == "" || db == "" {
		return ""
	}

	return fmt.Sprintf("%s:%s@/%s?parseTime=true", user, pass, db)
}
