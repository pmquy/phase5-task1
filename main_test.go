package main

import "testing"

func TestSum(t *testing.T) {
	a, b := 1, 2
	expected := 3
	result := sum(a, b)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
