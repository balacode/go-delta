// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-15 19:48:54 C85259                             go-delta/[diff.go]
// -----------------------------------------------------------------------------

package bdelta

// Diff __
type Diff struct {
	mode       byte
	sourceHash []byte
	targetHash []byte
	parts      []diffPart
} //                                                                        Diff

// diffPart __
type diffPart struct {
	sourceLoc int
	size      int
	data      []byte
} //                                                                    diffPart

//end
