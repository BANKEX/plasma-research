const BigNumber = require('bn.js');

// const EVMRevert = require('./helpers/EVMRevert');
const EVMThrow = require('./helpers/EVMThrow');

require('chai')
    .use(require('chai-as-promised'))
    .use(require('chai-bignumber')(web3.BigNumber))
    .should();

const PlasmaBlocks = artifacts.require('PlasmaBlocks');

contract('PlasmaBlocks', function ([_, wallet1, wallet2, wallet3, wallet4, wallet5]) {
    let plasma;

    beforeEach(async function () {
        plasma = await PlasmaBlocks.new();
    });

    it('should submit blocks sequentially', async function () {
        //
        // [1] [2] [3]
        //             [4]
        //                 [5] [6]
        //                         [ ]
        //                         [7] [8] [9]
        //
        await plasma.submitBlocks(0, [1, 2, 3]);
        await plasma.submitBlocks(3, [4]);
        await plasma.submitBlocks(4, [5, 6]);
        await plasma.submitBlocks(6, []);
        await plasma.submitBlocks(6, [7, 8, 9]);
    });

    it('should submit blocks with intersections', async function () {
        //
        // [1] [2] [3]
        //         [3] [4]
        //             [4] [5] [6]
        //                     [ ]
        //             [4] [5] [6]
        //                 [5] [6] [7] [8] [9]
        //
        await plasma.submitBlocks(0, [1, 2, 3]);
        await plasma.submitBlocks(2, [3, 4]);
        await plasma.submitBlocks(3, [4, 5, 6]);
        await plasma.submitBlocks(6, []);
        await plasma.submitBlocks(3, [4, 5, 6]);
        await plasma.submitBlocks(5, [5, 6, 7, 8, 9]);
    });

    it('should ignore old blocks', async function () {
        //
        // [1] [2] [3] [4]
        //     [2] [3]
        //     [8] [9]
        //
        await plasma.submitBlocks(0, [1, 2, 3, 4]);
        await plasma.submitBlocks(1, [2, 3]);
        await plasma.submitBlocks(1, [8, 9]);
    });

    it('should deny blocks after gap', async function () {
        //
        // [1] [2] [3]
        //             [ ]
        //                 [5] [6]
        //
        await plasma.submitBlocks(0, [1, 2, 3]);
        await plasma.submitBlocks(4, [5, 6]).should.be.rejectedWith(EVMThrow);
    });

    describe('getters', async function () {
        it('should work fine for blocks(i)', async function () {
            //
            // [1] [2] [3]
            //         [3] [4] [5]
            //

            await plasma.blocks.call(0).should.be.rejectedWith(EVMThrow);

            await plasma.submitBlocks(0, [1, 2, 3]);
            (await plasma.blocks.call(0)).should.be.bignumber.equal(1);
            (await plasma.blocks.call(1)).should.be.bignumber.equal(2);
            (await plasma.blocks.call(2)).should.be.bignumber.equal(3);
            await plasma.blocks.call(3).should.be.rejectedWith(EVMThrow);

            await plasma.submitBlocks(2, [3, 4, 5]);
            (await plasma.blocks.call(0)).should.be.bignumber.equal(1);
            (await plasma.blocks.call(1)).should.be.bignumber.equal(2);
            (await plasma.blocks.call(2)).should.be.bignumber.equal(3);
            (await plasma.blocks.call(3)).should.be.bignumber.equal(4);
            (await plasma.blocks.call(4)).should.be.bignumber.equal(5);
            await plasma.blocks.call(5).should.be.rejectedWith(EVMThrow);
        });

        it('should work fine for allBlocks()', async function () {
            //
            // [1] [2] [3]
            //         [3] [4] [5]
            //

            (await plasma.allBlocks.call()).should.be.deep.equal([]);

            await plasma.submitBlocks(0, [1, 2, 3]);
            (await plasma.allBlocks.call()).map(a => a.toNumber()).should.be.deep.equal([1, 2, 3]);

            await plasma.submitBlocks(2, [3, 4, 5]);
            (await plasma.allBlocks.call()).map(a => a.toNumber()).should.be.deep.equal([1, 2, 3, 4, 5]);
        });

        it('should work fine for blocksLength()', async function () {
            //
            // [1] [2] [3]
            //         [3] [4] [5]
            //

            (await plasma.blocksLength.call()).should.be.bignumber.equal(0);

            await plasma.submitBlocks(0, [1, 2, 3]);
            (await plasma.blocksLength.call()).should.be.bignumber.equal(3);

            await plasma.submitBlocks(2, [3, 4, 5]);
            (await plasma.blocksLength.call()).should.be.bignumber.equal(5);
        });
    });

    describe('malformed operator signature', async function () {
        it('should fail to submit blocks with malformed R', async function () {
            const messageHash = web3.sha3('0x' + [0, 1, 2, 3].map(a => new BigNumber(a).toString(16, 64)).reduce((a, b) => a + b), { encoding: 'hex' });
            const rsv = await web3.eth.sign(_, messageHash).substr(2);
            const r = '0x' + new BigNumber(rsv.substr(0, 64), 16).addn(1).toString(16, 64);
            const s = '0x' + rsv.substr(64, 64);
            const v = '0x' + rsv.substr(128, 2);
            await plasma.submitBlocksSigned(0, [1, 2, 3], r, s, v, { from: wallet1 }).should.be.rejectedWith(EVMThrow);
        });

        it('should fail to submit blocks with malformed S', async function () {
            const messageHash = web3.sha3('0x' + [0, 1, 2, 3].map(a => new BigNumber(a).toString(16, 64)).reduce((a, b) => a + b), { encoding: 'hex' });
            const rsv = await web3.eth.sign(_, messageHash).substr(2);
            const r = '0x' + new BigNumber(rsv.substr(0, 64), 16).addn(1).toString(16, 64);
            const s = '0x' + rsv.substr(64, 64);
            const v = '0x' + rsv.substr(128, 2);
            await plasma.submitBlocksSigned(0, [1, 2, 3], r, s, v, { from: wallet1 }).should.be.rejectedWith(EVMThrow);
        });

        it('should fail to submit blocks with malformed V', async function () {
            const messageHash = web3.sha3('0x' + [0, 1, 2, 3].map(a => new BigNumber(a).toString(16, 64)).reduce((a, b) => a + b), { encoding: 'hex' });
            const rsv = await web3.eth.sign(_, messageHash).substr(2);
            const r = '0x' + rsv.substr(0, 64);
            const s = '0x' + new BigNumber(rsv.substr(64, 128), 16).addn(1).toString(16, 64);
            const v = '0x' + rsv.substr(128, 2);
            await plasma.submitBlocksSigned(0, [1, 2, 3], r, s, v, { from: wallet1 }).should.be.rejectedWith(EVMThrow);
        });

        it('should fail to submit blocks with malformed R and S', async function () {
            const messageHash = web3.sha3('0x' + [0, 1, 2, 3].map(a => new BigNumber(a).toString(16, 64)).reduce((a, b) => a + b), { encoding: 'hex' });
            const rsv = await web3.eth.sign(_, messageHash).substr(2);
            const r = '0x' + rsv.substr(0, 64);
            const s = '0x' + rsv.substr(64, 64);
            const v = '0x' + rsv.substr(128, 2);
            await plasma.submitBlocksSigned(0, [1, 2, 3], s, r, v, { from: wallet1 }).should.be.rejectedWith(EVMThrow);
        });

        it('should fail to submit blocks with malformed short DATA', async function () {
            const messageHash = web3.sha3('0x' + [0, 1, 2, 3].map(a => new BigNumber(a).toString(16, 64)).reduce((a, b) => a + b), { encoding: 'hex' });
            const rsv = await web3.eth.sign(_, messageHash).substr(2);
            const r = '0x' + rsv.substr(0, 64);
            const s = '0x' + rsv.substr(64, 64);
            const v = '0x' + rsv.substr(128, 2);
            await plasma.submitBlocksSigned(0, [1, 2], r, s, v, { from: wallet1 }).should.be.rejectedWith(EVMThrow);
        });

        it('should fail to submit blocks with malformed long DATA', async function () {
            const messageHash = web3.sha3('0x' + [0, 1, 2, 3].map(a => new BigNumber(a).toString(16, 64)).reduce((a, b) => a + b), { encoding: 'hex' });
            const rsv = await web3.eth.sign(_, messageHash).substr(2);
            const r = '0x' + rsv.substr(0, 64);
            const s = '0x' + rsv.substr(64, 64);
            const v = '0x' + rsv.substr(128, 2);
            await plasma.submitBlocksSigned(0, [1, 2, 3, 4], r, s, v, { from: wallet1 }).should.be.rejectedWith(EVMThrow);
        });

        it('should fail to submit blocks with malformed reordered DATA', async function () {
            const messageHash = web3.sha3('0x' + [0, 1, 2, 3].map(a => new BigNumber(a).toString(16, 64)).reduce((a, b) => a + b), { encoding: 'hex' });
            const rsv = await web3.eth.sign(_, messageHash).substr(2);
            const r = '0x' + rsv.substr(0, 64);
            const s = '0x' + rsv.substr(64, 64);
            const v = '0x' + rsv.substr(128, 2);
            await plasma.submitBlocksSigned(0, [1, 3, 2], r, s, v, { from: wallet1 }).should.be.rejectedWith(EVMThrow);
        });

        it('should fail to submit blocks with malformed DATA with different offset', async function () {
            const messageHash = web3.sha3('0x' + [0, 1, 2, 3].map(a => new BigNumber(a).toString(16, 64)).reduce((a, b) => a + b), { encoding: 'hex' });
            const rsv = await web3.eth.sign(_, messageHash).substr(2);
            const r = '0x' + rsv.substr(0, 64);
            const s = '0x' + rsv.substr(64, 64);
            const v = '0x' + rsv.substr(128, 2);
            await plasma.submitBlocksSigned(1, [1, 2, 3], r, s, v, { from: wallet1 }).should.be.rejectedWith(EVMThrow);
        });
    });

    it('should be able to submit blocks signed by operator', async function () {
        //
        // [1] [2] [3]
        //         [3] [4] [5]
        //             [4] [5] [6] [7]
        //

        {
            const messageHash = web3.sha3('0x' + [0, 1, 2, 3].map(a => new BigNumber(a).toString(16, 64)).reduce((a, b) => a + b), { encoding: 'hex' });
            const rsv = await web3.eth.sign(_, messageHash).substr(2);
            const r = '0x' + rsv.substr(0, 64);
            const s = '0x' + rsv.substr(64, 64);
            const v = '0x' + rsv.substr(128, 2);
            await plasma.submitBlocksSigned(0, [1, 2, 3], r, s, v, { from: wallet1 });
        }

        {
            const messageHash = web3.sha3('0x' + [2, 3, 4, 5].map(a => new BigNumber(a).toString(16, 64)).reduce((a, b) => a + b), { encoding: 'hex' });
            const rsv = await web3.eth.sign(_, messageHash).substr(2);
            const r = '0x' + rsv.substr(0, 64);
            const s = '0x' + rsv.substr(64, 64);
            const v = '0x' + rsv.substr(128, 2);
            await plasma.submitBlocksSigned(2, [3, 4, 5], r, s, v, { from: wallet2 });
        }

        {
            const messageHash = web3.sha3('0x' + [3, 4, 5, 6, 7].map(a => new BigNumber(a).toString(16, 64)).reduce((a, b) => a + b), { encoding: 'hex' });
            const rsv = await web3.eth.sign(_, messageHash).substr(2);
            const r = '0x' + rsv.substr(0, 64);
            const s = '0x' + rsv.substr(64, 64);
            const v = '0x' + rsv.substr(128, 2);
            await plasma.submitBlocksSigned(3, [4, 5, 6, 7], r, s, v, { from: wallet3 });
        }
    });
});
