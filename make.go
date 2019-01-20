// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-20 06:55:02 51DDEA                             go-delta/[make.go]
// -----------------------------------------------------------------------------

package delta

// Make given two byte arrays 'a' and 'b', calculates the binary
// delta difference between the two arrays and returns it as a Delta.
// You can then use Delta.Apply() to generate 'b' from 'a' the Delta.
func Make(a, b []byte) Delta {
	if DebugTiming {
		tmr.Start("delta.Make")
		defer tmr.Stop("delta.Make")
	}
	var ret = Delta{
		sourceSize: len(a),
		sourceHash: makeHash(a),
		targetSize: len(b),
		targetHash: makeHash(b),
	}
	var lenB = len(b)
	if lenB < MatchSize {
		ret.parts = []deltaPart{{sourceLoc: -1, size: lenB, data: b}}
		return ret
	}
	var m = makeMap(a)
	var chunk [MatchSize]byte
	var tmc = 0 // timing counter
	for i := 0; i < lenB; {
		if DebugInfo && i-tmc >= 10000 {
			PL("delta.Make:", int(100.0/float32(lenB)*float32(i)), "%")
			tmc = i
		}
		if lenB-i < MatchSize {
			ret.appendPart(-1, lenB-i, b[i:])
			ret.newCount++
			break
		}
		var locs []int
		var found = false
		if lenB-i >= MatchSize {
			copy(chunk[:], b[i:])
			locs, found = m[chunk]
		}
		if found {
			var at, size = longestMatch(a, locs, b, i)
			ret.appendPart(at, size, nil)
			i += size
			ret.oldCount++
			continue
		}
		ret.appendPart(-1, MatchSize, chunk[:])
		i += MatchSize
		ret.newCount++
	}
	if DebugInfo {
		PL("delta.Make: finished writing parts. len(b) = ", len(b))
	}
	return ret
} //                                                                        Make

//end
