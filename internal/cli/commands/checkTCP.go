package commands

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/console_TCP/internal/config"
	"github.com/console_TCP/internal/models"
)

var (
	errUnavaliableServer = errors.New("can not connect to server")
	errInternal          = errors.New("internal client error")
	errParseTimeout      = errors.New("invalid timeout")
)

func CheckTCP(servCfg *config.ServerConfig, ip string, port string, timeoutServer, timeoutTCP string) (string, error) {

	timeOutDurationServer, err := time.ParseDuration(timeoutServer)
	if err != nil {
		return "", errParseTimeout
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeOutDurationServer)
	defer cancel()

	respChan := make(chan string)
	errChan := make(chan error)

	go func() {
		req := models.Request{
			IP:      ip,
			Port:    port,
			Timeout: timeoutTCP,
		}

		reqStr, err := json.Marshal(req)
		if err != nil {
			errChan <- err
			return
		}

		serAddr := fmt.Sprintf("http://%s:%s/check", servCfg.ServerAddress, servCfg.ServerPort)
		resp, err := http.Post(serAddr, "application/json", strings.NewReader(string(reqStr)))
		if err != nil {
			errChan <- errUnavaliableServer
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			errChan <- err
			return
		}

		respChan <- string(body)
	}()

	select {
	case <-ctx.Done():
		return "", errUnavaliableServer
	case err := <-errChan:
		if errors.Is(err, errUnavaliableServer) {
			return "", errUnavaliableServer
		} else {
			return "", errInternal
		}
	case resp := <-respChan:
		return resp, nil
	}
}

type FlagsForCheckTCP struct {
	IP               string
	Port             int
	TimeoutForTCP    string
	TimeoutForServer string
}

func ParseFlagsForCommandCheckTCP(cfg *config.CLIConfig) (*FlagsForCheckTCP, error) {
	checkCmd := flag.NewFlagSet("check", flag.ExitOnError)

	IP := checkCmd.String("ip", "", "ip address to check TCP conn")
	port := checkCmd.Int("p", -1, "port to check TCP conn")
	timeoutTCP := checkCmd.String("t", cfg.TCPConnectTimeout, "timeout trying TCP conn")
	timeoutServer := checkCmd.String("ts", cfg.ServerConnectTimeout, "timeout trying connect to server")

	if err := checkCmd.Parse(os.Args[2:]); err != nil {
		return nil, err
	}

	return &FlagsForCheckTCP{
		IP:               *IP,
		Port:             *port,
		TimeoutForTCP:    *timeoutTCP,
		TimeoutForServer: *timeoutServer,
	}, nil
}
