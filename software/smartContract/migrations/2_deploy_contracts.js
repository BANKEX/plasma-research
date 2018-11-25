const BankexPlasma = artifacts.require('BankexPlasma');
const fs = require('fs');
const path = require('path');
const filePath = path.join(__dirname, '../../client-operator/verifier/config.json');


module.exports = function (deployer) {
    deployer.then(async function () {
        const contractAddress = await BankexPlasma.new();
        const data = JSON.parse(fs.readFileSync(filePath, 'utf8'));
        data.smart = contractAddress.address;
        console.log("BankexPlasma address: " + contractAddress.address)
        fs.writeFileSync(filePath, JSON.stringify(data));
    });
};
