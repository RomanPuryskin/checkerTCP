package handlers

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/console_TCP/internal/models"
	"github.com/console_TCP/internal/service"
	"github.com/console_TCP/pkg/utils"
)

type CheckTCPHandler struct {
	CheckTCPService service.CheckerTCPService
	Logger          *slog.Logger
}

func NewCheckerTCPHandler(TCPService service.CheckerTCPService, logger *slog.Logger) *CheckTCPHandler {
	return &CheckTCPHandler{
		CheckTCPService: TCPService,
		Logger:          logger,
	}
}

func (ch *CheckTCPHandler) CheckTCP(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		ch.Logger.Error("method not allowed")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(models.ResponeError{
			Message: "method not allowed",
			Code:    http.StatusMethodNotAllowed,
		})
		return
	}

	var req models.Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ch.Logger.Error("failed decode request", "error", err, "request", req)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ResponeError{
			Message: "bad request",
			Code:    http.StatusBadRequest,
		})
		return
	}

	// провалидируем тело запроса на:
	// 1) присутствие всех обязательных полей
	if err := utils.ValidateStruct(req); err != nil {
		ch.Logger.Error("failed parse request", "error", err, "request", req)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ResponeError{
			Message: "bad request",
			Code:    http.StatusBadRequest,
		})
		return
	}
	// 2) корректный таймаут
	timeout, err := time.ParseDuration(req.Timeout)
	if err != nil {
		ch.Logger.Error("failed parse timeout", "error", err, "request", req)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ResponeError{
			Message: "invalid timeout",
			Code:    http.StatusBadRequest,
		})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	status, err := ch.CheckTCPService.CheckTCPConnection(ctx, req.IP, req.Port)

	w.Header().Set("Content-Type", "application/json")
	resp := models.Responce{
		IP:     req.IP,
		Port:   req.Port,
		Status: status,
		Error:  err.Error(),
	}

	json.NewEncoder(w).Encode(resp)
	slog.Info("success check TCP", "input", req, "responce", resp)
}
