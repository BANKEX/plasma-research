pragma solidity ^0.4.24;

import "./SafeMath32.sol";


library SumMerkleProof {
  using SafeMath32 for uint32;

  // here is 32-bit plasma
  struct Slice {
    uint32 begin;
    uint32 end;
  }

  // @dev data ordered from leaves to root.
  // @dev index bits: right bit correspond leaves
  struct Proof {
    uint32 index;
    Slice slice;
    address item;
    bytes data;
  }

  function item(bytes memory proof, uint i) internal pure returns(uint32 length, address result) {
    // solium-disable-next-line security/no-inline-assembly
    assembly {

      length := div(
        mload(
          add(
            proof,
            // 12 = index + begin + end
            // 20 = item (address)
            // 32 = 12 + 20
            add(32, mul(i, 24))
          )
        ),
        // Start from data offset, shift right to 28 bytes and return 4
        0x100000000000000000000000000000000000000000000000000000000
      )

      result := div(
        mload(
          add(
            proof,
            // 12 = index + begin + end
            // 20 = item (address)
            // 36 = 12 + 20 + 4 (slice len offset)
            add(36, mul(i, 24))
          )
        ),
        // Start from data offset + 4, shift right to 12 bytes and return first 20
        0x1000000000000000000000000
      )
    }
  }

  // @dev compute hash of the node from two child nodes
  function hash(uint32 l1, uint32 l2, address a1, address a2) internal pure returns(address) {
    return address(keccak256(abi.encodePacked(l1, l2, a1, a2)));
  }

  function sumMerkleProof(Proof memory proof, address root, uint32 rootLength) internal pure returns(bool) {
    uint depth = proof.data.length / 24;
    uint32 curLength = proof.slice.end.sub(proof.slice.begin);
    address curItem = proof.item;
    uint32 curLeft = proof.slice.begin;
    uint32 index = proof.index;

    for(uint8 i = 0; i < depth; i++) {
      (uint32 length, address result) = item(proof.data, i);
      if (index & 1 == 1) {
        curItem = hash(length, curLength, result, curItem);
        curLeft = curLeft.sub(length);
        curLength = curLength.add(length);
      } else {
        curItem = hash(curLength, length, curItem, result);
        curLength = curLength.add(length);
      }
      index >>= 1;
    }
    return curLeft == 0 && curLength == rootLength && curItem == root;
  }
}
