package main

import (
	"fmt"

	"github.com/si74/AlgorithmPractice/factorial"
)

func main() {

	fmt.Println("normal factorial")
	fmt.Println(factorial.Factorial(5))

	fmt.Println("tail recursive factorial")
	fmt.Println(factorial.TailFactorial(5, 1))

}
