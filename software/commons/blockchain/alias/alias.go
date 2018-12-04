package alias

type Uint160 [20]byte
type Uint256 [32]byte

// RSA Accumulator
type Uint2048 [256]byte

// Signature 65 bytes long ECDSA signature encoded in RSV format
// R(32) bytes S(32) bytes  V(1) byte
type Signature [65]byte

func ToUint160(slice []byte) Uint160 {
	var arr [20]byte
	copy(arr[:], slice[:20])
	return arr
}

func ToUint256(slice []byte) Uint256 {
	var arr [32]byte
	copy(arr[:], slice[:32])
	return arr
}
