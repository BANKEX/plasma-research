package verifier

import (
	"github.com/caarlos0/env"
)

type Config struct {
	VerifierPort          int    `env:"verifier_port" envDefault:"8080"`                                                                          // port where verifier server runs
	VerifierPrivateKey    string `env:"main_account_private_key" envDefault:"0xe4058d9c3a81b4e95d8e3a17a5f52486a7fc411e57dcd4f6c771dbc2428928e9"` // private key of account who deploy plasma contract and who push blocks to it (operator)
	VerifierPublicKey     string `env:"main_account_public_key" envDefault:"0x9cA4E1F69A3ABD60989864FAd1025095dFCC58F1"`                          // public key of account who deploy plasma contract and who push blocks to it (operator)
	PlasmaContractAddress string `env:"plasma_contract_address" envDefault:"0xb70f898520cd51d2a258ff0af25578289140c2fe"`                          // address of plasma smart contract
	GethHost              string `env:"geth_host" envDefault:"ws://127.0.0.1:8545"`
	OperatorHost          string `env:"operator_host" envDefault:"http://127.0.0.1:3001"`
}

// GetVerifier gets verifier config instance.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}