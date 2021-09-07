# CloudGontroller Documentation

This Documentation is built using [Docusaurus 2](https://docusaurus.io/), a modern static website generator.

# Write Content

### ADRs
Write adrs in the `adrs` subfolder. See Readme in that folder.

### Self Written Docs
Write docs in the `docs` subfolder.

# Run/Deploy Docusaurus
### Installation

From within this folder run
```console
brew install nodejs yarn
yarn install
```
to installs all dependencies

### Local Development

```console
yarn start --locale en
yarn start --locale de
```
This command starts a local development server and opens up a browser window. Most changes are reflected live without having to restart the server.

### Build

```console
yarn build
```

This command generates static content into the `build` directory and can be served using any static contents hosting service.

### Deployment

```console
GIT_USER=i507599 GIT_PASS=$pass GITHUB_HOST=github.tools.sap yarn deploy
```
This overwrites the `gh-pages` branch with a single commit delivering this page.