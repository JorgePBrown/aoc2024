package day9

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	i := strings.NewReader(`2333133121414131402`)

	expected := 1928
	actual, err := SolvePart1(i)
	if err != nil {
		t.Fatal(err)
	}
	if expected != actual {
		t.Fatalf("expected=%d got=%d", expected, actual)
	}
}
func TestPart2(t *testing.T) {
	i := strings.NewReader(`2333133121414131402`)

	expected := 2858
	actual, err := SolvePart2(i)
	if err != nil {
		t.Fatal(err)
	}
	if expected != actual {
		t.Fatalf("expected=%d got=%d", expected, actual)
	}
}
