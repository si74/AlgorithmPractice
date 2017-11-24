// Package factorial tests out tail recursive vs. non tail recursive way to write functions
package factorial

import "fmt"

// Factorial is normal recursive factorial
func Factorial(n int) int {
	defer fmt.Printf("step: %v * fact(%v - 1)\n", n, n)
	if n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// TailFactorial uses tail recursion for factorial
// note - some compilers will optimize for this (C, Lisp, etc)
// since it possibly tweaks the call stack requirements
// from O(n) to O(1)
func TailFactorial(n int, product int) int {
	defer fmt.Printf("step: %v * fact(%v - 1)\n", n, n)
	if n == 1 {
		return product
	}
	product *= n
	return TailFactorial(n-1, product)
}
