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


# TODO
- ~~Validate Connection String~~
- ~~Make Mysql Work~~
- ~~DB Migration~~
- Get Rid of Viper
- Better SQL Logging
- Store Memory instances in single Storage 
- Own Metrics in Prometheus endpoint
- Logging Redact Secrets (From SQL,Echo)
- Implement Controller Pattern
- Basic UI
- Rate Limits
- Tests
- GoDocs
