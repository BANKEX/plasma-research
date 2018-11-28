const BankexPlasma = artifacts.require('BankexPlasma');
const fs = require('fs');
const path = require('path');
const filePath = path.join(__dirname, '../../commons/config/config.verifier.json');
console.log(BankexPlasma)
module.exports = function (deployer) {
    deployer.then(async function () {
        const contract = await BankexPlasma.new();
        const data = JSON.parse(fs.readFileSync(filePath, 'utf8'));
        data.smart = contract.address;
        console.log("BankexPlasma address: " + contract.address)
        fs.writeFileSync(filePath, JSON.stringify(data));
    });
};