// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-20 23:32:11 F73C71                       go-delta/[delta_load.go]
// -----------------------------------------------------------------------------

package delta

import (
	"bytes"
	"encoding/binary"
)

// Load fills a new Delta structure from a byte
// array previously returned by Delta.Bytes().
func Load(data []byte) (Delta, error) {
	//
	// uncompress the delta
	if DebugInfo {
		PL("Load: compressed delta length:", len(data))
	}
	data = uncompressBytes(data)
	if DebugInfo {
		PL("Load: uncompressed delta length:", len(data))
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
} //                                                                        Load

//end
