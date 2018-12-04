package main

import "../commons/blockchain"
import a "../commons/blockchain/alias"
import p "./pool"

func assembleBlock(utxoPool p.UtxoPool, pendingTransactions []blockchain.Transaction, privateKey []byte) (blockchain.Block, p.UtxoPool) {
	block := blockchain.Block{}

	for _, transaction := range blockchain.HandleTxs(pendingTransactions) {
		block.Transactions = append(block.Transactions, transaction)
	}

	leafs := blockchain.PrepareLeafs(block.Transactions)
	tree := blockchain.NewSumMerkleTree(leafs)

	RSAAccumulator := a.Uint2048{}
	RSAInclusionProof := blockchain.GetRSAInclusionProof(block.Transactions)

	headerContent := blockchain.UnsignedBlockHeader{
		BlockNumber:      1, // TODO: current + 1
		TransactionsRoot: tree.GetRoot(),
		RSAAccumulator:   RSAAccumulator,
		RSAChainProof:    RSAInclusionProof,
	}

	signature := SingHeader(headerContent, privateKey)

	block.Header = blockchain.BlockHeader{
		UnsignedBlockHeader: headerContent,
		Signature:           signature,
	}

	return block, utxoPool
}
