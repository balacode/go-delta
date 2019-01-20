// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-20 14:19:30 87A8FE                        go-delta/[func_test.go]
// -----------------------------------------------------------------------------

package delta

/*
to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
)

const AtoM = "ABCDEFGHIJKLM"
const AtoS = "ABCDEFGHIJKLMNOPQRS"
const AtoZ = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const Nums = "0123456789"
const atoz = "abcdefghijklmnopqrstuvwxyz"

const PrintTestNames = true
const RunExperiments = true

var Line = strings.Repeat("#", 70)

// -----------------------------------------------------------------------------
// # Test Helper Functions

// ab converts s to a byte array.
func ab(s string) []byte {
	return []byte(s)
} //                                                                          ab

// printTestName prints the name of the calling unit test.
func printTestName() {
	if !PrintTestNames {
		return
	}
	var funcName = func() string {
		var programCounter, _, _, _ = runtime.Caller(2)
		var ret = runtime.FuncForPC(programCounter).Name()
		var i = strings.LastIndex(ret, ".")
		if i > -1 {
			ret = ret[i+1:]
		}
		ret += "()"
		return ret
	}
	fmt.Println("Running test:", funcName())
} //                                                               printTestName

// readData reads 'filename' and returns its contents as an array of bytes.
func readData(filename string) []byte {
	ret, err := ioutil.ReadFile(filename)
	if err != nil {
		PL("File reading error:", err)
		return nil
	}
	return ret
} //                                                                    readData

//end
