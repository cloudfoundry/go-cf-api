# Use [client_golang](https://github.com/prometheus/client_golang) prometheus library to expose metrics
* Status: accepted
* Deciders: Will Gant and Aftab Alam
* Date: 20.08.21

## Context and Problem Statement

We want our PoC to expose prometheus metrics to be consumed by a visualization dashbaord, e.g,. Grafana, and also to benchmark peformance against the exisitng Cloud Controller.

## Decision Drivers

* A metric system which can be easily extended with custom metrics to allow comparison with the performance of the existing Cloud Controller
* Community support, solutions should be a broadly adapted technology.

## Considered Options

* [client_golang](https://github.com/prometheus/client_golang)
* [VictoriaMetrics](https://github.com/VictoriaMetrics/metrics)
* [promenade](https://github.com/poblish/promenade)

## Decision Outcome

Chosen option: "client_golang", because it the official and by far most popular library, and appears to have most features.

## Pros and Cons of the Options <!-- optional -->

### [client_golang](https://github.com/prometheus/client_golang)

* Good, because it is the official library and has significant community support
* Good, because it has all features we wanted and has advanced functionality, e.g,. sqlDBStats out-of-the-box, easily customizable...

### [VictoriaMetrics](https://github.com/VictoriaMetrics/metrics)

* Good, because it depends on one exeternal package as compared to eight for client_golang
* Bad, because it has far fewer users and maintainers than client_golang
* Bad, because it is missing "advanced funcationality" from client_golang (their words)
* Bad, because it doesn't offer an implementation of collectors to register mulitple metrics at the same time

### [promenade](https://github.com/poblish/promenade)

* Bad, because it has client_golang as a dependency
* Bad, because it has only one author and one user
* Bad, because it doesn't offer an implementation of collectors to register mulitple metrics at the same time

