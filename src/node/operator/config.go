package operator

import (
	"github.com/caarlos0/env"
)

type Config struct {
	OperatorPort          int    `env:"operator_port" envDefault:"3001"`
	MainAccountPrivateKey string `env:"main_account_private_key" envDefault:"0x2bdd21761a483f71054e14f5b827213567971c676928d9a1808cbfa4b7501200"`
	MainAccountPublicKey  string `env:"main_account_public_key" envDefault:"0xDf08F82De32B8d460adbE8D72043E3a7e25A3B39"`
	PlasmaContractAddress string `env:"plasma_contract_address" envDefault:"0x08190c080ff35a9686674ee77d4e5e2f01064886"`
	GethHost              string `env:"geth_host" envDefault:"http://127.0.0.1:8545"`
	StartingBlock         uint64 `env:"starting_block" envDefault:"0"`
}

// GetOperator gets operator config instance.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
