// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-28 21:39:43 F4FE39                       go-delta/[delta_load.go]
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
	buf := bytes.NewBuffer(data)
	readInt := func() int {
		var i int32
		err := binary.Read(buf, binary.BigEndian, &i)
		if err != nil {
			mod.Error("readInt() failed:", err)
			return -1
		}
		return int(i)
	}
	readBytes := func() []byte {
		var size int32
		err := binary.Read(buf, binary.BigEndian, &size)
		if err != nil {
			mod.Error("readBytes() failed @1:", err)
		}
		ar := make([]byte, size)
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
	ret := Delta{
		sourceSize: readInt(),
		sourceHash: readBytes(),
		targetSize: readInt(),
		targetHash: readBytes(),
		newCount:   readInt(),
		oldCount:   readInt(),
	}
	// read the parts
	count := readInt()
	if count < 1 {
		return Delta{},
			mod.Error("readBytes() failed @4: invalid number of parts:", count)
	}
	ret.parts = make([]deltaPart, count)
	for i := range ret.parts {
		pt := &ret.parts[i]
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
