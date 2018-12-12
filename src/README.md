# About

Plasma Cashflow with history reduction with zk-SNARKS.

1. Plasma cashflow, based on specs https://hackmd.io/DgzmJIRjSzCYvl4lUjZXNQ?view and https://ethresear.ch/t/plasma-prime-design-proposal/4222
2. History voids compression on zkSNARK 

Overview more about voids and history reduction here:
https://ethresear.ch/t/plasma-prime-design-proposal/4222
https://ethresear.ch/t/short-s-nt-ark-exclusion-proofs-for-plasma/4438

Main goals for this prototype:
We are going to show plasma spec without lacks of MoreVP, Cash or RSA accumulators.
We think, potentially fully s[nt]ark-driven plasma is better. But this construction is closer to production usage because we use the very simple circuit to exclude voids from the history. It is enough to use only hash functions and Merkle proofs, without any signatures or transaction processing.

What can you see here:
1. SNARK, compressing the voids. For production, usage needs to implement Pedersen hash on bellman or libsnark.
2. prototypes of some kinds of challenges. We do not focus all our attention on challenges, because plasma cash-specific challenges are a solved problem and no novelty here.
3. we implement SumMerkleTree on solidity, golang, python, zokrates, covered by tests
4. Also, you can see the plasma with separate components: root node, thick and thin client, contracts and snark.

## Installation

1. Run install.sh
2. Run ganache
3. Deploy BankexPlasma contract
4. Copy address to config.go at node/config
5. Run operator.go with `go run operator.go`
6. Rum verifier.go with `go run verifier.go`
7. Open http://localhost:8080/frontend to see the dashboard

#### Don't forget to import your private key from ganache to config.go


