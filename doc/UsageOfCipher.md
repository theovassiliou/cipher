# Usage of cipher

Cipher reads a text from stdin and ciphers or deciphers the provided text to stdout. In the case of ciphering the text is
expected to be the plain text, in case of deciphering the ciphertext.

As [default plain alphabet](adr/0005-normalization-of-input-keyword-charakters.md) the Latin StdUppercaseAlphabet will be used, i.e. "ABCD...XYZ". Therefore a provided plain text will be first capitalized, and then encoded. The same will also be applied to any passed keyword.
As a result, the ciphertext will also contain only upper-case letters.

This can be changed via options.

## Options

For the [CLI](adr/0002-use-the-ciphers-via-command-line.md) the following options are available.

### ciphertool cipher

Command to cipher a text

    ciphertool cipher [text]

### ciphertool decipher

Command to decipher a text

    ciphertool decipher [text]

### ciphertool [cipher|decipher] --cipher

`--cipher`  option defines which cipher to use. Currently, the following ciphers are supported

    ciphertool --cipher caesar*
equivalent to `cipher --cipher rotation:3`

    ciphertool --cipher rotation:n
key alphabet is built by [left rotating](adr/0008-rotation-cipher.md) plain alphabet by n characters

    ciphertool --cipher rotation:-n
key alphabet is built by [right rotating](adr/0008-rotation-cipher.md) plain alphabet by n characters

    ciphertool --cipher reverse
key alphabet is built by reversing the order of characters of the plain alphabet

    ciphertool --cipher keyword:WEISSKOPFSEEADLER
key alphabet is build as WEISKOPFADLRBCGHJMNQTUVXYZ

### Other options

    ciphertool --raw

Plaintext and keywords are not preprocessed, i.e. not capitalized.

## Not yet implemented options

From time to time we are updated ciphertool. For this we are keeping here backlog of options that we would like to implement in the future.

    ciphertool --plainalphabet --keyalphabet

To pass hardcoded the `plainalphabet` and the `keyalphabet`. In addition, it should be considered to extend  the standard plain alphabets (and keyalphabets)

    ciphertool --strip

Strip any characters from the plaintext not included in the plain alphabet, before encoding. The default behaviour is that any characters not included in the plain alphabet are just passed.

    ciphertool --group n

ciphertext is split by a group of n characters. Implies `--strip`. Rationale: In order to communicate a ciphertext it might be helpful to pretty-print the output. Implies `--strip` as WS and CR/NL would impact the output.

## Definitions

See [README.md](../README.md#Definitions)
