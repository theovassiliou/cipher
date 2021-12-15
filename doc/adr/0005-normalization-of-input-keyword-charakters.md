# 5. Normalization of input / keyword charakters

Date: 2021-12-15

## Status

Accepted

Implements [2. Use the ciphers via command line](0002-use-the-ciphers-via-command-line.md)

## Context

In ciphering typically either lower case or upper case alphabets are being used for the plain and key alphabet. While there is no technical reason to limit this, this is often observed.

## Decision

As default plain alphabet all upper-case Latin characters (ABCD...XYZ) will be used

## Consequences

- It is required to give the possibility to use different defaults (only lower case or mixed case alphabets)
- It is unclear what this decision would mean in the case of non-Latin characters (Asian, greek, Hebrew, etc)
