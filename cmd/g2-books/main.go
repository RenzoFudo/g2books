package main

import (
	"github.com/RenzoFudo/g2books/cmd/g2-books/internal/config"
	"github.com/RenzoFudo/g2books/cmd/g2-books/internal/server"
	"github.com/RenzoFudo/g2books/cmd/g2-books/internal/storage"
	"log"
)

func main() {
	cfg := config.ReadConfig()
	log.Println(cfg)
	storage := storage.New()
	server := server.New(cfg.Host, storage)
	if err := server.Run(); err != nil {
		panic(err)

	}
}
