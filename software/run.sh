#!/bin/bash

# This is public address of operator at test geth node
ADDRESS="0x9b72b510f184e16bce51dfd7348ba474ce30b6ed"

# remove old node at demon
sudo docker rm -f geth

# run test geth node at demon (development config at truffle)
sudo docker run -it -d -p 9545:8545 --name geth -p 8546:8546 -p 30303:30303 erage/gethplasma

cd smartContract
sleep 3
# give plasma owner 100 ether
( echo "eth.sendTransaction({from: eth.accounts[0],to:'${ADDRESS}',value: web3.toWei(100, 'ether')})" ) | geth attach http://127.0.0.1:9545

# deploy bankex plasma contract to ethereum
truffle compile
node migrations/test_deploy.js

cd ../commons/operator

# run operator
go run operator.go &

cd ../verifier

# run verifier
go run verifier.go
