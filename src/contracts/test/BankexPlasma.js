const { BN, ether, balance } = require('openzeppelin-test-helpers');
const BankexPlasma = artifacts.require('BankexPlasma');
const Token = artifacts.require('Erc20Mock');

contract('BankexPlasma', function (accounts) {
    let owner = accounts[0];
    let user = accounts[2];

    describe('deposit', function () {
        let plasma;

        before(async function () {
            plasma = await BankexPlasma.new({ from: owner });
        });

        describe('ether', function () {
            let tx;
            let testEtherAmount = ether('42');
            let userBalance;
            let contractBalance;

            before(async function () {
                userBalance = await balance.current(user);
                contractBalance = await balance.current(plasma.address);
            });

            it('should deposit', async function () {
                tx = await plasma.deposit({ from: user, value: testEtherAmount, gasPrice: 0 });
            });

            it('should transfer asset from user', async function () {
                let currentBalance = await balance.current(user);
                assert.equal(userBalance.sub(testEtherAmount).toString(), currentBalance.toString());
            });

            it('should transfer asset to plasma contract', async function () {
                let currentBalance = await balance.current(plasma.address);
                assert.equal(contractBalance.add(testEtherAmount).toString(), currentBalance.toString());
            });

            it('should emit `CoinDeposited` event', function () {
                assert.equal(tx.logs.length, 2);
                assert.equal(tx.logs[0].event, 'CoinDeposited');
                assert.equal(tx.logs[0].args.who, user);
                assert.equal(tx.logs[0].args.amount.toString(), testEtherAmount.div(new BN(1e13)).toString());
            });

            it('should emit `AssetDeposited` event', function () {
                assert.equal(tx.logs[1].event, 'AssetDeposited');
                assert.equal(tx.logs[1].args.token, '0x0000000000000000000000000000000000000000');
                assert.equal(tx.logs[1].args.who, user);
                assert.equal(tx.logs[1].args.intervalId.toString(), '1');
                assert.equal(tx.logs[1].args.begin.toString(), '0');
                assert.equal(tx.logs[1].args.end.toString(), '4199999');
            });
        });

        describe('erc20', function () {
            let tx;
            let token;
            let testTokenAmount = ether('42');
            let userBalance;
            let contractBalance;

            before(async function () {
                token = await Token.new({ from: owner });
                await plasma.setAssetOffset(token.address, 24);

                await token.mint(user, testTokenAmount);
                await token.approve(plasma.address, testTokenAmount, { from: user });

                userBalance = await token.balanceOf(user);
                contractBalance = await token.balanceOf(plasma.address);
            });

            it('should deposit', async function () {
                tx = await plasma.depositERC20(token.address, testTokenAmount, { from: user });
            });

            it('should transfer tokens from user', async function () {
                let currentBalance = await token.balanceOf(user);
                assert.equal(userBalance.sub(testTokenAmount).toString(), currentBalance.toString());
            });

            it('should transfer tokens to plasma contract', async function () {
                let currentBalance = await token.balanceOf(plasma.address);
                assert.equal(contractBalance.add(testTokenAmount).toString(), currentBalance.toString());
            });

            it('should emit `CoinDeposited` event', function () {
                assert.equal(tx.logs.length, 2);
                assert.equal(tx.logs[0].event, 'ERC20Deposited');
                assert.equal(tx.logs[0].args.who, user);
                assert.equal(tx.logs[0].args.amount.toString(), testTokenAmount.toString());
            });

            it('should emit `AssetDeposited` event', function () {
                assert.equal(tx.logs[1].event, 'AssetDeposited');
                assert.equal(tx.logs[1].args.token, token.address);
                assert.equal(tx.logs[1].args.who, user);
                assert.equal(tx.logs[1].args.intervalId.toString(), '1');
                assert.equal(tx.logs[1].args.begin.toString(), '0');
                assert.equal(tx.logs[1].args.end.toString(), '4199999');
            });
        });
    });
});
