package db

import (
	"database/sql"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/FloThinksPi/golang-vuejs-bootstrap/internal/app/gopilot/config"
	"github.com/FloThinksPi/golang-vuejs-bootstrap/internal/app/gopilot/helpers"
	"github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.uber.org/zap"
)

var once sync.Once
var database *sql.DB
var databaseInfo DBInfo

type DBInfo struct {
	Host         string
	Port         string
	DatabaseName string
	User         string
	Password     string
	Type         string
}

func NewConnection(dbConfig config.DBConfig, connectDB bool) {
	once.Do(func() {
		if dbConfig.Type == "postgres" {
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
				Create(database, databaseInfo)
				database.Close()
				database, err = sql.Open("pgx", dbConfig.ConnectionString)
				helpers.CheckErrFatal(err)
			} else {
				helpers.CheckErrFatal(err)
			}
		} else {
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
				Create(database, databaseInfo)
				database.Close()
				database, err = sql.Open("mysql", dbConfig.ConnectionString)
				helpers.CheckErrFatal(err)
			} else {
				helpers.CheckErrFatal(err)
			}
		}

		// DB Connection Settings
		database.SetMaxIdleConns(5)
		database.SetMaxOpenConns(20)
		database.SetConnMaxLifetime(time.Hour)

		// Check Connection on StartUp
		helpers.CheckErrFatal(database.Ping())
	})
}

func GetConnection() *sql.DB {
	return database
}

func GetConnectionInfo() DBInfo {
	return databaseInfo
}
