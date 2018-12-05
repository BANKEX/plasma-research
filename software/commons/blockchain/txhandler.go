package blockchain

import a "../alias"
import "../../commons/utils"

func CheckSignatures(txHash []byte, uniqueOwners []a.Uint160, signatures []a.Signature) bool {

	if len(uniqueOwners) != len(signatures) && (len(uniqueOwners) > 0 && len(uniqueOwners) <= 2) {
		return false
	}

	// Shortcut
	IsValid := utils.IsSignatureValid

	if len(uniqueOwners) == 1 {
		return IsValid(uniqueOwners[0], txHash, signatures[0])
	}

	if len(uniqueOwners) == 2 {

		hasSignatureOfTheFirstOwner :=
			IsValid(uniqueOwners[0], txHash, signatures[0]) || IsValid(uniqueOwners[0], txHash, signatures[1])

		hasSignatureOfTheSecondOwner :=
			IsValid(uniqueOwners[1], txHash, signatures[0]) || IsValid(uniqueOwners[1], txHash, signatures[1])

		return hasSignatureOfTheFirstOwner && hasSignatureOfTheSecondOwner
	}

	return true
}

func isValidTx(utxoPool UtxoPool, transaction Transaction) bool {

	// Set of UTXO claimed by transaction
	var utxoSet = make(map[UTXO]bool)

	var ownersMap = make(map[a.Uint160Bytes]bool)

	for _, input := range transaction.Inputs {
		utxo := UTXO{input.GetPrevTxHash(), input.OutputIndex}

		_, outputIsUnspent := utxoPool[utxo]

		// TODO: think a bit, may be we also can remove double spend check if we have ranges check, we also need to check that rages are non negative
		_, alreadyClaimed := utxoSet[utxo]

		if !outputIsUnspent || alreadyClaimed {
			return false // (1) all transaction inputs has corresponding outputs in the pool and they aren't claimed twice by transaction
		}

		// Owner of the output of previous transaction,
		// Memorise it to check signatures further
		owner := a.ToUint160Bytes(utxoPool[utxo].Owner)
		ownersMap[owner] = true

		// Mark utxo as claimed
		utxoSet[utxo] = true
	}

	uniqueOwners := make([]a.Uint160, 2)
	for owner, _ := range ownersMap {
		uniqueOwners = append(uniqueOwners, owner[:])
	}

	return CheckSignatures(transaction.GetHash(), uniqueOwners, transaction.Signatures)
}

// Returns consistent set of transactions
func HandleTxs(utxoPool UtxoPool, possibleTxs []Transaction) []Transaction {

	// Accepted transactions that we choose from all possible transaction that we get
	var accepted []Transaction
	var utxoSet = make(map[UTXO]bool)

txLoop:
	for _, transaction := range possibleTxs {
		if !isValidTx(utxoPool, transaction) {
			continue
		}

		for _, input := range transaction.Inputs {
			utxo := UTXO{input.GetPrevTxHash(), input.OutputIndex}
			if utxoSet[utxo] {
				// Can't use this two transactions together, since they use the same output, let's go to next one
				continue txLoop
			}
		}

		for _, input := range transaction.Inputs {
			utxo := UTXO{input.GetPrevTxHash(), input.OutputIndex}
			utxoSet[utxo] = true
		}

		accepted = append(accepted, transaction)
	}

	// TODO: in case we split some range we should take and assign new prime numbers...
	// TODO: in case we just change ownership we should keep the same prime number mapping

	// Remove UTXO that we spent from the pool
	for _, transaction := range accepted {
		for _, input := range transaction.Inputs {
			utxo := UTXO{input.GetPrevTxHash(), input.OutputIndex}
			delete(utxoPool, utxo)
		}
	}

	// Add new UTXO that we produce to the pool
	for _, transaction := range accepted {
		for outputIdx, output := range transaction.Outputs {
			tx := transaction.GetHash()
			key := a.ToTxHashBytes(tx)
			utxo := UTXO{key, uint8(outputIdx)}
			utxoPool[utxo] = output
		}
	}

	return accepted
}
