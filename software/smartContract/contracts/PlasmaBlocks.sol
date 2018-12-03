pragma solidity ^0.4.24;

import { SafeMath } from "openzeppelin-solidity/contracts/math/SafeMath.sol";
import { Ownable } from "openzeppelin-solidity/contracts/ownership/Ownable.sol";
import { ECDSA } from "openzeppelin-solidity/contracts/cryptography/ECDSA.sol";


contract PlasmaBlocks is Ownable {
  using SafeMath for uint256;

  address[] private _blocks;

  event BlocksSubmitted(uint256 indexed length, uint256 time);

  function blocksLength() public view returns(uint) {
    return _blocks.length;
  }

  function blocks(uint i) public view returns(address) {
    return _blocks[i];
  }

  function submitBlocks(
    uint256 fromIndex,
    bytes newBlocks,
    uint256 protectedBlockNumber,
    address protectedBlockHash
  )
    public
    onlyOwner
    returns(uint256)
  {
    _submitBlocks(fromIndex, newBlocks, protectedBlockNumber, protectedBlockHash);
  }

  function submitBlocksSigned(
    uint256 fromIndex,
    bytes newBlocks,
    uint256 protectedBlockNumber,
    address protectedBlockHash,
    bytes rsv
  )
    public
    returns(uint256)
  {
    bytes32 messageHash = keccak256(
      abi.encodePacked(
        fromIndex,
        newBlocks,
        protectedBlockNumber,
        protectedBlockHash
      )
    );
    bytes32 signedHash = ECDSA.toEthSignedMessageHash(messageHash);
    require(owner() == ECDSA.recover(signedHash, rsv), "Invalid signature");
    return _submitBlocks(fromIndex, newBlocks, protectedBlockNumber, protectedBlockHash);
  }

  function _submitBlocks(
    uint256 fromIndex,
    bytes newBlocks,
    uint256 protectedBlockNumber,
    address protectedBlockHash
  )
    internal
    returns(uint256)
  {
    uint256 newBlocksLength = newBlocks.length / 20;

    require(fromIndex == _blocks.length, "Invalid fromIndex");
    require(fromIndex == 0 || _blocks[protectedBlockNumber] == protectedBlockHash, "Wrong protected block number");

    uint256 begin = _blocks.length.sub(fromIndex);
    _blocks.length = fromIndex.add(newBlocksLength);
    for (uint i = begin; i < newBlocksLength; i++) {
      address newBlock;
      uint256 offset = 32 + i * 20;
      // solium-disable-next-line security/no-inline-assembly
      assembly {
        // Load the current element of the proof
        newBlock := div(mload(add(newBlocks, offset)), 0x1000000000000000000000000)
      }

      _blocks[fromIndex + i] = newBlock;
    }

    if (begin < newBlocksLength) {
      // solium-disable-next-line security/no-block-members
      emit BlocksSubmitted(_blocks.length, block.timestamp);
    }

    return newBlocksLength - begin;
  }
}
