### sqlboiler templates

These templates are heavily modified versions of the ones found in [the sqlboiler repo](https://github.com/volatiletech/sqlboiler).

The significant changes are:
* Extract interfaces for each type of operation (e.g. `Inserter`, `Finisher`)
  and give all the concrete methods a struct receiver so we have an implementor
  if these interfaces. This allows us to mock these interfaces in tests without
  needing a DB
* Add a `Quote` helper to easily quote values for both MySQL and Postgres
* Add a `Subquery` helper to allow for building of subqueries in the same way as
  queries
