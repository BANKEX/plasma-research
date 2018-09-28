// const EVMRevert = require('./helpers/EVMRevert');
// const EVMThrow = require('./helpers/EVMThrow');

require('chai')
    .use(require('chai-as-promised'))
    .use(require('chai-bignumber')(web3.BigNumber))
    .should();

// const PlasmaMain = artifacts.require('PlasmaMain');

contract('PlasmaMain', function ([_, wallet1, wallet2, wallet3, wallet4, wallet5]) {
    // let plasma;

    beforeEach(async function () {
        // plasma = await PlasmaMain.new();
    });
});
