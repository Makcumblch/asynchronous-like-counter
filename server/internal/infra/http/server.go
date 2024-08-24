package http

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Makcumblch/asynchronous-like-counter/internal/infra/http/middleware"
	"github.com/Makcumblch/asynchronous-like-counter/internal/util/config"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run() {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			log.Fatal("error occured while running http server", err)
		}
	}()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func NewServer(config config.HttpConfig, handler http.Handler) Server {
	cors := middleware.CorsMW(handler)
	server := Server{
		httpServer: &http.Server{
			IdleTimeout:    30 * time.Minute,
			Addr:           ":" + config.Port,
			Handler:        cors,
			MaxHeaderBytes: 1 << 20, // 1 MB
			ReadTimeout:    30 * time.Minute,
			WriteTimeout:   30 * time.Minute,
		},
	}

	return server
}
