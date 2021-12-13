![lint](https://github.com/cloudfoundry/go-cf-api/workflows/Lint/badge.svg) ![unit tests](https://github.com/cloudfoundry/go-cf-api/workflows/Run%20unit%20tests/badge.svg) ![db tests](https://github.com/cloudfoundry/go-cf-api/workflows/Run%20database%20tests/badge.svg) ![build](https://github.com/cloudfoundry/go-cf-api/workflows/Build%20binaries/badge.svg)


# go-cf-api
A **proof of concept implementation** as an alternative to [cloud_controller_ng](https://github.com/cloudfoundry/cloud_controller_ng), written in Go.\
Currently, this project is intended to be deployed alongside the existing `cloud_controller_ng` with a L7 router to forward requests to endpoints that have been completed to the Go implementation (based on both path and HTTP method). A deployment using this setup should pass [CATS](https://github.com/cloudfoundry/cf-acceptance-tests).
> Since this project only implements a subset of the CF API endpoints, it should not be deployed standalone.

## Deploying with cf-deployment
See the [ADR](https://cloudfoundry.github.io/go-cf-api/adrs/ADR%20-%20Traffic%20Splitter#decision-outcome) for more information on this deployment approach.

To deploy `go-cf-api` into a [cf-deployment](https://github.com/cloudfoundry/cf-deployment), include
[this ops file](https://github.com/cloudfoundry-incubator/cf-performance-tests-pipeline/blob/main/operations/deploy-go-cf-api.yml).
This will use [HAProxy](https://github.com/cloudfoundry/haproxy-boshrelease) to route all requests to the `go-cf-api` where that endpoint and method has been implemented.
Any endpoints that are not yet implemented by `go-cf-api` will continue to be served by [cloud_controller_ng](https://github.com/cloudfoundry/cloud_controller_ng).
### Links

- [Documentation](https://cloudfoundry.github.io/go-cf-api/)
- [BOSH Release](https://github.com/cloudfoundry/go-cf-api-release)
- [Opsfile for cf-deployment](https://github.com/cloudfoundry-incubator/cf-performance-tests-pipeline/blob/main/operations/deploy-gontroller.yml)

## Development Setup
### Requirements
- [Go](https://golang.org/dl) version 1.17
- [Mage](https://github.com/magefile/mage) (Makefile alternative in go) - `go install github.com/magefile/mage`
- [Mysql](https://mariadb.com/kb/en/mysql-command-line-client/) and [Postgres](https://www.postgresql.org/docs/13/app-psql.html) CLI Tools (mysql, mysqldump, psql, pgdump etc.)
- Additional tooling/CLIs - `mage GetDependencies`
#### Optional
- [NodeJS](https://nodejs.org/en/) and [Yarn](https://yarnpkg.com/) (for working with the documentation in /docs)
- [Git LFS](https://git-lfs.github.com/) (to get assets resolved in /docs)
- [Docker + docker-compose](https://www.docker.com/) for running additional runtime dependencies

## Prepare dev database
To run go-cf-api requires a migrated/existing Cloud Controller database.
To help with this, there is a [docker-compose.yml](./docker-compose.yml) file to start the DB instances and [SQL dumps](./database_dumps) from an existing `cloud_controller_ng` that can be imported to create a DB for testing.

> There are `mage` commands for performing most common database operations. Run `mage help` to see the full list of available commands.

First start needed runtime components with:
```bash
docker compose up -d

[+] Running 3/0
 ⠿ Container go-cf-api_mariadb_1   Running                                                                                                                                                                                                                                                                   0.0s
 ⠿ Container go-cf-api_postgres_1  Running                                                                                                                                                                                                                                                                   0.0s
 ⠿ Container go-cf-api_uaa_1       Running
```
To ensure the DB is ready for use, it needs creating and loading with the relevant database dump.
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

### Useful database operations
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

To avoid having to type out tags every time, you can instead run:
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

## Starting the go-cf-api
* Postgres
	```bash
	go run --tags=psql cmd/main.go config_psql.yaml
	```
* MySQL/MariaDB
	```bash
	go run --tags=mysql cmd/main.go config_mysql.yaml
	```

The default values in the `config_{db}.yml` and `sqlboiler_{db}.toml` files should match the credentials for each database in `docker-compose.yml`.
You may need to modify these credentials to connect to an alternative DB setup.

The API should be accessible at
```
http://localhost:8080/v3
```

The API Documentation should be accessible at
```
http://localhost:8080/docs/v3
```

## Building the go-cf-api binaries

The two binaries (for `psql` and `mysql`) can then be compiled for your current OS/architecture with:
```bash
mage build
```
Binaries are output to the `/build` directory.

The binaries can be compiled for other architectures by exporting the `GOOS` and `GOARCH` environment variables.

## List of API endpoints
Following endpoints have implemented and should be equivalent to the endpoints from `cloud_controller_ng`.

#### go-cf-api Specific Endpoints
- http://localhost:8080/healhz (Health Endpoint)
- http://localhost:8080/metrics (Prometheus Metrics)
- http://localhost:8080/docs/v3 (API Documentation)
#### V3 API
- http://localhost:8080/
- http://localhost:8080/v3
- http://localhost:8080/v3/info
- http://localhost:8080/v3/buildpacks (GET)
- http://localhost:8080/v3/buildpacks/:guid (GET, POST)
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

## Documentation

The documentation is based on [docusaurus](https://docusaurus.io/) and can be found at https://pages.github.com/cloudfoundry/go-cf-api/.
It is defined under `/docs` in this project and is then served by github-pages from the `gh-pages` branch which is detached from the other branches.
A GitHub workflow ensures the `gh-pages` branch always reflects the state of the main branch's `/docs` folder.

Package documentation is auto generated using [gomarkdoc](https://github.com/princjef/gomarkdoc) and integrated into [docusaurus](https://docusaurus.io/), as is the API documentation which is generated from code comments via [swag](https://github.com/swaggo/swag).
Additionally, the `go-cf-api` binaries/`go run cmd/main.go` will also serve the versioned API documentation in a Swagger UI at:
```
http://localhost:8080/docs/v3
```

Some manual documentation such as project ADRs (architecture decision records) are included under `/doc/adrs` and `/docs/docs`.
The `/docs/docs/intro.md` file is a symlink to this `README.md`.
For further information how to build the documentation locally refer to the [Docs README](https://github.com/cloudfoundry/go-cf-api/blob/main/docs/README.md)

## Development tasks

Below is a summary of the most important development workflows and tools.

### Generating DB models
To support different databases (Postgres and MySQL/MariaDB), we generate two sets of database models with [sqlboiler](https://github.com/volatiletech/sqlboiler).
These are then included in the same package and given build tag comments to control when to compile each model set into the binary.
Two binaries are then produced - one for each database.
Attempting to use e.g. the `mysql` binary on a Postgres database will result in a startup error.

To facilitate the generating, combining and build tagging of the `sqlboiler` models, there is a mage task to automate it.
Firstly one must have a `mysql` and `psql` instance running with a desired database schema already on them(in the `ccdb` database). Then one can run:
```bash
mage GenerateSQLBoiler
```
This will scan both databases schemas and generate go code to `/internal/sotrage/db/models/` as well as add the relevant build tag comments to compile the correct models for each supported DB.

### Running the linter
This project uses [golangci-lint](https://golangci-lint.run/) to ensure code is formatted correctly and return values are checked etc.
To run the linter, run:
```
golangci-lint run --build-tags psql,unit,integration
```

There is a GitHub Action that runs the linter on every push. No code should be merged until it passes all the linter checks.

### Running tests
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