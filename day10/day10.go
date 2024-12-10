package day10

import (
	"errors"
	"io"
	"slices"
)

func SolvePart1(r io.Reader) (int, error) {
	tm, err := NewTopographicMap(r)
	if err != nil {
		return 0, err
	}

	trailheads := tm.FindPossibleTrailheads()

	visited := NewNestedPositionSet()
	sum := 0
	for _, pos := range trailheads {
		sum += tm.CountReachableNines(pos.x, pos.y, visited)
	}

	return sum, nil
}

func SolvePart2(r io.Reader) (int, error) {
	tm, err := NewTopographicMap(r)
	if err != nil {
		return 0, err
	}

	trailheads := tm.FindPossibleTrailheads()

	visited := NewPositionCounterSet()
	sum := 0
	for _, pos := range trailheads {
		sum += tm.CountReachableNines2(pos.x, pos.y, visited)
	}

	return sum, nil
}

type TopographicMap [][]byte

func NewTopographicMap(r io.Reader) (TopographicMap, error) {
	buf := make([]byte, 1024)
	tm := TopographicMap{}
	row := []byte{}
	for {
		n, err := r.Read(buf)
		i := 0
		for i < n {
			nlIdx := -1
			for j := i; j < n; j += 1 {
				if buf[j] < '0' || buf[j] > '9' {
					nlIdx = j
					break
				}
			}

			if nlIdx == -1 {
				row = slices.Concat(row, buf[i:n])
				i = n
				break
			} else {
				row = slices.Concat(row, buf[i:nlIdx])
				tm = append(tm, row)
				row = []byte{}
				i = nlIdx + 1
			}
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				if len(row) > 0 {
					tm = append(tm, row)
				}
				break
			}
			return nil, err
		}
	}
	return tm, nil
}

func (tm TopographicMap) FindPossibleTrailheads() []Position {
	pos := []Position{}
	for x, row := range tm {
		for y, b := range row {
			if b == '0' {
				pos = append(pos, Position{
					x,
					y,
				})
			}
		}
	}
	return pos
}

type Position struct {
	x int
	y int
}

type PositionSet map[int]map[int]struct{}
type PositionCounterSet map[int]map[int]int
type NestedPositionSet map[int]map[int]PositionSet

func NewPositionSet() PositionSet {
	return PositionSet{}
}
func NewPositionCounterSet() PositionCounterSet {
	return PositionCounterSet{}
}
func NewNestedPositionSet() NestedPositionSet {
	return NestedPositionSet{}
}

func (ps PositionSet) Set(x, y int) {
	if v, ok := ps[x]; ok {
		v[y] = struct{}{}
	} else {
		ps[x] = map[int]struct{}{
			y: struct{}{},
		}
	}
}
func (ps PositionCounterSet) Set(x, y, v int) {
	if inner, ok := ps[x]; ok {
		inner[y] = v
	} else {
		ps[x] = map[int]int{
			y: v,
		}
	}
}
func (ps PositionCounterSet) Get(x, y int) (int, bool) {
	if inner, ok := ps[x]; ok {
		v, ok := inner[y]
		return v, ok
	} else {
		return 0, false
	}
}
func (ps NestedPositionSet) Set(x, y, xv, yv int) {
	if v1, ok := ps[x]; ok {
		if v2, ok := v1[y]; ok {
			v2.Set(xv, yv)
		} else {
			set := NewPositionSet()
			set.Set(xv, yv)
			v1[y] = set
		}
	} else {
		set := NewPositionSet()
		set.Set(xv, yv)
		ps[x] = map[int]PositionSet{
			y: set,
		}
	}
}
func (ps PositionSet) Has(x, y int) bool {
	if v, ok := ps[x]; ok {
		_, ok := v[y]
		return ok
	} else {
		return false
	}
}
func (ps PositionSet) Len() int {
	l := 0
	for _, row := range ps {
		l += len(row)
	}
	return l
}
func (ps NestedPositionSet) Get(x, y int) (PositionSet, bool) {
	if v, ok := ps[x]; ok {
		value, ok := v[y]
		return value, ok
	} else {
		return nil, false
	}
}

func (tm TopographicMap) CountReachableNines(x, y int, visited NestedPositionSet) int {
	var fn func(x, y int, visited NestedPositionSet, current byte)
	fn = func(x, y int, visited NestedPositionSet, current byte) {
		if current == '9' {
			visited.Set(x, y, x, y)
			return
		}
		_, ok := visited.Get(x, y)
		if ok {
			return
		}
		target := current + 1
		if x > 0 && tm[x-1][y] == target {
			fn(x-1, y, visited, target)
			set, ok := visited.Get(x-1, y)
			if ok {
				for xt, row := range set {
					for yt := range row {
						visited.Set(x, y, xt, yt)
					}
				}
			}
		}
		if y > 0 && tm[x][y-1] == target {
			fn(x, y-1, visited, target)
			set, ok := visited.Get(x, y-1)
			if ok {
				for xt, row := range set {
					for yt := range row {
						visited.Set(x, y, xt, yt)
					}
				}
			}
		}
		if x < len(tm)-1 && tm[x+1][y] == target {
			fn(x+1, y, visited, target)
			set, ok := visited.Get(x+1, y)
			if ok {
				for xt, row := range set {
					for yt := range row {
						visited.Set(x, y, xt, yt)
					}
				}
			}
		}
		if y < len(tm[x])-1 && tm[x][y+1] == target {
			fn(x, y+1, visited, target)
			set, ok := visited.Get(x, y+1)
			if ok {
				for xt, row := range set {
					for yt := range row {
						visited.Set(x, y, xt, yt)
					}
				}
			}
		}
	}
	fn(x, y, visited, '0')
	set, _ := visited.Get(x, y)
	return set.Len()
}

func (tm TopographicMap) CountReachableNines2(x, y int, visited PositionCounterSet) int {
	var fn func(x, y int, visited PositionCounterSet, current byte) int
	fn = func(x, y int, visited PositionCounterSet, current byte) int {
		if current == '9' {
			visited.Set(x, y, 0)
			return 1
		}
		count, ok := visited.Get(x, y)
		if ok {
			return count
		} else {
			count = 0
		}

		target := current + 1

		if x > 0 && tm[x-1][y] == target {
			count += fn(x-1, y, visited, target)
		}
		if y > 0 && tm[x][y-1] == target {
			count += fn(x, y-1, visited, target)
		}
		if x < len(tm)-1 && tm[x+1][y] == target {
			count += fn(x+1, y, visited, target)
		}
		if y < len(tm[x])-1 && tm[x][y+1] == target {
			count += fn(x, y+1, visited, target)
		}

		visited.Set(x, y, count)
		return count
	}
	return fn(x, y, visited, '0')
}
