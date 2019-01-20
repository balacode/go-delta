// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-20 07:32:20 579F66                       go-delta/[delta_dump.go]
// -----------------------------------------------------------------------------

package delta

import (
	"fmt"
)

// Dump prints this object to the console in a human-friendly format.
func (ob *Delta) Dump() {
	var pl = fmt.Println
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

//end
