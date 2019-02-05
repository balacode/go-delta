// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-02-05 16:57:16 C9F2F9                             go-delta/[func.go]
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

// hashOfBytes returns the SHA-512 hash of byte slice 'ar'.
func hashOfBytes(ar []byte) []byte {
	if DebugTiming {
		tmr.Start("hashOfBytes")
		defer tmr.Stop("hashOfBytes")
	}
	if len(ar) == 0 {
		return nil
	}
	var ret = sha512.Sum512(ar)
	return ret[:]
} //                                                                 hashOfBytes

// hashOfReader returns the SHA-512 hash of the bytes from 'stream'.
func hashOfReader(stream io.Reader) []byte {
	if DebugTiming {
		tmr.Start("hashOfReader")
		defer tmr.Stop("hashOfReader")
	}
	var hasher = sha512.New()
	var buf = make([]byte, TempBufferSize)
	for first := true; ; first = false {
		var n, err = stream.Read(buf)
		if err == io.EOF && first {
			return nil
		}
		if err == io.EOF {
			if n != 0 {
				mod.Error("Expected zero: n =", n)
			}
			break
		}
		if err != nil {
			mod.Error("Failed reading:", err)
			return nil
		}
		if n == 0 {
			break
		}
		n, err = hasher.Write(buf[:n])
		if err != nil {
			mod.Error("Failed writing:", err)
			return nil
		}
	}
	var ret = hasher.Sum(nil)
	return ret
} //                                                                hashOfReader

//end
