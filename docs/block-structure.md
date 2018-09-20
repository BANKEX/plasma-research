# Plasma blockchain specifications

## block structure

block:

```
[
	blockNumber (uint32),
	previousBlockHash (uint256),
	merkleRoot (uint256),
	signature (65 bytes secp256k1),
	transaction 1,
	...,
	transaction N
]
```

transaction:

```
[
	input1, ..., input6,
	output1, ..., output6,
	metadata,
	signature1, signature2 (65 bytes secp256k1)
]
```

input:

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

output:

```
[
	owner (uint160), 
	assetId (uint160), 
	amount (uint256)
]
```

metadata:

```
[
	max_block_id (uint32)
]
```
