//+build mage

package main

import (
	"os"
	"path/filepath"

	"github.com/magefile/mage/sh"
	"go.uber.org/zap"

	"fmt"
	"runtime/debug"

	"github.com/creasty/defaults"
	"github.com/golobby/repl/interpreter"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/config"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/helpers"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/logging"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/storage/db"
	"gopkg.in/yaml.v2"

	"github.com/c-bata/go-prompt"
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

func completer(d prompt.Document) []prompt.Suggest {
	var s []prompt.Suggest
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func handler(input string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panic: %v\n%s", err, debug.Stack())
		}
	}()
	out, err := currentInterpreter.Eval(input)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Print(out)
}

// Opens a interactive golang console one can include internal modules and do queries/debugging interactively in go
func Console() error {
	symlink := filepath.Join("./", "pkg")
	os.Symlink("./internal", symlink)
	defer os.Remove(symlink)
	wd, err := os.Getwd()
	helpers.CheckErrFatal(err)

	currentInterpreter, err = interpreter.NewSession(wd)
	helpers.CheckErrFatal(err)

	_, err = currentInterpreter.Eval(":e 1")
	helpers.CheckErrFatal(err)

	_, err = currentInterpreter.Eval(":e import \"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/helpers\"")
	helpers.CheckErrFatal(err)

	_, err = currentInterpreter.Eval(":e import \"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/logging\"")
	helpers.CheckErrFatal(err)

	p := prompt.New(handler, completer, prompt.OptionPrefix("console> "))
	p.Run()
	return nil
}


// Starts a docker container running the database specified in the config file
func StartDBAndUaaContainers(configPath string) error {
	config := config.Get(configPath)
	var container string
	switch config.DB.Type {
	case "postgres":
		container = "postgres"
	case "mysql":
		container = "mariadb"
	default:
		return fmt.Errorf("Unrecognized DB type: %s", config.DB.Type)
	}
	err := sh.RunV("docker-compose", "-f", "docker-compose-dev.yaml", "up", "-d", container, "uaa")
	return err
}


// Starts a docker container running the database specified in the config file
func DBStart(configPath string) error {
	config := config.Get(configPath)
	var container string
	switch config.DB.Type {
	case "postgres":
		container = "postgres"
	case "mysql":
		container = "mariadb"
	default:
		return fmt.Errorf("Unrecognized DB type: %s", config.DB.Type)
	}
	err := sh.RunV("docker-compose", "-f", "docker-compose-dev.yaml", "up", "-d", container)
	return err
}

// Creates the Database that is specified in the config file
func DBCreate(configPath string) error {
	config := config.Get(configPath)
	logging.Setup(config)
	db.NewConnection(config.DB, false)
	db.Create(db.GetConnection(), db.GetConnectionInfo())
	return nil
}

// Deletes the Database that is specified in the config file
func DBDelete(configPath string) error {
	config := config.Get(configPath)
	logging.Setup(config)
	db.NewConnection(config.DB, false)
	db.Drop(db.GetConnection(), db.GetConnectionInfo())
	return nil
}

// Recreates (Delete,Create,Migrates) the Database that is specified in the config file
func DBRecreate(configPath string) error {
	if err := DBDelete(configPath); err != nil {
		return err
	}
	if err := DBCreate(configPath); err != nil {
		return err
	}
	if err := DBMigrate(configPath); err != nil {
		return err
	}
	return nil
}

// Migrates the Database that is specified in the config file to the newest schema
func DBMigrate(configPath string) error {
	config := config.Get(configPath)
	logging.Setup(config)
	db.NewConnection(config.DB, true)
	db.Migrate(config.DB.Type, db.GetConnection())
	return nil
}

// Load SQL file into the Database
func DBLoad(configPath string, sqlFilePath string) error {
	config := config.Get(configPath)
	logging.Setup(config)
	db.NewConnection(config.DB, false)
	if db.GetConnectionInfo().Type == "postgres" {
		if err := sh.RunV("psql", config.DB.ConnectionString, "-f", sqlFilePath); err != nil {
			return err
		}
	} else if db.GetConnectionInfo().Type == "mysql" {
		zap.L().Warn("Loading Mysql Dumps not yet implemented")
		if err := sh.RunV("mysql", "-h", db.GetConnectionInfo().Host, "-P", db.GetConnectionInfo().Port, "-u", db.GetConnectionInfo().User, fmt.Sprintf("--password=%s", db.GetConnectionInfo().Password), "--protocol=tcp", "-e", fmt.Sprintf("source %s", sqlFilePath)); err != nil {
			return err
		}
	}

	return nil
}

// Seeds the Database that is specified in the config file with a large number of entries for performance tests
func DBPerfSeed() error {
	zap.L().Warn("Not Yet Implemented")
	return nil
}

// Generate a config file with default values
func GenerateDefaultConfig(configPath string) error {
	tmp := &config.CloudgontrollerConfig{}
	err := defaults.Set(tmp)
	helpers.CheckErrFatal(err)
	data, err := yaml.Marshal(tmp)
	helpers.CheckErrFatal(err)
	err = os.WriteFile(configPath, data, 0611)
	helpers.CheckErrFatal(err)
	return nil
}
