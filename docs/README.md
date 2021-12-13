# Documentation

This documentation is built using [Docusaurus 2](https://docusaurus.io/) - a static website generator.
It is built and pushed to the `gh-pages` branch by [a GitHub Actions workflow](https://github.com/cloudfoundry/go-cf-api/actions/workflows/deploy-docs.yml) automatically on every commit to `main`

## Adding content

### ADRs
[ADRs](https://adr.github.io/madr/) are located in the `adrs` directory. See [ADR Readme](./adrs/README.md) for more information.

### Manually maintained documentation
Manually maintained docs are located in the `docs` directory in markdown format.

## Running/deploying Docusaurus
### Installation

Docusaurus requires [Node.js 16](https://nodejs.org/en/) and [yarn](https://classic.yarnpkg.com/lang/en/docs/install/)

To install all the dependencies, from within the `docs` folder run
```console
yarn install
```

### Local Development
To generate all the automated documentation (such as API docs and Go package docs)
```console
cd ..
mage GetDependencies
mage GenerateDocs
```
To serve the docs locally
```console
cd docs
yarn start --locale en
yarn start --locale de
```
This starts a local development server and opens up a new browser window.
Changes are rendered live without having to restart the server.

### Build

To generate the static content into the `build` directory and which can then be served using a static hosting service
```console
cd ..
mage GetDependencies
mage GenerateDocs
cd docs
yarn build
```

### Deployment to GitHub Pages

To deploy the docs to GitHub Pages
```console
cd ..
mage GetDependencies
mage GenerateDocs
cd docs
read -s ghtoken
GIT_USER="MYUSER" GIT_PASS=$ghtoken yarn deploy
```
This overwrites the `gh-pages` branch with a single commit containing all the docs content.
> Running these steps manually should not be necessary as there is a [GitHub Actions workflow](https://github.com/cloudfoundry/go-cf-api/actions/workflows/deploy-docs.yml) to do this automatically on every commit to `main`