# Gopilot Webserver
## Prerequisits
- Go 1.16 minimum: https://golang.org/dl/ (Install with GVM)
- Mage (Makefile alternative in go): https://github.com/magefile/mage


Install the rest of the used CLI Dependencies in this project
```
mage InstallDeps
```

All compile time dependencies will be downloaded when compiling (If you use the mage commands)

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
