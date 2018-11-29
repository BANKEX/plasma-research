package test

import (
	"fmt"
	"math/big"
	"math/rand"
	"reflect"
	"testing"

	"./plasmacrypto"
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

func mulAllUint32(i []uint32) *big.Int {
	res := big.NewInt(1)
	for _, item := range i {
		res.Mul(res, big.NewInt(int64(item)))
	}
	return res
}

func mulAllBigInt(i []*big.Int) *big.Int {
	res := big.NewInt(1)
	for _, item := range i {
		res.Mul(res, item)
	}
	return res
}

func TestRSAAccumulators(t *testing.T) {
	block_slices := [][]*slice.Slice{
		{
			&slice.Slice{Begin: 1, End: 2},
			&slice.Slice{Begin: 4, End: 100},
			&slice.Slice{Begin: 150, End: 170},
			&slice.Slice{Begin: 560, End: 800},
		}, {
			&slice.Slice{Begin: 120, End: 130},
			&slice.Slice{Begin: 300, End: 400},
			&slice.Slice{Begin: 1000, End: 1050},
			&slice.Slice{Begin: 1600, End: 1700},
		},
	}

	accuchain := make([]*plasmacrypto.Accumulator, 0)
	multipliers := make([]*big.Int, 0)
	accuchain = append(accuchain, new(plasmacrypto.Accumulator).SetInt(big.NewInt(17)))
	multipliers = append(multipliers, big.NewInt(1))

	for _, blocks := range block_slices {
		acc := accuchain[len(accuchain)-1]
		mult := big.NewInt(1)
		for _, s := range blocks {
			r := slice.LogProofInclusion(s.GetAlignedSlices())
			acc.BatchAccumulate(r)
			for _, m := range r {
				mult.Mul(mult, new(big.Int).SetUint64(uint64(m)))
			}
			accuchain = append(accuchain, acc)
			multipliers = append(multipliers, mult)
		}
	}

	r := slice.LogProofInclusion((&slice.Slice{Begin: 1, End: 2}).GetAlignedSlices())
	r_mul := mulAllUint32(r)
	p := plasmacrypto.GenProof(accuchain[0], accuchain[2], new(big.Int).Div(mulAllBigInt(multipliers[1:3]), r_mul), r_mul)
	// p.beta = 0 => included
	fmt.Println(p)

	r = slice.LogProofInclusion((&slice.Slice{Begin: 3, End: 4}).GetAlignedSlices())
	r_mul = mulAllUint32(r)
	p = plasmacrypto.GenProof(accuchain[0], accuchain[2], new(big.Int).Div(mulAllBigInt(multipliers[1:3]), r_mul), r_mul)
	// p.beta = 10523688520131956578713779864521885279237926185629405864639170870135220351587134108533202597008847969123508724992252667170 => not included

	fmt.Println(p)
}
