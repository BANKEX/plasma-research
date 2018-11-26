#!/bin/bash

# This is public address of operator at test geth node
ADDRESS="0x9B72B510F184e16BcE51Dfd7348Ba474cE30b6ed"

# remove old node at demon
sudo docker rm -f geth

# run test geth node at demon (development config at truffle)

sudo docker run -it -d -p 9545:8545 --name geth -p 8546:8546 -p 30303:30303 erage/gethplasma

cd smartContract

# clear abi before compiling new
sudo rm -rf build

# install all packages to truffle
npm i

# deploy bankex plasma contract to ethereum
truffle migrate --network development

# give plasma owner 100 ether
( echo "eth.sendTransaction({from: eth.accounts[0],to:'${ADDRESS}',value: web3.toWei(100, 'ether')})" ) | geth attach http://127.0.0.1:9545

cd ..
cd commons

# install all go lang packages
go get -d ./...

cd operator

# run operator
go run operator.go &

cd ..
cd verifier

# run verifier
go run verifier.go
