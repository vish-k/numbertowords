package numbertowords

import "errors"

func Convert(number int) (string, error) {
	if number < 0 || number > 5 {
		return "", errors.New("not in valid range")
	}

	words := [6]string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
	}
	return words[number], nil

}
