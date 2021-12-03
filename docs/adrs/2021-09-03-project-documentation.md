---
slug: ADR - Documentation Framework
title: Use Docusaurus to document our project
authors:
- name: Florian Braun
  title: go-cf-api Team
  url: https://github.com/FloThinksPi
  image_url: https://avatars1.githubusercontent.com/u/5863788?v=4
tags: [adr]
---

# Use [client_golang](https://github.com/prometheus/client_golang) prometheus library to expose metrics
* Status: proposed
* Deciders: [Florian Braun](https://github.com/FloThinksPi)
* Date: 03.09.21

## Context and Problem Statement

We want a documentation page in which we provide documentation from different sources:
* Manually written documentation
* GoDoc generated documentation from go code
* API Documentation from a swagger/openapi spec
* ADRs are nicely presented

We want to have a nice page for it as when going public with this project, a good (single)entry point for people to look into the project is essential.

## Decision Drivers

* Having the ability to have different sources of documentation, generated and manually written ones.
* Being able to statically deliver the documentation (e.g. via GitHub Pages)
* Nice Look
* Easy to write docs in
* Being able to use Continuous Delivery to serve the docs.

## Considered Options

* [Docusaurus](https://docusaurus.io/)
* [Jekyll](https://jekyllrb.com/)
* [mkdocs](https://www.mkdocs.org/)

## Decision Outcome

Chosen option: "[Docusaurus](https://docusaurus.io/)", because it was easy to use, produces very nice results, has the most features and seems to be widely used.

## Pros and Cons of the Options <!-- optional -->

### [Docusaurus](https://docusaurus.io/)

* Good, because it has a modern/nice look.
* Good, because it has a plugin system which extends functions and delivers e.g. swagger/openapi support.
* Good, because it is based on react and one is able to integrate own complex pages with that if one wants.
* Good, because it has a build in deploy mechanism to deploy into a second branch on the same project.
* Good, because it brings nice to have / advanced features for the future like versioned docs or multi-language support.

### [Jekyll](https://jekyllrb.com/)

* Good, because it is integrated into GitHub Enterprise. They will generate the page regularly for you.
* Bad, because it has few functionality. Hard to get e.g. Swagger/OpenAPI Docs to Display.
* Bad, because i diskliked the themes/page design in general.

### [mkdocs](https://www.mkdocs.org/)

* Good, because it has a plugin system which extends functions and delivers e.g. swagger/openapi support.
* Bad, because it doesn't offer an implementation of collectors to register multiple metrics at the same time
* Bad, because i diskliked the themes/page design in general.
