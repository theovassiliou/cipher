package main

import (
	"fmt"

	"github.com/theovassiliou/cyphering"
)

func main() {
	kc := cyphering.NewKeywordCipherer("WEISSKOPFSEEADLER", cyphering.StdUppercaseAlphabet)
	fmt.Println("Encrypted Text: ", kc.Cipher("NYT SEITE8 HEUTE 6PM BPPUTHAUS"))
	fmt.Println("Decrypted Text: ", kc.Decipher("CYQ NKAQK8 FKTQK 6HB EHHTQFWTN"))
	fmt.Println("Name: ", kc.Name())
	fmt.Println("Description: ", kc.Description())
}
