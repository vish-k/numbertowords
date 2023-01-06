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
		result, err := numbertowords.Convert(value)

		if err != nil {
			fmt.Println("ERROR:", fmt.Errorf("%d %s", value, err))
			continue
		}
		fmt.Println(value, result)

	}
}
