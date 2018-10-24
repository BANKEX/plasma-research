pragma solidity ^0.4.24;

import { MerkleProof } from "../MerkleProof.sol";


contract MerkleProofWrapper {

    function verify(
        bytes32[] proof,
        bytes32 root,
        bytes32 leaf,
        uint256 index
    )
        public
        pure
        returns (bool)
    {
        return MerkleProof.verify(proof, root, leaf, index);
    }
}
