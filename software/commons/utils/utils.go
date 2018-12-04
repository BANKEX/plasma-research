package utils

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

// StructureToBytes allows to convert structure into bytes array
func StructureToBytes(structure interface{}) []byte {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	// todo unhandled error
	enc.Encode(structure)
	return network.Bytes()
}

// PrivateKeyStringToBytes allows to convert hexadecimal private key string into bytes
func PrivateKeyStringToBytes(pvK string) []byte {
	p, err := hex.DecodeString(pvK)
	if err != nil {
		panic(err.Error())
	}
	return p
}

// BytesToHexString allows to convert byte array to hexadecimal string
func BytesToHexString(s []byte) string {
	return hex.EncodeToString(s)
}

// SliceToBytesArray allows to convert slice to byte array for signature
func SliceToBytesArray(s []byte) [65]byte {
	var dst [65]byte
	copy(dst[:], s[:65])
	return dst
}

// FromByteArrayToSlice allows to convert byte array to slice for signature
func FromByteArrayToSlice(s [65]byte) []byte {
	return s[:64]
}

func EncodeToRLP(obj interface{}) ([]byte, error) {
	b := new(bytes.Buffer)
	err := rlp.Encode(b, obj)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// Sign calculates an ECDSA signature.
func Sign(data, key []byte) (signature []byte, error error) {
	hash := crypto.Keccak256(data)
	privateKey, err := crypto.ToECDSA(key)
	if err != nil {
		return nil, err
	}
	return crypto.Sign(hash, privateKey)
}

// VerifySignature checks that the given public key created signature over hash.
// The public key should be in compressed (33 bytes) or uncompressed (65 bytes) format.
// The signature should have the 65 byte [R || S || V] format.
func VerifySignature(pubKey, hash, signature []byte) bool {
	signatureNoRecoverID := signature[:len(signature)-1]
	return crypto.VerifySignature(pubKey, hash, signatureNoRecoverID)
}
