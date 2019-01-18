package types

type Uint160 []byte //20
type Uint256 []byte //32

// RSA Accumulator
type Uint2048 []byte //256

// Signature 65 bytes long ECDSA signature encoded in RSV format
// R(32) bytes S(32) bytes  V(1) byte
type Signature []byte //65
