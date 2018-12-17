package utils

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"fmt"
	"testing"

	// "github.com/BANKEX/plasma-research/src/node/blockchain"
	"github.com/BANKEX/plasma-research/src/node/config"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func TestSign(t *testing.T) {
	data := []byte("hello")
	hash := Keccak256(data)
	privateKeyBytes, err := hex.DecodeString(config.GetOperator().MainAccountPrivateKey[2:])
	if err != nil {
		t.Fatal(err)
	}
	publicAddress, err := hex.DecodeString(config.GetOperator().MainAccountPublicKey[2:])
	if err != nil {
		t.Fatal(err)
	}
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		t.Fatal(err)
	}
	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		t.Fatal(err)
	}

	publicKeyBytes, err := crypto.Ecrecover(hash, signature)
	if err != nil {
		t.Fatal(err)
	}
	publicKey, err := crypto.UnmarshalPubkey(publicKeyBytes)
	if err != nil {
		t.Fatal(err)
	}
	addr := crypto.PubkeyToAddress(*publicKey).Bytes()
	if bytes.Compare(publicAddress, addr) != 0 {
		t.Fatal("Wrong signature")
	}
}

func ToECDSAPub(pub []byte) *ecdsa.PublicKey {
	if len(pub) == 0 {
		return nil
	}
	x, y := elliptic.Unmarshal(secp256k1.S256(), pub)
	return &ecdsa.PublicKey{Curve: secp256k1.S256(), X: x, Y: y}
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
	sig, _ := crypto.Sign(hashRaw, privateKey)
	rec, _ := crypto.Ecrecover(hashRaw, sig)
	pubKey := ToECDSAPub(rec)
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	fmt.Println(recoveredAddr)
}
