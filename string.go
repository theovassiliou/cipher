package cyphering

// ReverseUTF8 reverses a given UTF-8 string
// "abcde" --> "edcba"
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

// RotateUTF8 returns a left (delta >0) or right rotated (delta <0) string
// delta 2:  "abcde" --> "cdeab"
// delta -2: "abcde" --> "deabc"
func RotateUTF8(delta int, inputAlphabet string) string {
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
		return RotateUTF8(runeCount+delta, inputAlphabet)
	}
}
