package db

import (
	"database/sql"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/config"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/helpers"
	"go.uber.org/zap"
)

func NewMySQLConnection(dbConfig config.DBConfig, connectDB bool) {
	if dbConfig.Type != "mysql" {
		helpers.CheckErrFatal(fmt.Errorf("cannot use database of type %s with mysql binary", dbConfig.Type))
	}
	once.Do(func() {
		// Parse Infos from Connection String
		cfg, err := mysql.ParseDSN(dbConfig.ConnectionString)
		helpers.CheckErrFatal(err)
		databaseInfo.Host, databaseInfo.Port, err = net.SplitHostPort(cfg.Addr)
		helpers.CheckErrFatal(err)
		databaseInfo.DatabaseName = cfg.DBName
		databaseInfo.User = cfg.User
		databaseInfo.Password = cfg.Passwd
		databaseInfo.Type = dbConfig.Type

		if !connectDB {
			cfg.DBName = ""
			dbConfig.ConnectionString = cfg.FormatDSN()
		}

		zap.L().Debug(fmt.Sprintf("Using DB Connection String: %v",
			helpers.Redact(dbConfig.ConnectionString, []string{databaseInfo.Password})))
		database, err = sql.Open("mysql", dbConfig.ConnectionString)
		helpers.CheckErrFatal(err)
		// If DB is Missing, we Create it
		err = database.Ping()
		if err != nil {
			fmt.Print(err.Error())
		}
		if dbConfig.Create && connectDB && err != nil && strings.Contains(err.Error(), fmt.Sprintf("Error 1049: Unknown database '%s'", databaseInfo.DatabaseName)) {
			zap.L().Info(fmt.Sprintf("Database %s does not exist. Trying to create it", databaseInfo.DatabaseName))
			database.Close()
			dblessCfg := cfg.Clone()
			dblessCfg.DBName = ""
			dblessCfg.FormatDSN()
			database, err = sql.Open("mysql", dblessCfg.FormatDSN())
			helpers.CheckErrFatal(err)
			err = Create(database, databaseInfo)
			helpers.CheckErrFatal(err)
			database.Close()
			database, err = sql.Open("mysql", dbConfig.ConnectionString)
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
