// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-13 18:14:55 74EE90                        go-delta/[func_test.go]
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
	"fmt"
	"io/ioutil"
	"testing"
)

var PL = fmt.Println

// go test --run Test1
func Test1(t *testing.T) {
	PL("Test1 ################################################################")
	var m1 map[[8]byte][]int
	{
		// data, err := ioutil.ReadFile("lorem_ipsum.txt")
		data, err := ioutil.ReadFile("test1.zip")
		if err != nil {
			PL("ERROR:", err)
			return
		}
		m1 = MakeMap(data)
	}
	PL("Created m1 ###########################################################")
	PL("len(m1):", len(m1))
	//
	var m2 map[[8]byte][]int
	{
		data, err := ioutil.ReadFile("test2.zip")
		if err != nil {
			PL("ERROR:", err)
			return
		}
		m2 = MakeMap(data)
	}
	PL("Created m2 ###########################################################")
	PL("len(m2):", len(m2))
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
}

//end
