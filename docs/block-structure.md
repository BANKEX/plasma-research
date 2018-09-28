# Plasma blockchain specifications

## Block structure
We expect that you are familiar with notion of Input and Output in UTXO based model.     

Block is RLP Encoded object with the following structrue:

![block](https://raw.githubusercontent.com/BANKEX/plasma-research/master/docs/assets/block.svg?sanitize=true)
```
{
	blockNumber (uint32),
	previousBlockHash (uint256),
	merkleRoot (uint256),
	signature (65 bytes secp256k1),
	transactions: (array of Transaction)
}
```

## Transaction structure
Transaction RLP encoded object that has standard standard structure in terms of UTXO model.
```
{
	inputs (array of input),
	outputs (array of output),
	metadata (metadata),
	signatures (array of uint65 secp256k1)
}
```
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
Input is standard input object in terms of UTXO model extended with assetId property.
`assetId` is identifier of asset in our multi asset plasma.
TODO: add signature
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
Output is RLP encoded object extended with assetId property.
`assetId` is identifier of asset in our multi asset plasma.
```
{
	owner (uint160), 
	assetId (uint160), 
	amount (uint256)
}
```

#### Signatures

