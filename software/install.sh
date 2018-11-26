#!/bin/bash

cd smartContract

# clear abi before compiling new
sudo rm -rf build

# install all packages to truffle
npm i

cd ..
cd commons

# install all go lang packages
go get -d ./...