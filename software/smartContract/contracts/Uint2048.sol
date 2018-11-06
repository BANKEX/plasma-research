pragma solidity ^0.4.24;


library Uint2048 {
  using Uint2048 for uint256;
  using Uint2048 for uint256[2];
  using Uint2048 for uint256[4];
  using Uint2048 for uint256[8];
  using Uint2048 for uint256[16];

  function concat128(uint128 a, uint128 b) public pure returns(uint256) {
    return (uint256(a) << 128) | b;
  }

  function concat(uint256 a, uint256 /*b*/) public pure returns(uint256[2] c) {
    return [a, b];
  }

  function concat(uint256[2] a, uint256[2] /*b*/) public pure returns(uint256[4] c) {
    return [
      a[0], a[1],
      b[0], b[1]
    ];
  }

  function concat(uint256[4] a, uint256[4] /*b*/) public pure returns(uint256[8] c) {
    return [
      a[0], a[1], a[2], a[3],
      b[0], b[1], b[2], b[3]
    ];
  }

  function concat(uint256[8] a, uint256[8] /*b*/) public pure returns(uint256[16] c) {
    return [
      a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7],
      b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7]
    ];
  }

  //

  // position: 0 - high, 1 - mid, 2 - low
  function half(uint256 a, uint256 position) public pure returns(uint128 c) {
    return uint128(a >> (128 - position * 64));
  }

  // position: 0 - high, 1 - mid, 2 - low
  function half(uint256[2] a, uint256 position) public pure returns(uint256 c) {
    assembly {
      c := add(a, mul(position, 0x10))
    }
  }
  
  // position: 0 - high, 1 - mid, 2 - low
  function half(uint256[4] a, uint256 position) public pure returns(uint256[2] c) {
    assembly {
      c := add(a, mul(position, 0x20))
    }
  }
  
  // position: 0 - high, 1 - mid, 2 - low
  function half(uint256[8] a, uint256 position) public pure returns(uint256[4] c) {
    assembly {
      c := add(a, mul(position, 0x40))
    }
  }

  // position: 0 - high, 1 - mid, 2 - low
  function half(uint256[16] a, uint256 position) public pure returns(uint256[8] c) {
    assembly {
      c := add(a, mul(position, 0x40))
    }
  }

  //

  function lo(uint256 a) public pure returns(uint128) {
    return half(a, 2);
  }

  function lo(uint256[2] a) public pure returns(uint256) {
    return half(a, 2);
  }

  function lo(uint256[4] a) public pure returns(uint256[2]) {
    return half(a, 2);
  }

  function lo(uint256[8] a) public pure returns(uint256[4]) {
    return half(a, 2);
  }

  function lo(uint256[16] a) public pure returns(uint256[8]) {
    return half(a, 2);
  }

  //

  function hi(uint256 a) public pure returns(uint128) {
    return half(a, 0);
  }

  function hi(uint256[2] a) public pure returns(uint256) {
    return half(a, 0);
  }

  function hi(uint256[4] a) public pure returns(uint256[2]) {
    return half(a, 0);
  }

  function hi(uint256[8] a) public pure returns(uint256[4]) {
    return half(a, 0);
  }

  function hi(uint256[16] a) public pure returns(uint256[8]) {
    return half(a, 0);
  }

  //

  function mid(uint256 a) public pure returns(uint128 c) {
    return half(a, 1);
  }

  function mid(uint256[2] a) public pure returns(uint256 c) {
    return half(a, 1);
  }

  function mid(uint256[4] a) public pure returns(uint256[2]) {
    return half(a, 1);
  }

  function mid(uint256[8] a) public pure returns(uint256[4]) {
    return half(a, 1);
  }

  function mid(uint256[16] a) public pure returns(uint256[8]) {
    return half(a, 1);
  }

  ////////////////////////////////////////////////////////////////
  // Compare
  ////////////////////////////////////////////////////////////////

  function isZero(uint256 a) internal pure returns(bool) {
    return a == 0;
  }

  function isZero(uint256[2] a) internal pure returns(bool) {
    return isZero(a.lo()) && isZero(a.hi());
  }

  function isZero(uint256[4] a) internal pure returns(bool) {
    return isZero(a.lo()) && isZero(a.hi());
  }

  function isZero(uint256[8] a) internal pure returns(bool) {
    return isZero(a.lo()) && isZero(a.hi());
  }
  
  //

  function eq(uint256 a, uint256 b) internal pure returns(bool) {
    return a == b;
  }

  function eq(uint256[2] a, uint256[2] b) internal pure returns(bool) {
    return eq(a.hi(), b.hi()) && eq(a.lo(), b.lo());
  }

  function eq(uint256[4] a, uint256[4] b) internal pure returns(bool) {
    return eq(a.hi(), b.hi()) && eq(a.lo(), b.lo());
  }

  function eq(uint256[8] a, uint256[8] b) internal pure returns(bool) {
    return eq(a.hi(), b.hi()) && eq(a.lo(), b.lo());
  }

  //

  function lt(uint256 a, uint256 b) internal pure returns(bool) {
    return a.hi() < b.hi() || (a.hi() == b.hi() && a.lo() < b.lo());
  }

  function lt(uint256[2] a, uint256[2] b) internal pure returns(bool) {
    return lt(a.hi(), b.hi()) || (eq(a.hi(), b.hi()) && lt(a.lo(), b.lo()));
  }

  function lt(uint256[4] a, uint256[4] b) internal pure returns(bool) {
    return lt(a.hi(), b.hi()) || (eq(a.hi(), b.hi()) && lt(a.lo(), b.lo()));
  }

  function lt(uint256[8] a, uint256[8] b) internal pure returns(bool) {
    return lt(a.hi(), b.hi()) || (eq(a.hi(), b.hi()) && lt(a.lo(), b.lo()));
  }

  //

  function gt(uint256 a, uint256 b) internal pure returns(bool) {
    return !eq(a, b) && !lt(a, b);
  }

  function gt(uint256[2] a, uint256[2] b) internal pure returns(bool) {
    return !eq(a, b) && !lt(a, b);
  }

  function gt(uint256[4] a, uint256[4] b) internal pure returns(bool) {
    return !eq(a, b) && !lt(a, b);
  }

  function gt(uint256[8] a, uint256[8] b) internal pure returns(bool) {
    return !eq(a, b) && !lt(a, b);
  }

  //

  function le(uint256 a, uint256 b) internal pure returns(bool) {
    return lt(a, b) || eq(a, b);
  }

  function le(uint256[2] a, uint256[2] b) internal pure returns(bool) {
    return lt(a, b) || eq(a, b);
  }

  function le(uint256[4] a, uint256[4] b) internal pure returns(bool) {
    return lt(a, b) || eq(a, b);
  }

  function le(uint256[8] a, uint256[8] b) internal pure returns(bool) {
    return lt(a, b) || eq(a, b);
  }

  //

  function ge(uint256 a, uint256 b) internal pure returns(bool) {
    return !lt(a, b);
  }

  function ge(uint256[2] a, uint256[2] b) internal pure returns(bool) {
    return !lt(a, b);
  }

  function ge(uint256[4] a, uint256[4] b) internal pure returns(bool) {
    return !lt(a, b);
  }

  function ge(uint256[8] a, uint256[8] b) internal pure returns(bool) {
    return !lt(a, b);
  }

  ////////////////////////////////////////////////////////////////
  // Math
  ////////////////////////////////////////////////////////////////

  function inverted(uint256 a) internal pure returns(uint256) {
    return uint256(~a);
  }

  function inverted(uint256[2] a) internal pure returns(uint256[2]) {
    return concat(inverted(a.hi()), inverted(a.lo()));
  }

  function inverted(uint256[4] a) internal pure returns(uint256[4]) {
    return concat(inverted(a.hi()), inverted(a.lo()));
  }

  function inverted(uint256[8] a) internal pure returns(uint256[8]) {
    return concat(inverted(a.hi()), inverted(a.lo()));
  }

  //

  function incremented(uint256 a) internal pure returns(uint256) {
    return a + 1;
  }

  function incremented(uint256[2] a) internal pure returns(uint256[2] c) {
    c = concat(a.hi(), incremented(a.lo()));
    if (isZero(c.lo())) {
      c = concat(incremented(c.hi()), c.lo());
    }
  }

  function incremented(uint256[4] a) internal pure returns(uint256[4] c) {
    c = concat(a.hi(), incremented(a.lo()));
    if (isZero(c.lo())) {
      c = concat(incremented(c.hi()), c.lo());
    }
  }

  function incremented(uint256[8] a) internal pure returns(uint256[8] c) {
    c = concat(a.hi(), incremented(a.lo()));
    if (isZero(c.lo())) {
      c = concat(incremented(c.hi()), c.lo());
    }
  }

  //

  function add(uint256 a, uint256 b) internal pure returns(uint256) {
    return a + b;
  }

  function add(uint256[2] a, uint256[2] b) internal pure returns(uint256[2] c) {
    c = concat(
      add(a.hi(), b.hi()),
      add(a.lo(), b.lo())
    );
    if (lt(c.lo(), a.lo())) {
      c = concat(incremented(c.hi()), c.lo());
    }
  }

  function add(uint256[4] a, uint256[4] b) internal pure returns(uint256[4] c) {
    c = concat(
      add(a.hi(), b.hi()),
      add(a.lo(), b.lo())
    );
    if (lt(c.lo(), a.lo())) {
      c = concat(incremented(c.hi()), c.lo());
    }
  }

  function add(uint256[8] a, uint256[8] b) internal pure returns(uint256[8] c) {
    c = concat(
      add(a.hi(), b.hi()),
      add(a.lo(), b.lo())
    );
    if (lt(c.lo(), a.lo())) {
      c = concat(incremented(c.hi()), c.lo());
    }
  }

  //

  function sub(uint256 a, uint256 b) internal pure returns(uint256) {
    return a - b;
  }

  function sub(uint256[2] a, uint256[2] b) internal pure returns(uint256[2]) {
    return add(a, incremented(inverted(b)));
  }

  function sub(uint256[4] a, uint256[4] b) internal pure returns(uint256[4]) {
    return add(a, incremented(inverted(b)));
  }

  function sub(uint256[8] a, uint256[8] b) internal pure returns(uint256[8]) {
    return add(a, incremented(inverted(b)));
  }

  //

  // http://codepad.org/nB9HqWt1
  //
  // Karatsuba multiplication algorithm
  //
  //            +------+------+
  //            |  x1  |  x0  |
  //            +------+------+
  //                   *
  //            +------+------+
  //            |  y1  |  y0  |
  //            +------+------+
  //                   =
  //     +-------------+-------------+
  //  +  |    x1*y1    |    x0*y0    |
  //     +----+-+------+------+------+
  //          . .             .
  //          . .             .
  //          +-+------+------+
  //       +  | x0*y1 + x1*y0 |
  //          +-+------+------+
  //
  // x0*y1 + x1*y0 = (x0+x1)*(y0+y1) - (lo + hi);
  //

  function mul128(uint128 a, uint128 b) internal pure returns(uint256) {
    return uint256(a) * uint256(b);
  }

  function mul2(uint256 a, uint256 b) internal pure returns(uint256[2] c) {
    var hi = mul128(a.hi(), b.hi());
    var lo = mul128(a.lo(), b.lo());
    var mi = concat128(hi.lo(), lo.hi());

    var m = mi.add(
      add(a.lo(), a.hi())
        .mul2(add(b.lo(), b.hi())).lo()
      .sub(add(lo, hi))
    );

    var hihi = hi.hi();
    if (lt(m, mi)) {
      hihi = uint128(incremented(hihi));
    }
    
    return concat(
      concat128(hihi, m.hi()),
      concat128(m.lo(), lo.lo())
    );
  }

  function mul2(uint256[2] a, uint256[2] b) internal pure returns(uint256[4] c) {
    var hi = mul2(a.hi(), b.hi());
    var lo = mul2(a.lo(), b.lo());
    var mi = concat(hi.lo(), lo.hi());

    var m = mi.add(
      add(a.lo(), a.hi())
        .mul2(add(b.lo(), b.hi()))
      .sub(add(lo, hi))
    );

    var hihi = hi.hi();
    if (lt(m, mi)) {
      hihi = uint(incremented(hihi));
    }
    
    return concat(
      concat(hihi, m.hi()),
      concat(m.lo(), lo.lo())
    );
  }

  function mul2(uint256[4] a, uint256[4] b) internal pure returns(uint256[8] c) {
    var hi = mul2(a.hi(), b.hi());
    var lo = mul2(a.lo(), b.lo());
    var mi = concat(hi.lo(), lo.hi());

    var m = mi.add(
      add(a.lo(), a.hi())
        .mul2(add(b.lo(), b.hi()))
      .sub(add(lo, hi))
    );

    var hihi = hi.hi();
    if (lt(m, mi)) {
      hihi = incremented(hihi);
    }
    
    return concat(
      concat(hihi, m.hi()),
      concat(m.lo(), lo.lo())
    );
  }

  function mul2(uint256[8] a, uint256[8] b) internal pure returns(uint256[16] c) {
    var hi = mul2(a.hi(), b.hi());
    var lo = mul2(a.lo(), b.lo());
    var mi = concat(hi.lo(), lo.hi());

    var m = mi.add(
      add(a.lo(), a.hi())
        .mul2(add(b.lo(), b.hi()))
      .sub(add(lo, hi))
    );

    var hihi = hi.hi();
    if (lt(m, mi)) {
      hihi = incremented(hihi);
    }
    
    return concat(
      concat(hihi, m.hi()),
      concat(m.lo(), lo.lo())
    );
  }

  function mul(uint256 a, uint256 b) internal pure returns(uint256) {
    return a.mul2(b).lo();
  }

  function mul(uint256[2] a, uint256[2] b) internal pure returns(uint256[2]) {
    return a.mul2(b).lo();
  }

  function mul(uint256[4] a, uint256[4] b) internal pure returns(uint256[4]) {
    return a.mul2(b).lo();
  }

  function mul(uint256[8] a, uint256[8] b) internal pure returns(uint256[8]) {
    return a.mul2(b).lo();
  }

  //////

  // function div(uint256[LEN] a, uint256[LEN] b) internal pure returns(uint256[LEN] c) {
  //   // TODO: implement
  // }

  // // Modular

  // function mod(uint256[LEN] a, uint256[LEN] b) internal pure returns(uint256[LEN] c) {
  //   // TODO: implement
  // }

  // function addmod(uint256[LEN] a, uint256[LEN] b, uint256[LEN] m) internal pure returns(uint256[LEN] c) {
  //   // TODO: implement
  // }

  // function mulmod(uint256[LEN] a, uint256[LEN] b, uint256[LEN] m) internal pure returns(uint256[LEN] c) {
  //   // TODO: implement
  // }

  // function powmod(uint256[LEN] a, uint256 p, uint256[LEN] m) internal pure returns(uint256[LEN] c) {
  //   // TODO: implement
  // }
}