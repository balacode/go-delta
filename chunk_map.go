// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-02-05 16:18:45 9EFD0B                        go-delta/[chunk_map.go]
// -----------------------------------------------------------------------------

package delta

// makeMap creates a map of unique chunks in 'data'.
// The key specifies the unique chunk of bytes, while the
// values array returns the positions of the chunk in 'data'.
func makeMap(data []byte) map[[MatchSize]byte][]int {
	if DebugTiming {
		tmr.Start("makeMap")
		defer tmr.Stop("makeMap")
	}
	var lenData = len(data)
	if lenData < MatchSize {
		return map[[MatchSize]byte][]int{}
	}
	var ret = make(map[[MatchSize]byte][]int, lenData/4)
	var key [MatchSize]byte
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
} //                                                                     makeMap

//end
