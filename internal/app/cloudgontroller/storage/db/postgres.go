package db

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/config"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/helpers"
	"go.uber.org/zap"
)

func NewPostgresConnection(dbConfig config.DBConfig, connectDB bool) {
	if dbConfig.Type != "postgres" {
		helpers.CheckErrFatal(fmt.Errorf("cannot use database of type %s with postgres binary", dbConfig.Type))
	}
	once.Do(func() {
		// Parse Infos from Connection String
		cfg, err := pgx.ParseConfig(dbConfig.ConnectionString)
		helpers.CheckErrFatal(err)
		databaseInfo.Host = cfg.Host
		databaseInfo.Port = fmt.Sprint(cfg.Port)
		databaseInfo.DatabaseName = cfg.Database
		databaseInfo.User = cfg.User
		databaseInfo.Password = cfg.Password
		databaseInfo.Type = dbConfig.Type

		// Dont Connect to DB
		if !connectDB {
			dbConfig.ConnectionString = strings.ReplaceAll(dbConfig.ConnectionString, fmt.Sprintf("dbname=%s", databaseInfo.DatabaseName), "")
		}

		zap.L().Debug(fmt.Sprintf("Using DB Connection String: %v",
			helpers.Redact(dbConfig.ConnectionString, []string{databaseInfo.Password})))
		database, err = sql.Open("pgx", dbConfig.ConnectionString)
		helpers.CheckErrFatal(err)
		// If DB is Missing, we Create it
		err = database.Ping()
		if dbConfig.Create && connectDB && err != nil && strings.Contains(err.Error(), fmt.Sprintf("database \"%s\" does not exist (SQLSTATE 3D000)", databaseInfo.DatabaseName)) {
			zap.L().Info(fmt.Sprintf("Database %s does not exist. Trying to create it", databaseInfo.DatabaseName))
			database.Close()
			dblessConnectionString := strings.ReplaceAll(dbConfig.ConnectionString, fmt.Sprintf("dbname=%s", databaseInfo.DatabaseName), "")
			database, err = sql.Open("pgx", dblessConnectionString)
			helpers.CheckErrFatal(err)
			err = Create(database, databaseInfo)
			helpers.CheckErrFatal(err)
			database.Close()
			database, err = sql.Open("pgx", dbConfig.ConnectionString)
			helpers.CheckErrFatal(err)
		} else {
			helpers.CheckErrFatal(err)
		}

		// DB Connection Settings
		database.SetMaxIdleConns(5)
		database.SetMaxOpenConns(20)
		database.SetConnMaxLifetime(time.Hour)

		// Check Connection on StartUp
		helpers.CheckErrFatal(database.Ping())
	})
}
