package alias

type Uint160 []byte //20
type Uint256 []byte //32

// RSA Accumulator
type Uint2048 []byte //256

// Signature 65 bytes long ECDSA signature encoded in RSV format
// R(32) bytes S(32) bytes  V(1) byte
type Signature []byte //65

type Uint160Bytes [20]byte
type TxHashBytes [20]byte

func Copy20(slice []byte) [20]byte {
	var arr [20]byte
	copy(arr[:], slice[:20])
	return arr
}

func ToUint160Bytes(slice []byte) Uint160Bytes {
	return Copy20(slice)
}

func ToTxHashBytes(slice []byte) TxHashBytes {
	return Copy20(slice)
}
