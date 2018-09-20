# Open Plasma Architecture

## Description
Our plasma implementation target to be fast enough to provide speed and bandwidth applicable for the NFC payments, transactions from IoT devices and mid-frequency trading on decentralized exchanges.

Another goal is to provide support of multi assets features in plasma. It means we are going to embed such a notion as the asset Id into the low-level transaction objects. 

Below we are going to describe technical decisions and trade-offs that we take in our Plasma Implementation.

IMPORTANT
- TODO: describe process of creation different type of assets
- TODO: describe rules and process of exchange different type of assets
 

## Plasma owner application
Plasma owner application, i.e., block producer, is the core part of Plasma.
It receives signed transactions from plasma clients and builds blocks contains non-controversial transactions with the high speed.
Headers of that blocks are published on the smart contract with a certain frequency that is lower than a frequency of block producing in plasma.
- TODO: Some notions or specifications of block creations frequency in Plasma
 
Headers of the block include the root hash of Merkle tree that has the hash of transaction objects in the leaves.
- TODO: describe block header structure in details
Our implementation of block producer apps has written in Go and available in open source
The main bottle-neck that we meet is a database that should maintain an enormous amount of UTXO object and perform high-speed read-write operations and be fault tolerant at the same time.
We were able to resolve that issue by using multi-instance deployment of Foundation DB. 
Foundation DB is high-performance key-value storage that has almost linear performances scaling (linear to the amount of Foundation DB deployments)
- TODO: Describe the architecture of app the app - Foundation DB nodes, go app architecture - workers
- TODO: Design and add the reference to the separate page that should contain REST or GraphQL API of worker app


### Block structure
In our implementation Block is JSON object that includes transactions encoded on [RLP](https://github.com/ethereum/wiki/wiki/RLP) format
We use RLP because it supported in the ethereum smart contract.
The detailed structure of the block is provided [here](https://github.com/BANKEX/plasma-research/blob/master/docs/block-structure.md)

## Plasma client application
The client-side app is an integral part of Plasma Protocol it entrusted with functions of blocks verification that published by plasma owner.
In case of violations like double spend, plasma client can launch exit game procedure via triggering method on the smart contract.

- TODO: describe checks that verifier do, or put a reference to the standard verification algorithms that we had implemented
- TODO: describe the behavior of the plasma client if blocks are no published at all (no updates on smart contract)
- TODO: describe the behavior of the plasma client if blocks header are published on the smart contact but block itself no published in the plasma network
- TODO: describe withdraw and deposit from client-side point of view
- TODO: mention that the balance of the user is the total amount of UTXOx
- TODO: mention that is cheaper to do withdraw after merging serval UTXOs

Besides primary functionality in further versions of plasma, we are going to add fast withdraw feature to plasma.
- TODO: Describe the implementation of fast withdraw

## Ethereum Smart Contract
The smart contract is key part of that links Plasma side-chain and Ethereum main-net.
It guarantees that plasma operator can't steal funds of plasma participants and any of the participants can send withdraw the request at any time

TODO: Полная спецификация на все методы и структуры данных смарт контракта.

## Appendixes

### Exit game

We suppose to use standard [More Viable Plasma](https://ethresear.ch/t/more-viable-plasma/2160) exit game with some additions.
As we are going to support atomic swaps with multiple transaction source owners, one of a participant may have not all signatures of the transaction. 
So, we need to set up one more exit game branch with signatures collection.

![exit game schema](https://raw.githubusercontent.com/BANKEX/plasma-research/master/docs/assets/plasma_exit_game.svg?sanitize=true)


