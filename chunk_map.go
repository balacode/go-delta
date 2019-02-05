// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-02-05 16:50:31 73E4C3                        go-delta/[chunk_map.go]
// -----------------------------------------------------------------------------

package delta

const DebugChunkMap = false

type chunk [MatchSize]byte

// chunkMap __
type chunkMap struct {
	m map[chunk][]int
} //                                                                    chunkMap

// newChunkMap creates a map of unique chunks in 'data'.
// The key specifies the unique chunk of bytes, while the
// values array returns the positions of the chunk in 'data'.
func newChunkMap(data []byte) chunkMap {
	if DebugTiming {
		tmr.Start("newChunkMap")
		defer tmr.Stop("newChunkMap")
	}
	if DebugChunkMap {
		PL("newChunkMap init:", len(data), "bytes")
	}
	var lenData = len(data)
	if lenData < MatchSize {
		return chunkMap{m: map[chunk][]int{}}
	}
	var dbgN = 0
	var ret = chunkMap{m: make(map[chunk][]int, lenData/4)}
	var key chunk
	lenData -= MatchSize
	if DebugChunkMap {
		PL("newChunkMap begin loop")
	}
	for i := 0; i < lenData; {
		copy(key[:], data[i:])
		var ar, found = ret.m[key]
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
		if DebugChunkMap {
			dbgN++
			if dbgN < 10E6 {
				continue
			}
			dbgN = 0
			PL("i:", i, "len(m):", len(ret.m))
		}
	}
	return ret
} //                                                                 newChunkMap

// get __
func (ob *chunkMap) get(key chunk) (locs []int, found bool) {
	locs, found = ob.m[key]
	return
} //                                                                         get

//end
