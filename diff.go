// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-16 00:12:51 7A0BF8                             go-delta/[diff.go]
// -----------------------------------------------------------------------------

package bdelta

import (
	"github.com/balacode/zr"
)

// Diff stores the binary delta difference between two byte arrays
type Diff struct {
	sourceHash []byte
	// hash of the source byte array
	//
	targetHash []byte
	// expected hash of result after this Diff is applied to source
	//
	parts []diffPart
	// array of referring to chunks in source array, or new bytes to append
	//
	newCount int
	// number of chunks that could not be matched in source
	//
	oldCount int
	// number of chunks that were matched in source
} //                                                                        Diff

// diffPart stores references to chunks in the source array,
// or specifies bytes to append to result array directly
type diffPart struct {
	sourceLoc int
	// byte position of the chunk in source array,
	// or -1 when the bytes should be picked from 'data'
	//
	size int
	// size of the chunk in bytes
	//
	data []byte
	// optional bytes (only when sourceLoc is -1)
} //                                                                    diffPart

// -----------------------------------------------------------------------------
// # Public Properties

// NewCount __
func (ob *Diff) NewCount() int {
	return ob.newCount
} //                                                                    NewCount

// OldCount __
func (ob *Diff) OldCount() int {
	return ob.oldCount
} //                                                                    OldCount

// -----------------------------------------------------------------------------
// # Internal Methods

// appendPart appends binary difference data
func (ob *Diff) appendPart(sourceLoc, size int, data []byte) {
	if true {
		PL("appendPart",
			"sourceLoc:", sourceLoc,
			"size:", size,
			"data:", data, string(data))
	}
	// argument validations
	switch {
	case sourceLoc < -1:
		zr.Error("sourceLoc:", sourceLoc, " < -1")
		return
	case sourceLoc == -1 && len(data) == 0:
		zr.Error("sourceLoc == -1 && len(data) == 0")
		return
	case sourceLoc != -1 && len(data) != 0:
		zr.Error("sourceLoc != -1 && len(data):", len(data), "!= 0")
		return
	case size < 1:
		zr.Error("size:", size, " < 1")
		return
	}
	// if the previous part was embedded directly, append to that part's data
	if sourceLoc == -1 {
		var n = len(ob.parts)
		if n > 0 {
			var last = &ob.parts[n-1]
			if last.sourceLoc == -1 {
				last.size += len(data)
				last.data = append(last.data, data...)
				return
			}
		}
	}
	// append a new part
	var ar []byte
	if sourceLoc != -1 {
		var ar = make([]byte, len(data))
		copy(ar, data)
	}
	ob.parts = append(ob.parts,
		diffPart{sourceLoc: sourceLoc, size: size, data: ar})
} //                                                                  appendPart

//end