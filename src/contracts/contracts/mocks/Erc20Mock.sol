pragma solidity ^0.5.2;

import {Ownable} from "openzeppelin-solidity/contracts/ownership/Ownable.sol";
import {ERC20Mintable} from "openzeppelin-solidity/contracts/token/ERC20/ERC20Mintable.sol";

contract Erc20Mock is ERC20Mintable, Ownable {
    // TODO: function/constructor with parameters bellow
    string public name = "Mock token";

    string public symbol = "MCK";

    uint public decimals = 18;
}
