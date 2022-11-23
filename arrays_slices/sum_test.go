package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("sum of slice, variable length", func(t *testing.T) {
		slice := []int{3, 4, 65}
		sum := Sum(slice)
		expected := 72
		if sum != expected {
			t.Errorf("Expected %d but got %d. Slice -> %v", expected, sum, slice)
		}
	})

}
