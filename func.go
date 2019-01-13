// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-13 18:12:48 B32132                             go-delta/[func.go]
// -----------------------------------------------------------------------------

package bdelta

type Diff []byte

// ApplyDiff __
func ApplyDiff(source []byte, diff Diff) []byte {
	return []byte{}
} //                                                                   ApplyDiff

// MakeDiff __
func MakeDiff(a, b []byte) Diff {
	return Diff{}
} //                                                                    MakeDiff

const ChunkSize = 8

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
