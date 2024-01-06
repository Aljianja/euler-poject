package main

import (
	"fmt"
)

func LargestPrimeFactor(number int64) int64 {
	var factor int64 = 2
	var largest int64 = 0

	for number > 1 {
		if number%factor == 0 {
			number /= factor
			largest = factor
		} else {
			factor++
		}
	}
	return largest
}

func main() {
	number := int64(600851475143)
	fmt.Println("Largest Prime Factor:", LargestPrimeFactor(number))
}
