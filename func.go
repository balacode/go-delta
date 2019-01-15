// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-15 15:28:55 6ECE56                             go-delta/[func.go]
// -----------------------------------------------------------------------------

package bdelta

import (
	"bytes"
	"fmt"
	"github.com/balacode/zr"
)

const ChunkSize = 8

type Diff []byte

var PL = fmt.Println

// ApplyDiff __
func ApplyDiff(source []byte, diff Diff) []byte {
	return []byte{}
} //                                                                   ApplyDiff

// MakeDiff __
func MakeDiff(a, b []byte) Diff {
	if len(b) < ChunkSize {
		return Diff{}
	}
	var nfound = 0
	var nmiss = 0
	var chunk [ChunkSize]byte
	for i := 0; i < len(b)-ChunkSize; i += 1024 {
		copy(chunk[:], b[i:])
		if bytes.Index(a, chunk[:]) == -1 {
			nmiss++
		} else {
			nfound++
		}
	}
	PL("nfound:", nfound)
	PL("nmiss:", nmiss)
	return Diff{}
} //                                                                    MakeDiff

// -----------------------------------------------------------------------------
// # Helper Functions

// LongestMatch __
func LongestMatch(a []byte, aLocs []int, b []byte, bLoc int) (loc, size int) {
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
	var retLoc, retSize = -1, -1
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
} //                                                                LongestMatch

// MakeMap create a map of unique chunks in 'data'.
// The key specifies the unique chunk of bytes, while the
// values array returns the positions of the chunk in 'data'.
func MakeMap(data []byte) (ret map[[ChunkSize]byte][]int) {
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
} //                                                                     MakeMap

//end
