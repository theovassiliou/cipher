# 8. Rotation Cipher

Date: 2021-12-17

## Status

Accepted

Implements [2. Use the ciphers via command line](0002-use-the-ciphers-via-command-line.md)

## Context

A rotation (or shift) ciphers is a monoalphabetic cipherer where the key alphabet is created by rotating (shifting) the plain alphabet by a specified amount of characters.

Example: Shift Rotate the UpperCase Latin Alphabet by N=5 characters left

Plain Alphabet: ABCDEFGHIJKLMNOPQRSTUVWXYZ
  Key Alphabet: FGHIJKLMNOPQRSTUVWXYZABCDE

A special case for a shift left alphabet is the Caesars ciphers, with N=3.

## Decision

- Implement a shift/rotate left algorithm for arbitrary alphabets
- Implement a shift/rotate right algorithm for arbitrary alphabets
- Standard alphabet provided is the Latin upper case characters.

## Consequences

What becomes easier or more difficult to do and any risks introduced by the change that will need to be mitigated.
