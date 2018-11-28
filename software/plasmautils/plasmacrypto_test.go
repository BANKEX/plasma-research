package test

import (
	"math/big"
	"testing"

	"./plasmacrypto"
	"./testtools"
)

func TestRSAAccumulator(t *testing.T) {
	a := new(plasmacrypto.Accumulator).SetInt(big.NewInt(17))
	a.Accumulate(19)
	res, _ := new(big.Int).SetString("239072435685151324847153", 10)
	ttl.Assert(t, a.Value().Cmp(res) == 0, "Wrong accumulator value")

	a = new(plasmacrypto.Accumulator).SetInt(big.NewInt(17))
	a.BatchAccumulate([]uint32{3, 5})
	res, _ = new(big.Int).SetString("2862423051509815793", 10)
	ttl.Assert(t, a.Value().Cmp(res) == 0, "Wrong accumulator value")
}
