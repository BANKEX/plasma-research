#!/bin/bash

# This is public address of operator at test geth node
ADDRESS="0x9b72b510f184e16bce51dfd7348ba474ce30b6ed"

# remove old node at demon
docker rm -f geth

# run test geth node at demon (development config at truffle)
docker run -it -d -p 8545:8545 --name geth -p 8546:8546 -p 30303:30303 erage/gethplasma

cd ../contracts

#npm i

sleep 3
# give plasma owner 100 ether
( echo "eth.sendTransaction({from: eth.accounts[0],to:'0x9cA4E1F69A3ABD60989864FAd1025095dFCC58F1',value: web3.toWei(100, 'ether')})" ) | geth attach http://127.0.0.1:8545


# deploy bankex plasma contract to ethereum
truffle migrate --reset
#truffle compile
#node migrations/deploy.js



cd ../node/operator/

# run operator
GIN_MODE=release go run operator.go &

cd ../verifier/

# clear db before run - NOT FOR PRODUCTION
# sudo rm -rf database

# run verifier
GIN_MODE=release go run verifier.go
