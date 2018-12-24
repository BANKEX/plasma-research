package config

import (
	"log"

	"github.com/caarlos0/env"
)

type operatorConfig struct {
	OperatorPort          int    `env:"operator_port" envDefault:"3001"`
	MainAccountPrivateKey string `env:"main_account_private_key" envDefault:"0x240d6ad83930067d82e0803696996f743acd78d8fa6a5f6e4f148fd9def37c55"`
	MainAccountPublicKey  string `env:"main_account_public_key" envDefault:"0x9b72b510f184e16bce51dfd7348ba474ce30b6ed"`
	PlasmaContractAddress string `env:"plasma_contract_address" envDefault:"0x08190c080ff35a9686674ee77d4e5e2f01064886"`
	GethHost              string `env:"geth_host" envDefault:"http://127.0.0.1:8545"`
	GethAccount           string `env:"geth_account" envDefault:"ws://127.0.0.1:8545"`
	StartingBlock         uint64 `env:"starting_block" envDefault:"0"`
}

type verifierConfig struct {
	VerifierPort          int    `env:"verifier_port" envDefault:"8080"`                                                                          // port where verifier server runs
	VerifierPrivateKey    string `env:"main_account_private_key" envDefault:"0xe4058d9c3a81b4e95d8e3a17a5f52486a7fc411e57dcd4f6c771dbc2428928e9"` // private key of account who deploy plasma contract and who push blocks to it (operator)
	VerifierPublicKey     string `env:"main_account_public_key" envDefault:"0x9cA4E1F69A3ABD60989864FAd1025095dFCC58F1"`                          // public key of account who deploy plasma contract and who push blocks to it (operator)
	PlasmaContractAddress string `env:"plasma_contract_address" envDefault:"0xb70f898520cd51d2a258ff0af25578289140c2fe"`                          // address of plasma smart contract
	GethHost              string `env:"geth_host" envDefault:"ws://127.0.0.1:8545"`
	OperatorHost          string `env:"operator_host" envDefault:"http://127.0.0.1:3001"`
}

var (
	operatorConfigInstance *operatorConfig
	verifierConfigInstance *verifierConfig
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

// GetVerifier gets verifier config instance.
func GetVerifier() *verifierConfig {
	if verifierConfigInstance == nil {
		verifierConfigInstance = &verifierConfig{}
		err := env.Parse(verifierConfigInstance)
		if err != nil {
			log.Fatalf("error initializing config: %s", err)
		}
	}
	return verifierConfigInstance
}
