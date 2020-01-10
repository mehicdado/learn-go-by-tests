package arrayslice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6

		if got != expected {
			t.Errorf("expected %d but got %d given, %v", expected, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 9})
	expected := []int{3, 12}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %v but got %v", expected, got)
	}
}

func TestSumAllTails(t *testing.T) {
	assertSums := func(t *testing.T, got, expected []int) {
		t.Helper()

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	}

	t.Run("make sum of some slice tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 2}, []int{1, 9})
		expected := []int{4, 9}

		assertSums(t, got, expected)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 9})
		expected := []int{0, 9}

		assertSums(t, got, expected)
	})
}
