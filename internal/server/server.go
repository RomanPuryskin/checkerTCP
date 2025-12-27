package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/console_TCP/internal/config"
)

type HTTPServer struct {
	Server *http.Server
	Logger *slog.Logger
}

func NewHTTPServer(cfg *config.ServerConfig, router *http.ServeMux, logger *slog.Logger) *HTTPServer {
	serv := &http.Server{
		Addr:        fmt.Sprintf("%s:%s", cfg.ServerAddress, cfg.ServerPort),
		Handler:     router,
		ReadTimeout: cfg.ReadTimeout,
		IdleTimeout: cfg.IdleTimeout,
	}

	return &HTTPServer{
		Server: serv,
		Logger: logger,
	}
}

func (s *HTTPServer) Start() {
	s.Logger.Info("Server started", "addres:", s.Server.Addr)

	go func() {
		err := s.Server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			s.Logger.Error("Failed to start server", "error", err)
			os.Exit(1)
		}
	}()

}

func (s *HTTPServer) Stop(ctx context.Context) error {
	s.Logger.Info("[!] Shutting down...")

	if err := s.Server.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
