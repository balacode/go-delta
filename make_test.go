// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-28 21:39:43 F247CA                        go-delta/[make_test.go]
// -----------------------------------------------------------------------------

package delta

import (
	"testing"
)

// go test --run Test_Make_
func Test_Make_(t *testing.T) {
	if PrintTestNames {
		printTestName()
	}
	// func Make(a, b []byte) Delta
	//
	test := func(a, b []byte, expect Delta) {
		result := Make(a, b)
		if result.GoString() != expect.GoString() {
			t.Errorf("\n expect:\n\t%s\n result:\n\t%s\n",
				expect.GoString(), result.GoString())
		}
	}
	test(
		ab(AtoZ),
		ab(AtoZ),
		Delta{
			sourceSize: 26,
			sourceHash: makeHash(ab(AtoZ)),
			targetSize: 26,
			targetHash: makeHash(ab(AtoZ)),
			newCount:   0,
			oldCount:   1,
			parts: []deltaPart{
				{sourceLoc: 0, size: 26, data: nil},
			},
		},
	)
} //                                                                  Test_Make_

//end
