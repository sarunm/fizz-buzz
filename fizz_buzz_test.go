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

func Test3ShouldReturnFizz(t *testing.T) {
	want := "Fizz"
	got := fizzbuzz(3)

	if want != got {
		t.Errorf("I want %s, But I got %s", want, got)
	}
}

func Test4ShouldReturn4(t *testing.T) {
	want := "4"
	got := fizzbuzz(4)

	if want != got {
		t.Errorf("I want %s, But I got %s", want, got)
	}
}

func Test5ShouldReturnBuzz(t *testing.T) {
	want := "Buzz"
	got := fizzbuzz(5)

	if want != got {
		t.Errorf("I want %s, But I got %s", want, got)
	}
}

func Test6ShouldReturnFizz(t *testing.T) {
	want := "Fizz"
	got := fizzbuzz(6)

	if want != got {
		t.Errorf("I want %s, But I got %s", want, got)
	}
}
