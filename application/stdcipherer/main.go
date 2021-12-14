package main

import (
	"fmt"

	"github.com/theovassiliou/cipher"
)

func main() {

	// Cesar Cipher usees a ROT-3 key alphabet
	cc := cipher.NewCaesarCipher(cipher.StdUppercaseAlphabet)
	fmt.Println("Encrypted Text: ", cc.Cipher("HELLO WORLD"))
	fmt.Println("Decrypted Text: ", cc.Decipher("KHOOR ZRUOG"))
	fmt.Println("KeyAlphabet: ", cc.KeyAlphabet())
	fmt.Println("Name: ", cc.Name())
	fmt.Println("Description: ", cc.Description())

	// Create a ROT-3 cipher.
	cc = cipher.NewStdCipher(cipher.StdUppercaseAlphabet, cipher.RotateUTF8(3, cipher.StdUppercaseAlphabet))
	fmt.Println("Encrypted Text: ", cc.Cipher("HELLO WORLD"))
	fmt.Println("Decrypted Text: ", cc.Decipher("KHOOR ZRUOG"))
	fmt.Println("KeyAlphabet: ", cc.KeyAlphabet())
	fmt.Println("Name: ", cc.Name())
	fmt.Println("Description: ", cc.Description())

	// Create a ROT-13 cipher.
	kc := cipher.NewStdCipher(cipher.StdUppercaseAlphabet, cipher.RotateUTF8(13, cipher.StdUppercaseAlphabet))
	fmt.Println("Encrypted Text: ", kc.Cipher("HELLO WORLD"))
	fmt.Println("Decrypted Text: ", kc.Decipher("URYYB JBEYQ"))
	fmt.Println("KeyAlphabet: ", kc.KeyAlphabet())
	fmt.Println("Name: ", kc.Name())
	fmt.Println("Description: ", kc.Description())

	// For standard latin alphabet ROT-13 cipher and decipher is reversible.
	fmt.Println("Encrypted Text: ", kc.Cipher("URYYB JBEYQ"))
}
