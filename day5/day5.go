package day5

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

var (
	ErrMalformatedRule = errors.New("malformated rule: missing pipe")
)

func SolvePart1(r io.Reader) (int, error) {
	rules, pages, err := parse(r)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, page := range pages {
		if page.IsCorrect(rules) {
			sum += page.Middle()
		}
	}

	return sum, nil
}

func SolvePart2(r io.Reader) (int, error) {
	rules, pages, err := parse(r)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, page := range pages {
		if !page.IsCorrect(rules) {
			page.Sort(rules)
			sum += page.Middle()
		}
	}

	return sum, nil
}

func parse(r io.Reader) (*Rules, []Pages, error) {
	scanner := bufio.NewScanner(r)
	rules := NewRules()
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		pipeIdx := strings.IndexByte(line, '|')
		if pipeIdx == -1 {
			return nil, nil, ErrMalformatedRule
		}
		before, err := strconv.Atoi(line[:pipeIdx])
		if err != nil {
			return nil, nil, err
		}
		after, err := strconv.Atoi(line[pipeIdx+1:])
		if err != nil {
			return nil, nil, err
		}
		rules.Register(before, after)
	}

	pages := []Pages{}
	for scanner.Scan() {
		line := scanner.Text()
		page := Pages{}
		i := 0
		for j, v := range line {
			if v == ',' {
				n, err := strconv.Atoi(line[i:j])
				if err != nil {
					return nil, nil, err
				}
				page = append(page, n)
				i = j + 1
			}
		}
		n, err := strconv.Atoi(line[i:])
		if err != nil {
			return nil, nil, err
		}
		page = append(page, n)
		pages = append(pages, page)
	}

	return rules, pages, nil
}

type Rules struct {
	beforeMap map[int][]int
	afterMap  map[int][]int
}

func NewRules() *Rules {
	return &Rules{
		map[int][]int{},
		map[int][]int{},
	}
}

func (r *Rules) Register(before, after int) {
	if v, ok := r.beforeMap[before]; ok {
		r.beforeMap[before] = append(v, after)
	} else {
		r.beforeMap[before] = []int{after}
	}
	if v, ok := r.afterMap[after]; ok {
		r.afterMap[after] = append(v, before)
	} else {
		r.afterMap[after] = []int{before}
	}
}

func (r *Rules) Before(before, after int) bool {
	if v, ok := r.beforeMap[before]; ok {
		for _, elem := range v {
			if elem == after {
				return true
			}
		}
		return false
	}
	return false
}
func (r *Rules) After(after, before int) bool {
	if v, ok := r.afterMap[after]; ok {
		for _, elem := range v {
			if elem == before {
				return true
			}
		}
		return false
	}
	return false
}

type Pages []int

func (p Pages) IsCorrect(r *Rules) bool {
	for i, v := range p {
		for j := i - 1; j >= 0; j-- {
			if r.Before(v, p[j]) {
				return false
			}
		}
	}
	return true
}

func (p Pages) Middle() int {
	return p[len(p)/2]
}

func (p Pages) Sort(r *Rules) {
	c := make(Pages, 0, len(p))
	for i := 0; i < len(p); i++ {
		v1 := p[i]
		inserted := false
		for j, v2 := range c {
			if r.Before(v1, v2) {
				c = append(c, 0)
				for i := len(c) - 1; i > j; i-- {
					c[i] = c[i-1]
				}
				c[j] = v1
				inserted = true
				break
			}
		}
		if !inserted {
			c = append(c, v1)
		}
	}
	for i := range c {
		p[i] = c[i]
	}
}
