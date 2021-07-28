![unit tests](https://github.tools.sap/cloudfoundry/cloudgontroller/workflows/Run%20unit%20tests/badge.svg) ![db tests](https://github.tools.sap/cloudfoundry/cloudgontroller/workflows/Run%20database%20tests/badge.svg) ![build](https://github.tools.sap/cloudfoundry/cloudgontroller/workflows/Build%20binaries/badge.svg)
# cloudgontroller Webserver
## Prerequisits
- Go 1.16 minimum: https://golang.org/dl/, can either be installed with GVM(Go version Manager) or via `brew install go`
- Mage (Makefile alternative in go): https://github.com/magefile/mage, can be installed with `brew install mage`
- Mysql and Postgres CLI Tools (mysql, mysqldump, psql, pgdump etc..), can be installed with `brew install mysql postgres`

Install the rest of the used CLI Dependencies in this project
```bash
mage InstallDeps
```
All compile time dependencies will be downloaded when compiling (If you use the mage commands)

## Preparing Dev Database

To run cloudgontroller we need a propper ccdb.
To do this we have a docker-compose file to create the DBs and sql dumps from a ccdb that one can execute to create a ccdb.

```bash
// Start Databases (Mariadb + Postgres)
docker compose -f docker-compose-dev.yaml up -d
// Create Database (Just implemented for Postgres Currently)
// It is ok if psql cli reports insert errors (table already exists)
mage DBCreate config_psql.yaml
mage DBLoad config_psql.yaml database_dumps/3.102.0_psql_ccdb.sql
// For Mariadb use this
mage DBCreate config_mysql.yaml
mage DBLoad config_mysql.yaml database_dumps/3.102.0_mysql_ccdb.sql
```

Other noteable operations on the db are
```bash
// Drop DB
mage DBDelete config_mysql.yaml
// Drop and Recreate DB (will be empty)
mage DBRecreate config_mysql.yaml
```

## Running commands
Different database code and types of test are conditionally compiled and run based on [Go build tags](https://pkg.go.dev/cmd/go#hdr-Build_constraints).
To target a specific database and/or test set, use:
```
go -tags=psql,unit <go command>
```

To avoid having to type out tags every time, you can instead:
```
export GOFLAGS="-tags=psql,unit"
```

This will make `go test ./...` just run the unit tests and also allow you to just run `go run cmd/main.go config_psql.yaml`

## Starting It
Simply
```bash
// Make a binary to use against a mysql instance
go run --tags=psql cmd/main.go config_psql.yaml
// Make a binary to use against a postgres instance
go run --tags=mysql cmd/main.go config_psql.yaml
```

The values out of the `config_*.yml` files as well as `sqlboiler_*.toml` match the credentials in `docker-compose-dev.yaml`.
So as long as you use that docker compose file all credentials match and you dont need to mess around with the configs.

The method how we use sqlboiler for both technology stacks(psql and mysql) leads us to a problem which generated model to use.
We solved this by including build flags to include certain code on build time e.g. include psql code and exclude mysql code.
This will then produce two binaries, one for each db. Running on a psql config with a mysql binarie will result in a startup error.

There is also a mage command which outputs both binary files.
It will also run every code/doc/apidoc generator we have beforehand and then use
go build to produce a binary in the `build` folder.
```bash
mage build
```

Once Running Access it e.g. with:

http://localhost:8080/api/v3/buildpacks

http://localhost:8080/docs/v3


# Running Tests
Simply
```bash
// Run all unittests
go test -tags="unit,psql" ./...
// Run db tests for postgres
go test --tags="db,psql" ./internal/app/cloudgontroller/sqlboiler -test.config sqlboiler_psql.toml
// Run db tests for mysql
go test --tags="db,mysql" ./internal/app/cloudgontroller/sqlboiler -test.config sqlboiler_mysql.toml
```

# Current Feature List
- Structured Logging
- Prometheus Endpoint Metrics
- Operations Console (REPL)
- DBs (Postgres, MySQL) Autogenerates Go Code
- DB Migrations
- Read Config
- Webserver
- Auto API Docs
- Godocs Auto Gen
- Assets in memory
- Rate Limiting


# PoC ToDos
https://jtrack.wdf.sap.corp/browse/CFP-1731?jql=project%20%3D%20CFP%20AND%20labels%20%3D%20CCPoC
# MVP ToDos
https://jtrack.wdf.sap.corp/browse/CFP-1731?jql=project%20%3D%20CFP%20AND%20labels%20%3D%20CCMvP
