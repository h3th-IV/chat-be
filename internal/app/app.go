package app

import (
	"github.com/h3th-IV/chat-be/internal/config"
	"github.com/urfave/cli/v2"
)

func StartCommand() *cli.Command {
	var config = &config.Config{}

	cmd := &cli.Command{
		Name:  "start",
		Usage: "start the server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "listen-addr",
				EnvVars:     []string{"CHAT_SERVER_LISTEN_ADDR"},
				Usage:       "listening address for chat server",
				Destination: &config.ListenAddr,
				Value:       ":8080", //incase env vars not set
			},
			&cli.StringFlag{
				Name:        "mysql-database-host",
				EnvVars:     []string{"CHAT_DATABASE_HOST"},
				Usage:       "mysql database host",
				Destination: &config.MySQLDatabaseHost,
				Value:       "localhost",
			},
			&cli.StringFlag{
				Name:        "mysql-database-port",
				EnvVars:     []string{"CHAT_DATABASE_PORT"},
				Usage:       "mysql database port",
				Destination: &config.MySQLDatabasePort,
				Value:       "3306",
			},
			&cli.StringFlag{
				Name:        "mysql-database-user",
				EnvVars:     []string{"CHAT_DATABASE_USER"},
				Usage:       "mysql database user",
				Destination: &config.MySQLDatabaseUser,
				Value:       "root",
			},
			&cli.StringFlag{
				Name:        "mysql-database-password",
				EnvVars:     []string{"CHAT_DATABASE_PASSWORD"},
				Usage:       "mysql database password",
				Destination: &config.MySQLDatabasePassword,
				Value:       "password",
			},
			&cli.StringFlag{
				Name:        "mysql-database-name",
				EnvVars:     []string{"CHAT_DATABASE_NAME"},
				Usage:       "mysql database name",
				Destination: &config.MySQLDatabaseName,
				Value:       "chat",
			},
			&cli.StringFlag{
				Name:        "environment-name",
				EnvVars:     []string{"CHAT_ENVIRONMENT_NAME"},
				Usage:       "chat server environmemnt name",
				Destination: &config.MySQLDatabaseName,
				Value:       "",
			},
			&cli.StringFlag{
				Name:        "log-level",
				EnvVars:     []string{"CHAT_LOG_LEVEL"},
				Usage:       "chat server log level",
				Destination: &config.LogLevel,
				Value:       "info",
			},
			&cli.StringFlag{
				Name:        "info-log-output-path",
				EnvVars:     []string{"CHAT_LOG_OUTPUT_PATH"},
				Usage:       "chat server info log output path",
				Destination: &config.InfoLogOutputPath,
				Value:       "logs/chat.log",
			},
			&cli.StringFlag{
				Name:        "error-log-output-path",
				EnvVars:     []string{"CHAT_ERROR_LOG_OUTPUT_PATH"},
				Usage:       "chat server error log output path",
				Destination: &config.ErrorLogOutputPath,
				Value:       "logs/chat.log",
			},
		},
		Action: config.Run,
	}
	return cmd
}
