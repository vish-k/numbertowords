package numbertowords

func Convert(number int) (string, bool) {
	if number < 0 || number > 5 {
		return "", false
	}

	words := [6]string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
	}
	return words[number], true

}
