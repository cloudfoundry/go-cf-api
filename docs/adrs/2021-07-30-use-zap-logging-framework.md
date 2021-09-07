---
slug: ADR - Logging Framework
title: Use Zap logging as a performant and customizable structured logging framework
authors:
  - name: Florian Braun
    title: CloudGontroller Team
    url: https://github.com/FloThinksPi
    image_url: https://avatars1.githubusercontent.com/u/5863788?v=4
tags: [adr]
---

# Use [Zap](https://github.com/uber-go/zap) logging as a performant and customizable structured logging framework

* Status: accepted
* Deciders: [Florian Braun](https://github.com/FloThinksPi)
* Date: 30.7.21

## Context and Problem Statement

We want our PoC to log various data in a structured way so that it is parsable by e.g. logsearch.

## Decision Drivers

* Fast logging (no wasted cpu time)
* Customizable to make it compatible with the way log parsing in cloudfoundry works and maybe be able to mirror the log schema of the Cloud_Controller_NG if wanted.
* Community support, solutions should be a broadly adapted technology.

## Considered Options

* [logrus](https://github.com/sirupsen/logrus)
* [zap](https://github.com/uber-go/zap)
* [apex](https://github.com/apex/log)
* [zerolog](https://github.com/rs/zerolog)

## Decision Outcome

Chosen option: "", because it has bradly used, very fast and has all features we want.
As we started with zap, zerolog would be an alternative but we did not think it was worth it switching over as there was no obvious benefit.
So we sticked to use zap.

## Pros and Cons of the Options <!-- optional -->

### [logrus](https://github.com/sirupsen/logrus)

* Good, because has all features we like
* Bad, because dead project, recommends alternatives
* Bad, because bad performance

### [zap](https://github.com/uber-go/zap)

* Good, because very good performance,
* Good, because broadly used
* Good, because has all features we like

### [apex](https://github.com/apex/log)

* Good, because has all features we like
* Bad, because I dont like the interface
* Bad, because it is not used by as much people as zap
* Bad, because it is as slow as logrus

### [zerolog](https://github.com/rs/zerolog)

* Good, because its also very fast.
* Good, because has all features we like
* Good, because also broadly used but a bit less than zap