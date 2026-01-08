package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"mastery-project/internal/config"
	"mastery-project/internal/database"
	"net/http"
	"time"
)

type Server struct {
	Config     *config.Config
	Db         *database.Database
	httpServer *http.Server
}

func NewServer(config *config.Config) (*Server, error) {
	db, err := database.New(config)
	if err != nil {
		return nil, err
	}
	return &Server{
		Config: config,
		Db:     db,
	}, nil
}

func (s *Server) SetupHttpServer(handler http.Handler) {
	s.httpServer = &http.Server{
		Addr:         ":" + s.Config.Server.Port,
		Handler:      handler,
		IdleTimeout:  time.Duration(s.Config.Server.IdleTimeout) * time.Second,
		ReadTimeout:  time.Duration(s.Config.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(s.Config.Server.WriteTimeout) * time.Second,
	}
}

func (s *Server) Run() error {
	if s.httpServer == nil {
		return errors.New("http server not initialized")
	}

	slog.Info("http server started")
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("http server shutdown error: %v", err)
	}

	//close all services here
	err := s.Db.Close()
	if err != nil {
		return fmt.Errorf("database close error: %v", err)
	}
	return nil
}
