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

    event AssetDeposited(
        address indexed token,
        address indexed who,
        uint256 amount
    );

    function calculateAssetId(address token, uint256 tokenId) public pure returns(address) {
        return address(keccak256(abi.encodePacked(token, tokenId)));
    }

    // Deposits

    function deposit() public payable {
        emit AssetDeposited(MAIN_COIN_ADDRESS, msg.sender, msg.value);
    }

    function depositERC20(IERC20 token, uint256 amount) public {
        uint256 preBalance = token.balanceOf(this);
        token.safeTransferFrom(msg.sender, this, amount);
        uint256 deposited = token.balanceOf(this).sub(preBalance);

        emit AssetDeposited(token, msg.sender, deposited);
    }

    function depositERC721(IERC721 token, uint256 tokenId) public payable {
        address assetId = calculateAssetId(token, tokenId);
        _expectedAssetId = assetId;
        token.safeTransferFrom(msg.sender, this, tokenId);
        require(_expectedAssetId == address(0), "ERC721 token not received");

        emit AssetDeposited(assetId, msg.sender, tokenId);
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
