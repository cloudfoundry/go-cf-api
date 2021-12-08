//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/sh"
	"go.uber.org/zap"
	"os"

	_ "embed"
	"fmt"
	"github.com/cloudfoundry/go-cf-api/internal/config"
	"github.com/cloudfoundry/go-cf-api/internal/helpers"
	"github.com/cloudfoundry/go-cf-api/internal/logging"
	dbconfig "github.com/cloudfoundry/go-cf-api/internal/storage/db"
	"github.com/golobby/repl/interpreter"
)

//###################################//
// Mage Commands used my an operator //
//###################################//

/////////////////////
// RUNTIME HELPERS //
/////////////////////

var (
	currentInterpreter *interpreter.Interpreter
)

// DBCreate creates the Database that is specified in the config file
func DBCreate(configPath string) error {
	config, err := config.Get(configPath)
	if err != nil {
		return err
	}
	db, info, err := dbconfig.NewConnection(config.DB, false)
	if err != nil {
		return err
	}
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", info.DatabaseName))
	if err != nil {
		return err
	}
	zap.L().Info(fmt.Sprintf("Successfully created database %v", info.DatabaseName))
	return nil
}

// DBDelete deletes the Database that is specified in the config file
func DBDelete(configPath string) error {
	config, err := config.Get(configPath)
	if err != nil {
		return err
	}
	logging.Setup(config)
	db, info, err := dbconfig.NewConnection(config.DB, false)
	if err != nil {
		return err
	}
	_, err = db.Exec(fmt.Sprintf("DROP DATABASE %v", info.DatabaseName))
	if err != nil {
		return err
	}
	zap.L().Info(fmt.Sprintf("Successfully dropped database %v", info.DatabaseName))
	return nil
}

// DBRecreate recreates (Delete,Create,Migrates) the Database that is specified in the config file (Cf-Api Config is passed as first postitional parameter and sql file as second postitional parameter)
func DBRecreate(configPath string, sqlFilePath string) error {
	if err := DBDelete(configPath); err != nil {
		return err
	}
	if err := DBCreate(configPath); err != nil {
		return err
	}
	if err := DBLoad(configPath, sqlFilePath); err != nil {
		return err
	}
	return nil
}

// DBLoad loads a SQL file into the Database (Cf-Api Config is passed as first postitional parameter and sql file as second postitional parameter)
func DBLoad(configPath string, sqlFilePath string) error {
	config, err := config.Get(configPath)
	if err != nil {
		return err
	}
	logging.Setup(config)
	_, info, err := dbconfig.NewConnection(config.DB, false)
	if err != nil {
		return err
	}

	switch info.Type {
	case "postgres":
		if err := sh.RunV("psql", config.DB.ConnectionString, "-f", sqlFilePath); err != nil {
			return err
		}
	case "mysql":
		if err := sh.RunV("mysql",
			"-h", info.Host,
			"-P", info.Port,
			"-u", info.User,
			"-D", info.DatabaseName,
			fmt.Sprintf("--password=%s", info.Password),
			"--protocol=tcp",
			"-e", fmt.Sprintf("source %s", sqlFilePath),
		); err != nil {
			return err
		}
	}

	return nil
}

//go:embed internal/config/defaults.yml
var defaultConfig []byte

// GenerateDefaultConfig generate a config file with default values
func GenerateDefaultConfig(configPath string) error {
	err := os.WriteFile(configPath, defaultConfig, 0611)
	helpers.CheckErrFatal(err)
	return nil
}
