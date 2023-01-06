// Package numbertowords helps convert number to words
package numbertowords

import "errors"

var words = [20]string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
}

const Max = 19

// Convert converts a number to words.
// If number is out of range, an error is returned
func Convert(number int) (string, error) {
	if number < 0 || number > Max {
		return "", errors.New("not in valid range")
	}
	return words[number], nil

}
