package myfibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciSequence(t *testing.T) {
	assert.Equal(t, FibonacciSequence(-1), []int{})
	assert.Equal(t, FibonacciSequence(0), []int{})
	assert.Equal(t, FibonacciSequence(1), []int{1})
	assert.Equal(t, FibonacciSequence(2), []int{1, 1})
	assert.Equal(t, FibonacciSequence(5), []int{1, 1, 2, 3, 5})
}
