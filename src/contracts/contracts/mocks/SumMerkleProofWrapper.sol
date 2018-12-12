pragma solidity ^0.4.24;

import "../SumMerkleProof.sol";


library SumMerkleProofWrapper {
  using SumMerkleProof for SumMerkleProof.Proof;

  function sumMerkleProofTest(
    uint32 index,
    uint32 begin,
    uint32 end,
    uint256 item,
    bytes memory data,
    uint256 root, 
    uint32 rootLength
  )
    public
    pure
    returns(bool)
  {
    SumMerkleProof.Proof memory proof = SumMerkleProof.Proof({
      index: index,
      slice: SumMerkleProof.Slice({
        begin: begin,
        end: end
      }),
      item: address(item),
      data: data
    });
    return proof.sumMerkleProof(address(root), rootLength);
  }
}