package routes

import (
	"net/http"

	"github.com/console_TCP/internal/server/handlers"
)

func InitRoutes(handler handlers.CheckTCPHandler) *http.ServeMux {
	mux := http.NewServeMux()

	// роут для работы с TCP
	mux.HandleFunc("/check", handler.CheckTCP)

	return mux
}
