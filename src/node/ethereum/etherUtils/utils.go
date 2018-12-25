package etherUtils

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"regexp"
)

func IsValidAddress(iaddress interface{}) error {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	ok := false
	switch v := iaddress.(type) {
	case string:
		ok = re.MatchString(v)
		if !ok {
			err := errors.New("failed to validate address")
			return err
		}
		return nil
	case common.Address:
		ok = re.MatchString(v.Hex())
		if !ok {
			err := errors.New("failed to validate address")
			return err
		}
		return nil
	default:
		err := errors.New("failed to validate address")
		return err
	}
}
func ConvertStringPrivateKeyToRaw(privateKey string) (*ecdsa.PrivateKey, error) {
	rawPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}
	return rawPrivateKey, nil

}

func ConvertPublicKeyToAddress(publicKeyECDSA *ecdsa.PublicKey) common.Address {
	return crypto.PubkeyToAddress(*publicKeyECDSA)
}
