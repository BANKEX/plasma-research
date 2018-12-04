pragma solidity ^0.4.24;

import { SafeMath } from "openzeppelin-solidity/contracts/math/SafeMath.sol";
import { Ownable } from "openzeppelin-solidity/contracts/ownership/Ownable.sol";
import { IERC20 } from "openzeppelin-solidity/contracts/token/ERC20/IERC20.sol";
import { SafeERC20 } from "openzeppelin-solidity/contracts/token/ERC20/SafeERC20.sol";
import { IERC721 } from "openzeppelin-solidity/contracts/token/ERC721/IERC721.sol";
import { OrderedIntervalList } from "./OrderedIntervalList.sol";


contract PlasmaAssets is Ownable {
  using SafeMath for uint256;
  using SafeERC20 for IERC20;
  using OrderedIntervalList for OrderedIntervalList.Data;

  address constant public MAIN_COIN_ASSET_ID = address(0);
  address constant public ERC721_ASSET_ID = address(1);
  uint256 constant public ASSET_DECIMALS_TRUNCATION = 10e13; //TODO: will be different for tokens
  address constant public MAIN_COIN_ADDRESS = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

  bytes32 private _expectedTokenAndTokenIdHash;
  mapping (address => uint256) private _assetOffsets;
  mapping (address => OrderedIntervalList.Data) private _assetLists;
  mapping (address => bytes32[]) private _allDepositHashes;

  event AssetDeposited(
    address indexed token,
    address indexed who,
    uint256 intervalId,
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

  constructor() public {
    _assetLists[MAIN_COIN_ASSET_ID].initialize();
    _assetLists[ERC721_ASSET_ID].initialize();
  }

  function assetOffsets(address asset) public view returns(uint256) {
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
    (uint256 intervalId, uint64 begin, uint64 end) = _assetLists[MAIN_COIN_ADDRESS].append(amount);

    emit CoinDeposited(msg.sender, amount);
    emit AssetDeposited(MAIN_COIN_ADDRESS, msg.sender, intervalId, begin, end);
    bytes32 hash = keccak256(abi.encodePacked(msg.sender, amount));
    _allDepositHashes[msg.sender].push(hash);
  }

  function depositERC20(IERC20 token, uint256 amountArg) public {
    require(_assetLists[token].isInitialized(), "Operator should add this token first");

    uint64 amount = uint64(amountArg / ASSET_DECIMALS_TRUNCATION);
    (uint256 intervalId, uint64 begin, uint64 end) = _assetLists[token].append(amount);

    uint256 preBalance = token.balanceOf(this);
    token.safeTransferFrom(msg.sender, this, amount);
    uint256 deposited = token.balanceOf(this).sub(preBalance);
    
    emit ERC20Deposited(token, msg.sender, deposited);
    emit AssetDeposited(token, msg.sender, intervalId, begin, end);
    bytes32 hash = keccak256(abi.encodePacked(token, msg.sender, intervalId, begin, end));
    _allDepositHashes[msg.sender].push(hash);
  }

  function depositERC721(IERC721 token, uint256 tokenId) public {
    _expectedTokenAndTokenIdHash = keccak256(abi.encodePacked(token, tokenId));
    token.safeTransferFrom(msg.sender, this, tokenId);
    require(_expectedTokenAndTokenIdHash == bytes32(0), "ERC721 token not received");

    (uint256 intervalId, uint64 begin, uint64 end) = _assetLists[ERC721_ASSET_ID].append(1);

    emit ERC721Deposited(token, msg.sender, tokenId, begin);
    emit AssetDeposited(ERC721_ASSET_ID, msg.sender, intervalId, begin, end);
    bytes32 hash = keccak256(abi.encodePacked(ERC721_ASSET_ID, msg.sender, intervalId, begin, end));
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
    bytes32 receivedTokenAndTokenId = keccak256(abi.encodePacked(msg.sender, tokenId));

    require(operator == address(this), "Only this contract should deposit ERC721 tokens");
    require(receivedTokenAndTokenId == _expectedTokenAndTokenIdHash, "ERC721 token was not expected");
    delete _expectedTokenAndTokenIdHash;
    return this.onERC721Received.selector;
  }

  // Withdrawals

  // function withdrawalBegin(
  //   Input point,  
  //   RSAInclusionProof proof
  // )
  //   external
  //   payable
  //   returns(bool)
  // {
  // }

  // function withdrawalChallangeSpend(
  //   ExitState state, 
  //   Transaction tx,
  //   uint64 blockIndex,
  //   SumMerkleProof[] txProof, // serialized to bytes
  //   uint8 spendIndex,
  //   RSAInclusionProof spendInclusionProof
  // )
  //   external
  //   returns(bool)
  // {
  // }

  // function withdrawalChallangeExistance(
  //   ExitState state,
  //   SumMerkleProof txProof,
  //   MerkleProof inputProof, 
  //   uint64 maxBlockIndex,
  //   MerkleProof maxBlockIndexProof
  // )
  //   external
  //   returns(bool)
  // {
  // }
    
  // function withdrawalEnd(ExitState state) public {
  // }
}
