package app

import (
	"github.com/v-351/url-shortener/internal/postgres"
	"github.com/v-351/url-shortener/internal/server"
	"github.com/v-351/url-shortener/internal/service"

	"log"
)

func Run(postgresFlag *bool) {

	var storage service.Storage
	if *postgresFlag {
		storage = postgres.New()
		log.Println("Postgres as storage: flag ==", *postgresFlag)
	} else {
		storage = service.NewMemoryStorage(50)
		log.Println("InMemory as storage: flag ==", *postgresFlag)
	}
	service := &service.Service{Storage: storage}
	server := &server.Server{Service: service}

	server.Run()
	storage.Close()
}
