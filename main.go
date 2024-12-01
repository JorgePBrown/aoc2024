package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/jorgepbrown/aoc2024/day1"
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
	}
}

func mustLoadInput(day uint) (io.Reader, func() error) {
	f, err := os.Open(fmt.Sprintf("./day%d/input.txt", day))
	if err != nil {
		panic(err)
	}
	return f, f.Close
}
