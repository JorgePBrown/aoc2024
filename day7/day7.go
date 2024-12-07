package day7

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func SolvePart1(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	sum := 0
	for scanner.Scan() {
		equation, err := parseEquation(scanner.Text())
		if err != nil {
			return sum, err
		}
		if equation.IsPossible() {
			sum += equation.Result
		}

	}
	return sum, nil
}
func SolvePart2(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	sum := 0
	for scanner.Scan() {
		equation, err := parseEquation(scanner.Text())
		if err != nil {
			return sum, err
		}
		if equation.IsPossible2() {
			sum += equation.Result
		}

	}
	return sum, nil
}

type Equation struct {
	Result    int
	Arguments []int
}

func parseEquation(s string) (Equation, error) {
	eq := Equation{}
	colonIdx := strings.IndexByte(s, ':')
	if colonIdx == -1 {
		return eq, errors.New("invalid equation format")
	}
	result, err := strconv.Atoi(s[:colonIdx])
	if err != nil {
		return eq, err
	}
	eq.Result = result
	argsStrings := strings.Split(s[colonIdx+2:], " ")
	args := make([]int, len(argsStrings))

	for i := range argsStrings {
		v, err := strconv.Atoi(argsStrings[i])
		if err != nil {
			return eq, err
		}
		args[i] = v

	}
	eq.Arguments = args
	return eq, nil
}

func (e Equation) IsPossible() bool {
	return isPossible(e.Arguments, e.Result, e.Arguments[0], 1)
}
func (e Equation) IsPossible2() bool {
	return isPossible2(e.Arguments, e.Result, e.Arguments[0], 1)
}

func isPossible(i []int, result, current, idx int) bool {
	if idx >= len(i) {
		return result == current
	}

	newCurrent := current + i[idx]
	if isPossible(i, result, newCurrent, idx+1) {
		return true
	}

	newCurrent = current * i[idx]
	if isPossible(i, result, newCurrent, idx+1) {
		return true
	}

	return false
}
func isPossible2(i []int, result, current, idx int) bool {
	if idx >= len(i) {
		return result == current
	}

	newCurrent := current + i[idx]
	if isPossible2(i, result, newCurrent, idx+1) {
		return true
	}

	newCurrent = current * i[idx]
	if isPossible2(i, result, newCurrent, idx+1) {
		return true
	}

	var err error
	newCurrent, err = strconv.Atoi(fmt.Sprintf("%d%d", current, i[idx]))
	if err != nil {
		panic("how?")
	}
	if isPossible2(i, result, newCurrent, idx+1) {
		return true
	}

	return false
}
