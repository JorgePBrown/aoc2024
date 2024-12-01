package day1

import (
	"bufio"
	"errors"
	"io"
	"slices"
	"strconv"
	"strings"
)

var (
	ErrDifferentSizeLists = errors.New("lists have different sizes")
)

func SolvePart1(r io.Reader) (int, error) {
	l1, l2, err := parseLists(r)
	if err != nil {
		return 0, err
	}
	return diff(l1, l2)
}

func SolvePart2(r io.Reader) (int, error) {
	l1, l2, err := parseLists(r)
	if err != nil {
		return 0, err
	}
	return simm(l1, l2)
}

type List []int

func parseLists(r io.Reader) (List, List, error) {
	scanner := bufio.NewScanner(r)
	l1 := List{}
	l2 := List{}
	for scanner.Scan() {
		t := scanner.Text()
		t = strings.TrimSpace(t)
		firstSpaceIdx := strings.IndexByte(t, ' ')
		lastSpaceIdx := strings.LastIndexByte(t, ' ')
		l1Item, err := strconv.Atoi(t[:firstSpaceIdx])
		if err != nil {
			return nil, nil, err
		}
		l2Item, err := strconv.Atoi(t[lastSpaceIdx+1:])
		if err != nil {
			return nil, nil, err
		}
		l1 = append(l1, l1Item)
		l2 = append(l2, l2Item)
	}
	return l1, l2, nil
}

func diff(l1, l2 List) (int, error) {
	if len(l1) != len(l2) {
		return 0, ErrDifferentSizeLists
	}

	slices.Sort(l1)
	slices.Sort(l2)

	d := 0
	for i := range l1 {
		elemD := l1[i] - l2[i]
		if elemD < 0 {
			elemD = -elemD
		}
		d += elemD
	}

	return d, nil
}

func simm(l1, l2 List) (int, error) {
	if len(l1) != len(l2) {
		return 0, ErrDifferentSizeLists
	}
	slices.Sort(l2)

	count := map[int]int{}

	for _, e := range l2 {
		if v, ok := count[e]; ok {
			count[e] = v + 1
		} else {
			count[e] = 1
		}
	}

	s := 0
	for _, e := range l1 {
		if v, ok := count[e]; ok {
			s += e * v
		}
	}

	return s, nil
}
