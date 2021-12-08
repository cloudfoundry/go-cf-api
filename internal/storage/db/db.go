package db

import (
	"database/sql"
	"fmt"

	"github.com/cloudfoundry/go-cf-api/internal/config"
)

type Info struct {
	Host         string
	Port         string
	DatabaseName string
	User         string
	Password     string
	Type         string
}

func NewConnection(dbConfig config.DBConfig, connectDB bool) (*sql.DB, Info, error) {
	switch dbConfig.Type {
	case "postgres":
		return NewPostgresConnection(dbConfig, connectDB)
	case "mysql":
		return NewMySQLConnection(dbConfig, connectDB)
	}
	return nil, Info{}, fmt.Errorf("unrecognized database type %s", dbConfig.Type)
}
