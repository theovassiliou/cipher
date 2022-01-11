# Cipher

This package is a playground to experiment with some text ciphering. It includes a library for working with monoalphabetic ciphers as well as a command-line interface to directly use different ciphers.

## Installation

Go version 1.17+

    go get github.com/theovassiliou/cipher

## How to use the library

    import  "github.com/theovassiliou/cipher"

    func main() {
        cipher := cipher.NewStdCipher(StdUppercaseAlphabet, RotateUTF8(3, StdUppercaseAlphabet))
        cleartext := "HELLO WORLD"
        encryptedText := cipher.Cipher(cleartext)
        decryptedText := cipher.Decipher(encryptedText)

        // ABCDEFGHIJKLMNOPQRSTUVWXYZ
        fmt.Println(StdUppercaseAlphabet)
        // DEFGHIJKLMNOPQRSTUVWXYZABC
        fmt.Println(RotateUTF8(3, StdUppercaseAlphabet))
        // HELLO WORLD
        fmt.Println(cleartext)
        // KHOOR ZRUOG
        fmt.Println(encryptedText)
        // HELLO WORLD
        fmt.Println(decryptedText)
    }

The above example creates a standard monoalphabetic cipher and ciphers the string "HELLO WORLD" and deciphers the result.

The key alphabet is created by rotating the plaintext alphabet by 3 characters.

## Command Line Tool Install

To use the ciphering via command-line build and install the command-line tool.

    go install ./application/ciphertool

and make sure you have `$GOPATH/bin` in your `$PATH`

This creates an executable named `ciphertool` you can call directly.

```shell
$ ciphertool --help                  
chiphertool is CLI tool to cipher and decipher text using
        a set of different cipher algorithms.

Usage:
  ciphertool [command]

Available Commands:
  cipher      ciphers a text
  completion  generate the autocompletion script for the specified shell
  decipher    deciphers a ciphertext
  help        Help about any command

Flags:
  -c, --cipher string     name of the cipher and rotation. One of [reverse keyword rotation caesar] (default "rotation:3")
      --config string     config file (default is $HOME/.ciphertool.yaml)
  -f, --filename string   input filename
  -h, --help              help for ciphertool
      --raw               do not preprocess input string or keywords
  -t, --toggle            Help message for toggle

Use "ciphertool [command] --help" for more information about a command.
```

### Usage

Cipher reads a text from stdin and ciphers or deciphers the provided text to stdout. In the case of ciphering the text is
expected to be the plain text, in case of deciphering the ciphertext.

As [default plain alphabet](adr/0005-normalization-of-input-keyword-charakters.md) the Latin StdUppercaseAlphabet will be used, i.e. "ABCD...XYZ". Therefore a provided plain text will be first capitalized, and then encoded. The same will also be applied to any passed keyword.
As a result, the ciphertext will also contain only upper-case letters.

This can be changed via options.

The following examples show how to cipher a text. The standard cipher is the "Ceasar" cipher, aka Rotate the plaintext alphabet by 3 characters.

#### To cipher a passed text

    $ ciphertool cipher "HELLO WORLD"
    KHOOR ZRUOG

#### To decipher a passed text

    $ ciphertool decipher "KHOOR ZRUOG"
    HELLO WORLD

#### Ciphering from stdin

Echos the string `HELLO WORLD` via `stdin` to `ciphertool`

    echo "HELLO WORLD" | ciphertool cipher
    KHOOR ZRUOG

##### Ciphering from file

Creates a file `test.txt` and ciphers the text

    echo "HELLO WORLD" > test.txt && ciphertool cipher -f test.txt

Details on options for the CLI at [Usage of Cipher](doc/UsageOfCipher.md)

## Features

* Full UTF-8 support for alphabets and ciphertexts
* Reading and writing from stdin/stdout and files

Implemented ciphers are

* Rotation (Right and Left)
* Caesar (default)
* Reverse
* Keyword

## Definitions

The documentation uses the following terms and definitions

    plaintext - the readable text 
    ciphertext - the unreadable text
    [to] cipher - the action to transform a plaintext into a ciphertext
    [to] decipher - the action to transform ciphertext back to plaintext
    plain alphabet - the ordered sequence of characters (or runes in go) of which the plaintext is build
    key alphabet - the ordered sequence of characters (or runes in go) of which the ciphertext is build

## Outlook

As this project is a finger exercise occassionally we will extend the functionality.

Most importantently I would like to add support for polyalphabetic ciphers,both in the library and in the CLI.

Other ideas are to add a more user-interactive CLI that also supports a user to decipher a given ciphertext. No, I do not think that we will suport "breaking" ciphers, but to offer some support. How this could look like I do not know yet.

An additional idea is to expose the functionality also as a web-service via a REST API. Most probably this will be made available in a seperate repo. 
## History and Motivation

Motivated by an audio play ("Die Drei Frageezeichen - Folge 213 Der Fluch der Medusa") I wanted to experiment with ciphering and deciphering as a finger exercise in golang.
