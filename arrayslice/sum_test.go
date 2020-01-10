package arrayslice

import "testing"

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
