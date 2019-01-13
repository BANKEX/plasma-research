package verifier

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	VerifierPort            int    `env:"verifier_port" envDefault:"8080"`                                                                          // port where verifier server runs
	VerifierPrivateKey      string `env:"main_account_private_key" envDefault:"0x2bdd21761a483f71054e14f5b827213567971c676928d9a1808cbfa4b7501201"` // private key of account who deploy plasma contract and who push blocks to it (operator)
	VerifierEthereumAddress string `env:"main_account_public_key" envDefault:"0x6704Fbfcd5Ef766B287262fA2281C105d57246a6"`                          // public key of account who deploy plasma contract and who push blocks to it (operator)
	PlasmaContractAddress   string `env:"plasma_contract_address" envDefault:"0xb70f898520cd51d2a258ff0af25578289140c2fe"`                          // address of plasma smart contract
	GethHost                string `env:"geth_host" envDefault:"ws://127.0.0.1:8545"`
	OperatorHost            string `env:"operator_host" envDefault:"http://127.0.0.1:3001"`
}

// GetVerifier gets verifier config instance.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	if !common.IsHexAddress(cfg.PlasmaContractAddress) {
		return nil, fmt.Errorf("given contract address %s is invalid", cfg.PlasmaContractAddress)
	}

	return cfg, nil
}
