package day1

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	expected := 11

	actual, err := SolvePart1(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}
	if actual != expected {
		t.Fatalf("expected=%d got=%d", expected, actual)
	}
}
func TestPart2(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	expected := 31

	actual, err := SolvePart2(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}
	if actual != expected {
		t.Fatalf("expected=%d got=%d", expected, actual)
	}
}
