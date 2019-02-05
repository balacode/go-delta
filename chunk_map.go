// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-02-05 16:39:50 3F6F64                        go-delta/[chunk_map.go]
// -----------------------------------------------------------------------------

package delta

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
	var lenData = len(data)
	if lenData < MatchSize {
		return chunkMap{m: map[chunk][]int{}}
	}
	var ret = chunkMap{m: make(map[chunk][]int, lenData/4)}
	var key chunk
	lenData -= MatchSize
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
	}
	return ret
} //                                                                 newChunkMap

//end
