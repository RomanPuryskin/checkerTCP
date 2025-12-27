package handlers

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/console_TCP/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestHander_CheckTCP(t *testing.T) {

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(handler)
	TCPCheckService := service.NewCheckerTCPService()
	TCPCheckHandler := NewCheckerTCPHandler(TCPCheckService, logger)

	tests := []struct {
		Name         string
		Request      string
		ExpectedCode int
		ExpectedBody string
	}{
		{
			Name: "Error_not_enough_fields",
			Request: `{
			"ip": "0.0.0.0",
			"port": "8080"
			}`,
			ExpectedCode: 400,
			ExpectedBody: `{
			"message": "bad request",
			"code": 400
			}`,
		},
		{
			Name: "Error_not_invalid_timeout",
			Request: `{
			"ip": "0.0.0.0",
			"port": "8080",
			"timeout": "gfgr56"
			}`,
			ExpectedCode: 400,
			ExpectedBody: `{
			"message": "invalid timeout",
			"code": 400
			}`,
		},
		{
			Name: "Success_check_closed_TCP",
			Request: `{
			"ip": "127.0.0.1",
			"port": "3000",
			"timeout": "5s"
			}`,
			ExpectedCode: 200,
			ExpectedBody: `{
			"ip": "127.0.0.1",
			"port": "3000",
			"status": "closed",
			"error": "dial tcp 127.0.0.1:3000: connectex: No connection could be made because the target machine actively refused it."
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {

			req := httptest.NewRequest(http.MethodPost, "/check", strings.NewReader(tt.Request))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()

			TCPCheckHandler.CheckTCP(w, req)

			assert.Equal(t, tt.ExpectedCode, w.Code)

			body := w.Body.String()

			assert.JSONEq(t, tt.ExpectedBody, body)
		})
	}
}
