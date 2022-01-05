# 7. Raw input

Date: 2021-12-17

## Status

Accepted

Refines [5. Normalization of input / keyword charakters](0005-normalization-of-input-keyword-charakters.md)

## Context

[5. Normalization of input/keyword charakters](0005-normalization-of-input-keyword-charakters.md) requires the input to be preprocessed for default handling of input, concerning the plain alphabet, which is LatinUppercase alphabet. In cases where preprocessing is not adequate or wanted a way to avoid preprocessing is required.

## Decision

Offer a `--raw` option to avoid preprocessing.

## Consequences

None.