package day4

import (
	"bufio"
	"io"
)

func SolvePart1(r io.Reader) (int, error) {
	count := 0
	lines := NewLines(r, 4)
	for lines.Next() {
		line := lines.lines[0]
		for i, c := range line {
			if c == 'X' {
				//horizontal
				if lines.Check(0, i+1, 'M') && lines.Check(0, i+2, 'A') && lines.Check(0, i+3, 'S') {
					count += 1
				}
				//vertical
				if lines.Check(1, i, 'M') && lines.Check(2, i, 'A') && lines.Check(3, i, 'S') {
					count += 1
				}
				//diagonal /
				if lines.Check(1, i-1, 'M') && lines.Check(2, i-2, 'A') && lines.Check(3, i-3, 'S') {
					count += 1
				}
				//diagonal \
				if lines.Check(1, i+1, 'M') && lines.Check(2, i+2, 'A') && lines.Check(3, i+3, 'S') {
					count += 1
				}
			} else if c == 'S' {
				//horizontal
				if lines.Check(0, i+1, 'A') && lines.Check(0, i+2, 'M') && lines.Check(0, i+3, 'X') {
					count += 1
				}
				//vertical
				if lines.Check(1, i, 'A') && lines.Check(2, i, 'M') && lines.Check(3, i, 'X') {
					count += 1
				}
				//diagonal /
				if lines.Check(1, i-1, 'A') && lines.Check(2, i-2, 'M') && lines.Check(3, i-3, 'X') {
					count += 1
				}
				//diagonal \
				if lines.Check(1, i+1, 'A') && lines.Check(2, i+2, 'M') && lines.Check(3, i+3, 'X') {
					count += 1
				}
			}
		}
	}
	return count, nil
}

func SolvePart2(r io.Reader) (int, error) {
	count := 0
	lines := NewLines(r, 3)
	for lines.Next() {
		line := lines.lines[1]
		for i, c := range line {
			if c == 'A' {
				//diagonal \
				if (lines.Check(0, i-1, 'M') && lines.Check(2, i+1, 'S') || lines.Check(0, i-1, 'S') && lines.Check(2, i+1, 'M')) &&
					//diagonal /
					(lines.Check(0, i+1, 'M') && lines.Check(2, i-1, 'S') || lines.Check(0, i+1, 'S') && lines.Check(2, i-1, 'M')) {
					count += 1
				}
			}
		}
	}
	return count, nil
}

type Lines struct {
	scanner *bufio.Scanner
	lines   []string
}

func NewLines(r io.Reader, window int) *Lines {
	scanner := bufio.NewScanner(r)
	l := &Lines{scanner, make([]string, window)}
	i := 1
	for i < len(l.lines) && scanner.Scan() {
		l.lines[i] = scanner.Text()
		i += 1
	}
	for i < len(l.lines) {
		l.lines[i] = ""
		i += 1
	}
	return l
}

func (l *Lines) Next() bool {
	for i := 1; i < len(l.lines); i++ {
		l.lines[i-1] = l.lines[i]
	}

	if l.scanner.Scan() {
		l.lines[len(l.lines)-1] = l.scanner.Text()
		return true
	} else {
		l.lines[len(l.lines)-1] = ""
		for _, v := range l.lines {
			if v != "" {
				return true
			}
		}
		return false
	}
}
func (l *Lines) Offset(offset int) (string, bool) {
	if offset < 0 || offset >= len(l.lines) {
		return "", false
	}
	s := l.lines[offset]

	return s, s != ""
}

func (l *Lines) Check(hoffset, voffset int, c byte) bool {
	if s, ok := l.Offset(hoffset); ok {
		if voffset < 0 || voffset >= len(s) {
			return false
		}
		return s[voffset] == c
	}
	return false
}
