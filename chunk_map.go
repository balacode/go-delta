// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-02-05 16:24:59 DA513D                        go-delta/[chunk_map.go]
// -----------------------------------------------------------------------------

package delta

type Chunk [MatchSize]byte

// newChunkMap creates a map of unique chunks in 'data'.
// The key specifies the unique chunk of bytes, while the
// values array returns the positions of the chunk in 'data'.
func newChunkMap(data []byte) map[Chunk][]int {
	if DebugTiming {
		tmr.Start("newChunkMap")
		defer tmr.Stop("newChunkMap")
	}
	var lenData = len(data)
	if lenData < MatchSize {
		return map[Chunk][]int{}
	}
	var ret = make(map[Chunk][]int, lenData/4)
	var key Chunk
	lenData -= MatchSize
	for i := 0; i < lenData; {
		copy(key[:], data[i:])
		var ar, found = ret[key]
		if !found {
			ret[key] = []int{i}
			i++
			continue
		}
		if len(ar) >= MatchLimit {
			i++
			continue
		}
		ret[key] = append(ret[key], i)
		i += MatchSize
	}
	return ret
} //                                                                 newChunkMap

//end
