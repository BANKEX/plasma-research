pragma solidity ^0.4.24;


library Uint2048 {
  using Uint2048 for uint256[8];

  function eq(uint256[8] a, uint256[8] b) internal pure returns(bool) {
    return
      a[0] == b[0] && // solium-disable-line operator-whitespace
      a[1] == b[1] &&
      a[2] == b[2] &&
      a[3] == b[3] &&
      a[4] == b[4] &&
      a[5] == b[5] &&
      a[6] == b[6] &&
      a[7] == b[7];
  }

  function lt(uint256[8] a, uint256[8] b) internal pure returns(bool) {
    if (a[0] != b[0]) {
      return a[0] < b[0];
    }
    if (a[1] != b[1]) {
      return a[1] < b[1];
    }
    if (a[2] != b[2]) {
      return a[2] < b[2];
    }
    if (a[3] != b[3]) {
      return a[3] < b[3];
    }
    if (a[4] != b[4]) {
      return a[4] < b[4];
    }
    if (a[5] != b[5]) {
      return a[5] < b[5];
    }
    if (a[6] != b[6]) {
      return a[6] < b[6];
    }
    if (a[7] != b[7]) {
      return a[7] < b[7];
    }
  }

  function gt(uint256[8] a, uint256[8] b) internal pure returns(bool) {
    return !eq(a, b) && !lt(a, b);
  }

  function le(uint256[8] a, uint256[8] b) internal pure returns(bool) {
    return lt(a, b) || eq(a, b);
  }

  function ge(uint256[8] a, uint256[8] b) internal pure returns(bool) {
    return !lt(a, b);
  }

  // Math

  function add(uint256[8] a, uint256[8] b) internal pure returns(uint256[8] c) {
    c[7] = a[7] + b[7];
    c[6] = a[6] + b[6] + (a[7] > c[7] ? 1 : 0);
    c[5] = a[5] + b[5] + (a[6] > c[6] ? 1 : 0);
    c[4] = a[4] + b[4] + (a[5] > c[5] ? 1 : 0);
    c[3] = a[3] + b[3] + (a[4] > c[4] ? 1 : 0);
    c[2] = a[2] + b[2] + (a[3] > c[3] ? 1 : 0);
    c[1] = a[1] + b[1] + (a[2] > c[2] ? 1 : 0);
    c[0] = a[0] + b[0] + (a[1] > c[1] ? 1 : 0);
  }

  function sub(uint256[8] a, uint256[8] b) internal pure returns(uint256[8] c) {
    c[7] = a[7] - b[7];
    c[6] = a[6] - b[6] - (a[7] < c[7] ? 1 : 0);
    c[5] = a[5] - b[5] - (a[6] < c[6] ? 1 : 0);
    c[4] = a[4] - b[4] - (a[5] < c[5] ? 1 : 0);
    c[3] = a[3] - b[3] - (a[4] < c[4] ? 1 : 0);
    c[2] = a[2] - b[2] - (a[3] < c[3] ? 1 : 0);
    c[1] = a[1] - b[1] - (a[2] < c[2] ? 1 : 0);
    c[0] = a[0] - b[0] - (a[1] < c[1] ? 1 : 0);
  }

  function mul(uint256[8] a, uint256[8] b) internal pure returns(uint256[8] c) {
    // TODO: implement
  }

  function div(uint256[8] a, uint256[8] b) internal pure returns(uint256[8] c) {
    // TODO: implement
  }

  // Modular

  function mod(uint256[8] a, uint256[8] b) internal pure returns(uint256[8] c) {
    // TODO: implement
  }

  function addmod(uint256[8] a, uint256[8] b, uint256[8] m) internal pure returns(uint256[8] c) {
    // TODO: implement
  }

  function mulmod(uint256[8] a, uint256[8] b, uint256[8] m) internal pure returns(uint256[8] c) {
    // TODO: implement
  }

  function powmod(uint256[8] a, uint256 p, uint256[8] m) internal pure returns(uint256[8] c) {
    // TODO: implement
  }
}