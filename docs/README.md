# GO-CF-API Documentation

This Documentation is built using [Docusaurus 2](https://docusaurus.io/), a modern static website generator.
It is build and pushed to the `gh-pages` branch by [a github action](https://github.com/cloudfoundry/go-cf-api/actions/workflows/deploy-docs.yml) automatically on every commit on master.

# Write Content

### ADRs
Write [ADRs](https://adr.github.io/madr/) in the `adrs` subfolder. See [ADR Readme](./adrs/README.md) for more information.

### Self Written Docs
Write docs in the `docs` subfolder as markdown.

# Run/Deploy Docusaurus
### Installation

Prerequisite is you have nodejs 16 installed and [yarn](https://classic.yarnpkg.com/lang/en/docs/install/).

From within this folder run
```console
yarn install
```
to installs all dependencies

### Local Development
Firstly we want to generate all automatic created documentation(api docs and go package docs)
```console
cd ..
mage GetDependencies
mage GenerateDocs
```
Then serve them locally.
```console
cd docs
yarn start --locale en
yarn start --locale de
```
This command starts a local development server and opens up a browser window. 
Changes are rendered live without having to restart the server.

### Build

```console
cd ..
mage GetDependencies
mage GenerateDocs
cd docs
yarn build
```

This command generates static content into the `build` directory and can be served using any static contents hosting service.

### Deployment

```console
cd ..
mage GetDependencies
mage GenerateDocs
cd docs
read -s ghtoken
GIT_USER="MYUSER" GIT_PASS=$ghtoken yarn deploy
```
This overwrites the `gh-pages` branch with a single commit delivering this page.
This is usually not needed since we have a [GitHub action](https://github.com/cloudfoundry/go-cf-api/actions/workflows/deploy-docs.yml) in place that does this automatically.