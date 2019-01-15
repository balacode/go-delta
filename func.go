// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-15 23:18:20 F01335                             go-delta/[func.go]
// -----------------------------------------------------------------------------

package bdelta

import (
	"bytes"
	"crypto/sha512"
	"fmt"
	"github.com/balacode/zr"
)

const ChunkSize = 8
const Direct = -1

var PL = fmt.Println

// ApplyDiff __
func ApplyDiff(source []byte, diff Diff) []byte {
	return []byte{}
} //                                                                   ApplyDiff

// MakeDiff given two byte arrays 'a' and 'b', calculates the binary
// delta difference between the two arrays and returns it as a Diff.
// You can then use ApplyDiff() to generate 'b' from 'a' the Diff.
func MakeDiff(a, b []byte) Diff {
	var ret = Diff{targetHash: makeHash(b)}
	if len(b) < ChunkSize {
		ret.parts = []diffPart{{sourceLoc: Direct, size: len(b), data: b}}
		return ret
	}
	var m = makeMap(a)
	var chunk [ChunkSize]byte
	for i, end := 0, len(b); i < end; {
		if end-i < ChunkSize {
			ret.writePart(Direct, end-i, b[i:])
			ret.newCount++
			break
		}
		var locs []int
		var found = false
		if end-i >= ChunkSize {
			copy(chunk[:], b[i:])
			locs, found = m[chunk]
		}
		if found {
			var at, size = longestMatch(a, locs, b, i)
			ret.writePart(at, size, a[at:at+size])
			i += size
			ret.oldCount++
			continue
		}
		ret.writePart(Direct, ChunkSize, chunk[:])
		i += ChunkSize
		ret.newCount++
	}
	ret.sourceHash = makeHash(a)
	return ret
} //                                                                    MakeDiff

// -----------------------------------------------------------------------------
// # Helper Functions

// longestMatch __
func longestMatch(a []byte, aLocs []int, b []byte, bLoc int) (loc, size int) {
	if len(aLocs) < 1 {
		zr.Error("aLocs is empty")
		return -1, -1
	}
	var bEnd = len(b) - 1
	if bLoc < 0 || bLoc > bEnd {
		zr.Error("bLoc", bLoc, "out of range [0 -", len(b), "]")
		return -1, -1
	}
	var aEnd = len(a) - 1
	var retLoc = -1
	var retSize = -1
	for _, aLoc := range aLocs {
		var n = ChunkSize
		if !bytes.Equal(a[aLoc:aLoc+n], b[bLoc:bLoc+n]) {
			zr.Error("mismatch at aLoc:", aLoc, "bLoc:", bLoc)
			continue
		}
		// extend match forward
		for aLoc+n <= aEnd && bLoc+n <= bEnd && a[aLoc+n] == b[bLoc+n] {
			n++
		}
		if n > retSize {
			retLoc = aLoc
			retSize = n
		}
	}
	return retLoc, retSize
} //                                                                longestMatch

// makeHash returns the SHA-512 hash of byte slice 'data'.
func makeHash(data []byte) []byte {
	var ret = sha512.Sum512(data)
	return ret[:]
} //                                                                    makeHash

// makeMap creates a map of unique chunks in 'data'.
// The key specifies the unique chunk of bytes, while the
// values array returns the positions of the chunk in 'data'.
func makeMap(data []byte) (ret map[[ChunkSize]byte][]int) {
	ret = make(map[[ChunkSize]byte][]int, 0)
	if len(data) < ChunkSize {
		return ret
	}
	var key [ChunkSize]byte
	for i := 0; i < len(data)-ChunkSize; i++ {
		copy(key[:], data[i:])
		var _, found = ret[key]
		if found {
			ret[key] = append(ret[key], i)
		} else {
			ret[key] = []int{i}
		}
	}
	return ret
} //                                                                     makeMap

//end
