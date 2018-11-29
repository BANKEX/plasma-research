package primeset

import (
	"fmt"
	"io/ioutil"
	"os"

	crypto "../plasmacrypto"
	"github.com/snjax/gmp"
)

var (
	primeset []uint32
	dataset  []uint8
	diffset  []uint8
)

const PRIMES_LOG2 = 26

const CHUNK_SIZE_LOG2 = 7
const CHUNK_SIZE = 1 << CHUNK_SIZE_LOG2

func SayHelloPrimeset() {
	fmt.Println("Hello!")
}

func tailproof(proof *[]*crypto.Cipher, data []uint32, i int) {

	r := 1
	h := uint(0)
	for r < CHUNK_SIZE {
		if i>>h&1 == 1 {
			*proof = append(*proof, hashprimeset(data[i-r:i]))
			i -= r
		} else {
			*proof = append(*proof, hashprimeset(data[i+r:i+2*r]))
		}
		r <<= 1
		h++
	}
}

func hashprimeset(data []uint32) *crypto.Cipher {
	ldata := len(data)
	if len(data) > 1 {
		l := hashprimeset(data[0 : ldata/2]).Data
		r := hashprimeset(data[ldata/2 : ldata]).Data
		return crypto.Hash(append(l, r...))
	}
	x := data[0]
	return crypto.Hash([]uint8{byte(x >> 24), byte(x >> 16), byte(x >> 8), byte(x)})

}

func initFiles() {
	fmt.Println("Generating prime number assets. Wait some minutes")
	os.MkdirAll("assets", os.ModePerm)
	btmsz := 1 << PRIMES_LOG2 / CHUNK_SIZE
	ttlsz := btmsz*2 - 1

	primeset = make([]uint32, 1<<PRIMES_LOG2)
	diffset = make([]uint8, 1<<PRIMES_LOG2)
	dataset = make([]uint8, ttlsz*crypto.HASH_SIZE)

	var p = new(gmp.Int)
	primeset[0] = 5
	diffset[0] = 2
	for i := 1; i < 1<<PRIMES_LOG2; i++ {
		primeset[i] = p.SetUint64(uint64(primeset[i-1])).NextPrime().Uint32()
		diffset[i] = uint8((primeset[i] - primeset[i-1]) / 2)
	}

	ioutil.WriteFile("assets/primediffs", diffset[:], 0644)

	for i := 0; i < btmsz; i++ {
		copy(dataset[(btmsz-1+i)*crypto.HASH_SIZE:(btmsz+i)*crypto.HASH_SIZE], hashprimeset(primeset[i*CHUNK_SIZE:(i+1)*CHUNK_SIZE]).Data)
	}

	for i := btmsz - 2; i >= 0; i-- {
		copy(dataset[i*crypto.HASH_SIZE:(i+1)*crypto.HASH_SIZE], crypto.Hash(dataset[(2*i+1)*crypto.HASH_SIZE:(2*i+3)*crypto.HASH_SIZE]).Data)
	}

	ioutil.WriteFile("assets/primedataset", dataset[:], 0644)
}

func loadFiles() {
	primeset = make([]uint32, 1<<PRIMES_LOG2)
	diffset, _ = ioutil.ReadFile("assets/primediffs")
	dataset, _ = ioutil.ReadFile("assets/primedataset")
	p := uint32(1)
	for i := 0; i < 1<<PRIMES_LOG2; i++ {
		p += uint32(diffset[i]) * 2
		primeset[i] = p
	}
}

func PrimeN(i int) uint32 {
	return primeset[i]
}

// ProofN MerkleProof from leaf to root in raw binary
func ProofN(i int) []*crypto.Cipher {
	c := i / CHUNK_SIZE
	d := i % CHUNK_SIZE
	res := make([]*crypto.Cipher, 0)
	tailproof(&res, primeset[c*CHUNK_SIZE:(c+1)*CHUNK_SIZE], d)

	f := 1<<PRIMES_LOG2/CHUNK_SIZE + c - 1

	for h := 0; h < PRIMES_LOG2-CHUNK_SIZE_LOG2; h++ {
		if f&1 == 0 {
			res = append(res, new(crypto.Cipher).SetBytes(dataset[(f-1)*crypto.HASH_SIZE:f*crypto.HASH_SIZE]))
		} else {
			res = append(res, new(crypto.Cipher).SetBytes(dataset[(f+1)*crypto.HASH_SIZE:(f+2)*crypto.HASH_SIZE]))
		}
		f = (f - 1) >> 1
	}

	return res
}

func init() {
	if _, err := os.Stat("assets/primedataset"); os.IsNotExist(err) {
		initFiles()
		return
	}
	if _, err := os.Stat("assets/primediffs"); os.IsNotExist(err) {
		initFiles()
		return
	}
	loadFiles()
}
