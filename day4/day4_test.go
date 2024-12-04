package day4

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := strings.NewReader(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`)

	expected := 18

	actual, err := SolvePart1(input)
	if err != nil {
		t.Fatal(err)
	}
	if expected != actual {
		t.Fatalf("expected=%d got=%d", expected, actual)
	}
}
func TestPart2(t *testing.T) {
	input := strings.NewReader(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`)

	expected := 9

	actual, err := SolvePart2(input)
	if err != nil {
		t.Fatal(err)
	}
	if expected != actual {
		t.Fatalf("expected=%d got=%d", expected, actual)
	}
}
