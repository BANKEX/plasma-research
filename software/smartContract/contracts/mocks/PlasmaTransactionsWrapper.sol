pragma solidity ^0.4.24;

import { PlasmaTransactions as Tx } from "../PlasmaTransactions.sol";


contract PlasmaTransactionsWrapper {

  function validateTransaction(bytes rlpTransaction) public pure returns(bool) {
    return Tx.validateTransaction(rlpTransaction);
  }

  function validateInputOutput(bytes rlpInputOutput) public pure returns(bool) {
    return Tx.validateInputOutput(rlpInputOutput);
  }

  function verifyMerkleProof(bytes rlpMerkleProof) public pure returns(bool) {
    return Tx.verifyMerkleProof(rlpMerkleProof);
  }
}
