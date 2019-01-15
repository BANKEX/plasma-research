package config

import (
	"log"

	"github.com/caarlos0/env"
)

type operatorConfig struct {
	OperatorPort          int    `env:"operator_port" envDefault:"3001"`
	MainAccountPrivateKey string `env:"main_account_private_key" envDefault:"0x2bdd21761a483f71054e14f5b827213567971c676928d9a1808cbfa4b7501200"`
	MainAccountPublicKey  string `env:"main_account_public_key" envDefault:"0xDf08F82De32B8d460adbE8D72043E3a7e25A3B39"`
	PlasmaContractAddress string `env:"plasma_contract_address" envDefault:"0x7cc4b1851c35959d34e635a470f6b5c43ba3c9c9"`
	GethHost              string `env:"geth_host" envDefault:"http://127.0.0.1:8545"`
	GethAccount           string `env:"geth_account" envDefault:"ws://127.0.0.1:8545"`
	StartingBlock         uint64 `env:"starting_block" envDefault:"0"`
}

var (
	operatorConfigInstance *operatorConfig
)

// GetOperator gets operator config instance.
func GetOperator() *operatorConfig {
	if operatorConfigInstance == nil {
		operatorConfigInstance = &operatorConfig{}
		err := env.Parse(operatorConfigInstance)
		if err != nil {
			log.Fatalf("error initializing config: %s", err)
		}
	}
	return operatorConfigInstance
}
