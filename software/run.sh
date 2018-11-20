#!/usr/bin/env bash
cd smartContract
sudo docker rm -f ganache
sudo docker run -d -p 9545:8545 --name ganache trufflesuite/ganache-cli:latest
truffle migrate --network ganache
