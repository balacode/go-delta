// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-20 07:11:02 A08799                             go-delta/[func.go]
// -----------------------------------------------------------------------------

package delta

import (
	"bytes"
	"compress/zlib"
	"crypto/sha512"
	"io"
)

// -----------------------------------------------------------------------------
// # Helper Functions: Compression

// compressBytes compresses an array of bytes and
// returns the ZLIB-compressed array of bytes.
func compressBytes(data []byte) []byte {
	if DebugTiming {
		tmr.Start("compressBytes")
		defer tmr.Stop("compressBytes")
	}
	if len(data) == 0 {
		return nil
	}
	// zip data in standard manner
	var b bytes.Buffer
	var w = zlib.NewWriter(&b)
	var _, err = w.Write(data)
	w.Close()
	//
	// log any problem
	const ERRM = "Failed compressing data with zlib:"
	if err != nil {
		mod.Error(ERRM, err)
		return nil
	}
	var ret = b.Bytes()
	if len(ret) < 3 {
		mod.Error(ERRM, "length < 3")
		return nil
	}
	return ret
} //                                                               compressBytes

// uncompressBytes uncompresses a ZLIB-compressed array of bytes.
func uncompressBytes(data []byte) []byte {
	var readCloser, err = zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		mod.Error("uncompressBytes:", err)
		return nil
	}
	var ret = bytes.NewBuffer(make([]byte, 0, 8192))
	io.Copy(ret, readCloser)
	readCloser.Close()
	return ret.Bytes()
} //                                                             uncompressBytes

// -----------------------------------------------------------------------------
// # Helper Functions

// makeHash returns the SHA-512 hash of byte slice 'data'.
func makeHash(data []byte) []byte {
	if DebugTiming {
		tmr.Start("makeHash")
		defer tmr.Stop("makeHash")
	}
	if len(data) == 0 {
		return nil
	}
	var ret = sha512.Sum512(data)
	return ret[:]
} //                                                                    makeHash

//end
