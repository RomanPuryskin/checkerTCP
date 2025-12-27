package tcp

import (
	"context"
	"fmt"
	"net"
)

func CheckTCPConnectionWithContext(ctx context.Context, address, port string) error {

	addr := fmt.Sprintf("%s:%s", address, port)
	dialer := net.Dialer{}
	conn, err := dialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		return err
	}

	defer conn.Close()
	return nil
}
