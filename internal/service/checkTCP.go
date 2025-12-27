package service

import (
	"context"

	"github.com/console_TCP/pkg/tcp"
)

type CheckerTCPService interface {
	CheckTCPConnection(ctx context.Context, address, port string) (string, error)
}

type checkerTCPService struct {
}

func NewCheckerTCPService() *checkerTCPService {
	return &checkerTCPService{}
}

func (cs *checkerTCPService) CheckTCPConnection(ctx context.Context, address, port string) (string, error) {
	err := tcp.CheckTCPConnectionWithContext(ctx, address, port)
	if err != nil {
		return "closed", err
	}

	return "open", nil
}
