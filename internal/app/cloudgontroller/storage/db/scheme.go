package db

import (
	"database/sql"
	"embed"
	"fmt"
	migrate "github.com/rubenv/sql-migrate"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/config"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/helpers"
	"go.uber.org/zap"
	"os"
)

//go:embed migrations
var migrations embed.FS

func Migrate(db *sql.DB) {
	// Load migrations depending on DB Type
	folder := fmt.Sprintf("migrations/%s", GetConnectionInfo().Type)

	// Convert embed.FS to a usable migration source for sql-migrate
	migrations := &migrate.AssetMigrationSource{
		Asset: migrations.ReadFile,
		AssetDir: func() func(string) ([]string, error) {
			return func(path string) ([]string, error) {
				dirEntry, err := migrations.ReadDir(path)
				if err != nil {
					return nil, err
				}
				entries := make([]string, 0)
				for _, e := range dirEntry {
					entries = append(entries, e.Name())
				}

				return entries, nil
			}
		}(),
		Dir: folder,
	}

	// Execute the UP Migration
	n, err := migrate.Exec(db, config.Get().DB.Type, migrations, migrate.Up)
	zap.L().Info(fmt.Sprintf("Applied %v DB Migration Steps", n))
	helpers.CheckErrFatal(err)
}

func Create(db *sql.DB, info DBInfo) error {
	_, err := database.Exec(fmt.Sprintf("CREATE DATABASE %v", info.DatabaseName))
	if err != nil {
		return err
	}
	zap.L().Info(fmt.Sprintf("Sucessfully created database %v", info.DatabaseName))
	return nil
}

func Drop(db *sql.DB, info DBInfo) error {
	_, err := database.Exec(fmt.Sprintf("DROP DATABASE %v", info.DatabaseName))
	if err != nil {
		return err
	}
	zap.L().Info(fmt.Sprintf("Sucessfully droped database %v", info.DatabaseName))
	return nil
}

func Load(db *sql.DB, info DBInfo, sqlFilePath string) error {
	bytes, ioErr := os.ReadFile(sqlFilePath)
	if ioErr != nil {
		zap.L().Error(fmt.Sprintf("Loading SQL file %s failed", sqlFilePath))
	}
	_, err := database.Exec(string(bytes))
	if err != nil {
		return err
	}
	zap.L().Info(fmt.Sprintf("Sucessfully loaded SQL file into database %v", info.DatabaseName))
	return nil
}
