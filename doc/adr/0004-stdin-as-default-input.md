# 4. Stdin as Default input

Date: 2021-12-14

## Status

Accepted

Implements [2. Use the ciphers via command line](0002-use-the-ciphers-via-command-line.md)

## Context

The CLI has to read the plain text from somewhere. We anticipate the usage of the CLI tool in a tool chain.

## Decision

CLI tool shall read per default the plain text (when ciphering) or the cipher text (when deciphering) from stdin.
In addition we want to offer the possibility to read pass it as parameter.

Another optional possibility should be to read the plain/cipher text from a file.

## Consequences

None that I am aware of.
