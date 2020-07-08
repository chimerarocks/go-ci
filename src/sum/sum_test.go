package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	if sum(5, 5) != 10 {
		t.Errorf("Expected sum of %d and %d to be %d", 5,5,10)
	}
}