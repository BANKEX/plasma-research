# API

root: `/api/v1/`

---

`POST /send`

`raw RLP tx bytes`

Submit a transaction to plasma chain

results:

`200 OK` - tx receved and sucessfully written to DB and queued for the next block

`400` - tx is malformed

`403` - tx signature is invalid

`403` - absent or insufficient fee

`403` - sum of inputs is greater than sum of outputs

`403` - invalid tx input

---

`GET /blocks/<id>`

Download a block

result: raw block bytes or `404`

---

`GET /details`

Get plasma smart contract details

result:

```json
{
    address: "0x1233123",
    abi: [
    	...
    ]
}
```

---

`GET /status`

Get current system status

result:

```json
{
    lastBlock: 42
    lastEthereumBlock: 1337
    ...
}
```









