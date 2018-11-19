pragma solidity ^0.4.24;

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
  function recover(bytes32 hash, bytes signature)
    internal
    pure
    returns (address)
  {
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
    // solium-disable-next-line security/no-inline-assembly
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
      // solium-disable-next-line arg-overflow
      return ecrecover(hash, v, r, s);
    }
  }

  /**
   * toEthSignedMessageHash
   * @dev prefix a bytes32 value with "\x19Ethereum Signed Message:"
   * and hash the result
   */
  function toEthSignedMessageHash(bytes32 hash)
    internal
    pure
    returns (bytes32)
  {
    // 32 is the length in bytes of hash,
    // enforced by the type signature above
    return keccak256(
      abi.encodePacked("\x19Ethereum Signed Message:\n32", hash)
    );
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

  event OwnershipTransferred(
    address indexed previousOwner,
    address indexed newOwner
  );

  /**
   * @dev The Ownable constructor sets the original `owner` of the contract to the sender
   * account.
   */
  constructor() internal {
    _owner = msg.sender;
    emit OwnershipTransferred(address(0), _owner);
  }

  /**
   * @return the address of the owner.
   */
  function owner() public view returns(address) {
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
  function isOwner() public view returns(bool) {
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

// File: openzeppelin-solidity/contracts/math/SafeMath.sol

/**
 * @title SafeMath
 * @dev Math operations with safety checks that revert on error
 */
library SafeMath {

  /**
  * @dev Multiplies two numbers, reverts on overflow.
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
  * @dev Integer division of two numbers truncating the quotient, reverts on division by zero.
  */
  function div(uint256 a, uint256 b) internal pure returns (uint256) {
    require(b > 0); // Solidity only automatically asserts when dividing by 0
    uint256 c = a / b;
    // assert(a == b * c + a % b); // There is no case in which this doesn't hold

    return c;
  }

  /**
  * @dev Subtracts two numbers, reverts on overflow (i.e. if subtrahend is greater than minuend).
  */
  function sub(uint256 a, uint256 b) internal pure returns (uint256) {
    require(b <= a);
    uint256 c = a - b;

    return c;
  }

  /**
  * @dev Adds two numbers, reverts on overflow.
  */
  function add(uint256 a, uint256 b) internal pure returns (uint256) {
    uint256 c = a + b;
    require(c >= a);

    return c;
  }

  /**
  * @dev Divides two numbers and returns the remainder (unsigned integer modulo),
  * reverts when dividing by zero.
  */
  function mod(uint256 a, uint256 b) internal pure returns (uint256) {
    require(b != 0);
    return a % b;
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

// File: openzeppelin-solidity/contracts/token/ERC20/IERC20.sol

/**
 * @title ERC20 interface
 * @dev see https://github.com/ethereum/EIPs/issues/20
 */
interface IERC20 {
  function totalSupply() external view returns (uint256);

  function balanceOf(address who) external view returns (uint256);

  function allowance(address owner, address spender)
    external view returns (uint256);

  function transfer(address to, uint256 value) external returns (bool);

  function approve(address spender, uint256 value)
    external returns (bool);

  function transferFrom(address from, address to, uint256 value)
    external returns (bool);

  event Transfer(
    address indexed from,
    address indexed to,
    uint256 value
  );

  event Approval(
    address indexed owner,
    address indexed spender,
    uint256 value
  );
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

  function safeTransfer(
    IERC20 token,
    address to,
    uint256 value
  )
    internal
  {
    require(token.transfer(to, value));
  }

  function safeTransferFrom(
    IERC20 token,
    address from,
    address to,
    uint256 value
  )
    internal
  {
    require(token.transferFrom(from, to, value));
  }

  function safeApprove(
    IERC20 token,
    address spender,
    uint256 value
  )
    internal
  {
    // safeApprove should only be called when setting an initial allowance, 
    // or when resetting it to zero. To increase and decrease it, use 
    // 'safeIncreaseAllowance' and 'safeDecreaseAllowance'
    require((value == 0) || (token.allowance(msg.sender, spender) == 0));
    require(token.approve(spender, value));
  }

  function safeIncreaseAllowance(
    IERC20 token,
    address spender,
    uint256 value
  )
    internal
  {
    uint256 newAllowance = token.allowance(address(this), spender).add(value);
    require(token.approve(spender, newAllowance));
  }

  function safeDecreaseAllowance(
    IERC20 token,
    address spender,
    uint256 value
  )
    internal
  {
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
  function supportsInterface(bytes4 interfaceId)
    external
    view
    returns (bool);
}

// File: openzeppelin-solidity/contracts/token/ERC721/IERC721.sol

/**
 * @title ERC721 Non-Fungible Token Standard basic interface
 * @dev see https://github.com/ethereum/EIPs/blob/master/EIPS/eip-721.md
 */
contract IERC721 is IERC165 {

  event Transfer(
    address indexed from,
    address indexed to,
    uint256 indexed tokenId
  );
  event Approval(
    address indexed owner,
    address indexed approved,
    uint256 indexed tokenId
  );
  event ApprovalForAll(
    address indexed owner,
    address indexed operator,
    bool approved
  );

  function balanceOf(address owner) public view returns (uint256 balance);
  function ownerOf(uint256 tokenId) public view returns (address owner);

  function approve(address to, uint256 tokenId) public;
  function getApproved(uint256 tokenId)
    public view returns (address operator);

  function setApprovalForAll(address operator, bool _approved) public;
  function isApprovedForAll(address owner, address operator)
    public view returns (bool);

  function transferFrom(address from, address to, uint256 tokenId) public;
  function safeTransferFrom(address from, address to, uint256 tokenId)
    public;

  function safeTransferFrom(
    address from,
    address to,
    uint256 tokenId,
    bytes data
  )
    public;
}

// File: contracts/PlasmaAssets.sol

contract PlasmaAssets {
  using SafeMath for uint256;
  using SafeERC20 for IERC20;

  address constant public MAIN_COIN_ADDRESS = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

  address private _expectedAssetId;
  mapping (address => bytes32[]) private _allDepositHashes;

  event AssetDeposited(
    address indexed token,
    address indexed who,
    uint256 amount
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
    uint256 indexed tokenId
  );

  function calculateAssetId(address token, uint256 tokenId) public pure returns(address) {
    return address(keccak256(abi.encodePacked(token, tokenId)));
  }

  // Deposits

  function deposit() public payable {
    emit CoinDeposited(msg.sender, msg.value);
    emit AssetDeposited(MAIN_COIN_ADDRESS, msg.sender, msg.value);
    bytes32 hash = keccak256(abi.encodePacked(msg.sender, msg.value));
    _allDepositHashes[msg.sender].push(hash);
  }

  function depositERC20(IERC20 token, uint256 amount) public {
    uint256 preBalance = token.balanceOf(this);
    token.safeTransferFrom(msg.sender, this, amount);
    uint256 deposited = token.balanceOf(this).sub(preBalance);

    emit ERC20Deposited(token, msg.sender, deposited);
    emit AssetDeposited(token, msg.sender, deposited);
    bytes32 hash = keccak256(abi.encodePacked(token, msg.sender, deposited));
    _allDepositHashes[msg.sender].push(hash);
  }

  function depositERC721(IERC721 token, uint256 tokenId) public {
    address assetId = calculateAssetId(token, tokenId);
    _expectedAssetId = assetId;
    token.safeTransferFrom(msg.sender, this, tokenId);
    require(_expectedAssetId == address(0), "ERC721 token not received");

    emit ERC721Deposited(token, msg.sender, tokenId);
    emit AssetDeposited(assetId, msg.sender, tokenId);
    bytes32 hash = keccak256(abi.encodePacked(assetId, token, msg.sender, tokenId));
    _allDepositHashes[msg.sender].push(hash);
  }

  function onERC721Received(
    address operator,
    address /*from*/,
    uint256 tokenId,
    bytes /*data*/
  )
    public
    returns(bytes4)
  {
    require(operator == address(this), "Only this contract should deposit ERC721 tokens");
    require(calculateAssetId(msg.sender, tokenId) == _expectedAssetId, "ERC721 token was not expected");
    delete _expectedAssetId;
    return this.onERC721Received.selector;
  }
}

// File: contracts/BankexPlasma.sol

contract BankexPlasma is PlasmaBlocks, PlasmaAssets {

}
