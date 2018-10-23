package ethereum

import (
	"../utils"
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
)

type wallet struct {
	privateKey privateKey `json:"privateKey"`
	publicKey  publicKey  `json:"publicKey"`
}

type privateKey struct {
	k *ecdsa.PrivateKey
}

type publicKey struct {
	k *ecdsa.PublicKey
}

// Wallet allows to create new ecdsa private key or from existing string private key in []byte format
func Wallet(_pvK ...[]byte) (*wallet, error) {
	var pvK *ecdsa.PrivateKey
	var err error

	if _pvK != nil {
		pvK, err = crypto.ToECDSA(_pvK[0])
	} else {
		pvK, err = crypto.GenerateKey()
	}

	pbK := &ecdsa.PublicKey{Curve: pvK.Curve, X: pvK.X, Y: pvK.Y}

	if err != nil {
		return nil, err
	}

	return &wallet{
		privateKey: privateKey{pvK},
		publicKey: publicKey{pbK},
	}, nil
}

// PrivateKey allows to get string private key
func (w *wallet) PrivateKey() *privateKey  {
	return &w.privateKey
}

func (pvK *privateKey) Bytes() []byte {
	return pvK.k.D.Bytes()
}

func (pvK *privateKey) Hex() string {
	return hex.EncodeToString(pvK.k.D.Bytes())
}

// Allows to get public key associated with private key
func (w *wallet) PublicKey() *publicKey {
	return &w.publicKey
}

func (pubKey *publicKey) Bytes() []byte {
	return pubKeyToBytes(pubKey)
}

func (pubKey *publicKey) Hex() string {
	return utils.BytesToHexString(pubKeyToBytes(pubKey))
}

func pubKeyToBytes(pubKey *publicKey) []byte {
	return crypto.FromECDSAPub(pubKey.k)
}

// Allows to get string address
func (w *wallet) Address() []byte  {
	return crypto.PubkeyToAddress(w.privateKey.k.PublicKey).Bytes()
}

type Signature struct {
	raw       []byte			`json:"raw"`
	hash      []byte			`json:"hash"`
	R     	  []byte			`json:"R"`
	S         []byte			`json:"S"`
	V         []byte			`json:"V"`
}

// TODO: rewrite using stream
// Sign allows to sign any byte array via user private key
func (w *wallet) Sign(data []byte) (*Signature, error) {
	dataHash := crypto.Keccak256(data)
	s, err := crypto.Sign(dataHash, w.privateKey.k)
	if err != nil {
		return nil, err
	}

	return &Signature{
		s,
		dataHash,
		s[:31],
		s[32:63],
		s[64:],
	}, nil
}

// Bytes allows to get byte array of signature
func (s *Signature) Bytes() []byte {
	return s.raw
}

// HexString allows to get hexadecimal string of signature
func (s *Signature) HexString() string {
	return utils.BytesToHexString(s.raw)
}

// GetHash allows to get signed data hash
func (s *Signature) GetHash() []byte {
	return s.hash
}

// GetPublicKeyFromSignature allows to get Public Key from signature
//func GetPublicKeyFromSignature(s *Signature) *publicKey {
//	pubKey, err := crypto.SigToPub(s.hash, s.raw)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	return &publicKey{pubKey}
//}

func VerifySignature(pubKey []byte, hash []byte, signature []byte) bool {
	signatureNoRecoverID := signature[:len(signature)-1]
	return crypto.VerifySignature(pubKey, hash, signatureNoRecoverID)
}
