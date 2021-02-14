// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-28 21:39:43 739455                      go-delta/[delta_bytes.go]
// -----------------------------------------------------------------------------

package delta

import (
	"bytes"
	"encoding/binary"
)

// Bytes converts the Delta structure to a byte array
// (for serializing to a file, etc.)
func (ob *Delta) Bytes() []byte {
	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	//
	writeInt := func(i int) error {
		err := binary.Write(buf, binary.BigEndian, int32(i))
		if err != nil {
			return mod.Error("writeInt(", i, ") failed:", err)
		}
		return nil
	}
	writeBytes := func(data []byte) error {
		err := writeInt(len(data))
		if err != nil {
			return mod.Error("writeBytes([", len(data), "]) failed @1:", err)
		}
		var n int
		n, err = buf.Write(data)
		if err != nil {
			return mod.Error("writeBytes([", len(data), "]) failed @2:", err)
		}
		if n != len(data) {
			return mod.Error("writeBytes([", len(data), "]) failed @3:",
				"wrote wrong number of bytes:", n)
		}
		return nil
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
			writeBytes(part.data)
			continue
		}
		writeInt(part.size)
	}
	// compress the delta
	if DebugInfo {
		PL("uncompressed delta length:", len(buf.Bytes()))
	}
	ret := compressBytes(buf.Bytes())
	if DebugInfo {
		PL("compressed delta length:", len(ret))
	}
	return ret
} //                                                                       Bytes

// end
