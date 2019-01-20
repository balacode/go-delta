// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-01-20 23:14:07 DC874B                  go-delta/[integrated_test.go]
// -----------------------------------------------------------------------------

package delta

// This unit test checks the functioning of the entire module.
// It calls Make(), Delta.Apply(), Delta.Bytes() and loadDelta().

import (
	"bytes"
	"testing"
)

// go test --run TestIntegrated
func TestIntegrated(t *testing.T) {
	if !RunExperiments {
		return
	}
	if PrintTestNames {
		printTestName()
	}
	var vals = [][]byte{
		ab(""),
		ab(" "),
		ab(AtoZ),
		ab(AtoM),
		//
		ab("start" + Nums),
		ab(Nums + "middle" + Nums),
		ab(Nums + Nums + "end"),
		//
		ab(
			"Lorem ipsum dolor sit amet, consetetur sadipscing elitr," +
				" sed diam nonumy eirmod tempor invidunt ut labore et" +
				" dolore magna aliquyam erat, sed diam voluptua. At vero" +
				" eos et accusam et justo duo dolores et ea rebum. Stet" +
				" clita kasd gubergren, no sea takimata sanctus est Lorem" +
				" ipsum dolor sit amet. Lorem ipsum dolor sit amet," +
				" consetetur sadipscing elitr, sed diam nonumy eirmod" +
				" tempor invidunt ut labore et dolore magna aliquyam erat," +
				" sed diam voluptua. At vero eos et accusam et justo duo" +
				" dolores et ea rebum. Stet clita kasd gubergren, no sea" +
				" takimata sanctus est Lorem ipsum dolor sit amet. Lorem" +
				" ipsum dolor sit amet, consetetur sadipscing elitr, sed" +
				" diam nonumy eirmod tempor invidunt ut labore et dolore" +
				" magna aliquyam erat, sed diam voluptua. At vero eos et" +
				" accusam et justo duo dolores et ea rebum. Stet clita" +
				" kasd gubergren, no sea takimata sanctus est Lorem ipsum" +
				" dolor sit amet. suscipit lobortis nisl ut aliquip ex ea" +
				" commodo consequat"),
		ab(
			"Lorem ipsum dolor sit amet, consetetur sadipscing elitr"),
		ab(
			" consetetur sadipscing elitr, sed diam nonumy eirmod" +
				" magna aliquyam erat, sed diam voluptua. At vero eos et"),
		ab(
			"sit amet, consetetur sadipscing elitr" +
				" sed diam nonumy eirmod tempor"),
		ab(
			"suscipit lobortis nisl ut aliquip ex ea commodo consequat."),
		ab(
			"Lorem ipsum dolor sit amet, consetetur sadipscing elitr," +
				AtoZ +
				" sed diam voluptua. At vero eos et accusam et justo duo" +
				AtoM +
				" commodo consequat"),
	}
	for _, a := range vals {
		for _, b := range vals {
			var ar []byte
			{
				var d = Make(a, b)
				ar = d.Bytes()
			}
			var d Delta
			var err error
			d, err = loadDelta(ar)
			if err != nil {
				PL("FAILED @1")
				PL("SOURCE:", "\n", string(a))
				PL("TARGET:", "\n", string(b))
				PL("ERROR:", err)
				continue
			}
			var result []byte
			result, err = d.Apply(a)
			if err != nil {
				PL("FAILED @2")
				PL("SOURCE:", "\n", string(a))
				PL("TARGET:", "\n", string(b))
				PL("ERROR:", err)
				continue
			}
			if !bytes.Equal(result, b) {
				PL("FAILED @3")
				PL("SOURCE:", "\n", string(a))
				PL("TARGET:", "\n", string(b))
				PL("RETURNED:", "\n", string(result))
			}
		}
	}
} //                                                              TestIntegrated

//end
