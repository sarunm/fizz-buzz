package FizzBuzz

import (
	"testing"
)

func Test1ShouldReturn1(t *testing.T) {
	want := "1"
	got := fizzbuzz(1)

	if want != got {
		t.Errorf("I want %s, But I got %s", want, got)
	}
}

func Test2ShouldReturn2(t *testing.T) {
	want := "2"
	got := fizzbuzz(2)

	if want != got {
		t.Errorf("I want %s, But I got %s", want, got)
	}
}
