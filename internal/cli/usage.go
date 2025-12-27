package cli

import (
	"strings"
)

const (
	COMMAND_CHECK_TCP_CONN = "check"
	COMMAND_HELP           = "help"
)

func PortUsage() string {
	return "Valid port value must be in 0-65535"
}

func TimeoutUsage() string {
	var sb strings.Builder
	sb.WriteString("Examples of valid timeout value:\n")
	sb.WriteString("  30s    - 30 seconds\n")
	sb.WriteString("  5m     - 5 minutes\n")
	sb.WriteString("  1h     - 1 hour\n")
	sb.WriteString("  1h30m  - 1 hour 30 minutes\n")
	sb.WriteString("  500ms  - 500 milliseconds\n")
	sb.WriteString("  1.5h   - 1.5 hours\n")
	return sb.String()
}
