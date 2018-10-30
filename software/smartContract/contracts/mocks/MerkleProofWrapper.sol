pragma solidity ^0.4.24;

import { MerkleProof } from "../MerkleProof.sol";


contract MerkleProofWrapper {

  function verifyAtIndex(
    bytes32[] proof,
    bytes32 root,
    bytes32 leaf,
    uint256 index
  )
    public
    pure
    returns (bool)
  {
    return MerkleProof.verifyAtIndex(proof, root, leaf, index);
  }

  function verifyAtIndex160(
    bytes proof,
    address root,
    address leaf,
    uint256 index
  )
    public
    pure
    returns (bool)
  {
    return MerkleProof.verifyAtIndex160(proof, root, leaf, index);
  }
}
