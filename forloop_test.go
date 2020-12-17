package main

import "testing"

func TestTwoEqualsTwo(t *testing.T) {
	if 2 != 2 {
		t.Error(`2 should be equal to 2`)
	}
}
