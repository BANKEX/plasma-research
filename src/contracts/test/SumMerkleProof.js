const BigNumber = require('bn.js');
const { keccak256, bufferToHex } = require('ethereumjs-util');
const EVMRevert = require('./helpers/EVMRevert');
const EVMThrow = require('./helpers/EVMThrow');
var assert = require('assert');

require('chai')
  .use(require('chai-as-promised'))
  .use(require('chai-bignumber')(web3.BigNumber))
  .should();

const SumMerkleProofWrapper = artifacts.require('SumMerkleProofWrapper');
const SumMerkleProof = artifacts.require('SumMerkleProof');

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

contract('SumMerkleProofWrapper', function ([_, wallet1, wallet2, wallet3, wallet4, wallet5]) {
  let wrapper;

  beforeEach(async function () {
    wrapper = await SumMerkleProofWrapper.new();
  });

  it('should verify valid proof', async function () {
    // That proof was generated in Go code
    const root = '0x37c7f5efafd7761d94ec936360e27fbeae4dd877';
    const rootLength = '16777215';
    const index = 1;
    const begin = 1;
    const end = 2;
    const item = '0xfa61c529e022344b84ca026c1fb1214e8bac9afa';
    const proofSteps = '0x' +
      '00000001dcc703c0e500b653ca82273b7bfad8045d85a470' +
      '00000003d146cb615b8dac6a78641afd24d8c3296cf43a07' +
      '00fffffadcc703c0e500b653ca82273b7bfad8045d85a470';

    const result = await wrapper.sumMerkleProofTest(index, begin, end, item, proofSteps, root, rootLength);
    assert.strictEqual(result, true);
  });

  it('shouldn\'t verify invalid proof', async function () {
    // That proof was generated in Go code
    const root = '0x37c7f5efafd7761d94ec936360e27fbeae4dd877';
    const rootLength = '16777215';
    const index = 1;
    const begin = 1;
    const end = 2;
    const item = '0xfa61c529e022344b84ca026c1fb1214e8bac9afa';

    // Make proof wrong by replacing 4 bytes with 0xDEADBEEF
    const proofSteps = '0x' +
      '00000001dcc703c0e500b653ca82273b7bfad804DEADBEEF' +
      '00000003d146cb615b8dac6a78641afd24d8c329DEADBEEF' +
      '00fffffadcc703c0e500b653ca82273b7bfad804DEADBEEF';

    const result = await wrapper.sumMerkleProofTest(index, begin, end, item, proofSteps, root, rootLength);
    assert.strictEqual(result, false);
  });

  it('should verify valid proof represented as rlp bytes', async function () {
    const sumMerkleRoot = '0x37c7f5efafd7761d94ec936360e27fbeae4dd877';
    const rlpEncodedProof = '0xf86301c2010294fa61c529e022344b84ca026c1fb1214e8bac9afab84800000001dcc703c0e500b65' +
      '3ca82273b7bfad8045d85a47000000003d146cb615b8dac6a78641afd24d8c3296cf43a0700fffffadcc703c0e500b653ca82273b' +
      '7bfad8045d85a470';

    const result = await wrapper.sumMerkleProofFromBytesTest(sumMerkleRoot, rlpEncodedProof);
    assert.strictEqual(result, true);
  });
});
