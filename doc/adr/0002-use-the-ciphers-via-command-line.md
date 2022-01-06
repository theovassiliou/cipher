# 2. Use the ciphers via command line

Date: 2021-12-14

## Status

Accepted

Options [4. Stdin as Default input](0004-stdin-as-default-input.md)

Options [5. Normalization of input / keyword charakters](0005-normalization-of-input-keyword-charakters.md)

Implements [8. Rotation Cipher](0008-rotation-cipher.md)

Options [8. Rotation Cipher](0008-rotation-cipher.md)

Options [9. Keyword cipher](0009-keyword-cipher.md)

## Context

We have created a framework for developing text ciphers. To use them we would like to use them
via the command line. Thus a CLI interface to the existing ciphers should be made available.

## Decision

Implement a CLI that makes all available ciphers available via the command line.
As CLI framework we will use [Cobra](https://github.com/spf13/cobra)

## Consequences

By using golang we can support out-of-the-box multiple OS.
While we only have a macOS environment available for implementation we should consider future design decisions to be supported on multiple OSs.
