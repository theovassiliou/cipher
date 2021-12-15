# Usage of cipher

Cipher reads a text from stdin and qbbciphers or deciphers the provided text to stdout. In case of ciphering the text is
expected to be the plain text, in case of deciphering the cipher text.

As default plain alphabet the latin StdUppercaseAlphabet will be used, i.e. "ABCD...XYZ". Therefore a provided
plain text will be first capitalized, and then encoded. The same will also be applied to any passed keyword.
As a result the ciphertext will also contain only upper-case letters.

This can be changed via options.

## Options

For the [CLI](adr/0002-use-the-ciphers-via-command-line.md) the following options are available.

### ciphertool cipher

ciphertool cipher [text]

### ciphertool --decipher

cipher decipher [text]
### ciphertool [cipher|decipher] --cipher

--cipher options defines which cipher to use. Currently the following ciphers are supported

ciphertool --cipher rotation:n
key alphabet is build by left rotating plainalphabet by n characters

ciphertool --cipher rotation:-n
key alphabet is build by right rotating plainalphabet by n characters

ciphertool --cipher caesar*
aquivalent to cipher --cipher rotation:3

ciphertool --cipher reverse
key alphabet is build by reversing the order of characters of plain alphabet

ciphertool --cipher keyword:WEISSKOPFSEEADLER
key alphabet is build as WEISKOPFADLRBCGHJMNQTUVXYZ

### ciphertool --plainalphabet --keyalphabet

### ciphertool --raw

Plaintext and keywords are not preprocessed, i.e. capitalized.

### ciphertool --strip

Strips any characters from the plaintext not included in the plainalphabet, before encoding.

### ciphertool --group n

ciphertext is split by group of n characters. Implies --strip

## Definitions

See [README.md](../README.md#Definitions)
