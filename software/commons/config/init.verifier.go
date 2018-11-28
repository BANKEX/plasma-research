package config

import (
	"flag"
	"log"
	"path/filepath"
)

func InitConfig() VerifierConfig {
	defaultConfigPath, _ := filepath.Abs("../config/config.verifier.json")
	configFileName := flag.String("c", defaultConfigPath, "config file for verifier")
	flag.Parse()
	_, conf, err := ReadConfig(*configFileName, "v")
	if err != nil {
		log.Fatal(err)
	}
	if conf.Verifier_port == 0 {
		log.Fatal("Unmarshalling error!!!")
	}

	SmartContractAddress = conf.Plasma_contract_address

	return conf
}
