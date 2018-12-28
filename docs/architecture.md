# Plasma Cashflow data model

## Amounts, multipliers, differences from Prime and more


We use 250bit hashes (lesser bits of $Z_p$ numbers) and Pedersen hash function and 128bit for values.

Transaction hash algorithm and structure is simpliferd due zkSNARK optimization.

All signatures are made on baby jubjub curve with pedersen hashes. So, we need to map baby jubjub addresses to ethereum addresses.

Address is base58 encoded uint256 number


## Solidity


### Merkle trees and proofs

### General data types

We use this types as strucures or RLP structures and compute hash of them as is (without merkelization or something like this).
#### TXInput

Here is input to a transaction. You may consider it as pointer to any output in the blockchain. Note: amount may be lesser than the output. It is valid, if one output has many inputs with different parts of the segment.

```solidity
struct TXInput
{
    bool isNull,
    uint256 owner, 
    uint64 blockIndex, 
    uint32 txIndex,
    uint256 txContentHash,
    uint8 outputIndex, 
    Segment amount
}
```
#### TXOutput

You may consider it as owned segment.

```solidity
struct TXOutput
{
    bool IsNull,
    uint256 owner, 
    Segment amount
}
```

#### Segment

Segment with begin and end.

```solidity
struct Segment
{
    uint128 begin,
    uint128 end
}
```

#### Signature

```solidity
struct Signature {
    bool isNull,
    uint8 v,
    uint256 r,
    uint256 s
}
```



### Complex data types





#### Transaction

Transaction may be passed into functions as following object:

``` solidity
struct Transaction {
    Input[2] inputs,
    Output[2] outputs,
    uint64 maxBlockIndex,
    Signature[2] signatures
}
```

Some of inputs, outputs or signatures may be NULL. 

##### Transaction hash computation

Arguments are limited by 2736030358979909402780800718157159386076813972158567259200215660948447373041 (it is about 250 bits)

```PedersenHash(x1,x2,x3) = PedersenHash(PedersenHash(x1, x2), x3)```

We linearize TransactionContent and compute the hash.
It is enough to store only TransactionContentHashes at leaves of SumMerkleTree, because signatures are cryptographically bounded to inputs of the transactions. If the operator do not provide signatures, blocks are considered to be withheld.





### Plasma state

#### Plasma chain


```solidity
hashmap(uint64 => uint240) sumMerkleRoot;
```

First 32 bits are used for transactional blocks. Second 32 bits are used for deposit or withdrawal blocks (mined on the contract).

It is enough for 8000 years of plasma lifetime (if we mine transactional block for each 5 minutes) and 9 years to exhaust exitability if the operator stops blocks producing.

Deposit blocks creates assets from nothing and withdrawal blocks burn assets.

We do not publish slice here, because it must corresponds total plasma space.



#### Exit state

```solidity
hashmap(uint256 => bool) exitStateHashmap;
```

The list of unchallanged and not finalized exits


#### Deposit state

```solidity
OrderedLinkedList deposit;
```

Here is list of deposited segments.



### General methods

```solidity
function deposit(OrderedLinkedListItem depositSlot) external payable returns(bool);
```

Transaction may be banned and transaction may be increased priority for txContentHash 

### Priority increasing game

#### PriorityState
``` Solidity
struct PriorityState {
    TransactionContent txContent, //unsigned transaction
    uint256 txContentHash, 
    Signature s, // at least one signature
    uint256 timestamp
}

//storage

mapping (uint256 => bool) priorityChallenges // mapping keccak256(priorityState) => bool of all active priority challenges

mapping (uint256 => uint64) txFix // priority fix for transacion. 2^64-1 for banned transactions, zero is default value

```



```Solidity
function priorityBegin(
    TransactionContent txContent,
    uint256txContentHash,
    Signature s
)

struct priorityChallengeHash(
    PriorityState state,
    Groth16Proof snarkProof // proof that hash is invalid
)

// spend of one input of state.txContent
// txContentHash, txBlockIndex is information abot spending tx
struct priorityChallengeSpend(
    PriorityState state,
    uint256[3] txContentHash, //ContentHash of 2 inputs and spending tx
    uint64[3] txBlockIndex, // BlockIndex of 2 inputs and spending tx
    Groth16Proof snarkProof // proof of inclusion tx into , spending part of state.point
) external returns (bool);


// accept signature differs from state.s
function prioritySignatureCollect(
    PriorityState state,
    Signature s
)

```

### Exit game

#### ExitState

Here is exit state. `TXInput` is not an input of any included transaction. You may consider it as pointer to any output or input of withdrawal transaction.

Hashing algorithm is standard (keccak256 of blob). It is enough

```
struct ExitState {
    TXInput point,
    uint256 timestamp
}
```

```Solidity
function withdrawalBegin(
    Input point 
) external payable returns (bool);

function withdrawalChallangeSpend(
    ExitState state, 
    uint256[3] txContentHash,
    uint64[3] txBlockIndex, 
    Groth16Proof snarkProof // proof of inclusion tx into , spending part of state.point
) external returns (bool);

function withdrawalChallangeExistance(
    ExitState state,
    Groth16Proof snarkProof // proof of exclusion state.point from state.point.blockIndex
) external returns (bool);

function withdrawalChallangeHash(
    ExitState state,
    Groth16Proof snarkProof // proof that txContextHash is wrong
) external returns (bool);

function withdrawalChallangeConcurrent(

) returns (bool);


// finalize. If transaction is banned by txContentHash, reject exit procedure
function withdrawalEnd(ExitState state)
```


## SNARKs

### zkSNARK proof

```solidity
struct Groth16Proof
{
    uint256[2] A;
    uint256[2][2] B;
    uint256[2] C;
}
```
