package db

import (
	"database/sql"
	"fmt"
	"net"
	"time"

	"github.com/cloudfoundry/go-cf-api/internal/config"
	"github.com/cloudfoundry/go-cf-api/internal/helpers"
	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func NewMySQLConnection(dbConfig config.DBConfig, connectDB bool) (*sql.DB, Info, error) {
	if dbConfig.Type != "mysql" {
		return nil, Info{}, fmt.Errorf("cannot use database of type %s with mysql binary", dbConfig.Type)
	}
	// Parse Infos from Connection String
	cfg, err := mysql.ParseDSN(dbConfig.ConnectionString)
	if err != nil {
		return nil, Info{}, err
	}
	host, port, err := net.SplitHostPort(cfg.Addr)
	if err != nil {
		return nil, Info{}, err
	}
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
	if err != nil {
		return nil, Info{}, err
	}

	// DB Connection Settings
	database.SetMaxIdleConns(dbConfig.MinConnections)
	database.SetMaxOpenConns(dbConfig.MaxConnections)
	database.SetConnMaxLifetime(time.Hour)

	// Check Connection on StartUp
	err = database.Ping()

	return database, info, err
}
