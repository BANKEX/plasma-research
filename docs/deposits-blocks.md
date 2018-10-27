### Working with deposit blocks

We need to work with deposits at both plasma chain and mainnet chain and these chains must be synced. Also, outputs from deposits must be exitable and useful for simple and special plasma exit protocol. Let's use the following approach to reach the target:

* increase blockIndex space from 32 bits up to 64 bits.
* allow plasma operator to fill only blocks from 2^63 up to 2^64-1
* keep blocks from 0 up to 2^63-1 for deposit transactions. One block for one transaction
* Set timestamp of deposit blocks equal to zero to speedup withdraw procedure from deposits and keep consistency between block number and timestamp.
 

Then we get the following UX:

1. User send a deposit to the contract
2. Contract compute single deposit block and put it into plasma chain (we need to spend about 10k gas for hashing)
3. Other UX of contract and plasma can work with such block like with regular plasma block

Why do we need to separate namespace of deposit and regular blocks of plasma:

Deposit blocks are fulfilled at the contract and it is not vulnerable to [blockchain reorgs](https://github.com/BANKEX/plasma-research/blob/master/docs/reorgs-protection.md) and do not depend on regular blocks. On the contrary, regular plasma blocks are vulnerable to reorgs and strictly depend on deposit blocks (if we have any input from deposit block to regular block).
