// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-20 07:25:22 4A054E                           go-delta/[module.go]
// -----------------------------------------------------------------------------

package delta

import (
	"fmt"
	"github.com/balacode/zr"
)

// -----------------------------------------------------------------------------
// # Module Constants / Variables

// MatchLimit specifies the maximum number of positions tracked
// for each unique key in the map of source data. See makeMap().
const MatchLimit = 50

// MatchSize specifies the size of unique chunks being searched for, in bytes.
const MatchSize = 8

// DebugInfo when set, causes printing of messages helpful for debugging.
var DebugInfo = false

// DebugTiming controls timing (benchmarking) of time spent in each function.
var DebugTiming = true

// DebugWriteArgs when set, prints the arguments passed to write()
var DebugWriteArgs = false

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
