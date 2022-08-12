package main_test

import (
	"array_slices/server"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := main.Sum(numbers)
	want := 15
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}

}
