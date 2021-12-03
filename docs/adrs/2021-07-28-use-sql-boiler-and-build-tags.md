---
slug: ADR - SQL Framework
title: Use sqlboiler together with build tags to support multiple databases
authors:
  - name: Florian Braun
    title: go-cf-api Team
    url: https://github.com/FloThinksPi
    image_url: https://avatars1.githubusercontent.com/u/5863788?v=4
  - name: Marc Misoch
    title: go-cf-api Team
    url: https://github.com/MMisoch
    image_url: https://avatars1.githubusercontent.com/u/47423110?v=4
  - name: Andrew Paine
    title: go-cf-api Team
    url: https://github.com/andy-paine
    image_url: https://avatars1.githubusercontent.com/u/14118619?v=4
tags: [adr]
---

# Use [sqlboiler](https://github.com/volatiletech/sqlboiler) together with build tags to support multiple databases

* Status: accepted
* Deciders: [Florian Braun](https://github.com/FloThinksPi), [Andy Paine](https://github.com/andy-paine), [Marc Misoch](https://github.com/mmisoch)
* Date: 2021-07-28

## Context and Problem Statement

The existing Cloud Controller supports both Postgres and MySQL as a storage backend and both options are used in real production deployments today.
In order to replace the existing Cloud Controller implementation, this project should also be able to support both

## Decision Drivers

* Community support: should be able to work with any CC DB
* Performance: want a minimal runtime overhead
* Performance: want control over SQL queries used to be able to optimise

## Considered Options

1. [sqlboiler](https://github.com/volatiletech/sqlboiler) with separate packages for Postgres and MySQL
1. [sqlboiler](https://github.com/volatiletech/sqlboiler) with an extracted interface that is implemented by both Postgres and MySQL
1. [sqlboiler](https://github.com/volatiletech/sqlboiler) in a shared package with different build tags for each implementation
1. [GORM](https://gorm.io/index.html)
1. [xo/xo](https://github.com/xo/xo)

## Decision Outcome

Chosen option: 3 ([sqlboiler](https://github.com/volatiletech/sqlboiler) in a shared package with different build tags for each implementation) because this allows all of the model code to be generated from an existing schema and lets the compiler type check that the generated implementations have the same signatures without needing to manipulate the generated code much.

### Positive Consequences

* Model code can be regularly regenerated when CC DB schema changes (due to new migrations in existing implementation)
* Smaller binaries as each only contains code relevant to that DB backend
* Compile time checks that the generated code has the same function signatures for both implementations (or at least for all functions that are actually used)
* Generated code can be easily extended to support optimisations

### Negative Consequences

* Extra tooling is required to rename and combine the files into a single package and add build tags
* Directory containing generated files is extremely large
* Developers need to supply build tags in order to browse, lint, compile and test code

## Pros and Cons of other options

### Option 1 ([sqlboiler](https://github.com/volatiletech/sqlboiler) with separate packages for Postgres and MySQL)

* Good, because `sqlboiler` can be run to generate the code without modification
* Bad, because all controller code would need to do an `if`/`switch` statement on database type
* Bad, because binaries will contain redundant code for other databases

### Option 2 ([sqlboiler](https://github.com/volatiletech/sqlboiler) with an extracted interface that is implemented by both Postgres and MySQL)

* Good, because the interface could abstract between the two database
* Good, because compiler would check that both generated implementations satisfy interface
* Bad, because requires significant effort to extract complex `sqlboiler` interfaces
* Bad, because some functions are static and cannot be extracted into an interface

### Option 4 [GORM](https://gorm.io/index.html)

* Good, because good documentation
* Good, because controller code that interacts with database models is easy to write
* Bad, because little/no support for generating models from existing schema
* Bad, because has runtime overhead of using reflection

### Option 5 [xo/xo](https://github.com/xo/xo)


* Good, because generated models are extremely simple
* Good, because templates are easy to customise
* Bad, because generated models have no support for eager loading