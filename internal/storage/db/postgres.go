package db

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v4"

	// We need to include this for compatibility with database/sql.
	"github.com/cloudfoundry/go-cf-api/internal/config"
	"github.com/cloudfoundry/go-cf-api/internal/helpers"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.uber.org/zap"
)

func NewPostgresConnection(dbConfig config.DBConfig, connectDB bool) (*sql.DB, Info, error) {
	if dbConfig.Type != "postgres" {
		return nil, Info{}, fmt.Errorf("cannot use database of type %s with postgres binary", dbConfig.Type)
	}
	// Parse Infos from Connection String
	cfg, err := pgx.ParseConfig(dbConfig.ConnectionString)
	if err != nil {
		return nil, Info{}, err
	}
	info := Info{
		Host:         cfg.Host,
		Port:         fmt.Sprint(cfg.Port),
		DatabaseName: cfg.Database,
		User:         cfg.User,
		Password:     cfg.Password,
		Type:         dbConfig.Type,
	}

	// Dont Connect to DB
	if !connectDB {
		dbConfig.ConnectionString = strings.ReplaceAll(dbConfig.ConnectionString, fmt.Sprintf("dbname=%s", info.DatabaseName), "")
	}

	zap.L().Debug(fmt.Sprintf("Using DB Connection String: %v",
		helpers.Redact(dbConfig.ConnectionString, []string{info.Password})))
	database, err := sql.Open("pgx", dbConfig.ConnectionString)
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
