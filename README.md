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
go test ./...
// Run all integration tests (But not DB tests)
go test --tags=integration ./...
// Run Integration Tests for Postgres
go test --tags=mysql ./internal/app/cloudgontroller/sqlboiler -test.config $PWD/sqlboiler_mysql.toml
// Run Integration Tests for Postgres
go test --tags=psql ./internal/app/cloudgontroller/sqlboilero -test.config $PWD/sqlboiler_psql.toml
```

# Existing Features
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


# Needed for POC Demonstration
Local Dev Demo
- Get Public Keys from UAA
- Authentication (JWT)
- Rollen
- VCAP-REQUEST_ID middleware
- Metric Middleware
- Logging Redact Secrets (From SQL,Echo)
- Tests
- GoDocs
- Swagger Docs

For Deployment
- Health Endpoint
- Extra Bosh Release
  - Gorouter Routing Adaptions
  - CC Config
  - Promscraper


# MVP

- Cloud Gontroller
  - Implement two simple GET endpoints (accessible for all roles): GET /v3/buildpacks/:guid + GET /v3/buildpacks (incl. query parameters, e.g. page, per_page)
  - Handle authentication (no roles)
  - Add VCAP-Request-ID header
  - Adhere to rate limits (authenticated + unauthenticated)
  - Write logs (should show up in Kibana)
  - Emit metrics (should show up in Grafana)
  - 100% test coverage, well-formatted code, no (severe) linting issues

- Product integration
  - Create a BOSH release (Cloud Gontroller as submodule, incl. route-registrar)
  - Add release to cf component
  - Enable deployment via config option (+ scaling options)
  - Routing done by Gorouter (list of supported endpoints as config option)
