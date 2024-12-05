package day5

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := strings.NewReader(`47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`)

	expected := 143

	actual, err := SolvePart1(input)
	if err != nil {
		t.Fatal(err)
	}
	if actual != expected {
		t.Fatalf("expected=%d got=%d", expected, actual)
	}
}
func TestPart2(t *testing.T) {
	input := strings.NewReader(`47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`)

	expected := 123

	actual, err := SolvePart2(input)
	if err != nil {
		t.Fatal(err)
	}
	if actual != expected {
		t.Fatalf("expected=%d got=%d", expected, actual)
	}
}
