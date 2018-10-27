## Reorgs protection

Plasma blocks must be strictly bounded with previous blocks and deposits. We need to provide resistance for blockchain reorgs and censorship of transactions.

Let's protect from reorgs and miners using the link to the mainnet block, including all deposits used inside the current plasma block. Also, we can minimize the reorgs of plasma chain ignoring fresh deposits.

```soldity

function sendPlasmaMerkleRoot(uint currentBlockNumber, address currentBlock, 
  uint protectedBlockNumber, uint protectedBlockValue) onlyOwner public returns(bool)

```

Here we need to check, that plasma block with number `currentBlockNumber - 1` exits. Also, we need to check, that `block.blockhash[protectedBlockNumber] == protectedBlockValue`. 

Plasma operator can use nonce mechanics to help miners to arrange elements in the right order.

That guarantees us that plasma blocks included into the blockchain one by one into the right order and plasma blocks are rightly connected to corresponded deposits.
