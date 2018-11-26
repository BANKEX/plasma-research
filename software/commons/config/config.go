package config

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"log"
)

type VerifierConfig struct {
	Verifier_port            int    `json:verifier_port`
	Main_account_private_key string `json:main_account_private_key`
	Main_account_public_key  string `json:main_account_public_key`
	Plasma_operator_address  string `json:plasma_operator_address`
	Geth_account             string `json:geth_account`
}

type OperatorConfig struct {
	Operator_port            int    `json:operator_port`
	Main_account_private_key string `json:main_account_private_key`
	Main_account_public_key  string `json:main_account_public_key`
	Geth_account             string `json:geth_account`
}


// role: v - verifier, o - operator
func ReadConfig(fileName string, role string) (OperatorConfig, VerifierConfig, error) {
	var vConf VerifierConfig
	var oConf OperatorConfig

	f, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	byteValue, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println(err)
	}

	if role == "v" {
		err = json.Unmarshal(byteValue, &vConf)
		if err != nil {
			log.Println(err)
		}

	} else if role == "o"{
		err = json.Unmarshal(byteValue, &oConf)
		if err != nil {
			log.Println(err)
		}
	}


	return oConf, vConf, nil

}


