# 3. Uppercase Standard Plain Alphabet

Date: 2021-12-14

## Status

Accepted

## Context

When dealing with cipher the ciphertext is built by mapping characters from the plaintext via the plain alphabet to one or more key alphabets. The ciphertext then consists of characters from the key alphabet.

## Decision

As the Standard Plain Text Alphabet, we will use the Latin upper case characters, i.e. "ABCD...XYZ"

## Consequences

Each cipher has to be able to cope with plaintext characters that are not present in the plain alphabet. Handling can be "ignoring", or using "as is". In case of ignoring deciphering is not 100% reversible.
