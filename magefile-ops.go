// +build mage

package main

import (
	"embed"
	"os"
	"path/filepath"

	"github.com/magefile/mage/sh"
	migrate "github.com/rubenv/sql-migrate"
	"go.uber.org/zap"

	"fmt"
	"runtime/debug"

	"github.com/creasty/defaults"
	"github.com/golobby/repl/interpreter"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/config"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/helpers"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/logging"
	dbconfig "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/storage/db"
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
	config, err := config.Get(configPath)
	if err != nil {
		return err
	}
	var container string
	switch config.DB.Type {
	case "postgres":
		container = "postgres"
	case "mysql":
		container = "mariadb"
	default:
		return fmt.Errorf("Unrecognized DB type: %s", config.DB.Type)
	}
	err = sh.RunV("docker-compose", "-f", "docker-compose.yml", "up", "-d", container, "uaa")
	return err
}

// Starts a docker container running the database specified in the config file
func DBStart(configPath string) error {
	config, err := config.Get(configPath)
	if err != nil {
		return err
	}
	var container string
	switch config.DB.Type {
	case "postgres":
		container = "postgres"
	case "mysql":
		container = "mariadb"
	default:
		return fmt.Errorf("Unrecognized DB type: %s", config.DB.Type)
	}
	err = sh.RunV("docker-compose", "-f", "docker-compose.yml", "up", "-d", container)
	return err
}

// Creates the Database that is specified in the config file
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

// Deletes the Database that is specified in the config file
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

//nolint:gochecknoglobals
//go:embed internal/app/cloudgontroller/storage/db/migrations
var migrations embed.FS

// Migrates the Database that is specified in the config file to the newest schema
func DBMigrate(configPath string) error {
	config, err := config.Get(configPath)
	if err != nil {
		return err
	}
	logging.Setup(config)
	db, info, err := dbconfig.NewConnection(config.DB, false)
	if err != nil {
		return err
	}

	// Load migrations depending on DB Type
	folder := fmt.Sprintf("migrations/%s", info.Type)

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
	n, err := migrate.Exec(db, info.Type, migrations, migrate.Up)
	zap.L().Info(fmt.Sprintf("Applied %v DB Migration Steps", n))
	helpers.CheckErrFatal(err)
	return nil
}

// Load SQL file into the Database
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
