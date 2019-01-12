package operator

import (
	"github.com/caarlos0/env"
)

type Config struct {
	OperatorPort          int    `env:"operator_port" envDefault:"3001"`
	MainAccountPrivateKey string `env:"main_account_private_key" envDefault:"0x240d6ad83930067d82e0803696996f743acd78d8fa6a5f6e4f148fd9def37c55"`
	MainAccountPublicKey  string `env:"main_account_public_key" envDefault:"0x9b72b510f184e16bce51dfd7348ba474ce30b6ed"`
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
