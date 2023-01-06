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

	result, err = numbertowords.Convert(6)
	if result != "" || err == nil {
		t.Fatal("failed test for 6")
	}

}

func TestValidInput(t *testing.T) {
	validrange := map[int]string{
		0: "zero",
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
	}

	for num, word := range validrange {
		result, err := numbertowords.Convert(num)
		if result != word || err != nil {
			t.Fatal("failed test for valid number")
		}
	}
}
