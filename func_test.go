// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-15 16:24:49 47EABF                        go-delta/[func_test.go]
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
	"testing"

	"github.com/balacode/zr"
)

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
	PL("Test1 ################################################################")
	//
	var m1 = MakeMap(readData("test1.zip"))
	PL("Created m1. len(m1):", len(m1))
	//
	var m2 = MakeMap(readData("test2.zip"))
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
	var a = []byte("ABCDEFGHIJKLM" + " " +
		"ABCDEFGHIJKLMNOPQRSTUVWX" + " " +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	)
	var b = []byte("0x0x0xABCDEFGHIJKLMNOPQRSTUVWXYZ" + " " +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" + " " +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" + " " + "0123456789",
	)
	var tmr zr.Timer
	tmr.Start("MakeDiff()")
	{
		MakeDiff(a, b)
	}
	tmr.Stop("MakeDiff()")
	tmr.Print()
} //                                                              Test_MakeDiff_

//end
