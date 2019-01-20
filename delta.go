// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-20 07:49:17 2CE007                            go-delta/[delta.go]
// -----------------------------------------------------------------------------

package delta

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
// # Read-Only Properties

// NewCount returns the number of chunks not matched in source array.
func (ob *Delta) NewCount() int {
	return ob.newCount
} //                                                                    NewCount

// OldCount returns the number of matched chunks in source array.
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

//end
