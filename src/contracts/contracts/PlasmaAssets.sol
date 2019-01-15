pragma solidity ^0.5.0;

import {SafeMath} from "openzeppelin-solidity/contracts/math/SafeMath.sol";
import {Ownable} from "openzeppelin-solidity/contracts/ownership/Ownable.sol";
import {IERC20} from "openzeppelin-solidity/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "openzeppelin-solidity/contracts/token/ERC20/SafeERC20.sol";
import {IERC721} from "openzeppelin-solidity/contracts/token/ERC721/IERC721.sol";
import {OrderedIntervalList} from "./OrderedIntervalList.sol";
import {SumMerkleProof} from "./SumMerkleProof.sol";
import {PlasmaDecoder} from "./PlasmaDecoder.sol";
import {PlasmaBlocks} from "./PlasmaBlocks.sol";


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
