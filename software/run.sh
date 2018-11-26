#!/bin/bash
export ADDRESS="0x9B72B510F184e16BcE51Dfd7348Ba474cE30b6ed"
echo ${ADDRESS}

sudo docker rm -f geth
# docker rm -f op
sudo docker run -it -d -p 9545:8545 --name geth -p 8546:8546 -p 30303:30303 erage/gethplasma
cd smartContract
sudo rm -rf build
npm i
truffle migrate --network development
( echo "eth.sendTransaction({from: eth.accounts[0],to:'${ADDRESS}',value: web3.toWei(100, 'ether')})" ) | geth attach http://127.0.0.1:9545

cd ..
cd client-operator
cd operator
go run operator.go &
cd ..
# docker build . -t operator
# docker run --name op -p 8080:8080 -d operator
cd verifier
cp config.json ../../test
cd ../../test
go run main.go &
cd ../client-operator/verifier
go run verifier.go
