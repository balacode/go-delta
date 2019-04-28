// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-28 21:39:43 AB8AB9                             go-delta/[func.go]
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
	w := zlib.NewWriter(&b)
	_, err := w.Write(data)
	w.Close()
	//
	// log any problem
	const ERRM = "Failed compressing data with zlib:"
	if err != nil {
		mod.Error(ERRM, err)
		return nil
	}
	ret := b.Bytes()
	if len(ret) < 3 {
		mod.Error(ERRM, "length < 3")
		return nil
	}
	return ret
} //                                                               compressBytes

// uncompressBytes uncompresses a ZLIB-compressed array of bytes.
func uncompressBytes(data []byte) []byte {
	readCloser, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		mod.Error("uncompressBytes:", err)
		return nil
	}
	ret := bytes.NewBuffer(make([]byte, 0, 8192))
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
	ret := sha512.Sum512(data)
	return ret[:]
} //                                                                    makeHash

// readHash returns the SHA-512 hash of the bytes from 'stream'.
func readHash(stream io.Reader) []byte {
	if DebugTiming {
		tmr.Start("readHash")
		defer tmr.Stop("readHash")
	}
	hasher := sha512.New()
	buf := make([]byte, TempBufferSize)
	for first := true; ; first = false {
		n, err := stream.Read(buf)
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
	ret := hasher.Sum(nil)
	return ret
} //                                                                    readHash

// readLen returns the total size of 'stream' in bytes.
// After a call to readLen, the current reading
// position returns to the start or the stream.
func readLen(stream io.ReadSeeker) int {
	ret, _ := stream.Seek(0, io.SeekEnd)
	stream.Seek(0, io.SeekStart)
	return int(ret)
} //                                                                     readLen

// readStream __
func readStream(from io.ReadSeeker, to []byte) (n int64, err error) {
	// read from the stream
	{
		var num int
		num, err = from.Read(to)
		n = int64(num)
	}
	if err == io.EOF {
		if n != 0 {
			mod.Error("Expected zero: n =", n)
		}
		return -1, nil
	}
	if err != nil {
		return -1, mod.Error("Failed reading:", err)
	}
	return n, err
} //                                                                  readStream

//end
