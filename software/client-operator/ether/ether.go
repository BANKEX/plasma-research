package ether

import "github.com/ethereum/go-ethereum/crypto"

func GetTxHash(tx []byte) []byte {
	return crypto.Keccak256(tx)
}
