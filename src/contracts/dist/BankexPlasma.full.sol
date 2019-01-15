pragma solidity ^0.5.0;

// File: openzeppelin-solidity/contracts/math/SafeMath.sol

/**
 * @title SafeMath
 * @dev Unsigned math operations with safety checks that revert on error
 */
library SafeMath {
    /**
    * @dev Multiplies two unsigned integers, reverts on overflow.
    */
    function mul(uint256 a, uint256 b) internal pure returns (uint256) {
        // Gas optimization: this is cheaper than requiring 'a' not being zero, but the
        // benefit is lost if 'b' is also tested.
        // See: https://github.com/OpenZeppelin/openzeppelin-solidity/pull/522
        if (a == 0) {
            return 0;
        }

        uint256 c = a * b;
        require(c / a == b);

        return c;
    }

    /**
    * @dev Integer division of two unsigned integers truncating the quotient, reverts on division by zero.
    */
    function div(uint256 a, uint256 b) internal pure returns (uint256) {
        // Solidity only automatically asserts when dividing by 0
        require(b > 0);
        uint256 c = a / b;
        // assert(a == b * c + a % b); // There is no case in which this doesn't hold

        return c;
    }

    /**
    * @dev Subtracts two unsigned integers, reverts on overflow (i.e. if subtrahend is greater than minuend).
    */
    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b <= a);
        uint256 c = a - b;

        return c;
    }

    /**
    * @dev Adds two unsigned integers, reverts on overflow.
    */
    function add(uint256 a, uint256 b) internal pure returns (uint256) {
        uint256 c = a + b;
        require(c >= a);

        return c;
    }

    /**
    * @dev Divides two unsigned integers and returns the remainder (unsigned integer modulo),
    * reverts when dividing by zero.
    */
    function mod(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b != 0);
        return a % b;
    }
}

// File: openzeppelin-solidity/contracts/ownership/Ownable.sol

/**
 * @title Ownable
 * @dev The Ownable contract has an owner address, and provides basic authorization control
 * functions, this simplifies the implementation of "user permissions".
 */
contract Ownable {
    address private _owner;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    /**
     * @dev The Ownable constructor sets the original `owner` of the contract to the sender
     * account.
     */
    constructor () internal {
        _owner = msg.sender;
        emit OwnershipTransferred(address(0), _owner);
    }

    /**
     * @return the address of the owner.
     */
    function owner() public view returns (address) {
        return _owner;
    }

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyOwner() {
        require(isOwner());
        _;
    }

    /**
     * @return true if `msg.sender` is the owner of the contract.
     */
    function isOwner() public view returns (bool) {
        return msg.sender == _owner;
    }

    /**
     * @dev Allows the current owner to relinquish control of the contract.
     * @notice Renouncing to ownership will leave the contract without an owner.
     * It will not be possible to call the functions with the `onlyOwner`
     * modifier anymore.
     */
    function renounceOwnership() public onlyOwner {
        emit OwnershipTransferred(_owner, address(0));
        _owner = address(0);
    }

    /**
     * @dev Allows the current owner to transfer control of the contract to a newOwner.
     * @param newOwner The address to transfer ownership to.
     */
    function transferOwnership(address newOwner) public onlyOwner {
        _transferOwnership(newOwner);
    }

    /**
     * @dev Transfers control of the contract to a newOwner.
     * @param newOwner The address to transfer ownership to.
     */
    function _transferOwnership(address newOwner) internal {
        require(newOwner != address(0));
        emit OwnershipTransferred(_owner, newOwner);
        _owner = newOwner;
    }
}

// File: openzeppelin-solidity/contracts/token/ERC20/IERC20.sol

/**
 * @title ERC20 interface
 * @dev see https://github.com/ethereum/EIPs/issues/20
 */
interface IERC20 {
    function transfer(address to, uint256 value) external returns (bool);

    function approve(address spender, uint256 value) external returns (bool);

    function transferFrom(address from, address to, uint256 value) external returns (bool);

    function totalSupply() external view returns (uint256);

    function balanceOf(address who) external view returns (uint256);

    function allowance(address owner, address spender) external view returns (uint256);

    event Transfer(address indexed from, address indexed to, uint256 value);

    event Approval(address indexed owner, address indexed spender, uint256 value);
}

// File: openzeppelin-solidity/contracts/token/ERC20/SafeERC20.sol

/**
 * @title SafeERC20
 * @dev Wrappers around ERC20 operations that throw on failure.
 * To use this library you can add a `using SafeERC20 for ERC20;` statement to your contract,
 * which allows you to call the safe operations as `token.safeTransfer(...)`, etc.
 */
library SafeERC20 {
    using SafeMath for uint256;

    function safeTransfer(IERC20 token, address to, uint256 value) internal {
        require(token.transfer(to, value));
    }

    function safeTransferFrom(IERC20 token, address from, address to, uint256 value) internal {
        require(token.transferFrom(from, to, value));
    }

    function safeApprove(IERC20 token, address spender, uint256 value) internal {
        // safeApprove should only be called when setting an initial allowance,
        // or when resetting it to zero. To increase and decrease it, use
        // 'safeIncreaseAllowance' and 'safeDecreaseAllowance'
        require((value == 0) || (token.allowance(msg.sender, spender) == 0));
        require(token.approve(spender, value));
    }

    function safeIncreaseAllowance(IERC20 token, address spender, uint256 value) internal {
        uint256 newAllowance = token.allowance(address(this), spender).add(value);
        require(token.approve(spender, newAllowance));
    }

    function safeDecreaseAllowance(IERC20 token, address spender, uint256 value) internal {
        uint256 newAllowance = token.allowance(address(this), spender).sub(value);
        require(token.approve(spender, newAllowance));
    }
}

// File: openzeppelin-solidity/contracts/introspection/IERC165.sol

/**
 * @title IERC165
 * @dev https://github.com/ethereum/EIPs/blob/master/EIPS/eip-165.md
 */
interface IERC165 {
    /**
     * @notice Query if a contract implements an interface
     * @param interfaceId The interface identifier, as specified in ERC-165
     * @dev Interface identification is specified in ERC-165. This function
     * uses less than 30,000 gas.
     */
    function supportsInterface(bytes4 interfaceId) external view returns (bool);
}

// File: openzeppelin-solidity/contracts/token/ERC721/IERC721.sol

/**
 * @title ERC721 Non-Fungible Token Standard basic interface
 * @dev see https://github.com/ethereum/EIPs/blob/master/EIPS/eip-721.md
 */
contract IERC721 is IERC165 {
    event Transfer(address indexed from, address indexed to, uint256 indexed tokenId);
    event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId);
    event ApprovalForAll(address indexed owner, address indexed operator, bool approved);

    function balanceOf(address owner) public view returns (uint256 balance);
    function ownerOf(uint256 tokenId) public view returns (address owner);

    function approve(address to, uint256 tokenId) public;
    function getApproved(uint256 tokenId) public view returns (address operator);

    function setApprovalForAll(address operator, bool _approved) public;
    function isApprovedForAll(address owner, address operator) public view returns (bool);

    function transferFrom(address from, address to, uint256 tokenId) public;
    function safeTransferFrom(address from, address to, uint256 tokenId) public;

    function safeTransferFrom(address from, address to, uint256 tokenId, bytes memory data) public;
}

// File: contracts/OrderedIntervalList.sol

/**
 * @title OrderedIntervalList
 * @dev List of ordered by intervals with non intersections checks.
*/
library OrderedIntervalList {
  struct Interval {
    uint64 begin; // inclusive
    uint64 end;   // exclusive

    uint64 next;
    uint64 prev;
  }

  struct Data {
    Interval[] intervals; // sparsed array
    uint64 firstIndex;
    uint64 lastIndex;
  }

  function isInitialized(Data storage self) internal view returns(bool) {
    return self.intervals.length > 0;
  }

  function getFirstIndex(Data storage self) internal view returns(uint64) {
    return self.firstIndex;
  }

  function getLastIndex(Data storage self) internal view returns(uint64) {
    return self.lastIndex;
  }

  /**
   * @notice Check if OrderedIntervalList initialized
   * @return was initialized or not
   */
  function initialize(Data storage self) internal {
    require(self.intervals.length == 0, "OrderedIntervalList was already initialized");
    self.intervals.push(Interval(0,0,0,0));
  }

  /**
   * @notice Get interval by the index
   * @param id interval index in the list
   * @return interval tuple
   */
  function get(Data storage self, uint64 id) internal view returns(Interval storage interval) {
    require(id < self.intervals.length, "interval id doesn't exists in interval set");
    interval = self.intervals[id];
    //require(interval.end != 0, "interval id doesn't exsits in interval set");
  }

  /**
   * @notice Check interval existance by the index
   * @param id interval index in the list
   * @return is existing or not
   */
  function exist(Data storage self, uint64 id) internal view returns (bool) {
    return self.intervals[id].end != 0;
  }

  /**
   * @notice Add interval after the lates interval
   * @param size length of the new interval
   * @return id of the latest interval
   */
  function append(
    Data storage self,
    uint64 size
  )
    internal
    returns(
      uint64 id,
      uint64 begin,
      uint64 end
    )
  {
    Interval storage lastInterval = self.intervals[self.lastIndex];
    begin = lastInterval.end;
    end = lastInterval.end + size - 1;
    id = insert(self, self.lastIndex, 0, begin, end);
  }

  /**
   * @notice Insert interval in the specific place in a list
   * @dev Method also check that new interval doesn't intersect with existed intervals in list
   * @param prev id of the prev interval in the list. Zero if it's a first interval.
   * @param next id of the next interval in the list. Zero if it's a last interval.
   * @param begin left bound of the new interval (inclusive)
   * @param end right bound of the new interval (exclusive)
   * @return id of the interval that contain new interval. Could be a new interval or an existed with
   * extended bounds in case of adjacent bounds of the inserted interval with his neighbors.
   */
  function insert(
    Data storage self,
    uint64 prev,
    uint64 next,
    uint64 begin,
    uint64 end
  )
    internal
    returns(uint64 id)
  {
    return _insert(
      self,
      prev,
      next,
      begin,
      end,
      false
    );
  }

  /**
   * @notice Remove range in interval by index
   * @param index interval index in list
   * @param begin left range bound
   * @param end right range bound
   * @return index of the new interval if new one was created (was made a hole insided existed interval) or zero.
   */
  function remove(
    Data storage self,
    uint64 index,
    uint64 begin,
    uint64 end
  )
    internal
    returns(uint64 newInterval)
  {
    require(begin < end, "right bound less than left bound");
    require(index < self.intervals.length, "valid index bounds");

    Interval storage modifiedInterval = self.intervals[index];
    Interval storage prevInterval = self.intervals[modifiedInterval.prev];
    Interval storage nextInterval = self.intervals[modifiedInterval.next];
    require(modifiedInterval.end != 0, "removed interval doesn't exists");
    require(modifiedInterval.begin <= begin && end <= modifiedInterval.end, "incorrect removed range bounds");

    bool shrinkBegin = (begin == modifiedInterval.begin);
    bool shrinkEnd = (end == modifiedInterval.end);

    if (shrinkBegin && shrinkEnd) {
      // Remove whole interval

      if (modifiedInterval.prev > 0) {
        prevInterval.next = modifiedInterval.next;
      } else {
        self.firstIndex = modifiedInterval.next;
      }

      if (modifiedInterval.next > 0) {
        nextInterval.prev = modifiedInterval.prev;
      } else {
        self.lastIndex = modifiedInterval.prev;
      }

      delete self.intervals[index];
    } else
    if (shrinkBegin) {
      // Shrink from left side
      modifiedInterval.begin = end;
    } else
    if (shrinkEnd) {
      // Shrink from right side
      modifiedInterval.end = begin;
    } else {
      // Make a hole
      uint64 oldEnd = modifiedInterval.end;
      modifiedInterval.end = begin;
      modifiedInterval.next = _insert(
        self,
        index,
        modifiedInterval.next,
        end,
        oldEnd,
        true
      );
      newInterval = modifiedInterval.next;
    }
  }

  function _insert(
    Data storage self,
    uint64 prev,
    uint64 next,
    uint64 begin,
    uint64 end,
    bool allowGapAfterLast
  )
    private
    returns(uint64 id)
  {
    require(begin < end, "right bound less or equal to left bound");
    require((prev != 0 || next != 0) == (self.firstIndex > 0), "prev and next could be zero iff no intervals");

    if (!isInitialized(self)) {
      initialize(self);
    }

    Interval storage prevInterval = self.intervals[prev];
    Interval storage nextInterval = self.intervals[next];

    require(prev == 0 || begin >= prevInterval.end, "begin could not intersect prev interval");
    require(next == 0 || end <= nextInterval.begin, "end could not intersect next interval");

    if ((prev > 0) == (next > 0)) {
      // Adding between existing intervals or very first interval
      require(
        prevInterval.next == next && nextInterval.prev == prev,
        "prev and next should refer to the neighboring intervals"
      );
    } else
    if (next > 0) {
      // Adding before first interval
      require(
        self.firstIndex == next && nextInterval.prev == 0,
        "next should refer to the first interval"
      );
    } else
    if (prev > 0) {
      // Adding after last interval
      require(
        self.lastIndex == prev && prevInterval.next == 0,
        "prev should refer to the last interval"
      );
      require(
        allowGapAfterLast || prev != self.lastIndex || prevInterval.end == begin, 
        "should begin from the end of latest interval when adding to the end"
      );
    }

    bool concatPrev = (prev > 0 && begin == prevInterval.end);
    bool concatNext = (next > 0 && end == nextInterval.begin);

    if (!concatPrev && !concatNext) {
      id = uint64(self.intervals.length);
      self.intervals.push(Interval({
        begin: begin,
        end: end,
        prev: prev,
        next: next
      }));

      if (next > 0) {
        nextInterval.prev = id;
      } else {
        self.lastIndex = id;
      }

      if (prev > 0) {
        prevInterval.next = id;
      } else {
        self.firstIndex = id;
      }
    } else
    if (concatPrev && concatNext) {
      prevInterval.end = nextInterval.end;
      prevInterval.next = nextInterval.next;
      id = prev;

      // When attaching pre last to last
      if (next == self.lastIndex) {
        self.lastIndex = id;
      } else {
        self.intervals[nextInterval.next].prev = id;
      }

      delete self.intervals[next];
    } else
    if (concatPrev) {
      prevInterval.end = end;
      id = prev;
    } else
    if (concatNext) {
      nextInterval.begin = begin;
      id = next;
    }
  }
}

// File: contracts/SafeMath32.sol

/**
 * @title SafeMath
 * @dev Math operations with safety checks that revert on error
 */
library SafeMath32 {

  /**
  * @dev Multiplies two numbers, reverts on overflow.
  */
  function mul(uint32 a, uint32 b) internal pure returns (uint32) {
    // Gas optimization: this is cheaper than requiring 'a' not being zero, but the
    // benefit is lost if 'b' is also tested.
    // See: https://github.com/OpenZeppelin/openzeppelin-solidity/pull/522
    if (a == 0) {
      return 0;
    }

    uint32 c = a * b;
    require(c / a == b);

    return c;
  }

  /**
  * @dev Integer division of two numbers truncating the quotient, reverts on division by zero.
  */
  function div(uint32 a, uint32 b) internal pure returns (uint32) {
    require(b > 0); // Solidity only automatically asserts when dividing by 0
    uint32 c = a / b;
    // assert(a == b * c + a % b); // There is no case in which this doesn't hold

    return c;
  }

  /**
  * @dev Subtracts two numbers, reverts on overflow (i.e. if subtrahend is greater than minuend).
  */
  function sub(uint32 a, uint32 b) internal pure returns (uint32) {
    require(b <= a);
    uint32 c = a - b;

    return c;
  }

  /**
  * @dev Adds two numbers, reverts on overflow.
  */
  function add(uint32 a, uint32 b) internal pure returns (uint32) {
    uint32 c = a + b;
    require(c >= a);

    return c;
  }

  /**
  * @dev Divides two numbers and returns the remainder (unsigned integer modulo),
  * reverts when dividing by zero.
  */
  function mod(uint32 a, uint32 b) internal pure returns (uint32) {
    require(b != 0);
    return a % b;
  }
}

// File: contracts/SumMerkleProof.sol

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
    return address(uint256(keccak256(abi.encodePacked(l1, l2, a1, a2))));
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

// File: solidity-rlp/contracts/RLPReader.sol

/*
* @author Hamdi Allam hamdi.allam97@gmail.com
* Please reach out with any questions or concerns
*/
pragma solidity ^0.5.0;

library RLPReader {
    uint8 constant STRING_SHORT_START = 0x80;
    uint8 constant STRING_LONG_START  = 0xb8;
    uint8 constant LIST_SHORT_START   = 0xc0;
    uint8 constant LIST_LONG_START    = 0xf8;

    uint8 constant WORD_SIZE = 32;

    struct RLPItem {
        uint len;
        uint memPtr;
    }

    /*
    * @param item RLP encoded bytes
    */
    function toRlpItem(bytes memory item) internal pure returns (RLPItem memory) {
        uint memPtr;
        assembly {
            memPtr := add(item, 0x20)
        }

        return RLPItem(item.length, memPtr);
    }

    /*
    * @param item RLP encoded bytes
    */
    function size(RLPItem memory item) internal pure returns (uint) {
        return item.len;
    }

    /*
    * @param item RLP encoded list in bytes
    */
    function toList(RLPItem memory item) internal pure returns (RLPItem[] memory result) {
        require(isList(item));

        uint items = numItems(item);
        result = new RLPItem[](items);

        uint memPtr = item.memPtr + _payloadOffset(item.memPtr);
        uint dataLen;
        for (uint i = 0; i < items; i++) {
            dataLen = _itemLength(memPtr);
            result[i] = RLPItem(dataLen, memPtr); 
            memPtr = memPtr + dataLen;
        }
    }

    // @return indicator whether encoded payload is a list. negate this function call for isData.
    function isList(RLPItem memory item) internal pure returns (bool) {
        if (item.len == 0) return false;

        uint8 byte0;
        uint memPtr = item.memPtr;
        assembly {
            byte0 := byte(0, mload(memPtr))
        }

        if (byte0 < LIST_SHORT_START)
            return false;
        return true;
    }

    /** RLPItem conversions into data types **/

    // @returns raw rlp encoding in bytes
    function toRlpBytes(RLPItem memory item) internal pure returns (bytes memory) {
        bytes memory result = new bytes(item.len);
        if (result.length == 0) return result;
        
        uint ptr;
        assembly {
            ptr := add(0x20, result)
        }

        copy(item.memPtr, ptr, item.len);
        return result;
    }

    // any non-zero byte is considered true
    function toBoolean(RLPItem memory item) internal pure returns (bool) {
        require(item.len == 1);
        uint result;
        uint memPtr = item.memPtr;
        assembly {
            result := byte(0, mload(memPtr))
        }

        return result == 0 ? false : true;
    }

    function toAddress(RLPItem memory item) internal pure returns (address) {
        // 1 byte for the length prefix according to RLP spec
        require(item.len <= 21);

        return address(toUint(item));
    }

    function toUint(RLPItem memory item) internal pure returns (uint) {
        require(item.len > 0);

        uint offset = _payloadOffset(item.memPtr);
        uint len = item.len - offset;
        uint memPtr = item.memPtr + offset;

        uint result;
        assembly {
            result := div(mload(memPtr), exp(256, sub(32, len))) // shift to the correct location
        }

        return result;
    }

    function toBytes(RLPItem memory item) internal pure returns (bytes memory) {
        require(item.len > 0);

        uint offset = _payloadOffset(item.memPtr);
        uint len = item.len - offset; // data length
        bytes memory result = new bytes(len);

        uint destPtr;
        assembly {
            destPtr := add(0x20, result)
        }

        copy(item.memPtr + offset, destPtr, len);
        return result;
    }

    /*
    * Private Helpers
    */

    // @return number of payload items inside an encoded list.
    function numItems(RLPItem memory item) private pure returns (uint) {
        if (item.len == 0) return 0;

        uint count = 0;
        uint currPtr = item.memPtr + _payloadOffset(item.memPtr);
        uint endPtr = item.memPtr + item.len;
        while (currPtr < endPtr) {
           currPtr = currPtr + _itemLength(currPtr); // skip over an item
           count++;
        }

        return count;
    }

    // @return entire rlp item byte length
    function _itemLength(uint memPtr) private pure returns (uint len) {
        uint byte0;
        assembly {
            byte0 := byte(0, mload(memPtr))
        }

        if (byte0 < STRING_SHORT_START)
            return 1;
        
        else if (byte0 < STRING_LONG_START)
            return byte0 - STRING_SHORT_START + 1;

        else if (byte0 < LIST_SHORT_START) {
            assembly {
                let byteLen := sub(byte0, 0xb7) // # of bytes the actual length is
                memPtr := add(memPtr, 1) // skip over the first byte
                
                /* 32 byte word size */
                let dataLen := div(mload(memPtr), exp(256, sub(32, byteLen))) // right shifting to get the len
                len := add(dataLen, add(byteLen, 1))
            }
        }

        else if (byte0 < LIST_LONG_START) {
            return byte0 - LIST_SHORT_START + 1;
        } 

        else {
            assembly {
                let byteLen := sub(byte0, 0xf7)
                memPtr := add(memPtr, 1)

                let dataLen := div(mload(memPtr), exp(256, sub(32, byteLen))) // right shifting to the correct length
                len := add(dataLen, add(byteLen, 1))
            }
        }
    }

    // @return number of bytes until the data
    function _payloadOffset(uint memPtr) private pure returns (uint) {
        uint byte0;
        assembly {
            byte0 := byte(0, mload(memPtr))
        }

        if (byte0 < STRING_SHORT_START) 
            return 0;
        else if (byte0 < STRING_LONG_START || (byte0 >= LIST_SHORT_START && byte0 < LIST_LONG_START))
            return 1;
        else if (byte0 < LIST_SHORT_START)  // being explicit
            return byte0 - (STRING_LONG_START - 1) + 1;
        else
            return byte0 - (LIST_LONG_START - 1) + 1;
    }

    /*
    * @param src Pointer to source
    * @param dest Pointer to destination
    * @param len Amount of memory to copy from the source
    */
    function copy(uint src, uint dest, uint len) private pure {
        if (len == 0) return;

        // copy as many word sizes as possible
        for (; len >= WORD_SIZE; len -= WORD_SIZE) {
            assembly {
                mstore(dest, mload(src))
            }

            src += WORD_SIZE;
            dest += WORD_SIZE;
        }

        // left over bytes. Mask is used to remove unwanted bytes from the word
        uint mask = 256 ** (WORD_SIZE - len) - 1;
        assembly {
            let srcpart := and(mload(src), not(mask)) // zero out src
            let destpart := and(mload(dest), mask) // retrieve the bytes
            mstore(dest, or(destpart, srcpart))
        }
    }
}

// File: contracts/PlasmaDecoder.sol

library PlasmaDecoder {
  using RLPReader for RLPReader.RLPItem;
  using RLPReader for bytes;

  struct Input {
    address payable owner;
    uint32 blockIndex;
    uint32 txIndex;
    uint8 outputIndex;
    address assetId;
    uint64 begin;
    uint64 end;
  }

  struct Output {
    address payable owner;
    address assetId;
    uint64 begin;
    uint64 end;
  }

  struct Metadata {
    uint32 maxBlockId;
  }

  struct Signature {
    uint256 r;
    uint256 s;
    uint8 v;
  }

  struct Transaction {
    Input[] inputs;
    Output[] outputs;
    Metadata metadata;
    Signature[] signatures;
  }

  struct Block {
    uint32 blockNumber;
    uint256 previousBlockHash;
    uint256 merkleRoot;
    Signature signature;
    Transaction[] transactions;
  }

  function decodeProof(bytes memory rlpBytes) internal pure returns(SumMerkleProof.Proof memory) {
    return _decodeProof(rlpBytes.toRlpItem().toList());
  }

  function decodeInput(bytes memory rlpBytes) internal pure returns(Input memory) {
    return _decodeInput(rlpBytes.toRlpItem().toList());
  }

  function decodeOutput(bytes memory rlpBytes) internal pure returns(Output memory) {
    return _decodeOutput(rlpBytes.toRlpItem().toList());
  }

  function decodeMetadata(bytes memory rlpBytes) internal pure returns(Metadata memory) {
    return _decodeMetadata(rlpBytes.toRlpItem().toList());
  }

  function decodeSignature(bytes memory rlpBytes) internal pure returns(Signature memory) {
    return _decodeSignature(rlpBytes.toRlpItem().toList());
  }

  function decodeTransaction(bytes memory rlpBytes) internal pure returns(Transaction memory) {
    return _decodeTransaction(rlpBytes.toRlpItem().toList());
  }

  function decodeBlock(bytes memory rlpBytes) internal pure returns(Block memory) {
    return _decodeBlock(rlpBytes.toRlpItem().toList());
  }

  // Private methods

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

  function _decodeSlice(RLPReader.RLPItem[] memory items) private pure returns(SumMerkleProof.Slice memory) {
    return SumMerkleProof.Slice({
      begin: uint32(items[0].toUint()),
      end: uint32(items[1].toUint())
    });
  }

  function _decodeProof(RLPReader.RLPItem[] memory items) private pure returns(SumMerkleProof.Proof memory) {
    return SumMerkleProof.Proof({
      index: uint32(items[0].toUint()),
      slice: _decodeSlice(items[1].toList()),
      item: items[2].toAddress(),
      data: items[3].toBytes()
    });
  }

  function _decodeInput(RLPReader.RLPItem[] memory items) private pure returns(Input memory) {
    return Input({
      owner: address(uint160(items[0].toAddress())),
      blockIndex: uint32(items[1].toUint()),
      txIndex: uint32(items[2].toUint()),
      outputIndex: uint8(items[3].toUint()),
      assetId: items[4].toAddress(),
      begin: uint64(items[5].toUint()),
      end: uint64(items[6].toUint())
    });
  }

  function _decodeInputs(RLPReader.RLPItem[] memory items) private pure returns(Input[] memory) {
    Input[] memory inputs = new Input[](items.length);
    for (uint i = 0; i < items.length; i++) {
      inputs[i] = _decodeInput(items[i].toList());
    }
    return inputs;
  }

  function _decodeOutput(RLPReader.RLPItem[] memory items) private pure returns(Output memory) {
    return Output({
      owner: address(uint160(items[0].toAddress())),
      assetId: items[1].toAddress(),
      begin: uint64(items[2].toUint()),
      end: uint64(items[3].toUint())
    });
  }

  function _decodeOutputs(RLPReader.RLPItem[] memory items) private pure returns(Output[] memory) {
    Output[] memory outputs = new Output[](items.length);
    for (uint i = 0; i < items.length; i++) {
      outputs[i] = _decodeOutput(items[i].toList());
    }
    return outputs;
  }

  function _decodeMetadata(RLPReader.RLPItem[] memory items) private pure returns(Metadata memory) {
    return Metadata({
      maxBlockId: uint32(items[0].toUint())
    });
  }

  function _decodeSignature(RLPReader.RLPItem[] memory items) internal pure returns(Signature memory) {
    return Signature({
      r: items[0].toUint(),
      s: items[0].toUint(),
      v: uint8(items[0].toUint())
    });
  }

  function _decodeSignatures(RLPReader.RLPItem[] memory items) private pure returns(Signature[] memory) {
    Signature[] memory signatures = new Signature[](items.length);
    for (uint i = 0; i < items.length; i++) {
      signatures[i] = _decodeSignature(items[i].toList());
    }
    return signatures;
  }

  function _decodeTransaction(RLPReader.RLPItem[] memory items) private pure returns(Transaction memory) {
    return Transaction({
      inputs: _decodeInputs(items[0].toList()),
      outputs: _decodeOutputs(items[1].toList()),
      metadata: _decodeMetadata(items[2].toList()),
      signatures: _decodeSignatures(items[3].toList())
    });
  }

  function _decodeTransactions(RLPReader.RLPItem[] memory items) private pure returns(Transaction[] memory) {
    Transaction[] memory transactions = new Transaction[](items.length);
    for (uint i = 0; i < items.length; i++) {
      transactions[i] = _decodeTransaction(items[i].toList());
    }
    return transactions;
  }

  function _decodeBlock(RLPReader.RLPItem[] memory items) private pure returns(Block memory) {
    return Block({
      blockNumber: uint32(items[0].toUint()),
      previousBlockHash: items[1].toUint(),
      merkleRoot: items[2].toUint(),
      signature: _decodeSignature(items[3].toList()),
      transactions: _decodeTransactions(items[4].toList())
    });
  }
}

// File: openzeppelin-solidity/contracts/cryptography/ECDSA.sol

/**
 * @title Elliptic curve signature operations
 * @dev Based on https://gist.github.com/axic/5b33912c6f61ae6fd96d6c4a47afde6d
 * TODO Remove this library once solidity supports passing a signature to ecrecover.
 * See https://github.com/ethereum/solidity/issues/864
 */

library ECDSA {
    /**
     * @dev Recover signer address from a message by using their signature
     * @param hash bytes32 message, the hash is the signed message. What is recovered is the signer address.
     * @param signature bytes signature, the signature is generated using web3.eth.sign()
     */
    function recover(bytes32 hash, bytes memory signature) internal pure returns (address) {
        bytes32 r;
        bytes32 s;
        uint8 v;

        // Check the signature length
        if (signature.length != 65) {
            return (address(0));
        }

        // Divide the signature in r, s and v variables
        // ecrecover takes the signature parameters, and the only way to get them
        // currently is to use assembly.
        // solhint-disable-next-line no-inline-assembly
        assembly {
            r := mload(add(signature, 0x20))
            s := mload(add(signature, 0x40))
            v := byte(0, mload(add(signature, 0x60)))
        }

        // Version of signature should be 27 or 28, but 0 and 1 are also possible versions
        if (v < 27) {
            v += 27;
        }

        // If the version is correct return the signer address
        if (v != 27 && v != 28) {
            return (address(0));
        } else {
            return ecrecover(hash, v, r, s);
        }
    }

    /**
     * toEthSignedMessageHash
     * @dev prefix a bytes32 value with "\x19Ethereum Signed Message:"
     * and hash the result
     */
    function toEthSignedMessageHash(bytes32 hash) internal pure returns (bytes32) {
        // 32 is the length in bytes of hash,
        // enforced by the type signature above
        return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hash));
    }
}

// File: contracts/PlasmaBlocks.sol

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
    bytes memory newBlocks
  )
    public
    onlyOwner
    returns(uint256)
  {
    _submitBlocks(fromIndex, newBlocks);
  }

  function submitBlocksSigned(
    uint256 fromIndex,
    bytes memory newBlocks,
    bytes memory rsv
  )
    public
    returns(uint256)
  {
    bytes32 messageHash = keccak256(
      abi.encodePacked(
        fromIndex,
        newBlocks
      )
    );

    bytes32 signedHash = ECDSA.toEthSignedMessageHash(messageHash);
    require(owner() == ECDSA.recover(signedHash, rsv), "Invalid signature");
    return _submitBlocks(fromIndex, newBlocks);
  }

  function _submitBlocks(
    uint256 fromIndex,
    bytes memory newBlocks
  )
    internal
    returns(uint256)
  {
    uint256 newBlocksLength = newBlocks.length / 20;

    require(fromIndex == _blocks.length, "Invalid fromIndex");

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

// File: contracts/PlasmaAssets.sol

contract PlasmaAssets is Ownable, PlasmaBlocks {
  using SafeMath for uint256;
  using SafeERC20 for IERC20;
  using PlasmaDecoder for bytes;
  using OrderedIntervalList for OrderedIntervalList.Data;
  using SumMerkleProof for SumMerkleProof.Proof;

  address constant public MAIN_COIN_ASSET_ID = address(0);
  address constant public ERC721_ASSET_ID = address(1);
  uint256 constant public ASSET_DECIMALS_TRUNCATION = 10e13; //TODO: will be different for tokens
  uint32 constant public PLASMA_ASSETS_TOTAL_SIZE = 2 ** 24 - 1;

  bytes32 private _expectedTokenAndTokenIdHash;
  mapping(address => uint256) private _assetOffsets;
  mapping(address => OrderedIntervalList.Data) private _assetLists;
  mapping(address => bytes32[]) private _allDepositHashes;
  mapping(bytes32 => bool) private _allWithdrawalHashes;
  mapping(bytes32 => bool) private _erc721Deposits;

  event AssetDeposited(
    address indexed token,
    address indexed who,
    uint64 intervalId,
    uint64 begin,
    uint64 end
  );

  event CoinDeposited(
    address indexed who,
    uint256 amount
  );

  event ERC20Deposited(
    address indexed token,
    address indexed who,
    uint256 amount
  );

  event ERC721Deposited(
    address indexed token,
    address indexed who,
    uint256 tokenId,
    uint64 indexed begin
  );

  event WithdrawalBegin(
    address owner,
    uint32 blockIndex,
    uint32 txIndex,
    uint8 outputIndex,
    address assetId,
    uint64 begin,
    uint64 end
  );

  //

  constructor() public {
    _assetLists[MAIN_COIN_ASSET_ID].initialize();
    _assetLists[ERC721_ASSET_ID].initialize();
  }

  function assetOffsets(address asset) public view returns (uint256) {
    return _assetOffsets[asset];
  }

  function setAssetOffset(address asset, uint256 assetOffset) public onlyOwner {
    require(assetOffset > uint(ERC721_ASSET_ID) && assetOffset < 256, "assetOffset should be in range [2, 255]");
    require(_assetOffsets[asset] == 0, "assetOffset was already set");
    _assetOffsets[asset] = assetOffset;
    _assetLists[asset].initialize();
  }

  // Deposits

  function deposit() public payable {
    uint64 amount = uint64(msg.value / ASSET_DECIMALS_TRUNCATION);
    (uint64 intervalId, uint64 begin, uint64 end) = _assetLists[MAIN_COIN_ASSET_ID].append(amount);

    emit CoinDeposited(msg.sender, amount);
    emit AssetDeposited(MAIN_COIN_ASSET_ID, msg.sender, intervalId, begin, end);
    bytes32 hash = keccak256(abi.encodePacked(msg.sender, amount));
    _allDepositHashes[msg.sender].push(hash);
  }

  function depositERC20(IERC20 token, uint256 amountArg) public {
    require(_assetLists[address(token)].isInitialized(), "Operator should add this token first");

    uint64 amount = uint64(amountArg / ASSET_DECIMALS_TRUNCATION);
    (uint64 intervalId, uint64 begin, uint64 end) = _assetLists[address(token)].append(amount);

    uint256 preBalance = token.balanceOf(address(this));
    token.safeTransferFrom(msg.sender, address(this), amount);
    uint256 deposited = token.balanceOf(address(this)).sub(preBalance);

    emit ERC20Deposited(address(token), msg.sender, deposited);
    emit AssetDeposited(address(token), msg.sender, intervalId, begin, end);
    bytes32 hash = keccak256(abi.encodePacked(token, msg.sender, intervalId, begin, end));
    _allDepositHashes[msg.sender].push(hash);
  }

  function depositERC721(IERC721 token, uint256 tokenId) public {
    _expectedTokenAndTokenIdHash = keccak256(abi.encodePacked(token, tokenId));
    token.safeTransferFrom(msg.sender, address(this), tokenId);
    require(_expectedTokenAndTokenIdHash == bytes32(0), "ERC721 token not received");

    (uint64 intervalId, uint64 begin, uint64 end) = _assetLists[ERC721_ASSET_ID].append(1);

    emit ERC721Deposited(address(token), msg.sender, tokenId, begin);
    emit AssetDeposited(ERC721_ASSET_ID, msg.sender, intervalId, begin, end);

    bytes32 hash = keccak256(abi.encodePacked(ERC721_ASSET_ID, msg.sender, intervalId, begin, end));
    _allDepositHashes[msg.sender].push(hash);

    bytes32 erc721Hash = keccak256(abi.encodePacked(token, tokenId, begin));
    _erc721Deposits[erc721Hash] = true;
  }

  function onERC721Received(
    address operator,
    address /*from*/,
    uint256 tokenId,
    bytes memory /*data*/
  )
  public
  returns (bytes4)
  {
    bytes32 receivedTokenAndTokenId = keccak256(abi.encodePacked(msg.sender, tokenId));

    require(operator == address(this), "Only this contract should deposit ERC721 tokens");
    require(receivedTokenAndTokenId == _expectedTokenAndTokenIdHash, "ERC721 token was not expected");
    delete _expectedTokenAndTokenIdHash;
    return this.onERC721Received.selector;
  }

  // Withdrawals

  function withdrawalBegin(
    bytes memory inputBytes // PlasmaDecoder.Input
  )
  public
  payable //TODO: Bonds
  returns (bool)
  {
    PlasmaDecoder.Input memory input = inputBytes.decodeInput();

    emit WithdrawalBegin(
      input.owner,
      input.blockIndex,
      input.txIndex,
      input.outputIndex,
      input.assetId,
      input.begin,
      input.end
    );

    bytes32 inputHash = keccak256(
      abi.encodePacked(
        input.owner,
        input.blockIndex,
        input.txIndex,
        input.outputIndex,
        input.assetId,
        input.begin,
        input.end
      ));
    _allWithdrawalHashes[inputHash] = true;

    return true;
  }

  function withdrawalChallangeSpend(
    bytes memory inputBytes, // PlasmaDecoder.Input
    bytes memory txProofBytes, // SumMerkleProof.Proof
    uint64 blockIndex,
    uint8 /*spendIndex*/
  )
  public
  returns (bool)
  {
    PlasmaDecoder.Input memory input = inputBytes.decodeInput();
    SumMerkleProof.Proof memory txProof = txProofBytes.decodeProof();

    bytes32 inputHash = keccak256(
      abi.encodePacked(
        input.owner,
        input.blockIndex,
        input.txIndex,
        input.outputIndex,
        input.assetId,
        input.begin,
        input.end
      ));
    require(_allWithdrawalHashes[inputHash], "You should start withdrawal first");

    require(txProof.sumMerkleProof(blocks(blockIndex), PLASMA_ASSETS_TOTAL_SIZE));

    // Cancel widthraw
    delete _allWithdrawalHashes[inputHash];

    return true;
  }

  // function withdrawalChallangeExistance(
  //   ExitState state,
  //   SumMerkleProof txProof,
  //   MerkleProof inputProof,
  //   uint64 maxBlockIndex,
  //   MerkleProof maxBlockIndexProof
  // )
  //   public
  //   returns(bool)
  // {
  // }

  function withdrawalEnd(
    bytes memory inputBytes, // PlasmaDecoder.Input
    uint64 intervalId,
    address token,
    uint256 tokenId
  ) public {
    PlasmaDecoder.Input memory input = inputBytes.decodeInput();

    bytes32 inputHash = keccak256(
      abi.encodePacked(
        input.owner,
        input.blockIndex,
        input.txIndex,
        input.outputIndex,
        input.assetId,
        input.begin,
        input.end
      ));
    require(_allWithdrawalHashes[inputHash], "You should start withdrawal first");
    delete _allWithdrawalHashes[inputHash];

    // Update interval and check it exist
    _assetLists[input.assetId].remove(intervalId, input.begin, input.end);

    if (input.assetId == MAIN_COIN_ASSET_ID) {
      input.owner.transfer(uint256(input.end).sub(input.begin).mul(ASSET_DECIMALS_TRUNCATION));
      return;
    }

    if (input.assetId == ERC721_ASSET_ID) {
      require(input.end == input.begin + 1, "It is allowed to withdraw only 1 ERC721 per transaction");
      bytes32 depositHash = keccak256(abi.encodePacked(token, tokenId, input.begin));
      require(_erc721Deposits[depositHash], "Invalid token or tokeId arguments");
      delete _erc721Deposits[depositHash];
      IERC721(token).approve(msg.sender, tokenId);
      return;
    }

    IERC20(token).transfer(msg.sender, uint256(input.end).sub(input.begin));
  }
}

// File: contracts/BankexPlasma.sol

contract BankexPlasma is PlasmaAssets {

}
