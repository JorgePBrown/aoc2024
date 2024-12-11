package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"text/template"

	"github.com/jorgepbrown/aoc2024/day1"
	"github.com/jorgepbrown/aoc2024/day10"
	"github.com/jorgepbrown/aoc2024/day11"
	"github.com/jorgepbrown/aoc2024/day3"
	"github.com/jorgepbrown/aoc2024/day4"
	"github.com/jorgepbrown/aoc2024/day5"
	"github.com/jorgepbrown/aoc2024/day6"
	"github.com/jorgepbrown/aoc2024/day7"
	"github.com/jorgepbrown/aoc2024/day8"
	"github.com/jorgepbrown/aoc2024/day9"
)

var (
	initDay = flag.Bool("init", false, "--init")
	day     = flag.Uint("day", 0, "--day=1")
	part    = flag.Uint("part", 0, "--part=1|2")
)

func main() {
	flag.Parse()
	if *day == 0 {
		panic("day must be greater than 0")
	}

	if *initDay {
		if *day != 0 {
			setupDay(*day)
		}
		return
	}

	i, c := mustLoadInput(*day)
	defer c()
	switch *day {
	case 1:
		if *part == 1 {
			fmt.Println(day1.SolvePart1(i))
		} else {
			fmt.Println(day1.SolvePart2(i))
		}
	case 3:
		if *part == 1 {
			fmt.Println(day3.SolvePart1(i))
		} else {
			fmt.Println(day3.SolvePart2(i))
		}
	case 4:
		if *part == 1 {
			fmt.Println(day4.SolvePart1(i))
		} else {
			fmt.Println(day4.SolvePart2(i))
		}
	case 5:
		if *part == 1 {
			fmt.Println(day5.SolvePart1(i))
		} else {
			fmt.Println(day5.SolvePart2(i))
		}
	case 6:
		if *part == 1 {
			fmt.Println(day6.SolvePart1(i))
		} else {
			fmt.Println(day6.SolvePart2(i))
		}
	case 7:
		if *part == 1 {
			fmt.Println(day7.SolvePart1(i))
		} else {
			fmt.Println(day7.SolvePart2(i))
		}
	case 8:
		if *part == 1 {
			fmt.Println(day8.SolvePart1(i))
		} else {
			fmt.Println(day8.SolvePart2(i))
		}
	case 9:
		if *part == 1 {
			fmt.Println(day9.SolvePart1(i))
		} else {
			fmt.Println(day9.SolvePart2(i))
		}
	case 10:
		if *part == 1 {
			fmt.Println(day10.SolvePart1(i))
		} else {
			fmt.Println(day10.SolvePart2(i))
		}
	case 11:
		if *part == 1 {
			fmt.Println(day11.SolvePart1(i))
		} else {
			fmt.Println(day11.SolvePart2(i))
		}
	}
}

func mustLoadInput(day uint) (io.Reader, func() error) {
	f, err := os.Open(fmt.Sprintf("./day%d/input.txt", day))
	if err != nil {
		panic(err)
	}
	return f, f.Close
}

type TemplateInfo struct {
	Day uint
}

func setupDay(day uint) {
	err := os.Mkdir(fmt.Sprintf("day%d", day), 0777)
	if err != nil {
		panic(err)
	}

	ti := TemplateInfo{
		day,
	}

	ft, err := template.New("source file").Parse(`package day{{.Day}}

import "io"

func SolvePart1(r io.Reader) (int, error) {
	return 0, nil
}
func SolvePart2(r io.Reader) (int, error) {
	return 0, nil
}`)

	if err != nil {
		panic(err)
	}

	sourceF, err := os.Create(fmt.Sprintf("./day%d/day%d.go", day, day))
	if err != nil {
		panic(err)
	}

	err = ft.Execute(sourceF, &ti)
	if err != nil {
		panic(err)
	}

	ft, err = template.New("test file").Parse(`package day{{.Day}}

import (
	"testing"
	"strings"
)

func TestPart1(t *testing.T) {
	input := strings.NewReader("")

	expected := 0

	actual, err := SolvePart1(input)

	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Fatalf("expected=%d got=%d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	input := strings.NewReader("")

	expected := 0

	actual, err := SolvePart2(input)

	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Fatalf("expected=%d got=%d", expected, actual)
	}
}`)

	if err != nil {
		panic(err)
	}

	sourceF, err = os.Create(fmt.Sprintf("./day%d/day%d_test.go", day, day))
	if err != nil {
		panic(err)
	}

	err = ft.Execute(sourceF, ti)
	if err != nil {
		panic(err)
	}
}
