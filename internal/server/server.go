// ! Server struct + constructor
package server

import (
	config "github.com/TusharChauhan09/flashcache/internal/config"
	"github.com/TusharChauhan09/flashcache/internal/store"
)

type Server struct{
	config *config.Config
	store *store.Store
}

func New(cfg *config.Config, st *store.Store) *Server{
	return &Server{
		config: cfg,
		store: st,
	};
}