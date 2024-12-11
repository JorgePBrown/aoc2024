package day11

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := strings.NewReader("125 17")

	expected := 55312

	actual, err := SolvePart1(input)

	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Fatalf("expected=%d got=%d", expected, actual)
	}
}

func BenchmarkSolvePart1(b *testing.B) {
	for _ = range b.N {
		b.StopTimer()
		input := strings.NewReader("4610211 4 0 59 3907 201586 929 33750")
		b.StartTimer()
		SolvePart1(input)
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	for _ = range b.N {
		b.StopTimer()
		input := strings.NewReader("4610211 4 0 59 3907 201586 929 33750")
		b.StartTimer()
		SolvePart2(input)
	}
}
