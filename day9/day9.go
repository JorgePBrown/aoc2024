package day9

import (
	"io"
)

func SolvePart1(r io.Reader) (int, error) {
	dm := NewDiskMap(r)

	dm.Compact()

	return dm.Checksum(), nil
}

func SolvePart2(r io.Reader) (int, error) {
	dm := NewDiskMap(r)

	dm.Compact2()

	return dm.Checksum2(), nil
}

type DiskMap struct {
	data []*DiskMapEntry
}

type DiskMapEntry struct {
	free bool
	id   int
}

func NewDiskMap(r io.Reader) *DiskMap {
	buffer := make([]byte, 1024)
	free := false
	id := 0
	dm := DiskMap{
		data: []*DiskMapEntry{},
	}
	for {
		n, err := r.Read(buffer)

		for i := 0; i < n; i += 1 {
			c := buffer[i] - '0'
			if c < 0 || c > 9 {
				return &dm
			}

			for _ = range c {
				dm.data = append(dm.data, &DiskMapEntry{
					free: free,
					id:   id,
				})
			}
			if !free {
				id += 1
			}
			free = !free
		}

		if err != nil {
			break
		}
	}
	return &dm
}

func (dm *DiskMap) Compact() {
	freeIdx := dm.IndexFree(0)
	nonFreeIdx := dm.LastIndexNonFree(len(dm.data))

	for freeIdx > -1 && nonFreeIdx > 1 && freeIdx < nonFreeIdx {
		swap := dm.data[freeIdx]
		dm.data[freeIdx] = dm.data[nonFreeIdx]
		dm.data[nonFreeIdx] = swap

		freeIdx = dm.IndexFree(freeIdx + 1)
		nonFreeIdx = dm.LastIndexNonFree(nonFreeIdx - 1)
	}
}

func (dm *DiskMap) Checksum() int {
	checksum := 0
	for pos, e := range dm.data {
		if e.free {
			break
		}
		checksum += pos * e.id
	}
	return checksum
}

func (dm *DiskMap) Compact2() {
	idx := dm.LastIndexNonFree(len(dm.data))
	if idx == -1 {
		return
	}
	id := dm.data[idx].id + 1
	for idx >= 0 {
		newId := dm.data[idx].id
		if newId < id {
			swapped := dm.CompactOnce(idx)
			if swapped {
				idx = dm.LastIndexNonFree(idx)
			} else {
				start, _ := dm.IndexSegment(idx, newId, false)
				idx = dm.LastIndexNonFree(start - 1)
			}
			id = newId
		} else {
			start, _ := dm.IndexSegment(idx, newId, false)
			idx = dm.LastIndexNonFree(start - 1)
		}
	}
}

func (dm *DiskMap) CompactOnce(nonFreeEndIdx int) bool {
	nonFreeStartIdx, nonFreeEndIdx := dm.IndexSegment(nonFreeEndIdx, dm.data[nonFreeEndIdx].id, false)
	nonFreeLen := nonFreeEndIdx - nonFreeStartIdx + 1

	if nonFreeStartIdx < 0 || nonFreeEndIdx < 0 {
		return false
	}

	freeStartIdx := dm.IndexFree(0)
	for freeStartIdx > -1 && freeStartIdx < nonFreeStartIdx {
		freeEndIdx := dm.IndexNonFree(freeStartIdx)
		freeLen := freeEndIdx - freeStartIdx
		if freeLen >= nonFreeLen {
			for i := range nonFreeLen {
				swap := dm.data[freeStartIdx+i]
				dm.data[freeStartIdx+i] = dm.data[nonFreeStartIdx+i]
				dm.data[nonFreeStartIdx+i] = swap
			}
			return true
		}

		freeStartIdx = dm.IndexFree(freeEndIdx + 1)
	}

	return false
}

func (dm *DiskMap) Checksum2() int {
	checksum := 0
	for pos, e := range dm.data {
		if !e.free {
			checksum += pos * e.id
		}
	}
	return checksum
}

func (dm *DiskMap) IndexFree(start int) int {
	for i := start; i < len(dm.data); i += 1 {
		v := dm.data[i]
		if v.free {
			return i
		}
	}
	return -1
}
func (dm *DiskMap) IndexNonFree(start int) int {
	for i := start; i < len(dm.data); i += 1 {
		v := dm.data[i]
		if !v.free {
			return i
		}
	}
	return -1
}

func (dm *DiskMap) LastIndexNonFree(start int) int {
	if start >= len(dm.data) {
		start = len(dm.data) - 1
	}
	for i := start; i >= 0; i -= 1 {
		v := dm.data[i]
		if !v.free {
			return i
		}
	}
	return -1
}

func (dm *DiskMap) IndexSegment(i int, id int, free bool) (int, int) {
	if i < 0 || i >= len(dm.data) {
		panic("start over length of data")
	}

	var start int
	for start = i; start >= 0; start -= 1 {
		v := dm.data[start]
		if free != v.free {
			break
		}
		if !free && id != v.id {
			break
		}
	}

	var end int
	for end = i; end < len(dm.data); end += 1 {
		v := dm.data[end]
		if free != v.free {
			break
		}
		if !free && id != v.id {
			break
		}
	}
	return start + 1, end - 1
}
