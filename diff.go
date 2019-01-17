// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-17 14:54:17 520860                             go-delta/[diff.go]
// -----------------------------------------------------------------------------

package bdelta

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Diff stores the binary delta difference between two byte arrays
type Diff struct {
	sourceSize int        // size of the source array
	sourceHash []byte     // hash of the source byte array
	targetSize int        // size of the target array
	targetHash []byte     // hash of the result after this Diff is applied
	newCount   int        // number of chunks not matched in source array
	oldCount   int        // number of matched chunks in source array
	parts      []diffPart // array referring to chunks in source array,
	//                       or new bytes to append
} //                                                                        Diff

// diffPart stores references to chunks in the source array,
// or specifies bytes to append to result array directly
type diffPart struct {
	sourceLoc int // byte position of the chunk in source array,
	//               or -1 when 'data' supplies the bytes directly
	//
	size int    // size of the chunk in bytes
	data []byte // optional bytes (only when sourceLoc is -1)
} //                                                                    diffPart

// -----------------------------------------------------------------------------
// # Public Properties

// Bytes converts the Delta structure to a byte array
// (for serializing to a file, etc.)
func (ob *Diff) Bytes() []byte {
	var buf = bytes.NewBuffer(make([]byte, 0, 1024))
	var writeInt = func(i int) {
		err := binary.Write(buf, binary.BigEndian, int32(i))
		if err != nil {
			mod.Error("writeInt(", i, ") failed:", err)
		}
	}
	var writeBytes = func(data []byte) {
		err := binary.Write(buf, binary.BigEndian, int32(len(data)))
		if err != nil {
			mod.Error("writeBytes(", data, ") failed:", err)
		}
		buf.Write(data)
	}
	// write the header
	writeInt(ob.sourceSize)
	writeBytes(ob.sourceHash)
	writeInt(ob.targetSize)
	writeBytes(ob.targetHash)
	writeInt(ob.newCount)
	writeInt(ob.oldCount)
	writeInt(len(ob.parts))
	//
	// write the parts
	for _, part := range ob.parts {
		writeInt(part.sourceLoc)
		if part.sourceLoc == -1 {
			writeInt(len(part.data))
			buf.Write(part.data)
			continue
		}
		writeInt(part.size)
	}
	// compress the delta
	if DebugInfo {
		PL("unsipped delta length:", len(buf.Bytes()))
	}
	var ret = compressBytes(buf.Bytes())
	if DebugInfo {
		PL("zipped delta length:", len(ret))
	}
	return ret
} //                                                                       Bytes

// NewCount __
func (ob *Diff) NewCount() int {
	return ob.newCount
} //                                                                    NewCount

// OldCount __
func (ob *Diff) OldCount() int {
	return ob.oldCount
} //                                                                    OldCount

// SourceSize returns the size of the source byte array, in bytes.
func (ob *Diff) SourceSize() int {
	return ob.sourceSize
} //                                                                  SourceSize

// TargetSize returns the size of the target byte array, in bytes.
func (ob *Diff) TargetSize() int {
	return ob.targetSize
} //                                                                  TargetSize

// -----------------------------------------------------------------------------
// # Public Method

// Dump prints this object to the console in a human-friendly format.
func (ob *Diff) Dump() {
	var pl = fmt.Println
	pl()
	pl("sourceHash:", ob.sourceHash)
	pl("targetHash:", ob.targetHash)
	pl("newCount:", ob.newCount)
	pl("oldCount:", ob.oldCount)
	pl("len(parts):", len(ob.parts))
	pl()
	for i, part := range ob.parts {
		pl("part:", i, "sourceLoc:", part.sourceLoc,
			"size:", part.size,
			"data:", part.data, string(part.data))
	}
} //                                                                        Dump

// -----------------------------------------------------------------------------
// # Internal Methods

// appendPart appends binary difference data
func (ob *Diff) appendPart(sourceLoc, size int, data []byte) {
	if DebugTiming {
		tmr.Start("appendPart")
		defer tmr.Stop("appendPart")
	}
	if DebugInfo && DebugAppendPartArgs {
		PL("appendPart",
			"sourceLoc:", sourceLoc,
			"size:", size,
			"data:", data, string(data))
	}
	// argument validations
	switch {
	case sourceLoc < -1:
		mod.Error("sourceLoc:", sourceLoc, " < -1")
		return
	case sourceLoc == -1 && len(data) == 0:
		mod.Error("sourceLoc == -1 && len(data) == 0")
		return
	case sourceLoc != -1 && len(data) != 0:
		mod.Error("sourceLoc != -1 && len(data):", len(data), "!= 0")
		return
	case size < 1:
		mod.Error("size:", size, " < 1")
		return
	}
	// if the previous part was embedded directly, append to that part's data
	if sourceLoc == -1 {
		var n = len(ob.parts)
		if n > 0 {
			var last = &ob.parts[n-1]
			if last.sourceLoc == -1 {
				last.size += len(data)
				last.data = append(last.data, data...)
				return
			}
		}
	}
	// append a new part
	var ar []byte
	if sourceLoc == -1 {
		ar = make([]byte, len(data))
		copy(ar, data)
	}
	ob.parts = append(ob.parts,
		diffPart{sourceLoc: sourceLoc, size: size, data: ar})
} //                                                                  appendPart

//end
