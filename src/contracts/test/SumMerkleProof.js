let assert = require('assert');
const SumMerkleProofWrapper = artifacts.require('SumMerkleProofWrapper');

require('chai')
  .use(require('chai-as-promised'))
  .use(require('chai-bignumber')(web3.BigNumber))
  .should();

const testResult = {
  'rootHash': '37c7f5efafd7761d94ec936360e27fbeae4dd877',
  'rootLength': 'ffffff',
  'begin': '1',
  'end': '2',
  'item': 'fa61c529e022344b84ca026c1fb1214e8bac9afa',
  'itemsLenAndHash': ['00000001dcc703c0e500b653ca82273b7bfad8045d85a470', '00000003d146cb615b8dac6a78641afd24d8c3296cf43a07', '00fffffadcc703c0e500b653ca82273b7bfad8045d85a470'],
  // eslint-disable-next-line max-len
  'rlpEncoded': 'f86301c2010294fa61c529e022344b84ca026c1fb1214e8bac9afab84800000001dcc703c0e500b653ca82273b7bfad8045d85a47000000003d146cb615b8dac6a78641afd24d8c3296cf43a0700fffffadcc703c0e500b653ca82273b7bfad8045d85a470',
};

contract('SumMerkleProofWrapper', function () {
  let wrapper;

  beforeEach(async function () {
    wrapper = await SumMerkleProofWrapper.new();
  });

  it('should verify valid proof', async function () {
    // That proof was generated in Go code
    const root = '0x' + testResult.rootHash;
    const rootLength = '0x' + testResult.rootLength;
    const index = 1;
    const begin = testResult.begin;
    const end = testResult.end;
    const item = '0x' + testResult.item;
    const proofSteps = '0x' +
      testResult.itemsLenAndHash[0] +
      testResult.itemsLenAndHash[1] +
      testResult.itemsLenAndHash[2];

    const result = await wrapper.sumMerkleProofTest(index, begin, end, item, proofSteps, root, rootLength);
    assert.strictEqual(result, true);
  });

  it('should not verify invalid proof', async function () {
    // That proof was generated in Go code
    const root = '0x' + testResult.rootHash;
    const rootLength = '0x' + testResult.rootLength;
    const index = 1;
    const begin = testResult.begin;
    const end = testResult.end;
    const item = '0x' + testResult.item;

    // Make proof wrong by replacing 4 bytes with 0xDEADBEEF
    let stepOne = testResult.itemsLenAndHash[0].slice(0, 40) + 'DEADBEEF';
    let stepTwo = testResult.itemsLenAndHash[1].slice(0, 40) + 'DEADBEEF';
    let stepThree = testResult.itemsLenAndHash[2].slice(0, 40) + 'DEADBEEF';
    const proofSteps = '0x' + stepOne + stepTwo + stepThree;

    const result = await wrapper.sumMerkleProofTest(index, begin, end, item, proofSteps, root, rootLength);
    assert.strictEqual(result, false);
  });

  it('should verify valid proof represented as rlp bytes', async function () {
    const sumMerkleRoot = '0x' + testResult.rootHash;
    const rlpEncodedProof = '0x' + testResult.rlpEncoded;

    const result = await wrapper.sumMerkleProofFromBytesTest(sumMerkleRoot, rlpEncodedProof);
    assert.strictEqual(result, true);
  });
});
