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

    it('should submit blocks sequentially', async function() {
        //
        // [1] [2] [3]
        //             [4]
        //                 [5] [6]
        //                         [ ]
        //                         [7] [8] [9]
        //
        await plasma.submitBlocks(0, [1,2,3]);
        await plasma.submitBlocks(3, [4]);
        await plasma.submitBlocks(4, [5,6]);
        await plasma.submitBlocks(6, []);
        await plasma.submitBlocks(6, [7,8,9]);
    });

    it('should submit blocks with intersections', async function() {
        //
        // [1] [2] [3]
        //         [3] [4]
        //             [4] [5] [6]
        //                     [ ]
        //             [4] [5] [6]
        //                 [5] [6] [7] [8] [9]
        //
        await plasma.submitBlocks(0, [1,2,3]);
        await plasma.submitBlocks(2, [3,4]);
        await plasma.submitBlocks(3, [4,5,6]);
        await plasma.submitBlocks(6, []);
        await plasma.submitBlocks(3, [4,5,6]);
        await plasma.submitBlocks(5, [5,6,7,8,9]);
    });

    it('should fail to submit old blocks', async function() {
        //
        // [1] [2] [3] [4]
        //     [2] [3]
        //
        await plasma.submitBlocks(0, [1,2,3,4]);
        await plasma.submitBlocks(1, [2,3]).should.be.rejectedWith(EVMThrow);
    });

    it('should ignore blocks after gap', async function() {
        //
        // [1] [2] [3]
        //             [ ]
        //                 [5] [6]
        //
        await plasma.submitBlocks(0, [1,2,3]);
        await plasma.submitBlocks(4, [5,6]).should.be.rejectedWith(EVMThrow);
    });

    it('should be able to submit blocks signed by operator', async function() {
        //
        // [1] [2] [3]
        //         [3] [4] [5]
        //

        {
            const messageHash = web3.sha3('0x' + [0,1,2,3].map(a => new BigNumber(a).toString(16,64)).reduce((a,b) => a + b), { encoding: 'hex' });
            const rsv = await web3.eth.sign(_, messageHash).substr(2);
            const r = '0x' + rsv.substr(0,64);
            const s = '0x' + rsv.substr(64,64);
            const v = '0x' + rsv.substr(128,2);
            await plasma.submitBlocksSigned(0, [1,2,3], r, s, v, { from: wallet1 });
        }

        {
            const messageHash = web3.sha3('0x' + [2,3,4,5].map(a => new BigNumber(a).toString(16,64)).reduce((a,b) => a + b), { encoding: 'hex' });
            const rsv = await web3.eth.sign(_, messageHash).substr(2);
            const r = '0x' + rsv.substr(0,64);
            const s = '0x' + rsv.substr(64,64);
            const v = '0x' + rsv.substr(128,2);
            await plasma.submitBlocksSigned(2, [3,4,5], r, s, v, { from: wallet2 });
        }
    });
});
