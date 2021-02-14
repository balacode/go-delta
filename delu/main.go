// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-28 21:39:43 3FBED6                      go-delta/deltau/[main.go]
// -----------------------------------------------------------------------------

package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/balacode/go-delta"
)

const Usage = `delu - delta update utility
Usage:

To create a delta update file:
    delu make <source file> <target file> <delta-file>

To apply a delta update:
    delu apply <source file> <delta file> <target-file>
`

var printError = fmt.Println

func main() {
	a := os.Args[1:]
	n := len(a)
	switch {
	case n == 0:
		fmt.Println(Usage)
	case a[0] == "-help" || a[0] == "--help" || a[0] == "/?":
		fmt.Println(Usage)
	case n != 4:
		printError("You specified the wrong number of parameters!")
		fmt.Println(Usage)
	case a[0] == "apply":
		applyDelta(a[1], a[2], a[3]) // source, delta, target
	case a[0] == "make":
		makeDelta(a[1], a[2], a[3]) // source, target, delta
	}
} //                                                                        main

// -----------------------------------------------------------------------------
// # Helper Functions

// applyDelta creates 'targetFile' by applying 'deltaFile' to 'sourceFile'.
func applyDelta(sourceFile, deltaFile, targetFile string) {
	//
	// make sure the target file does not exist
	if fileExists(targetFile) {
		printError("Target exists already:", targetFile)
		return
	}
	var err error
	//
	// read the source file into a byte array
	var sourceAr []byte
	sourceAr, err = ioutil.ReadFile(sourceFile)
	if err != nil {
		printError("Failed reading", sourceFile, ":\n", err)
		return
	}
	// read the delta file into a byte array
	var deltaAr []byte
	deltaAr, err = ioutil.ReadFile(deltaFile)
	if err != nil {
		printError("Failed reading", deltaFile, ":\n", err)
		return
	}
	// create a Delta from the delta bytes
	var d delta.Delta
	d, err = delta.Load(deltaAr)
	if err != nil {
		printError("Failed to apply delta to source:\n", err)
	}
	// create target data from source and delta
	var targetAr []byte
	targetAr, err = d.Apply(sourceAr)
	if err != nil {
		printError("Failed to apply delta to source:\n", err)
	}
	// save the target
	err = ioutil.WriteFile(targetFile, targetAr, 0644)
	if err != nil {
		printError("Failed saving", targetFile, ":\n", err)
	}
} //                                                                  applyDelta

// fileExists returns true if the file given by 'path' exists.
func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	printError("Error while checking if", path, "exists:\n", err)
	return false
} //                                                                  fileExists

// makeDelta creates 'deltaFile', using 'sourceFile' and 'targetFile'.
// The delta file only stores the differences between source and target.
func makeDelta(sourceFile, targetFile, deltaFile string) {
	//
	// make sure the delta file does not exist
	if fileExists(deltaFile) {
		printError("Delta file exists already:", deltaFile)
		return
	}
	var err error
	//
	// read the source file into a byte array
	var sourceAr []byte
	sourceAr, err = ioutil.ReadFile(sourceFile)
	if err != nil {
		printError("Failed reading", sourceFile, ":\n", err)
		return
	}
	// read the target file into a byte array
	var targetAr []byte
	targetAr, err = ioutil.ReadFile(targetFile)
	if err != nil {
		printError("Failed reading", targetFile, ":\n", err)
		return
	}
	// create a Delta from the difference between source and target
	d := delta.Make(sourceAr, targetAr)
	deltaAr := d.Bytes()
	//
	// save the delta
	err = ioutil.WriteFile(deltaFile, deltaAr, 0644)
	if err != nil {
		printError("Failed saving", deltaFile, ":\n", err)
	}
} //                                                                   makeDelta

// end
