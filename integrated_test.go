// -----------------------------------------------------------------------------
// github.com/balacode/go-delta                    go-delta/[integrated_test.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package delta

// This unit test checks the functioning of the entire module.
// It calls Make(), Delta.Apply(), Delta.Bytes() and delta.Load().

import (
	"bytes"
	"testing"
)

// go test --run Test_Integrated_
func Test_Integrated_(t *testing.T) {
	if PrintTestNames {
		printTestName()
	}
	vals := [][]byte{
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
				d := Make(a, b)
				ar = d.Bytes()
			}
			var d Delta
			var err error
			d, err = Load(ar)
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
} //                                                            Test_Integrated_

// end
