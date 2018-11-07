pragma solidity ^0.4.24;

import { Uint2048 } from "../Uint2048.sol";


contract Uint2048Wrapper {
  function eq(uint256[8] a, uint256[8] b) public pure returns(bool) {
    return Uint2048.eq(a,b);
  }

  function lt(uint256[8] a, uint256[8] b) public pure returns(bool) {
    return Uint2048.lt(a,b);
  }

  function gt(uint256[8] a, uint256[8] b) public pure returns(bool) {
    return Uint2048.gt(a,b);
  }

  function le(uint256[8] a, uint256[8] b) public pure returns(bool) {
    return Uint2048.le(a,b);
  }

  function ge(uint256[8] a, uint256[8] b) public pure returns(bool) {
    return Uint2048.ge(a,b);
  }

  // Math

  function add(uint256[8] a, uint256[8] b) public pure returns(uint256[8] c) {
    return Uint2048.add(a,b);
  }

  function sub(uint256[8] a, uint256[8] b) public pure returns(uint256[8] c) {
    return Uint2048.sub(a,b);
  }

  function mul(uint256[8] a, uint256[8] b) public pure returns(uint256[8] c) {
    return Uint2048.mul(a,b);
  }

  function div(uint256[8] a, uint256[8] b) public pure returns(uint256[8] c) {
    return Uint2048.div(a,b);
  }

  // Modular

  function mod(uint256[8] a, uint256[8] b) public pure returns(uint256[8] c) {
    return Uint2048.mod(a,b);
  }

  function addmod(uint256[8] a, uint256[8] b, uint256[8] m) public pure returns(uint256[8] c) {
    return Uint2048.addmod(a,b,m);
  }

  function mulmod(uint256[8] a, uint256[8] b, uint256[8] m) public pure returns(uint256[8] c) {
    return Uint2048.mulmod(a,b,m);
  }

  function powmod(uint256[8] a, uint256 p, uint256[8] m) public pure returns(uint256[8] c) {
    return Uint2048.powmod(a,p,m);
  }
}
