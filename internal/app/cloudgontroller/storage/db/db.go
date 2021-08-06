package db

import (
	"database/sql"
	"fmt"
	"sync"

	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/config"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/helpers"
)

// nolint:gochecknoglobals
var (
	once         sync.Once
	database     *sql.DB
	databaseInfo Info
)

type Info struct {
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

func GetConnectionInfo() Info {
	return databaseInfo
}
