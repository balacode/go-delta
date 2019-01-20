// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-20 07:34:51 5F174D                            go-delta/[delta.go]
// -----------------------------------------------------------------------------

package delta

import (
	"bytes"
	"encoding/binary"
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
// # Internal Methods

// loadDelta __
func loadDelta(data []byte) (Delta, error) {
	//
	// uncompress the delta
	if DebugInfo {
		PL("loadDelta: compressed delta length:", len(data))
	}
	data = uncompressBytes(data)
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
