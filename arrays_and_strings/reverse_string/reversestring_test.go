package reverse_string

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	input := []byte{'h', 'e', 'l', 'l', 'o'}
	got := ReverseString(input)
	want := []byte{'o', 'l', 'l', 'e', 'h'}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}
