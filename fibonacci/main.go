package main

import (
	"fmt"
)

func main() {
	sum := 0
	first, second := 1, 2

	for second <= 4000000 {
		if second%2 == 0 {
			sum += second
		}

		first, second = second, first+second
	}

	fmt.Printf("Sum of even-valued terms in the Fibonacci sequence (not exceeding four million): %d\n", sum)
}
