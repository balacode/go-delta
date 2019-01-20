// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-20 07:16:23 66DDEC                             go-delta/[make.go]
// -----------------------------------------------------------------------------

package delta

import (
	"bytes"
)

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

// -----------------------------------------------------------------------------
// # Helper Functions

// longestMatch __
func longestMatch(a []byte, aLocs []int, b []byte, bLoc int) (loc, size int) {
	if DebugTiming {
		tmr.Start("longestMatch")
		defer tmr.Stop("longestMatch")
	}
	if len(aLocs) < 1 {
		mod.Error("aLocs is empty")
		return -1, -1
	}
	var bEnd = len(b) - 1
	if bLoc < 0 || bLoc > bEnd {
		mod.Error("bLoc", bLoc, "out of range [0 -", len(b), "]")
		return -1, -1
	}
	var aEnd = len(a) - 1
	var retLoc = -1
	var retSize = -1
	for _, ai := range aLocs {
		var n = MatchSize
		var bi = bLoc
		if !bytes.Equal(a[ai:ai+n], b[bi:bi+n]) {
			mod.Error("mismatch at ai:", ai, "bi:", bi)
			continue
		}
		/*
			DISABLED: EXTENDING MATCH BACKWARD OVERLAPS PREVIOUSLY-WRITTEN PARTS
			// extend match backward
			for ai-1 >= 0 && bi-1 >= 0 && a[ai-1] == b[bi-1] {
				ai--
				bi--
				n++
			}
		*/
		// extend match forward
		for ai+n <= aEnd && bi+n <= bEnd && a[ai+n] == b[bi+n] {
			n++
		}
		if n > retSize {
			retLoc = ai
			retSize = n
		}
	}
	return retLoc, retSize
} //                                                                longestMatch

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
