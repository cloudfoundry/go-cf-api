//+build mage

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/magefile/mage/sh"
)

//###################################//
// Mage Commands used my a developer //
//###################################//

//////////////////
// Dependencies //
//////////////////

// Installs used CLI Tools in this project other than nodejs,yarn and go itself. Is a requirement for many mage commands to run.
func InstallDeps() error {
	if err := sh.Run("go", "get", "github.com/swaggo/swag/cmd/swag"); err != nil {
		return err
	}
	if err := sh.Run("go", "get", "golang.org/x/tools/cmd/godoc"); err != nil {
		return err
	}
	if err := sh.Run("go", "get", "gitlab.com/tslocum/godoc-static"); err != nil {
		return err
	}
	if err := sh.Run("go", "get", "github.com/volatiletech/sqlboiler/v4@latest"); err != nil {
		return err
	}
	if err := sh.Run("go", "get", "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest"); err != nil {
		return err
	}
	if err := sh.Run("go", "get", " github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest"); err != nil {
		return err
	}
	return nil
}

///////////////////////////
// Software Dev Commands //
///////////////////////////

// Generates sql builerplate code from the specified Database in the config and places it in internal
func GenerateSQLBoiler() error {
	if err := sh.Rm("./internal/app/gopilot/sqlboiler/mysql"); err != nil {
		return err
	}
	if err := sh.Rm("./internal/app/gopilot/sqlboiler/postgres"); err != nil {
		return err
	}
	if err := sh.Run("sqlboiler", "psql", "-c", "sqlboiler_postgres.toml"); err != nil {
		return err
	}
	if err := sh.Run("sqlboiler", "mysql", "-c", "sqlboiler_mysql.toml"); err != nil {
		return err
	}
	return nil
}

// Run a development web server for the frontend, that updates automatically on code changes
func FrontendDevServer() error {
	os.Chdir("web")
	defer os.Chdir("..")
	fmt.Println("Please connect to http://localhost:8080")
	fmt.Println("Abort with CTRL-C")
	if err := sh.RunV("yarn", "serve"); err != nil {
		return err
	}
	return nil
}

///////////////////////
// Software Building //
///////////////////////

// Runs go mod download and then installs the binary.
func Build() error {
	if err := sh.Rm("./build/gopilot"); err != nil {
		return err
	}
	if err := BuildFrontend(); err != nil {
		return err
	}
	if err := CreateAPIDocs(); err != nil {
		return err
	}
	if err := sh.RunV("go", "mod", "download"); err != nil {
		return err
	}
	if err := sh.RunV("go", "install", "./..."); err != nil {
		return err
	}
	return sh.RunV("go", "build", "-o", "build/gopilot", "cmd/main.go")
}

// Build the VueJS Frontend so it can be integrated into the go binary
func BuildFrontend() error {
	if err := sh.Rm("./cmd/frontend"); err != nil {
		return err
	}
	os.Chdir("web")
	defer os.Chdir("..")
	if err := sh.RunV("yarn", "install"); err != nil {
		return err
	}
	return sh.RunV("yarn", "build", "--dest", "../cmd/frontend")
}

///////////////////////////
// DOCUMENTATION HELPERS //
///////////////////////////

// Generates the swagger apidoc spec that later is included into the go binary and served on a productive system.
func CreateAPIDocs() error {
	if err := sh.Rm("./docs/swagger"); err != nil {
		return err
	}
	return sh.Run("swag", "init", "-o", "./docs/swagger", "--parseInternal", "--dir", "./cmd")
}

// Generates Godocs one can then set as a github page so devs can look at godocs in github
func CreateGoDocs() error {
	if err := sh.Rm("./docs/godocs"); err != nil {
		return err
	}
	time.Sleep(2 * time.Second)
	if err := sh.Run("mkdir", "-p", "./docs/godocs"); err != nil {
		return err
	}
	if err := sh.Run("godoc-static", "-site-name='Rocket Nine Labs Documentation'", "-site-description-file=./docs/description.md", "-destination=./docs/godocs", "."); err != nil {
		return err
	}
	return sh.Rm("./docs/godocs/docs.zip")
}

//////////
// MISC //
//////////

// Removes all generated files from the project
func Clean() error {
	if err := sh.Rm("./build"); err != nil {
		return err
	}
	if err := sh.Rm("./cmd/frontend"); err != nil {
		return err
	}
	if err := sh.Rm("./vendor"); err != nil {
		return err
	}
	if err := sh.Rm("./web/node_modules"); err != nil {
		return err
	}
	if err := sh.Rm("./repl"); err != nil {
		return err
	}
	if err := sh.Rm("./mage_output_file.go"); err != nil {
		return err
	}
	return nil
}
