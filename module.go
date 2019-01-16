// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-16 15:18:00 845B74                           go-delta/[module.go]
// -----------------------------------------------------------------------------

package bdelta

import (
	"fmt"
	"github.com/balacode/zr"
)

// -----------------------------------------------------------------------------
// # Module Constants / Variables

const ChunkSize = 8

const DebugInfo = true
const DebugTiming = true

// PL is fmt.Println() but is used only for debugging.
var PL = fmt.Println

// tmr is used for timing all methods/functions during tuning.
var tmr zr.Timer

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
