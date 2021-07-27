package db

import (
	"database/sql"
	"fmt"
	"sync"

	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/config"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/helpers"
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
	switch dbConfig.Type {
	case "postgres":
		NewPostgresConnection(dbConfig, connectDB)
	case "mysql":
		NewMySQLConnection(dbConfig, connectDB)
	default:
		helpers.CheckErrFatal(fmt.Errorf("unrecognized database type %s", dbConfig.Type))
	}
}

func GetConnection() *sql.DB {
	return database
}

func GetConnectionInfo() DBInfo {
	return databaseInfo
}
