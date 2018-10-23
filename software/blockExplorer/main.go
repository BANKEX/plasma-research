package main

import (
	"./ethereum"
	"./rlp"
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strconv"
)

var secretKey, _ = hexutil.Decode("0x3d345c1036f325e046c8a013707def71a8854f563f7b17f7c61d81975c7de479")
var wallet, _ = ethereum.Wallet(secretKey)

var operatorSecretKey, _ = hexutil.Decode("0xa982390dadc0b9e0fb8df22236fb8211e0b2a6c1f61f0f67c268b11d1f1bae52")
var operatorWallet, _ = ethereum.Wallet(operatorSecretKey)

var someAssetId = big.NewInt(123).Bytes()
var samplePubKey = wallet.PublicKey().Bytes()

var amount1 = big.NewInt(1)
var amount2 = big.NewInt(1)
var sum = big.NewInt(0).Add(amount1, amount2)

func sampleInputs() []rlp.TransactionInput {

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

func sampleOutputs() [] rlp.TransactionOutput {
	output := rlp.TransactionOutput{};
	output.Owner = samplePubKey
	output.AssetId = someAssetId
	output.Amount = sum.Bytes()
	return []rlp.TransactionOutput{output};
}

func sampleTransaction() rlp.Transaction {

	transaction := rlp.Transaction{}
	transaction.UnsignedContent = rlp.UnsignedTransactionContent{
		Inputs:   sampleInputs(),
		Outputs:  sampleOutputs(),
		Metadata: rlp.Metadata{MaxBlockId: 7},
	}

	rlpEncoded, _ := rlp.EncodeToRLP(transaction.UnsignedContent)

	// TODO: replace with Sprintf
	var data []byte
	data = append(data, []byte("\x19Ethereum Signed Message:\n")...)
	data = append(data, []byte(strconv.Itoa(len(rlpEncoded)))...)
	data = append(data, rlpEncoded...)

	// TODO: add error handling
	var signature, _ = wallet.Sign(data)

	R := signature.R
	S := signature.S
	V := signature.V

	transaction.Signatures = []rlp.Signature{
		rlp.Signature{R: R, S: S, V: V},
	}

	return transaction
}

func sampleHeader() rlp.BlockHeader {
	header := rlp.BlockHeader{}

	header.BlockNumber = 5
	header.PreviousBlockHash = big.NewInt(0xDEADBEFF).Bytes()
	header.MerkleRoot = big.NewInt(0xCAFEBABE).Bytes()

	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, header.BlockNumber)

	var data []byte
	data = append(data, b...)
	data = append(data, header.PreviousBlockHash...)
	data = append(data, header.MerkleRoot...)
	var signature, _ = operatorWallet.Sign(data)

	R := signature.R
	S := signature.S
	V := signature.V

	header.Signature = rlp.Signature{R: R, S: S, V: V}
	return header
}

func main() {

	block := rlp.Block{}
	block.Transactions = []rlp.Transaction{
		sampleTransaction(),
	}
	block.BlockHeader = sampleHeader()

	var rlpData1, _ = rlp.EncodeToRLP(block)

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
