package docs

//go:generate sh -c "cd ../../../ && ln -s cmd/main.go main.go"
//go:generate sh -c "cd ../../../ && swag init -o ./internal/api/docs --parseInternal --parseDepth 1 --parseDependency --parseVendor"
//go:generate sh -c "cd ../../../ && unlink main.go && cp -f ./internal/api/docs/swagger.yaml ./docs/swagger.yaml"
