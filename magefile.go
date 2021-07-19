//+build mage

package main

import (
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"os"
	"text/tabwriter"

	"github.com/FloThinksPi/golang-vuejs-bootstrap/internal/app/gopilot/helpers"
)

//#######################//
// General Mage Commands //
//#######################//

// Prints this helptext for available commands
func Help() error {

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug|tabwriter.DiscardEmptyColumns)

	pack := getPackage("./magefile.go")
	fmt.Fprintf(w, "----- General Commands, defined in %v -----\n", pack.Filenames[0])
	printHelpText(pack, w)
	fmt.Fprintln(w)

	pack = getPackage("./magefile-ops.go")
	fmt.Fprintf(w, "----- Commands for Operators, defined in %v -----\n", pack.Filenames[0])
	printHelpText(pack, w)
	fmt.Fprintln(w)

	pack = getPackage("./magefile-dev.go")
	fmt.Fprintf(w, "----- Commands for Developers, defined in %v -----\n", pack.Filenames[0])
	printHelpText(pack, w)

	w.Flush()
	return nil
}

func getPackage(path string) *doc.Package {
	fset := token.NewFileSet() // positions are relative to fset
	obj, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	helpers.CheckErrFatal(err)
	objFile := []*ast.File{obj}
	pack, err := doc.NewFromFiles(fset, objFile, path, doc.AllMethods)
	helpers.CheckErrFatal(err)
	return pack
}

func printHelpText(pack *doc.Package, w *tabwriter.Writer) {
	for _, f := range pack.Funcs {
		fmt.Fprintf(w, "%v \t %v", f.Name, f.Doc)
	}
}
