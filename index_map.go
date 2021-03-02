// -----------------------------------------------------------------------------
// github.com/balacode/go-delta                          go-delta/[index_map.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package delta

const DebugIndex = false

type chunk [MatchSize]byte

// indexMap _ _
type indexMap struct {
	m map[chunk][]int
} //                                                                    indexMap

// makeMap creates a map of unique chunks in 'data'.
// The key specifies the unique chunk of bytes, while the
// values array returns the positions of the chunk in 'data'.
func makeMap(data []byte) indexMap {
	if DebugTiming {
		tmr.Start("makeMap")
		defer tmr.Stop("makeMap")
	}
	if DebugIndex {
		PL("makeMap init:", len(data), "bytes")
	}
	lenData := len(data)
	if lenData < MatchSize {
		return indexMap{m: map[chunk][]int{}}
	}
	dbgN := 0
	ret := indexMap{m: make(map[chunk][]int, lenData/4)}
	var key chunk
	lenData -= MatchSize
	if DebugIndex {
		PL("makeMap begin loop")
	}
	for i := 0; i < lenData; {
		copy(key[:], data[i:])
		ar, found := ret.m[key]
		if !found {
			ret.m[key] = []int{i}
			i++
			continue
		}
		if len(ar) >= MatchLimit {
			i++
			continue
		}
		ret.m[key] = append(ret.m[key], i)
		i += MatchSize
		if DebugIndex {
			dbgN++
			if dbgN < 10e6 {
				continue
			}
			dbgN = 0
			PL("i:", i, "len(m):", len(ret.m))
		}
	}
	return ret
} //                                                                     makeMap

// get _ _
func (ob *indexMap) get(key chunk) (locs []int, found bool) {
	locs, found = ob.m[key]
	return
} //                                                                         get

// end
