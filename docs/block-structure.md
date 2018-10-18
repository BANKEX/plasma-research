# Plasma blockchain specifications

## Block structure
We expect that you are familiar with the notion of Input and Output in UTXO based model.

Block is RLP Encoded object with the following structure:

![block](https://raw.githubusercontent.com/BANKEX/plasma-research/master/docs/assets/block.svg?sanitize=true)
```
{
    blockNumber (uint32),
    previousBlockHash (uint256),
    merkleRoot (uint256),
    signature (Signature),
    transactions (array of Transaction)
}
```

## Transaction structure
Transaction RLP encoded object that has a standard structure in terms of UTXO model.
```
{
    inputs (array of input),
    outputs (array of output),
    metadata (metadata),
    signatures (array of signatures)
}
```
Note: Lenght of signatures array must be equal to 1 or two. A single transaction can't include inputs from more than two different owners

#### Metadata
At the moment, that object contains only one field `max_block_id`. 
That field says that transaction should be included into the block which number is less or equal to `max_block_id`.
```
{
    max_block_id (uint32)
}
```
That field will protect the end user in the following scenario:
1) User send valid signed transaction
2) Plasma operator keep transaction for a long time
3) User see that transaction wasn't included in a block for a long time
4) User send exit request on the smart contract
5) Plasma operator sees that and put the transaction into a block that was withheld and challenge exit of the honest user.

#### Input
Input is standard input object in terms of UTXO model extended with the assetId property.
`assetId` is the identifier of an asset in our multi-asset plasma.
TODO: add the signature
```
{
    owner (uint160), 
    blockIndex (uint32), 
    txIndex (uint32), 
    outputIndex (uint8), 
    assetId (uint160), 
    amount (uint256)
}
```

#### Output:
The output is RLP encoded object extended with an assetId property.
`assetId` is the identifier of an asset in our multi-asset plasma should have corresponding value on Smart Contract
```
{
    owner (uint160), 
    assetId (uint160), 
    amount (uint256)
}
```

#### Signature
Signature is 65 bytes long secp256k1 based signature with the following structure.
```
{
    R (uint256), 
    S (uint256), 
    V (uint8)
}
```
