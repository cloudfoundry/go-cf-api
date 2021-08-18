package db

import (
	"database/sql"
	"fmt"

	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/config"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/helpers"
)

type Info struct {
	Host         string
	Port         string
	DatabaseName string
	User         string
	Password     string
	Type         string
}

func NewConnection(dbConfig config.DBConfig, connectDB bool) (*sql.DB, Info) {
	switch dbConfig.Type {
	case "postgres":
		return NewPostgresConnection(dbConfig, connectDB)
	case "mysql":
		return NewMySQLConnection(dbConfig, connectDB)
	}
	helpers.CheckErrFatal(fmt.Errorf("unrecognized database type %s", dbConfig.Type))
	return nil, Info{}
}
