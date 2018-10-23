package utils

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
)

// StructureToBytes allows to convert structure into bytes array
func StructureToBytes(structure interface{}) []byte {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
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