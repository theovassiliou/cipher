package cyphering

const StdLowercaseAlphabet = "abcdefghijklmnopqrstuvwxyz"
const StdUppercaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const StdAlphabet = StdLowercaseAlphabet + StdUppercaseAlphabet

// ReverseUTF8 reverses a given UTF-8 string
func ReverseUTF8(input string) string {
	n := 0
	runes := make([]rune, len(input))
	for _, r := range input {
		runes[n] = r
		n++
	}
	runes = runes[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	// Convert back to UTF-8.
	return string(runes)
}

// ShiftUTF8 returns a left (delta >0) or right shifted (delta <0) string
// delta 2:  "abcde" --> "cdeab"
// delta -2: "abcde" --> "deabc"
func ShiftUTF8(delta int, inputAlphabet string) string {
	if delta == 0 {
		return inputAlphabet
	} else if delta > 0 {
		tail := []rune{}
		outputAlphabet := []rune{}
		runeNo := 0

		for _, runeValue := range inputAlphabet {
			if runeNo < delta {
				tail = append(tail, runeValue)
			} else {
				outputAlphabet = append(outputAlphabet, runeValue)
			}
			runeNo++
		}
		outputAlphabet = append(outputAlphabet, tail...)

		return string(outputAlphabet)
	} else {
		runeCount := countRunes(inputAlphabet)
		return ShiftUTF8(runeCount+delta, inputAlphabet)
	}
}

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

func str2Runes(std string) []rune {
	n := 0
	runes := make([]rune, len(std))
	for _, r := range std {
		runes[n] = r
		n++
	}
	return runes[0:n]
}

func countRunes(in string) int {
	runes := 0
	for range in {
		runes++
	}
	return runes
}
