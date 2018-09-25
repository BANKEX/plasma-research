# Plasma blockchain specifications

## Block structure
We expect that you are familiar with notion of Input and Output in UTXO based model.     

Block is JSON object that consist of two fields header of the block and array of transaction 
```
{
	header: Blockheader,
	transactions: Array<Transaction>
}
```

Block header:
- TODO: we can't encode uint256 in JSON directly. What encoding we are going to use here base64?
- TODO: the same for signature
MetaInfo: rootMerkleHash (uint256, based on keccak256) was removed by @EnoRage
```
{
	blockNumber (uint32),
	previousBlockHash (uint256),
	merkleRoot (uint256),
	signature (65 bytes secp256k1),
	transactions: Array<Transaction>
}
```

#### Transaction
Transaction RLP encoded object that has standard standard structure in terms of UTXO model.
- TODO: describe what metadata contains
- TODO: describe type of input1...
- TODO: describe type of output1...
```
[
	input1, ..., input6,
	output1, ..., output6,
	metadata,
	signature1, signature2 (65 bytes secp256k1)
]
```

#### Input
Input is standard input object in terms of UTXO model extended with assetId property.
`assetId` is identifier of asset in our multi asset plasma.
TODO: add signature
```
[
	owner (uint160), 
	blockIndex (uint32), 
	txIndex (uint32), 
	outputIndex (uint8), 
	assetId (uint160), 
	amount (uint256)
]
```

#### Output:
Output is RLP encoded object extended with assetId property.
`assetId` is identifier of asset in our multi asset plasma.
```
[
	owner (uint160), 
	assetId (uint160), 
	amount (uint256)
]
```

Metadata
- TODO: describe why we need max_block_id
- TODO: if it's only shell we include id directly to the transaction object 
```
[
	max_block_id (uint32)
]
```
