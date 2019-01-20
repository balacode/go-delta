// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-20 13:55:14 82D8D5                 go-delta/[delta_apply_test.go]
// -----------------------------------------------------------------------------

package delta

import (
	"bytes"
	"testing"
)

// go test --run Test_Delta_Apply_
func Test_Delta_Apply_(t *testing.T) {
	if PrintTestNames {
		printTestName()
	}
	var test = func(src []byte, d Delta, expect []byte) {
		var result, err = d.Apply(src)
		if err != nil {
			t.Errorf("\n encountered error: %s\n", err)
			return
		}
		if !bytes.Equal(result, expect) {
			t.Errorf("\n expect:\n\t%v\n\t'%s'\n result:\n\t%v\n\t'%s'\n",
				expect, expect, result, result)
		}
	}
	test(
		// source:
		nil,
		//
		// delta:
		Delta{
			sourceHash: nil,
			targetHash: makeHash(ab("abc")),
			parts: []deltaPart{
				{sourceLoc: -1, size: 3, data: ab("abc")},
			},
		},
		// expect:
		ab("abc"),
	)
	test(
		// source:
		ab("abc"),
		//
		// delta:
		Delta{
			sourceHash: makeHash(ab("abc")),
			sourceSize: 3,
			targetHash: makeHash(ab("abc")),
			targetSize: 3,
			parts: []deltaPart{
				{sourceLoc: -1, size: 3, data: ab("abc")},
			},
		},
		// expect:
		ab("abc"),
	)
} //                                                           Test_Delta_Apply_

//end
