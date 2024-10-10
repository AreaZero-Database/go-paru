package server

import (
	"context"
	"errors"
	"github.com/AreaZero-Database/go-paru/server"
	"github.com/kataras/iris/v12"
	"net/http"
	"time"
)

type TypeServer = server.Server

const DefaultShutdownTimeout = time.Minute

var _ TypeServer = (*Server)(nil)

type Server struct {
	ShutdownTimeout time.Duration
	srv             *http.Server
}

type Options func(*Server)

// NewServer creates a new server instance with default settings
func NewServer(e *iris.Application, addr string, options ...Options) *Server {
	ser := Server{
		ShutdownTimeout: DefaultShutdownTimeout,
		srv: &http.Server{
			Addr:    addr,
			Handler: e,
		},
	}

	for _, option := range options {
		option(&ser)
	}

	return &ser
}

// WithShutdownTimeout duration of graceful shutdown
func WithShutdownTimeout(duration time.Duration) Options {
	return func(server *Server) {
		server.ShutdownTimeout = duration
	}
}

// Start to start the server and wait for it to listen on the given address
func (s *Server) Start() (err error) {
	err = s.srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Shutdown shuts down the server and close with graceful shutdown duration
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.ShutdownTimeout)
	defer cancel()
	return s.srv.Shutdown(ctx)
}
