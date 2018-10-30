pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/token/ERC20/IERC20.sol";
import "openzeppelin-solidity/contracts/token/ERC20/SafeERC20.sol";
import "openzeppelin-solidity/contracts/token/ERC721/IERC721.sol";
import "openzeppelin-solidity/contracts/math/SafeMath.sol";


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
