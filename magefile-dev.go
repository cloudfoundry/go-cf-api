// +build mage

package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path"
	"strings"

	"github.com/magefile/mage/sh"
)

//###################################//
// Mage Commands used my a developer //
//###################################//

//////////////////
// Dependencies //
//////////////////

// Installs used CLI Tools in this project other than nodejs,yarn and go itself. Is a requirement for many mage commands to run.
func Install() error {
	if err := sh.Run("go", "install", "github.com/golang/mock/mockgen"); err != nil {
		return err
	}
	if err := sh.Run("go", "install", "github.com/golangci/golangci-lint/cmd/golangci-lint"); err != nil {
		return err
	}
	if err := sh.Run("go", "install", "github.com/swaggo/swag/cmd/swag"); err != nil {
		return err
	}
	if err := sh.Run("go", "install", "golang.org/x/tools/cmd/godoc"); err != nil {
		return err
	}
	if err := sh.Run("go", "install", "github.com/princjef/gomarkdoc/cmd/gomarkdoc"); err != nil {
		return err
	}
	if err := sh.Run("go", "install", "github.com/volatiletech/sqlboiler/v4"); err != nil {
		return err
	}
	if err := sh.Run("go", "install", "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql"); err != nil {
		return err
	}
	if err := sh.Run("go", "install", "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql"); err != nil {
		return err
	}
	return nil
}

///////////////////////////
// Software Dev Commands //
///////////////////////////

// Generates sql builerplate code from the specified Database in the config and places it in internal
func GenerateSQLBoiler() error {
	if err := sh.Rm("./internal/storage/db/models"); err != nil {
		return err
	}
	if err := os.MkdirAll("internal/storage/db/models/mocks", 0755); err != nil {
		return err
	}

	if err := runSQLBoiler(); err != nil {
		return err
	}

	if err := modifySQLBoilerFiles("psql", true); err != nil {
		return err
	}
	if err := modifySQLBoilerFiles("mysql", false); err != nil {
		return err
	}

	if err := sh.Rm("./internal/storage/db/models/psql"); err != nil {
		return err
	}
	if err := sh.Rm("./internal/storage/db/models/mysql"); err != nil {
		return err
	}

	if err := sh.Run("go", "generate", "--tags=psql,mysql,unit", "./..."); err != nil {
		return err
	}

	return nil
}

func runSQLBoiler() error {
	if err := sh.Run("sqlboiler", "psql", "-c", "sqlboiler_psql.toml",
		"--no-driver-templates",
		"--templates", "sqlboiler-templates/shared",
		"--templates", "sqlboiler-templates/psql",
	); err != nil {
		return err
	}
	if err := sh.Run("sqlboiler", "mysql", "-c", "sqlboiler_mysql.toml",
		"--no-driver-templates",
		"--templates", "sqlboiler-templates/shared",
		"--templates", "sqlboiler-templates/mysql",
	); err != nil {
		return err
	}
	return nil
}

// Parses the AST to add additional comments and remove certain tests that fail
func modifySQLBoilerFiles(dbEngine string, addGoGenerate bool) error {
	// Append build tags to every generated file so we can use e.g. "go test -tags=db,psql" to switch between unit and db tests
	fileSet := token.NewFileSet()
	packages, err := parser.ParseDir(
		fileSet,
		fmt.Sprintf("internal/storage/db/models/%s", dbEngine),
		func(fi fs.FileInfo) bool { return true },
		parser.ParseComments,
	)
	for filepath, file := range packages["models"].Files {
		commentMap := ast.NewCommentMap(fileSet, file, file.Comments)
		_, filename := path.Split(filepath)

		comments := []*ast.Comment{{Text: fmt.Sprintf("// +build %s", dbEngine)}}
		if addGoGenerate && !strings.HasPrefix(filename, "boil_") {
			comments = append(comments, &ast.Comment{
				Text: fmt.Sprintf("//go:generate sh -c \"echo '\\x2bbuild unit' > ../../../../buildtags.txt && mockgen -source=$GOFILE -destination=mocks/%s -copyright_file=../../../../buildtags.txt && rm -f ../../../../buildtags.txt\"", filename),
			})
		}
		commentMap[file] = append([]*ast.CommentGroup{{List: comments}}, commentMap[file]...)
		file.Comments = commentMap.Comments()

		buf := &bytes.Buffer{}
		err = format.Node(buf, fileSet, file)
		if err != nil {
			return err
		}
		err = os.WriteFile(fmt.Sprintf("internal/storage/db/models/%s_%s", dbEngine, filename), buf.Bytes(), 0644)
		if err != nil {
			return err
		}
	}
	return err
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
	if err := sh.RunV("go", "build", "--tags=mysql", "-o", "build/cloudgontroller_mysql", "cmd/main.go"); err != nil {
		return err
	}
	return sh.RunV("go", "build", "--tags=psql", "-o", "build/cloudgontroller_psql", "cmd/main.go")
}

// Runs generators whose result is included in cloudgontroller and runs cloudgontroller.
func Run() error {
	if err := createAPIDocs(); err != nil {
		return err
	}
	return sh.RunV("go", "run", "cmd/main.go", "config_psql.yaml")
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
	if err := sh.Rm("./internal/api/swagger"); err != nil {
		return fmt.Errorf("failed to remove swagger output directory: %+v", err)
	}
	if err := sh.Run("swag", "init", "-o", "./internal/api/swagger", "--parseInternal", "--parseDepth", "1", "--parseDependency", "--parseVendor"); err != nil {
		return fmt.Errorf("failed to run swagger generation: %+v", err)
	}
	if err := sh.Rm("./docs/swagger.yaml"); err != nil {
		return fmt.Errorf("failed to remove swagger output directory: %+v", err)
	}
	if err := sh.Copy("./docs/swagger.yaml", "./internal/api/swagger/swagger.yaml"); err != nil {
		return err
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
	const godocdir = "./docs/godocs/Packages"
	if err := sh.Rm(godocdir); err != nil {
		return err
	}
	if err := sh.Run("mkdir", "-p", godocdir); err != nil {
		return err
	}
	if err := sh.Run("gomarkdoc", "-u", "--output", fmt.Sprintf("%s/{{.ImportPath}}.md", godocdir), "./..."); err != nil {
		return err
	}
	return nil
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
