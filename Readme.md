# Gopilot Webserver
## Prerequisits
- Go 1.16 minimum: https://golang.org/dl/
- NodeJS 14.17.0 minimum: https://nodejs.org/en/download/
- Yarn (Any version should do):  https://classic.yarnpkg.com/en/docs/install
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
