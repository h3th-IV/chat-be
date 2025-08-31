package config

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	database "github.com/h3th-IV/chat-be/internal/database/mysql"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

type Config struct {
	//System vars
	ListenAddr string

	//Logger vars
	Environment        string // --development / --production
	LogLevel           string
	InfoLogOutputPath  string
	ErrorLogOutputPath string

	//DB vars
	MySQLDatabaseHost     string
	MySQLDatabasePort     string
	MySQLDatabaseUser     string
	MySQLDatabasePassword string
	MySQLDatabaseName     string
}

func (configure *Config) Run(c *cli.Context) error {
	var (
		mysqlDBInstance     *sql.DB
		err                 error
		mysqlDatabaseClient database.Database
		loggerConfig        zap.Config
		logger              *zap.Logger
	)

	//In Production, write Logs to file
	if configure.Environment == "production" {
		loggerConfig = zap.NewProductionConfig()
		loggerConfig.OutputPaths = []string{configure.InfoLogOutputPath}
		loggerConfig.ErrorOutputPaths = []string{configure.ErrorLogOutputPath}
	} else {
		loggerConfig = zap.NewDevelopmentConfig()
	}

	if logger, err = loggerConfig.Build(); err != nil {
		return nil
	}
	logger.Sync()

	databaseConfig := &mysql.Config{
		User:                 configure.MySQLDatabaseUser,
		Passwd:               configure.MySQLDatabasePassword,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", configure.MySQLDatabaseHost, configure.MySQLDatabasePort),
		DBName:               configure.MySQLDatabaseName,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	//connect to the database
	if mysqlDBInstance, err = sql.Open("msql", databaseConfig.FormatDSN()); err != nil {
		return fmt.Errorf("unable to connect to the database :(")
	}

	//create the database client
	if mysqlDatabaseClient, err = database.NewChatDB(mysqlDBInstance); err != nil {
		return fmt.Errorf("unable to create database client :(")
	}
	logger.Info("Chat database connected successfully :)")

	mysqlDatabaseClient.Close()

	server := Server{
		HTTPListenAddr: configure.ListenAddr,
	}
	server.StartServer()
	logger.Info("Chat server running @", zap.String("port", configure.ListenAddr))
	return nil
}
