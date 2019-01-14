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

type verifierConfig struct {
	VerifierPort          int    `env:"verifier_port" envDefault:"8080"`                                                                          // port where verifier server runs
	VerifierPrivateKey    string `env:"main_account_private_key" envDefault:"0x2bdd21761a483f71054e14f5b827213567971c676928d9a1808cbfa4b7501201"` // private key of account who deploy plasma contract and who push blocks to it (operator)
	VerifierPublicKey     string `env:"main_account_public_key" envDefault:"0x6704Fbfcd5Ef766B287262fA2281C105d57246a6"`                          // public key of account who deploy plasma contract and who push blocks to it (operator)
	PlasmaContractAddress string `env:"plasma_contract_address" envDefault:"0x7cc4b1851c35959d34e635a470f6b5c43ba3c9c9"`                          // address of plasma smart contract
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
