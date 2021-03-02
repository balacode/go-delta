// -----------------------------------------------------------------------------
// github.com/balacode/go-delta                    go-delta/[delta_go_string.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package delta

import (
	"bytes"
	"fmt"
)

// GoString returns a Go-syntax representation of the Delta structure.
// It implements the fmt.GoStringer interface.
func (ob Delta) GoString() string {
	var buf bytes.Buffer
	write := func(args ...string) {
		for _, s := range args {
			buf.WriteString(s)
		}
	}
	str := func(v interface{}) string {
		return fmt.Sprintf("%#v", v)
	}
	write("Delta{", "\n",
		"\t", "sourceSize: ", str(ob.sourceSize), ",\n",
		"\t", "sourceHash: ", str(ob.sourceHash), ",\n",
		"\t", "targetSize: ", str(ob.targetSize), ",\n",
		"\t", "targetHash: ", str(ob.targetHash), ",\n",
		"\t", "newCount:   ", str(ob.newCount), ",\n",
		"\t", "oldCount:   ", str(ob.oldCount), ",\n",
		"\t", "parts: []deltaPart{\n",
	)
	for _, pt := range ob.parts {
		write("\t\t{",
			"sourceLoc: ", str(pt.sourceLoc), ", ",
			"size: ", str(pt.size), ", ",
			"data: ", str(pt.data), "}\n")
	}
	write("\t},\n}")
	return buf.String()
} //                                                                    GoString

// end
