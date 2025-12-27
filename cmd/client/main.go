package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/console_TCP/internal/cli"
	"github.com/console_TCP/internal/cli/commands"
	"github.com/console_TCP/internal/config"
	"github.com/console_TCP/pkg/utils"
)

func main() {
	cliConfig, err := config.LoadCLICofig()
	if err != nil {
		log.Println(err)
		return
	}
	servConfig, err := config.LoadSeverCofig()
	if err != nil {
		log.Panicln(err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println(commands.Help(cliConfig))
		os.Exit(1)
	}

	switch os.Args[1] {
	case cli.COMMAND_CHECK_TCP_CONN:

		if len(os.Args) == 2 {
			fmt.Println(commands.Help(cliConfig))
			os.Exit(1)
		}

		parametrs, err := commands.ParseFlagsForCommandCheckTCP(cliConfig)
		if err != nil {
			fmt.Println("error:", err)
		}

		if parametrs.IP == "" {
			fmt.Println("ip value is required")
			fmt.Println(commands.Help(cliConfig))
			os.Exit(1)
		}

		if parametrs.Port == -1 {
			fmt.Println("port value is required")
			fmt.Println(commands.Help(cliConfig))
			os.Exit(1)
		}

		if err := utils.ValidatePort(parametrs.Port); err != nil {
			fmt.Println("invalid port value")
			fmt.Println(cli.PortUsage())
			os.Exit(1)
		}

		if err := utils.ValidateTimeout(parametrs.TimeoutForTCP); err != nil {
			fmt.Println("invalid timeout value")
			fmt.Println(cli.TimeoutUsage())
			os.Exit(1)
		}

		if err := utils.ValidateTimeout(parametrs.TimeoutForServer); err != nil {
			fmt.Println("invalid timeout value")
			fmt.Println(cli.TimeoutUsage())
			os.Exit(1)
		}

		respJSON, err := commands.CheckTCP(servConfig, parametrs.IP, strconv.Itoa(parametrs.Port), parametrs.TimeoutForServer, parametrs.TimeoutForTCP)
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
		fmt.Println(respJSON)
		os.Exit(0)
	case cli.COMMAND_HELP:
		fmt.Println(commands.Help(cliConfig))
	default:
		fmt.Println("unknown command")
		fmt.Println(commands.Help(cliConfig))
		os.Exit(1)
	}
}
