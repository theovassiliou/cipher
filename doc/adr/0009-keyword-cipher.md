# 9. Keyword cipher

Date: 2022-01-05

## Status

Accepted

Implements [2. Use the ciphers via command line](0002-use-the-ciphers-via-command-line.md)

## Context

A keyword cipher is a monoalphabetic cipherer where the key alphabet is created based on a keyword and a base alphabet

First, the keyword will be stripped to contain unique characters, only.
Second, the base alphabet will be appended with no characters from the keyword.

 // Example:
 // Input: ASECRETKEYWORD, StdUppercaseAlphabet
 // ASECRTKYWODBFGHIJLMNPQRUVXYZ

## Decision

- Implement keyword cipherer for arbitrary alphabets
- Implement a shift/rotate right algorithm for arbitrary alphabets
- Standard alphabet provided is the Latin upper case characters.
- No preprocessing of the keyword, i.e. no capitalization or the lie

## Consequences

None.
