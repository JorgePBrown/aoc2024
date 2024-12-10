package day10

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	i := strings.NewReader(`89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`)

	expected := 36
	actual, err := SolvePart1(i)
	if err != nil {
		t.Fatal(err)
	}
	if expected != actual {
		t.Fatalf("expected=%d got=%d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	i := strings.NewReader(`89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`)

	expected := 81
	actual, err := SolvePart2(i)
	if err != nil {
		t.Fatal(err)
	}
	if expected != actual {
		t.Fatalf("expected=%d got=%d", expected, actual)
	}
}
