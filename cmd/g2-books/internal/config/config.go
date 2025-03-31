package config

import (
	"flag"
	"log"
	"os"
)

type Config struct {
	Host        string
	DbDsn       string
	MigratePath string
	Debug       bool
}

const (
	defaultDbDSN       = "postgres://postgres:6406655@localhost:5432/cours-db"
	defaultMigratePath = "migrations"
	defaultHost        = ":8080"
)

func ReadConfig() Config {
	var host string
	var dbDsn string
	var migratePath string
	flag.StringVar(&host, "host", defaultHost, "server host")
	flag.StringVar(&dbDsn, "db", defaultDbDSN, "data base adress")
	flag.StringVar(&migratePath, "m", defaultMigratePath, "path to migrations")
	debug := flag.Bool("debug", false, "enable debug logging level")
	flag.Parse()

	hostEnv := os.Getenv("SERVER_HOST")
	dbDsnEnv := os.Getenv("DB_DSN")
	migratePathEnv := os.Getenv("MIGRATE_PATH")

	log.Println(hostEnv)
	if hostEnv != "" && host == defaultHost {
		host = hostEnv
	}
	if dbDsnEnv != "" && dbDsn == defaultDbDSN {
		dbDsn = dbDsnEnv
	}
	if migratePathEnv != "" && migratePath == defaultHost {
		migratePath = migratePathEnv
	}

	return Config{
		Host:        host,
		DbDsn:       dbDsn,
		MigratePath: migratePathEnv,
		Debug:       *debug,
	}
}
