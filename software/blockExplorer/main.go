package main

import (
	"fmt"
	"math/big"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"./rlp"
)

var someAssetId = big.NewInt(123).Bytes()
var samplePubKey = hexutil.MustDecodeBig("0x12345").Bytes()
var amount1 = big.NewInt(1)
var amount2 = big.NewInt(1)
var sum = (big.NewInt(0).Add(amount1, amount2))

func sampleInputs() []rlp.TransactionInput{

	input1 := rlp.TransactionInput{}
	input1.Amount = amount1.Bytes()
	input1.BlockIndex = 0
	input1.TxIndex = 0
	input1.TxIndex = 0
	input1.AssetId = someAssetId
	input1.Owner = samplePubKey
	input1.OutputIndex = 0

	input2 := rlp.TransactionInput{}
	input2.Amount = amount2.Bytes()
	input2.BlockIndex = 0
	input2.TxIndex = 0
	input2.AssetId = samplePubKey
	input2.Owner = samplePubKey
	input2.OutputIndex = 0

	return []rlp.TransactionInput{input1, input2};
}

func sampleOutputs() [] rlp.TransactionOutput{
	output := rlp.TransactionOutput{};
	output.Owner = samplePubKey
	output.AssetId = someAssetId
	output.Amount  = sum.Bytes()
	return []rlp.TransactionOutput{output};
}

func sampleTransaction() rlp.Transaction {

	transaction := rlp.Transaction{}
	transaction.Inputs = sampleInputs()
	transaction.Outputs = sampleOutputs()
	transaction.Metadata = rlp.Metadata{5}

	R := big.NewInt(1).Bytes()
	S := big.NewInt(1).Bytes()
	V := big.NewInt(1).Bytes()
	transaction.Signatures = []rlp.Signature{
		rlp.Signature{R,S,V},
	}

	return transaction
}

func sampleHeader() rlp.BlockHeader {
	header := rlp.BlockHeader{}

	header.BlockNumber = 5
	header.PreviousBlockHash = big.NewInt(0xDEADBEFF).Bytes()

	r := big.NewInt(1).Bytes()
	s := big.NewInt(1).Bytes()
	v := big.NewInt(1).Bytes()

	header.Signature = rlp.Signature{r,s,v}
	header.MerkleRoot = big.NewInt(0xCAFEBABE).Bytes()

	return header
}

func main() {

	block := rlp.Block{}
	block.Transactions = []rlp.Transaction{
		sampleTransaction(),
	}
	block.BlockHeader = sampleHeader()


	var rlpData1, _ = rlp.EncodeBlock(block)

	fmt.Printf("src = %x \n", block)
	fmt.Printf("encoded1= %x\n", rlpData1)

	{
		// var block, err := rlp.DecodeBlock(bytes.NewReader(rlpData1))
		block, err := rlp.DecodeBlock(rlpData1)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		fmt.Printf("dst = %x \n", block)
	}
}
