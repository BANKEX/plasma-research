const BigNumber = require('bn.js');
const { keccak256 } = require('ethereumjs-util');

const be32 = function (x) {
  return x.toBuffer('be', 32).slice(-4);
};

const be160 = function (x) {
  return x.toBuffer('be', 32).slice(-20);
};

const hash4 = function (l1, l2, a1, a2) {
  return new BigNumber(keccak256(Buffer.concat([be32(l1), be32(l2), be160(a1), be160(a2)])).slice(12));
};

const randBits = function (bits) {
  let value = new BigNumber();
  for (let i = 0; i < bits; i++) {
    if (Math.random() > 0.5) {
      value = value.or((new BigNumber(2)).pow(new BigNumber(i)));
    }
  }
  return value;
};

const genRandProof = function (depth) {
  const item = randBits(160);
  const begin = randBits(30);
  const end = begin.add(randBits(24));

  let index = 0;
  let proof = new Buffer('');
  let curItem = item;
  let curLength = end.sub(begin);
  let curLeft = begin;

  for (let i = 0; i < depth; i++) {
    let b = randBits(1);
    let ci = randBits(160);
    let cl = randBits(24);
    if (i === (depth - 2)) {
      b = new BigNumber(1);
      cl = curLeft;
    }
    if (i === (depth - 1)) {
      b = new BigNumber(0);
      cl = new BigNumber(2 ** 32 - 1).sub(curLength);
    }

    proof = Buffer.concat([proof, be32(cl), be160(ci)]);
    index = index.add(b.mul((new BigNumber(2)).pow(i)));
    if (b === 1) {
      curItem = hash4(cl, curLength, ci, curItem);
      curLeft = curLeft.sub(cl);
      curLength = curLength.add(cl);
    } else {
      curItem = hash4(curLength, cl, curItem, ci);
      curLength = curLength.add(cl);
    }
  }

  return {
    index,
    begin,
    end,
    item,
    proof,
    curItem,
    curLength,
  };
};
