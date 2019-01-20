// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-20 07:15:01 789CF3                            go-delta/[delta.go]
// -----------------------------------------------------------------------------

package delta

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Delta stores the binary delta difference between two byte arrays
type Delta struct {
	sourceSize int         // size of the source array
	sourceHash []byte      // hash of the source byte array
	targetSize int         // size of the target array
	targetHash []byte      // hash of the result after this Delta is applied
	newCount   int         // number of chunks not matched in source array
	oldCount   int         // number of matched chunks in source array
	parts      []deltaPart // array referring to chunks in source array,
	//                        or new bytes to append
} //                                                                       Delta

// deltaPart stores references to chunks in the source array,
// or specifies bytes to append to result array directly
type deltaPart struct {
	sourceLoc int // byte position of the chunk in source array,
	//               or -1 when 'data' supplies the bytes directly
	//
	size int    // size of the chunk in bytes
	data []byte // optional bytes (only when sourceLoc is -1)
} //                                                                   deltaPart

// -----------------------------------------------------------------------------
// # Public Properties

// GoString returns a Go-syntax representation of the Delta structure.
// It implements the GoStringer interface.
func (ob Delta) GoString() string {
	var buf bytes.Buffer
	var write = func(args ...string) {
		for _, s := range args {
			buf.WriteString(s)
		}
	}
	var str = func(v interface{}) string {
		return fmt.Sprintf("%#v", v)
	}
	write("Delta{", "\n",
		"\t", "sourceSize: ", str(ob.sourceSize), ",\n",
		"\t", "sourceHash: ", str(ob.sourceHash), ",\n",
		"\t", "targetSize: ", str(ob.targetSize), ",\n",
		"\t", "targetHash: ", str(ob.targetHash), ",\n",
		"\t", "newCount:   ", str(ob.newCount), ",\n",
		"\t", "oldCount:   ", str(ob.oldCount), ",\n",
		"\t", "parts: []deltaPart{\n",
	)
	for _, pt := range ob.parts {
		write("\t\t{",
			"sourceLoc: ", str(pt.sourceLoc), ", ",
			"size: ", str(pt.size), ", ",
			"data: ", str(pt.data), "}\n")
	}
	write("\t},\n}")
	return buf.String()
} //                                                                    GoString

// NewCount __
func (ob *Delta) NewCount() int {
	return ob.newCount
} //                                                                    NewCount

// OldCount __
func (ob *Delta) OldCount() int {
	return ob.oldCount
} //                                                                    OldCount

// SourceSize returns the size of the source byte array, in bytes.
func (ob *Delta) SourceSize() int {
	return ob.sourceSize
} //                                                                  SourceSize

// TargetSize returns the size of the target byte array, in bytes.
func (ob *Delta) TargetSize() int {
	return ob.targetSize
} //                                                                  TargetSize

// -----------------------------------------------------------------------------
// # Public Method

// Dump prints this object to the console in a human-friendly format.
func (ob *Delta) Dump() {
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
func (ob *Delta) appendPart(sourceLoc, size int, data []byte) {
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
		deltaPart{sourceLoc: sourceLoc, size: size, data: ar})
} //                                                                  appendPart

// loadDelta __
func loadDelta(delta []byte) (Delta, error) {
	//
	// uncompress the delta
	if DebugInfo {
		PL("loadDelta: compressed delta length:", len(delta))
	}
	var data = uncompressBytes(delta)
	if DebugInfo {
		PL("loadDelta: uncompressed delta length:", len(data))
	}
	var buf = bytes.NewBuffer(data)
	var readInt = func() int {
		var i int32
		err := binary.Read(buf, binary.BigEndian, &i)
		if err != nil {
			mod.Error("readInt() failed:", err)
			return -1
		}
		return int(i)
	}
	var readBytes = func() []byte {
		var size int32
		err := binary.Read(buf, binary.BigEndian, &size)
		if err != nil {
			mod.Error("readBytes() failed @1:", err)
		}
		var ar = make([]byte, size)
		var nread int
		nread, err = buf.Read(ar)
		if err != nil {
			mod.Error("readBytes() failed @2:", err)
		}
		if nread != int(size) {
			mod.Error("readBytes() failed @3: size:", size, "nread:", nread)
		}
		return ar
	}
	// read the header
	var ret = Delta{
		sourceSize: readInt(),
		sourceHash: readBytes(),
		targetSize: readInt(),
		targetHash: readBytes(),
		newCount:   readInt(),
		oldCount:   readInt(),
	}
	// read the parts
	var count = readInt()
	if count < 1 {
		return Delta{},
			mod.Error("readBytes() failed @4: invalid number of parts:", count)
	}
	ret.parts = make([]deltaPart, count)
	for i := range ret.parts {
		var pt = &ret.parts[i]
		pt.sourceLoc = readInt()
		if pt.sourceLoc == -1 {
			pt.data = readBytes()
			pt.size = len(pt.data)
			continue
		}
		pt.size = readInt()
	}
	return ret, nil
} //                                                                   loadDelta

//end
