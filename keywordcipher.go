package cipher

type KeywordCipher struct {
	name          string
	plainAlphabet string
	keyAlphabet   string
	keyword       string
}

func NewKeywordCipherer(keyword, plainAlphabet string) CiphererDecipherer {
	return KeywordCipher{
		name:          "KeywordCipherer",
		plainAlphabet: plainAlphabet,
		keyAlphabet:   NewKeywordAlphabet(keyword, plainAlphabet),
		keyword:       keyword,
	}
}

func (c KeywordCipher) Cipher(plaintext string) string {
	return stdCipher(c.plainAlphabet, c.keyAlphabet, plaintext)
}

func (c KeywordCipher) Decipher(cipherText string) string {
	return stdDecipher(c.plainAlphabet, c.keyAlphabet, cipherText)
}

func (c KeywordCipher) Name() string {
	return c.name
}

func (c KeywordCipher) KeyAlphabet() string {
	return c.keyAlphabet
}

func (c KeywordCipher) Description() string {
	return `Keyword Cipher encodes a plaintext based on a given keyword.
The key alphabet will be constructed as follows.
First, the keyword will be stripped to contain only unique characters.
Second, the alphabet will be appended with no characters from the keyword.

	// Example:
	// Input: ASECRETKEYWORD, StdUppercaseAlphabet
	// ASECRTKYWODBFGHIJLMNPQRUVXYZ
`
}

// NewKeywordAlphabet create a new, supposed secret key alphabet based on a given keyword,
// white spaces and any other special character are not treated special
// Characters of keyword must be included in plainAlphabet, otherwise character
// will be ignored in key alphabet.

// Example for key alphabet:
// Input: ASECRETKEYWORD, StdUppercaseAlphabet
// ASECRTKYWODBFGHIJLMNPQRUVXYZ
func NewKeywordAlphabet(keyword, plainAlphabet string) string {
	var keyAlphabet []rune
	plainAlphabetRunes := []rune(plainAlphabet)

	strippedKeyword := StripDuplicates(keyword)

	for _, c := range strippedKeyword {
		if contains(plainAlphabetRunes, c) {
			keyAlphabet = append(keyAlphabet, c)
		}
	}

	for _, c := range plainAlphabet {
		if !contains(keyAlphabet, c) {
			keyAlphabet = append(keyAlphabet, c)
		}
	}

	return string(keyAlphabet)
}
