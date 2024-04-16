package FizzBuzz

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	want := "1"
	got := fizzbuzz(1)

	if want != got {
		t.Errorf("I want %s, But I got %s", want, got)
	}
}
