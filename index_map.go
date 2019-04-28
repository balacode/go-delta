// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-28 21:39:43 F20099                        go-delta/[index_map.go]
// -----------------------------------------------------------------------------

package delta

const DebugIndex = false

type chunk [MatchSize]byte

// indexMap __
type indexMap struct {
	m map[chunk][]int
} //                                                                    indexMap

// newIndexMap creates a map of unique chunks in 'data'.
// The key specifies the unique chunk of bytes, while the
// values array returns the positions of the chunk in 'data'.
func newIndexMap(data []byte) indexMap {
	if DebugTiming {
		tmr.Start("newIndexMap")
		defer tmr.Stop("newIndexMap")
	}
	if DebugIndex {
		PL("newIndexMap init:", len(data), "bytes")
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
		PL("newIndexMap begin loop")
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
			if dbgN < 10E6 {
				continue
			}
			dbgN = 0
			PL("i:", i, "len(m):", len(ret.m))
		}
	}
	return ret
} //                                                                 newIndexMap

// get __
func (ob *indexMap) get(key chunk) (locs []int, found bool) {
	locs, found = ob.m[key]
	return
} //                                                                         get

//end
