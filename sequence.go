// Package myfibonacci provides example functions for generating sequences of numbers.
package myfibonacci

// FibonacciSequence generates a Fibonacci sequence up to the given number of terms.
func FibonacciSequence(n int) []int {
	if n <= 0 {
		return []int{}
	} else if n == 1 {
		return []int{1}
	}

	sequence := make([]int, n)
	sequence[0] = 1
	sequence[1] = 1

	for i := 2; i < n; i++ {
		sequence[i] = sequence[i-2] + sequence[i-1]
	}

	return sequence
}
