package main

import "testing"

func TestSumUp(t *testing.T) {
	sum := sumUp(100)
	if sum != 5050 {
		t.Error(`sumUp(100) should be equal to 5050`)
	}
}

// 1 + 2 + ... + n
func sumUp(n int) int {
	var total int
	for i := 1; i <= n; i++ {
		total += i
	}
	return total
}
