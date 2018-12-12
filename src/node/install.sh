#!/bin/bash

cd ././../contracts/

# clear abi before compiling new
sudo rm -rf build

# install all packages to truffle
npm i

cd ../node

# install all go lang packages
go get -d ./...