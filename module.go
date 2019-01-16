// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-16 12:37:51 7262F5                           go-delta/[module.go]
// -----------------------------------------------------------------------------

package bdelta

import (
	"fmt"
	"github.com/balacode/zr"
)

// -----------------------------------------------------------------------------
// # Module Constants / Variables

const ChunkSize = 8

// PL is fmt.Println() but is used only for debugging.
var PL = fmt.Println

// -----------------------------------------------------------------------------
// # Function Proxy Variables (for mocking)

type thisMod struct {
	Error func(args ...interface{}) error
}

var mod = thisMod{
	Error: zr.Error,
}

// ModReset restores all mocked functions to the original standard functions.
func (ob *thisMod) Reset() {
	ob.Error = zr.Error
} //                                                                       Reset

//end
