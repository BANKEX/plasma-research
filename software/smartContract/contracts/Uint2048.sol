pragma solidity ^0.4.24;


library Uint2048 {
  using Uint2048 for uint256[LEN];

  uint256 constant public LEN = 8;
  uint256 constant public MAX_UINT256 = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF;

  function eq(uint256[LEN] a, uint256[LEN] b) internal pure returns(bool) {
    for (uint i = 0; i < LEN; i++) {
      if (a[i] != b[i]) {
        return false;
      }
    }
    return true;
  }

  function lt(uint256[LEN] a, uint256[LEN] b) internal pure returns(bool) {
    for (uint i = 0; i < LEN; i++) {
      if (a[i] != b[i]) {
        return a[i] < b[i];
      }
    }
    return false;
  }

  function gt(uint256[LEN] a, uint256[LEN] b) internal pure returns(bool) {
    return !eq(a, b) && !lt(a, b);
  }

  function le(uint256[LEN] a, uint256[LEN] b) internal pure returns(bool) {
    return lt(a, b) || eq(a, b);
  }

  function ge(uint256[LEN] a, uint256[LEN] b) internal pure returns(bool) {
    return !lt(a, b);
  }

  // Math

  function inverted(uint256[LEN] a) internal pure returns(uint256[LEN] c) {
    for (uint i = 0; i < LEN; i++) {
      c[i] = ~a[i];
    }
  }

  function incremented(uint256[LEN] a) internal pure returns(uint256[LEN]) {
    for (uint i = 0; i < LEN; i++) {
      if (a[i] < MAX_UINT256) {
        a[i] += 1;
        break;
      }
    }

    while (i > 0) {
      a[--i] = 0;
    }

    return a;
  }

  function add(uint256[LEN] a, uint256[LEN] b) internal pure returns(uint256[LEN] c) {
    c[0] = a[0] + b[0];
    for (uint i = 1; i < LEN; i++) {
      c[i] = a[i] + b[i] + (a[i - 1] > c[i - 1] ? 1 : 0);
    }
  }

  function sub(uint256[LEN] a, uint256[LEN] b) internal pure returns(uint256[LEN] c) {
    //return a.add(b.inverted().incremented());
    c[0] = a[0] - b[0];
    for (uint i = 1; i < LEN; i++) {
      c[i] = a[i] - b[i] - (a[i - 1] < c[i - 1] ? 1 : 0);
    }
  }

  function mul(uint256[LEN] a, uint256[LEN] b) internal pure returns(uint256[LEN] c) {
    // TODO: implement
  }

  function div(uint256[LEN] a, uint256[LEN] b) internal pure returns(uint256[LEN] c) {
    // TODO: implement
  }

  // Modular

  function mod(uint256[LEN] a, uint256[LEN] b) internal pure returns(uint256[LEN] c) {
    // TODO: implement
  }

  function addmod(uint256[LEN] a, uint256[LEN] b, uint256[LEN] m) internal pure returns(uint256[LEN] c) {
    // TODO: implement
  }

  function mulmod(uint256[LEN] a, uint256[LEN] b, uint256[LEN] m) internal pure returns(uint256[LEN] c) {
    // TODO: implement
  }

  function powmod(uint256[LEN] a, uint256 p, uint256[LEN] m) internal pure returns(uint256[LEN] c) {
    // TODO: implement
  }
}