package plasmacrypto

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto/sha3"
)

const HASH_SIZE = 20

type Cipher struct {
	Data []byte
}

func (s *Cipher) SetBytes(c []byte) *Cipher {
	s.Data = c
	return s
}

func (s *Cipher) String() string {
	return hex.EncodeToString(s.Data)
}

func Hash(data []byte) *Cipher {
	res := []byte{}
	hash := sha3.NewKeccak256()
	hash.Write(data)
	res = hash.Sum(res)[0:HASH_SIZE]
	return new(Cipher).SetBytes(res)
}
