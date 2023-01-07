package numbertowords_test

import (
	"testing"

	"github.com/vish-k/numbertowords"
)

func TestInvalidInput(t *testing.T) {
	result, err := numbertowords.Convert(-1)
	if result != "" || err == nil {
		t.Fatal("failed test for negative number")
	}

	result, err = numbertowords.Convert(100000)
	if result != "" || err == nil {
		t.Fatal("failed test for 100000")
	}

}

func TestValidInputTo19(t *testing.T) {
	validrange := map[int]string{
		0:  "zero",
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
	}

	for num, word := range validrange {
		result, err := numbertowords.Convert(num)
		if result != word || err != nil {
			t.Fatal("failed test for valid number", num)
		}
	}
}

func Test20To99(t *testing.T) {
	result, err := numbertowords.Convert(20)
	if result != "twenty" || err != nil {
		t.Fatal("failed test for 20")
	}

	result, err = numbertowords.Convert(45)
	if result != "forty five" || err != nil {
		t.Fatal("failed test for 99")
	}

	result, err = numbertowords.Convert(99)
	if result != "ninety nine" || err != nil {
		t.Fatal("failed test for 99")
	}
}

func Test100To999(t *testing.T) {
	result, err := numbertowords.Convert(100)
	if result != "one hundred " || err != nil {
		t.Fatal("failed test for 100")
	}

	result, err = numbertowords.Convert(445)
	if result != "four hundred forty five" || err != nil {
		t.Fatal("failed test for 445")
	}

	result, err = numbertowords.Convert(999)
	if result != "nine hundred ninety nine" || err != nil {
		t.Fatal("failed test for 99")
	}
}

func Test1000To99999(t *testing.T) {
	result, err := numbertowords.Convert(1000)
	if result != "one thousand " || err != nil {
		t.Fatal("failed test for 100")
	}

	result, err = numbertowords.Convert(4445)
	if result != "four thousand four hundred forty five" || err != nil {
		t.Fatal("failed test for 445")
	}

	result, err = numbertowords.Convert(99999)
	if result != "ninety nine thousand nine hundred ninety nine" || err != nil {
		t.Fatal("failed test for 99")
	}
}
