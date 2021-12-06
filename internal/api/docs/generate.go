package docs

//go:generate sh -c "cd ../../../ && ln -s cmd/main.go main.go  && swag init -o ./internal/api/docs --parseInternal --parseDepth 1 --parseDependency --parseVendor && unlink main.go && cp -f ./internal/api/docs/swagger.yaml ./docs/swagger.yaml "
