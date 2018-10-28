pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/cryptography/ECDSA.sol";
import "openzeppelin-solidity/contracts/ownership/Ownable.sol";
import "openzeppelin-solidity/contracts/math/SafeMath.sol";


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
        address[] newBlocks,
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
        address[] newBlocks,
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
        address[] newBlocks,
        uint256 protectedBlockNumber,
        address protectedBlockHash
    )
        internal
        returns(uint256)
    {
        require(fromIndex == _blocks.length, "Invalid fromIndex");
        require(fromIndex == 0 || _blocks[protectedBlockNumber] == protectedBlockHash, "Wrong protected block number");

        uint256 begin = _blocks.length.sub(fromIndex);
        _blocks.length = fromIndex.add(newBlocks.length);
        for (uint i = begin; i < newBlocks.length; i++) {
            _blocks[fromIndex + i] = newBlocks[i];
        }

        if (begin < newBlocks.length) {
            // solium-disable-next-line security/no-block-members
            emit BlocksSubmitted(_blocks.length, block.timestamp);
        }

        return newBlocks.length - begin;
    }
}
