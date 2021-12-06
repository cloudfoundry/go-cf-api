//go:build mage
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

// GetDependencies installs needed Go CLI Tools and Dependencies. Required to run go:generate or other generators.
func GetDependencies() error {
	// Check Required Tools
	requirements := [][]string{
		{"go", "version"},
		{"node", "--version"},
		{"yarn", "--version"},
	}
	for _, requirement := range requirements {
		if err := sh.Run(requirement[0], requirement[1:]...); err != nil {
			fmt.Println("Please make sure this tool is installed. It is a prerequisite.")
			return err
		}
	}
	// Install Doc Dependencies
	if err := sh.Run("yarn", "install", "--cwd", "./docs"); err != nil {
		return err
	}
	// Get Go Mod Packages
	if err := sh.Run("go", "get", "--tags=psql,unit,integration", "./..."); err != nil {
		return err
	}
	// Install CLIs
	packages := []string{
		"github.com/golang/mock/mockgen",
		"github.com/golangci/golangci-lint/cmd/golangci-lint",
		"github.com/swaggo/swag/cmd/swag",
		"golang.org/x/tools/cmd/godoc",
		"github.com/princjef/gomarkdoc/cmd/gomarkdoc",
		"github.com/volatiletech/sqlboiler/v4",
		"github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql",
		"github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql",
	}

	for _, pkg := range packages {
		if err := sh.Run("go", "install", pkg); err != nil {
			return err
		}
	}
	return nil
}

///////////////////////////
// Software Dev Commands //
///////////////////////////

// GenerateSQLBoiler generates sql builerplate code from the specified Database in the config and places it in ./internal/storage/db/models. Needs a running mysql and psql instance with ccdb database and a fully deployed schema to generate the code from.
func GenerateSQLBoiler() error {
	if err := sh.Rm("./internal/storage/db/models"); err != nil {
		return err
	}
	if err := os.MkdirAll("internal/storage/db/models/mocks", 0755); err != nil {
		return err
	}

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

// modifySQLBoilerFiles adds buildtags for multi db support. Also add go generated for test mocks. TODO -> Do this with SQLBoiler Natively(Needs Contribution)
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

// Build Runs go:generate and builds both binary(psql and mysql).
func Build() error {
	if err := sh.Rm("./build/cfapi*"); err != nil {
		return err
	}
	if err := sh.Run("go", "generate", "./..."); err != nil {
		return err
	}
	if err := sh.RunV("go", "build", "--tags=mysql", "-o", "build/cfapi_mysql", "cmd/main.go"); err != nil {
		return err
	}
	return sh.RunV("go", "build", "--tags=psql", "-o", "build/cfapi_psql", "cmd/main.go")
}

// Run runs go:generate and starts cfapi with config_psql.yaml.
func Run() error {
	if err := sh.Run("go", "generate", "./..."); err != nil {
		return err
	}
	return sh.RunV("go", "run", "cmd/main.go", "config_psql.yaml")
}

///////////////////////////
// DOCUMENTATION HELPERS //
///////////////////////////

// GenerateDocs generates Markdown Docs for Go Packages in ./docs/godocs
func GenerateDocs() error {
	if err := sh.Run("go", "generate", "./..."); err != nil {
		return err
	}
	const godocdir = "./docs/godocs/Packages"
	if err := sh.Rm(godocdir); err != nil {
		return err
	}
	if err := sh.Run("mkdir", "-p", godocdir); err != nil {
		return err
	}
	if err := sh.Run("gomarkdoc", "-u", "--tags=psql", "--output", fmt.Sprintf("%s/{{.ImportPath}}.md", godocdir), "./..."); err != nil {
		return err
	}
	return nil
}

//////////
// MISC //
//////////

// Clean removes all generated files from the project
func Clean() error {
	paths := []string{
		"./build",
		"./vendor",
		"./repl",
		"./internal/api/docs/docs.go",
		"./internal/api/docs/swagger.json",
		"./internal/api/docs/swagger.yaml",
		"./docs/swagger.yaml",
		"./docs/build",
		"./docs/.docusaurus",
		"./docs/godocs/Packages",
		"./docs/godocs/nmode_modules",
		"./main.go",
		"./mage_output_file.go",
	}

	for _, p := range paths {
		if err := sh.Rm(p); err != nil {
			return err
		}
	}

	return nil
}
