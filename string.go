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
func RotateUTF8(delta int, input string) string {
	if delta == 0 {
		return input
	} else if delta > 0 {
		tail := []rune{}
		output := []rune{}
		runeCount := 0

		for _, runeValue := range input {
			if runeCount < delta {
				tail = append(tail, runeValue)
			} else {
				output = append(output, runeValue)
			}
			runeCount++
		}
		output = append(output, tail...)

		return string(output)
	} else {
		runeCount := countRunes(input)
		return RotateUTF8(runeCount+delta, input)
	}
}

func StripDuplicates(input string) string {
	var output []rune

	for _, c := range input {
		if !contains(output, c) {
			output = append(output, c)
		}
	}
	return string(output)
}

func contains(s []rune, e rune) bool {
	for _, c := range s {
		if c == e {
			return true
		}
	}
	return false
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
