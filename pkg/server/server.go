package server

import (
	"context"
	"net/http"
	"time"
)

const (
	defaultAddr            = ":8080"
	defaultReadTimeout     = 5 * time.Second
	defaultWriteTimeout    = 5 * time.Second
	defaultShutdownTimeOut = 7 * time.Second
)

type Server struct {
	Server *http.Server
}

func (s *Server) RunServer(handler http.Handler) error {
	s.Server = &http.Server{
		Addr:           defaultAddr,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    defaultReadTimeout,
		WriteTimeout:   defaultWriteTimeout,
	}

	return s.Server.ListenAndServe()
}

func (s *Server) ShutdownServer() error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultShutdownTimeOut)
	defer cancel()

	return s.Server.Shutdown(ctx)
}
