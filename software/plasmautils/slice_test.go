package test

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"./slice"
	"./testtools"
	"github.com/snjax/gmp"
)

func TestSlice(t *testing.T) {
	s := &slice.Slice{Begin: 1, End: 8}
	ttl.Assert(t, reflect.DeepEqual(s.GetAlignedSlices(), []uint32{16777216, 8388608, 4194304}), "Wrong prime number generation")
}

func TestInclusionProof(t *testing.T) {
	s := &slice.Slice{Begin: 1, End: 108327}
	r := slice.LogProofInclusion(s.GetAlignedSlices())
	fmt.Println(r)
}

func TestPerformanceVerybig(t *testing.T) {
	n := new(gmp.Int).SetInt64(1)
	for i := 0; i < 100000; i++ {
		n.Mul(n, new(gmp.Int).SetUint64(rand.Uint64()))
	}
	m := new(gmp.Int)
	_, _ = n.DivMod(n, new(gmp.Int).SetUint64(rand.Uint64()), m)

}

func TestGmpBytes(t *testing.T) {
	n := new(gmp.Int).SetUint64(0xffffffffffffffff)
	n = n.Exp(n, gmp.NewInt(2), gmp.NewInt(0))
	fmt.Println(n.String())
	fmt.Println(n.Bytes())

}
