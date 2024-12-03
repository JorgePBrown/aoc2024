package day3

import (
	"io"
)

const (
	M uint = iota
	U
	L
	LPAREN
	FIRST_NUMBER
	SECOND_NUMBER

	D
	O
	DOBRANCH
	APO
	T
	RPARENDO
	LPARENDONT
	RPARENDONT
)

func SolvePart1(r io.Reader) (int, error) {
	bs, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	sum := 0
	stage := M
	n1 := 0
	n2 := 0
	for _, b := range bs {
		switch b {
		case 'm':
			if stage == M {
				n1 = 0
				n2 = 0
				stage = U
			}
		case 'u':
			if stage == U {
				stage = L
			} else {
				stage = M
			}
		case 'l':
			if stage == L {
				stage = LPAREN
			} else {
				stage = M
			}
		case '(':
			if stage == LPAREN {
				stage = FIRST_NUMBER
			} else {
				stage = M
			}
		case ',':
			if stage == FIRST_NUMBER {
				stage = SECOND_NUMBER
			} else {
				stage = M
			}
		case ')':
			if stage == SECOND_NUMBER {
				sum += n1 * n2
			}
			stage = M
		default:
			if b >= '0' && b <= '9' {
				if stage == FIRST_NUMBER {
					n1 *= 10
					n := b - '0'
					n1 += int(n)
				} else if stage == SECOND_NUMBER {
					n2 *= 10
					n := b - '0'
					n2 += int(n)
				}
			} else {
				stage = M
			}
		}
	}

	return sum, nil
}

func SolvePart2(r io.Reader) (int, error) {
	bs, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	sum := 0
	stage := M
	n1 := 0
	n2 := 0
	enabled := true
	for _, b := range bs {
		switch b {
		case 'm':
			if stage == M {
				n1 = 0
				n2 = 0
				stage = U
			}
		case 'u':
			if stage == U {
				stage = L
			} else {
				stage = M
			}
		case 'l':
			if stage == L {
				stage = LPAREN
			} else {
				stage = M
			}
		case '(':
			if stage == LPAREN {
				stage = FIRST_NUMBER
			} else if stage == DOBRANCH {
				stage = RPARENDO
			} else if stage == LPARENDONT {
				stage = RPARENDONT
			} else {
				stage = M
			}
		case ',':
			if stage == FIRST_NUMBER {
				stage = SECOND_NUMBER
			} else {
				stage = M
			}
		case ')':
			if stage == SECOND_NUMBER {
				if enabled {
					sum += n1 * n2
				}
			} else if stage == RPARENDO {
				enabled = true
			} else if stage == RPARENDONT {
				enabled = false
			}
			stage = M
		case 'd':
			if stage == M {
				stage = O
			}
		case 'o':
			if stage == O {
				stage = DOBRANCH
			} else {
				stage = M
			}
		case 'n':
			if stage == DOBRANCH {
				stage = APO
			} else {
				stage = M
			}
		case '\'':
			if stage == APO {
				stage = T
			} else {
				stage = M
			}
		case 't':
			if stage == T {
				stage = LPARENDONT
			} else {
				stage = M
			}
		default:
			if b >= '0' && b <= '9' {
				if stage == FIRST_NUMBER {
					n1 *= 10
					n := b - '0'
					n1 += int(n)
				} else if stage == SECOND_NUMBER {
					n2 *= 10
					n := b - '0'
					n2 += int(n)
				}
			} else {
				stage = M
			}
		}
	}

	return sum, nil
}
