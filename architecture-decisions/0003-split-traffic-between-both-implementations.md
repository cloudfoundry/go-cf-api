# Deploy [HAProxy](https://www.haproxy.org/) to route specific endpoints/methods to the new implementation

* Status: accepted
* Deciders: Andy Paine, Sven Krieger, Philipp Thun
* Date: 2021-08-04

## Context and Problem Statement

The existing Cloud Controller is the reference implementation of the CF v3 API.
This project should avoid replacing the entire `cloud_controller_ng` project in a single "big bang" migration.
In order to build the new implementation iteratively, this project should be deployable in parallel with the existing implementation.
It should be possible to route individual API calls to the new Cloud Controller as soon as each endpoint is complete.

## Decision Drivers <!-- optional -->

* Discover bugs early
* Deliver value from reimplementation quickly
* Minimise mean-time-to-recovery when bugs are discovered

## Considered Options

1. Deploy new CC alongside existing CC in same instance group, using `nginx` for routing
1. Complete an entire endpoint at once (all HTTP methods) and use `gorouter` for path based routing
1. Deploy a dedicated path and HTTP method based router/proxy in front of old and new implementations and split traffic based on that

## Decision Outcome

Chosen option: 3 (deploy a path and HTTP method based router/proxy in front of old and new implementations and split traffic based on that),
because it is the only option that allows for separate scaling of old and new implementations as well as routing based on HTTP method and path.
Following images shows a rough routing example with just the verry specific `GET /v3/buildpacks/:guid` endpoint beeing routed to the go implementation. Everything else will be routed to the cloudcontroller_ng. 
![image](https://media.github.tools.sap/user/9608/files/9fd7fd80-fc39-11eb-934c-0aff81887456)

### Positive Consequences <!-- optional -->

* New implementation can be built in small units (endpoint + HTTP method)
* Proxy can be registered only for certain routes, minimising throughput
* Networking such as TLS in new implementation can be delayed until closer to completion (as proxy can perform this function)
* Good support for HAProxy BOSH release as it is maintained by SAP team

### Negative Consequences <!-- optional -->

* Additional software to manage
* HAProxy BOSH release is not that well suited to this use case

## Pros and Cons of the Options <!-- optional -->

### Option 1 (Deploy new CC alongside existing CC in same instance group, using `nginx` for routing)

* Good, because does not change network architecture
* Good, because does not add any new VMs
* Bad, because cannot independently scale each implementation
* Bad, because would require changes to the CAPI release to support this use case

### Option 2 (Complete an entire endpoint at once (all HTTP methods) and use `gorouter` for path based routing)

* Good, because does not change network architecture
* Good, because can independently scale each implementation
* Bad, because requires large amount of work to complete a whole endpoint
* Bad, because requires new implementation to support TLS etc. for secure communication

## Links <!-- optional -->

* [HAProxy BOSH release](https://github.com/cloudfoundry-incubator/haproxy-boshrelease)
