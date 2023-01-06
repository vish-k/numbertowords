package main

import (
	"fmt"

	"github.com/vish-k/numbertowords/numbertowords"
)

func main() {
	testvalues := []int{
		-1,
		0,
		1,
		2,
		3,
		4,
		5,
		10,
	}

	for _, value := range testvalues {
		result, valid := numbertowords.Convert(value)
		fmt.Println(value, result, valid)

	}
}
