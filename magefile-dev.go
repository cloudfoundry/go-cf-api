// +build mage

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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
	if err := sh.Run("go", "get", "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest"); err != nil {
		return err
	}
	return nil
}

///////////////////////////
// Software Dev Commands //
///////////////////////////

// Generates sql builerplate code from the specified Database in the config and places it in internal
func GenerateSQLBoiler() error {
	if err := sh.Rm("./internal/app/cloudgontroller/sqlboiler"); err != nil {
		return err
	}
	if err := sh.Run("sqlboiler", "psql", "-c", "sqlboiler_postgres.toml"); err != nil {
		return err
	}
	if err := sh.Run("sqlboiler", "mysql", "-c", "sqlboiler_mysql.toml"); err != nil {
		return err
	}
	// Append build tag "integration" to every generated file so we can use "go test -tag integration" to switch between unit and integration tests
	r, _ := regexp.Compile(".*_test.go$")
	err := filepath.Walk("./internal/app/cloudgontroller/sqlboiler/mysql",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			if info.Mode().IsRegular() && r.MatchString(path) {
				err = addBuildTags(path, "./internal/app/cloudgontroller/sqlboiler/mysql_", []string{"integration", "mysql"})
				if err != nil {
					return err
				}
			} else if info.Mode().IsRegular() && !r.MatchString(path) {
				err = addBuildTags(path, "./internal/app/cloudgontroller/sqlboiler/mysql_", []string{"mysql"})
				if err != nil {
					return err
				}
			}
			return nil
		})
	if err != nil {
		return err
	}
	err = filepath.Walk("./internal/app/cloudgontroller/sqlboiler/postgres",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			if info.Mode().IsRegular() && r.MatchString(path) {
				err = addBuildTags(path, "./internal/app/cloudgontroller/sqlboiler/postgres_", []string{"integration", "postgres"})
				if err != nil {
					return err
				}
			} else if info.Mode().IsRegular() && !r.MatchString(path) {
				err = addBuildTags(path, "./internal/app/cloudgontroller/sqlboiler/postgres_", []string{"postgres"})
				if err != nil {
					return err
				}
			}
			return nil
		})
	if err != nil {
		return err
	}
	if err := sh.Rm("./internal/app/cloudgontroller/sqlboiler/postgres"); err != nil {
		return err
	}
	if err := sh.Rm("./internal/app/cloudgontroller/sqlboiler/mysql"); err != nil {
		return err
	}
	return nil
}

func addBuildTags(path string, outputPrefix string, tags []string) error {
	// Add build tag
	var data []byte
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	data = append(data, []byte(fmt.Sprintf("// +build %s\n", strings.Join(tags, ",")))...)
	data = append(data, content...)
	err = os.Remove(path)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("%s%s", outputPrefix, filepath.Base(path)), data, 0644)
	if err != nil {
		return err
	}
	return nil
}

///////////////////////
// Software Building //
///////////////////////

// Runs all generators we have and builds a binary.
func Build() error {
	if err := sh.Rm("./build/cloudgontroller"); err != nil {
		return err
	}
	if err := Generate(); err != nil {
		return err
	}
	if err := sh.RunV("go", "mod", "download"); err != nil {
		return err
	}
	if err := sh.RunV("go", "install", "./..."); err != nil {
		return err
	}
	return sh.RunV("go", "build", "-o", "build/cloudgontroller", "cmd/main.go")
}

// Runs generators whose result is included in cloudgontroller and runs cloudgontroller.
func Run() error {
	if err := createAPIDocs(); err != nil {
		return err
	}
	return sh.RunV("go", "run", "cmd/main.go", "config_postgres.yaml")
}

///////////////////////////
// DOCUMENTATION HELPERS //
///////////////////////////

// Runs all the generators we have that produce code/docs etc.
func Generate() error {
	if err := createAPIDocs(); err != nil {
		return err
	}
	if err := createGoDocs(); err != nil {
		return err
	}
	return nil
}

// Generates the swagger apidoc spec that later is included into the go binary and served on a productive system.
func createAPIDocs() error {
	symlinkPath := "./main.go"
	filePath := "./cmd/main.go"

	// Create symlink main.go in root
	err := os.Symlink(filePath, symlinkPath)
	if err != nil {
		return fmt.Errorf("failed to create symlink: %+v", err)
	}
	// Delete it after everything finished
	defer delSymLink(symlinkPath)

	// Generate Doc
	if err := sh.Rm("./internal/app/cloudgontroller/api/swagger"); err != nil {
		return fmt.Errorf("failed to remove swagger output directory: %+v", err)
	}
	if err := sh.Run("swag", "init", "-o", "./internal/app/cloudgontroller/api/swagger", "--parseInternal", "--parseDepth", "1", "--parseDependency", "--parseVendor"); err != nil {
		return fmt.Errorf("failed to run swagger generation: %+v", err)
	}
	return nil
}

func delSymLink(symlinkPath string) error {
	// Remove Symlink after generation
	if _, err := os.Lstat(symlinkPath); err == nil {
		if err := os.Remove(symlinkPath); err != nil {
			return fmt.Errorf("failed to unlink: %+v", err)
		}
	} else if os.IsNotExist(err) {
		return fmt.Errorf("failed to check symlink: %+v", err)
	}
	return nil
}

// Generates Godocs one can then set as a github page so devs can look at godocs in github
func createGoDocs() error {
	if err := sh.Rm("./docs"); err != nil {
		return err
	}
	time.Sleep(2 * time.Second)
	if err := sh.Run("mkdir", "-p", "./docs"); err != nil {
		return err
	}
	if err := sh.Run("godoc-static", "-site-name='CloudGontroller'", "-site-description-file=./Readme.md", "-destination=./docs", "."); err != nil {
		return err
	}
	return sh.Rm("./docs/docs.zip")
}

//////////
// MISC //
//////////

// Removes all generated files from the project
func Clean() error {
	if err := sh.Rm("./build"); err != nil {
		return err
	}
	if err := sh.Rm("./vendor"); err != nil {
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
