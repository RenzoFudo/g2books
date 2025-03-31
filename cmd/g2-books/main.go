package main

import (
	"context"
	"github.com/RenzoFudo/g2books/cmd/g2-books/internal/config"
	"github.com/RenzoFudo/g2books/cmd/g2-books/internal/server"
	"github.com/RenzoFudo/g2books/cmd/g2-books/internal/storage"
	"log"
)

func main() {
	cfg := config.ReadConfig()
	log.Println(cfg)
	var stor server.Storage
	stor, err := storage.NewRepo(context.Background(), cfg.DbDsn)
	if err != nil {
		log.Fatal(err.Error())
		stor = storage.New()
	}
	if err = storage.Migrations(cfg.DbDsn, cfg.MigratePath); err != nil {
		log.Fatal(err.Error())
	}

	server := server.New(cfg.Host, stor)
	if err := server.Run(); err != nil {
		panic(err)

	}
}
