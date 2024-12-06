package day6

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

type Direction int8

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

var (
	directionMap = map[byte]Direction{
		'^': UP,
		'>': RIGHT,
		'v': DOWN,
		'<': LEFT,
	}
)

func SolvePart1(r io.Reader) (int, error) {
	m := parse(r)

	x, y, direction, err := m.StartingPosition()
	if err != nil {
		return 0, err
	}

	out := false
	for !out {
		x, y, direction, out = m.Move(x, y, direction)
	}
	return m.walked.Len(), nil
}

func SolvePart2(r io.Reader) (int, error) {
	m := parse(r)

	x, y, direction, err := m.StartingPosition()
	if err != nil {
		return 0, err
	}

	m.LookForLoops(x, y, direction)

	return m.obstacles.Len(), nil
}

func parse(r io.Reader) *Map {
	return &Map{
		bufio.NewScanner(r),
		[]string{},
		NewPositionSet(),
		NewPositionSet(),
	}
}

type Map struct {
	scanner   *bufio.Scanner
	lines     []string
	walked    *PositionSet
	obstacles *PositionSet
}

func (m *Map) StartingPosition() (int, int, Direction, error) {
	i := 0
	for m.scanner.Scan() {
		line := m.scanner.Text()
		m.lines = append(m.lines, line)
		startCol := strings.IndexAny(line, "^<>v")
		if startCol != -1 {
			direction := directionMap[line[startCol]]
			return i, startCol, direction, nil
		}
		i += 1
	}
	return 0, 0, UP, errors.New("no starting position found")
}

func (m *Map) Move(x, y int, d Direction) (int, int, Direction, bool) {
	newX, newY, ok := m.FindObstacle(x, y, d)
	m.MarkTraveled(x, y, d, newX, newY)
	return newX, newY, rotate(d), !ok
}

func (m *Map) FindObstacle(x, y int, d Direction) (int, int, bool) {
	vecX, vecY := vec(d)

	if vecX == 0 {
		line := m.lines[x]
		var i int
		for i = y + vecY; i >= 0 && i < len(line); i += vecY {
			if line[i] == '#' {
				return x, i - vecY, true
			}
		}
		return x, i, false
	} else if vecX == -1 {
		var i int
		for i = x + vecX; i >= 0; i += vecX {
			line := m.lines[i]
			if line[y] == '#' {
				return i - vecX, y, true
			}
		}
		return i, y, false
	} else {
		var i int
		ok := true
		for i = x + vecX; ok; i += vecX {
			var line string
			line, ok = m.Line(i)
			if !ok {
				break
			}
			if line[y] == '#' {
				return i - vecX, y, true
			}
		}
		return i, y, false
	}
}

func (m *Map) Line(x int) (string, bool) {
	if x < 0 {
		return "", false
	}
	if x < len(m.lines) {
		return m.lines[x], true
	}
	i := len(m.lines)
	for m.scanner.Scan() {
		line := m.scanner.Text()
		m.lines = append(m.lines, line)
		if i == x {
			return line, true
		}
		i += 1
	}
	return "", false
}

func (m *Map) MarkTraveled(x, y int, d Direction, newX, newY int) {
	vecX, vecY := vec(d)

	for x != newX || y != newY {
		m.walked.Add(x, y)
		x += vecX
		y += vecY
	}
}

func (m *Map) LookForLoops(x, y int, d Direction) {
	initX, initY, initD := x, y, d
	ok := true
	for ok {
		vecX, vecY := vec(d)
		obsX, obsY := x+vecX, y+vecY
		if !m.IsOutOfBounds(obsX, obsY) && !m.IsObstacle(obsX, obsY) && m.IsLoop(initX, initY, initD, obsX, obsY) {
			m.obstacles.Add(obsX, obsY)
		}
		x, y, d, ok = m.MoveOnce(x, y, d)
	}
}

func (m *Map) IsLoop(x, y int, d Direction, obsX, obsY int) bool {
	walked := NewPositionDirectionSet()

	for {
		if walked.Has(x, y, d) {
			return true
		}

		walked.Add(x, y, d)

		vecX, vecY := vec(d)
		peekX := x + vecX
		peekY := y + vecY

		// obstacle?
		if (peekX == obsX && peekY == obsY) || m.IsObstacle(peekX, peekY) {
			// rotate
			d = rotate(d)
		} else {
			if m.IsOutOfBounds(peekX, peekY) {
				break
			}
			// move
			x = peekX
			y = peekY
		}
	}
	return false
}

func (m *Map) IsObstacle(x, y int) bool {
	l, ok := m.Line(x)
	if !ok {
		return false
	}
	return y >= 0 && y < len(l) && l[y] == '#'
}

func (m *Map) IsOutOfBounds(x, y int) bool {
	l, ok := m.Line(x)
	if !ok {
		return true
	}
	return y < 0 || y >= len(l)
}

func (m *Map) MoveOnce(x, y int, d Direction) (int, int, Direction, bool) {
	vecX, vecY := vec(d)
	peekX := x + vecX
	peekY := y + vecY

	if m.IsObstacle(peekX, peekY) {
		d = rotate(d)
	} else {
		x = peekX
		y = peekY
	}
	return x, y, d, !m.IsOutOfBounds(x, y)
}

func rotate(d Direction) Direction {
	switch d {
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	default:
		panic("unknown direction")
	}
}
func vec(d Direction) (int, int) {
	switch d {
	case UP:
		return -1, 0
	case RIGHT:
		return 0, 1
	case DOWN:
		return 1, 0
	case LEFT:
		return 0, -1
	default:
		panic("unknown direction")
	}
}

type PositionSet struct {
	m map[int]map[int]struct{}
}
type PositionDirectionSet struct {
	m map[Direction]*PositionSet
}

func NewPositionSet() *PositionSet {
	return &PositionSet{
		map[int]map[int]struct{}{},
	}
}
func NewPositionDirectionSet() *PositionDirectionSet {
	return &PositionDirectionSet{
		map[Direction]*PositionSet{
			UP:    NewPositionSet(),
			RIGHT: NewPositionSet(),
			DOWN:  NewPositionSet(),
			LEFT:  NewPositionSet(),
		},
	}
}

func (s *PositionSet) Add(x, y int) {
	if _, ok := s.m[x]; ok {
		s.m[x][y] = struct{}{}
	} else {
		s.m[x] = map[int]struct{}{
			y: struct{}{},
		}
	}
}
func (s *PositionSet) Has(x, y int) bool {
	if m, ok := s.m[x]; ok {
		_, ok := m[y]
		return ok
	}
	return false
}
func (s *PositionSet) Len() int {
	sum := 0
	for _, m := range s.m {
		sum += len(m)
	}
	return sum
}

func (s *PositionDirectionSet) Add(x, y int, d Direction) {
	s.m[d].Add(x, y)
}
func (s *PositionDirectionSet) Has(x, y int, d Direction) bool {
	return s.m[d].Has(x, y)
}
func (s *PositionDirectionSet) Len() int {
	sum := 0
	for _, m := range s.m {
		sum += m.Len()
	}
	return sum
}
