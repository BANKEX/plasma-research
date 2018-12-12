package blockchain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func Sign(hash []byte, prv *ecdsa.PrivateKey) (sig []byte, err error) {
	if len(hash) != 32 {
		return nil, fmt.Errorf("hash is required to be exactly 32 bytes (%d)", len(hash))
	}

	sig, err = secp256k1.Sign(hash, common.LeftPadBytes(prv.D.Bytes(), prv.Params().BitSize/8))
	return
}

func Ecrecover(hash, sig []byte) ([]byte, error) {
	return secp256k1.RecoverPubkey(hash, sig)
}

func ToECDSAPub(pub []byte) *ecdsa.PublicKey {
	if len(pub) == 0 {
		return nil
	}
	x, y := elliptic.Unmarshal(secp256k1.S256(), pub)
	return &ecdsa.PublicKey{Curve: secp256k1.S256(), X: x, Y: y}
}

func TestSMT(t *testing.T) {
	// res:=[]Transaction{Transaction{UnsignedTransaction{[]Input{}, []Output{}, Metadata}, []Signature{}}}
	data := []byte(`
	[
		{
			"inputs": [
				{
					"blockNumber":1,
					"txNumber":1,
					"outputNumber":1,
					"owner":[212, 94, 140, 187, 90, 4, 197, 233, 140, 235, 41, 216, 173, 145, 71, 238, 13, 15, 62, 194],
					"slice": {
						"begin": 1,
						"end": 2
					}
				}

			],
			"outputs": [
				{
					"owner":[212, 94, 140, 187, 90, 4, 197, 233, 140, 235, 41, 216, 173, 145, 71, 238, 13, 15, 62, 194],
					"slice": {
						"begin": 1,
						"end": 2
					}
				}
			],
			"metadata": {
				"maxBlockNumber":100
			}
		},

		{
			"inputs": [
				{
					"blockNumber":1,
					"txNumber":1,
					"outputNumber":1,
					"owner":[212, 94, 140, 187, 90, 4, 197, 233, 140, 235, 41, 216, 173, 145, 71, 238, 13, 15, 62, 194],
					"slice": {
						"begin": 4,
						"end": 5
					}
				}

			],
			"outputs": [
				{
					"owner":[212, 94, 140, 187, 90, 4, 197, 233, 140, 235, 41, 216, 173, 145, 71, 238, 13, 15, 62, 194],
					"slice": {
						"begin": 4,
						"end": 5
					}
				}
			],
			"metadata": {
				"maxBlockNumber":100
			}
		}
	]
	`)

	q := make([]Transaction, 0)
	json.Unmarshal(data, &q)
	var r *SumMerkleTree
	r, _ = NewSumMerkleTree(q)
	_ = r.GetLeaves()

	// fmt.Println(r.MerkleProof(5))

}

func TestSig(t *testing.T) {
	message := "hello"
	hashRaw := crypto.Keccak256([]byte(message))
	privateKey, _ := crypto.HexToECDSA("69b39aa2fb86c7172d77d4b87b459ed7643c1e4b052536561e08d7d25592b373")

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println(fromAddress)
	_ = fromAddress
	sig, _ := Sign(hashRaw, privateKey)
	rec, _ := Ecrecover(hashRaw, sig)
	pubKey := ToECDSAPub(rec)
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	fmt.Println(recoveredAddr)
}
