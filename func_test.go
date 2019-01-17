// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-17 15:01:54 5B1BD9                        go-delta/[func_test.go]
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

// readData reads 'filename' and returns its contents as an array of bytes
func readData(filename string) []byte {
	ret, err := ioutil.ReadFile(filename)
	if err != nil {
		PL("File reading error:", err)
		return []byte{}
	}
	return ret
} //                                                                    readData

// go test --run Test1
func Test1(t *testing.T) {
	PL("Test1 " + Line)
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
} //                                                                       Test1

// go test --run Test_MakeDiff_
func Test_MakeDiff_(t *testing.T) {
	var a, b []byte
	switch 3 {
	case 1:
		a = []byte(AtoM + " " + AtoS + " " + AtoZ)
		b = []byte("0x0x0x" + AtoZ + " " + AtoZ + " " + AtoZ + " " + Nums)
	case 2:
		a = []byte(AtoM + " " + AtoS + " " + AtoZ)
		b = []byte(atoz + " " + atoz + " " + atoz + " " + Nums)
	case 3:
		/*
			Target array's size: 16,994,304 bytes

			Before optimizing makeMap():
			--------------------------------------------------------------
			unsipped delta length: 1,855,440 bytes
			zipped delta length:     704,583 (4.15% of target's size)
			elapsed time:              171.4 seconds
			--------------------------------------------------------------
			171.25880: MakeDiff
			  0.16411: makeHash
			  3.78551: makeMap
			165.82172: longestMatch
			  0.09878: appendPart
			  0.13109: compressBytes

			After optimizing makeMap():
			--------------------------------------------------------------
			unsipped delta length: 1,952,772 bytes
			zipped delta length:     729,574 (4.29% of target's size)
			elapsed time:                2.4 seconds
			--------------------------------------------------------------
			  2.40135: MakeDiff
			  0.11608: makeHash
			  1.28985: makeMap
			  0.14999: longestMatch
			  0.07882: appendPart
			  0.09806: compressBytes

			After adding backward-scanning in longestMatch()
			--------------------------------------------------------------
			unsipped delta length: 1,675,811 bytes
			zipped delta length:     666,880 (3.92% of target's size)
			elapsed time:                2.4 seconds
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
			target size:    10,356,821
			unzipped delta:  5,414,754
			zipped delta:    5,258,684 (50.7% of file size)
			elapsed time:          6.2 seconds
		*/
		a = readData("test1.zip")
		b = readData("test2.zip")
		PL("loaded data")
	}
	if DebugTiming {
		tmr.Start("MakeDiff")
	}
	{
		var diff = MakeDiff(a, b)
		diff.Bytes()
	}
	if DebugTiming {
		tmr.Stop("MakeDiff")
		tmr.Print()
	}
} //                                                              Test_MakeDiff_

//end
