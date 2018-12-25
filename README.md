# plasma-research [![Build Status](https://travis-ci.org/BANKEX/plasma-research.svg?branch=master)](https://travis-ci.org/BANKEX/plasma-research)
That repo contains docs and code snippets related to BANKEX Plasma implementation.
At the moment we focused on implementation of plasma cashflow model improved by zk-S[NT]ARK exclusion proofs for history compression

# History compression with zk-S[NT]ARK exclusion proofs
At the moment when plasma user is getting a transfer from another user, he wants to ensure that the history of that transfer is valid. To do that ask other parties to provide exclusion proof for the particular slice that was transferred on most of the published plasma blocks.
The size of such exclusion proof based on Merkle proof grows excessively, during the plasma lifetime. That problem was described [here](https://ethresear.ch/t/rsa-accumulators-for-plasma-cash-history-reduction/3739).

In our implementation, we are going to use S[NT]ARK exclusion proofs to compress it. You can find more details [here](https://ethresear.ch/t/short-s-nt-ark-exclusion-proofs-for-plasma/4438)

# Multiassets plasma
By Multi-asset we mean support of deposit, withdraw, transfer and exchange of a wide range of Ethereum tokens standards like ERC20, ERC721, ERC888, etc.. as well as Ethereum it self

To support that we are going to implement the more sophisticated parent smart contract that will be able to take ownership of different tokens types.

Plasma parent contract can't guarantee that actual token contract is implemented safely and fully correctly.
It means that end user should check that token contract he was offered doesn't have a backdors or bugs in the implementation by himself.

Each token type that was deposit to plasma smart contract will get a unique id. That id will be used further in the transactions to describe transfer and exchange operations.

To implement an exchange operation, we need a single transaction signed by both parties. The protocol that will allow users to share transactions for signatures will be mainly implemented in the plasma client.

All the research we do, we going to document in the form of Yellow Paper.
At the moment we have following separate docs that cover different parts of our application:
- [Architecture overview](https://github.com/BANKEX/plasma-research/blob/master/docs/architecture.md)
- [Block structure](https://github.com/BANKEX/plasma-research/blob/master/docs/block-structure.md)
