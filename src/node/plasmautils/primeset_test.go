package test

import (
	"testing"
)

func TestPrimeGeneration(t *testing.T) {
	// TODO: test package not working
	// ttl.Assert(t, primeset.PrimeN(3) == 11, "Wrong prime number generation")
}

// func main() {
// 	n := 5637
// 	x := primeset.PrimeN(n)
// 	r := primeset.ProofN(n)
// 	q := plasmacrypto.Hash([]uint8{byte(x >> 24), byte(x >> 16), byte(x >> 8), byte(x)})
// 	for j, item := range r {
// 		fmt.Println(q, "\t", item)
// 		if n>>uint(j)&1 == 1 {
// 			q = plasmacrypto.Hash(append(item.Data, q.Data...))
// 		} else {
// 			q = plasmacrypto.Hash(append(q.Data, item.Data...))
// 		}
// 	}
// 	fmt.Println(q.String())

// }
