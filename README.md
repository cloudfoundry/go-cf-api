[![CodeGen](https://github.com/cloudfoundry/go-cf-api/actions/workflows/code-gen.yml/badge.svg)](https://github.com/cloudfoundry/go-cf-api/actions/workflows/code-gen.yml)
[![Build Documentation](https://github.com/cloudfoundry/go-cf-api/actions/workflows/build-docs.yml/badge.svg)](https://github.com/cloudfoundry/go-cf-api/actions/workflows/build-docs.yml)
[![Deploy Documentation](https://github.com/cloudfoundry/go-cf-api/actions/workflows/deploy-docs.yml/badge.svg)](https://github.com/cloudfoundry/go-cf-api/actions/workflows/deploy-docs.yml)

[![Build Binaries](https://github.com/cloudfoundry/go-cf-api/actions/workflows/build-binaries.yml/badge.svg)](https://github.com/cloudfoundry/go-cf-api/actions/workflows/build-binaries.yml)
[![Integration Tests](https://github.com/cloudfoundry/go-cf-api/actions/workflows/integration-tests.yml/badge.svg)](https://github.com/cloudfoundry/go-cf-api/actions/workflows/integration-tests.yml)
[![Unit Tests](https://github.com/cloudfoundry/go-cf-api/actions/workflows/unit-tests.yml/badge.svg)](https://github.com/cloudfoundry/go-cf-api/actions/workflows/unit-tests.yml)

[![Linter](https://github.com/cloudfoundry/go-cf-api/actions/workflows/lint.yml/badge.svg)](https://github.com/cloudfoundry/go-cf-api/actions/workflows/lint.yml)
[![CodeQL Analyze](https://github.com/cloudfoundry/go-cf-api/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/cloudfoundry/go-cf-api/actions/workflows/codeql-analysis.yml)

# GO-CF-API
A proof of concept implementation for [cloud_controller_ng](https://github.com/cloudfoundry/cloud_controller_ng), written in Go

## Development Setup
- [Go](https://golang.org/dl) version 1.17
- [Mage](https://github.com/magefile/mage) (Makefile alternative in go) --> `go install github.com/magefile/mage`
- [Mysql](https://mariadb.com/kb/en/mysql-command-line-client/) and [Postgres](https://www.postgresql.org/docs/13/app-psql.html) CLI Tools (mysql, mysqldump, psql, pgdump etc.)
- Packages and other needed command line tools --> `mage GetDependencies`
#### Optional
- [NodeJS](https://nodejs.org/en/) and [Yarn](https://yarnpkg.com/) (for working with the documentation in /docs)
- [Git LFS](https://git-lfs.github.com/) (to get assets resolved in /docs)
- [Docker](https://www.docker.com/) to make your life easier setting up runtime dependecies

## Prepare dev database
To run go-cf-api we need a proper Cloud Controller database.
To do this we have a docker-compose file to start the DB instances and SQL dumps from an existing CloudController_NG that can be imported to create a DB.
There are `mage` commands for most database operations.

Run `mage help` to see the full list of available commands.

First start needed runtime components with:
```bash
docker compose up -d 

[+] Running 3/0
 ⠿ Container go-cf-api_mariadb_1   Running                                                                                                                                                                                                                                                                   0.0s
 ⠿ Container go-cf-api_postgres_1  Running                                                                                                                                                                                                                                                                   0.0s
 ⠿ Container go-cf-api_uaa_1       Running 
```
To create a usable db we import a prepared database dump.
### Postgres
```bash
mage DBCreate config_psql.yaml
mage DBLoad config_psql.yaml database_dumps/3.102.0_psql_ccdb.sql
```

### MySQL/MariaDB
```bash
mage DBCreate config_mysql.yaml
mage DBLoad config_mysql.yaml database_dumps/3.102.0_mysql_ccdb.sql
```

##### Useful database operations
* Drop database
	```bash
	mage DBDelete config_psql.yaml
	```
* Drop and create database and import new dump
	```bash
	mage DBRecreate config_psql.yaml database_dumps/3.102.0_psql_ccdb.sql
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

Available Tags are:
- mysql (adds mysql database support)
- psql (adds psql database support)
- unit (includes unittests)
- integration (includes integration tests)

Note here that you can build/run go-cf-api just with either `mysql` or `psql` support/tag.
The intended way to support both dbs is to build two separate binaries and depending on the backend use the fitting one.
The `unit` and `integration` tags are just relevant for running `go test` but either `mysql` or `psql` must be always provided otherwise there are missing imports. 

## Starting the GO-CF-API
* Postgres
	```bash
	go run --tags=psql cmd/main.go config_psql.yaml
	```
* MySQL/MariaDB
	```bash
	go run --tags=mysql cmd/main.go config_mysql.yaml
	```

The default values in the `config_{db}.yml` and `sqlboiler_{db}.toml` files do match the credentials for each database in `docker-compose.yml`.
If you use other DB setup the credentials and config must be adopted.

The API should be accessible at
```
http://localhost:8080/v3
```

The API Documentation should be accessible at
```
http://localhost:8080/docs/v3
```

## List of Implemented Endpoints
Following endpoints are implemented and feature equal(just v3 endpoints) to cloud_controller_ng.


#### GO-CF-API Specific Endpoints
- http://localhost:8080/healhz (Health Endpoint)
- http://localhost:8080/metrics (Prometheus Metrics)
- http://localhost:8080/docs/v3 (API Documentation)
#### V3 API
- http://localhost:8080/
- http://localhost:8080/v3
- http://localhost:8080/v3/info
- http://localhost:8080/v3/buildpacks (GET)
- http://localhost:8080/v3/buildpacks/:guid (GET,POST)
- http://localhost:8080/v3/security_groups (GET)
- http://localhost:8080/v3/security_groups/:guid (GET)


## Querying locally
To query most endpoints locally, you will need a JWT. The easiest way to get this is to create a UAA client with the [uaac](https://github.com/cloudfoundry/cf-uaac) CLI then login with then `cf` CLI. To create a user called `bob`:
```bash
uaac target http://localhost:8095
uaac token client get admin -s adminsecret
uaac user add bob -p password --emails none@none.local
cf api http://localhost:8080
cf auth bob password
```

Now various endpoints can be queried using e.g. `cf curl /v3/buildpacks`

To authenticate as admin (for example to assign roles to `bob`):
```bash
cf auth admin adminsecret --client-credentials
```

## SQL Boiler
To support different databases (Postgres and MySQL/MariaDB), we generate two sets of database models with `sqlboiler`. These are then included in the same package and given build tags to control when to compile each model set into the binary. Two binaries are then produced - one for each database. Attempting to use e.g. the `mysql` binary on a Postgres database will result in a startup error.

To facilitate the generating, combining and build tagging of the `sqlboiler` models, there is a mage task to automate it:
```bash
mage GenerateSQLBoiler
```

The two binaries can then be compiled for your current OS/architecture with:
```bash
mage build
```

The binaries can be compiled for other architectures by exporting the `GOOS` and `GOARCH` environment variables.

## Documentation

The documentation is based on [docusaurus]() and can be found at https://pages.github.com/cloudfoundry/go-cf-api/.
It is defined under /docs in this project and is then served by github-pages from the `gh-pages` branch which is detached from the other branches.
A GitHub workflow ensures the `gh-pages` branch always reflects the state of the main branch /docs folder.

Package documentation is auto generated using [gomarcdoc]() and integrated into [docusaurus]().
Similarly the API documentation is also integrated into [docusaurus]().
Additionally, binaries will also serve the API documentation as swagger ui at:
```
http://localhost:8080/docs/v3
```

## Running linter
This project uses [golangci-lint](https://golangci-lint.run/) to ensure code is formatted correctly and return values are checked etc.
To run the linter, run:
```
golangci-lint run --build-tags psql,unit,integration
```

There is a GitHub Action that runs the linter on every push. No code should be pushed until it passes the linter.

## Running tests
Different tags are used to control which tests are run:
* Unit tests
	```bash
	go test -tags="unit,psql" ./...
	```
* Integration tests (Postgres) - require running database
	```bash
	go test -tags=integration,psql -parallel=1 -p=1 ./... -args $PWD/config_psql.yaml
	```
* Integration tests (MySQL/MariaDB) - require running database
	```bash
	go test -tags=integration,mysql -parallel=1 -p=1 ./... -args $PWD/config_mysql.yaml
	```