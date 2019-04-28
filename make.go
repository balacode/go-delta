// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-28 21:39:43 C9C4BE                             go-delta/[make.go]
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
	ret := Delta{
		sourceSize: len(a),
		sourceHash: makeHash(a),
		targetSize: len(b),
		targetHash: makeHash(b),
	}
	lenB := len(b)
	if lenB < MatchSize {
		ret.parts = []deltaPart{{sourceLoc: -1, size: lenB, data: b}}
		return ret
	}
	cmap := newIndexMap(a)
	var key chunk
	tmc := 0 // timing counter
	for i := 0; i < lenB; {
		if DebugInfo && i-tmc >= 10000 {
			PL("delta.Make:", int(100.0/float32(lenB)*float32(i)), "%")
			tmc = i
		}
		if lenB-i < MatchSize {
			ret.write(-1, lenB-i, b[i:])
			ret.newCount++
			break
		}
		var locs []int
		found := false
		if lenB-i >= MatchSize {
			copy(key[:], b[i:])
			locs, found = cmap.get(key)
		}
		if found {
			at, size := longestMatch(a, locs, b, i)
			ret.write(at, size, nil)
			i += size
			ret.oldCount++
			continue
		}
		ret.write(-1, MatchSize, key[:])
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

// longestMatch is called by Make() to determine the longest
// matching block of bytes between the source array 'a'
// and target array 'b' out of limited choices.
//
// 'bLoc' specifies the position (in 'b') of the chunk to match.
// The MatchSize global constant specifies the length of each
// chunk in bytes, usually 8 bytes.
//
// 'aLocs' is an array of positions (in 'a') at which the chunk is found.
// This array is produced by newIndexMap() before longestMatch() is called.
//
// Returns the location ('loc') of the match in 'a'
// and the length of the match in 'b' ('size').
//
func longestMatch(a []byte, aLocs []int, b []byte, bLoc int) (loc, size int) {
	if DebugTiming {
		tmr.Start("longestMatch")
		defer tmr.Stop("longestMatch")
	}
	if len(aLocs) < 1 {
		mod.Error("aLocs is empty")
		return -1, -1
	}
	bEnd := len(b) - 1
	if bLoc < 0 || bLoc > bEnd {
		mod.Error("bLoc", bLoc, "out of range [0 -", len(b), "]")
		return -1, -1
	}
	var (
		aEnd    = len(a) - 1
		retLoc  = -1
		retSize = -1
	)
	for _, ai := range aLocs {
		n := MatchSize
		bi := bLoc
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

//end
