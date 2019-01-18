// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-18 17:53:01 A7C5CF                        go-delta/[func_test.go]
// -----------------------------------------------------------------------------

package bdelta

/*
to test all items in dir_watcher_windows.go use:
    go test --run Test_gdlt_

to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

const AtoM = "ABCDEFGHIJKLM"
const AtoS = "ABCDEFGHIJKLMNOPQRS"
const AtoZ = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const Nums = "0123456789"
const atoz = "abcdefghijklmnopqrstuvwxyz"

var Line = strings.Repeat("#", 70)

// go test --run Test_MakeDiff_
func Test_MakeDiff_(t *testing.T) {
	//
	// func MakeDiff(a, b []byte) Diff
	//
	var test = func(a, b []byte, expect Diff) {
		var result = MakeDiff(a, b)
		if result.GoString() != expect.GoString() {
			t.Errorf("\n expect:\n\t%s\n result:\n\t%s\n",
				expect.GoString(), result.GoString())
		}
	}
	test(
		ab(AtoZ),
		ab(AtoZ),
		Diff{
			sourceHash: makeHash(ab(AtoZ)),
			targetHash: makeHash(ab(AtoZ)),
			newCount:   0,
			oldCount:   1,
			parts: []diffPart{
				{sourceLoc: 0, size: 26, data: nil},
			},
		},
	)
} //                                                              Test_MakeDiff_

// go test --run Test_ApplyDiff_
func Test_ApplyDiff_(t *testing.T) {
	var test = func(src []byte, d Diff, expect []byte) {
		var result, err = ApplyDiff(src, d)
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
		nil,
		Diff{
			sourceHash: nil,
			targetHash: makeHash(ab("abc")),
			parts: []diffPart{
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
		// diff:
		Diff{
			sourceHash: makeHash(ab("abc")),
			sourceSize: 3,
			targetHash: makeHash(ab("abc")),
			targetSize: 3,
			parts: []diffPart{
				{sourceLoc: -1, size: 3, data: ab("abc")},
			},
		},
		// expect:
		ab("abc"),
	)
} //                                                             Test_ApplyDiff_

// -----------------------------------------------------------------------------
// # Auxiliary / Temporary Unit Tests

// go test --run Test_01_
func Test_01_(t *testing.T) {
	PL("Test_01_ " + Line)
	//
	var m1 = makeMap(readData("test1.zip"))
	PL("Created m1. len(m1):", len(m1))
	//
	var m2 = makeMap(readData("test2.zip"))
	PL("Created m2. len(m2):", len(m2))
	//
	if false {
		const MaxLines = 0
		var i = 1
		for k, v := range m1 {
			PL("key:", k, "val:", v)
			i++
			if i > MaxLines {
				break
			}
		}
	}
	if true {
		for k, v := range m2 {
			_, exist := m1[k]
			PL("key:", k, "val:", v, "exist:", exist)
		}
	}
} //                                                                    Test_01_

// go test --run Test_02_
func Test_02_(t *testing.T) {
	var a, b []byte
	switch 5 {
	case 1:
		a = ab(AtoM + " " + AtoS + " " + AtoZ)
		b = ab("0x0x0x" + AtoZ + " " + AtoZ + " " + AtoZ + " " + Nums)
	case 2:
		a = ab(AtoM + " " + AtoS + " " + AtoZ)
		b = ab(atoz + " " + atoz + " " + atoz + " " + Nums)
	case 3:
		/*
			Target array's size: 16,994,304 bytes
			-
			Before optimizing makeMap():
			--------------------------------------------------------------
			uncompressed delta length: 1,855,440 bytes
			compressed delta length:     704,583 (4.15% of target's size)
			elapsed time:              171.4 seconds
			--------------------------------------------------------------
			171.25880: MakeDiff
			  0.16411: makeHash
			  3.78551: makeMap
			165.82172: longestMatch
			  0.09878: appendPart
			  0.13109: compressBytes
			-
			After optimizing makeMap():
			--------------------------------------------------------------
			uncompressed delta length: 1,952,772 bytes
			compressed delta length:     729,574 (4.29% of target's size)
			elapsed time:                2.4 seconds
			--------------------------------------------------------------
			  2.40135: MakeDiff
			  0.11608: makeHash
			  1.28985: makeMap
			  0.14999: longestMatch
			  0.07882: appendPart
			  0.09806: compressBytes
			-
			After adding backward-scanning in longestMatch()
			--------------------------------------------------------------
			uncompressed delta length: 1,675,811 bytes
			compressed delta length:     666,880 (3.92% of target's size)
			elapsed time:                    2.4 seconds
			--------------------------------------------------------------
			  2.45898: MakeDiff
			  0.15910: makeHash
			  1.49399: makeMap
			  0.16595: longestMatch
			  0.07311: appendPart
			  0.12408: compressBytes
		*/
		a = readData("test1.file")
		b = readData("test2.file")
		PL("loaded data")
	case 4:
		/*
			target size:        10,356,821
			uncompressed delta:  5,414,754
			compressed delta:    5,258,684 (50.7% of file size)
			elapsed time:              6.2 seconds
		*/
		a = readData("test1.zip")
		b = readData("test2.zip")
		PL("loaded data")
	case 5:
		/*
				target size:        17,096,704 bytes
				uncompressed delta:     64,081 bytes
				compressed delta:       25,967 (50.7% of file size)
				elapsed time:             2.06 seconds
				--------------------------------------------------------------
			  	  2.06019: MakeDiff
				  0.11507: makeHash
				  1.44146: makeMap
				  0.05109: longestMatch
				  0.00349: appendPart
				  0.00600: compressBytes
				  3.67731
		*/
		a = readData("day1.data")
		b = readData("day2.data")
		PL("loaded data")
	}
	if DebugTiming {
		tmr.Start("MakeDiff")
	}
	{
		var dif = MakeDiff(a, b)
		dif.Bytes()
	}
	if DebugTiming {
		tmr.Stop("MakeDiff")
		tmr.Print()
	}
} //                                                                    Test_02_

// go test --run Test_03_
func Test_03_(t *testing.T) {
	var a, b []byte
	switch 1 {
	case 1:
		a = ab(AtoM + " " + AtoS + " " + AtoZ)
		b = ab("000" + AtoZ + " " + AtoZ + " " + AtoZ + " " + Nums)
	}
	PL("start Test_03_")
	// -------------------------------------------------------------------------
	PL("\n" + Line)
	var d1 = MakeDiff(a, b)
	PL("CREATED d1:")
	d1.Dump()
	//
	var dbytes = d1.Bytes()
	PL("got 'dbytes'")
	// -------------------------------------------------------------------------
	PL("\n" + Line)
	if DebugTiming {
		tmr.Start("loadDiff")
	}
	var d2, err = loadDiff(dbytes)
	PL("CREATED d2: err:", err)
	d2.Dump()
	if DebugTiming {
		tmr.Stop("loadDiff")
		tmr.Print()
	}
} //                                                                    Test_03_

// go test --run Test_04_
func Test_04_(t *testing.T) {
	PL("Test_04_")
	var dif = Diff{
		sourceSize: 111,
		sourceHash: []byte("SOURCE"),
		targetSize: 222,
		targetHash: []byte("TARGET"),
		newCount:   333,
		oldCount:   444,
		parts: []diffPart{
			{},
			{},
		},
	}
	PL(dif.GoString())
} //                                                                    Test_04_

// -----------------------------------------------------------------------------
// # Test Helper Function

// readData reads 'filename' and returns its contents as an array of bytes
func readData(filename string) []byte {
	ret, err := ioutil.ReadFile(filename)
	if err != nil {
		PL("File reading error:", err)
		return nil
	}
	return ret
} //                                                                    readData

// ab converts s to a byte array
func ab(s string) []byte {
	return []byte(s)
} //                                                                          ab

//end
