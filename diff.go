// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-15 20:27:06 7E19B3                             go-delta/[diff.go]
// -----------------------------------------------------------------------------

package bdelta

// Diff stores the binary delta difference between two byte arrays
type Diff struct {
	sourceHash []byte
	// hash of the source byte array
	//
	targetHash []byte
	// expected hash of result after this Diff is applied to source
	//
	parts []diffPart
	// array of referring to chunks in source array, or new bytes to append
	//
	newCount int
	// number of chunks that could not be matched in source
	//
	oldCount int
	// number of chunks that were matched in source
	//
} //                                                                        Diff

// diffPart stores references to chunks in the source array,
// or specifies bytes to append to result array directly
type diffPart struct {
	sourceLoc int
	// byte position of the chunk in source array,
	// or -1 when the bytes should be picked from 'data'
	//
	size int
	// size of the chunk in bytes
	//
	data []byte
	// optional bytes (only when sourceLoc is -1)
} //                                                                    diffPart

// -----------------------------------------------------------------------------
// # Public Properties

// NewCount __
func (ob *Diff) NewCount() int {
	return ob.newCount
} //                                                                    NewCount

// OldCount __
func (ob *Diff) OldCount() int {
	return ob.oldCount
} //                                                                    OldCount

// -----------------------------------------------------------------------------
// # Internal Methods

// writePart appends binary difference data
func (ob *Diff) writePart(sourceLoc, size int, data []byte) {
	PL("WP",
		"sourceLoc:", sourceLoc,
		"size:", size,
		"data:", data, string(data))
} //                                                                   writePart

//end
