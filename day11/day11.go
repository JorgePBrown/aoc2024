package day11

import (
	"io"
	"strconv"
	"strings"
)

func Solve(r io.Reader, blinks int) (int, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	ss := strings.Split(strings.TrimSpace(string(b)), " ")
	is := make([]int64, len(ss))
	for i, s := range ss {
		n, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		is[i] = int64(n)
	}

	count := 0
	memo := map[int64][]int{}
	for _, n := range is {
		count += Expand(n, blinks, memo)
	}

	return count, nil
}
func SolvePart1(r io.Reader) (int, error) {
	return Solve(r, 25)
}

func SolvePart2(r io.Reader) (int, error) {
	return Solve(r, 75)
}

func Expand(n int64, blinks int, memo map[int64][]int) int {
	var fn func(int64, int) int
	fn = func(n int64, b int) int {
		if b == 0 {
			return 1
		}

		mem, ok := memo[n]
		if !ok {
			mem = make([]int, blinks)
			memo[n] = mem
		}

		count := mem[b-1]
		if count != 0 {
			return count
		}

		if n == 0 {
			count = fn(1, b-1)
			mem[b-1] = count
			return count
		} else if s := strconv.FormatInt(n, 10); len(s)%2 == 0 {
			half := len(s) / 2
			n1, err := strconv.ParseInt(s[:half], 10, 64)
			if err != nil {
				panic(err)
			}
			n2, err := strconv.ParseInt(s[half:], 10, 64)
			if err != nil {
				panic(err)
			}

			count = fn(n1, b-1) + fn(n2, b-1)
			mem[b-1] = count
			return count
		} else {
			count = fn(n*2024, b-1)
			mem[b-1] = count
			return count
		}
	}
	return fn(n, blinks)
}
