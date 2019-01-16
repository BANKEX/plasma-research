pragma solidity ^0.5.2;

import "../SumMerkleProof.sol";
import {PlasmaDecoder} from "../PlasmaDecoder.sol";

library SumMerkleProofWrapper {
  using SumMerkleProof for SumMerkleProof.Proof;
  using PlasmaDecoder for bytes;
  uint32 constant public PLASMA_ASSETS_TOTAL_SIZE = 2 ** 24 - 1;

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

  function sumMerkleProofFromBytesTest(uint256 root, bytes memory txProofBytes) public pure returns (bool)
  {
    SumMerkleProof.Proof memory txProof = txProofBytes.decodeProof();
    return txProof.sumMerkleProof(address(root), PLASMA_ASSETS_TOTAL_SIZE);
  }
}