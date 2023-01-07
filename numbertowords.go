// Package numbertowords helps convert number to words
package numbertowords

import "errors"

var basewords = [20]string{
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

var tens = [10]string{
	"",
	"",
	"twenty",
	"thirty",
	"forty",
	"fifty",
	"sixty",
	"seventy",
	"eighty",
	"ninety",
}

const Max = 99999

// Convert converts a number to words.
// If number is out of range, an error is returned
func Convert(number int) (string, error) {
	if number < 0 || number > Max {
		return "", errors.New("not in valid range")
	}

	words := ""

	//convert the thousandth digit
	k := number / 1000
	if k > 0 {
		words, _ = Convert(k)
		words = words + " thousand "
		number = number % 1000
		if number == 0 {
			return words, nil
		}
	}

	//convert the thousandth digit
	h := number / 100 //convert hundredth digit
	if h > 0 {
		words = words + basewords[h] + " hundred "
		number = number % 100
		if number == 0 {
			return words, nil
		}
	}

	//convert the tenth digit
	t := number / 10
	if t < 2 { //up to 19
		return words + basewords[number], nil
	}

	//convert tens higher than 19
	n := number % 10
	if n == 0 {
		return words + tens[t], nil
	}

	return words + tens[t] + " " + basewords[n], nil

}
