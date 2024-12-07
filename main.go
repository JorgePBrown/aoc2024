package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/jorgepbrown/aoc2024/day1"
	"github.com/jorgepbrown/aoc2024/day3"
	"github.com/jorgepbrown/aoc2024/day4"
	"github.com/jorgepbrown/aoc2024/day5"
	"github.com/jorgepbrown/aoc2024/day6"
	"github.com/jorgepbrown/aoc2024/day7"
)

var (
	day  = flag.Uint("day", 0, "--day=1")
	part = flag.Uint("part", 0, "--part=1|2")
)

func main() {
	flag.Parse()
	if *day == 0 {
		panic("day must be greater than 0")
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
	}
}

func mustLoadInput(day uint) (io.Reader, func() error) {
	f, err := os.Open(fmt.Sprintf("./day%d/input.txt", day))
	if err != nil {
		panic(err)
	}
	return f, f.Close
}
