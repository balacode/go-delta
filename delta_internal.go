// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-28 21:39:43 CEDAE0                   go-delta/[delta_internal.go]
// -----------------------------------------------------------------------------

package delta

// write appends binary difference data
func (ob *Delta) write(sourceLoc, size int, data []byte) {
	if DebugTiming {
		tmr.Start("write")
		defer tmr.Stop("write")
	}
	if DebugInfo && DebugWriteArgs {
		PL("write",
			"sourceLoc:", sourceLoc,
			"size:", size,
			"data:", data, string(data))
	}
	// argument validations
	switch {
	case sourceLoc < -1:
		mod.Error("sourceLoc:", sourceLoc, " < -1")
		return
	case sourceLoc == -1 && len(data) == 0:
		mod.Error("sourceLoc == -1 && len(data) == 0")
		return
	case sourceLoc != -1 && len(data) != 0:
		mod.Error("sourceLoc != -1 && len(data):", len(data), "!= 0")
		return
	case size < 1:
		mod.Error("size:", size, " < 1")
		return
	}
	// if the previous part was embedded directly, append to that part's data
	if sourceLoc == -1 {
		n := len(ob.parts)
		if n > 0 {
			last := &ob.parts[n-1]
			if last.sourceLoc == -1 {
				last.size += len(data)
				last.data = append(last.data, data...)
				return
			}
		}
	}
	// append a new part
	var ar []byte
	if sourceLoc == -1 {
		ar = make([]byte, len(data))
		copy(ar, data)
	}
	ob.parts = append(ob.parts,
		deltaPart{sourceLoc: sourceLoc, size: size, data: ar})
} //                                                                       write

// end
