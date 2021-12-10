package cipher

// ReverseUTF8 reverses a given UTF-8 string
// "abcde" --> "edcba"
func ReverseUTF8(input string) string {
	runes := []rune(input)
	n := len(runes)
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
		in := []rune(input)
		lin := len(in)
		output := make([]rune, lin)

		for i := delta; i < lin; i++ {
			output[i-delta] = in[i]
		}
		for i := 0; i < delta; i++ {
			output[lin-delta+i] = in[i]
		}
		return string(output)
	} else {
		runeCount := len([]rune(input))
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
