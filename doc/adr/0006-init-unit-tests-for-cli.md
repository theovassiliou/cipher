# 6. init Unit Tests for CLI

Date: 2021-12-16

## Status

Accepted

## Context

Unit tests are essential for the QA of SW and therefore should also be created and executed for CLIs. cipherer uses Cobra as a CLI framework. Using the golang test framework should be targeted for the unit testing of cipherer

## Decision

Create and maintain unit tests within the golang test framework
Organize tests around CLI commands

## Consequences

The cobra-based commands have to be restructured to be efficiently unit testable. Thus the default created templates can not be used as they are.

