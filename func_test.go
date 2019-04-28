// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-28 21:39:43 19D4DF                        go-delta/[func_test.go]
// -----------------------------------------------------------------------------

package delta

/*
to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
	"testing"
)

const AtoM = "ABCDEFGHIJKLM"
const AtoS = "ABCDEFGHIJKLMNOPQRS"
const AtoZ = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const Nums = "0123456789"
const atoz = "abcdefghijklmnopqrstuvwxyz"

const PrintTestNames = true

var Line = strings.Repeat("#", 70)

// -----------------------------------------------------------------------------
// # Function Unit Tests

// go test --run Test_readHash_
func Test_readHash_(t *testing.T) {
	if PrintTestNames {
		printTestName()
	}
	var test = func(input []byte) {
		var resultHash []byte
		{
			buf := bytes.NewBuffer(input)
			resultHash = readHash(buf)
		}
		var expectHash []byte
		{
			buf := bytes.NewBuffer(input)
			expectHash = makeHash(buf.Bytes())
		}
		if !bytes.Equal(resultHash, expectHash) {
			t.Errorf("\n input:\n\t%v\n%s\n expect:%v\n\t result:\n\t%v\n",
				input, string(input), expectHash, resultHash)
		}
	}
	TempBufferSize = 100
	test(nil)
	test([]byte("abc"))
	test([]byte(strings.Repeat("abc", 1024)))
} //                                                              Test_readHash_

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
	funcName := func() string {
		programCounter, _, _, _ := runtime.Caller(2)
		ret := runtime.FuncForPC(programCounter).Name()
		i := strings.LastIndex(ret, ".")
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
