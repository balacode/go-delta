// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-20 12:49:30 02CB94                        go-delta/[func_test.go]
// -----------------------------------------------------------------------------

package delta

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

// go test --run Test_Delta_Apply_
func Test_Delta_Apply_(t *testing.T) {
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
			171.25880: delta.Make
			  0.16411: makeHash
			  3.78551: makeMap
			165.82172: longestMatch
			  0.09878: write
			  0.13109: compressBytes
			-
			After optimizing makeMap():
			--------------------------------------------------------------
			uncompressed delta length: 1,952,772 bytes
			compressed delta length:     729,574 (4.29% of target's size)
			elapsed time:                2.4 seconds
			--------------------------------------------------------------
			  2.40135: delta.Make
			  0.11608: makeHash
			  1.28985: makeMap
			  0.14999: longestMatch
			  0.07882: write
			  0.09806: compressBytes
			-
			After adding backward-scanning in longestMatch()
			--------------------------------------------------------------
			uncompressed delta length: 1,675,811 bytes
			compressed delta length:     666,880 (3.92% of target's size)
			elapsed time:                    2.4 seconds
			--------------------------------------------------------------
			  2.45898: delta.Make
			  0.15910: makeHash
			  1.49399: makeMap
			  0.16595: longestMatch
			  0.07311: write
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
			  	  2.06019: delta.Make
				  0.11507: makeHash
				  1.44146: makeMap
				  0.05109: longestMatch
				  0.00349: write
				  0.00600: compressBytes
				  3.67731
		*/
		a = readData("day1.data")
		b = readData("day2.data")
		PL("loaded data")
	}
	if DebugTiming {
		tmr.Start("delta.Make")
	}
	{
		var d = Make(a, b)
		d.Bytes()
	}
	if DebugTiming {
		tmr.Stop("delta.Make")
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
	var d1 = Make(a, b)
	PL("CREATED d1:")
	d1.Dump()
	//
	var dbytes = d1.Bytes()
	PL("got 'dbytes'")
	// -------------------------------------------------------------------------
	PL("\n" + Line)
	if DebugTiming {
		tmr.Start("loadDelta")
	}
	var d2, err = loadDelta(dbytes)
	PL("CREATED d2: err:", err)
	d2.Dump()
	if DebugTiming {
		tmr.Stop("loadDelta")
		tmr.Print()
	}
} //                                                                    Test_03_

// go test --run Test_04_
func Test_04_(t *testing.T) {
	PL("Test_04_")
	var d = Delta{
		sourceSize: 111,
		sourceHash: []byte("SOURCE"),
		targetSize: 222,
		targetHash: []byte("TARGET"),
		newCount:   333,
		oldCount:   444,
		parts: []deltaPart{
			{},
			{},
		},
	}
	PL(d.GoString())
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
