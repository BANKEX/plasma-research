package plasmacrypto

import (
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

const HashSize = 20

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
	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	res = hash.Sum(res)[0:HashSize]
	return new(Cipher).SetBytes(res)
}

func Hash256(data []byte) *Cipher {
	res := []byte{}
	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	res = hash.Sum(res)[:]
	return new(Cipher).SetBytes(res)
}
