package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Sum of 6 digits", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		got := Sum(numbers)
		want := 21

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}

	})

	t.Run("Any size sum", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 6}
		got := Sum(numbers)
		want := 16

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
