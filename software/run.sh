#!/bin/bash
docker stop $(docker ps -aq)
docker run -it -d -p 9545:8545 -p 8546:8546 -p 30303:30303  erage/gethplasma
cd smartContract 
sudo rm -rf build
truffle migrate --network development
