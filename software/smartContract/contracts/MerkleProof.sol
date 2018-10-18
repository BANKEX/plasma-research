pragma solidity ^0.4.24;

//
// Added `index` argument to this version:
// https://github.com/OpenZeppelin/openzeppelin-solidity/blob/master/contracts/cryptography/MerkleProof.sol
//


/**
 * @title MerkleProof
 * @dev Merkle proof verification based on
 * https://github.com/ameensol/merkle-tree-solidity/blob/master/src/MerkleProof.sol
 */
library MerkleProof {
    /**
    * @dev Verifies a Merkle proof proving the existence of a leaf in a Merkle tree. Assumes that each pair of leaves
    * and each pair of pre-images are sorted.
    * @param proof Merkle proof containing sibling hashes on the branch from the leaf to the root of the Merkle tree
    * @param root Merkle root
    * @param leaf Leaf of Merkle tree
    */
    function verify(
        bytes32[] proof,
        bytes32 root,
        bytes32 leaf,
        uint256 index
    )
        internal
        pure
        returns (bool)
    {
        bytes32 computedHash = leaf;

        uint256 currentIndex = index;
        for (uint256 i = 0; i < proof.length; i++) {
            bytes32 proofElement = proof[i];

            if ((currentIndex & 1) == 0) {
                // Hash(current computed hash + current element of the proof)
                computedHash = keccak256(abi.encodePacked(computedHash, proofElement));
            } else {
                // Hash(current element of the proof + current computed hash)
                computedHash = keccak256(abi.encodePacked(proofElement, computedHash));
            }

            currentIndex = currentIndex / 2;
        }

        // Check if the computed hash (root) is equal to the provided root
        return computedHash == root;
    }
}