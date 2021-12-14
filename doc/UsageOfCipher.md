# Usage of cipher

Cipher reads a text from stdin and qbbciphers or deciphers the provided text to stdout. In case of ciphering the text is
expected to be the plain text, in case of deciphering the cipher text.

As default plain alphabet the latin StdUppercaseAlphabet will be used, i.e. "ABCD...XYZ". Therefore a provided
plain text will be first capitalized, and then encoded. The same will also be applied to any passed keyword.
As a result the ciphertext will also contain only upper-case letters.

This can be changed via options.

## Options

For the [CLI](adr/0002-use-the-ciphers-via-command-line.md) the following options are available.

### cipher

cipher [text]

### cipher --decipher

cipher --decipher [text]

### cipher --cipher

--cipher options defines which cipher to use. Currently the following ciphers are supported

cipher --cipher rotation:n
key alphabet is build by left rotating plainalphabet by n characters

cipher --cipher rotation:-n
key alphabet is build by right rotating plainalphabet by n characters

cipher --cipher caesar*
aquivalent to cipher --cipher rotation:3

cipher --cipher reverse
key alphabet is build by reversing the order of characters of plain alphabet

cipher --cipher keyword:WEISSKOPFSEEADLER
key alphabet is build as WEISKOPFADLRBCGHJMNQTUVXYZ

### cipher --plainalphabet --keyalphabet

### ciper --raw

Plaintext and keywords are not preprocessed, i.e. capitalized.

### ciper --strip

Strips any characters from the plaintext not included in the plainalphabet, before encoding.

### cipher --group n

ciphertext is split by group of n characters. Implies --strip

## Definitions

See [README.md](../README.md#Definitions)
