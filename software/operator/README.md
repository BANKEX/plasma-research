# Backend subsystems:

## Transaction processor:

Acepts signed transactions from clients using the REST API, validates them, and stores in DB

```
      / processor
users - ...         --> DB
      \ processor

        (tx processing)
```

1. Check that transaction is well formed
2. Check that transaction fee is valid
3. Check in DB that inputs are valid and unspent
4. Add transaction to DB


## Block generator

Periodically assembles pending transactions into a block and uploads them to the durable storage

```
DB -> block generator --> S3

     (assemble block)
```

1. Periodically pull all transactions to be included in the next block
2. Assemble the block
3. Upload to a public durable storage (AWS S3 / DO Dpaces)
4. Upload block header to the Plasma contract on Ethereum

## Block submitter

Watches for new blocks uploaded to the durable storage and submits their headers to Plasma smart contract

```
S3 -> block generator --> Smart Contract

     (submit block)
```


## UTXO uploader

When a new block is successfully uploaded to the storage, add its outputs to DB

```
S3 -> block watcher -> DB
 
 (put new utxos into db)
```

## Event listener

Listens for events on Plasma smart contract and processes them

* Creates `Deposit` transactions to the plasma chain
* Removes Withdrawn UTXOs from the plasma chain
* Submits challenges (TODO)


```
Smart Contract -> event listener --> DB
								 \-> Smart Contract challenge

         (process deposit/withdraw/challenge events)
```
