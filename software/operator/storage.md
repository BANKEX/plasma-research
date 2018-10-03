Schema
======

This documents describes the storage structure and operations on it. Schema and operations are expressed in generic language, 
assuming we are working with ACID compliant NoSQL storage

UTXOs
-----

This directory contains only valid unspent outputs

`/utxo/block:id:output -> { owner, assetId, amount }`

Transactions
------------

Transaction blobs for inclusion in the next block:

`/tx/blockId:hash -> raw tx bytes`

Deposits
--------

Deposit objects corresponding to deposits on smart contract

`/deposit/id -> true`

// TODO - what to do with withdrawn deposits

Misc
----

`/currentBlock -> int32 (current block counter)`
`/lastUploadedBlock -> int32 (last block that was uploaded to S3)`
`/lastProcessedBlock -> int32 (last block which utxos were added to the db)`
`/lastProcessedEthereumBlock -> int32 (last block from which all ethereum events were successfully processed, 
on DB creation initialize this with PlasmaParent contract creation block)` 


DB Operations
=============

Insert tx
---------

```
READ /utxo/block:id:output (1-6x, read inputs for verification)
DELETE /utxo/block:id:output (1-6x, spend inputs)
READ /currentBlock
INSERT /tx/currentBlock:hash (add transaction to the next block)
```

Remove UTXO when `Withdraw` event is detected on Ethereum network
-----------------------------------------------------------------

```
DELETE /utxo/block:id:output
```

### Create UTXO when `Deposit` event is detected on Ethereum network
--------------------------------------------------------------------

```
READ /deposits/id (to make sure we didn't already insert it)
READ /currentBlock
INSERT /tx/currentBlock:hash
INSERT /deposits/id
```

Insert UTXO after block is uploaded
-----------------------------------

After the block is successfully published to S3. This can become a performance bottleneck when number of utxos in block is too large, 
but can be solved by chunking the block and uploading in parallel (and introducing a new `success` variable for each chunk).

```
READ /lastProcessedBlock (to make sure we are not double inserting and are inserting in order)
BULK INSERT /utxo/...
INCREMENT /lastProcessedBlock
```

Create block
------------

All operations here are done in a separate transaction (assuming we have only one block creator).
`lastUploadedBlock` can eqal `block - 1` if block uploader has crashed after incrementing the current block but before uploading it

```
READ /block
READ /lastUploadedBlock
if (lastUploadedBlock == block) {
	INCREMENT /block
}
RANGE READ /tx/<id:...> (id = lastUploadedBlock)
INCREMENT /lastUploadedBlock
```

Cleanup old tx after the block is successfully uploaded to the storage
----------------------------------------------------------------------

```
DELETE /tx/<past block tx range>
```