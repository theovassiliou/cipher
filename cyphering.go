package cyphering

const StdLowercaseAlphabet = "abcdefghijklmnopqrstuvwxyz"
const StdUppercaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const StdAlphabet = StdLowercaseAlphabet + StdUppercaseAlphabet

func Decypher(inputAlphabet, secretAlphabet, cyphertext string) (input string) {
	iaRunes := str2Runes(inputAlphabet)
	saRunes := str2Runes(secretAlphabet)
	cyphertextRunes := str2Runes(cyphertext)

	inputRunes := make([]rune, len(cyphertextRunes))

	for p1, r1 := range cyphertextRunes {
		// find position of rune in secret alphabet
		var cleartextRune rune
		for p2, r2 := range saRunes {
			if r1 == r2 {
				// find the rune at position in secret alphabet
				cleartextRune = iaRunes[p2]
				break
			} else {
				cleartextRune = r1
			}

		}
		// write secreteRune to outputString
		inputRunes[p1] = cleartextRune
	}

	return string(inputRunes)
}

func Cypher(inputAlphabet, secretAlphabet, input string) (cyphertext string) {
	iaRunes := str2Runes(inputAlphabet)
	saRunes := str2Runes(secretAlphabet)
	inputRunes := str2Runes(input)
	cyphertextRunes := make([]rune, len(inputRunes))

	if len(iaRunes) == 0 || len(saRunes) == 0 {
		return ""
	}

	for p1, r1 := range inputRunes {
		// find position of rune in secret alphabet
		var cleartextRune rune
		for p2, r2 := range iaRunes {
			if r1 == r2 {
				// find the rune at position in secret alphabet
				cleartextRune = saRunes[p2]
				break
			}
			cleartextRune = r1
		}

		// write secreteRune to outputString
		cyphertextRunes[p1] = cleartextRune
	}

	return string(cyphertextRunes)
}

// NewAlphabet create a new alphabet based on a given keyword, white spaces and
// any other special character are not treated special
// Example:
// Input: ASECRETKEYWORD, StdUppercaseAlphabet
// ASECRTKYWODBFGHIJLMNPQRUVXYZ
func NewAlphabet(keyword, baseAlphabet string) string {
	var output []rune
	cleanKeyword := StripDuplicates(keyword)

	for _, c := range cleanKeyword {
		output = append(output, c)
	}

	for _, c := range baseAlphabet {
		if !contains(output, c) {
			output = append(output, c)
		}
	}

	return string(output)
}
