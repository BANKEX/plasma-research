#!/bin/bash

cd ../contracts

truffle migrate --network ganache

cd build/contracts

cd ..

abigen --abi=BankexPlasma.abi -bin=BankexPlasma.bin --pkg=store --out=plasmacontract.go

cat plasmacontract.go > ../../node/ethereum/plasmacontract/plasmacontract.go
#rm  plasmacontract.go
# https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/en/smart-contract-read/README.md