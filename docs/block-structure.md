# Plasma blockchain specifications

## block structure

block:

```
[
	blockNumber (uint32),
	previousBlockHash (uint256),
	merkleRoot (uint256),
	signature (65 bytes secp256k1),
	transaction1, ..., transactionN
l
```

transaction:

```
[
	input1, ..., input6,
	output1, ..., output6,
	metadata,
	signature1, signature2 (65 bytes secp256k1)
l
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
l
```

output:

```
[
	owner (uint160), 
	assetId (uint160), 
	amount (uint256)
l
```

metadata:

```
[
	max_block_id (uint32)
l
```
