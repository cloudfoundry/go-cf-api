package db

import (
	"database/sql"
	"fmt"
	"net"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/config"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/helpers"
	"go.uber.org/zap"
)

func NewMySQLConnection(dbConfig config.DBConfig, connectDB bool) (*sql.DB, Info) {
	if dbConfig.Type != "mysql" {
		helpers.CheckErrFatal(fmt.Errorf("cannot use database of type %s with mysql binary", dbConfig.Type))
	}
	// Parse Infos from Connection String
	cfg, err := mysql.ParseDSN(dbConfig.ConnectionString)
	helpers.CheckErrFatal(err)
	host, port, err := net.SplitHostPort(cfg.Addr)
	helpers.CheckErrFatal(err)
	info := Info{
		Host:         host,
		Port:         port,
		DatabaseName: cfg.DBName,
		User:         cfg.User,
		Password:     cfg.Passwd,
		Type:         dbConfig.Type,
	}

	if !connectDB {
		cfg.DBName = ""
		dbConfig.ConnectionString = cfg.FormatDSN()
	}

	zap.L().Debug(fmt.Sprintf("Using DB Connection String: %v",
		helpers.Redact(dbConfig.ConnectionString, []string{info.Password})))
	database, err := sql.Open("mysql", dbConfig.ConnectionString)
	helpers.CheckErrFatal(err)

	// DB Connection Settings
	database.SetMaxIdleConns(5)  //nolint:gomnd // This will be configurable at some point
	database.SetMaxOpenConns(20) //nolint:gomnd // This will be configurable at some point
	database.SetConnMaxLifetime(time.Hour)

	// Check Connection on StartUp
	helpers.CheckErrFatal(database.Ping())

	return database, info
}
