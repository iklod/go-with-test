package arrslice

import "testing"

func TestSum(t *testing.T) {
	t.Run("Add from array of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6
		assert(t, got, want, numbers)
	})
}

func assert(t testing.TB, got, want int, numbers []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
