pragma solidity ^0.4.24;


library Uint2048 {
  using Uint2048 for uint256[1];
  using Uint2048 for uint256[2];
  using Uint2048 for uint256[4];
  using Uint2048 for uint256[8];
  using Uint2048 for uint256[16];

  function concat(uint256[1] memory a, uint256[1] memory /*b*/) public pure returns(uint256[2] memory c) {
    assembly {
      c := a
    }
  }

  function concat(uint256[2] memory a, uint256[2] memory /*b*/) public pure returns(uint256[4] memory c) {
    assembly {
      c := a
    }
  }

  function concat(uint256[4] memory a, uint256[4] memory /*b*/) public pure returns(uint256[8] memory c) {
    assembly {
      c := a
    }
  }

  function concat(uint256[8] memory a, uint256[8] memory /*b*/) public pure returns(uint256[16] memory c) {
    assembly {
      c := a
    }
  }

  //

  // position: 0 - high, 1 - mid, 2 - low
  function half(uint256[1] memory a, uint256 position) public pure returns(uint256 c) {
    return uint128(a[0] >> (128 - position * 64));
  }

  // position: 0 - high, 1 - mid, 2 - low
  function half(uint256[2] memory a, uint256 position) public pure returns(uint256[1] memory c) {
    assembly {
      c := add(a, mul(position, 0x10))
    }
  }
  
  // position: 0 - high, 1 - mid, 2 - low
  function half(uint256[4] memory a, uint256 position) public pure returns(uint256[2] memory c) {
    assembly {
      c := add(a, mul(position, 0x20))
    }
  }
  
  // position: 0 - high, 1 - mid, 2 - low
  function half(uint256[8] memory a, uint256 position) public pure returns(uint256[4] memory c) {
    assembly {
      c := add(a, mul(position, 0x40))
    }
  }

  // position: 0 - high, 1 - mid, 2 - low
  function half(uint256[16] memory a, uint256 position) public pure returns(uint256[8] memory c) {
    assembly {
      c := add(a, mul(position, 0x40))
    }
  }

  //

  function lo(uint256[1] memory a) public pure returns(uint256) {
    return half(a, 2);
  }

  function lo(uint256[2] memory a) public pure returns(uint256[1] memory) {
    return half(a, 2);
  }

  function lo(uint256[4] memory a) public pure returns(uint256[2] memory) {
    return half(a, 2);
  }

  function lo(uint256[8] memory a) public pure returns(uint256[4] memory) {
    return half(a, 2);
  }

  function lo(uint256[16] memory a) public pure returns(uint256[8] memory) {
    return half(a, 2);
  }

  //

  function hi(uint256[1] memory a) public pure returns(uint256) {
    return half(a, 0);
  }

  function hi(uint256[2] memory a) public pure returns(uint256[1] memory) {
    return half(a, 0);
  }

  function hi(uint256[4] memory a) public pure returns(uint256[2] memory) {
    return half(a, 0);
  }

  function hi(uint256[8] memory a) public pure returns(uint256[4] memory) {
    return half(a, 0);
  }

  function hi(uint256[16] memory a) public pure returns(uint256[8] memory) {
    return half(a, 0);
  }

  //

  function mid(uint256[1] memory a) public pure returns(uint256 c) {
    return half(a, 1);
  }

  function mid(uint256[2] memory a) public pure returns(uint256[1] memory c) {
    return half(a, 1);
  }

  function mid(uint256[4] memory a) public pure returns(uint256[2] memory) {
    return half(a, 1);
  }

  function mid(uint256[8] memory a) public pure returns(uint256[4] memory) {
    return half(a, 1);
  }

  function mid(uint256[16] memory a) public pure returns(uint256[8] memory) {
    return half(a, 1);
  }

  ////////////////////////////////////////////////////////////////
  // Compare
  ////////////////////////////////////////////////////////////////

  function isZero(uint256[1] memory a) internal pure returns(bool) {
    return a[0] == 0;
  }

  function isZero(uint256[2] memory a) internal pure returns(bool) {
    return isZero(a.lo()) && isZero(a.hi());
  }

  function isZero(uint256[4] memory a) internal pure returns(bool) {
    return isZero(a.lo()) && isZero(a.hi());
  }

  function isZero(uint256[8] memory a) internal pure returns(bool) {
    return isZero(a.lo()) && isZero(a.hi());
  }

  function isZero(uint256[16] memory a) internal pure returns(bool) {
    return isZero(a.lo()) && isZero(a.hi());
  }
  
  //

  function eq(uint256[1] memory a, uint256[1] memory b) internal pure returns(bool) {
    return a[0] == b[0];
  }

  function eq(uint256[2] memory a, uint256[2] memory b) internal pure returns(bool) {
    return eq(a.hi(), b.hi()) && eq(a.lo(), b.lo());
  }

  function eq(uint256[4] memory a, uint256[4] memory b) internal pure returns(bool) {
    return eq(a.hi(), b.hi()) && eq(a.lo(), b.lo());
  }

  function eq(uint256[8] memory a, uint256[8] memory b) internal pure returns(bool) {
    return eq(a.hi(), b.hi()) && eq(a.lo(), b.lo());
  }

  function eq(uint256[16] memory a, uint256[16] memory b) internal pure returns(bool) {
    return eq(a.hi(), b.hi()) && eq(a.lo(), b.lo());
  }

  //

  function lt(uint256[1] memory a, uint256[1] memory b) internal pure returns(bool) {
    return a.hi() < b.hi() || (a.hi() == b.hi() && a.lo() < b.lo());
  }

  function lt(uint256[2] memory a, uint256[2] memory b) internal pure returns(bool) {
    return lt(a.hi(), b.hi()) || (eq(a.hi(), b.hi()) && lt(a.lo(), b.lo()));
  }

  function lt(uint256[4] memory a, uint256[4] memory b) internal pure returns(bool) {
    return lt(a.hi(), b.hi()) || (eq(a.hi(), b.hi()) && lt(a.lo(), b.lo()));
  }

  function lt(uint256[8] memory a, uint256[8] memory b) internal pure returns(bool) {
    return lt(a.hi(), b.hi()) || (eq(a.hi(), b.hi()) && lt(a.lo(), b.lo()));
  }

  function lt(uint256[16] memory a, uint256[16] memory b) internal pure returns(bool) {
    return lt(a.hi(), b.hi()) || (eq(a.hi(), b.hi()) && lt(a.lo(), b.lo()));
  }

  //

  function gt(uint256[1] memory a, uint256[1] memory b) internal pure returns(bool) {
    return !eq(a, b) && !lt(a, b);
  }

  function gt(uint256[2] memory a, uint256[2] memory b) internal pure returns(bool) {
    return !eq(a, b) && !lt(a, b);
  }

  function gt(uint256[4] memory a, uint256[4] memory b) internal pure returns(bool) {
    return !eq(a, b) && !lt(a, b);
  }

  function gt(uint256[8] memory a, uint256[8] memory b) internal pure returns(bool) {
    return !eq(a, b) && !lt(a, b);
  }

  function gt(uint256[16] memory a, uint256[16] memory b) internal pure returns(bool) {
    return !eq(a, b) && !lt(a, b);
  }

  //

  function le(uint256[1] memory a, uint256[1] memory b) internal pure returns(bool) {
    return lt(a, b) || eq(a, b);
  }

  function le(uint256[2] memory a, uint256[2] memory b) internal pure returns(bool) {
    return lt(a, b) || eq(a, b);
  }

  function le(uint256[4] memory a, uint256[4] memory b) internal pure returns(bool) {
    return lt(a, b) || eq(a, b);
  }

  function le(uint256[8] memory a, uint256[8] memory b) internal pure returns(bool) {
    return lt(a, b) || eq(a, b);
  }

  function le(uint256[16] memory a, uint256[16] memory b) internal pure returns(bool) {
    return lt(a, b) || eq(a, b);
  }

  //

  function ge(uint256[1] memory a, uint256[1] memory b) internal pure returns(bool) {
    return !lt(a, b);
  }

  function ge(uint256[2] memory a, uint256[2] memory b) internal pure returns(bool) {
    return !lt(a, b);
  }

  function ge(uint256[4] memory a, uint256[4] memory b) internal pure returns(bool) {
    return !lt(a, b);
  }

  function ge(uint256[8] memory a, uint256[8] memory b) internal pure returns(bool) {
    return !lt(a, b);
  }

  function ge(uint256[16] memory a, uint256[16] memory b) internal pure returns(bool) {
    return !lt(a, b);
  }

  ////////////////////////////////////////////////////////////////
  // Math
  ////////////////////////////////////////////////////////////////

  function inverted(uint256[1] memory a) internal pure returns(uint256[1] memory) {
    return [uint256(~a[0])];
  }

  function inverted(uint256[2] memory a) internal pure returns(uint256[2] memory) {
    return concat(inverted(a.hi()), inverted(a.lo()));
  }

  function inverted(uint256[4] memory a) internal pure returns(uint256[4] memory) {
    return concat(inverted(a.hi()), inverted(a.lo()));
  }

  function inverted(uint256[8] memory a) internal pure returns(uint256[8] memory) {
    return concat(inverted(a.hi()), inverted(a.lo()));
  }

  function inverted(uint256[16] memory a) internal pure returns(uint256[16] memory) {
    return concat(inverted(a.hi()), inverted(a.lo()));
  }

  //

  function invert(uint256[1] memory a) internal pure returns(uint256[1] memory) {
    a[0] = ~a[0];
    return a;
  }

  function invert(uint256[2] memory a) internal pure returns(uint256[2] memory) {
    invert(a.lo());
    invert(a.hi());
    return a;
  }

  function invert(uint256[4] memory a) internal pure returns(uint256[4] memory) {
    invert(a.lo());
    invert(a.hi());
    return a;
  }

  function invert(uint256[8] memory a) internal pure returns(uint256[8] memory) {
    invert(a.lo());
    invert(a.hi());
    return a;
  }

  function invert(uint256[16] memory a) internal pure returns(uint256[16] memory) {
    invert(a.lo());
    invert(a.hi());
    return a;
  }

  //

  function incremented(uint256[1] memory a) internal pure returns(uint256[1] memory) {
    return [a[0] + 1];
  }

  function incremented(uint256[2] memory a) internal pure returns(uint256[2] memory c) {
    c = concat(a.hi(), incremented(a.lo()));
    if (isZero(c.lo())) {
      c = concat(incremented(c.hi()), c.lo());
    }
  }

  function incremented(uint256[4] memory a) internal pure returns(uint256[4] memory c) {
    c = concat(a.hi(), incremented(a.lo()));
    if (isZero(c.lo())) {
      c = concat(incremented(c.hi()), c.lo());
    }
  }

  function incremented(uint256[8] memory a) internal pure returns(uint256[8] memory c) {
    c = concat(a.hi(), incremented(a.lo()));
    if (isZero(c.lo())) {
      c = concat(incremented(c.hi()), c.lo());
    }
  }

  //

  function increment(uint256[1] memory a) internal pure returns(uint256[1] memory) {
    a[0] += 1;
    return a;
  }

  function increment(uint256[2] memory a) internal pure returns(uint256[2] memory) {
    increment(a.lo());
    if (a.lo().isZero()) {
      increment(a.hi());
    }
    return a;
  }

  function increment(uint256[4] memory a) internal pure returns(uint256[4] memory) {
    increment(a.lo());
    if (a.lo().isZero()) {
      increment(a.hi());
    }
    return a;
  }

  function increment(uint256[8] memory a) internal pure returns(uint256[8] memory) {
    increment(a.lo());
    if (a.lo().isZero()) {
      increment(a.hi());
    }
    return a;
  }

  function increment(uint256[16] memory a) internal pure returns(uint256[16] memory) {
    increment(a.lo());
    if (a.lo().isZero()) {
      increment(a.hi());
    }
    return a;
  }

  //

  function add(uint256[1] memory a, uint256[1] memory b) internal pure returns(uint256[1] memory) {
    return [a[0] + b[0]];
  }

  function add(uint256[2] memory a, uint256[2] memory b) internal pure returns(uint256[2] memory c) {
    c = concat(
      add(a.hi(), b.hi()),
      add(a.lo(), b.lo())
    );
    if (lt(c.lo(), a.lo())) {
      c = concat(incremented(c.hi()), c.lo());
    }
  }

  function add(uint256[4] memory a, uint256[4] memory b) internal pure returns(uint256[4] memory c) {
    c = concat(
      add(a.hi(), b.hi()),
      add(a.lo(), b.lo())
    );
    if (lt(c.lo(), a.lo())) {
      c = concat(incremented(c.hi()), c.lo());
    }
  }

  function add(uint256[8] memory a, uint256[8] memory b) internal pure returns(uint256[8] memory c) {
    c = concat(
      add(a.hi(), b.hi()),
      add(a.lo(), b.lo())
    );
    if (lt(c.lo(), a.lo())) {
      c = concat(incremented(c.hi()), c.lo());
    }
  }

  //

  function addto(uint256[1] memory a, uint256[1] memory b) internal pure returns(uint256[1] memory) {
    a[0] += b[0];
    return a;
  }

  function addto(uint256[2] memory a, uint256[2] memory b) internal pure returns(uint256[2] memory) {
    addto(a.lo(). b.lo());
    addto(a.hi(). b.hi());
    if (lt(a.lo(), b.lo())) {
      increment(a.hi());
    }
    return a;
  }

  function addto(uint256[4] memory a, uint256[4] memory b) internal pure returns(uint256[4] memory) {
    addto(a.lo(). b.lo());
    addto(a.hi(). b.hi());
    if (lt(a.lo(), b.lo())) {
      increment(a.hi());
    }
    return a;
  }

  function addto(uint256[8] memory a, uint256[8] memory b) internal pure returns(uint256[8] memory) {
    addto(a.lo(). b.lo());
    addto(a.hi(). b.hi());
    if (lt(a.lo(), b.lo())) {
      increment(a.hi());
    }
    return a;
  }

  function addto(uint256[16] memory a, uint256[16] memory b) internal pure returns(uint256[16] memory) {
    addto(a.lo(). b.lo());
    addto(a.hi(). b.hi());
    if (lt(a.lo(), b.lo())) {
      increment(a.hi());
    }
    return a;
  }

  //

  function sub(uint256[1] memory a, uint256[1] memory b) internal pure returns(uint256[1] memory) {
    return [a[0] - b[0]];
  }

  function sub(uint256[2] memory a, uint256[2] memory b) internal pure returns(uint256[2] memory) {
    return add(a, incremented(inverted(b)));
  }

  function sub(uint256[4] memory a, uint256[4] memory b) internal pure returns(uint256[4] memory) {
    return add(a, incremented(inverted(b)));
  }

  function sub(uint256[8] memory a, uint256[8] memory b) internal pure returns(uint256[8] memory) {
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

  function mul2(uint256[1] memory a, uint256[1] memory b) internal pure returns(uint256[2] memory c) {
    c = [
      a.hi() * b.hi(),
      a.lo() * b.lo()
    ];
    uint256[1] memory mi = mid(c);

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

  function mul2(uint256[2] memory a, uint256[2] memory b) internal pure returns(uint256[4] memory c) {
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

  function mul2(uint256[4] memory a, uint256[4] memory b) internal pure returns(uint256[8] memory c) {
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

  function mul2(uint256[8] memory a, uint256[8] memory b) internal pure returns(uint256[16] memory c) {
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

  function mul(uint256[1] memory a, uint256[1] memory b) internal pure returns(uint256[1] memory) {
    return a.mul2(b).lo();
  }

  function mul(uint256[2] memory a, uint256[2] memory b) internal pure returns(uint256[2] memory) {
    return a.mul2(b).lo();
  }

  function mul(uint256[4] memory a, uint256[4] memory b) internal pure returns(uint256[4] memory) {
    return a.mul2(b).lo();
  }

  function mul(uint256[8] memory a, uint256[8] memory b) internal pure returns(uint256[8] memory) {
    return a.mul2(b).lo();
  }

  //////

  // function div(uint256[LEN] memory a, uint256[LEN] memory b) internal pure returns(uint256[LEN] memory c) {
  //   // TODO: implement
  // }

  // // Modular

  // function mod(uint256[LEN] memory a, uint256[LEN] memory b) internal pure returns(uint256[LEN] memory c) {
  //   // TODO: implement
  // }

  // function addmod(uint256[LEN] memory a, uint256[LEN] memory b, uint256[LEN] memory m) internal pure returns(uint256[LEN] memory c) {
  //   // TODO: implement
  // }

  // function mulmod(uint256[LEN] memory a, uint256[LEN] memory b, uint256[LEN] memory m) internal pure returns(uint256[LEN] memory c) {
  //   // TODO: implement
  // }

  // function powmod(uint256[LEN] memory a, uint256 p, uint256[LEN] memory m) internal pure returns(uint256[LEN] memory c) {
  //   // TODO: implement
  // }
}