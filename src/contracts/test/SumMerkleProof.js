const BigNumber = require('bn.js');
const { keccak256, bufferToHex } = require('ethereumjs-util');
const EVMRevert = require('./helpers/EVMRevert');
const EVMThrow = require('./helpers/EVMThrow');
var assert = require('assert');
var fs = require('fs');

require('chai')
  .use(require('chai-as-promised'))
  .use(require('chai-bignumber')(web3.BigNumber))
  .should();

var resultJSON = JSON.parse(fs.readFileSync('test/result.json'));

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
    const root = '0x' + resultJSON.rootHash;
    const rootLength = '0x' + resultJSON.rootLength;
    const index = 1;
    const begin = resultJSON.begin;
    const end = resultJSON.end;
    const item = '0x' + resultJSON.item;
    const proofSteps = '0x' +
      resultJSON.itemsLenAndHash[0] +
      resultJSON.itemsLenAndHash[1] +
      resultJSON.itemsLenAndHash[2];

    const result = await wrapper.sumMerkleProofTest(index, begin, end, item, proofSteps, root, rootLength);
    assert.strictEqual(result, true);
  });

  it('shouldn\'t verify invalid proof', async function () {
    // That proof was generated in Go code
    const root = '0x' + resultJSON.rootHash;
    const rootLength = '0x' + resultJSON.rootLength;
    const index = 1;
    const begin = resultJSON.begin;
    const end = resultJSON.end;
    const item = '0x' + resultJSON.item;

    // Make proof wrong by replacing 4 bytes with 0xDEADBEEF
    let stepOne = resultJSON.itemsLenAndHash[0].slice(0, 40) + 'DEADBEEF';
    let stepTwo = resultJSON.itemsLenAndHash[1].slice(0, 40) + 'DEADBEEF';
    let stepThree = resultJSON.itemsLenAndHash[2].slice(0, 40) + 'DEADBEEF';
    const proofSteps = '0x' + stepOne + stepTwo + stepThree;

    const result = await wrapper.sumMerkleProofTest(index, begin, end, item, proofSteps, root, rootLength);
    assert.strictEqual(result, false);
  });

  it('should verify valid proof represented as rlp bytes', async function () {
    const sumMerkleRoot = '0x' + resultJSON.rootHash;
    const rlpEncodedProof = '0x' + resultJSON.rlpEncoded;

    const result = await wrapper.sumMerkleProofFromBytesTest(sumMerkleRoot, rlpEncodedProof);
    assert.strictEqual(result, true);
  });
});
