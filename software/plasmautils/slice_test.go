package test

import (
	"fmt"
	"reflect"
	"testing"

	"./slice"
	"./testtools"
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
