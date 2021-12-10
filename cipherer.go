package cyphering

const StdLowercaseAlphabet = "abcdefghijklmnopqrstuvwxyz"
const StdUppercaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const StdAlphabet = StdLowercaseAlphabet + StdUppercaseAlphabet

type Describer interface {
	Name() string
	Description() string
}
type Cipherer interface {
	Cipher(plaintext string) string
	Describer
}
type Decipherer interface {
	Decipher(cipherText string) string
	Describer
}

type CiphererDecipherer interface {
	Cipherer
	Decipherer
	Describer
}
