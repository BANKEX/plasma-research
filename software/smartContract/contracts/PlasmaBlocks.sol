pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/ECRecovery.sol";
import "openzeppelin-solidity/contracts/ownership/Ownable.sol";
import "openzeppelin-solidity/contracts/math/SafeMath.sol";


contract PlasmaBlocks is Ownable {
    using SafeMath for uint256;

    uint256[] private _blocks;

    event BlocksSubmitted(uint256 indexed length);

    function blocksLength() public view returns(uint) {
        return _blocks.length;
    }

    function blocks(uint i) public view returns(uint256) {
        return _blocks[i];
    }

    function allBlocks() public view returns(uint256[]) {
        return _blocks;
    }

    function submitBlocks(uint256 fromIndex, uint256[] newBlocks) public onlyOwner returns(uint) {
        return _submitBlocks(fromIndex, newBlocks);
    }

    function submitBlocksSigned(uint256 fromIndex, uint256[] newBlocks, bytes32 r, bytes32 s, uint8 v) public returns(uint) {
        bytes32 messageHash = keccak256(abi.encodePacked(fromIndex, newBlocks));
        bytes32 signedHash = ECRecovery.toEthSignedMessageHash(messageHash);
        require(owner == ecrecover(signedHash, v < 27 ? v + 27 : v, r, s), "Invalid signature");
        return _submitBlocks(fromIndex, newBlocks);
    }

    function _submitBlocks(uint256 fromIndex, uint256[] newBlocks) internal returns(uint) {
        uint256 begin = _blocks.length.sub(fromIndex);
        uint256 end = newBlocks.length.sub(begin);
        _blocks.length = fromIndex.add(newBlocks.length);
        for (uint i = begin; i < end; i++) {
            _blocks[fromIndex + i] = newBlocks[i];
        }

        if (begin < end) {
            emit BlocksSubmitted(_blocks.length);
        }
    }
}
