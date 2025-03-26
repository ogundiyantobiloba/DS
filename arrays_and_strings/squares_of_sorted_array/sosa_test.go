package squares_of_sorted_array

import (
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	got := SortedSquares([]int{-4, -1, 0, 3, 10})
	want := []int{0, 1, 9, 16, 100}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
