package cipher

type StdCipher struct {
	name          string
	plainAlphabet string
	keyAlphabet   string
}

func NewStdCipher(plainAlphabet, keyAlphabet string) CiphererDecipherer {
	return StdCipher{
		name:          "StandardCipherer",
		plainAlphabet: plainAlphabet,
		keyAlphabet:   keyAlphabet,
	}
}

func NewCaesarCipher(plainAlphabet string) CiphererDecipherer {
	return StdCipher{
		name:          "CeasarCipherer",
		plainAlphabet: plainAlphabet,
		keyAlphabet:   RotateUTF8(3, plainAlphabet),
	}
}

func (c StdCipher) Cipher(plaintext string) string {
	return stdCipher(c.plainAlphabet, c.keyAlphabet, plaintext)
}

func (c StdCipher) Decipher(cipherText string) string {
	return stdDecipher(c.plainAlphabet, c.keyAlphabet, cipherText)
}

func (c StdCipher) Name() string {
	return c.name
}

func (c StdCipher) Description() string {
	return `Standard Cipher encodes a plaintext by using a 
shift rotated alphabet. 

	// Example of Shift rotated alphabet:
	// Input: 13, StdUppercaseAlphabet
	// NOPQRSTUVWXYZABCDEFGHIJKLM
`
}

func stdDecipher(plainAlphabet, keyAlphabet, ciphertext string) (input string) {
	paRunes := []rune(plainAlphabet)
	kaRunes := []rune(keyAlphabet)
	ciphertextRunes := []rune(ciphertext)

	plaintextRunes := make([]rune, len(ciphertextRunes))

	for p1, r1 := range ciphertextRunes {
		// find position of rune in key alphabet
		var plaintextRune rune
		for p2, r2 := range kaRunes {
			if r1 == r2 {
				// find the rune at position in secret alphabet
				plaintextRune = paRunes[p2]
				break
			} else {
				plaintextRune = r1
			}

		}
		plaintextRunes[p1] = plaintextRune
	}

	return string(plaintextRunes)
}

func stdCipher(plainAlphabet, keyAlphabet, plaintext string) (ciphertext string) {
	paRunes := []rune(plainAlphabet)
	kaRunes := []rune(keyAlphabet)
	plaintextRunes := []rune(plaintext)
	ciphertextRunes := make([]rune, len(plaintextRunes))

	if len(paRunes) == 0 || len(kaRunes) == 0 {
		return ""
	}

	for p1, r1 := range plaintextRunes {
		// find position of rune in secret alphabet
		var cipherRune rune
		for p2, r2 := range paRunes {
			if r1 == r2 {
				// find the rune at position in key alphabet
				cipherRune = kaRunes[p2]
				break
			}
			cipherRune = r1
		}

		// write cipherRune to outputString
		ciphertextRunes[p1] = cipherRune
	}

	return string(ciphertextRunes)
}
