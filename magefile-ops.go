//+build mage

package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

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
	"github.com/magefile/mage/sh"
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
	db.Migrate(db.GetConnection())
	return nil
}

// Load SQL file into the Database
func DBLoad(configPath string, sqlFilePath string) error {
	config := config.Get(configPath)
	logging.Setup(config)
	if err := sh.Run("psql", config.DB.ConnectionString, "-f", sqlFilePath); err != nil {
		return err
	}
	return nil
}

// Seeds the Database that is specified in the config file with a large number of entries for performance tests
func DBPerfSeed() error {
	fmt.Println("Not Yet Implemented")
	return nil
}

// Generate a config file with default values
func GenerateDefaultConfig(configPath string) error {
	tmp := &config.CloudgontrollerConfig{}
	err := defaults.Set(tmp)
	helpers.CheckErrFatal(err)
	data, err := yaml.Marshal(tmp)
	helpers.CheckErrFatal(err)
	err = ioutil.WriteFile(configPath, data, 0611)
	helpers.CheckErrFatal(err)
	return nil
}
