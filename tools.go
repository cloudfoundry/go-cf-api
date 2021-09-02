// +build tools

package tools

// This file does not get compiled but keeps the tools versioned in go.mod/sum when running go mod tidy
import (
	_ "github.com/golang/mock/mockgen"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/magefile/mage"
	_ "github.com/swaggo/swag/cmd/swag"
	_ "github.com/volatiletech/sqlboiler/v4"
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql"
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql"
	_ "gitlab.com/tslocum/godoc-static"
	_ "golang.org/x/tools/cmd/godoc"
)
