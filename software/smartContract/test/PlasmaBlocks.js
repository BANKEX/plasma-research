const BigNumber = require('bn.js');

const EVMRevert = require('./helpers/EVMRevert');
const EVMThrow = require('./helpers/EVMThrow');

require('chai')
    .use(require('chai-as-promised'))
    .use(require('chai-bignumber')(web3.BigNumber))
    .should();

const PlasmaBlocks = artifacts.require('PlasmaBlocks');

const to256bits = function (blocks) {
    return blocks.map(b => (new BigNumber(b)).toString(16, 64)).concat(['']).reduce((a, b) => a + b);
};

const to160bits = function (blocks) {
    return blocks.map(b => (new BigNumber(b)).toString(16, 40)).concat(['']).reduce((a, b) => a + b);
};

contract('PlasmaBlocks', function ([_, wallet1, wallet2, wallet3, wallet4, wallet5]) {
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
        await plasma.submitBlocks(0, '0x' + to160bits([1, 2, 3]), 0, 0);
        await plasma.submitBlocks(3, '0x' + to160bits([4]), 2, 3);
        await plasma.submitBlocks(4, '0x' + to160bits([5, 6]), 3, 4);
        await plasma.submitBlocks(6, '0x' + to160bits([]), 5, 6);
        await plasma.submitBlocks(6, '0x' + to160bits([7, 8, 9]), 5, 6);
    });

    it('should revert on wrong protected block index', async function () {
        //
        // [1] [2] [3] [ ] [ ] [ ]
        //              ^
        //             [4]
        //             [4]--^
        //             [4]------^
        //
        await plasma.submitBlocks(0, '0x' + to160bits([1, 2, 3]), 0, 0);
        await plasma.submitBlocks(3, '0x' + to160bits([4]), 3, 3).should.be.rejectedWith(EVMRevert);
        await plasma.submitBlocks(3, '0x' + to160bits([4]), 4, 3).should.be.rejectedWith(EVMRevert);
        await plasma.submitBlocks(3, '0x' + to160bits([4]), 5, 3).should.be.rejectedWith(EVMRevert);
    });

    it('should handle not latest protected block', async function () {
        //
        // [1] [2] [3]
        //      ^------[4]
        //
        await plasma.submitBlocks(0, '0x' + to160bits([1, 2, 3]), 0, 0);
        await plasma.submitBlocks(3, '0x' + to160bits([4]), 1, 2);
    });

    it('should revert on wrong protected block hash', async function () {
        //
        // [1] [2] [3]
        //          ^--[4]
        //          ^--[4]
        //
        await plasma.submitBlocks(0, '0x' + to160bits([1, 2, 3]), 0, 0);
        await plasma.submitBlocks(3, '0x' + to160bits([4]), 2, 2).should.be.rejectedWith(EVMRevert);
        await plasma.submitBlocks(3, '0x' + to160bits([4]), 2, 4).should.be.rejectedWith(EVMRevert);
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
        await plasma.submitBlocks(0, '0x' + to160bits([1, 2, 3]), 0, 0);
        await plasma.submitBlocks(2, '0x' + to160bits([3, 4]), 2, 3).should.be.rejectedWith(EVMRevert);
        await plasma.submitBlocks(3, '0x' + to160bits([4, 5, 6]), 2, 3);
        await plasma.submitBlocks(6, '0x' + to160bits([]), 5, 6);
        await plasma.submitBlocks(3, '0x' + to160bits([4, 5, 6]), 2, 3).should.be.rejectedWith(EVMRevert);
        await plasma.submitBlocks(3, '0x' + to160bits([4, 5, 6]), 1, 2).should.be.rejectedWith(EVMRevert);
        await plasma.submitBlocks(6, '0x' + to160bits([7, 8]), 3, 4);
        await plasma.submitBlocks(8, '0x' + to160bits([9]), 0, 1);
    });

    it('should revert old blocks', async function () {
        //
        // [1] [2] [3] [4]
        //  ^--[2] [3]
        //  ^--[8] [9]
        //
        await plasma.submitBlocks(0, '0x' + to160bits([1, 2, 3, 4]), 0, 0);
        await plasma.submitBlocks(1, '0x' + to160bits([2, 3]), 0, 1).should.be.rejectedWith(EVMRevert);
        await plasma.submitBlocks(1, '0x' + to160bits([8, 9]), 0, 1).should.be.rejectedWith(EVMRevert);
    });

    it('should deny blocks after gap', async function () {
        //
        // [1] [2] [3]
        //             [ ]
        //          ^------[5] [6]
        //
        await plasma.submitBlocks(0, '0x' + to160bits([1, 2, 3]), 0, 0);
        await plasma.submitBlocks(4, '0x' + to160bits([5, 6]), 2, 3).should.be.rejectedWith(EVMRevert);
    });

    describe('getters', async function () {
        it('should work fine for blocks(i)', async function () {
            //
            // [1] [2] [3]
            //          ^--[4] [5]
            //

            await plasma.blocks.call(0).should.be.rejectedWith(EVMThrow);

            await plasma.submitBlocks(0, '0x' + to160bits([1, 2, 3]), 0, 0);
            (await plasma.blocks.call(0)).should.be.bignumber.equal(1);
            (await plasma.blocks.call(1)).should.be.bignumber.equal(2);
            (await plasma.blocks.call(2)).should.be.bignumber.equal(3);
            await plasma.blocks.call(3).should.be.rejectedWith(EVMThrow);

            await plasma.submitBlocks(3, '0x' + to160bits([4, 5]), 2, 3);
            (await plasma.blocks.call(0)).should.be.bignumber.equal(1);
            (await plasma.blocks.call(1)).should.be.bignumber.equal(2);
            (await plasma.blocks.call(2)).should.be.bignumber.equal(3);
            (await plasma.blocks.call(3)).should.be.bignumber.equal(4);
            (await plasma.blocks.call(4)).should.be.bignumber.equal(5);
            await plasma.blocks.call(5).should.be.rejectedWith(EVMThrow);
        });

        it('should work fine for blocksLength()', async function () {
            //
            // [1] [2] [3]
            //          ^--[4] [5]
            //

            (await plasma.blocksLength.call()).should.be.bignumber.equal(0);

            await plasma.submitBlocks(0, '0x' + to160bits([1, 2, 3]), 0, 0);
            (await plasma.blocksLength.call()).should.be.bignumber.equal(3);

            await plasma.submitBlocks(3, '0x' + to160bits([4, 5]), 2, 3);
            (await plasma.blocksLength.call()).should.be.bignumber.equal(5);
        });
    });

    describe('malformed operator signature', async function () {
        it('should fail to submit blocks with malformed R', async function () {
            const messageHash = web3.sha3('0x' + to256bits([0]) + to160bits([1, 2, 3]), { encoding: 'hex' });
            const rsv = await web3.eth.sign(_, messageHash).substr(2);
            const r = new BigNumber(rsv.substr(0, 64), 16).addn(1).toString(16, 64);
            const s = rsv.substr(64, 64);
            const v = rsv.substr(128, 2);
            await plasma.submitBlocksSigned(0, '0x' + to160bits([1, 2, 3]), '0x' + r + s + v, { from: wallet1 }).should.be.rejectedWith(EVMThrow);
        });

        it('should fail to submit blocks with malformed S', async function () {
            const messageHash = web3.sha3('0x' + to256bits([0]) + to160bits([1, 2, 3]), { encoding: 'hex' });
            const rsv = await web3.eth.sign(_, messageHash).substr(2);
            const r = rsv.substr(0, 64);
            const s = new BigNumber(rsv.substr(64, 128), 16).addn(1).toString(16, 64);
            const v = rsv.substr(128, 2);
            await plasma.submitBlocksSigned(0, '0x' + to160bits([1, 2, 3]), '0x' + r + s + v, { from: wallet1 }).should.be.rejectedWith(EVMThrow);
        });

        it('should fail to submit blocks with malformed V', async function () {
            const messageHash = web3.sha3('0x' + to256bits([0]) + to160bits([1, 2, 3]), { encoding: 'hex' });
            const rsv = await web3.eth.sign(_, messageHash).substr(2);
            const r = rsv.substr(0, 64);
            const s = rsv.substr(64, 64);
            const v = new BigNumber(rsv.substr(128, 2), 16).addn(1).toString(16, 64);
            await plasma.submitBlocksSigned(0, '0x' + to160bits([1, 2, 3]), '0x' + r + s + v, { from: wallet1 }).should.be.rejectedWith(EVMThrow);
        });

        it('should fail to submit blocks with malformed R and S', async function () {
            const messageHash = web3.sha3('0x' + to256bits([0]) + to160bits([1, 2, 3]), { encoding: 'hex' });
            const rsv = await web3.eth.sign(_, messageHash).substr(2);
            const r = rsv.substr(0, 64);
            const s = rsv.substr(64, 64);
            const v = rsv.substr(128, 2);
            await plasma.submitBlocksSigned(0, '0x' + to160bits([1, 2, 3]), '0x' + s + r + v, { from: wallet1 }).should.be.rejectedWith(EVMThrow);
        });

        it('should fail to submit blocks with malformed short DATA', async function () {
            const messageHash = web3.sha3('0x' + to256bits([0]) + to160bits([1, 2, 3]), { encoding: 'hex' });
            const rsv = await web3.eth.sign(_, messageHash).substr(2);
            const r = rsv.substr(0, 64);
            const s = rsv.substr(64, 64);
            const v = rsv.substr(128, 2);
            await plasma.submitBlocksSigned(0, '0x' + to160bits([1, 2]), '0x' + r + s + v, { from: wallet1 }).should.be.rejectedWith(EVMThrow);
        });

        it('should fail to submit blocks with malformed long DATA', async function () {
            const messageHash = web3.sha3('0x' + to256bits([0]) + to160bits([1, 2, 3]), { encoding: 'hex' });
            const rsv = await web3.eth.sign(_, messageHash).substr(2);
            const r = rsv.substr(0, 64);
            const s = rsv.substr(64, 64);
            const v = rsv.substr(128, 2);
            await plasma.submitBlocksSigned(0, '0x' + to160bits([1, 2, 3, 4]), '0x' + r + s + v, { from: wallet1 }).should.be.rejectedWith(EVMThrow);
        });

        it('should fail to submit blocks with malformed reordered DATA', async function () {
            const messageHash = web3.sha3('0x' + to256bits([0]) + to160bits([1, 2, 3]), { encoding: 'hex' });
            const rsv = await web3.eth.sign(_, messageHash).substr(2);
            const r = rsv.substr(0, 64);
            const s = rsv.substr(64, 64);
            const v = rsv.substr(128, 2);
            await plasma.submitBlocksSigned(0, '0x' + to160bits([1, 3, 2]), '0x' + r + s + v, { from: wallet1 }).should.be.rejectedWith(EVMThrow);
        });

        it('should fail to submit blocks with malformed DATA with different offset', async function () {
            const messageHash = web3.sha3('0x' + to256bits([0]) + to160bits([1, 2, 3]), { encoding: 'hex' });
            const rsv = await web3.eth.sign(_, messageHash).substr(2);
            const r = rsv.substr(0, 64);
            const s = rsv.substr(64, 64);
            const v = rsv.substr(128, 2);
            await plasma.submitBlocksSigned(1, '0x' + to160bits([1, 2, 3]), '0x' + r + s + v, { from: wallet1 }).should.be.rejectedWith(EVMThrow);
        });
    });

    it('should be able to submit blocks signed by operator', async function () {
        //
        // [1] [2] [3]
        //          ^--[4]
        //              ^--[5] [6] [7]
        //

        {
            const messageHash = web3.sha3(
                '0x' + to256bits([0]) + to160bits([1, 2, 3]) + to256bits([0]) + to160bits([0]),
                { encoding: 'hex' }
            );
            const rsv = await web3.eth.sign(_, messageHash);
            await plasma.submitBlocksSigned(0, '0x' + to160bits([1, 2, 3]), 0, 0, rsv, { from: wallet1 });
        }

        {
            const messageHash = web3.sha3(
                '0x' + to256bits([3]) + to160bits([4]) + to256bits([2]) + to160bits([3]),
                { encoding: 'hex' }
            );
            const rsv = await web3.eth.sign(_, messageHash);
            await plasma.submitBlocksSigned(3, '0x' + to160bits([4]), 2, 3, rsv, { from: wallet2 });
        }

        {
            const messageHash = web3.sha3(
                '0x' + to256bits([4]) + to160bits([5, 6, 7]) + to256bits([3]) + to160bits([4]),
                { encoding: 'hex' }
            );
            const rsv = await web3.eth.sign(_, messageHash);
            await plasma.submitBlocksSigned(4, '0x' + to160bits([5, 6, 7]), 3, 4, rsv, { from: wallet3 });
        }
    });
});
