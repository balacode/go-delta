// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-05-25 22:34:13 FA0B42                           go-delta/[module.go]
// -----------------------------------------------------------------------------

package delta

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/balacode/zr"
)

// -----------------------------------------------------------------------------
// # Module Constants / Variables

const (
	// MatchLimit specifies the maximum number of positions tracked
	// for each unique key in the map of source data. See makeMap().
	MatchLimit = 50

	// MatchSize specifies the size of unique
	// chunks being searched for, in bytes.
	MatchSize = 9
)

var (
	// PL is fmt.Println() but is used only for debugging.
	PL = fmt.Println

	// TempBufferSize sets the size of memory buffers for reading files and other
	// streams. This memory is not fixed but allocated/released transiently.
	TempBufferSize = 32 * 1024 * 1024 // 32 MB

	// tmr is used for timing all methods/functions during tuning.
	tmr zr.Timer
)

// -----------------------------------------------------------------------------
// # Debugging Flags

var (
	// DebugInfo when set, causes printing of messages helpful for debugging.
	DebugInfo = false

	// DebugTiming controls timing (benchmarking) of time spent in each function.
	DebugTiming = true

	// DebugWriteArgs when set, prints the arguments passed to write()
	DebugWriteArgs = false
)

// -----------------------------------------------------------------------------
// # Error Handler

// SetErrorFunc changes the error-handling function, so that
// all errors in this package will be sent to this handler,
// which is useful for custom logging and mocking during unit tests.
// To restore the default error handler use SetErrorFunc(nil).
func SetErrorFunc(fn func(args ...interface{}) error) {
	if fn == nil {
		mod.Error = defaultErrorFunc
		return
	}
	mod.Error = fn
} //                                                                SetErrorFunc

// defaultErrorFunc is the default error
// handling function assigned to mod.Error
func defaultErrorFunc(args ...interface{}) error {
	//
	// write all args to a message string (add spaces between args)
	var buf bytes.Buffer
	for i, arg := range args {
		if i > 0 {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprint(arg))
	}
	msg := buf.String()
	//
	// if DebugInfo is on, print the message to the console
	if DebugInfo {
		fmt.Println("ERROR:\n", msg)
	}
	// return error based on message
	return errors.New(msg)
} //                                                            defaultErrorFunc

// -----------------------------------------------------------------------------
// # Module Global

// mod variable though wich mockable functions are called
var mod = thisMod{Error: defaultErrorFunc}

// thisMod specifies mockable functions
type thisMod struct {
	Error func(args ...interface{}) error
}

// ModReset restores all mocked functions to the original standard functions.
func (ob *thisMod) Reset() { ob.Error = defaultErrorFunc }

// end
