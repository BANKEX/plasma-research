package blockchain

import "log"

func isValidTx(utxoPool UtxoPool, transaction Transaction) bool {

	hasValidRanges(transaction)

	var utxoSet = make(map[UTXO]bool)

	for _, input := range transaction.Inputs {
		utxo := UTXO{input.PrevTxHash, input.OutputIndex}

		_, outputIsUnspent := utxoPool[utxo]
		_, alreadyClaimed := utxoSet[utxo]

		if !outputIsUnspent || alreadyClaimed {
			return false // (1) all transaction inputs has corresponding outputs in the pool and they aren't claimed twice by transaction
		}

		output := utxoPool[utxo]

		// TODO: Since we know the owner of output we can check specific signature - not both
		isValidSignature := verifySignature(output.Owner, transaction.UnsignedTransaction, transaction.Signatures[0])
		if !isValidSignature && len(transaction.Signatures) == 1 {
			// Only one signature and it doesn't belong to the output Owner
			return false
		} else if len(transaction.Signatures) == 2 && !verifySignature(output.Owner, transaction.UnsignedTransaction, transaction.Signatures[1]) {
			// Two signatures and both of them don't belong to the output Owner
			return false
		}

		// txInputSum += output.value;
		utxoSet[utxo] = true
	}

	// (4) all of {@code tx}s output values are non-negative
	//double txOutputSum = 0;
	//for (Transaction.Output output: tx.getOutputs()) {
	//	if( output.value < 0)
	//	return false;
	//
	//	txOutputSum += output.value;
	//}

	return true // (txInputSum >= txOutputSum);
}

// Returns consistent set of transactions
func HandleTxs(possibleTxs []Transaction) []Transaction {

	// Accepted transactions that we choose from all possible transaction that we get
	var accepted []Transaction
	var utxoSet = make(map[UTXO]bool)

txLoop:
	for _, transaction := range possibleTxs {
		if !isValidTx(transaction) {
			continue
		}

		for _, input := range transaction.Inputs {
			utxo := UTXO{input.PrevTxHash, input.OutputIndex}
			if utxoSet[utxo] {
				// Can't use this two transactions together, since they use the same output, let's go to next one
				continue txLoop
			}
		}

		for _, input := range transaction.Inputs {
			utxo := UTXO{input.PrevTxHash, input.OutputIndex}
			utxoSet[utxo] = true
		}

		accepted = append(accepted, transaction)
	}

	// TODO: in case we split some range we should take and assign new prime numbers...
	// TODO: in case we just change ownership we should keep the same prime number mapping

	// Remove UXTO that we spent from the pool
	for _, transaction := range accepted {
		for _, input := range transaction.Inputs {
			utxo := UTXO{input.PrevTxHash, input.OutputIndex}
			delete(utxoPool, utxo)
		}
	}

	// Add new UXTO that we produce to the pool
	for _, transaction := range accepted {
		for outputIdx, output := range transaction.Outputs {
			tx, err := transaction.GetHash()
			log.Fatal(err)
			utxo := UTXO{tx, uint8(outputIdx)}
			utxoPool[utxo] = output
		}
	}

	return accepted
}
