// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-14 19:18:48 E18215                             go-delta/[func.go]
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

// MakeMap __
func MakeMap(ar []byte) (m map[[ChunkSize]byte][]int) {
	m = make(map[[ChunkSize]byte][]int, 0)
	if len(ar) < ChunkSize {
		return
	}
	var k [ChunkSize]byte
	for i := 0; i < len(ar)-ChunkSize; i++ {
		copy(k[:], ar[i:])
		var _, exist = m[k]
		if exist {
			m[k] = append(m[k], i)
		} else {
			m[k] = []int{i}
		}
	}
	return m
} //                                                                     MakeMap

//end
