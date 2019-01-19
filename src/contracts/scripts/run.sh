#!/usr/bin/env bash

# Exit script as soon as a command fails.
set -o errexit

ganache_port=8545

ganache_running() {
  nc -z localhost "$ganache_port"
}

# Ganache use first address to make migration (0x2bdd21761a483f71054e14f5b827213567971c676928d9a1808cbfa4b7501200)
# Since contract is ownable we will assign address to operator
# The second one(0x2bdd21761a483f71054e14f5b827213567971c676928d9a1808cbfa4b7501201) will be assigned to verifier
#
#
# Private Key to Address mapping 
# 0x2bdd21761a483f71054e14f5b827213567971c676928d9a1808cbfa4b7501200 -> 0xDf08F82De32B8d460adbE8D72043E3a7e25A3B39
# 0x2bdd21761a483f71054e14f5b827213567971c676928d9a1808cbfa4b7501201 -> 0x6704Fbfcd5Ef766B287262fA2281C105d57246a6
start_ganache() {
  # We define 10 accounts with balance 1M ether, needed for high-value tests.
  local accounts=(
    --account="0x2bdd21761a483f71054e14f5b827213567971c676928d9a1808cbfa4b7501200,1000000000000000000000000"
    --account="0x2bdd21761a483f71054e14f5b827213567971c676928d9a1808cbfa4b7501201,1000000000000000000000000"
    --account="0x2bdd21761a483f71054e14f5b827213567971c676928d9a1808cbfa4b7501202,1000000000000000000000000"
    --account="0x2bdd21761a483f71054e14f5b827213567971c676928d9a1808cbfa4b7501203,1000000000000000000000000"
    --account="0x2bdd21761a483f71054e14f5b827213567971c676928d9a1808cbfa4b7501204,1000000000000000000000000"
    --account="0x2bdd21761a483f71054e14f5b827213567971c676928d9a1808cbfa4b7501205,1000000000000000000000000"
    --account="0x2bdd21761a483f71054e14f5b827213567971c676928d9a1808cbfa4b7501206,1000000000000000000000000"
    --account="0x2bdd21761a483f71054e14f5b827213567971c676928d9a1808cbfa4b7501207,1000000000000000000000000"
    --account="0x2bdd21761a483f71054e14f5b827213567971c676928d9a1808cbfa4b7501208,1000000000000000000000000"
    --account="0x2bdd21761a483f71054e14f5b827213567971c676928d9a1808cbfa4b7501209,1000000000000000000000000"
  )

  node_modules/.bin/ganache-cli --db="./data/" --gasLimit 0xfffffffffff --host 0.0.0.0 --port "$ganache_port" "${accounts[@]}" --networkId 5777 &
  node_modules/.bin/truffle migrate --network=ganache
  ganache_pid=$!
}

if ganache_running; then
  echo "Using existing ganache instance"
else
  echo "Starting our own ganache instance"
  start_ganache
fi

wait ${ganache_pid}

