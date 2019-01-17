// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-17 14:52:20 99E81A                             go-delta/[func.go]
// -----------------------------------------------------------------------------

package bdelta

import (
	"bytes"
	"compress/zlib"
	"crypto/sha512"
	"io"
)

// ApplyDiff __
func ApplyDiff(source []byte, diff Diff) []byte {
	if DebugTiming {
		tmr.Start("ApplyDiff")
		defer tmr.Stop("ApplyDiff")
	}
	return []byte{}
} //                                                                   ApplyDiff

// MakeDiff given two byte arrays 'a' and 'b', calculates the binary
// delta difference between the two arrays and returns it as a Diff.
// You can then use ApplyDiff() to generate 'b' from 'a' the Diff.
func MakeDiff(a, b []byte) Diff {
	if DebugTiming {
		tmr.Start("MakeDiff")
		defer tmr.Stop("MakeDiff")
	}
	var ret = Diff{targetHash: makeHash(b)}
	var lenB = len(b)
	if lenB < MatchSize {
		ret.parts = []diffPart{{sourceLoc: -1, size: lenB, data: b}}
		return ret
	}
	var m = makeMap(a)
	var chunk [MatchSize]byte
	var tmc = 0 // timing counter
	for i := 0; i < lenB; {
		if DebugInfo && i-tmc >= 10000 {
			PL("MakeDiff:", int(100.0/float32(lenB)*float32(i)), "%")
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
		PL("MakeDiff: finished writing parts. len(b) = ", len(b))
	}
	ret.sourceHash = makeHash(a)
	return ret
} //                                                                    MakeDiff

// -----------------------------------------------------------------------------
// # Helper Functions: Compression

// compressBytes compresses an array of bytes and
// returns the ZLIB-compressed array of bytes.
func compressBytes(data []byte) []byte {
	if DebugTiming {
		tmr.Start("compressBytes")
		defer tmr.Stop("compressBytes")
	}
	if len(data) == 0 {
		return []byte{}
	}
	// zip data in standard manner
	var b bytes.Buffer
	var w = zlib.NewWriter(&b)
	var _, err = w.Write(data)
	w.Close()
	//
	// log any problem
	const ERRM = "Failed compressing data with zlib:"
	if err != nil {
		mod.Error(ERRM, err)
		return []byte{}
	}
	var ret = b.Bytes()
	if len(ret) < 3 {
		mod.Error(ERRM, "length < 3")
		return []byte{}
	}
	return ret
} //                                                               compressBytes

// uncompressBytes uncompresses a ZLIB-compressed array of bytes.
func uncompressBytes(data []byte) []byte {
	var readCloser, err = zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		mod.Error("uncompressBytes:", err)
		return []byte{}
	}
	var ret = bytes.NewBuffer(make([]byte, 0, 8192))
	io.Copy(ret, readCloser)
	readCloser.Close()
	return ret.Bytes()
} //                                                             uncompressBytes

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

// makeHash returns the SHA-512 hash of byte slice 'data'.
func makeHash(data []byte) []byte {
	if DebugTiming {
		tmr.Start("makeHash")
		defer tmr.Stop("makeHash")
	}
	var ret = sha512.Sum512(data)
	return ret[:]
} //                                                                    makeHash

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
