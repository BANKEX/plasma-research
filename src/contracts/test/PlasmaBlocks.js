const { BN, shouldFail } = require('openzeppelin-test-helpers');

const PlasmaBlocks = artifacts.require('PlasmaBlocks');

const to256bits = function (blocks) {
  return blocks.map(b => (new BN(b)).toString(16, 64)).concat(['']).reduce((a, b) => a + b);
};

const to160bits = function (blocks) {
  return blocks.map(b => (new BN(b)).toString(16, 40)).concat(['']).reduce((a, b) => a + b);
};

contract('PlasmaBlocks', function ([_, wallet1, wallet2, wallet3]) {
  let plasma;

  beforeEach(async function () {
    plasma = await PlasmaBlocks.new();
  });

  it('should submit blocks sequentially', async function () {
    //
    // [1] [2] [3]
    //          ^--[4]
    //              ^--[5] [6]
    //                      ^--[ ]
    //                      ^--[7] [8] [9]
    //
    await plasma.submitBlocks(0, '0x' + to160bits([1, 2, 3]));
    await plasma.submitBlocks(3, '0x' + to160bits([4]));
    await plasma.submitBlocks(4, '0x' + to160bits([5, 6]));
    await plasma.submitBlocks(6, '0x' + to160bits([]));
    await plasma.submitBlocks(6, '0x' + to160bits([7, 8, 9]));
  });

  it('should handle not latest protected block', async function () {
    //
    // [1] [2] [3]
    //      ^------[4]
    //
    await plasma.submitBlocks(0, '0x' + to160bits([1, 2, 3]));
    await plasma.submitBlocks(3, '0x' + to160bits([4]));
  });

  it('should revert on submit blocks with intersections', async function () {
    //
    // [1] [2] [3]
    //      ^--[3] [4]
    //          ^--[4] [5] [6]
    //                  ^--[ ]
    //          ^--[4] [5] [6]
    //      ^------[4] [5] [6]
    //              ^----------[7] [8]
    //  ^------------------------------[9]
    //
    await plasma.submitBlocks(0, '0x' + to160bits([1, 2, 3]));
    await shouldFail.reverting(
      plasma.submitBlocks(2, '0x' + to160bits([3, 4]))
    );
    await plasma.submitBlocks(3, '0x' + to160bits([4, 5, 6]));
    await plasma.submitBlocks(6, '0x' + to160bits([]));
    await shouldFail.reverting(
      plasma.submitBlocks(3, '0x' + to160bits([4, 5, 6]))
    );
    await shouldFail.reverting(
      plasma.submitBlocks(3, '0x' + to160bits([4, 5, 6]))
    );
    await plasma.submitBlocks(6, '0x' + to160bits([7, 8]));
    await plasma.submitBlocks(8, '0x' + to160bits([9]));
  });

  it('should revert old blocks', async function () {
    //
    // [1] [2] [3] [4]
    //  ^--[2] [3]
    //  ^--[8] [9]
    //
    await plasma.submitBlocks(0, '0x' + to160bits([1, 2, 3, 4]));
    await shouldFail.reverting(
      plasma.submitBlocks(1, '0x' + to160bits([2, 3]))
    );
    await shouldFail.reverting(
      plasma.submitBlocks(1, '0x' + to160bits([8, 9]))
    );
  });

  it('should deny blocks after gap', async function () {
    //
    // [1] [2] [3]
    //             [ ]
    //          ^------[5] [6]
    //
    await plasma.submitBlocks(0, '0x' + to160bits([1, 2, 3]));
    await shouldFail.reverting(
      plasma.submitBlocks(4, '0x' + to160bits([5, 6]))
    );
  });

  describe('getters', async function () {
    it('should work fine for blocks(i)', async function () {
      //
      // [1] [2] [3]
      //          ^--[4] [5]
      //

      await shouldFail.throwing(
        plasma.blocks.call(0)
      );

      await plasma.submitBlocks(0, '0x' + to160bits([1, 2, 3]));

      web3.utils.toBN(await plasma.blocks.call(0)).should.be.bignumber.equal('1');
      web3.utils.toBN(await plasma.blocks.call(1)).should.be.bignumber.equal('2');
      web3.utils.toBN(await plasma.blocks.call(2)).should.be.bignumber.equal('3');

      await shouldFail.throwing(
        plasma.blocks.call(3)
      );

      await plasma.submitBlocks(3, '0x' + to160bits([4, 5]));
      web3.utils.toBN(await plasma.blocks.call(0)).should.be.bignumber.equal('1');
      web3.utils.toBN(await plasma.blocks.call(1)).should.be.bignumber.equal('2');
      web3.utils.toBN(await plasma.blocks.call(2)).should.be.bignumber.equal('3');
      web3.utils.toBN(await plasma.blocks.call(3)).should.be.bignumber.equal('4');
      web3.utils.toBN(await plasma.blocks.call(4)).should.be.bignumber.equal('5');
      await shouldFail.throwing(
        plasma.blocks.call(5)
      );
    });

    it('should work fine for blocksLength()', async function () {
      //
      // [1] [2] [3]
      //          ^--[4] [5]
      //

      (await plasma.blocksLength.call()).should.be.bignumber.equal('0');

      await plasma.submitBlocks(0, '0x' + to160bits([1, 2, 3]));
      (await plasma.blocksLength.call()).should.be.bignumber.equal('3');

      await plasma.submitBlocks(3, '0x' + to160bits([4, 5]));
      (await plasma.blocksLength.call()).should.be.bignumber.equal('5');
    });
  });

  describe('malformed operator signature', async function () {
    it('should fail to submit blocks with malformed R', async function () {
      const messageHash = web3.utils.keccak256('0x' + to256bits([0]) + to160bits([1, 2, 3]), { encoding: 'hex' });
      const rsv = (await web3.eth.sign(messageHash, _)).substr(2);
      const r = new BN(rsv.substr(0, 64), 16).addn(1).toString(16, 64);
      const s = rsv.substr(64, 64);
      const v = rsv.substr(128, 2);
      await shouldFail.reverting(
        plasma.submitBlocksSigned(0, '0x' + to160bits([1, 2, 3]), '0x' + r + s + v, { from: wallet1 })
      );
    });

    it('should fail to submit blocks with malformed S', async function () {
      const messageHash = web3.utils.keccak256('0x' + to256bits([0]) + to160bits([1, 2, 3]), { encoding: 'hex' });
      const rsv = (await web3.eth.sign(messageHash, _)).substr(2);
      const r = rsv.substr(0, 64);
      const s = new BN(rsv.substr(64, 128), 16).addn(1).toString(16, 64);
      const v = rsv.substr(128, 2);
      await shouldFail.reverting(
        plasma.submitBlocksSigned(0, '0x' + to160bits([1, 2, 3]), '0x' + r + s + v, { from: wallet1 })
      );
    });

    it('should fail to submit blocks with malformed V', async function () {
      const messageHash = web3.utils.keccak256('0x' + to256bits([0]) + to160bits([1, 2, 3]), { encoding: 'hex' });
      const rsv = (await web3.eth.sign(messageHash, _)).substr(2);
      const r = rsv.substr(0, 64);
      const s = rsv.substr(64, 64);
      const v = new BN(rsv.substr(128, 2), 16).addn(1).toString(16, 64);
      await shouldFail.reverting(
        plasma.submitBlocksSigned(0, '0x' + to160bits([1, 2, 3]), '0x' + r + s + v, { from: wallet1 })
      );
    });

    it('should fail to submit blocks with malformed R and S', async function () {
      const messageHash = web3.utils.keccak256('0x' + to256bits([0]) + to160bits([1, 2, 3]), { encoding: 'hex' });
      const rsv = (await web3.eth.sign(messageHash, _)).substr(2);
      const r = rsv.substr(0, 64);
      const s = rsv.substr(64, 64);
      const v = rsv.substr(128, 2);
      await shouldFail.reverting(
        plasma.submitBlocksSigned(0, '0x' + to160bits([1, 2, 3]), '0x' + s + r + v, { from: wallet1 })
      );
    });

    it('should fail to submit blocks with malformed short DATA', async function () {
      const messageHash = web3.utils.keccak256('0x' + to256bits([0]) + to160bits([1, 2, 3]), { encoding: 'hex' });
      const rsv = (await web3.eth.sign(messageHash, _)).substr(2);
      const r = rsv.substr(0, 64);
      const s = rsv.substr(64, 64);
      const v = rsv.substr(128, 2);
      await shouldFail.reverting(
        plasma.submitBlocksSigned(0, '0x' + to160bits([1, 2]), '0x' + r + s + v, { from: wallet1 })
      );
    });

    it('should fail to submit blocks with malformed long DATA', async function () {
      const messageHash = web3.utils.keccak256('0x' + to256bits([0]) + to160bits([1, 2, 3]), { encoding: 'hex' });
      const rsv = (await web3.eth.sign(messageHash, _)).substr(2);
      const r = rsv.substr(0, 64);
      const s = rsv.substr(64, 64);
      const v = rsv.substr(128, 2);
      await shouldFail.reverting(
        plasma.submitBlocksSigned(0, '0x' + to160bits([1, 2, 3, 4]), '0x' + r + s + v, { from: wallet1 })
      );
    });

    it('should fail to submit blocks with malformed reordered DATA', async function () {
      const messageHash = web3.utils.keccak256('0x' + to256bits([0]) + to160bits([1, 2, 3]), { encoding: 'hex' });
      const rsv = (await web3.eth.sign(messageHash, _)).substr(2);
      const r = rsv.substr(0, 64);
      const s = rsv.substr(64, 64);
      const v = rsv.substr(128, 2);
      await shouldFail.reverting(
        plasma.submitBlocksSigned(0, '0x' + to160bits([1, 3, 2]), '0x' + r + s + v, { from: wallet1 })
      );
    });

    it('should fail to submit blocks with malformed DATA with different offset', async function () {
      const messageHash = web3.utils.keccak256('0x' + to256bits([0]) + to160bits([1, 2, 3]), { encoding: 'hex' });
      const rsv = (await web3.eth.sign(messageHash, _)).substr(2);
      const r = rsv.substr(0, 64);
      const s = rsv.substr(64, 64);
      const v = rsv.substr(128, 2);
      await shouldFail.reverting(
        plasma.submitBlocksSigned(1, '0x' + to160bits([1, 2, 3]), '0x' + r + s + v, { from: wallet1 })
      );
    });
  });

  it('should be able to submit blocks signed by operator', async function () {
    //
    // [1] [2] [3]
    //          ^--[4]
    //              ^--[5] [6] [7]
    //

    {
      const messageHash = web3.utils.keccak256(
        '0x' + to256bits([0]) + to160bits([1, 2, 3]),
        { encoding: 'hex' }
      );
      const rsv = await web3.eth.sign(messageHash, _);
      await plasma.submitBlocksSigned(0, '0x' + to160bits([1, 2, 3]), rsv, { from: wallet1 });
    }

    {
      const messageHash = web3.utils.keccak256(
        '0x' + to256bits([3]) + to160bits([4]),
        { encoding: 'hex' }
      );
      const rsv = await web3.eth.sign(messageHash, _);
      await plasma.submitBlocksSigned(3, '0x' + to160bits([4]), rsv, { from: wallet2 });
    }

    {
      const messageHash = web3.utils.keccak256(
        '0x' + to256bits([4]) + to160bits([5, 6, 7]),
        { encoding: 'hex' }
      );
      const rsv = await web3.eth.sign(messageHash, _);
      await plasma.submitBlocksSigned(4, '0x' + to160bits([5, 6, 7]), rsv, { from: wallet3 });
    }
  });
});
