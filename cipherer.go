package cipher

const StdLowercaseAlphabet = "abcdefghijklmnopqrstuvwxyz"
const StdUppercaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const StdGreekLowercaseAlphabet = "αβγδεζηθικλμνξοπρστυφχψω"
const StdGreekΘππερcaseAlphabet = "ΑΒΓΔΕΖΗΘΙΚΛΜΝΞΟΠΡΣΤΥΦΧΨΩ"
const StdAlphabet = StdLowercaseAlphabet + StdUppercaseAlphabet

type Describer interface {
	Name() string
	Description() string
	KeyAlphabet() string
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
