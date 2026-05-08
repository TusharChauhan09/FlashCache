package main

import (
	"log"

	config "github.com/TusharChauhan09/flashcache/internal/config"
	"github.com/TusharChauhan09/flashcache/internal/server"
	"github.com/TusharChauhan09/flashcache/internal/store"
)

func main(){

	cfg := config.LoadConfig()

	st := store.New()

	srv := server.New(cfg,st)

	log.Fatal(srv.Start())
}