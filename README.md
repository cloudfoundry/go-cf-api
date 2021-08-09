![lint](https://github.tools.sap/cloudfoundry/cloudgontroller/workflows/Lint/badge.svg) ![unit tests](https://github.tools.sap/cloudfoundry/cloudgontroller/workflows/Run%20unit%20tests/badge.svg) ![db tests](https://github.tools.sap/cloudfoundry/cloudgontroller/workflows/Run%20database%20tests/badge.svg) ![build](https://github.tools.sap/cloudfoundry/cloudgontroller/workflows/Build%20binaries/badge.svg)
# cloudgontroller
A replacement for [cloud_controller_ng](https://github.com/cloudfoundry/cloud_controller_ng), written in Go
## Prerequisites
- [Go](https://golang.org/dl) version 1.16
	```bash
	// Use Go version manager (GVM) or
	brew install go@1.16
	```
- Mysql and Postgres CLI Tools (mysql, mysqldump, psql, pgdump etc.)
	```bash
	brew install mysql postgres
	```
- [Mage](https://github.com/magefile/mage) (Makefile alternative in go)
	```bash
	go install github.com/magefile/mage
	```
- Other command line tools:
	```bash
	mage install
	```

## Prepare dev database

To run cloudgontroller we need a proper Cloud Controller database.
To do this we have a docker-compose file to create the DBs and SQL dumps from an existing Cloud Controller that can be used to build an DB for testing.
There are `mage` commands for most database operations.

Run `mage -l` to see the full list of available commands.

### Postgres
```bash
mage DBStart config_psql.yaml
mage DBCreate config_psql.yaml
mage DBLoad config_psql.yaml database_dumps/3.102.0_psql_ccdb.sql
```

### MySQL/MariaDB
```bash
mage DBStart config_mysql.yaml
mage DBCreate config_mysql.yaml
mage DBLoad config_mysql.yaml database_dumps/3.102.0_mysql_ccdb.sql
```

### Misc database operations
* Drop database
	```bash
	mage DBDelete config_psql.yaml
	```
* Drop and recreate (empty) database
	```bash
	mage DBRecreate config_psql.yaml
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

## Starting the Cloud Controller
* Postgres
	```bash
	go run --tags=psql cmd/main.go config_psql.yaml
	```
* MySQL/MariaDB
	```bash
	go run --tags=mysql cmd/main.go config_mysql.yaml
	```

The default values in the `config_{db}.yml` and `sqlboiler_{db}.toml` files should match the credentials for each database in `docker-compose-dev.yaml`.

The API should be accessible at e.g.
```
http://localhost:8080/v3/buildpacks
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
Documentation is auto generated using `godoc` and `godoc-static` and served via GitHub pages. The binaries compiled by the `mage build` command will also serve the documentation at:
```
http://localhost:8080/docs/v3
```


## Running linter
This project uses [golangci-lint](https://golangci-lint.run/) to ensure code is formatted correctly and return values are checked etc.
To run the linter, run:
```
golangci-lint run --build-tags psql,unit
```

There is a GitHub Action that runs the linter on every push. No code should be pushed until it passes the linter.

## Running tests
Different tags are used to control which tests are run:
* Unit tests
	```bash
	go test -tags="unit,psql" ./...
	```
* DB tests (Postgres) - require running database
	```bash
	go test -tags="db,psql" ./internal/app/cloudgontroller/sqlboiler -test.config sqlboiler_psql.toml
	```
* DB tests (MySQL/MariaDB) - require running database
	```bash
	go test -tags="db,mysql" ./internal/app/cloudgontroller/sqlboiler -test.config sqlboiler_mysql.toml
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

# Contributing
## Jwt authentication
Uaa container gets started automatically with docker-compose up. When uaa container is up you can use uaac to get a valid jwt token which is signed with the uaa private key, written down in the uaa.yml

```
uaac target http://localhost:8095
uaac token client get admin -s adminsecret
uaac contexts
```
To access restricted endpoint, a curl would look like:
```
curl -i http://localhost:8080/api/v3/buildpacks -H "Authorization: Bearer [Add token from uaac context here]"
```

This is handled by the echo framework buildin middleware. To make an endpoint only usable when authenticated:
Example from v3/handlers.go
```
    var publicKey, _ = jwt.ParseRSAPublicKeyFromPEM([]byte(uaakey))
    // Restricted group
    r := e.Group(prefix)
    {
	config := middleware.JWTConfig{
		SigningKey: publicKey,
		SigningMethod: "RS256",
	}
	r.Use(middleware.JWTWithConfig(config))

	// Buildpacks
	r.GET(fmt.Sprintf("/buildpacks"), controllers.GetBuildpacks)
	r.GET(fmt.Sprintf("/buildpacks/:guid"), controllers.GetBuildpack)
    }
}
```


Accessing the userdata from the context during a request:
```
func GetBuildpacks(c echo.Context) error {
    var user = c.Get("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    name := claims["client_id"].(string)

    fmt.Println(user) //contains the token itself
    fmt.Println(name) //contains the username read from the client_id claim (can be anything from the jwt token middle part)
```


