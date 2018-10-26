### Block structure for solidity

Cut hashes up to 160 bits. It provides us 80bit reliability.

```solidity
struct MerkleProof {
    uint160 from,
    bytes proof,
    uint index
}
```

Use following structure of transactions:

![plasma block](https://gist.githubusercontent.com/snjax/3c99fbed8393a7d4b03afd4361998039/raw/44ad5c9ad47f5109192628001efdc9711b251e1d/plasma_block.svg?sanitize=true)

### Merkle proof index bit map

1. 32 bits for  block to tx merkle proof
1. 1 bit for tx netto proof (must be zero for outputs, inputs, max_blockid and must be one for signatures)
1. 5 bits for merkle proof for outputs, inputs and max_block_id or for signatures.
1. 3 bits for merkle proof for signatures.

So, maximal outputNumber is limited by 31 (must checked on the contract for input structures).

### Data types encoding

``` leaf = keccak256(concat(datatype_byte, abi.encode(data)) ```


### Simple exit challenges

Use following requests:

```

struct SimpleExit {
  Input point
}

struct ExitQueueItem {
  uint160 exitHash,
  uint SFT
}


function withdrawal(Input point, address queuePtr) external payable returns (bool);

function withdrawalChallangeSpend(SimpleExit exit, address queuePtr, 
    Input spend, MerkleProof txProof, MerkleProof spendProof, 
    uint32 maxBlockId, MerkleProof maxBlockIdProof,
    Signature sign, MerkleProof signProof) external returns (bool);

function withdrawalChallangeBlock(SimpleExit exit, address queuePtr, MerkleProof proof) external returns (bool);
```


1. Somebody published exit and the bond. We check the validity of input on the contract and put it into the exit queue with finalization time ordered by `SFT = max(now+REP+MFT, input.timestamp+MFT)`, where `REP=MFT=1` week. Also `withdrawal` emit `SFT` in event.
1. anybody can challenge the spend of the input.
1. anybody can challenge the existence of input inside the blockchain.



Signer guarantees validity of the transaction (or he have a risk to burn his money). That's why it is enough to check the signature and maxBlockId to run `withdrawalChallangeSpend`.

2nd, 3rd, 4th merkle proof are short (linked to tx hash, not to block hash).
.

 
