package commands

import (
	"fmt"

	"github.com/console_TCP/internal/cli"
	"github.com/console_TCP/internal/config"
)

func Help(cfg *config.CLIConfig) string {
	str := fmt.Sprintf(`checkTCP - network resource checking utility
COMMANDS:
  checkTCP help                    - show this reference
  checkTCP check --ip=IP --p=PORT  - check network resource

COMMAND CHECK:
  Checks the availability of a network resource at the specified IP and port

  REQUIRED FLAGS:
    --ip=ADDRESS    IP address
    --p=PORT        port

  OPTIONAL FLAGS:
    --t=TIMEOUT     TCP connection establishment timeout (current: %s)
    --ts=TIMEOUT    server connection timeout (current: %s)

	%s`, cfg.TCPConnectTimeout, cfg.ServerConnectTimeout, cli.TimeoutUsage())

	return str
}
