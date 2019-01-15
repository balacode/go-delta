// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-15 18:58:56 94BCA1                             go-delta/[func.go]
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
	if len(b) < ChunkSize {
		/// WRITE MODE 0 TO RETURNED VALUE
		/// WRITE b TO RETURNED VALUE
		return Diff{}
	}
	var m = makeMap(a)
	var nmatch = 0
	var nmiss = 0
	var step = 1024
	if step > len(b) {
		step = 1
	}
	var i = 0
	var end = len(b)
	var chunk [ChunkSize]byte
	for i < end {
		if end-i < ChunkSize {
			writeDiff(Direct, end-i, b[i:])
			nmiss++
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
			writeDiff(at, size, a[at:at+size])
			i += size
			nmatch++
			continue
		}
		writeDiff(Direct, ChunkSize, chunk[:])
		i += ChunkSize
		nmiss++
	}
	PL("nmatch:", nmatch)
	PL("nmiss:", nmiss)
	return Diff{}
} //                                                                    MakeDiff

// -----------------------------------------------------------------------------
// # Helper Functions

// hashOfBytes returns the SHA-512 hash of a byte slice.
// It also requires a 'salt' argument.
func hashOfBytes(ar []byte, salt []byte) []byte {
	var input []byte
	input = append(input, salt[:]...)
	input = append(input, ar...)
	var hash = sha512.Sum512(input)
	return hash[:]
} //                                                                 hashOfBytes

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
} //                                                                longestMatch

// makeMap create a map of unique chunks in 'data'.
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

// writeDiff __
func writeDiff(offset, size int, data []byte) {
	PL("WD", "offset:", offset, "size:", size, "data:", data, string(data))
} //                                                                   writeDiff

//end
