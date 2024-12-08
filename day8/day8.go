package day8

import (
	"bufio"
	"fmt"
	"io"
)

func SolvePart1(r io.Reader) (int, error) {
	m := NewMap(r)

	err := m.FindAntinodes()
	if err != nil {
		return 0, err
	}

	return m.Antinodes.Len(), nil
}

func SolvePart2(r io.Reader) (int, error) {
	m := NewMap(r)

	err := m.FindAntinodes2()
	if err != nil {
		return 0, err
	}

	return m.Antinodes.Len(), nil
}

type Map struct {
	scanner   *bufio.Scanner
	m         []string
	Antinodes *PositionSet
	Nodes     map[byte]*PositionSet
}

func NewMap(r io.Reader) *Map {
	scanner := bufio.NewScanner(r)
	if scanner.Scan() {
		t := scanner.Text()
		return &Map{
			scanner,
			[]string{t},
			NewPositionSet(len(t)),
			map[byte]*PositionSet{},
		}
	} else {
		panic("reader is empty")
	}
}

func (m *Map) FindAntinodes() error {
	x := 0
	for {
		line, err := m.Line(x)
		if err != nil {
			break
		}
		for y := range line {
			c := line[y]
			if c != '.' {
				m.MarkAntinodes(c, x, y)
			}
		}
		x += 1
	}

	for i := range m.Antinodes.data {
		if i >= x {
			delete(m.Antinodes.data, i)
		}
	}

	return nil
}

func (m *Map) FindAntinodes2() error {
	x := 0
	for {
		line, err := m.Line(x)
		if err != nil {
			break
		}
		for y := range line {
			c := line[y]
			if c != '.' {
				m.MarkAntinodes2(c, x, y)
			}
		}
		x += 1
	}

	return nil
}

func (m *Map) Line(i int) (string, error) {
	if i < 0 {
		return "", fmt.Errorf("index out of bounds %d", i)
	}

	if i < len(m.m) {
		return m.m[i], nil
	}

	j := len(m.m)
	for m.scanner.Scan() {
		text := m.scanner.Text()
		m.m = append(m.m, text)
		if j == i {
			return text, nil
		}
		j += 1
	}

	return "", io.EOF
}

func (m *Map) MarkAntinodes(b byte, x1, y1 int) {
	var set *PositionSet
	if s, ok := m.Nodes[b]; ok {
		set = s
	} else {
		set = NewPositionSet(len(m.m[0]))
		m.Nodes[b] = set
	}
	for x2, v := range set.data {
		for y2 := range v {
			vec := diff(x1, y1, x2, y2)

			// mark
			m.Antinodes.Set(move(x2, y2, vec))
			m.Antinodes.Set(move(x1, y1, compl(vec)))
		}
	}
	set.Set(x1, y1)
}

func (m *Map) MarkAntinodes2(b byte, x1, y1 int) {
	var set *PositionSet
	if s, ok := m.Nodes[b]; ok {
		set = s
	} else {
		set = NewPositionSet(len(m.m[0]))
		m.Nodes[b] = set
	}
	for x2, v := range set.data {
		for y2 := range v {
			vec := diff(x1, y1, x2, y2)

			m.Antinodes.Set(x1, y1)
			m.Antinodes.Set(x2, y2)
			// mark

			x2, y2 = move(x2, y2, vec)
			for x2 >= 0 && y2 >= 0 && m.IsLine(x2) && y2 < m.Antinodes.yLimit {
				m.Antinodes.Set(x2, y2)
				x2, y2 = move(x2, y2, vec)
			}

			cmp := compl(vec)
			x3, y3 := move(x1, y1, cmp)
			for x3 >= 0 && y3 >= 0 && m.IsLine(x3) && y3 < m.Antinodes.yLimit {
				m.Antinodes.Set(x3, y3)
				x3, y3 = move(x3, y3, cmp)
			}
		}
	}
	set.Set(x1, y1)
}

func (m *Map) IsLine(i int) bool {
	_, err := m.Line(i)
	return err == nil
}

type PositionSet struct {
	data   map[int]map[int]struct{}
	yLimit int
}

func NewPositionSet(yLimit int) *PositionSet {
	return &PositionSet{
		data:   map[int]map[int]struct{}{},
		yLimit: yLimit,
	}
}

func (s *PositionSet) Set(x, y int) {
	if x < 0 || y < 0 || y >= s.yLimit {
		return
	}
	if v, ok := s.data[x]; ok {
		v[y] = struct{}{}
	} else {
		s.data[x] = map[int]struct{}{
			y: struct{}{},
		}
	}
}
func (s *PositionSet) Has(x, y int) bool {
	if v, ok := s.data[x]; ok {
		_, ok := v[y]
		return ok
	} else {
		return false
	}
}

func (s *PositionSet) Len() int {
	sum := 0
	for _, innerS := range s.data {
		sum += len(innerS)
	}
	return sum
}

type Vector struct {
	x int
	y int
}

func move(x, y int, v Vector) (int, int) {
	return x + v.x, y + v.y
}
func diff(x1, y1 int, x2, y2 int) Vector {
	return Vector{
		x2 - x1,
		y2 - y1,
	}
}
func compl(v Vector) Vector {
	return Vector{
		-v.x,
		-v.y,
	}
}
