package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/console_TCP/internal/config"
	"github.com/console_TCP/internal/server"
	"github.com/console_TCP/internal/server/handlers"
	"github.com/console_TCP/internal/server/routes"
	"github.com/console_TCP/internal/service"
	"github.com/console_TCP/pkg/logger"
)

func main() {
	servConfig, err := config.LoadSeverCofig()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	logConfig, err := config.LoadLoggerCofig()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// инициализация всех зависимостей
	logger := logger.InitLogger(logConfig)
	TCPCheckService := service.NewCheckerTCPService()
	TCPCheckHandler := handlers.NewCheckerTCPHandler(TCPCheckService, logger)
	router := routes.InitRoutes(*TCPCheckHandler)
	HTTPServer := server.NewHTTPServer(servConfig, router, logger)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	// запуск сервера
	HTTPServer.Start()

	<-ctx.Done()
	ctxTimeOut, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = HTTPServer.Stop(ctxTimeOut)
	if err != nil {
		HTTPServer.Logger.Error("Failed shutdown", "error", err)
		os.Exit(1)
	}

	HTTPServer.Logger.Info("Gracefully stopped")
}
