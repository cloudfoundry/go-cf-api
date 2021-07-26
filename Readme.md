# cloudgontroller Webserver
## Prerequisits
- Go 1.16 minimum: https://golang.org/dl/ (Install with GVM)
- Mage (Makefile alternative in go): https://github.com/magefile/mage

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
mage DBCreate config.yaml
mage DBLoad config.yaml database_dumps/3.102.0_postgres_ccdb.sql
mage DBLoad config.yaml database_dumps/buildpacks_table_pgdump_aws_staging.sql
```

Other noteable operations on the db are
```bash
// Drop and Recreate DB
mage DBRecreate
```

## Starting It
Simly
```bash
go run cmd/main.go config.yaml
```
is sufficient

There is also a mage command which outputs a binary file. 
It will also run every code/doc/apidoc generator we have beforehand and then use
go build to produce a binary in the `build` folder.
```bash
mage build
```

Once Running Access it e.g. with:

http://localhost:8080/api/v3/buildpacks

http://localhost:8080/docs/v3


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
