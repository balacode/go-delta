// -----------------------------------------------------------------------------
// github.com/balacode/go-delta                         go-delta/[delta_dump.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package delta

import (
	"fmt"
)

// Dump prints this object to the console in a human-friendly format.
func (ob *Delta) Dump() {
	pl := fmt.Println
	pl()
	pl("sourceHash:", ob.sourceHash)
	pl("targetHash:", ob.targetHash)
	pl("newCount:", ob.newCount)
	pl("oldCount:", ob.oldCount)
	pl("len(parts):", len(ob.parts))
	pl()
	for i, part := range ob.parts {
		pl("part:", i, "sourceLoc:", part.sourceLoc,
			"size:", part.size,
			"data:", part.data, string(part.data))
	}
} //                                                                        Dump

// end
