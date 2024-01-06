package main

import (
	"fmt"
	"strconv"
)

// isPalindrome checks if a number is a palindrome.
func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// findLargestPalindrome finds the largest palindrome made from the product of two 3-digit numbers.
func findLargestPalindrome() int {
	maxPalindrome := 0
	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			product := i * j
			if isPalindrome(product) && product > maxPalindrome {
				maxPalindrome = product
			}
		}
	}
	return maxPalindrome
}

func main() {
	fmt.Println("Largest Palindrome:", findLargestPalindrome())
}
