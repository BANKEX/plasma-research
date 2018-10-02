# Backend subsystems:

## Transaction processor:

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

3 + 4 are atomic


## Block generator

```
DB -> block generator --> S3
                       \-> Smart Contract

     (assemble block)
```

1. Periodically pull all transactions to be included in the next block
2. Assemble the block
3. Upload to a public durable storage (AWS S3 / DO Dpaces)
4. Upload block header to the Plasma contract on Ethereum

## UTXO uploader

When a new block is successfully uploaded to the storage, add its outputs to DB

```
S3 -> block watcher -> DB
 
 (put new utxos into db)
```


## Event listener

```
Smart Contract -> event listener --> DB
								 \-> Smart Contract challenge

         (process deposit/withdraw/challenge events)
```

1. Get new events from Plasma contract
2. Add new deposits to the plasma chain
3. Flag (censor) UTXOs withdrawed in the contract
4. Challenge invalid withdrawals

