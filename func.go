// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-15 11:55:50 457CCB                             go-delta/[func.go]
// -----------------------------------------------------------------------------

package bdelta

import (
	"bytes"
	"fmt"
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

// MakeMap create a map of unique chunks in 'data'.
// The key specifies the unique chunk of bytes, while the
// values array returns the positions of the chunks in 'data'.
func MakeMap(data []byte) (ret map[[ChunkSize]byte][]int) {
	ret = make(map[[ChunkSize]byte][]int, 0)
	if len(data) < ChunkSize {
		return ret
	}
	var key [ChunkSize]byte
	for i := 0; i < len(data)-ChunkSize; i++ {
		copy(key[:], data[i:])
		var _, exist = ret[key]
		if exist {
			ret[key] = append(ret[key], i)
		} else {
			ret[key] = []int{i}
		}
	}
	return ret
} //                                                                     MakeMap

//end
