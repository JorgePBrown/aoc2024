package day3

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := strings.NewReader(`xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`)
	expected := 161
	actual, err := SolvePart1(input)
	if err != nil {
		t.Fatal(err)
	}
	if expected != actual {
		t.Fatalf("expected=%d got=%d", expected, actual)
	}
}
func TestPart2(t *testing.T) {
	input := strings.NewReader(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`)
	expected := 48
	actual, err := SolvePart2(input)
	if err != nil {
		t.Fatal(err)
	}
	if expected != actual {
		t.Fatalf("expected=%d got=%d", expected, actual)
	}
}
