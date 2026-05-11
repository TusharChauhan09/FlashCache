// ! Server struct + constructor
package server

import (
	config "github.com/TusharChauhan09/flashcache/internal/config"
	"github.com/TusharChauhan09/flashcache/internal/metrics"
	"github.com/TusharChauhan09/flashcache/internal/store"
)

type Server struct{
	config *config.Config
	store *store.Store
	metrics *metrics.Metrics
}

func New(
	cfg *config.Config, 
	st *store.Store,
	mt *metrics.Metrics,
	) *Server{
	return &Server{
		config: cfg,
		store: st,
		metrics: mt,
	};
}