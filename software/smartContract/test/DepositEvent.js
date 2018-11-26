const BankexPlasma = artifacts.require("./BankexPlasma.sol");

const web3 = global.web3;

contract('BankexPlasma', (accounts) => {

    //initial params for testing

    it("deploy", async () => {
        bp = await BankexPlasma.new({from: accounts[0]});

        console.log(bp);
        await bp.deposit()
    })

});