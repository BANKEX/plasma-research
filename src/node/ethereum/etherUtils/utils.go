package etherUtils

import (
	"errors"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
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
