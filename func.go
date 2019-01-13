// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-13 18:06:53 C7CD93                             go-delta/[func.go]
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

func MakeMap(ar []byte) (m map[[8]byte][]int) {
	m = make(map[[8]byte][]int, 0)
	if len(ar) < 8 {
		return
	}
	var k [8]byte
	for i := 0; i < len(ar)-8; i++ {
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
